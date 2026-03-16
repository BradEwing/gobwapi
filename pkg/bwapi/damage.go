package bwapi

// damageMultiplier256 is the DamageType vs UnitSizeType multiplier table.
// Values are scaled by 256 to avoid floating point (256 = 1.0).
// Indexed as [DamageType][UnitSizeType].
var damageMultiplier256 = [7][6]int{
	// Independent:  Ind  Small  Med  Large  None  Unknown
	{256, 256, 256, 256, 256, 256},
	// Explosive:
	{256, 128, 192, 256, 256, 256},
	// Concussive:
	{256, 256, 128, 64, 256, 256},
	// Normal:
	{256, 256, 256, 256, 256, 256},
	// Ignore_Armor:
	{256, 256, 256, 256, 256, 256},
	// None:
	{0, 0, 0, 0, 0, 0},
	// Unknown:
	{256, 256, 256, 256, 256, 256},
}

// GetDamageFrom calculates the damage that fromType deals to toType,
// accounting for upgrades and armor. Returns damage per hit.
func (g *Game) GetDamageFrom(fromType, toType UnitType, fromPlayer, toPlayer *Player) int {
	// Determine weapon based on target type
	var weapon WeaponType
	if toType.IsFlyer() {
		weapon = fromType.AirWeapon()
	} else {
		weapon = fromType.GroundWeapon()
	}
	if weapon == WeaponTypeNone {
		return 0
	}

	// Base damage + upgrade bonus
	baseDamage := weapon.DamageAmount()
	if fromPlayer != nil {
		upgradeLevel := fromPlayer.UpgradeLevel(weapon.GetUpgradeType())
		baseDamage += weapon.DamageBonus() * upgradeLevel
	}

	// Target armor + upgrade bonus
	armor := toType.Armor()
	if toPlayer != nil {
		armor += toPlayer.UpgradeLevel(toType.ArmorUpgrade())
	}

	// Apply armor reduction (ignore armor for IgnoreArmor damage type)
	dmgType := weapon.GetDamageType()
	reduced := baseDamage
	if dmgType != DamageTypeIgnoreArmor {
		reduced = baseDamage - armor
		if reduced < 1 && baseDamage > 0 {
			reduced = 1
		}
	}

	// Apply damage type vs unit size multiplier
	sizeType := toType.Size()
	if int(dmgType) < len(damageMultiplier256) && int(sizeType) < len(damageMultiplier256[0]) {
		reduced = reduced * damageMultiplier256[dmgType][sizeType] / 256
	}

	// Apply damage factor (e.g., Zealot hits twice = factor 2)
	reduced *= weapon.DamageFactor()

	if reduced < 0 {
		return 0
	}
	return reduced
}

// GetDamageTo calculates the damage that toType receives from fromType.
// This is the same as GetDamageFrom with parameters in the more natural order.
func (g *Game) GetDamageTo(toType, fromType UnitType, toPlayer, fromPlayer *Player) int {
	return g.GetDamageFrom(fromType, toType, fromPlayer, toPlayer)
}

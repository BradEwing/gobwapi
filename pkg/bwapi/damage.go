package bwapi

// damageMultiplier256 is the DamageType vs UnitSizeType multiplier table.
// Values are scaled by 256 to avoid floating point (256 = 1.0).
// Indexed as [DamageType][UnitSizeType].
// Source: rsbwapi DAMAGE_RATIO / BWAPI DamageType vs UnitSizeType table.
var damageMultiplier256 = [7][6]int{
	{0, 0, 0, 0, 0, 0},       // Independent
	{0, 128, 192, 256, 0, 0}, // Explosive: 0.5x Small, 0.75x Medium, 1.0x Large
	{0, 256, 128, 64, 0, 0},  // Concussive: 1.0x Small, 0.5x Medium, 0.25x Large
	{0, 256, 256, 256, 0, 0}, // Normal
	{0, 256, 256, 256, 0, 0}, // IgnoreArmor
	{0, 0, 0, 0, 0, 0},       // None
	{0, 0, 0, 0, 0, 0},       // Unknown
}

// GetDamageFrom calculates the damage that fromType deals to toType,
// accounting for upgrades, armor, damage type, and unit size. Returns damage per hit.
func (g *Game) GetDamageFrom(fromType, toType UnitType, fromPlayer, toPlayer *Player) int {
	var weapon WeaponType
	if toType.IsFlyer() {
		weapon = fromType.AirWeapon()
	} else {
		weapon = fromType.GroundWeapon()
	}
	if weapon == WeaponTypeNone {
		return 0
	}

	baseDamage := weapon.DamageAmount() * weapon.DamageFactor()
	if fromPlayer != nil {
		baseDamage += weapon.DamageBonus() * weapon.DamageFactor() * fromPlayer.UpgradeLevel(weapon.GetUpgradeType())
	}

	armor := toType.Armor()
	if toPlayer != nil {
		armor += toPlayer.UpgradeLevel(toType.ArmorUpgrade())
	}

	dmgType := weapon.GetDamageType()
	reduced := baseDamage
	if dmgType != DamageTypeIgnoreArmor {
		reduced = baseDamage - armor
		if reduced < 0 {
			reduced = 0
		}
	}

	sizeType := toType.Size()
	if int(dmgType) < len(damageMultiplier256) && int(sizeType) < len(damageMultiplier256[0]) {
		reduced = reduced * damageMultiplier256[dmgType][sizeType] / 256
	}

	if reduced < 1 && baseDamage > 0 {
		reduced = 1
	}

	return reduced
}

// GetDamageTo calculates the damage that toType receives from fromType.
func (g *Game) GetDamageTo(toType, fromType UnitType, toPlayer, fromPlayer *Player) int {
	return g.GetDamageFrom(fromType, toType, fromPlayer, toPlayer)
}

package bwapi

// UpgradeType represents a BWAPI upgrade type.
type UpgradeType int32

const (
	UpgradeTypeTerranInfantryArmor   UpgradeType = 0
	UpgradeTypeTerranVehiclePlating  UpgradeType = 1
	UpgradeTypeTerranShipPlating     UpgradeType = 2
	UpgradeTypeZergCarapace          UpgradeType = 3
	UpgradeTypeZergFlyerCarapace     UpgradeType = 4
	UpgradeTypeProtossGroundArmor    UpgradeType = 5
	UpgradeTypeProtossAirArmor       UpgradeType = 6
	UpgradeTypeTerranInfantryWeapons UpgradeType = 7
	UpgradeTypeTerranVehicleWeapons  UpgradeType = 8
	UpgradeTypeTerranShipWeapons     UpgradeType = 9
	UpgradeTypeZergMeleeAttacks      UpgradeType = 10
	UpgradeTypeZergMissileAttacks    UpgradeType = 11
	UpgradeTypeZergFlyerAttacks      UpgradeType = 12
	UpgradeTypeProtossGroundWeapons  UpgradeType = 13
	UpgradeTypeProtossAirWeapons     UpgradeType = 14
	UpgradeTypeProtossPlasmaShields  UpgradeType = 15
	UpgradeTypeU238Shells            UpgradeType = 16
	UpgradeTypeIonThrusters          UpgradeType = 17
	// 18 is skipped
	UpgradeTypeTitanReactor          UpgradeType = 19
	UpgradeTypeOcularImplants        UpgradeType = 20
	UpgradeTypeMoebiusReactor        UpgradeType = 21
	UpgradeTypeApolloReactor         UpgradeType = 22
	UpgradeTypeColossusReactor       UpgradeType = 23
	UpgradeTypeVentralSacs           UpgradeType = 24
	UpgradeTypeAntennae              UpgradeType = 25
	UpgradeTypePneumatizedCarapace   UpgradeType = 26
	UpgradeTypeMetabolicBoost        UpgradeType = 27
	UpgradeTypeAdrenalGlands         UpgradeType = 28
	UpgradeTypeMuscularAugments      UpgradeType = 29
	UpgradeTypeGroovedSpines         UpgradeType = 30
	UpgradeTypeGameteMeiosis         UpgradeType = 31
	UpgradeTypeMetasynapticNode      UpgradeType = 32
	UpgradeTypeSingularityCharge     UpgradeType = 33
	UpgradeTypeLegEnhancements       UpgradeType = 34
	UpgradeTypeScarabDamage          UpgradeType = 35
	UpgradeTypeReaverCapacity        UpgradeType = 36
	UpgradeTypeGraviticDrive         UpgradeType = 37
	UpgradeTypeSensorArray           UpgradeType = 38
	UpgradeTypeGraviticBoosters      UpgradeType = 39
	UpgradeTypeKhaydarinAmulet       UpgradeType = 40
	UpgradeTypeApialSensors          UpgradeType = 41
	UpgradeTypeGraviticThrusters     UpgradeType = 42
	UpgradeTypeCarrierCapacity       UpgradeType = 43
	UpgradeTypeKhaydarinCore         UpgradeType = 44
	// 45-46 are skipped
	UpgradeTypeArgusJewel            UpgradeType = 47
	// 48 is skipped
	UpgradeTypeArgusTalisman         UpgradeType = 49
	// 50 is skipped
	UpgradeTypeCaduceusReactor       UpgradeType = 51
	UpgradeTypeChitinousPlating      UpgradeType = 52
	UpgradeTypeAnabolicSynthesis     UpgradeType = 53
	UpgradeTypeCharonBoosters        UpgradeType = 54
	// 55-59 are skipped
	UpgradeTypeUpgrade60             UpgradeType = 60
	UpgradeTypeNone                  UpgradeType = 61
	UpgradeTypeUnknown               UpgradeType = 62
)

func (u UpgradeType) String() string {
	switch u {
	case UpgradeTypeTerranInfantryArmor:
		return "TerranInfantryArmor"
	case UpgradeTypeTerranVehiclePlating:
		return "TerranVehiclePlating"
	case UpgradeTypeTerranShipPlating:
		return "TerranShipPlating"
	case UpgradeTypeZergCarapace:
		return "ZergCarapace"
	case UpgradeTypeZergFlyerCarapace:
		return "ZergFlyerCarapace"
	case UpgradeTypeProtossGroundArmor:
		return "ProtossGroundArmor"
	case UpgradeTypeProtossAirArmor:
		return "ProtossAirArmor"
	case UpgradeTypeTerranInfantryWeapons:
		return "TerranInfantryWeapons"
	case UpgradeTypeTerranVehicleWeapons:
		return "TerranVehicleWeapons"
	case UpgradeTypeTerranShipWeapons:
		return "TerranShipWeapons"
	case UpgradeTypeZergMeleeAttacks:
		return "ZergMeleeAttacks"
	case UpgradeTypeZergMissileAttacks:
		return "ZergMissileAttacks"
	case UpgradeTypeZergFlyerAttacks:
		return "ZergFlyerAttacks"
	case UpgradeTypeProtossGroundWeapons:
		return "ProtossGroundWeapons"
	case UpgradeTypeProtossAirWeapons:
		return "ProtossAirWeapons"
	case UpgradeTypeProtossPlasmaShields:
		return "ProtossPlasmaShields"
	case UpgradeTypeU238Shells:
		return "U238Shells"
	case UpgradeTypeIonThrusters:
		return "IonThrusters"
	case UpgradeTypeTitanReactor:
		return "TitanReactor"
	case UpgradeTypeOcularImplants:
		return "OcularImplants"
	case UpgradeTypeMoebiusReactor:
		return "MoebiusReactor"
	case UpgradeTypeApolloReactor:
		return "ApolloReactor"
	case UpgradeTypeColossusReactor:
		return "ColossusReactor"
	case UpgradeTypeVentralSacs:
		return "VentralSacs"
	case UpgradeTypeAntennae:
		return "Antennae"
	case UpgradeTypePneumatizedCarapace:
		return "PneumatizedCarapace"
	case UpgradeTypeMetabolicBoost:
		return "MetabolicBoost"
	case UpgradeTypeAdrenalGlands:
		return "AdrenalGlands"
	case UpgradeTypeMuscularAugments:
		return "MuscularAugments"
	case UpgradeTypeGroovedSpines:
		return "GroovedSpines"
	case UpgradeTypeGameteMeiosis:
		return "GameteMeiosis"
	case UpgradeTypeMetasynapticNode:
		return "MetasynapticNode"
	case UpgradeTypeSingularityCharge:
		return "SingularityCharge"
	case UpgradeTypeLegEnhancements:
		return "LegEnhancements"
	case UpgradeTypeScarabDamage:
		return "ScarabDamage"
	case UpgradeTypeReaverCapacity:
		return "ReaverCapacity"
	case UpgradeTypeGraviticDrive:
		return "GraviticDrive"
	case UpgradeTypeSensorArray:
		return "SensorArray"
	case UpgradeTypeGraviticBoosters:
		return "GraviticBoosters"
	case UpgradeTypeKhaydarinAmulet:
		return "KhaydarinAmulet"
	case UpgradeTypeApialSensors:
		return "ApialSensors"
	case UpgradeTypeGraviticThrusters:
		return "GraviticThrusters"
	case UpgradeTypeCarrierCapacity:
		return "CarrierCapacity"
	case UpgradeTypeKhaydarinCore:
		return "KhaydarinCore"
	case UpgradeTypeArgusJewel:
		return "ArgusJewel"
	case UpgradeTypeArgusTalisman:
		return "ArgusTalisman"
	case UpgradeTypeCaduceusReactor:
		return "CaduceusReactor"
	case UpgradeTypeChitinousPlating:
		return "ChitinousPlating"
	case UpgradeTypeAnabolicSynthesis:
		return "AnabolicSynthesis"
	case UpgradeTypeCharonBoosters:
		return "CharonBoosters"
	case UpgradeTypeUpgrade60:
		return "Upgrade60"
	case UpgradeTypeNone:
		return "None"
	case UpgradeTypeUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

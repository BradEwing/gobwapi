package bwapi

// DamageType represents a BWAPI damage type.
type DamageType int32

const (
	DamageTypeIndependent DamageType = 0
	DamageTypeExplosive   DamageType = 1
	DamageTypeConcussive  DamageType = 2
	DamageTypeNormal      DamageType = 3
	DamageTypeIgnoreArmor DamageType = 4
	DamageTypeNone        DamageType = 5
	DamageTypeUnknown     DamageType = 6
)

func (d DamageType) String() string {
	switch d {
	case DamageTypeIndependent:
		return "Independent"
	case DamageTypeExplosive:
		return "Explosive"
	case DamageTypeConcussive:
		return "Concussive"
	case DamageTypeNormal:
		return "Normal"
	case DamageTypeIgnoreArmor:
		return "Ignore_Armor"
	case DamageTypeNone:
		return "None"
	case DamageTypeUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

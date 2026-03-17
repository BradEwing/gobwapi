package bwapi

// UnitSizeType represents a BWAPI unit size type (affects damage calculations).
type UnitSizeType int32

const (
	UnitSizeTypeIndependent UnitSizeType = 0
	UnitSizeTypeSmall       UnitSizeType = 1
	UnitSizeTypeMedium      UnitSizeType = 2
	UnitSizeTypeLarge       UnitSizeType = 3
	UnitSizeTypeNone        UnitSizeType = 4
	UnitSizeTypeUnknown     UnitSizeType = 5
)

func (s UnitSizeType) String() string {
	switch s {
	case UnitSizeTypeIndependent:
		return "Independent"
	case UnitSizeTypeSmall:
		return "Small"
	case UnitSizeTypeMedium:
		return "Medium"
	case UnitSizeTypeLarge:
		return "Large"
	case UnitSizeTypeNone:
		return "None"
	case UnitSizeTypeUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

package bwapi

// ExplosionType represents a BWAPI explosion type.
type ExplosionType int32

const (
	ExplosionTypeNone             ExplosionType = 0
	ExplosionTypeNormal           ExplosionType = 1
	ExplosionTypeRadialSplash     ExplosionType = 2
	ExplosionTypeEnemySplash      ExplosionType = 3
	ExplosionTypeLockdown         ExplosionType = 4
	ExplosionTypeNuclearMissile   ExplosionType = 5
	ExplosionTypeParasite         ExplosionType = 6
	ExplosionTypeBroodlings       ExplosionType = 7
	ExplosionTypeEMPShockwave     ExplosionType = 8
	ExplosionTypeIrradiate        ExplosionType = 9
	ExplosionTypeEnsnare          ExplosionType = 10
	ExplosionTypePlague           ExplosionType = 11
	ExplosionTypeStasisField      ExplosionType = 12
	ExplosionTypeDarkSwarm        ExplosionType = 13
	ExplosionTypeConsume          ExplosionType = 14
	ExplosionTypeYamatoGun        ExplosionType = 15
	ExplosionTypeRestoration      ExplosionType = 16
	ExplosionTypeDisruptionWeb    ExplosionType = 17
	ExplosionTypeCorrosiveAcid    ExplosionType = 18
	ExplosionTypeMindControl      ExplosionType = 19
	ExplosionTypeFeedback         ExplosionType = 20
	ExplosionTypeOpticalFlare     ExplosionType = 21
	ExplosionTypeMaelstrom        ExplosionType = 22
	ExplosionTypeAirSplash        ExplosionType = 25
	ExplosionTypeUnknown          ExplosionType = 26
)

func (e ExplosionType) String() string {
	switch e {
	case ExplosionTypeNone:
		return "None"
	case ExplosionTypeNormal:
		return "Normal"
	case ExplosionTypeRadialSplash:
		return "Radial_Splash"
	case ExplosionTypeEnemySplash:
		return "Enemy_Splash"
	case ExplosionTypeLockdown:
		return "Lockdown"
	case ExplosionTypeNuclearMissile:
		return "Nuclear_Missile"
	case ExplosionTypeParasite:
		return "Parasite"
	case ExplosionTypeBroodlings:
		return "Broodlings"
	case ExplosionTypeEMPShockwave:
		return "EMP_Shockwave"
	case ExplosionTypeIrradiate:
		return "Irradiate"
	case ExplosionTypeEnsnare:
		return "Ensnare"
	case ExplosionTypePlague:
		return "Plague"
	case ExplosionTypeStasisField:
		return "Stasis_Field"
	case ExplosionTypeDarkSwarm:
		return "Dark_Swarm"
	case ExplosionTypeConsume:
		return "Consume"
	case ExplosionTypeYamatoGun:
		return "Yamato_Gun"
	case ExplosionTypeRestoration:
		return "Restoration"
	case ExplosionTypeDisruptionWeb:
		return "Disruption_Web"
	case ExplosionTypeCorrosiveAcid:
		return "Corrosive_Acid"
	case ExplosionTypeMindControl:
		return "Mind_Control"
	case ExplosionTypeFeedback:
		return "Feedback"
	case ExplosionTypeOpticalFlare:
		return "Optical_Flare"
	case ExplosionTypeMaelstrom:
		return "Maelstrom"
	case ExplosionTypeAirSplash:
		return "Air_Splash"
	case ExplosionTypeUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

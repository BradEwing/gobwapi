package bwapi

// TechType represents a BWAPI technology type.
type TechType int32

const (
	TechTypeStimPacks        TechType = 0
	TechTypeLockdown         TechType = 1
	TechTypeEMPShockwave     TechType = 2
	TechTypeSpiderMines      TechType = 3
	TechTypeScannerSweep     TechType = 4
	TechTypeTankSiegeMode    TechType = 5
	TechTypeDefensiveMatrix  TechType = 6
	TechTypeIrradiate        TechType = 7
	TechTypeYamatoGun        TechType = 8
	TechTypeCloakingField    TechType = 9
	TechTypePersonnelCloaking TechType = 10
	TechTypeBurrowing        TechType = 11
	TechTypeInfestation      TechType = 12
	TechTypeSpawnBroodlings  TechType = 13
	TechTypeDarkSwarm        TechType = 14
	TechTypePlague           TechType = 15
	TechTypeConsume          TechType = 16
	TechTypeEnsnare          TechType = 17
	TechTypeParasite         TechType = 18
	TechTypePsionicStorm     TechType = 19
	TechTypeHallucination    TechType = 20
	TechTypeRecall           TechType = 21
	TechTypeStasisField      TechType = 22
	TechTypeArchonWarp       TechType = 23
	TechTypeRestoration      TechType = 24
	TechTypeDisruptionWeb    TechType = 25
	TechTypeUnused26         TechType = 26
	TechTypeMindControl      TechType = 27
	TechTypeDarkArchonMeld   TechType = 28
	TechTypeFeedback         TechType = 29
	TechTypeOpticalFlare     TechType = 30
	TechTypeMaelstrom        TechType = 31
	TechTypeLurkerAspect     TechType = 32
	TechTypeUnused33         TechType = 33
	TechTypeHealing          TechType = 34
	TechTypeNone             TechType = 44
	TechTypeNuclearStrike    TechType = 45
	TechTypeUnknown          TechType = 46
)

func (t TechType) String() string {
	switch t {
	case TechTypeStimPacks:
		return "StimPacks"
	case TechTypeLockdown:
		return "Lockdown"
	case TechTypeEMPShockwave:
		return "EMPShockwave"
	case TechTypeSpiderMines:
		return "SpiderMines"
	case TechTypeScannerSweep:
		return "ScannerSweep"
	case TechTypeTankSiegeMode:
		return "TankSiegeMode"
	case TechTypeDefensiveMatrix:
		return "DefensiveMatrix"
	case TechTypeIrradiate:
		return "Irradiate"
	case TechTypeYamatoGun:
		return "YamatoGun"
	case TechTypeCloakingField:
		return "CloakingField"
	case TechTypePersonnelCloaking:
		return "PersonnelCloaking"
	case TechTypeBurrowing:
		return "Burrowing"
	case TechTypeInfestation:
		return "Infestation"
	case TechTypeSpawnBroodlings:
		return "SpawnBroodlings"
	case TechTypeDarkSwarm:
		return "DarkSwarm"
	case TechTypePlague:
		return "Plague"
	case TechTypeConsume:
		return "Consume"
	case TechTypeEnsnare:
		return "Ensnare"
	case TechTypeParasite:
		return "Parasite"
	case TechTypePsionicStorm:
		return "PsionicStorm"
	case TechTypeHallucination:
		return "Hallucination"
	case TechTypeRecall:
		return "Recall"
	case TechTypeStasisField:
		return "StasisField"
	case TechTypeArchonWarp:
		return "ArchonWarp"
	case TechTypeRestoration:
		return "Restoration"
	case TechTypeDisruptionWeb:
		return "DisruptionWeb"
	case TechTypeUnused26:
		return "Unused26"
	case TechTypeMindControl:
		return "MindControl"
	case TechTypeDarkArchonMeld:
		return "DarkArchonMeld"
	case TechTypeFeedback:
		return "Feedback"
	case TechTypeOpticalFlare:
		return "OpticalFlare"
	case TechTypeMaelstrom:
		return "Maelstrom"
	case TechTypeLurkerAspect:
		return "LurkerAspect"
	case TechTypeUnused33:
		return "Unused33"
	case TechTypeHealing:
		return "Healing"
	case TechTypeNone:
		return "None"
	case TechTypeNuclearStrike:
		return "NuclearStrike"
	case TechTypeUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

package bwapi

// WeaponType represents a BWAPI weapon type.
type WeaponType int32

const (
	WeaponTypeGaussRifle                     WeaponType = 0
	WeaponTypeGaussRifleJimRaynor            WeaponType = 1
	WeaponTypeC10CanisterRifle               WeaponType = 2
	WeaponTypeC10CanisterRifleSarahKerrigan  WeaponType = 3
	WeaponTypeFragmentationGrenade           WeaponType = 4
	WeaponTypeFragmentationGrenadeJimRaynor  WeaponType = 5
	WeaponTypeSpiderMines                    WeaponType = 6
	WeaponTypeTwinAutocannons                WeaponType = 7
	WeaponTypeHellfireMissilePack            WeaponType = 8
	WeaponTypeTwinAutocannonsAlanSchezar     WeaponType = 9
	WeaponTypeHellfireMissilePackAlanSchezar WeaponType = 10
	WeaponTypeArcliteCannon                  WeaponType = 11
	WeaponTypeArcliteCannonEdmundDuke        WeaponType = 12
	WeaponTypeFusionCutter                   WeaponType = 13
	// 14 is skipped
	WeaponTypeGeminiMissiles               WeaponType = 15
	WeaponTypeBurstLasers                  WeaponType = 16
	WeaponTypeGeminiMissilesTomKazansky    WeaponType = 17
	WeaponTypeBurstLasersTomKazansky       WeaponType = 18
	WeaponTypeATSLaserBattery              WeaponType = 19
	WeaponTypeATALaserBattery              WeaponType = 20
	WeaponTypeATSLaserBatteryHero          WeaponType = 21
	WeaponTypeATALaserBatteryHero          WeaponType = 22
	WeaponTypeATSLaserBatteryHyperion      WeaponType = 23
	WeaponTypeATALaserBatteryHyperion      WeaponType = 24
	WeaponTypeFlameThrower                 WeaponType = 25
	WeaponTypeFlameThrowerGuiMontag        WeaponType = 26
	WeaponTypeArcliteShockCannon           WeaponType = 27
	WeaponTypeArcliteShockCannonEdmundDuke WeaponType = 28
	WeaponTypeLongboltMissile              WeaponType = 29
	WeaponTypeYamatoGun                    WeaponType = 30
	WeaponTypeNuclearStrike                WeaponType = 31
	WeaponTypeLockdown                     WeaponType = 32
	WeaponTypeEMPShockwave                 WeaponType = 33
	WeaponTypeIrradiate                    WeaponType = 34
	WeaponTypeClaws                        WeaponType = 35
	WeaponTypeClawsDevouringOne            WeaponType = 36
	WeaponTypeClawsInfestedKerrigan        WeaponType = 37
	WeaponTypeNeedleSpines                 WeaponType = 38
	WeaponTypeNeedleSpinesHunterKiller     WeaponType = 39
	WeaponTypeKaiserBlades                 WeaponType = 40
	WeaponTypeKaiserBladesTorrasque        WeaponType = 41
	WeaponTypeToxicSpores                  WeaponType = 42
	WeaponTypeSpines                       WeaponType = 43
	// 44-45 are skipped
	WeaponTypeAcidSpore        WeaponType = 46
	WeaponTypeAcidSporeKukulza WeaponType = 47
	WeaponTypeGlaveWurm        WeaponType = 48
	WeaponTypeGlaveWurmKukulza WeaponType = 49
	// 50-51 are skipped
	WeaponTypeSeekerSpores          WeaponType = 52
	WeaponTypeSubterraneanTentacle  WeaponType = 53
	WeaponTypeSuicideInfestedTerran WeaponType = 54
	WeaponTypeSuicideScourge        WeaponType = 55
	WeaponTypeParasite              WeaponType = 56
	WeaponTypeSpawnBroodlings       WeaponType = 57
	WeaponTypeEnsnare               WeaponType = 58
	WeaponTypeDarkSwarm             WeaponType = 59
	WeaponTypePlague                WeaponType = 60
	WeaponTypeConsume               WeaponType = 61
	WeaponTypeParticleBeam          WeaponType = 62
	// 63 is skipped
	WeaponTypePsiBlades           WeaponType = 64
	WeaponTypePsiBladesFenix      WeaponType = 65
	WeaponTypePhaseDisruptor      WeaponType = 66
	WeaponTypePhaseDisruptorFenix WeaponType = 67
	// 68 is skipped
	WeaponTypePsiAssault               WeaponType = 69
	WeaponTypePsionicShockwave         WeaponType = 70
	WeaponTypePsionicShockwaveTZArchon WeaponType = 71
	// 72 is skipped
	WeaponTypeDualPhotonBlasters           WeaponType = 73
	WeaponTypeAntiMatterMissiles           WeaponType = 74
	WeaponTypeDualPhotonBlastersMojo       WeaponType = 75
	WeaponTypeAntiMatterMissilesMojo       WeaponType = 76
	WeaponTypePhaseDisruptorCannon         WeaponType = 77
	WeaponTypePhaseDisruptorCannonDanimoth WeaponType = 78
	WeaponTypePulseCannon                  WeaponType = 79
	WeaponTypeSTSPhotonCannon              WeaponType = 80
	WeaponTypeSTAPhotonCannon              WeaponType = 81
	WeaponTypeScarab                       WeaponType = 82
	WeaponTypeStasisField                  WeaponType = 83
	WeaponTypePsionicStorm                 WeaponType = 84
	WeaponTypeWarpBladesZeratul            WeaponType = 85
	WeaponTypeWarpBladesHero               WeaponType = 86
	// 87-91 are skipped
	WeaponTypePlatformLaserBattery    WeaponType = 92
	WeaponTypeIndependantLaserBattery WeaponType = 93
	// 94-95 are skipped
	WeaponTypeTwinAutocannonsFloorTrap     WeaponType = 96
	WeaponTypeHellfireMissilePackWallTrap  WeaponType = 97
	WeaponTypeFlameThrowerWallTrap         WeaponType = 98
	WeaponTypeHellfireMissilePackFloorTrap WeaponType = 99
	WeaponTypeNeutronFlare                 WeaponType = 100
	WeaponTypeDisruptionWeb                WeaponType = 101
	WeaponTypeRestoration                  WeaponType = 102
	WeaponTypeHaloRockets                  WeaponType = 103
	WeaponTypeCorrosiveAcid                WeaponType = 104
	WeaponTypeMindControl                  WeaponType = 105
	WeaponTypeFeedback                     WeaponType = 106
	WeaponTypeOpticalFlare                 WeaponType = 107
	WeaponTypeMaelstrom                    WeaponType = 108
	WeaponTypeSubterraneanSpines           WeaponType = 109
	// 110 is skipped
	WeaponTypeWarpBlades                    WeaponType = 111
	WeaponTypeC10CanisterRifleSamirDuran    WeaponType = 112
	WeaponTypeC10CanisterRifleInfestedDuran WeaponType = 113
	WeaponTypeDualPhotonBlastersArtanis     WeaponType = 114
	WeaponTypeAntiMatterMissilesArtanis     WeaponType = 115
	WeaponTypeC10CanisterRifleAlexeiStukov  WeaponType = 116
	// 117-129 are skipped
	WeaponTypeNone    WeaponType = 130
	WeaponTypeUnknown WeaponType = 131
)

func (w WeaponType) String() string {
	switch w {
	case WeaponTypeGaussRifle:
		return "GaussRifle"
	case WeaponTypeGaussRifleJimRaynor:
		return "GaussRifleJimRaynor"
	case WeaponTypeC10CanisterRifle:
		return "C10CanisterRifle"
	case WeaponTypeC10CanisterRifleSarahKerrigan:
		return "C10CanisterRifleSarahKerrigan"
	case WeaponTypeFragmentationGrenade:
		return "FragmentationGrenade"
	case WeaponTypeFragmentationGrenadeJimRaynor:
		return "FragmentationGrenadeJimRaynor"
	case WeaponTypeSpiderMines:
		return "SpiderMines"
	case WeaponTypeTwinAutocannons:
		return "TwinAutocannons"
	case WeaponTypeHellfireMissilePack:
		return "HellfireMissilePack"
	case WeaponTypeTwinAutocannonsAlanSchezar:
		return "TwinAutocannonsAlanSchezar"
	case WeaponTypeHellfireMissilePackAlanSchezar:
		return "HellfireMissilePackAlanSchezar"
	case WeaponTypeArcliteCannon:
		return "ArcliteCannon"
	case WeaponTypeArcliteCannonEdmundDuke:
		return "ArcliteCannonEdmundDuke"
	case WeaponTypeFusionCutter:
		return "FusionCutter"
	case WeaponTypeGeminiMissiles:
		return "GeminiMissiles"
	case WeaponTypeBurstLasers:
		return "BurstLasers"
	case WeaponTypeGeminiMissilesTomKazansky:
		return "GeminiMissilesTomKazansky"
	case WeaponTypeBurstLasersTomKazansky:
		return "BurstLasersTomKazansky"
	case WeaponTypeATSLaserBattery:
		return "ATSLaserBattery"
	case WeaponTypeATALaserBattery:
		return "ATALaserBattery"
	case WeaponTypeATSLaserBatteryHero:
		return "ATSLaserBatteryHero"
	case WeaponTypeATALaserBatteryHero:
		return "ATALaserBatteryHero"
	case WeaponTypeATSLaserBatteryHyperion:
		return "ATSLaserBatteryHyperion"
	case WeaponTypeATALaserBatteryHyperion:
		return "ATALaserBatteryHyperion"
	case WeaponTypeFlameThrower:
		return "FlameThrower"
	case WeaponTypeFlameThrowerGuiMontag:
		return "FlameThrowerGuiMontag"
	case WeaponTypeArcliteShockCannon:
		return "ArcliteShockCannon"
	case WeaponTypeArcliteShockCannonEdmundDuke:
		return "ArcliteShockCannonEdmundDuke"
	case WeaponTypeLongboltMissile:
		return "LongboltMissile"
	case WeaponTypeYamatoGun:
		return "YamatoGun"
	case WeaponTypeNuclearStrike:
		return "NuclearStrike"
	case WeaponTypeLockdown:
		return "Lockdown"
	case WeaponTypeEMPShockwave:
		return "EMPShockwave"
	case WeaponTypeIrradiate:
		return "Irradiate"
	case WeaponTypeClaws:
		return "Claws"
	case WeaponTypeClawsDevouringOne:
		return "ClawsDevouringOne"
	case WeaponTypeClawsInfestedKerrigan:
		return "ClawsInfestedKerrigan"
	case WeaponTypeNeedleSpines:
		return "NeedleSpines"
	case WeaponTypeNeedleSpinesHunterKiller:
		return "NeedleSpinesHunterKiller"
	case WeaponTypeKaiserBlades:
		return "KaiserBlades"
	case WeaponTypeKaiserBladesTorrasque:
		return "KaiserBladesTorrasque"
	case WeaponTypeToxicSpores:
		return "ToxicSpores"
	case WeaponTypeSpines:
		return "Spines"
	case WeaponTypeAcidSpore:
		return "AcidSpore"
	case WeaponTypeAcidSporeKukulza:
		return "AcidSporeKukulza"
	case WeaponTypeGlaveWurm:
		return "GlaveWurm"
	case WeaponTypeGlaveWurmKukulza:
		return "GlaveWurmKukulza"
	case WeaponTypeSeekerSpores:
		return "SeekerSpores"
	case WeaponTypeSubterraneanTentacle:
		return "SubterraneanTentacle"
	case WeaponTypeSuicideInfestedTerran:
		return "SuicideInfestedTerran"
	case WeaponTypeSuicideScourge:
		return "SuicideScourge"
	case WeaponTypeParasite:
		return "Parasite"
	case WeaponTypeSpawnBroodlings:
		return "SpawnBroodlings"
	case WeaponTypeEnsnare:
		return "Ensnare"
	case WeaponTypeDarkSwarm:
		return "DarkSwarm"
	case WeaponTypePlague:
		return "Plague"
	case WeaponTypeConsume:
		return "Consume"
	case WeaponTypeParticleBeam:
		return "ParticleBeam"
	case WeaponTypePsiBlades:
		return "PsiBlades"
	case WeaponTypePsiBladesFenix:
		return "PsiBladesFenix"
	case WeaponTypePhaseDisruptor:
		return "PhaseDisruptor"
	case WeaponTypePhaseDisruptorFenix:
		return "PhaseDisruptorFenix"
	case WeaponTypePsiAssault:
		return "PsiAssault"
	case WeaponTypePsionicShockwave:
		return "PsionicShockwave"
	case WeaponTypePsionicShockwaveTZArchon:
		return "PsionicShockwaveTZArchon"
	case WeaponTypeDualPhotonBlasters:
		return "DualPhotonBlasters"
	case WeaponTypeAntiMatterMissiles:
		return "AntiMatterMissiles"
	case WeaponTypeDualPhotonBlastersMojo:
		return "DualPhotonBlastersMojo"
	case WeaponTypeAntiMatterMissilesMojo:
		return "AntiMatterMissilesMojo"
	case WeaponTypePhaseDisruptorCannon:
		return "PhaseDisruptorCannon"
	case WeaponTypePhaseDisruptorCannonDanimoth:
		return "PhaseDisruptorCannonDanimoth"
	case WeaponTypePulseCannon:
		return "PulseCannon"
	case WeaponTypeSTSPhotonCannon:
		return "STSPhotonCannon"
	case WeaponTypeSTAPhotonCannon:
		return "STAPhotonCannon"
	case WeaponTypeScarab:
		return "Scarab"
	case WeaponTypeStasisField:
		return "StasisField"
	case WeaponTypePsionicStorm:
		return "PsionicStorm"
	case WeaponTypeWarpBladesZeratul:
		return "WarpBladesZeratul"
	case WeaponTypeWarpBladesHero:
		return "WarpBladesHero"
	case WeaponTypePlatformLaserBattery:
		return "PlatformLaserBattery"
	case WeaponTypeIndependantLaserBattery:
		return "IndependantLaserBattery"
	case WeaponTypeTwinAutocannonsFloorTrap:
		return "TwinAutocannonsFloorTrap"
	case WeaponTypeHellfireMissilePackWallTrap:
		return "HellfireMissilePackWallTrap"
	case WeaponTypeFlameThrowerWallTrap:
		return "FlameThrowerWallTrap"
	case WeaponTypeHellfireMissilePackFloorTrap:
		return "HellfireMissilePackFloorTrap"
	case WeaponTypeNeutronFlare:
		return "NeutronFlare"
	case WeaponTypeDisruptionWeb:
		return "DisruptionWeb"
	case WeaponTypeRestoration:
		return "Restoration"
	case WeaponTypeHaloRockets:
		return "HaloRockets"
	case WeaponTypeCorrosiveAcid:
		return "CorrosiveAcid"
	case WeaponTypeMindControl:
		return "MindControl"
	case WeaponTypeFeedback:
		return "Feedback"
	case WeaponTypeOpticalFlare:
		return "OpticalFlare"
	case WeaponTypeMaelstrom:
		return "Maelstrom"
	case WeaponTypeSubterraneanSpines:
		return "SubterraneanSpines"
	case WeaponTypeWarpBlades:
		return "WarpBlades"
	case WeaponTypeC10CanisterRifleSamirDuran:
		return "C10CanisterRifleSamirDuran"
	case WeaponTypeC10CanisterRifleInfestedDuran:
		return "C10CanisterRifleInfestedDuran"
	case WeaponTypeDualPhotonBlastersArtanis:
		return "DualPhotonBlastersArtanis"
	case WeaponTypeAntiMatterMissilesArtanis:
		return "AntiMatterMissilesArtanis"
	case WeaponTypeC10CanisterRifleAlexeiStukov:
		return "C10CanisterRifleAlexeiStukov"
	case WeaponTypeNone:
		return "None"
	case WeaponTypeUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

package bwapi

// BulletType represents a BWAPI bullet/projectile type.
type BulletType int32

const (
	BulletTypeMelee                 BulletType = 0
	BulletTypeFusionCutterHit       BulletType = 141
	BulletTypeGaussRifleHit         BulletType = 142
	BulletTypeC10CanisterRifleHit   BulletType = 143
	BulletTypeGeminiMissiles        BulletType = 144
	BulletTypeFragmentationGrenade  BulletType = 145
	BulletTypeLongboltMissile       BulletType = 146
	BulletTypeUnusedLockdown        BulletType = 147
	BulletTypeATSATALaserBattery    BulletType = 148
	BulletTypeBurstLasers           BulletType = 149
	BulletTypeArcliteShockCannonHit BulletType = 150
	BulletTypeEMPMissile            BulletType = 151
	BulletTypeDualPhotonBlastersHit BulletType = 152
	BulletTypeParticleBeamHit       BulletType = 153
	BulletTypeAntiMatterMissile     BulletType = 154
	BulletTypePulseCannon           BulletType = 155
	BulletTypePsionicShockwaveHit   BulletType = 156
	BulletTypePsionicStorm          BulletType = 157
	BulletTypeYamatoGun             BulletType = 158
	BulletTypePhaseDisruptor        BulletType = 159
	BulletTypeSTASTSCannonOverlay   BulletType = 160
	BulletTypeSunkenColonyTentacle  BulletType = 161
	BulletTypeVenomUnused           BulletType = 162
	BulletTypeAcidSpore             BulletType = 163
	BulletTypePlasmaDripUnused      BulletType = 164
	BulletTypeGlaveWurm             BulletType = 165
	BulletTypeSeekerSpores          BulletType = 166
	BulletTypeQueenSpellCarrier     BulletType = 167
	BulletTypePlagueCloud           BulletType = 168
	BulletTypeConsume               BulletType = 169
	BulletTypeEnsnare               BulletType = 170
	BulletTypeNeedleSpineHit        BulletType = 171
	BulletTypeInvisible             BulletType = 172
	BulletTypeOpticalFlareGrenade   BulletType = 201
	BulletTypeHaloRockets           BulletType = 202
	BulletTypeSubterraneanSpines    BulletType = 203
	BulletTypeCorrosiveAcidShot     BulletType = 204
	BulletTypeCorrosiveAcidHit      BulletType = 205
	BulletTypeNeutronFlare          BulletType = 206
	BulletTypeNone                  BulletType = 209
	BulletTypeUnknown               BulletType = 210
)

func (b BulletType) String() string {
	switch b {
	case BulletTypeMelee:
		return "Melee"
	case BulletTypeFusionCutterHit:
		return "FusionCutterHit"
	case BulletTypeGaussRifleHit:
		return "GaussRifleHit"
	case BulletTypeC10CanisterRifleHit:
		return "C10CanisterRifleHit"
	case BulletTypeGeminiMissiles:
		return "GeminiMissiles"
	case BulletTypeFragmentationGrenade:
		return "FragmentationGrenade"
	case BulletTypeLongboltMissile:
		return "LongboltMissile"
	case BulletTypeUnusedLockdown:
		return "UnusedLockdown"
	case BulletTypeATSATALaserBattery:
		return "ATSATALaserBattery"
	case BulletTypeBurstLasers:
		return "BurstLasers"
	case BulletTypeArcliteShockCannonHit:
		return "ArcliteShockCannonHit"
	case BulletTypeEMPMissile:
		return "EMPMissile"
	case BulletTypeDualPhotonBlastersHit:
		return "DualPhotonBlastersHit"
	case BulletTypeParticleBeamHit:
		return "ParticleBeamHit"
	case BulletTypeAntiMatterMissile:
		return "AntiMatterMissile"
	case BulletTypePulseCannon:
		return "PulseCannon"
	case BulletTypePsionicShockwaveHit:
		return "PsionicShockwaveHit"
	case BulletTypePsionicStorm:
		return "PsionicStorm"
	case BulletTypeYamatoGun:
		return "YamatoGun"
	case BulletTypePhaseDisruptor:
		return "PhaseDisruptor"
	case BulletTypeSTASTSCannonOverlay:
		return "STASTSCannonOverlay"
	case BulletTypeSunkenColonyTentacle:
		return "SunkenColonyTentacle"
	case BulletTypeVenomUnused:
		return "VenomUnused"
	case BulletTypeAcidSpore:
		return "AcidSpore"
	case BulletTypePlasmaDripUnused:
		return "PlasmaDripUnused"
	case BulletTypeGlaveWurm:
		return "GlaveWurm"
	case BulletTypeSeekerSpores:
		return "SeekerSpores"
	case BulletTypeQueenSpellCarrier:
		return "QueenSpellCarrier"
	case BulletTypePlagueCloud:
		return "PlagueCloud"
	case BulletTypeConsume:
		return "Consume"
	case BulletTypeEnsnare:
		return "Ensnare"
	case BulletTypeNeedleSpineHit:
		return "NeedleSpineHit"
	case BulletTypeInvisible:
		return "Invisible"
	case BulletTypeOpticalFlareGrenade:
		return "OpticalFlareGrenade"
	case BulletTypeHaloRockets:
		return "HaloRockets"
	case BulletTypeSubterraneanSpines:
		return "SubterraneanSpines"
	case BulletTypeCorrosiveAcidShot:
		return "CorrosiveAcidShot"
	case BulletTypeCorrosiveAcidHit:
		return "CorrosiveAcidHit"
	case BulletTypeNeutronFlare:
		return "NeutronFlare"
	case BulletTypeNone:
		return "None"
	case BulletTypeUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

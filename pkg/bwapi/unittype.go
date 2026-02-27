package bwapi

// UnitType represents a StarCraft unit type.
type UnitType int32

const (
	UnitTypeTerranMarine                   UnitType = 0
	UnitTypeTerranGhost                    UnitType = 1
	UnitTypeTerranVulture                  UnitType = 2
	UnitTypeTerranGoliath                  UnitType = 3
	UnitTypeTerranGoliathTurret            UnitType = 4
	UnitTypeTerranSiegeTankTankMode        UnitType = 5
	UnitTypeTerranSiegeTankTankModeTurret  UnitType = 6
	UnitTypeTerranSCV                      UnitType = 7
	UnitTypeTerranWraith                   UnitType = 8
	UnitTypeTerranScienceVessel            UnitType = 9
	UnitTypeHeroGuiMontag                  UnitType = 10
	UnitTypeTerranDropship                 UnitType = 11
	UnitTypeTerranBattlecruiser            UnitType = 12
	UnitTypeTerranVultureSpiderMine        UnitType = 13
	UnitTypeTerranNuclearMissile           UnitType = 14
	UnitTypeTerranCivilian                 UnitType = 15
	UnitTypeHeroSarahKerrigan              UnitType = 16
	UnitTypeHeroAlanSchezar                UnitType = 17
	UnitTypeHeroAlanSchezarTurret          UnitType = 18
	UnitTypeHeroJimRaynorVulture           UnitType = 19
	UnitTypeHeroJimRaynorMarine            UnitType = 20
	UnitTypeHeroTomKazansky                UnitType = 21
	UnitTypeHeroMagellan                   UnitType = 22
	UnitTypeHeroEdmundDukeTankMode         UnitType = 23
	UnitTypeHeroEdmundDukeTankModeTurret   UnitType = 24
	UnitTypeHeroEdmundDukeSiegeMode        UnitType = 25
	UnitTypeHeroEdmundDukeSiegeModeTurret  UnitType = 26
	UnitTypeHeroArcturusMengsk             UnitType = 27
	UnitTypeHeroHyperion                   UnitType = 28
	UnitTypeHeroNoradII                    UnitType = 29
	UnitTypeTerranSiegeTankSiegeMode       UnitType = 30
	UnitTypeTerranSiegeTankSiegeModeTurret UnitType = 31
	UnitTypeTerranFirebat                  UnitType = 32
	UnitTypeSpellScannerSweep              UnitType = 33
	UnitTypeTerranMedic                    UnitType = 34
	UnitTypeZergLarva                      UnitType = 35
	UnitTypeZergEgg                        UnitType = 36
	UnitTypeZergZergling                   UnitType = 37
	UnitTypeZergHydralisk                  UnitType = 38
	UnitTypeZergUltralisk                  UnitType = 39
	UnitTypeZergBroodling                  UnitType = 40
	UnitTypeZergDrone                      UnitType = 41
	UnitTypeZergOverlord                   UnitType = 42
	UnitTypeZergMutalisk                   UnitType = 43
	UnitTypeZergGuardian                   UnitType = 44
	UnitTypeZergQueen                      UnitType = 45
	UnitTypeZergDefiler                    UnitType = 46
	UnitTypeZergScourge                    UnitType = 47
	UnitTypeHeroTorrasque                  UnitType = 48
	UnitTypeHeroMatriarch                  UnitType = 49
	UnitTypeZergInfestedTerran             UnitType = 50
	UnitTypeHeroInfestedKerrigan           UnitType = 51
	UnitTypeHeroUncleanOne                 UnitType = 52
	UnitTypeHeroHunterKiller               UnitType = 53
	UnitTypeHeroDevouringOne               UnitType = 54
	UnitTypeHeroKukulzaMutalisk            UnitType = 55
	UnitTypeHeroKukulzaGuardian            UnitType = 56
	UnitTypeHeroYggdrasill                 UnitType = 57
	UnitTypeTerranValkyrie                 UnitType = 58
	UnitTypeZergCocoon                     UnitType = 59
	UnitTypeProtossCorsair                 UnitType = 60
	UnitTypeProtossDarkTemplar             UnitType = 61
	UnitTypeZergDevourer                   UnitType = 62
	UnitTypeProtossDarkArchon              UnitType = 63
	UnitTypeProtossProbe                   UnitType = 64
	UnitTypeProtossZealot                  UnitType = 65
	UnitTypeProtossDragoon                 UnitType = 66
	UnitTypeProtossHighTemplar             UnitType = 67
	UnitTypeProtossArchon                  UnitType = 68
	UnitTypeProtossShuttle                 UnitType = 69
	UnitTypeProtossScout                   UnitType = 70
	UnitTypeProtossArbiter                 UnitType = 71
	UnitTypeProtossCarrier                 UnitType = 72
	UnitTypeProtossInterceptor             UnitType = 73
	UnitTypeHeroDarkTemplar                UnitType = 74
	UnitTypeHeroZeratul                    UnitType = 75
	UnitTypeHeroTassadarZeratulArchon      UnitType = 76
	UnitTypeHeroFenixZealot                UnitType = 77
	UnitTypeHeroFenixDragoon               UnitType = 78
	UnitTypeHeroTassadar                   UnitType = 79
	UnitTypeHeroMojo                       UnitType = 80
	UnitTypeHeroWarbringer                 UnitType = 81
	UnitTypeHeroGantrithor                 UnitType = 82
	UnitTypeProtossReaver                  UnitType = 83
	UnitTypeProtossObserver                UnitType = 84
	UnitTypeProtossScarab                  UnitType = 85
	UnitTypeHeroDanimoth                   UnitType = 86
	UnitTypeHeroAldaris                    UnitType = 87
	UnitTypeHeroArtanis                    UnitType = 88
	UnitTypeCritterRhynadon                UnitType = 89
	UnitTypeCritterBengalaas               UnitType = 90
	UnitTypeSpecialCargoShip               UnitType = 91
	UnitTypeSpecialMercenaryGunship        UnitType = 92
	UnitTypeCritterScantid                 UnitType = 93
	UnitTypeCritterKakaru                  UnitType = 94
	UnitTypeCritterRagnasaur               UnitType = 95
	UnitTypeCritterUrsadon                 UnitType = 96
	UnitTypeZergLurkerEgg                  UnitType = 97
	UnitTypeHeroRaszagal                   UnitType = 98
	UnitTypeHeroSamirDuran                 UnitType = 99
	UnitTypeHeroAlexeiStukov               UnitType = 100
	UnitTypeSpecialMapRevealer             UnitType = 101
	UnitTypeHeroGerardDuGalle              UnitType = 102
	UnitTypeZergLurker                     UnitType = 103
	UnitTypeHeroInfestedDuran              UnitType = 104
	UnitTypeSpellDisruptionWeb             UnitType = 105
	UnitTypeTerranCommandCenter            UnitType = 106
	UnitTypeTerranComsatStation            UnitType = 107
	UnitTypeTerranNuclearSilo              UnitType = 108
	UnitTypeTerranSupplyDepot              UnitType = 109
	UnitTypeTerranRefinery                 UnitType = 110
	UnitTypeTerranBarracks                 UnitType = 111
	UnitTypeTerranAcademy                  UnitType = 112
	UnitTypeTerranFactory                  UnitType = 113
	UnitTypeTerranStarport                 UnitType = 114
	UnitTypeTerranControlTower             UnitType = 115
	UnitTypeTerranScienceFacility          UnitType = 116
	UnitTypeTerranCovertOps                UnitType = 117
	UnitTypeTerranPhysicsLab               UnitType = 118
	UnitTypeUnusedTerran1                  UnitType = 119
	UnitTypeTerranMachineShop              UnitType = 120
	UnitTypeUnusedTerran2                  UnitType = 121
	UnitTypeTerranEngineeringBay           UnitType = 122
	UnitTypeTerranArmory                   UnitType = 123
	UnitTypeTerranMissileTurret            UnitType = 124
	UnitTypeTerranBunker                   UnitType = 125
	UnitTypeSpecialCrashedNoradII          UnitType = 126
	UnitTypeSpecialIonCannon               UnitType = 127
	UnitTypePowerupUrajCrystal             UnitType = 128
	UnitTypePowerupKhalisCrystal           UnitType = 129
	UnitTypeZergInfestedCommandCenter      UnitType = 130
	UnitTypeZergHatchery                   UnitType = 131
	UnitTypeZergLair                       UnitType = 132
	UnitTypeZergHive                       UnitType = 133
	UnitTypeZergNydusCanal                 UnitType = 134
	UnitTypeZergHydraliskDen               UnitType = 135
	UnitTypeZergDefilerMound               UnitType = 136
	UnitTypeZergGreaterSpire               UnitType = 137
	UnitTypeZergQueensNest                 UnitType = 138
	UnitTypeZergEvolutionChamber           UnitType = 139
	UnitTypeZergUltraliskCavern            UnitType = 140
	UnitTypeZergSpire                      UnitType = 141
	UnitTypeZergSpawningPool               UnitType = 142
	UnitTypeZergCreepColony                UnitType = 143
	UnitTypeZergSporeColony                UnitType = 144
	UnitTypeUnusedZerg1                    UnitType = 145
	UnitTypeZergSunkenColony               UnitType = 146
	UnitTypeSpecialOvermindWithShell       UnitType = 147
	UnitTypeSpecialOvermind                UnitType = 148
	UnitTypeZergExtractor                  UnitType = 149
	UnitTypeSpecialMatureChrysalis         UnitType = 150
	UnitTypeSpecialCerebrate               UnitType = 151
	UnitTypeSpecialCerebrateDaggoth        UnitType = 152
	UnitTypeUnusedZerg2                    UnitType = 153
	UnitTypeProtossNexus                   UnitType = 154
	UnitTypeProtossRoboticsFacility        UnitType = 155
	UnitTypeProtossPylon                   UnitType = 156
	UnitTypeProtossAssimilator             UnitType = 157
	UnitTypeUnusedProtoss1                 UnitType = 158
	UnitTypeProtossObservatory             UnitType = 159
	UnitTypeProtossGateway                 UnitType = 160
	UnitTypeUnusedProtoss2                 UnitType = 161
	UnitTypeProtossPhotonCannon            UnitType = 162
	UnitTypeProtossCitadelOfAdun           UnitType = 163
	UnitTypeProtossCyberneticsCore         UnitType = 164
	UnitTypeProtossTemplarArchives         UnitType = 165
	UnitTypeProtossForge                   UnitType = 166
	UnitTypeProtossStargate                UnitType = 167
	UnitTypeSpecialStasisCellPrison        UnitType = 168
	UnitTypeProtossFleetBeacon             UnitType = 169
	UnitTypeProtossArbiterTribunal         UnitType = 170
	UnitTypeProtossRoboticsSupportBay      UnitType = 171
	UnitTypeProtossShieldBattery           UnitType = 172
	UnitTypeSpecialKhaydarinCrystalForm    UnitType = 173
	UnitTypeSpecialProtossTemple           UnitType = 174
	UnitTypeSpecialXelNagaTemple           UnitType = 175
	UnitTypeResourceMineralField           UnitType = 176
	UnitTypeResourceMineralFieldType2      UnitType = 177
	UnitTypeResourceMineralFieldType3      UnitType = 178
	UnitTypeUnusedCave                     UnitType = 179
	UnitTypeUnusedCaveIn                   UnitType = 180
	UnitTypeUnusedCantina                  UnitType = 181
	UnitTypeUnusedMiningPlatform           UnitType = 182
	UnitTypeUnusedIndependantCommandCenter UnitType = 183
	UnitTypeSpecialIndependantStarport     UnitType = 184
	UnitTypeUnusedIndependantJumpGate      UnitType = 185
	UnitTypeUnusedRuins                    UnitType = 186
	UnitTypeUnusedKhaydarinCrystalFormation UnitType = 187
	UnitTypeResourceVespeneGeyser          UnitType = 188
	UnitTypeSpecialWarpGate                UnitType = 189
	UnitTypeSpecialPsiDisrupter            UnitType = 190
	UnitTypeUnusedZergMarker               UnitType = 191
	UnitTypeUnusedTerranMarker             UnitType = 192
	UnitTypeUnusedProtossMarker            UnitType = 193
	UnitTypeSpecialZergBeacon              UnitType = 194
	UnitTypeSpecialTerranBeacon            UnitType = 195
	UnitTypeSpecialProtossBeacon           UnitType = 196
	UnitTypeSpecialZergFlagBeacon          UnitType = 197
	UnitTypeSpecialTerranFlagBeacon        UnitType = 198
	UnitTypeSpecialProtossFlagBeacon       UnitType = 199
	UnitTypeSpecialPowerGenerator          UnitType = 200
	UnitTypeSpecialOvermindCocoon          UnitType = 201
	UnitTypeSpellDarkSwarm                 UnitType = 202
	UnitTypeSpecialFloorMissileTrap        UnitType = 203
	UnitTypeSpecialFloorHatch              UnitType = 204
	UnitTypeSpecialUpperLevelDoor          UnitType = 205
	UnitTypeSpecialRightUpperLevelDoor     UnitType = 206
	UnitTypeSpecialPitDoor                 UnitType = 207
	UnitTypeSpecialRightPitDoor            UnitType = 208
	UnitTypeSpecialFloorGunTrap            UnitType = 209
	UnitTypeSpecialWallMissileTrap         UnitType = 210
	UnitTypeSpecialWallFlameTrap           UnitType = 211
	UnitTypeSpecialRightWallMissileTrap    UnitType = 212
	UnitTypeSpecialRightWallFlameTrap      UnitType = 213
	UnitTypeSpecialStartLocation           UnitType = 214
	UnitTypePowerupFlag                    UnitType = 215
	UnitTypePowerupYoungChrysalis          UnitType = 216
	UnitTypePowerupPsiEmitter              UnitType = 217
	UnitTypePowerupDataDisk                UnitType = 218
	UnitTypePowerupKhaydarinCrystal        UnitType = 219
	UnitTypePowerupMineralClusterType1     UnitType = 220
	UnitTypePowerupMineralClusterType2     UnitType = 221
	UnitTypePowerupProtossGasOrbType1      UnitType = 222
	UnitTypePowerupProtossGasOrbType2      UnitType = 223
	UnitTypePowerupZergGasSacType1         UnitType = 224
	UnitTypePowerupZergGasSacType2         UnitType = 225
	UnitTypePowerupTerranGasTankType1      UnitType = 226
	UnitTypePowerupTerranGasTankType2      UnitType = 227
	UnitTypeNone                           UnitType = 228
	UnitTypeAllUnits                       UnitType = 229
	UnitTypeMen                            UnitType = 230
	UnitTypeBuildings                      UnitType = 231
	UnitTypeFactories                      UnitType = 232
	UnitTypeUnknown                        UnitType = 233
)

var unitTypeNames = [234]string{
	"Terran_Marine",
	"Terran_Ghost",
	"Terran_Vulture",
	"Terran_Goliath",
	"Terran_Goliath_Turret",
	"Terran_Siege_Tank_Tank_Mode",
	"Terran_Siege_Tank_Tank_Mode_Turret",
	"Terran_SCV",
	"Terran_Wraith",
	"Terran_Science_Vessel",
	"Hero_Gui_Montag",
	"Terran_Dropship",
	"Terran_Battlecruiser",
	"Terran_Vulture_Spider_Mine",
	"Terran_Nuclear_Missile",
	"Terran_Civilian",
	"Hero_Sarah_Kerrigan",
	"Hero_Alan_Schezar",
	"Hero_Alan_Schezar_Turret",
	"Hero_Jim_Raynor_Vulture",
	"Hero_Jim_Raynor_Marine",
	"Hero_Tom_Kazansky",
	"Hero_Magellan",
	"Hero_Edmund_Duke_Tank_Mode",
	"Hero_Edmund_Duke_Tank_Mode_Turret",
	"Hero_Edmund_Duke_Siege_Mode",
	"Hero_Edmund_Duke_Siege_Mode_Turret",
	"Hero_Arcturus_Mengsk",
	"Hero_Hyperion",
	"Hero_Norad_II",
	"Terran_Siege_Tank_Siege_Mode",
	"Terran_Siege_Tank_Siege_Mode_Turret",
	"Terran_Firebat",
	"Spell_Scanner_Sweep",
	"Terran_Medic",
	"Zerg_Larva",
	"Zerg_Egg",
	"Zerg_Zergling",
	"Zerg_Hydralisk",
	"Zerg_Ultralisk",
	"Zerg_Broodling",
	"Zerg_Drone",
	"Zerg_Overlord",
	"Zerg_Mutalisk",
	"Zerg_Guardian",
	"Zerg_Queen",
	"Zerg_Defiler",
	"Zerg_Scourge",
	"Hero_Torrasque",
	"Hero_Matriarch",
	"Zerg_Infested_Terran",
	"Hero_Infested_Kerrigan",
	"Hero_Unclean_One",
	"Hero_Hunter_Killer",
	"Hero_Devouring_One",
	"Hero_Kukulza_Mutalisk",
	"Hero_Kukulza_Guardian",
	"Hero_Yggdrasill",
	"Terran_Valkyrie",
	"Zerg_Cocoon",
	"Protoss_Corsair",
	"Protoss_Dark_Templar",
	"Zerg_Devourer",
	"Protoss_Dark_Archon",
	"Protoss_Probe",
	"Protoss_Zealot",
	"Protoss_Dragoon",
	"Protoss_High_Templar",
	"Protoss_Archon",
	"Protoss_Shuttle",
	"Protoss_Scout",
	"Protoss_Arbiter",
	"Protoss_Carrier",
	"Protoss_Interceptor",
	"Hero_Dark_Templar",
	"Hero_Zeratul",
	"Hero_Tassadar_Zeratul_Archon",
	"Hero_Fenix_Zealot",
	"Hero_Fenix_Dragoon",
	"Hero_Tassadar",
	"Hero_Mojo",
	"Hero_Warbringer",
	"Hero_Gantrithor",
	"Protoss_Reaver",
	"Protoss_Observer",
	"Protoss_Scarab",
	"Hero_Danimoth",
	"Hero_Aldaris",
	"Hero_Artanis",
	"Critter_Rhynadon",
	"Critter_Bengalaas",
	"Special_Cargo_Ship",
	"Special_Mercenary_Gunship",
	"Critter_Scantid",
	"Critter_Kakaru",
	"Critter_Ragnasaur",
	"Critter_Ursadon",
	"Zerg_Lurker_Egg",
	"Hero_Raszagal",
	"Hero_Samir_Duran",
	"Hero_Alexei_Stukov",
	"Special_Map_Revealer",
	"Hero_Gerard_DuGalle",
	"Zerg_Lurker",
	"Hero_Infested_Duran",
	"Spell_Disruption_Web",
	"Terran_Command_Center",
	"Terran_Comsat_Station",
	"Terran_Nuclear_Silo",
	"Terran_Supply_Depot",
	"Terran_Refinery",
	"Terran_Barracks",
	"Terran_Academy",
	"Terran_Factory",
	"Terran_Starport",
	"Terran_Control_Tower",
	"Terran_Science_Facility",
	"Terran_Covert_Ops",
	"Terran_Physics_Lab",
	"Unused_Terran1",
	"Terran_Machine_Shop",
	"Unused_Terran2",
	"Terran_Engineering_Bay",
	"Terran_Armory",
	"Terran_Missile_Turret",
	"Terran_Bunker",
	"Special_Crashed_Norad_II",
	"Special_Ion_Cannon",
	"Powerup_Uraj_Crystal",
	"Powerup_Khalis_Crystal",
	"Zerg_Infested_Command_Center",
	"Zerg_Hatchery",
	"Zerg_Lair",
	"Zerg_Hive",
	"Zerg_Nydus_Canal",
	"Zerg_Hydralisk_Den",
	"Zerg_Defiler_Mound",
	"Zerg_Greater_Spire",
	"Zerg_Queens_Nest",
	"Zerg_Evolution_Chamber",
	"Zerg_Ultralisk_Cavern",
	"Zerg_Spire",
	"Zerg_Spawning_Pool",
	"Zerg_Creep_Colony",
	"Zerg_Spore_Colony",
	"Unused_Zerg1",
	"Zerg_Sunken_Colony",
	"Special_Overmind_With_Shell",
	"Special_Overmind",
	"Zerg_Extractor",
	"Special_Mature_Chrysalis",
	"Special_Cerebrate",
	"Special_Cerebrate_Daggoth",
	"Unused_Zerg2",
	"Protoss_Nexus",
	"Protoss_Robotics_Facility",
	"Protoss_Pylon",
	"Protoss_Assimilator",
	"Unused_Protoss1",
	"Protoss_Observatory",
	"Protoss_Gateway",
	"Unused_Protoss2",
	"Protoss_Photon_Cannon",
	"Protoss_Citadel_of_Adun",
	"Protoss_Cybernetics_Core",
	"Protoss_Templar_Archives",
	"Protoss_Forge",
	"Protoss_Stargate",
	"Special_Stasis_Cell_Prison",
	"Protoss_Fleet_Beacon",
	"Protoss_Arbiter_Tribunal",
	"Protoss_Robotics_Support_Bay",
	"Protoss_Shield_Battery",
	"Special_Khaydarin_Crystal_Form",
	"Special_Protoss_Temple",
	"Special_XelNaga_Temple",
	"Resource_Mineral_Field",
	"Resource_Mineral_Field_Type_2",
	"Resource_Mineral_Field_Type_3",
	"Unused_Cave",
	"Unused_Cave_In",
	"Unused_Cantina",
	"Unused_Mining_Platform",
	"Unused_Independant_Command_Center",
	"Special_Independant_Starport",
	"Unused_Independant_Jump_Gate",
	"Unused_Ruins",
	"Unused_Khaydarin_Crystal_Formation",
	"Resource_Vespene_Geyser",
	"Special_Warp_Gate",
	"Special_Psi_Disrupter",
	"Unused_Zerg_Marker",
	"Unused_Terran_Marker",
	"Unused_Protoss_Marker",
	"Special_Zerg_Beacon",
	"Special_Terran_Beacon",
	"Special_Protoss_Beacon",
	"Special_Zerg_Flag_Beacon",
	"Special_Terran_Flag_Beacon",
	"Special_Protoss_Flag_Beacon",
	"Special_Power_Generator",
	"Special_Overmind_Cocoon",
	"Spell_Dark_Swarm",
	"Special_Floor_Missile_Trap",
	"Special_Floor_Hatch",
	"Special_Upper_Level_Door",
	"Special_Right_Upper_Level_Door",
	"Special_Pit_Door",
	"Special_Right_Pit_Door",
	"Special_Floor_Gun_Trap",
	"Special_Wall_Missile_Trap",
	"Special_Wall_Flame_Trap",
	"Special_Right_Wall_Missile_Trap",
	"Special_Right_Wall_Flame_Trap",
	"Special_Start_Location",
	"Powerup_Flag",
	"Powerup_Young_Chrysalis",
	"Powerup_Psi_Emitter",
	"Powerup_Data_Disk",
	"Powerup_Khaydarin_Crystal",
	"Powerup_Mineral_Cluster_Type_1",
	"Powerup_Mineral_Cluster_Type_2",
	"Powerup_Protoss_Gas_Orb_Type_1",
	"Powerup_Protoss_Gas_Orb_Type_2",
	"Powerup_Zerg_Gas_Sac_Type_1",
	"Powerup_Zerg_Gas_Sac_Type_2",
	"Powerup_Terran_Gas_Tank_Type_1",
	"Powerup_Terran_Gas_Tank_Type_2",
	"None",
	"AllUnits",
	"Men",
	"Buildings",
	"Factories",
	"Unknown",
}

func (u UnitType) String() string {
	if u >= 0 && int(u) < len(unitTypeNames) {
		return unitTypeNames[u]
	}
	return "Unknown"
}

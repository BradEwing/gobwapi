package bwapi

// Order represents a BWAPI unit order.
type Order int32

const (
	OrderDie                      Order = 0
	OrderStop                     Order = 1
	OrderGuard                    Order = 2
	OrderPlayerGuard              Order = 3
	OrderTurretGuard              Order = 4
	OrderBunkerGuard              Order = 5
	OrderMove                     Order = 6
	OrderReaverStop               Order = 7
	OrderAttack1                  Order = 8
	OrderAttack2                  Order = 9
	OrderAttackUnit               Order = 10
	OrderAttackFixedRange         Order = 11
	OrderAttackTile               Order = 12
	OrderHover                    Order = 13
	OrderAttackMove               Order = 14
	OrderInfestedCommandCenter    Order = 15
	OrderUnusedNothing            Order = 16
	OrderUnusedPowerup            Order = 17
	OrderTowerGuard               Order = 18
	OrderTowerAttack              Order = 19
	OrderVultureMine              Order = 20
	OrderStayInRange              Order = 21
	OrderTurretAttack             Order = 22
	OrderNothing                  Order = 23
	OrderUnused24                 Order = 24
	OrderDroneStartBuild          Order = 25
	OrderDroneBuild               Order = 26
	OrderCastInfestation          Order = 27
	OrderMoveToInfest             Order = 28
	OrderInfestingCommandCenter   Order = 29
	OrderPlaceBuilding            Order = 30
	OrderPlaceProtossBuilding     Order = 31
	OrderCreateProtossBuilding    Order = 32
	OrderConstructingBuilding     Order = 33
	OrderRepair                   Order = 34
	OrderMoveToRepair             Order = 35
	OrderPlaceAddon               Order = 36
	OrderBuildAddon               Order = 37
	OrderTrain                    Order = 38
	OrderRallyPointUnit           Order = 39
	OrderRallyPointTile           Order = 40
	OrderZergBirth                Order = 41
	OrderZergUnitMorph            Order = 42
	OrderZergBuildingMorph        Order = 43
	OrderIncompleteBuilding       Order = 44
	OrderIncompleteMorphing       Order = 45
	OrderBuildNydusExit           Order = 46
	OrderEnterNydusCanal          Order = 47
	OrderIncompleteWarping        Order = 48
	OrderFollow                   Order = 49
	OrderCarrier                  Order = 50
	OrderReaverCarrierMove        Order = 51
	OrderCarrierStop              Order = 52
	OrderCarrierAttack            Order = 53
	OrderCarrierMoveToAttack      Order = 54
	OrderCarrierIgnore2           Order = 55
	OrderCarrierFight             Order = 56
	OrderCarrierHoldPosition      Order = 57
	OrderReaver                   Order = 58
	OrderReaverAttack             Order = 59
	OrderReaverMoveToAttack       Order = 60
	OrderReaverFight              Order = 61
	OrderReaverHoldPosition       Order = 62
	OrderTrainFighter             Order = 63
	OrderInterceptorAttack        Order = 64
	OrderScarabAttack             Order = 65
	OrderRechargeShieldsUnit      Order = 66
	OrderRechargeShieldsBattery   Order = 67
	OrderShieldBattery            Order = 68
	OrderInterceptorReturn        Order = 69
	OrderDroneLand                Order = 70
	OrderBuildingLand             Order = 71
	OrderBuildingLiftOff          Order = 72
	OrderDroneLiftOff             Order = 73
	OrderLiftingOff               Order = 74
	OrderResearchTech             Order = 75
	OrderUpgrade                  Order = 76
	OrderLarva                    Order = 77
	OrderSpawningLarva            Order = 78
	OrderHarvest1                 Order = 79
	OrderHarvest2                 Order = 80
	OrderMoveToGas                Order = 81
	OrderWaitForGas               Order = 82
	OrderHarvestGas               Order = 83
	OrderReturnGas                Order = 84
	OrderMoveToMinerals           Order = 85
	OrderWaitForMinerals          Order = 86
	OrderMiningMinerals           Order = 87
	OrderHarvest3                 Order = 88
	OrderHarvest4                 Order = 89
	OrderReturnMinerals           Order = 90
	OrderInterrupted              Order = 91
	OrderEnterTransport           Order = 92
	OrderPickupIdle               Order = 93
	OrderPickupTransport          Order = 94
	OrderPickupBunker             Order = 95
	OrderPickup4                  Order = 96
	OrderPowerupIdle              Order = 97
	OrderSieging                  Order = 98
	OrderUnsieging                Order = 99
	OrderWatchTarget              Order = 100
	OrderInitCreepGrowth          Order = 101
	OrderSpreadCreep              Order = 102
	OrderStoppingCreepGrowth      Order = 103
	OrderGuardianAspect           Order = 104
	OrderArchonWarp               Order = 105
	OrderCompletingArchonSummon   Order = 106
	OrderHoldPosition             Order = 107
	OrderQueenHoldPosition        Order = 108
	OrderCloak                    Order = 109
	OrderDecloak                  Order = 110
	OrderUnload                   Order = 111
	OrderMoveUnload               Order = 112
	OrderFireYamatoGun            Order = 113
	OrderMoveToFireYamatoGun      Order = 114
	OrderCastLockdown             Order = 115
	OrderBurrowing                Order = 116
	OrderBurrowed                 Order = 117
	OrderUnburrowing              Order = 118
	OrderCastDarkSwarm            Order = 119
	OrderCastParasite             Order = 120
	OrderCastSpawnBroodlings      Order = 121
	OrderCastEMPShockwave         Order = 122
	OrderNukeWait                 Order = 123
	OrderNukeTrain                Order = 124
	OrderNukeLaunch               Order = 125
	OrderNukePaint                Order = 126
	OrderNukeUnit                 Order = 127
	OrderCastNuclearStrike        Order = 128
	OrderNukeTrack                Order = 129
	OrderInitializeArbiter        Order = 130
	OrderCloakNearbyUnits         Order = 131
	OrderPlaceMine                Order = 132
	OrderRightClickAction         Order = 133
	OrderSuicideUnit              Order = 134
	OrderSuicideLocation          Order = 135
	OrderSuicideHoldPosition      Order = 136
	OrderCastRecall               Order = 137
	OrderTeleport                 Order = 138
	OrderCastScannerSweep         Order = 139
	OrderScanner                  Order = 140
	OrderCastDefensiveMatrix      Order = 141
	OrderCastPsionicStorm         Order = 142
	OrderCastIrradiate            Order = 143
	OrderCastPlague               Order = 144
	OrderCastConsume              Order = 145
	OrderCastEnsnare              Order = 146
	OrderCastStasisField          Order = 147
	OrderCastHallucination        Order = 148
	OrderHallucination2           Order = 149
	OrderResetCollision           Order = 150
	OrderResetHarvestCollision    Order = 151
	OrderPatrol                   Order = 152
	OrderCTFCOPInit               Order = 153
	OrderCTFCOPStarted            Order = 154
	OrderCTFCOP2                  Order = 155
	OrderComputerAI               Order = 156
	OrderAtkMoveEP                Order = 157
	OrderHarassMove               Order = 158
	OrderAIPatrol                 Order = 159
	OrderGuardPost                Order = 160
	OrderRescuePassive            Order = 161
	OrderNeutral                  Order = 162
	OrderComputerReturn           Order = 163
	OrderInitializePsiProvider    Order = 164
	OrderSelfDestructing          Order = 165
	OrderCritter                  Order = 166
	OrderHiddenGun                Order = 167
	OrderOpenDoor                 Order = 168
	OrderCloseDoor                Order = 169
	OrderHideTrap                 Order = 170
	OrderRevealTrap               Order = 171
	OrderEnableDoodad             Order = 172
	OrderDisableDoodad            Order = 173
	OrderWarpIn                   Order = 174
	OrderMedic                    Order = 175
	OrderMedicHeal                Order = 176
	OrderHealMove                 Order = 177
	OrderMedicHoldPosition        Order = 178
	OrderMedicHealToIdle          Order = 179
	OrderCastRestoration          Order = 180
	OrderCastDisruptionWeb        Order = 181
	OrderCastMindControl          Order = 182
	OrderDarkArchonMeld           Order = 183
	OrderCastFeedback             Order = 184
	OrderCastOpticalFlare         Order = 185
	OrderCastMaelstrom            Order = 186
	OrderJunkYardDog              Order = 187
	OrderFatal                    Order = 188
	OrderNone                     Order = 189
	OrderUnknown                  Order = 190
)

var orderNames = [191]string{
	"Die",
	"Stop",
	"Guard",
	"PlayerGuard",
	"TurretGuard",
	"BunkerGuard",
	"Move",
	"ReaverStop",
	"Attack1",
	"Attack2",
	"AttackUnit",
	"AttackFixedRange",
	"AttackTile",
	"Hover",
	"AttackMove",
	"InfestedCommandCenter",
	"UnusedNothing",
	"UnusedPowerup",
	"TowerGuard",
	"TowerAttack",
	"VultureMine",
	"StayInRange",
	"TurretAttack",
	"Nothing",
	"Unused_24",
	"DroneStartBuild",
	"DroneBuild",
	"CastInfestation",
	"MoveToInfest",
	"InfestingCommandCenter",
	"PlaceBuilding",
	"PlaceProtossBuilding",
	"CreateProtossBuilding",
	"ConstructingBuilding",
	"Repair",
	"MoveToRepair",
	"PlaceAddon",
	"BuildAddon",
	"Train",
	"RallyPointUnit",
	"RallyPointTile",
	"ZergBirth",
	"ZergUnitMorph",
	"ZergBuildingMorph",
	"IncompleteBuilding",
	"IncompleteMorphing",
	"BuildNydusExit",
	"EnterNydusCanal",
	"IncompleteWarping",
	"Follow",
	"Carrier",
	"ReaverCarrierMove",
	"CarrierStop",
	"CarrierAttack",
	"CarrierMoveToAttack",
	"CarrierIgnore2",
	"CarrierFight",
	"CarrierHoldPosition",
	"Reaver",
	"ReaverAttack",
	"ReaverMoveToAttack",
	"ReaverFight",
	"ReaverHoldPosition",
	"TrainFighter",
	"InterceptorAttack",
	"ScarabAttack",
	"RechargeShieldsUnit",
	"RechargeShieldsBattery",
	"ShieldBattery",
	"InterceptorReturn",
	"DroneLand",
	"BuildingLand",
	"BuildingLiftOff",
	"DroneLiftOff",
	"LiftingOff",
	"ResearchTech",
	"Upgrade",
	"Larva",
	"SpawningLarva",
	"Harvest1",
	"Harvest2",
	"MoveToGas",
	"WaitForGas",
	"HarvestGas",
	"ReturnGas",
	"MoveToMinerals",
	"WaitForMinerals",
	"MiningMinerals",
	"Harvest3",
	"Harvest4",
	"ReturnMinerals",
	"Interrupted",
	"EnterTransport",
	"PickupIdle",
	"PickupTransport",
	"PickupBunker",
	"Pickup4",
	"PowerupIdle",
	"Sieging",
	"Unsieging",
	"WatchTarget",
	"InitCreepGrowth",
	"SpreadCreep",
	"StoppingCreepGrowth",
	"GuardianAspect",
	"ArchonWarp",
	"CompletingArchonSummon",
	"HoldPosition",
	"QueenHoldPosition",
	"Cloak",
	"Decloak",
	"Unload",
	"MoveUnload",
	"FireYamatoGun",
	"MoveToFireYamatoGun",
	"CastLockdown",
	"Burrowing",
	"Burrowed",
	"Unburrowing",
	"CastDarkSwarm",
	"CastParasite",
	"CastSpawnBroodlings",
	"CastEMPShockwave",
	"NukeWait",
	"NukeTrain",
	"NukeLaunch",
	"NukePaint",
	"NukeUnit",
	"CastNuclearStrike",
	"NukeTrack",
	"InitializeArbiter",
	"CloakNearbyUnits",
	"PlaceMine",
	"RightClickAction",
	"SuicideUnit",
	"SuicideLocation",
	"SuicideHoldPosition",
	"CastRecall",
	"Teleport",
	"CastScannerSweep",
	"Scanner",
	"CastDefensiveMatrix",
	"CastPsionicStorm",
	"CastIrradiate",
	"CastPlague",
	"CastConsume",
	"CastEnsnare",
	"CastStasisField",
	"CastHallucination",
	"Hallucination2",
	"ResetCollision",
	"ResetHarvestCollision",
	"Patrol",
	"CTFCOPInit",
	"CTFCOPStarted",
	"CTFCOP2",
	"ComputerAI",
	"AtkMoveEP",
	"HarassMove",
	"AIPatrol",
	"GuardPost",
	"RescuePassive",
	"Neutral",
	"ComputerReturn",
	"InitializePsiProvider",
	"SelfDestructing",
	"Critter",
	"HiddenGun",
	"OpenDoor",
	"CloseDoor",
	"HideTrap",
	"RevealTrap",
	"EnableDoodad",
	"DisableDoodad",
	"WarpIn",
	"Medic",
	"MedicHeal",
	"HealMove",
	"MedicHoldPosition",
	"MedicHealToIdle",
	"CastRestoration",
	"CastDisruptionWeb",
	"CastMindControl",
	"DarkArchonMeld",
	"CastFeedback",
	"CastOpticalFlare",
	"CastMaelstrom",
	"JunkYardDog",
	"Fatal",
	"None",
	"Unknown",
}

func (o Order) String() string {
	if o >= 0 && int(o) < len(orderNames) {
		return orderNames[o]
	}
	return "Unknown"
}

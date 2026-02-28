package bwapi

// CommandType represents a BWAPI game command type.
type CommandType int32

const (
	CommandTypeNone                     CommandType = 0
	CommandTypeSetScreenPosition        CommandType = 1
	CommandTypePingMinimap              CommandType = 2
	CommandTypeEnableFlag               CommandType = 3
	CommandTypePrintf                   CommandType = 4
	CommandTypeSendText                 CommandType = 5
	CommandTypePauseGame                CommandType = 6
	CommandTypeResumeGame               CommandType = 7
	CommandTypeLeaveGame                CommandType = 8
	CommandTypeRestartGame              CommandType = 9
	CommandTypeSetLocalSpeed            CommandType = 10
	CommandTypeSetLatCom                CommandType = 11
	CommandTypeSetGui                   CommandType = 12
	CommandTypeSetFrameSkip             CommandType = 13
	CommandTypeSetMap                   CommandType = 14
	CommandTypeSetAllies                CommandType = 15
	CommandTypeSetVision                CommandType = 16
	CommandTypeSetCommandOptimizerLevel CommandType = 17
	CommandTypeSetRevealAll             CommandType = 18
)

func (c CommandType) String() string {
	switch c {
	case CommandTypeNone:
		return "None"
	case CommandTypeSetScreenPosition:
		return "SetScreenPosition"
	case CommandTypePingMinimap:
		return "PingMinimap"
	case CommandTypeEnableFlag:
		return "EnableFlag"
	case CommandTypePrintf:
		return "Printf"
	case CommandTypeSendText:
		return "SendText"
	case CommandTypePauseGame:
		return "PauseGame"
	case CommandTypeResumeGame:
		return "ResumeGame"
	case CommandTypeLeaveGame:
		return "LeaveGame"
	case CommandTypeRestartGame:
		return "RestartGame"
	case CommandTypeSetLocalSpeed:
		return "SetLocalSpeed"
	case CommandTypeSetLatCom:
		return "SetLatCom"
	case CommandTypeSetGui:
		return "SetGui"
	case CommandTypeSetFrameSkip:
		return "SetFrameSkip"
	case CommandTypeSetMap:
		return "SetMap"
	case CommandTypeSetAllies:
		return "SetAllies"
	case CommandTypeSetVision:
		return "SetVision"
	case CommandTypeSetCommandOptimizerLevel:
		return "SetCommandOptimizerLevel"
	case CommandTypeSetRevealAll:
		return "SetRevealAll"
	default:
		return "Unknown"
	}
}

// UnitCommandType represents a BWAPI unit command type.
type UnitCommandType int32

const (
	UnitCommandTypeAttackMove         UnitCommandType = 0
	UnitCommandTypeAttackUnit         UnitCommandType = 1
	UnitCommandTypeBuild              UnitCommandType = 2
	UnitCommandTypeBuildAddon         UnitCommandType = 3
	UnitCommandTypeTrain              UnitCommandType = 4
	UnitCommandTypeMorph              UnitCommandType = 5
	UnitCommandTypeResearch           UnitCommandType = 6
	UnitCommandTypeUpgrade            UnitCommandType = 7
	UnitCommandTypeSetRallyPosition   UnitCommandType = 8
	UnitCommandTypeSetRallyUnit       UnitCommandType = 9
	UnitCommandTypeMove               UnitCommandType = 10
	UnitCommandTypePatrol             UnitCommandType = 11
	UnitCommandTypeHoldPosition       UnitCommandType = 12
	UnitCommandTypeStop               UnitCommandType = 13
	UnitCommandTypeFollow             UnitCommandType = 14
	UnitCommandTypeGather             UnitCommandType = 15
	UnitCommandTypeReturnCargo        UnitCommandType = 16
	UnitCommandTypeRepair             UnitCommandType = 17
	UnitCommandTypeBurrow             UnitCommandType = 18
	UnitCommandTypeUnburrow           UnitCommandType = 19
	UnitCommandTypeCloak              UnitCommandType = 20
	UnitCommandTypeDecloak            UnitCommandType = 21
	UnitCommandTypeSiege              UnitCommandType = 22
	UnitCommandTypeUnsiege            UnitCommandType = 23
	UnitCommandTypeLift               UnitCommandType = 24
	UnitCommandTypeLand               UnitCommandType = 25
	UnitCommandTypeLoad               UnitCommandType = 26
	UnitCommandTypeUnload             UnitCommandType = 27
	UnitCommandTypeUnloadAll          UnitCommandType = 28
	UnitCommandTypeUnloadAllPosition  UnitCommandType = 29
	UnitCommandTypeRightClickPosition UnitCommandType = 30
	UnitCommandTypeRightClickUnit     UnitCommandType = 31
	UnitCommandTypeHaltConstruction   UnitCommandType = 32
	UnitCommandTypeCancelConstruction UnitCommandType = 33
	UnitCommandTypeCancelAddon        UnitCommandType = 34
	UnitCommandTypeCancelTrain        UnitCommandType = 35
	UnitCommandTypeCancelTrainSlot    UnitCommandType = 36
	UnitCommandTypeCancelMorph        UnitCommandType = 37
	UnitCommandTypeCancelResearch     UnitCommandType = 38
	UnitCommandTypeCancelUpgrade      UnitCommandType = 39
	UnitCommandTypeUseTech            UnitCommandType = 40
	UnitCommandTypeUseTechPosition    UnitCommandType = 41
	UnitCommandTypeUseTechUnit        UnitCommandType = 42
	UnitCommandTypePlaceCOP           UnitCommandType = 43
	UnitCommandTypeNone               UnitCommandType = 44
	UnitCommandTypeUnknown            UnitCommandType = 45
)

func (u UnitCommandType) String() string {
	switch u {
	case UnitCommandTypeAttackMove:
		return "AttackMove"
	case UnitCommandTypeAttackUnit:
		return "AttackUnit"
	case UnitCommandTypeBuild:
		return "Build"
	case UnitCommandTypeBuildAddon:
		return "BuildAddon"
	case UnitCommandTypeTrain:
		return "Train"
	case UnitCommandTypeMorph:
		return "Morph"
	case UnitCommandTypeResearch:
		return "Research"
	case UnitCommandTypeUpgrade:
		return "Upgrade"
	case UnitCommandTypeSetRallyPosition:
		return "SetRallyPosition"
	case UnitCommandTypeSetRallyUnit:
		return "SetRallyUnit"
	case UnitCommandTypeMove:
		return "Move"
	case UnitCommandTypePatrol:
		return "Patrol"
	case UnitCommandTypeHoldPosition:
		return "HoldPosition"
	case UnitCommandTypeStop:
		return "Stop"
	case UnitCommandTypeFollow:
		return "Follow"
	case UnitCommandTypeGather:
		return "Gather"
	case UnitCommandTypeReturnCargo:
		return "ReturnCargo"
	case UnitCommandTypeRepair:
		return "Repair"
	case UnitCommandTypeBurrow:
		return "Burrow"
	case UnitCommandTypeUnburrow:
		return "Unburrow"
	case UnitCommandTypeCloak:
		return "Cloak"
	case UnitCommandTypeDecloak:
		return "Decloak"
	case UnitCommandTypeSiege:
		return "Siege"
	case UnitCommandTypeUnsiege:
		return "Unsiege"
	case UnitCommandTypeLift:
		return "Lift"
	case UnitCommandTypeLand:
		return "Land"
	case UnitCommandTypeLoad:
		return "Load"
	case UnitCommandTypeUnload:
		return "Unload"
	case UnitCommandTypeUnloadAll:
		return "UnloadAll"
	case UnitCommandTypeUnloadAllPosition:
		return "UnloadAllPosition"
	case UnitCommandTypeRightClickPosition:
		return "RightClickPosition"
	case UnitCommandTypeRightClickUnit:
		return "RightClickUnit"
	case UnitCommandTypeHaltConstruction:
		return "HaltConstruction"
	case UnitCommandTypeCancelConstruction:
		return "CancelConstruction"
	case UnitCommandTypeCancelAddon:
		return "CancelAddon"
	case UnitCommandTypeCancelTrain:
		return "CancelTrain"
	case UnitCommandTypeCancelTrainSlot:
		return "CancelTrainSlot"
	case UnitCommandTypeCancelMorph:
		return "CancelMorph"
	case UnitCommandTypeCancelResearch:
		return "CancelResearch"
	case UnitCommandTypeCancelUpgrade:
		return "CancelUpgrade"
	case UnitCommandTypeUseTech:
		return "UseTech"
	case UnitCommandTypeUseTechPosition:
		return "UseTechPosition"
	case UnitCommandTypeUseTechUnit:
		return "UseTechUnit"
	case UnitCommandTypePlaceCOP:
		return "PlaceCOP"
	case UnitCommandTypeNone:
		return "None"
	case UnitCommandTypeUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

package bwapi

// PlayerType represents a BWAPI player type.
type PlayerType int32

const (
	PlayerTypeNone            PlayerType = 0
	PlayerTypeComputer        PlayerType = 1
	PlayerTypePlayer          PlayerType = 2
	PlayerTypeRescuePassive   PlayerType = 3
	PlayerTypeRescueActive    PlayerType = 4
	PlayerTypeEitherPreferred PlayerType = 5
	PlayerTypeHuman           PlayerType = 6
	PlayerTypeNeutral         PlayerType = 7
	PlayerTypeClosed          PlayerType = 8
	PlayerTypeObserver        PlayerType = 9
	PlayerTypePlayerLeft      PlayerType = 10
	PlayerTypeComputerLeft    PlayerType = 11
	PlayerTypeUnknown         PlayerType = 12
)

func (pt PlayerType) String() string {
	switch pt {
	case PlayerTypeNone:
		return "None"
	case PlayerTypeComputer:
		return "Computer"
	case PlayerTypePlayer:
		return "Player"
	case PlayerTypeRescuePassive:
		return "RescuePassive"
	case PlayerTypeRescueActive:
		return "RescueActive"
	case PlayerTypeEitherPreferred:
		return "EitherPreferred"
	case PlayerTypeHuman:
		return "Human"
	case PlayerTypeNeutral:
		return "Neutral"
	case PlayerTypeClosed:
		return "Closed"
	case PlayerTypeObserver:
		return "Observer"
	case PlayerTypePlayerLeft:
		return "PlayerLeft"
	case PlayerTypeComputerLeft:
		return "ComputerLeft"
	case PlayerTypeUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

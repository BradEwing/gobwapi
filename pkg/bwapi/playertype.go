package bwapi

// PlayerType represents a BWAPI player type.
type PlayerType int32

const (
	PlayerTypeNone            PlayerType = 0
	PlayerTypeComputer        PlayerType = 1
	PlayerTypePlayer          PlayerType = 2
	PlayerTypeRescuePassive   PlayerType = 3
	PlayerTypeRescueActive    PlayerType = 4 // Unused
	PlayerTypeEitherPreferred PlayerType = 5
	PlayerTypeHuman           PlayerType = 6
	PlayerTypeNeutral         PlayerType = 7
	PlayerTypeClosed          PlayerType = 8
	PlayerTypeObserver        PlayerType = 9
	PlayerTypePlayerLeft      PlayerType = 10
	PlayerTypeComputerLeft    PlayerType = 11
	PlayerTypeUnknown         PlayerType = 12
)

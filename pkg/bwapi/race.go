package bwapi

// Race represents a StarCraft race.
type Race int32

const (
	RaceZerg    Race = 0
	RaceTerran  Race = 1
	RaceProtoss Race = 2
	RaceOther   Race = 3
	RaceUnused  Race = 4
	RaceSelect  Race = 5
	RaceRandom  Race = 6
	RaceNone    Race = 7
	RaceUnknown Race = 8
)

func (r Race) String() string {
	switch r {
	case RaceZerg:
		return "Zerg"
	case RaceTerran:
		return "Terran"
	case RaceProtoss:
		return "Protoss"
	case RaceOther:
		return "Other"
	case RaceUnused:
		return "Unused"
	case RaceSelect:
		return "Select"
	case RaceRandom:
		return "Random"
	case RaceNone:
		return "None"
	case RaceUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

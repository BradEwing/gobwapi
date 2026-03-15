package bwem

import "github.com/bradewing/gobwapi/pkg/bwapi"

func isMineralField(ut bwapi.UnitType) bool {
	return ut == bwapi.UnitTypeResourceMineralField ||
		ut == bwapi.UnitTypeResourceMineralFieldType2 ||
		ut == bwapi.UnitTypeResourceMineralFieldType3
}

func isGeyser(ut bwapi.UnitType) bool {
	return ut == bwapi.UnitTypeResourceVespeneGeyser
}

func isSpecialBuilding(ut bwapi.UnitType) bool {
	switch ut {
	case bwapi.UnitTypeSpecialIndependantStarport:
		return true
	}
	// Check if the unit is a non-resource, non-critter neutral building.
	// Critters are mobile units, not buildings — they don't block terrain.
	return false
}

// unitTypeTileSize returns the tile dimensions (width, height) for unit types
// relevant to BWEM terrain analysis.
func unitTypeTileSize(ut bwapi.UnitType) (w, h int) {
	switch ut {
	case bwapi.UnitTypeResourceMineralField,
		bwapi.UnitTypeResourceMineralFieldType2,
		bwapi.UnitTypeResourceMineralFieldType3:
		return 2, 1
	case bwapi.UnitTypeResourceVespeneGeyser:
		return 4, 2
	case bwapi.UnitTypeSpecialIndependantStarport:
		return 4, 3
	default:
		return 4, 3
	}
}

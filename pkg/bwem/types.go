package bwem

import "github.com/bradewing/gobwapi/pkg/bwapi"

// AreaId identifies an Area.
// Positive = valid area, negative = undersized area, 0 = unassigned.
type AreaId int16

// Altitude represents pixel distance to the nearest sea boundary.
type Altitude int16

// GroupId identifies a connected component of mutually accessible areas.
type GroupId int16

// MiniTile represents a single walk-tile (8x8 pixels).
type MiniTile struct {
	Walkable bool
	AreaID   AreaId
	Alt      Altitude
	Sea      bool
	Lake     bool
}

// Tile represents a build-tile (32x32 pixels, contains 4x4 MiniTiles).
type Tile struct {
	Buildable    bool
	GroundHeight int8
	Doodad       bool
	AreaID       AreaId
	NeutralIdx   int
	MinAltitude  Altitude
}

// Neutral represents a neutral unit on the map (mineral, geyser, or static building).
type Neutral struct {
	Unit     *bwapi.Unit
	UnitType bwapi.UnitType
	Pos      bwapi.Position
	TilePos  bwapi.TilePosition
	TileW    int
	TileH    int
	Blocking bool
}

// Mineral is a mineral patch tracked by BWEM.
type Mineral struct {
	NeutralIdx int
	Resources  int
	BaseIdx    int
}

// Geyser is a vespene geyser tracked by BWEM.
type Geyser struct {
	NeutralIdx int
	Resources  int
	BaseIdx    int
}

// StaticBuilding is an unbuildable neutral building on the map.
type StaticBuilding struct {
	NeutralIdx int
}

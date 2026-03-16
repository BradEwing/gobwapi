package bwem

import "github.com/bradewing/gobwapi/pkg/bwapi"

// Area represents a connected region of walkable terrain.
type Area struct {
	ID          AreaId
	GroupID     GroupId
	Top         bwapi.WalkPosition
	TopAltitude Altitude

	MiniTileCount       int
	TileCount           int
	BuildableTiles      int
	HighGroundTiles     int
	VeryHighGroundTiles int

	ChokePointIdxs []int
	NeighborIDs    []AreaId

	BaseIdxs    []int
	MineralIdxs []int
	GeyserIdxs  []int

	TopLeft     bwapi.TilePosition
	BottomRight bwapi.TilePosition
}

// ChokePoints returns the chokepoints bordering this area.
func (a *Area) ChokePoints(m *Map) []*ChokePoint {
	cps := make([]*ChokePoint, len(a.ChokePointIdxs))
	for i, idx := range a.ChokePointIdxs {
		cps[i] = &m.chokePoints[idx]
	}
	return cps
}

// Bases returns the bases within this area.
func (a *Area) Bases(m *Map) []*Base {
	bases := make([]*Base, len(a.BaseIdxs))
	for i, idx := range a.BaseIdxs {
		bases[i] = &m.bases[idx]
	}
	return bases
}

// Minerals returns the minerals within this area.
func (a *Area) Minerals(m *Map) []*Mineral {
	mins := make([]*Mineral, len(a.MineralIdxs))
	for i, idx := range a.MineralIdxs {
		mins[i] = &m.minerals[idx]
	}
	return mins
}

// Geysers returns the geysers within this area.
func (a *Area) Geysers(m *Map) []*Geyser {
	gs := make([]*Geyser, len(a.GeyserIdxs))
	for i, idx := range a.GeyserIdxs {
		gs[i] = &m.geysers[idx]
	}
	return gs
}

// IsAccessibleFrom returns true if this area is ground-connected to other.
func (a *Area) IsAccessibleFrom(other *Area) bool {
	return a.GroupID == other.GroupID && a.GroupID > 0
}

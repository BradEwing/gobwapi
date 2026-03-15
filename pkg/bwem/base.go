package bwem

import "github.com/bradewing/gobwapi/pkg/bwapi"

// Base represents a potential base location (Command Center placement).
type Base struct {
	Index           int
	Location        bwapi.TilePosition // top-left of CC footprint
	Center          bwapi.Position     // pixel center of CC
	AreaID          AreaId
	IsStartLocation bool
	MineralIdxs     []int // indices into Map.minerals
	GeyserIdxs      []int // indices into Map.geysers
}

// Area returns this base's area.
func (b *Base) Area(m *Map) *Area {
	return m.Area(b.AreaID)
}

// Minerals returns the minerals assigned to this base.
func (b *Base) Minerals(m *Map) []*Mineral {
	mins := make([]*Mineral, len(b.MineralIdxs))
	for i, idx := range b.MineralIdxs {
		mins[i] = &m.minerals[idx]
	}
	return mins
}

// Geysers returns the geysers assigned to this base.
func (b *Base) Geysers(m *Map) []*Geyser {
	gs := make([]*Geyser, len(b.GeyserIdxs))
	for i, idx := range b.GeyserIdxs {
		gs[i] = &m.geysers[idx]
	}
	return gs
}

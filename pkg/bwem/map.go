package bwem

import (
	"github.com/bradewing/gobwapi/pkg/bwapi"
)

// Map is the result of BWEM terrain analysis.
type Map struct {
	tileWidth  int
	tileHeight int
	walkWidth  int
	walkHeight int

	startLocations []bwapi.TilePosition

	miniTiles []MiniTile
	tiles     []Tile

	areas       []Area
	chokePoints []ChokePoint
	bases       []Base
	neutrals    []Neutral
	minerals    []Mineral
	geysers     []Geyser
	statics     []StaticBuilding

	gr graph
}

// Analyze runs the full BWEM terrain analysis and returns the completed Map.
func Analyze(game *bwapi.Game) *Map {
	m := &Map{}
	m.analyze(game)
	return m
}

func (m *Map) TileWidth() int  { return m.tileWidth }
func (m *Map) TileHeight() int { return m.tileHeight }
func (m *Map) WalkWidth() int  { return m.walkWidth }
func (m *Map) WalkHeight() int { return m.walkHeight }

// StartLocations returns the map's player start positions.
func (m *Map) StartLocations() []bwapi.TilePosition {
	return m.startLocations
}

// GetMiniTile returns the MiniTile at a WalkPosition.
func (m *Map) GetMiniTile(wp bwapi.WalkPosition) *MiniTile {
	if !m.validWalk(wp) {
		return nil
	}
	return &m.miniTiles[m.miniTileIndex(wp)]
}

// GetTile returns the Tile at a TilePosition.
func (m *Map) GetTile(tp bwapi.TilePosition) *Tile {
	if !m.validTile(tp) {
		return nil
	}
	return &m.tiles[m.tileIndex(tp)]
}

// Areas returns all analyzed areas.
func (m *Map) Areas() []Area {
	return m.areas
}

// Area returns the area with the given ID, or nil if invalid.
func (m *Map) Area(id AreaId) *Area {
	if id <= 0 || int(id) > len(m.areas) {
		return nil
	}
	return &m.areas[id-1]
}

// AreaAt returns the area at a TilePosition, or nil.
func (m *Map) AreaAt(tp bwapi.TilePosition) *Area {
	t := m.GetTile(tp)
	if t == nil {
		return nil
	}
	return m.Area(t.AreaID)
}

// AreaAtWalk returns the area at a WalkPosition, or nil.
func (m *Map) AreaAtWalk(wp bwapi.WalkPosition) *Area {
	mt := m.GetMiniTile(wp)
	if mt == nil {
		return nil
	}
	return m.Area(mt.AreaID)
}

// NearestArea finds the nearest valid area to a WalkPosition via BFS.
func (m *Map) NearestArea(wp bwapi.WalkPosition) *Area {
	if a := m.AreaAtWalk(wp); a != nil {
		return a
	}
	visited := make([]bool, len(m.miniTiles))
	queue := []bwapi.WalkPosition{wp}
	visited[m.miniTileIndex(wp)] = true
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, n := range m.walkNeighbors8(cur) {
			ni := m.miniTileIndex(n)
			if visited[ni] {
				continue
			}
			visited[ni] = true
			if m.miniTiles[ni].AreaID > 0 {
				return m.Area(m.miniTiles[ni].AreaID)
			}
			queue = append(queue, n)
		}
	}
	return nil
}

// ChokePoints returns all chokepoints.
func (m *Map) ChokePoints() []ChokePoint {
	return m.chokePoints
}

// Bases returns all base locations.
func (m *Map) Bases() []Base {
	return m.bases
}

// StartingBases returns only bases that are player start locations.
func (m *Map) StartingBases() []*Base {
	var result []*Base
	for i := range m.bases {
		if m.bases[i].IsStartLocation {
			result = append(result, &m.bases[i])
		}
	}
	return result
}

func (m *Map) Neutrals() []Neutral { return m.neutrals }
func (m *Map) Minerals() []Mineral { return m.minerals }
func (m *Map) Geysers() []Geyser   { return m.geysers }

// GetPath returns the shortest path between two positions as a sequence of
// chokepoints, along with the total ground distance in pixels.
// Returns nil, -1 if no path exists.
func (m *Map) GetPath(from, to bwapi.Position) ([]*ChokePoint, int) {
	fromArea := m.NearestArea(from.ToWalkPosition())
	toArea := m.NearestArea(to.ToWalkPosition())
	if fromArea == nil || toArea == nil {
		return nil, -1
	}
	if fromArea.ID == toArea.ID {
		return nil, walkDist(from.ToWalkPosition(), to.ToWalkPosition())
	}
	if fromArea.GroupID != toArea.GroupID {
		return nil, -1
	}

	bestDist := -1
	bestFrom := -1
	bestTo := -1
	for _, fi := range fromArea.ChokePointIdxs {
		for _, ti := range toArea.ChokePointIdxs {
			d := m.gr.getDistance(fi, ti)
			if d >= 0 && (bestDist < 0 || d < bestDist) {
				bestDist = d
				bestFrom = fi
				bestTo = ti
			}
		}
	}
	if bestDist < 0 {
		return nil, -1
	}

	cpIdxs := m.gr.getPath(bestFrom, bestTo)
	result := make([]*ChokePoint, len(cpIdxs))
	for i, idx := range cpIdxs {
		result[i] = &m.chokePoints[idx]
	}
	return result, bestDist
}

// OnMineralDestroyed should be called when a mineral patch is destroyed.
func (m *Map) OnMineralDestroyed(unit *bwapi.Unit) {
	for i := range m.minerals {
		n := &m.neutrals[m.minerals[i].NeutralIdx]
		if n.Unit == unit {
			if m.minerals[i].BaseIdx >= 0 {
				base := &m.bases[m.minerals[i].BaseIdx]
				base.MineralIdxs = removeInt(base.MineralIdxs, i)
			}
			if n.Blocking {
				m.unblockNeutral(m.minerals[i].NeutralIdx)
			}
			return
		}
	}
}

// OnStaticBuildingDestroyed should be called when a neutral static building is destroyed.
func (m *Map) OnStaticBuildingDestroyed(unit *bwapi.Unit) {
	for i := range m.statics {
		n := &m.neutrals[m.statics[i].NeutralIdx]
		if n.Unit == unit {
			if n.Blocking {
				m.unblockNeutral(m.statics[i].NeutralIdx)
			}
			return
		}
	}
}

func (m *Map) unblockNeutral(neutralIdx int) {
	for i := range m.chokePoints {
		if m.chokePoints[i].NeutralIdx == neutralIdx {
			m.chokePoints[i].Blocked = false
		}
	}
}

func removeInt(s []int, val int) []int {
	for i, v := range s {
		if v == val {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

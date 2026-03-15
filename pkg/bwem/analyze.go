package bwem

import (
	"container/heap"
	"math"
	"sort"

	"github.com/bradewing/gobwapi/pkg/bwapi"
)

// analyze runs the full 13-step BWEM pipeline.
func (m *Map) analyze(game *bwapi.Game) {
	m.initializeTerrainData(game)      // 1
	m.markUnwalkableMiniTiles(game)    // 2
	m.markBuildableTilesAndGH(game)    // 3
	m.decideSeasOrLakes()              // 4
	m.initializeNeutralData(game)      // 5
	m.computeAltitude()                // 6
	m.processBlockingNeutrals()        // 7
	temps, fronts := m.computeTempAreas() // 8
	m.computeAreas(temps)              // 9
	m.createChokePoints(fronts)        // 10
	m.gr.computeChokePointDistanceMatrix(m) // 11
	m.gr.computeGroupIds(m)            // 12
	m.createBases()                    // 13
}

// --- Step 1: Initialize terrain data ---

func (m *Map) initializeTerrainData(game *bwapi.Game) {
	m.tileWidth = game.MapWidth()
	m.tileHeight = game.MapHeight()
	m.walkWidth = m.tileWidth * 4
	m.walkHeight = m.tileHeight * 4
	m.startLocations = game.GetStartLocations()
	m.miniTiles = make([]MiniTile, m.walkWidth*m.walkHeight)
	m.tiles = make([]Tile, m.tileWidth*m.tileHeight)
	for i := range m.tiles {
		m.tiles[i].NeutralIdx = -1
	}
}

// --- Step 2: Mark unwalkable MiniTiles ---

func (m *Map) markUnwalkableMiniTiles(game *bwapi.Game) {
	// First pass: read walkability from BWAPI.
	for y := 0; y < m.walkHeight; y++ {
		for x := 0; x < m.walkWidth; x++ {
			m.miniTiles[y*m.walkWidth+x].Walkable = game.IsWalkable(x, y)
		}
	}

	// Second pass: spread unwalkability to 8 neighbors.
	// Use a snapshot to avoid cascade.
	snapshot := make([]bool, len(m.miniTiles))
	for i, mt := range m.miniTiles {
		snapshot[i] = mt.Walkable
	}

	for y := 0; y < m.walkHeight; y++ {
		for x := 0; x < m.walkWidth; x++ {
			if snapshot[y*m.walkWidth+x] {
				continue // walkable, skip
			}
			// Mark all 8 neighbors as unwalkable.
			wp := bwapi.WalkPosition{X: int32(x), Y: int32(y)}
			for _, n := range m.walkNeighbors8(wp) {
				m.miniTiles[m.miniTileIndex(n)].Walkable = false
			}
		}
	}
}

// --- Step 3: Mark buildable tiles and ground height ---

func (m *Map) markBuildableTilesAndGH(game *bwapi.Game) {
	for ty := 0; ty < m.tileHeight; ty++ {
		for tx := 0; tx < m.tileWidth; tx++ {
			ti := ty*m.tileWidth + tx
			t := &m.tiles[ti]
			t.Buildable = game.IsBuildable(tx, ty)

			raw := game.GetGroundHeight(tx, ty)
			t.GroundHeight = int8(raw / 2)
			t.Doodad = raw%2 == 1

			// If buildable, force all 4x4 sub-MiniTiles walkable.
			if t.Buildable {
				for dy := 0; dy < 4; dy++ {
					for dx := 0; dx < 4; dx++ {
						wx := tx*4 + dx
						wy := ty*4 + dy
						if wx < m.walkWidth && wy < m.walkHeight {
							m.miniTiles[wy*m.walkWidth+wx].Walkable = true
						}
					}
				}
			}
		}
	}
}

// --- Step 4: Decide seas or lakes ---

func (m *Map) decideSeasOrLakes() {
	visited := make([]bool, len(m.miniTiles))

	for y := 0; y < m.walkHeight; y++ {
		for x := 0; x < m.walkWidth; x++ {
			wp := bwapi.WalkPosition{X: int32(x), Y: int32(y)}
			idx := m.miniTileIndex(wp)
			if visited[idx] || m.miniTiles[idx].Walkable {
				continue
			}

			// BFS over connected unwalkable MiniTiles.
			component := m.bfsFloodFill(wp, visited, func(_ bwapi.WalkPosition, mt *MiniTile) bool {
				return !mt.Walkable
			})

			// Compute bounding box.
			minX, maxX := int32(math.MaxInt32), int32(math.MinInt32)
			minY, maxY := int32(math.MaxInt32), int32(math.MinInt32)
			for _, p := range component {
				if p.X < minX {
					minX = p.X
				}
				if p.X > maxX {
					maxX = p.X
				}
				if p.Y < minY {
					minY = p.Y
				}
				if p.Y > maxY {
					maxY = p.Y
				}
			}
			w := maxX - minX + 1
			h := maxY - minY + 1

			// Classify: sea if large or wide, lake otherwise.
			isSea := len(component) > lakeMaxMiniTiles ||
				int(w) > lakeMaxWidthMiniTiles ||
				int(h) > lakeMaxWidthMiniTiles

			for _, p := range component {
				mt := &m.miniTiles[m.miniTileIndex(p)]
				if isSea {
					mt.Sea = true
					mt.Alt = 0
				} else {
					mt.Lake = true
				}
			}
		}
	}
}

// --- Step 5: Initialize neutral data ---

func (m *Map) initializeNeutralData(game *bwapi.Game) {
	neutralPlayer := game.Neutral()
	if neutralPlayer == nil {
		return
	}
	neutralIdx := neutralPlayer.Index()

	for _, u := range game.GetAllUnits() {
		p := u.GetPlayer()
		if p == nil || p.Index() != neutralIdx {
			continue
		}

		ut := u.GetType()
		tw, th := unitTypeTileSize(ut)
		tp := u.GetTilePosition()

		neutral := Neutral{
			Unit:     u,
			UnitType: ut,
			Pos:      u.GetPosition(),
			TilePos:  tp,
			TileW:    tw,
			TileH:    th,
		}
		nIdx := len(m.neutrals)
		m.neutrals = append(m.neutrals, neutral)

		// Mark tile footprint.
		for dy := 0; dy < th; dy++ {
			for dx := 0; dx < tw; dx++ {
				ttp := bwapi.TilePosition{X: tp.X + int32(dx), Y: tp.Y + int32(dy)}
				if m.validTile(ttp) {
					t := &m.tiles[m.tileIndex(ttp)]
					if t.NeutralIdx < 0 {
						t.NeutralIdx = nIdx
					}
				}
			}
		}

		res := u.Resources()
		if isMineralField(ut) && res >= 40 {
			m.minerals = append(m.minerals, Mineral{
				NeutralIdx: nIdx,
				Resources:  res,
				BaseIdx:    -1,
			})
		} else if isGeyser(ut) && res >= 300 {
			m.geysers = append(m.geysers, Geyser{
				NeutralIdx: nIdx,
				Resources:  res,
				BaseIdx:    -1,
			})
		} else if isSpecialBuilding(ut) {
			m.statics = append(m.statics, StaticBuilding{
				NeutralIdx: nIdx,
			})
		}
	}
}

// --- Step 6: Compute altitude ---

func (m *Map) computeAltitude() {
	// Multi-source Dijkstra from all MiniTiles adjacent to Sea.
	pq := &priorityQueue{}
	heap.Init(pq)

	altSet := make([]bool, len(m.miniTiles)) // true if altitude already finalized

	// Seed: all walkable MiniTiles that are adjacent to a Sea MiniTile.
	for y := 0; y < m.walkHeight; y++ {
		for x := 0; x < m.walkWidth; x++ {
			wp := bwapi.WalkPosition{X: int32(x), Y: int32(y)}
			idx := m.miniTileIndex(wp)
			mt := &m.miniTiles[idx]
			if mt.Sea {
				mt.Alt = 0
				altSet[idx] = true
				continue
			}
			if mt.Lake {
				// Lakes keep altitude > 0 but don't participate in Dijkstra.
				mt.Alt = 1
				altSet[idx] = true
				continue
			}
			if !mt.Walkable {
				// Non-sea, non-lake unwalkable: treat like sea for altitude seeding.
				mt.Alt = 0
				altSet[idx] = true
				continue
			}
		}
	}

	// Seed from all walkable tiles adjacent to altitude-0 tiles.
	for y := 0; y < m.walkHeight; y++ {
		for x := 0; x < m.walkWidth; x++ {
			wp := bwapi.WalkPosition{X: int32(x), Y: int32(y)}
			idx := m.miniTileIndex(wp)
			if !m.miniTiles[idx].Walkable || altSet[idx] {
				continue
			}
			// Check if adjacent to any altitude-0 tile.
			for _, n := range m.walkNeighbors8(wp) {
				ni := m.miniTileIndex(n)
				if m.miniTiles[ni].Alt == 0 && altSet[ni] {
					// Distance to neighbor: 8 for cardinal, 11 for diagonal.
					dx := n.X - wp.X
					dy := n.Y - wp.Y
					d := 8
					if dx != 0 && dy != 0 {
						d = 11
					}
					heap.Push(pq, pqItem{wp: wp, dist: d})
					break
				}
			}
		}
	}

	// Dijkstra: propagate altitude inward.
	for pq.Len() > 0 {
		item := heap.Pop(pq).(pqItem)
		idx := m.miniTileIndex(item.wp)
		if altSet[idx] {
			continue
		}
		altSet[idx] = true
		m.miniTiles[idx].Alt = Altitude(item.dist)

		for _, n := range m.walkNeighbors8(item.wp) {
			ni := m.miniTileIndex(n)
			if altSet[ni] || !m.miniTiles[ni].Walkable {
				continue
			}
			dx := n.X - item.wp.X
			dy := n.Y - item.wp.Y
			cost := 8
			if dx != 0 && dy != 0 {
				cost = 11
			}
			heap.Push(pq, pqItem{wp: n, dist: item.dist + cost})
		}
	}
}

// --- Step 7: Process blocking neutrals ---

func (m *Map) processBlockingNeutrals() {
	for i := range m.neutrals {
		n := &m.neutrals[i]
		if isMineralField(n.UnitType) || isGeyser(n.UnitType) || isSpecialBuilding(n.UnitType) {
			doors := m.countDoors(n)
			if doors >= 2 {
				n.Blocking = true
			}
		}
	}
}

// countDoors counts distinct walkable "doors" around a neutral's tile footprint.
// A door is a walkable MiniTile on the border that leads to a sufficiently large area.
func (m *Map) countDoors(n *Neutral) int {
	// Collect all walkable MiniTiles on the 1-tile border around the neutral.
	visited := make(map[int]bool)
	var seeds []bwapi.WalkPosition

	for dy := -1; dy <= n.TileH; dy++ {
		for dx := -1; dx <= n.TileW; dx++ {
			// Only border tiles.
			if dy >= 0 && dy < n.TileH && dx >= 0 && dx < n.TileW {
				continue
			}
			tp := bwapi.TilePosition{X: n.TilePos.X + int32(dx), Y: n.TilePos.Y + int32(dy)}
			if !m.validTile(tp) {
				continue
			}
			// Check all 4x4 sub-MiniTiles.
			for sy := 0; sy < 4; sy++ {
				for sx := 0; sx < 4; sx++ {
					wp := bwapi.WalkPosition{
						X: tp.X*4 + int32(sx),
						Y: tp.Y*4 + int32(sy),
					}
					if !m.validWalk(wp) {
						continue
					}
					idx := m.miniTileIndex(wp)
					if m.miniTiles[idx].Walkable && !visited[idx] {
						visited[idx] = true
						seeds = append(seeds, wp)
					}
				}
			}
		}
	}

	if len(seeds) == 0 {
		return 0
	}

	// BFS from each unvisited seed to count distinct connected components
	// that reach a significant number of tiles.
	doorVisited := make([]bool, len(m.miniTiles))
	doors := 0
	for _, seed := range seeds {
		si := m.miniTileIndex(seed)
		if doorVisited[si] {
			continue
		}
		comp := m.bfsFloodFill(seed, doorVisited, func(_ bwapi.WalkPosition, mt *MiniTile) bool {
			return mt.Walkable
		})
		if len(comp) >= 10 {
			doors++
		}
	}
	return doors
}

// --- Step 8: Compute temporary areas ---

type tempArea struct {
	id          AreaId
	top         bwapi.WalkPosition
	topAltitude Altitude
	count       int
	mergedInto  AreaId // 0 if not merged
}

type frontier struct {
	wp    bwapi.WalkPosition
	area1 AreaId
	area2 AreaId
}

func (m *Map) computeTempAreas() ([]tempArea, []frontier) {
	// Collect all walkable MiniTiles and sort by altitude descending.
	type wpAlt struct {
		wp  bwapi.WalkPosition
		alt Altitude
	}
	var walkable []wpAlt
	for y := 0; y < m.walkHeight; y++ {
		for x := 0; x < m.walkWidth; x++ {
			idx := y*m.walkWidth + x
			if m.miniTiles[idx].Walkable {
				walkable = append(walkable, wpAlt{
					wp:  bwapi.WalkPosition{X: int32(x), Y: int32(y)},
					alt: m.miniTiles[idx].Alt,
				})
			}
		}
	}
	sort.Slice(walkable, func(i, j int) bool {
		return walkable[i].alt > walkable[j].alt
	})

	var temps []tempArea
	var fronts []frontier
	nextID := AreaId(1)

	// resolveID follows merge chains to find the current area ID.
	resolveID := func(id AreaId) AreaId {
		for {
			if id <= 0 || int(id) > len(temps) {
				return id
			}
			t := &temps[id-1]
			if t.mergedInto == 0 {
				return id
			}
			id = t.mergedInto
		}
	}

	for _, wa := range walkable {
		idx := m.miniTileIndex(wa.wp)

		// Find neighboring areas.
		neighborAreas := make(map[AreaId]bool)
		for _, n := range m.walkNeighbors8(wa.wp) {
			ni := m.miniTileIndex(n)
			aid := m.miniTiles[ni].AreaID
			if aid > 0 {
				resolved := resolveID(aid)
				neighborAreas[resolved] = true
			}
		}

		switch len(neighborAreas) {
		case 0:
			// No assigned neighbors: create a new tempArea.
			temps = append(temps, tempArea{
				id:          nextID,
				top:         wa.wp,
				topAltitude: wa.alt,
				count:       1,
			})
			m.miniTiles[idx].AreaID = nextID
			nextID++

		case 1:
			// Exactly one neighbor: extend that area.
			var areaID AreaId
			for id := range neighborAreas {
				areaID = id
			}
			temps[areaID-1].count++
			m.miniTiles[idx].AreaID = areaID

		default:
			// Multiple neighbors: pick the two largest, decide merge or frontier.
			ids := make([]AreaId, 0, len(neighborAreas))
			for id := range neighborAreas {
				ids = append(ids, id)
			}
			// Sort by size descending.
			sort.Slice(ids, func(i, j int) bool {
				return temps[ids[i]-1].count > temps[ids[j]-1].count
			})

			larger := ids[0]
			smaller := ids[1]
			lt := &temps[larger-1]
			st := &temps[smaller-1]

			shouldMerge := st.count < smallAreaMaxMiniTiles ||
				(lt.topAltitude > 0 && float64(st.topAltitude)/float64(lt.topAltitude) >= altitudeMergeRatio) ||
				m.containsStartLocation(st)

			if shouldMerge {
				// Merge smaller into larger.
				lt.count += st.count
				st.mergedInto = larger
				m.miniTiles[idx].AreaID = larger

				// Re-merge any remaining areas too.
				for _, id := range ids[2:] {
					t := &temps[id-1]
					lt.count += t.count
					t.mergedInto = larger
				}
			} else {
				// Record frontier between the two largest areas.
				fronts = append(fronts, frontier{
					wp:    wa.wp,
					area1: larger,
					area2: smaller,
				})
				// Assign to the larger area.
				lt.count++
				m.miniTiles[idx].AreaID = larger
			}
		}
	}

	// Resolve all merge chains in MiniTiles.
	for i := range m.miniTiles {
		if m.miniTiles[i].AreaID > 0 {
			m.miniTiles[i].AreaID = resolveID(m.miniTiles[i].AreaID)
		}
	}

	// Resolve merge chains in frontiers.
	for i := range fronts {
		fronts[i].area1 = resolveID(fronts[i].area1)
		fronts[i].area2 = resolveID(fronts[i].area2)
	}

	// Filter out frontiers where both sides merged into the same area.
	var filtered []frontier
	for _, f := range fronts {
		if f.area1 != f.area2 {
			filtered = append(filtered, f)
		}
	}

	return temps, filtered
}

// containsStartLocation checks if a tempArea contains any start location.
func (m *Map) containsStartLocation(ta *tempArea) bool {
	for _, sl := range m.startLocations {
		// Check center of the start location (CC footprint center).
		wp := bwapi.WalkPosition{X: sl.X*4 + 2, Y: sl.Y*4 + 2}
		if m.validWalk(wp) {
			idx := m.miniTileIndex(wp)
			if m.miniTiles[idx].AreaID == ta.id {
				return true
			}
		}
	}
	return false
}

// --- Step 9: Compute final areas ---

func (m *Map) computeAreas(temps []tempArea) {
	// Filter tempAreas that are large enough.
	// Build a mapping from old tempArea ID to new Area ID.
	oldToNew := make(map[AreaId]AreaId)
	newID := AreaId(1)

	for i := range temps {
		t := &temps[i]
		if t.mergedInto != 0 {
			continue // merged away
		}
		if t.count >= areaMinMiniTiles {
			oldToNew[t.id] = newID
			m.areas = append(m.areas, Area{
				ID:          newID,
				Top:         t.top,
				TopAltitude: t.topAltitude,
				TopLeft:     bwapi.TilePosition{X: int32(m.tileWidth), Y: int32(m.tileHeight)},
				BottomRight: bwapi.TilePosition{X: 0, Y: 0},
			})
			newID++
		}
	}

	// Reassign MiniTile AreaIDs: valid → new ID, undersized → negative.
	for i := range m.miniTiles {
		mt := &m.miniTiles[i]
		if mt.AreaID <= 0 {
			continue
		}
		if newAID, ok := oldToNew[mt.AreaID]; ok {
			mt.AreaID = newAID
		} else {
			mt.AreaID = -mt.AreaID // undersized
		}
	}

	// Compute area statistics from MiniTiles.
	for y := 0; y < m.walkHeight; y++ {
		for x := 0; x < m.walkWidth; x++ {
			mt := &m.miniTiles[y*m.walkWidth+x]
			if mt.AreaID <= 0 {
				continue
			}
			area := &m.areas[mt.AreaID-1]
			area.MiniTileCount++
		}
	}

	// Compute Tile-level area assignment and statistics.
	for ty := 0; ty < m.tileHeight; ty++ {
		for tx := 0; tx < m.tileWidth; tx++ {
			ti := ty*m.tileWidth + tx
			t := &m.tiles[ti]

			// Find most common area among 4x4 sub-MiniTiles and min altitude.
			counts := make(map[AreaId]int)
			minAlt := Altitude(math.MaxInt16)
			for dy := 0; dy < 4; dy++ {
				for dx := 0; dx < 4; dx++ {
					wx := tx*4 + dx
					wy := ty*4 + dy
					if wx < m.walkWidth && wy < m.walkHeight {
						mt := &m.miniTiles[wy*m.walkWidth+wx]
						if mt.AreaID > 0 {
							counts[mt.AreaID]++
						}
						if mt.Alt < minAlt {
							minAlt = mt.Alt
						}
					}
				}
			}
			t.MinAltitude = minAlt

			if len(counts) == 1 {
				for aid := range counts {
					t.AreaID = aid
				}
			} else if len(counts) > 1 {
				// Mixed: pick the most common, but mark as 0 if truly mixed.
				bestID := AreaId(0)
				bestCount := 0
				for aid, c := range counts {
					if c > bestCount {
						bestCount = c
						bestID = aid
					}
				}
				t.AreaID = bestID
			}

			// Update area bounding box and tile counts.
			if t.AreaID > 0 {
				area := &m.areas[t.AreaID-1]
				area.TileCount++
				tp := bwapi.TilePosition{X: int32(tx), Y: int32(ty)}
				if tp.X < area.TopLeft.X {
					area.TopLeft.X = tp.X
				}
				if tp.Y < area.TopLeft.Y {
					area.TopLeft.Y = tp.Y
				}
				if tp.X > area.BottomRight.X {
					area.BottomRight.X = tp.X
				}
				if tp.Y > area.BottomRight.Y {
					area.BottomRight.Y = tp.Y
				}
				if t.Buildable {
					area.BuildableTiles++
				}
				if t.GroundHeight == 1 {
					area.HighGroundTiles++
				} else if t.GroundHeight == 2 {
					area.VeryHighGroundTiles++
				}
			}
		}
	}

	// Assign minerals and geysers to areas.
	for i := range m.minerals {
		n := &m.neutrals[m.minerals[i].NeutralIdx]
		tp := n.TilePos
		if m.validTile(tp) {
			t := &m.tiles[m.tileIndex(tp)]
			if t.AreaID > 0 {
				area := &m.areas[t.AreaID-1]
				area.MineralIdxs = append(area.MineralIdxs, i)
			}
		}
	}
	for i := range m.geysers {
		n := &m.neutrals[m.geysers[i].NeutralIdx]
		tp := n.TilePos
		if m.validTile(tp) {
			t := &m.tiles[m.tileIndex(tp)]
			if t.AreaID > 0 {
				area := &m.areas[t.AreaID-1]
				area.GeyserIdxs = append(area.GeyserIdxs, i)
			}
		}
	}
}

// --- Step 10: Create chokepoints ---

func (m *Map) createChokePoints(fronts []frontier) {
	// Update frontier area IDs to use the new (post-step-9) IDs.
	// The MiniTiles already have new IDs, so re-read from them.
	for i := range fronts {
		idx := m.miniTileIndex(fronts[i].wp)
		// The frontier WalkPosition was assigned to the larger area in step 8.
		// We need the two areas on either side. Look at neighbors.
		neighborAreas := make(map[AreaId]bool)
		for _, n := range m.walkNeighbors8(fronts[i].wp) {
			ni := m.miniTileIndex(n)
			aid := m.miniTiles[ni].AreaID
			if aid > 0 {
				neighborAreas[aid] = true
			}
		}
		// Also include the frontier tile's own area.
		own := m.miniTiles[idx].AreaID
		if own > 0 {
			neighborAreas[own] = true
		}

		ids := make([]AreaId, 0, len(neighborAreas))
		for id := range neighborAreas {
			ids = append(ids, id)
		}
		if len(ids) >= 2 {
			sort.Slice(ids, func(a, b int) bool { return ids[a] < ids[b] })
			fronts[i].area1 = ids[0]
			fronts[i].area2 = ids[1]
		}
	}

	// Group frontiers by area pair.
	type areaPair struct{ a, b AreaId }
	groups := make(map[areaPair][]bwapi.WalkPosition)
	for _, f := range fronts {
		a, b := f.area1, f.area2
		if a <= 0 || b <= 0 || a == b {
			continue
		}
		if a > b {
			a, b = b, a
		}
		groups[areaPair{a, b}] = append(groups[areaPair{a, b}], f.wp)
	}

	// For each group, cluster frontier positions and create ChokePoints.
	for pair, positions := range groups {
		clusters := clusterPositions(positions, chokeClusterDistSq)

		for _, cluster := range clusters {
			cp := m.makeChokePoint(pair.a, pair.b, cluster)
			cp.Index = len(m.chokePoints)
			m.chokePoints = append(m.chokePoints, cp)

			// Register with areas.
			if a := m.Area(pair.a); a != nil {
				a.ChokePointIdxs = append(a.ChokePointIdxs, cp.Index)
			}
			if b := m.Area(pair.b); b != nil {
				b.ChokePointIdxs = append(b.ChokePointIdxs, cp.Index)
			}
		}
	}

	// Create pseudo-chokepoints for blocking neutrals.
	for i := range m.neutrals {
		n := &m.neutrals[i]
		if !n.Blocking {
			continue
		}

		// Find two areas on either side of the neutral.
		areasAround := m.areasAroundNeutral(n)
		if len(areasAround) < 2 {
			continue
		}

		ids := make([]AreaId, 0, len(areasAround))
		for id := range areasAround {
			ids = append(ids, id)
		}
		sort.Slice(ids, func(a, b int) bool { return ids[a] < ids[b] })

		// Create pseudo-CP between the two largest surrounding areas.
		centerWP := n.Pos.ToWalkPosition()
		cp := ChokePoint{
			Index:      len(m.chokePoints),
			AreaIDs:    [2]AreaId{ids[0], ids[1]},
			End1:       centerWP,
			Middle:     centerWP,
			End2:       centerWP,
			Blocked:    true,
			Pseudo:     true,
			NeutralIdx: i,
		}
		m.chokePoints = append(m.chokePoints, cp)

		if a := m.Area(ids[0]); a != nil {
			a.ChokePointIdxs = append(a.ChokePointIdxs, cp.Index)
		}
		if b := m.Area(ids[1]); b != nil {
			b.ChokePointIdxs = append(b.ChokePointIdxs, cp.Index)
		}
	}
}

// makeChokePoint creates a ChokePoint from a cluster of frontier positions.
func (m *Map) makeChokePoint(a, b AreaId, cluster []bwapi.WalkPosition) ChokePoint {
	// Middle = highest altitude point.
	middle := cluster[0]
	middleAlt := m.miniTiles[m.miniTileIndex(middle)].Alt
	for _, wp := range cluster[1:] {
		alt := m.miniTiles[m.miniTileIndex(wp)].Alt
		if alt > middleAlt {
			middleAlt = alt
			middle = wp
		}
	}

	// End1, End2 = the two points farthest from Middle.
	end1, end2 := cluster[0], cluster[0]
	maxDist1, maxDist2 := 0, 0
	for _, wp := range cluster {
		d := walkDistSq(wp, middle)
		if d > maxDist1 {
			maxDist2 = maxDist1
			end2 = end1
			maxDist1 = d
			end1 = wp
		} else if d > maxDist2 {
			maxDist2 = d
			end2 = wp
		}
	}

	return ChokePoint{
		AreaIDs:    [2]AreaId{a, b},
		End1:       end1,
		Middle:     middle,
		End2:       end2,
		Geometry:   cluster,
		NeutralIdx: -1,
	}
}

// clusterPositions groups WalkPositions into clusters where each position
// is within distSq of at least one other position in the cluster.
func clusterPositions(positions []bwapi.WalkPosition, distSq int) [][]bwapi.WalkPosition {
	n := len(positions)
	if n == 0 {
		return nil
	}

	// Union-Find.
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[ra] = rb
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if walkDistSq(positions[i], positions[j]) <= distSq {
				union(i, j)
			}
		}
	}

	groups := make(map[int][]bwapi.WalkPosition)
	for i, p := range positions {
		r := find(i)
		groups[r] = append(groups[r], p)
	}

	result := make([][]bwapi.WalkPosition, 0, len(groups))
	for _, g := range groups {
		result = append(result, g)
	}
	return result
}

// areasAroundNeutral finds the distinct areas bordering a neutral's tile footprint.
func (m *Map) areasAroundNeutral(n *Neutral) map[AreaId]bool {
	areas := make(map[AreaId]bool)
	for dy := -1; dy <= n.TileH; dy++ {
		for dx := -1; dx <= n.TileW; dx++ {
			if dy >= 0 && dy < n.TileH && dx >= 0 && dx < n.TileW {
				continue // inside, not border
			}
			tp := bwapi.TilePosition{X: n.TilePos.X + int32(dx), Y: n.TilePos.Y + int32(dy)}
			if !m.validTile(tp) {
				continue
			}
			t := &m.tiles[m.tileIndex(tp)]
			if t.AreaID > 0 {
				areas[t.AreaID] = true
			}
		}
	}
	return areas
}

// --- Step 13: Create bases ---

func (m *Map) createBases() {
	for areaIdx := range m.areas {
		area := &m.areas[areaIdx]
		m.createBasesForArea(area)
	}

	// Match start locations to nearest bases.
	for _, sl := range m.startLocations {
		bestDist := math.MaxFloat64
		bestIdx := -1
		slCenter := bwapi.Position{
			X: sl.X*32 + ccTileWidth*16,
			Y: sl.Y*32 + ccTileHeight*16,
		}
		for i := range m.bases {
			dx := float64(m.bases[i].Center.X - slCenter.X)
			dy := float64(m.bases[i].Center.Y - slCenter.Y)
			d := dx*dx + dy*dy
			if d < bestDist {
				bestDist = d
				bestIdx = i
			}
		}
		if bestIdx >= 0 && bestDist < float64(3*32*3*32) {
			m.bases[bestIdx].IsStartLocation = true
		}
	}
}

func (m *Map) createBasesForArea(area *Area) {
	// Collect unassigned minerals and geysers in this area.
	availMinerals := make([]int, 0)
	for _, mi := range area.MineralIdxs {
		if m.minerals[mi].BaseIdx < 0 {
			availMinerals = append(availMinerals, mi)
		}
	}
	availGeysers := make([]int, 0)
	for _, gi := range area.GeyserIdxs {
		if m.geysers[gi].BaseIdx < 0 {
			availGeysers = append(availGeysers, gi)
		}
	}

	for len(availMinerals) > 0 || len(availGeysers) > 0 {
		// Compute resource centroid.
		cx, cy := 0.0, 0.0
		count := 0
		for _, mi := range availMinerals {
			n := &m.neutrals[m.minerals[mi].NeutralIdx]
			cx += float64(n.Pos.X)
			cy += float64(n.Pos.Y)
			count++
		}
		for _, gi := range availGeysers {
			n := &m.neutrals[m.geysers[gi].NeutralIdx]
			cx += float64(n.Pos.X)
			cy += float64(n.Pos.Y)
			count++
		}
		if count == 0 {
			break
		}
		cx /= float64(count)
		cy /= float64(count)

		// Search for best CC placement.
		bestScore := 0.0
		bestTP := bwapi.TilePosition{X: -1, Y: -1}
		searchRadius := maxTilesBetweenCCAndRes + ccTileWidth

		centerTile := bwapi.TilePosition{
			X: int32(cx) / 32,
			Y: int32(cy) / 32,
		}

		for dy := -searchRadius; dy <= searchRadius; dy++ {
			for dx := -searchRadius; dx <= searchRadius; dx++ {
				tp := bwapi.TilePosition{
					X: centerTile.X + int32(dx),
					Y: centerTile.Y + int32(dy),
				}
				if !m.canPlaceCC(tp, area.ID) {
					continue
				}

				// Score: sum of 1/distance to nearby resources.
				score := 0.0
				ccCenter := bwapi.Position{
					X: tp.X*32 + ccTileWidth*16,
					Y: tp.Y*32 + ccTileHeight*16,
				}

				for _, mi := range availMinerals {
					n := &m.neutrals[m.minerals[mi].NeutralIdx]
					d := pixelDist(ccCenter, n.Pos)
					if d > 0 && d < float64(maxTilesBetweenCCAndRes*32) {
						score += 1.0 / d
					}
				}
				for _, gi := range availGeysers {
					n := &m.neutrals[m.geysers[gi].NeutralIdx]
					d := pixelDist(ccCenter, n.Pos)
					if d > 0 && d < float64(maxTilesBetweenCCAndRes*32) {
						score += 3.0 / d // geysers weighted 3x
					}
				}

				if score > bestScore {
					bestScore = score
					bestTP = tp
				}
			}
		}

		if bestTP.X < 0 {
			break // no valid placement found
		}

		// Create base.
		base := Base{
			Index:    len(m.bases),
			Location: bestTP,
			Center: bwapi.Position{
				X: bestTP.X*32 + ccTileWidth*16,
				Y: bestTP.Y*32 + ccTileHeight*16,
			},
			AreaID: area.ID,
		}

		// Assign nearby resources to this base.
		var remainMinerals []int
		for _, mi := range availMinerals {
			n := &m.neutrals[m.minerals[mi].NeutralIdx]
			d := tileDist(bestTP, n.TilePos)
			if d <= maxTilesBetweenCCAndRes {
				m.minerals[mi].BaseIdx = base.Index
				base.MineralIdxs = append(base.MineralIdxs, mi)
			} else {
				remainMinerals = append(remainMinerals, mi)
			}
		}
		availMinerals = remainMinerals

		var remainGeysers []int
		for _, gi := range availGeysers {
			n := &m.neutrals[m.geysers[gi].NeutralIdx]
			d := tileDist(bestTP, n.TilePos)
			if d <= maxTilesBetweenCCAndRes {
				m.geysers[gi].BaseIdx = base.Index
				base.GeyserIdxs = append(base.GeyserIdxs, gi)
			} else {
				remainGeysers = append(remainGeysers, gi)
			}
		}
		availGeysers = remainGeysers

		// Only add base if it has resources.
		if len(base.MineralIdxs) > 0 || len(base.GeyserIdxs) > 0 {
			area.BaseIdxs = append(area.BaseIdxs, base.Index)
			m.bases = append(m.bases, base)
		} else {
			break
		}
	}
}

// canPlaceCC checks if a CC can be placed at a TilePosition within the given area.
func (m *Map) canPlaceCC(tp bwapi.TilePosition, areaID AreaId) bool {
	for dy := 0; dy < ccTileHeight; dy++ {
		for dx := 0; dx < ccTileWidth; dx++ {
			check := bwapi.TilePosition{X: tp.X + int32(dx), Y: tp.Y + int32(dy)}
			if !m.validTile(check) {
				return false
			}
			t := &m.tiles[m.tileIndex(check)]
			if !t.Buildable {
				return false
			}
			if t.NeutralIdx >= 0 {
				return false // occupied by a neutral
			}
		}
	}
	return true
}

func pixelDist(a, b bwapi.Position) float64 {
	dx := float64(a.X - b.X)
	dy := float64(a.Y - b.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

func tileDist(a, b bwapi.TilePosition) int {
	dx := int(a.X - b.X)
	dy := int(a.Y - b.Y)
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	if dx > dy {
		return dx
	}
	return dy
}

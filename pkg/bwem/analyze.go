package bwem

import (
	"container/heap"
	"math"
	"sort"

	"github.com/bradewing/gobwapi/pkg/bwapi"
)

func (m *Map) analyze(game *bwapi.Game) {
	m.initializeTerrainData(game)
	m.markUnwalkableMiniTiles(game)
	m.markBuildableTilesAndGH(game)
	m.decideSeasOrLakes()
	m.initializeNeutralData(game)
	m.computeAltitude()
	m.processBlockingNeutrals()
	temps, fronts := m.computeTempAreas()
	m.computeAreas(temps)
	m.createChokePoints(fronts)
	m.gr.computeChokePointDistanceMatrix(m)
	m.gr.computeGroupIds(m)
	m.createBases()
}

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

func (m *Map) markUnwalkableMiniTiles(game *bwapi.Game) {
	for y := 0; y < m.walkHeight; y++ {
		for x := 0; x < m.walkWidth; x++ {
			m.miniTiles[y*m.walkWidth+x].Walkable = game.IsWalkable(x, y)
		}
	}

	snapshot := make([]bool, len(m.miniTiles))
	for i, mt := range m.miniTiles {
		snapshot[i] = mt.Walkable
	}

	for y := 0; y < m.walkHeight; y++ {
		for x := 0; x < m.walkWidth; x++ {
			if snapshot[y*m.walkWidth+x] {
				continue
			}
			wp := bwapi.WalkPosition{X: int32(x), Y: int32(y)}
			for _, n := range m.walkNeighbors8(wp) {
				m.miniTiles[m.miniTileIndex(n)].Walkable = false
			}
		}
	}
}

func (m *Map) markBuildableTilesAndGH(game *bwapi.Game) {
	for ty := 0; ty < m.tileHeight; ty++ {
		for tx := 0; tx < m.tileWidth; tx++ {
			ti := ty*m.tileWidth + tx
			t := &m.tiles[ti]
			t.Buildable = game.IsBuildable(tx, ty)

			raw := game.GetGroundHeight(tx, ty)
			t.GroundHeight = int8(raw / 2)
			t.Doodad = raw%2 == 1

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

func (m *Map) decideSeasOrLakes() {
	visited := make([]bool, len(m.miniTiles))

	for y := 0; y < m.walkHeight; y++ {
		for x := 0; x < m.walkWidth; x++ {
			wp := bwapi.WalkPosition{X: int32(x), Y: int32(y)}
			idx := m.miniTileIndex(wp)
			if visited[idx] || m.miniTiles[idx].Walkable {
				continue
			}

			component := m.bfsFloodFill(wp, visited, func(_ bwapi.WalkPosition, mt *MiniTile) bool {
				return !mt.Walkable
			})

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

func (m *Map) initializeNeutralData(game *bwapi.Game) {
	for _, u := range game.GetStaticNeutralUnits() {
		ut := u.GetType()
		if !isMineralField(ut) && !isGeyser(ut) && !isSpecialBuilding(ut) {
			continue
		}

		tw, th := unitTypeTileSize(ut)
		pos := u.GetPosition()
		tp := neutralTilePosition(pos, tw, th)

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

func (m *Map) computeAltitude() {
	pq := &priorityQueue{}
	heap.Init(pq)

	altSet := make([]bool, len(m.miniTiles))

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
				mt.Alt = 1
				altSet[idx] = true
				continue
			}
			if !mt.Walkable {
				mt.Alt = 0
				altSet[idx] = true
				continue
			}
		}
	}

	for y := 0; y < m.walkHeight; y++ {
		for x := 0; x < m.walkWidth; x++ {
			wp := bwapi.WalkPosition{X: int32(x), Y: int32(y)}
			idx := m.miniTileIndex(wp)
			if !m.miniTiles[idx].Walkable || altSet[idx] {
				continue
			}
			for _, n := range m.walkNeighbors8(wp) {
				ni := m.miniTileIndex(n)
				if m.miniTiles[ni].Alt == 0 && altSet[ni] {
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

func (m *Map) countDoors(n *Neutral) int {
	visited := make(map[int]bool)
	var seeds []bwapi.WalkPosition

	for dy := -1; dy <= n.TileH; dy++ {
		for dx := -1; dx <= n.TileW; dx++ {
			if dy >= 0 && dy < n.TileH && dx >= 0 && dx < n.TileW {
				continue
			}
			tp := bwapi.TilePosition{X: n.TilePos.X + int32(dx), Y: n.TilePos.Y + int32(dy)}
			if !m.validTile(tp) {
				continue
			}
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

type tempArea struct {
	id          AreaId
	top         bwapi.WalkPosition
	topAltitude Altitude
	count       int
	mergedInto  AreaId
}

type frontier struct {
	wp    bwapi.WalkPosition
	area1 AreaId
	area2 AreaId
}

func (m *Map) computeTempAreas() ([]tempArea, []frontier) {
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
			temps = append(temps, tempArea{
				id:          nextID,
				top:         wa.wp,
				topAltitude: wa.alt,
				count:       1,
			})
			m.miniTiles[idx].AreaID = nextID
			nextID++

		case 1:
			var areaID AreaId
			for id := range neighborAreas {
				areaID = id
			}
			temps[areaID-1].count++
			m.miniTiles[idx].AreaID = areaID

		default:
			ids := make([]AreaId, 0, len(neighborAreas))
			for id := range neighborAreas {
				ids = append(ids, id)
			}
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
				lt.count += st.count
				st.mergedInto = larger
				m.miniTiles[idx].AreaID = larger

				for _, id := range ids[2:] {
					t := &temps[id-1]
					lt.count += t.count
					t.mergedInto = larger
				}
			} else {
				fronts = append(fronts, frontier{
					wp:    wa.wp,
					area1: larger,
					area2: smaller,
				})
				lt.count++
				m.miniTiles[idx].AreaID = larger
			}
		}
	}

	for i := range m.miniTiles {
		if m.miniTiles[i].AreaID > 0 {
			m.miniTiles[i].AreaID = resolveID(m.miniTiles[i].AreaID)
		}
	}

	for i := range fronts {
		fronts[i].area1 = resolveID(fronts[i].area1)
		fronts[i].area2 = resolveID(fronts[i].area2)
	}

	var filtered []frontier
	for _, f := range fronts {
		if f.area1 != f.area2 {
			filtered = append(filtered, f)
		}
	}

	return temps, filtered
}

func (m *Map) containsStartLocation(ta *tempArea) bool {
	for _, sl := range m.startLocations {
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

func (m *Map) computeAreas(temps []tempArea) {
	oldToNew := make(map[AreaId]AreaId)
	newID := AreaId(1)

	for i := range temps {
		t := &temps[i]
		if t.mergedInto != 0 {
			continue
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

	for i := range m.miniTiles {
		mt := &m.miniTiles[i]
		if mt.AreaID <= 0 {
			continue
		}
		if newAID, ok := oldToNew[mt.AreaID]; ok {
			mt.AreaID = newAID
		} else {
			mt.AreaID = -mt.AreaID
		}
	}

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

	for ty := 0; ty < m.tileHeight; ty++ {
		for tx := 0; tx < m.tileWidth; tx++ {
			ti := ty*m.tileWidth + tx
			t := &m.tiles[ti]

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

	for i := range m.minerals {
		n := &m.neutrals[m.minerals[i].NeutralIdx]
		if area := m.findNeutralArea(n); area != nil {
			area.MineralIdxs = append(area.MineralIdxs, i)
		}
	}
	for i := range m.geysers {
		n := &m.neutrals[m.geysers[i].NeutralIdx]
		if area := m.findNeutralArea(n); area != nil {
			area.GeyserIdxs = append(area.GeyserIdxs, i)
		}
	}
}

func (m *Map) findNeutralArea(n *Neutral) *Area {
	for dy := -1; dy <= n.TileH; dy++ {
		for dx := -1; dx <= n.TileW; dx++ {
			tp := bwapi.TilePosition{X: n.TilePos.X + int32(dx), Y: n.TilePos.Y + int32(dy)}
			if !m.validTile(tp) {
				continue
			}
			t := &m.tiles[m.tileIndex(tp)]
			if t.AreaID > 0 {
				return &m.areas[t.AreaID-1]
			}
		}
	}
	return m.NearestArea(n.Pos.ToWalkPosition())
}

func (m *Map) createChokePoints(fronts []frontier) {
	for i := range fronts {
		idx := m.miniTileIndex(fronts[i].wp)
		neighborAreas := make(map[AreaId]bool)
		for _, n := range m.walkNeighbors8(fronts[i].wp) {
			ni := m.miniTileIndex(n)
			aid := m.miniTiles[ni].AreaID
			if aid > 0 {
				neighborAreas[aid] = true
			}
		}
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

	for pair, positions := range groups {
		clusters := clusterPositions(positions, chokeClusterDistSq)

		for _, cluster := range clusters {
			cp := m.makeChokePoint(pair.a, pair.b, cluster)
			cp.Index = len(m.chokePoints)
			m.chokePoints = append(m.chokePoints, cp)

			if a := m.Area(pair.a); a != nil {
				a.ChokePointIdxs = append(a.ChokePointIdxs, cp.Index)
			}
			if b := m.Area(pair.b); b != nil {
				b.ChokePointIdxs = append(b.ChokePointIdxs, cp.Index)
			}
		}
	}

	for i := range m.neutrals {
		n := &m.neutrals[i]
		if !n.Blocking {
			continue
		}

		areasAround := m.areasAroundNeutral(n)
		if len(areasAround) < 2 {
			continue
		}

		ids := make([]AreaId, 0, len(areasAround))
		for id := range areasAround {
			ids = append(ids, id)
		}
		sort.Slice(ids, func(a, b int) bool { return ids[a] < ids[b] })

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

func (m *Map) makeChokePoint(a, b AreaId, cluster []bwapi.WalkPosition) ChokePoint {
	middle := cluster[0]
	middleAlt := m.miniTiles[m.miniTileIndex(middle)].Alt
	for _, wp := range cluster[1:] {
		alt := m.miniTiles[m.miniTileIndex(wp)].Alt
		if alt > middleAlt {
			middleAlt = alt
			middle = wp
		}
	}

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

func clusterPositions(positions []bwapi.WalkPosition, distSq int) [][]bwapi.WalkPosition {
	n := len(positions)
	if n == 0 {
		return nil
	}

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

func (m *Map) areasAroundNeutral(n *Neutral) map[AreaId]bool {
	areas := make(map[AreaId]bool)
	for dy := -1; dy <= n.TileH; dy++ {
		for dx := -1; dx <= n.TileW; dx++ {
			if dy >= 0 && dy < n.TileH && dx >= 0 && dx < n.TileW {
				continue
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

func (m *Map) createBases() {
	for areaIdx := range m.areas {
		area := &m.areas[areaIdx]
		m.createBasesForArea(area)
	}

	for _, sl := range m.startLocations {
		bestDist := math.MaxInt32
		bestIdx := -1
		for i := range m.bases {
			d := queenWiseDist(m.bases[i].Location, sl)
			if d < bestDist {
				bestDist = d
				bestIdx = i
			}
		}
		if bestIdx >= 0 && bestDist <= maxTilesBetweenStartAndBase {
			m.bases[bestIdx].IsStartLocation = true
		}
	}
}

type baseResource struct {
	neutralIdx int
	isGeyser   bool
	resIdx     int
}

func (m *Map) createBasesForArea(area *Area) {
	var remaining []baseResource
	for _, mi := range area.MineralIdxs {
		if m.minerals[mi].BaseIdx < 0 && !m.neutrals[m.minerals[mi].NeutralIdx].Blocking {
			remaining = append(remaining, baseResource{m.minerals[mi].NeutralIdx, false, mi})
		}
	}
	for _, gi := range area.GeyserIdxs {
		if m.geysers[gi].BaseIdx < 0 && !m.neutrals[m.geysers[gi].NeutralIdx].Blocking {
			remaining = append(remaining, baseResource{m.geysers[gi].NeutralIdx, true, gi})
		}
	}

	for len(remaining) > 0 {
		searchTopLeft, searchBottomRight := m.resourceSearchBox(remaining, area)

		m.markResourceScores(remaining, area.ID)

		bestScore := 0
		bestTP := bwapi.TilePosition{X: -1, Y: -1}

		for ty := searchTopLeft.Y; ty <= searchBottomRight.Y; ty++ {
			for tx := searchTopLeft.X; tx <= searchBottomRight.X; tx++ {
				loc := bwapi.TilePosition{X: tx, Y: ty}
				score := m.computeBaseLocationScore(loc, area.ID)
				if score <= 0 {
					continue
				}
				if score > bestScore && m.validateBaseLocation(loc, area.ID) {
					bestScore = score
					bestTP = loc
				}
			}
		}

		m.clearResourceScores(remaining)

		if bestTP.X < 0 {
			break
		}

		base := Base{
			Index:    len(m.bases),
			Location: bestTP,
			Center: bwapi.Position{
				X: bestTP.X*32 + ccTileWidth*16,
				Y: bestTP.Y*32 + ccTileHeight*16,
			},
			AreaID: area.ID,
		}

		ccTopLeft := bwapi.Position{X: bestTP.X * 32, Y: bestTP.Y * 32}
		ccSize := bwapi.Position{X: ccTileWidth * 32, Y: ccTileHeight * 32}

		var kept []baseResource
		for _, r := range remaining {
			n := &m.neutrals[r.neutralIdx]
			d := distToRectangle(n.Pos, ccTopLeft, ccSize)
			if d+2 <= maxTilesBetweenCCAndRes*32 {
				if r.isGeyser {
					m.geysers[r.resIdx].BaseIdx = base.Index
					base.GeyserIdxs = append(base.GeyserIdxs, r.resIdx)
				} else {
					m.minerals[r.resIdx].BaseIdx = base.Index
					base.MineralIdxs = append(base.MineralIdxs, r.resIdx)
				}
			} else {
				kept = append(kept, r)
			}
		}
		remaining = kept

		if len(base.MineralIdxs) > 0 || len(base.GeyserIdxs) > 0 {
			area.BaseIdxs = append(area.BaseIdxs, base.Index)
			m.bases = append(m.bases, base)
		} else {
			break
		}
	}
}

func (m *Map) resourceSearchBox(resources []baseResource, area *Area) (bwapi.TilePosition, bwapi.TilePosition) {
	minX := int32(m.tileWidth)
	minY := int32(m.tileHeight)
	maxX := int32(0)
	maxY := int32(0)
	for _, r := range resources {
		n := &m.neutrals[r.neutralIdx]
		if n.TilePos.X < minX {
			minX = n.TilePos.X
		}
		if n.TilePos.Y < minY {
			minY = n.TilePos.Y
		}
		rx := n.TilePos.X + int32(n.TileW) - 1
		ry := n.TilePos.Y + int32(n.TileH) - 1
		if rx > maxX {
			maxX = rx
		}
		if ry > maxY {
			maxY = ry
		}
	}

	topLeft := bwapi.TilePosition{
		X: minX - int32(ccTileWidth) - int32(maxTilesBetweenCCAndRes),
		Y: minY - int32(ccTileHeight) - int32(maxTilesBetweenCCAndRes),
	}
	bottomRight := bwapi.TilePosition{
		X: maxX + 1 + int32(maxTilesBetweenCCAndRes),
		Y: maxY + 1 + int32(maxTilesBetweenCCAndRes),
	}

	if topLeft.X < area.TopLeft.X {
		topLeft.X = area.TopLeft.X
	}
	if topLeft.Y < area.TopLeft.Y {
		topLeft.Y = area.TopLeft.Y
	}
	areaMaxX := area.BottomRight.X - int32(ccTileWidth) + 1
	areaMaxY := area.BottomRight.Y - int32(ccTileHeight) + 1
	if bottomRight.X > areaMaxX {
		bottomRight.X = areaMaxX
	}
	if bottomRight.Y > areaMaxY {
		bottomRight.Y = areaMaxY
	}

	return topLeft, bottomRight
}

func (m *Map) markResourceScores(resources []baseResource, areaID AreaId) {
	for _, r := range resources {
		n := &m.neutrals[r.neutralIdx]
		resTopLeft := bwapi.Position{X: n.TilePos.X * 32, Y: n.TilePos.Y * 32}
		resSize := bwapi.Position{X: int32(n.TileW) * 32, Y: int32(n.TileH) * 32}

		rangeX := int(ccTileWidth) + maxTilesBetweenCCAndRes
		rangeY := int(ccTileHeight) + maxTilesBetweenCCAndRes

		for dy := -int(ccTileHeight) - maxTilesBetweenCCAndRes; dy < int(n.TileH)+rangeY; dy++ {
			for dx := -int(ccTileWidth) - maxTilesBetweenCCAndRes; dx < int(n.TileW)+rangeX; dx++ {
				tp := bwapi.TilePosition{X: n.TilePos.X + int32(dx), Y: n.TilePos.Y + int32(dy)}
				if !m.validTile(tp) {
					continue
				}
				t := &m.tiles[m.tileIndex(tp)]
				if t.AreaID != areaID {
					continue
				}

				tileCenter := bwapi.Position{X: tp.X*32 + 16, Y: tp.Y*32 + 16}
				pixDist := distToRectangle(tileCenter, resTopLeft, resSize)
				tileDist := (pixDist + 16) / 32
				score := maxTilesBetweenCCAndRes + resourceExclusionGap - tileDist
				if score <= 0 {
					continue
				}
				if r.isGeyser {
					score *= 3
				}
				t.internalData += score
			}
		}
	}

	for _, r := range resources {
		n := &m.neutrals[r.neutralIdx]
		for dy := -resourceExclusionGap; dy < int(n.TileH)+resourceExclusionGap; dy++ {
			for dx := -resourceExclusionGap; dx < int(n.TileW)+resourceExclusionGap; dx++ {
				tp := bwapi.TilePosition{X: n.TilePos.X + int32(dx), Y: n.TilePos.Y + int32(dy)}
				if m.validTile(tp) {
					m.tiles[m.tileIndex(tp)].internalData = -1
				}
			}
		}
	}
}

func (m *Map) clearResourceScores(resources []baseResource) {
	for _, r := range resources {
		n := &m.neutrals[r.neutralIdx]
		rangeX := int(ccTileWidth) + maxTilesBetweenCCAndRes
		rangeY := int(ccTileHeight) + maxTilesBetweenCCAndRes

		for dy := -int(ccTileHeight) - maxTilesBetweenCCAndRes; dy < int(n.TileH)+rangeY; dy++ {
			for dx := -int(ccTileWidth) - maxTilesBetweenCCAndRes; dx < int(n.TileW)+rangeX; dx++ {
				tp := bwapi.TilePosition{X: n.TilePos.X + int32(dx), Y: n.TilePos.Y + int32(dy)}
				if m.validTile(tp) {
					m.tiles[m.tileIndex(tp)].internalData = 0
				}
			}
		}
	}
}

func (m *Map) computeBaseLocationScore(loc bwapi.TilePosition, areaID AreaId) int {
	score := 0
	for dy := 0; dy < ccTileHeight; dy++ {
		for dx := 0; dx < ccTileWidth; dx++ {
			tp := bwapi.TilePosition{X: loc.X + int32(dx), Y: loc.Y + int32(dy)}
			if !m.validTile(tp) {
				return -1
			}
			t := &m.tiles[m.tileIndex(tp)]
			if !t.Buildable {
				return -1
			}
			if t.AreaID != areaID {
				return -1
			}
			if t.internalData < 0 {
				return -1
			}
			if t.NeutralIdx >= 0 {
				n := &m.neutrals[t.NeutralIdx]
				if isSpecialBuilding(n.UnitType) {
					return -1
				}
			}
			score += t.internalData
		}
	}
	return score
}

func (m *Map) validateBaseLocation(loc bwapi.TilePosition, areaID AreaId) bool {
	for dy := -resourceExclusionGap; dy < ccTileHeight+resourceExclusionGap; dy++ {
		for dx := -resourceExclusionGap; dx < ccTileWidth+resourceExclusionGap; dx++ {
			tp := bwapi.TilePosition{X: loc.X + int32(dx), Y: loc.Y + int32(dy)}
			if !m.validTile(tp) {
				continue
			}
			t := &m.tiles[m.tileIndex(tp)]
			if t.NeutralIdx >= 0 {
				n := &m.neutrals[t.NeutralIdx]
				if isGeyser(n.UnitType) {
					return false
				}
				if isMineralField(n.UnitType) {
					nRes := m.findMineralByNeutralIdx(t.NeutralIdx)
					if nRes != nil && nRes.Resources > 8 {
						return false
					}
				}
			}
		}
	}

	for i := range m.bases {
		if roundedDist(m.bases[i].Location, loc) < minTilesBetweenBases {
			return false
		}
	}

	return true
}

func (m *Map) findMineralByNeutralIdx(neutralIdx int) *Mineral {
	for i := range m.minerals {
		if m.minerals[i].NeutralIdx == neutralIdx {
			return &m.minerals[i]
		}
	}
	return nil
}

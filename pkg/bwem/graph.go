package bwem

import (
	"container/heap"

	"github.com/bradewing/gobwapi/pkg/bwapi"
)

// graph holds the precomputed chokepoint distance matrix and path lookup.
type graph struct {
	// dist[i][j] = shortest ground distance (pixels) between chokepoints i and j.
	// -1 if unreachable.
	dist [][]int

	// nextCP[i][j] = next chokepoint index on the shortest path from i to j.
	// Used for O(1) path reconstruction.
	nextCP [][]int
}

// getDistance returns the precomputed distance between two chokepoints.
func (g *graph) getDistance(fromCP, toCP int) int {
	if fromCP >= len(g.dist) || toCP >= len(g.dist) {
		return -1
	}
	return g.dist[fromCP][toCP]
}

// getPath reconstructs the shortest path between two chokepoints as a slice
// of chokepoint indices (inclusive of both endpoints).
func (g *graph) getPath(fromCP, toCP int) []int {
	if fromCP == toCP {
		return []int{fromCP}
	}
	if g.dist[fromCP][toCP] < 0 {
		return nil
	}

	var path []int
	cur := fromCP
	for cur != toCP {
		path = append(path, cur)
		next := g.nextCP[cur][toCP]
		if next < 0 || next == cur {
			return nil // should not happen
		}
		cur = next
	}
	path = append(path, toCP)
	return path
}

// computeChokePointDistanceMatrix builds the all-pairs shortest path matrix.
func (g *graph) computeChokePointDistanceMatrix(m *Map) {
	n := len(m.chokePoints)
	if n == 0 {
		return
	}

	// Initialize matrices.
	g.dist = make([][]int, n)
	g.nextCP = make([][]int, n)
	for i := 0; i < n; i++ {
		g.dist[i] = make([]int, n)
		g.nextCP[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				g.dist[i][j] = 0
				g.nextCP[i][j] = j
			} else {
				g.dist[i][j] = -1
				g.nextCP[i][j] = -1
			}
		}
	}

	// Phase 1: For each area, compute within-area distances between all
	// chokepoints bordering that area using BFS on walkable MiniTiles.
	for areaIdx := range m.areas {
		area := &m.areas[areaIdx]
		if len(area.ChokePointIdxs) < 2 {
			continue
		}
		m.computeWithinAreaDistances(area, g)
	}

	// Phase 2: Floyd-Warshall for all-pairs shortest paths.
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if g.dist[i][k] < 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if g.dist[k][j] < 0 {
					continue
				}
				newDist := g.dist[i][k] + g.dist[k][j]
				if g.dist[i][j] < 0 || newDist < g.dist[i][j] {
					g.dist[i][j] = newDist
					g.nextCP[i][j] = g.nextCP[i][k]
				}
			}
		}
	}
}

// computeWithinAreaDistances computes ground distances between chokepoints
// within a single area using Dijkstra from each chokepoint's middle position.
func (m *Map) computeWithinAreaDistances(area *Area, g *graph) {
	// For each chokepoint in this area, BFS/Dijkstra to find distances to
	// other chokepoints in the same area.
	cpMiddles := make(map[int]int) // miniTileIndex -> chokepoint index
	for _, cpIdx := range area.ChokePointIdxs {
		cp := &m.chokePoints[cpIdx]
		mi := m.miniTileIndex(cp.Middle)
		cpMiddles[mi] = cpIdx
	}

	for _, srcCPIdx := range area.ChokePointIdxs {
		srcCP := &m.chokePoints[srcCPIdx]

		// Dijkstra from srcCP.Middle within this area.
		dist := make(map[int]int) // miniTileIndex -> distance
		pq := &graphPQ{}
		heap.Init(pq)

		startIdx := m.miniTileIndex(srcCP.Middle)
		dist[startIdx] = 0
		heap.Push(pq, graphPQItem{node: startIdx, dist: 0})

		found := 0
		target := len(area.ChokePointIdxs) - 1 // exclude self

		for pq.Len() > 0 && found < target {
			item := heap.Pop(pq).(graphPQItem)
			if item.dist > dist[item.node] {
				continue // stale entry
			}

			// Check if we reached another chokepoint.
			if dstCPIdx, ok := cpMiddles[item.node]; ok && dstCPIdx != srcCPIdx {
				// Record the distance (symmetric).
				if g.dist[srcCPIdx][dstCPIdx] < 0 || item.dist < g.dist[srcCPIdx][dstCPIdx] {
					g.dist[srcCPIdx][dstCPIdx] = item.dist
					g.dist[dstCPIdx][srcCPIdx] = item.dist
					g.nextCP[srcCPIdx][dstCPIdx] = dstCPIdx
					g.nextCP[dstCPIdx][srcCPIdx] = srcCPIdx
				}
				found++
			}

			// Expand neighbors within this area.
			curWP := indexToWalk(item.node, m.walkWidth)
			for _, n := range m.walkNeighbors8(curWP) {
				ni := m.miniTileIndex(n)
				mt := &m.miniTiles[ni]
				if !mt.Walkable || (mt.AreaID != area.ID && mt.AreaID > 0) {
					continue
				}
				// Distance: 8 pixels for cardinal, 11 (~8*sqrt2) for diagonal.
				dx := n.X - curWP.X
				dy := n.Y - curWP.Y
				cost := 8
				if dx != 0 && dy != 0 {
					cost = 11
				}
				newDist := item.dist + cost
				if d, ok := dist[ni]; !ok || newDist < d {
					dist[ni] = newDist
					heap.Push(pq, graphPQItem{node: ni, dist: newDist})
				}
			}
		}
	}
}

// indexToWalk converts a flat miniTile index back to a WalkPosition.
func indexToWalk(idx, walkWidth int) bwapi.WalkPosition {
	return bwapi.WalkPosition{
		X: int32(idx % walkWidth),
		Y: int32(idx / walkWidth),
	}
}

// computeGroupIds assigns GroupIds to areas via DFS over chokepoint connectivity.
func (g *graph) computeGroupIds(m *Map) {
	nextGroup := GroupId(1)
	for i := range m.areas {
		if m.areas[i].GroupID > 0 {
			continue
		}
		// DFS from this area.
		m.areas[i].GroupID = nextGroup
		stack := []AreaId{m.areas[i].ID}
		for len(stack) > 0 {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curArea := m.Area(cur)
			if curArea == nil {
				continue
			}
			for _, cpIdx := range curArea.ChokePointIdxs {
				cp := &m.chokePoints[cpIdx]
				if cp.Blocked {
					continue
				}
				otherID := cp.OtherArea(cur)
				other := m.Area(otherID)
				if other == nil || other.GroupID > 0 {
					continue
				}
				other.GroupID = nextGroup
				stack = append(stack, otherID)
			}
		}
		nextGroup++
	}

	// Populate NeighborIDs for each area.
	for i := range m.areas {
		area := &m.areas[i]
		seen := make(map[AreaId]bool)
		for _, cpIdx := range area.ChokePointIdxs {
			cp := &m.chokePoints[cpIdx]
			if cp.Blocked {
				continue
			}
			otherID := cp.OtherArea(area.ID)
			if !seen[otherID] {
				seen[otherID] = true
				area.NeighborIDs = append(area.NeighborIDs, otherID)
			}
		}
	}
}

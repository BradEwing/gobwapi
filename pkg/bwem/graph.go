package bwem

import (
	"container/heap"

	"github.com/bradewing/gobwapi/pkg/bwapi"
)

type graph struct {
	dist   [][]int
	nextCP [][]int
}

func (g *graph) getDistance(fromCP, toCP int) int {
	if fromCP >= len(g.dist) || toCP >= len(g.dist) {
		return -1
	}
	return g.dist[fromCP][toCP]
}

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
			return nil
		}
		cur = next
	}
	path = append(path, toCP)
	return path
}

func (g *graph) computeChokePointDistanceMatrix(m *Map) {
	n := len(m.chokePoints)
	if n == 0 {
		return
	}

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

	for areaIdx := range m.areas {
		area := &m.areas[areaIdx]
		if len(area.ChokePointIdxs) < 2 {
			continue
		}
		m.computeWithinAreaDistances(area, g)
	}

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

func (m *Map) computeWithinAreaDistances(area *Area, g *graph) {
	cpMiddles := make(map[int]int)
	for _, cpIdx := range area.ChokePointIdxs {
		cp := &m.chokePoints[cpIdx]
		mi := m.miniTileIndex(cp.Middle)
		cpMiddles[mi] = cpIdx
	}

	for _, srcCPIdx := range area.ChokePointIdxs {
		srcCP := &m.chokePoints[srcCPIdx]

		dist := make(map[int]int)
		pq := &graphPQ{}
		heap.Init(pq)

		startIdx := m.miniTileIndex(srcCP.Middle)
		dist[startIdx] = 0
		heap.Push(pq, graphPQItem{node: startIdx, dist: 0})

		found := 0
		target := len(area.ChokePointIdxs) - 1

		for pq.Len() > 0 && found < target {
			item := heap.Pop(pq).(graphPQItem)
			if item.dist > dist[item.node] {
				continue
			}

			if dstCPIdx, ok := cpMiddles[item.node]; ok && dstCPIdx != srcCPIdx {
				if g.dist[srcCPIdx][dstCPIdx] < 0 || item.dist < g.dist[srcCPIdx][dstCPIdx] {
					g.dist[srcCPIdx][dstCPIdx] = item.dist
					g.dist[dstCPIdx][srcCPIdx] = item.dist
					g.nextCP[srcCPIdx][dstCPIdx] = dstCPIdx
					g.nextCP[dstCPIdx][srcCPIdx] = srcCPIdx
				}
				found++
			}

			curWP := indexToWalk(item.node, m.walkWidth)
			for _, n := range m.walkNeighbors8(curWP) {
				ni := m.miniTileIndex(n)
				mt := &m.miniTiles[ni]
				if !mt.Walkable || (mt.AreaID != area.ID && mt.AreaID > 0) {
					continue
				}
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

func indexToWalk(idx, walkWidth int) bwapi.WalkPosition {
	return bwapi.WalkPosition{
		X: int32(idx % walkWidth),
		Y: int32(idx / walkWidth),
	}
}

func (g *graph) computeGroupIds(m *Map) {
	nextGroup := GroupId(1)
	for i := range m.areas {
		if m.areas[i].GroupID > 0 {
			continue
		}
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

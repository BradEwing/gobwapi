package bwem

import (
	"container/heap"
	"math"

	"github.com/bradewing/gobwapi/pkg/bwapi"
)

func (m *Map) miniTileIndex(wp bwapi.WalkPosition) int {
	return int(wp.Y)*m.walkWidth + int(wp.X)
}

func (m *Map) tileIndex(tp bwapi.TilePosition) int {
	return int(tp.Y)*m.tileWidth + int(tp.X)
}

func (m *Map) validWalk(wp bwapi.WalkPosition) bool {
	return wp.X >= 0 && wp.Y >= 0 && int(wp.X) < m.walkWidth && int(wp.Y) < m.walkHeight
}

func (m *Map) validTile(tp bwapi.TilePosition) bool {
	return tp.X >= 0 && tp.Y >= 0 && int(tp.X) < m.tileWidth && int(tp.Y) < m.tileHeight
}

var dirs8 = [8][2]int32{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

var dirs4 = [4][2]int32{
	{0, -1}, {-1, 0}, {1, 0}, {0, 1},
}

func (m *Map) walkNeighbors8(wp bwapi.WalkPosition) []bwapi.WalkPosition {
	var result []bwapi.WalkPosition
	for _, d := range dirs8 {
		n := bwapi.WalkPosition{X: wp.X + d[0], Y: wp.Y + d[1]}
		if m.validWalk(n) {
			result = append(result, n)
		}
	}
	return result
}

func (m *Map) walkNeighbors4(wp bwapi.WalkPosition) []bwapi.WalkPosition {
	var result []bwapi.WalkPosition
	for _, d := range dirs4 {
		n := bwapi.WalkPosition{X: wp.X + d[0], Y: wp.Y + d[1]}
		if m.validWalk(n) {
			result = append(result, n)
		}
	}
	return result
}

func (m *Map) bfsFloodFill(
	seed bwapi.WalkPosition,
	visited []bool,
	predicate func(bwapi.WalkPosition, *MiniTile) bool,
) []bwapi.WalkPosition {
	idx := m.miniTileIndex(seed)
	if visited[idx] || !predicate(seed, &m.miniTiles[idx]) {
		return nil
	}

	var component []bwapi.WalkPosition
	queue := []bwapi.WalkPosition{seed}
	visited[idx] = true

	for len(queue) > 0 {
		wp := queue[0]
		queue = queue[1:]
		component = append(component, wp)

		for _, n := range m.walkNeighbors8(wp) {
			ni := m.miniTileIndex(n)
			if !visited[ni] && predicate(n, &m.miniTiles[ni]) {
				visited[ni] = true
				queue = append(queue, n)
			}
		}
	}
	return component
}

func walkDist(a, b bwapi.WalkPosition) int {
	dx := float64(a.X - b.X)
	dy := float64(a.Y - b.Y)
	return int(math.Sqrt(dx*dx+dy*dy) * 8)
}

func walkDistSq(a, b bwapi.WalkPosition) int {
	dx := int(a.X - b.X)
	dy := int(a.Y - b.Y)
	return dx*dx + dy*dy
}

type pqItem struct {
	wp   bwapi.WalkPosition
	dist int
}

type priorityQueue []pqItem

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqItem)) }
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

type graphPQItem struct {
	node int
	dist int
}

type graphPQ []graphPQItem

func (pq graphPQ) Len() int            { return len(pq) }
func (pq graphPQ) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq graphPQ) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *graphPQ) Push(x interface{}) { *pq = append(*pq, x.(graphPQItem)) }
func (pq *graphPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

var _ heap.Interface = (*priorityQueue)(nil)
var _ heap.Interface = (*graphPQ)(nil)

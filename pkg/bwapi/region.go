package bwapi

import (
	"github.com/bradewing/gobwapi/internal/shm"
)

// Region wraps a BWAPI pathfinding region.
type Region struct {
	data *shm.RegionData
	game *Game
}

func (r *Region) ID() int              { return int(r.data.ID()) }
func (r *Region) IslandID() int        { return int(r.data.IslandID()) }
func (r *Region) IsAccessible() bool   { return r.data.IsAccessible() }
func (r *Region) IsHigherGround() bool { return r.data.IsHigherGround() }
func (r *Region) Priority() int        { return int(r.data.Priority()) }

func (r *Region) GetCenter() Position {
	return Position{X: r.data.CenterX(), Y: r.data.CenterY()}
}

func (r *Region) BoundsLeft() int   { return int(r.data.LeftMost()) }
func (r *Region) BoundsRight() int  { return int(r.data.RightMost()) }
func (r *Region) BoundsTop() int    { return int(r.data.TopMost()) }
func (r *Region) BoundsBottom() int { return int(r.data.BottomMost()) }

// GetNeighbors returns the IDs of neighboring regions.
func (r *Region) GetNeighbors() []int {
	count := int(r.data.NeighborCount())
	neighbors := make([]int, count)
	for i := 0; i < count; i++ {
		neighbors[i] = int(r.data.Neighbor(i))
	}
	return neighbors
}

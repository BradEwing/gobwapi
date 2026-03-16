package bwapi

import (
	"github.com/bradewing/gobwapi/internal/shm"
)

// Force wraps a BWAPI force (team grouping).
type Force struct {
	data  *shm.ForceData
	game  *Game
	index int
}

func (f *Force) Index() int   { return f.index }
func (f *Force) Name() string { return f.data.Name() }

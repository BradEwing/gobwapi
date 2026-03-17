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

// GetPlayers returns all players belonging to this force.
func (f *Force) GetPlayers() []*Player {
	count := f.game.data.PlayerCount()
	players := make([]*Player, 0)
	for i := 0; i < count; i++ {
		pd := f.game.data.Player(i)
		if int(pd.ForceIndex()) == f.index {
			players = append(players, &Player{data: pd, game: f.game, index: i})
		}
	}
	return players
}

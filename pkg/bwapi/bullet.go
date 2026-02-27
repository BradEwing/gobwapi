package bwapi

import (
	"github.com/bradewing/gobwapi/internal/shm"
)

// Bullet wraps a BWAPI bullet's shared memory data.
type Bullet struct {
	data *shm.BulletData
	game *Game
}

func (b *Bullet) ID() int              { return int(b.data.ID()) }
func (b *Bullet) Exists() bool         { return b.data.Exists() }
func (b *Bullet) GetType() BulletType  { return BulletType(b.data.TypeID()) }

func (b *Bullet) GetPlayer() *Player {
	idx := int(b.data.PlayerIndex())
	if idx < 0 || idx >= shm.MaxPlayers {
		return nil
	}
	return &Player{data: b.game.data.Player(idx), game: b.game, index: idx}
}

func (b *Bullet) GetSource() *Unit {
	idx := int(b.data.SourceIndex())
	if idx < 0 || idx >= shm.MaxUnits {
		return nil
	}
	return b.game.GetUnit(idx)
}

func (b *Bullet) GetPosition() Position {
	return Position{X: b.data.PositionX(), Y: b.data.PositionY()}
}

func (b *Bullet) Angle() float64     { return b.data.Angle() }
func (b *Bullet) VelocityX() float64 { return b.data.VelocityX() }
func (b *Bullet) VelocityY() float64 { return b.data.VelocityY() }

func (b *Bullet) GetTarget() *Unit {
	idx := int(b.data.TargetIndex())
	if idx < 0 || idx >= shm.MaxUnits {
		return nil
	}
	return b.game.GetUnit(idx)
}

func (b *Bullet) GetTargetPosition() Position {
	return Position{X: b.data.TargetPositionX(), Y: b.data.TargetPositionY()}
}

func (b *Bullet) RemoveTimer() int { return int(b.data.RemoveTimer()) }

func (b *Bullet) IsVisibleTo(playerIndex int) bool {
	return b.data.IsVisibleTo(playerIndex)
}

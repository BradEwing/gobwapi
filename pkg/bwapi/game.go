package bwapi

import (
	"github.com/bradewing/gobwapi/internal/shm"
)

// Game provides the main interface for reading game state and issuing commands.
// It wraps the BWAPI shared memory GameData struct.
type Game struct {
	data *shm.GameData

	// Cached object slices, rebuilt each frame.
	units   []*Unit
	players []*Player
	bullets []*Bullet
	regions []*Region
}

// NewGame creates a Game wrapping the given shared memory data.
func NewGame(data *shm.GameData) *Game {
	return &Game{data: data}
}

// Data returns the underlying shared memory data for advanced access.
func (g *Game) Data() *shm.GameData {
	return g.data
}

// --- Frame Info ---

func (g *Game) FrameCount() int     { return int(g.data.FrameCount()) }
func (g *Game) FPS() int            { return int(g.data.FPS()) }
func (g *Game) AverageFPS() float64 { return g.data.AverageFPS() }
func (g *Game) IsInGame() bool      { return g.data.IsInGame() }
func (g *Game) IsMultiplayer() bool { return g.data.IsMultiplayer() }
func (g *Game) IsReplay() bool      { return g.data.IsReplay() }
func (g *Game) IsPaused() bool      { return g.data.IsPaused() }
func (g *Game) IsBattleNet() bool   { return g.data.IsBattleNet() }

func (g *Game) Latency() int          { return int(g.data.Latency()) }
func (g *Game) LatencyFrames() int    { return int(g.data.LatencyFrames()) }
func (g *Game) HasLatCom() bool       { return g.data.HasLatCom() }
func (g *Game) HasGUI() bool          { return g.data.HasGUI() }
func (g *Game) RandomSeed() uint32    { return g.data.RandomSeed() }
func (g *Game) ReplayFrameCount() int { return int(g.data.ReplayFrameCount()) }
func (g *Game) ElapsedTime() int      { return int(g.data.ElapsedTime()) }
func (g *Game) CountdownTimer() int   { return int(g.data.CountdownTimer()) }

// --- Map Info ---

func (g *Game) MapWidth() int       { return int(g.data.MapWidth()) }
func (g *Game) MapHeight() int      { return int(g.data.MapHeight()) }
func (g *Game) MapFileName() string { return g.data.MapFileName() }
func (g *Game) MapPathName() string { return g.data.MapPathName() }
func (g *Game) MapName() string     { return g.data.MapName() }
func (g *Game) MapHash() string     { return g.data.MapHash() }

func (g *Game) GetGroundHeight(tileX, tileY int) int {
	return int(g.data.GetGroundHeight(tileX, tileY))
}

func (g *Game) IsWalkable(walkX, walkY int) bool {
	return g.data.IsWalkable(walkX, walkY)
}

func (g *Game) IsBuildable(tileX, tileY int) bool {
	return g.data.IsBuildable(tileX, tileY)
}

func (g *Game) IsVisible(tileX, tileY int) bool {
	return g.data.IsVisible(tileX, tileY)
}

func (g *Game) IsExplored(tileX, tileY int) bool {
	return g.data.IsExplored(tileX, tileY)
}

func (g *Game) HasCreep(tileX, tileY int) bool {
	return g.data.HasCreep(tileX, tileY)
}

// GetStartLocations returns the map's start positions.
func (g *Game) GetStartLocations() []TilePosition {
	count := g.data.StartLocationCount()
	locs := make([]TilePosition, count)
	for i := 0; i < count; i++ {
		locs[i] = TilePosition{
			X: g.data.StartLocationX(i),
			Y: g.data.StartLocationY(i),
		}
	}
	return locs
}

// --- Players ---

// Self returns the player representing this bot.
func (g *Game) Self() *Player {
	idx := int(g.data.SelfIndex())
	if idx < 0 || idx >= shm.MaxPlayers {
		return nil
	}
	return &Player{data: g.data.Player(idx), game: g, index: idx}
}

// Enemy returns the enemy player (first enemy found).
func (g *Game) Enemy() *Player {
	idx := int(g.data.EnemyIndex())
	if idx < 0 || idx >= shm.MaxPlayers {
		return nil
	}
	return &Player{data: g.data.Player(idx), game: g, index: idx}
}

// Neutral returns the neutral player.
func (g *Game) Neutral() *Player {
	idx := int(g.data.NeutralIndex())
	if idx < 0 || idx >= shm.MaxPlayers {
		return nil
	}
	return &Player{data: g.data.Player(idx), game: g, index: idx}
}

// GetPlayers returns all players in the game.
func (g *Game) GetPlayers() []*Player {
	count := g.data.PlayerCount()
	players := make([]*Player, count)
	for i := 0; i < count; i++ {
		players[i] = &Player{data: g.data.Player(i), game: g, index: i}
	}
	return players
}

// --- Units ---

// GetAllUnits returns all visible units from the unit array.
func (g *Game) GetAllUnits() []*Unit {
	units := make([]*Unit, 0, shm.UnitArraySize)
	for i := 0; i < shm.UnitArraySize; i++ {
		idx := int(g.data.UnitArrayEntry(i))
		if idx < 0 {
			continue
		}
		ud := g.data.Unit(idx)
		if !ud.Exists() {
			continue
		}
		units = append(units, &Unit{data: ud, game: g, index: idx})
	}
	return units
}

// GetUnit returns a specific unit by its index.
func (g *Game) GetUnit(index int) *Unit {
	if index < 0 || index >= shm.MaxUnits {
		return nil
	}
	ud := g.data.Unit(index)
	return &Unit{data: ud, game: g, index: index}
}

// --- Bullets ---

// GetBullets returns all active bullets.
func (g *Game) GetBullets() []*Bullet {
	bullets := make([]*Bullet, 0)
	for i := 0; i < shm.MaxBullets; i++ {
		bd := g.data.Bullet(i)
		if !bd.Exists() {
			continue
		}
		bullets = append(bullets, &Bullet{data: bd, game: g})
	}
	return bullets
}

// --- Regions ---

// GetRegions returns all pathfinding regions.
func (g *Game) GetRegions() []*Region {
	count := g.data.RegionCount()
	regions := make([]*Region, count)
	for i := 0; i < count; i++ {
		regions[i] = &Region{data: g.data.Region(i), game: g}
	}
	return regions
}

// --- Nuke Dots ---

// GetNukeDots returns positions of detected nuclear strikes.
func (g *Game) GetNukeDots() []Position {
	count := g.data.NukeDotCount()
	dots := make([]Position, count)
	for i := 0; i < count; i++ {
		dots[i] = Position{X: g.data.NukeDotX(i), Y: g.data.NukeDotY(i)}
	}
	return dots
}

// --- Input ---

func (g *Game) MouseX() int  { return int(g.data.MouseX()) }
func (g *Game) MouseY() int  { return int(g.data.MouseY()) }
func (g *Game) ScreenX() int { return int(g.data.ScreenX()) }
func (g *Game) ScreenY() int { return int(g.data.ScreenY()) }

// --- Commands ---

// SetScreenPosition moves the screen to the given pixel coordinates.
func (g *Game) SetScreenPosition(x, y int) {
	g.data.AddCommand(int32(CommandTypeSetScreenPosition), int32(x), int32(y))
}

// PingMinimap pings the minimap at pixel coordinates.
func (g *Game) PingMinimap(x, y int) {
	g.data.AddCommand(int32(CommandTypePingMinimap), int32(x), int32(y))
}

// EnableFlag enables a BWAPI flag.
func (g *Game) EnableFlag(flag int) {
	g.data.AddCommand(int32(CommandTypeEnableFlag), int32(flag), 0)
}

// Printf prints text to the game screen.
func (g *Game) Printf(text string) {
	idx := g.data.AddString(text)
	g.data.AddCommand(int32(CommandTypePrintf), idx, 0)
}

// SendText sends a text message in game.
func (g *Game) SendText(text string) {
	idx := g.data.AddString(text)
	g.data.AddCommand(int32(CommandTypeSendText), idx, 0)
}

// PauseGame pauses the game.
func (g *Game) PauseGame() {
	g.data.AddCommand(int32(CommandTypePauseGame), 0, 0)
}

// ResumeGame resumes a paused game.
func (g *Game) ResumeGame() {
	g.data.AddCommand(int32(CommandTypeResumeGame), 0, 0)
}

// LeaveGame leaves the current game.
func (g *Game) LeaveGame() {
	g.data.AddCommand(int32(CommandTypeLeaveGame), 0, 0)
}

// SetLocalSpeed sets the game speed.
func (g *Game) SetLocalSpeed(speed int) {
	g.data.AddCommand(int32(CommandTypeSetLocalSpeed), int32(speed), 0)
}

// SetFrameSkip sets the number of frames to skip rendering.
func (g *Game) SetFrameSkip(frameSkip int) {
	g.data.AddCommand(int32(CommandTypeSetFrameSkip), int32(frameSkip), 0)
}

// SetLatCom enables or disables latency compensation.
func (g *Game) SetLatCom(enabled bool) {
	v := int32(0)
	if enabled {
		v = 1
	}
	g.data.AddCommand(int32(CommandTypeSetLatCom), v, 0)
}

// SetGUI enables or disables the GUI.
func (g *Game) SetGUI(enabled bool) {
	v := int32(0)
	if enabled {
		v = 1
	}
	g.data.AddCommand(int32(CommandTypeSetGui), v, 0)
}

package bwapi

import (
	"github.com/bradewing/gobwapi/internal/shm"
)

// initialUnitState stores a snapshot of a unit's state at game start.
type initialUnitState struct {
	unitType     UnitType
	position     Position
	tilePosition TilePosition
	hitPoints    int
	resources    int
}

// Game provides the main interface for reading game state and issuing commands.
// It wraps the BWAPI shared memory GameData struct.
type Game struct {
	data *shm.GameData

	// Cached object slices, rebuilt each frame.
	units   []*Unit
	players []*Player
	bullets []*Bullet
	regions []*Region

	// Snapshot of initial unit states, populated by SnapshotInitialState.
	initialStates map[int]initialUnitState
}

// NewGame creates a Game wrapping the given shared memory data.
func NewGame(data *shm.GameData) *Game {
	return &Game{data: data}
}

// Data returns the underlying shared memory data for advanced access.
func (g *Game) Data() *shm.GameData {
	return g.data
}

// SnapshotInitialState captures the initial state of all units.
// Should be called during OnStart before any game frames are processed.
func (g *Game) SnapshotInitialState() {
	count := g.data.InitialUnitCount()
	g.initialStates = make(map[int]initialUnitState, count)
	for i := 0; i < count; i++ {
		ud := g.data.Unit(i)
		g.initialStates[i] = initialUnitState{
			unitType:     UnitType(ud.TypeID()),
			position:     Position{X: ud.PositionX(), Y: ud.PositionY()},
			tilePosition: TilePosition{X: ud.PositionX() / 32, Y: ud.PositionY() / 32},
			hitPoints:    int(ud.HitPoints()),
			resources:    int(ud.Resources()),
		}
	}
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

// GetAllUnits returns all units visible to the self player.
func (g *Game) GetAllUnits() []*Unit {
	selfIdx := int(g.data.SelfIndex())
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
		if !ud.IsVisibleTo(selfIdx) {
			continue
		}
		units = append(units, &Unit{data: ud, game: g, index: idx})
	}
	return units
}

// GetInitialUnits returns all units that existed at game start
// (minerals, geysers, static buildings, player units). Available during OnStart.
func (g *Game) GetInitialUnits() []*Unit {
	count := g.data.InitialUnitCount()
	units := make([]*Unit, 0, count)
	for i := 0; i < count; i++ {
		ud := g.data.Unit(i)
		units = append(units, &Unit{data: ud, game: g, index: i})
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

// --- API Info ---

func (g *Game) GetRevision() int        { return int(g.data.Revision()) }
func (g *Game) GetClientVersion() int   { return int(g.data.ClientVersion()) }
func (g *Game) IsDebug() bool           { return g.data.IsDebug() }
func (g *Game) GetInstanceNumber() int  { return int(g.data.InstanceID()) }
func (g *Game) LatencyTime() int        { return int(g.data.LatencyTime()) }
func (g *Game) RemainingLatencyFrames() int { return int(g.data.RemainingLatencyFrames()) }
func (g *Game) RemainingLatencyTime() int   { return int(g.data.RemainingLatencyTime()) }

// GetAPM returns the bot's actions per minute.
// If includeSelects is true, selection commands are included in the count.
func (g *Game) GetAPM(includeSelects bool) int {
	if includeSelects {
		return int(g.data.BotAPMSelects())
	}
	return int(g.data.BotAPMNoSelects())
}

// --- Input ---

// GetMousePosition returns the mouse position as a Position.
func (g *Game) GetMousePosition() Position {
	return Position{X: g.data.MouseX(), Y: g.data.MouseY()}
}

// GetScreenPosition returns the screen position as a Position.
func (g *Game) GetScreenPosition() Position {
	return Position{X: g.data.ScreenX(), Y: g.data.ScreenY()}
}

// GetMouseState returns whether a mouse button is pressed.
func (g *Game) GetMouseState(button int) bool { return g.data.MouseState(button) }

// GetKeyState returns whether a keyboard key is pressed.
func (g *Game) GetKeyState(key int) bool { return g.data.KeyState(key) }

// IsFlagEnabled returns whether a BWAPI flag is currently enabled.
func (g *Game) IsFlagEnabled(flag int) bool { return g.data.Flag(flag) }

// IsOccupied returns whether a build tile is occupied by a unit.
func (g *Game) IsOccupied(tileX, tileY int) bool {
	return g.data.IsOccupied(tileX, tileY)
}

// --- Additional Commands ---

// RestartGame restarts the current match.
func (g *Game) RestartGame() {
	g.data.AddCommand(int32(CommandTypeRestartGame), 0, 0)
}

// SetMap changes the map for the next game.
func (g *Game) SetMap(filename string) {
	idx := g.data.AddString(filename)
	g.data.AddCommand(int32(CommandTypeSetMap), idx, 0)
}

// SetAlliance changes alliance status with another player.
func (g *Game) SetAlliance(playerID int, allied, alliedVictory bool) {
	v1 := int32(playerID)
	v2 := int32(0)
	if allied {
		v2 = 1
	}
	if alliedVictory {
		v2 = 2
	}
	g.data.AddCommand(int32(CommandTypeSetAllies), v1, v2)
}

// SetVision shares or unshares vision with another player.
func (g *Game) SetVision(playerID int, enabled bool) {
	v := int32(0)
	if enabled {
		v = 1
	}
	g.data.AddCommand(int32(CommandTypeSetVision), int32(playerID), v)
}

// SetCommandOptimizationLevel sets the command optimization level.
func (g *Game) SetCommandOptimizationLevel(level int) {
	g.data.AddCommand(int32(CommandTypeSetCommandOptimizerLevel), int32(level), 0)
}

// SetRevealAll reveals or unreveals the entire map.
func (g *Game) SetRevealAll(reveal bool) {
	v := int32(0)
	if reveal {
		v = 1
	}
	g.data.AddCommand(int32(CommandTypeSetRevealAll), v, 0)
}

// SendTextEx sends a text message, optionally only to allies.
func (g *Game) SendTextEx(toAllies bool, text string) {
	idx := g.data.AddString(text)
	v2 := int32(0)
	if toAllies {
		v2 = 1
	}
	g.data.AddCommand(int32(CommandTypeSendText), idx, v2)
}

// --- Player Collections ---

// Allies returns all allied players.
func (g *Game) Allies() []*Player {
	self := g.Self()
	if self == nil {
		return nil
	}
	count := g.data.PlayerCount()
	allies := make([]*Player, 0)
	for i := 0; i < count; i++ {
		pd := g.data.Player(i)
		if i == self.index || pd.IsNeutral() {
			continue
		}
		if pd.IsAlly(self.index) {
			allies = append(allies, &Player{data: pd, game: g, index: i})
		}
	}
	return allies
}

// Enemies returns all enemy players.
func (g *Game) Enemies() []*Player {
	self := g.Self()
	if self == nil {
		return nil
	}
	count := g.data.PlayerCount()
	enemies := make([]*Player, 0)
	for i := 0; i < count; i++ {
		pd := g.data.Player(i)
		if i == self.index || pd.IsNeutral() {
			continue
		}
		if pd.IsEnemy(self.index) {
			enemies = append(enemies, &Player{data: pd, game: g, index: i})
		}
	}
	return enemies
}

// Observers returns all observer players.
func (g *Game) Observers() []*Player {
	count := g.data.PlayerCount()
	observers := make([]*Player, 0)
	for i := 0; i < count; i++ {
		pd := g.data.Player(i)
		if pd.TypeID() == int32(PlayerTypeObserver) {
			observers = append(observers, &Player{data: pd, game: g, index: i})
		}
	}
	return observers
}

// --- Forces ---

// GetForces returns all forces (team groupings) in the game.
func (g *Game) GetForces() []*Force {
	count := g.data.ForceCount()
	forces := make([]*Force, count)
	for i := 0; i < count; i++ {
		forces[i] = &Force{data: g.data.Force(i), game: g, index: i}
	}
	return forces
}

// GetForce returns a force by its index.
func (g *Game) GetForce(index int) *Force {
	if index < 0 || index >= g.data.ForceCount() {
		return nil
	}
	return &Force{data: g.data.Force(index), game: g, index: index}
}

// --- Region Lookup ---

// GetRegion returns a region by its ID.
func (g *Game) GetRegion(id int) *Region {
	if id < 0 || id >= g.data.RegionCount() {
		return nil
	}
	return &Region{data: g.data.Region(id), game: g}
}

// GetRegionAt returns the region at a pixel position.
func (g *Game) GetRegionAt(x, y int) *Region {
	tileX := x / 32
	tileY := y / 32
	if tileX < 0 || tileX >= int(g.data.MapWidth()) || tileY < 0 || tileY >= int(g.data.MapHeight()) {
		return nil
	}
	regionID := int(g.data.MapTileRegionID(tileX, tileY))
	if regionID >= 5000 {
		// Split tile: look up from split tile data
		splitIdx := regionID - 5000
		miniTileMask := g.data.MapSplitTilesMiniTileMask(splitIdx)
		// Determine which mini-tile within the build tile
		miniX := (x % 32) / 8
		miniY := (y % 32) / 8
		bit := uint16(1) << uint(miniX+miniY*4)
		if miniTileMask&bit != 0 {
			regionID = int(g.data.MapSplitTilesRegion2(splitIdx))
		} else {
			regionID = int(g.data.MapSplitTilesRegion1(splitIdx))
		}
	}
	if regionID < 0 || regionID >= g.data.RegionCount() {
		return nil
	}
	return &Region{data: g.data.Region(regionID), game: g}
}

// --- Unit Collections ---

// GetSelectedUnits returns the units currently selected by the player.
func (g *Game) GetSelectedUnits() []*Unit {
	count := g.data.SelectedUnitCount()
	units := make([]*Unit, 0, count)
	for i := 0; i < count; i++ {
		idx := int(g.data.SelectedUnit(i))
		if idx < 0 || idx >= shm.MaxUnits {
			continue
		}
		units = append(units, &Unit{data: g.data.Unit(idx), game: g, index: idx})
	}
	return units
}

// GetMinerals returns all visible mineral field units.
func (g *Game) GetMinerals() []*Unit {
	allUnits := g.GetAllUnits()
	minerals := make([]*Unit, 0)
	for _, u := range allUnits {
		if u.GetType().IsMineralField() {
			minerals = append(minerals, u)
		}
	}
	return minerals
}

// GetGeysers returns all visible vespene geyser units.
func (g *Game) GetGeysers() []*Unit {
	allUnits := g.GetAllUnits()
	geysers := make([]*Unit, 0)
	for _, u := range allUnits {
		t := u.GetType()
		if t == UnitTypeResourceVespeneGeyser {
			geysers = append(geysers, u)
		}
	}
	return geysers
}

// GetNeutralUnits returns all visible units owned by the neutral player.
func (g *Game) GetNeutralUnits() []*Unit {
	neutralIdx := int(g.data.NeutralIndex())
	allUnits := g.GetAllUnits()
	neutrals := make([]*Unit, 0)
	for _, u := range allUnits {
		if int(u.data.PlayerIndex()) == neutralIdx {
			neutrals = append(neutrals, u)
		}
	}
	return neutrals
}

// GetStaticMinerals returns mineral units that existed at game start.
// Must be called after OnStart (uses InitialUnitCount).
func (g *Game) GetStaticMinerals() []*Unit {
	count := g.data.InitialUnitCount()
	minerals := make([]*Unit, 0)
	for i := 0; i < count; i++ {
		ud := g.data.Unit(i)
		t := UnitType(ud.TypeID())
		if t.IsMineralField() {
			minerals = append(minerals, &Unit{data: ud, game: g, index: i})
		}
	}
	return minerals
}

// GetStaticGeysers returns geyser units that existed at game start.
// Must be called after OnStart (uses InitialUnitCount).
func (g *Game) GetStaticGeysers() []*Unit {
	count := g.data.InitialUnitCount()
	geysers := make([]*Unit, 0)
	for i := 0; i < count; i++ {
		ud := g.data.Unit(i)
		t := UnitType(ud.TypeID())
		if t == UnitTypeResourceVespeneGeyser {
			geysers = append(geysers, &Unit{data: ud, game: g, index: i})
		}
	}
	return geysers
}

// GetStaticNeutralUnits returns all neutral units that existed at game start.
// Must be called after OnStart (uses InitialUnitCount).
func (g *Game) GetStaticNeutralUnits() []*Unit {
	neutralIdx := int(g.data.NeutralIndex())
	count := g.data.InitialUnitCount()
	neutrals := make([]*Unit, 0)
	for i := 0; i < count; i++ {
		ud := g.data.Unit(i)
		if int(ud.PlayerIndex()) == neutralIdx {
			neutrals = append(neutrals, &Unit{data: ud, game: g, index: i})
		}
	}
	return neutrals
}

// --- Spatial Unit Queries ---

// GetUnitsOnTile returns all visible units on a specific build tile.
func (g *Game) GetUnitsOnTile(tileX, tileY int) []*Unit {
	left := int32(tileX * 32)
	top := int32(tileY * 32)
	right := left + 32
	bottom := top + 32
	allUnits := g.GetAllUnits()
	result := make([]*Unit, 0)
	for _, u := range allUnits {
		pos := u.GetPosition()
		if pos.X >= left && pos.X < right && pos.Y >= top && pos.Y < bottom {
			result = append(result, u)
		}
	}
	return result
}

// GetUnitsInRectangle returns all visible units within a pixel bounding box.
func (g *Game) GetUnitsInRectangle(left, top, right, bottom int) []*Unit {
	allUnits := g.GetAllUnits()
	result := make([]*Unit, 0)
	for _, u := range allUnits {
		pos := u.GetPosition()
		if int(pos.X) >= left && int(pos.X) <= right &&
			int(pos.Y) >= top && int(pos.Y) <= bottom {
			result = append(result, u)
		}
	}
	return result
}

// getUnitsInRadiusFiltered returns visible units within a radius, with optional filter.
func (g *Game) getUnitsInRadiusFiltered(x, y, radius int, filter func(*Unit) bool) []*Unit {
	r2 := int64(radius) * int64(radius)
	allUnits := g.GetAllUnits()
	result := make([]*Unit, 0)
	for _, u := range allUnits {
		if filter != nil && !filter(u) {
			continue
		}
		pos := u.GetPosition()
		dx := int64(pos.X) - int64(x)
		dy := int64(pos.Y) - int64(y)
		if dx*dx+dy*dy <= r2 {
			result = append(result, u)
		}
	}
	return result
}

// GetUnitsInRadius returns all visible units within a pixel radius from a point.
func (g *Game) GetUnitsInRadius(x, y, radius int) []*Unit {
	r2 := int64(radius) * int64(radius)
	allUnits := g.GetAllUnits()
	result := make([]*Unit, 0)
	for _, u := range allUnits {
		pos := u.GetPosition()
		dx := int64(pos.X) - int64(x)
		dy := int64(pos.Y) - int64(y)
		if dx*dx+dy*dy <= r2 {
			result = append(result, u)
		}
	}
	return result
}

// GetClosestUnit returns the closest visible unit to a pixel position
// that matches the optional filter. Pass nil for no filter.
func (g *Game) GetClosestUnit(x, y int, filter func(*Unit) bool) *Unit {
	allUnits := g.GetAllUnits()
	var closest *Unit
	bestDist := int64(1<<62 - 1)
	for _, u := range allUnits {
		if filter != nil && !filter(u) {
			continue
		}
		pos := u.GetPosition()
		dx := int64(pos.X) - int64(x)
		dy := int64(pos.Y) - int64(y)
		dist := dx*dx + dy*dy
		if dist < bestDist {
			bestDist = dist
			closest = u
		}
	}
	return closest
}

// --- Path Queries ---

// HasPath returns whether there is a ground path between two positions,
// based on BWAPI region island connectivity.
func (g *Game) HasPath(source, dest Position) bool {
	r1 := g.GetRegionAt(int(source.X), int(source.Y))
	r2 := g.GetRegionAt(int(dest.X), int(dest.Y))
	if r1 == nil || r2 == nil {
		return false
	}
	if !r1.IsAccessible() || !r2.IsAccessible() {
		return false
	}
	return r1.IslandID() == r2.IslandID()
}

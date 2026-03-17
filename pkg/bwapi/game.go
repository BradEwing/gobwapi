package bwapi

import (
	"github.com/bradewing/gobwapi/internal/shm"
)

// initialUnitState stores a snapshot of a unit's state at game start.
type initialUnitState struct {
	unitType  UnitType
	position  Position
	hitPoints int
	resources int
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
			unitType:  UnitType(ud.TypeID()),
			position:  Position{X: ud.PositionX(), Y: ud.PositionY()},
			hitPoints: int(ud.HitPoints()),
			resources: int(ud.Resources()),
		}
	}
}

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

// GetRegions returns all pathfinding regions.
func (g *Game) GetRegions() []*Region {
	count := g.data.RegionCount()
	regions := make([]*Region, count)
	for i := 0; i < count; i++ {
		regions[i] = &Region{data: g.data.Region(i), game: g}
	}
	return regions
}

// GetNukeDots returns positions of detected nuclear strikes.
func (g *Game) GetNukeDots() []Position {
	count := g.data.NukeDotCount()
	dots := make([]Position, count)
	for i := 0; i < count; i++ {
		dots[i] = Position{X: g.data.NukeDotX(i), Y: g.data.NukeDotY(i)}
	}
	return dots
}

func (g *Game) MouseX() int  { return int(g.data.MouseX()) }
func (g *Game) MouseY() int  { return int(g.data.MouseY()) }
func (g *Game) ScreenX() int { return int(g.data.ScreenX()) }
func (g *Game) ScreenY() int { return int(g.data.ScreenY()) }

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

func (g *Game) Revision() int        { return int(g.data.Revision()) }
func (g *Game) ClientVersion() int   { return int(g.data.ClientVersion()) }
func (g *Game) IsDebug() bool        { return g.data.IsDebug() }
func (g *Game) InstanceNumber() int  { return int(g.data.InstanceID()) }
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
// BWAPI encodes: 0=unallied, 1=allied, 2=allied+shared victory.
func (g *Game) SetAlliance(playerID int, allied, alliedVictory bool) {
	v2 := int32(0)
	if allied && alliedVictory {
		v2 = 2
	} else if allied {
		v2 = 1
	}
	g.data.AddCommand(int32(CommandTypeSetAllies), int32(playerID), v2)
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
	return g.getUnitsInRadiusFiltered(x, y, radius, nil)
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

// CanBuildHere checks whether a building of the given type can be placed at
// the given tile position. If builder is non-nil, its position is excluded
// from the occupation check. If checkExplored is true, all tiles must be explored.
func (g *Game) CanBuildHere(pos TilePosition, unitType UnitType, builder *Unit, checkExplored bool) bool {
	tw := unitType.TileWidth()
	th := unitType.TileHeight()
	if tw == 0 || th == 0 {
		return false
	}

	tx := int(pos.X)
	ty := int(pos.Y)
	mapW := int(g.data.MapWidth())
	mapH := int(g.data.MapHeight())

	if tx < 0 || ty < 0 || tx+tw > mapW || ty+th > mapH {
		return false
	}

	if unitType.IsRefinery() {
		for _, geyser := range g.GetGeysers() {
			gtp := geyser.GetTilePosition()
			if int(gtp.X) == tx && int(gtp.Y) == ty {
				return true
			}
		}
		return false
	}

	needsCreep := unitType.RequiresCreep()
	for x := tx; x < tx+tw; x++ {
		for y := ty; y < ty+th; y++ {
			if !g.IsBuildable(x, y) {
				return false
			}
			if checkExplored && !g.IsExplored(x, y) {
				return false
			}
			if g.IsOccupied(x, y) {
				if builder != nil {
					btp := builder.GetTilePosition()
					bx, by := int(btp.X), int(btp.Y)
					if x == bx && y == by {
						continue
					}
				}
				return false
			}
			if needsCreep && !g.HasCreep(x, y) {
				return false
			}
			if !needsCreep && unitType.GetRace() != RaceZerg && g.HasCreep(x, y) {
				return false
			}
		}
	}

	if unitType.RequiresPsi() && !g.HasPowerForType(tx, ty, unitType) {
		return false
	}

	if unitType.IsResourceDepot() {
		for _, mineral := range g.GetStaticMinerals() {
			mtp := mineral.GetTilePosition()
			mx, my := int(mtp.X), int(mtp.Y)
			if mx >= tx-5 && mx <= tx+tw+2 && my >= ty-4 && my <= ty+th+2 {
				return false
			}
		}
		for _, geyser := range g.GetStaticGeysers() {
			gtp := geyser.GetTilePosition()
			gx, gy := int(gtp.X), int(gtp.Y)
			if gx >= tx-7 && gx <= tx+tw+4 && gy >= ty-5 && gy <= ty+th+2 {
				return false
			}
		}
	}

	return true
}

// GetBuildLocation finds a valid placement for the given building type near the
// specified tile position. It scans tiles in a spiral outward from near,
// checking CanBuildHere and ground path connectivity for each candidate.
// The creep parameter is accepted for API compatibility with BWAPI but is
// currently unused (creep requirements are enforced by CanBuildHere).
// If builder is non-nil, it is passed to CanBuildHere for collision exclusion.
// Returns the found position and true, or a zero TilePosition and false if
// unitType is not a building or no valid spot exists within maxRange tiles.
func (g *Game) GetBuildLocation(unitType UnitType, near TilePosition, maxRange int, creep bool, builder *Unit) (TilePosition, bool) {
	if !unitType.IsBuilding() {
		return TilePosition{}, false
	}
	if maxRange <= 0 {
		maxRange = 64
	}

	nearCenter := near.ToPosition()
	tw := unitType.TileWidth()
	th := unitType.TileHeight()

	var fallback TilePosition
	fallbackFound := false
	fallbackDist := int64(1<<62 - 1)

	for radius := int32(0); radius < int32(maxRange); radius++ {
		for dx := -radius; dx <= radius; dx++ {
			for dy := -radius; dy <= radius; dy++ {
				if dx != -radius && dx != radius && dy != -radius && dy != radius {
					continue
				}
				tp := TilePosition{X: near.X + dx, Y: near.Y + dy}
				if !g.CanBuildHere(tp, unitType, builder, false) {
					continue
				}
				// Check ground connectivity: building center must be path-reachable
				// from near's center.
				bldCenter := Position{
					X: tp.X*32 + int32(tw)*16,
					Y: tp.Y*32 + int32(th)*16,
				}
				if !g.HasPath(nearCenter, bldCenter) {
					continue
				}
				// Track as candidate; pick by approximate distance.
				cdx := int64(bldCenter.X) - int64(nearCenter.X)
				cdy := int64(bldCenter.Y) - int64(nearCenter.Y)
				dist := cdx*cdx + cdy*cdy
				if !fallbackFound || dist < fallbackDist {
					fallback = tp
					fallbackFound = true
					fallbackDist = dist
				}
				// First ring that has any valid tile wins — return immediately
				// since the spiral guarantees we've checked all tiles at this
				// Chebyshev distance.
			}
		}
		if fallbackFound {
			return fallback, true
		}
	}
	// Fallback: if a valid tile was found beyond maxRange during iteration
	// (not possible with current logic, but guard for future changes).
	if fallbackFound {
		return fallback, true
	}
	return TilePosition{}, false
}

// CanMake checks whether the self player can produce a unit of the given type.
// If builder is non-nil, checks that the builder is of the correct type.
func (g *Game) CanMake(unitType UnitType, builder *Unit) bool {
	self := g.Self()
	if self == nil {
		return false
	}

	if !self.IsUnitAvailable(unitType) {
		return false
	}

	builderType, _ := unitType.WhatBuilds()
	if builder != nil {
		if builder.GetType() != builderType {
			return false
		}
		if builderType.IsBuilding() && builderType.GetRace() == RaceZerg {
			if unitType == UnitTypeZergLarva {
				return true
			}
			if builder.GetLarva() == nil || len(builder.GetLarva()) == 0 {
				return false
			}
		}
	}

	if self.Minerals() < unitType.MineralPrice() {
		return false
	}
	if self.Gas() < unitType.GasPrice() {
		return false
	}

	supplyReq := unitType.SupplyRequired()
	if unitType.IsTwoUnitsInOneEgg() {
		supplyReq *= 2
	}
	if supplyReq > 0 {
		race := unitType.GetRace()
		supplyFree := builderType.SupplyRequired()
		if builderType.GetRace() != race {
			supplyFree = 0
		}
		if self.SupplyUsedForRace(race)-supplyFree+supplyReq > self.SupplyTotalForRace(race) {
			return false
		}
	}

	for reqType, reqCount := range unitType.RequiredUnits() {
		if self.CompletedUnitCount(reqType) < reqCount {
			return false
		}
	}

	reqTech := unitType.RequiredTech()
	if reqTech != TechTypeNone && !self.HasResearched(reqTech) {
		return false
	}

	return true
}

// CanResearch checks whether the self player can research the given technology.
// If unit is non-nil, checks that the unit is of the correct building type.
func (g *Game) CanResearch(tech TechType, unit *Unit) bool {
	self := g.Self()
	if self == nil {
		return false
	}

	// Check availability
	if !self.IsResearchAvailable(tech) {
		return false
	}

	// Already researched
	if self.HasResearched(tech) {
		return false
	}

	// Currently researching
	if self.IsResearching(tech) {
		return false
	}

	// Check unit type
	if unit != nil {
		if unit.GetType() != tech.WhatResearches() {
			return false
		}
	}

	// Check resources
	if self.Minerals() < tech.MineralPrice() {
		return false
	}
	if self.Gas() < tech.GasPrice() {
		return false
	}

	return true
}

// CanUpgrade checks whether the self player can perform the given upgrade.
// If unit is non-nil, checks that the unit is of the correct building type.
func (g *Game) CanUpgrade(upgrade UpgradeType, unit *Unit) bool {
	self := g.Self()
	if self == nil {
		return false
	}

	// Check current level
	currentLevel := self.UpgradeLevel(upgrade)
	if currentLevel >= upgrade.MaxRepeats() {
		return false
	}

	// Check already upgrading
	if self.IsUpgrading(upgrade) {
		return false
	}

	// Check unit type
	if unit != nil {
		if unit.GetType() != upgrade.WhatUpgrades() {
			return false
		}
	}

	// Check resources for next level
	nextLevel := currentLevel + 1
	if self.Minerals() < upgrade.MineralPrice(nextLevel) {
		return false
	}
	if self.Gas() < upgrade.GasPrice(nextLevel) {
		return false
	}

	// Check required building for this level
	reqBuilding := upgrade.WhatsRequired(nextLevel)
	if reqBuilding != UnitTypeNone {
		if self.CompletedUnitCount(reqBuilding) < 1 {
			return false
		}
	}

	return true
}

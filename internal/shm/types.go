package shm

/*
#include "gamedata.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"
import "unsafe"

// Type aliases for C structs. These map directly to BWAPI shared memory layout.
type GameData C.GameData
type GameTable C.GameTable
type GameInstance C.GameInstance
type UnitData C.UnitData
type PlayerData C.PlayerData
type BulletData C.BulletData
type RegionData C.RegionData
type ForceData C.ForceData
type Position C.Position
type UnitFinderEntry C.UnitFinder
type Event C.Event
type Command C.Command
type Shape C.Shape
type UnitCommand C.UnitCommand

// CastGameData casts raw memory to a GameData pointer.
func CastGameData(ptr unsafe.Pointer) *GameData {
	return (*GameData)(ptr)
}

// CastGameTable casts raw memory to a GameTable pointer.
func CastGameTable(ptr unsafe.Pointer) *GameTable {
	return (*GameTable)(ptr)
}

// --- Helper cast methods ---

func (gd *GameData) c() *C.GameData         { return (*C.GameData)(unsafe.Pointer(gd)) }
func (gt *GameTable) c() *C.GameTable        { return (*C.GameTable)(unsafe.Pointer(gt)) }
func (gi *GameInstance) c() *C.GameInstance   { return (*C.GameInstance)(unsafe.Pointer(gi)) }
func (ud *UnitData) c() *C.UnitData          { return (*C.UnitData)(unsafe.Pointer(ud)) }
func (pd *PlayerData) c() *C.PlayerData      { return (*C.PlayerData)(unsafe.Pointer(pd)) }
func (bd *BulletData) c() *C.BulletData      { return (*C.BulletData)(unsafe.Pointer(bd)) }
func (rd *RegionData) c() *C.RegionData      { return (*C.RegionData)(unsafe.Pointer(rd)) }
func (fd *ForceData) c() *C.ForceData        { return (*C.ForceData)(unsafe.Pointer(fd)) }

// cString converts a C char array (pointer + max length) to a Go string.
func cString(p *C.char, maxLen int) string {
	return C.GoStringN(p, C.int(maxLen))
}

// --- GameTable / GameInstance accessors ---

func (gt *GameTable) Instance(i int) *GameInstance {
	return (*GameInstance)(unsafe.Pointer(&gt.c().gameInstances[i]))
}

func (gi *GameInstance) ServerProcessID() uint32    { return uint32(gi.c().serverProcessID) }
func (gi *GameInstance) IsConnected() bool          { return gi.c().isConnected != 0 }
func (gi *GameInstance) LastKeepAliveTime() uint32   { return uint32(gi.c().lastKeepAliveTime) }
func (gi *GameInstance) SetIsConnected(v bool) {
	if v {
		gi.c().isConnected = 1
	} else {
		gi.c().isConnected = 0
	}
}

// --- GameData general accessors ---

func (gd *GameData) ClientVersion() int32          { return int32(gd.c().client_version) }
func (gd *GameData) Revision() int32               { return int32(gd.c().revision) }
func (gd *GameData) IsDebug() bool                 { return gd.c().isDebug != 0 }
func (gd *GameData) InstanceID() int32             { return int32(gd.c().instanceID) }
func (gd *GameData) BotAPMNoSelects() int32        { return int32(gd.c().botAPM_noselects) }
func (gd *GameData) BotAPMSelects() int32          { return int32(gd.c().botAPM_selects) }

// Forces
func (gd *GameData) ForceCount() int               { return int(gd.c().forceCount) }
func (gd *GameData) Force(i int) *ForceData {
	return (*ForceData)(unsafe.Pointer(&gd.c().forces[i]))
}

// Players
func (gd *GameData) PlayerCount() int              { return int(gd.c().playerCount) }
func (gd *GameData) Player(i int) *PlayerData {
	return (*PlayerData)(unsafe.Pointer(&gd.c().players[i]))
}

// Units
func (gd *GameData) InitialUnitCount() int         { return int(gd.c().initialUnitCount) }
func (gd *GameData) Unit(i int) *UnitData {
	return (*UnitData)(unsafe.Pointer(&gd.c().units[i]))
}
func (gd *GameData) UnitArrayEntry(i int) int32    { return int32(gd.c().unitArray[i]) }

// Bullets
func (gd *GameData) Bullet(i int) *BulletData {
	return (*BulletData)(unsafe.Pointer(&gd.c().bullets[i]))
}

// Nuke dots
func (gd *GameData) NukeDotCount() int             { return int(gd.c().nukeDotCount) }
func (gd *GameData) NukeDotX(i int) int32          { return int32(gd.c().nukeDots[i].x) }
func (gd *GameData) NukeDotY(i int) int32          { return int32(gd.c().nukeDots[i].y) }

// Game state
func (gd *GameData) GameType() int32               { return int32(gd.c().gameType) }
func (gd *GameData) Latency() int32                { return int32(gd.c().latency) }
func (gd *GameData) LatencyFrames() int32          { return int32(gd.c().latencyFrames) }
func (gd *GameData) LatencyTime() int32            { return int32(gd.c().latencyTime) }
func (gd *GameData) RemainingLatencyFrames() int32 { return int32(gd.c().remainingLatencyFrames) }
func (gd *GameData) RemainingLatencyTime() int32   { return int32(gd.c().remainingLatencyTime) }
func (gd *GameData) HasLatCom() bool               { return gd.c().hasLatCom != 0 }
func (gd *GameData) HasGUI() bool                  { return gd.c().hasGUI != 0 }
func (gd *GameData) ReplayFrameCount() int32       { return int32(gd.c().replayFrameCount) }
func (gd *GameData) RandomSeed() uint32            { return uint32(gd.c().randomSeed) }
func (gd *GameData) FrameCount() int32             { return int32(gd.c().frameCount) }
func (gd *GameData) ElapsedTime() int32            { return int32(gd.c().elapsedTime) }
func (gd *GameData) CountdownTimer() int32         { return int32(gd.c().countdownTimer) }
func (gd *GameData) FPS() int32                    { return int32(gd.c().fps) }
func (gd *GameData) AverageFPS() float64           { return float64(gd.c().averageFPS) }

// Input
func (gd *GameData) MouseX() int32                 { return int32(gd.c().mouseX) }
func (gd *GameData) MouseY() int32                 { return int32(gd.c().mouseY) }
func (gd *GameData) MouseState(button int) bool    { return gd.c().mouseState[button] != 0 }
func (gd *GameData) KeyState(key int) bool         { return gd.c().keyState[key] != 0 }
func (gd *GameData) ScreenX() int32                { return int32(gd.c().screenX) }
func (gd *GameData) ScreenY() int32                { return int32(gd.c().screenY) }
func (gd *GameData) Flag(i int) bool               { return gd.c().flags[i] != 0 }

// Map info
func (gd *GameData) MapWidth() int32               { return int32(gd.c().mapWidth) }
func (gd *GameData) MapHeight() int32              { return int32(gd.c().mapHeight) }
func (gd *GameData) MapFileName() string {
	return C.GoString(&gd.c().mapFileName[0])
}
func (gd *GameData) MapPathName() string {
	return C.GoString(&gd.c().mapPathName[0])
}
func (gd *GameData) MapName() string {
	return C.GoString(&gd.c().mapName[0])
}
func (gd *GameData) MapHash() string {
	return C.GoString(&gd.c().mapHash[0])
}

// Map tile data
func (gd *GameData) GetGroundHeight(x, y int) int32 {
	return int32(gd.c().getGroundHeight[x][y])
}
func (gd *GameData) IsWalkable(x, y int) bool {
	return gd.c().isWalkable[x][y] != 0
}
func (gd *GameData) IsBuildable(x, y int) bool {
	return gd.c().isBuildable[x][y] != 0
}
func (gd *GameData) IsVisible(x, y int) bool {
	return gd.c().isVisible[x][y] != 0
}
func (gd *GameData) IsExplored(x, y int) bool {
	return gd.c().isExplored[x][y] != 0
}
func (gd *GameData) HasCreep(x, y int) bool {
	return gd.c().hasCreep[x][y] != 0
}
func (gd *GameData) IsOccupied(x, y int) bool {
	return gd.c().isOccupied[x][y] != 0
}
func (gd *GameData) MapTileRegionID(x, y int) uint16 {
	return uint16(gd.c().mapTileRegionId[x][y])
}

// Regions
func (gd *GameData) RegionCount() int              { return int(gd.c().regionCount) }
func (gd *GameData) Region(i int) *RegionData {
	return (*RegionData)(unsafe.Pointer(&gd.c().regions[i]))
}

// Start locations
func (gd *GameData) StartLocationCount() int       { return int(gd.c().startLocationCount) }
func (gd *GameData) StartLocationX(i int) int32    { return int32(gd.c().startLocations[i].x) }
func (gd *GameData) StartLocationY(i int) int32    { return int32(gd.c().startLocations[i].y) }

// Game flags
func (gd *GameData) IsInGame() bool                { return gd.c().isInGame != 0 }
func (gd *GameData) IsMultiplayer() bool            { return gd.c().isMultiplayer != 0 }
func (gd *GameData) IsBattleNet() bool              { return gd.c().isBattleNet != 0 }
func (gd *GameData) IsPaused() bool                 { return gd.c().isPaused != 0 }
func (gd *GameData) IsReplay() bool                 { return gd.c().isReplay != 0 }

// Selected units
func (gd *GameData) SelectedUnitCount() int        { return int(gd.c().selectedUnitCount) }
func (gd *GameData) SelectedUnit(i int) int32      { return int32(gd.c().selectedUnits[i]) }

// Player references
func (gd *GameData) SelfIndex() int32              { return int32(gd.c().self) }
func (gd *GameData) EnemyIndex() int32             { return int32(gd.c().enemy) }
func (gd *GameData) NeutralIndex() int32           { return int32(gd.c().neutral) }

// Events
func (gd *GameData) EventCount() int               { return int(gd.c().eventCount) }
func (gd *GameData) EventType(i int) int32         { return int32(gd.c().events[i]._type) }
func (gd *GameData) EventV1(i int) int32           { return int32(gd.c().events[i].v1) }
func (gd *GameData) EventV2(i int) int32           { return int32(gd.c().events[i].v2) }

// Event strings
func (gd *GameData) EventStringCount() int         { return int(gd.c().eventStringCount) }
func (gd *GameData) EventString(i int) string {
	return C.GoString(&gd.c().eventStrings[i][0])
}

// General strings
func (gd *GameData) StringCount() int              { return int(gd.c().stringCount) }
func (gd *GameData) GetString(i int) string {
	return C.GoString(&gd.c().strings[i][0])
}

// Shape/command/unit command counts
func (gd *GameData) ShapeCount() int               { return int(gd.c().shapeCount) }
func (gd *GameData) CommandCount() int             { return int(gd.c().commandCount) }
func (gd *GameData) UnitCommandCount() int         { return int(gd.c().unitCommandCount) }

// Unit search
func (gd *GameData) UnitSearchSize() int           { return int(gd.c().unitSearchSize) }
func (gd *GameData) XUnitSearchIndex(i int) int32  { return int32(gd.c().xUnitSearch[i].unitIndex) }
func (gd *GameData) XUnitSearchValue(i int) int32  { return int32(gd.c().xUnitSearch[i].searchValue) }
func (gd *GameData) YUnitSearchIndex(i int) int32  { return int32(gd.c().yUnitSearch[i].unitIndex) }
func (gd *GameData) YUnitSearchValue(i int) int32  { return int32(gd.c().yUnitSearch[i].searchValue) }

// --- GameData write methods (for staging commands during a frame) ---

// AddCommand adds a client command to the command buffer.
func (gd *GameData) AddCommand(cmdType, value1, value2 int32) {
	c := gd.c()
	idx := c.commandCount
	c.commands[idx]._type = C.int32_t(cmdType)
	c.commands[idx].value1 = C.int32_t(value1)
	c.commands[idx].value2 = C.int32_t(value2)
	c.commandCount++
}

// AddUnitCommand adds a unit command to the unit command buffer.
func (gd *GameData) AddUnitCommand(cmdType, unitIndex, targetIndex, x, y, extra int32) {
	c := gd.c()
	idx := c.unitCommandCount
	c.unitCommands[idx]._type = C.int32_t(cmdType)
	c.unitCommands[idx].unitIndex = C.int32_t(unitIndex)
	c.unitCommands[idx].targetIndex = C.int32_t(targetIndex)
	c.unitCommands[idx].x = C.int32_t(x)
	c.unitCommands[idx].y = C.int32_t(y)
	c.unitCommands[idx].extra = C.int32_t(extra)
	c.unitCommandCount++
}

// AddShape adds a debug drawing shape to the shape buffer.
func (gd *GameData) AddShape(shapeType, ctype, x1, y1, x2, y2, extra1, extra2, color int32, isSolid bool) {
	c := gd.c()
	idx := c.shapeCount
	c.shapes[idx]._type = C.int32_t(shapeType)
	c.shapes[idx].ctype = C.int32_t(ctype)
	c.shapes[idx].x1 = C.int32_t(x1)
	c.shapes[idx].y1 = C.int32_t(y1)
	c.shapes[idx].x2 = C.int32_t(x2)
	c.shapes[idx].y2 = C.int32_t(y2)
	c.shapes[idx].extra1 = C.int32_t(extra1)
	c.shapes[idx].extra2 = C.int32_t(extra2)
	c.shapes[idx].color = C.int32_t(color)
	if isSolid {
		c.shapes[idx].isSolid = 1
	} else {
		c.shapes[idx].isSolid = 0
	}
	c.shapeCount++
}

// AddString adds a string to the string table and returns its index.
func (gd *GameData) AddString(s string) int32 {
	c := gd.c()
	idx := int(c.stringCount)
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	// Copy into the fixed-size buffer, leaving room for null terminator
	C.strncpy(&c.strings[idx][0], cs, C.size_t(BW_STRING_SIZE-1))
	c.strings[idx][BW_STRING_SIZE-1] = 0
	c.stringCount++
	return int32(idx)
}

// ResetCommands clears all command buffers for a new frame.
func (gd *GameData) ResetCommands() {
	c := gd.c()
	c.commandCount = 0
	c.unitCommandCount = 0
	c.shapeCount = 0
	c.stringCount = 0
}

// --- UnitData accessors ---

func (ud *UnitData) ClearanceLevel() int32         { return int32(ud.c().clearanceLevel) }
func (ud *UnitData) ID() int32                     { return int32(ud.c().id) }
func (ud *UnitData) PlayerIndex() int32            { return int32(ud.c().player) }
func (ud *UnitData) TypeID() int32                 { return int32(ud.c()._type) }
func (ud *UnitData) PositionX() int32              { return int32(ud.c().positionX) }
func (ud *UnitData) PositionY() int32              { return int32(ud.c().positionY) }
func (ud *UnitData) Angle() float64                { return float64(ud.c().angle) }
func (ud *UnitData) VelocityX() float64            { return float64(ud.c().velocityX) }
func (ud *UnitData) VelocityY() float64            { return float64(ud.c().velocityY) }
func (ud *UnitData) HitPoints() int32              { return int32(ud.c().hitPoints) }
func (ud *UnitData) LastHitPoints() int32          { return int32(ud.c().lastHitPoints) }
func (ud *UnitData) Shields() int32                { return int32(ud.c().shields) }
func (ud *UnitData) Energy() int32                 { return int32(ud.c().energy) }
func (ud *UnitData) Resources() int32              { return int32(ud.c().resources) }
func (ud *UnitData) ResourceGroup() int32          { return int32(ud.c().resourceGroup) }
func (ud *UnitData) KillCount() int32              { return int32(ud.c().killCount) }
func (ud *UnitData) AcidSporeCount() int32         { return int32(ud.c().acidSporeCount) }
func (ud *UnitData) ScarabCount() int32            { return int32(ud.c().scarabCount) }
func (ud *UnitData) InterceptorCount() int32       { return int32(ud.c().interceptorCount) }
func (ud *UnitData) SpiderMineCount() int32        { return int32(ud.c().spiderMineCount) }
func (ud *UnitData) GroundWeaponCooldown() int32   { return int32(ud.c().groundWeaponCooldown) }
func (ud *UnitData) AirWeaponCooldown() int32      { return int32(ud.c().airWeaponCooldown) }
func (ud *UnitData) SpellCooldown() int32          { return int32(ud.c().spellCooldown) }
func (ud *UnitData) DefenseMatrixPoints() int32    { return int32(ud.c().defenseMatrixPoints) }
func (ud *UnitData) DefenseMatrixTimer() int32     { return int32(ud.c().defenseMatrixTimer) }
func (ud *UnitData) EnsnareTimer() int32           { return int32(ud.c().ensnareTimer) }
func (ud *UnitData) IrradiateTimer() int32         { return int32(ud.c().irradiateTimer) }
func (ud *UnitData) LockdownTimer() int32          { return int32(ud.c().lockdownTimer) }
func (ud *UnitData) MaelstromTimer() int32         { return int32(ud.c().maelstromTimer) }
func (ud *UnitData) OrderTimer() int32             { return int32(ud.c().orderTimer) }
func (ud *UnitData) PlagueTimer() int32            { return int32(ud.c().plagueTimer) }
func (ud *UnitData) RemoveTimer() int32            { return int32(ud.c().removeTimer) }
func (ud *UnitData) StasisTimer() int32            { return int32(ud.c().stasisTimer) }
func (ud *UnitData) StimTimer() int32              { return int32(ud.c().stimTimer) }
func (ud *UnitData) BuildTypeID() int32            { return int32(ud.c().buildType) }
func (ud *UnitData) TrainingQueueCount() int32     { return int32(ud.c().trainingQueueCount) }
func (ud *UnitData) TrainingQueueEntry(i int) int32 { return int32(ud.c().trainingQueue[i]) }
func (ud *UnitData) TechID() int32                 { return int32(ud.c().tech) }
func (ud *UnitData) UpgradeID() int32              { return int32(ud.c().upgrade) }
func (ud *UnitData) RemainingBuildTime() int32     { return int32(ud.c().remainingBuildTime) }
func (ud *UnitData) RemainingTrainTime() int32     { return int32(ud.c().remainingTrainTime) }
func (ud *UnitData) RemainingResearchTime() int32  { return int32(ud.c().remainingResearchTime) }
func (ud *UnitData) RemainingUpgradeTime() int32   { return int32(ud.c().remainingUpgradeTime) }
func (ud *UnitData) BuildUnitIndex() int32         { return int32(ud.c().buildUnit) }
func (ud *UnitData) TargetIndex() int32            { return int32(ud.c().target) }
func (ud *UnitData) TargetPositionX() int32        { return int32(ud.c().targetPositionX) }
func (ud *UnitData) TargetPositionY() int32        { return int32(ud.c().targetPositionY) }
func (ud *UnitData) OrderID() int32                { return int32(ud.c().order) }
func (ud *UnitData) OrderTargetIndex() int32       { return int32(ud.c().orderTarget) }
func (ud *UnitData) OrderTargetPositionX() int32   { return int32(ud.c().orderTargetPositionX) }
func (ud *UnitData) OrderTargetPositionY() int32   { return int32(ud.c().orderTargetPositionY) }
func (ud *UnitData) SecondaryOrderID() int32       { return int32(ud.c().secondaryOrder) }
func (ud *UnitData) RallyPositionX() int32         { return int32(ud.c().rallyPositionX) }
func (ud *UnitData) RallyPositionY() int32         { return int32(ud.c().rallyPositionY) }
func (ud *UnitData) RallyUnitIndex() int32         { return int32(ud.c().rallyUnit) }
func (ud *UnitData) AddonIndex() int32             { return int32(ud.c().addon) }
func (ud *UnitData) NydusExitIndex() int32         { return int32(ud.c().nydusExit) }
func (ud *UnitData) PowerUpIndex() int32           { return int32(ud.c().powerUp) }
func (ud *UnitData) TransportIndex() int32         { return int32(ud.c().transport) }
func (ud *UnitData) CarrierIndex() int32           { return int32(ud.c().carrier) }
func (ud *UnitData) HatcheryIndex() int32          { return int32(ud.c().hatchery) }
func (ud *UnitData) Exists() bool                  { return ud.c().exists != 0 }
func (ud *UnitData) HasNuke() bool                 { return ud.c().hasNuke != 0 }
func (ud *UnitData) IsAccelerating() bool          { return ud.c().isAccelerating != 0 }
func (ud *UnitData) IsAttacking() bool             { return ud.c().isAttacking != 0 }
func (ud *UnitData) IsAttackFrame() bool           { return ud.c().isAttackFrame != 0 }
func (ud *UnitData) IsBeingGathered() bool         { return ud.c().isBeingGathered != 0 }
func (ud *UnitData) IsBlind() bool                 { return ud.c().isBlind != 0 }
func (ud *UnitData) IsBraking() bool               { return ud.c().isBraking != 0 }
func (ud *UnitData) IsBurrowed() bool              { return ud.c().isBurrowed != 0 }
func (ud *UnitData) CarryResourceType() int32      { return int32(ud.c().carryResourceType) }
func (ud *UnitData) IsCloaked() bool               { return ud.c().isCloaked != 0 }
func (ud *UnitData) IsCompleted() bool             { return ud.c().isCompleted != 0 }
func (ud *UnitData) IsConstructing() bool          { return ud.c().isConstructing != 0 }
func (ud *UnitData) IsDetected() bool              { return ud.c().isDetected != 0 }
func (ud *UnitData) IsGathering() bool             { return ud.c().isGathering != 0 }
func (ud *UnitData) IsHallucination() bool         { return ud.c().isHallucination != 0 }
func (ud *UnitData) IsIdle() bool                  { return ud.c().isIdle != 0 }
func (ud *UnitData) IsInterruptible() bool         { return ud.c().isInterruptible != 0 }
func (ud *UnitData) IsInvincible() bool            { return ud.c().isInvincible != 0 }
func (ud *UnitData) IsLifted() bool                { return ud.c().isLifted != 0 }
func (ud *UnitData) IsMorphing() bool              { return ud.c().isMorphing != 0 }
func (ud *UnitData) IsMoving() bool                { return ud.c().isMoving != 0 }
func (ud *UnitData) IsParasited() bool             { return ud.c().isParasited != 0 }
func (ud *UnitData) IsSelected() bool              { return ud.c().isSelected != 0 }
func (ud *UnitData) IsStartingAttack() bool        { return ud.c().isStartingAttack != 0 }
func (ud *UnitData) IsStuck() bool                 { return ud.c().isStuck != 0 }
func (ud *UnitData) IsTraining() bool              { return ud.c().isTraining != 0 }
func (ud *UnitData) IsUnderStorm() bool            { return ud.c().isUnderStorm != 0 }
func (ud *UnitData) IsUnderDarkSwarm() bool        { return ud.c().isUnderDarkSwarm != 0 }
func (ud *UnitData) IsUnderDWeb() bool             { return ud.c().isUnderDWeb != 0 }
func (ud *UnitData) IsPowered() bool               { return ud.c().isPowered != 0 }
func (ud *UnitData) IsVisibleTo(playerIndex int) bool {
	return ud.c().isVisible[playerIndex] != 0
}
func (ud *UnitData) ButtonSet() int32              { return int32(ud.c().buttonset) }
func (ud *UnitData) LastAttackerPlayer() int32     { return int32(ud.c().lastAttackerPlayer) }
func (ud *UnitData) RecentlyAttacked() bool        { return ud.c().recentlyAttacked != 0 }
func (ud *UnitData) ReplayID() int32               { return int32(ud.c().replayID) }

// --- PlayerData accessors ---

func (pd *PlayerData) Name() string {
	return C.GoString(&pd.c().name[0])
}
func (pd *PlayerData) RaceID() int32               { return int32(pd.c().race) }
func (pd *PlayerData) TypeID() int32               { return int32(pd.c()._type) }
func (pd *PlayerData) ForceIndex() int32           { return int32(pd.c().force) }
func (pd *PlayerData) IsAlly(playerIndex int) bool  { return pd.c().isAlly[playerIndex] != 0 }
func (pd *PlayerData) IsEnemy(playerIndex int) bool { return pd.c().isEnemy[playerIndex] != 0 }
func (pd *PlayerData) IsNeutral() bool             { return pd.c().isNeutral != 0 }
func (pd *PlayerData) StartLocationX() int32       { return int32(pd.c().startLocationX) }
func (pd *PlayerData) StartLocationY() int32       { return int32(pd.c().startLocationY) }
func (pd *PlayerData) IsVictorious() bool          { return pd.c().isVictorious != 0 }
func (pd *PlayerData) IsDefeated() bool            { return pd.c().isDefeated != 0 }
func (pd *PlayerData) LeftGame() bool              { return pd.c().leftGame != 0 }
func (pd *PlayerData) IsParticipating() bool       { return pd.c().isParticipating != 0 }
func (pd *PlayerData) Minerals() int32             { return int32(pd.c().minerals) }
func (pd *PlayerData) Gas() int32                  { return int32(pd.c().gas) }
func (pd *PlayerData) GatheredMinerals() int32     { return int32(pd.c().gatheredMinerals) }
func (pd *PlayerData) GatheredGas() int32          { return int32(pd.c().gatheredGas) }
func (pd *PlayerData) RepairedMinerals() int32     { return int32(pd.c().repairedMinerals) }
func (pd *PlayerData) RepairedGas() int32          { return int32(pd.c().repairedGas) }
func (pd *PlayerData) RefundedMinerals() int32     { return int32(pd.c().refundedMinerals) }
func (pd *PlayerData) RefundedGas() int32          { return int32(pd.c().refundedGas) }
func (pd *PlayerData) SupplyTotal(raceIndex int) int32  { return int32(pd.c().supplyTotal[raceIndex]) }
func (pd *PlayerData) SupplyUsed(raceIndex int) int32   { return int32(pd.c().supplyUsed[raceIndex]) }
func (pd *PlayerData) AllUnitCount(unitTypeID int) int32 {
	return int32(pd.c().allUnitCount[unitTypeID])
}
func (pd *PlayerData) VisibleUnitCount(unitTypeID int) int32 {
	return int32(pd.c().visibleUnitCount[unitTypeID])
}
func (pd *PlayerData) CompletedUnitCount(unitTypeID int) int32 {
	return int32(pd.c().completedUnitCount[unitTypeID])
}
func (pd *PlayerData) DeadUnitCount(unitTypeID int) int32 {
	return int32(pd.c().deadUnitCount[unitTypeID])
}
func (pd *PlayerData) KilledUnitCount(unitTypeID int) int32 {
	return int32(pd.c().killedUnitCount[unitTypeID])
}
func (pd *PlayerData) UpgradeLevel(upgradeTypeID int) int32 {
	return int32(pd.c().upgradeLevel[upgradeTypeID])
}
func (pd *PlayerData) HasResearched(techTypeID int) bool {
	return pd.c().hasResearched[techTypeID] != 0
}
func (pd *PlayerData) IsResearching(techTypeID int) bool {
	return pd.c().isResearching[techTypeID] != 0
}
func (pd *PlayerData) IsUpgrading(upgradeTypeID int) bool {
	return pd.c().isUpgrading[upgradeTypeID] != 0
}
func (pd *PlayerData) Color() int32                { return int32(pd.c().color) }
func (pd *PlayerData) TotalUnitScore() int32       { return int32(pd.c().totalUnitScore) }
func (pd *PlayerData) TotalKillScore() int32       { return int32(pd.c().totalKillScore) }
func (pd *PlayerData) TotalBuildingScore() int32   { return int32(pd.c().totalBuildingScore) }
func (pd *PlayerData) TotalRazingScore() int32     { return int32(pd.c().totalRazingScore) }
func (pd *PlayerData) CustomScore() int32          { return int32(pd.c().customScore) }
func (pd *PlayerData) MaxUpgradeLevel(upgradeTypeID int) int32 {
	return int32(pd.c().maxUpgradeLevel[upgradeTypeID])
}
func (pd *PlayerData) IsResearchAvailable(techTypeID int) bool {
	return pd.c().isResearchAvailable[techTypeID] != 0
}
func (pd *PlayerData) IsUnitAvailable(unitTypeID int) bool {
	return pd.c().isUnitAvailable[unitTypeID] != 0
}

// --- BulletData accessors ---

func (bd *BulletData) ID() int32                   { return int32(bd.c().id) }
func (bd *BulletData) PlayerIndex() int32          { return int32(bd.c().player) }
func (bd *BulletData) TypeID() int32               { return int32(bd.c()._type) }
func (bd *BulletData) SourceIndex() int32          { return int32(bd.c().source) }
func (bd *BulletData) PositionX() int32            { return int32(bd.c().positionX) }
func (bd *BulletData) PositionY() int32            { return int32(bd.c().positionY) }
func (bd *BulletData) Angle() float64              { return float64(bd.c().angle) }
func (bd *BulletData) VelocityX() float64          { return float64(bd.c().velocityX) }
func (bd *BulletData) VelocityY() float64          { return float64(bd.c().velocityY) }
func (bd *BulletData) TargetIndex() int32          { return int32(bd.c().target) }
func (bd *BulletData) TargetPositionX() int32      { return int32(bd.c().targetPositionX) }
func (bd *BulletData) TargetPositionY() int32      { return int32(bd.c().targetPositionY) }
func (bd *BulletData) RemoveTimer() int32          { return int32(bd.c().removeTimer) }
func (bd *BulletData) Exists() bool                { return bd.c().exists != 0 }
func (bd *BulletData) IsVisibleTo(playerIndex int) bool {
	return bd.c().isVisible[playerIndex] != 0
}

// --- RegionData accessors ---

func (rd *RegionData) ID() int32                   { return int32(rd.c().id) }
func (rd *RegionData) IslandID() int32             { return int32(rd.c().islandID) }
func (rd *RegionData) CenterX() int32              { return int32(rd.c().center_x) }
func (rd *RegionData) CenterY() int32              { return int32(rd.c().center_y) }
func (rd *RegionData) Priority() int32             { return int32(rd.c().priority) }
func (rd *RegionData) LeftMost() int32             { return int32(rd.c().leftMost) }
func (rd *RegionData) RightMost() int32            { return int32(rd.c().rightMost) }
func (rd *RegionData) TopMost() int32              { return int32(rd.c().topMost) }
func (rd *RegionData) BottomMost() int32           { return int32(rd.c().bottomMost) }
func (rd *RegionData) NeighborCount() int32        { return int32(rd.c().neighborCount) }
func (rd *RegionData) Neighbor(i int) int32        { return int32(rd.c().neighbors[i]) }
func (rd *RegionData) IsAccessible() bool          { return rd.c().isAccessible != 0 }
func (rd *RegionData) IsHigherGround() bool        { return rd.c().isHigherGround != 0 }

// --- ForceData accessors ---

func (fd *ForceData) Name() string {
	return C.GoString(&fd.c().name[0])
}

// BW_STRING_SIZE for use in AddString.
const BW_STRING_SIZE = C.BW_STRING_SIZE

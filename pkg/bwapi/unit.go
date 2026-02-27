package bwapi

import (
	"github.com/bradewing/gobwapi/internal/shm"
)

// Unit wraps a BWAPI unit's shared memory data.
type Unit struct {
	data  *shm.UnitData
	game  *Game
	index int
}

// --- Identity ---

func (u *Unit) ID() int            { return int(u.data.ID()) }
func (u *Unit) Index() int         { return u.index }
func (u *Unit) Exists() bool       { return u.data.Exists() }
func (u *Unit) GetType() UnitType  { return UnitType(u.data.TypeID()) }

// GetPlayer returns the owning player.
func (u *Unit) GetPlayer() *Player {
	idx := int(u.data.PlayerIndex())
	if idx < 0 || idx >= shm.MaxPlayers {
		return nil
	}
	return &Player{data: u.game.data.Player(idx), game: u.game, index: idx}
}

// --- Position ---

func (u *Unit) GetPosition() Position {
	return Position{X: u.data.PositionX(), Y: u.data.PositionY()}
}

func (u *Unit) GetTilePosition() TilePosition {
	return TilePosition{X: u.data.PositionX() / 32, Y: u.data.PositionY() / 32}
}

func (u *Unit) Angle() float64      { return u.data.Angle() }
func (u *Unit) VelocityX() float64  { return u.data.VelocityX() }
func (u *Unit) VelocityY() float64  { return u.data.VelocityY() }

// --- Stats ---

func (u *Unit) HitPoints() int       { return int(u.data.HitPoints()) }
func (u *Unit) LastHitPoints() int    { return int(u.data.LastHitPoints()) }
func (u *Unit) Shields() int          { return int(u.data.Shields()) }
func (u *Unit) Energy() int           { return int(u.data.Energy()) }
func (u *Unit) Resources() int        { return int(u.data.Resources()) }
func (u *Unit) KillCount() int        { return int(u.data.KillCount()) }
func (u *Unit) ScarabCount() int      { return int(u.data.ScarabCount()) }
func (u *Unit) InterceptorCount() int { return int(u.data.InterceptorCount()) }
func (u *Unit) SpiderMineCount() int  { return int(u.data.SpiderMineCount()) }

// --- Timers ---

func (u *Unit) GroundWeaponCooldown() int { return int(u.data.GroundWeaponCooldown()) }
func (u *Unit) AirWeaponCooldown() int    { return int(u.data.AirWeaponCooldown()) }
func (u *Unit) SpellCooldown() int         { return int(u.data.SpellCooldown()) }
func (u *Unit) DefenseMatrixPoints() int   { return int(u.data.DefenseMatrixPoints()) }
func (u *Unit) DefenseMatrixTimer() int    { return int(u.data.DefenseMatrixTimer()) }
func (u *Unit) StimTimer() int             { return int(u.data.StimTimer()) }
func (u *Unit) RemainingBuildTime() int    { return int(u.data.RemainingBuildTime()) }
func (u *Unit) RemainingTrainTime() int    { return int(u.data.RemainingTrainTime()) }

// --- Orders ---

func (u *Unit) GetOrder() Order          { return Order(u.data.OrderID()) }
func (u *Unit) GetSecondaryOrder() Order  { return Order(u.data.SecondaryOrderID()) }
func (u *Unit) GetBuildType() UnitType    { return UnitType(u.data.BuildTypeID()) }

func (u *Unit) GetTarget() *Unit {
	idx := int(u.data.TargetIndex())
	if idx < 0 || idx >= shm.MaxUnits {
		return nil
	}
	return u.game.GetUnit(idx)
}

func (u *Unit) GetTargetPosition() Position {
	return Position{X: u.data.TargetPositionX(), Y: u.data.TargetPositionY()}
}

func (u *Unit) GetOrderTarget() *Unit {
	idx := int(u.data.OrderTargetIndex())
	if idx < 0 || idx >= shm.MaxUnits {
		return nil
	}
	return u.game.GetUnit(idx)
}

func (u *Unit) GetRallyPosition() Position {
	return Position{X: u.data.RallyPositionX(), Y: u.data.RallyPositionY()}
}

// --- Training ---

func (u *Unit) TrainingQueueCount() int { return int(u.data.TrainingQueueCount()) }

func (u *Unit) GetTrainingQueue() []UnitType {
	count := int(u.data.TrainingQueueCount())
	queue := make([]UnitType, count)
	for i := 0; i < count; i++ {
		queue[i] = UnitType(u.data.TrainingQueueEntry(i))
	}
	return queue
}

// --- Flags ---

func (u *Unit) IsCompleted() bool      { return u.data.IsCompleted() }
func (u *Unit) IsConstructing() bool   { return u.data.IsConstructing() }
func (u *Unit) IsMorphing() bool       { return u.data.IsMorphing() }
func (u *Unit) IsTraining() bool       { return u.data.IsTraining() }
func (u *Unit) IsIdle() bool           { return u.data.IsIdle() }
func (u *Unit) IsMoving() bool         { return u.data.IsMoving() }
func (u *Unit) IsAttacking() bool      { return u.data.IsAttacking() }
func (u *Unit) IsGathering() bool      { return u.data.IsGathering() }
func (u *Unit) IsBurrowed() bool       { return u.data.IsBurrowed() }
func (u *Unit) IsCloaked() bool        { return u.data.IsCloaked() }
func (u *Unit) IsDetected() bool       { return u.data.IsDetected() }
func (u *Unit) IsLifted() bool         { return u.data.IsLifted() }
func (u *Unit) IsSelected() bool       { return u.data.IsSelected() }
func (u *Unit) IsPowered() bool        { return u.data.IsPowered() }
func (u *Unit) IsInvincible() bool     { return u.data.IsInvincible() }
func (u *Unit) IsHallucination() bool  { return u.data.IsHallucination() }
func (u *Unit) HasNuke() bool          { return u.data.HasNuke() }
func (u *Unit) IsUnderStorm() bool     { return u.data.IsUnderStorm() }
func (u *Unit) IsUnderDarkSwarm() bool { return u.data.IsUnderDarkSwarm() }

func (u *Unit) IsVisibleTo(playerIndex int) bool {
	return u.data.IsVisibleTo(playerIndex)
}

// --- Related Units ---

func (u *Unit) GetAddon() *Unit {
	idx := int(u.data.AddonIndex())
	if idx < 0 || idx >= shm.MaxUnits {
		return nil
	}
	return u.game.GetUnit(idx)
}

func (u *Unit) GetTransport() *Unit {
	idx := int(u.data.TransportIndex())
	if idx < 0 || idx >= shm.MaxUnits {
		return nil
	}
	return u.game.GetUnit(idx)
}

func (u *Unit) GetCarrier() *Unit {
	idx := int(u.data.CarrierIndex())
	if idx < 0 || idx >= shm.MaxUnits {
		return nil
	}
	return u.game.GetUnit(idx)
}

func (u *Unit) GetHatchery() *Unit {
	idx := int(u.data.HatcheryIndex())
	if idx < 0 || idx >= shm.MaxUnits {
		return nil
	}
	return u.game.GetUnit(idx)
}

// --- Commands ---

func (u *Unit) issueCommand(cmdType UnitCommandType, targetIdx, x, y, extra int32) {
	u.game.data.AddUnitCommand(int32(cmdType), int32(u.index), targetIdx, x, y, extra)
}

func (u *Unit) Attack(target *Unit) {
	u.issueCommand(UnitCommandTypeAttackUnit, int32(target.index), 0, 0, 0)
}

func (u *Unit) AttackMove(pos Position) {
	u.issueCommand(UnitCommandTypeAttackMove, -1, pos.X, pos.Y, 0)
}

func (u *Unit) Move(pos Position) {
	u.issueCommand(UnitCommandTypeMove, -1, pos.X, pos.Y, 0)
}

func (u *Unit) Patrol(pos Position) {
	u.issueCommand(UnitCommandTypePatrol, -1, pos.X, pos.Y, 0)
}

func (u *Unit) HoldPosition() {
	u.issueCommand(UnitCommandTypeHoldPosition, -1, 0, 0, 0)
}

func (u *Unit) Stop() {
	u.issueCommand(UnitCommandTypeStop, -1, 0, 0, 0)
}

func (u *Unit) Follow(target *Unit) {
	u.issueCommand(UnitCommandTypeFollow, int32(target.index), 0, 0, 0)
}

func (u *Unit) Gather(target *Unit) {
	u.issueCommand(UnitCommandTypeGather, int32(target.index), 0, 0, 0)
}

func (u *Unit) ReturnCargo() {
	u.issueCommand(UnitCommandTypeReturnCargo, -1, 0, 0, 0)
}

func (u *Unit) Repair(target *Unit) {
	u.issueCommand(UnitCommandTypeRepair, int32(target.index), 0, 0, 0)
}

func (u *Unit) Build(unitType UnitType, pos TilePosition) {
	u.issueCommand(UnitCommandTypeBuild, -1, pos.X, pos.Y, int32(unitType))
}

func (u *Unit) BuildAddon(unitType UnitType) {
	u.issueCommand(UnitCommandTypeBuildAddon, -1, 0, 0, int32(unitType))
}

func (u *Unit) Train(unitType UnitType) {
	u.issueCommand(UnitCommandTypeTrain, -1, 0, 0, int32(unitType))
}

func (u *Unit) Morph(unitType UnitType) {
	u.issueCommand(UnitCommandTypeMorph, -1, 0, 0, int32(unitType))
}

func (u *Unit) Research(tech TechType) {
	u.issueCommand(UnitCommandTypeResearch, -1, 0, 0, int32(tech))
}

func (u *Unit) Upgrade(upgrade UpgradeType) {
	u.issueCommand(UnitCommandTypeUpgrade, -1, 0, 0, int32(upgrade))
}

func (u *Unit) SetRallyPosition(pos Position) {
	u.issueCommand(UnitCommandTypeSetRallyPosition, -1, pos.X, pos.Y, 0)
}

func (u *Unit) SetRallyUnit(target *Unit) {
	u.issueCommand(UnitCommandTypeSetRallyUnit, int32(target.index), 0, 0, 0)
}

func (u *Unit) Burrow() {
	u.issueCommand(UnitCommandTypeBurrow, -1, 0, 0, 0)
}

func (u *Unit) Unburrow() {
	u.issueCommand(UnitCommandTypeUnburrow, -1, 0, 0, 0)
}

func (u *Unit) Cloak() {
	u.issueCommand(UnitCommandTypeCloak, -1, 0, 0, 0)
}

func (u *Unit) Decloak() {
	u.issueCommand(UnitCommandTypeDecloak, -1, 0, 0, 0)
}

func (u *Unit) Siege() {
	u.issueCommand(UnitCommandTypeSiege, -1, 0, 0, 0)
}

func (u *Unit) Unsiege() {
	u.issueCommand(UnitCommandTypeUnsiege, -1, 0, 0, 0)
}

func (u *Unit) Lift() {
	u.issueCommand(UnitCommandTypeLift, -1, 0, 0, 0)
}

func (u *Unit) Land(pos TilePosition) {
	u.issueCommand(UnitCommandTypeLand, -1, pos.X, pos.Y, 0)
}

func (u *Unit) Load(target *Unit) {
	u.issueCommand(UnitCommandTypeLoad, int32(target.index), 0, 0, 0)
}

func (u *Unit) Unload(target *Unit) {
	u.issueCommand(UnitCommandTypeUnload, int32(target.index), 0, 0, 0)
}

func (u *Unit) UnloadAll() {
	u.issueCommand(UnitCommandTypeUnloadAll, -1, 0, 0, 0)
}

func (u *Unit) UnloadAllPosition(pos Position) {
	u.issueCommand(UnitCommandTypeUnloadAllPosition, -1, pos.X, pos.Y, 0)
}

func (u *Unit) RightClick(target *Unit) {
	u.issueCommand(UnitCommandTypeRightClickUnit, int32(target.index), 0, 0, 0)
}

func (u *Unit) RightClickPosition(pos Position) {
	u.issueCommand(UnitCommandTypeRightClickPosition, -1, pos.X, pos.Y, 0)
}

func (u *Unit) HaltConstruction() {
	u.issueCommand(UnitCommandTypeHaltConstruction, -1, 0, 0, 0)
}

func (u *Unit) CancelConstruction() {
	u.issueCommand(UnitCommandTypeCancelConstruction, -1, 0, 0, 0)
}

func (u *Unit) CancelAddon() {
	u.issueCommand(UnitCommandTypeCancelAddon, -1, 0, 0, 0)
}

func (u *Unit) CancelTrain(slot int) {
	u.issueCommand(UnitCommandTypeCancelTrainSlot, -1, 0, 0, int32(slot))
}

func (u *Unit) CancelMorph() {
	u.issueCommand(UnitCommandTypeCancelMorph, -1, 0, 0, 0)
}

func (u *Unit) CancelResearch() {
	u.issueCommand(UnitCommandTypeCancelResearch, -1, 0, 0, 0)
}

func (u *Unit) CancelUpgrade() {
	u.issueCommand(UnitCommandTypeCancelUpgrade, -1, 0, 0, 0)
}

func (u *Unit) UseTech(tech TechType) {
	u.issueCommand(UnitCommandTypeUseTech, -1, 0, 0, int32(tech))
}

func (u *Unit) UseTechPosition(tech TechType, pos Position) {
	u.issueCommand(UnitCommandTypeUseTechPosition, -1, pos.X, pos.Y, int32(tech))
}

func (u *Unit) UseTechUnit(tech TechType, target *Unit) {
	u.issueCommand(UnitCommandTypeUseTechUnit, int32(target.index), 0, 0, int32(tech))
}

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

func (u *Unit) ID() int           { return int(u.data.ID()) }
func (u *Unit) Index() int        { return u.index }
func (u *Unit) Exists() bool      { return u.data.Exists() }
func (u *Unit) GetType() UnitType { return UnitType(u.data.TypeID()) }

// GetPlayer returns the owning player.
func (u *Unit) GetPlayer() *Player {
	idx := int(u.data.PlayerIndex())
	if idx < 0 || idx >= shm.MaxPlayers {
		return nil
	}
	return &Player{data: u.game.data.Player(idx), game: u.game, index: idx}
}

func (u *Unit) GetPosition() Position {
	return Position{X: u.data.PositionX(), Y: u.data.PositionY()}
}

func (u *Unit) GetTilePosition() TilePosition {
	return TilePosition{X: u.data.PositionX() / 32, Y: u.data.PositionY() / 32}
}

func (u *Unit) Angle() float64     { return u.data.Angle() }
func (u *Unit) VelocityX() float64 { return u.data.VelocityX() }
func (u *Unit) VelocityY() float64 { return u.data.VelocityY() }

func (u *Unit) HitPoints() int        { return int(u.data.HitPoints()) }
func (u *Unit) LastHitPoints() int    { return int(u.data.LastHitPoints()) }
func (u *Unit) Shields() int          { return int(u.data.Shields()) }
func (u *Unit) Energy() int           { return int(u.data.Energy()) }
func (u *Unit) Resources() int        { return int(u.data.Resources()) }
func (u *Unit) KillCount() int        { return int(u.data.KillCount()) }
func (u *Unit) ScarabCount() int      { return int(u.data.ScarabCount()) }
func (u *Unit) InterceptorCount() int { return int(u.data.InterceptorCount()) }
func (u *Unit) SpiderMineCount() int  { return int(u.data.SpiderMineCount()) }

func (u *Unit) GroundWeaponCooldown() int { return int(u.data.GroundWeaponCooldown()) }
func (u *Unit) AirWeaponCooldown() int    { return int(u.data.AirWeaponCooldown()) }
func (u *Unit) SpellCooldown() int        { return int(u.data.SpellCooldown()) }
func (u *Unit) DefenseMatrixPoints() int  { return int(u.data.DefenseMatrixPoints()) }
func (u *Unit) DefenseMatrixTimer() int   { return int(u.data.DefenseMatrixTimer()) }
func (u *Unit) StimTimer() int              { return int(u.data.StimTimer()) }
func (u *Unit) EnsnareTimer() int           { return int(u.data.EnsnareTimer()) }
func (u *Unit) IrradiateTimer() int         { return int(u.data.IrradiateTimer()) }
func (u *Unit) LockdownTimer() int          { return int(u.data.LockdownTimer()) }
func (u *Unit) MaelstromTimer() int         { return int(u.data.MaelstromTimer()) }
func (u *Unit) OrderTimer() int             { return int(u.data.OrderTimer()) }
func (u *Unit) PlagueTimer() int            { return int(u.data.PlagueTimer()) }
func (u *Unit) RemoveTimer() int            { return int(u.data.RemoveTimer()) }
func (u *Unit) StasisTimer() int            { return int(u.data.StasisTimer()) }
func (u *Unit) RemainingBuildTime() int     { return int(u.data.RemainingBuildTime()) }
func (u *Unit) RemainingTrainTime() int     { return int(u.data.RemainingTrainTime()) }
func (u *Unit) RemainingResearchTime() int  { return int(u.data.RemainingResearchTime()) }
func (u *Unit) RemainingUpgradeTime() int   { return int(u.data.RemainingUpgradeTime()) }

func (u *Unit) GetOrder() Order          { return Order(u.data.OrderID()) }
func (u *Unit) GetSecondaryOrder() Order { return Order(u.data.SecondaryOrderID()) }
func (u *Unit) GetBuildType() UnitType   { return UnitType(u.data.BuildTypeID()) }

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

func (u *Unit) TrainingQueueCount() int { return int(u.data.TrainingQueueCount()) }

func (u *Unit) GetTrainingQueue() []UnitType {
	count := int(u.data.TrainingQueueCount())
	queue := make([]UnitType, count)
	for i := 0; i < count; i++ {
		queue[i] = UnitType(u.data.TrainingQueueEntry(i))
	}
	return queue
}

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
func (u *Unit) IsUnderDarkSwarm() bool      { return u.data.IsUnderDarkSwarm() }
func (u *Unit) IsUnderDisruptionWeb() bool  { return u.data.IsUnderDWeb() }
func (u *Unit) IsAccelerating() bool        { return u.data.IsAccelerating() }
func (u *Unit) IsAttackFrame() bool         { return u.data.IsAttackFrame() }
func (u *Unit) IsBeingGathered() bool       { return u.data.IsBeingGathered() }
func (u *Unit) IsBlind() bool               { return u.data.IsBlind() }
func (u *Unit) IsBraking() bool             { return u.data.IsBraking() }
func (u *Unit) IsInterruptible() bool       { return u.data.IsInterruptible() }
func (u *Unit) IsParasited() bool           { return u.data.IsParasited() }
func (u *Unit) IsStartingAttack() bool      { return u.data.IsStartingAttack() }
func (u *Unit) IsStuck() bool { return u.data.IsStuck() }

// IsDefenseMatrixed returns whether this unit is under a Defense Matrix.
func (u *Unit) IsDefenseMatrixed() bool { return u.data.DefenseMatrixTimer() > 0 }

// IsEnsnared returns whether this unit is under Ensnare.
func (u *Unit) IsEnsnared() bool { return u.data.EnsnareTimer() > 0 }

// IsIrradiated returns whether this unit is under Irradiate.
func (u *Unit) IsIrradiated() bool { return u.data.IrradiateTimer() > 0 }

// IsLockedDown returns whether this unit is under Lockdown.
func (u *Unit) IsLockedDown() bool { return u.data.LockdownTimer() > 0 }

// IsMaelstrommed returns whether this unit is under Maelstrom.
func (u *Unit) IsMaelstrommed() bool { return u.data.MaelstromTimer() > 0 }

// IsPlagued returns whether this unit is under Plague.
func (u *Unit) IsPlagued() bool { return u.data.PlagueTimer() > 0 }

// IsStasised returns whether this unit is under Stasis Field.
func (u *Unit) IsStasised() bool { return u.data.StasisTimer() > 0 }

// IsStimmed returns whether this unit is under Stim Pack.
func (u *Unit) IsStimmed() bool { return u.data.StimTimer() > 0 }

// IsUnderAttack returns whether this unit was recently attacked.
func (u *Unit) IsUnderAttack() bool { return u.data.RecentlyAttacked() }

// IsCarryingGas returns whether this worker is carrying gas.
func (u *Unit) IsCarryingGas() bool { return u.data.CarryResourceType() == 1 }

// IsCarryingMinerals returns whether this worker is carrying minerals.
func (u *Unit) IsCarryingMinerals() bool { return u.data.CarryResourceType() == 2 }

// IsLoaded returns whether this unit is loaded inside a transport.
func (u *Unit) IsLoaded() bool { return u.GetTransport() != nil }

// IsSieged returns whether this unit is in siege mode.
func (u *Unit) IsSieged() bool {
	t := u.GetType()
	return t == UnitTypeTerranSiegeTankSiegeMode || t == UnitTypeHeroEdmundDukeSiegeMode
}

// IsFollowing returns whether this unit is executing a Follow order.
func (u *Unit) IsFollowing() bool { return u.GetOrder() == OrderFollow }

// IsHoldingPosition returns whether this unit is executing a Hold Position order.
func (u *Unit) IsHoldingPosition() bool { return u.GetOrder() == OrderHoldPosition }

// IsPatrolling returns whether this unit is executing a Patrol order.
func (u *Unit) IsPatrolling() bool { return u.GetOrder() == OrderPatrol }

// IsRepairing returns whether this unit is repairing.
func (u *Unit) IsRepairing() bool {
	o := u.GetOrder()
	return o == OrderRepair || o == OrderMoveToRepair
}

// IsResearching returns whether this building is researching a technology.
func (u *Unit) IsResearching() bool { return u.GetOrder() == OrderResearchTech }

// IsUpgrading returns whether this building is performing an upgrade.
func (u *Unit) IsUpgrading() bool { return u.GetOrder() == OrderUpgrade }

// IsGatheringGas returns whether this worker is gathering gas.
func (u *Unit) IsGatheringGas() bool {
	if !u.data.IsGathering() {
		return false
	}
	o := u.GetOrder()
	if o == OrderHarvest1 || o == OrderHarvest2 {
		return true
	}
	if o == OrderResetHarvestCollision {
		return u.data.CarryResourceType() == 1
	}
	return o == OrderMoveToGas || o == OrderWaitForGas ||
		o == OrderHarvestGas || o == OrderReturnGas
}

// IsGatheringMinerals returns whether this worker is gathering minerals.
func (u *Unit) IsGatheringMinerals() bool {
	if !u.data.IsGathering() {
		return false
	}
	o := u.GetOrder()
	if o == OrderHarvest1 || o == OrderHarvest2 {
		return true
	}
	if o == OrderResetHarvestCollision {
		return u.data.CarryResourceType() == 2
	}
	return o == OrderMoveToMinerals || o == OrderWaitForMinerals ||
		o == OrderMiningMinerals || o == OrderReturnMinerals
}

// IsBeingConstructed returns whether this unit/building is being constructed.
// Matches JBWAPI: morphing units always count; for Terran buildings, only if
// an SCV is actively building; non-Terran incomplete buildings always count.
func (u *Unit) IsBeingConstructed() bool {
	if u.data.IsMorphing() {
		return true
	}
	if u.data.IsCompleted() {
		return false
	}
	if u.GetType().GetRace() != RaceTerran {
		return true
	}
	return u.GetBuildUnit() != nil
}

// IsFlying returns whether this unit is airborne.
func (u *Unit) IsFlying() bool {
	return u.GetType().IsFlyer() || u.data.IsLifted()
}

// IsTargetable returns whether this unit can be targeted by commands.
func (u *Unit) IsTargetable() bool {
	if !u.data.Exists() {
		return false
	}
	t := u.GetType()
	if !u.data.IsCompleted() && !t.IsBuilding() && !u.data.IsMorphing() &&
		t != UnitTypeProtossArchon && t != UnitTypeProtossDarkArchon {
		return false
	}
	return t != UnitTypeSpellScannerSweep && t != UnitTypeSpellDarkSwarm &&
		t != UnitTypeSpellDisruptionWeb && t != UnitTypeSpecialMapRevealer
}

// IsBeingHealed returns whether this unit is being healed by a Medic.
func (u *Unit) IsBeingHealed() bool {
	for _, unit := range u.game.GetAllUnits() {
		if unit.GetType() != UnitTypeTerranMedic || unit.GetOrder() != OrderMedicHeal {
			continue
		}
		if target := unit.GetTarget(); target != nil && target.Index() == u.index {
			return true
		}
	}
	return false
}

// IsInWeaponRange returns whether the target is within this unit's weapon range.
// Accounts for player weapon range upgrades.
func (u *Unit) IsInWeaponRange(target *Unit) bool {
	if target == nil {
		return false
	}
	var weapon WeaponType
	if target.IsFlying() {
		weapon = u.GetType().AirWeapon()
	} else {
		weapon = u.GetType().GroundWeapon()
	}
	if weapon == WeaponTypeNone {
		return false
	}
	maxRange := weapon.MaxRange()
	if p := u.GetPlayer(); p != nil {
		maxRange += p.WeaponMaxRangeUpgrade(weapon)
	}
	minRange := weapon.MinRange()
	dist := u.GetDistanceToUnit(target)
	return dist <= maxRange && (minRange == 0 || dist > minRange)
}

func (u *Unit) IsVisibleTo(playerIndex int) bool {
	return u.data.IsVisibleTo(playerIndex)
}

func (u *Unit) IsVisible() bool {
	selfIdx := int(u.game.data.SelfIndex())
	return u.data.IsVisibleTo(selfIdx)
}

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

// GetRallyUnit returns the unit this building rallies to.
func (u *Unit) GetRallyUnit() *Unit {
	idx := int(u.data.RallyUnitIndex())
	if idx < 0 || idx >= shm.MaxUnits {
		return nil
	}
	return u.game.GetUnit(idx)
}

// GetNydusExit returns the paired Nydus Canal exit unit.
func (u *Unit) GetNydusExit() *Unit {
	idx := int(u.data.NydusExitIndex())
	if idx < 0 || idx >= shm.MaxUnits {
		return nil
	}
	return u.game.GetUnit(idx)
}

// GetPowerUp returns the powerup this unit is carrying.
func (u *Unit) GetPowerUp() *Unit {
	idx := int(u.data.PowerUpIndex())
	if idx < 0 || idx >= shm.MaxUnits {
		return nil
	}
	return u.game.GetUnit(idx)
}

// GetBuildUnit returns the unit this worker/building is currently constructing.
func (u *Unit) GetBuildUnit() *Unit {
	idx := int(u.data.BuildUnitIndex())
	if idx < 0 || idx >= shm.MaxUnits {
		return nil
	}
	return u.game.GetUnit(idx)
}

// GetTech returns the TechType currently being researched by this building.
func (u *Unit) GetTech() TechType { return TechType(u.data.TechID()) }

// GetUpgrade returns the UpgradeType currently being upgraded by this building.
func (u *Unit) GetUpgrade() UpgradeType { return UpgradeType(u.data.UpgradeID()) }

// ResourceGroup returns the resource group ID for mineral/gas grouping.
func (u *Unit) ResourceGroup() int { return int(u.data.ResourceGroup()) }

// AcidSporeCount returns the number of acid spores on this unit.
func (u *Unit) AcidSporeCount() int { return int(u.data.AcidSporeCount()) }

// GetLastAttackingPlayer returns the player whose unit last attacked this unit.
func (u *Unit) GetLastAttackingPlayer() *Player {
	idx := int(u.data.LastAttackerPlayer())
	if idx < 0 || idx >= shm.MaxPlayers {
		return nil
	}
	return &Player{data: u.game.data.Player(idx), game: u.game, index: idx}
}

// ReplayID returns this unit's replay ID (only valid in replays).
func (u *Unit) ReplayID() int { return int(u.data.ReplayID()) }

// GetSpaceRemaining returns the remaining transport space.
func (u *Unit) GetSpaceRemaining() int {
	total := u.GetType().SpaceProvided()
	if total == 0 {
		return 0
	}
	loaded := u.GetLoadedUnits()
	used := 0
	for _, lu := range loaded {
		used += lu.GetType().SpaceRequired()
	}
	return total - used
}

// GetRegion returns the map region this unit is in.
func (u *Unit) GetRegion() *Region {
	return u.game.GetRegionAt(int(u.data.PositionX()), int(u.data.PositionY()))
}

// GetUnitsInRadius returns all visible units within the given pixel radius.
func (u *Unit) GetUnitsInRadius(radius int, filter func(*Unit) bool) []*Unit {
	return u.game.getUnitsInRadiusFiltered(int(u.data.PositionX()), int(u.data.PositionY()), radius, filter)
}

// These methods return the unit's state at game start. Requires Game.SnapshotInitialState().

// GetInitialType returns this unit's type at game start.
func (u *Unit) GetInitialType() UnitType {
	if s, ok := u.game.initialStates[u.index]; ok {
		return s.unitType
	}
	return UnitTypeUnknown
}

// GetInitialPosition returns this unit's position at game start.
func (u *Unit) GetInitialPosition() Position {
	if s, ok := u.game.initialStates[u.index]; ok {
		return s.position
	}
	return Position{}
}

// GetInitialTilePosition returns this unit's tile position at game start.
func (u *Unit) GetInitialTilePosition() TilePosition {
	if s, ok := u.game.initialStates[u.index]; ok {
		return s.position.ToTilePosition()
	}
	return TilePosition{}
}

// GetInitialHitPoints returns this unit's hit points at game start.
func (u *Unit) GetInitialHitPoints() int {
	if s, ok := u.game.initialStates[u.index]; ok {
		return s.hitPoints
	}
	return 0
}

// GetInitialResources returns this unit's resource amount at game start.
func (u *Unit) GetInitialResources() int {
	if s, ok := u.game.initialStates[u.index]; ok {
		return s.resources
	}
	return 0
}

// GetLeft returns the left pixel edge of this unit's collision box.
func (u *Unit) GetLeft() int {
	return int(u.data.PositionX()) - u.GetType().DimensionLeft()
}

// GetTop returns the top pixel edge of this unit's collision box.
func (u *Unit) GetTop() int {
	return int(u.data.PositionY()) - u.GetType().DimensionUp()
}

// GetRight returns the right pixel edge of this unit's collision box.
func (u *Unit) GetRight() int {
	return int(u.data.PositionX()) + u.GetType().DimensionRight()
}

// GetBottom returns the bottom pixel edge of this unit's collision box.
func (u *Unit) GetBottom() int {
	return int(u.data.PositionY()) + u.GetType().DimensionDown()
}

// GetDistance returns the edge-to-edge distance to a position in pixels.
// Uses BWAPI's approximate distance (matches weapon range checks).
func (u *Unit) GetDistance(pos Position) int {
	l := u.GetLeft()
	t := u.GetTop()
	r := u.GetRight() + 1
	b := u.GetBottom() + 1

	px := int(pos.X)
	py := int(pos.Y)

	var xDist, yDist int32
	if px < l {
		xDist = int32(l - px)
	} else if px > r {
		xDist = int32(px - r)
	}
	if py < t {
		yDist = int32(t - py)
	} else if py > b {
		yDist = int32(py - b)
	}

	return int(approxDistance(xDist, yDist))
}

// GetDistanceToUnit returns the edge-to-edge distance to another unit in pixels.
// Uses BWAPI's box expansion: only the target box is expanded by ±1.
func (u *Unit) GetDistanceToUnit(other *Unit) int {
	if other == nil {
		return 0
	}
	l1, t1, r1, b1 := u.GetLeft(), u.GetTop(), u.GetRight(), u.GetBottom()
	l2, t2, r2, b2 := other.GetLeft()-1, other.GetTop()-1, other.GetRight()+1, other.GetBottom()+1

	var xDist, yDist int32
	if l1 > r2 {
		xDist = int32(l1 - r2)
	} else if l2 > r1 {
		xDist = int32(l2 - r1)
	}
	if t1 > b2 {
		yDist = int32(t1 - b2)
	} else if t2 > b1 {
		yDist = int32(t2 - b1)
	}

	return int(approxDistance(xDist, yDist))
}

// HasPath returns whether there is a ground path from this unit to a position.
func (u *Unit) HasPath(pos Position) bool {
	return u.game.HasPath(u.GetPosition(), pos)
}

// HasPathToUnit returns whether there is a ground path from this unit to another.
func (u *Unit) HasPathToUnit(other *Unit) bool {
	if other == nil {
		return false
	}
	return u.game.HasPath(u.GetPosition(), other.GetPosition())
}

// GetLarva returns all larva units associated with this hatchery/lair/hive.
func (u *Unit) GetLarva() []*Unit {
	allUnits := u.game.GetAllUnits()
	larva := make([]*Unit, 0)
	for _, unit := range allUnits {
		if unit.GetType() == UnitTypeZergLarva && unit.GetHatchery() != nil &&
			unit.GetHatchery().Index() == u.index {
			larva = append(larva, unit)
		}
	}
	return larva
}

// GetLoadedUnits returns all units loaded inside this transport/bunker.
func (u *Unit) GetLoadedUnits() []*Unit {
	allUnits := u.game.GetAllUnits()
	loaded := make([]*Unit, 0)
	for _, unit := range allUnits {
		if unit.GetTransport() != nil && unit.GetTransport().Index() == u.index {
			loaded = append(loaded, unit)
		}
	}
	return loaded
}

// GetInterceptors returns all interceptor units belonging to this carrier.
func (u *Unit) GetInterceptors() []*Unit {
	allUnits := u.game.GetAllUnits()
	interceptors := make([]*Unit, 0)
	for _, unit := range allUnits {
		if unit.GetType() == UnitTypeProtossInterceptor && unit.GetCarrier() != nil &&
			unit.GetCarrier().Index() == u.index {
			interceptors = append(interceptors, unit)
		}
	}
	return interceptors
}

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

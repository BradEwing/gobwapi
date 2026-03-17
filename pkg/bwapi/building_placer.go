package bwapi

const placerMaxRange = 64

var bpDirections = [8]TilePosition{
	{1, 1}, {0, 1}, {-1, 1},
	{1, 0}, {-1, 0},
	{1, -1}, {0, -1}, {-1, -1},
}

type buildTemplate struct {
	startX, startY, stepX, stepY int
}

var buildTemplates = [14]buildTemplate{
	{32, 0, 0, 1},
	{0, 32, 1, 0},
	{31, 0, 0, 1},
	{0, 31, 1, 0},
	{33, 0, 0, 1},
	{0, 33, 1, 0},
	{30, 0, 0, 1},
	{29, 0, 0, 1},
	{0, 30, 1, 0},
	{28, 0, 0, 1},
	{0, 29, 1, 0},
	{27, 0, 0, 1},
	{0, 28, 1, 0},
	{-1, 0, 0, 0}, // sentinel
}

// placementReserve is a 64x64 bitmap used by the building placement algorithm.
// It matches BWAPI's PlacementReserve class.
type placementReserve struct {
	data      [placerMaxRange][placerMaxRange]byte
	save      [placerMaxRange][placerMaxRange]byte
	maxSearch int
}

func newPlacementReserve(maxRange int) *placementReserve {
	if maxRange < 0 {
		maxRange = 0
	}
	if maxRange > placerMaxRange {
		maxRange = placerMaxRange
	}
	pr := &placementReserve{maxSearch: maxRange}
	// data and save are zero-initialized by Go
	pr.backup()
	return pr
}

func prIsValidPos(x, y int) bool {
	return x >= 0 && x < placerMaxRange && y >= 0 && y < placerMaxRange
}

func prIsValidPosTile(p TilePosition) bool {
	return prIsValidPos(int(p.X), int(p.Y))
}

func (pr *placementReserve) setValue(x, y int, value byte) {
	if prIsValidPos(x, y) {
		pr.data[y][x] = value
	}
}

func (pr *placementReserve) setValueTile(p TilePosition, value byte) {
	pr.setValue(int(p.X), int(p.Y), value)
}

func (pr *placementReserve) setRange(left, top, right, bottom int, value byte) {
	for y := top; y < bottom; y++ {
		for x := left; x < right; x++ {
			pr.setValue(x, y, value)
		}
	}
}

func (pr *placementReserve) setRangeTile(lt, rb TilePosition, value byte) {
	pr.setRange(int(lt.X), int(lt.Y), int(rb.X), int(rb.Y), value)
}

func (pr *placementReserve) getValue(x, y int) byte {
	if prIsValidPos(x, y) {
		return pr.data[y][x]
	}
	return 0
}

func (pr *placementReserve) getValueTile(p TilePosition) byte {
	return pr.getValue(int(p.X), int(p.Y))
}

func (pr *placementReserve) iterate(proc func(pr *placementReserve, x, y int)) {
	min := placerMaxRange/2 - pr.maxSearch/2
	max := min + pr.maxSearch
	for y := min; y < max; y++ {
		for x := min; x < max; x++ {
			proc(pr, x, y)
		}
	}
}

func (pr *placementReserve) hasValidSpace() bool {
	min := placerMaxRange/2 - pr.maxSearch/2
	max := min + pr.maxSearch
	for y := min; y < max; y++ {
		for x := min; x < max; x++ {
			if pr.data[y][x] == 1 {
				return true
			}
		}
	}
	return false
}

func (pr *placementReserve) backup() {
	pr.save = pr.data
}

func (pr *placementReserve) restore() {
	pr.data = pr.save
}

func (pr *placementReserve) restoreIfInvalid() {
	if !pr.hasValidSpace() {
		pr.restore()
	}
}

// getSelfUnits returns all units owned by the self player.
func getSelfUnits(g *Game) []*Unit {
	self := g.Self()
	if self == nil {
		return nil
	}
	allUnits := g.GetAllUnits()
	result := make([]*Unit, 0, len(allUnits))
	for _, u := range allUnits {
		if u.GetPlayer() != nil && u.GetPlayer().Index() == self.Index() {
			result = append(result, u)
		}
	}
	return result
}

func bpAssignBuildableLocations(reserve *placementReserve, unitType UnitType, desiredPosition TilePosition, g *Game) {
	start := desiredPosition.Sub(TilePosition{placerMaxRange, placerMaxRange}.Div(2))
	hasAddon := unitType.CanBuildAddon()

	reserve.iterate(func(pr *placementReserve, x, y int) {
		if hasAddon && !g.CanBuildHere(start.Add(TilePosition{int32(x + 4), int32(y + 1)}), UnitTypeTerranMissileTurret, nil, false) {
			return
		}
		if g.CanBuildHere(start.Add(TilePosition{int32(x), int32(y)}), unitType, nil, false) {
			pr.setValue(x, y, 1)
		}
	})
}

func bpRemoveDisconnected(reserve *placementReserve, desiredPosition TilePosition, g *Game) {
	start := desiredPosition.Sub(TilePosition{placerMaxRange, placerMaxRange}.Div(2))

	reserve.iterate(func(pr *placementReserve, x, y int) {
		if !g.HasPath(desiredPosition.ToPosition(), start.Add(TilePosition{int32(x), int32(y)}).ToPosition()) {
			pr.setValue(x, y, 0)
		}
	})
}

func bpReserveGroundHeight(reserve *placementReserve, desiredPosition TilePosition, g *Game) {
	start := desiredPosition.Sub(TilePosition{placerMaxRange, placerMaxRange}.Div(2))

	reserve.backup()
	targetHeight := g.GetGroundHeight(int(desiredPosition.X), int(desiredPosition.Y))
	reserve.iterate(func(pr *placementReserve, x, y int) {
		tp := start.Add(TilePosition{int32(x), int32(y)})
		if g.GetGroundHeight(int(tp.X), int(tp.Y)) != targetHeight {
			pr.setValue(x, y, 0)
		}
	})

	reserve.restoreIfInvalid()
}

func bpReserveStructureWithPadding(reserve *placementReserve, currentPosition, sizeExtra TilePosition, padding int, unitType UnitType, desiredPosition TilePosition) {
	paddingTP := TilePosition{int32(padding), int32(padding)}
	paddingSize := sizeExtra.Add(paddingTP.Mul(2))

	topLeft := currentPosition.Sub(unitType.TileSize()).Sub(paddingSize.Div(2)).Sub(TilePosition{1, 1})
	topLeftRelative := topLeft.Sub(desiredPosition).Add(TilePosition{placerMaxRange, placerMaxRange}.Div(2))
	maxSize := topLeftRelative.Add(unitType.TileSize()).Add(paddingSize).Add(TilePosition{1, 1})

	reserve.setRangeTile(topLeftRelative, maxSize, 0)
}

func bpReserveStructure(reserve *placementReserve, unit *Unit, padding int, unitType UnitType, desiredPosition TilePosition) {
	bpReserveStructureWithPadding(reserve,
		unit.GetPosition().ToTilePosition(),
		unit.GetType().TileSize(),
		padding, unitType, desiredPosition)
}

func bpReserveAllStructures(reserve *placementReserve, unitType UnitType, desiredPosition TilePosition, g *Game) {
	if unitType.IsAddon() {
		return
	}
	reserve.backup()

	// Reserve space around owned resource depots and refineries
	myUnits := getSelfUnits(g)
	for _, u := range myUnits {
		if !u.Exists() {
			continue
		}
		ut := u.GetType()
		if !ut.IsBuilding() {
			continue
		}
		if !(u.IsCompleted() || (ut.ProducesLarva() && u.IsMorphing())) {
			continue
		}
		if !(ut.IsResourceDepot() || ut.IsRefinery()) {
			continue
		}
		bpReserveStructure(reserve, u, 2, unitType, desiredPosition)
	}

	// Reserve space around neutral resources
	if unitType != UnitTypeTerranBunker {
		neutrals := g.GetNeutralUnits()
		for _, u := range neutrals {
			if u.Exists() && u.GetType().IsResourceContainer() {
				bpReserveStructure(reserve, u, 2, unitType, desiredPosition)
			}
		}
	}

	reserve.restoreIfInvalid()
}

func bpReserveExistingAddonPlacement(reserve *placementReserve, desiredPosition TilePosition, g *Game) {
	start := desiredPosition.Sub(TilePosition{placerMaxRange, placerMaxRange}.Div(2))

	reserve.backup()
	myUnits := getSelfUnits(g)
	for _, u := range myUnits {
		if u.Exists() && u.GetType().CanBuildAddon() {
			addonPos := u.GetTilePosition().Add(TilePosition{4, 1}).Sub(start)
			reserve.setRangeTile(addonPos, addonPos.Add(TilePosition{2, 2}), 0)
		}
	}

	reserve.restoreIfInvalid()
}

func bpReserveDefault(reserve *placementReserve, unitType UnitType, desiredPosition TilePosition, g *Game) {
	reserve.backup()

	// Save a copy of the current state before structure padding
	var original [placerMaxRange][placerMaxRange]byte
	original = reserve.data

	myUnits := getSelfUnits(g)
	for _, it := range myUnits {
		if !it.Exists() {
			continue
		}
		switch it.GetType() {
		case UnitTypeTerranBarracks, UnitTypeTerranBunker, UnitTypeZergCreepColony:
			bpReserveStructure(reserve, it, 1, unitType, desiredPosition)
		default:
			bpReserveStructure(reserve, it, 2, unitType, desiredPosition)
		}
	}

	switch unitType {
	case UnitTypeTerranBarracks, UnitTypeTerranFactory,
		UnitTypeTerranMissileTurret, UnitTypeTerranBunker,
		UnitTypeProtossRoboticsFacility, UnitTypeProtossGateway,
		UnitTypeProtossPhotonCannon:
		for y := 0; y < 64; y++ {
			for x := 0; x < 64; x++ {
				for dir := 0; dir < 8; dir++ {
					p := TilePosition{int32(x), int32(y)}.Add(bpDirections[dir])
					if !prIsValidPosTile(p) || original[p.Y][p.X] == 0 {
						reserve.setValueTile(p, 0)
					}
				}
			}
		}
	}

	reserve.restoreIfInvalid()
}

func bpReserveTemplateSpacing(reserve *placementReserve) {
	reserve.backup()

	for j := 0; buildTemplates[j].startX != -1; j++ {
		t := buildTemplates[j]
		x, y := t.startX, t.startY
		for i := 0; i < 64; i++ {
			reserve.setValue(x, y, 0)
			x += t.stepX
			y += t.stepY
		}
	}

	reserve.restoreIfInvalid()
}

func bpReservePlacement(reserve *placementReserve, unitType UnitType, desiredPosition TilePosition, g *Game) {
	reserve.data = [placerMaxRange][placerMaxRange]byte{}

	bpAssignBuildableLocations(reserve, unitType, desiredPosition, g)
	bpRemoveDisconnected(reserve, desiredPosition, g)

	// Exclude positions off the map
	start := desiredPosition.Sub(TilePosition{placerMaxRange, placerMaxRange}.Div(2))
	reserve.iterate(func(pr *placementReserve, x, y int) {
		if !start.Add(TilePosition{int32(x), int32(y)}).IsValid(g) {
			pr.setValue(x, y, 0)
		}
	})

	if !reserve.hasValidSpace() {
		return
	}

	bpReserveGroundHeight(reserve, desiredPosition, g)

	if !unitType.IsResourceDepot() {
		bpReserveAllStructures(reserve, unitType, desiredPosition, g)
		bpReserveExistingAddonPlacement(reserve, desiredPosition, g)
	}

	switch unitType {
	case UnitTypeProtossPylon:
		// @TODO: reservePylonPlacement
	case UnitTypeTerranBunker:
		// @TODO: reserveBunkerPlacement
	case UnitTypeTerranMissileTurret, UnitTypeProtossPhotonCannon:
		// @TODO: reserveTurretPlacement
	case UnitTypeZergCreepColony:
		// @TODO: reserveCreepColonyPlacement
	default:
		if !unitType.IsResourceDepot() {
			bpReserveDefault(reserve, unitType, desiredPosition, g)
		}
	}
}

// buildingPlacerGetBuildLocation implements BWAPI's getBuildLocation algorithm
// using a 64x64 PlacementReserve bitmap with multiple reservation passes.
func buildingPlacerGetBuildLocation(unitType UnitType, desiredPosition TilePosition, maxRange int, creep bool, g *Game) TilePosition {
	if !unitType.IsBuilding() {
		return TilePosition{-1, -1} // TilePositions::Invalid
	}

	// Do type-specific checks
	trimPlacement := true
	switch unitType {
	case UnitTypeProtossPylon:
		pos := desiredPosition.ToPosition()
		target := g.GetClosestUnit(int(pos.X), int(pos.Y), func(u *Unit) bool {
			return u.GetPlayer() != nil && g.Self() != nil &&
				u.GetPlayer().Index() == g.Self().Index() && !u.IsPowered()
		})
		if target != nil {
			desiredPosition = target.GetPosition().ToTilePosition()
			trimPlacement = false
		}
	case UnitTypeTerranCommandCenter, UnitTypeProtossNexus,
		UnitTypeZergHatchery, UnitTypeSpecialStartLocation:
		trimPlacement = false
	case UnitTypeZergCreepColony, UnitTypeTerranBunker:
		// @TODO
	}

	reserve := newPlacementReserve(maxRange)
	bpReservePlacement(reserve, unitType, desiredPosition, g)

	if trimPlacement {
		bpReserveTemplateSpacing(reserve)
	}

	centerPosition := desiredPosition.Sub(TilePosition{placerMaxRange, placerMaxRange}.Div(2))

	bestDistance := int32(999999)
	fallbackDistance := int32(999999)
	bestPosition := TilePosition{-1, -1}  // None
	fallbackPosition := TilePosition{-1, -1}

	for y := 0; y < placerMaxRange; y++ {
		for x := 0; x < placerMaxRange; x++ {
			if reserve.getValue(x, y) == 0 {
				continue
			}

			currentPosition := TilePosition{int32(x), int32(y)}.Add(centerPosition)
			currentDistance := desiredPosition.GetApproxDistance(currentPosition)

			if currentDistance < bestDistance {
				if currentDistance <= int32(maxRange) {
					bestDistance = currentDistance
					bestPosition = currentPosition
				} else if currentDistance < fallbackDistance {
					fallbackDistance = currentDistance
					fallbackPosition = currentPosition
				}
			}
		}
	}

	if bestPosition.X != -1 || bestPosition.Y != -1 {
		return bestPosition
	}
	if fallbackPosition.X != -1 || fallbackPosition.Y != -1 {
		return fallbackPosition
	}
	return TilePosition{-1, -1}
}

package bwapi

// bPsiFieldMask is the Pylon power template from BWAPI Templates.h.
// 10 rows x 16 columns, each cell represents one 32-pixel build tile.
// Offset: Pylon center ± 256px horizontal, ± 160px vertical.
var bPsiFieldMask = [10][16]bool{
	{false, false, false, false, false, true, true, true, true, true, true, false, false, false, false, false},
	{false, false, true, true, true, true, true, true, true, true, true, true, true, true, false, false},
	{false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false},
	{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
	{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
	{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
	{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
	{false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false},
	{false, false, true, true, true, true, true, true, true, true, true, true, true, true, false, false},
	{false, false, false, false, false, true, true, true, true, true, true, false, false, false, false, false},
}

// HasPowerPrecise returns whether a pixel position is powered by any completed Pylon.
func (g *Game) HasPowerPrecise(x, y int) bool {
	return g.hasPowerCheck(x, y, UnitTypeNone)
}

// HasPower returns whether a build tile is powered by a Pylon.
func (g *Game) HasPower(tileX, tileY int) bool {
	return g.hasPowerCheck(tileX*32, tileY*32, UnitTypeNone)
}

// HasPowerForType returns whether a building of the given type can be powered
// at the given tile position. Uses the center of the building footprint.
func (g *Game) HasPowerForType(tileX, tileY int, unitType UnitType) bool {
	if utInRange(unitType) && (!unitType.RequiresPsi() || !unitType.IsBuilding()) {
		return true
	}
	px := tileX*32 + unitType.TileWidth()*16
	py := tileY*32 + unitType.TileHeight()*16
	return g.hasPowerCheck(px, py, unitType)
}

// hasPowerCheck tests if pixel (x,y) falls inside any completed Pylon's power field.
func (g *Game) hasPowerCheck(x, y int, unitType UnitType) bool {
	if utInRange(unitType) && unitType != UnitTypeNone &&
		(!unitType.RequiresPsi() || !unitType.IsBuilding()) {
		return true
	}

	selfIdx := int(g.data.SelfIndex())
	for _, u := range g.GetAllUnits() {
		if u.GetType() != UnitTypeProtossPylon || !u.IsCompleted() {
			continue
		}
		if int(u.data.PlayerIndex()) != selfIdx {
			continue
		}

		px := int(u.data.PositionX())
		py := int(u.data.PositionY())

		dx := x - px
		dy := y - py
		if dx < -256 || dx >= 256 || dy < -160 || dy >= 160 {
			continue
		}

		row := (dy + 160) / 32
		col := (dx + 256) / 32
		if row >= 0 && row < 10 && col >= 0 && col < 16 && bPsiFieldMask[row][col] {
			return true
		}
	}
	return false
}

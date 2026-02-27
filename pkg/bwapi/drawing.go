package bwapi

// CoordinateType specifies the coordinate system for drawing.
type CoordinateType int32

const (
	CoordinateTypeNone   CoordinateType = 0
	CoordinateTypeScreen CoordinateType = 1
	CoordinateTypeMap    CoordinateType = 2
	CoordinateTypeMouse  CoordinateType = 3
)

// ShapeType identifies the type of debug drawing shape.
type ShapeType int32

const (
	ShapeTypeNone     ShapeType = 0
	ShapeTypeText     ShapeType = 1
	ShapeTypeBox      ShapeType = 2
	ShapeTypeTriangle ShapeType = 3
	ShapeTypeCircle   ShapeType = 4
	ShapeTypeEllipse  ShapeType = 5
	ShapeTypeDot      ShapeType = 6
	ShapeTypeLine     ShapeType = 7
)

// DrawTextScreen draws text at screen coordinates.
func (g *Game) DrawTextScreen(x, y int, text string) {
	idx := g.data.AddString(text)
	g.data.AddShape(int32(ShapeTypeText), int32(CoordinateTypeScreen), int32(x), int32(y), 0, 0, int32(idx), 0, 0, false)
}

// DrawTextMap draws text at map pixel coordinates.
func (g *Game) DrawTextMap(x, y int, text string) {
	idx := g.data.AddString(text)
	g.data.AddShape(int32(ShapeTypeText), int32(CoordinateTypeMap), int32(x), int32(y), 0, 0, int32(idx), 0, 0, false)
}

// DrawBoxScreen draws a box at screen coordinates.
func (g *Game) DrawBoxScreen(left, top, right, bottom, color int, isSolid bool) {
	g.data.AddShape(int32(ShapeTypeBox), int32(CoordinateTypeScreen), int32(left), int32(top), int32(right), int32(bottom), 0, 0, int32(color), isSolid)
}

// DrawBoxMap draws a box at map pixel coordinates.
func (g *Game) DrawBoxMap(left, top, right, bottom, color int, isSolid bool) {
	g.data.AddShape(int32(ShapeTypeBox), int32(CoordinateTypeMap), int32(left), int32(top), int32(right), int32(bottom), 0, 0, int32(color), isSolid)
}

// DrawCircleScreen draws a circle at screen coordinates.
func (g *Game) DrawCircleScreen(x, y, radius, color int, isSolid bool) {
	g.data.AddShape(int32(ShapeTypeCircle), int32(CoordinateTypeScreen), int32(x), int32(y), 0, 0, int32(radius), 0, int32(color), isSolid)
}

// DrawCircleMap draws a circle at map pixel coordinates.
func (g *Game) DrawCircleMap(x, y, radius, color int, isSolid bool) {
	g.data.AddShape(int32(ShapeTypeCircle), int32(CoordinateTypeMap), int32(x), int32(y), 0, 0, int32(radius), 0, int32(color), isSolid)
}

// DrawLineScreen draws a line at screen coordinates.
func (g *Game) DrawLineScreen(x1, y1, x2, y2, color int) {
	g.data.AddShape(int32(ShapeTypeLine), int32(CoordinateTypeScreen), int32(x1), int32(y1), int32(x2), int32(y2), 0, 0, int32(color), false)
}

// DrawLineMap draws a line at map pixel coordinates.
func (g *Game) DrawLineMap(x1, y1, x2, y2, color int) {
	g.data.AddShape(int32(ShapeTypeLine), int32(CoordinateTypeMap), int32(x1), int32(y1), int32(x2), int32(y2), 0, 0, int32(color), false)
}

// DrawDotScreen draws a dot at screen coordinates.
func (g *Game) DrawDotScreen(x, y, color int) {
	g.data.AddShape(int32(ShapeTypeDot), int32(CoordinateTypeScreen), int32(x), int32(y), 0, 0, 0, 0, int32(color), false)
}

// DrawDotMap draws a dot at map pixel coordinates.
func (g *Game) DrawDotMap(x, y, color int) {
	g.data.AddShape(int32(ShapeTypeDot), int32(CoordinateTypeMap), int32(x), int32(y), 0, 0, 0, 0, int32(color), false)
}

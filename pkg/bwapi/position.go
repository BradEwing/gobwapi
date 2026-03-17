package bwapi

import "math"

// Position represents pixel coordinates.
type Position struct {
	X, Y int32
}

// WalkPosition represents walk tile coordinates (1 walk tile = 8 pixels).
type WalkPosition struct {
	X, Y int32
}

// TilePosition represents build tile coordinates (1 tile = 32 pixels).
type TilePosition struct {
	X, Y int32
}

func (p Position) ToWalkPosition() WalkPosition {
	return WalkPosition{X: p.X / 8, Y: p.Y / 8}
}

func (p Position) ToTilePosition() TilePosition {
	return TilePosition{X: p.X / 32, Y: p.Y / 32}
}

func (wp WalkPosition) ToPosition() Position {
	return Position{X: wp.X * 8, Y: wp.Y * 8}
}

func (wp WalkPosition) ToTilePosition() TilePosition {
	return TilePosition{X: wp.X / 4, Y: wp.Y / 4}
}

func (tp TilePosition) ToPosition() Position {
	return Position{X: tp.X * 32, Y: tp.Y * 32}
}

func (tp TilePosition) ToWalkPosition() WalkPosition {
	return WalkPosition{X: tp.X * 4, Y: tp.Y * 4}
}

// Center returns the pixel position at the center of this build tile.
func (tp TilePosition) Center() Position {
	return Position{X: tp.X*32 + 16, Y: tp.Y*32 + 16}
}

// Center returns the pixel position at the center of this walk tile.
func (wp WalkPosition) Center() Position {
	return Position{X: wp.X*8 + 4, Y: wp.Y*8 + 4}
}

// GetDistance returns the Euclidean distance to another Position.
func (p Position) GetDistance(other Position) float64 {
	dx := float64(p.X - other.X)
	dy := float64(p.Y - other.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

// GetDistance returns the Euclidean distance to another WalkPosition.
func (wp WalkPosition) GetDistance(other WalkPosition) float64 {
	dx := float64(wp.X - other.X)
	dy := float64(wp.Y - other.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

// GetDistance returns the Euclidean distance to another TilePosition.
func (tp TilePosition) GetDistance(other TilePosition) float64 {
	dx := float64(tp.X - other.X)
	dy := float64(tp.Y - other.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

// GetLength returns the Euclidean distance from the origin.
func (p Position) GetLength() float64 {
	return math.Sqrt(float64(p.X)*float64(p.X) + float64(p.Y)*float64(p.Y))
}

// GetLength returns the Euclidean distance from the origin.
func (wp WalkPosition) GetLength() float64 {
	return math.Sqrt(float64(wp.X)*float64(wp.X) + float64(wp.Y)*float64(wp.Y))
}

// GetLength returns the Euclidean distance from the origin.
func (tp TilePosition) GetLength() float64 {
	return math.Sqrt(float64(tp.X)*float64(tp.X) + float64(tp.Y)*float64(tp.Y))
}

// ChebyshevDistance returns max(|dx|, |dy|) to another Position.
func (p Position) ChebyshevDistance(other Position) int32 {
	dx := p.X - other.X
	dy := p.Y - other.Y
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	if dx > dy {
		return dx
	}
	return dy
}

// ChebyshevDistance returns max(|dx|, |dy|) to another WalkPosition.
func (wp WalkPosition) ChebyshevDistance(other WalkPosition) int32 {
	dx := wp.X - other.X
	dy := wp.Y - other.Y
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	if dx > dy {
		return dx
	}
	return dy
}

// ChebyshevDistance returns max(|dx|, |dy|) to another TilePosition.
func (tp TilePosition) ChebyshevDistance(other TilePosition) int32 {
	dx := tp.X - other.X
	dy := tp.Y - other.Y
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	if dx > dy {
		return dx
	}
	return dy
}

// GetApproxDistance returns StarCraft's built-in fast approximate distance.
// This matches the exact integer result used by the game engine.
func (p Position) GetApproxDistance(other Position) int32 {
	return approxDistance(p.X-other.X, p.Y-other.Y)
}

// GetApproxDistance returns StarCraft's built-in fast approximate distance
// in walk tile units.
func (wp WalkPosition) GetApproxDistance(other WalkPosition) int32 {
	return approxDistance(wp.X-other.X, wp.Y-other.Y)
}

// GetApproxDistance returns StarCraft's built-in fast approximate distance
// in build tile units.
func (tp TilePosition) GetApproxDistance(other TilePosition) int32 {
	return approxDistance(tp.X-other.X, tp.Y-other.Y)
}

// approxDistance computes SC:BW's integer approximate Euclidean distance.
func approxDistance(dx, dy int32) int32 {
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	if dx < dy {
		dx, dy = dy, dx
	}
	if dy <= (dx >> 2) {
		return dx
	}
	minCalc := (3 * dy) >> 3
	return (minCalc >> 5) + minCalc + dx - (dx >> 4) - (dx >> 6)
}

// Add returns the component-wise sum of two Positions.
func (p Position) Add(other Position) Position {
	return Position{X: p.X + other.X, Y: p.Y + other.Y}
}

// Sub returns the component-wise difference of two Positions.
func (p Position) Sub(other Position) Position {
	return Position{X: p.X - other.X, Y: p.Y - other.Y}
}

// Mul returns the Position scaled by n.
func (p Position) Mul(n int32) Position {
	return Position{X: p.X * n, Y: p.Y * n}
}

// Div returns the Position divided by n. Panics if n is zero.
func (p Position) Div(n int32) Position {
	return Position{X: p.X / n, Y: p.Y / n}
}

// Add returns the component-wise sum of two WalkPositions.
func (wp WalkPosition) Add(other WalkPosition) WalkPosition {
	return WalkPosition{X: wp.X + other.X, Y: wp.Y + other.Y}
}

// Sub returns the component-wise difference of two WalkPositions.
func (wp WalkPosition) Sub(other WalkPosition) WalkPosition {
	return WalkPosition{X: wp.X - other.X, Y: wp.Y - other.Y}
}

// Mul returns the WalkPosition scaled by n.
func (wp WalkPosition) Mul(n int32) WalkPosition {
	return WalkPosition{X: wp.X * n, Y: wp.Y * n}
}

// Div returns the WalkPosition divided by n. Panics if n is zero.
func (wp WalkPosition) Div(n int32) WalkPosition {
	return WalkPosition{X: wp.X / n, Y: wp.Y / n}
}

// Add returns the component-wise sum of two TilePositions.
func (tp TilePosition) Add(other TilePosition) TilePosition {
	return TilePosition{X: tp.X + other.X, Y: tp.Y + other.Y}
}

// Sub returns the component-wise difference of two TilePositions.
func (tp TilePosition) Sub(other TilePosition) TilePosition {
	return TilePosition{X: tp.X - other.X, Y: tp.Y - other.Y}
}

// Mul returns the TilePosition scaled by n.
func (tp TilePosition) Mul(n int32) TilePosition {
	return TilePosition{X: tp.X * n, Y: tp.Y * n}
}

// Div returns the TilePosition divided by n. Panics if n is zero.
func (tp TilePosition) Div(n int32) TilePosition {
	return TilePosition{X: tp.X / n, Y: tp.Y / n}
}

// IsValid returns whether this Position is within the map bounds.
func (p Position) IsValid(g *Game) bool {
	return p.X >= 0 && p.Y >= 0 &&
		p.X < int32(g.MapWidth())*32 && p.Y < int32(g.MapHeight())*32
}

// IsValid returns whether this WalkPosition is within the map bounds.
func (wp WalkPosition) IsValid(g *Game) bool {
	return wp.X >= 0 && wp.Y >= 0 &&
		wp.X < int32(g.MapWidth())*4 && wp.Y < int32(g.MapHeight())*4
}

// IsValid returns whether this TilePosition is within the map bounds.
func (tp TilePosition) IsValid(g *Game) bool {
	return tp.X >= 0 && tp.Y >= 0 &&
		tp.X < int32(g.MapWidth()) && tp.Y < int32(g.MapHeight())
}

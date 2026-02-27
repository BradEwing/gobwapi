package bwapi

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

package bwem

// Algorithm constants matching BWEM C++/Rust originals.
const (
	lakeMaxMiniTiles        = 300
	lakeMaxWidthMiniTiles   = 32
	areaMinMiniTiles        = 64
	smallAreaMaxMiniTiles   = 80
	altitudeMergeRatio      = 0.90
	minTilesBetweenBases       = 10
	maxTilesBetweenCCAndRes    = 10
	maxTilesBetweenStartAndBase = 3
	resourceExclusionGap       = 3
	chokeClusterDistSq         = 300

	ccTileWidth  = 4
	ccTileHeight = 3
)

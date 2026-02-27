// Package shm provides CGo bindings for the BWAPI shared memory protocol.
package shm

const (
	// BWAPIVersion is the client/server protocol version (BWAPI 4.4.0).
	BWAPIVersion = 10003

	// GameDataSize is the expected size of the GameData shared memory region.
	GameDataSize = 33_017_048

	// GameTableSize is the expected size of the GameTable shared memory region.
	GameTableSize = 96

	// Array size constants matching BWAPI 4.4.0.
	MaxForces         = 5
	MaxPlayers        = 12
	MaxUnits          = 10000
	MaxBullets        = 100
	MaxRegions        = 5000
	MaxUnitTypes      = 234
	MaxUpgradeTypes   = 63
	MaxTechTypes      = 47
	MaxEvents         = 10000
	MaxEventStrings   = 1000
	MaxStrings        = 20000
	MaxShapes         = 20000
	MaxCommands       = 20000
	MaxUnitCommands   = 20000
	MaxNukeDots       = 200
	MaxStartLocations = 8
	MaxGameInstances  = 8
	TrainingQueueSize = 5
	MaxNeighbors      = 256
	MapMaxTileSize    = 256
	MapMaxWalkSize    = 1024
	UnitArraySize     = 1700
	UnitSearchSize    = 3400
	KMax              = 255
	MMax              = 3
	FlagMax           = 2
	MaxSelectedUnits  = 12

	// String buffer sizes.
	ForceNameSize   = 32
	PlayerNameSize  = 25
	MapFileNameSize = 261
	MapPathNameSize = 261
	MapNameSize     = 33
	MapHashSize     = 41
	StringSize      = 1024
	EventStringSize = 256

	// Sync protocol bytes.
	SyncSend    byte = 0x01
	SyncReceive byte = 0x02
)

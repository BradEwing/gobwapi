# CLAUDE.md — gobwapi

## Project Overview

Go bindings for BWAPI (Brood War API) — a C++ library for writing StarCraft: Brood War AI bots.
Uses CGo to interface with BWAPI's shared memory client protocol.

## Tech Stack

- **Language**: Go (1.21+)
- **Build**: CGo with GCC toolchain
- **Protocol**: BWAPI 4.4.0 shared memory IPC (version ID: 10003)
- **Platforms**: Windows (native BWAPI), Linux (OpenBW via POSIX shm)

## Architecture

### BWAPI Client/Server Model

BWAPI injects a DLL into StarCraft that exposes game state via a ~33 MB shared memory region
(`GameData`). This library is a **client** that maps that shared memory to read game state and
write commands. Synchronization is lock-step per frame via named pipes (Windows) or Unix domain
sockets (Linux/OpenBW).

### Binding Strategy

Use CGo to:
1. Define struct layouts matching BWAPI's `GameData.h` shared memory format
2. Handle platform-specific shared memory mapping (Win32 / POSIX `shm_open`)
3. Handle pipe/socket synchronization

Game logic types (`Game`, `Unit`, `Player`, etc.) are implemented in pure Go on top of the
mapped memory.

### Key Reference Files

When implementing, reference these source files from other BWAPI implementations:

- **Shared memory layout**: `bwapi/include/BWAPI/Client/GameData.h` (C++ original)
- **Data structs**: `bwapi/include/BWAPI/Client/UnitData.h`, `PlayerData.h`, `BulletData.h`, `RegionData.h`
- **Client connection**: JBWAPI `src/main/java/bwapi/Client.java`, `ClientData.java`
- **Rust struct bindings**: rsbwapi `bwapi_wrapper/src/lib.rs`
- **Unit type data**: rsbwapi `bwapi_wrapper/src/unit_type.rs`
- **Game logic**: rsbwapi `src/game.rs`, JBWAPI `src/main/java/bwapi/Game.java`

### Shared Memory Constants

```
GameTable name (Windows): "Local\\bwapi_shared_memory_game_list"
GameTable name (POSIX):   "/bwapi_shared_memory_game_list"
GameData name (Windows):  "Local\\bwapi_shared_memory_{pid}"
GameData name (POSIX):    "/bwapi_shared_memory_{pid}"
Pipe (Windows):           "\\\\.\\pipe\\bwapi_pipe_{pid}"
Socket (POSIX):           "/tmp/bwapi_socket_{pid}"
Sync protocol:            Client sends 0x01, server responds 0x02
BWAPI version:            10003 (4.4.0)
GameData size:            ~33,017,048 bytes
GameTable:                8 instances × 12 bytes
```

### Core Type Hierarchy

```
AIModule (interface)      — 17 event callbacks (OnStart, OnFrame, OnUnitCreate, etc.)
Game                      — Central API, wraps GameData shared memory
  ├── Unit                — 336 bytes per unit, max 10,000
  ├── Player              — 5,788 bytes per player, max 12
  ├── Bullet              — Projectile state, max 100
  ├── Region              — Pathfinding region, max 5,000
  └── Force               — Alliance grouping, max 5
Position types:
  ├── Position            — Pixel coordinates (scale 1)
  ├── WalkPosition        — Walk tiles (scale 8, = Position/8)
  └── TilePosition        — Build tiles (scale 32, = Position/32)
Enum types:
  ├── UnitType            — 234 unit types
  ├── Race                — Zerg, Terran, Protoss
  ├── Order               — Unit orders
  ├── TechType            — Research technologies
  ├── UpgradeType         — Unit upgrades
  └── WeaponType          — Weapon stats
```

## Conventions

- Module import path: `github.com/bradewing/gobwapi`
- Follow standard Go project layout (`pkg/`, `cmd/`, `internal/`)
- Exported Go types use PascalCase matching BWAPI naming (e.g., `UnitType`, `TilePosition`)
- BWAPI enum values use Go constants with the type prefix (e.g., `UnitTypeTerranMarine`)
- The game loop MUST run on a single goroutine — BWAPI is not thread-safe
- Use `runtime.LockOSThread()` for the game loop goroutine when shared memory APIs require it
- Platform-specific code goes in `_windows.go` / `_linux.go` build-tagged files

## Commands

```bash
# Build
go build ./...

# Test
go test ./...

# Vet
go vet ./...
```

## Reference Implementations

- **C++ (original)**: https://github.com/bwapi/bwapi
- **Rust**: https://github.com/Bytekeeper/rsbwapi
- **Java**: https://github.com/JavaBWAPI/JBWAPI

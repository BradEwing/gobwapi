# gobwapi

Go bindings for [BWAPI](https://github.com/bwapi/bwapi) (Brood War API), enabling StarCraft: Brood War AI development in Go.

## Overview

gobwapi is a Go library that wraps BWAPI's client protocol using CGo bindings, allowing you to write Brood War bots in Go. It connects to the BWAPI server (injected into StarCraft or running via [OpenBW](https://github.com/OpenBW/openbw)) through shared memory and named pipes/Unix domain sockets.

## Architecture

BWAPI uses a **client/server architecture** with shared memory IPC:

1. **Server**: A DLL injected into StarCraft (or OpenBW) that exposes game state via a ~33 MB shared memory region (`GameData`)
2. **Client** (this library): Maps the shared memory, reads game state, and writes commands back each frame
3. **Synchronization**: Lock-step frame-by-frame via named pipes (Windows) or Unix domain sockets (Linux/OpenBW)

### Connection Protocol

A secondary shared memory region (`GameTable`) holds up to 8 game instances. The client scans this table to find an available server, then opens the corresponding shared memory and pipe/socket for that process.

```
GameTable:  "Local\bwapi_shared_memory_game_list"    (Windows)
            "/bwapi_shared_memory_game_list"          (POSIX)

GameData:   "Local\bwapi_shared_memory_{pid}"         (Windows)
            "/bwapi_shared_memory_{pid}"              (POSIX)

Pipe:       "\\.\pipe\bwapi_pipe_{pid}"               (Windows)
Socket:     "/tmp/bwapi_socket_{pid}"                 (POSIX)

Sync:       Client writes 0x01 → Server responds 0x02
```

## Core Types

| Type | Description |
|------|-------------|
| `Game` | Central interface — map queries, unit collections, build validation, drawing commands |
| `Unit` | Individual unit state and commands (attack, move, build, train, etc.) |
| `Player` | Player resources, upgrades, tech, and unit counts |
| `UnitType` | Enum of all 234 unit types with stats (costs, build time, abilities) |
| `Position` | Pixel coordinates (x, y) |
| `WalkPosition` | Walk tile coordinates (Position / 8) |
| `TilePosition` | Build tile coordinates (Position / 32) |
| `Region` | Pathfinding regions from BW's internal navigation mesh |
| `Bullet` | Projectile state |

## Shared Memory Layout

The `GameData` struct (~33 MB) contains fixed-size arrays:

| Data | Max Count | Per-Item Size |
|------|-----------|---------------|
| Units | 10,000 | 336 bytes |
| Players | 12 | 5,788 bytes |
| Bullets | 100 | small |
| Regions | 5,000 | small |
| Forces | 5 | small |
| Events | 10,000 | small |
| Commands | 20,000 | small |
| Map tiles | 256x256 / 1024x1024 | varies |

## Usage

```go
package main

import "github.com/bradewing/gobwapi"

type MyBot struct{}

func (b *MyBot) OnStart(game *gobwapi.Game) {
    for _, loc := range game.GetStartLocations() {
        fmt.Printf("Start location: %v\n", loc)
    }
}

func (b *MyBot) OnFrame(game *gobwapi.Game) {
    player := game.Self()
    for _, unit := range player.GetUnits() {
        // Bot logic here
    }
}

func main() {
    gobwapi.Start(&MyBot{})
}
```

## Event Callbacks

Implement the `AIModule` interface to handle game events:

- `OnStart` / `OnEnd` — Game lifecycle
- `OnFrame` — Called each game frame (~42ms at fastest speed)
- `OnUnitCreate` / `OnUnitDestroy` / `OnUnitMorph` / `OnUnitComplete` — Unit lifecycle
- `OnUnitShow` / `OnUnitHide` / `OnUnitRenegade` — Visibility and ownership changes
- `OnSendText` / `OnReceiveText` — Chat
- `OnNukeDetect` — Nuke launch detection
- `OnPlayerLeft` — Player disconnection
- `OnSaveGame` — Game saved

## Building

Requires:
- Go 1.21+
- GCC/MinGW (for CGo on Windows) or GCC (Linux)
- BWAPI 4.4.0 headers (included as a git submodule)

```bash
go build ./...
```

## Running

1. Start StarCraft: Brood War with BWAPI injected, **or** start an OpenBW instance
2. Run your bot binary — it will connect to the first available BWAPI server

## Reference Implementations

- [bwapi/bwapi](https://github.com/bwapi/bwapi) — Original C++ BWAPI
- [Bytekeeper/rsbwapi](https://github.com/Bytekeeper/rsbwapi) — Rust BWAPI client
- [JavaBWAPI/JBWAPI](https://github.com/JavaBWAPI/JBWAPI) — Pure Java BWAPI client

## License

MIT

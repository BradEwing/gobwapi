package bwapi

import (
	"log"
	"runtime"

	"github.com/bradewing/gobwapi/internal/client"
)

// BWClient is the top-level entry point for running a BWAPI bot.
type BWClient struct {
	client *client.Client
	game   *Game
}

// NewBWClient creates a new BWClient.
func NewBWClient() *BWClient {
	return &BWClient{
		client: client.NewClient(),
	}
}

// Run connects to BWAPI and runs the game loop, dispatching events to the given module.
// This function blocks until the connection is permanently lost.
// The game loop runs on a locked OS thread as required by BWAPI.
func (bw *BWClient) Run(module AIModule) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	log.Println("Waiting for BWAPI server...")
	bw.client.Reconnect()

	bw.game = NewGame(bw.client.Data())

	for {
		// Dispatch all events for this frame.
		dispatchEvents(bw.game, module)

		// Clear command buffers for next frame.
		bw.game.data.ResetCommands()

		// If no longer in game, reconnect for the next game.
		if !bw.game.IsInGame() {
			bw.client.Disconnect()
			log.Println("Game ended. Waiting for next game...")
			bw.client.Reconnect()
			bw.game = NewGame(bw.client.Data())
			continue
		}

		// Synchronize with the server: send commands, wait for next frame.
		if err := bw.client.Update(); err != nil {
			log.Printf("Connection lost: %v", err)
			bw.client.Disconnect()
			log.Println("Reconnecting...")
			bw.client.Reconnect()
			bw.game = NewGame(bw.client.Data())
		}
	}
}

// Game returns the current Game instance. Only valid during callbacks.
func (bw *BWClient) Game() *Game {
	return bw.game
}

// Close releases all resources.
func (bw *BWClient) Close() {
	bw.client.Close()
}

// Package client handles the BWAPI shared memory connection and frame synchronization.
package client

import (
	"errors"
	"fmt"
	"io"
	"log"
	"time"
	"unsafe"

	"github.com/bradewing/gobwapi/internal/shm"
)

// Client manages the connection to a BWAPI server instance.
type Client struct {
	gameTable      *shm.GameTable
	gameTableClose func()

	data      *shm.GameData
	dataClose func()

	pipe io.ReadWriteCloser

	connected bool
	serverPID uint32
}

// NewClient creates a new unconnected BWAPI client.
func NewClient() *Client {
	return &Client{}
}

// IsConnected returns whether the client is connected to a server.
func (c *Client) IsConnected() bool {
	return c.connected
}

// Data returns the shared GameData. Only valid when connected.
func (c *Client) Data() *shm.GameData {
	return c.data
}

// Connect scans the game table and connects to an available server instance.
func (c *Client) Connect() error {
	if c.connected {
		return nil
	}

	if c.gameTable == nil {
		ptr, cleanup, err := mapGameTable()
		if err != nil {
			return fmt.Errorf("map game table: %w", err)
		}
		c.gameTable = shm.CastGameTable(ptr)
		c.gameTableClose = cleanup
	}

	bestIdx := -1
	var bestTime uint32 = ^uint32(0)

	for i := 0; i < shm.MaxGameInstances; i++ {
		inst := c.gameTable.Instance(i)
		pid := inst.ServerProcessID()
		if pid == 0 {
			continue
		}
		if inst.IsConnected() {
			continue
		}
		if t := inst.LastKeepAliveTime(); t < bestTime {
			bestTime = t
			bestIdx = i
		}
	}

	if bestIdx < 0 {
		return errors.New("no available BWAPI server instance found")
	}

	inst := c.gameTable.Instance(bestIdx)
	c.serverPID = inst.ServerProcessID()

	dataPtr, dataClose, err := mapGameData(c.serverPID)
	if err != nil {
		return fmt.Errorf("map game data for PID %d: %w", c.serverPID, err)
	}
	c.data = shm.CastGameData(dataPtr)
	c.dataClose = dataClose

	if v := c.data.ClientVersion(); v != shm.BWAPIVersion {
		c.dataClose()
		c.data = nil
		c.dataClose = nil
		return fmt.Errorf("BWAPI version mismatch: got %d, want %d", v, shm.BWAPIVersion)
	}

	pipe, err := dialPipe(c.serverPID)
	if err != nil {
		c.dataClose()
		c.data = nil
		c.dataClose = nil
		return fmt.Errorf("dial pipe for PID %d: %w", c.serverPID, err)
	}
	c.pipe = pipe

	if err := c.waitForServer(); err != nil {
		c.pipe.Close()
		c.dataClose()
		c.pipe = nil
		c.data = nil
		c.dataClose = nil
		return fmt.Errorf("initial handshake: %w", err)
	}

	inst.SetIsConnected(true)
	c.connected = true

	log.Printf("Connected to BWAPI server PID %d", c.serverPID)
	return nil
}

// Disconnect closes the connection to the server.
func (c *Client) Disconnect() {
	if !c.connected {
		return
	}

	if c.pipe != nil {
		c.pipe.Close()
		c.pipe = nil
	}
	if c.dataClose != nil {
		c.dataClose()
		c.dataClose = nil
	}
	c.data = nil
	c.connected = false

	log.Printf("Disconnected from BWAPI server PID %d", c.serverPID)
}

// Reconnect blocks until a connection is established, retrying periodically.
func (c *Client) Reconnect() {
	for {
		err := c.Connect()
		if err == nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

// Update performs one frame synchronization cycle:
// sends the client-ready signal and waits for the server to complete the next frame.
func (c *Client) Update() error {
	if !c.connected {
		return errors.New("not connected")
	}
	if err := c.sendClientData(); err != nil {
		return fmt.Errorf("send client data: %w", err)
	}
	if err := c.waitForServer(); err != nil {
		return fmt.Errorf("wait for server: %w", err)
	}
	return nil
}

// Close releases all resources.
func (c *Client) Close() {
	c.Disconnect()
	if c.gameTableClose != nil {
		c.gameTableClose()
		c.gameTableClose = nil
	}
	c.gameTable = nil
}

// sendClientData writes the sync byte to the pipe.
func (c *Client) sendClientData() error {
	_, err := c.pipe.Write([]byte{shm.SyncSend})
	return err
}

// waitForServer reads from the pipe until the server sync byte is received.
func (c *Client) waitForServer() error {
	buf := make([]byte, 1)
	for {
		n, err := c.pipe.Read(buf)
		if err != nil {
			return err
		}
		if n > 0 && buf[0] == shm.SyncReceive {
			return nil
		}
	}
}

// GameTableSize returns the size of GameTable for external use.
func GameTableSize() int {
	return int(unsafe.Sizeof(shm.GameTable{}))
}

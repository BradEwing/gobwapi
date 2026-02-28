package main

import (
	"fmt"
	"log"

	"github.com/bradewing/gobwapi/pkg/bwapi"
)

// ExampleBot implements a simple resource-gathering AI.
type ExampleBot struct {
	bwapi.BaseModule
}

func (b *ExampleBot) OnStart(game *bwapi.Game) {
	log.Printf("Game started on map: %s (%dx%d)", game.MapName(), game.MapWidth(), game.MapHeight())

	self := game.Self()
	if self != nil {
		log.Printf("Playing as: %s (%s)", self.Name(), self.GetRace())
	}

	game.EnableFlag(0)
	game.EnableFlag(1)

	game.SetLocalSpeed(20)
}

func (b *ExampleBot) OnEnd(game *bwapi.Game, isWinner bool) {
	if isWinner {
		log.Println("I won!")
	} else {
		log.Println("I lost!")
	}
}

func (b *ExampleBot) OnFrame(game *bwapi.Game) {
	self := game.Self()
	if self == nil {
		return
	}

	game.DrawTextScreen(10, 10, fmt.Sprintf("Frame: %d", game.FrameCount()))
	game.DrawTextScreen(10, 20, fmt.Sprintf("Minerals: %d  Gas: %d", self.Minerals(), self.Gas()))
	game.DrawTextScreen(10, 30, fmt.Sprintf("Supply: %d/%d", self.SupplyUsed(), self.SupplyTotal()))

	units := game.GetAllUnits()
	for _, unit := range units {
		if unit.GetPlayer() == nil || unit.GetPlayer().Index() != self.Index() {
			continue
		}
		if !unit.IsIdle() || !unit.IsCompleted() {
			continue
		}

		unitType := unit.GetType()
		isWorker := unitType == bwapi.UnitTypeTerranSCV ||
			unitType == bwapi.UnitTypeProtossProbe ||
			unitType == bwapi.UnitTypeZergDrone

		if !isWorker {
			continue
		}

		var bestMineral *bwapi.Unit
		bestDist := int32(1<<31 - 1)
		for _, other := range units {
			if other.GetType() != bwapi.UnitTypeResourceMineralField &&
				other.GetType() != bwapi.UnitTypeResourceMineralFieldType2 &&
				other.GetType() != bwapi.UnitTypeResourceMineralFieldType3 {
				continue
			}
			dx := unit.GetPosition().X - other.GetPosition().X
			dy := unit.GetPosition().Y - other.GetPosition().Y
			dist := dx*dx + dy*dy
			if dist < bestDist {
				bestDist = dist
				bestMineral = other
			}
		}

		if bestMineral != nil {
			unit.Gather(bestMineral)
		}
	}
}

func (b *ExampleBot) OnUnitCreate(game *bwapi.Game, unit *bwapi.Unit) {
	self := game.Self()
	if self != nil && unit.GetPlayer() != nil && unit.GetPlayer().Index() == self.Index() {
		log.Printf("Unit created: %s", unit.GetType())
	}
}

func main() {
	log.Println("Starting example bot...")
	client := bwapi.NewBWClient()
	defer client.Close()
	client.Run(&ExampleBot{})
}

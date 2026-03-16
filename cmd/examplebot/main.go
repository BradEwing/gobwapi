package main

import (
	"fmt"
	"log"

	"github.com/bradewing/gobwapi/pkg/bwapi"
	"github.com/bradewing/gobwapi/pkg/bwem"
)

// ExampleBot implements a simple resource-gathering AI.
type ExampleBot struct {
	bwapi.BaseModule
	bwem *bwem.Map
}

func (b *ExampleBot) OnStart(game *bwapi.Game) {
	log.Printf("Game started on map: %s (%dx%d)", game.MapName(), game.MapWidth(), game.MapHeight())

	self := game.Self()
	if self != nil {
		log.Printf("Playing as: %s (%s)", self.Name(), self.GetRace())
	}

	game.EnableFlag(0) // UserInput — allow user keyboard/mouse input

	game.SetLocalSpeed(20)

	allUnits := game.GetAllUnits()
	log.Printf("DEBUG: GetAllUnits returned %d units", len(allUnits))
	typeCounts := make(map[bwapi.UnitType]int)
	for _, u := range allUnits {
		typeCounts[u.GetType()]++
	}
	for ut, count := range typeCounts {
		log.Printf("  UnitType %d (%s): %d", int(ut), ut, count)
	}

	b.bwem = bwem.Analyze(game)
	log.Printf("BWEM: %d areas, %d chokepoints, %d bases",
		len(b.bwem.Areas()), len(b.bwem.ChokePoints()), len(b.bwem.Bases()))
	log.Printf("BWEM: %d minerals, %d geysers, %d neutrals",
		len(b.bwem.Minerals()), len(b.bwem.Geysers()), len(b.bwem.Neutrals()))
	for _, area := range b.bwem.Areas() {
		log.Printf("  Area %d: %d minitiles, %d minerals, %d geysers, %d bases",
			area.ID, area.MiniTileCount, len(area.MineralIdxs), len(area.GeyserIdxs), len(area.BaseIdxs))
	}
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

	b.drawBases(game)
	b.drawDepots(game)

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

func (b *ExampleBot) drawBases(game *bwapi.Game) {
	if b.bwem == nil {
		return
	}
	for _, base := range b.bwem.Bases() {
		left := int(base.Location.X) * 32
		top := int(base.Location.Y) * 32
		right := left + 4*32
		bottom := top + 3*32
		color := bwapi.ColorGreen
		if base.IsStartLocation {
			color = bwapi.ColorPurple
		}
		game.DrawBoxMap(left, top, right, bottom, color, false)
		game.DrawTextMap(left+4, top+4, fmt.Sprintf("Base (area %d)", base.AreaID))
	}
}

func (b *ExampleBot) drawDepots(game *bwapi.Game) {
	for _, unit := range game.GetAllUnits() {
		ut := unit.GetType()
		if ut != bwapi.UnitTypeTerranCommandCenter &&
			ut != bwapi.UnitTypeProtossNexus &&
			ut != bwapi.UnitTypeZergHatchery {
			continue
		}
		pos := unit.GetPosition()
		left := int(pos.X) - 4*16
		top := int(pos.Y) - 3*16
		game.DrawBoxMap(left, top, left+4*32, top+3*32, bwapi.ColorCyan, false)
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

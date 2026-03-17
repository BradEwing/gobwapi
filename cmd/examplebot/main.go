package main

import (
	"fmt"
	"log"
	"math"

	"github.com/bradewing/gobwapi/pkg/bwapi"
	"github.com/bradewing/gobwapi/pkg/bwem"
)

// ExampleBot implements a 4-pool zergling rush.
type ExampleBot struct {
	bwapi.BaseModule
	bwem *bwem.Map

	poolOrdered      bool                // true once we've issued a build-pool command
	enemyBase        *bwapi.TilePosition // discovered enemy start location
	scoutTargets     []bwapi.TilePosition
	scoutAssignments map[int]int // unit index -> scoutTargets index
}

func (b *ExampleBot) OnStart(game *bwapi.Game) {
	log.Printf("Game started on map: %s (%dx%d)", game.MapName(), game.MapWidth(), game.MapHeight())

	self := game.Self()
	if self != nil {
		log.Printf("Playing as: %s (%s)", self.Name(), self.GetRace())
	}

	game.EnableFlag(0) // UserInput
	game.SetLocalSpeed(20)

	b.bwem = bwem.Analyze(game)
	log.Printf("BWEM: %d areas, %d chokepoints, %d bases",
		len(b.bwem.Areas()), len(b.bwem.ChokePoints()), len(b.bwem.Bases()))

	// Initialize scout targets: all start locations except our own.
	myStart := self.GetStartLocation()
	for _, loc := range game.GetStartLocations() {
		if loc.X == myStart.X && loc.Y == myStart.Y {
			continue
		}
		b.scoutTargets = append(b.scoutTargets, loc)
	}
	b.scoutAssignments = make(map[int]int)
	log.Printf("Scout targets: %d enemy start locations", len(b.scoutTargets))
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

	var (
		drones    []*bwapi.Unit
		larva     []*bwapi.Unit
		zerglings []*bwapi.Unit
		overlords []*bwapi.Unit
		hatchery  *bwapi.Unit
		poolCount int // spawning pools (building or complete)
		enemies   []*bwapi.Unit
		minerals  []*bwapi.Unit
	)

	units := game.GetAllUnits()
	for _, u := range units {
		ut := u.GetType()
		player := u.GetPlayer()

		if player != nil && player.Index() == self.Index() {
			switch ut {
			case bwapi.UnitTypeZergDrone:
				drones = append(drones, u)
			case bwapi.UnitTypeZergLarva:
				larva = append(larva, u)
			case bwapi.UnitTypeZergZergling:
				zerglings = append(zerglings, u)
			case bwapi.UnitTypeZergOverlord:
				overlords = append(overlords, u)
			case bwapi.UnitTypeZergHatchery:
				if hatchery == nil {
					hatchery = u
				}
			case bwapi.UnitTypeZergSpawningPool:
				poolCount++
			}
		} else if player != nil && player.Index() != self.Index() && !player.IsNeutral() {
			enemies = append(enemies, u)
		}

		// Mineral fields (neutral).
		if ut == bwapi.UnitTypeResourceMineralField ||
			ut == bwapi.UnitTypeResourceMineralFieldType2 ||
			ut == bwapi.UnitTypeResourceMineralFieldType3 {
			minerals = append(minerals, u)
		}
	}

	poolComplete := self.CompletedUnitCount(bwapi.UnitTypeZergSpawningPool) > 0

	game.DrawTextScreen(10, 40, fmt.Sprintf("Drones: %d  Lings: %d  Larva: %d  Pool: %d",
		len(drones), len(zerglings), len(larva), poolCount))
	if b.enemyBase != nil {
		game.DrawTextScreen(10, 50, fmt.Sprintf("Enemy base: (%d, %d)", b.enemyBase.X, b.enemyBase.Y))
	}

	if poolCount == 0 && !b.poolOrdered && self.Minerals() >= 200 && hatchery != nil {
		// Find an idle drone to build.
		var builder *bwapi.Unit
		for _, d := range drones {
			if d.IsGathering() {
				builder = d
				break
			}
		}
		if builder != nil {
			loc, ok := game.GetBuildLocation(bwapi.UnitTypeZergSpawningPool, hatchery.GetTilePosition(), 15, false, builder)
			if ok {
				builder.Build(bwapi.UnitTypeZergSpawningPool, loc)
				b.poolOrdered = true
				log.Printf("Building spawning pool at (%d, %d)", loc.X, loc.Y)
			}
		}
	}

	if len(drones) < 4 && self.Minerals() >= 50 && len(larva) > 0 {
		larva[0].Morph(bwapi.UnitTypeZergDrone)
		larva = larva[1:]
	}

	if poolComplete && self.Minerals() >= 50 && self.SupplyUsed() < self.SupplyTotal() {
		for _, l := range larva {
			if self.Minerals() >= 50 {
				l.Morph(bwapi.UnitTypeZergZergling)
			}
		}
	}

	b.handleScouting(game, overlords, zerglings, enemies)

	b.handleAttack(game, zerglings, enemies)

	for _, d := range drones {
		if !d.IsIdle() || !d.IsCompleted() {
			continue
		}
		best := nearestMineral(d, minerals)
		if best != nil {
			d.Gather(best)
		}
	}
}

func (b *ExampleBot) handleScouting(game *bwapi.Game, overlords, zerglings []*bwapi.Unit, enemies []*bwapi.Unit) {
	if len(b.scoutTargets) == 0 {
		return
	}

	// Check if any scout target has been reached / enemy found.
	b.checkScoutResults(game, enemies)

	// Assign overlord to first unscouted target.
	if len(overlords) > 0 {
		ov := overlords[0]
		targetIdx, ok := b.scoutAssignments[ov.Index()]
		if !ok {
			targetIdx = b.nextUnassignedTarget(ov.Index())
			if targetIdx >= 0 {
				b.scoutAssignments[ov.Index()] = targetIdx
			}
		}
		if targetIdx >= 0 && targetIdx < len(b.scoutTargets) && ov.IsIdle() {
			target := b.scoutTargets[targetIdx].ToPosition()
			ov.Move(target)
		}
	}

	// Assign first zergling to scout a different target.
	if len(zerglings) > 0 && b.enemyBase == nil {
		scout := zerglings[0]
		targetIdx, ok := b.scoutAssignments[scout.Index()]
		if !ok {
			targetIdx = b.nextUnassignedTarget(scout.Index())
			if targetIdx >= 0 {
				b.scoutAssignments[scout.Index()] = targetIdx
			}
		}
		if targetIdx >= 0 && targetIdx < len(b.scoutTargets) && scout.IsIdle() {
			target := b.scoutTargets[targetIdx].ToPosition()
			scout.Move(target)
		}
	}
}

func (b *ExampleBot) checkScoutResults(game *bwapi.Game, enemies []*bwapi.Unit) {
	// Check if we can see enemy buildings at any scout target.
	for i, target := range b.scoutTargets {
		targetPos := target.ToPosition()
		for _, e := range enemies {
			if e.GetPosition().GetDistance(targetPos) < 600 {
				b.enemyBase = &b.scoutTargets[i]
				log.Printf("Enemy base found at (%d, %d)", target.X, target.Y)
				return
			}
		}
	}

	// If there's only one target left, assume it's the enemy base.
	if len(b.scoutTargets) > 1 {
		// Remove targets we've explored but found empty.
		remaining := b.scoutTargets[:0]
		for _, target := range b.scoutTargets {
			tp := target.ToPosition().ToTilePosition()
			if !game.IsExplored(int(tp.X), int(tp.Y)) {
				remaining = append(remaining, target)
			}
		}
		if len(remaining) == 0 {
			// All explored, none had enemies — shouldn't happen, but keep last target.
			remaining = b.scoutTargets[:1]
		}
		b.scoutTargets = remaining
	}
	if len(b.scoutTargets) == 1 && b.enemyBase == nil {
		// Check if we've explored all other locations — this must be it.
		b.enemyBase = &b.scoutTargets[0]
		log.Printf("Enemy base deduced at (%d, %d)", b.scoutTargets[0].X, b.scoutTargets[0].Y)
	}
}

func (b *ExampleBot) nextUnassignedTarget(excludeUnit int) int {
	assigned := make(map[int]bool)
	for unitIdx, targetIdx := range b.scoutAssignments {
		if unitIdx != excludeUnit {
			assigned[targetIdx] = true
		}
	}
	for i := range b.scoutTargets {
		if !assigned[i] {
			return i
		}
	}
	// All assigned — return first target as fallback.
	if len(b.scoutTargets) > 0 {
		return 0
	}
	return -1
}

func (b *ExampleBot) handleAttack(game *bwapi.Game, zerglings []*bwapi.Unit, enemies []*bwapi.Unit) {
	if b.enemyBase == nil {
		return
	}

	attackPos := b.enemyBase.ToPosition()

	for i, ling := range zerglings {
		// Skip the scouting zergling (index 0) if enemy base not yet confirmed by vision.
		if i == 0 && len(enemies) == 0 {
			continue
		}

		if !ling.IsIdle() {
			continue
		}

		// If enemies nearby, focus-fire the nearest one.
		nearest := nearestUnit(ling, enemies)
		if nearest != nil {
			if ling.GetPosition().GetDistance(nearest.GetPosition()) < 300 {
				ling.Attack(nearest)
				continue
			}
		}

		// Otherwise attack-move toward enemy base.
		ling.AttackMove(attackPos)
	}
}

func (b *ExampleBot) OnUnitDestroy(game *bwapi.Game, unit *bwapi.Unit) {
	self := game.Self()
	if self == nil {
		return
	}
	if unit.GetPlayer() != nil && unit.GetPlayer().Index() == self.Index() {
		if unit.GetType() == bwapi.UnitTypeZergSpawningPool {
			b.poolOrdered = false
			log.Println("Spawning pool destroyed, will rebuild")
		}
		// Clean up scout assignment if this unit was scouting.
		delete(b.scoutAssignments, unit.Index())
	}
}

func nearestUnit(from *bwapi.Unit, candidates []*bwapi.Unit) *bwapi.Unit {
	var best *bwapi.Unit
	bestDist := math.MaxFloat64
	pos := from.GetPosition()
	for _, c := range candidates {
		d := pos.GetDistance(c.GetPosition())
		if d < bestDist {
			bestDist = d
			best = c
		}
	}
	return best
}

func nearestMineral(from *bwapi.Unit, minerals []*bwapi.Unit) *bwapi.Unit {
	var best *bwapi.Unit
	bestDist := math.MaxFloat64
	pos := from.GetPosition()
	for _, m := range minerals {
		d := pos.GetDistance(m.GetPosition())
		if d < bestDist {
			bestDist = d
			best = m
		}
	}
	return best
}

func main() {
	log.Println("Starting 4-pool bot...")
	client := bwapi.NewBWClient()
	defer client.Close()
	client.Run(&ExampleBot{})
}

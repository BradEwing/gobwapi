package bwapi

import (
	"github.com/bradewing/gobwapi/internal/shm"
)

// Player wraps a BWAPI player's shared memory data.
type Player struct {
	data  *shm.PlayerData
	game  *Game
	index int
}

func (p *Player) Index() int      { return p.index }
func (p *Player) Name() string    { return p.data.Name() }
func (p *Player) GetRace() Race   { return Race(p.data.RaceID()) }
func (p *Player) IsNeutral() bool { return p.data.IsNeutral() }

func (p *Player) GetStartLocation() TilePosition {
	return TilePosition{X: p.data.StartLocationX(), Y: p.data.StartLocationY()}
}

func (p *Player) IsVictorious() bool { return p.data.IsVictorious() }
func (p *Player) IsDefeated() bool   { return p.data.IsDefeated() }
func (p *Player) LeftGame() bool     { return p.data.LeftGame() }

// --- Resources ---

func (p *Player) Minerals() int         { return int(p.data.Minerals()) }
func (p *Player) Gas() int              { return int(p.data.Gas()) }
func (p *Player) GatheredMinerals() int { return int(p.data.GatheredMinerals()) }
func (p *Player) GatheredGas() int      { return int(p.data.GatheredGas()) }

// --- Supply ---

func (p *Player) SupplyTotal() int {
	return int(p.data.SupplyTotal(int(p.GetRace())))
}

func (p *Player) SupplyUsed() int {
	return int(p.data.SupplyUsed(int(p.GetRace())))
}

func (p *Player) SupplyTotalForRace(race Race) int {
	return int(p.data.SupplyTotal(int(race)))
}

func (p *Player) SupplyUsedForRace(race Race) int {
	return int(p.data.SupplyUsed(int(race)))
}

// --- Unit Counts ---

func (p *Player) AllUnitCount(unitType UnitType) int {
	return int(p.data.AllUnitCount(int(unitType)))
}

func (p *Player) VisibleUnitCount(unitType UnitType) int {
	return int(p.data.VisibleUnitCount(int(unitType)))
}

func (p *Player) CompletedUnitCount(unitType UnitType) int {
	return int(p.data.CompletedUnitCount(int(unitType)))
}

func (p *Player) DeadUnitCount(unitType UnitType) int {
	return int(p.data.DeadUnitCount(int(unitType)))
}

func (p *Player) KilledUnitCount(unitType UnitType) int {
	return int(p.data.KilledUnitCount(int(unitType)))
}

// --- Upgrades / Research ---

func (p *Player) UpgradeLevel(upgrade UpgradeType) int {
	return int(p.data.UpgradeLevel(int(upgrade)))
}

func (p *Player) HasResearched(tech TechType) bool {
	return p.data.HasResearched(int(tech))
}

func (p *Player) IsResearching(tech TechType) bool {
	return p.data.IsResearching(int(tech))
}

func (p *Player) IsUpgrading(upgrade UpgradeType) bool {
	return p.data.IsUpgrading(int(upgrade))
}

func (p *Player) MaxUpgradeLevel(upgrade UpgradeType) int {
	return int(p.data.MaxUpgradeLevel(int(upgrade)))
}

func (p *Player) IsResearchAvailable(tech TechType) bool {
	return p.data.IsResearchAvailable(int(tech))
}

func (p *Player) IsUnitAvailable(unitType UnitType) bool {
	return p.data.IsUnitAvailable(int(unitType))
}

// --- Scores ---

func (p *Player) TotalUnitScore() int     { return int(p.data.TotalUnitScore()) }
func (p *Player) TotalKillScore() int     { return int(p.data.TotalKillScore()) }
func (p *Player) TotalBuildingScore() int { return int(p.data.TotalBuildingScore()) }
func (p *Player) TotalRazingScore() int   { return int(p.data.TotalRazingScore()) }
func (p *Player) CustomScore() int        { return int(p.data.CustomScore()) }

// --- Alliance ---

func (p *Player) IsAlly(other *Player) bool {
	return p.data.IsAlly(other.index)
}

func (p *Player) IsEnemy(other *Player) bool {
	return p.data.IsEnemy(other.index)
}

func (p *Player) Color() int { return int(p.data.Color()) }

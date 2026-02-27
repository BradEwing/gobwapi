package bwapi

import (
	"github.com/bradewing/gobwapi/internal/shm"
)

// dispatchEvents reads events from shared memory and calls the appropriate AIModule callbacks.
func dispatchEvents(game *Game, module AIModule) {
	count := game.data.EventCount()
	for i := 0; i < count; i++ {
		evtType := EventType(game.data.EventType(i))
		v1 := game.data.EventV1(i)
		v2 := game.data.EventV2(i)

		switch evtType {
		case EventTypeMatchStart:
			module.OnStart(game)

		case EventTypeMatchEnd:
			module.OnEnd(game, v1 != 0)

		case EventTypeMatchFrame:
			module.OnFrame(game)

		case EventTypeSendText:
			text := game.data.EventString(int(v1))
			module.OnSendText(game, text)

		case EventTypeReceiveText:
			player := getPlayerByIndex(game, int(v1))
			text := game.data.EventString(int(v2))
			module.OnReceiveText(game, player, text)

		case EventTypePlayerLeft:
			player := getPlayerByIndex(game, int(v1))
			module.OnPlayerLeft(game, player)

		case EventTypeNukeDetect:
			pos := Position{X: v1, Y: v2}
			module.OnNukeDetect(game, pos)

		case EventTypeUnitDiscover:
			unit := getUnitByIndex(game, int(v1))
			module.OnUnitDiscover(game, unit)

		case EventTypeUnitEvade:
			unit := getUnitByIndex(game, int(v1))
			module.OnUnitEvade(game, unit)

		case EventTypeUnitShow:
			unit := getUnitByIndex(game, int(v1))
			module.OnUnitShow(game, unit)

		case EventTypeUnitHide:
			unit := getUnitByIndex(game, int(v1))
			module.OnUnitHide(game, unit)

		case EventTypeUnitCreate:
			unit := getUnitByIndex(game, int(v1))
			module.OnUnitCreate(game, unit)

		case EventTypeUnitDestroy:
			unit := getUnitByIndex(game, int(v1))
			module.OnUnitDestroy(game, unit)

		case EventTypeUnitMorph:
			unit := getUnitByIndex(game, int(v1))
			module.OnUnitMorph(game, unit)

		case EventTypeUnitRenegade:
			unit := getUnitByIndex(game, int(v1))
			module.OnUnitRenegade(game, unit)

		case EventTypeSaveGame:
			text := game.data.EventString(int(v1))
			module.OnSaveGame(game, text)

		case EventTypeUnitComplete:
			unit := getUnitByIndex(game, int(v1))
			module.OnUnitComplete(game, unit)
		}
	}
}

func getUnitByIndex(game *Game, index int) *Unit {
	if index < 0 || index >= shm.MaxUnits {
		return nil
	}
	return &Unit{data: game.data.Unit(index), game: game, index: index}
}

func getPlayerByIndex(game *Game, index int) *Player {
	if index < 0 || index >= shm.MaxPlayers {
		return nil
	}
	return &Player{data: game.data.Player(index), game: game, index: index}
}

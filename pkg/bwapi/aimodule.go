package bwapi

// AIModule is the interface for bot callbacks.
// Implement the methods you need and embed BaseModule for the rest.
type AIModule interface {
	OnStart(game *Game)
	OnEnd(game *Game, isWinner bool)
	OnFrame(game *Game)
	OnSendText(game *Game, text string)
	OnReceiveText(game *Game, player *Player, text string)
	OnPlayerLeft(game *Game, player *Player)
	OnNukeDetect(game *Game, target Position)
	OnUnitDiscover(game *Game, unit *Unit)
	OnUnitEvade(game *Game, unit *Unit)
	OnUnitShow(game *Game, unit *Unit)
	OnUnitHide(game *Game, unit *Unit)
	OnUnitCreate(game *Game, unit *Unit)
	OnUnitDestroy(game *Game, unit *Unit)
	OnUnitMorph(game *Game, unit *Unit)
	OnUnitRenegade(game *Game, unit *Unit)
	OnSaveGame(game *Game, gameName string)
	OnUnitComplete(game *Game, unit *Unit)
}

// BaseModule provides default no-op implementations for all AIModule callbacks.
// Embed this in your bot struct and override only the methods you need.
type BaseModule struct{}

func (BaseModule) OnStart(*Game)                        {}
func (BaseModule) OnEnd(*Game, bool)                    {}
func (BaseModule) OnFrame(*Game)                        {}
func (BaseModule) OnSendText(*Game, string)             {}
func (BaseModule) OnReceiveText(*Game, *Player, string) {}
func (BaseModule) OnPlayerLeft(*Game, *Player)          {}
func (BaseModule) OnNukeDetect(*Game, Position)         {}
func (BaseModule) OnUnitDiscover(*Game, *Unit)          {}
func (BaseModule) OnUnitEvade(*Game, *Unit)             {}
func (BaseModule) OnUnitShow(*Game, *Unit)              {}
func (BaseModule) OnUnitHide(*Game, *Unit)              {}
func (BaseModule) OnUnitCreate(*Game, *Unit)            {}
func (BaseModule) OnUnitDestroy(*Game, *Unit)           {}
func (BaseModule) OnUnitMorph(*Game, *Unit)             {}
func (BaseModule) OnUnitRenegade(*Game, *Unit)          {}
func (BaseModule) OnSaveGame(*Game, string)             {}
func (BaseModule) OnUnitComplete(*Game, *Unit)          {}

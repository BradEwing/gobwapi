package bwapi

// EventType represents a BWAPI game event type.
type EventType int32

const (
	EventTypeMatchStart  EventType = 0
	EventTypeMatchEnd    EventType = 1
	EventTypeMatchFrame  EventType = 2
	EventTypeMenuFrame   EventType = 3
	EventTypeSendText    EventType = 4
	EventTypeReceiveText EventType = 5
	EventTypePlayerLeft  EventType = 6
	EventTypeNukeDetect  EventType = 7
	EventTypeUnitDiscover EventType = 8
	EventTypeUnitEvade   EventType = 9
	EventTypeUnitShow    EventType = 10
	EventTypeUnitHide    EventType = 11
	EventTypeUnitCreate  EventType = 12
	EventTypeUnitDestroy EventType = 13
	EventTypeUnitMorph   EventType = 14
	EventTypeUnitRenegade EventType = 15
	EventTypeSaveGame    EventType = 16
	EventTypeUnitComplete EventType = 17
	EventTypeNone        EventType = 18
)

func (e EventType) String() string {
	switch e {
	case EventTypeMatchStart:
		return "MatchStart"
	case EventTypeMatchEnd:
		return "MatchEnd"
	case EventTypeMatchFrame:
		return "MatchFrame"
	case EventTypeMenuFrame:
		return "MenuFrame"
	case EventTypeSendText:
		return "SendText"
	case EventTypeReceiveText:
		return "ReceiveText"
	case EventTypePlayerLeft:
		return "PlayerLeft"
	case EventTypeNukeDetect:
		return "NukeDetect"
	case EventTypeUnitDiscover:
		return "UnitDiscover"
	case EventTypeUnitEvade:
		return "UnitEvade"
	case EventTypeUnitShow:
		return "UnitShow"
	case EventTypeUnitHide:
		return "UnitHide"
	case EventTypeUnitCreate:
		return "UnitCreate"
	case EventTypeUnitDestroy:
		return "UnitDestroy"
	case EventTypeUnitMorph:
		return "UnitMorph"
	case EventTypeUnitRenegade:
		return "UnitRenegade"
	case EventTypeSaveGame:
		return "SaveGame"
	case EventTypeUnitComplete:
		return "UnitComplete"
	case EventTypeNone:
		return "None"
	default:
		return "Unknown"
	}
}

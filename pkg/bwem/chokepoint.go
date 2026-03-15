package bwem

import "github.com/bradewing/gobwapi/pkg/bwapi"

// ChokePoint represents a narrow passage between two Areas.
type ChokePoint struct {
	Index   int
	AreaIDs [2]AreaId

	End1   bwapi.WalkPosition
	Middle bwapi.WalkPosition
	End2   bwapi.WalkPosition

	Geometry []bwapi.WalkPosition

	Blocked    bool
	Pseudo     bool
	NeutralIdx int
}

// MiddlePos returns the middle node as a pixel Position.
func (cp *ChokePoint) MiddlePos() bwapi.Position {
	return cp.Middle.ToPosition()
}

// End1Pos returns the first endpoint as a pixel Position.
func (cp *ChokePoint) End1Pos() bwapi.Position {
	return cp.End1.ToPosition()
}

// End2Pos returns the second endpoint as a pixel Position.
func (cp *ChokePoint) End2Pos() bwapi.Position {
	return cp.End2.ToPosition()
}

// OtherArea returns the area on the other side of this chokepoint.
func (cp *ChokePoint) OtherArea(id AreaId) AreaId {
	if cp.AreaIDs[0] == id {
		return cp.AreaIDs[1]
	}
	return cp.AreaIDs[0]
}

// IsPseudo returns true if this is a pseudo-chokepoint created for a blocking neutral.
func (cp *ChokePoint) IsPseudo() bool {
	return cp.Pseudo
}

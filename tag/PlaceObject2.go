package tag

import (
	"github.com/judah-caruso/swf.go/subtypes"
	"github.com/judah-caruso/swf.go/types"
)

type PlaceObject2 struct {
	_    struct{} `swfFlags:"root,align"`
	Flag struct {
		HasClipActions    bool
		HasClipDepth      bool
		HasName           bool
		HasRatio          bool
		HasColorTransform bool
		HasMatrix         bool
		HasCharacter      bool
		Move              bool
	}
	Depth          uint16
	CharacterId    uint16                `swfCondition:"Flag.HasCharacter"`
	Matrix         types.MATRIX          `swfCondition:"Flag.HasMatrix"`
	ColorTransform types.CXFORMWITHALPHA `swfCondition:"Flag.HasColorTransform"`
	Ratio          uint16                `swfCondition:"Flag.HasRatio"`
	Name           string                `swfCondition:"Flag.HasName"`
	ClipDepth      uint16                `swfCondition:"Flag.HasClipDepth"`
	ClipActions    subtypes.CLIPACTIONS  `swfCondition:"Flag.HasClipActions"`
}

func (t *PlaceObject2) Code() Code {
	return RecordPlaceObject2
}

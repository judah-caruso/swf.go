package tag

import (
	"github.com/judah-caruso/swf.go/subtypes"
	"github.com/judah-caruso/swf.go/types"
)

type DefineMorphShape struct {
	_                      struct{} `swfFlags:"root"`
	CharacterId            uint16
	StartBounds, EndBounds types.RECT
	Offset                 uint32
	MorphFillStyles        subtypes.MORPHFILLSTYLEARRAY
	MorphLineStyles        subtypes.MORPHLINESTYLEARRAY
	StartEdges, EndEdges   subtypes.SHAPE
}

func (t *DefineMorphShape) Code() Code {
	return RecordDefineMorphShape
}

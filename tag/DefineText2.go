package tag

import (
	"github.com/judah-caruso/swf.go/subtypes"
	"github.com/judah-caruso/swf.go/types"
)

type DefineText2 struct {
	_           struct{} `swfFlags:"root"`
	CharacterId uint16
	Bounds      types.RECT
	Matrix      types.MATRIX
	GlyphBits   uint8
	AdvanceBits uint8
	TextRecords subtypes.TEXTRECORDS `swfFlags:"Text2"`
}

func (t *DefineText2) Code() Code {
	return RecordDefineText2
}

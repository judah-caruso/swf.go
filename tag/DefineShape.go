package tag

import (
	"github.com/judah-caruso/swf.go/subtypes"
	"github.com/judah-caruso/swf.go/types"
)

type DefineShape struct {
	_           struct{} `swfFlags:"root"`
	ShapeId     uint16
	ShapeBounds types.RECT
	Shapes      subtypes.SHAPEWITHSTYLE
}

func (t *DefineShape) Code() Code {
	return RecordDefineShape
}

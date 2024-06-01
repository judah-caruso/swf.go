package tag

import (
	"github.com/judah-caruso/swf.go/subtypes"
	"github.com/judah-caruso/swf.go/types"
)

type DefineShape2 struct {
	_           struct{} `swfFlags:"root,align"`
	ShapeId     uint16
	ShapeBounds types.RECT
	Shapes      subtypes.SHAPEWITHSTYLE
}

func (t *DefineShape2) Code() Code {
	return RecordDefineShape2
}

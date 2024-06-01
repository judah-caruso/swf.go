package tag

import (
	"github.com/judah-caruso/swf.go/subtypes"
	"github.com/judah-caruso/swf.go/types"
)

type DefineShape3 struct {
	_           struct{} `swfFlags:"root,align"`
	ShapeId     uint16
	ShapeBounds types.RECT
	Shapes      subtypes.SHAPEWITHSTYLE `swfFlags:"Shape3"`
}

func (t *DefineShape3) Code() Code {
	return RecordDefineShape3
}

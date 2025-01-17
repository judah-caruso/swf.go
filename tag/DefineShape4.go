package tag

import (
	"github.com/judah-caruso/swf.go/subtypes"
	"github.com/judah-caruso/swf.go/types"
)

type DefineShape4 struct {
	_                     struct{} `swfFlags:"root,align"`
	ShapeId               uint16
	ShapeBounds           types.RECT
	EdgeBounds            types.RECT
	Reserved              uint8 `swfBits:",5"`
	UsesFillWindingRule   bool
	UsesNonScalingStrokes bool
	UsesScalingStrokes    bool
	Shapes                subtypes.SHAPEWITHSTYLE `swfFlags:"Shape4"`
}

func (t *DefineShape4) Code() Code {
	return RecordDefineShape4
}

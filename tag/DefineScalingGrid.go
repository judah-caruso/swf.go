package tag

import "github.com/judah-caruso/swf.go/types"

type DefineScalingGrid struct {
	CharacterId uint16
	Splitter    types.RECT
}

func (t *DefineScalingGrid) Code() Code {
	return RecordDefineScalingGrid
}

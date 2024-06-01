package tag

import (
	"github.com/judah-caruso/swf.go/types"
)

type DefineBits struct {
	_           struct{} `swfFlags:"root"`
	CharacterId uint16
	Data        types.UntilEndBytes
}

func (t *DefineBits) Code() Code {
	return RecordDefineBits
}

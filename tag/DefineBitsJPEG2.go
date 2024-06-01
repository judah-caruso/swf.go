package tag

import "github.com/judah-caruso/swf.go/types"

type DefineBitsJPEG2 struct {
	_           struct{} `swfFlags:"root"`
	CharacterId uint16
	Data        types.UntilEndBytes
}

func (t *DefineBitsJPEG2) Code() Code {
	return RecordDefineBitsJPEG2
}

package tag

import "github.com/judah-caruso/swf.go/types"

type DefineBinaryData struct {
	CharacterId uint16
	Reserved    uint32
	Data        types.UntilEndBytes
}

func (t *DefineBinaryData) Code() Code {
	return RecordDefineBinaryData
}

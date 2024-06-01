package tag

import (
	"github.com/judah-caruso/swf.go/types"
)

type DoABC struct {
	Flags uint32
	Name  string
	Data  types.UntilEndBytes
}

func (t *DoABC) Code() Code {
	return RecordDoABC
}

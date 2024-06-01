package tag

import (
	"github.com/judah-caruso/swf.go/types"
)

type JPEGTables struct {
	_    struct{} `swfFlags:"root"`
	Data types.UntilEndBytes
}

func (t *JPEGTables) Code() Code {
	return RecordJPEGTables
}

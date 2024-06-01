package tag

import "github.com/judah-caruso/swf.go/types"

type SetBackgroundColor struct {
	BackgroundColor types.RGB
}

func (s *SetBackgroundColor) Code() Code {
	return RecordSetBackgroundColor
}

package tag

import (
	"github.com/judah-caruso/swf.go/subtypes"
	"github.com/judah-caruso/swf.go/types"
)

type DoAction struct {
	Actions []subtypes.ACTIONRECORD
}

func (t *DoAction) SWFRead(r types.DataReader, ctx types.ReaderContext) (err error) {
	for {
		var record subtypes.ACTIONRECORD
		err = types.ReadType(r, ctx, &record)
		if err != nil {
			return err
		}
		if record.ActionCode == 0 {
			break
		}
		t.Actions = append(t.Actions, record)
	}

	return nil
}

func (t *DoAction) Code() Code {
	return RecordDoAction
}

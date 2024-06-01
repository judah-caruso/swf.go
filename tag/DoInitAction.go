package tag

import (
	"github.com/judah-caruso/swf.go/subtypes"
	"github.com/judah-caruso/swf.go/types"
)

type DoInitAction struct {
	SpriteID uint16
	Actions  []subtypes.ACTIONRECORD
}

func (t *DoInitAction) SWFRead(r types.DataReader, ctx types.ReaderContext) (err error) {
	err = types.ReadU16(r, &t.SpriteID)
	if err != nil {
		return err
	}
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

func (t *DoInitAction) Code() Code {
	return RecordDoInitAction
}

package tag

import (
	"github.com/judah-caruso/swf.go/subtypes"
	"github.com/judah-caruso/swf.go/types"
)

type DefineFont struct {
	_                struct{} `swfFlags:"root"`
	FontId           uint16
	NumGlyphsEntries uint16
	OffsetTable      []uint16         `swfCount:"OffsetTableLength()"`
	ShapeTable       []subtypes.SHAPE `swfCount:"ShapeTableLength()"`
}

func (t *DefineFont) Scale() float64 {
	return 1024
}

func (t *DefineFont) OffsetTableLength(ctx types.ReaderContext) uint64 {
	return uint64(t.NumGlyphsEntries/2) - 1
}

func (t *DefineFont) ShapeTableLength(ctx types.ReaderContext) uint64 {
	return uint64(t.NumGlyphsEntries / 2)
}

func (t *DefineFont) Code() Code {
	return RecordDefineFont
}

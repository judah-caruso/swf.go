package tag

import (
	"image"

	"github.com/judah-caruso/swf.go/subtypes"
	"github.com/judah-caruso/swf.go/types"
)

type DefineBitsLossless struct {
	_              struct{} `swfFlags:"root"`
	CharacterId    uint16
	Format         subtypes.ImageBitsFormat
	Width, Height  uint16
	ColorTableSize uint8 `swfCondition:"HasColorTableSize()"`
	ZlibData       types.UntilEndBytes
}

func (t *DefineBitsLossless) HasColorTableSize(ctx types.ReaderContext) bool {
	return t.Format == 3
}

func (t *DefineBitsLossless) GetImage() (image.Image, error) {
	return subtypes.DecodeImageBits(t.ZlibData, int(t.Width), int(t.Height), t.Format, int(t.ColorTableSize)+1, false)
}

func (t *DefineBitsLossless) Code() Code {
	return RecordDefineBitsLossless
}

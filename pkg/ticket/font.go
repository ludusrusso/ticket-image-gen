package ticket

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

func setFont(dc *gg.Context, f *truetype.Font, size float64) {
	face := truetype.NewFace(f, &truetype.Options{
		Size: size,
	})
	dc.SetFontFace(face)
}

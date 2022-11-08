package job

import (
	"fmt"
	"image"
	"net/http"

	"github.com/fogleman/gg"
	"github.com/ludusrusso/ticket-image-gen/pkg/palette"
)

type Job struct {
	Title         string
	Company       string
	Ral           string
	Location      string
	Type          string
	Palette       palette.ColorPalette
	BgTransparent bool
}

func (j Job) Draw() (*gg.Context, error) {
	dc := gg.NewContext(1200, 630)
	prepareBG(dc, j.Palette, j.BgTransparent)
	dc.Translate(150, 65)

	if err := j.drawJob(dc); err != nil {
		return nil, fmt.Errorf("failed to draw event: %w", err)
	}

	return dc, nil
}

func loadRemoteImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	image, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	return image, nil
}

package job

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/ludusrusso/ticket-image-gen/pkg/fonts"
)

func (j Job) drawJob(dc *gg.Context) error {
	dc.Push()
	defer dc.Pop()

	dc.Translate(0, 50)

	if err := j.drawEventName(dc); err != nil {
		return err
	}
	if err := j.drawEventLogo(dc); err != nil {
		return err
	}

	return nil
}

func (j Job) drawEventName(dc *gg.Context) error {
	setFont(dc, fonts.RobotoBlack, 42)
	dc.SetHexColor(j.Palette.V600)
	dc.DrawStringWrapped(j.Title, 64, 0, 0, 0, float64(dc.Width()-300), 1, gg.AlignLeft)

	_, h := dc.MeasureString(j.Title)
	// dc.Translate(0, h)

	setFont(dc, fonts.RobotMedium, 30)

	dc.Translate(64, 145)
	dc.DrawString("Azienda: "+j.Company, 0, 0)

	dc.Translate(0, 25)
	setFont(dc, fonts.RobotoItalicRegular, 30)
	_, h = dc.MeasureString(j.Location)
	dc.DrawString(j.Ral, 0, h)
	dc.DrawString(j.Location, 0, 2*h+10)
	dc.DrawString(j.Type, 0, 3*h+20)

	setFont(dc, fonts.RobotoRegularl, 30)

	return nil
}

//go:embed logo.png
var logoB []byte

func (j Job) drawEventLogo(dc *gg.Context) error {
	dc.Translate(0, 150)
	logo, _, err := image.Decode(bytes.NewReader(logoB))
	if err != nil {
		return fmt.Errorf("failed to load background image: %w", err)
	}

	size := 90
	logo = imaging.Resize(logo, size, size, imaging.Lanczos)
	m := gg.NewContext(1000, size)

	m.DrawImage(logo, 0, 0)
	msak := m.AsMask()
	ac := gg.NewContext(1000, size)
	ac.SetMask(msak)

	ac.SetHexColor(j.Palette.V600)
	ac.DrawRectangle(0, 0, float64(size), float64(size))
	ac.Fill()
	ac.InvertMask()
	ac.Translate(float64(size)+15, 0)
	setFont(ac, fonts.RobotMedium, 35)

	w, h := ac.MeasureString("Farmaceutica")

	ac.SetHexColor(j.Palette.V600)
	ac.DrawString("Farmaceutica", 0, 90/2+h/2)

	setFont(ac, fonts.RobotoItalicRegular, 35)
	ac.DrawString("Younger", w+10, 90/2+h/2)

	dc.DrawImage(ac.Image(), 0, 0)
	return nil
}

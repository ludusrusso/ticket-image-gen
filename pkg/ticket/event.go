package ticket

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/ludusrusso/ticket-image-gen/pkg/fonts"
)

func (t *Ticket) drawEvent(dc *gg.Context) error {
	dc.Push()
	defer dc.Pop()

	dc.Translate(64, 145)

	if err := t.drawEventName(dc); err != nil {
		return err
	}
	if err := t.drawEventLogo(dc); err != nil {
		return err
	}

	return nil
}

func (t Ticket) drawEventName(dc *gg.Context) error {
	setFont(dc, fonts.RobotoBlack, 42)
	dc.SetHexColor(t.Palette.V600)

	_, h := dc.MeasureString(t.EventName)
	dc.Translate(0, h)

	dc.DrawString(t.EventName, 0, 0)

	dc.Translate(0, 25)
	setFont(dc, fonts.RobotMediumItalic, 35)
	_, h = dc.MeasureString(t.EventLocation)
	dc.DrawString(t.EventLocation, 0, h)

	setFont(dc, fonts.RobotoRegularl, 30)

	dc.Translate(0, 60)
	_, h = dc.MeasureString(t.EventDate)
	dc.DrawString(t.EventDate, 0, h)
	dc.Translate(0, 35)
	_, h = dc.MeasureString(t.EventHours)
	dc.DrawString(t.EventHours, 0, h)

	return nil
}

//go:embed logo.png
var logoB []byte

func (t Ticket) drawEventLogo(dc *gg.Context) error {
	dc.Translate(0, 50)
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

	ac.SetHexColor(t.Palette.V600)
	ac.DrawRectangle(0, 0, float64(size), float64(size))
	ac.Fill()
	ac.InvertMask()
	ac.Translate(float64(size)+15, 0)
	setFont(ac, fonts.RobotMedium, 35)

	w, h := ac.MeasureString("Farmaceutica")

	ac.SetHexColor(t.Palette.V600)
	ac.DrawString("Farmaceutica", 0, 90/2+h/2)

	setFont(ac, fonts.RobotoItalicRegular, 35)
	ac.DrawString("Younger", w+10, 90/2+h/2)

	dc.DrawImage(ac.Image(), 0, 0)
	return nil
}

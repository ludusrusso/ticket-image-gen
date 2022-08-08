package ticket

import (
	"fmt"
	"image/color"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/ludusrusso/ticket-image-gen/pkg/fonts"
)

func (t Ticket) drawUser(dc *gg.Context) error {
	dc.Push()
	defer dc.Pop()
	dc.Translate(64, 28)
	if err := drawUserAvatar(dc, t.UserAvatar); err != nil {
		return fmt.Errorf("failed to draw user avatar: %w", err)
	}

	if err := drawUserInfo(dc, t.UserName); err != nil {
		return fmt.Errorf("failed to draw user info: %w", err)
	}
	return nil
}

func drawUserInfo(dc *gg.Context, name string) error {
	dc.Push()
	defer dc.Pop()
	dc.Translate(90, 45)
	setFont(dc, fonts.RobotMedium, 36)
	dc.SetColor(color.Black)
	_, textHeight := dc.MeasureString(name)
	x := 16.0
	y := textHeight / 4
	dc.DrawString(name, x, y)

	return nil
}

func drawUserAvatar(dc *gg.Context, url string) error {
	image, err := loadRemoteImage(url)
	if err != nil {
		return err
	}

	size := 90

	image = imaging.Resize(image, size, size, imaging.Lanczos)

	ac := gg.NewContext(size, size)
	ac.DrawRoundedRectangle(0, 0, float64(size), float64(size), float64(size)/2)
	ac.Clip()
	ac.DrawImage(image, 0, 0)

	dc.DrawImage(ac.Image(), 0, 0)
	return nil
}

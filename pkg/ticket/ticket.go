package ticket

import (
	"fmt"
	"image"
	"math"
	"net/http"

	"github.com/fogleman/gg"
	"github.com/ludusrusso/ticket-image-gen/pkg/fonts"
	"github.com/ludusrusso/ticket-image-gen/pkg/palette"
)

type Ticket struct {
	UserName      string
	UserAvatar    string
	EventName     string
	TicketNum     uint
	EventLocation string
	EventDate     string
	EventHours    string
	Palette       palette.ColorPalette
	BgTransparent bool
}

func (t Ticket) Draw() (*gg.Context, error) {
	dc := gg.NewContext(1200, 630)
	prepareBG(dc, t.Palette, t.BgTransparent)
	dc.Translate(150, 65)

	if err := t.drawEvent(dc); err != nil {
		return nil, fmt.Errorf("failed to draw event: %w", err)
	}

	if err := t.drawUser(dc); err != nil {
		return nil, fmt.Errorf("failed to draw user: %w", err)
	}

	if err := t.drawTicketNumber(dc); err != nil {
		return nil, fmt.Errorf("failed to draw ticket: %w", err)
	}

	return dc, nil
}

func (t Ticket) drawTicketNumber(dc *gg.Context) error {
	dc.Push()
	defer dc.Pop()
	dc.Translate(-150, -65)

	setFont(dc, fonts.RobotMedium, 50)
	dc.Rotate(math.Pi / 2)
	dc.SetHexColor(t.Palette.V300)
	number := fmt.Sprintf("N° %010d", t.TicketNum)
	textWidth, _ := dc.MeasureString(number)
	x := (float64(dc.Height()) - textWidth) / 2
	y := -920.0
	dc.DrawString(number, x, y)
	return nil
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

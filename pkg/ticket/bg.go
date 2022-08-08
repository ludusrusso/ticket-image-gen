package ticket

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/ludusrusso/ticket-image-gen/pkg/palette"
)

func prepareBG(dc *gg.Context, p palette.ColorPalette, drawBg bool) error {
	if drawBg {
		dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
		dc.SetHexColor(p.V300)
		dc.Fill()
	}

	px := 150.0
	py := 65.0
	drawTicketShape(dc, px, py, float64(dc.Width())-2*px, float64(dc.Height())-2*py, 30, 20, 0.75)
	dc.SetColor(color.White)
	dc.Fill()
	drawTicketShape(dc, px, py, float64(dc.Width())-2*px, float64(dc.Height())-2*py, 30, 20, 0.75)

	dc.SetHexColor(p.V600)
	dc.SetLineWidth(6)
	dc.Stroke()

	drawLine(dc, px, py, float64(dc.Width())-2*px, float64(dc.Height())-2*py, 20, 0.75)

	dc.SetHexColor(p.V600)
	dc.SetLineWidth(4)
	dc.SetDash(15, 15)

	dc.Stroke()
	return nil
}

func drawLine(dc *gg.Context, x, y, w, h, sr, pw float64) {
	y0, y1 := y+sr, y+h-sr

	xp := w*pw + x

	dc.NewSubPath()
	dc.MoveTo(xp, y0)
	dc.LineTo(xp, y1)

}

func drawTicketShape(dc *gg.Context, x, y, w, h, r, sr, pw float64) {
	x0, x1, x2, x3 := x, x+r, x+w-r, x+w
	y0, y1, y2, y3 := y, y+r, y+h-r, y+h

	xp := w*pw + x
	xp1, xp2 := xp-sr, xp+sr

	dc.NewSubPath()
	dc.MoveTo(x1, y0)
	dc.LineTo(xp1, y0)
	dc.DrawArc(xp, y0, sr, gg.Radians(180), gg.Radians(0))
	dc.LineTo(x2, y0)
	dc.DrawArc(x2, y1, r, gg.Radians(270), gg.Radians(360))
	dc.LineTo(x3, y2)
	dc.DrawArc(x2, y2, r, gg.Radians(0), gg.Radians(90))
	dc.LineTo(xp2, y3)
	dc.DrawArc(xp, y3, sr, gg.Radians(0), gg.Radians(-180))
	dc.LineTo(x1, y3)
	dc.DrawArc(x1, y2, r, gg.Radians(90), gg.Radians(180))
	dc.LineTo(x0, y2)
	dc.DrawArc(x1, y1, r, gg.Radians(180), gg.Radians(270))
	dc.ClosePath()
}

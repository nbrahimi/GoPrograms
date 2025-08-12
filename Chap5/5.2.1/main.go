package main

import (
	"image"
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

// Test colors.
var (
	red   = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
	green = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: 0xFF}
	blue  = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("5.2.1-Stacked"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		var ops op.Ops
		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return stacked(gtx)
				})
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func stacked(gtx layout.Context) layout.Dimensions {
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, gtx.Constraints.Min, red)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(200, 40), green)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(60, 100), blue)
		}),
	)
}

// ColorBox creates a widget with the specified dimensions and color.
func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

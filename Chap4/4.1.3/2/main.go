package main

import (
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("Custom Editor"))
		w.Option(app.Size(unit.Dp(500), unit.Dp(300)))
		th := material.NewTheme()
		var ops op.Ops
		var editor widget.Editor
		editor.SingleLine = true
		InputMargins := layout.Inset{
			Top:    unit.Dp(30),
			Bottom: unit.Dp(30),
			Right:  unit.Dp(20),
			Left:   unit.Dp(20),
		}
		EditBorder := widget.Border{
			Color:        color.NRGBA{R: 255, G: 0, B: 0, A: 255}, // Red border
			CornerRadius: unit.Dp(2),
			Width:        unit.Dp(2),
		}

		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)

				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx C) D {
						return InputMargins.Layout(gtx, func(gtx C) D {
							ed := material.Editor(th, &editor, "Blue text here ...")
							ed.TextSize = unit.Sp(25)
							ed.Color = color.NRGBA{R: 0, G: 120, B: 255, A: 255} // Blue text color
							return EditBorder.Layout(gtx, ed.Layout)
						},
						)
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

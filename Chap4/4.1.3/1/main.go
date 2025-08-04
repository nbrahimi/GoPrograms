package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("4.1.3.1-Editor with border"))
		w.Option(app.Size(unit.Dp(200), unit.Dp(300)))
		if err := run(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

type (
	C = layout.Context
	D = layout.Dimensions
)

func run(w *app.Window) error {
	th := material.NewTheme()
	var editor widget.Editor
	var ops op.Ops
	editor.SingleLine = true

	border := widget.Border{
		Color:        color.NRGBA{R: 0, G: 0, B: 0, A: 255}, // Black border
		Width:        unit.Dp(2),
		CornerRadius: unit.Dp(4),
	}

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			layout.Center.Layout(gtx,
				func(gtx C) D {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							// Adding a border around the editor
							return border.Layout(gtx, func(gtx C) D {
								// Draw the editor inside the border
								ed := material.Editor(th, &editor, "Enter text...")
								ed.TextSize = unit.Sp(25)
								return ed.Layout(gtx)
							})
						}),
					)
				},
			)
			e.Frame(gtx.Ops)
		}
	}
}

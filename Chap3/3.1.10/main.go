package main

import (
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("3.1.10-Label Alignment"))
		w.Option(app.Size(unit.Dp(600), unit.Dp(300)))
		th := material.NewTheme()
		var ops op.Ops

		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				// Root Flex to split into two equal sides
				layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					// Left Container (50%)
					layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
						// Stack two labels vertically
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							// Rigid widget to hold the first label
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								lbl1 := material.Label(th, unit.Sp(20), "This Label is arranged right")
								lbl1.Alignment = text.End
								return lbl1.Layout(gtx)
							}),
							// Rigid widget to hold the second label
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								lbl2 := material.Label(th, unit.Sp(20), "This one too")
								lbl2.Alignment = text.End
								return lbl2.Layout(gtx)
							}),
						)
					}),
					// Right Container (50%) - Empty for now
					layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
						return layout.Dimensions{}
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

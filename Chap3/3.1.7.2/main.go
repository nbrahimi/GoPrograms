package main

import (
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
		w.Option(app.Title("3.1.7.2-80%_vs_20% with Inset"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		th := material.NewTheme()
		var ops op.Ops
		var Flexed20Button, flexed80Button widget.Clickable
		margins := layout.Inset{Top: unit.Dp(10), Bottom: unit.Dp(10), Left: unit.Dp(10), Right: unit.Dp(10)}
		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				layout.Flex{
					Axis: layout.Horizontal,
				}.Layout(gtx,
					layout.Flexed(0.8, func(gtx layout.Context) layout.Dimensions {
						return margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return material.Button(th, &flexed80Button, "0.8 Flexed").Layout(gtx)
						})
					}),
					layout.Flexed(0.2, func(gtx layout.Context) layout.Dimensions {
						return margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return material.Button(th, &Flexed20Button, "0.2 Flexed").Layout(gtx)
						})
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

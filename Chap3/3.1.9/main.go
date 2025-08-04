package main

import (
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("3.1.9-Spacer"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		th := material.NewTheme()
		var ops op.Ops
		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Label(th, unit.Sp(20), "Top Label").Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout), // 20dp space
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Label(th, unit.Sp(20), "Middle Label").Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout), // 20dp space
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Label(th, unit.Sp(20), "Bottom Label").Layout(gtx)
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

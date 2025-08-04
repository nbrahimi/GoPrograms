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
		w.Option(app.Title("3.1.11-layout.Widget"))
		w.Option(app.Size(unit.Dp(300), unit.Dp(400)))
		th := material.NewTheme()
		var ops op.Ops

		var myLabel layout.Widget = func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, unit.Sp(20), "Reusable Label").Layout(gtx)
		}

		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				layout.Center.Layout(gtx, myLabel)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

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

// Tracks the switch state
var switchState = new(widget.Bool)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("4.4.1-Basic Switch"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		th := material.NewTheme()
		var ops op.Ops
		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return material.Switch(th, switchState, "Toggle").Layout(gtx)
				})
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

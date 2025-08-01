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
		w.Option(app.Title("3.1.6-Centered button"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		th := material.NewTheme()
		var btn widget.Clickable
		var ops op.Ops
		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, &btn, "I'm in the Center").Layout(gtx)
				})

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

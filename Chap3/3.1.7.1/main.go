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
		w.Option(app.Title("3.1.7.1-layout.Inset"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		th := material.NewTheme()
		var btn widget.Clickable
		var ops op.Ops
		margins := layout.Inset{Top: unit.Dp(10), Left: unit.Dp(20)}
		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, &btn, "Click me!").Layout(gtx)
				})
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

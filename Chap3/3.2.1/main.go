package main

import (
	"fmt"
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
		w.Option(app.Title("Scrollable list"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		th := material.NewTheme()
		var ops op.Ops
		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				list := layout.List{Axis: layout.Vertical}
				list.Layout(gtx, 10, func(gtx layout.Context, index int) layout.Dimensions {
					return material.Label(th, unit.Sp(20), fmt.Sprintf("Item %d", index+1)).Layout(gtx)
				})

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

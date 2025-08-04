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

var switchState = new(widget.Bool)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("4.4.2-Dynamic Switch"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		th := material.NewTheme() // Default theme
		var ops op.Ops
		margins := layout.UniformInset(unit.Dp(20))

		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
								layout.Rigid(material.Body1(th, "Enable Feature: ").Layout),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return material.Switch(th, switchState, "").Layout(gtx)
								}),
							)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							if switchState.Value {
								return material.Body1(th, "Switch is ON!").Layout(gtx)
							}
							return layout.Dimensions{}
						}),
					)
				})
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

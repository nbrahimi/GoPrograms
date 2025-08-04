package main

import (
	"fmt"
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
		w.Option(app.Title("3.3-Toolbar"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		th := material.NewTheme()
		var ops op.Ops
		var buttons [4]widget.Clickable
		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				margins := layout.UniformInset(unit.Dp(2))
				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Flexed(0.1, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									if buttons[0].Clicked(gtx) {
										fmt.Println("Home clicked!")
									}
									return material.Button(th, &buttons[0], "Home").Layout(gtx)
								})
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									if buttons[1].Clicked(gtx) {
										fmt.Println("Settings clicked!")
									}
									return material.Button(th, &buttons[1], "Settings").Layout(gtx)
								})
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									if buttons[2].Clicked(gtx) {
										fmt.Println("Profile clicked!")
									}
									return material.Button(th, &buttons[2], "Profile").Layout(gtx)
								})
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									if buttons[3].Clicked(gtx) {
										fmt.Println("Tools clicked!")
									}
									return material.Button(th, &buttons[3], "Tools").Layout(gtx)
								})
							}),
						)
					}),
					layout.Flexed(0.9, func(gtx layout.Context) layout.Dimensions {
						return layout.Dimensions{}
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

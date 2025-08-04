package main

import (
	"log"
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
		w.Option(app.Title("4.3.1-Spaced Radiobuttons"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(400)))
		if err := run(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func run(w *app.Window) error {
	th := material.NewTheme()
	var selectedOption widget.Enum
	var ops op.Ops

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceBetween}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Bottom: unit.Dp(10)}.Layout(gtx,
							material.Body1(th, "Select a Theme:").Layout,
						)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Bottom: unit.Dp(5)}.Layout(gtx,
							material.RadioButton(th, &selectedOption, "light", "Light Mode").Layout,
						)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Bottom: unit.Dp(5)}.Layout(gtx,
							material.RadioButton(th, &selectedOption, "dark", "Dark Mode").Layout,
						)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Bottom: unit.Dp(5)}.Layout(gtx,
							material.RadioButton(th, &selectedOption, "auto", "Auto Mode").Layout,
						)
					}),
				)
			})
			e.Frame(gtx.Ops)
		}
	}
}

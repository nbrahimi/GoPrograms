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

type (
	C = layout.Context
	D = layout.Dimensions
)

var buttons [3]widget.Clickable

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("Toolbar"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		if err := run(w); err != nil {
			os.Exit(1)
		}
	}()
	app.Main()
}

func run(w *app.Window) error {
	var ops op.Ops
	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			toolbarLayouts(gtx)
			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			os.Exit(0)
		}
	}
}

func toolbarLayouts(gtx C) D {
	margins := layout.UniformInset(unit.Dp(2))
	th := material.NewTheme()
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Flexed(0.1, func(gtx C) D {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return margins.Layout(gtx, func(gtx C) D {
						return toolbarButton(th, gtx, &buttons[0], "Home")
					})
				}),
				layout.Rigid(func(gtx C) D {
					return margins.Layout(gtx, func(gtx C) D {
						return toolbarButton(th, gtx, &buttons[1], "Settings")
					})
				}),
				layout.Rigid(func(gtx C) D {
					return margins.Layout(gtx, func(gtx C) D {
						return toolbarButton(th, gtx, &buttons[2], "Profile")
					})
				}),
			)
		}),
		layout.Flexed(0.9, func(gtx C) D {
			return D{}
		}),
	)
}

func toolbarButton(theme *material.Theme, gtx C, btn *widget.Clickable, label string) D {
	if btn.Clicked(gtx) {
		fmt.Println(label + " clicked!")
	}
	return material.Button(theme, btn, label).Layout(gtx)
}

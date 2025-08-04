package main

import (
	"fmt"
	"image/color"
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
		w.Option(app.Title("3.2.2.2-List of clickable labels"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		th := material.NewTheme()
		var list widget.List = widget.List{
			List: layout.List{Axis: layout.Vertical},
		}
		var ops op.Ops
		// Create 10 clickable labels
		var labels [10]widget.Clickable
		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return list.Layout(gtx, len(labels), func(gtx layout.Context, index int) layout.Dimensions {
							lbl := &labels[index]
							if lbl.Clicked(gtx) {
								fmt.Println("Clicked label:", index)
							}
							return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return lbl.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									lbl := material.Label(th, unit.Sp(20), fmt.Sprintf("Label %d", index))
									lbl.Color = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
									
									return lbl.Layout(gtx)
								})
							})
						})
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

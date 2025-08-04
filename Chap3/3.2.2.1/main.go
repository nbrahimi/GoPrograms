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
		w.Option(app.Title("3.2.2.1-List of flat buttons"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))

		th := material.NewTheme()
		var list widget.List = widget.List{
			List: layout.List{Axis: layout.Vertical},
		}
		// Create 10 clickables
		var buttons [10]widget.Clickable
		var ops op.Ops
		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)

				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return list.Layout(gtx, len(buttons), func(gtx layout.Context, index int) layout.Dimensions {
							btn := &buttons[index]
							if btn.Clicked(gtx) {
								fmt.Println("Clicked item:", index)
							}
							return layout.UniformInset(unit.Dp(1)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								btn := material.Button(th, btn, fmt.Sprintf("Item %d", index+1))
								// Optional: make it rectangular
								btn.CornerRadius = unit.Dp(0)
								// Transparent background
								btn.Background = color.NRGBA{R: 255, G: 255, B: 255, A: 0}
								// Black text
								btn.Color = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
								return btn.Layout(gtx)
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

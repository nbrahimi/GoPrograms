package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("4.5.4-Circle-UniformRRect"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

var circleArea widget.Clickable

func loop(w *app.Window) error {
	var ops op.Ops
	circleRadius := float32(50) // Radius of our circle
	colorHovered := color.NRGBA{R: 4, G: 85, B: 4, A: 128}
	colorNormal := color.NRGBA{R: 3, G: 247, B: 3, A: 128}
	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// Handle clicks
			for circleArea.Clicked(gtx) {
				log.Println("Circle clicked!")
			}
			// Layout the clickable area
			circleArea.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				// Define the circle's dimensions
				diameter := int(2 * circleRadius)
				size := image.Pt(diameter, diameter)
				// Create a circular clip path
				defer clip.UniformRRect(
					image.Rectangle{Max: size},
					int(circleRadius),
				).Push(gtx.Ops).Pop()
				// Change color when hovering
				if circleArea.Hovered() {
					paint.Fill(gtx.Ops, colorHovered)
				} else {
					paint.Fill(gtx.Ops, colorNormal)
				}
				return layout.Dimensions{Size: size}
			})
			e.Frame(gtx.Ops)
		}
	}
}

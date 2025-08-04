package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("4.5.5-Circle-Ellipse"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func loop(w *app.Window) error {
	var ops op.Ops
	var hovered bool
	radius := float32(50)
	circlePos := image.Point{X: 100, Y: 100} // Position of circle center
	var tag = "My Input Routing Tag"

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// Define the circular area
			circleArea := clip.Ellipse{
				Min: image.Pt(
					circlePos.X-int(radius),
					circlePos.Y-int(radius)),
				Max: image.Pt(
					circlePos.X+int(radius),
					circlePos.Y+int(radius)),
			}.Op(gtx.Ops)

			// Getting and handling mouse events in this loop
			for {
				ev, ok := gtx.Event(
					pointer.Filter{
						Target: tag,
						Kinds:  pointer.Press, // filter mouse clicks only
					},
				)
				if !ok {
					break
				}

				// Handling mouse events
				if pe, ok := ev.(pointer.Event); ok && pe.Kind == pointer.Press {
					// Capture the click position
					lastClick := pe.Position.Round() // image.Point with x, y coordinates

					// Check if the click is inside the circle or outside
					dx := float32(lastClick.X) - float32(circlePos.X)
					dy := float32(lastClick.Y) - float32(circlePos.Y)
					if dx*dx+dy*dy <= radius*radius {
						// We're inside the circle. Check if it's a click
						if pe.Kind == pointer.Press {
							log.Println("Circle clicked! at coordinate ", lastClick)
						}
						hovered = true
					} else {
						if pe.Kind == pointer.Press {
							log.Println("Clicked outside the Circle! at coordinate ", lastClick)
						}
						hovered = false
					}
				}
			}
			// Draw the circle
			paint.FillShape(gtx.Ops, circleColor(hovered), circleArea)

			// Registering events using the tag
			event.Op(&ops, tag)

			e.Frame(gtx.Ops)
		}
	}
}

func circleColor(hovered bool) color.NRGBA {
	if hovered {
		return color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	}
	return color.NRGBA{R: 0, G: 255, B: 0, A: 255}
}

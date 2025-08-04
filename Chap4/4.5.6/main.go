package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("4.5.6-Rectangular confined area"))
		w.Option(app.Size(unit.Dp(300), unit.Dp(400)))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func loop(w *app.Window) error {
	var ops op.Ops
	var hovered bool

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// Create a rectangular area to be confined for mouse events
			rectangularArea := clip.Rect(image.Rect(0, 0, 100, 100)).Push(gtx.Ops)

			// Register the rectangular area
			event.Op(gtx.Ops, rectangularArea)

			// Getting and handling mouse events in this loop
			for {
				ev, ok := gtx.Event(pointer.Filter{
					// Confine the area as a target
					Target: rectangularArea,
					// Filtering clicks, entering and leaving the rectangular area
					Kinds: pointer.Press | pointer.Enter | pointer.Leave,
				})
				if !ok {
					break
				}
				pe, ok := ev.(pointer.Event)
				if !ok {
					continue
				}

				switch pe.Kind {
				case pointer.Press:
					// Getting the click position
					lastClick := pe.Position.Round()
					log.Println("Click at coordinate ", lastClick)
				case pointer.Enter:
					hovered = true
				case pointer.Leave:
					hovered = false
				}
			}
			rectangularArea.Pop()

			// Draw the rectangular area and pick the color based on the cursor positon
			areaColor := color.NRGBA{R: 0x80, A: 0xFF}
			if hovered {
				areaColor = color.NRGBA{G: 0x80, A: 0xFF}
			}
			drawRectangle(gtx.Ops, areaColor)

			e.Frame(gtx.Ops)
		}
	}
}

func drawRectangle(ops *op.Ops, areaColor color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: image.Pt(100, 100)}.Push(ops).Pop()
	paint.ColorOp{Color: areaColor}.Add(ops)
	paint.PaintOp{}.Add(ops)
	return layout.Dimensions{Size: image.Pt(100, 100)}
}

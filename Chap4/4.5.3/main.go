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
		w.Option(app.Title("4.5.3-Rectangle"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

var rectArea widget.Clickable

func loop(w *app.Window) error {
	var ops op.Ops
	colorHovered := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	colorNormal := color.NRGBA{R: 0, G: 255, B: 0, A: 255}

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// Handle clicks
			for rectArea.Clicked(gtx) {
				log.Println("Rectangle clicked!")
			}
			// Layout
			rectArea.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				// Change the color when hovering the area
				if rectArea.Hovered() {
					// Different color when hovered
					paint.FillShape(gtx.Ops, colorHovered, clip.Rect{Max: image.Pt(200, 100)}.Op())
				} else {
					// Normal color
					paint.FillShape(gtx.Ops, colorNormal, clip.Rect{Max: image.Pt(200, 100)}.Op())
				}

				return layout.Dimensions{Size: image.Pt(200, 100)}
			})
			e.Frame(gtx.Ops)
		}
	}
}

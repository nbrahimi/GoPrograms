package main

import (
	"fmt"
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
		w.Option(app.Title("4.5.1-Buttons clicks"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		if err := run(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func run(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops
	var count int
	var button widget.Clickable
	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			// Handle button clicks
			for button.Clicked(gtx) {
				count++
				fmt.Println("Button clicked! Count:", count)
			}
			// UI Layout
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(material.Button(th, &button, fmt.Sprintf("Click me! (%d)", count+1)).Layout),
				)
			})
			e.Frame(gtx.Ops)
		}
	}
}

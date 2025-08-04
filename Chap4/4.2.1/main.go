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
		w.Option(app.Title("4.2.1-Basic checkbox"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(300)))
		if err := run(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func run(w *app.Window) error {
	var ops op.Ops
	theme := material.NewTheme()
	var checkbox widget.Bool

	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.CheckBox(theme, &checkbox, "Enable feature").Layout(gtx)
			})
			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			os.Exit(0)
		}
	}
}

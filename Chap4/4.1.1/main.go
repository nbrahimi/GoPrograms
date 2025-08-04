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
		w.Option(app.Title("Editor"))
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
	var editor widget.Editor
	editor.SingleLine = true // Set single-line mode

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			// UI Layout
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(material.Editor(th, &editor, "Type here...").Layout), // Input field
					layout.Rigid(material.Body1(th, "Copy: "+editor.Text()).Layout),   // Non-editable field. Contain a copy of the previous
				)
			})
			e.Frame(gtx.Ops)
		}
	}
}

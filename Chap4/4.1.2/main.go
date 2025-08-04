package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("Editor"))
		w.Option(app.Size(unit.Dp(300), unit.Dp(200)))
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
	editor.SingleLine = false // Multi-line input
	// editor.WrapPolicy = text.WrapGraphemes
	// editor.WrapPolicy = text.WrapWords
	// editor.WrapPolicy = text.WrapHeuristically

	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(material.Body1(th, "WrapGraphemes").Layout),
					layout.Rigid(material.Editor(th, &editor, "Multi-line text...").Layout),
				)
			})
			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			os.Exit(0)
		}
	}
}

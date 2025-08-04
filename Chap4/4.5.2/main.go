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
		w.Option(app.Title("4.5.2-Editor key event"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		if err := run(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

var editor widget.Editor

func textInputEvents(gtx layout.Context, th *material.Theme) layout.Dimensions {
	// Set editor properties (typically in initialization)
	editor.SingleLine = false // or true for single line
	editor.Submit = true      // Enable submit on Enter

	// Check for changes and submit
	if ev, ok := editor.Update(gtx); ok {
		if _, ok := ev.(widget.SubmitEvent); ok {
			fmt.Println("Text submitted:", editor.Text())
		}
	}

	// Check for changes only
	if _, ok := editor.Update(gtx); ok {
		fmt.Println("Text changed:", editor.Text())
	}

	// Layout the editor
	return material.Editor(th, &editor, "Type here...").Layout(gtx)
}

func run(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			// UI Layout
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return textInputEvents(gtx, th)
			})
			e.Frame(gtx.Ops)
		}
	}
}

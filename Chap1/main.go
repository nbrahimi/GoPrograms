package main

import (
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := new(app.Window)
		th := material.NewTheme()
		var ops op.Ops
		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				material.Label(th, 16, "Hello, World!").Layout(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

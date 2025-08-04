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
		w.Option(app.Title("4.2.2.-Multi checkbox"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		if err := run(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

var (
	checkbox1 widget.Bool
	checkbox2 widget.Bool
	checkbox3 widget.Bool
)

func layoutCheckboxes(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(material.CheckBox(th, &checkbox1, "Option 1").Layout),
		layout.Rigid(material.CheckBox(th, &checkbox2, "Option 2").Layout),
		layout.Rigid(material.CheckBox(th, &checkbox3, "Option 3").Layout),
	)
}

func run(w *app.Window) error {
	var ops op.Ops
	theme := material.NewTheme()

	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layoutCheckboxes(gtx, theme)
			})
			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			os.Exit(0)
		}
	}
}

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
		w.Option(app.Title("4.3.2-Radiobuttons-Conditional Display"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme()

	var ops op.Ops
	var selectedOption widget.Enum
	var advancedOption widget.Enum

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				children := []layout.FlexChild{
					layout.Rigid(material.RadioButton(th, &selectedOption, "basic", "Basic Mode").Layout),
					layout.Rigid(material.RadioButton(th, &selectedOption, "advanced", "Advanced Mode").Layout),
				}

				// Show additional options if "Advanced Mode" is selected
				if selectedOption.Value == "advanced" {
					children = append(children,
						layout.Rigid(material.RadioButton(th, &advancedOption, "pro", "Pro Features").Layout),
						layout.Rigid(material.RadioButton(th, &advancedOption, "expert", "Expert Features").Layout),
					)
				}
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
			})
			e.Frame(gtx.Ops)
		}
	}
}

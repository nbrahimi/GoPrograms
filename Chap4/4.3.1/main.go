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
		w.Option(app.Title("Grouped checkbox"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		if err := run(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

var (
	selectAll widget.Bool
	item1     widget.Bool
	item2     widget.Bool
	item3     widget.Bool
)

func updateSelectAll() {
	if item1.Value && item2.Value && item3.Value {
		selectAll.Value = true
	} else {
		selectAll.Value = false
	}
}

func toggleAll() {
	if selectAll.Value {
		item1.Value, item2.Value, item3.Value = true, true, true
	} else {
		item1.Value, item2.Value, item3.Value = false, false, false
	}
}

func layoutGroupedCheckboxes(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			dims := material.CheckBox(th, &selectAll, "Select All").Layout(gtx)
			if selectAll.Value {
				toggleAll()
			} else {
				updateSelectAll()
			}
			return dims
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			dims := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(material.CheckBox(th, &item1, "Item 1").Layout),
				layout.Rigid(material.CheckBox(th, &item2, "Item 2").Layout),
				layout.Rigid(material.CheckBox(th, &item3, "Item 3").Layout),
			)
			clicked := item1.Value || item2.Value || item3.Value
			if clicked {
				updateSelectAll()
			}
			return dims
		}),
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
				return layoutGroupedCheckboxes(gtx, theme)
			})
			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			os.Exit(0)
		}
	}
}

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
		width := 700
		length := 500
		w.Option(app.Title("4.6-Login"))
		w.Option(app.Size(unit.Dp(width), unit.Dp(length)))
		w.Option(app.MaxSize(unit.Dp(width), unit.Dp(length)))
		w.Option(app.MinSize(unit.Dp(width), unit.Dp(length)))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

var (
	usernameEdt, passwordEdt, pinEdt widget.Editor
	submitBtn, cancelBtn             widget.Clickable
	adminCb                          widget.Bool
	phoneCallOrMsg                   widget.Enum
	ops                              op.Ops
)

type (
	C = layout.Context
	D = layout.Dimensions
)

func loop(w *app.Window) error {
	th := material.NewTheme()
	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			layoutWindowTest(gtx, th)
			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			os.Exit(0)
		}
	}
}

func layoutWindowTest(gtx C, th *material.Theme) {
	margins := layout.UniformInset(unit.Dp(5))
	layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		layout.Flexed(0.7, func(gtx C) D {
			return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
				layout.Flexed(0.3, func(gtx C) D {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Flexed(0.5, func(gtx C) D {
							return margins.Layout(gtx, func(gtx C) D {
								return material.Button(th, &cancelBtn, "LeftTopLeft").Layout(gtx)
							})
						}),
						layout.Flexed(0.5, func(gtx C) D {
							return margins.Layout(gtx, func(gtx C) D {
								return material.Button(th, &cancelBtn, "LeftTopRight").Layout(gtx)
							})
						}),
					)
				}),
				layout.Flexed(0.3, func(gtx C) D {
					return margins.Layout(gtx, func(gtx C) D {
						return material.Button(th, &cancelBtn, "LeftMiddle").Layout(gtx)
					})
				}),
				layout.Flexed(0.4, func(gtx C) D {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						layout.Flexed(0.5, func(gtx C) D {
							return margins.Layout(gtx, func(gtx C) D {
								return material.Button(th, &cancelBtn, "LeftBottomRight").Layout(gtx)
							})
						}),
						layout.Flexed(0.5, func(gtx C) D {
							return margins.Layout(gtx, func(gtx C) D {
								return material.Button(th, &cancelBtn, "LeftBottomLeft").Layout(gtx)
							})
						}),
					)
				}),
			)
		}),
		layout.Flexed(0.3, func(gtx C) D {
			return margins.Layout(gtx, func(gtx C) D {
				return material.Button(th, &cancelBtn, "MostRight").Layout(gtx)
			})
		}),
	)
}

package main

import (
	"image/color"
	"log"
	"os"
	"unicode"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("TextField Editor"))
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
	var nameField, limitedField, priceField, numberField component.TextField

	for {
		switch e := w.Event().(type) {

		case app.DestroyEvent:
			os.Exit(0)

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,

					// Name field - Show the UI
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return nameField.Layout(gtx, th, "Name*")
					}),

					// Field with limited number of character
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if limitedField.TextTooLong() {
							limitedField.SetError("Exceeded the max number of characters ")
						} else {
							limitedField.ClearError()
						}
						limitedField.CharLimit = 10
						limitedField.Helper = "This field have a limited character count: "
						return limitedField.Layout(gtx, th, "Max number of character is 10")
					}),

					// Price field. Shows the currency and number of decimals
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						priceField.Prefix = func(gtx layout.Context) layout.Dimensions {
							th := *th
							th.Palette.Fg = color.NRGBA{R: 100, G: 100, B: 100, A: 255}
							return material.Label(&th, th.TextSize, "$").Layout(gtx)
						}
						priceField.Suffix = func(gtx layout.Context) layout.Dimensions {
							th := *th
							th.Palette.Fg = color.NRGBA{R: 100, G: 100, B: 100, A: 255}
							return material.Label(&th, th.TextSize, ".00").Layout(gtx)
						}
						priceField.SingleLine = true
						return priceField.Layout(gtx, th, "Price")
					}),
					// Number field
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if err := func() string {
							for _, r := range numberField.Text() {
								if !unicode.IsDigit(r) {
									return "Only digits are accepted"
								}
							}
							return ""
						}(); err != "" {
							numberField.SetError(err)
						} else {
							numberField.ClearError()
						}
						numberField.SingleLine = true
						return numberField.Layout(gtx, th, "Numbers")
					}),
				)
			})
			e.Frame(gtx.Ops)
		}
	}
}

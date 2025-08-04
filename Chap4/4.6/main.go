package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"strings"

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
		width := 800
		length := 400
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
	switchVIP                        widget.Bool
	phoneCallOrMsg                   widget.Enum
	ops                              op.Ops
	admin                            bool = false
	vip                              bool = false
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
			if cancelBtn.Clicked(gtx) {
				os.Exit(0)
			}
			if submitBtn.Clicked(gtx) {
				submitFct()
			}
			layoutWindow(gtx, th)
			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			os.Exit(0)
		}
	}
}

func submitFct() {
	username := strings.TrimSpace(usernameEdt.Text())
	password := strings.TrimSpace(passwordEdt.Text())
	if len(username) == 0 || len(password) == 0 {
		fmt.Println("Error: Either Username or Password field or both are empty")
	} else {
		phone := phoneCallOrMsg.Value
		if len(phone) == 0 {
			fmt.Println("Error: You need to pick how you want to receive the PIN")
		} else {
			pin := strings.TrimSpace(pinEdt.Text())
			if len(pin) == 0 {
				fmt.Println("Error: Pin is empty")
			} else {
				user := "Guest"
				if adminCb.Value {
					user = "Admin"
				}
				if switchVIP.Value {
					fmt.Printf("\n\nSuccess: Username %s is connected as %s with VIP authorities\nPin %s was received via phone %s\n", username, user, pin, phone)
				} else {
					fmt.Printf("\n\nSuccess: Username %s is connected as %s \nPin %s was received via phone %s\n", username, user, pin, phone)
				}
				os.Exit(0)
			}
		}
	}
}

func layoutWindow(gtx C, th *material.Theme) {
	layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		// Left Side: Will be divided horizontally to 3 sides, top, middle and bottom
		layout.Flexed(0.7, func(gtx C) D {
			return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
				// Left top side. Will be divided horizontally to 2 sides: Used for labels and editors
				layout.Flexed(0.3, func(gtx C) D {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						// Left top left side: Used for labels
						layout.Flexed(0.5, func(gtx C) D {
							return layoutLeftTopLeftSide(gtx, th)
						}),
						// Left top right side: Used for editors
						layout.Flexed(0.5, func(gtx C) D {
							return layoutLeftTopRightSide(gtx, th)
						}),
					)
				}),
				// Left middle side: Used for the checkbox and the switch
				layout.Flexed(0.3, func(gtx C) D {
					return layoutLeftMiddleSide(gtx, th)
				}),
				// Left bottom left side: Used for radiobuttons, checkbox and switch
				layout.Flexed(0.4, func(gtx C) D {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
						// Left bottom left side: Used for options
						layout.Flexed(0.5, func(gtx C) D {
							return layoutLeftBottomLeftSide(gtx, th)
						}),
						// Left bottom right side: Used for checkbox and switch
						layout.Flexed(0.5, func(gtx C) D {
							return layoutLeftBottomRightSide(gtx, th)
						}),
					)
				}),
			)
		}),
		// Most Right Side: Used for buttons
		layout.Flexed(0.3, func(gtx C) D {
			return layoutMostRightSide(gtx, th)
		}),
	)
}

func layoutMostRightSide(gtx C, th *material.Theme) D {
	margins := layout.Inset{Top: unit.Dp(15), Left: unit.Dp(15), Right: unit.Dp(3), Bottom: unit.Dp(3)}
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return margins.Layout(gtx, func(gtx C) D {
				return material.Button(th, &submitBtn, "Submit").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx C) D {
			return margins.Layout(gtx, func(gtx C) D {
				return material.Button(th, &cancelBtn, "Cancel").Layout(gtx)
			})
		}),
	)
}

func layoutLeftTopLeftSide(gtx C, th *material.Theme) D {
	return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle | layout.Baseline}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return layout.Inset{Top: unit.Dp(12)}.Layout(gtx, func(gtx C) D {
				return layoutLabel(gtx, th, unit.Sp(16), "Username: ")
			})
		}),
		layout.Rigid(func(gtx C) D {
			return layout.Inset{Top: unit.Dp(20)}.Layout(gtx, func(gtx C) D {
				return layoutLabel(gtx, th, unit.Sp(16), "Password: ")
			})
		}),
	)
}

func layoutLeftMiddleSide(gtx C, th *material.Theme) D {
	return layout.Center.Layout(gtx, func(gtx C) D {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return layout.Inset{Top: unit.Dp(12)}.Layout(gtx, func(gtx C) D {
					return layoutLabel(gtx, th, unit.Sp(16), "Pin #: ")
				})
			}),
			layout.Rigid(func(gtx C) D {
				return layout.Inset{Top: unit.Dp(10), Bottom: unit.Dp(30)}.Layout(gtx, func(gtx C) D {
					return layoutEditor(gtx, th, &pinEdt, "Enter PIN")
				})
			}),
		)
	})
}

func layoutLeftTopRightSide(gtx C, th *material.Theme) D {
	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Start, // Align editors to the left
	}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return layout.Inset{Top: unit.Dp(10)}.Layout(gtx, func(gtx C) D {
				return layoutEditor(gtx, th, &usernameEdt, "Enter username")
			})
		}),
		layout.Rigid(func(gtx C) D {
			return layout.Inset{Top: unit.Dp(10)}.Layout(gtx, func(gtx C) D {
				ed := layoutEditor(gtx, th, &passwordEdt, "Enter password")
				passwordEdt.Mask = '*'
				return ed
			})
		}),
	)
}

func layoutLeftBottomLeftSide(gtx C, th *material.Theme) D {
	return layout.Center.Layout(gtx, func(gtx C) D {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Bottom: unit.Dp(5)}.Layout(gtx,
					material.Body1(th, "Receive the PIN via:").Layout,
				)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Bottom: unit.Dp(5)}.Layout(gtx,
					material.RadioButton(th, &phoneCallOrMsg, "call", "Phone call").Layout,
				)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Bottom: unit.Dp(5)}.Layout(gtx,
					material.RadioButton(th, &phoneCallOrMsg, "sms", "Phone SMS").Layout,
				)
			}),
		)
	})
}

func layoutLeftBottomRightSide(gtx C, th *material.Theme) D {
	margins := layout.Inset{Top: unit.Dp(1), Bottom: unit.Dp(3), Right: unit.Dp(5)}
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(1), Right: unit.Dp(5)}.Layout(gtx, func(gtx C) D {
						return material.Label(th, unit.Sp(16), "Admin? ").Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return margins.Layout(gtx, func(gtx C) D {
						return material.CheckBox(th, &adminCb, "").Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(1), Right: unit.Dp(5)}.Layout(gtx, func(gtx C) D {
						if adminCb.Value {
							return material.Body1(th, "Admin access").Layout(gtx)
						}
						return D{}
					})
				}),
			)
		}),
		layout.Rigid(func(gtx C) D {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(1), Right: unit.Dp(5)}.Layout(gtx, func(gtx C) D {
						return material.Label(th, unit.Sp(16), "VIP? ").Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return margins.Layout(gtx, func(gtx C) D {
						return material.Switch(th, &switchVIP, "vip").Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(1), Right: unit.Dp(5)}.Layout(gtx, func(gtx C) D {
						if switchVIP.Value {
							return material.Body1(th, "VIP access").Layout(gtx)
						}
						return D{}
					})
				}),
			)
		}),
	)
}

func layoutLabel(gtx C, th *material.Theme, textSize unit.Sp, label string) D {
	lbl := material.Label(th, textSize, label)
	lbl.Alignment = text.End
	return lbl.Layout(gtx)
}

func layoutEditor(gtx C, th *material.Theme, editor *widget.Editor, hint string) D {
	editor.SingleLine = true
	border := widget.Border{
		Color:        color.NRGBA{R: 0, G: 0, B: 0, A: 255}, // Black border
		Width:        unit.Dp(2),                            // Border width
		CornerRadius: unit.Dp(2),
	}
	return border.Layout(gtx, func(gtx C) D {
		return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx C) D {
			return material.Editor(th, editor, hint).Layout(gtx)
		})
	})
}

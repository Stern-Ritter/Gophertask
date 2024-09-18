package client

import (
	"github.com/rivo/tview"
)

func (c *ClientImpl) ShowInfoModal(text string, currentView tview.Primitive) tview.Primitive {
	modal := tview.NewModal().
		SetText(text).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				selectView(c.app, currentView)
			}
		})

	selectView(c.app, modal)

	return modal
}

func (c *ClientImpl) ShowConfirmModal(text string, currentView tview.Primitive, handler func()) tview.Primitive {
	modal := tview.NewModal().
		SetText(text).
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Yes" {
				handler()
			} else {
				selectView(c.app, currentView)
			}
		})

	selectView(c.app, modal)

	return modal
}

func (c *ClientImpl) ShowRetryModal(text string, currentView tview.Primitive, previousView tview.Primitive) tview.Primitive {
	modal := tview.NewModal().
		SetText(text).
		AddButtons([]string{"Retry", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Retry" {
				selectView(c.app, currentView)
			} else {
				selectView(c.app, previousView)
			}
		})

	selectView(c.app, modal)

	return modal
}

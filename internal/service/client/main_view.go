package client

import "github.com/rivo/tview"

func (c *ClientImpl) MainView() tview.Primitive {
	menu := tview.NewList()
	menu.AddItem("Add new task: task, epic, subtask", "",
		'a', func() { c.AddView() }).
		AddItem("View tasks: tasks, epics", "",
			'v', func() { c.DataView() }).
		AddItem("Application version info", "", 'i', func() { c.VersionView() }).
		AddItem("Logout", "", 'q', func() { c.AuthView() })

	selectView(c.app, menu)

	return menu
}

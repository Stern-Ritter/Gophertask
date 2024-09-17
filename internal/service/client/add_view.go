package client

import (
	"fmt"

	"github.com/rivo/tview"
	"google.golang.org/grpc/status"

	"github.com/Stern-Ritter/gophertask/internal/utils"
)

func (c *ClientImpl) AddView() tview.Primitive {
	menu := tview.NewList()
	menu.AddItem("Add task", "", 't', func() { c.AddTaskView() }).
		AddItem("Add epic", "", 'e', func() { c.AddEpicView() }).
		AddItem("Back", "", 'b', func() { c.MainView() })

	c.SelectView(menu)

	return menu
}

func (c *ClientImpl) AddTaskView() tview.Primitive {
	form := tview.NewForm()
	form.AddInputField("Name", "", 20, nil, nil).
		AddInputField("Description", "", 20, nil, nil).
		AddInputField("Duration (HH:MM:SS)", "", 20, nil, nil).
		AddInputField("Start date (DD.MM.YYYY)", "", 20, nil, nil).
		AddButton("Add", func() { addTaskHandler(c, c.taskService, form) }).
		AddButton("Cancel", func() { c.MainView() })

	c.SelectView(form)

	return form
}

func addTaskHandler(c Client, taskService TaskService, form *tview.Form) {
	name := form.GetFormItemByLabel("Name").(*tview.InputField).GetText()
	description := form.GetFormItemByLabel("Description").(*tview.InputField).GetText()
	durationStr := form.GetFormItemByLabel("Duration (HH:MM:SS)").(*tview.InputField).GetText()
	startedAtStr := form.GetFormItemByLabel("Start date (DD.MM.YYYY)").(*tview.InputField).GetText()

	duration, err := utils.ParseDuration(durationStr)
	if err != nil {
		c.ShowInfoModal(fmt.Sprintf("Error adding task: %s", err.Error()), form)
		return
	}
	startedAt, err := utils.ParseDate(startedAtStr)
	if err != nil {
		c.ShowInfoModal("Error adding task: invalid date format", form)
	}

	err = taskService.CreateTask(name, description, duration, startedAt)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			c.ShowInfoModal(fmt.Sprintf("Error adding task: %s", err.Error()), form)
		} else {
			errMsg := st.Message()
			c.ShowInfoModal(fmt.Sprintf("Error adding task: %s", errMsg), form)
		}
	} else {
		c.ShowInfoModal("Success added task", c.AddTaskView())
	}
}

func (c *ClientImpl) AddEpicView() tview.Primitive {
	form := tview.NewForm()
	form.AddInputField("Name", "", 20, nil, nil).
		AddInputField("Description", "", 20, nil, nil).
		AddInputField("Start date (DD.MM.YYYY)", "", 20, nil, nil).
		AddButton("Add", func() { addEpicHandler(c, c.epicService, form) }).
		AddButton("Cancel", func() { c.MainView() })

	c.SelectView(form)

	return form
}

func addEpicHandler(c Client, epicService EpicService, form *tview.Form) {
	name := form.GetFormItemByLabel("Name").(*tview.InputField).GetText()
	description := form.GetFormItemByLabel("Description").(*tview.InputField).GetText()
	startedAtStr := form.GetFormItemByLabel("Start date (DD.MM.YYYY)").(*tview.InputField).GetText()

	startedAt, err := utils.ParseDate(startedAtStr)
	if err != nil {
		c.ShowInfoModal("Error adding epic: invalid date format", form)
	}

	err = epicService.CreateEpic(name, description, startedAt)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			c.ShowInfoModal(fmt.Sprintf("Error adding epic: %s", err.Error()), form)
		} else {
			errMsg := st.Message()
			c.ShowInfoModal(fmt.Sprintf("Error adding epic: %s", errMsg), form)
		}
	} else {
		c.ShowInfoModal("Success added epic", c.AddEpicView())
	}
}

func (c *ClientImpl) AddSubtaskView(epicID string) tview.Primitive {
	form := tview.NewForm()
	form.AddInputField("Name", "", 20, nil, nil).
		AddInputField("Description", "", 20, nil, nil).
		AddInputField("Duration (HH:MM:SS)", "", 20, nil, nil).
		AddInputField("Start date (DD.MM.YYYY)", "", 20, nil, nil).
		AddButton("Add", func() { addSubtaskHandler(c, c.subtaskService, epicID, form) }).
		AddButton("Cancel", func() { c.MainView() })

	c.SelectView(form)

	return form
}

func addSubtaskHandler(c Client, subtaskService SubtaskService, epicID string, form *tview.Form) {
	name := form.GetFormItemByLabel("Name").(*tview.InputField).GetText()
	description := form.GetFormItemByLabel("Description").(*tview.InputField).GetText()
	durationStr := form.GetFormItemByLabel("Duration (HH:MM:SS)").(*tview.InputField).GetText()
	startedAtStr := form.GetFormItemByLabel("Start date (DD.MM.YYYY)").(*tview.InputField).GetText()

	duration, err := utils.ParseDuration(durationStr)
	if err != nil {
		c.ShowInfoModal(fmt.Sprintf("Error adding subtask: %s", err.Error()), form)
		return
	}
	startedAt, err := utils.ParseDate(startedAtStr)
	if err != nil {
		c.ShowInfoModal("Error adding subtask: invalid date format", form)
	}

	err = subtaskService.CreateSubtask(epicID, name, description, duration, startedAt)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			c.ShowInfoModal(fmt.Sprintf("Error adding subtask: %s", err.Error()), form)
		} else {
			errMsg := st.Message()
			c.ShowInfoModal(fmt.Sprintf("Error adding subtask: %s", errMsg), form)
		}
	} else {
		c.ShowInfoModal("Success added subtask", c.AddTaskView())
	}
}

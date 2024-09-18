package client

import (
	"fmt"

	"github.com/rivo/tview"
	"google.golang.org/grpc/status"

	"github.com/Stern-Ritter/gophertask/internal/model"
	pb "github.com/Stern-Ritter/gophertask/proto/gen/gophertask/gophertaskapi/v1"

	"github.com/Stern-Ritter/gophertask/internal/utils"
)

var statusOptions = []string{"NEW", "IN PROGRESS", "DONE"}

func (c *ClientImpl) UpdateTaskView(task *pb.TaskV1) tview.Primitive {
	form := tview.NewForm()
	form.AddInputField("Name", task.GetName(), 20, nil, nil).
		AddInputField("Description", task.GetDescription(), 20, nil, nil).
		AddDropDown("Status", statusOptions, getStatusSelectedIndex(task.GetStatus()), nil).
		AddInputField("Duration (HH:MM:SS)", utils.FormatDuration(task.GetDuration()), 20, nil, nil).
		AddInputField("Start date (DD.MM.YYYY)", utils.FormatTimestamp(task.GetStartedAt()), 10, nil, nil).
		AddButton("Update", func() { updateTaskHandler(c, c.taskService, task.GetId(), form) }).
		AddButton("Cancel", func() { c.MainView() })

	c.SelectView(form)

	return form
}

func updateTaskHandler(c Client, taskService TaskService, id string, form *tview.Form) {
	name := form.GetFormItemByLabel("Name").(*tview.InputField).GetText()
	description := form.GetFormItemByLabel("Description").(*tview.InputField).GetText()
	option, _ := form.GetFormItemByLabel("Status").(*tview.DropDown).GetCurrentOption()
	statusOption := model.TaskStatus(statusOptions[option])
	durationStr := form.GetFormItemByLabel("Duration (HH:MM:SS)").(*tview.InputField).GetText()
	startedAtStr := form.GetFormItemByLabel("Start date (DD.MM.YYYY)").(*tview.InputField).GetText()

	duration, err := utils.ParseDuration(durationStr)
	if err != nil {
		c.ShowInfoModal(fmt.Sprintf("Error updating task: %s", err.Error()), form)
		return
	}
	startedAt, err := utils.ParseDate(startedAtStr)
	if err != nil {
		c.ShowInfoModal("Error updating task: invalid date format", form)
	}

	err = taskService.UpdateTask(id, name, description, statusOption, duration, startedAt)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			c.ShowInfoModal(fmt.Sprintf("Error updating task: %s", err.Error()), form)
		} else {
			errMsg := st.Message()
			c.ShowInfoModal(fmt.Sprintf("Error updating task: %s", errMsg), form)
		}
	} else {
		c.ShowInfoModal("Success updated task", c.TasksView(c.MainView()))
	}
}

func (c *ClientImpl) UpdateEpicView(epic *pb.EpicV1) tview.Primitive {
	form := tview.NewForm()
	form.AddInputField("Name", epic.GetName(), 20, nil, nil).
		AddInputField("Description", epic.GetDescription(), 20, nil, nil).
		AddInputField("Start date (DD.MM.YYYY)", utils.FormatTimestamp(epic.GetStartedAt()), 10, nil, nil).
		AddButton("Update", func() { updateEpicHandler(c, c.epicService, epic.GetId(), form) }).
		AddButton("Cancel", func() { c.MainView() })

	c.SelectView(form)

	return form
}

func updateEpicHandler(c Client, epicService EpicService, id string, form *tview.Form) {
	name := form.GetFormItemByLabel("Name").(*tview.InputField).GetText()
	description := form.GetFormItemByLabel("Description").(*tview.InputField).GetText()
	startedAtStr := form.GetFormItemByLabel("Start date (DD.MM.YYYY)").(*tview.InputField).GetText()

	startedAt, err := utils.ParseDate(startedAtStr)
	if err != nil {
		c.ShowInfoModal("Error updating epic: invalid date format", form)
	}

	err = epicService.UpdateEpic(id, name, description, startedAt)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			c.ShowInfoModal(fmt.Sprintf("Error updating epic: %s", err.Error()), form)
		} else {
			errMsg := st.Message()
			c.ShowInfoModal(fmt.Sprintf("Error updating epic: %s", errMsg), form)
		}
	} else {
		c.ShowInfoModal("Success updated epic", c.EpicsView(c.MainView()))
	}
}

func (c *ClientImpl) UpdateSubtaskView(subtask *pb.SubtaskV1) tview.Primitive {
	form := tview.NewForm()
	form.AddInputField("Name", subtask.GetName(), 20, nil, nil).
		AddInputField("Description", subtask.GetDescription(), 20, nil, nil).
		AddDropDown("Status", statusOptions, getStatusSelectedIndex(subtask.GetStatus()), nil).
		AddInputField("Duration (HH:MM:SS)", utils.FormatDuration(subtask.GetDuration()), 20, nil, nil).
		AddInputField("Start date (DD.MM.YYYY)", utils.FormatTimestamp(subtask.GetStartedAt()), 10, nil, nil).
		AddButton("Update", func() {
			updateSubtaskHandler(c, c.subtaskService, subtask.GetId(), subtask.GetEpicId(), form)
		}).
		AddButton("Cancel", func() { c.MainView() })

	c.SelectView(form)

	return form
}

func updateSubtaskHandler(c Client, subtaskService SubtaskService, id string, epicID string, form *tview.Form) {
	name := form.GetFormItemByLabel("Name").(*tview.InputField).GetText()
	description := form.GetFormItemByLabel("Description").(*tview.InputField).GetText()
	option, _ := form.GetFormItemByLabel("Status").(*tview.DropDown).GetCurrentOption()
	statusOption := model.TaskStatus(statusOptions[option])
	durationStr := form.GetFormItemByLabel("Duration (HH:MM:SS)").(*tview.InputField).GetText()
	startedAtStr := form.GetFormItemByLabel("Start date (DD.MM.YYYY)").(*tview.InputField).GetText()

	duration, err := utils.ParseDuration(durationStr)
	if err != nil {
		c.ShowInfoModal(fmt.Sprintf("Error updating subtask: %s", err.Error()), form)
		return
	}
	startedAt, err := utils.ParseDate(startedAtStr)
	if err != nil {
		c.ShowInfoModal("Error updating subtask: invalid date format", form)
	}

	err = subtaskService.UpdateSubtask(id, epicID, name, description, statusOption, duration, startedAt)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			c.ShowInfoModal(fmt.Sprintf("Error updating subtask: %s", err.Error()), form)
		} else {
			errMsg := st.Message()
			c.ShowInfoModal(fmt.Sprintf("Error updating subtask: %s", errMsg), form)
		}
	} else {
		c.ShowInfoModal("Success updated subtask", c.SubtasksView(epicID, c.EpicsView(c.MainView())))
	}
}

func getStatusSelectedIndex(status pb.TaskStatus) int {
	currentStatus := model.MapMessageTaskStatusToTaskStatus(status)
	selectedIndex := 0
	for i, statusOption := range statusOptions {
		if string(currentStatus) == statusOption {
			selectedIndex = i
			break
		}
	}

	return selectedIndex
}

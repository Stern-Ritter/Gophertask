package client

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"google.golang.org/grpc/status"

	"github.com/Stern-Ritter/gophertask/internal/model"
	"github.com/Stern-Ritter/gophertask/internal/utils"
)

const (
	subtasksBtnText = "[blue::bl]SUBTASKS"
	updateBtnText   = "[green::bl]UPDATE"
	deleteBtnText   = "[yellow::bl]DELETE"
)

func (c *ClientImpl) DataView() tview.Primitive {
	menu := tview.NewList()
	menu.AddItem("View tasks", "", 't', func() { c.TasksView(menu) }).
		AddItem("View epics", "", 'e', func() { c.EpicsView(menu) }).
		AddItem("Back", "", 'b', func() { c.MainView() })

	c.SelectView(menu)

	return menu
}

func (c *ClientImpl) TasksView(previousView tview.Primitive) tview.Primitive {
	table := tview.NewTable().SetSelectable(true, true).SetBorders(true)
	setTableHeader(table, []string{"Name", "Description", "Status", "Duration", "Started at", "Update", "Delete"})
	addTableListeners(table, c.app, previousView)

	tasks, err := c.taskService.GetAllTasks()
	if err != nil {
		c.ShowInfoModal(fmt.Sprintf("Get tasks data error: %s", err.Error()), previousView)
		return table
	}

	for i, task := range tasks {
		row := i + 1
		column := getColumnCounter()
		table.SetCell(row, column(), tview.NewTableCell(task.GetName()))
		table.SetCell(row, column(), tview.NewTableCell(task.GetDescription()))
		table.SetCell(row, column(), tview.NewTableCell(string(model.MapMessageTaskStatusToTaskStatus(task.GetStatus()))))
		table.SetCell(row, column(), tview.NewTableCell(utils.FormatDuration(task.GetDuration())))
		table.SetCell(row, column(), tview.NewTableCell(utils.FormatTimestamp(task.GetStartedAt())))
		table.SetCell(row, column(), newClickableCell(updateBtnText, func() {
			c.UpdateTaskView(task)
		}))
		table.SetCell(row, column(), newClickableCell(deleteBtnText, func() {
			deleteTaskHandler(c, c.taskService, task.GetId(), table, previousView)
		}))
	}

	c.SelectView(table)

	return table
}

func deleteTaskHandler(c Client, taskService TaskService, taskID string, currentView tview.Primitive,
	previousView tview.Primitive) {
	c.ShowConfirmModal("Are you sure you want to delete this task?", currentView,
		func() {
			err := taskService.DeleteTask(taskID)
			if err != nil {
				st, ok := status.FromError(err)
				if !ok {
					c.ShowInfoModal(fmt.Sprintf("Failed to delete task: %s", err.Error()), currentView)
					return
				} else {
					errMsg := st.Message()
					c.ShowInfoModal(fmt.Sprintf("Failed to delete task: %s", errMsg), currentView)
					return
				}
			}
			c.ShowInfoModal("Task deleted successfully", c.TasksView(previousView))
		})
}

func (c *ClientImpl) EpicsView(previousView tview.Primitive) tview.Primitive {
	table := tview.NewTable().SetSelectable(true, true).SetBorders(true)
	setTableHeader(table, []string{"Name", "Description", "Status", "Duration", "Started at", "Subtask", "Update", "Delete"})
	addTableListeners(table, c.app, previousView)

	epics, err := c.epicService.GetAllEpics()
	if err != nil {
		c.ShowInfoModal(fmt.Sprintf("Get epics data error: %s", err.Error()), previousView)
		return table
	}

	for i, epic := range epics {
		row := i + 1
		column := getColumnCounter()
		table.SetCell(row, column(), tview.NewTableCell(epic.GetName()))
		table.SetCell(row, column(), tview.NewTableCell(epic.GetDescription()))
		table.SetCell(row, column(), tview.NewTableCell(string(model.MapMessageTaskStatusToTaskStatus(epic.GetStatus()))))
		table.SetCell(row, column(), tview.NewTableCell(utils.FormatDuration(epic.GetDuration())))
		table.SetCell(row, column(), tview.NewTableCell(utils.FormatTimestamp(epic.GetStartedAt())))
		table.SetCell(row, column(), newClickableCell(subtasksBtnText, func() {
			c.SubtasksView(epic.GetId(), table)
		}))
		table.SetCell(row, column(), newClickableCell(updateBtnText, func() {
			c.UpdateEpicView(epic)
		}))
		table.SetCell(row, column(), newClickableCell(deleteBtnText, func() {
			deleteEpicHandler(c, c.epicService, epic.GetId(), table, previousView)
		}))
	}

	c.SelectView(table)

	return table
}

func deleteEpicHandler(c Client, epicService EpicService, epicID string, currentView tview.Primitive,
	previousView tview.Primitive) {
	c.ShowConfirmModal("Are you sure you want to delete this epic with all subtasks?", currentView,
		func() {
			err := epicService.DeleteEpic(epicID)
			if err != nil {
				st, ok := status.FromError(err)
				if !ok {
					c.ShowInfoModal(fmt.Sprintf("Failed to delete epic: %s", err.Error()), currentView)
					return
				} else {
					errMsg := st.Message()
					c.ShowInfoModal(fmt.Sprintf("Failed to delete epic: %s", errMsg), currentView)
					return
				}
			}
			c.ShowInfoModal("Epic deleted successfully", c.EpicsView(previousView))
		})
}

func (c *ClientImpl) SubtasksView(epicID string, previousView tview.Primitive) tview.Primitive {
	table := tview.NewTable().SetSelectable(true, true).SetBorders(true)
	setTableHeader(table, []string{"Name", "Description", "Status", "Duration", "Started at", "Update", "Delete"})
	addTableListeners(table, c.app, previousView)

	epic, err := c.epicService.GetEpicByID(epicID)
	if err != nil {
		c.ShowInfoModal(fmt.Sprintf("Get epic subtasks data error: %s", err.Error()), previousView)
		return table
	}
	if epic == nil {
		c.ShowInfoModal("Get epic subtasks data error", previousView)
		return table
	}

	subtasks := epic.GetSubtasks()
	for i, subtask := range subtasks {
		row := i + 1
		column := getColumnCounter()
		table.SetCell(row, column(), tview.NewTableCell(subtask.GetName()))
		table.SetCell(row, column(), tview.NewTableCell(subtask.GetDescription()))
		table.SetCell(row, column(), tview.NewTableCell(string(model.MapMessageTaskStatusToTaskStatus(subtask.GetStatus()))))
		table.SetCell(row, column(), tview.NewTableCell(utils.FormatDuration(subtask.GetDuration())))
		table.SetCell(row, column(), tview.NewTableCell(utils.FormatTimestamp(subtask.GetStartedAt())))
		table.SetCell(row, column(), newClickableCell(updateBtnText, func() {
			c.UpdateSubtaskView(subtask)
		}))
		table.SetCell(row, column(), newClickableCell(deleteBtnText, func() {
			deleteSubtaskHandler(c, c.subtaskService, epicID, subtask.GetId(), table, previousView)
		}))
	}

	addButton := tview.NewButton("Add subtask").SetSelectedFunc(func() {
		c.AddSubtaskView(epicID)
	})

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(table, 0, 1, true).
		AddItem(addButton, 1, 10, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			if c.app.GetFocus() == table {
				c.app.SetFocus(addButton)
			} else {
				c.app.SetFocus(table)
			}
			return nil
		}
		return event
	})

	c.SelectView(flex)

	return flex
}

func deleteSubtaskHandler(c Client, subtaskService SubtaskService, epicID string, subtaskID string, currentView tview.Primitive,
	previousView tview.Primitive) {
	c.ShowConfirmModal("Are you sure you want to delete this subtask?", currentView,
		func() {
			err := subtaskService.DeleteSubtask(subtaskID)
			if err != nil {
				st, ok := status.FromError(err)
				if !ok {
					c.ShowInfoModal(fmt.Sprintf("Failed to delete subtask: %s", err.Error()), currentView)
					return
				} else {
					errMsg := st.Message()
					c.ShowInfoModal(fmt.Sprintf("Failed to delete subtask: %s", errMsg), currentView)
					return
				}
			}
			c.ShowInfoModal("Subtask deleted successfully", c.SubtasksView(epicID, c.SubtasksView(epicID, previousView)))
		})
}

func addTableListeners(table *tview.Table, app *tview.Application, previousView tview.Primitive) {
	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			row, column := table.GetSelection()
			cell := table.GetCell(row, column)
			switch cell.Text {
			case subtasksBtnText, updateBtnText, deleteBtnText:
				if cell.Clicked != nil {
					cell.Clicked()
				}
			}
		case tcell.KeyEsc:
			app.SetRoot(previousView, true)
			return nil
		}
		return event
	})
}

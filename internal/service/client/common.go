package client

import "github.com/rivo/tview"

func (c *ClientImpl) SelectView(view tview.Primitive) {
	selectView(c.app, view)
}

func selectView(app *tview.Application, view tview.Primitive) {
	app.SetRoot(view, true)
	app.SetFocus(view)
}

func setTableHeader(table *tview.Table, columns []string) {
	for idx, column := range columns {
		table.SetCell(0, idx, tview.NewTableCell(column))
	}
}

func getColumnCounter() func() int {
	currentColumn := -1
	return func() int {
		currentColumn++
		return currentColumn
	}
}

func newClickableCell(text string, handler func()) *tview.TableCell {
	cell := tview.NewTableCell(text).
		SetAlign(tview.AlignCenter).
		SetClickedFunc(func() bool {
			handler()
			return false
		})

	return cell
}

func stopApp(app *tview.Application) {
	app.Stop()
}

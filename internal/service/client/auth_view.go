package client

import (
	"fmt"

	"github.com/rivo/tview"
	"google.golang.org/grpc/status"
)

func (c *ClientImpl) AuthView() tview.Primitive {
	form := tview.NewForm()
	form.AddInputField("Login", "", 20, nil, nil).
		AddInputField("Password", "", 20, nil, nil).
		AddButton("Sign in", func() { c.signInHandler(form) }).
		AddButton("Sign up", func() { c.signUpHandler(form) }).
		AddButton("Quit", func() { stopApp(c.app) })

	c.SelectView(form)

	return form
}

func (c *ClientImpl) signUpHandler(form *tview.Form) {
	login := form.GetFormItemByLabel("Login").(*tview.InputField).GetText()
	password := form.GetFormItemByLabel("Password").(*tview.InputField).GetText()

	token, err := c.authService.SignUp(login, password)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			c.ShowRetryModal(fmt.Sprintf("Registration failed: %s. Try again or cancel.", err.Error()), form, form)
		} else {
			errMsg := st.Message()
			c.ShowRetryModal(fmt.Sprintf("Registration failed: %s. Try again or cancel.", errMsg), form, form)
		}
		return
	}

	c.authToken = token
	c.ShowInfoModal("Registration success", c.MainView())
}

func (c *ClientImpl) signInHandler(form *tview.Form) {
	login := form.GetFormItemByLabel("Login").(*tview.InputField).GetText()
	password := form.GetFormItemByLabel("Password").(*tview.InputField).GetText()

	token, err := c.authService.SignIn(login, password)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			c.ShowRetryModal(fmt.Sprintf("Login failed: %s. Try again or cancel.", err.Error()), form, form)
		} else {
			errMsg := st.Message()
			c.ShowRetryModal(fmt.Sprintf("Login failed: %s. Try again or cancel.", errMsg), form, form)
		}
		return
	}

	c.authToken = token
	c.MainView()
}

package controllers

type ExitController struct {
	BaseController
}

func (e *ExitController) Get() {
	e.DelSession("loginuser")
	e.Redirect("/", 302)
}

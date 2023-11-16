package controllers

import "github.com/labstack/echo/v4"

type StudentController struct {
}

func (controller StudentController) getData(c echo.Context) error {

}

func NewStudentController() StudentController {
	return StudentController{}
}

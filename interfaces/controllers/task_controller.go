package controllers

import (
	"github.com/labstack/echo"
	"github.com/moko-poi/go-api-todo/domain"
	"github.com/moko-poi/go-api-todo/interfaces/database"
	"github.com/moko-poi/go-api-todo/usecase"
	"strconv"
)

type TaskController struct {
	Interactor usecase.TaskInteractor
}

func NewTaskController(sqlHandler database.SqlHandler) *TaskController {
	return &TaskController{
		Interactor: usecase.TaskInteractor{
			TaskRepository: &database.TaskRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TaskController) Create(c echo.Context) (err error) {
	u := domain.Task{}
	c.Bind(&u)
	err = controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, "completed")
	return
}

func (controller *TaskController) Index(c echo.Context) (err error) {
	tasks, err := controller.Interactor.Tasks()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, tasks)
	return
}

func (controller *TaskController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := controller.Interactor.TaskById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, task)
	return
}

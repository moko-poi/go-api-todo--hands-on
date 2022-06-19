package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/moko-poi/go-api-todo/interfaces/controllers"
)

func Init() {
	// Echo instance
	e := echo.New()

	taskController := controllers.NewTaskController(NewSqlHandler())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// CORSの設定追加
	e.Use(middleware.CORS())

	e.GET("/tasks", func(c echo.Context) error { return taskController.Index(c) })
	e.GET("/tasks/:id", func(c echo.Context) error { return taskController.Show(c) })
	e.POST("/create", func(c echo.Context) error { return taskController.Create(c) })

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

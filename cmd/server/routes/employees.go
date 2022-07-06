package routes

import (
	// "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/database"
	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	employees "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee"

	// deletar depois
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store"

	"github.com/gin-gonic/gin"
	// _ "github.com/go-sql-driver/mysql"
)

func Employees(routerGroup *gin.RouterGroup) {

	// substituir depois
	// employeesRepository := employees.NewRepository(database.GetInstance())

	file := store.New(store.FileType, "../../internal/employee/employees.json")
	employeesRepository := employees.NewRepository(file)
	employeesService := employees.NewService(employeesRepository)
	employeesHandler := handler.NewEmployee(employeesService)

	employeesRouterGroup := routerGroup.Group("/employees")
	{
		employeesRouterGroup.POST("/", employeesHandler.Create())
		employeesRouterGroup.GET("/", employeesHandler.GetAll())
		employeesRouterGroup.GET("/:id", employeesHandler.GetById())
		employeesRouterGroup.PATCH("/:id", employeesHandler.Update())
		employeesRouterGroup.DELETE("/:id", employeesHandler.Delete())
	}
}

package routes

import (
	"employee_app/internal/app/employee/controller"
	"employee_app/internal/app/employee/repository"
	"employee_app/internal/app/employee/usecase"

	_ "employee_app/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	EmployeeRepository := repository.NewEmployeeRepository(db)
	employeeUseCase := usecase.NewEmployeeUseCase(*EmployeeRepository)
	employeeController := controller.NewEmployeeController(employeeUseCase)

	employees := r.Group("/employees")
	{
		employees.GET("/", employeeController.GetEmployees)

		employees.GET("/:id", employeeController.GetEmployee)

		employees.POST("/", employeeController.CreateEmployee)

		employees.PUT("/:id", employeeController.UpdateEmployee)

		employees.DELETE("/:id", employeeController.DeleteEmployee)

		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	}
}

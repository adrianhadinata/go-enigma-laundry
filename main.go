package main

import (
	"submission-project-enigma-laundry/config"
	"submission-project-enigma-laundry/entity"
	"submission-project-enigma-laundry/test"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	apiGroup := router.Group("/api")
	{
		v1Api := apiGroup.Group("/v1")
		{
			customersGroup := v1Api.Group("/customers")
			{
				customersGroup.POST("/", entity.InsertCustomer)
				customersGroup.PUT("/:id", entity.UpdateCustomer)
				customersGroup.DELETE("/:id", entity.DeleteCustomer)
				customersGroup.GET("/", entity.FindAllCustomers)
			}

			employeesGroup := v1Api.Group("/employees")
			{
				employeesGroup.POST("/", entity.InsertEmployee)
				employeesGroup.PUT("/:id", entity.UpdateEmployee)
				employeesGroup.DELETE("/:id", entity.DeleteEmployee)
				employeesGroup.GET("/", entity.FindAllEmployees)
			}

			servicesGroup := v1Api.Group("/services")
			{
				servicesGroup.POST("/", entity.InsertService)
				servicesGroup.PUT("/:id", entity.UpdateService)
				servicesGroup.DELETE("/:id", entity.DeleteService)
				servicesGroup.GET("/", entity.FindAllServices)
			}

			transactionsGroup := v1Api.Group("/transactions")
			{
				transactionsGroup.POST("/", entity.InsertTransaction)
				transactionsGroup.GET("/", entity.FindAllTransaction)
			}

			testDBConnection := v1Api.Group("/db")
			{
				testDBConnection.GET("/", test.TestConnectDB)
			}
		}
	}

	err := router.Run(config.ApiPort())

	if err != nil {
		panic(err)
	}
}
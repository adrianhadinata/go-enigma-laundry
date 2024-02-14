package entity

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"
	"submission-project-enigma-laundry/config"

	"github.com/gin-gonic/gin"
)

type Service struct {
	Id          int		`json:"id"`
	ServiceName string	`json:"serviceName"`
	Unit        string	`json:"unit"`
	Price       int		`json:"price"`
}

func FindAllServices(c *gin.Context) {
	config.TableName = "mst_services"

	db := config.ConnectDB()
	defer db.Close()

	// SQL query
	query := "SELECT * FROM " + config.TableName

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	services := ScanDataservice(rows)

	name := c.Query("name")

	if name == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Success Get Services!", "data": services})
		return
	}

	var matchedServices []Service

	for _, service := range services {
		if strings.Contains(strings.ToLower(service.ServiceName), strings.ToLower(name)) {
			matchedServices = append(matchedServices, service)
		}
	}

	if len(matchedServices) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Success Get Service(s)!", "data": matchedServices})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error" : "Service not found"})
}

func FindOneService(id int) (Service, error) {
	config.TableName = "mst_services"
	db := config.ConnectDB()
	defer db.Close()

	query := "SELECT * FROM " + config.TableName + " WHERE id = $1"
	service := Service{}

	var err error
	err = db.QueryRow(query, id).Scan(&service.Id, &service.ServiceName, &service.Unit, &service.Price)
	
	return service, err
}

func InsertService(c *gin.Context) {
	config.TableName = "mst_services"

	db := config.ConnectDB()
	defer db.Close()

	// SQL query
	query := "INSERT INTO " + config.TableName + "(id, service_name, unit, price) VALUES ($1, $2, $3, $4)"

	var newService Service

	err := c.ShouldBind(&newService)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}

	// Duplicated Id Validation
	isIdDuplicated, err := FindOneService(newService.Id)

	if isIdDuplicated.Id == newService.Id {
		err = errors.New("Service Id Already Exist")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}

	// Service Name Validation
	if len(newService.ServiceName) >= 100 || len(newService.ServiceName) == 0 {
		err = errors.New("Service Name length must greater than 0 and under 100 character")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}

	_, err = db.Exec(query, newService.Id, newService.ServiceName, newService.Unit, newService.Price)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() })
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Success Insert New Service!", "data": newService})
}

func DeleteService(c *gin.Context) {
	config.TableName = "mst_services"

	db := config.ConnectDB()
	defer db.Close()
	var err error
	
	id := c.Param("id")

	query := "DELETE FROM "+ config.TableName +" WHERE id = $1;"
	_, err = db.Exec(query, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete Success!"})
}

func UpdateService(c *gin.Context) {
	config.TableName = "mst_services"

	db := config.ConnectDB()
	defer db.Close()

	var updatedService Service
	err := c.ShouldBind(&updatedService)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() })
		return
	}

	// SQL query
	query := "UPDATE " + config.TableName + " SET service_name = $2, unit = $3, price = $4 WHERE id = $1;"

	if len(updatedService.ServiceName) >= 100 || len(updatedService.ServiceName) == 0 {
		err = errors.New("Service Name length must greater than 0 and under 100 character")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() })
		return
	}

	if len(updatedService.Unit) >= 100 || len(updatedService.Unit) == 0 {
		err = errors.New("Service Unit length must greater than 0 and under 100 character")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() })
		return
	}

	_, err = db.Exec(query, updatedService.Id, updatedService.ServiceName, updatedService.Unit, updatedService.Price)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() })
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success Insert uPDATE Service!", "data": updatedService})
	c.JSON(http.StatusOK, updatedService)
}

func ScanDataservice(rows *sql.Rows) []Service {
	services := []Service{}
	var err error

	for rows.Next() {
		service := Service{}
		err := rows.Scan(&service.Id, &service.ServiceName, &service.Unit, &service.Price)

		if err != nil {
			panic(err)
		}

		services = append(services, service)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return services
}
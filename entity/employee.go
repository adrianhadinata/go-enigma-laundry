package entity

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"
	"submission-project-enigma-laundry/config"
	"time"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Id           int		`json:"id"`
	EmployeeName string		`json:"employeeName"`
	DateCreated  time.Time	`json:"dateCreated"`
}

func FindAllEmployees(c *gin.Context) {
	config.TableName = "mst_employees"

	db := config.ConnectDB()
	defer db.Close()

	// SQL query
	query := "SELECT * FROM " + config.TableName

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	employees := ScanDataEmployee(rows)
	
	name := c.Query("name")

	if name == "" {
		c.JSON(http.StatusOK, employees)
		return
	}

	var matchedEmployees []Employee

	for _, employee := range employees {
		if strings.Contains(strings.ToLower(employee.EmployeeName), strings.ToLower(name)) {
			matchedEmployees = append(matchedEmployees, employee)
		}
	}

	if len(matchedEmployees) > 0 {
		c.JSON(http.StatusOK, matchedEmployees)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error" : "Employee not found"})
}

func FindOneEmployee(id int) (Employee, error) {
	config.TableName = "mst_employees"
	db := config.ConnectDB()
	defer db.Close()

	query := "SELECT * FROM " + config.TableName + " WHERE id = $1"
	employee := Employee{}

	var err error
	err = db.QueryRow(query, id).Scan(&employee.Id, &employee.EmployeeName, &employee.DateCreated)
	
	return employee, err
}

func InsertEmployee(c *gin.Context) {
	config.TableName = "mst_employees"

	db := config.ConnectDB()
	defer db.Close()

	// SQL query
	query := "INSERT INTO " + config.TableName + "(id, employee_name, date_created) VALUES ($1, $2, $3)"

	var newEmployee Employee

	err := c.ShouldBind(&newEmployee)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}

	// Duplicated Id Validation
	isIdDuplicated, err := FindOneEmployee(newEmployee.Id)

	if isIdDuplicated.Id == newEmployee.Id {
		err = errors.New("Employee Id Already Exist")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}

	// Employee Name Validation
	if len(newEmployee.EmployeeName) >= 100 || len(newEmployee.EmployeeName) == 0 {
		err = errors.New("Employee Name length must greater than 0 and under 100 character")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}

	_, err = db.Exec(query, newEmployee.Id, newEmployee.EmployeeName, newEmployee.DateCreated)

	if err != nil {
		err = errors.New("Cannot Insert to Database")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() })
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Success Insert New Employee!", "data": newEmployee})
}

func DeleteEmployee(c *gin.Context) {
	config.TableName = "mst_employees"

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

func UpdateEmployee(c *gin.Context) {
	config.TableName = "mst_Employees"

	db := config.ConnectDB()
	defer db.Close()

	var updatedEmployee Employee
	err := c.ShouldBind(&updatedEmployee)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() })
		return
	}

	// SQL query
	query := "UPDATE " + config.TableName + " SET employee_name = $2, date_created = $3 WHERE id = $1;"

	// Employee Name Validation
	if len(updatedEmployee.EmployeeName) >= 100 || len(updatedEmployee.EmployeeName) == 0 {
		err = errors.New("Employee Name length must greater than 0 and under 100 character")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() })
		return
	}

	_, err = db.Exec(query, updatedEmployee.Id, updatedEmployee.EmployeeName, updatedEmployee.DateCreated)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() })
		return
	}

	c.JSON(http.StatusOK, updatedEmployee)
}

func ScanDataEmployee(rows *sql.Rows) []Employee {
	employees := []Employee{}
	var err error

	for rows.Next() {
		employee := Employee{}
		err := rows.Scan(&employee.Id, &employee.EmployeeName, &employee.DateCreated)

		if err != nil {
			panic(err)
		}

		employees = append(employees, employee)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return employees
}
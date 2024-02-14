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

type Customer struct {
	Id           int		`json:"id"`
	CustomerName string		`json:"customerName"`
	PhoneNumber  string		`json:"phoneNumber"`
	DateCreated  time.Time	`json:"dateCreated"`
}

func FindAllCustomers(c *gin.Context) {
	config.TableName = "mst_customers"

	db := config.ConnectDB()
	defer db.Close()

	// SQL query
	query := "SELECT * FROM " + config.TableName

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	customers := ScanDataCustomer(rows)

	name := c.Query("name")

	if name == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Success Get Customer(s)!", "data": customers})
		return
	}

	var matchedCustomers []Customer

	for _, customer := range customers {
		if strings.Contains(strings.ToLower(customer.CustomerName), strings.ToLower(name)) {
			matchedCustomers = append(matchedCustomers, customer)
		}
	}

	if len(matchedCustomers) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Success Get Customer(s)!", "data": matchedCustomers})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error" : "Customer not found"})
}

func FindOneCustomer(id int) (Customer, error) {
	config.TableName = "mst_customers"
	db := config.ConnectDB()
	defer db.Close()

	query := "SELECT * FROM " + config.TableName +" WHERE id = $1"
	customer := Customer{}

	var err error
	err = db.QueryRow(query, id).Scan(&customer.Id, &customer.CustomerName, &customer.PhoneNumber, &customer.DateCreated)

	return customer, err
}

func InsertCustomer(c *gin.Context) {
	config.TableName = "mst_customers"

	db := config.ConnectDB()
	defer db.Close()

	// SQL query
	query := "INSERT INTO " + config.TableName + "(id, customer_name, phone_number, date_created) VALUES ($1, $2, $3, $4)"

	var newCustomer Customer

	err := c.ShouldBind(&newCustomer)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}

	// Duplicated Id Validation
	isIdDuplicated, err := FindOneCustomer(newCustomer.Id)

	if isIdDuplicated.Id == newCustomer.Id {
		err = errors.New("Customer Id Already Exist")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}

	// Customer Name Validation
	if len(newCustomer.CustomerName) >= 100 || len(newCustomer.CustomerName) == 0 {
		err = errors.New("Customer Name length must greater than 0 and under 100 character")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}

	_, err = db.Exec(query, newCustomer.Id, newCustomer.CustomerName, newCustomer.PhoneNumber, newCustomer.DateCreated)

	if err != nil {
		err = errors.New("Cannot Insert to Database")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() })
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Success Insert New Customer!", "data": newCustomer})
}

func DeleteCustomer(c *gin.Context) {
	config.TableName = "mst_customers"

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

func UpdateCustomer(c *gin.Context) {
	config.TableName = "mst_customers"

	db := config.ConnectDB()
	defer db.Close()

	var updatedCustomer Customer
	err := c.ShouldBind(&updatedCustomer)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() })
		return
	}

	// SQL query
	query := "UPDATE " + config.TableName + " SET Customer_name = $2, phone_number = $3, date_created = $4 WHERE id = $1;"

	// Customer Name Validation
	if len(updatedCustomer.CustomerName) >= 100 || len(updatedCustomer.CustomerName) == 0 {
		err = errors.New("Customer Name length must greater than 0 and under 100 character")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}

	_, err = db.Exec(query, updatedCustomer.Id, updatedCustomer.CustomerName, updatedCustomer.PhoneNumber, updatedCustomer.DateCreated)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() })
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Success Update Customer!", "data": updatedCustomer})
}

func ScanDataCustomer(rows *sql.Rows) []Customer {
	customers := []Customer{}
	var err error

	for rows.Next() {
		customer := Customer{}
		err := rows.Scan(&customer.Id, &customer.CustomerName, &customer.PhoneNumber, &customer.DateCreated)

		if err != nil {
			panic(err)
		}

		customers = append(customers, customer)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return customers
}
package entity

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"submission-project-enigma-laundry/config"
	"time"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id            	int			`json:"id"`
	IdEmployee    	int			`json:"idEmployee"`
	IdCustomer    	int			`json:"idCustomer"`
	IdService     	int			`json:"idService"`
	Amount 			int			`json:"amount"`
	TransactionIn 	time.Time	`json:"transactionIn"`
	TransactionOut 	time.Time	`json:"transactionOut"`
}

func FindAllTransaction(c *gin.Context) {
	config.TableName = "tx_enigma_laundry"

	db := config.ConnectDB()
	defer db.Close()

	// SQL query
	query := "SELECT id, id_employee, id_customer, id_service, amount, transaction_in, transaction_out FROM " + config.TableName

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	transactions := ScanDataTransaction(rows)

	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusOK, transactions)
		return
	}

	integerId, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	var matchedTransactions []Transaction

	for _, transaction := range transactions {
		if transaction.Id == integerId {
			matchedTransactions = append(matchedTransactions, transaction)
		}
	}

	if len(matchedTransactions) > 0 {
		c.JSON(http.StatusOK, matchedTransactions)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error" : "Transaction not found"})
}

func FindOneTransaction(id int) (Transaction, error) {
	config.TableName = "tx_enigma_laundry"
	db := config.ConnectDB()
	defer db.Close()

	query := "SELECT * FROM " + config.TableName + " WHERE id = $1"
	transaction := Transaction{}

	var err error
	err = db.QueryRow(query, id).Scan(&transaction.Id, &transaction.IdCustomer, &transaction.IdEmployee, &transaction.IdService, &transaction.TransactionIn, &transaction.TransactionOut, &transaction.Amount)
	
	return transaction, err
}

func InsertTransaction(c *gin.Context) {
	config.TableName = "tx_enigma_laundry"

	db := config.ConnectDB()
	defer db.Close()
	
	var newTransaction Transaction
	err := c.ShouldBind(&newTransaction)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}

	// SQL query
	query := "INSERT INTO " + config.TableName + "(id, id_employee, id_customer, id_service, amount, transaction_in, transaction_out) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	// Duplicated Id Validation
	isIdDuplicated, err := FindOneTransaction(newTransaction.Id)

	if isIdDuplicated.Id == newTransaction.Id {
		err = errors.New("Transaction Id Already Exist")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}

	// Employee Validation
	isEmployeeExist, err := FindOneEmployee(newTransaction.IdEmployee)

	if isEmployeeExist.Id != newTransaction.IdEmployee {
		err = errors.New("Employee with id: " + strconv.Itoa(newTransaction.IdEmployee) + " doesnt exist!")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}

	// Customer Validation
	isCustomerExist, err := FindOneEmployee(newTransaction.IdCustomer)

	if isCustomerExist.Id != newTransaction.IdCustomer {
		err = errors.New("Customer with id: " + strconv.Itoa(newTransaction.IdCustomer) + " doesnt exist!")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}

	// Customer Validation
	isServiceExist, err := FindOneEmployee(newTransaction.IdService)

	if isServiceExist.Id != newTransaction.IdService {
		err = errors.New("Service with id: " + strconv.Itoa(newTransaction.IdService) + " doesnt exist!")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}

	// Amount Validation
	if newTransaction.Amount <= 0 {
		err = errors.New("Amount must be greater than 0!")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}

	_, err = db.Exec(query, newTransaction.Id, newTransaction.IdEmployee, newTransaction.IdCustomer, newTransaction.IdService, newTransaction.Amount, newTransaction.TransactionIn, newTransaction.TransactionOut)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() })
		return
	}

	c.JSON(http.StatusCreated, newTransaction)
}

func ScanDataTransaction(rows *sql.Rows) []Transaction {
	transactions := []Transaction{}
	var err error

	for rows.Next() {
		transaction := Transaction{}

		err := rows.Scan(&transaction.Id, &transaction.IdEmployee, &transaction.IdCustomer, &transaction.IdService, &transaction.Amount, &transaction.TransactionIn, &transaction.TransactionOut)

		if err != nil {
			panic(err)
		}

		transactions = append(transactions, transaction)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return transactions
}
package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/izaakdale/goBank/errs"
	"github.com/izaakdale/utils-go/logger"
	"github.com/jmoiron/sqlx"
)

type CustomerRepoDb struct {
	client *sqlx.DB
}

func NewCustomerRepoDb(dbClient *sqlx.DB) CustomerRepoDb {
	return CustomerRepoDb{
		dbClient,
	}
}

func (crdb CustomerRepoDb) FindAll(status string) ([]Customer, *errs.AppError) {

	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select * from customers"
		err = crdb.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select * from customers where status=?"
		err = crdb.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB error")
	}

	return customers, nil
}

func (crdb CustomerRepoDb) FindById(id string) (*Customer, *errs.AppError) {

	var customer Customer
	var err error
	customerSql := "select * from customers where customer_id = ?"
	err = crdb.client.Get(&customer, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("No Rows " + err.Error())
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error scanning customer " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB error")
	}

	return &customer, nil
}

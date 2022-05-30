package domain

import (
	"database/sql"
	"strconv"

	"github.com/izaakdale/goBank/errs"
	"github.com/izaakdale/utils-go/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepoDb struct {
	client *sqlx.DB
}

func NewAccountRepoDb(dbClient *sqlx.DB) AccountRepoDb {
	return AccountRepoDb{
		dbClient,
	}
}

func (ardb AccountRepoDb) SaveAccount(account Account) (*Account, *errs.AppError) {

	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?,?,?,?,?)"
	result, err := ardb.client.Exec(sqlInsert, account.CustomerId, account.OpeningDate, account.AccountType, account.Balance, account.Status)

	if err != nil {
		logger.Error("Error while creating account " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB Error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting lastInsertId for new account " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB Error")
	}
	account.AccountId = strconv.FormatInt(id, 10)

	return &account, nil
}

func (ardb AccountRepoDb) SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError) {

	tranXBlock, err := ardb.client.Begin()
	if err != nil {
		logger.Error("Error while start db transaction block " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB Error")
	}
	sqlInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES (?,?,?,?)"
	result, err := tranXBlock.Exec(sqlInsert, transaction.AccountId, transaction.Amount, transaction.TransactionType, transaction.TransactionDate)
	if err != nil {
		logger.Error("Error executing insert on transactions table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB Error")
	}

	if transaction.IsWithdrawal() {
		sqlUpdate := "UPDATE accounts SET amount = amount - ?  WHERE account_id = ?"
		_, err = tranXBlock.Exec(sqlUpdate, transaction.Amount, transaction.AccountId)
	} else {
		sqlUpdate := "UPDATE accounts SET amount = amount + ?  WHERE account_id = ?"
		_, err = tranXBlock.Exec(sqlUpdate, transaction.Amount, transaction.AccountId)
	}

	if err != nil {
		tranXBlock.Rollback()
		logger.Error("Error when updating account amount after transaction, rolling back " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB Error")
	}
	err = tranXBlock.Commit()
	if err != nil {
		logger.Error("Error committing transaction changes to accounts and transactions tables " + err.Error())
		return nil, errs.NewNotFoundError("Unexpected DB Error")
	}

	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error getting transaction ID " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB Error")
	}

	account, appError := ardb.FindById(transaction.AccountId)
	if appError != nil {
		return nil, appError
	}
	transaction.TransactionId = strconv.FormatInt(transactionId, 10)
	transaction.Amount = account.Balance

	return &transaction, nil
}

func (ardb AccountRepoDb) FindById(id string) (*Account, *errs.AppError) {

	var account Account
	var err error
	customerSql := "select * from accounts where account_id = ?"
	err = ardb.client.Get(&account, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("No Rows " + err.Error())
			return nil, errs.NewNotFoundError("Account not found")
		}
		logger.Error("Error scanning customer " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB error")
	}

	return &account, nil
}

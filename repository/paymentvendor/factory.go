package paymentvendor

import (
	"backend_capstone/services/paymentvendor"
	"backend_capstone/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) paymentvendor.Repository {
	var transactionRepo paymentvendor.Repository

	if dbCon.Driver == utils.Postgres {
		transactionRepo = NewPostgresRepository(dbCon.Postgres)
	}

	return transactionRepo
}
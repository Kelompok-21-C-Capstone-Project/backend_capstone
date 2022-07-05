package paymentmethod

import (
	"backend_capstone/services/paymentmethod"
	"backend_capstone/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) paymentmethod.Repository {
	var transactionRepo paymentmethod.Repository

	if dbCon.Driver == utils.Postgres {
		transactionRepo = NewPostgresRepository(dbCon.Postgres)
	} else if dbCon.Driver == utils.MySQL {
		transactionRepo = NewPostgresRepository(dbCon.Postgres)
	}

	return transactionRepo
}

package transactions

import (
	"backend_capstone/services/transactions"
	"backend_capstone/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) transactions.Repository {
	var transactionsRepo transactions.Repository

	if dbCon.Driver == utils.Postgres {
		transactionsRepo = NewPostgresRepository(dbCon.Postgres)
	}

	return transactionsRepo
}

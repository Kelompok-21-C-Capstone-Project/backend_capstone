package transaction

import (
	"backend_capstone/services/transaction"
	"backend_capstone/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) transaction.Repository {
	var transactionRepo transaction.Repository

	if dbCon.Driver == utils.Postgres {
		transactionRepo = NewPostgresRepository(dbCon.Postgres)
	} else if dbCon.Driver == utils.MySQL {
		transactionRepo = NewPostgresRepository(dbCon.Postgres)
	}

	return transactionRepo
}

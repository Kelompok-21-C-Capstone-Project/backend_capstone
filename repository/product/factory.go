package products

import (
	"backend_capstone/services/product"
	"backend_capstone/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) product.Repository {
	var transactionsRepo product.Repository

	if dbCon.Driver == utils.Postgres {
		transactionsRepo = NewPostgresRepository(dbCon.Postgres)
	} else if dbCon.Driver == utils.MySQL {
		transactionsRepo = NewPostgresRepository(dbCon.MySQL)
	}

	return transactionsRepo
}

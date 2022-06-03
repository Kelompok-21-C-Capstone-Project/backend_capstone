package products

import (
	"backend_capstone/services/products"
	"backend_capstone/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) products.Repository {
	var transactionsRepo products.Repository

	if dbCon.Driver == utils.Postgres {
		transactionsRepo = NewPostgresRepository(dbCon.Postgres)
	}

	return transactionsRepo
}

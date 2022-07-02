package productcategory

import (
	"backend_capstone/services/productcategory"
	"backend_capstone/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) productcategory.Repository {
	var transactionRepo productcategory.Repository

	if dbCon.Driver == utils.Postgres {
		transactionRepo = NewPostgresRepository(dbCon.Postgres)
	}
	return transactionRepo
}

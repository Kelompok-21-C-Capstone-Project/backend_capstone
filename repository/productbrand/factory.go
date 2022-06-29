package productbrand

import (
	"backend_capstone/services/productbrand"
	"backend_capstone/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) productbrand.Repository {
	var transactionRepo productbrand.Repository

	if dbCon.Driver == utils.Postgres {
		transactionRepo = NewPostgresRepository(dbCon.Postgres)
	}

	return transactionRepo
}

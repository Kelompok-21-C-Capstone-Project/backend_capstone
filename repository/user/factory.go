package user

import (
	"backend_capstone/services/user"
	"backend_capstone/utils"
	"log"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) user.Repository {
	var transactionsRepo user.Repository
	log.Print(dbCon)

	if dbCon.Driver == utils.Postgres {
		transactionsRepo = NewPostgresRepository(dbCon.Postgres)
	} else if dbCon.Driver == utils.MySQL {
		transactionsRepo = NewPostgresRepository(dbCon.MySQL)
	}

	return transactionsRepo
}

package services

import (
	"database/sql"

	"github.com/morgances/hydra/server/services/mysql"
)

var (
	// AdminService -
	AdminService *mysql.AdminServiceImpl
)

// Load all services.
func Load(db *sql.DB) error {
	AdminService = &mysql.AdminServiceImpl{
		DB: db,
	}

	if err := AdminService.Initialize(); err != nil {
		panic(err)
	}

	AdminService.Create("administrator", "111111", "", "", "")

	return nil
}

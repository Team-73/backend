package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GuiaBolso/darwin"
	"github.com/Team-73/backend/data/migrations"
	"github.com/Team-73/backend/domain/contract"
	"github.com/Team-73/backend/infra/config"

	_ "github.com/go-sql-driver/mysql" //Used to connect to database
)

// DBManager is the MySQL connection manager
type DBManager struct {
	db *sql.DB
}

//Instance retunrs an instance of a RepoManager
func Instance() (contract.RepoManager, error) {
	cfg := config.GetDBConfig()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
	)

	log.Println("Connecting to database...")

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Database successfully configured")

	log.Println("Running the migrations")
	driver := darwin.NewGenericDriver(db, darwin.MySQLDialect{})

	d := darwin.New(driver, migrations.Migrations, nil)

	err = d.Migrate()
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	log.Println("Migrations executed")

	instance := &DBManager{
		db: db,
	}

	return instance, nil
}

//Ping returns a session to use cassadra querys
func (c *DBManager) Ping() contract.PingRepo {
	return nil
}

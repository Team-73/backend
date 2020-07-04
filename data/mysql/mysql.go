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

//Ping returns a session to use mysql querys
func (c *DBManager) Ping() contract.PingRepo {
	return nil
}

//Business returns a session to use mysql querys
func (c *DBManager) Business() contract.BusinessRepo {
	return newBusinessRepo(c.db)
}

//Category returns a session to use mysql querys
func (c *DBManager) Category() contract.CategoryRepo {
	return newCategoryRepo(c.db)
}

//Company returns a session to use mysql querys
func (c *DBManager) Company() contract.CompanyRepo {
	return newCompanyRepo(c.db)
}

//Product returns a session to use mysql querys
func (c *DBManager) Product() contract.ProductRepo {
	return newProductRepo(c.db)
}

//User returns a session to use mysql querys
func (c *DBManager) User() contract.UserRepo {
	return newUserRepo(c.db)
}

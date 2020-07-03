package migrations

import "github.com/GuiaBolso/darwin"

//Migrations to execute our queries that changes database structure
var (
	Migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Creating table users",
			Script: `CREATE TABLE IF NOT EXISTS users (
				id INT AUTO_INCREMENT,
				name VARCHAR(100) NOT NULL,
				email VARCHAR(50) NOT NULL,
				password VARCHAR(100) NOT NULL,
				document_number VARCHAR(50) NOT NULL,
				area_code VARCHAR(3) NOT NULL,
				phone_number VARCHAR(9) NOT NULL,
				birthdate DATE NOT NULL,
				gender VARCHAR(1) NOT NULL,
				revenue INT(11) NOT NULL DEFAULT 0,
				active TINYINT(1) NOT NULL DEFAULT 1,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				PRIMARY KEY (id),
				UNIQUE INDEX ID_UNIQUE (id ASC),
				UNIQUE INDEX EMAIL_UNIQUE (email ASC)
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
	}
)

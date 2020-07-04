package migrations

import "github.com/GuiaBolso/darwin"

//Migrations to execute our queries that changes database structure
var (
	Migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Creating table tab_user",
			Script: `CREATE TABLE IF NOT EXISTS tab_user (
				id INT NOT NULL AUTO_INCREMENT,
				name VARCHAR(100) NOT NULL,
				email VARCHAR(50) NOT NULL,
				password VARCHAR(100) NOT NULL,
				document_number VARCHAR(50) NULL,
				country_code VARCHAR(4) NOT NULL DEFAULT 55,
				area_code VARCHAR(3) NOT NULL,
				phone_number VARCHAR(9) NOT NULL,
				birthdate DATE NULL,
				gender CHAR(1) NULL,
				revenue DECIMAL(8,2) NOT NULL DEFAULT '0.00',
				active TINYINT(1) NOT NULL DEFAULT 1,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				PRIMARY KEY (id),
				UNIQUE INDEX ID_UNIQUE (id ASC),
				UNIQUE INDEX EMAIL_UNIQUE (email ASC)
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     2,
			Description: "Creating table tab_category",
			Script: `CREATE TABLE IF NOT EXISTS tab_category (
				id INT NOT NULL AUTO_INCREMENT,
				name VARCHAR(100) NOT NULL,

				PRIMARY KEY (id),
				UNIQUE INDEX ID_UNIQUE (id ASC)
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     3,
			Description: "Creating table tab_product",
			Script: `CREATE TABLE IF NOT EXISTS tab_product (
				id INT NOT NULL AUTO_INCREMENT,
				name VARCHAR(100) NOT NULL,
				description VARCHAR(500) NULL,
				price DECIMAL(7,2) NOT NULL DEFAULT '0.00',
				discount_price DECIMAL(7,2) NOT NULL DEFAULT '0.00',
				category_id INT(11) NOT NULL,
				minimum_age_for_consumption INT(3) NOT NULL DEFAULT 0,
				product_image_url VARCHAR(3000) NULL,
				time_for_preparing_minutes INT(11) NULL,

				PRIMARY KEY (id),
				INDEX fk_tab_product_tab_category_idx (id ASC),
				UNIQUE INDEX ID_UNIQUE (id ASC),
				CONSTRAINT fk_product_category
					FOREIGN KEY (category_id)
					REFERENCES chefia_db.tab_category (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     4,
			Description: "Creating table tab_business",
			Script: `CREATE TABLE IF NOT EXISTS tab_business (
				id INT NOT NULL AUTO_INCREMENT,
				name VARCHAR(100) NOT NULL,

				PRIMARY KEY (id),
				UNIQUE INDEX ID_UNIQUE (id ASC)
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     5,
			Description: "Creating table tab_company",
			Script: `CREATE TABLE IF NOT EXISTS tab_company (
				id INT NOT NULL AUTO_INCREMENT,
				name VARCHAR(100) NOT NULL,
				email VARCHAR(50) NOT NULL,
				country_code VARCHAR(4) NOT NULL DEFAULT 55,
				area_code VARCHAR(3) NOT NULL,
				phone_number VARCHAR(9) NOT NULL,
				document_number VARCHAR(50) NULL,
				website VARCHAR(3000) NULL,
				business_id INT(11) NOT NULL,
				country VARCHAR(100) NULL DEFAULT 'Brazil',
				street VARCHAR(3000) NULL,
				street_number INT(11) NULL,
				complement VARCHAR(100) NULL,
				zip_code INT(11) NULL,
				neighborhood VARCHAR(100) NULL,
				city VARCHAR(100) NULL,
				federative_unit CHAR(2) NULL,
				instagram_url VARCHAR(3000) NULL DEFAULT 'https://www.instagram.com/ambev/',
				facebook_url VARCHAR(3000) NULL DEFAULT 'https://www.facebook.com/cervejariaambev/',
				linkedin_url VARCHAR(3000) NULL DEFAULT 'https://www.linkedin.com/company/ambev/',
				twitter_url VARCHAR(3000) NULL DEFAULT 'https://twitter.com/cervejariaambev',

				PRIMARY KEY (id),
				INDEX fk_tab_company_tab_business_idx (id ASC),
				UNIQUE INDEX ID_UNIQUE (id ASC),
				CONSTRAINT fk_company_business
					FOREIGN KEY (business_id)
					REFERENCES chefia_db.tab_business (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
	}
)

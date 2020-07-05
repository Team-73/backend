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
				country_code VARCHAR(4) NOT NULL,
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
			Description: "Creating table tab_product_category",
			Script: `CREATE TABLE IF NOT EXISTS tab_product_category (
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
				time_for_preparing_minutes INT(11) NOT NULL DEFAULT 0,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				PRIMARY KEY (id),
				INDEX fk_tab_product_tab_product_category_idx (id ASC),
				UNIQUE INDEX ID_UNIQUE (id ASC),
				CONSTRAINT fk_product_category
					FOREIGN KEY (category_id)
					REFERENCES chefia_db.tab_product_category (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     4,
			Description: "Creating table tab_company_business",
			Script: `CREATE TABLE IF NOT EXISTS tab_company_business (
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
				email VARCHAR (50) NOT NULL,
				country_code VARCHAR(4) NULL,
				area_code VARCHAR(3) NOT NULL,
				phone_number VARCHAR(9) NOT NULL,
				document_number VARCHAR(50) NULL,
				website VARCHAR(3000) NOT NULL DEFAULT 'https://www.ambev.com.br/',
				business_id INT(11) NOT NULL,
				country VARCHAR(100) NULL DEFAULT 'Brazil',
				street VARCHAR(3000) NULL,
				street_number VARCHAR(11) NULL,
				complement VARCHAR(100) NULL,
				zip_code INT(11) NULL,
				neighborhood VARCHAR(100) NULL,
				city VARCHAR(100) NULL,
				federative_unit CHAR(2) NULL,
				instagram_url VARCHAR(3000) NULL DEFAULT 'https://www.instagram.com/ambev/',
				facebook_url VARCHAR(3000) NULL DEFAULT 'https://www.facebook.com/cervejariaambev/',
				linkedin_url VARCHAR(3000) NULL DEFAULT 'https://www.linkedin.com/company/ambev/',
				twitter_url VARCHAR(3000) NULL DEFAULT 'https://twitter.com/cervejariaambev',
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				PRIMARY KEY (id),
				INDEX fk_tab_company_tab_company_business_idx (id ASC),
				UNIQUE INDEX ID_UNIQUE (id ASC),
				CONSTRAINT fk_company_business
					FOREIGN KEY (business_id)
					REFERENCES chefia_db.tab_company_business (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     6,
			Description: "Creating table tab_order",
			Script: `CREATE TABLE IF NOT EXISTS tab_order (
				id INT NOT NULL AUTO_INCREMENT,
				user_id INT(11) NOT NULL,
				company_id INT(11) NULL,
				rating DECIMAL(1,1) NOT NULL DEFAULT '0.00',
				accept_tip TINYINT(1) NOT NULL,
				total_tip DECIMAL(7,2) NOT NULL DEFAULT '0.00',
				total_price DECIMAL(7,2) NOT NULL DEFAULT '0.00',
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				PRIMARY KEY (id),
				UNIQUE INDEX ID_UNIQUE (id ASC),
				INDEX fk_tab_order_tab_user_idx (id ASC),
				INDEX fk_tab_order_tab_company_idx (id ASC),
				CONSTRAINT fk_user
					FOREIGN KEY (user_id)
					REFERENCES chefia_db.tab_user (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION,
				CONSTRAINT fk_company
					FOREIGN KEY (company_id)
					REFERENCES chefia_db.tab_company (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     7,
			Description: "Creating table tab_order_product",
			Script: `CREATE TABLE IF NOT EXISTS tab_order_product (
				id INT NOT NULL AUTO_INCREMENT,
				order_id INT(11) NOT NULL,
				product_id INT(11) NOT NULL,
				quantity INT(11) NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				PRIMARY KEY (id),
				UNIQUE INDEX ID_UNIQUE (id ASC),
				INDEX fk_tab_order_product_tab_order_idx (id ASC),
				INDEX fk_tab_order_product_tab_product_idx (id ASC),
				CONSTRAINT fk_order
					FOREIGN KEY (order_id)
					REFERENCES chefia_db.tab_order (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION,
				CONSTRAINT fk_product
					FOREIGN KEY (product_id)
					REFERENCES chefia_db.tab_product (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     8,
			Description: "Creating table tab_company_rating",
			Script: `CREATE TABLE IF NOT EXISTS tab_company_rating (
				id INT NOT NULL AUTO_INCREMENT,
				user_id INT(11) NULL,
				company_id INT(11) NULL,
				customer_service INT(11) NULL,
				company_clean INT(11) NULL,
				ice_beer INT(11) NULL,
				good_food INT(11) NULL,
				would_go_back INT(11) NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				PRIMARY KEY (id),
				UNIQUE INDEX ID_UNIQUE (id ASC),
				INDEX fk_tab_company_rating_tab_user_idx (id ASC),
				INDEX fk_tab_company_rating_tab_company_idx (id ASC),
				CONSTRAINT fk_user
					FOREIGN KEY (user_id)
					REFERENCES chefia_db.tab_user (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION,
				CONSTRAINT fk_company
					FOREIGN KEY (company_id)
					REFERENCES chefia_db.tab_company (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		// ============================================== ADD FAKE DATA ==============================================
		{
			Version:     9,
			Description: "Inserting data on table tab_user",
			Script: `
				INSERT INTO tab_user 
					(id,name,email,password,document_number,country_code,area_code,phone_number,birthdate,gender,revenue)
				VALUES
				(1,'Diego Clair','diego@gmail.com','12345678','01234567890','055','011','991212475','1993-10-25','M',3500.55),
				(2,'Fernando Camilo','fernando@gmail.com','12345678','01234567890','055','011','991212475','1993-10-25','M',3500.55),
				(3,'João Hortale','joao@gmail.com','12345678','01234567890','055','011','991212475','1993-10-25','M',3500.55),
				(4,'Gisele Karpinski','gisele@gmail.com','12345678','01234567890','055','011','991212475','1993-10-25','F',3500.55),
				(5,'Guilherme Castro','guilherme@gmail.com','12345678','01234567890','055','011','991212475','1993-10-25','M',3500.55);
			`,
		},
		{
			Version:     10,
			Description: "Inserting data on table tab_product_category",
			Script: `
				INSERT INTO tab_product_category 
					(id,name)
				VALUES
				(1,'Águas'),
				(2,'Cervejas'),
				(3,'Lanches'),
				(4,'Sobremesas'),
				(5,'Sucos');
			`,
		},
		{
			Version:     11,
			Description: "Inserting data on table tab_company_business",
			Script: `
				INSERT INTO tab_company_business 
					(id,name)
				VALUES
				(1,'Bar'),
				(2,'Restaurante'),
				(3,'Bar e Restaurante'),
				(4,'Balada'),
				(5,'Lanchonete');
			`,
		},
		{
			Version:     12,
			Description: "Inserting data on table tab_product",
			Script: `
				INSERT INTO tab_product 
					(id,name,description,price,category_id,minimum_age_for_consumption,product_image_url)
				VALUES
				(1,'Stella Artois','Cerveja Stella Artois 269 ml Lata',2.09,2,18,'https://cdn.shopify.com/s/files/1/0010/3150/3987/products/Collection__Stella269.jpg?v=1592954179?nocache=0.4872827889854048'),
				(2,'Skol Puro Malte','Cerveja Skol Puro Malte 269 ml Lata',1.99,2,18,'https://cdn.shopify.com/s/files/1/0010/3150/3987/products/Collection__SkolPM269_1_0f322231-a07f-442b-8183-d0bc8b535e37.jpg?v=1593621454?nocache=0.14057180978378314'),
				(3,'Bohemia','Cerveja Bohemia 350 ml Lata (Puro Malte)',2.19,2,18,'https://cdn.shopify.com/s/files/1/0010/3150/3987/products/Collection__Bohemia_2.jpg?v=1593010929?nocache=0.06681187903758978'),
				(4,'Tônica Antárctica','Água Tônica Antárctica 350 ml Lata',1.99,1,0,'https://cdn.shopify.com/s/files/1/0010/3150/3987/products/Agua_Tonica_Antarctica_350_ml_Lata_2x_06b235e1-76cc-4da2-a0e9-cc07a77d6ad1.png?v=1565713553?nocache=0.6236592114019643'),
				(5,'Do Bem','Água de Coco Do Bem 1 L Tetra Pak',7.99,1,0,'https://cdn.shopify.com/s/files/1/0010/3150/3987/products/Suco_Do_Bem_Agua_de_Coco_1L_Tetra_Pak_2x_6f937dbb-3b6d-48ea-8c2f-42e93a86f3dd.png?v=1567752696?nocache=0.01789523325641751'),
				(6,'Do Bem','Suco Do Bem Uva Integral 1 L Vidro',9.99,5,0,'https://cdn.shopify.com/s/files/1/0010/3150/3987/products/Suco_Do_Bem_Uva_Integral_1_L_Garrafa_vidro_2x_28855c5a-2691-42ab-aff0-7bec921814b5.png?v=1567752751?nocache=0.2591075763069701'),
				(7,'X-tudo','1 Carne, Bacon, Alface, Tomate, Queijo e Cheddar',11.99,3,0,'https://img.olx.com.br/images/93/938011001780366.jpg'),
				(8,'Brigadeiro','',3.99,4,0,'https://upload.wikimedia.org/wikipedia/commons/thumb/a/a4/Brigadeiro.jpg/1200px-Brigadeiro.jpg');
			`,
		},
		{
			Version:     13,
			Description: "Inserting data on table tab_company",
			Script: `
				INSERT INTO tab_company 
					(id,name,email,country_code,area_code,phone_number,document_number,business_id,street,street_number,complement,zip_code,neighborhood,city,federative_unit)
				VALUES
				(1,'TAJ BAR','ambev@ambev.com','55','11','48732123','27281922000170',1,'Av. Otaviano Alves de Lima','1824','Empresa Ambev',02701-000,'Vila Albertina','São Paulo','SP'),
				(2,'TAJ RESTAURANTE','ambev@ambev.com','55','11','48732123','27281922000170',2,'Av. Otaviano Alves de Lima','1724','',02701-000,'Vila Albertina','São Paulo','SP'),
				(3,'TAJ BAR E RESTAURANTE','ambev@ambev.com','55','11','48732123','27281922000170',3,'Av. Otaviano Alves de Lima','1625','',02701-000,'Vila Albertina','São Paulo','SP'),
				(4,'TAJ BALADA','ambev@ambev.com','55','11','48732123','27281922000170',4,'Av. Otaviano Alves de Lima','1323','',02701-000,'Vila Albertina','São Paulo','SP'),
				(5,'TAJ LANCHONETE','ambev@ambev.com','55','11','48732123','27281922000170',5,'Av. Otaviano Alves de Lima','1854','',02701-000,'Vila Albertina','São Paulo','SP'),
				(6,'AMBEV BAR','ambev@ambev.com','55','11','48732123','27281922000170',1,'Av. Otaviano Alves de Lima','1884','',02701-000,'Vila Albertina','São Paulo','SP'),
				(7,'AMBEV RESTAURANTE','ambev@ambev.com','55','11','48732123','27281922000170',2,'Av. Otaviano Alves de Lima','1814','',02701-000,'Vila Albertina','São Paulo','SP');
			`,
		},
	}
)

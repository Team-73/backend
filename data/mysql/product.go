package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Team-73/backend/domain/entity"
	"github.com/Team-73/backend/utils/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type productRepo struct {
	db *sql.DB
}

// newProductRepo returns a instance of dbrepo
func newProductRepo(db *sql.DB) *productRepo {
	return &productRepo{
		db: db,
	}
}

const queryProductSelectBase = `
					SELECT 	p.id,
									p.name,
									p.description,
									p.price,
									p.discount_price,
									p.category_id,
									p.minimum_age_for_consumption,
									p.product_image_url,
									p.time_for_preparing_minutes
								
					FROM 		tab_product 		p 
					`

func (s *productRepo) parseProductSet(rows *sql.Rows) (products []entity.Product, err error) {
	for rows.Next() {
		product := entity.Product{}
		product, err = s.parseProduct(rows)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (s *productRepo) parseProduct(row scanner) (product entity.Product, err error) {

	err = row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.DiscountPrice,
		&product.CategoryID,
		&product.MinimumAgeForConsumption,
		&product.ProductImageURL,
		&product.TimeForPreparingMinutes,
	)

	if err != nil {
		return product, err
	}

	return product, nil
}

//GetProducts - return a list os products
func (s *productRepo) GetProducts() (*[]entity.Product, *resterrors.RestErr) {

	query := queryProductSelectBase

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0021: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetProductByID", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	var products []entity.Product

	rows, err := stmt.Query()
	if err != nil {
		errorCode := "Error 0022: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in GetProducts", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer rows.Close()

	products, err = s.parseProductSet(rows)
	if err != nil {
		errorCode := "Error 0023: "
		log.Println(fmt.Sprintf("%sError when trying to parse result in parseProductSet", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &products, nil
}

//GetProductByID - get a product by ID
func (s *productRepo) GetProductByID(id int64) (*entity.Product, *resterrors.RestErr) {

	query := queryProductSelectBase + `
		WHERE 	p.id 		= ?;`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0012: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetProductByID", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	var product entity.Product

	result := stmt.QueryRow(id)
	product, err = s.parseProduct(result)
	if err != nil {
		errorCode := "Error 0013: "
		log.Println(fmt.Sprintf("%sError when trying to execute QueryRow in GetProductByID", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &product, nil
}

// Create - to create a product on database
func (s *productRepo) Create(product entity.Product) (int64, *resterrors.RestErr) {

	query := `
		INSERT INTO tab_product (
			name,
			description,
			price,
			discount_price,
			category_id,
			minimum_age_for_consumption,
			product_image_url,
			time_for_preparing_minutes) VALUES
		(?, ?, ?, ?, ?, ?, ?, ?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0014: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Create a product", errorCode), err)
		return 0, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(
		product.Name,
		product.Description,
		product.Price,
		product.DiscountPrice,
		product.CategoryID,
		product.MinimumAgeForConsumption,
		product.ProductImageURL,
		product.TimeForPreparingMinutes)
	if err != nil {
		errorCode := "Error 0015: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Create product", errorCode), err)
		return 0, mysqlutils.HandleMySQLError(errorCode, err)
	}

	productID, err := insertResult.LastInsertId()
	if err != nil {
		errorCode := "Error 0016: "
		log.Println(fmt.Sprintf("%sError when trying to get LastInsertId in the Create product", errorCode), err)
		return 0, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return productID, nil
}

// Update - to update a product on database
func (s *productRepo) Update(product entity.Product) (*entity.Product, *resterrors.RestErr) {

	query := `
		UPDATE tab_product
			SET	name 												= ?,
					description									= ?,
					price												= ?,
					discount_price							= ?,
					category_id									= ?,
					minimum_age_for_consumption = ?,
					product_image_url 					= ?,
					time_for_preparing_minutes	= ?
			
		WHERE id	= ?;
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0017: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Update a product", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		product.Name,
		product.Description,
		product.Price,
		product.DiscountPrice,
		product.CategoryID,
		product.MinimumAgeForConsumption,
		product.ProductImageURL,
		product.TimeForPreparingMinutes,
		product.ID)
	if err != nil {
		errorCode := "Error 0018: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Update product", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &product, nil
}

// Delete - to delete a product on database
func (s *productRepo) Delete(id int64) *resterrors.RestErr {

	query := `
		DELETE FROM tab_product
		WHERE 	id			= ?;
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0019: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Delete product", errorCode), err)
		return resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		errorCode := "Error 0020: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Delete product", errorCode), err)
		return mysqlutils.HandleMySQLError(errorCode, err)
	}

	return nil
}

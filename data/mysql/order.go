package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Team-73/backend/domain/entity"
	"github.com/Team-73/backend/utils/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type orderRepo struct {
	db *sql.DB
}

// newOrderRepo returns a instance of dbrepo
func newOrderRepo(db *sql.DB) *orderRepo {
	return &orderRepo{
		db: db,
	}
}

//GetOrdersByUserID to get all orders by userID
func (s *orderRepo) GetOrdersByUserID(userID int64) (*[]entity.OrdersByUserID, *resterrors.RestErr) {

	query := `
		SELECT
			o.id,
			o.company_id,
			c.name,
			o.total_price,
			IFNULL(cr.total_rating, 0) 			AS total_rating,
			COUNT(op.id)		 								AS total_items,
			o.created_at

		FROM 				tab_order 		o

		INNER JOIN 	tab_company						c
				ON 			o.company_id					= c.id

		LEFT JOIN 	tab_rating		cr
				ON 			o.company_id					= cr.company_id
				AND			cr.user_id						= o.user_id

		LEFT JOIN 	tab_order_product			op
				ON 			o.id									= op.order_id

		WHERE 			o.user_id 						= ?
		
		GROUP BY		
			o.id, 
			o.company_id, 
			c.name, 
			o.total_price, 
			cr.total_rating, 
			o.created_at;
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0004: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetOrdersByUserID", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		errorCode := "Error 0004: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetOrdersByUserID", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}

	var orders []entity.OrdersByUserID
	var order entity.OrdersByUserID
	for rows.Next() {
		err = rows.Scan(
			&order.OrderID,
			&order.CompanyID,
			&order.CompanyName,
			&order.TotalPrice,
			&order.TotalRating,
			&order.TotalItems,
			&order.CreatedAt,
		)
		if err != nil {
			errorCode := "Error 0005: "
			log.Println(fmt.Sprintf("%sError when trying to execute Query in GetOrdersByUserID", errorCode), err)
			return nil, mysqlutils.HandleMySQLError(errorCode, err)
		}
		orders = append(orders, order)
	}

	return &orders, nil
}

//GetOrderDetail to get order detail
func (s *orderRepo) GetOrderDetail(orderID int64) (*entity.OrderDetail, *resterrors.RestErr) {

	query := `
		SELECT
			o.company_id,
			c.name,
			o.total_price,
			o.total_tip,
			o.created_at

		FROM 				tab_order 		o

		INNER JOIN 	tab_company		c
				ON 			o.company_id	= c.id

		WHERE 			o.id 					= ?;
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0004: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetOrderDetail", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	row := stmt.QueryRow(orderID)

	var order entity.OrderDetail

	err = row.Scan(
		&order.CompanyID,
		&order.CompanyName,
		&order.TotalPrice,
		&order.TotalTip,
		&order.CreatedAt,
	)
	if err != nil {
		errorCode := "Error 0005: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in GetOrderDetail", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &order, nil
}

//GetOrderProducts to get all products detail of a orderID
func (s *orderRepo) GetOrderProducts(orderID int64) ([]entity.ProductsDetail, *resterrors.RestErr) {

	query := `
		SELECT
			op.quantity,
			p.name,
		 (p.price * op.quantity)					AS total_product_price

		FROM 				tab_order_product 		op

		INNER JOIN 	tab_product						p
				ON 			op.product_id					= p.id

		WHERE 			op.order_id 					= ?;
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0004: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetOrderProducts", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	rows, err := stmt.Query(orderID)
	if err != nil {
		errorCode := "Error 0004: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetOrderProducts", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}

	var products []entity.ProductsDetail
	var product entity.ProductsDetail
	for rows.Next() {
		err = rows.Scan(
			&product.Quantity,
			&product.ProductName,
			&product.TotalProductPrice,
		)
		if err != nil {
			errorCode := "Error 0005: "
			log.Println(fmt.Sprintf("%sError when trying to execute Query in GetOrderProducts", errorCode), err)
			return nil, mysqlutils.HandleMySQLError(errorCode, err)
		}
		products = append(products, product)
	}

	return products, nil
}

// CreateOrder - to create a order on database
func (s *orderRepo) CreateOrder(order entity.Order) (int64, *resterrors.RestErr) {

	query := `
		INSERT INTO tab_order (
			user_id,company_id,accept_tip
		) VALUES (?,?,?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0006: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Create a order", errorCode), err)
		return 0, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(
		order.UserID,
		order.CompanyID,
		order.AcceptTip)
	if err != nil {
		errorCode := "Error 0007: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Create order", errorCode), err)
		return 0, mysqlutils.HandleMySQLError(errorCode, err)
	}

	orderID, err := insertResult.LastInsertId()
	if err != nil {
		errorCode := "Error 0008: "
		log.Println(fmt.Sprintf("%sError when trying to get LastInsertId in the Create order", errorCode), err)
		return 0, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return orderID, nil
}

// CreateOrderProductAndReturnProductPrice - to create a order product on database and return the product price
func (s *orderRepo) CreateOrderProductAndReturnProductPrice(orderID int64, oderProduct entity.OrderProduct) (float64, *resterrors.RestErr) {

	var price float64

	query := `
		INSERT INTO tab_order_product (
			order_id,product_id,quantity) VALUES
		(?,?,?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0006: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Create a order products", errorCode), err)
		return price, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		orderID,
		oderProduct.ProductID,
		oderProduct.Quantity)
	if err != nil {
		errorCode := "Error 0007: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Create a order products", errorCode), err)
		return price, mysqlutils.HandleMySQLError(errorCode, err)
	}

	query = `
		SELECT 	p.price
		FROM 		tab_product	p
		WHERE 	p.id = ?;
	`

	stmt, err = s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0004: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in CreateOrderProductAndReturnProductPrice", errorCode), err)
		return 0, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	row := stmt.QueryRow(orderID)
	err = row.Scan(
		&price,
	)
	if err != nil {
		errorCode := "Error 0005: "
		log.Println(fmt.Sprintf("%sError when trying to execute QueryRow in CreateOrderProductAndReturnProductPrice", errorCode), err)
		return price, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return price, nil
}

// Update - to update a order on database
func (s *orderRepo) UpdateOrder(orderID int64, order entity.Order) *resterrors.RestErr {

	query := `
		UPDATE tab_order
			SET	total_tip 		= ?,
					total_price		= ?
			
		WHERE id	= ?;
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0009: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Update a order", errorCode), err)
		return resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		order.TotalTip,
		order.TotalPrice,
		orderID)
	if err != nil {
		errorCode := "Error 0010: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Update order", errorCode), err)
		return mysqlutils.HandleMySQLError(errorCode, err)
	}

	return nil
}

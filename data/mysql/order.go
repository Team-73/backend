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

const queryOrderSelectBase = `
					SELECT 	o.id,
									o.user_id,
									o.company_id,
									o.total_tip,
									o.total_price,
									o.created_at

								
					FROM 		tab_order 		o
					`

func (s *orderRepo) parseOrderSet(rows *sql.Rows) (orders []entity.Order, err error) {
	for rows.Next() {
		order := entity.Order{}
		order, err = s.parseOrder(rows)
		if err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (s *orderRepo) parseOrder(row scanner) (order entity.Order, err error) {

	err = row.Scan(
		&order.ID,
		&order.UserID,
		&order.UserID,
		&order.UserID,
		&order.UserID,
		&order.UserID,
	)

	if err != nil {
		return order, err
	}

	return order, nil
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

func (s *orderRepo) GetOrderByUserID(userID int64) ([]entity.Order, *resterrors.RestErr) {

	query := `
		SELECT 
			o.id,
			o.user_id,
			o.company_id,
			o.accept_tip,
			o.total_tip,
			o.total_price,
			o.created_at

		FROM tab_order o 
		WHERE o.user_id = ?;
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0004: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetOrderByUserID", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		errorCode := "Error 0004: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetOrderByUserID", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}

	var orders []entity.Order
	for rows.Next() {
		order := entity.Order{}
		err = rows.Scan(
			&order.ID,
			&order.UserID,
			&order.CompanyID,
			&order.AcceptTip,
			&order.TotalTip,
			&order.TotalPrice,
			&order.CreatedAt,
		)
		if err != nil {
			errorCode := "Error 0005: "
			log.Println(fmt.Sprintf("%sError when trying to execute Query in GetOrderByUserID", errorCode), err)
			return nil, mysqlutils.HandleMySQLError(errorCode, err)
		}
		orders = append(orders, order)
	}

	return orders, nil
}

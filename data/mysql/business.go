package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Team-73/backend/domain/entity"
	"github.com/Team-73/backend/utils/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type businessRepo struct {
	db *sql.DB
}

// newBusinessRepo returns a instance of dbrepo
func newBusinessRepo(db *sql.DB) *businessRepo {
	return &businessRepo{
		db: db,
	}
}

const queryBusinessSelectBase = `
					SELECT 	b.id,
									b.name
								
					FROM 		tab_business 		b 
					`

func (s *businessRepo) parseBusinessSet(rows *sql.Rows) (businesss []entity.Business, err error) {
	for rows.Next() {
		business := entity.Business{}
		business, err = s.parseBusiness(rows)
		if err != nil {
			return businesss, err
		}
		businesss = append(businesss, business)
	}

	return businesss, nil
}

func (s *businessRepo) parseBusiness(row scanner) (business entity.Business, err error) {

	err = row.Scan(
		&business.ID,
		&business.Name,
	)

	if err != nil {
		return business, err
	}

	return business, nil
}

//GetBusinesses - return a list os businesses
func (s *businessRepo) GetBusinesses() (*[]entity.Business, *resterrors.RestErr) {

	query := queryBusinessSelectBase

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0027: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetBusinessByID", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	var businesss []entity.Business

	rows, err := stmt.Query()
	if err != nil {
		errorCode := "Error 0028: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in GetBusinesss", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer rows.Close()

	businesss, err = s.parseBusinessSet(rows)
	if err != nil {
		errorCode := "Error 0029: "
		log.Println(fmt.Sprintf("%sError when trying to parse result in parseBusinessSet", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &businesss, nil
}

//GetBusinessByID - get a business by ID
func (s *businessRepo) GetBusinessByID(id int64) (*entity.Business, *resterrors.RestErr) {

	query := queryBusinessSelectBase + `
		WHERE 	b.id 		= ?;`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0030: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetBusinessByID", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	var business entity.Business

	result := stmt.QueryRow(id)
	business, err = s.parseBusiness(result)
	if err != nil {
		errorCode := "Error 0031: "
		log.Println(fmt.Sprintf("%sError when trying to execute QueryRow in GetBusinessByID", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &business, nil
}

// Create - to create a business on database
func (s *businessRepo) Create(business entity.Business) (int64, *resterrors.RestErr) {

	query := `
		INSERT INTO tab_business (
			name) VALUES
		(?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0032: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Create a business", errorCode), err)
		return 0, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(
		business.Name)
	if err != nil {
		errorCode := "Error 0033: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Create business", errorCode), err)
		return 0, mysqlutils.HandleMySQLError(errorCode, err)
	}

	businessID, err := insertResult.LastInsertId()
	if err != nil {
		errorCode := "Error 0034: "
		log.Println(fmt.Sprintf("%sError when trying to get LastInsertId in the Create business", errorCode), err)
		return 0, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return businessID, nil
}

// Update - to update a business on database
func (s *businessRepo) Update(business entity.Business) (*entity.Business, *resterrors.RestErr) {

	query := `
		UPDATE tab_business
			SET	name 												= ?
			
		WHERE id	= ?;
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0035: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Update a business", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		business.Name,
		business.ID)
	if err != nil {
		errorCode := "Error 0036: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Update business", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &business, nil
}

// Delete - to delete a business on database
func (s *businessRepo) Delete(id int64) *resterrors.RestErr {

	query := `
		DELETE FROM tab_business
		WHERE 	id			= ?;
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0037: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Delete business", errorCode), err)
		return resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		errorCode := "Error 0038: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Delete business", errorCode), err)
		return mysqlutils.HandleMySQLError(errorCode, err)
	}

	return nil
}

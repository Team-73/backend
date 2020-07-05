package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Team-73/backend/domain/entity"
	"github.com/Team-73/backend/utils/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type companyRepo struct {
	db *sql.DB
}

// newCompanyRepo returns a instance of dbrepo
func newCompanyRepo(db *sql.DB) *companyRepo {
	return &companyRepo{
		db: db,
	}
}

const queryCompanySelectBase = `
					SELECT 	c.id,
									c.name,
									c.email,
									c.country_code,
									c.area_code,
									c.phone_number,
									c.document_number,
									c.website,
									c.business_id,
									c.country,
									c.street,
									c.street_number,
									c.complement,
									c.zip_code,
									c.neighborhood,
									c.city,
									c.federative_unit,
									c.instagram_url,
									c.facebook_url,
									c.linkedin_url,
									c.twitter_url

					FROM 		tab_company 		c 
					`

func (s *companyRepo) parseCompanySet(rows *sql.Rows) (companies []entity.Company, err error) {
	for rows.Next() {
		company := entity.Company{}
		company, err = s.parseCompany(rows)
		if err != nil {
			return companies, err
		}
		companies = append(companies, company)
	}

	return companies, nil
}

func (s *companyRepo) parseCompany(row scanner) (company entity.Company, err error) {

	err = row.Scan(
		&company.ID,
		&company.Name,
		&company.Email,
		&company.CountryCode,
		&company.AreaCode,
		&company.PhoneNumber,
		&company.DocumentNumber,
		&company.Website,
		&company.BusinessID,
		&company.Address.Country,
		&company.Address.Street,
		&company.Address.Number,
		&company.Address.Complement,
		&company.Address.ZipCode,
		&company.Address.Neighborhood,
		&company.Address.City,
		&company.Address.FederativeUnit,
		&company.SocialNetwork.InstagramURL,
		&company.SocialNetwork.FacebookURL,
		&company.SocialNetwork.LinkedinURL,
		&company.SocialNetwork.TwitterURL,
	)

	if err != nil {
		return company, err
	}

	return company, nil
}

//GetCompanies - return a list of companies
func (s *companyRepo) GetCompanies() (*[]entity.Company, *resterrors.RestErr) {

	query := queryCompanySelectBase

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0001: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetCompanyByID", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	var companies []entity.Company

	rows, err := stmt.Query()
	if err != nil {
		errorCode := "Error 0002: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in GetCompanys", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer rows.Close()

	companies, err = s.parseCompanySet(rows)
	if err != nil {
		errorCode := "Error 0003: "
		log.Println(fmt.Sprintf("%sError when trying to parse result in parseCompanySet", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &companies, nil
}

//GetCompanyByID - get a company by ID
func (s *companyRepo) GetCompanyByID(id int64) (*entity.Company, *resterrors.RestErr) {

	query := queryCompanySelectBase + `
		WHERE 	c.id 		= ?;`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0004: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetCompanyByID", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	var company entity.Company

	result := stmt.QueryRow(id)
	company, err = s.parseCompany(result)
	if err != nil {
		errorCode := "Error 0005: "
		log.Println(fmt.Sprintf("%sError when trying to execute QueryRow in GetCompanyByID", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &company, nil
}

// Create - to create a company on database
func (s *companyRepo) Create(company entity.Company) (int64, *resterrors.RestErr) {

	query := `
		INSERT INTO tab_company (
			name,
			email,
			country_code,
			area_code,
			phone_number,
			document_number,
			website,
			business_id,
			country,
			street,
			street_number,
			complement,
			zip_code,
			neighborhood,
			city,
			federative_unit,
			instagram_url,
			facebook_url,
			linkedin_url,
			twitter_url) VALUES	
		(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0006: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Create company", errorCode), err)
		return 0, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(
		company.Name,
		company.Email,
		company.CountryCode,
		company.AreaCode,
		company.PhoneNumber,
		company.DocumentNumber,
		company.Website,
		company.BusinessID,
		company.Address.Country,
		company.Address.Street,
		company.Address.Number,
		company.Address.Complement,
		company.Address.ZipCode,
		company.Address.Neighborhood,
		company.Address.City,
		company.Address.FederativeUnit,
		company.SocialNetwork.InstagramURL,
		company.SocialNetwork.FacebookURL,
		company.SocialNetwork.LinkedinURL,
		company.SocialNetwork.TwitterURL)
	if err != nil {
		errorCode := "Error 0007: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Create company", errorCode), err)
		return 0, mysqlutils.HandleMySQLError(errorCode, err)
	}

	companyID, err := insertResult.LastInsertId()
	if err != nil {
		errorCode := "Error 0008: "
		log.Println(fmt.Sprintf("%sError when trying to get LastInsertId in the Create company", errorCode), err)
		return 0, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return companyID, nil
}

// Update - to update a company on database
func (s *companyRepo) Update(company entity.Company) (*entity.Company, *resterrors.RestErr) {

	query := `
		UPDATE tab_company
			SET	name 					= ?,
					email					= ?,
					country_code	= ?,
					area_code			= ?,
					phone_number	= ?,
					business_id		= ?
			
		WHERE id	= ?;
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0009: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Update company", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		company.Name,
		company.Email,
		company.CountryCode,
		company.AreaCode,
		company.PhoneNumber,
		company.BusinessID,
		company.ID)
	if err != nil {
		errorCode := "Error 0010: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Update company", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &company, nil
}

// Delete - to delete a company on database
func (s *companyRepo) Delete(id int64) *resterrors.RestErr {

	query := `
		DELETE FROM tab_company
		WHERE 	id			= ?;
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0011: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Delete company", errorCode), err)
		return resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		errorCode := "Error 0012: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Delete company", errorCode), err)
		return mysqlutils.HandleMySQLError(errorCode, err)
	}

	return nil
}

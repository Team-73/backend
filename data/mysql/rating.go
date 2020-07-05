package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Team-73/backend/domain/entity"
	"github.com/Team-73/backend/utils/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type ratingRepo struct {
	db *sql.DB
}

// newRatingRepo returns a instance of dbrepo
func newRatingRepo(db *sql.DB) *ratingRepo {
	return &ratingRepo{
		db: db,
	}
}

// GetCompanyUserRating - to get a company user rating on database
func (s *ratingRepo) GetCompanyUserRating(companyID, userID int64) (*entity.Rating, *resterrors.RestErr) {

	var rating entity.Rating

	query := `
		SELECT 
			cr.id,
			cr.user_id,
			cr.company_id,
			cr.total_rating,
			cr.customer_service,
			cr.company_clean,
			cr.ice_beer,
			cr.good_food,
			cr.would_go_back,
			cr.created_at

		FROM tab_company_rating cr
		WHERE cr.user_id 		= ?
		AND 	cr.company_id = ?
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0004: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetUserRating", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	row := stmt.QueryRow(userID, companyID)

	err = row.Scan(
		&rating.ID,
		&rating.UserID,
		&rating.CompanyID,
		&rating.TotalRating,
		&rating.CustomerService,
		&rating.CompanyClean,
		&rating.IceBeer,
		&rating.GoodFood,
		&rating.WouldGoBack,
		&rating.CreatedAt,
	)

	if err != nil {
		errorCode := "Error 0005: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in GetUserRating", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &rating, nil
}

// CreateRating - to create a company rating on database
func (s *ratingRepo) CreateRating(rating entity.Rating) (*entity.Rating, *resterrors.RestErr) {

	query := `
		INSERT INTO tab_company_rating (
			user_id,
			company_id,
			total_rating,
			customer_service,
			company_clean,
			ice_beer,
			good_food,
			would_go_back) VALUES	
		(?, ?, ?, ?, ?, ?, ?, ?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0006: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in CreateRating", errorCode), err)
		return &rating, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		rating.UserID,
		rating.CompanyID,
		rating.TotalRating,
		rating.CustomerService,
		rating.CompanyClean,
		rating.IceBeer,
		rating.GoodFood,
		rating.WouldGoBack)
	if err != nil {
		errorCode := "Error 0007: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in CreateRating", errorCode), err)
		return &rating, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &rating, nil
}

// UpdateRating - to update a company rating on database
func (s *ratingRepo) UpdateRating(rating entity.Rating) (*entity.Rating, *resterrors.RestErr) {

	query := `
		UPDATE tab_company_rating
			SET	customer_service	= ?,
					company_clean			= ?,
					ice_beer					= ?,
					good_food					= ?,
					would_go_back			= ?,
					total_rating			= ?
			
		WHERE user_id						= ?
		AND   company_id 				= ?;
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0009: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in UpdateRating", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		rating.CustomerService,
		rating.CompanyClean,
		rating.IceBeer,
		rating.GoodFood,
		rating.WouldGoBack,
		rating.TotalRating,
		rating.UserID,
		rating.CompanyID)
	if err != nil {
		errorCode := "Error 0010: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in UpdateRating", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &rating, nil
}

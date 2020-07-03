package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Team-73/backend/domain/entity"
	"github.com/Team-73/backend/utils/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type userRepo struct {
	db *sql.DB
}

// newUserRepo returns a instance of dbrepo
func newUserRepo(db *sql.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

const querySelectBase = `
					SELECT 	u.id,
									u.name,
									u.email,
									u.password,
									u.document_number,
									u.area_code,
									u.phone_number,
									u.birthdate,
									u.gender,
									u.revenue,
									u.active

					FROM 		users 		u 
					`

func (s *userRepo) parseUserSet(rows *sql.Rows) (users []entity.User, err error) {
	for rows.Next() {
		user := entity.User{}
		user, err = s.parseUser(rows)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *userRepo) parseUser(row scanner) (user entity.User, err error) {

	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.DocumentNumber,
		&user.AreaCode,
		&user.PhoneNumber,
		&user.Birthdate,
		&user.Gender,
		&user.Revenue,
		&user.Active,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}

//GetByID - get a user by ID
func (s *userRepo) GetByID(id int64) (*entity.User, *resterrors.RestErr) {

	query := querySelectBase + `
		WHERE 	u.id 		= ?;`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0001: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetByID", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	var user entity.User

	// If we use result, err := stmt.Query(user.ID) instead QueryRow, then we need to close the connection defer result.Close() and check if we have some err
	// to get only one register on database, is better to use queryRow
	result := stmt.QueryRow(id)
	user, err = s.parseUser(result)
	if err != nil {
		errorCode := "Error 0002: "
		log.Println(fmt.Sprintf("%sError when trying to execute QueryRow in GetByID", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &user, nil
}

// Create - to create a user on database
func (s *userRepo) Create(user entity.User) (int64, *resterrors.RestErr) {

	query := `
		INSERT INTO users (
			name,
			email,
			password,
			document_number,
			area_code,
			phone_number,
			birthdate,
			gender) VALUES	
		(?, ?, ?, ?, ?, ?, ?, ?);
		`

	// When you use prepare, you not already execute the query on database, it's like to validate the query first
	// its is (more fast) than when you get an error directly on your database
	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0006: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Create user", errorCode), err)
		return 0, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.Name, user.Email, user.Password, user.DocumentNumber, user.AreaCode, user.PhoneNumber, user.Birthdate, user.Gender)
	if err != nil {
		errorCode := "Error 0007: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Create user", errorCode), err)
		return 0, mysqlutils.HandleMySQLError(errorCode, err)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		errorCode := "Error 0008: "
		log.Println(fmt.Sprintf("%sError when trying to get LastInsertId in the Create user", errorCode), err)
		return 0, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return userID, nil
}

// Update - to update a user on database
func (s *userRepo) Update(user entity.User) (*entity.User, *resterrors.RestErr) {

	query := `
		UPDATE users
			SET	name 					= ?,
					email					= ?,
					area_code			= ?,
					phone_number	= ?
			
		WHERE id	= ?;
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0009: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Update user", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.AreaCode, user.PhoneNumber, user.ID)
	if err != nil {
		errorCode := "Error 0010: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Update user", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &user, nil
}

// Delete - to delete a user on database
func (s *userRepo) Delete(id int64) *resterrors.RestErr {

	query := `
		DELETE FROM users
		WHERE 	id			= ?;
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0011: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Delete user", errorCode), err)
		return resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		errorCode := "Error 0012: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Delete user", errorCode), err)
		return mysqlutils.HandleMySQLError(errorCode, err)
	}

	return nil
}

//GetByEmailAndPassword - get a user by their email and password
func (s *userRepo) GetByEmailAndPassword(userRequest entity.LoginRequest) (*entity.User, *resterrors.RestErr) {

	query := querySelectBase + `
	
		WHERE 	u.email 		= ?
		  AND   u.password	= ?
		  AND   u.active		= 1;` //1 - active

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0013: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetByEmailAndPassword", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	// If we use result, err := stmt.Query(user.ID) instead QueryRow, then we need to close the connection defer result.Close() and check if we have some err
	// to get only one register on database, is better to use queryRow
	result := stmt.QueryRow(userRequest.Email, userRequest.Password)

	var user entity.User
	user, err = s.parseUser(result)
	if err != nil {
		errorCode := "Error 0014: "
		log.Println(fmt.Sprintf("%sError when trying to execute QueryRow in GetByDocumentNumberAndPassword", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &user, nil
}

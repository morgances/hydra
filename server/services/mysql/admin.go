package mysql

import (
	"database/sql"
	"errors"

	"github.com/fengyfei/go-box/encrypt/salt"
)

type (
	// AdminServiceImpl implements the operations on administrator.
	AdminServiceImpl struct {
		DB *sql.DB
	}
)

const (
	sqlAdminCreateTable = iota
	sqlAdminInsert
	sqlAdminLogin
	sqlAdminModifyEmail
	sqlAdminModifyMobile
	sqlAdminGetPassword
	sqlAdminModifyPassword
	sqlAdminModifyActive
	sqlAdminGetIsActive
)

var (
	errAdminLoginFailed = errors.New("invalid username or password")

	adminSqls = []string{
		`CREATE TABLE IF NOT EXISTS admin (
			id 	        BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
			name		VARCHAR(512) UNIQUE NOT NULL DEFAULT ' ',
			pwd			VARCHAR(512) NOT NULL DEFAULT ' ',
			real_name	VARCHAR(512) NOT NULL DEFAULT ' ',
			mobile		VARCHAR(32) UNIQUE DEFAULT NULL,
			email		VARCHAR(128) UNIQUE DEFAULT NULL,
			active		BOOLEAN DEFAULT TRUE,
			created_at 	DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`,
		`INSERT INTO admin(name,pwd,real_name,mobile,email,active) VALUES (?,?,?,?,?,?)`,
		`SELECT id,pwd FROM admin WHERE name = ? AND active=true LOCK IN SHARE MODE`,
		`UPDATE admin SET email=? WHERE id = ? LIMIT 1`,
		`UPDATE admin SET mobile=? WHERE id = ? LIMIT 1`,
		`SELECT pwd FROM admin WHERE id = ? AND active = true LOCK IN SHARE MODE`,
		`UPDATE admin SET pwd = ? WHERE id = ? LIMIT 1`,
		`UPDATE admin SET active = ? WHERE id = ? LIMIT 1`,
		`SELECT active FROM admin WHERE id = ? LOCK IN SHARE MODE`,
	}
)

// Initialize -
func (as *AdminServiceImpl) Initialize() error {
	_, err := as.DB.Exec(adminSqls[sqlAdminCreateTable])
	return err
}

// Create -
func (as *AdminServiceImpl) Create(name, pwd, realName, mobile, email string) error {
	hash, err := salt.Generate(&pwd)
	if err != nil {
		return err
	}

	result, err := as.DB.Exec(adminSqls[sqlAdminInsert], name, hash, realName, mobile, email, true)
	if err != nil {
		return err
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return errQueryFailed
	}

	return nil
}

// Login return user id and nil if login success.
func (as *AdminServiceImpl) Login(name, pwd *string) (uint32, error) {
	var (
		id       uint32
		password string
	)

	err := as.DB.QueryRow(adminSqls[sqlAdminLogin], name).Scan(&id, &password)
	if err != nil {
		return 0, err
	}

	if !salt.Compare([]byte(password), pwd) {
		return 0, errAdminLoginFailed
	}

	return id, nil
}

// ModifyEmail -
func (as *AdminServiceImpl) ModifyEmail(id uint32, email *string) error {
	result, err := as.DB.Exec(adminSqls[sqlAdminModifyEmail], email, id)
	if err != nil {
		return err
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return errQueryFailed
	}

	return nil
}

// ModifyMobile -
func (as *AdminServiceImpl) ModifyMobile(id uint32, mobile *string) error {
	result, err := as.DB.Exec(adminSqls[sqlAdminModifyMobile], mobile, id)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errQueryFailed
	}

	return nil
}

// ModifyPassword -
func (as *AdminServiceImpl) ModifyPassword(id uint32, pwd, newPwd *string) error {
	var (
		password string
	)
	err := as.DB.QueryRow(adminSqls[sqlAdminGetPassword], id).Scan(&password)
	if err != nil {
		return err
	}

	if !salt.Compare([]byte(password), pwd) {
		return errAdminLoginFailed
	}

	hash, err := salt.Generate(newPwd)
	if err != nil {
		return err
	}

	_, err = as.DB.Exec(adminSqls[sqlAdminModifyPassword], hash, id)

	return err
}

func (as *AdminServiceImpl) changeActive(id uint32, active bool) error {
	result, err := as.DB.Exec(adminSqls[sqlAdminModifyActive], active, id)
	if err != nil {
		return err
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return errQueryFailed
	}

	return nil

}

// Active -
func (as *AdminServiceImpl) Active(id uint32) error {
	return as.changeActive(id, true)
}

// Inactive -
func (as *AdminServiceImpl) Inactive(id uint32) error {
	return as.changeActive(id, false)
}

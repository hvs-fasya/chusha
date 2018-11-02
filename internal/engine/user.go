package engine

import (
	"github.com/hvs-fasya/chusha/internal/models"
	"github.com/hvs-fasya/chusha/internal/utils"
)

//UserCreate create new user in db
func (db *PgDB) UserCreate(user *models.UserNewInput) error {
	//todo: role?
	var e error
	user.PswdHash, e = utils.HashAndSalt([]byte(user.Password))
	if e != nil {
		return e
	}
	q := `INSERT INTO users (email, phone, nickname, name, lastname, role_id)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id`
	e = db.Conn.QueryRow(q, user.Email, user.Phone, user.Nickname, user.Name, user.LastName, user.Role.ID).Scan(
		&user.ID,
	)
	if e != nil {
		return e
	}
	return e
}

//UserCheck check username/password pair exists
func (db *PgDB) UserCheck(login string, pwd string) (*models.UserDB, error) {
	user := new(models.UserDB)
	user.Role = new(models.RoleDB)
	pwdHash, e := utils.HashAndSalt([]byte(pwd))
	if e != nil {
		return user, e
	}
	user.PswdHash = pwdHash
	q := `SELECT u.id, u.email, u.phone, u.nickname, u.name, u.lastname, u.pswd_hash,
			r.id, r.role
		FROM users u
		JOIN roles r ON r.id=u.role_id
		WHERE nickname=$1 AND pswd_hash=$2`
	err := db.Conn.QueryRow(q, login, pwdHash).Scan(
		&user.ID,
		&user.Email,
		&user.Phone,
		&user.Nickname,
		&user.Name,
		&user.LastName,
		&user.Role.ID,
		&user.Role.Role,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}

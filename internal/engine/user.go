package engine

import "github.com/hvs-fasya/chusha/internal/models"

//UserCheck check username/password pair exists
func (db *PgDB) UserCheck(login string, pwd string) (*models.User, error) {
	user := new(models.User)
	pwdHash := pwd.MakeHash()
	return user, nil
}

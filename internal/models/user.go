package models

//UserDB user db object structure
type UserDB struct {
	ID       int64   `json:"id"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Nickname string  `json:"nickname"`
	Name     string  `json:"name"`
	LastName string  `json:"last_name"`
	PswdHash string  `json:"-"`
	Role     *RoleDB `json:"role"`
}

//UserNewInput new user input structure
type UserNewInput struct {
	Password string `json:"password"`
	UserDB
}

//UserOut user data output structure
type UserOut struct {
	UserDB
	RoleDB
}

//RoleDB role db object structure
type RoleDB struct {
	ID   int64  `json:"id"`
	Role string `json:"role"`
}

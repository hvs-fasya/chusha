package models

//User user object structure
type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Role     *Role  `json:"role"`
}

//Role role object structure
type Role struct {
	ID   int64  `json:"id"`
	Role string `json:"role"`
}

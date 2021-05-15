package book_store

type User struct {
	ID       int    `json:"-"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

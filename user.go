package book_store

type User struct {
	ID       int    `json:"-" db:"id"`
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

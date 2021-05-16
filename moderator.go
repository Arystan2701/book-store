package book_store

type Moderator struct {
	ID       int    `json:"-" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password"`
	Role     string `json:"role" db:"role"`
}

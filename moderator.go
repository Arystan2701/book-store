package book_store

type Moderator struct {
	ID       int    `json:"-"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

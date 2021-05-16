package book_store

type Moderator struct {
	ID       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"-"`
	Role     string `json:"role" db:"role"`
}

type CreateModeratorInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

package users

type UserResponse struct {
	Code int      `json:"code"`
	Data UserData `json:"data"`
}

type UserData struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}

type UsersListResponse struct {
	Code int        `json:"code"`
	Meta Meta       `json:"meta"`
	Data []UserData `json:"data"`
}

type Meta struct {
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Total float64 `json:"total"`
	Pages int     `json:"pages"`
	Page  int     `json:"page"`
	Limit int     `json:"limit"`
}

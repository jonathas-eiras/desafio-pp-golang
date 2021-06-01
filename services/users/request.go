package users

type UserRequest struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}

type UserRequestUpdate struct {
	Name string `json:"name"`
}

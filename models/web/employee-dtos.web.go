package web

type RegisterEmployeeDTO struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Username string `json:"username" form:"username"`
	Role     string `json:"role" form:"role"`
}

type EmployeeBodyResponse struct {
	Email    string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
	Role     string `json:"role" form:"role"`
}

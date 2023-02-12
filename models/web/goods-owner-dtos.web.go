package web

type GoodsOwnerRegisterDTO struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
}

type GoodsOwnerResponseBody struct {
	Email string `json:"email" form:"email"`
	Name  string `json:"name" form:"name"`
}

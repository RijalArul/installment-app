package web

type LoanLimitResponseDTO struct {
	UserID      uint `json:"user_id" form:"user_id"`
	FirstMonth  int  `json:"first_month" form:"first_month"`
	SecondMonth int  `json:"second_month" form:"second_month"`
	ThirdMonth  int  `json:"third_month" form:"third_month"`
	FourthMonth int  `json:"fourth_month" form:"fourth_month"`
}

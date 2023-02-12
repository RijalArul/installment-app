package web

type InstallmentRequestDTO struct {
	OTR               int    `json:"otr" form:"otr"`
	AdminFee          int    `json:"admin_fee" form:"admin_fee"`
	AmountInstallment int    `json:"amount_installment" form:"amount_installment"`
	AmountRate        int    `json:"amount_rate" form:"amount_rate"`
	GoodName          string `json:"good_name" form:"good_name"`
	UserID            uint
	GoodID            uint
	LoanLimitID       uint
}

type InstallmentResponseDTO struct {
	OTR               int    `json:"otr" form:"otr"`
	AdminFee          int    `json:"admin_fee" form:"admin_fee"`
	AmountInstallment int    `json:"amount_installment" form:"amount_installment"`
	AmountRate        int    `json:"amount_rate" form:"amount_rate"`
	GoodName          string `json:"good_name" form:"good_name"`
	UserID            uint
	GoodID            uint
	LoanLimitID       uint
}

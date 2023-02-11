package web

import "mime/multipart"

type CheckAccountDTO struct {
	RekKoran *multipart.FileHeader `json:"rek_koran" form:"rek_koran"`
}

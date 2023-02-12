package web

import (
	"mime/multipart"
	"test-kr-sigma/models/entities"
)

type UserRegisterDTO struct {
	NIK        string                `json:"nik" form:"nik"`
	Email      string                `json:"email" form:"email"`
	Password   string                `json:"password" form:"password"`
	Salary     int                   `json:"salary" form:"salary"`
	BirthPlace string                `json:"birth_place" form:"birth_place"`
	BirthDate  string                `json:"birth_date" form:"birth_date"`
	KTP        *multipart.FileHeader `json:"ktp" form:"ktp"`
	Selfie     *multipart.FileHeader `json:"selfie" form:"selfie"`
}

type UserResponseBodyDTO struct {
	NIK          string `json:"nik" form:"nik"`
	Email        string `json:"email" form:"email"`
	Salary       int    `json:"salary" form:"salary"`
	BirthPlace   string `json:"birth_place" form:"birth_place"`
	BirthDate    string `json:"birth_date" form:"birth_date"`
	KTP          string `json:"ktp" form:"ktp"`
	Selfie       string `json:"selfie" form:"selfie"`
	CheckAccount []entities.CheckAccount
}

type LoginRequestDTO struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserLoginResponseBody struct {
	AccessToken string `json:"access_token" form:"access_token"`
}

type UpdateExpends struct {
	ExpendAverage int `json:"expend_average" form:"expend_average"`
}

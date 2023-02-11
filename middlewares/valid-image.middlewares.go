package middlewares

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	maxUploadSize = 255000
	jpeg          = "image/jpeg"
	png           = "image/png"
	jpg           = "image/jpg"
	pdf           = "application/pdf"
)

func ValidateKTP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, header, err := ctx.Request.FormFile("KTP")
		if err != nil {
			// log.Fatal(err)
			return
		}

		bytes, _ := ioutil.ReadAll(file)

		mimeType := http.DetectContentType(bytes)
		if header.Size > int64(maxUploadSize) || mimeType != jpeg && mimeType != png && mimeType != jpg && mimeType != pdf {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"err": "Upload size limited 255KB, filetype must be jpeg/png/jpg",
			})
		}

		ctx.Next()
	}
}

func ValidateSelfie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, header, err := ctx.Request.FormFile("Selfie")
		if err != nil {
			// log.Fatal(err)
			return
		}

		bytes, _ := ioutil.ReadAll(file)

		mimeType := http.DetectContentType(bytes)
		if header.Size > int64(maxUploadSize) || mimeType != jpeg && mimeType != png && mimeType != jpg && mimeType != pdf {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"err": "Upload size limited 255KB, filetype must be jpeg/png/jpg",
			})
		}

		ctx.Next()
	}
}

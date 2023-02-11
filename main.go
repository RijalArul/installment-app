package main

// import (

// )

// func main() {
// 	databases.StartDB()
// }

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"test-kr-sigma/databases"
	"time"

	"github.com/gin-gonic/gin"
)

func uploadSingleFile(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}

	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	filePath := "http://localhost:8080/images/" + filename

	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"filepath": filePath})
}

func uploadMultipleFile(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["images"]
	filePaths := []string{}
	for _, file := range files {
		fileExt := filepath.Ext(file.Filename)
		originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
		now := time.Now()
		filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
		filePath := "http://localhost:8080/images/" + filename

		filePaths = append(filePaths, filePath)
		out, err := os.Create("./public/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		readerFile, _ := file.Open()
		_, err = io.Copy(out, readerFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"filepath": filePaths})
}
func init() {
	if _, err := os.Stat("public"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll("public", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}
func main() {
	router := gin.Default()
	databases.StartDB()

	router.POST("/uploads", uploadSingleFile)
	router.StaticFS("/images", http.Dir("public"))
	router.Run()
}

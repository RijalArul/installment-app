package main

import (
	"errors"
	"log"
	"os"
	"test-kr-sigma/databases"
	"test-kr-sigma/routers"
)

func init() {
	if _, err := os.Stat("public"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll("public", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}
func main() {
	// router := gin.Default()
	databases.StartDB()
	routers.Routes()

	// router.POST("/uploads", middlewares.ValidateKTP(), uploadFiles)

	// router.Run()
}

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/seetharamu/urlShortner/controller"
	"github.com/seetharamu/urlShortner/model"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/urlshort?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&model.URL{})

	router := gin.Default()

	urlController := controller.NewURLController(db)
	router.POST("/shorten", urlController.Shorten)
	router.GET("/:shortened", urlController.Redirect)

	if err := http.ListenAndServe(":8082", router); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

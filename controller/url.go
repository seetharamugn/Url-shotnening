package controller

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/seetharamu/urlShortner/model"
)

type URLController struct {
	db *gorm.DB
}

func NewURLController(db *gorm.DB) *URLController {
	return &URLController{db: db}
}

func (u *URLController) Shorten(c *gin.Context) {
	var url model.URL
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var count int64
	u.db.Model(&model.URL{}).Count(&count)
	if count >= 20000 {
		c.JSON(http.StatusTooManyRequests, gin.H{"message": "maximum number of links reached"})
		return
	}

	shortened := u.generateShortenedURL()
	expiresAt := time.Now().Add(24 * time.Hour).Unix()

	u.db.Create(&model.URL{Original: url.Original, Shortened: "http://localhost:8082/" + shortened, ExpiresAt: expiresAt})

	c.JSON(http.StatusOK, gin.H{"shortened": "http://localhost:8082/" + shortened,
		"ExpiredAt":   expiresAt,
		"originalUrl": url.Original})
}

func isExpired(url model.URL) bool {
	return time.Now().Unix() > url.ExpiresAt
}

func (u *URLController) Redirect(c *gin.Context) {
	var url model.URL
	u.db.Where("shortened = ?", c.Param("shortened")).First(&url)

	if url.Original != "" {
		if isExpired(url) {
			u.db.Delete(&url)
			c.AbortWithStatus(http.StatusGone)
		} else {
			c.Redirect(http.StatusMovedPermanently, url.Original)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (u *URLController) generateShortenedURL() string {
	rand.Seed(time.Now().UnixNano())

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

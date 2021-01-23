package handler

import (
	"github.com/gin-gonic/gin"
	"go-gin-gorm/config"
	"go-gin-gorm/model"
	"go-gin-gorm/service"
	"log"
	"net/http"
	"strconv"
)



func GetPost(c *gin.Context) {

	query := c.Query("id")
	if len(query) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid param"})
		return
	}
	id, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid param"})
		return
	}
	post, err := service.GetPostById(id)
	if err == config.NotFoundErr {
		c.Status(http.StatusNoContent)
		return
	}
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, post)
}

func SavePost(c *gin.Context) {
	post := &model.Post{}
	err := c.Bind(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid param"})
		return
	}
	err = service.SavePost(post)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

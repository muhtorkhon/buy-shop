package controllers

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Data interface{} `json:"data"`
}

func HandleResponse(c *gin.Context, statusCode int, message interface{}) {
	c.JSON(statusCode, Response{Data: message})
}

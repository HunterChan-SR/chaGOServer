package controllers

import (
	"github.com/gin-gonic/gin"
	"os"
)

type SaveFileController struct{}

func (sfc SaveFileController) GetList(c *gin.Context) {
	files, _ := os.ReadDir("saveFiles")
	var data []string
	for _, file := range files {
		data = append(data, file.Name())
	}
	ReturnSuccess(c, OK, "success", data, len(data))
}
func (sfc SaveFileController) Get(c *gin.Context) {
	param := c.Param("filename")
	filename := "saveFiles/" + param
	c.File(filename)
}

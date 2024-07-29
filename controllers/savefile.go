package controllers

import (
	"chag/bean"
	"chag/db"
	"github.com/gin-gonic/gin"
)

type SavedFileController struct{}

func (sfc SavedFileController) GetList(c *gin.Context) {
	var savedfiles []bean.SavedFile
	db.DB.Find(&savedfiles)
	ReturnSuccess(c, OK, "success", savedfiles, len(savedfiles))
}
func (sfc SavedFileController) Get(c *gin.Context) {
	param := c.Param("filename")
	fileName := "savedFiles/" + param
	c.File(fileName)
}

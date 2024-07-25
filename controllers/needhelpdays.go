package controllers

import (
	"chag/bean"
	"chag/db"
	"github.com/gin-gonic/gin"
)

type NeedHelpDaysController struct{}

func (nhd NeedHelpDaysController) GetList(c *gin.Context) {
	var needhelpdays []bean.NeedHelpDays
	db.DB.Find(&needhelpdays)
	ReturnSuccess(c, OK, "success", needhelpdays, len(needhelpdays))
}

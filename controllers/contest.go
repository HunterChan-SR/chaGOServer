package controllers

import (
	"chag/bean"
	"chag/db"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ContestController struct{}

func (co ContestController) GetList(c *gin.Context) {
	var contests []bean.Contest
	db.DB.Find(&contests)
	ReturnSuccess(c, OK, "success", contests, len(contests))
}
func (co ContestController) Get(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	var problems []bean.Problem
	db.DB.Where("contestid = ?", id).Order("createtime").Find(&problems)
	ReturnSuccess(c, OK, "success", problems, len(problems))
}

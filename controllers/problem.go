package controllers

import (
	"chag/bean"
	"chag/db"
	"chag/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ProblemController struct{}

func (p ProblemController) Get(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	problem := bean.Problem{Id: id}
	db.DB.Find(&problem)
	fileName := "files/" + param + ".html"

	c.File(fileName)
}
func (p ProblemController) Post(c *gin.Context) {
	//获取表单form-data
	var problem bean.Problem
	_ = c.Bind(&problem)
	util.GetP(strconv.Itoa(problem.Id))
	util.Tran(strconv.Itoa(problem.Id))
	db.DB.Create(&problem)
	ReturnSuccess(c, OK, "success", nil, 0)
}

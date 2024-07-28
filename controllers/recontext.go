package controllers

import (
	"chag/bean"
	"chag/db"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type RecontextController struct{}

func (rc RecontextController) GetList(c *gin.Context) {
	param := c.Param("needhelpid")
	needhelpid, _ := strconv.Atoi(param)
	var recontexts []bean.RecontextView
	db.DB.Where("needhelpid = ?", needhelpid).Find(&recontexts)
	ReturnSuccess(c, OK, "success", recontexts, len(recontexts))
}

func (rc RecontextController) Post(c *gin.Context) {
	var recontext bean.Recontext
	_ = c.Bind(&recontext)
	currentTime := time.Now()
	recontext.Createtime = currentTime.Format("2006-01-02 15:04:05")
	db.DB.Create(&recontext)
	ReturnSuccess(c, OK, "success", "success", 0)
}

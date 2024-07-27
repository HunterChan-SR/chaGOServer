package controllers

import (
	"chag/bean"
	"chag/db"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type NeedHelpController struct{}

//create view needhelpview as
//select needhelp.id,
//       needhelp.days,
//       user.nickname,
//       needhelp.problemtitle,
//       needhelp.subcode,
//       needhelp.context,
//       needhelp.recontext,
//       needhelp.createtime
//from needhelp,
//     user
//where needhelp.userid = user.id;

type NeedHelpGetList struct {
	Id           int    `json:"id"`
	Days         int    `json:"days"`
	Nickname     string `json:"nickname"`
	Problemtitle string `json:"problemtitle"`
	Subcode      string `json:"subcode"`
	Context      string `json:"context"`
	Createtime   string `json:"createtime"`
}

func (nh NeedHelpController) GetList(c *gin.Context) {
	param := c.Param("days")
	days, _ := strconv.Atoi(param)
	var needhelps []NeedHelpGetList
	db.DB.Table("needhelpview").Where("days = ?", days).Find(&needhelps)
	ReturnSuccess(c, OK, "success", needhelps, len(needhelps))
}

func (nh NeedHelpController) PostContext(c *gin.Context) {
	var needhelp bean.NeedHelp
	_ = c.Bind(&needhelp)
	//fmt.Println(needhelp)
	//获取当前时间
	currentTime := time.Now()
	needhelp.Createtime = currentTime.Format("2006-01-02 15:04:05")
	//插入
	db.DB.Create(&needhelp)
	ReturnSuccess(c, OK, "success", "success", 0)
}

//func (nh NeedHelpController) PostRecontext(c *gin.Context) {
//	var needhelp bean.NeedHelp
//	_ = c.Bind(&needhelp)
//	//fmt.Println("##########")
//	//fmt.Println(needhelp.Id)
//	//更新
//	db.DB.Model(&needhelp).Update("recontext", needhelp.Recontext)
//	ReturnSuccess(c, OK, "success", "success", 0)
//}

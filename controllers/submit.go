package controllers

import (
	"chag/bean"
	"chag/db"
	"chag/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type SubmitController struct{}

var subFlag bool

func init() {
	subFlag = false
}

func (s SubmitController) Get(c *gin.Context) {
	paramUserid := c.Param("userid")
	paramProblemid := c.Param("problemid")

	UserId, _ := strconv.Atoi(paramUserid)
	ProblemId, _ := strconv.Atoi(paramProblemid)

	var submits []bean.Submit
	db.DB.Where("userid = ? and problemid = ?", UserId, ProblemId).Find(&submits)
	ReturnSuccess(c, OK, "success", submits, len(submits))
}

func (s SubmitController) Post(c *gin.Context) {
	type subCode struct {
		Code      string
		Userid    string
		Problemid string
	}
	var subcode subCode
	_ = c.Bind(&subcode)
	//fmt.Println(subcode.Code, "\n", subcode.Userid, "\n", subcode.Problemid)

	submit := bean.Submit{}
	submit.Userid, _ = strconv.Atoi(subcode.Userid)
	submit.Problemid, _ = strconv.Atoi(subcode.Problemid)

	var submit2 []bean.Submit

	haveRating := true
	wrongCnt := 0
	db.DB.Where("userid = ? and problemid = ?", submit.Userid, submit.Problemid).Find(&submit2)
	for _, v := range submit2 {
		if strings.Contains(v.State, "通过") {
			haveRating = false
		} else {
			wrongCnt++
		}
	}

	//循环等待subFlag 变成false
	for subFlag {
		//等待20毫秒
		time.Sleep(20 * time.Millisecond)
	}
	subFlag = true
	submit.Dfbyid = util.GetId(subcode.Problemid, subcode.Code)
	subFlag = false

	getRes := util.GetResult(submit.Dfbyid)
	for strings.Contains(getRes, "正在") || strings.Contains(getRes, "等待") {
		time.Sleep(5 * time.Millisecond)
		getRes = util.GetResult(submit.Dfbyid)
	}
	submit.State = getRes

	if !haveRating {
		ReturnSuccess(c, OK, "success", getRes+"+0", 1)
		return
	} else {
		db.DB.Create(&submit)
		if strings.Contains(getRes, "通过") {
			db.DB.Model(&bean.User{}).Where("id = ?", submit.Userid).Update("rating", gorm.Expr("rating + ?", 100))
			ReturnSuccess(c, OK, "success", getRes+"+100", 1)
			return
		} else {
			db.DB.Model(&bean.User{}).Where("id = ?", submit.Userid).Update("rating", gorm.Expr("rating - ?", 33))
			if wrongCnt >= 3 {
				//设置ranking为7
				db.DB.Model(&bean.User{}).Where("id = ?", submit.Userid).Update("ranking", 7)
			}
			ReturnSuccess(c, OK, "success", getRes+"-33", 1)
			return
		}
	}

}

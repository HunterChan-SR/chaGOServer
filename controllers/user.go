package controllers

import (
	"chag/bean"
	"chag/db"
	"chag/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type UserController struct{}

func (u UserController) Get(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)
	use := bean.User{Id: id}
	db.DB.Find(&use)
	use.Password = ""
	ReturnSuccess(c, OK, "success", use, 1)
}

func (u UserController) GetList(c *gin.Context) {
	var userList []bean.User
	//先按照ranking升序 再按照rating降序
	db.DB.Order("ranking asc").Order("rating desc").Find(&userList)
	for i := 0; i < len(userList); i++ {
		userList[i].Password = ""
	}
	ReturnSuccess(c, OK, "success", userList, len(userList))
}

func (u UserController) PostRanking(c *gin.Context) {
	var userList []bean.User
	db.DB.Find(&userList)
	sum := 0
	cnt := len(userList)
	for i := 0; i < cnt; i++ {
		if strings.Contains(userList[i].Username, "admin") {
			continue
		}
		sum += userList[i].Rating
	}
	avg := float64(sum) / float64(cnt-1)
	for i := 0; i < cnt; i++ {
		if strings.Contains(userList[i].Username, "admin") {
			db.DB.Model(&bean.User{}).Where("id = ?", userList[i].Id).Update("rating", int(avg)).Update("ranking", 4)
			continue
		}
		rating := float64(userList[i].Rating)
		if rating >= avg*2.0 {
			db.DB.Model(&bean.User{}).Where("id = ?", userList[i].Id).Update("ranking", 1)
		} else if rating >= avg*1.6 {
			db.DB.Model(&bean.User{}).Where("id = ?", userList[i].Id).Update("ranking", 2)
		} else if rating >= avg*1.2 {
			db.DB.Model(&bean.User{}).Where("id = ?", userList[i].Id).Update("ranking", 3)
		} else if rating >= avg*0.8 {
			db.DB.Model(&bean.User{}).Where("id = ?", userList[i].Id).Update("ranking", 4)
		} else if rating >= avg*0.4 {
			db.DB.Model(&bean.User{}).Where("id = ?", userList[i].Id).Update("ranking", 5)
		} else {
			db.DB.Model(&bean.User{}).Where("id = ?", userList[i].Id).Update("ranking", 6)
		}
	}
	ReturnSuccess(c, OK, "success", "success", 1)
}

type loginUser struct {
	Token  string `json:"token"`
	Userid int    `json:"userid"`
}

func (u UserController) PostLogin(c *gin.Context) {
	//获取表达 Username Password
	var user bean.User
	_ = c.Bind(&user)
	if len(user.Username) == 0 || len(user.Password) == 0 {
		ReturnError(c, ERROR, "用户名或密码不能为空")
		return
	}
	if db.DB.Where("username = ?", user.Username).Where("password = ?", user.Password).First(&user).RowsAffected == 0 {
		ReturnError(c, ERROR, "用户名或密码错误")
		return
	}
	token := strings.Join([]string{user.Username, user.Password}, ":")
	token = util.Encrypt(token)
	//设置多个cookie myidentify 和 userid 任意domain
	//c.SetCookie("myidentify", myidentify, 3600*24*3, "/", "", false, true)
	//c.SetCookie("userid", strconv.Itoa(user.Id), 3600*24*3, "/", "", false, true)
	loginuser := loginUser{Token: token, Userid: user.Id}
	ReturnSuccess(c, OK, "success", loginuser, 1)
}

type modifyPassword struct {
	Id          string `json:"id"`
	Password    string `json:"password"`
	Newpassword string `json:"newpassword"`
}

func (u UserController) ModifyPassword(c *gin.Context) {

	var modifyUser modifyPassword
	err := c.Bind(&modifyUser)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("\nmodifyUser.Id", modifyUser.Id, "\nmodifyUser.Password", modifyUser.Password, "\nmodifyUser.Newpassword", modifyUser.Newpassword)
	//
	//fmt.Println("###DB")
	id, _ := strconv.Atoi(modifyUser.Id)
	user := bean.User{
		Id:       id,
		Password: modifyUser.Password,
	}
	cnt := db.DB.Where("id = ?", id).Where("password = ?", modifyUser.Password).Find(&user).RowsAffected
	//fmt.Println(cnt)
	if cnt == 0 {
		//fmt.Println("###RETURN")
		ReturnError(c, ERROR, "原密码错误")
		return
	}
	db.DB.Model(&bean.User{}).Where("id = ?", modifyUser.Id).Update("password", modifyUser.Newpassword)
	//c.SetCookie("myidentify", "", -1, "/", "", false, true)
	//c.SetCookie("userid", "", -1, "/", "", false, true)
	ReturnSuccess(c, OK, "success", "请重新登录", 1)
}

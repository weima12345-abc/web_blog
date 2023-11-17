package Http

import (
	"webBlog/app/Helpers"
	"webBlog/app/Service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LoginPost_xiazhaoxuan struct {
	UserName string `form:"username" binding:"required"`
	PassWord string `form:"password" binding:"required"`
}

type RegisterPost_xiazhaoxuan struct {
	UserName string `form:"username" binding:"required"`
	NickName string `form:"nickname" binding:"required"`
	PassWord string `form:"password" binding:"required"`
}

// 用户中心
func Users(c *gin.Context) {
	assign := AssignData(c, "会员中心")
	userId, _ := strconv.Atoi(c.Param("uid"))
	page, _ := strconv.Atoi(c.Query("page"))
	if userId <= 0 {
		c.HTML(http.StatusBadRequest, "error/400", nil)
		return
	}
	// 查询用户信息
	user, err := Service.GetUserInfo(userId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error/500", nil)
		return
	}
	data := Service.HandleArticleList(userId, page)
	assign["paginator"] = data.Paginator
	assign["list"] = data.ArticleList
	assign["user"] = user
	c.HTML(http.StatusOK, "users/index", assign)
}

func Edit(c *gin.Context) {
	assign := AssignData(c, "编辑个人资料")
	c.HTML(http.StatusOK, "users/edit", assign)
}

// 用户登录
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login/login", AssignData(c, "会员登录"))
}

// 用户注册
func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "login/register", AssignData(c, "会员注册"))
}

// 退出登录
func Logout(c *gin.Context) {
	Helpers.ClearSession(c)
	c.Redirect(http.StatusMovedPermanently, "/")
}

// 用户登录动作
func LoginAction(c *gin.Context) {
	var login LoginPost_xiazhaoxuan
	assign := AssignData(c, "会员登录")
	if err := c.ShouldBind(&login); err != nil {
		assign["errorMsg"] = err
		c.HTML(http.StatusBadRequest, "login/login", assign)
		return
	}
	UserId, err := Service.Login(login.UserName, Helpers.Md5(login.PassWord))
	if err != nil {
		assign["errorMsg"] = err
		c.HTML(http.StatusBadRequest, "login/login", assign)
		return
	}
	Helpers.SetSession(c, "UserId", UserId)
	c.Redirect(http.StatusMovedPermanently, "/")
}

// 用户注册动作
func RegisterAction(c *gin.Context) {
	var register RegisterPost_xiazhaoxuan
	assign := AssignData(c, "会员注册")
	if err := c.ShouldBind(&register); err != nil {
		assign["errorMsg"] = err
		c.HTML(http.StatusBadRequest, "login/register", assign)
		return
	}
	err := Service.Register(register.UserName, Helpers.Md5(register.PassWord), register.NickName)
	if err != nil {
		assign["errorMsg"] = err
		c.HTML(http.StatusOK, "login/register", assign)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}

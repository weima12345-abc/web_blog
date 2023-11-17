package Http

import (
	"webBlog/app/Service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	articles, err := Service.GetIndexShow()
	if err != nil {
		c.HTML(http.StatusOK, "error/500", nil)
		return
	}
	assign := AssignData(c, "首页")
	assign["article"] = articles
	c.HTML(http.StatusOK, "index/index", assign)
}

func Time(c *gin.Context) {
	assign := AssignData(c, "时间线")
	c.HTML(http.StatusOK, "index/time", assign)
}

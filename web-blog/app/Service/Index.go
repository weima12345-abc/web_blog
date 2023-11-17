package Service

import (
	"webBlog/app/Models"
)

func GetIndexShow() (list []*Models.Article_xiazhaoxuan, err error) {
	list, err = Models.GetTopArticleList(6)
	if err != nil {
		return
	}
	return
}
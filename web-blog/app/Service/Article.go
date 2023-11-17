package Service

import (
	"webBlog/app/Helpers"
	"webBlog/app/Models"
	"github.com/jinzhu/gorm"
)

// 组装返回数据，有点冗余，待功力提升回头优化
// 循环组装文章以及所属分类可使用sql 联查优化
type ReturnData_xiazhaoxuan struct {
	ArticleTopList []*Models.Article_xiazhaoxuan
	ArticleList    []*Models.ArticleRecord_xiazhaoxuan
	Paginator      map[string]interface{}
}

// 组合文章列表、分类列表带分页
func HandleArticleList(UserId, Page int) (data *ReturnData_xiazhaoxuan) {
	var list []*Models.ArticleRecord_xiazhaoxuan
	topList, err := Models.GetTopArticleList(10)
	if err != nil {
		return
	}
	articleList, total, err := Models.GetArticleList(UserId, Page)
	if err != nil {
		return
	}
	// 组合文章与分类信息
	for _, article := range articleList {
		// 将文章信息写入
		recordInfo := &Models.ArticleRecord_xiazhaoxuan{
			Article_xiazhaoxuan: *article,
		}
		list = append(list, recordInfo)
	}
	paginator := Helpers.Paginator(Page, Models.PageSize, total)
	return &ReturnData_xiazhaoxuan{
		ArticleTopList: topList,
		ArticleList:    list,
		Paginator:      paginator,
	}
}

// 查询单个文章信息
func GetArticleDetail(id int) (details *Models.ArticleShow_xiazhaoxuan, err error) {
	article := &Models.Article_xiazhaoxuan{
		Model: gorm.Model{ID: uint(id)},
	}
	err = article.ShowArticleInfo()
	if err != nil {
		return
	} // 文章信息
	uinfo, err := GetUserBaseInfo(int(article.UserId))
	if err != nil {
		return
	} // 用户信息
	err = article.AddArticleView(id)
	if err != nil {
		return
	} // 浏览量+1
	return &Models.ArticleShow_xiazhaoxuan{
		Article_xiazhaoxuan:  *article,
		UserInfo_xiazhaoxuan: *uinfo,
	}, nil // 组装返回信息
}

func CreateArticle(Title, Content string, UserId uint) error {
	article := Models.Article_xiazhaoxuan{
		Title:   Title,
		UserId:  UserId,
		Content: Content,
	}
	if err := article.CreateBlog(); err != nil {
		return err
	}
	return nil
}

func CheckArticleAuthor(ArticleID, UserId uint) bool {
	article := &Models.Article_xiazhaoxuan{
		Model: gorm.Model{ID: ArticleID},
	}
	err := article.ShowArticleInfo()
	if err != nil {
		return false
	}
	if article.UserId != UserId {
		return false
	}
	return true
}

func DeleteArticleById(ArticleID uint) bool {
	article := &Models.Article_xiazhaoxuan{
		Model: gorm.Model{ID: ArticleID},
	}
	err := article.DeleteArticle()
	if err != nil {
		return false
	}
	return true
}

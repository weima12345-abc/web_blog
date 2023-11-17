package main

import (
	"webBlog/app/Models"
	"webBlog/routes"
)

func main() {
	blog := routes.InitWebRouter()
	defer Models.DB.Close()
	_ = blog.Run(":8888")
}

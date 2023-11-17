# web_blog
web博客：基于Go, Gin, Mysql的个人博客

启动方式

1.将项目文件夹放在%GOPATH%/src目录下

2.编辑env文件中的DB_USERNAME用户名以及DB_PASSWORD密码，
和DB_PORT端口（默认为3306），连接本地数据库；

3.启动Mysql，然后执行CREATE DATABASE `webblogdatabase`; 

即可在MySQL中创建名为webblogdatabase的数据库

4.将sql文件导入该数据库（这一步是可选的，程序会自动创造表）

5.如果import中多处报红，需要在终端下执行go env -w GO111MODULE=auto，

然后执行go mod vendor下载所需依赖包

6.启动webBlog.exe可执行文件

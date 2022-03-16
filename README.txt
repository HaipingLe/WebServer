1. 该Demo主要模拟了web服务器和相应的客户端，服务器和客户端都加载了证书，可实现双向认证.
   1.1> MyServer.go   web服务器，采用了Gin作为路由框架
   1.2> client文件夹下的client_get.go, client_put.go，client_post.go和client_del.go分别模拟了HTTP的GET,PUT,POST和DELETE.
   1.3> sqlUtil文件夾下的sqlHelper.go 封装了MySql数据库的基本操作

2. 该Demo需要下载并导入Gin和MySql包
   go get github.com/gin-gonic/gin
   go get github.com/go-sql-driver/mysql
   下载后的包存放在$GOPATH/src目录下
   如果想学习Golang，建议搭建环境、了解几个重要的环境变量，比如GOROOT、GOPATH. 推荐《Go语言编程》,我已将其放在该仓库的根目录下.
## 在 Go 程序中使用外部库
下载并安装 Go 客户端库:
将通过 go install 实现。但是首先要验证环境变量中是否含有 `GOPATH` 变量，因为外部源码将被下载到 `$GOPATH/src` 目录下并被安装到 `$GOPATH/PKG/"machine_arch"/` 目录下。

我们将通过在终端调用以下命令来下载 API:

	go get github.com/go-sql-driver/mysql

我们将通过在终端调用以下命令来安装 API:

	go install github.com/go-sql-driver/mysql

go install 将下载源码，编译并安装包

使用 urlshortener 服务的 web 程序:
现在我们可以通过导入并赋予别名来使用已安装的包：

	import  "github.com/go-sql-driver/mysql"


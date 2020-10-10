## 在 Go 程序中使用外部库
下载并安装 Go 客户端库:
将通过 go install 实现。但是首先要验证环境变量中是否含有 `GOPATH` 变量，因为外部源码将被下载到 `$GOPATH/src` 目录下并被安装到 `$GOPATH/PKG/"machine_arch"/` 目录下。

我们将通过在终端调用以下命令来下载 API:

	go get google.golang.org/api/urlshortener/v1

我们将通过在终端调用以下命令来安装 API:

	go install google.golang.org/api/urlshortener/v1

go install 将下载源码，编译并安装包

使用 urlshortener 服务的 web 程序:
现在我们可以通过导入并赋予别名来使用已安装的包：

	import  "google.golang.org/api/urlshortener/v1"


要调用 urlshortener 接口必须先通过 http 包中的默认客户端创建一个服务实例 urlshortenerSvc：  
```go
urlshortenerSvc, _ := urlshortener.New(http.DefaultClient)
```

我们通过调用服务中的 `Url.Insert` 中的 `Do` 方法传入包含长地址的 `Url` 数据结构从而获取短地址：

```go
url, _ := urlshortenerSvc.Url.Insert(&urlshortener.Url{LongUrl: longUrl}).Do()
```

返回 `url` 的 `Id` 便是我们需要的短地址。

我们通过调用服务中的 `Url.Get` 中的 `Do` 方法传入包含短地址的Url数据结构从而获取长地址：

```go
url, error := urlshortenerSvc.Url.Get(shwortUrl).Do()
```

返回的长地址便是转换前的原始地址。


示例

```go
package main

import (
	 "fmt"
	 "net/http"
	 "text/template"

	 "google.golang.org/api/urlshortener/v1"
)
func main() {
	 http.HandleFunc("/", root)
	 http.HandleFunc("/short", short)
	 http.HandleFunc("/long", long)

	 http.ListenAndServe("localhost:8080", nil)
}
// the template used to show the forms and the results web page to the user
var rootHtmlTmpl = template.Must(template.New("rootHtml").Parse(`
<html><body>
<h1>URL SHORTENER</h1>
{{if .}}{{.}}<br /><br />{{end}}
<form action="/short" type="POST">
Shorten this: <input type="text" name="longUrl" />
<input type="submit" value="Give me the short URL" />
</form>
<br />
<form action="/long" type="POST">
Expand this: http://goo.gl/<input type="text" name="shortUrl" />
<input type="submit" value="Give me the long URL" />
</form>
</body></html>
`))
func root(w http.ResponseWriter, r *http.Request) {
	rootHtmlTmpl.Execute(w, nil)
}
func short(w http.ResponseWriter, r *http.Request) {
	 longUrl := r.FormValue("longUrl")
	 urlshortenerSvc, _ := urlshortener.New(http.DefaultClient)
	 url, _ := urlshortenerSvc.Url.Insert(&urlshortener.Url{LongUrl:
	 longUrl,}).Do()
	 rootHtmlTmpl.Execute(w, fmt.Sprintf("Shortened version of %s is : %s",
	 longUrl, url.Id))
}

func long(w http.ResponseWriter, r *http.Request) {
	 shortUrl := "http://goo.gl/" + r.FormValue("shortUrl")
	 urlshortenerSvc, _ := urlshortener.New(http.DefaultClient)
	 url, err := urlshortenerSvc.Url.Get(shortUrl).Do()
	 if err != nil {
		 fmt.Println("error: %v", err)
		 return

	 }
	 rootHtmlTmpl.Execute(w, fmt.Sprintf("Longer version of %s is : %s",
	 shortUrl, url.LongUrl))
}
```
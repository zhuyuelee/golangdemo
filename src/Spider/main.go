package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
)

func main() {
	var start, end int
	fmt.Print("请输入开始页码：")
	fmt.Scanln(&start)
	fmt.Print("请输入结束页码：")
	fmt.Scanln(&end)
	task := make(chan int, end-start+1)
	for i := start; i <= end; i++ {
		go spider(i, task)
	}

	for i := start; i <= end; i++ {
		page := <-task
		fmt.Printf("第%d页采集结束\n", page)
	}
}

//https://www.pengfue.com/index_1.html
func spider(page int, task chan<- int) {
	fmt.Printf("第%d页开始采集\n", page)
	body, err := httpGet(fmt.Sprintf("https://www.pengfue.com/index_%d.html", page))
	if err != nil {
		fmt.Printf("%d 采集出错 err=%v\n", page, err)
	}
	f, _ := os.Create(fmt.Sprintf("%d.html", page))
	jokes := getJokeStruct(body)
	for _, joke := range jokes {
		f.WriteString(fmt.Sprintf("title:%s content:%s \n", joke.title, joke.content))

	}
	defer f.Close()
	task <- page
}

func httpGet(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s 采集出错 err=%v\n", url, err)
	}
	defer res.Body.Close()
	var body string

	buf := make([]byte, 1024*4)
	for {
		n, _ := res.Body.Read(buf)
		if n == 0 {
			break
		}
		body += string(buf[:n])
	}

	return body, err
}

// getJoke 批量采集
func getJoke(body string) {
	repxTitle := `<h1 class="dp-b"><a href="https://www.pengfue.com/content_\d+_1.html" target="_blank">(.*?)</a>`
	repxContent := `<div class="content-img clearfix pt10 relative">\s*(.*?)\s*</div>`

	regTitle := regexp.MustCompile(repxTitle)
	regContent := regexp.MustCompile(repxContent)

	resTitles := regTitle.FindAllStringSubmatch(body, -1)
	resContents := regContent.FindAllStringSubmatch(body, -1)

	for _, result := range resTitles {
		fmt.Println(result[1])
	}
	for _, result := range resContents {
		fmt.Println(result[1])
	}
}

// getJokeStruct 采集成结构体
func getJokeStruct(body string) []joke {
	jokes := make([]joke, 10)
	regJoks := regexp.MustCompile(`<h1 class="dp-b">(?s:(.*?))</dd>`)
	resJoks := regJoks.FindAllStringSubmatch(body, -1)

	repxTitle := `<a href="https://www.pengfue.com/content_\d+_1.html" target="_blank">(.*?)</a>`
	repxContent := `<div class="content-img clearfix pt10 relative">\s*(.*?)\s*</div>`

	regTitle := regexp.MustCompile(repxTitle)
	regContent := regexp.MustCompile(repxContent)

	index := 0
	for _, result := range resJoks {
		var j joke
		resTitle := regTitle.FindAllStringSubmatch(result[1], 1)
		if len(resTitle) == 1 {
			j.title = resTitle[0][1]
		}
		resContent := regContent.FindAllStringSubmatch(result[1], 1)
		if len(resContent) == 1 {
			j.content = resContent[0][1]
		}
		jokes[index] = j
		index++
	}
	return jokes
}

type joke struct {
	title, content string
}

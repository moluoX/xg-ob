package smzdm

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/moluoX/xg-ob/dataaccess"
	"github.com/moluoX/xg-ob/model"
)

//Crawl auto
func Crawl() {
	for {
		crawlOnce()
	}
}

func crawlOnce() {
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		crawlPage(i)
	}
}

func crawlPage(page int) {
	res, err := http.Get(fmt.Sprintf("https://www.smzdm.com/jingxuan/json_more?filter=s0f0t0b0d0r0p%d", page))
	handleErr(err)
	go analyzePage(res.Body)
}

func analyzePage(body io.ReadCloser) {
	jsonBody, err := ioutil.ReadAll(body)
	body.Close()
	handleErr(err)
	var page model.SmzdmArticlePage
	handleErr(json.Unmarshal(jsonBody, &page))
	for _, m := range page.ArticleList {
		analyzeArticle(m)
	}
}

func analyzeArticle(m model.SmzdmArticle) {
	if m.Worthy < 32 {
		return
	}
	if m.Unworthy > 0 && m.Worthy/m.Unworthy < 4 {
		return
	}
	m.Time = time.Unix(m.TimeSort, 0)
	dataaccess.SaveArticle(m)
}

func handleErr(err error) {
	if err != nil && err.Error() != "EOF" {
		log.Printf("[crawl smzdm error] %v\n", err)
	}
}

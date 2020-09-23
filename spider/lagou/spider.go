package lagou

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/moluoX/xg-ob/xlog"
)

//Crawl lagou
func Crawl() {
	crawlOnce()
}

func crawlOnce() {
	//var client http.Client
	res, err := http.PostForm("https://www.lagou.com/jobs/positionAjax.json?city=%E6%B7%B1%E5%9C%B3&needAddtionalResult", url.Values{"first": {"true"}, "pn": {"1"}, "kd": {".net"}})
	handleErr(err)
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	handleErr(err)
	fmt.Printf("%s", robots)
}

func handleErr(err error) {
	if err != nil && err.Error() != "EOF" {
		xlog.SugarLogger.Errorf("[crawl lagou error] %v\n", err)
	}
}

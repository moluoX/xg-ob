package main

import (
	"github.com/moluoX/xg-ob/spider/smzdm"
	"github.com/moluoX/xg-ob/web"
)

func main() {
	go smzdm.Crawl()
	web.Run()
}

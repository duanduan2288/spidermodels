package main

import (
	"time"

	"github.com/duanduan2288/spidermodels/spider"
)

func main() {
	go spider.SpiderDou()
	time.Sleep(time.Second * 20)
}

package main

import (
	"goCrawler/engine"
	"goCrawler/parser/zhenai"
)

func main() {
	engine.Run(engine.Request{
		Url:       "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParseFunc: zhenai.ParseCityList,
	})
}

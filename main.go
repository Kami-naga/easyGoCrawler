package main

import (
	"goCrawler/engine"
	"goCrawler/parser/zhenai"
	"goCrawler/scheduler"
)

func main() {
	//engine.SingleEngine{}.Run(engine.Request{
	//	Url:       "http://localhost:8080/mock/www.zhenai.com/zhenghun",
	//	ParseFunc: zhenai.ParseCityList,
	//})
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url:       "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParseFunc: zhenai.ParseCityList,
	})
}

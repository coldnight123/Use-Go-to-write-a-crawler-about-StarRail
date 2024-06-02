package main

import (
	"project/railroad/engine"
	"project/railroad/persist"
	"project/railroad/rail/parser"
	"project/railroad/scheduler"
)

const url = "https://api-static.mihoyo.com/common/blackboard/sr_wiki/v1/home/content/list?app_sn=sr_wiki&channel_id=17"

func main() {
	//engine.concurrent
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{}, //QueuedScheduler为方案Ⅲ实现,SimpleScheduler为方案Ⅱ实现
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseFigureList,
	})
	/*//engine.simple
	engine.SimpleEngine{}.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseFigureList},
	)
	*/
}

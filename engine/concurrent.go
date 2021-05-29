package engine

import (
	"goCrawler/model"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request

	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDup(r.Url) {
			log.Printf("Duplicate request: %s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	profileCnt := 0
	for {
		result := <-out
		for _, item := range result.Items {
			if _, ok := item.(model.Profile); ok {
				log.Printf("Got item # %d: %v", profileCnt, item)
				profileCnt++
			}
		}

		// url dedup
		for _, r := range result.Requests {
			if isDup(r.Url) {
				continue
			}
			e.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, notifier ReadyNotifier) {
	go func() {
		for {
			notifier.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

// some issues: 1.it will take too much memory to save those long urls
// how to solve it? make urls short -> md5, md5 encoding is too slow? -> bloom filter!
// memory is still not enough? maybe you need some databases like redis
// another issue: all the map will be lost when system shut down, how to solve it?
// save the map into files before shut down the system, or just save it every some time
// or just save it into some external databases like redis
var visitedUrls = make(map[string]bool)

func isDup(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}

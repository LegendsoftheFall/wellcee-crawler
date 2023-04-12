package engine

import "crawler/model"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChannel chan model.ApartmentDetail
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	InitWorkerChannel() chan Request
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	// 创建worker
	for i := 0; i < e.WorkerCount; i++ {
		// queue or simple
		createWorker(e.Scheduler.InitWorkerChannel(), out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			// 存储数据
			go func() {
				e.ItemChannel <- item
			}()
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request) // 新的request加入Scheduler
		}
	}
}

//共用一个Scheduler

func createWorker(in chan Request, out chan ParseResult, r ReadyNotifier) {
	go func() {
		for {
			// 告诉scheduler自己是空闲的
			r.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result // 没有人收这个out就会阻塞在这
		}
	}()
}

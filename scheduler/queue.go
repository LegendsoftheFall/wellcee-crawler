package scheduler

import "crawler/engine"

type QueueScheduler struct {
	RequestChannel chan engine.Request      // 收发request
	WorkerChannel  chan chan engine.Request // queue通过channel调度独立的worker,每个worker有对应的channel
}

func (q *QueueScheduler) InitWorkerChannel() chan engine.Request {
	// 对于队列 每个worker有属于自己的channel
	return make(chan engine.Request)
}

func (q *QueueScheduler) Submit(request engine.Request) {
	q.RequestChannel <- request
}

func (q *QueueScheduler) WorkerReady(w chan engine.Request) {
	q.WorkerChannel <- w
}

func (q *QueueScheduler) Run() {
	// 创建worker channel 和 request channel
	q.RequestChannel = make(chan engine.Request)
	q.WorkerChannel = make(chan chan engine.Request)
	go func() {
		var reqQueue []engine.Request
		var wrkQueue []chan engine.Request
		for {
			// 拿到合法的request和worker
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(reqQueue) > 0 && len(wrkQueue) > 0 {
				activeRequest = reqQueue[0]
				activeWorker = wrkQueue[0]
			}
			select {
			case req := <-q.RequestChannel:
				// 收到request加入队列
				reqQueue = append(reqQueue, req)
			case idleWorker := <-q.WorkerChannel:
				// 收到一个空闲的worker加入队列
				wrkQueue = append(wrkQueue, idleWorker)
			case activeWorker <- activeRequest:
				// 触发发送事件、从队列中取出
				reqQueue = reqQueue[1:]
				wrkQueue = wrkQueue[1:]
			}

		}
	}()
}

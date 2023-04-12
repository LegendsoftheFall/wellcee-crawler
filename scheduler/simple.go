package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	WorkerChannel chan engine.Request // 所有worker共用一个channel
}

func (s *SimpleScheduler) WorkerReady(requests chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	//创建一个master worker channel
	s.WorkerChannel = make(chan engine.Request)
}

func (s *SimpleScheduler) InitWorkerChannel() chan engine.Request {
	// 返回master worker channel
	return s.WorkerChannel
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// 发送request到worker channel
	go func() { s.WorkerChannel <- request }()

}

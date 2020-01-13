package service_core

type WorkerStruct struct {
	RunningFlag bool
	Stop        chan int
	Suspend     chan int
}

type FireService interface {
	Fire()
}

type AutoService interface {
	Start()
	Stop()
	Suspend()
}

type BroadCastService struct {
	BroadCastWorkerName string
	NoticeChan          chan interface{}
	QuitChan            chan struct{}
}

type BroadCastProcess interface {
	Listen()
}

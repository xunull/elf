package service_core

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

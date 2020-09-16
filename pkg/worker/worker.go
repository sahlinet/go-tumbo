package worker

type Config struct {
}

type RequestMessage struct {
}

type ResponseMessage struct {
}

type Workers interface {
	Start() error
	Stop() error

	HandleMessage(RequestMessage) (ResponseMessage, error)
	ReceiveConfiguration(Config) error
}

type Worker struct {
	Running bool
}

func (w *Worker) Start() error {
	w.Running = true
	return nil
}

func (w *Worker) Stop() error {
	w.Running = false
	return nil
}

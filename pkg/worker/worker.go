package worker

import (
	"gorm.io/gorm"
)

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
	gorm.Model
	Name    string `gorm:"unique" json:name`
	Running bool   `json:"running"`
}

func (w *Worker) Start() error {
	w.Running = true
	return nil
}

func (w *Worker) Stop() error {
	w.Running = false
	return nil
}

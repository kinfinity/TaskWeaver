package worker

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"

	"TaskWeaver/task"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("Stats!")
}

func (w *Worker) RunTask() {
	fmt.Println("Run Task")
}

func (w *Worker) StopTask() {
	fmt.Println("Stop Task")
}

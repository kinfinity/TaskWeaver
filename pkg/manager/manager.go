package manager

import (
	"taskweaver/pkg/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]task.Task
	EventDb       map[string][]task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
}

// Restart failed tasks
// update all tasks to new state
func (m *Manager) UpdateTasks() {
	// for _, worker := range m.Workers {
	// 	if len(m.WorkerTaskMap[worker]) > 0
	// 	m.WorkerTaskMap[worker]
	// }
}

func (m *Manager) SendWork() {
	// for _, worker := range m.Workers {}
}

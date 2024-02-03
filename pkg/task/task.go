package task

import (
	"time"

	"github.com/google/uuid"

)

// Task 
type Task struct {
	ID          uuid.UUID // Unique identifier for the task.
	Title       string    // Title or name of the task.
	Description string    // Additional details about the task.
	DueDate     time.Time // Deadline for completing the task.
	IsCompleted bool      // Indicates whether the task is completed or not.
	State TaskState
	// -- Track Task Time
	CreatedAt   time.Time // Timestamp of when the task was created.
	UpdatedAt   time.Time // Timestamp of when the task was last updated.
	StartTime 	time.Time
	Finish 		time.Time
	// -- Container properties
	// ExposedPorts nat.PortSet
	// PortBindings map[string]string
	// RestartPolicy string
	// Image 	string
	// Command	string
}


type TState int
type TaskState string

const (
	Pending TState = iota
	Scheduled
	Running
	Completed
	Failed
)

const (
	PendingStr TaskState = "Pending"
    ScheduledStr = "Scheduled"
    RunningStr = "Running"
    CompletedStr = "Completed"
    FailedStr = "Failed"
)
 
// Transition tasks to new state 
type TaskEvent struct{ 
	ID uuid.UUID
	Task *Task
	State TaskState // Transition to state
	Timestamp time.Time
}

// Forward and Reverse Maps
var TaskStateToTStateMap = map[TaskState]TState{
	PendingStr:   Pending,
	ScheduledStr: Scheduled,
	RunningStr:   Running,
	CompletedStr: Completed,
	FailedStr:    Failed,
}
var TStateToTaskStateMap = map[TState]TaskState{
	Pending:   PendingStr,
	Scheduled: ScheduledStr,
	Running:   RunningStr,
	Completed: CompletedStr,
	Failed:    FailedStr,
}
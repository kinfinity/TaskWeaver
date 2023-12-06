/*
Copyright © 2023 EGBEWATT M. KOKOU	 kokou.egbewatt@gmail.com
*/
package main

import (
	"TaskWeaver/manager"
	"TaskWeaver/node"
	"TaskWeaver/node/worker"
	"TaskWeaver/task"
	"TaskWeaver/utils"
	"fmt"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

func main() {
	// Setup Task ->  Event
	taskID, eventID := uuid.New(), uuid.New()
	t := task.Task{
		ID:          taskID,
		Title:       "TaskWeaver Task - 1",
		Description: "First Task for our Distributed Task Orchestration System",
		State:       task.TStateToTaskStateMap[task.Pending],
	}
	utils.PrettyPrintStruct(t)

	te := task.TaskEvent{
		ID:        eventID,
		Task:      &t,
		State:     task.TStateToTaskStateMap[task.Scheduled],
		Timestamp: time.Now(),
	}

	fmt.Println("Schedule task")
	utils.PrettyPrintStruct(te)

	// Setup Worker
	w := worker.Worker{
		Name:  "Prime Worker",
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]task.Task),
	}
	w.Db[taskID] = t

	w.CollectStats()
	w.RunTask()
	w.StopTask()

	// Setup Manager
	m := manager.Manager{
		Pending: *queue.New(),
		TaskDb: map[string][]task.Task{
			taskID.String(): {t},
		},
		EventDb: map[string][]task.TaskEvent{
			eventID.String(): {te},
		},
		Workers:       []string{w.Name},
		WorkerTaskMap: make(map[string][]uuid.UUID),
		TaskWorkerMap: make(map[uuid.UUID]string),
	}

	m.SelectWorker()
	m.UpdateTasks()
	m.SendWork()

	// Setup Node - This should just involve picking up an IP and connecting to it
	// polling it's information and installing the weaver
	n := node.Node{
		Name:   "Node-1",
		Ip:     "192.168.1.1",
		Cores:  4,
		Memory: 1024,
		Disk:   25,
		Role:   "worker",
	}
	utils.PrettyPrintStruct(n)

}

// Need threads for each system?
// figure out how to setup etcd & kafka for communication
// use GRPC or GraphQL for master - worker communication
// how are worker nodes synced to  master? install worker client on node
//

// Do we decide the type of machine our Master Node Runs? Need it for ETCD setup
// single ETCD or cluster in case of Multiple master Nodes with a leader
// control plane components and etcd communicate over secure channels using TLS certificates.
// backup etcd data stored on master node directories for data durability and consistency
// operators to automate etcd backup and restore procedures
// scale, upgrade & monitor health and performance of etcd

//

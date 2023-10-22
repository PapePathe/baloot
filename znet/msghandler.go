package znet

import (
	"fmt"
	"strconv"

	"pathe.co/zinx/utils"
	"pathe.co/zinx/ziface"
)

type MsgHandle struct {
	Apis           map[uint32]ziface.IRouter // Map to store the handler methods for each MsgID
	WorkerPoolSize uint32                    // Number of workers in the business worker pool
	TaskQueue      []chan ziface.IRequest    // Message queues that workers are responsible for picking tasks from

}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
}

// Process messages in a non-blocking manner
func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgId =", request.GetMsgID(), "is not FOUND!")
		return
	}

	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

// Add specific handling logic for a message
func (mh *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := mh.Apis[msgId]; ok {
		panic("repeated api, msgId = " + strconv.Itoa(int(msgId)))
	}
	mh.Apis[msgId] = router
	fmt.Println("Add api msgId =", msgId)
}

// Start a Worker workflow
func (mh *MsgHandle) StartOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	fmt.Println("Worker ID =", workerID, "is started.")
	for {
		request := <-taskQueue
		mh.DoMsgHandler(request)
	}
}

// Start the worker pool
func (mh *MsgHandle) StartWorkerPool() {
	// Start the required number of workers
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		// A worker is started
		// Allocate space for the current worker's task queue
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		// Start the current worker, blocking and waiting for messages in the corresponding task queue
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

// Send the message to the TaskQueue to be processed by the worker
func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	// Assign the current connection to the worker responsible for processing this connection based on ConnID
	// Round-robin average allocation policy

	// Get the workerID responsible for processing this connection
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize

	fmt.Println("Add ConnID=", request.GetConnection().GetConnID(), " request msgID=", request.GetMsgID(), "to workerID=", workerID)

	// Send the request message to the task queue
	mh.TaskQueue[workerID] <- request
}

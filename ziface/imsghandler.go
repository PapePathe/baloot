package ziface

/*
Abstract layer for message management
*/
type IMsgHandle interface {
	DoMsgHandler(request IRequest)          // Process messages immediately in a non-blocking manner
	AddRouter(msgId uint32, router IRouter) // Add specific handling logic for a message
	StartWorkerPool()                       // Start the worker pool
	SendMsgToTaskQueue(request IRequest)    // Send the message to the TaskQueue to be processed by the worker
}

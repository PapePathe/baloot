package ziface

/*
   Router interface.
   Routers are used by framework users to configure custom business methods for a connection.
   The IRequest in the router contains the connection information and the request data of that connection.
*/

type IRouter interface {
	PreHandle(request IRequest)  // Hook method executed before handling the conn business
	Handle(request IRequest)     // Method to handle the conn business
	PostHandle(request IRequest) // Hook method executed after handling the conn business
}

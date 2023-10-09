package ziface

/*
   IRequest interface:
   This interface encapsulates the client's connection
   information and request data into a Request.
*/

type IRequest interface {
	GetConnection() IConnection // Get the connection information of the request
	GetData() []byte            // Get the data of the request message
	GetMsgID() uint32
}

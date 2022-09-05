package k8s


// Handler defines a specific action that should be taken.
// Experimental.
type Handler struct {
	// One and only one of the following should be specified.
	//
	// Exec specifies the action to take.
	// Experimental.
	Exec *ExecAction `field:"optional" json:"exec" yaml:"exec"`
	// HTTPGet specifies the http request to perform.
	// Experimental.
	HttpGet *HttpGetAction `field:"optional" json:"httpGet" yaml:"httpGet"`
	// TCPSocket specifies an action involving a TCP port.
	//
	// TCP hooks not yet supported.
	// Experimental.
	TcpSocket *TcpSocketAction `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}


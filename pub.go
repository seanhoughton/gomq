package gomq

import (
	"net"

	"github.com/zeromq/gomq/zmtp"
)

// PubSocket is a ZMQ_PUB socket type.
// See: http://rfc.zeromq.org/spec:41
type PubSocket struct {
	*Socket
}

// NewPush accepts a zmtp.SecurityMechanism and returns
// a PushSocket as a gomq.Push interface.
func NewPub(mechanism zmtp.SecurityMechanism) *PubSocket {
	return &PubSocket{
		Socket: NewSocket(false, zmtp.PubSocketType, nil, mechanism),
	}
}

// Bind accepts a zeromq endpoint and binds the
// push socket to it. Currently the only transport
// supported is TCP. The endpoint string should be
// in the format "tcp://<address>:<port>".
func (s *PubSocket) Bind(endpoint string) (net.Addr, error) {
	return BindServer(s, endpoint)
}

// Connect accepts a zeromq endpoint and connects the
// client socket to it. Currently the only transport
// supported is TCP. The endpoint string should be
// in the format "tcp://<address>:<port>".
func (s *PubSocket) Connect(endpoint string) error {
	return ConnectClient(s, endpoint)
}

var (
	_ Client = (*PubSocket)(nil)
	_ Server = (*PubSocket)(nil)
)

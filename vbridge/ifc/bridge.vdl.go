// This file was auto-generated by the vanadium vdl tool.
// Source: bridge.vdl

package ifc

import (
	// VDL system imports
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/vdl"
)

type Message struct {
	Topic     string
	MessageId uint16
	Payload   []byte
}

func (Message) __VDLReflect(struct {
	Name string `vdl:"github.com/jeffallen/mqtt/vbridge/ifc.Message"`
}) {
}

type Topic string

func (Topic) __VDLReflect(struct {
	Name string `vdl:"github.com/jeffallen/mqtt/vbridge/ifc.Topic"`
}) {
}

func init() {
	vdl.Register((*Message)(nil))
	vdl.Register((*Topic)(nil))
}

// BridgeClientMethods is the client interface
// containing Bridge methods.
type BridgeClientMethods interface {
	// Subscribes to a set of topics. All messages on those topics are
	// returned by the stream.
	Subscribe(_ *context.T, topics []Topic, _ ...rpc.CallOpt) (BridgeSubscribeClientCall, error)
}

// BridgeClientStub adds universal methods to BridgeClientMethods.
type BridgeClientStub interface {
	BridgeClientMethods
	rpc.UniversalServiceMethods
}

// BridgeClient returns a client stub for Bridge.
func BridgeClient(name string) BridgeClientStub {
	return implBridgeClientStub{name}
}

type implBridgeClientStub struct {
	name string
}

func (c implBridgeClientStub) Subscribe(ctx *context.T, i0 []Topic, opts ...rpc.CallOpt) (ocall BridgeSubscribeClientCall, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Subscribe", []interface{}{i0}, opts...); err != nil {
		return
	}
	ocall = &implBridgeSubscribeClientCall{ClientCall: call}
	return
}

// BridgeSubscribeClientStream is the client stream for Bridge.Subscribe.
type BridgeSubscribeClientStream interface {
	// RecvStream returns the receiver side of the Bridge.Subscribe client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() Message
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// BridgeSubscribeClientCall represents the call returned from Bridge.Subscribe.
type BridgeSubscribeClientCall interface {
	BridgeSubscribeClientStream
	// Finish blocks until the server is done, and returns the positional return
	// values for call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
}

type implBridgeSubscribeClientCall struct {
	rpc.ClientCall
	valRecv Message
	errRecv error
}

func (c *implBridgeSubscribeClientCall) RecvStream() interface {
	Advance() bool
	Value() Message
	Err() error
} {
	return implBridgeSubscribeClientCallRecv{c}
}

type implBridgeSubscribeClientCallRecv struct {
	c *implBridgeSubscribeClientCall
}

func (c implBridgeSubscribeClientCallRecv) Advance() bool {
	c.c.valRecv = Message{}
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implBridgeSubscribeClientCallRecv) Value() Message {
	return c.c.valRecv
}
func (c implBridgeSubscribeClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implBridgeSubscribeClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// BridgeServerMethods is the interface a server writer
// implements for Bridge.
type BridgeServerMethods interface {
	// Subscribes to a set of topics. All messages on those topics are
	// returned by the stream.
	Subscribe(_ *context.T, _ BridgeSubscribeServerCall, topics []Topic) error
}

// BridgeServerStubMethods is the server interface containing
// Bridge methods, as expected by rpc.Server.
// The only difference between this interface and BridgeServerMethods
// is the streaming methods.
type BridgeServerStubMethods interface {
	// Subscribes to a set of topics. All messages on those topics are
	// returned by the stream.
	Subscribe(_ *context.T, _ *BridgeSubscribeServerCallStub, topics []Topic) error
}

// BridgeServerStub adds universal methods to BridgeServerStubMethods.
type BridgeServerStub interface {
	BridgeServerStubMethods
	// Describe the Bridge interfaces.
	Describe__() []rpc.InterfaceDesc
}

// BridgeServer returns a server stub for Bridge.
// It converts an implementation of BridgeServerMethods into
// an object that may be used by rpc.Server.
func BridgeServer(impl BridgeServerMethods) BridgeServerStub {
	stub := implBridgeServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implBridgeServerStub struct {
	impl BridgeServerMethods
	gs   *rpc.GlobState
}

func (s implBridgeServerStub) Subscribe(ctx *context.T, call *BridgeSubscribeServerCallStub, i0 []Topic) error {
	return s.impl.Subscribe(ctx, call, i0)
}

func (s implBridgeServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implBridgeServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{BridgeDesc}
}

// BridgeDesc describes the Bridge interface.
var BridgeDesc rpc.InterfaceDesc = descBridge

// descBridge hides the desc to keep godoc clean.
var descBridge = rpc.InterfaceDesc{
	Name:    "Bridge",
	PkgPath: "github.com/jeffallen/mqtt/vbridge/ifc",
	Methods: []rpc.MethodDesc{
		{
			Name: "Subscribe",
			Doc:  "// Subscribes to a set of topics. All messages on those topics are\n// returned by the stream.",
			InArgs: []rpc.ArgDesc{
				{"topics", ``}, // []Topic
			},
		},
	},
}

// BridgeSubscribeServerStream is the server stream for Bridge.Subscribe.
type BridgeSubscribeServerStream interface {
	// SendStream returns the send side of the Bridge.Subscribe server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item Message) error
	}
}

// BridgeSubscribeServerCall represents the context passed to Bridge.Subscribe.
type BridgeSubscribeServerCall interface {
	rpc.ServerCall
	BridgeSubscribeServerStream
}

// BridgeSubscribeServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements BridgeSubscribeServerCall.
type BridgeSubscribeServerCallStub struct {
	rpc.StreamServerCall
}

// Init initializes BridgeSubscribeServerCallStub from rpc.StreamServerCall.
func (s *BridgeSubscribeServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// SendStream returns the send side of the Bridge.Subscribe server stream.
func (s *BridgeSubscribeServerCallStub) SendStream() interface {
	Send(item Message) error
} {
	return implBridgeSubscribeServerCallSend{s}
}

type implBridgeSubscribeServerCallSend struct {
	s *BridgeSubscribeServerCallStub
}

func (s implBridgeSubscribeServerCallSend) Send(item Message) error {
	return s.s.Send(item)
}

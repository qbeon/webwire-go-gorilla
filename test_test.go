package gorilla_test

import (
	"context"
	"testing"

	webwire "github.com/qbeon/webwire-go"
	wwrgorilla "github.com/qbeon/webwire-go-gorilla"
	"github.com/stretchr/testify/require"
)

type TestSrvImpl struct{}

func (tsi *TestSrvImpl) OnClientConnected(
	_ webwire.ConnectionOptions,
	_ webwire.Connection,
) {
}
func (tsi *TestSrvImpl) OnClientDisconnected(_ webwire.Connection, _ error) {}
func (tsi *TestSrvImpl) OnSignal(
	_ context.Context,
	_ webwire.Connection,
	_ webwire.Message,
) {
}
func (tsi *TestSrvImpl) OnRequest(
	_ context.Context,
	_ webwire.Connection,
	_ webwire.Message,
) (payload webwire.Payload, err error) {
	return webwire.Payload{}, nil
}

// TestNewServer tests whether the implementation of the transport interface is
// accepted by the server constructor
func TestNewServer(t *testing.T) {
	srv, err := webwire.NewServer(
		&TestSrvImpl{},
		webwire.ServerOptions{},
		&wwrgorilla.Transport{
			Host: "127.0.0.1:",
		},
	)
	require.NoError(t, err)
	require.NotNil(t, srv)
}

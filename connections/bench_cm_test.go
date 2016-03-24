package connections

import (
	. "github.com/vivowares/eywa/configs"
	. "github.com/vivowares/eywa/utils"
	"strconv"
	"testing"
	"time"
)

func BenchmarkNewHttpConnection(b *testing.B) {
	cm, _ := NewConnectionManager("default")
	defer CloseConnectionManager("default")

	for n := 0; n < b.N; n++ {
		cm.NewHttpConnection(strconv.Itoa(n), "", make(chan []byte), func(Connection, *Message, error) {}, nil)
	}
}

func BenchmarkNewWsConnection(b *testing.B) {
	SetConfig(&Conf{
		Connections: &ConnectionsConf{
			Websocket: &WsConnectionConf{
				RequestQueueSize: 8,
				Timeouts: &WsConnectionTimeoutConf{
					Write:    &JSONDuration{2 * time.Second},
					Read:     &JSONDuration{300 * time.Second},
					Request:  &JSONDuration{1 * time.Second},
					Response: &JSONDuration{2 * time.Second},
				},
				BufferSizes: &WsConnectionBufferSizeConf{
					Write: 1024,
					Read:  1024,
				},
			},
		},
	})

	cm, _ := NewConnectionManager("default")
	defer CloseConnectionManager("default")

	for n := 0; n < b.N; n++ {
		cm.NewWebsocketConnection(strconv.Itoa(n), "", &fakeWsConn{}, func(Connection, *Message, error) {}, nil)
	}
}

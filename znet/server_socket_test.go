package znet

import (
	"encoding/json"
	"net"
	"testing"

	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/game"
)

func TestNewSocketHandler2(t *testing.T) {
	t.Parallel()

	setupTestApp("7778")

	conn, resp, err := websocket.DefaultDialer.Dial("ws://localhost:7778/ws/gm100", nil)
	require.NoError(t, err)

	defer conn.Close()

	assert.Equal(t, fiber.StatusSwitchingProtocols, resp.StatusCode)
	assert.Equal(t, "websocket", resp.Header.Get("Upgrade"))

	var msg game.ReceiveTakeHandMsg

	_, rawmsg, err := conn.ReadMessage()
	require.NoError(t, err)
	err = json.Unmarshal(rawmsg, &msg)
	require.NoError(t, err)
}

func TestNewSocketHandler(t *testing.T) {
	t.Parallel()
	setupTestApp("7779")

	conn, resp, err := websocket.DefaultDialer.Dial("ws://localhost:7779/ws/gm100", nil)

	require.NoError(t, err)

	defer conn.Close()

	assert.Equal(t, fiber.StatusSwitchingProtocols, resp.StatusCode)
	assert.Equal(t, "websocket", resp.Header.Get("Upgrade"))

	var msg game.ReceiveTakeHandMsg

	_, rawmsg, err := conn.ReadMessage()
	require.NoError(t, err)
	err = json.Unmarshal(rawmsg, &msg)
	require.NoError(t, err)

	assert.Len(t, msg.Player.Hand.Cards, 5)
	assert.Equal(t, game.ReceiveTakeHand, msg.ID)
	assert.Equal(t, gametake.AllTakeNames, msg.AvailableTakes)

	err = conn.WriteMessage(websocket.TextMessage, []byte("{\"id\":\"2\",\"gametake\":\"Tout\", \"player_id\":\"0\"}"))
	require.NoError(t, err)

	_, rawmsg, err = conn.ReadMessage()
	require.NoError(t, err)

	var playMsg game.ReceivePlayingHandMsg

	err = json.Unmarshal(rawmsg, &playMsg)
	require.NoError(t, err)
	assert.Equal(t, game.ReceivePlayingHand, playMsg.ID)
	assert.Equal(t, "Tout", playMsg.Take)
	assert.Len(t, playMsg.Cards, 8)
}

func setupTestApp(port string) SocketApp {
	s := NewSocketHandler()
	app := NewSocketApp(port)

	app.SetupRoutes(s)
	go app.Start()

	readyCh := make(chan struct{})

	go func() {
		for {
			conn, err := net.Dial("tcp", "localhost:7777")
			if err != nil {
				continue
			}

			if conn != nil {
				readyCh <- struct{}{}

				conn.Close()

				break
			}
		}
	}()

	<-readyCh

	return app
}

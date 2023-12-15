package znet

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http/httptest"
	"testing"

	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
	"pathe.co/zinx/pkg/game"
)

func TestIndexPage(t *testing.T) {
	t.Parallel()

	app := setupTestApp("7780")
	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.app.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)
}

func assertPlayerReceiveTakingHand(t *testing.T, i int, take string) *websocket.Conn {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:7778/ws/gm100", nil)
	require.NoError(t, err)

	var msg game.ReceiveTakeHandMsg

	_, rawmsg, err := conn.ReadMessage()
	require.NoError(t, err)
	err = json.Unmarshal(rawmsg, &msg)
	require.NoError(t, err)

	clientMsg := fmt.Sprintf("{\"id\":\"2\",\"gametake\":\"%s\", \"player_id\":\"%d\"}", take, i)
	err = conn.WriteMessage(websocket.TextMessage, []byte(clientMsg))
	require.NoError(t, err)

	assert.Len(t, msg.Player.GetHand().Cards, 5)
	assert.Equal(t, game.ReceiveTakeHand, msg.ID)
	assert.Equal(t, gametake.AllTakeNames, msg.AvailableTakes)

	return conn
}

func assertPlayerReceivesPlayingHand(t *testing.T, conn *websocket.Conn, take string) []cards.Card {
	_, rawmsg, err := conn.ReadMessage()
	require.NoError(t, err)

	var playMsg game.ReceivePlayingHandMsg

	err = json.Unmarshal(rawmsg, &playMsg)
	require.NoError(t, err)
	assert.Equal(t, game.ReceivePlayingHand, playMsg.ID)
	assert.Equal(t, take, playMsg.Take)
	assert.Len(t, playMsg.Cards, 8)

	fmt.Println("PLAYER CARDS", playMsg.Cards)

	return playMsg.Cards
}

// func TestNewSocketHandler2(t *testing.T) {
// 	t.Parallel()

// 	setupTestApp("7778")

// 	takes := []string{"Trefle", "Carreau", "Passe", "Cent"}
// 	conns := []*websocket.Conn{}
// 	playerCards := make([][]cards.Card, 4)

// 	for i := 0; i < 4; i++ {
// 		t.Run(fmt.Sprintf("start new player connection %d", i), func(t *testing.T) {
// 			c := assertPlayerReceiveTakingHand(t, i, takes[i])
// 			conns = append(conns, c)
// 		})
// 	}

// 	for i := 0; i < 4; i++ {
// 		t.Run(fmt.Sprintf("receive playing hand %d", i), func(t *testing.T) {
// 			playerCards[i] = assertPlayerReceivesPlayingHand(t, conns[i], takes[3])
// 		})
// 	}

// 	for i := 0; i < 4; i++ {
// 		conn := conns[i]

// 		cardToPlay := playerCards[i][0]

// 		clientMsg := fmt.Sprintf("{\"id\":\"4\",\"genre\":\"%s\",\"color\":\"%s\", \"player_id\":\"%d\"}", cardToPlay.Genre, cardToPlay.Couleur, i)
// 		err := conn.WriteMessage(websocket.TextMessage, []byte(clientMsg))
// 		require.NoError(t, err)
// 	}
// }

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

	assert.Len(t, msg.Player.GetHand().Cards, 5)
	assert.Equal(t, game.ReceiveTakeHand, msg.ID)
	assert.Equal(t, gametake.AllTakeNames, msg.AvailableTakes)

	clientMsg := "{\"id\":\"2\",\"gametake\":\"Tout\", \"player_id\":\"0\"}"
	err = conn.WriteMessage(websocket.TextMessage, []byte(clientMsg))
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
			conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", port))
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

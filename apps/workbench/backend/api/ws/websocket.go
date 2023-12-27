package ws

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/deamgo/workbench/auth/jwt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Connection struct {
	WS   *websocket.Conn
	Send chan []byte
}

// Connection pool
var connectionPool = struct {
	sync.RWMutex
	pool map[string]*Connection
}{pool: make(map[string]*Connection)}

var (
	upgrader = websocket.Upgrader{
		// The size of the storage space is read
		ReadBufferSize: 1024,
		// Write storage space size
		WriteBufferSize: 1024,
		// Allow cross-domain
		CheckOrigin: func(r *http.Request) bool {
			token := r.Header.Get("Authorization")
			_, err := jwt.ParseToken(strings.SplitN(token, " ", 2)[1])
			fmt.Println(token)
			return err == nil
		},
	}
)

func WSHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			wbsCon *websocket.Conn
			err    error
			data   []byte
		)
		// To complete the http reply, drop the following parameters in the httpheader
		if wbsCon, err = upgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
			return // Failure to obtain a connection is returned directly
		}
		token := c.Request.Header.Get("Authorization")
		developerID, err := jwt.ExtractIDFromToken(token)
		if err != nil {
			goto ERR
		}
		connectionPool.Lock()
		connectionPool.pool[developerID] = &Connection{
			WS:   wbsCon,
			Send: make(chan []byte),
		}
		connectionPool.Unlock()

		for {
			// Only Text and Binary data can be sent
			if _, data, err = wbsCon.ReadMessage(); err != nil {
				goto ERR // Jump to close the connection
			}
			fmt.Println(data)
			if err = wbsCon.WriteMessage(websocket.TextMessage, data); err != nil {
				goto ERR // Failed to send the message. Closed the connection
			}
		}

	ERR:
		// Close connection
		wbsCon.Close()
	}
}

func SendMsgToDeveloper(developerID string, msg []byte) error {
	connectionPool.Lock()
	connection := connectionPool.pool[developerID]
	err := connection.WS.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return err
	}
	connectionPool.Unlock()
	return nil
}

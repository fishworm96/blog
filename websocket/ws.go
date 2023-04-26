package websocket

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
}

func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zap.L().Error("upgrade(c.Writer, c.Request, nil) failed", zap.Error(err))
		return
	}
	defer conn.Close() // 关闭webSocket连接

	// 事件轮巡，处理 WebSocket 连接
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			zap.L().Error("conn.ReadMessage() failed", zap.Error(err))
			break
		}

		zap.L().Debug("messageType: ", zap.Any("messageType", messageType))
		zap.L().Debug("Received: ", zap.Any("message", message))

		// 输出 WebSocket 消息内容
		// message = append(message, strings(message))
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			zap.L().Error("conn.WriteMessage(messageType, message)", zap.Any("messageType", messageType), zap.Any("message", message), zap.Error(err))
			break
		}
	}
}
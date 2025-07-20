package metrics

import (
	"context"

	"Castle_Go/utils"
	"Castle_Go/websocket"
)

// 訂閱 Redis 頻道並推送到 WebSocket
func StartRedisSubscriber(hub *websocket.Hub) {
	pubsub := utils.RedisClient.Subscribe(context.Background(), "metrics_channel")
	ch := pubsub.Channel()

	go func() {
		for msg := range ch {
			hub.Broadcast <- []byte(msg.Payload)
		}
	}()
}

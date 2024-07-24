package chat

import (
	"first-go-server/chatUser"
	"time"
)

type Chat struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	CreatedAt int64  `json:"created_at"`
	UserIp    string `json:"user_ip"`
}

var Queue = make([]Chat, 0)

func CreateChat(user chatUser.User, message string) {
	// 채팅 메시지 생성
	// chatQueue 에 추가

	chat := Chat{
		ID:        len(Queue) + 1,
		Message:   message,
		CreatedAt: time.Now().Unix(),
		UserIp:    user.Ip,
	}

	Queue = append(Queue, chat)
}

func GetChats() []Chat {
	// 모든 채팅 메시지 얻어오기
	return Queue
}

func MaintainChats() {
	// 3초마다 오래된 채팅 메시지를 삭제하는 함수
	for {
		time.Sleep(3 * time.Second)
		now := time.Now().Unix()
		for i, chat := range Queue {
			if chat.CreatedAt+10 < now {
				Queue = append(Queue[:i], Queue[i+1:]...)
			}
		}
	}
}

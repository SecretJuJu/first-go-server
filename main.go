package main

import (
	"encoding/json"
	"first-go-server/chat"
	chatUser "first-go-server/chatUser"
	"fmt"
	"net/http"
	"time"
)

/**
post /chats 으로 채팅 메시지를 생성하고,
get /chats 으로 최근 10초동안 생성된 채팅을 얻어오는 간단한 채팅서버 입니다.
*/

func main() {
	http.HandleFunc("/chats", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createChat(w, r)
		case http.MethodGet:
			getCurrentChats(w)
		}
	})

	// go routine 으로 3초마다 오래된 채팅 메시지를 삭제하는 함수를 실행
	go chat.MaintainChats()

	fmt.Println("Server is running on 8080 port")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func createChat(w http.ResponseWriter, r *http.Request) {
	// chatQueue 에 추가
	user := chatUser.FindOrCreateUser(r.RemoteAddr)
	now := time.Now().Unix()
	if user.LastChatAt+1 > now {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	chatUser.UpdateLastChatAt(r.RemoteAddr, now)
	chat.CreateChat(user, req.Message)

	w.WriteHeader(http.StatusCreated)
	return
}

func getCurrentChats(w http.ResponseWriter) {
	// 채팅 메시지 얻어오기
	chats := chat.GetChats()

	// json 모듈 사용
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string][]chat.Chat{"chats": chats})
	if err != nil {
		return
	}
	return
}

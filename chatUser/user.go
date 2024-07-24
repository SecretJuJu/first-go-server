package chatUser

type User struct {
	Ip         string `json:"ip"`
	LastChatAt int64  `json:"last_chat_at"`
}

var Map = make(map[string]User)

func FindOrCreateUser(ip string) User {
	user, exist := Map[ip]
	if exist {
		return user
	}

	Map[ip] = User{Ip: ip, LastChatAt: 0}
	return User{Ip: ip, LastChatAt: 0}
}

func UpdateLastChatAt(ip string, now int64) {
	user := Map[ip]
	user.LastChatAt = now
	Map[ip] = user
}

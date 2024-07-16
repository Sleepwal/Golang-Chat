package chat

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register: // 用户进入聊天室
			h.clients[client] = true

		case client := <-h.unregister: // 用户退出聊天室
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
		case massage := <-h.broadcast: // 广播消息
			for client := range h.clients { // 遍历所有用户
				select {
				case client.Send <- massage: // 发送消息
				default: // 发送失败
					delete(h.clients, client)
					close(client.Send)
				}
			}
		}
	}
}

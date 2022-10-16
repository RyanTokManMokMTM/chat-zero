package ws

type ClientHub struct {
	clients    map[int]Client //user id -> client
	register   chan Client
	unregister chan Client
	broadcast  chan []byte
}

func NewClientHub() *ClientHub {
	return &ClientHub{
		clients:    make(map[int]Client),
		register:   make(chan Client),
		unregister: make(chan Client),
		broadcast:  make(chan []byte),
	}
}

func (h *ClientHub) Start() {
	for {
		select {
		case client := <-h.register:
			//register to map
			h.clients[client.Id] = client
		case client := <-h.unregister:
			//remove from map
			if _, ok := h.clients[client.Id]; ok {
				delete(h.clients, client.Id)
				close(client.send)
			}
		case message := <-h.broadcast:
			//send message to client
			for _, client := range h.clients {
				select {
				case client.send <- message:
				}
			}
		}
	}
}

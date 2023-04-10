package events

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"sync"
)

type EventMessage struct {
	EventName string
	Data      string
}

type HandlerEvent struct {
	m       sync.RWMutex
	clients map[string]*client
}

func NewHandlerEvent() *HandlerEvent {
	return &HandlerEvent{
		clients: make(map[string]*client),
	}
}

func (h *HandlerEvent) Handler(c echo.Context) error {
	// Configura la respuesta HTTP como un evento de servidor
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Response().Header().Set("Content-Type", "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")

	id := c.QueryParam("id")
	client := newClient(id)
	h.AddClient(client)
	fmt.Println("connected id: ", id)
	client.OnLine(c)
	fmt.Println("disconnected id: ", id)
	h.RemoveClient(client)
	return nil
}

func (h *HandlerEvent) AddClient(c *client) {
	h.m.Lock()
	defer h.m.Unlock()
	h.clients[c.ID] = c
}

func (h *HandlerEvent) RemoveClient(c *client) {
	h.m.Lock()
	h.m.Unlock()
	delete(h.clients, c.ID)
}

func (h *HandlerEvent) Broadcast(msg EventMessage) {
	fmt.Println("broadcast message", h.clients)
	h.m.RLock()
	defer h.m.RUnlock()
	for _, c := range h.clients {
		fmt.Println("send message to: ", c.ID)
		c.SendMessage <- msg
	}
}

func (h *HandlerEvent) ListClients() []string {
	h.m.RLock()
	defer h.m.RUnlock()
	var list []string
	for _, c := range h.clients {
		list = append(list, c.ID)
	}
	return list
}

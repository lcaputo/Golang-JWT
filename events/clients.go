package events

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type client struct {
	ID          string
	SendMessage chan EventMessage
}

func newClient(id string) *client {
	return &client{
		ID:          id,
		SendMessage: make(chan EventMessage),
	}
}

func (c *client) OnLine(ctx echo.Context) {
	flusher, ok := ctx.Response().Writer.(http.Flusher)
	if !ok {
		return
	}
	for {
		select {
		case msg := <-c.SendMessage:
			// Format SSE Event
			event := msg.EventName
			id := ""
			data, _ := json.Marshal(msg.Data)
			response := fmt.Sprintf("event: %s\nid: %s\ndata: %s\n\n", event, id, data)
			_, err := fmt.Fprint(ctx.Response().Writer, response)
			if err != nil {
				return
			}
			flusher.Flush()
		case <-ctx.Request().Context().Done():
			return
		}

	}
}

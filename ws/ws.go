package ws

import (
	"golang.org/x/net/websocket"
	"io"
	"fmt"
	"github.com/JILeXanDR/parser/repositories"
	"github.com/JILeXanDR/parser/models"
	"log"
	"strconv"
)

type webSocketMessage struct {
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"`
}

func NewWebSocketMessage(name string, payload interface{}) *webSocketMessage {
	return &webSocketMessage{
		Event:   name,
		Payload: payload,
	}
}

var WebSocketClients Clients

type Clients map[string]*websocket.Conn

func (c Clients) add(connection *websocket.Conn) {
	c[connection.Request().Header.Get("Sec-WebSocket-Key")] = connection
}

func (c Clients) remove(connection *websocket.Conn) {
	delete(c, connection.Request().Header.Get("Sec-WebSocket-Key"))
}

func SendMessageToAllClients(msg *webSocketMessage) {
	log.Println("Send ws message to all clients. Active clients count is " + strconv.Itoa(len(WebSocketClients)))
	for _, ws := range WebSocketClients {
		if err := websocket.JSON.Send(ws, msg); err != nil {
			fmt.Println("Can't send", err.Error())
		}
	}
}

func SendMessageToClient(ws *websocket.Conn, msg *webSocketMessage) {
	if err := websocket.JSON.Send(ws, msg); err != nil {
		fmt.Println("Can't send", err.Error())
	}
}

func waitIncomingMessage(ws *websocket.Conn) error {

	var err error
	var message interface{}

	// получение сообщения от клиента
	if err = websocket.Message.Receive(ws, &message); err != nil {
		// TODO соединение было закрыто?
		if err == io.EOF {
			WebSocketClients.remove(ws)
			fmt.Println("Connection was closed by client")
		} else {
			fmt.Println("Can't receive: ", err.Error())
		}
		return err
	}
	return nil
}

// получение сообщений от клиента
func WebSocketHandler(ws *websocket.Conn) {

	// установка соединения, добавляем клиента в массив
	WebSocketClients.add(ws)

	// отправка сообщения после подключения
	SendMessageToClient(ws, NewWebSocketMessage("main", "Привет!"))

	last, _ := repositories.LastStatistics(models.CURRENCY_USD)
	SendMessageToClient(ws, NewWebSocketMessage("rates", last))

	for {
		if err := waitIncomingMessage(ws); err != nil {
			break
		}
	}
}

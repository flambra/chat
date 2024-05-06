// server/websocket.go

package server

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type Connection struct {
	WS   *websocket.Conn
	Send chan []byte
}

var Upgrade = func(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		log.Println("Upgrade")
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}
// WebSocketHandler é um manipulador WebSocket simplificado
func WebSocketHandler() func(*websocket.Conn) {
	return func(conn *websocket.Conn) {
		
		conn.SetPongHandler(func(appData string) error {
			log.Println("Received pong from client")
			return nil
		})

		client := &Connection{
			WS:   conn,
			Send: make(chan []byte, 256),
		}

		// Configura o CloseHandler para capturar desconexões
		conn.SetCloseHandler(func(code int, text string) error {
			log.Printf("Client disconnected with code: %d, message: %s", code, text)
			return nil
		})

		log.Println("New WebSocket client connected")

		// Iniciar goroutines para leitura e escrita
		go client.readPump()
		go client.writePump()
	}
}

func (c *Connection) readPump() {
	defer func() {
		if c.WS != nil {
			c.WS.Close()
		}
		log.Println("Closed connection for client")
	}()

	for {
		// Verificar se a conexão está ativa
		if c.WS == nil {
			log.Println("Attempted to read from a nil WebSocket connection")
			break
		}

		// Ler mensagem do WebSocket
		// _, message, err := c.WS.ReadMessage()
		_, message, err := c.WS.ReadMessage()
		if err != nil {
			// Identificar diferentes tipos de erros
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Println("Client closed the WebSocket connection")
			} else {
				log.Printf("Error reading from WebSocket: %v", err)
			}
			break
		}

		// Processar a mensagem recebida
		log.Printf("Received message from client: %s", string(message))

		// Adicionar a mensagem ao canal de envio
		c.Send <- append([]byte{websocket.TextMessage}, message...)
	}
}

func (c *Connection) writePump() {
	for message := range c.Send {
		msgType := int(message[0])
		if err := c.WS.WriteMessage(msgType, message[1:]); err != nil {
			log.Println("Error writing to WebSocket:", err)
			return
		}
		log.Println("Message successfully sent to WebSocket client:", string(message[1:]))
	}
}

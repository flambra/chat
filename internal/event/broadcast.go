package event

import (
	"bufio"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Channel = make(chan primitive.M, 10)

func Broadcast(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	c.Status(fiber.StatusOK).Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		for {
			select {
			case change := <-Channel:
				eventParsed, err := Parse(change)
				if err != nil {
					log.Printf("Error parsing change: %v", err)
					continue
				}
				jsonData, err := json.Marshal(eventParsed)
				if err != nil {
					log.Printf("Error marshaling event data to JSON: %v", err)
					continue
				}
				w.WriteString("data: " + string(jsonData) + "\n")
				log.Println("Sent Change")
			case <-time.After(30 * time.Second):
				w.WriteString(":keep-alive\n")
			}
			if err := w.Flush(); err != nil {
				log.Printf("Error while flushing: %v. Closing HTTP connection.\n", err)
				break
			}
			time.Sleep(1 * time.Second)
		}
	}))

	return nil
}

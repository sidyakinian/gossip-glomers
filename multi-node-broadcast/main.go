package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type MessageBody struct {
    Type    string `json:"type"`
    Message int    `json:"message"`
}

func main() {
	n := maelstrom.NewNode()
	// create an array to store messages
	var messages []int
	
	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var body MessageBody
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
	
		messages = append(messages, body.Message)
	
		// Echo the original message back with the updated message type.
		response_body := map[string]string{
			"type": "broadcast_ok",
		}
		return n.Reply(msg, response_body)
	})

	n.Handle("read", func(msg maelstrom.Message) error {
		response_body := map[string]any{
			"type": "read_ok",
			"messages": messages,
		}
		return n.Reply(msg, response_body)
	})

	n.Handle("topology", func (msg maelstrom.Message) error {
		response_body := map[string]any{
			"type": "topology_ok",
		}
		return n.Reply(msg, response_body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
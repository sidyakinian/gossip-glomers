package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type MessageBody struct {
    Type    string `json:"type"`
    Element int    `json:"message"`
}

type server struct {
	node *maelstrom.Node
	elements []int
}

func main() {
	node := maelstrom.NewNode()
	server := &server{node: node, elements: []int{}}
	
	node.Handle("broadcast", server.handleBroadcast)
	node.Handle("read", server.handleRead)
	node.Handle("topology", server.handleTopology)

	if err := node.Run(); err != nil {
		log.Fatal(err)
	}
}

func (server *server) handleBroadcast(message maelstrom.Message) error {
	var body MessageBody
	if err := json.Unmarshal(message.Body, &body); err != nil {
		return err
	}

	server.elements = append(server.elements, body.Element)

	response_body := map[string]string{
		"type": "broadcast_ok",
	}
	return server.node.Reply(message, response_body)
}

func (server *server) handleRead(message maelstrom.Message) error {
	response_body := map[string]any{
		"type": "read_ok",
		"messages": server.elements,
	}
	return server.node.Reply(message, response_body)
}

func (server *server) handleTopology(message maelstrom.Message) error {
	response_body := map[string]any{
		"type": "topology_ok",
	}
	return server.node.Reply(message, response_body)
}
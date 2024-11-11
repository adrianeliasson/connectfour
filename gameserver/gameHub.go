// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type GameHub struct {
	// Registered clients.
	queue  *Player
	gaming map[*Player]bool

	// Register requests from the clients.
	register chan *Player

	// Unregister requests from clients.
	unregister chan *Player
}

func newGameHub() *GameHub {
	return &GameHub{
		queue:      nil,
		register:   make(chan *Player),
		unregister: make(chan *Player),
		gaming:     make(map[*Player]bool),
	}
}

func (h *GameHub) run() {
	for {
		select {
		case client := <-h.register:
			if h.queue != nil {
				h.queue.opponent = client
				client.opponent = h.queue
				h.gaming[client] = true
				h.gaming[h.queue] = true
				h.queue = nil
			} else {
				h.queue = client
			}
		case client := <-h.unregister:
			if h.queue == client {
				h.queue.conn.Close()
				h.queue = nil
			}
			if _, ok := h.gaming[client]; ok {
				delete(h.gaming, client)
				close(client.inboundMoves)
				close(client.outboundMoves)
			}
		}
	}
}

package main

import "strings"

// hub is the central entity with which the clients connect and register with. It manages all the channels,
// messages braodcast, registration-degregistration of client etc.
// channels is the map of channel name and channel.
// client : a map of the clients(connected users), with the username as key and *client as value.
// commands : a channel of command that are flowing from the client to hub, that hub will validate and execute.
// deregistration : a channel of *client through which client will unregister itself from the hub, and hub will delete all
// its references.
// registrationi : a channel of *client through which client will register itself to the hub. It will add the client to the usermap.
type hub struct {
	channels       map[string]*channel
	clients        map[string]*client
	commands       chan command
	deregistration chan *client
	registration   chan *client
}

// newHub returns new hub instance.
func newHub() *hub {
	return &hub{
		registration:   make(chan *client),
		deregistration: make(chan *client),
		clients:        make(map[string]*client),
		channels:       make(map[string]*channel),
		commands:       make(chan command),
	}
}

// run is the central method of hub. It takes command from the channels and calls the specific method for that channel.
// run runs infinetly. and it run in the main using goroutine.
func (h *hub) run() {
	for {
		select {
		case client := <-h.registration:
			h.register(client)
		case client := <-h.deregistration:
			h.unregister(client)
		case cmd := <-h.commands:
			switch cmd.id {
			case JOIN:
				h.joinChannel(cmd.sender, cmd.recipient)
			case LEAVE:
				h.leaveChannel(cmd.sender, cmd.recipient)
			case MSG:
				h.message(cmd.sender, cmd.recipient, cmd.body)
			case USRS:
				h.listUsers(cmd.sender)
			case CHNS:
				h.listChannels(cmd.sender)
			default:
				panic("Incorrect command")
			}
		}
	}
}

// register is use to check if the user is alread registered else it adds the user in the h.clients map.
// and write OK to the connection.
func (h *hub) register(c *client) {
	if _, exists := h.clients[c.username]; exists {
		c.username = ""
		c.conn.Write([]byte("Err: Username already taken \n"))
	} else {
		h.clients[c.username] = c
		c.conn.Write([]byte("OK\n"))
	}
}

// unregister is used to delete the client from the server.
// it first deletes the user from h.client map then deletes the user from all the channels present.
func (h *hub) unregister(c *client) {
	if _, exists := h.clients[c.username]; exists {
		delete(h.clients, c.username)
		for _, channel := range h.channels {
			delete(channel.clients, c)
		}
	}
}

// joinChannel is used add a user to the channel, if the channel is not already present then it adds a
// new channel and write `OK` to the connection.
func (h *hub) joinChannel(u string, c string) {
	if client, ok := h.clients[u]; ok {
		if channel, ok := h.channels[c]; ok {
			// Channel exists, Join
			channel.clients[client] = true
		} else {
			// Channel doesn't exists, create and join
			ch := newChannel(c)
			ch.clients[client] = true
			h.channels[c] = ch
		}
		client.conn.Write([]byte("OK\n"))
	}
}

// leave channel is used to remove a user from the channel. It checks for the user in the map,
// if it is present then it checks for the user in the channel and if it present than it deletes it from that
// channel.
func (h *hub) leaveChannel(u string, c string) {
	if client, ok := h.clients[u]; ok {
		if channel, ok := h.channels[c]; ok {
			delete(channel.clients, client)
		}
	}
}

// message is used to write the message to the channel or user. For writing into the channel is uses
// broadcast method of channel.
func (h *hub) message(u, r string, m []byte) {
	if sender, ok := h.clients[u]; ok {
		switch r[0] {
		case '#':
			if channel, ok := h.channels[r]; ok {
				if _, ok := channel.clients[sender]; ok {
					channel.broadcast(sender.username, m)
				} else {
					sender.conn.Write([]byte("ERR : No such channel exists\n"))
				}

			}
		case '@':
			if user, ok := h.clients[r]; ok {
				msg := append([]byte(user.username+":"), m...)
				msg = append(msg, '\n')
				user.conn.Write(msg)
			} else {
				sender.conn.Write([]byte("ERR : No such user is present\n"))
			}
		default:
			sender.conn.Write([]byte("ERR : Invalid MSG command\n"))
		}

	}
}

// listChannels is used to write all the channels that are present in the hub.
func (h *hub) listChannels(u string) {
	if client, ok := h.clients[u]; ok {
		var names []string
		if len(h.channels) == 0 {
			client.conn.Write([]byte("ERR : No such channel is found\n"))
		}
		for c := range h.channels {
			names = append(names, c)
		}
		resp := strings.Join(names, ", ")
		client.conn.Write([]byte(resp + "\n"))
	}
}

// listUsers is used to print all the users that present in the hub.
func (h *hub) listUsers(u string) {
	if client, ok := h.clients[u]; ok {
		var names []string
		for c := range h.clients {
			names = append(names, c)
		}
		resp := strings.Join(names, ", ")
		client.conn.Write([]byte(resp + "\n"))
	}
}

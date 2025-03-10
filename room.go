package main

import "net"

type Room struct {
	name    string
	members map[net.Addr]*Client
}

func (r *Room) broadcast(sender *Client, msg string) {
	for addr, client := range r.members {
		if addr != sender.conn.RemoteAddr() {
			client.msg(msg)
		}
	}
}

package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Server struct {
	Listener net.Listener
	rooms    map[string]*Room
	commands chan command
}

func NewServer() *Server {
	return &Server{
		rooms:    make(map[string]*Room),
		commands: make(chan command),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", ":4000")
	if err != nil {
		return err
	}
	defer ln.Close()

	fmt.Println("Listening on addr", ln.Addr())
	s.Listener = ln

	s.acceptLoop()
	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			fmt.Println("unable to accept connection", err)
			continue
		}
		go s.newClient(conn)
	}
}

func (s *Server) handleCmds() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client)
		case CMD_MSG:
			s.message(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

func (s *Server) newClient(conn net.Conn) {
	fmt.Println("New client has connected:", conn.RemoteAddr())

	c := &Client{
		conn:     conn,
		nick:     "anonymous",
		commands: s.commands,
	}

	c.readInput()
}

func (s *Server) nick(c *Client, args []string) {
	c.nick = args[1]
	c.msg(fmt.Sprintf("your nickname is now %s", c.nick))
}

func (s *Server) join(c *Client, args []string) {
	roomName := args[1]
	room, ok := s.rooms[roomName]
	if !ok {
		room = &Room{
			name:    roomName,
			members: make(map[net.Addr]*Client),
		}
		s.rooms[roomName] = room
	}
	room.members[c.conn.RemoteAddr()] = c

	if c.room != room {
		s.quitCurrentRoom(c)
	}
	c.room = room
	room.broadcast(c, fmt.Sprintf("%s has joined the room", c.nick))
	c.msg(fmt.Sprintf("you joined the room %s", room.name))
}

func (s *Server) listRooms(c *Client) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}
	c.msg(fmt.Sprintf("Created rooms are: \n%s", strings.Join(rooms, "\n")))
}

func (s *Server) message(c *Client, args []string) {
	if c.room == nil {
		c.msg("you must join a room before you message")
		return
	}

	c.room.broadcast(c, fmt.Sprintf("%s: %s\n", c.nick, strings.Join(args[1:], " ")))
}

func (s *Server) quit(c *Client) {
	if c.room != nil {
		log.Printf("client %s has disconnected from the room %s", c.conn.RemoteAddr().String(), c.room.name)
		c.msg(fmt.Sprintf("you disconnected from the room %s", c.room.name))
		s.quitCurrentRoom(c)
	}
	c.msg("no room joined")
}

func (s *Server) quitCurrentRoom(c *Client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf("%s has left the chat", c.nick))
	}
}

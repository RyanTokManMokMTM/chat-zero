package server

import (
	"errors"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"sync"
)

////Server option
//type ServerOption struct {
//}

//WebSocket server struct
type Server struct {
	once sync.Once
	id   string
	addr string
	sync sync.Mutex

	//User group
	users map[string]net.Conn
}

func NewServer(id, addr string) *Server {
	return newServer(id, addr)
}

func newServer(id, addr string) *Server {
	return &Server{
		id:    id,
		addr:  addr,
		users: make(map[string]net.Conn, 100), //allow at most 100 user to connect
	}
}

//Start server
func (s *Server) Start() error {
	mux := http.NewServeMux() //get new server mux
	log := logrus.WithFields(logrus.Fields{
		"module": "Server",
		"listen": s.addr,
		"id":     s.id,
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//upgrade http to web socket
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			conn.Close()
			return
		}

		//get username from uri
		name := r.URL.Query().Get("name")
		if len(name) == 0 {
			conn.Close()
			return
		}

		//handling
		old, ok := s.addUser(name, r.RemoteAddr, conn)
		// if this user already connected
		// disconnect the old one
		if ok {
			old.Close()
		}
		log.Infof("user: %s loged in", name)

		//open a goroutine for receive and send message
		go func(name string, conn net.Conn) {
			//read message
			err := s.EventLoop(name, conn)
			if err != nil {
				log.Error(err)
			}

			conn.Close()

			//remove the user
			s.removeUser(name)
			log.Infof("user %s disconnected", name)
		}(name, conn)
	})
	log.Info("Started")
	return http.ListenAndServe(s.addr, mux)
}

func (s *Server) EventLoop(user string, conn net.Conn) error {
	for {
		//read tcp frame from conn
		frame, err := ws.ReadFrame(conn)
		if err != nil {
			return err
		}

		//user is leave?
		if frame.Header.OpCode == ws.OpClose {
			return errors.New("user is closed the connection")
		}

		//Data is encode?
		//Data is added a mask to data
		if frame.Header.Masked {
			ws.Cipher(frame.Payload, frame.Header.Mask, 0)
		}

		if frame.Header.OpCode == ws.OpText {
			//get a text
			s.handleMessage(user, string(frame.Payload))
		}

	}
}

func (s *Server) handleMessage(user, message string) {
	logrus.Infof("Recv message %s from %s ", message, user)
	s.sync.Lock()
	defer s.sync.Unlock()

	broadCast := fmt.Sprintf("%s -- From %s", message, user)
	//send to all user except itself
	for u, conn := range s.users {
		if u == user {
			continue
		}
		err := s.WriteText(conn, broadCast)
		if err != nil {
			logrus.Errorf("write to %s failed,err : %v", u, err)
		}
	}
}

func (s *Server) WriteText(conn net.Conn, message string) error {
	f := ws.NewTextFrame([]byte(message))
	return ws.WriteFrame(conn, f)
}

func (s *Server) addUser(name, addr string, conn net.Conn) (net.Conn, bool) {
	s.sync.Lock()
	defer s.sync.Unlock()
	//map is thread unsafe
	old, ok := s.users[name] //Get old connection??
	s.users[name] = conn

	return old, ok
}

func (s *Server) removeUser(name string) {
	s.sync.Lock()
	defer s.sync.Unlock()
	//map is thread unsafe
	delete(s.users, name)
}

func (s *Server) Shutdown() {
	s.once.Do(func() { //do once only
		s.sync.Lock()
		defer s.sync.Lock()

		//close all connection
		for _, conn := range s.users {
			conn.Close()
		}

	})
}

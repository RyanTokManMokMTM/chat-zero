package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/sirupsen/logrus"
	"net"
	"net/url"
	"time"
)

type StartOpt struct {
	addr string
	name string
}

type handler struct {
	conn  net.Conn
	close chan struct{}
	msg   chan []byte
}

func connect(addr string) (*handler, error) {
	logrus.Info("client is connecting to sever")
	_, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	conn, _, _, err := ws.Dial(context.Background(), addr)
	if err != nil {
		return nil, err
	}

	han := handler{
		conn:  conn,
		close: make(chan struct{}),
		msg:   make(chan []byte, 10), //at most 10 byte(10 word)
	}

	//receive msg
	go func() {
		err := han.EventLoop(conn)
		if err != nil {
			logrus.Warn(err)
		}

		han.close <- struct{}{} //close connection
	}()

	return &han, nil
}

func (h *handler) EventLoop(conn net.Conn) error {
	logrus.Info("EventLoop started")
	for {
		frame, err := ws.ReadFrame(h.conn)
		if err != nil {
			return err
		}

		if frame.Header.OpCode == ws.OpClose {
			return errors.New("connection is closed")
		}

		if frame.Header.OpCode == ws.OpText {
			h.msg <- frame.Payload //
		}
	}
}

func (h *handler) SendText(msg string) error {
	logrus.Info("send message : ", msg)

	return wsutil.WriteClientText(h.conn, []byte(msg))
}

func run(ctx context.Context, opts *StartOpt) error {
	url := fmt.Sprintf("%s?name=%s", opts.addr, opts.name)
	logrus.Info("client connects to ", url)
	h, err := connect(url)
	if err != nil {
		return err
	}

	//go routine to read message
	go func() {
		for msg := range h.msg {
			logrus.Info("Received message :", string(msg))
		}
	}()

	tk := time.NewTicker(time.Second * 5)

	for {
		select {
		case <-tk.C:
			err := h.SendText("hello")
			if err != nil {
				logrus.Error(err)
			}
		case <-h.close:
			return nil
		}
	}
}

package service

import (
	"bufio"
	"context"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"
	"net"
	"os"

	"tcp_echo_client/internal/config"
)

type service struct {
	logger           *logan.Entry
	cfg 			 config.Config
}

func NewService(cfg config.Config) *service {
	return &service{
		logger:    cfg.Log(),
		cfg:	   cfg,
	}
}

func (s *service) Run(ctx context.Context) error{
	for {
		conn, _ := net.Dial("tcp", "127.0.0.1:5101")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text + "\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)
	}

	return nil
}


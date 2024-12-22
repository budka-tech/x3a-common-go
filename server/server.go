package server

import (
	"github.com/budka-tech/snip-common-go/port"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	listener net.Listener
	port     port.Port
}

func New(port port.Port) *Server {
	return &Server{
		port: port,
	}
}

func (receiver *Server) Init(beforeLister, afterListener func() error) {
	var err error

	if beforeLister != nil {
		err = beforeLister()
		if err != nil {
			log.Fatal(err)
		}
	}

	receiver.listener, err = net.Listen("tcp", port.FormatTCP(port.Users))
	if err != nil {
		log.Fatalf("Ошибка прослушивания порта: %v", err)
	}
	s := grpc.NewServer()
	log.Printf("Сервис запущен на %v", receiver.listener.Addr())

	if afterListener != nil {
		err = afterListener()
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := s.Serve(receiver.listener); err != nil {
		log.Fatalf("Ошибка сервиса: %v", err)
	}
}

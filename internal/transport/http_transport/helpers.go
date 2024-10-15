package http_transport

import (
	"fmt"
)

// Listen запускает соответсвующую конфигурацию для http\https
func (s *Server) Listen() error {

	fmt.Printf("Application started (bound on host and port " + s.cfg.Host + ")")
	err := s.a.Listen(s.cfg.Host)
	return err
}

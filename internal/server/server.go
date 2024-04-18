package server

import (
	"net/http"
	"time"
	"web-forum/pkg"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	pkg.InfoLog.Printf("Server run on http://localhost%s", port)

	return s.httpServer.ListenAndServe()
}

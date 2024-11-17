package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

func NewServer(l net.Listener, mux http.Handler) *Server {
	return &Server{
		l:   l,
		srv: &http.Server{Handler: mux},
	}
}

func (s *Server) Run(ctx context.Context) error {
	// Interrupt signal 신호 처리를 위한 context 생성
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop() // signal.NotifyContext로 등록된 signal 핸들러를 해제

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := s.srv.Serve(s.l); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close %+v", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("Failed to shutdown: %+v", err)
	}
	return eg.Wait()
}

package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to listen port %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx, l)
	})

	in := "message"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)
	t.Logf("try reqeust to %q", url)
	rsp, err := http.Get(url)

	if err != nil {
		t.Errorf("Failed to get response: %v", err)
	}
	defer rsp.Body.Close()

	got, err := io.ReadAll(rsp.Body)

	if err != nil {
		t.Errorf("Failed to read response body: %v", err)
	}
	want := fmt.Sprintf("Hello, %s", in)
	if string(got) != want {
		t.Errorf("want %q, got %q", want, got)
	}
	cancel()
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}

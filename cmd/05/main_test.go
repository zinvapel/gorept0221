package main

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		select {
		case <-ctx.Done():
		case <-time.After(time.Second):
			t.Error("Timeout exceed")
			os.Exit(1)
		}
	}()

	Do()

	cancelFunc()
}

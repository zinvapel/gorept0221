package main

import (
	"context"
	"time"
)

func ExampleDo() {
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	Do(ctx)
	// Output:
	// First
	// Third
}
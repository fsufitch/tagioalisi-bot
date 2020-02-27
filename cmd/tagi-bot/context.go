package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// InterruptContext is a channel containing interrupts to stop the application
type InterruptContext context.Context

// ProvideInterruptContext builds an InterruptContext from SIGTERM signals
func ProvideInterruptContext() InterruptContext {
	interrupt := make(chan os.Signal, 2)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-interrupt
		cancel()
	}()
	return InterruptContext(ctx)
}

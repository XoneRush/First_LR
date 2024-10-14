package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"server/tcp_util"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	wg.Add(1)
	go func() {
		defer wg.Done()
		tcp_util.Start("localhost", "443", stop)
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// Perform application shutdown with a maximum timeout of 5 seconds.
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	select {
	case <-timeoutCtx.Done():
		if timeoutCtx.Err() == context.DeadlineExceeded {
			log.Fatalln("timeout exceeded, forcing shutdown")
		}

		os.Exit(0)
	}
}

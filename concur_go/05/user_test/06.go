package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type Limit float64

// NewLimiter creates a new rate limiter.
func NewLimiter(r Limit, b int) *rate.Limiter {
	return rate.NewLimiter(rate.Limit(r), b)
}

func Every(interval time.Duration) Limit {
	return Limit(float64(time.Second) / float64(interval))
}

// 1秒に一つのアクセスしか許さないようにした
func Open() *APIConnection {
	return &APIConnection{
		rateLimiter: rate.NewLimiter(rate.Limit(1), 1),
	}
}

type APIConnection struct {
	rateLimiter *rate.Limiter
}

func (a *APIConnection) ReadFile(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	// Simulate file reading operation
	return nil
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	// Simulate address resolution operation
	return nil
}

func main() {
	defer log.Printf("Done.")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("cannot ReadFile: %v\n", err)
				return
			}
			log.Printf("ReadFile")
		}(i)
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("cannot ResolveAddress: %v\n", err)
				return
			}
			log.Printf("ResolveAddress")
		}(i)
	}
	wg.Wait()
}

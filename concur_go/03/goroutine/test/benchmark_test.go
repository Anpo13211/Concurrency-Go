package main

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- struct{}{}
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer() // Start timing the benchmark
	close(begin)   // Start the sender and receiver
	wg.Wait()      // Wait for them to finish
	b.StopTimer()  // Stop timing the benchmark
}

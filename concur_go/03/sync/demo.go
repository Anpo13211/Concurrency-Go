package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	countCh := make(chan int) // インクリメント/デクリメント操作の要求を送信するためのチャネル
	doneCh := make(chan bool) // 操作が完了したことを送信するためのチャネル

	increment := func(ch chan<- int) {
		ch <- 1 // インクリメントの要求をチャネルに送信
	}
	decrement := func(ch chan<- int) {
		ch <- -1 // デクリメントの要求をチャネルに送信
	}

	var arithmetic sync.WaitGroup

	// カウントを管理するゴルーチン
	go func() {
		for {
			select {
			case x := <-countCh: // x が countCh からの情報を読み込む（書き込まれるまでポーズされる）
				count += x
				if x > 0 {
					fmt.Printf("Incrementing: %d\n", count)
				} else {
					fmt.Printf("Decrementing: %d\n", count)
				}
			case <-doneCh: // 処理完了の通知を受けたら終了
				close(countCh)
				return
			}
		}
	}()

	// インクリメントのゴルーチンを起動
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment(countCh)
		}()
	}
	// デクリメントのゴルーチンを起動
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement(countCh)
		}()
	}

	// 全てのゴルーチンの完了を待つ
	arithmetic.Wait()
	doneCh <- true // 管理ゴルーチンに終了を通知
	fmt.Println("Arithmetic complete.")
}

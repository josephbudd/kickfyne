package main

import (
	"context"
	"fmt"
)

func startDumper(ctx context.Context) (dumperCh chan string) {
	dumperCh = make(chan string)
	go func(x context.Context, ch chan string) {
		for {
			select {
			case <-x.Done():
				// Drain the buffered channel.
				for {
					if len(ch) > 0 {
						msg := <-ch
						fmt.Println(msg)
					} else {
						break
					}
				}
				fmt.Println("startDumper: Done()")
				return
			case msg := <-ch:
				fmt.Println(msg)
			}
		}
	}(ctx, dumperCh)

	return
}

package partie6

import (
	"fmt"
	"time"
)

func Partie6() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(500 * time.Millisecond)
		}
	}()
	go func() {
		for {
			c2 <- "Every 2s"
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		select {
		case res1 := <-c1:
			fmt.Println(res1)
		case res2 := <-c2:
			fmt.Println(res2)
		}
	}
}
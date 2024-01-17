package partie5

import (
	"fmt"
	"time"
	"strconv"
)

func Count(thing string, time2 int, c chan string) {
	for i := 0; i < time2; i++ {
		c <- strconv.Itoa(i) + thing
		time.Sleep(500 * time.Millisecond)
	}
	close(c)
}

func Partie5() {
	c := make(chan string, 2)	// chanel buffer
	c <- "hello"
	fmt.Println("j'ai passé la chan")
	c <- "re coucou"
	msg := <-c
	fmt.Println(msg)
	msg = <-c
	fmt.Println(msg)
}
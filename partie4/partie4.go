package partie4

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

func Partie4() {
	chanel := make(chan string)

	go Count("Mouton(s)", 3, chanel)
	for {
		message, isOpen := <-chanel
		if !isOpen {
			break
		}
		fmt.Println(message)
	}
}
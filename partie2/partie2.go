package partie2

import (
	"fmt"
	"time"
)

func Count(thing string) {
	for i := 0; i < 15; i++ {
		fmt.Println(i, thing)
		time.Sleep(500 * time.Millisecond)
	}
}

func Partie2() {
	go Count("Mouton(s)")
	go Count("Poisson(s)")
}
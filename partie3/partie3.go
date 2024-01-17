package partie3

import (
	"fmt"
	"time"
	"Sync"
)

func Count(thing string, time2 int) string {
	for i := 0; i < time2; i++ {
		fmt.Println(i, thing)
		time.Sleep(500 * time.Millisecond)
	}
}

func Partie3() {
	var wg sync.WaitGroup
	var data string
	var data2 string

	wg.Add(2)
	go func() {
		data = Count("Mouton(s)")
		wg.Done()
	}
	go func() {
		data2 = Count("Poisson(s)")
		wg.Done()
	}
}
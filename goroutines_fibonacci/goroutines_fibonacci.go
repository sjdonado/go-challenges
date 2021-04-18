package goroutines_fibonacci

import (
	"fmt"
	"time"
)

func fib(n float64, ch chan [2]float64) {
	x, y := 1.0, 1.0
	for i := 0; i < int(n); i++ {
		x, y = y, x+y
	}

	ch <- [2]float64{n, x}
}

func Test(size int) {
	start := time.Now()

	ch := make(chan [2]float64, size)

	for i := 1; i < size; i++ {
		go fib(float64(i), ch)
	}

	for i := 1; i < size; i++ {
		res := <-ch
		fmt.Printf("Fib(%v): %v\n", res[0], res[1])
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func fibInfinite(ch chan float64, quit chan bool) {
	x, y := 1.0, 1.0
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Done!")
			return
		}
	}
}

func InfiniteLoop() {
	start := time.Now()

	ch := make(chan float64)
	quit := make(chan bool)

	go fibInfinite(ch, quit)

	var command string

	fmt.Println("(Press enter to continue or type quit to exit)")

	for {
		fmt.Println(<-ch)
		fmt.Scanf("%s", &command)
		if command == "quit" {
			quit <- true
			break
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

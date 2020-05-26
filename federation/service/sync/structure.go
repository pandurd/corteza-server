package sync

import "github.com/davecgh/go-spew/spew"

func Doit() {
	ch := make(chan int)

	for i := 0; i < 100; i++ {
		go func() {
			ch <- i
		}()
	}

	// go func() {
	// 	ch <- 1
	// 	ch <- 2
	// }()

	b := <-ch

	spew.Dump("B", b)
}

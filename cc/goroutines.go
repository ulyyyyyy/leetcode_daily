package cc

import (
	"fmt"
	"sync"
)

var (
	number      = 100
	catChannel  = make(chan struct{}, 1)
	dogChannel  = make(chan struct{}, 1)
	fishChannel = make(chan struct{}, 1)
	wg          sync.WaitGroup
)

func do() {
	catChannel <- struct{}{}
	wg.Add(number)
	for i := 0; i < number; i++ {
		go catFunc()
		go dogFunc()
		go fishFunc()
	}
	wg.Wait()
	fmt.Println("over")
}

func catFunc() {
	select {
	case <-catChannel:
		fmt.Println("cat")
		dogChannel <- struct{}{}
	}
	return
}

func dogFunc() {
	select {
	case <-dogChannel:
		fmt.Println("dog")
		fishChannel <- struct{}{}
	}
	return
}

func fishFunc() {
	select {
	case <-fishChannel:
		fmt.Println("fish")
		catChannel <- struct{}{}
		wg.Done()
	}
	return
}

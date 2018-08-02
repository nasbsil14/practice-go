package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func say_async(s string, ch chan bool) {
	for i := 0; i < 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(s)
	}
	ch <- true
}

func pub(ch1, ch2 chan int) {
	x, y := 0, 1
	for {
		select {
		case ch1 <- x:
			x, y = y, x + y
		case <- ch2:
			fmt.Println("end")
			return
		default:
			fmt.Println("default")
			time.Sleep(10 * time.Millisecond)
		}
	}

}
func main() {
	//ch := make(chan bool)
	//go say_async("async", ch)
	//say("sync")
	//
	////同期（同期なしの場合say終了後にsay_asyncの完了を待たずmainが終了となるため
	//x := <- ch
	//fmt.Println(x)

	ch1 := make(chan int, 3)
	ch2 := make(chan int)

	go pub(ch1, ch2)
	for i := 0; i < 10; i++ {
		fmt.Println(<- ch1)
	}
	ch2 <- 0
}

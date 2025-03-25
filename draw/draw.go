package draw

import "fmt"

var ch chan string = make(chan string)

func Start() {
	for {
		draw()
		msg := <-ch
		fmt.Println(msg)
	}
}

func draw() {

}

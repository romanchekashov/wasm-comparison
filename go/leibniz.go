package main

import (
	"fmt"
	"syscall/js"
	"time"
)

func calcPi(this js.Value, p []js.Value) interface{} {
	start := time.Now()

	rounds := p[0].Int()

	x := 1.0
	pi := 1.0

	rounds += 2 // do this outside the loop

	for i := 2; i < rounds; i++ {
		x *= -1
		pi += x / float64(2*i-1)
	}

	pi *= 4

	fmt.Println(time.Since(start))
	fmt.Println(pi)

	return js.ValueOf(pi)
}

func arrayManipulations(this js.Value, p []js.Value) interface{} {
	run := func(num int) {
		var arr = make([]string, num)

		for i := range arr {
			arr[i] = "a"
		}

		for j := 0; j < num; j++ {
			arr = make([]string, num)

			for i := range arr {
				arr[i] = "a" + "1"
			}
		}

		for len(arr) > 0 {
			arr = arr[1:]
		}
	}

	rounds := p[0].Int()
	var count = rounds

	for count > 0 {
		run(rounds)
		count--
	}

	return js.ValueOf(0)
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("goCalcPi", js.FuncOf(calcPi))
	js.Global().Set("goArrayManipulations", js.FuncOf(arrayManipulations))

	<-c
}

package main

import (
	"fmt"
	"concurrency/src/libs/conc"
	)


type ArrayData struct {
	Value int
}

func main() {
	fmt.Println("Here")
	var data []conc.ArrayConcurrencyData

	for i := 0; i < 30; i++  {
		data = append(data, conc.ArrayConcurrencyData{Data : ArrayData{Value : i}})
	}

	fmt.Println(len(data))

	arc := conc.ArrayConcurrency{}
	arc.Set(data)

	worker := func(d conc.ArrayConcurrencyData) conc.ArrayConcurrencyData {
		return d
	}

	listener := func(d conc.ArrayConcurrencyData) {
		fmt.Println(d.Data)
	}

	arc.AddProcess("Process 1", worker, listener)
	arc.AddProcess("Process 2", worker, listener)
	arc.AddProcess("Process 3", worker, listener)
	arc.AddProcess("Process 4", worker, listener)
	arc.AddProcess("Process 5", worker, listener)
	arc.AddProcess("Process 6", worker, listener)
	arc.AddProcess("Process 7", worker, listener)
	arc.AddProcess("Process 8", worker, listener)

	arc.Run()

	var input string
	fmt.Scanln(&input)
}

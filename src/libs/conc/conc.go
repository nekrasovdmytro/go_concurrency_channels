package conc

import (
	"fmt"
)

type ArrayConcurrencyData struct {
	worked bool
	Data interface{}
}

func (ad *ArrayConcurrencyData) Worked() bool {
	return ad.worked
}

func (ad *ArrayConcurrencyData) SetWorked() {
	ad.worked = true
}

func (ad *ArrayConcurrencyData) SetData(data interface{}) {
	ad.Data = data
}

type Process struct {
	Name string
	Worker func(data ArrayConcurrencyData) ArrayConcurrencyData
	Listener func(data ArrayConcurrencyData)
}

type ArrayConcurrency struct {
	ArrayData []ArrayConcurrencyData
	Processes []Process
}

func (a *ArrayConcurrency) Set(array []ArrayConcurrencyData) {
	a.ArrayData = array
}

func (a *ArrayConcurrency) AddProcess(name string, worker func(data ArrayConcurrencyData) ArrayConcurrencyData, listener func(data ArrayConcurrencyData)) {
	a.Processes = append(a.Processes, Process{Name: name, Worker: worker, Listener: listener})
}

func (a *ArrayConcurrency) Run() {
	if a.ArrayData == nil {
		//error here
	}

	if a.Processes == nil {
		//error here
	}

	for i := 0; i < len(a.Processes) ; i++ {
		c := make(chan ArrayConcurrencyData)

		process := a.Processes[i]

		basicWorker := func (c chan ArrayConcurrencyData) {

			for j := 0; j < len(a.ArrayData); j++ {
				if a.ArrayData[j].Worked() {
					continue
				}

				a.ArrayData[j].SetWorked()

				c <- process.Worker(a.ArrayData[j])
			}
		}

		basicListener := func (c chan ArrayConcurrencyData) {
			for {
				data := <-c

				process.Listener(data)
			}
		}

		go basicWorker(c)
		go basicListener(c)


	}
}
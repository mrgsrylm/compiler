package main

import (
	"fmt"
	"sync"
)


func main() {
    var data interface{}
    var coba interface{}
	var wg sync.WaitGroup
	var mx sync.Mutex
    
	data = "[data1, data2, data3]"
    coba = "[coba1, coba2, coba3]"

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		mx.Lock()
		go printData(i, data, &wg)
		mx.Unlock()
		mx.Lock()
		go printCoba(i, coba, &wg)
		mx.Unlock()
	}

	wg.Wait()
}

func printData(index int, data interface{}, wg *sync.WaitGroup) {
	fmt.Printf("%s : %d\n",data, index)
	wg.Done()
}

func printCoba(index int, coba interface{}, wg *sync.WaitGroup) {
	fmt.Printf("%s : %d\n", coba, index)
	wg.Done()
}
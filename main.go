package main

import "fmt"
import "github.com/movebean/affinity"

func Loop() {
	var cpuSet int64
	affinity.GetAffinity(0, &cpuSet)
	fmt.Println(cpuSet)

	cpuSet = cpuSet & 0XFFFFFE
	affinity.SetAffinity(0, &cpuSet)

	affinity.GetAffinity(0, &cpuSet)
	fmt.Println(cpuSet)
	sum := 0
	for {
		sum++
		if sum%100000000 == 0 {
			sum = 0
		}
	}
}

func main() {
	var done chan int = make(chan int)

	//even two goroutine are running, they wolln't run on CPU0
	go Loop()
	go Loop()
	<-done
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const COUNT = 1000000
	var arr [6]int = [6]int{0, 0, 0, 0, 0, 0}

	t := time.Now().UnixNano()
	fmt.Println(t)
	rand.Seed(t)

	for i := 0; i < COUNT; i++ {
		s := rand.Intn(6)
		arr[s]++
	}

	for i := 0; i < len(arr); i++ {
		avg := (float32(arr[i]) / float32(COUNT)) * 100.0
		fmt.Printf("%d = %.2f％\n", i+1, avg)
	}
	fmt.Println(arr)
}

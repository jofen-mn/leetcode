package main

import (
	"fmt"
	"sync"
)

func main() {
	test111()
}

func test11() {
	const N = 10
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			mu.Lock()
			m[i] = i
			mu.Unlock()
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	println(len(m))
}

func test111() {
	const N = 10
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	wg.Add(N)

	ch := make(chan int, 10)

	for i := 0; i < N; i++ {
		ch <- i
	}

	for i := 0; i < N; i++ {
		go func() {
			k := <-ch
			fmt.Println(k)
			m[k] = k
			defer wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(len(m))
}

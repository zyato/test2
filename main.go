package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

func main() {
	pool, err := ants.NewPool(8000)
	if err != nil {
		return
	}
	defer pool.Release()
	var sum int64
	var wg sync.WaitGroup
	add := func(i int64) {
		defer wg.Done()
		time.Sleep(time.Second * 5)
		_ = atomic.AddInt64(&sum, i)
	}
	for i := 0; i <= 50000; i++ {
		wg.Add(1)
		v := i
		_ = pool.Submit(func() {
			add(int64(v))
		})
	}
	fmt.Println("running goroutines ", pool.Running())
	wg.Wait()
	fmt.Println("end goroutines ", pool.Running())
	fmt.Println("result = ", sum)

}
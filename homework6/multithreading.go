package homework6

import (
	"fmt"
	"sync"
	"time"
)

func TestSync(num int) {
	defer timeMeter(time.Now())
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			factorial(num)
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestAsync(num int) {
	defer timeMeter(time.Now())
	for i := 0; i < num; i++ {
		factorial(i)
	}
}

func TestBalance(num int, goRuts int) {
	defer timeMeter(time.Now())
	sem := make(chan struct{}, goRuts)
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			factorial(num)
			<-sem
			wg.Done()
		}()
	}
}

func factorial(num int) {
	x, y := 0, 1
	for i := 0; i < num; i++ {
		x, y = y, x+y
	}
}

func timeMeter(t time.Time) {
	fmt.Println(time.Since(t))
}

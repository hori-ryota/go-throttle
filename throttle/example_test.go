package throttle

import (
	"fmt"
	"time"
)

func Example() {
	c := time.Tick(100 * time.Millisecond)

	t := New()
	for i := 0; i < 9; i++ {
		t.Do("key", 301*time.Millisecond, func() {
			fmt.Println("do")
		})
		<-c
	}
	fmt.Println("waiting...")
	time.Sleep(400 * time.Millisecond)
	fmt.Println("finished")
	// Output:
	// do
	// do
	// do
	// waiting...
	// do
	// finished
}

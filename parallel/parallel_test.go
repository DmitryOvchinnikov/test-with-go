package parallel

import (
	"fmt"
	"testing"
	"time"
)

func TestSomething(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second)
}

func TestA(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second)
}

func TestB(t *testing.T) {
	fmt.Println("setup")
	defer fmt.Println("deferred teardown")
	t.Run("sub1", func(t *testing.T) {
		t.Parallel()
		// run sub1
		time.Sleep(time.Second)
		fmt.Println("sub1 done")
	})
	t.Run("sub2", func(t *testing.T) {
		t.Parallel()
		// run sub2
		time.Sleep(time.Second)
		fmt.Println("sub2 done")
	})
	fmt.Println("teardown")
}

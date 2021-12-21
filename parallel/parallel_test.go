package parallel

import (
	"fmt"
	"testing"
)

func TestGotcha(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("i=%d", i), func(t *testing.T) {
			t.Parallel()
			t.Logf("Testing with i=%d", i)
		})
	}
}

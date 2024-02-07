package cloudantv1_test

import (
	"slices"
	"sync"
	"testing"
)

func TestLoopvar(t *testing.T) {
	var wg sync.WaitGroup
	values := []int{1, 2, 3, 4, 5}
	pipe := make(chan int, len(values))
	for _, val := range values {
		wg.Add(1)
		go func() {
			defer wg.Done()
			t.Logf("%d ", val)
			pipe <- val
		}()
	}
	wg.Wait()
	close(pipe)

	acc := make([]int, 0)
	for val := range pipe {
		acc = append(acc, val)
	}
	slices.Sort(acc)
	if !slices.Equal(values, acc) {
		t.Errorf("Expected %d, got %d", values, acc)
	}
}

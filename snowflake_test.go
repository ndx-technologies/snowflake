package snowflake_test

import (
	"sync"
	"testing"

	"github.com/ndx-technologies/snowflake"
)

func BenchmarkGenerator(b *testing.B) {
	g := snowflake.NewGenerator(1)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			g.Next()
		}
	})
}

func TestGenerator(t *testing.T) {
	g := snowflake.NewGenerator(1)

	wg := sync.WaitGroup{}
	wg.Add(10)

	for range 10 {
		go func() {
			defer wg.Done()
			for range 1_000_000 {
				g.Next()
			}
		}()
	}

	wg.Wait()
}

package snowflake_test

import (
	"sync"
	"testing"
	"testing/synctest"

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
	synctest.Test(t, func(t *testing.T) {
		g := snowflake.NewGenerator(1)

		var wg sync.WaitGroup

		for range 10 {
			wg.Go(func() {
				for range 1_000_000 {
					g.Next()
				}
			})
		}

		wg.Wait()
	})
}

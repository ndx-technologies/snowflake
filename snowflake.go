package snowflake

import (
	"sync/atomic"
	"time"
)

const (
	timeMask      = (1 << 42) - 1
	generatorMask = (1 << 10) - 1
	sequenceMask  = (1 << 12) - 1
	timeShift     = 10 + 12
)

// Generator of snowflake ids.
// Lock free and thread safe.
// Loops back sequence when it reaches the limit within the same time and generator id.
// 42b time | 10b generator | 12b sequence
type Generator struct{ state uint64 }

func NewGenerator(generatorID uint16) *Generator {
	return &Generator{state: (uint64(generatorID) & generatorMask) << 12}
}

func (g *Generator) Next() uint64 {
	var state uint64

	for {
		t := uint64(time.Now().UnixMilli()) & timeMask

		current := atomic.LoadUint64(&g.state)

		currentTime := (current >> timeShift) & timeMask
		currentSeq := current & sequenceMask

		if t > currentTime || currentSeq == sequenceMask {
			state = (t << timeShift) | (current & (generatorMask << 12))
		} else {
			state = current + 1
		}

		if atomic.CompareAndSwapUint64(&g.state, current, state) {
			break
		}
	}

	return state
}

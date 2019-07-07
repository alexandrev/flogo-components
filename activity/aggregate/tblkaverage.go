package aggregate

import (
	"fmt"
	"sync"
	"time"
)

type TimeBlockAverage struct {
	windowSize   time.Duration
	values       [][]float64
	operations   []string
	windowMtx    *sync.Mutex
	startMtx     *sync.RWMutex
	windowActive bool
}

func init() {
	RegisterFactory("timeblock", NewTimeBlockAverage)
}

func (ta *TimeBlockAverage) Add(operation []string, value []float64) (bool, []float64) {

	ta.windowMtx.Lock()
	ta.values = append(ta.values, value)
	fmt.Sprintf("Total Values: %s", ta.values)
	ta.operations = operation
	ta.windowMtx.Unlock()

	if ta.startWindow() {
		time.Sleep(ta.windowSize * time.Millisecond)
		return true, ta.average()
	} else {
		return false, make([]float64, ta.windowSize)
	}
}

func (ta *TimeBlockAverage) average() []float64 {

	ta.windowMtx.Lock()

	var total = make([]float64, len(ta.values[0]))

	count := len(ta.values)
	fmt.Sprintf("Total Value Length: %d.\n", count)
	for i := 0; i < count; i++ {

		fmt.Sprintf("Value Length: %d.\n", ta.values[i])
		for j := 0; j < len(ta.values[i]); j++ {
			if ta.operations[j] == "min" {
				if total[j] > ta.values[i][j] {
					total[j] = ta.values[i][j]
				}
			} else if ta.operations[j] == "max" {
				if total[j] < ta.values[i][j] {
					total[j] = ta.values[i][j]
				}
			} else {
				total[j] += ta.values[i][j]
			}
		}

	}

	for j := 0; j < len(total); j++ {
		if ta.operations[j] == "avg" {
			total[j] = total[j] / float64(count)
		}
	}

	ta.resetWindow()

	return total
}

func (ta *TimeBlockAverage) startWindow() bool {

	ta.startMtx.RLock()

	if ta.windowActive {
		ta.startMtx.RUnlock()
		return false
	}
	ta.startMtx.RUnlock()

	ta.startMtx.Lock()
	defer ta.startMtx.Unlock()

	if !ta.windowActive {
		ta.windowActive = true
		return true
	}

	return false
}

func (ta *TimeBlockAverage) resetWindow() {
	ta.values = nil
	ta.windowMtx.Unlock()

	ta.startMtx.Lock()
	ta.windowActive = false
	ta.startMtx.Unlock()
}

func NewTimeBlockAverage(windowSize int) Aggregator {
	return &TimeBlockAverage{
		windowSize: time.Duration(windowSize),
		windowMtx:  &sync.Mutex{},
		startMtx:   &sync.RWMutex{},
	}
}

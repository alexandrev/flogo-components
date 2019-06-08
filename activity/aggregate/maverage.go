package aggregate

import "sync"

type MovingAverage struct {
	windowSize   int
	values       [][]float64
	operations   []string
	nextValueIdx int
	full         bool
	mutex        *sync.Mutex
}

func init() {
	RegisterFactory("moving", NewMovingAverage)
}

func (ma *MovingAverage) Add(operation []string, value []float64) (bool, []float64) {

	ma.mutex.Lock()
	defer ma.mutex.Unlock()

	//if ma.full && ma.nextValueIdx == 0 && ma.autoReset {
	//	ma.full = false
	//}

	ma.values[ma.nextValueIdx] = value
	ma.operations = operation

	ma.nextValueIdx = (ma.nextValueIdx + 1) % ma.windowSize

	if !ma.full && ma.nextValueIdx == 0 {
		ma.full = true
	}

	if ma.full {
		return true, ma.result()
	}

	return false, make([]float64, len(value))
}

func (ma *MovingAverage) result() []float64 {

	var count = ma.windowSize

	if !ma.full {
		if ma.nextValueIdx == 0 {
			return make([]float64, ma.windowSize)
		}

		count = ma.nextValueIdx
	}

	var total = make([]float64, len(ma.values[0]))
	for i := 0; i < count; i++ {

		for j := 0; j < len(ma.values[i]); j++ {
			if ma.operations[j] == "min" {
				if total[j] > ma.values[i][j] {
					total[j] = ma.values[i][j]
				}
			} else if ma.operations[j] == "max" {
				if total[j] < ma.values[i][j] {
					total[j] = ma.values[i][j]
				}
			} else {
				total[j] += ma.values[i][j]
			}
		}

	}

	for j := 0; j < len(total); j++ {
		if ma.operations[j] == "avg" {
			total[j] = total[j] / float64(count)
		}
	}

	return total
}

func NewMovingAverage(windowSize int) Aggregator {
	return &MovingAverage{
		windowSize: windowSize,
		values:     make([][]float64, windowSize),
		mutex:      &sync.Mutex{},
	}
}

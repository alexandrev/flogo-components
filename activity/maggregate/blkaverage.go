package aggregate

import "sync"

type BlockAverage struct {
	windowSize   int
	values       [][]float64
	operations   []string
	nextValueIdx int
	mutex        *sync.Mutex
}

func init() {

}

func (ba *BlockAverage) Get() (bool, []float64) {

	if len(ba.operations) >= 1 {
		return true, ba.average()
	}
	return false, make([]float64, 1)
}

func (ba *BlockAverage) Add(operation []string, value []float64) (bool, []float64) {

	ba.mutex.Lock()
	defer ba.mutex.Unlock()

	ba.values[ba.nextValueIdx] = value
	ba.operations = operation

	ba.nextValueIdx = ba.nextValueIdx + 1

	return false, make([]float64, len(value))
}

func (ba *BlockAverage) average() []float64 {

	var total = make([]float64, len(ba.values[0]))

	for i := 0; i < len(ba.values); i++ {

		for j := 0; j < len(ba.values[i]); j++ {
			if ba.operations[j] == "min" {
				if total[j] > ba.values[i][j] {
					total[j] = ba.values[i][j]
				}
			} else if ba.operations[j] == "max" {
				if total[j] < ba.values[i][j] {
					total[j] = ba.values[i][j]
				}
			} else {
				total[j] += ba.values[i][j]
			}
		}

	}

	for j := 0; j < len(total); j++ {
		if ba.operations[j] == "avg" {
			total[j] = total[j] / float64(len(ba.values))
		}
	}

	ba.nextValueIdx = 0

	return total
}

func NewBlockAverage() *BlockAverage {
	return &BlockAverage{
		values: make([][]float64, 1),
		mutex:  &sync.Mutex{}}
}

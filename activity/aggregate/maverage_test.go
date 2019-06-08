package aggregate

import (
	"fmt"
	"testing"
)

func TestMovingAverage_Add(t *testing.T) {

	agg := NewMovingAverage(5)

	report, avg := agg.Add([]string{"avg", "avg"}, []float64{10, 10})
	if report {
		t.Error("Window should not report yet")
	}
	report, avg = agg.Add([]string{"avg", "avg"}, []float64{20, 20})
	if report {
		t.Error("Window should not report yet")
	}
	report, avg = agg.Add([]string{"avg", "avg"}, []float64{30, 30})
	if report {
		t.Error("Window should not report yet")
	}
	report, avg = agg.Add([]string{"avg", "avg"}, []float64{40, 40})
	if report {
		t.Error("Window should not report yet")
	}
	report, avg = agg.Add([]string{"avg", "avg"}, []float64{50, 50})

	if avg[0] != 30.0 {
		t.Error("Average should be 30")
	}

	report, avg = agg.Add([]string{"avg", "avg"}, []float64{60, 60})

	fmt.Println("avg:", avg)

	if avg[0] != 40.0 {
		t.Error("Average should be 40")
	}
}

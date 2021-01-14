package chi_square

import (
	"testing"
)

func searchIndex(v []interface{}, target interface{}) int {
    for i, j := range v {
        if j == target {
            return i
        }
    }
    return 0
}

func observationFrequency(init_value []float64, targetArray []interface{}) []float64 {
	targetCopy := make([]interface{}, len(targetArray))
	copy(targetCopy, targetArray)
	for i := 0; i < 100; i++ {
		v := Shuffle(targetArray)
        init_value[searchIndex(v, targetCopy[0])]++
	}
	return init_value
}

func TestShuffle(t *testing.T) {
	observation_frequency := observationFrequency([]float64{0, 0, 0, 0, 0}, []interface{}{"a", "b", "c", "d", "e"})
	expected_frequency := []float64{20, 20, 20, 20, 20}
    if !ChiSquare(5, len(expected_frequency)-1, observation_frequency, expected_frequency) {
	    t.Error("not shuffled")
	}
}

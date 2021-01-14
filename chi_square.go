package chi_square

import (
	"math"
	"fmt"
)

func ChiSquare(significance_level int, freedom_degree int, observation_frequency []float64, expected_frequency []float64) bool {
	x := 0.0
	for i := 0; i <= freedom_degree; i++ {
		x += (math.Pow(observation_frequency[i] - expected_frequency[i], 2)) / expected_frequency[i]
	}
	if significance_level == 5 {
		fmt.Printf("x: %g, rejection_area: %g \n", x, fiveRejectionArea(freedom_degree))
		return jadge(x, fiveRejectionArea(freedom_degree))
	} else {
		return jadge(x, oneRejectionArea(freedom_degree))
	}
}

func jadge(x float64, rejection_area float64) bool {
	if x > rejection_area {
		return false
	} else {
		return true
	}
}

func fiveRejectionArea(freedom_degree int) float64 {
	five_rejection_area := []float64{
		3.84,
		5.99,
		7.81,
		9.49,
		11.07,
		12.6,
		14.1,
		15.5,
		16.9,
		18.3,
	}
	return five_rejection_area[freedom_degree-1]
}

func oneRejectionArea(freedom_degree int) float64{
	one_rejection_area := []float64{
		6.63,
		9.21,
		11.3,
		13.3,
		15.1,
		16.8,
		18.5,
		20.1,
		21.7,
		23.2,
	}
	return one_rejection_area[freedom_degree-1]
}

package rootfinding

import (
	"math"
	"testing"
)

func TestBrent(t *testing.T) {
	cases := []struct {
		f             func(float64) float64
		intervalStart float64
		intervalEnd   float64
		precision     int
		roots         []float64
		expectedErr   error
	}{
		{
			func(x float64) float64 {
				return (x + 3) * math.Pow(x-1, 2)
			},
			-100000, 100000, 4,
			[]float64{-3, 1},
			nil,
		},
		{
			func(x float64) float64 {
				return (x + 3) * math.Pow(x-2, 2)
			},
			-2, 1.5, 5,
			[]float64{-3, 2},
			ErrRootIsNotBracketed,
		},
		{
			func(x float64) float64 {
				return math.Pow(x, 4) - 2*math.Pow(x, 2) + 0.25
			},
			0, 1, 6,
			[]float64{0.366025403784438},
			nil,
		},
		{
			func(x float64) float64 {
				return -10 + math.Pow(x, 2)
			},
			-10000, 10000, 5,
			[]float64{-3.162278, 3.162278},
			nil,
		},
		{
			func(x float64) float64 {
				return -10 + 100*math.Pow(x, 2)
			},
			-1, 1, 5,
			[]float64{-0.316227, 0.316227},
			nil,
		},
	}
	for _, c := range cases {
		root, err := Brent(c.f, c.intervalStart, c.intervalEnd, c.precision)
		if err != c.expectedErr {
			t.Errorf("expected %v, got %v", c.expectedErr, err)
		}
		if err != nil {
			continue
		}
		matched := false
		i := 0
		acceptance := math.Pow10(-c.precision)
		for i < len(c.roots) && !matched {
			matched = c.roots[i]-acceptance <= root && root <= c.roots[i]+acceptance
			i++
		}
		if !matched {
			t.Errorf("expected roots are %v, got %f", c.roots, root)
		}
	}
}

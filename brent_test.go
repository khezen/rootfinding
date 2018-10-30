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
		precision     float64
		roots         []float64
	}{
		{
			func(x float64) float64 {
				return (x + 3) * math.Pow(x-1, 2)
			},
			-100000, 100000, 0.0001,
			[]float64{-3, 1},
		},
		{
			func(x float64) float64 {
				return math.Pow(x, 4) - 2*math.Pow(x, 2) + 0.25
			},
			0, 1, 0.000001,
			[]float64{0.366025403784438},
		},
	}
	for _, c := range cases {
		root, err := Brent(c.f, c.intervalStart, c.intervalEnd, c.precision)
		if err != nil {
			panic(err)
		}
		matched := false
		i := 0
		for i < len(c.roots) && !matched {
			matched = c.roots[i]-c.precision <= root && root <= c.roots[i]+c.precision
			i++
		}
		if !matched {
			t.Errorf("expected roots are %v, got %f", c.roots, root)
		}
	}
}

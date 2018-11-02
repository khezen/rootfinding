package rootfinding

import (
	"math"
)

// Brent - Brent's Method finds the root of the given quadratic function in [a,b]
// reference: https://en.wikipedia.org/wiki/Brent%27s_method
func Brent(f func(x float64) float64, a, b, precision float64) (r float64, err error) {
	var (
		delta            = EpsilonF64 * (b - a) // numerical tolerance
		fa               = f(a)
		fb               = f(b)
		c                = a
		fc               = f(c)
		s                float64
		fs               = f(s)
		d                float64
		wasBisectionUsed = true
		absBMinusC       float64
		absCMinusD       float64
		absSMinusB       float64
		tmp              float64
		// swap - a becomes b, b becomes a
		swap = func() {
			tmp = a
			a = b
			b = tmp
			tmp = fa
			fa = fb
			fb = tmp
		}
	)
	if fa*fb > 0 {
		return 0, ErrRootIsNotBracketed
	}
	if math.Abs(fa) < math.Abs(fb) {
		swap()
	}
	for fb != 0 && math.Abs(b-a) > precision {
		if fa != fc && fb != fc { // inverse quadratic interpolation
			s = (a*fb*fc)/((fa-fb)*(fa-fc)) + (b*fa*fc)/((fb-fa)*(fb-fc)) + (c*fa*fb)/((fc-fa)*(fc-fb))
		} else { // secant method
			s = b - fb*(b-a)/(fb-fa)
		}
		absBMinusC = math.Abs(b - c)
		absCMinusD = math.Abs(c - d)
		absSMinusB = math.Abs(s - b)
		switch {
		case s < (3*a+b)/4 || s > b,
			wasBisectionUsed && absSMinusB >= absBMinusC/2,
			!wasBisectionUsed && absSMinusB >= absCMinusD/2,
			wasBisectionUsed && absBMinusC < delta,
			!wasBisectionUsed && absCMinusD < delta: // bisection method
			s = (a + b) / 2
			wasBisectionUsed = true
			break
		default:
			wasBisectionUsed = false
			break

		}
		fs = f(s)
		d = c // d is first defined here; is not use in the first step above because wasBisectionUsed set to true
		c = b
		fc = fb
		if fa*fs < 0 {
			b = s
			fb = fs
		} else {
			a = s
			fa = fs
		}
		if math.Abs(fa) < math.Abs(fb) {
			swap()
		}
	}
	return s, nil
}

/*

*/

package unit

type Kilometers float64
type Miles float64

type Distance interface {
	Miles() float64
	Kilometers() float64

	ToMiles() Miles
	ToKilometers() Kilometers
}

func (m Miles) Miles() float64 {
	return float64(m)
}

func (m Miles) Kilometers() float64 {
	return float64(m * 0.6)
}

func (m Miles) ToMiles() Miles {
	return m
}

func (m Miles) ToKilometers() Kilometers {
	return Kilometers(m.Kilometers())
}

func (k Kilometers) Miles() float64 {
	return float64(k / 0.6)
}

func (k Kilometers) Kilometers() float64 {
	return float64(k)
}

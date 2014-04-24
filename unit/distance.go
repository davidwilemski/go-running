package unit

/*
Kilometers distance unit
*/
type Kilometers float64

/*
Miles distance unit
*/
type Miles float64

/*
A Distance interface implements conversion between various distance units
*/
type Distance interface {
	Miles() float64
	Kilometers() float64

	ToMiles() Miles
	ToKilometers() Kilometers
}

/*
Miles returns a distance value (in miles)
*/
func (m Miles) Miles() float64 {
	return float64(m)
}

/*
Kilometers returns a distance value (in kilometers)
*/
func (m Miles) Kilometers() float64 {
	return float64(m * 0.6)
}

/*
ToMiles converts the distance value to a Miles type
*/
func (m Miles) ToMiles() Miles {
	return m
}

/*
ToKilometers converts the distance value to a Kilometerss type
*/
func (m Miles) ToKilometers() Kilometers {
	return Kilometers(m.Kilometers())
}

/*
Miles returns a distance value (in miles)
*/
func (k Kilometers) Miles() float64 {
	return float64(k / 0.6)
}

/*
Kilometers returns a distance value (in kilometers)
*/
func (k Kilometers) Kilometers() float64 {
	return float64(k)
}

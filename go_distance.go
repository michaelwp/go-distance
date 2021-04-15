package go_distance

import "math"

func (l LatLon) count() Miles {
	var theta = l.LatStart - l.LatEnd

	var dist = math.Sin(Degree(l.LatStart).ToRadians().Float64())*
		math.Sin(Degree(l.LatEnd).ToRadians().Float64()) +
		math.Cos(Degree(l.LatStart).ToRadians().Float64())*
			math.Cos(Degree(l.LatEnd).ToRadians().Float64())*
			math.Cos(Degree(theta).ToRadians().Float64())

	dist = math.Acos(dist)
	dist = Radian(dist).ToDegrees().Float64()

	var miles = dist * 60 * 1.1515

	return Miles(miles)
}

func (deg Degree) ToRadians() Radian {
	return Radian(float64(deg) * (math.Pi / 180.0))
}

func (rad Radian) ToDegrees() Degree {
	return Degree(float64(rad) * (180.0 / math.Pi))
}

func (rad Radian) Float64() float64 {
	return float64(rad)
}

func (deg Degree) Float64() float64 {
	return float64(deg)
}

func (m Miles) Kilometre() Kilometre {
	return Kilometre(m * 1.609344)
}

func (k Kilometre) Miles() Miles {
	return Miles(k * 0.621371)
}
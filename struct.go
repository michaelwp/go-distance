package go_distance

type (
	LatLon struct {
		LatStart float64 `json:"lat_start"`
		LonStart float64 `json:"lon_start"`
		LatEnd   float64 `json:"lat_end"`
		LonEnd   float64 `json:"lon_end"`
	}

	Degree    float64
	Radian    float64
	Miles     float64
	Kilometre float64
)

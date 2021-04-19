package go_distance

type (
	LatLon struct {
		OriginLatitude       float64 `json:"origin_latitude"`
		OriginLongitude      float64 `json:"origin_longitude"`
		DestinationLatitude  float64 `json:"destination_latitude"`
		DestinationLongitude float64 `json:"destination_longitude"`
	}

	Degree    float64
	Radian    float64
	Miles     float64
	Kilometre float64
)

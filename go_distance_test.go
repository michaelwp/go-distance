package go_distance

import "testing"

func TestDegree_Float64(t *testing.T) {
	tests := []struct {
		name string
		deg  Degree
		want float64
	}{
		{"Degree To Float64", 10, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.deg.ToFloat64(); got != tt.want {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDegree_ToRadians(t *testing.T) {
	tests := []struct {
		name string
		deg  Degree
		want Radian
	}{
		{"Degree To Radians", 10, 0.17453292519943295},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.deg.ToRadians(); got != tt.want {
				t.Errorf("ToRadians() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKilometre_Miles(t *testing.T) {
	tests := []struct {
		name string
		k    Kilometre
		want Miles
	}{
		{"Kilometre To Miles", 10, 6.21371},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.ToMiles(); got != tt.want {
				t.Errorf("Miles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLatLon_count(t *testing.T) {
	type fields struct {
		OriginLatitude       float64
		OriginLongitude      float64
		DestinationLatitude  float64
		DestinationLongitude float64
	}
	tests := []struct {
		name   string
		fields fields
		want   Miles
	}{
		{"Distance Count", fields{
			OriginLatitude:       -6.2973856,
			OriginLongitude:      106.6388177,
			DestinationLatitude:  -6.3027637,
			DestinationLongitude: 106.6410986,
		}, 0.40323836641733635},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LatLon{
				OriginLatitude:       tt.fields.OriginLatitude,
				OriginLongitude:      tt.fields.OriginLongitude,
				DestinationLatitude:  tt.fields.DestinationLatitude,
				DestinationLongitude: tt.fields.DestinationLongitude,
			}
			if got := l.Count(); got != tt.want {
				t.Errorf("count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiles_Kilometre(t *testing.T) {
	tests := []struct {
		name string
		m    Miles
		want Kilometre
	}{
		{"Miles To Kilometres",
			0.40323836641733635, 0.6489492455635418},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.ToKilometre(); got != tt.want {
				t.Errorf("Kilometre() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRadian_Float64(t *testing.T) {
	tests := []struct {
		name string
		rad  Radian
		want float64
	}{
		{"Radian To Float64", 10, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rad.ToFloat64(); got != tt.want {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRadian_ToDegrees(t *testing.T) {
	tests := []struct {
		name string
		rad  Radian
		want Degree
	}{
		{"Radian To Degrees", 10, 572.9577951308232},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rad.ToDegrees(); got != tt.want {
				t.Errorf("ToDegrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

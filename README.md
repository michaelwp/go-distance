# go-distance v1.2.0

Ini adalah package bahasa pemrograman go sederhana untuk menghitung jarak antara dua titik latitude dan longitude.

### func Count()
```go
func (l LatLon) Count() Miles
```
### func ToKilometre()
```go
func (m Miles) ToKilometre() Kilometre
```
### func ToMiles()
```go
func (k Kilometre) ToMiles() Miles
```

### Instalasi
`go get -u github.com/michaelwp/go-distance`

### Contoh penggunaan

```go 
package main

import (
	"fmt"
	go_distance "github.com/michaelwp/go-distance"
)

func main()  {
	var LatLon = go_distance.LatLon{
		LatStart: -6.2973856,
		LonStart: 106.6388177,
		LatEnd:   -6.3027637,
		LonEnd:   106.6410986,
	}

	fmt.Println(LatLon.Count(), "mil")
	fmt.Println(LatLon.Count().ToKilometre(), "km")
}
```
#### Hasil
```text
0.5238991320272929 mil
0.8431339247333317 km
```
# go-distance v1.0.0

Ini adalah package bahasa pemrograman go sederhana untuk menghitung jarak antara dua titik latitude dan longitude.

### func Count()
```go
func (l LatLon) Count() Miles
```
### func Kilometre()
```go
func (m Miles) Kilometre() Kilometre
```
### func Ts()
```go
func (f JFile) Ts(key string) (message interface{})
```

### Instalasi
`go get github.com/michaelwp/linggo` atau
`go get github.com/michaelwp/linggo@v1.1.1`

### Contoh penggunaan
A. Buat Json file untuk menampung bahasa yang digunakan
- id.json (bahasa Indonesia)
```json 
{
  "hello world": "halo dunia",
  "welcome": "selamat datang",
  "please": "silahkan"
}
```
- en.json (bahasa Inggris)
```json 
{
  "hello world": "hello world"
}
```
- fr.json (bahasa Prancis)
```json 
{
  "hello world": "Bonjour le monde"
}
```
- en.json (bahasa China)
```json 
{
  "hello world": "你好，世界"
}
```
B. memanggil file json setiap kali proses translate

```go 
package main

import (
    "github.com/michaelwp/linggo"  // import linggo package
    "log"
)

func main(){
    // contoh translate ke dalam bahasa China
    msg, err := linggo.Tr("ch", "hello world")
    // handle error (disarankan untuk menghandle error dengan cara yang lebih baik)
    if err != nil {
        log.Fatal(err)
    }
    // menampilkan hasil ke layar
    log.Println(msg)
}
```
#### Hasil
```text
你好，世界
```

C. memanggil file json hanya sekali diawal.

```go 
package main

import (
    "github.com/michaelwp/linggo"  // import linggo package
    "log"
)

// load json file (bahasa indonesia/ id.json)
var f, _ = linggo.Set("id")

func main() {
    welcome := f.Ts("welcome")
    please := f.Ts("please")

    log.Println(welcome)
    log.Println(please)
}
```
#### Hasil
```text
selamat datang
silahkan
```
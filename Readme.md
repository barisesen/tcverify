# TC Kimlik Numarası Kontrolü ve Doğrulaması

Golang kullanılarak TC kontrolü ve doğrulama işlemini yapmayı sağlayan bir pakettir.

## Kurulum

```go
    go get github.com/barisesen/tcverify 
```

## Kullanım

### Algoritmik olarak kontrol etmek

```go
    package main

    import (
        "fmt"
        "github.com/barisesen/tcverify"
    )

    func main() {
        resp, err := tcverify.Validate("xxxxxxxxxxx")
        fmt.Println(resp, err)
        // true <nil>
        // false xxxxxxxxxxx tc numarası algoritmik olarak doğrulanamadı.
    }
```
TC kimlik numarası algoritmik doğrulama için [Hakan Ersu](https://github.com/hakanersu/)'nun oluşturduğu [tcvalidate](https://github.com/hakanersu/tcvalidate) paketi kullanılmıştır.

-----------------

### Api desteği ile doğrulama

```go
     package main

    import (
        "fmt"
        "github.com/barisesen/tcverify"
    )

    func main() {
        tc := "xxxxxxxxxxx"
        isim := "BARIŞ"
        soyisim := "ESEN"
        dogumTarihi := "1996"

        resp, err := tcverify.Check(tc, isim, soyisim, dogumTarihi)
        fmt.Println(resp, err)
        // true <nil>
        // false Bu bilgileri ait vatandaşlık doğrulanamadı.
    }   
````

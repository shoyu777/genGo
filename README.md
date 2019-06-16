# shoyu777/genGo
This package is to convert Anno Domini(YYYY) to Japanese Era called Gengo.

## How to use
```bash
$ go get github.com/shoyu777/genGo
```
Please call New function with any years with as int YYYY. It returns new struct with result.

The input year has to be 645 or above. Otherwise error returned.

```go
import "genGo"
```

```go
g, err := genGo.New(2019)
if err != nil {
  fmt.Println(err)
}
fmt.Println(g.Romaji)
//output -> REIWA GAN NEN

fmt.Println(g.Kanji)
//output -> 令和 元年

fmt.Println(g.Romaji)
//output -> れいわ がんねん

```

```go
g, err := genGo.New(760)
if err != nil {
  fmt.Println(err)
}
fmt.Println(g.Kanji)
//output -> 天平宝字 4年
```

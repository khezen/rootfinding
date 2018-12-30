[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/rootfinding)
[![Build Status](http://img.shields.io/travis/khezen/rootfinding.svg?style=flat-square)](https://travis-ci.org/khezen/rootfinding) [![codecov](https://img.shields.io/codecov/c/github/khezen/rootfinding/master.svg?style=flat-square)](https://codecov.io/gh/khezen/rootfinding)
[![Go Report Card](https://goreportcard.com/badge/github.com/khezen/rootfinding?style=flat-square)](https://goreportcard.com/report/github.com/khezen/rootfinding)

# *rootfinding*

`github.com/khezen/rootfinding`

* Brent's Method

## Example

```golang
package main

import(
    "fmt"
    "github.com/khezen/rootfinding"
)

func f(x float64) float64 {
	return math.Pow(x, 4) - 2*math.Pow(x, 2) + 0.25
}

const(
    intervalStart = -100
    intervalEnd = 100
    precision = 6
)
func main(){
    root, err := rootfinding.Brent(f, intervalStart, intervalEnd, precision)
    if err != nil {
        panic(err)
    }
    fmt.Println(root)
}		
```

```bash
0.366025403784438
```
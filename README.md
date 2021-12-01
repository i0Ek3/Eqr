# eqr

Everything is qrcode.

## Overview

```console
Input Link  --->   Process(Link)   --->   Generate & Show Image
```

## Usage

### 1. Run Command `./eqr`

### 2. Build & Run

```Shell
$ go build eqr.go 

// or

$ go run eqr.go

// or

make clean ; make install 
```

Before that, run command `go mod tidy` first and make sure you have installed `imgcat`(brew install imgcat).

### 3. Import

```Go
// main.go

package main

import (
    "github.com/i0Ek3/eqr"
)

func main() {
    // Offer your own string field link and output,
    // you will get an image generate by Process().
    eqr.Process(link, output)
}

```

## TODO

- [ ] cmd arguments support


## Credit

- [ship2/go-qrcode](https://github.com/skip2/go-qrcode)

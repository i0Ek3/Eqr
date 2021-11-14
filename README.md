# Eqr

Everything is qrcode.

## Overview

```console
Input Link  --->   Process(Link)   --->   Generate & Show Image
```

## Usage

### 1. Run command `./eqr`

### 2. Build and run

```Shell
$ go build eqr.go 

// or

$ go run eqr.go

// or

make clean ; make install 
```

Before that, run command `go mod tidy` first and make sure you have installed `imgcat`(brew install imgcat).

### 3. Import It Into Go File

```Go
// main.go

package main

import (
    eqr "github.com/i0Ek3/Eqr"
)

func main() {
    // offer your own string field link and output
    // you will get an image of your link and a png
    // file named output
    eqr.Process(link, output)
}

```


## Credit

- [ship2/go-qrcode](https://github.com/skip2/go-qrcode)

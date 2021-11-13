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

make
```

Before that, run command `go mod tidy` first and make sure you have installed `imgcat`(brew install imgcat).


## Credit

- [ship2/go-qrcode](https://github.com/skip2/go-qrcode)

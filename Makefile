all: clean install

install: eqr.go
	go build . 

test:
	go test -v

clean: 
	rm eqr

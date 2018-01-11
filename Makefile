all: gen
	
gen:
	c-for-go -out .. cairo.yml

clean:
	rm -f cgo_helpers.go cgo_helpers.h cgo_helpers.c
	rm -f doc.go types.go const.go
	rm -f cairo.go

test:
	go build

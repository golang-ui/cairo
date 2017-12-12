all:
	c-for-go cairo.yml

clean:
	rm -f cairo/cgo_helpers.go cairo/cgo_helpers.h cairo/cgo_helpers.c
	rm -f cairo/doc.go cairo/types.go cairo/const.go
	rm -f cairo/cairo.go

test:
	cd cairo && go build

install:
	cd cairo && go install

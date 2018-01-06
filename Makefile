all: gen-cairo gen-pugl

gen-cairo:
	c-for-go cairo.yml
gen-pugl:
	c-for-go pugl.yml

clean: clean-cairo clean-pugl

clean-cairo:
	rm -f cairo/cgo_helpers.go cairo/cgo_helpers.h cairo/cgo_helpers.c
	rm -f cairo/doc.go cairo/types.go cairo/const.go
	rm -f cairo/cairo.go

clean-pugl:
	rm -f pugl/cgo_helpers.go pugl/cgo_helpers.h pugl/cgo_helpers.c
	rm -f pugl/doc.go pugl/types.go pugl/const.go
	rm -f pugl/pugl.go

test: test-cairo test-pugl

test-cairo:
	cd cairo && go build

test-pugl:
	cd pugl && go build

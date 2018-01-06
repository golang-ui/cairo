// +build darwin
package cairo

// #cgo CFLAGS: -DCAIRO_HAS_QUARTZ_SURFACE=1
// #include <cairo-quartz.h>
import "C"

import "unsafe"

func QuartzSurfaceCreate(format Format, width uint32, height uint32) *Surface {
	cformat, _ := (C.cairo_format_t)(format), cgoAllocsUnknown
	cwidth, _ := (C.uint)(width), cgoAllocsUnknown
	cheight, _ := (C.uint)(height), cgoAllocsUnknown
	__ret := C.cairo_quartz_surface_create(cformat, cwidth, cheight)
	__v := *(**Surface)(unsafe.Pointer(&__ret))
	return __v
}

func QuartzSurfaceCreateForCGContext(ctx C.CGContextRef, width uint32, height uint32) *Surface {
	cwidth, _ := (C.uint)(width), cgoAllocsUnknown
	cheight, _ := (C.uint)(height), cgoAllocsUnknown
	__ret := C.cairo_quartz_surface_create_for_cg_context(ctx, cwidth, cheight)
	__v := *(**Surface)(unsafe.Pointer(&__ret))
	return __v
}

// +build darwin
// +build external
package cairo

/*
#cgo darwin pkg-config: cairo
#cgo darwin CFLAGS: -DCAIRO_HAS_QUARTZ_FONT
#cgo darwin CFLAGS: -DCAIRO_HAS_QUARTZ_IMAGE_SURFACE
#cgo darwin CFLAGS: -DCAIRO_HAS_QUARTZ_SURFACE

#include "include/cairo-quartz.h"
*/
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

// QuartzSurfaceGetCgContext function as declared in cairo/cairo-quartz.h:57
func QuartzSurfaceGetCgContext(surface *Surface) C.CGContextRef {
	csurface, _ := (*C.cairo_surface_t)(unsafe.Pointer(surface)), cgoAllocsUnknown
	__ret := C.cairo_quartz_surface_get_cg_context(csurface)
	__v := (C.CGContextRef)(__ret)
	return __v
}

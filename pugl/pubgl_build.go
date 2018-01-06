// +build darwin

package pugl

/*
#cgo CFLAGS: -DPUGL_HAVE_GL -DPUGL_HAVE_CAIRO -DPUGL_VERSION="0.2.0" -DPUGL_INTERNAL
#cgo LDFLAGS: -fvisibility=hidden -lm

#cgo darwin,!ios CFLAGS: -Wno-deprecated-declarations
#cgo darwin,!ios LDFLAGS: -framework Cocoa -framework OpenGL
#cgo darwin,!ios pkg-config: cairo

#cgo windows LDFLAGS: -lopengl32 -lgdi32 -luser32
#cgo linux LDFLAGS: -lX11 -lGL
*/
import "C"

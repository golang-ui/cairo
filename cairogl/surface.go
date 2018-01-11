package cairogl

import (
	"sync"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/golang-ui/cairo"
)

type Surface struct {
	mux      *sync.RWMutex
	platform *platformSurface
	context  *cairo.Cairo
	width    int
	height   int
	tex_id   uint32
}

func NewSurface(width, height int) *Surface {
	s := &Surface{
		mux:      new(sync.RWMutex),
		width:    width,
		height:   height,
		platform: newPlatformSurface(width, height),
	}
	s.context = cairo.Create(s.platform.Surface)

	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.TEXTURE_RECTANGLE_ARB)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(-1.0, 1.0, -1.0, 1.0, -1.0, 1.0)

	gl.DeleteTextures(1, &s.tex_id)
	gl.GenTextures(1, &s.tex_id)
	gl.BindTexture(gl.TEXTURE_RECTANGLE_ARB, s.tex_id)
	gl.TexEnvi(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.DECAL)

	return s
}

func (s *Surface) Surface() *cairo.Surface {
	s.mux.RLock()
	ss := s.platform.Surface
	s.mux.RUnlock()
	return ss
}

func (s *Surface) Context() *cairo.Cairo {
	s.mux.RLock()
	cr := s.context
	s.mux.RUnlock()
	return cr
}

func (s *Surface) Destroy() {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.context != nil {
		cairo.Destroy(s.context)
		s.context = nil
	}
	if s.platform != nil {
		s.platform.Destroy()
	}
}

func (s *Surface) Size() (int, int) {
	s.mux.RLock()
	w, h := s.width, s.height
	s.mux.RUnlock()
	return w, h
}

func (s *Surface) Width() int {
	s.mux.RLock()
	w := s.width
	s.mux.RUnlock()
	return w
}

func (s *Surface) Height() int {
	s.mux.RLock()
	h := s.height
	s.mux.RUnlock()
	return h
}

func (s *Surface) Update(width, height int) {
	s.mux.RLock()
	w, h := s.width, s.height
	s.mux.RUnlock()

	if width != w || height != h {
		s.Destroy()
		s.mux.Lock()
		tmp := s.mux
		*s = *NewSurface(width, height)
		s.mux = tmp
		s.mux.Unlock()
	}
}

func (s *Surface) Draw() {
	s.mux.RLock()

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Viewport(0, 0, int32(s.width), int32(s.height))
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.PushMatrix()
	gl.Enable(gl.TEXTURE_RECTANGLE_ARB)
	gl.Enable(gl.TEXTURE_2D)

	gl.TexImage2D(gl.TEXTURE_RECTANGLE_ARB, 0, gl.RGBA8,
		int32(s.width), int32(s.height), 0,
		gl.BGRA, gl.UNSIGNED_BYTE, s.platform.Buffer)

	gl.Begin(gl.QUADS)
	gl.TexCoord2f(0.0, float32(s.height))
	gl.Vertex2f(-1.0, -1.0)

	gl.TexCoord2f(float32(s.width), float32(s.height))
	gl.Vertex2f(1.0, -1.0)

	gl.TexCoord2f(float32(s.width), 0.0)
	gl.Vertex2f(1.0, 1.0)

	gl.TexCoord2f(0.0, 0.0)
	gl.Vertex2f(-1.0, 1.0)
	gl.End()

	gl.Disable(gl.TEXTURE_2D)
	gl.Disable(gl.TEXTURE_RECTANGLE_ARB)
	gl.PopMatrix()

	s.mux.RUnlock()
}

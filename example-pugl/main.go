package main

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/golang-ui/cairo-go/cairo"
	"github.com/golang-ui/cairo-go/pugl"
	"github.com/xlab/closer"
)

func init() {
	runtime.LockOSThread()
}

type canvas struct {
	mux    sync.RWMutex
	cr     *cairo.Cairo
	width  int
	height int
}

func (c *canvas) Resize(view *pugl.View) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cr = (*cairo.Cairo)(pugl.GetContext(view))
	var width, height int32
	pugl.GetSize(view, &width, &height)
	c.width = int(width)
	c.height = int(height)
}

func (c *canvas) Update(fn func(cr *cairo.Cairo, width, height int)) {
	c.mux.Lock()
	fn(c.cr, c.width, c.height)
	c.mux.Unlock()
}

func main() {
	view := pugl.Init(nil, nil)
	if view == nil {
		closer.Fatalln("failed to init PUGL")
	}
	cv := &canvas{}

	exitC := make(chan struct{}, 1)
	doneC := make(chan struct{}, 1)
	closer.Bind(func() {
		close(exitC)
		<-doneC
	})

	pugl.InitWindowSize(view, 800, 600)
	pugl.InitResizable(view, 1)
	pugl.InitContextType(view, pugl.ContextCairoGL)
	pugl.SetEventFunc(view, func(view *pugl.View, event *pugl.Event) {
		log.Println("event:", event.Type())
		switch event.Type() {
		case pugl.Expose:
			cv.Resize(view)
			cv.Update(func(cr *cairo.Cairo, width, height int) {
				gfxMain(cr)
			})
		case pugl.Close:
			close(exitC)
		}
	})
	pugl.CreateWindow(view, "Pugl Test")
	pugl.ShowWindow(view)

	fpsTicker := time.NewTicker(time.Second / 30)
	for {
		select {
		case <-exitC:
			log.Println("closing")
			pugl.Destroy(view)
			log.Println("destroyed ")
			fpsTicker.Stop()
			log.Println("stopped")
			close(doneC)
			log.Println("doneC")
			return
		case <-fpsTicker.C:
			log.Println("processing events (stuck on OS X)")
			pugl.ProcessEvents(view)
			log.Println("end events processing")
			cv.Update(func(cr *cairo.Cairo, width, height int) {
				gfxMain(cr)
			})
		}
	}
}

func gfxMain(cr *cairo.Cairo) {
	cairo.SetSourceRgba(cr, 0, 0, 0, 1)
	cairo.SetLineWidth(cr, 5)
	cairo.MakeRectangle(cr, 100, 100, 200, 200)
	cairo.Stroke(cr)
}

package main

import (
	"log"

	"github.com/golang-ui/cairo"
)

func main() {
	surface := cairo.QuartzSurfaceCreate(cairo.FormatArgb32, 800, 600)
	defer cairo.SurfaceDestroy(surface)
	cr := cairo.Create(surface)
	defer cairo.Destroy(cr)

	cairo.SetSourceRgb(cr, 1.0, 1.0, 1.0)
	cairo.Paint(cr)
	setScene(cr)
	cairo.SurfaceWriteToPng(surface, "out.png")
	log.Println(cairo.GetStatus(cr), cairo.StatusToString(cairo.GetStatus(cr)))
}

const PI = 3.1415926

func setScene(cr *cairo.Cairo) {
	cairo.Save(cr)
	defer cairo.Restore(cr)

	cairo.SetLineWidth(cr, 6.0)
	cairo.SetSourceRgba(cr, 0, 0, 0, 1)
	cairo.NewPath(cr)
	cairo.MoveTo(cr, 50, 50)
	cairo.CurveTo(cr, 50, 50, 300, 250, 50, 500)
	cairo.ClosePath(cr)
	cairo.Stroke(cr)

	xc := 500.0
	yc := 200.0
	radius := 100.0
	angle1 := 45.0 * (PI / 180.0)
	angle2 := 180.0 * (PI / 180.0)

	cairo.SetLineWidth(cr, 3.0)
	cairo.SetSourceRgba(cr, 0, 0, 0, 1)
	cairo.Arc(cr, xc, yc, radius, angle1, angle2)
	cairo.Stroke(cr)

	cairo.SetSourceRgba(cr, 1, 0.2, 0.2, 0.6)
	cairo.SetLineWidth(cr, 6.0)

	cairo.Arc(cr, xc, yc, 10.0, 0, 2*PI)
	cairo.Fill(cr)

	cairo.Arc(cr, xc, yc, radius, angle1, angle2)
	cairo.LineTo(cr, xc, yc)
	cairo.Arc(cr, xc, yc, radius, angle1, angle2)
	cairo.LineTo(cr, xc, yc)
	cairo.Stroke(cr)
}

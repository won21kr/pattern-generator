package body

import (
	"math"
	"github.com/tobyjsullivan/dxf/drawing"
)

type Waist struct {
	Circumference float64
	Height float64
}

func (w *Waist) DrawDXF(d *drawing.Drawing) error {
	_, err := d.Circle(0.0, 0.0, w.Height, w.Circumference / (2.0 * math.Pi))
	return err
}


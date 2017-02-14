package geometry

type Text struct {
	Content  string
	Position *Point
	Rotation *Angle
}

func (t *Text) Move(x, y float64) *Text {
	var out Text = *t
	out.Position = t.Position.Move(x, y)
	return &out
}
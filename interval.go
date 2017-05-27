package interval

type Intr struct {
	x float64
	y float64
}

func NewIntr(x, y float64) *Intr {
	return &Intr{x, y}
}

func (i *Intr) Add(j *Intr) *Intr {
	return &Intr{
		x: i.x + j.x,
		y: i.y + j.y,
	}
}

func (i *Intr) Sub(j *Intr) *Intr {
	return &Intr{
		x: i.x - j.y,
		y: i.y - j.x,
	}
}

func (i *Intr) Mul(j *Intr) *Intr {
	muls := [4]float64{i.x * j.x, i.x * j.y, i.y * j.x, i.y * j.y}
	return &Intr{
		x: min(min(muls[0], muls[1]), min(muls[2], muls[3])),
		y: max(max(muls[0], muls[1]), max(muls[2], muls[3])),
	}
}

func (i *Intr) Div(j *Intr) *Intr {
	inv := &Intr{1 / j.x, 1 / j.y}
	return i.Mul(inv)
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	s.start.x += int(s.a)
	s.start.y += int(s.a)

	return s.start
}

func (s Square) Area() uint {
	res := s.a * s.a

	return res
}

func (s Square) Perimeter() uint {
	res := s.a * 4

	return res
}

package main

import (
	"errors"
	"fmt"
)

type Point struct {
	X int
	Y int
}

type Triangle struct {
	A *Point
	B *Point
	C *Point
}

func main() {
	Do()
}

func Do() error {
	t := Triangle{
		A: &Point{X: 1, Y: 2},
		B: &Point{X: 2, Y: 3},
		C: &Point{X: -3, Y: 6},
	}

	// Написать код между...
	tp := t
	// ...этими комментариями

	Scale(tp)

	if fmt.Sprintf("%p", &tp) == fmt.Sprintf("%p", &t) {
		return errors.New("same object")
	}

	if tp.A.X == t.A.X {
		return errors.New("scale failed")
	}

	return nil
}

func Scale(t Triangle) {
	(*t.A).X = (*t.A).X * 2
	(*t.A).Y = (*t.A).Y * 2
	(*t.B).X = (*t.B).X * 2
	(*t.B).Y = (*t.B).Y * 2
	(*t.C).X = (*t.C).X * 2
	(*t.C).Y = (*t.C).Y * 2
}
package sys

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

type System struct {
	Nx    int
	Nu    int
	Time  float64
	State *mat.VecDense
}

func NewSystem(nx, nu int) *System {
	return &System{
		Nx:    nx,
		Nu:    nu,
		Time:  0.0,
		State: mat.NewVecDense(nx, nil),
	}
}

func (s *System) String() string {
	return fmt.Sprintf(
		"Nx: %d\nNu: %d\nTime: %v\nState:\n%v",
		s.Nx, s.Nu, s.Time, mat.Formatted(s.State),
	)
}

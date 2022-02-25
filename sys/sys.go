package sys

import (
    "gonum.org/v1/gonum/mat"
)

type System struct {
	Time  float64
	State *mat.VecDense
	Nx    int
	Nu    int
}

func NewSystem(nx, nu int) *System {
    return &System{
        Time: 0.0,
        State: mat.NewVecDense(nx, nil),
        Nx: nx,
        Nu: nu,
    }
}

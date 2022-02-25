package sim

import (
	"gonum.org/v1/gonum/mat"
)

type Simulation interface {
	Update(u *mat.VecDense, dt float64) *mat.VecDense
    State() *mat.VecDense
    Time() float64
}

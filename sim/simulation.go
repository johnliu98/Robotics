package sim

import (
	"gonum.org/v1/gonum/mat"
)

type Simulation interface {
	Update(u mat.Vector, dt float64) mat.Vector
    State() mat.Vector
    Time() float64
}

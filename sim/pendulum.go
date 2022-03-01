package sim

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

const g float64 = 9.81
const l float64 = 0.5
const m float64 = 0.15
const eta float64 = 0.1

type Pendulum struct {
	nx    int
	nu    int
	time  float64
	state *mat.VecDense
}

func NewPendulum() *Pendulum {
	return &Pendulum{
		nx:    2,
		nu:    1,
		time:  0.0,
		state: mat.NewVecDense(2, nil),
	}
}

func (p *Pendulum) Update(u mat.Vector, dt float64) mat.Vector {
	if n, _ := u.Dims(); n != p.nu {
		fmt.Println("Warning: pen: argument does not match input dimensions")
	}

	p.time += dt

	p.state = mat.NewVecDense(
		p.nx,
		[]float64{
			p.state.AtVec(0) + dt*p.state.AtVec(1),
			p.state.AtVec(1) +
				-dt*g/l*math.Sin(p.state.AtVec(0)) +
				-dt*eta/m/math.Pow(l, 2)*p.state.AtVec(1) +
				+dt/m/math.Pow(l, 2)*u.AtVec(0),
		},
	)

	return p.State()
}

func (p *Pendulum) Time() float64 {
	return p.time
}

func (p *Pendulum) State() mat.Vector {
	s := &mat.VecDense{}
	s.CloneFromVec(p.state)
	return s
}

func (p *Pendulum) String() string {
	return fmt.Sprintf(
		"Nx: %d\nNu: %d\nTime: %v\nState: %v",
		p.nx, p.nu, p.time, p.state,
	)
}

package sys

import (
	"fmt"
	"log"

	"github.com/jinzhu/copier"
	"gonum.org/v1/gonum/mat"
)

type System struct {
	nx    int
	nu    int
	time  float64
	state *mat.VecDense
}

func NewSystem(nx, nu int) *System {
	return &System{
		nx:    nx,
		nu:    nu,
		time:  0.0,
		state: mat.NewVecDense(nx, nil),
	}
}

func (s *System) Nx() int {
	var nx int
	err := copier.Copy(nx, s.nx)
	if err != nil {
		log.Fatal(err)
	}
	return nx
}

func (s *System) Nu() int {
	var nu int
	err := copier.Copy(nu, s.nu)
	if err != nil {
		log.Fatal(err)
	}
	return nu
}

func (s *System) Time() float64 {
	var t float64
	err := copier.Copy(t, s.time)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func (s *System) State() *mat.VecDense {
	st := &mat.VecDense{}
	st.CloneFromVec(s.state)
	return st
}

func (s *System) String() string {
	return fmt.Sprintf(
		"Nx: %d\nNu: %d\nTime: %v\nState:\n%v",
		s.nx, s.nu, s.time, mat.Formatted(s.state),
	)
}

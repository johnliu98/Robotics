package main

import (
	"fmt"
	"github.com/johnliu98/robotics/sim/pen"
	"github.com/johnliu98/robotics/sim"
    "gonum.org/v1/gonum/mat"
)

func main() {
	p := pen.NewPendulum()
	fmt.Println(p)

    var s sim.Simulation = pen.NewPendulum()

    u := mat.NewVecDense(1, []float64{1})
    fmt.Println(s.Update(u, 0.02))
    fmt.Println(s.Update(u, 0.02))
}

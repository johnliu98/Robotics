package main

import (
	"fmt"
    "image/color"
	"time"

	"github.com/johnliu98/robotics/sim"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const T float64 = 10

func main() {
	var s sim.Simulation = sim.NewPendulum()

	u := mat.NewVecDense(1, []float64{0.7})
    dt := 0.02
    N := int(T / dt)

    thData := make(plotter.XYs, N)
    dthData := make(plotter.XYs, N)

    // simulate pendulum
	start := time.Now()
	for i := 0; i < int(N); i++ {
        v := s.Update(u, dt)
		fmt.Println(i, v)

        thData[i].X = s.Time()
        thData[i].Y = v.AtVec(0)
        dthData[i].X = s.Time()
        dthData[i].Y = v.AtVec(1)
	}

    // print computation time
	d := float64(time.Since(start)) / float64(N)
    var unit string
    switch i := d; {
    case i > 1e9: unit = "s "; d /= 1e9
    case i > 1e6: unit = "ms"; d /= 1e6
    case i > 1e3: unit = "µs"; d /= 1e3
    case i > 1e0: unit = "ns";
    }
	fmt.Printf("Computation time: %.2f %s", d, unit)

    // plot data
    p := plot.New()
    p.Title.Text = "Pendulum"
    p.X.Label.Text = "time [s]"
    p.Add(plotter.NewGrid())

    thLine, err := plotter.NewLine(thData)
    if err != nil {
        panic(err)
    }
    thLine.Color = color.RGBA{31, 119, 180, 255}
	thLine.LineStyle.Width = vg.Points(2)

    dthLine, err := plotter.NewLine(dthData)
    if err != nil {
        panic(err)
    }
    dthLine.Color = color.RGBA{255, 127, 14, 255}
	dthLine.LineStyle.Width = vg.Points(2)

    p.Add(thLine, dthLine)
	p.Save(30*vg.Centimeter, 30*vg.Centimeter, "states.png")
}

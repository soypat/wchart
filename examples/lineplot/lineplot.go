package main

import (
	"fmt"
	"syscall/js"

	"github.com/soypat/wchart"
)

func main() {
	wchart.AddScript("https://cdnjs.cloudflare.com/ajax/libs/Chart.js/3.7.1/chart.js", "Chart")
	wchart.Init()
	canvas := js.Global().Get("document").Call("createElement", "canvas")
	js.Global().Get("document").Get("body").Call("appendChild", canvas)
	lp := wchart.NewLineFromXYers(canvas, &Func{
		F: func(x float64) float64 { return x * x },
		N: 10,
	})
	cfg := lp.GetConfig()
	js.Global().Set("chart", lp)
	js.Global().Set("cfg", cfg)
	fmt.Println("added chart!", cfg.Get("data").Get("datasets").Get("length"))
	select {}
}

type Func struct {
	F            func(x float64) float64
	Xi, XfMinus1 float64
	N            int
}

func (f *Func) XY(i int) (x, y float64) {
	x = f.Xi + float64(i)*(f.XfMinus1+1-f.Xi)/float64(f.N)
	return x, f.F(x)
}

func (f *Func) Len() int { return f.N }

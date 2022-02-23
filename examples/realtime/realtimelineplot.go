package main

import (
	"fmt"
	"math"
	"syscall/js"
	"time"

	"github.com/soypat/wchart"
)

func main() {
	wchart.AddScript("https://cdnjs.cloudflare.com/ajax/libs/Chart.js/3.7.1/chart.js", "Chart")
	wchart.Init()
	canvas := js.Global().Get("document").Call("createElement", "canvas")
	js.Global().Get("document").Get("body").Call("appendChild", canvas)
	rtlp := wchart.NewRealtimeLinePlot(canvas, []wchart.Dataset{{Label: "data"}}, nil)
	go func() {
		x := 0.0
		for {
			x += .1
			y := math.Sin(x)
			label := fmt.Sprintf("%g", x)
			rtlp.AddData(label, []float64{y})
			rtlp.Update()
			time.Sleep(time.Second * 1000 / 1618) // the "golden" frequency?
		}
	}()
	select {}
}

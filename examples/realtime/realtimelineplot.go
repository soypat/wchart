package main

import (
	"fmt"
	"math"
	"syscall/js"
	"time"

	"github.com/soypat/gwasm"
	"github.com/soypat/wchart"
)

func main() {
	gwasm.AddScript("https://cdnjs.cloudflare.com/ajax/libs/Chart.js/3.7.1/chart.js", "Chart", time.Second)
	wchart.Init()
	canvas := js.Global().Get("document").Call("createElement", "canvas")
	js.Global().Get("document").Get("body").Call("appendChild", canvas)

	rtlp := wchart.NewRealtimeLinePlot(canvas, nil, []wchart.Dataset{
		{Label: "data", BorderColor: js.ValueOf("red")},
	})
	js.Global().Set("plot", rtlp.Value)
	go func() {
		x := 0.0
		for {
			x += math.Pi / 10
			y := math.Sin(x)
			label := fmt.Sprintf("%g", x)
			rtlp.AddData(label, []float64{y})
			rtlp.Update()
			time.Sleep(time.Second * 1000 / 1618) // the "golden" frequency?
			if x > 2*math.Pi {
				rtlp.Decimate(3)
				x = 0
			}
		}
	}()
	select {}
}

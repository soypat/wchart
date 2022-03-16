package main

import (
	"fmt"
	"strconv"
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
	opts := &wchart.Options{}
	pgt := wchart.PluginTitle{
		Text:    "Hello world!",
		Display: true,
		Align:   "center",
	}
	opts.AddPlugins(pgt)

	// Generate data
	var labels []string
	data := js.Global().Get("Array").New()
	for i := 0; i < 10; i++ {
		data.Call("push", i)
		labels = append(labels, strconv.Itoa(i))
	}
	config := &wchart.Config{
		Type:    "line",
		Options: opts,
		Data: wchart.Data{
			Labels: labels,
			Datasets: []wchart.Dataset{
				{Label: "The data", Data: data},
			},
		},
	}
	lp := wchart.NewChart(canvas, config)
	cfg := lp.GetConfig()
	js.Global().Set("chart", lp.Value)
	js.Global().Set("cfg", cfg.Value)
	fmt.Println("added chart!", cfg.Get("data").Get("datasets").Get("length"))
	select {}
}

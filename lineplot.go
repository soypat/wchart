package wchart

import "syscall/js"

type LinePlot struct {
	Chart
}

// Documentation says this should work.
// I guess fuck me for trusting a js library's documentation.
func NewLineFromXYers(ctx js.Value, xyers ...XYer) LinePlot {
	cfg := Config{
		Type: "line",
		Data: Data{
			Datasets: make([]Dataset, len(xyers)),
		},
	}

	for ds, xyer := range xyers {
		cfg.Data.Datasets[ds].Data = js.Global().Get("Array").New()
		for i := 0; i < xyer.Len(); i++ {
			x, y := xyer.XY(i)
			xyobj := js.Global().Get("Object").New()
			xyobj.Set("x", x)
			xyobj.Set("y", y)
			cfg.Data.Datasets[ds].Data.Call("push", xyobj)
		}
	}
	return LinePlot{
		Chart: NewChart(ctx, &cfg),
	}
}

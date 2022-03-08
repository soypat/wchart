package wchart

import "syscall/js"

type LinePlot struct {
	Chart
}

// NewLineFromXYers plots several XYers on a plot. If an xyer implements
// the LabelledXYer interface then the dataset is labelled.
func NewLineFromXYers(ctx js.Value, xyers ...XYer) LinePlot {
	if len(xyers) == 0 {
		panic("expected at least one XYer")
	}
	n := xyers[0].Len()

	for _, xyer := range xyers {
		if n != xyer.Len() {
			panic("all XYers must be of same length")
		}
	}
	cfg := Config{
		Type: "line",
		Data: Data{
			Datasets: make([]Dataset, len(xyers)),
		},
	}

	labels := make([]string, n)
	for ds, xyer := range xyers {
		data := js.Global().Get("Array").New()
		for i := 0; i < n; i++ {
			x, y := xyer.XY(i)
			if ds == 0 {
				labels[i] = js.Global().Get("Number").New(x).Call("toString").String()
			}
			data.Call("push", y)
		}
		cfg.Data.Datasets[ds].Data = data
		cfg.Data.Labels = labels
		if lxyerL, ok := xyer.(LabelledXYer); ok {
			cfg.Data.Datasets[ds].Label = lxyerL.Label()
		}
	}
	return LinePlot{
		Chart: NewChart(ctx, &cfg),
	}
}

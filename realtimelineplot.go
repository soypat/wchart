package wchart

import "syscall/js"

type RealtimeLinePlot struct {
	Chart
}

func NewRealtimeLinePlot(ctx js.Value, xLabels []string, datasets []Dataset) RealtimeLinePlot {
	switch {
	case len(datasets) == 0:
		panic("cannot create plot from 0 length datasets")
	}
	if xLabels == nil {
		xLabels = []string{}
	}
	cfg := Config{
		Type: "line",
		Data: Data{
			Datasets: datasets,
		},
		Labels: xLabels,
	}

	return RealtimeLinePlot{
		Chart: NewChart(ctx, &cfg),
	}
}

// Call Update after calling AddData to update graph. data must be of length
func (r RealtimeLinePlot) AddData(xLabel string, data []float64) {
	r.GetConfig().AppendFloats(xLabel, data)
}

func (r RealtimeLinePlot) DiscardDatasetData() {
	r.GetConfig().Get("data").Set("labels", js.Global().Get("Array").New())
	dsets := r.GetConfig().Datasets()
	for _, dset := range dsets {
		dset.Set("data", js.Global().Get("Array").New())
	}
}

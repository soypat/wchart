package wchart

import "syscall/js"

type Chart struct {
	js.Value
}

func NewChart(ctx js.Value, config *Config) Chart {
	if !ctx.Truthy() {
		panic("ctx argument to NewChart should be a DOM element node.")
	}
	if chart.IsUndefined() {
		panic("Chart not defined. Did you add the script to DOM?")
	}
	return Chart{
		Value: chart.Get("Chart").New(ctx, objectify(config)),
	}
}

func (c Chart) Update() {
	c.Call("update")
}

func (c Chart) GetConfig() ConfigHandle {
	return ConfigHandle{
		Value: c.Get("config"),
	}
}

type ConfigHandle struct {
	js.Value
}

func (ch ConfigHandle) AppendFloat(label string, data []float64) {
	datasets := ch.Datasets()
	if len(data) != len(datasets) {
		panic("length of incoming data must match length of existing datasets")
	}
	ch.Get("data").Get("labels").Call("push", label)
	for i, d := range datasets {
		d.AppendFloat(data[i])
	}
}

func (ch ConfigHandle) Datasets() []DatasetHandle {
	D := ch.Get("data").Get("datasets")
	N := D.Get("length").Int()
	datasets := make([]DatasetHandle, N)
	for i := 0; i < N; i++ {
		datasets[i] = DatasetHandle{
			Value: D.Index(i),
		}
	}
	return datasets
}

type DatasetHandle struct {
	js.Value
}

func (dh DatasetHandle) AppendFloat(f float64) {
	dh.Get("data").Call("push", f)
}

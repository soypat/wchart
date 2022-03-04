package wchart

import (
	"image/color"
	"syscall/js"

	"github.com/soypat/gwasm"
)

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

// IsHiddenMeta is null until user clicks on legend, then it is
// set to true if user has hidden dataset, or false if dataset is visible.
func (c Chart) IsHiddenMeta(datasetIdx int) js.Value {
	return c.Call("getDatasetMeta", datasetIdx).Get("hidden")
}

// IsHidden returns true if dataset hidden property is set to true.
func (c Chart) IsHidden(datasetIdx int) bool {
	hidden := c.GetConfig().Datasets()[datasetIdx].Get("hidden")
	if hidden.IsNull() || hidden.IsUndefined() {
		return false
	}
	return hidden.Bool()
}

type ConfigHandle struct {
	js.Value
}

func (ch ConfigHandle) AppendFloats(label string, data []float64) {
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

// ClearChartData destroys dataset `data` fields and config data.labels and sets them to empty array
func (ch ConfigHandle) ClearChartData() {
	ch.Get("data").Set("labels", js.Global().Get("Array").New())
	dsets := ch.Datasets()
	for _, dset := range dsets {
		dset.Set("data", js.Global().Get("Array").New())
	}
}

// SetAnimation enables or disables animations for the chart.
// Usually used for performance reasons or when animations are distracting.
func (ch ConfigHandle) SetAnimation(enableAnimation bool) {
	setOption(ch.Value, "animation", enableAnimation)
}

// SetSpanGaps If you have a lot of data points, it can be more performant to enable spanGaps.
// This disables segmentation of the line, which can be an unneeded step.
func (ch ConfigHandle) SetSpanGaps(enableSpanGaps bool) {
	setOption(ch.Value, "spanGaps", enableSpanGaps)
}

// SetShowLine enables/disables line drawing for the chart.
func (ch ConfigHandle) SetShowLine(enableShowLine bool) {
	setOption(ch.Value, "showLine", enableShowLine)
}

// SetPointRadius sets the point radius of the chart. Set to zero for performance gains.
func (ch ConfigHandle) SetPointRadius(pointRadius float64) {
	opts := ch.Get("options")
	if opts.IsUndefined() {
		opts = js.Global().Get("Object").New()
		ch.Set("options", opts)
	}
	elems := opts.Get("elements")
	if elems.IsUndefined() {
		elems = js.Global().Get("Object").New()
		setOption(ch.Value, "elements", elems)
	}
	point := elems.Get("point")
	if point.IsUndefined() {
		point = js.Global().Get("Object").New()
		elems.Set("point", point)
	}
	point.Set("radius", pointRadius)
}

type DatasetHandle struct {
	js.Value
}

func (dh DatasetHandle) AppendFloat(f float64) {
	data := dh.Get("data")
	if !data.Truthy() {
		dh.Set("data", js.Global().Get("Array").New())
		data = dh.Get("data")
	}
	data.Call("push", f)
}

func (dh DatasetHandle) SetBackgroundColor(c color.Color) {
	dh.Set("backgroundColor", gwasm.JSColor(c))
}

// SetColor sets font color
func (dh DatasetHandle) SetColor(c color.Color) {
	dh.Set("color", gwasm.JSColor(c))
}

// SetBorderColor sets outline color of the chart data.
func (dh DatasetHandle) SetBorderColor(c color.Color) {
	dh.Set("borderColor", gwasm.JSColor(c))
}

// SetAnimation enables or disables animations for the dataset.
// Usually used for performance reasons or when animations are distracting.
func (dh DatasetHandle) SetAnimation(enableAnimation bool) {
	setOption(dh.Value, "animation", enableAnimation)
}

// SetSpanGaps if the dataset has a lot of data points, it can be more performant to enable spanGaps.
// This disables segmentation of the line, which can be an unneeded step.
func (dh DatasetHandle) SetSpanGaps(enableSpanGaps bool) {
	setOption(dh.Value, "spanGaps", enableSpanGaps)
}

// SetShowLine disables/enables dataset line drawing.
// func (dh DatasetHandle) SetShowLine(enableShowLine bool) {
// 	setOption(dh.Value, "showLine", enableShowLine) // Not working.
// }

// SetPointRadius sets the point radius of dataset. Set to zero for performance gains.
func (dh DatasetHandle) SetPointRadius(pointRadius float64) {
	dh.Set("pointRadius", pointRadius)
}

// SetHidden hides/shows dataset data in chart.
func (dh DatasetHandle) SetHidden(hidden bool) {
	dh.Set("hidden", hidden)
}

func setOption(jsv js.Value, prop string, x interface{}) {
	opts := jsv.Get("options")
	if opts.IsUndefined() {
		opts = js.Global().Get("Object").New()
		jsv.Set("options", opts)
	}
	opts.Set(prop, x)
}

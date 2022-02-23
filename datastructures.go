package wchart

import "syscall/js"

type XYer interface {
	XY(int) (x, y float64)
	Len() int
}

type LabelledXYer interface {
	XYer
	Label() string
}

type Config struct {
	Type   string   `js:"type"`
	Data   Data     `js:"data"`
	Labels []string `js:"labels"`
}

type Data struct {
	Datasets []Dataset `js:"datasets"`
}

type Dataset struct {
	Data js.Value `js:"data"`
	// The label for the dataset which appears in the legend and tooltips.
	Label string `js:"label"`
	// How to clip relative to chartArea. Positive value allows overflow, negative value clips that many pixels inside chartArea. 0 = clip at chartArea.
	// Clipping can also be configured per side: clip: {left: 5, top: false, right: -2, bottom: 0}
	Clip Clip `js:"clip"`
	// The drawing order of dataset. Also affects order for stacking, tooltip and legend.
	Order int `js:"order"`
	// The ID of the group to which this dataset belongs to (when stacked, each group will be a separate stack). Defaults to dataset `type`.
	Stack   string `js:"stack"`
	Parsing bool   `js:"parsing"`
	Hidden  bool   `js:"hidden"`
}

type Clip struct {
	Left   float64 `js:"left"`
	Top    float64 `js:"top"`
	Right  float64 `js:"right"`
	Bottom float64 `js:"bottom"`
}

type labeller string

func (l labeller) Label() string { return string(l) }

func LabelXYer(label string, xyer XYer) LabelledXYer {
	type lxyer struct {
		XYer
		labeller
	}
	return lxyer{
		XYer:     xyer,
		labeller: labeller(label),
	}
}

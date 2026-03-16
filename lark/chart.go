package lark

// Chart renders VChart-based visualizations (tag: "chart").
type Chart struct {
	ElementID   string `json:"element_id,omitempty"`
	Margin      string `json:"margin,omitempty"`
	AspectRatio string `json:"aspect_ratio,omitempty"`
	ColorTheme  string `json:"color_theme,omitempty"`
	Height      string `json:"height,omitempty"`
	ChartSpec   any    `json:"chart_spec,omitempty"`
}

func (*Chart) cardTag() string { return "chart" }

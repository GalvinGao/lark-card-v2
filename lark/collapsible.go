package lark

// CollapsiblePanel is an expandable/collapsible container (tag: "collapsible_panel").
type CollapsiblePanel struct {
	ElementID         string        `json:"element_id,omitempty"`
	Expanded          bool          `json:"expanded,omitempty"`
	Direction         string        `json:"direction,omitempty"`
	Margin            string        `json:"margin,omitempty"`
	Padding           string        `json:"padding,omitempty"`
	VerticalSpacing   string        `json:"vertical_spacing,omitempty"`
	HorizontalSpacing string        `json:"horizontal_spacing,omitempty"`
	VerticalAlign     string        `json:"vertical_align,omitempty"`
	HorizontalAlign   string        `json:"horizontal_align,omitempty"`
	BackgroundColor   string        `json:"background_color,omitempty"`
	Header            *PanelHeader  `json:"header,omitempty"`
	Border            *BorderConfig `json:"border,omitempty"`
	Elements          Elements      `json:"elements,omitempty"`
}

func (*CollapsiblePanel) cardTag() string { return "collapsible_panel" }

// PanelHeader is the header section of a CollapsiblePanel.
type PanelHeader struct {
	Title             *Text  `json:"title,omitempty"`
	BackgroundColor   string `json:"background_color,omitempty"`
	VerticalAlign     string `json:"vertical_align,omitempty"`
	Padding           string `json:"padding,omitempty"`
	Width             string `json:"width,omitempty"`
	Icon              *Icon  `json:"icon,omitempty"`
	IconPosition      string `json:"icon_position,omitempty"`
	IconExpandedAngle int    `json:"icon_expanded_angle,omitempty"`
}

// BorderConfig defines border appearance.
type BorderConfig struct {
	Color        string `json:"color,omitempty"`
	CornerRadius string `json:"corner_radius,omitempty"`
}

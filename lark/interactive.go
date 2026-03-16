package lark

// InteractiveContainer provides unified interactive behavior (tag: "interactive_container").
type InteractiveContainer struct {
	ElementID         string     `json:"element_id,omitempty"`
	Width             string     `json:"width,omitempty"`
	Height            string     `json:"height,omitempty"`
	Margin            string     `json:"margin,omitempty"`
	Direction         string     `json:"direction,omitempty"`
	HorizontalSpacing string     `json:"horizontal_spacing,omitempty"`
	VerticalSpacing   string     `json:"vertical_spacing,omitempty"`
	HorizontalAlign   string     `json:"horizontal_align,omitempty"`
	VerticalAlign     string     `json:"vertical_align,omitempty"`
	BackgroundStyle   string     `json:"background_style,omitempty"`
	HasBorder         bool       `json:"has_border,omitempty"`
	BorderColor       string     `json:"border_color,omitempty"`
	CornerRadius      string     `json:"corner_radius,omitempty"`
	Padding           string     `json:"padding,omitempty"`
	Behaviors         []Behavior `json:"behaviors,omitempty"`
	HoverTips         *Text      `json:"hover_tips,omitempty"`
	Disabled          bool       `json:"disabled,omitempty"`
	DisabledTips      *Text      `json:"disabled_tips,omitempty"`
	Confirm           *Confirm   `json:"confirm,omitempty"`
	Elements          Elements   `json:"elements,omitempty"`
}

func (*InteractiveContainer) cardTag() string { return "interactive_container" }

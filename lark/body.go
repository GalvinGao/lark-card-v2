package lark

// Body is the main content area of a card.
type Body struct {
	Direction         string   `json:"direction,omitempty"`
	Padding           string   `json:"padding,omitempty"`
	HorizontalSpacing string   `json:"horizontal_spacing,omitempty"`
	VerticalSpacing   string   `json:"vertical_spacing,omitempty"`
	HorizontalAlign   string   `json:"horizontal_align,omitempty"`
	VerticalAlign     string   `json:"vertical_align,omitempty"`
	Elements          Elements `json:"elements,omitempty"`
}

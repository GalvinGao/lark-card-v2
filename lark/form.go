package lark

// Form is a form container component (tag: "form").
type Form struct {
	ElementID         string   `json:"element_id,omitempty"`
	Name              string   `json:"name,omitempty"`
	Direction         string   `json:"direction,omitempty"`
	Margin            string   `json:"margin,omitempty"`
	Padding           string   `json:"padding,omitempty"`
	HorizontalSpacing string   `json:"horizontal_spacing,omitempty"`
	VerticalSpacing   string   `json:"vertical_spacing,omitempty"`
	HorizontalAlign   string   `json:"horizontal_align,omitempty"`
	VerticalAlign     string   `json:"vertical_align,omitempty"`
	Elements          Elements `json:"elements,omitempty"`
}

func (*Form) cardTag() string { return "form" }

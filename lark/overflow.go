package lark

// Overflow is a collapsible button group (tag: "overflow").
type Overflow struct {
	ElementID string           `json:"element_id,omitempty"`
	Margin    string           `json:"margin,omitempty"`
	Width     string           `json:"width,omitempty"`
	Options   []OverflowOption `json:"options,omitempty"`
	Value     map[string]any   `json:"value,omitempty"`
	Confirm   *Confirm         `json:"confirm,omitempty"`
}

func (*Overflow) cardTag() string { return "overflow" }

// OverflowOption is an option within an Overflow button group.
type OverflowOption struct {
	Text     *Text     `json:"text,omitempty"`
	MultiURL *MultiURL `json:"multi_url,omitempty"`
	Value    string    `json:"value,omitempty"`
}

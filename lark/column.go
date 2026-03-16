package lark

// ColumnSet arranges columns horizontally (tag: "column_set").
type ColumnSet struct {
	ElementID         string   `json:"element_id,omitempty"`
	Margin            string   `json:"margin,omitempty"`
	HorizontalSpacing string   `json:"horizontal_spacing,omitempty"`
	HorizontalAlign   string   `json:"horizontal_align,omitempty"`
	FlexMode          string   `json:"flex_mode,omitempty"`
	BackgroundStyle   string   `json:"background_style,omitempty"`
	Action            *Action  `json:"action,omitempty"`
	Columns           []Column `json:"columns,omitempty"`
}

func (*ColumnSet) cardTag() string { return "column_set" }

// Column is a single column within a ColumnSet (tag: "column").
type Column struct {
	ElementID         string   `json:"element_id,omitempty"`
	BackgroundStyle   string   `json:"background_style,omitempty"`
	Width             string   `json:"width,omitempty"`
	Weight            int      `json:"weight,omitempty"`
	HorizontalSpacing string   `json:"horizontal_spacing,omitempty"`
	HorizontalAlign   string   `json:"horizontal_align,omitempty"`
	VerticalAlign     string   `json:"vertical_align,omitempty"`
	VerticalSpacing   string   `json:"vertical_spacing,omitempty"`
	Direction         string   `json:"direction,omitempty"`
	Padding           string   `json:"padding,omitempty"`
	Margin            string   `json:"margin,omitempty"`
	Action            *Action  `json:"action,omitempty"`
	Elements          Elements `json:"elements,omitempty"`
}

// Action wraps a multi_url for column/column_set click actions.
type Action struct {
	MultiURL *MultiURL `json:"multi_url,omitempty"`
}

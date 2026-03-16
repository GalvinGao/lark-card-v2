package lark

// Divider is a horizontal rule component (tag: "hr").
type Divider struct {
	ElementID string `json:"element_id,omitempty"`
	Margin    string `json:"margin,omitempty"`
}

func (*Divider) cardTag() string { return "hr" }

package lark

// Input is a text input component (tag: "input").
type Input struct {
	ElementID     string     `json:"element_id,omitempty"`
	Margin        string     `json:"margin,omitempty"`
	Name          string     `json:"name,omitempty"`
	Required      bool       `json:"required,omitempty"`
	Placeholder   *Text      `json:"placeholder,omitempty"`
	DefaultValue  string     `json:"default_value,omitempty"`
	Disabled      bool       `json:"disabled,omitempty"`
	DisabledTips  *Text      `json:"disabled_tips,omitempty"`
	Width         string     `json:"width,omitempty"`
	Behaviors     []Behavior `json:"behaviors,omitempty"`
	MaxLength     int        `json:"max_length,omitempty"`
	InputType     string     `json:"input_type,omitempty"`
	Rows          int        `json:"rows,omitempty"`
	AutoResize    bool       `json:"auto_resize,omitempty"`
	MaxRows       int        `json:"max_rows,omitempty"`
	ShowIcon      *bool      `json:"show_icon,omitempty"`
	Label         *Text      `json:"label,omitempty"`
	LabelPosition string     `json:"label_position,omitempty"`
	Value         any        `json:"value,omitempty"`
	Confirm       *Confirm   `json:"confirm,omitempty"`
}

func (*Input) cardTag() string { return "input" }

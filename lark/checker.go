package lark

// Checker is a task-checking interactive component (tag: "checker").
type Checker struct {
	ElementID        string             `json:"element_id,omitempty"`
	Name             string             `json:"name,omitempty"`
	Checked          bool               `json:"checked,omitempty"`
	Text             *Text              `json:"text,omitempty"`
	OverallCheckable *bool              `json:"overall_checkable,omitempty"`
	ButtonArea       *CheckerButtonArea `json:"button_area,omitempty"`
	CheckedStyle     *CheckedStyle      `json:"checked_style,omitempty"`
	Margin           string             `json:"margin,omitempty"`
	Padding          string             `json:"padding,omitempty"`
	Confirm          *Confirm           `json:"confirm,omitempty"`
	Behaviors        []Behavior         `json:"behaviors,omitempty"`
	HoverTips        *Text              `json:"hover_tips,omitempty"`
	Disabled         bool               `json:"disabled,omitempty"`
	DisabledTips     *Text              `json:"disabled_tips,omitempty"`
}

func (*Checker) cardTag() string { return "checker" }

// CheckerButtonArea configures buttons within a Checker.
type CheckerButtonArea struct {
	PCDisplayRule string   `json:"pc_display_rule,omitempty"`
	Buttons       []Button `json:"buttons,omitempty"`
}

// CheckedStyle configures the visual style when a Checker is checked.
type CheckedStyle struct {
	ShowStrikethrough bool    `json:"show_strikethrough,omitempty"`
	Opacity           float64 `json:"opacity,omitempty"`
}

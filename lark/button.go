package lark

// Button is an interactive button component (tag: "button").
type Button struct {
	ElementID      string         `json:"element_id,omitempty"`
	Margin         string         `json:"margin,omitempty"`
	Type           string         `json:"type,omitempty"`
	Size           string         `json:"size,omitempty"`
	Width          string         `json:"width,omitempty"`
	Text           *Text          `json:"text,omitempty"`
	Icon           *Icon          `json:"icon,omitempty"`
	HoverTips      *Text          `json:"hover_tips,omitempty"`
	Disabled       bool           `json:"disabled,omitempty"`
	DisabledTips   *Text          `json:"disabled_tips,omitempty"`
	Confirm        *Confirm       `json:"confirm,omitempty"`
	Behaviors      []Behavior     `json:"behaviors,omitempty"`
	Name           string         `json:"name,omitempty"`
	FormActionType string         `json:"form_action_type,omitempty"`
	ActionType     string         `json:"action_type,omitempty"`
	URL            string         `json:"url,omitempty"`
	MultiURL       *MultiURL      `json:"multi_url,omitempty"`
	Value          map[string]any `json:"value,omitempty"`
}

func (*Button) cardTag() string { return "button" }

// Confirm configures a secondary confirmation dialog.
type Confirm struct {
	Title *Text `json:"title,omitempty"`
	Text  *Text `json:"text,omitempty"`
}

// Behavior defines an interaction behavior (open_url or callback).
type Behavior struct {
	Type       string `json:"type,omitempty"`
	DefaultURL string `json:"default_url,omitempty"`
	AndroidURL string `json:"android_url,omitempty"`
	IOSURL     string `json:"ios_url,omitempty"`
	PCURL      string `json:"pc_url,omitempty"`
	Value      any    `json:"value,omitempty"`
}

// MultiURL provides platform-specific jump links.
type MultiURL struct {
	URL        string `json:"url,omitempty"`
	AndroidURL string `json:"android_url,omitempty"`
	IOSURL     string `json:"ios_url,omitempty"`
	PCURL      string `json:"pc_url,omitempty"`
}

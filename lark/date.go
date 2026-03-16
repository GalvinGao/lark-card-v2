package lark

// DatePicker provides date selection (tag: "date_picker").
type DatePicker struct {
	ElementID   string         `json:"element_id,omitempty"`
	Margin      string         `json:"margin,omitempty"`
	Name        string         `json:"name,omitempty"`
	Required    bool           `json:"required,omitempty"`
	Disabled    bool           `json:"disabled,omitempty"`
	Width       string         `json:"width,omitempty"`
	Behaviors   []Behavior     `json:"behaviors,omitempty"`
	InitialDate string         `json:"initial_date,omitempty"`
	Placeholder *Text          `json:"placeholder,omitempty"`
	Value       map[string]any `json:"value,omitempty"`
	Confirm     *Confirm       `json:"confirm,omitempty"`
}

func (*DatePicker) cardTag() string { return "date_picker" }

// TimePicker provides time selection (tag: "picker_time").
type TimePicker struct {
	ElementID   string         `json:"element_id,omitempty"`
	Margin      string         `json:"margin,omitempty"`
	Name        string         `json:"name,omitempty"`
	Required    bool           `json:"required,omitempty"`
	Disabled    bool           `json:"disabled,omitempty"`
	Width       string         `json:"width,omitempty"`
	Behaviors   []Behavior     `json:"behaviors,omitempty"`
	InitialTime string         `json:"initial_time,omitempty"`
	Placeholder *Text          `json:"placeholder,omitempty"`
	Value       map[string]any `json:"value,omitempty"`
	Confirm     *Confirm       `json:"confirm,omitempty"`
}

func (*TimePicker) cardTag() string { return "picker_time" }

// DateTimePicker provides combined date and time selection (tag: "picker_datetime").
type DateTimePicker struct {
	ElementID       string         `json:"element_id,omitempty"`
	Margin          string         `json:"margin,omitempty"`
	Name            string         `json:"name,omitempty"`
	Required        bool           `json:"required,omitempty"`
	Disabled        bool           `json:"disabled,omitempty"`
	Width           string         `json:"width,omitempty"`
	Behaviors       []Behavior     `json:"behaviors,omitempty"`
	InitialDatetime string         `json:"initial_datetime,omitempty"`
	Placeholder     *Text          `json:"placeholder,omitempty"`
	Value           map[string]any `json:"value,omitempty"`
	Confirm         *Confirm       `json:"confirm,omitempty"`
}

func (*DateTimePicker) cardTag() string { return "picker_datetime" }

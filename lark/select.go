package lark

// SelectStatic is a single-select dropdown (tag: "select_static").
type SelectStatic struct {
	ElementID     string         `json:"element_id,omitempty"`
	Margin        string         `json:"margin,omitempty"`
	Type          string         `json:"type,omitempty"`
	Name          string         `json:"name,omitempty"`
	Required      bool           `json:"required,omitempty"`
	Disabled      bool           `json:"disabled,omitempty"`
	InitialOption string         `json:"initial_option,omitempty"`
	InitialIndex  int            `json:"initial_index,omitempty"`
	Placeholder   *Text          `json:"placeholder,omitempty"`
	Width         string         `json:"width,omitempty"`
	Behaviors     []Behavior     `json:"behaviors,omitempty"`
	Options       []SelectOption `json:"options,omitempty"`
	Confirm       *Confirm       `json:"confirm,omitempty"`
}

func (*SelectStatic) cardTag() string { return "select_static" }

// MultiSelectStatic is a multi-select dropdown (tag: "multi_select_static").
type MultiSelectStatic struct {
	ElementID      string         `json:"element_id,omitempty"`
	Margin         string         `json:"margin,omitempty"`
	Type           string         `json:"type,omitempty"`
	Name           string         `json:"name,omitempty"`
	Required       bool           `json:"required,omitempty"`
	Disabled       bool           `json:"disabled,omitempty"`
	SelectedValues []string       `json:"selected_values,omitempty"`
	Placeholder    *Text          `json:"placeholder,omitempty"`
	Width          string         `json:"width,omitempty"`
	Behaviors      []Behavior     `json:"behaviors,omitempty"`
	Options        []SelectOption `json:"options,omitempty"`
	Confirm        *Confirm       `json:"confirm,omitempty"`
}

func (*MultiSelectStatic) cardTag() string { return "multi_select_static" }

// SelectOption is a single option in a select dropdown.
type SelectOption struct {
	Text  *Text  `json:"text,omitempty"`
	Icon  *Icon  `json:"icon,omitempty"`
	Value string `json:"value,omitempty"`
}

// SelectPerson is a single-select person picker (tag: "select_person").
type SelectPerson struct {
	ElementID     string         `json:"element_id,omitempty"`
	Margin        string         `json:"margin,omitempty"`
	Type          string         `json:"type,omitempty"`
	Name          string         `json:"name,omitempty"`
	Required      bool           `json:"required,omitempty"`
	Disabled      bool           `json:"disabled,omitempty"`
	InitialOption string         `json:"initial_option,omitempty"`
	Placeholder   *Text          `json:"placeholder,omitempty"`
	Width         string         `json:"width,omitempty"`
	Behaviors     []Behavior     `json:"behaviors,omitempty"`
	Options       []PersonOption `json:"options,omitempty"`
	Confirm       *Confirm       `json:"confirm,omitempty"`
}

func (*SelectPerson) cardTag() string { return "select_person" }

// MultiSelectPerson is a multi-select person picker (tag: "multi_select_person").
type MultiSelectPerson struct {
	ElementID      string         `json:"element_id,omitempty"`
	Margin         string         `json:"margin,omitempty"`
	Type           string         `json:"type,omitempty"`
	Name           string         `json:"name,omitempty"`
	Required       bool           `json:"required,omitempty"`
	Disabled       bool           `json:"disabled,omitempty"`
	SelectedValues []string       `json:"selected_values,omitempty"`
	Placeholder    *Text          `json:"placeholder,omitempty"`
	Width          string         `json:"width,omitempty"`
	Behaviors      []Behavior     `json:"behaviors,omitempty"`
	Options        []PersonOption `json:"options,omitempty"`
	Confirm        *Confirm       `json:"confirm,omitempty"`
}

func (*MultiSelectPerson) cardTag() string { return "multi_select_person" }

// PersonOption is an option in a person picker.
type PersonOption struct {
	Value string `json:"value,omitempty"`
}

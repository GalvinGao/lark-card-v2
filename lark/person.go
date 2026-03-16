package lark

// Person displays a single user's name and/or avatar (tag: "person").
type Person struct {
	ElementID  string `json:"element_id,omitempty"`
	Margin     string `json:"margin,omitempty"`
	Size       string `json:"size,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	ShowAvatar *bool  `json:"show_avatar,omitempty"`
	ShowName   *bool  `json:"show_name,omitempty"`
	Style      string `json:"style,omitempty"`
}

func (*Person) cardTag() string { return "person" }

// PersonList displays multiple users (tag: "person_list").
type PersonList struct {
	ElementID         string       `json:"element_id,omitempty"`
	Margin            string       `json:"margin,omitempty"`
	DropInvalidUserID bool         `json:"drop_invalid_user_id,omitempty"`
	Lines             int          `json:"lines,omitempty"`
	ShowName          *bool        `json:"show_name,omitempty"`
	ShowAvatar        *bool        `json:"show_avatar,omitempty"`
	Size              string       `json:"size,omitempty"`
	Persons           []PersonItem `json:"persons,omitempty"`
	Icon              *Icon        `json:"icon,omitempty"`
	UdIcon            *UdIcon      `json:"ud_icon,omitempty"`
}

func (*PersonList) cardTag() string { return "person_list" }

// PersonItem identifies a person by their ID.
type PersonItem struct {
	ID string `json:"id,omitempty"`
}

// UdIcon is a custom icon with style properties.
type UdIcon struct {
	Token string            `json:"token,omitempty"`
	Style map[string]string `json:"style,omitempty"`
}

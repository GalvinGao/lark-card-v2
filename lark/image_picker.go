package lark

// ImagePicker allows selecting from image options (tag: "select_img").
type ImagePicker struct {
	ElementID    string              `json:"element_id,omitempty"`
	Margin       string              `json:"margin,omitempty"`
	MultiSelect  bool                `json:"multi_select,omitempty"`
	Layout       string              `json:"layout,omitempty"`
	Name         string              `json:"name,omitempty"`
	Required     bool                `json:"required,omitempty"`
	CanPreview   *bool               `json:"can_preview,omitempty"`
	AspectRatio  string              `json:"aspect_ratio,omitempty"`
	Disabled     bool                `json:"disabled,omitempty"`
	DisabledTips *Text               `json:"disabled_tips,omitempty"`
	Value        any                 `json:"value,omitempty"`
	Options      []ImagePickerOption `json:"options,omitempty"`
	Behaviors    []Behavior          `json:"behaviors,omitempty"`
	Confirm      *Confirm            `json:"confirm,omitempty"`
}

func (*ImagePicker) cardTag() string { return "select_img" }

// ImagePickerOption is a single image option in an ImagePicker.
type ImagePickerOption struct {
	ImgKey       string            `json:"img_key,omitempty"`
	I18nImgKey   map[string]string `json:"i18n_img_key,omitempty"`
	Value        string            `json:"value,omitempty"`
	Disabled     bool              `json:"disabled,omitempty"`
	DisabledTips *Text             `json:"disabled_tips,omitempty"`
	HoverTips    *Text             `json:"hover_tips,omitempty"`
}

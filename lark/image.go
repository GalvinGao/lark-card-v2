package lark

// Image is the image display component (tag: "img").
type Image struct {
	ElementID    string            `json:"element_id,omitempty"`
	Margin       string            `json:"margin,omitempty"`
	ImgKey       string            `json:"img_key,omitempty"`
	I18nImgKey   map[string]string `json:"i18n_img_key,omitempty"`
	Alt          *Text             `json:"alt,omitempty"`
	Title        *Text             `json:"title,omitempty"`
	CornerRadius string            `json:"corner_radius,omitempty"`
	ScaleType    string            `json:"scale_type,omitempty"`
	Size         string            `json:"size,omitempty"`
	Transparent  bool              `json:"transparent,omitempty"`
	Preview      *bool             `json:"preview,omitempty"`
	Mode         string            `json:"mode,omitempty"`
	CustomWidth  string            `json:"custom_width,omitempty"`
	CompactWidth bool              `json:"compact_width,omitempty"`
}

func (*Image) cardTag() string { return "img" }

// MultiImageLayout displays multiple images in a layout (tag: "img_combination").
type MultiImageLayout struct {
	ElementID              string      `json:"element_id,omitempty"`
	Margin                 string      `json:"margin,omitempty"`
	CombinationMode        string      `json:"combination_mode,omitempty"`
	CombinationTransparent bool        `json:"combination_transparent,omitempty"`
	CornerRadius           string      `json:"corner_radius,omitempty"`
	ImgList                []ImageItem `json:"img_list,omitempty"`
}

func (*MultiImageLayout) cardTag() string { return "img_combination" }

// ImageItem is a single image within a MultiImageLayout.
type ImageItem struct {
	ImgKey      string            `json:"img_key,omitempty"`
	I18nImgKey  map[string]string `json:"i18n_img_key,omitempty"`
	Transparent bool              `json:"transparent,omitempty"`
}

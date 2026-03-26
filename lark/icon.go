package lark

// Icon represents a standard or custom icon.
type Icon struct {
	Tag        string            `json:"tag,omitempty"`
	Token      string            `json:"token,omitempty"`
	Color      string            `json:"color,omitempty"`
	ImgKey     string            `json:"img_key,omitempty"`
	I18nImgKey map[string]string `json:"i18n_img_key,omitempty"`
}

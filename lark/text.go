package lark

// Text represents a plain_text or lark_md text object.
// Used in headers, placeholders, button labels, etc.
type Text struct {
	Tag         string            `json:"tag,omitempty"`
	Content     string            `json:"content,omitempty"`
	I18nContent map[string]string `json:"i18n_content,omitempty"`
	TextSize    string            `json:"text_size,omitempty"`
	TextColor   string            `json:"text_color,omitempty"`
	TextAlign   string            `json:"text_align,omitempty"`
	Lines       int               `json:"lines,omitempty"`
}

// Div is the plain text display component (tag: "div").
type Div struct {
	ElementID string `json:"element_id,omitempty"`
	Margin    string `json:"margin,omitempty"`
	Width     string `json:"width,omitempty"`
	Text      *Text  `json:"text,omitempty"`
	Icon      *Icon  `json:"icon,omitempty"`
}

func (*Div) cardTag() string { return "div" }

// Markdown is the rich text display component (tag: "markdown").
type Markdown struct {
	ElementID   string            `json:"element_id,omitempty"`
	Margin      string            `json:"margin,omitempty"`
	Content     string            `json:"content,omitempty"`
	I18nContent map[string]string `json:"i18n_content,omitempty"`
	TextSize    string            `json:"text_size,omitempty"`
	TextColor   string            `json:"text_color,omitempty"`
	TextAlign   string            `json:"text_align,omitempty"`
	Icon        *Icon             `json:"icon,omitempty"`
}

func (*Markdown) cardTag() string { return "markdown" }

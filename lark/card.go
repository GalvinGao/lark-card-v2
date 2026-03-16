package lark

import "encoding/json"

// Card represents a Lark Card JSON v2.0 structure.
type Card struct {
	Config   *Config   `json:"config,omitempty"`
	CardLink *CardLink `json:"card_link,omitempty"`
	Header   Header    `json:"header,omitempty"`
	Body     Body      `json:"body,omitempty"`
}

// Config configures global card behavior.
type Config struct {
	StreamingMode            bool     `json:"streaming_mode,omitempty"`
	StreamingConfig          any      `json:"streaming_config,omitempty"`
	Summary                  *Summary `json:"summary,omitempty"`
	Locales                  []string `json:"locales,omitempty"`
	EnableForward            *bool    `json:"enable_forward,omitempty"`
	UpdateMulti              *bool    `json:"update_multi,omitempty"`
	WidthMode                string   `json:"width_mode,omitempty"`
	UseCustomTranslation     bool     `json:"use_custom_translation,omitempty"`
	EnableForwardInteraction bool     `json:"enable_forward_interaction,omitempty"`
	Style                    *Style   `json:"style,omitempty"`
}

// Summary holds card summary information for chat bar preview.
type Summary struct {
	Content     string            `json:"content,omitempty"`
	I18nContent map[string]string `json:"i18n_content,omitempty"`
}

// Style defines custom font sizes and colors.
type Style struct {
	TextSize map[string]TextSizeConfig `json:"text_size,omitempty"`
	Color    map[string]ColorConfig    `json:"color,omitempty"`
}

// TextSizeConfig defines per-platform font sizes.
type TextSizeConfig struct {
	Default string `json:"default,omitempty"`
	PC      string `json:"pc,omitempty"`
	Mobile  string `json:"mobile,omitempty"`
}

// ColorConfig defines per-theme colors.
type ColorConfig struct {
	LightMode string `json:"light_mode,omitempty"`
	DarkMode  string `json:"dark_mode,omitempty"`
}

// CardLink specifies the overall jump link for the card.
type CardLink struct {
	URL        string `json:"url,omitempty"`
	PCURL      string `json:"pc_url,omitempty"`
	IOSURL     string `json:"ios_url,omitempty"`
	AndroidURL string `json:"android_url,omitempty"`
}

// Header configures the card title area.
type Header struct {
	Title           *Text                `json:"title,omitempty"`
	Subtitle        *Text                `json:"subtitle,omitempty"`
	TextTagList     []TextTag            `json:"text_tag_list,omitempty"`
	I18nTextTagList map[string][]TextTag `json:"i18n_text_tag_list,omitempty"`
	Template        string               `json:"template,omitempty"`
	Icon            *Icon                `json:"icon,omitempty"`
	Padding         string               `json:"padding,omitempty"`
}

// TextTag represents a suffix tag on the header title.
type TextTag struct {
	ElementID string `json:"element_id,omitempty"`
	Text      *Text  `json:"text,omitempty"`
	Color     string `json:"color,omitempty"`
}

// MarshalJSON injects "tag": "text_tag" into the JSON output.
func (t TextTag) MarshalJSON() ([]byte, error) {
	type textTagAlias TextTag
	return json.Marshal(struct {
		Tag string `json:"tag"`
		textTagAlias
	}{
		Tag:          "text_tag",
		textTagAlias: textTagAlias(t),
	})
}

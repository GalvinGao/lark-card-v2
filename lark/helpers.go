package lark

// Md creates a Markdown element with the given content.
func Md(content string) *Markdown {
	return &Markdown{Content: content}
}

// PlainText creates a Text with tag "plain_text".
func PlainText(content string) *Text {
	return &Text{Tag: "plain_text", Content: content}
}

// MdText creates a Text with tag "lark_md" for use in header titles, placeholders, etc.
func MdText(content string) *Text {
	return &Text{Tag: "lark_md", Content: content}
}

// Btn creates a Button with plain text.
func Btn(text string) *Button {
	return &Button{Text: PlainText(text)}
}

// Hr creates a Divider (horizontal rule).
func Hr() *Divider {
	return &Divider{}
}

// Img creates an Image with the given key.
func Img(key string) *Image {
	return &Image{ImgKey: key}
}

# Rich Text / Markdown — `markdown`

Tag: `markdown`

## JSON Example

```json
{
  "tag": "markdown",
  "content": "# Heading\n**Bold** *italic* ~~strike~~\n- list item",
  "text_size": "normal",
  "text_align": "left",
  "icon": { "tag": "standard_icon", "token": "info_outlined", "color": "blue" },
  "margin": "0px"
}
```

## Key Fields

- **content**: Full CommonMark markdown string plus Lark HTML extensions (see below)
- **text_size**: Same size values as plain text (`"heading-0"` through `"x-small"`, `"normal"`, etc.)
- **text_align**: `"left"` | `"center"` | `"right"`
- **icon**: Optional prefix icon (standard_icon or custom_icon)
- **margin**: Component margin

## Supported Markdown Syntax (v2.0)

### Standard Markdown

- **Headings**: `#` through `######`
- **Bold**: `**text**`
- **Italic**: `*text*`
- **Strikethrough**: `~~text~~`
- **Links**: `[text](url)`
- **Images**: `![alt](img_key)`
- **Code blocks**: ` ```language\ncode\n``` `
- **Inline code**: `` `code` ``
- **Ordered lists**: `1. item`
- **Unordered lists**: `- item`
- **Blockquotes**: `> text`
- **Tables**: `| col1 | col2 |` with `|---|---|` header separator
- **Dividers**: `---`

### Lark HTML Extensions

| Tag | Usage | Description |
|---|---|---|
| `<br>` | `<br>` | Line break |
| `<hr>` | `<hr>` | Horizontal rule |
| `<font>` | `<font color='red'>text</font>` | Colored text |
| `<at>` | `<at id=open_id></at>` | @mention user |
| `<at>` | `<at id=all></at>` | @mention everyone |
| `<text_tag>` | `<text_tag color='blue'>tag</text_tag>` | Inline tag badge |
| `<person>` | `<person id='user_id' show_name=true show_avatar=true style='normal'></person>` | Inline person display |
| `<link>` | `<link icon='chat_outlined' url='...'>text</link>` | Link with icon |
| `<number_tag>` | `<number_tag>1</number_tag>` | Numeric badge |
| `<local_datetime>` | `<local_datetime millisecond='' format_type='date_num'></local_datetime>` | Localized datetime |

## Line Break Behavior

- Single `\n` = soft break (may render as space on some clients)
- Double `\n\n` = hard break / new paragraph

## Escaping

Use HTML entities for special characters: `&#42;` for `*`, `&#95;` for `_`, `&#96;` for `` ` ``, etc.

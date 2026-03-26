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
- **text_size**: `"heading-0"` through `"x-small"`, `"normal"`, etc. (see size table below). Also accepts custom size names defined in `config.style.text_size`.
- **text_align**: `"left"` | `"center"` | `"right"`
- **icon**: Optional prefix icon (`standard_icon` or `custom_icon`)
- **margin**: Component margin, range [-99,99]px

## Supported Markdown Syntax (v2.0)

### Standard Markdown

- **Headings**: `#` through `######` (26px, 22px, 20px, 18px, 17px, 14px)
- **Bold**: `**text**` or `__text__`
- **Italic**: `*text*`
- **Strikethrough**: `~~text~~`
- **Links**: `[text](url)` — must include `http://` or `https://`
- **Phone links**: `[display text](tel://phone-number)` — mobile only
- **Images**: `![hover_text](img_key)` — `img_key` from Upload Image API
- **Code blocks**: ` ```language\ncode\n``` ` — case-insensitive language name
- **Inline code**: `` `code` ``
- **Ordered lists**: `1. item` — 4 spaces per indent level
- **Unordered lists**: `- item` — 4 spaces per indent level
- **Blockquotes**: `> text`
- **Tables**: `| col1 | col2 |` with `|---|---|` separator — max 5 data rows, max 4 tables per component, JSON 2.0 only
- **Dividers**: `---` or `<hr>` — must be on its own line
- **Lark emoji**: `:DONE:` `:THUMBSUP:` etc.

### Lark HTML Extensions

| Tag | Usage | Notes |
|-----|-------|-------|
| `<br>` | `<br>` or `<br/>` | Line break. In JSON, `\n` also works. |
| `<hr>` | `<hr>` or `<hr/>` | Horizontal rule. Must be on its own line. |
| `<font>` | `<font color='red'>text</font>` | Colored text. Supports color enums and RGBA. Does not apply to link text. Supports nesting other tags. |
| `<at>` | `<at id=open_id></at>` | @mention user. Also: `<at ids=id1,id2></at>`, `<at email=x@y.com></at>`. |
| `<at>` | `<at id=all></at>` | @mention everyone. Requires group owner permission. |
| `<text_tag>` | `<text_tag color='blue'>tag</text_tag>` | Inline tag badge. Colors: `neutral` `blue` `turquoise` `lime` `orange` `violet` `indigo` `wathet` `green` `yellow` `red` `purple` `carmine` |
| `<person>` | `<person id='user_id' show_name=true show_avatar=true style='normal'></person>` | Inline person display. `id` accepts open_id/union_id/user_id. `style`: `normal` (default) or `capsule`. |
| `<link>` | `<link icon='token' url='...' pc_url='' ios_url='' android_url=''>text</link>` | Link with prefix icon. Icon color fixed blue. Per-platform URLs override `url`. |
| `<number_tag>` | `<number_tag>1</number_tag>` | Numeric badge (0-99). Optional attrs: `background_color`, `font_color`, `url`, `pc_url`, `ios_url`, `android_url`. |
| `<local_datetime>` | `<local_datetime millisecond='' format_type='date_num' link='...'></local_datetime>` | Localized datetime. `format_type`: `date_num` `date_short` `date` `week` `week_short` `time` `time_sec` `timezone`. |
| `<a>` | `<a href='url'></a>` | Hyperlink (HTTP/HTTPS only). |
| `<raw>` | `<raw>content</raw>` | Raw content passthrough. |

## Line Break Behavior

- Single `\n` = soft break (may render as space on some clients)
- Double `\n\n` = hard break / new paragraph

## Escaping

Use HTML entities for special characters: `&#42;` for `*`, `&#95;` for `_`, `&#96;` for `` ` ``, `&#60;` for `<`, `&#62;` for `>`, etc. Format: `&#entity_number;`. Full list: [HTML escape reference](https://www.w3school.com.cn/charsets/ref_html_8859.asp).

## Text Size Values

| Value | PC | Mobile |
|-------|----|--------|
| `heading-0` | 30px | 26px |
| `heading-1` | 24px | 24px |
| `heading-2` | 20px | 20px |
| `heading-3` | 18px | 17px |
| `heading-4` | 16px | 16px |
| `heading` | 16px | 16px |
| `normal` | 14px | 14px |
| `notation` | 12px | 12px |
| `xxxx-large` | 30px | 26px |
| `xxx-large` | 24px | 24px |
| `xx-large` | 20px | 20px |
| `x-large` | 18px | 18px |
| `large` | 16px | 17px |
| `medium` | 14px | 14px |
| `small` | 12px | 12px |
| `x-small` | 10px | 10px |

Unrecognized values fall back to `normal`.

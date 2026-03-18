# Plain Text — `div`

Tag: `div`

## JSON Example

```json
{
  "tag": "div",
  "text": {
    "tag": "plain_text",
    "content": "Hello world",
    "text_size": "normal",
    "text_color": "default",
    "text_align": "left",
    "lines": 2
  },
  "icon": {
    "tag": "standard_icon",
    "token": "info_outlined",
    "color": "blue"
  },
  "width": "fill",
  "margin": "0px"
}
```

## Key Fields

- **text.tag**: `"plain_text"` (plain text and emoji) or `"lark_md"` (partial markdown with inline formatting)
- **text.content**: The text content string
- **text.text_size**: Predefined sizes (PC / mobile px):

  | Value | PC | Mobile |
  |-------|-----|--------|
  | `heading-0` / `xxxx-large` | 30px | 26px |
  | `heading-1` / `xxx-large` | 24px | 24px |
  | `heading-2` / `xx-large` | 20px | 20px |
  | `heading-3` / `x-large` | 18px | 17px |
  | `heading-4` / `heading` / `large` | 16px | 16px |
  | `normal` / `medium` | 14px | 14px |
  | `notation` / `small` | 12px | 12px |
  | `x-small` | 10px | 10px |

  Custom per-platform sizes via `config.style.text_size` (see `card-structure.md`).
- **text.text_color**: `"default"` or a color enum value. Only works when `text.tag` is `"plain_text"`.
- **text.text_align**: `"left"` | `"center"` | `"right"`
- **text.lines**: Max display lines (integer). Content exceeding this is truncated with `"..."`.
- **width**: `"fill"` | `"auto"` | `"[16,999]px"`
- **icon**: Optional prefix icon (standard_icon or custom_icon)
- **margin**: Component margin

## lark_md Syntax

When `text.tag` is `"lark_md"`, the following inline formatting is supported:

- `*italic*` — italic text
- `**bold**` — bold text
- `~~strikethrough~~` — strikethrough text
- `[link text](url)` — hyperlink
- `<at id=open_id></at>` — @mention a user
- `<font color='red'>text</font>` — colored text
- `<text_tag color='blue'>tag</text_tag>` — inline tag badge

Note: `lark_md` in `div` supports only inline formatting. For full markdown (headings, lists, code blocks, etc.), use the `markdown` component instead.

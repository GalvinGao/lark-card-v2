# Plain Text

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/plain-text>

The plain text component supports adding plain text and prefix icons, and setting display styles such as text size, color, and alignment.

## Notes

- This document covers the JSON 2.0 structure of the plain text component. For the historical JSON 1.0 structure, refer to the Lark documentation.
- In the Lark card building tool, only the `plain_text` tag type is supported. Use the rich text component to add Markdown-formatted text.

## JSON Structure

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "div",
        "element_id": "custom_id",
        "margin": "0px 0px 0px 0px",
        "width": "fill",
        "text": {
          "tag": "plain_text",
          "content": "",
          "text_size": "normal",
          "text_color": "default",
          "text_align": "left",
          "lines": 2
        },
        "icon": {
          "tag": "standard_icon",
          "token": "chat-forbidden_outlined",
          "color": "orange",
          "img_key": "img_v2_38811724"
        }
      }
    ]
  }
}
```

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | Tag of the component. Fixed to `div` for the plain text component. |
| `element_id` | No | String | Empty | Unique identifier for the component. Used to specify the component in interface calls. Must be globally unique within the same card. Only letters, numbers, and underscores are allowed; must start with a letter and not exceed 20 characters. |
| `margin` | No | String | `0` | Margin of the component. Value range is `[-99,99]px`. Single value (e.g. `"10px"`) sets all four margins. Double values (e.g. `"4px 0"`) set top/bottom and left/right. Four values (e.g. `"4px 0 4px 0"`) set top, right, bottom, and left respectively. |
| `width` | No | String | `fill` | Text width. `fill`: matches component width. `auto`: adapts to text content length. `[16,999]px`: custom width. |
| `text` | No | Object | — | Configures the plain text information of the card. |
| `text.tag` | Yes | String | `plain_text` | Tag for text type. `plain_text`: plain text content or emojis. `lark_md`: text content supporting partial Markdown syntax (see below). |
| `text.content` | Yes | String | — | Text content. When `tag` is `lark_md`, supports partial Markdown syntax. |
| `text.text_size` | No | String | `normal` | Text size. See the text size values table below. You can also define different font sizes for mobile and desktop (see below). |
| `text.text_color` | No | String | `default` | Text color. Only effective when `tag` is `plain_text`. `default`: black in light theme, white in dark theme. Also accepts color enumeration values. |
| `text.text_align` | No | String | `left` | Text alignment. `left`, `center`, or `right`. |
| `text.lines` | No | Int | — | Maximum number of lines to display. Content exceeding this limit is truncated with `...`. |
| `icon` | No | Object | — | Adds an icon as a text prefix. Supports custom or library icons. |
| `icon.tag` | No | String | — | Icon type. `standard_icon`: use an icon from the library. `custom_icon`: use a custom image as an icon. |
| `icon.token` | No | String | — | Token of the icon from the library. Effective when `icon.tag` is `standard_icon`. See the Icon Library for values. |
| `icon.color` | No | String | — | Icon color. Supports setting colors for linear and filled icons (tokens ending with `outlined` or `filled`). Effective when `icon.tag` is `standard_icon`. See Color Enumeration Values. |
| `icon.img_key` | No | String | — | Custom prefix icon image key. Effective when `icon.tag` is `custom_icon`. Obtain the key by calling the Upload Image API. |

### Text Size Values

| Value | Description |
|-------|-------------|
| `heading-0` | Extra large title (30px) |
| `heading-1` | First-level title (24px) |
| `heading-2` | Second-level title (20px) |
| `heading-3` | Third-level title (18px) |
| `heading-4` | Fourth-level title (16px) |
| `heading` | Title (16px) |
| `normal` | Text (14px) |
| `notation` | Supplementary information (12px) |
| `xxxx-large` | 30px |
| `xxx-large` | 24px |
| `xx-large` | 20px |
| `x-large` | 18px |
| `large` | 16px |
| `medium` | 14px |
| `small` | 12px |
| `x-small` | 10px |

## Markdown Syntax Supported by lark_md

| Feature | Syntax | Effect |
|---------|--------|--------|
| New Line | `First line\nSecond line` | First line / Second line |
| Italic | `*Italic*` | *Italic* |
| Bold | `**Bold**` or `__Bold__` | **Bold** |
| Strikethrough | `~~Strikethrough~~` | ~~Strikethrough~~ |
| Text Link | `[Text Link](https://www.larksuite.com)` | Text Link |
| Hyperlink | `<a href='https://open.larksuite.com'></a>` | https://open.larksuite.com |
| @ Mention | `<at id=all></at>`, `<at id={{open_id}}></at>`, `<at id={{user_id}}></at>`, `<at email=test@email.com></at>` | @all, @test |
| Colored Text | `<font color=red>Red</font>` (see Color Enumeration Values) | Red |
| Emoji | `😁😢🌞💼🏆❌✅` (copy emojis directly) | 😁😢🌞💼🏆❌✅ |
| Lark Emoticons | `:OK:` (see Emoticon Descriptions) | — |
| Tags | `<text_tag color='neutral'> neutral </text_tag>` (colors: `neutral`, `blue`, `turquoise`, `lime`, `orange`, `violet`, `indigo`, `wathet`, `green`, `yellow`, `red`, `purple`, `carmine`) | — |

## Example

### plain_text Type

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "div",
        "text": {
          "tag": "plain_text",
          "content": "这是示例文本。",
          "text_size": "normal",
          "text_align": "center",
          "text_color": "default"
        },
        "icon": {
          "tag": "standard_icon",
          "token": "reply-cn_filled",
          "color": "blue"
        },
        "margin": "0px 0px 0px 0px",
        "element_id": "plaintext01"
      }
    ]
  }
}
```

### lark_md Type

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "div",
        "text": {
          "tag": "plain_text",
          "content": "text-lark_md",
          "lines": 1
        },
        "fields": [
          {
            "is_short": false,
            "text": {
              "tag": "lark_md",
              "content": "<a>https://open.larksuite.com</a>"
            }
          },
          {
            "is_short": false,
            "text": {
              "tag": "lark_md",
              "content": "ready\nnew line"
            }
          },
          {
            "is_short": false,
            "text": {
              "tag": "lark_md",
              "content": "*Italic*"
            }
          },
          {
            "is_short": false,
            "text": {
              "tag": "lark_md",
              "content": "**Bold**"
            }
          },
          {
            "is_short": false,
            "text": {
              "tag": "lark_md",
              "content": "~~delete line~~"
            }
          },
          {
            "is_short": false,
            "text": {
              "tag": "lark_md",
              "content": "<at id=all></at>"
            }
          }
        ]
      }
    ]
  }
}
```

## Defining Different Font Sizes for Mobile and Desktop

In both plain text and rich text components, you can define different font sizes for the same text on mobile and desktop platforms.

### Custom Font Size Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `text_size` | No | Object | — | Text size. Allows customizing different font sizes for mobile and desktop platforms. |
| `text_size.<custom_name>` | No | Object | — | Custom font size entry. Define a name such as `cus-0`, `cus-1`, etc. |
| `text_size.<custom_name>.default` | No | String | — | Font size for old Lark clients that do not support per-platform configuration. Recommended. Accepts the same text size values listed above. |
| `text_size.<custom_name>.pc` | No | String | — | Desktop platform font size. Same values as above. |
| `text_size.<custom_name>.mobile` | No | String | — | Mobile platform font size. Note: some mobile sizes differ from desktop (e.g. `heading-0` is 26px on mobile vs 30px on desktop; `heading-3` is 17px vs 18px; `large` is 17px vs 16px). |

### Step 1: Define custom font sizes in the card config

```json
{
  "config": {
    "style": {
      "text_size": {
        "cus-0": {
          "default": "medium",
          "pc": "medium",
          "mobile": "large"
        },
        "cus-1": {
          "default": "medium",
          "pc": "normal",
          "mobile": "x-large"
        }
      }
    }
  }
}
```

### Step 2: Apply custom font size in the component

```json
{
  "i18n_elements": {
    "zh_cn": [
      {
        "tag": "column_set",
        "flex_mode": "none",
        "horizontal_spacing": "default",
        "background_style": "default",
        "columns": [
          {
            "tag": "column",
            "elements": [
              {
                "tag": "div",
                "text": {
                  "tag": "plain_text",
                  "content": "This is an example of plain text.",
                  "text_size": "cus-0",
                  "text_align": "center",
                  "text_color": "default"
                },
                "icon": {
                  "tag": "standard_icon",
                  "token": "app-default_filled",
                  "color": "blue"
                }
              }
            ],
            "width": "weighted",
            "weight": 1
          }
        ]
      }
    ],
    "i18n_header": {}
  }
}
```

# Rich Text (Markdown)

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/rich-text>

The rich text (Markdown) component supports rendering text, images, dividers, and other elements.

## Notes

- This document covers the JSON 2.0 structure of the rich text component. For the historical JSON 1.0 structure, refer to the Lark documentation.
- The JSON 2.0 structure no longer supports the differentiated jump syntax (`"href": { "urlVal": { ... } }`). Use the link syntax with an icon as a replacement:

```json
{
  "tag": "markdown",
  "content": "<link icon='chat_outlined' url='https://applink.larksuite.com/client/chat/xxxxx' pc_url='' ios_url='' android_url=''>Strategy Seminar</link>"
}
```

## JSON Structure

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "markdown",
        "element_id": "custom_id",
        "margin": "0px 0px 0px 0px",
        "content": "Personnel<person id='ou_449b53ad6aee526f7ed311b216aabcef' show_name=true show_avatar=true style='normal'></person>",
        "text_size": "normal",
        "text_align": "left",
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
| `tag` | Yes | String | — | Tag of the component. Fixed to `markdown` for the rich text component. |
| `element_id` | No | String | Empty | Unique identifier for the component. Used to specify the component in interface calls. Must be globally unique within the same card. Only letters, numbers, and underscores are allowed; must start with a letter and not exceed 20 characters. |
| `margin` | No | String | `0` | Margin of the component. Value range is `[-99,99]px`. Single value (e.g. `"10px"`) sets all four margins. Double values (e.g. `"4px 0"`) set top/bottom and left/right. Four values (e.g. `"4px 0 4px 0"`) set top, right, bottom, and left respectively. |
| `text_align` | No | String | `left` | Text alignment. `left`, `center`, or `right`. |
| `text_size` | No | String | `normal` | Text size. See the text size values table below. You can also define different font sizes for mobile and desktop (see below). |
| `icon` | No | Object | — | Adds an icon as a text prefix. Supports custom or library icons. |
| `icon.tag` | No | String | — | Icon type. `standard_icon`: use an icon from the library. `custom_icon`: use a custom image as an icon. |
| `icon.token` | No | String | — | Token of the icon from the library. Effective when `icon.tag` is `standard_icon`. See the Icon Library for values. |
| `icon.color` | No | String | — | Icon color. Supports setting colors for linear and filled icons (tokens ending with `outlined` or `filled`). Effective when `icon.tag` is `standard_icon`. See Color Enumeration Values. |
| `icon.img_key` | No | String | — | Custom prefix icon image key. Effective when `icon.tag` is `custom_icon`. Obtain the key by calling the Upload Image API. |
| `content` | Yes | String | — | Markdown text content. For supported syntax, see below. |

### Text Size Values

| Value | Description |
|-------|-------------|
| `heading-0` | Extra large title (30px) |
| `heading-1` | First-level title (24px) |
| `heading-2` | Second-level title (20px) |
| `heading-3` | Third-level title (18px) |
| `heading-4` | Fourth-level title (16px) |
| `heading` | Title (16px) |
| `normal` | Body text (14px) |
| `notation` | Auxiliary information (12px) |
| `xxxx-large` | 30px |
| `xxx-large` | 24px |
| `xx-large` | 20px |
| `x-large` | 18px |
| `large` | 16px |
| `medium` | 14px |
| `small` | 12px |
| `x-small` | 10px |

## Example

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "markdown",
        "content": "# 一级标题",
        "margin": "0px 0px 0px 0px",
        "text_align": "left",
        "text_size": "normal"
      },
      {
        "tag": "markdown",
        "content": "标准emoji 😁😢🌞💼🏆❌✅\nLarkemoji :OK::THUMBSUP:\n*斜体* **粗体** ~~删除线~~ \n<font color='red'>这是红色文本</font>\n<text_tag color=\"blue\">标签</text_tag>\n[文字链接](https://open.feishu.cn/document/server-docs/im-v1/message-reaction/emojis-introduce)\n<link icon='chat_outlined' url='https://open.feishu.cn' pc_url='' ios_url='' android_url=''>带图标的链接</link>\n<at id=all></at>\n- 无序列表1\n    - 无序列表 1.1\n- 无序列表2\n1. 有序列表1\n    1. 有序列表 1.1\n2. 有序列表2\n```JSON\n{\"This is\": \"JSON demo\"}\n```"
      },
      {
        "tag": "markdown",
        "content": "行内引用`code`"
      },
      {
        "tag": "markdown",
        "content": "数字角标，支持 1-99 数字<number_tag background_color='grey' font_color='white' url='https://open.larksuite.com' pc_url='https://open.larksuite.com' android_url='https://open.larksuite.com' ios_url='https://open.larksuite.com'>1</number_tag>"
      },
      {
        "tag": "markdown",
        "content": "默认数字角标展示<number_tag>1</number_tag>"
      },
      {
        "tag": "markdown",
        "content": "人员<person id='ou_449b53ad6aee526f7ed311b216a8f88f' show_name=true show_avatar=true style='normal'></person>"
      },
      {
        "tag": "markdown",
        "content": "> 这是一段引用文字\n引用内换行 \n"
      }
    ]
  }
}
```

## Supported Markdown Syntax

Card JSON 2.0 supports all standard Markdown syntax and some HTML syntax (except HTMLBlock). Refer to the [CommonMark Spec](https://spec.commonmark.org/) for standard Markdown syntax.

**Notes on rendering differences from CommonMark:**
- One Enter key acts as a soft break; two Enter keys act as a hard break. Soft breaks may be ignored during rendering; hard breaks always render as a new line.

**Supported HTML syntax in 2.0:**

- `<br>`, `<br/>`
- `<hr>`, `<hr/>`
- `<person></person>`
- `<local_datetime></local_datetime>`
- `<at></at>`
- `<a></a>`
- `<text_tag></text_tag>`
- `<raw></raw>`
- `<link></link>`
- `<font></font>` (supports nesting other tags such as `<at>`, `<a>`, `<link>`, `<local_datetime>`, and other `<font>` tags)

### Syntax Reference

| Name | Syntax | Notes |
|------|--------|-------|
| Line Break | `Row 1<br>Row 2` or `Row 1<br />Row 2` | In card JSON, you can also use `\n` for line breaks. |
| Italic | `*Italic*` | — |
| Bold | `**Bold**` or `__Bold__` | Do not use 4 consecutive `*` or `_`. Ensure spaces exist before and after bold syntax if the effect does not display. |
| Strikethrough | `~~Strikethrough~~` | — |
| Mention Person | `<at id=open_id></at>`, `<at id=user_id></at>`, `<at email=test@email.com></at>`, `<at ids=id_01,id_02,xxx></at>` | Mentioned users receive a notification. Custom robots only support `open_id` and `user_id`. Use `<at ids=...>` to mention multiple people. |
| Mention Everyone | `<at id=all></at>` | Requires group owner permission. If not enabled, the card will fail to send. |
| Hyperlink | `<a href='https://open.larksuite.com'></a>` | Must include HTTP/HTTPS schema. Link text color cannot be customized. |
| Colored Text | `<font color='green'>Green text</font>` | `default`: white background, black text. Supports color enumerations and RGBA syntax. Does not apply to text in links. |
| Clickable Phone | `[Display text](tel://phone_number)` | Effective only on mobile devices. |
| Text Link | `[Open Platform](https://open.larksuite.com/)` | Must include HTTP/HTTPS schema. |
| Differential Redirect | `[Label]($urlVal)` with `href.urlVal` object | Use only when different links are needed for PC and mobile. Must include HTTP/HTTPS schema. |
| Image | `![hover_text](image_key)` | `hover_text` is shown on PC hover. Obtain `image_key` via the Upload Image API. |
| Divider | `---` | Must be on its own line with line breaks before and after if surrounded by text. |
| Lark Emoji | `:DONE:` | See Emoji Text Description for supported keys. |
| Tag | `<text_tag color='red'>Tag text</text_tag>` | Colors: `neutral`, `blue`, `turquoise`, `lime`, `orange`, `violet`, `indigo`, `wathet`, `green`, `yellow`, `red`, `purple`, `carmine`. |
| Ordered List | `1. Item 1\n    1. Sub-item 1.1\n2. Item 2` | Numbers at start of line; 4 spaces per indentation level. Requires Lark 7.6+. |
| Unordered List | `- Item 1\n    - Sub-item 1.1\n- Item 2` | Marks at start of line; 4 spaces per indentation level. Requires Lark 7.6+. |
| Code Block | `` ```JSON\n{...}\n``` `` | Must start at the beginning of a line. Supports specifying a programming language. Four or more leading spaces also trigger a code block. |
| Link with Icon | `<link icon='chat_outlined' url='...' pc_url='' ios_url='' android_url=''>Text</link>` | `icon`: from the icon library (fixed blue color, optional). `url`: default link (required). `pc_url`, `ios_url`, `android_url`: device-specific overrides (optional). |
| Personnel | `<person id='user_id' show_name=true show_avatar=true style='normal'></person>` | `id`: supports `open_id`, `union_id`, `user_id`. `show_name`/`show_avatar` default to `true`. `style`: `normal` (default) or `capsule`. |
| Heading | `# H1` through `###### H6` | Font sizes: H1=26px, H2=22px, H3=20px, H4=18px, H5=17px, H6=14px. JSON 2.0 only. |
| Blockquote | `> This is a blockquote` | JSON 2.0 only. |
| Inline Code | `` `code` `` | JSON 2.0 only. |
| Number Badge | `<number_tag>1</number_tag>` | Supports 0-99. Optional attributes: `background_color`, `font_color`, `url`, `pc_url`, `ios_url`, `android_url`. JSON 2.0 only. |
| Table | `\| Header \| Header \|\n\| --- \| --- \|\n\| Cell \| Cell \|` | Max 5 data rows shown (paginated beyond that). Max 4 tables per component. Use the Table component for column width control. |
| Internationalization Time | `<local_datetime millisecond='' format_type='date_num' link='...'></local_datetime>` | Displays time in the user's local timezone. `format_type` values: `date_num`, `date_short`, `date`, `week`, `week_short`, `time`, `time_sec`, `timezone`. |

## Special Character Escaping

If text contains characters used in Markdown syntax (such as `*`, `~`, `>`, `<`), use HTML escape sequences.

| Character | Escape | Description |
|-----------|--------|-------------|
| (space) | `&nbsp;` | Non-breaking space |
| (space) | `&ensp;` | Half-width space |
| (space) | `&emsp;` | Full-width space |
| `>` | `&#62;` | Greater than |
| `<` | `&#60;` | Less than |
| `~` | `&sim;` | Tilde |
| `-` | `&#45;` | Hyphen |
| `!` | `&#33;` | Exclamation mark |
| `*` | `&#42;` | Asterisk |
| `/` | `&#47;` | Forward slash |
| `\` | `&#92;` | Backslash |
| `[` | `&#91;` | Left square bracket |
| `]` | `&#93;` | Right square bracket |
| `(` | `&#40;` | Left parenthesis |
| `)` | `&#41;` | Right parenthesis |
| `#` | `&#35;` | Hash |
| `:` | `&#58;` | Colon |
| `+` | `&#43;` | Plus sign |
| `"` | `&#34;` | Double quote |
| `'` | `&#39;` | Single quote |
| `` ` `` | `&#96;` | Backtick |
| `$` | `&#36;` | Dollar sign |
| `_` | `&#95;` | Underscore |

The escaped format is `&#entity_number;`. Refer to the HTML escape standard for additional characters.

## Supported Code Block Languages

The rich text component supports specifying a programming language for code block syntax highlighting. Language names are case-insensitive.

`plain_text`, `abap`, `ada`, `apache`, `apex`, `assembly`, `bash`, `c_sharp`, `cpp`, `c`, `cmake`, `cobol`, `css`, `coffee_script`, `d`, `dart`, `delphi`, `diff`, `django`, `docker_file`, `erlang`, `fortran`, `gherkin`, `go`, `graphql`, `groovy`, `html`, `htmlbars`, `http`, `haskell`, `json`, `java`, `javascript`, `julia`, `kotlin`, `latex`, `lisp`, `lua`, `matlab`, `makefile`, `markdown`, `nginx`, `objective_c`, `opengl_shading_language`, `php`, `perl`, `powershell`, `prolog`, `properties`, `protobuf`, `python`, `r`, `ruby`, `rust`, `sas`, `scss`, `sql`, `scala`, `scheme`, `shell`, `solidity`, `swift`, `toml`, `thrift`, `typescript`, `vbscript`, `visual_basic`, `xml`, `yaml`

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
  "elements": [
    {
      "tag": "markdown",
      "text_size": "cus-0",
      "href": {
        "urlVal": {
          "url": "xxx1",
          "pc_url": "xxx2",
          "ios_url": "xxx3",
          "android_url": "xxx4"
        }
      },
      "content": "Plain text\nStandard emoji😁😢🌞💼🏆❌✅\n*Italic*\n**Bold**\n~~Strikethrough~~\nText link\nDifferentiated redirection\n<at id=all></at>"
    },
    {
      "tag": "hr"
    },
    {
      "tag": "markdown",
      "content": "The line above is a divider\n!hover_text\nThe line above is an image tag"
    }
  ],
  "header": {
    "template": "blue",
    "title": {
      "content": "This is the card title bar",
      "tag": "plain_text"
    }
  }
}
```

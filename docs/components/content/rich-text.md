# Rich Text (Markdown)

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/rich-text>

The rich text (Markdown) component in JSON 2.0 cards supports rendering headings, emojis, tables, images, code blocks, dividers, and more.

> **Note:** This document covers the JSON 2.0 structure. For the legacy JSON 1.0 structure, see [Rich Text (Markdown)](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/rich-text).

## Notes

The JSON 2.0 structure no longer supports the differentiated jump link syntax. Use the icon link syntax instead:

```
<link icon='chat_outlined' url='https://applink.larksuite.com/client/chat/xxxxx' pc_url='' ios_url='' android_url=''>Link text</link>
```

Deprecated syntax:

```json
{
  "tag": "markdown",
  "href": {
    "urlVal": {
      "url": "xxx",
      "pc_url": "xxx",
      "ios_url": "xxx",
      "android_url": "xxx"
    }
  },
  "content": "[Differentiated jump]($urlVal)"
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
        "content": "Person <person id='ou_449b53ad6aee526f7ed311b216aabcef' show_name=true show_avatar=true style='normal'></person>",
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
| `tag` | Yes | String | — | Fixed value `markdown`. |
| `element_id` | No | String | — | Unique component identifier (new in JSON 2.0). Must be globally unique within the card, letters/numbers/underscores only, start with a letter, max 20 chars. Used with the [component API](https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/cardkit-v1/card-element/create). |
| `margin` | No | String | `0` | Outer margin, range [-99,99]px. Accepts single `"10px"`, double `"4px 0"`, or quad `"4px 0 4px 0"`. |
| `content` | Yes | String | — | Markdown text content. See supported syntax below. |
| `text_align` | No | String | `left` | Text alignment: `left` / `center` / `right` |
| `text_size` | No | String | `normal` | Text size. See "Text Size Values" below. Also accepts custom size names. |
| `icon` | No | Object | — | Prefix icon. |
| `icon.tag` | No | String | — | `standard_icon` (icon library) or `custom_icon` (custom image). |
| `icon.token` | No | String | — | Icon token. Only effective when `tag` is `standard_icon`. Values listed in the [icon library](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/enumerations-for-icons). |
| `icon.color` | No | String | — | Icon color. Only effective when `tag` is `standard_icon`. Values listed in [color enumerations](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/enumerations-for-fields-related-to-color). |
| `icon.img_key` | No | String | — | Custom icon image key. Only effective when `tag` is `custom_icon`. Obtain via the [Upload Image](https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create) API. |

## Example

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "markdown",
        "content": "# Heading 1",
        "margin": "0px 0px 0px 0px",
        "text_align": "left",
        "text_size": "normal"
      },
      {
        "tag": "markdown",
        "content": "Standard emoji 😁😢🌞💼🏆❌✅\nLark emoji :OK::THUMBSUP:\n*italic* **bold** ~~strikethrough~~\n<font color='red'>Red text</font>\n<text_tag color=\"blue\">Tag</text_tag>\n[Text link](https://open.feishu.cn/document/server-docs/im-v1/message-reaction/emojis-introduce)\n<link icon='chat_outlined' url='https://open.feishu.cn' pc_url='' ios_url='' android_url=''>Link with icon</link>\n<at id=all></at>\n- Unordered item 1\n    - Nested item 1.1\n- Unordered item 2\n1. Ordered item 1\n    1. Nested item 1.1\n2. Ordered item 2\n```JSON\n{\"This is\": \"JSON demo\"}\n```"
      },
      {
        "tag": "markdown",
        "content": "Inline code `code`"
      },
      {
        "tag": "markdown",
        "content": "Number badge <number_tag background_color='grey' font_color='white' url='https://open.larksuite.com' pc_url='https://open.larksuite.com' android_url='https://open.larksuite.com' ios_url='https://open.larksuite.com'>1</number_tag>"
      },
      {
        "tag": "markdown",
        "content": "Default number badge <number_tag>1</number_tag>"
      },
      {
        "tag": "markdown",
        "content": "Person <person id='ou_449b53ad6aee526f7ed311b216a8f88f' show_name=true show_avatar=true style='normal'></person>"
      },
      {
        "tag": "markdown",
        "content": "> This is a blockquote\nLine break within quote\n"
      }
    ]
  }
}
```

---

## Supported Markdown Syntax

[Card JSON 2.0](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-structure) supports all standard Markdown syntax (except `HTMLBlock`) and some HTML syntax. For standard syntax, see [CommonMark Spec](https://spec.commonmark.org/0.31.2/) and [CommonMark playground](https://spec.commonmark.org/dingus/).

**Differences from CommonMark:**
- Single Enter = soft break (may be ignored by renderer); double Enter = hard break (always renders as new line).
- Supported HTML tags: `<br>` `<br/>` `<hr>` `<hr/>` `<person>` `<local_datetime>` `<at>` `<a>` `<text_tag>` `<raw>` `<link>` `<font>` (supports nesting).

---

### Line Breaks

```
Line 1<br />Line 2
```

In card JSON you can also use `\n`; in the building tool, use the Enter key.

### Italic / Bold / Strikethrough

```
*italic*
**bold** or __bold__
~~strikethrough~~
```

> Do not use 4 consecutive `*` or `_` — non-standard usage may cause rendering errors. If bold doesn't render, ensure there is a space before and after the syntax.

### @Mention

```html
<at id=open_id></at>
<at id=user_id></at>
<at ids=id_01,id_02,xxx></at>
<at email=test@email.com></at>
<at id=all></at>       <!-- @everyone — requires group owner permission -->
```

- Mentioned users receive a notification (except on forwarded cards).
- To display username/avatar/profile card without notification, use the [Person](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/user-profile) or [Person List](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/user-list) component.
- [Custom bots](https://open.larksuite.com/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN) only support `open_id` and `user_id`.
- See [How to get user IDs](https://open.larksuite.com/document/home/user-identity-introduction/open-id).

### Hyperlinks / Text Links

```markdown
[Open Platform](https://open.larksuite.com/)
<a href='https://open.larksuite.com'></a>
```

Links must include a schema (only HTTP/HTTPS supported). Link text color cannot be customized.

### Icon Links

```html
<link icon='chat_outlined' url='https://open.larksuite.com' pc_url='' ios_url='' android_url=''>Strategy meeting</link>
```

- `icon`: Icon token from the icon library, color is fixed to blue. Optional. See [icon library](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/enumerations-for-icons).
- `url`: Default link URL. Required.
- `pc_url` / `ios_url` / `android_url`: Per-platform URLs, take priority over `url`. Optional.

### Clickable Phone Number

```markdown
[Display text](tel://phone-number)
```

Only works on mobile.

### Colored Text

```html
<font color='red'>Red text</font>
<font color='green'>Green text</font>
```

- `color` supports `default`, all [color enumeration values](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/enumerations-for-fields-related-to-color), and RGBA custom colors.
- Does not apply to link text.
- `<font>` supports nesting other tags (`<at>` `<a>` `<link>` `<local_datetime>` `<font>`).

### Images

```markdown
![hover_text](image_key)
```

- `hover_text`: Text shown on hover (PC only).
- `image_key`: Obtained via the [Upload Image](https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create) API.

### Dividers

```markdown
<hr>
---
```

`<hr>` is recommended. The divider must be on its own line with line breaks before and after.

### Headings

```markdown
# Heading 1    (26px)
## Heading 2   (22px)
### Heading 3  (20px)
#### Heading 4 (18px)
##### Heading 5 (17px)
###### Heading 6 (14px)
```

### Blockquotes / Inline Code

```markdown
> This is a blockquote
`inline code`
```

### Lists

```markdown
- Unordered item 1
    - Nested item 1.1
- Unordered item 2

1. Ordered item 1
    1. Nested item 1.1
2. Ordered item 2
```

Items must start at the beginning of a line. 4 spaces = one indent level. In JSON, use `\n` for line breaks.

### Tables

```markdown
| Syntax | Description |
| -------- | -------- |
| Paragraph | Text |
```

- Maximum 5 data rows (excluding header); excess rows are paginated.
- Only supported in JSON 2.0; max 4 tables per rich text component.
- Column width etc. cannot be configured — use the [Table](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/table) component for more control.

### Code Blocks

````markdown
```JSON
{"This is": "JSON demo"}
```
````

Specify a language (case-insensitive); defaults to Plain Text if omitted. Four or more leading spaces ([indented code blocks](https://spec.commonmark.org/0.30/#indented-code-blocks)) also trigger code block rendering.

### Lark Emoji

```
:DONE:  :THUMBSUP:
```

See the [emoji reference](https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce) for supported emoji keys.

### Text Tags

```html
<text_tag color='red'>Tag text</text_tag>
```

`color` values: `neutral` `blue` `turquoise` `lime` `orange` `violet` `indigo` `wathet` `green` `yellow` `red` `purple` `carmine`

### Number Badges

```html
<number_tag>1</number_tag>
<number_tag background_color='grey' font_color='white' url='https://...' pc_url='...' android_url='...' ios_url='...'>99</number_tag>
```

Supports numbers 0-99. `background_color` / `font_color` / `url` / `pc_url` / `ios_url` / `android_url` are all optional.

### Person

```html
<person id='user_id' show_name=true show_avatar=true style='normal'></person>
```

- `id`: User ID (supports open_id / union_id / user_id). Shows "Unknown user" fallback if empty or invalid.
- `show_name`: Whether to show the username. Default `true`.
- `show_avatar`: Whether to show the avatar. Default `true`.
- `style`: `normal` (default) or `capsule`.

### Localized Date/Time

```html
<local_datetime millisecond='' format_type='date_num' link='https://www.feishu.com'></local_datetime>
```

- `millisecond`: Unix millisecond timestamp. Defaults to send/publish time if empty.
- `format_type`: Display format:
  - `date_num` — `2019-03-15`
  - `date_short` — `Mar 15`
  - `date` — `Mar 15, 2019`
  - `week` — `Tuesday`
  - `week_short` — `Tue`
  - `time` — `13:42`
  - `time_sec` — `13:42:53`
  - `timezone` — `GMT+8:00`
- `link`: URL to open when the time is clicked.

---

## Reference

### Special Character Escaping

Markdown special characters (`*` `~` `>` `<` etc.) must be HTML-escaped. Format: `&#entity;`. Full list: [HTML escape reference](https://www.w3school.com.cn/charsets/ref_html_8859.asp).

Common escapes: `&nbsp;`(non-breaking space) `&ensp;`(en space) `&emsp;`(em space) `&#62;`(>) `&#60;`(<) `&#42;`(\*) `&#91;`([) `&#93;`(]) `&#40;`(() `&#41;`()) `&#35;`(#) `&#96;`(\`) `&#92;`(\\) `&#47;`(/) `&#45;`(-) `&#33;`(!) `&#43;`(+) `&#58;`(:) `&#34;`(") `&#39;`(') `&#36;`($) `&#95;`(\_)

### Supported Code Block Languages

Case-insensitive: `plain_text` `abap` `ada` `apache` `apex` `assembly` `bash` `c` `c_sharp` `cpp` `cmake` `cobol` `coffee_script` `css` `d` `dart` `delphi` `diff` `django` `docker_file` `erlang` `fortran` `gherkin` `go` `graphql` `groovy` `haskell` `html` `htmlbars` `http` `java` `javascript` `json` `julia` `kotlin` `latex` `lisp` `lua` `makefile` `markdown` `matlab` `nginx` `objective_c` `opengl_shading_language` `perl` `php` `powershell` `prolog` `properties` `protobuf` `python` `r` `ruby` `rust` `sas` `scala` `scheme` `scss` `shell` `solidity` `sql` `swift` `thrift` `toml` `typescript` `vbscript` `visual_basic` `xml` `yaml`

### Text Size Values

| Value | Label | PC | Mobile |
|-------|-------|----|--------|
| `heading-0` | Extra-large heading | 30px | 26px |
| `heading-1` | Heading 1 | 24px | 24px |
| `heading-2` | Heading 2 | 20px | 20px |
| `heading-3` | Heading 3 | 18px | 17px |
| `heading-4` | Heading 4 | 16px | 16px |
| `heading` | Heading | 16px | 16px |
| `normal` | Body | 14px | 14px |
| `notation` | Caption | 12px | 12px |
| `xxxx-large` | — | 30px | 26px |
| `xxx-large` | — | 24px | 24px |
| `xx-large` | — | 20px | 20px |
| `x-large` | — | 18px | 18px |
| `large` | — | 16px | 17px |
| `medium` | — | 14px | 14px |
| `small` | — | 12px | 12px |
| `x-small` | — | 10px | 10px |

Unrecognized values fall back to `normal`.

### Custom Font Sizes for Mobile and Desktop

Define custom sizes in `config.style.text_size`, then reference them in component `text_size`:

```json
{
  "config": {
    "style": {
      "text_size": {
        "cus-0": {
          "default": "medium",
          "pc": "medium",
          "mobile": "large"
        }
      }
    }
  }
}
```

Usage in a component:

```json
{
  "tag": "markdown",
  "text_size": "cus-0",
  "content": "Custom font size text"
}
```

| Field | Required | Type | Description |
|-------|----------|------|-------------|
| `text_size.{name}` | No | Object | Custom size object. Name is user-defined (e.g. `cus-0`). |
| `text_size.{name}.default` | No | String | Fallback size for older clients. Recommended. |
| `text_size.{name}.pc` | No | String | Desktop font size. |
| `text_size.{name}.mobile` | No | String | Mobile font size (some values differ from PC). |

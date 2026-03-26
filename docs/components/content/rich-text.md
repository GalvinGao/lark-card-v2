# 富文本 (Markdown)

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/rich-text>

JSON 2.0 结构卡片的富文本（Markdown）组件支持渲染标题、表情、表格、图片、代码块、分割线等元素。

> **注意**：本文档介绍富文本组件的 JSON 2.0 结构。历史 JSON 1.0 结构参考[富文本（Markdown）](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/rich-text)。

## 注意事项

富文本 JSON 2.0 结构不再支持差异化跳转语法。请改用含图标的链接语法：

```
<link icon='chat_outlined' url='https://applink.larksuite.com/client/chat/xxxxx' pc_url='' ios_url='' android_url=''>差异化链接</link>
```

已废弃的语法：

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
  "content": "[差异化跳转]($urlVal)"
}
```

## JSON 结构

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "markdown",
        "element_id": "custom_id",
        "margin": "0px 0px 0px 0px",
        "content": "人员<person id='ou_449b53ad6aee526f7ed311b216aabcef' show_name=true show_avatar=true style='normal'></person>",
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

## 字段说明

| 字段 | 必填 | 类型 | 默认值 | 说明 |
|------|------|------|--------|------|
| `tag` | 是 | String | — | 固定值 `markdown` |
| `element_id` | 否 | String | — | 组件唯一标识（JSON 2.0 新增）。同一卡片内全局唯一，仅字母/数字/下划线，字母开头，≤20 字符。用于[组件相关接口](https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/cardkit-v1/card-element/create)。 |
| `margin` | 否 | String | `0` | 外边距，范围 [-99,99]px。支持单值 `"10px"`、双值 `"4px 0"`、四值 `"4px 0 4px 0"`。 |
| `content` | 是 | String | — | Markdown 文本内容。支持的语法见下文。 |
| `text_align` | 否 | String | `left` | 对齐方式：`left` / `center` / `right` |
| `text_size` | 否 | String | `normal` | 文本大小，可选值见下方「文本字号枚举」。也支持自定义字号名称。 |
| `icon` | 否 | Object | — | 前缀图标 |
| `icon.tag` | 否 | String | — | `standard_icon`（图标库）或 `custom_icon`（自定义图片） |
| `icon.token` | 否 | String | — | 图标 token，`tag` 为 `standard_icon` 时生效。枚举值见[图标库](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/enumerations-for-icons)。 |
| `icon.color` | 否 | String | — | 图标颜色，`tag` 为 `standard_icon` 时生效。枚举值见[颜色枚举值](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/enumerations-for-fields-related-to-color)。 |
| `icon.img_key` | 否 | String | — | 自定义图标的图片 key，`tag` 为 `custom_icon` 时生效。通过[上传图片](https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create)接口获取。 |

## 示例

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
        "content": "标准emoji 😁😢🌞💼🏆❌✅\nLarkemoji :OK::THUMBSUP:\n*斜体* **粗体** ~~删除线~~ \n这是红色文本</font>\n<text_tag color=\"blue\">标签</text_tag>\n[文字链接](https://open.feishu.cn/document/server-docs/im-v1/message-reaction/emojis-introduce)\n<link icon='chat_outlined' url='https://open.feishu.cn' pc_url='' ios_url='' android_url=''>带图标的链接</link>\n<at id=all></at>\n- 无序列表1\n    - 无序列表 1.1\n- 无序列表2\n1. 有序列表1\n    1. 有序列表 1.1\n2. 有序列表2\n```JSON\n{\"This is\": \"JSON demo\"}\n```"
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

---

## 支持的 Markdown 语法

[卡片 JSON 2.0 结构](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-structure)支持除 `HTMLBlock` 外所有标准 Markdown 语法及部分 HTML 语法。标准语法参考 [CommonMark Spec](https://spec.commonmark.org/0.31.2/) 和 [CommonMark playground](https://spec.commonmark.org/dingus/)。

**与 CommonMark 的差异：**
- 一个 Enter = 软换行（可能被忽略）；两个 Enter = 硬换行（始终新行）。
- 支持的 HTML 标签：`<br>` `<br/>` `<hr>` `<hr/>` `<person>` `<local_datetime>` `<at>` `<a>` `<text_tag>` `<raw>` `<link>` `<font>`（支持嵌套）。

---

### 换行

```
第一行<br />第二行
```

在卡片 JSON 中也可用 `\n` 换行；在搭建工具中可用回车键。

### 斜体 / 粗体 / 删除线

```
*斜体*
**粗体** 或 __粗体__
~~删除线~~
```

> 不要连续使用 4 个 `*` 或 `_`，语法不规范可能导致渲染错误。粗体效果未显示时，确保前后保留空格。

### @指定人

```html
<at id=open_id></at>
<at id=user_id></at>
<at ids=id_01,id_02,xxx></at>
<at email=test@email.com></at>
<at id=all></at>       <!-- @所有人，需群主开启权限 -->
```

- 被 @ 的用户将收到提及通知（转发的卡片除外）。
- 要展示用户名/头像/名片但不发通知，用[人员](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/user-profile)或[人员列表](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/user-list)组件。
- [自定义机器人](https://open.larksuite.com/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN)仅支持 `open_id`、`user_id`。
- 获取 user_id / open_id 参考[如何获取不同的用户 ID](https://open.larksuite.com/document/home/user-identity-introduction/open-id)。

### 超链接 / 文字链接

```markdown
[开放平台](https://open.larksuite.com/)
<a href='https://open.larksuite.com'></a>
```

超链接必须包含 schema（仅支持 HTTP/HTTPS），文本颜色不支持自定义。

### 含图标的链接

```html
<link icon='chat_outlined' url='https://open.larksuite.com' pc_url='' ios_url='' android_url=''>战略研讨会</link>
```

- `icon`：图标库中的图标 token，颜色固定蓝色。可选。枚举值见[图标库](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/enumerations-for-icons)。
- `url`：默认链接地址。必填。
- `pc_url` / `ios_url` / `android_url`：各端链接，优先级高于 `url`。可选。

### 可点击的电话号码

```markdown
[显示文案](tel://电话号码)
```

仅在移动端生效。

### 彩色文本

```html
<font color='red'>红色文本</font>
<font color='green'>绿色文本</font>
```

- `color` 支持 `default` 及所有[颜色枚举值](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/enumerations-for-fields-related-to-color)和 RGBA 自定义颜色。
- 不支持对链接文本生效。
- `<font>` 支持嵌套其他标签（`<at>` `<a>` `<link>` `<local_datetime>` `<font>`）。

### 图片

```markdown
![hover_text](image_key)
```

- `hover_text`：PC 端悬浮时展示的文案。
- `image_key`：通过[上传图片](https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create)接口获取。

### 分割线

```markdown
<hr>
---
```

推荐使用 `<hr>`。分割线必须单独一行使用，前后需有换行符。

### 标题

```markdown
# 一级标题    (26px)
## 二级标题   (22px)
### 三级标题  (20px)
#### 四级标题 (18px)
##### 五级    (17px)
###### 六级   (14px)
```

### 引用 / 行内引用

```markdown
> 这是一段引用文字
`行内代码`
```

### 列表

```markdown
- 无序列表1
    - 无序列表 1.1
- 无序列表2

1. 有序列表1
    1. 有序列表 1.1
2. 有序列表2
```

序号需在行首使用，4 个空格 = 一层缩进。JSON 中需添加 `\n` 换行符。

### 表格

```markdown
| Syntax | Description |
| -------- | -------- |
| Paragraph | Text |
```

- 除标题行外，最多展示五行数据，超出分页展示。
- 仅支持 JSON 2.0 结构；单个富文本组件最多 4 个表格。
- 不支持设置列宽等，需要更多控制请用[表格](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/table)组件。

### 代码块

````markdown
```JSON
{"This is": "JSON demo"}
```
````

支持指定编程语言（大小写不敏感），未指定默认 Plain Text。四个及以上空格的[缩进式代码块](https://spec.commonmark.org/0.30/#indented-code-blocks)也会触发代码块效果。

### Lark 表情

```
:DONE:  :THUMBSUP:
```

支持的 Emoji Key 列表参考[表情文案说明](https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)。

### 标签

```html
<text_tag color='red'>标签文本</text_tag>
```

`color` 枚举值：`neutral` `blue` `turquoise` `lime` `orange` `violet` `indigo` `wathet` `green` `yellow` `red` `purple` `carmine`

### 数字角标

```html
<number_tag>1</number_tag>
<number_tag background_color='grey' font_color='white' url='https://...' pc_url='...' android_url='...' ios_url='...'>99</number_tag>
```

支持 0-99 之间的数字。`background_color` / `font_color` / `url` / `pc_url` / `ios_url` / `android_url` 均为可选。

### 人员

```html
<person id='user_id' show_name=true show_avatar=true style='normal'></person>
```

- `id`：用户 ID（支持 open_id / union_id / user_id）。不填或错误时展示"未知用户"。
- `show_name`：是否展示用户名。默认 `true`。
- `show_avatar`：是否展示头像。默认 `true`。
- `style`：`normal`（默认）或 `capsule`（胶囊样式）。

### 国际化时间

```html
<local_datetime millisecond='' format_type='date_num' link='https://www.feishu.com'></local_datetime>
```

- `millisecond`：Unix 毫秒时间戳。不填则默认使用发送/发布时间。
- `format_type`：时间格式，枚举值：
  - `date_num` — `2019-03-15`
  - `date_short` — `3月15日` / `Mar 15`
  - `date` — `2019年3月15日` / `Mar 15, 2019`
  - `week` — `星期二` / `Tuesday`
  - `week_short` — `周二` / `Tue`
  - `time` — `13:42`
  - `time_sec` — `13:42:53`
  - `timezone` — `GMT+8:00`
- `link`：点击时间时跳转的链接。

---

## 参考

### 特殊字符转义

命中 Markdown 特殊字符（`*` `~` `>` `<` 等）时需 HTML 转义。格式：`&#实体编号;`。完整列表参考 [HTML 转义标准](https://www.w3school.com.cn/charsets/ref_html_8859.asp)。

常用转义：`&nbsp;`(不换行空格) `&ensp;`(半角空格) `&emsp;`(全角空格) `&#62;`(>) `&#60;`(<) `&#42;`(\*) `&#91;`([) `&#93;`(]) `&#40;`(() `&#41;`()) `&#35;`(#) `&#96;`(\`) `&#92;`(\\) `&#47;`(/) `&#45;`(-) `&#33;`(!) `&#43;`(+) `&#58;`(:) `&#34;`(") `&#39;`(') `&#36;`($) `&#95;`(\_)

### 代码块支持的编程语言

大小写不敏感：`plain_text` `abap` `ada` `apache` `apex` `assembly` `bash` `c` `c_sharp` `cpp` `cmake` `cobol` `coffee_script` `css` `d` `dart` `delphi` `diff` `django` `docker_file` `erlang` `fortran` `gherkin` `go` `graphql` `groovy` `haskell` `html` `htmlbars` `http` `java` `javascript` `json` `julia` `kotlin` `latex` `lisp` `lua` `makefile` `markdown` `matlab` `nginx` `objective_c` `opengl_shading_language` `perl` `php` `powershell` `prolog` `properties` `protobuf` `python` `r` `ruby` `rust` `sas` `scala` `scheme` `scss` `shell` `solidity` `sql` `swift` `thrift` `toml` `typescript` `vbscript` `visual_basic` `xml` `yaml`

### 文本字号枚举

| 枚举值 | 含义 | PC 端 | 移动端 |
|--------|------|-------|--------|
| `heading-0` | 特大标题 | 30px | 26px |
| `heading-1` | 一级标题 | 24px | 24px |
| `heading-2` | 二级标题 | 20px | 20px |
| `heading-3` | 三级标题 | 18px | 17px |
| `heading-4` | 四级标题 | 16px | 16px |
| `heading` | 标题 | 16px | 16px |
| `normal` | 正文 | 14px | 14px |
| `notation` | 辅助信息 | 12px | 12px |
| `xxxx-large` | — | 30px | 26px |
| `xxx-large` | — | 24px | 24px |
| `xx-large` | — | 20px | 20px |
| `x-large` | — | 18px | 18px |
| `large` | — | 16px | 17px |
| `medium` | — | 14px | 14px |
| `small` | — | 12px | 12px |
| `x-small` | — | 10px | 10px |

填写其它值时，卡片将以 `normal` 字号展示。

### 为移动端和桌面端定义不同字号

在 `config.style.text_size` 中添加自定义字号，然后在组件的 `text_size` 属性中引用：

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

在组件中使用：

```json
{
  "tag": "markdown",
  "text_size": "cus-0",
  "content": "自定义字号文本"
}
```

| 字段 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `text_size.{name}` | 否 | Object | 自定义字号对象，名称自定义（如 `cus-0`） |
| `text_size.{name}.default` | 否 | String | 旧版客户端兜底字号。建议填写。 |
| `text_size.{name}.pc` | 否 | String | 桌面端字号 |
| `text_size.{name}.mobile` | 否 | String | 移动端字号（部分枚举值大小与 PC 端有差异） |

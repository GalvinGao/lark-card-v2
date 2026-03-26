# Title

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/title>

The title component supports adding a main title, subtitle, suffix tags, and a title icon.

Only one title component can be added to the same card.

## JSON Structure

```json
{
  "schema": "2.0",
  "header": {
    "title": {
      "tag": "plain_text",
      "content": "Example Title"
    },
    "subtitle": {
      "tag": "plain_text",
      "content": "Example Text"
    },
    "text_tag_list": [
      {
        "tag": "text_tag",
        "element_id": "custom_id",
        "text": {
          "tag": "plain_text",
          "content": "Tag 1"
        },
        "color": "neutral"
      }
    ],
    "i18n_text_tag_list": {
      "zh_cn": [],
      "en_us": [],
      "ja_jp": [],
      "zh_hk": [],
      "zh_tw": []
    },
    "template": "blue",
    "icon": {
      "tag": "standard_icon",
      "token": "chat-forbidden_outlined",
      "color": "orange",
      "img_key": "img_v2_38811724"
    },
    "padding": "12px 8px 12px 8px"
  }
}
```

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `title` | Yes | Object | — | Main title of the card. If only the subtitle is configured, it will be displayed as the main title. |
| `title.tag` | Yes | String | — | Text type tag. Values: `plain_text` (plain text or emoji) or `lark_md` (partial Markdown syntax). |
| `title.content` | No | String | — | Main title content. Up to four lines; content exceeding four lines is truncated with `...`. To configure in multiple languages, refer to the multilingual card configuration document. |
| `subtitle` | No | Object | — | Subtitle of the card. If only the subtitle is configured, it will be displayed as the main title. |
| `subtitle.tag` | Yes | String | — | Text type tag. Values: `plain_text` or `lark_md`. |
| `subtitle.content` | No | String | — | Subtitle content. Up to one line; content exceeding one line is truncated with `...`. |
| `text_tag_list` | No | TextTagList | — | Suffix tags for the title. Up to 3 tags; excess tags are not displayed. Display order matches array order. Cannot be used together with `i18n_text_tag_list`; if both are configured, only `i18n_text_tag_list` takes effect. |
| `text_tag_list[].tag` | Yes | String | — | Fixed value: `text_tag`. |
| `text_tag_list[].element_id` | No | String | — | Unique identifier for the component (JSON 2.0). Must be globally unique within the card. Only letters, numbers, and underscores; must start with a letter and not exceed 20 characters. |
| `text_tag_list[].text` | No | Text Object | — | Content of the suffix tag, defined in `plain_text` mode. |
| `text_tag_list[].color` | No | String | `"blue"` | Color of the suffix tag. See suffix tag color enumeration below. |
| `i18n_text_tag_list` | No | Object | — | Multilingual suffix tags. Up to 3 tags per language. Supported locales: `zh_cn`, `en_us`, `ja_jp`, `zh_hk`, `zh_tw`, `id_id`, `vi_vn`, `th_th`, `pt_br`, `es_es`, `ko_kr`, `de_de`, `fr_fr`, `it_it`, `ru_ru`, `ms_my`. |
| `template` | No | String | `"default"` | Title theme color. See theme style enumeration below. |
| `icon` | No | Object | — | Prefix icon. Supports custom or icon library usage. |
| `icon.tag` | No | String | — | Icon type. Values: `standard_icon` (icon library) or `custom_icon` (custom image). |
| `icon.token` | No | String | — | Icon token from the icon library. Effective when `icon.tag` is `standard_icon`. |
| `icon.color` | No | String | — | Icon color. Supports line and surface icons (tokens ending in `outlined` or `filled`). Effective when `icon.tag` is `standard_icon`. |
| `icon.img_key` | No | String | — | Image key for a custom prefix icon. Effective when `icon.tag` is `custom_icon`. Obtain via the Upload Image API. |
| `padding` | No | String | `"12px"` | Padding of the title component (JSON 2.0). Value range is [0,99]px. Accepts a single value (e.g. `"10px"`), double value (e.g. `"4px 0"`), or four values (e.g. `"4px 0 4px 0"`) for top, right, bottom, left. |

## Enumerations

### Title Theme Style (`template`)

| Value | Description |
|-------|-------------|
| `blue` | Blue theme |
| `wathet` | Wathet theme |
| `turquoise` | Turquoise theme |
| `green` | Green theme |
| `yellow` | Yellow theme |
| `orange` | Orange theme |
| `red` | Red theme |
| `carmine` | Carmine theme |
| `violet` | Violet theme |
| `purple` | Purple theme |
| `indigo` | Indigo theme |
| `grey` | Grey theme |
| `default` | Default theme |

### Suffix Tag Color (`color`)

| Value |
|-------|
| `neutral` |
| `blue` |
| `turquoise` |
| `lime` |
| `orange` |
| `violet` |
| `indigo` |
| `wathet` |
| `green` |
| `yellow` |
| `red` |
| `purple` |
| `carmine` |

### Icon

For the enumeration of `icon.token`, refer to the Icon Library. For `icon.color`, refer to the Color Enumeration Values.

## Example

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "div",
        "text": {
          "tag": "plain_text",
          "content": "示例文本"
        }
      }
    ]
  },
  "header": {
    "title": {
      "tag": "lark_md",
      "content": ":Partying:卡片主标题 "
    },
    "subtitle": {
      "tag": "plain_text",
      "content": "卡片副标题"
    },
    "text_tag_list": [
      {
        "tag": "text_tag",
        "text": {
          "tag": "plain_text",
          "content": "标签 1"
        },
        "color": "blue"
      },
      {
        "tag": "text_tag",
        "text": {
          "tag": "plain_text",
          "content": "标签 2"
        },
        "color": "turquoise"
      },
      {
        "tag": "text_tag",
        "text": {
          "tag": "plain_text",
          "content": "标签 3"
        },
        "color": "orange"
      }
    ],
    "template": "blue",
    "icon": {
      "tag": "standard_icon",
      "token": "larkcommunity_colorful"
    },
    "padding": "12px"
  }
}
```

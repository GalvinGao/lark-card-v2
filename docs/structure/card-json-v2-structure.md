# Card JSON 2.0 Structure

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-structure>

This document introduces the overall structure and attribute descriptions of Card JSON 2.0. There are many incompatible differences and new attributes between the structure of JSON 2.0 and JSON 1.0. For more details, refer to Card JSON 2.0 Incompatible Changes & Updates.

## Concept Description

Card JSON 2.0 refers to the version in which the `schema` attribute is declared as `"2.0"` in the card JSON data. Compared to version 1.0, version 2.0 has many incompatible differences and new attributes. For details, refer to Card JSON 2.0 version update notes.

## Restrictions

- The Card JSON 2.0 structure does not support building and generating through the Lark Card Builder Tool. It only supports implementation through Card JSON code.
- The Card JSON 2.0 structure supports up to 200 elements (e.g., text elements where tag is `plain_text`) or components for a card.
- Card Schema 2.0 structure is supported from Lark client version 7.20 onwards. When a card using Schema v2.0 structure is sent to clients with versions lower than 7.20, the card title will display correctly, but the content will show fallback upgrade prompts.
- Currently, Lark does not support the new Feishu card builder tool.

## JSON Structure

The following is the overall structure of Card JSON 2.0.

```json
{
    "schema": "2.0",
    "config": {
        "streaming_mode": true,
        "streaming_config": {},
        "summary": {
            "content": "Custom content",
            "i18n_content": {
                "zh_cn": "",
                "en_us": "",
                "ja_jp": ""
            }
        },
        "locales": [
            "en_us",
            "ja_jp"
        ],
        "enable_forward": true,
        "update_multi": true,
        "width_mode": "fill",
        "use_custom_translation": false,
        "enable_forward_interaction": false,
        "style": {
            "text_size": {
                "cus-0": {
                    "default": "medium",
                    "pc": "medium",
                    "mobile": "large"
                }
            },
            "color": {
                "cus-0": {
                    "light_mode": "rgba(5,157,178,0.52)",
                    "dark_mode": "rgba(78,23,108,0.49)"
                }
            }
        }
    },
    "card_link": {
        "url": "https://www.baidu.com",
        "android_url": "https://developer.android.com/",
        "ios_url": "https://developer.apple.com/",
        "pc_url": "https://www.windows.com"
    },
    "header": {
        "title": {
            "tag": "plain_text",
            "content": "Example title"
        },
        "subtitle": {
            "tag": "plain_text",
            "content": "Example text"
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
    },
    "body": {
        "direction": "vertical",
        "padding": "12px 8px 12px 8px",
        "horizontal_spacing": "3px",
        "horizontal_align": "left",
        "vertical_spacing": "4px",
        "vertical_align": "center",
        "elements": [
            {
                "tag": "xxx",
                "margin": "4px",
                "element_id": "custom_id"
            }
        ]
    }
}
```

**Key notes on the JSON structure:**

- `schema`: The version of the card JSON structure. The default is `"1.0"`. To use the JSON 2.0 structure, you must explicitly declare `"2.0"`.
- `config.streaming_mode`: Whether the card is in streaming update mode. Default is `false`.
- `config.streaming_config`: Streaming update configuration.
- `config.summary`: Card summary information. Can be used to customize the display text in the client chat bar message preview. If streaming update mode is enabled, this defaults to "Generating." but can be customized.
- `config.summary.i18n_content`: Multi-language configuration of summary information. Refer to the multi-language configuration document for cards.
- `config.locales`: New attribute in JSON 2.0. Used to specify effective languages. If configured, only the languages in locales will be effective.
- `config.enable_forward`: Whether to support forwarding the card. Default is `true`.
- `config.update_multi`: Whether it is a shared card. Default is `true`. JSON 2.0 currently only supports setting it to `true`, meaning that updating the card's content is visible to all recipients.
- `config.width_mode`: Card width mode. Supports `"compact"` (400px), `"fill"` (fills chat window width), or default (600px).
- `config.use_custom_translation`: Whether to use custom translation data. Default is `false`. When `true`, after the user clicks message translation, the corresponding target language of i18n is used as the translation result.
- `config.enable_forward_interaction`: Whether the forwarded card still supports interaction. Default is `false`.
- `config.style.text_size`: Add custom font sizes for mobile and desktop, with a fallback font size. Supports adding multiple custom font size objects.
- `config.style.color`: Add RGBA syntax for light and dark themes. Supports adding multiple custom color objects.
- `card_link`: Specify the overall jump link of the card.
- `header.title`: Main title of the card (required). Tag optional values: `plain_text` and `lark_md`.
- `header.subtitle`: Subtitle of the card (optional). Tag optional values: `plain_text` and `lark_md`.
- `header.text_tag_list`: Suffix tags for the title, up to 3 tags (optional).
- `header.i18n_text_tag_list`: Multi-language suffix tags for the title, up to 3 per language. If both the original field and the i18n field are configured, the multi-language configuration takes precedence.
- `header.template`: Title theme style color. Supports `"blue"` | `"wathet"` | `"turquoise"` | `"green"` | `"yellow"` | `"orange"` | `"red"` | `"carmine"` | `"violet"` | `"purple"` | `"indigo"` | `"grey"` | `"default"`. Default is `"default"`.
- `header.icon`: Prefix icon. `tag` is the icon type. `token` is the icon token (only for `standard_icon`). `color` is the icon color (only for `standard_icon`). `img_key` is the image key (only for `custom_icon`).
- `header.padding`: Padding of the title component. New in JSON 2.0. Default is `"12px"`, supports range [0,99]px.
- `body`: Card body.
- `body.direction`: Arrangement direction. `"vertical"` (default) or `"horizontal"`.
- `body.padding`: Padding, supports range [0,99]px.
- `body.horizontal_spacing`: Horizontal spacing. Values: `"small"` (4px), `"medium"` (8px), `"large"` (12px), `"extra_large"` (16px), or [0,99]px.
- `body.horizontal_align`: Horizontal alignment. `"left"` (default), `"center"`, or `"right"`.
- `body.vertical_spacing`: Vertical spacing. Values: `"small"` (4px), `"medium"` (8px), `"large"` (12px), `"extra_large"` (16px), or [0,99]px.
- `body.vertical_align`: Vertical alignment. `"top"` (default), `"center"`, or `"bottom"`.
- `body.elements`: JSON data of each component, arranged in array order.
- `body.elements[].tag`: Tag of the component.
- `body.elements[].margin`: Margin, default `"0"`, range [-99,99]px. New in JSON 2.0.
- `body.elements[].element_id`: Unique identifier for the component. New in JSON 2.0. Must be globally unique within the same card, letters/numbers/underscores only, must start with a letter, max 20 characters.

## Field Descriptions

This section provides detailed descriptions of each field in the card structure.

### Global Properties

Global properties of the card include the following fields.

```json
{
    "schema": "2.0",
    "config": {},
    "card_link": {},
    "header": {},
    "body": {
        "elements": []
    }
}
```

If none of these fields are provided, the card JSON will be `"{}"`. The Lark Open Platform supports sending a blank card with the card JSON as `"{}"`.

| Field | Required | Description |
|-------|----------|-------------|
| `schema` | No | Version declaration of the card structure. Default is `"1.0"`. Values: `"1.0"` (Card v1.0 structure) or `"2.0"` (Card v2.0 structure with more fields and capabilities such as streaming updates and extended Markdown syntax). |
| `config` | No | Used to configure the global behavior of the card, including streaming update mode, whether forwarding is allowed, whether it is a shared card, etc. |
| `card_link` | No | Used to specify the overall jump link of the card. You can configure a default link, or configure different jump links for PC, Android, and iOS. |
| `header` | No | Title component configuration. For details, refer to the Title component. |
| `body` | No | The card body, containing an `elements` array used to place various components. |

### Card Global Behavior Settings (`config`)

The `config` field is used to configure the global behaviors of the card, including whether the card can be forwarded and whether it is a shared card.

```json
{
    "config": {
        "streaming_mode": true,
        "summary": {
            "content": "Custom content",
            "i18n_content": {
                "zh_cn": "",
                "en_us": "",
                "ja_jp": ""
            }
        },
        "enable_forward": true,
        "update_multi": true,
        "width_mode": "fill",
        "use_custom_translation": false,
        "enable_forward_interaction": false,
        "style": {
            "text_size": {
                "cus-0": {
                    "default": "medium",
                    "pc": "medium",
                    "mobile": "large"
                }
            },
            "color": {
                "cus-0": {
                    "light_mode": "rgba(5,157,178,0.52)",
                    "dark_mode": "rgba(78,23,108,0.49)"
                }
            }
        }
    }
}
```

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `streaming_mode` | No | Boolean | `false` | Whether the card is in streaming update mode. Refer to Streaming updates overview for details. |
| `summary` | No | Object | — | Card summary information. Can be used to customize the display text in the client chat bar message preview. |
| `enable_forward` | No | Boolean | `true` | Whether to allow card forwarding. `true`: allow, `false`: do not allow. |
| `update_multi` | No | Boolean | `true` | Whether it is a shared card. `true`: shared card (updates visible to all recipients), `false`: non-shared card (only the operating user sees updates). |
| `width_mode` | No | String | `"default"` | Card width mode. `"default"`: max width 600px on PC/iPad. `"compact"`: compact width 400px. `"fill"`: adaptive screen width. Note: not supported in the Card Construction Tool. |
| `use_custom_translation` | No | Boolean | `false` | Whether to use custom translation data. `true`: use i18n target language as translation result (falls back to machine translation if unavailable). `false`: use machine translation directly. |
| `enable_forward_interaction` | No | Boolean | `false` | Whether forwarded cards still support interactive feedback. |
| `style` | No | Object | — | Add custom font size and color. Can be applied to component JSON data to set font size and color properties. |

### Card Global Jump Link (`card_link`)

The `card_link` field is used to specify the overall click-through link for the card. You can configure a default link or separate links for PC, Android, and iOS platforms.

```json
{
    "card_link": {
        "url": "https://www.baidu.com",
        "android_url": "https://developer.android.com/",
        "ios_url": "https://developer.apple.com/",
        "pc_url": "https://www.windows.com"
    }
}
```

> **Note:** Either `url` or platform-specific links (`android_url`, `ios_url`, `pc_url`) must be provided. If `url` is not provided, then all three platform-specific links must be provided. If both `url` and platform-specific links are provided, `url` takes effect. To disable redirection for a specific platform, set the corresponding parameter value to `lark://msgcard/unsupported_action`.

| Field | Required | Type | Description |
|-------|----------|------|-------------|
| `url` | No | String | Default link address. |
| `pc_url` | No | String | Link address for PC platform. |
| `ios_url` | No | String | Link address for iOS platform. |
| `android_url` | No | String | Link address for Android platform. |

### Card Title (`header`)

The `header` field is used to configure the title of the card. For details on the header field, refer to the Title component documentation.

```json
{
    "header": {}
}
```

### Card Body (`body`)

In the `body` field of the card, you add card components as the content of the card. Components are vertically arranged in streaming fashion according to the array order. For details on card components, refer to the component JSON v2.0 overview.

In the card JSON 2.0 structure, all components have a new `element_id` attribute, which serves as a unique identifier for interacting with components.

```json
{
    "body": {
        "elements": [
            {
                "tag": "xxx",
                "element_id": "custom_id"
            }
        ]
    }
}
```

- `element_id`: Must be globally unique within the same card. Only letters, numbers, and underscores are allowed. Must start with a letter and not exceed 20 characters.

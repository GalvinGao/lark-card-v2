# Overflow

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/overflow>

The overflow (folding button group) component supports adding multiple buttons and folding them into a single trigger. Clicking on the button group displays all buttons within the group.

This document describes the JSON 2.0 structure. For the historical JSON 1.0 structure, refer to Folding Button Group.

## Notes

- The `confirm` field only triggers the secondary confirmation popup when a button containing a submit attribute is clicked.
- Within the same overflow component, each option's `value` should be unique for proper identification.

## Nesting Rules

The overflow component supports nesting within form containers, folding panels, loop containers, interactive containers, and column components. In the card building tool, it does not support nesting within interactive containers.

## JSON Structure

```json
{
  "tag": "overflow",
  "element_id": "custom_id",
  "margin": "0px 0px 0px 0px",
  "width": "fill",
  "options": [
    {
      "text": {
        "tag": "plain_text",
        "content": "This is a link jump"
      },
      "multi_url": {
        "url": "/ssl:ttdoc/home/index",
        "pc_url": "",
        "ios_url": "",
        "android_url": ""
      },
      "value": "document"
    }
  ],
  "value": {
    "key_1": "value_1"
  },
  "confirm": {
    "title": {
      "tag": "plain_text",
      "content": "title"
    },
    "text": {
      "tag": "plain_text",
      "content": "content"
    }
  }
}
```

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | Component tag. Fixed value: `overflow`. |
| `element_id` | No | String | — | Unique identifier for the component (JSON 2.0). Must be globally unique within the card. Only letters, numbers, and underscores allowed; must start with a letter; max 20 characters. |
| `margin` | No | String | `0` | Margin of the component (JSON 2.0). Value range is `[-99,99]px`. Accepts a single value (e.g. `"10px"`), double value (e.g. `"4px 0"`), or four values (e.g. `"4px 0 4px 0"`). |
| `width` | No | String | `default` | Component width. `default`: default width. `fill`: fills card width. `[100,∞)px`: custom width (clips to card width if exceeded). |
| `options` | Yes | Array | — | Option buttons in the overflow group. See the options fields below. |
| `value` | No | Object | — | Callback data for the entire component. Returned to the server when users click the overflow button. Must be a key-value JSON structure with String keys. |
| `confirm` | No | Object | — | Secondary confirmation popup configuration. Only triggers when a button with a submit attribute is clicked. |
| `confirm.title` | Yes | Object | — | Popup title. |
| `confirm.title.tag` | Yes | String | `plain_text` | Popup title tag. Fixed value: `plain_text`. |
| `confirm.title.content` | Yes | String | — | Popup title content. |
| `confirm.text` | Yes | Object | — | Popup body text. |
| `confirm.text.tag` | Yes | String | `plain_text` | Popup body tag. Fixed value: `plain_text`. |
| `confirm.text.content` | Yes | String | — | Popup body content. |

### Options Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `text` | No | Object | — | Text on the button. |
| `text.tag` | No | String | `plain_text` | Text tag. Fixed value: `plain_text`. |
| `text.content` | No | String | — | Text content. Supports up to 100 characters. |
| `multi_url` | No | Object | — | Jump links for multiple platforms. |
| `multi_url.url` | No | String | — | Default jump link. |
| `multi_url.android_url` | No | String | — | Android jump link. Set to `lark://msgcard/unsupported_action` to disallow jumping on this platform. |
| `multi_url.ios_url` | No | String | — | iOS jump link. Set to `lark://msgcard/unsupported_action` to disallow jumping on this platform. |
| `multi_url.pc_url` | No | String | — | PC jump link. Set to `lark://msgcard/unsupported_action` to disallow jumping on this platform. |
| `value` | No | String | — | Return value for this button. When users click the option, the application returns this value to the card request address. |

## Callback Structure

After configuring interactions for the component, the request address configured in the developer backend will receive callback data when a user interacts with the component.

- If you use the new card callback interaction (`card.action.trigger`), refer to Card Callback Interaction for the callback structure.
- If you use the old card callback interaction (`card.action.trigger_v1`), refer to Message Card Callback Interaction (Old) for the callback structure.

## Example

```json
{
  "schema": "2.0",
  "body": {
    "direction": "vertical",
    "padding": "12px 12px 12px 12px",
    "elements": [
      {
        "tag": "column_set",
        "columns": [
          {
            "tag": "column",
            "width": "auto",
            "elements": [
              {
                "tag": "button",
                "text": {
                  "tag": "plain_text",
                  "content": "按钮 1"
                },
                "type": "default",
                "size": "medium"
              },
              {
                "tag": "button",
                "text": {
                  "tag": "plain_text",
                  "content": "按钮 2"
                },
                "type": "default",
                "size": "medium"
              },
              {
                "tag": "button",
                "text": {
                  "tag": "plain_text",
                  "content": "按钮 3"
                },
                "type": "default",
                "size": "medium"
              },
              {
                "tag": "overflow",
                "width": "",
                "options": [
                  {
                    "text": {
                      "tag": "plain_text",
                      "content": "按钮 4"
                    }
                  },
                  {
                    "text": {
                      "tag": "plain_text",
                      "content": "按钮 5"
                    }
                  }
                ]
              }
            ],
            "direction": "horizontal",
            "vertical_align": "top"
          }
        ],
        "margin": "0px 0px 0px 0px"
      }
    ]
  }
}
```

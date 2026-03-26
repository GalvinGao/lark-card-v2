# Single Select Dropdown Menu

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/single-select-dropdown-menu>

The single select dropdown menu component supports customizing the option text, icon, and callback parameters. It is an interactive component.

This document describes the JSON 2.0 structure. For the historical JSON 1.0 structure, refer to Dropdown Select-Single.

## Notes

- The `initial_option` configuration overrides the `placeholder` text.
- When nested in a form container, the `name` field is required and must be unique within the card.
- The `confirm` field only triggers the confirmation popup when the user clicks a button containing the submit attribute.

## Nesting Rules

The single select dropdown supports nesting within columns, form containers, folding panels, loop containers, and interactive containers. In the card building tool, it does not currently support nesting within interactive containers.

## JSON Structure

```json
{
  "tag": "select_static",
  "element_id": "custom_id",
  "margin": "0px 0px 0px 0px",
  "type": "text",
  "name": "select_static1",
  "required": false,
  "disabled": false,
  "initial_option": "Option 1",
  "placeholder": {
    "tag": "plain_text",
    "content": "Default placeholder text"
  },
  "width": "default",
  "behaviors": [
    {
      "type": "callback",
      "value": {
        "key": "value"
      }
    }
  ],
  "options": [
    {
      "text": {
        "tag": "plain_text",
        "content": "Option text"
      },
      "icon": {
        "tag": "standard_icon",
        "token": "chat-forbidden_outlined",
        "color": "orange",
        "img_key": "img_v2_38811724"
      },
      "value": "selectDemo1"
    }
  ],
  "confirm": {
    "title": {
      "tag": "plain_text",
      "content": "Pop-up title"
    },
    "text": {
      "tag": "plain_text",
      "content": "Pop-up body text"
    }
  }
}
```

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | Component tag. Fixed value: `select_static`. |
| `element_id` | No | String | — | Unique identifier for the component (JSON 2.0). Must be globally unique within the card. Only letters, numbers, and underscores allowed; must start with a letter; max 20 characters. |
| `margin` | No | String | `0` | Margin of the component (JSON 2.0). Value range is `[-99,99]px`. Accepts a single value (e.g. `"10px"`), double value (e.g. `"4px 0"`), or four values (e.g. `"4px 0 4px 0"`). |
| `type` | No | String | `default` | Border style. `default`: with border. `text`: plain text style without border. |
| `name` | No | String | — | Unique identifier within a form container. Required and must be card-unique when nested in a form container. |
| `required` | No | Boolean | `false` | Whether the component is required. Only effective when nested in a form container; other situations will error or have no effect. |
| `disabled` | No | Boolean | `false` | Whether to disable the component. Supported by Lark clients version 7.4 and above. |
| `initial_option` | No | String | — | Initial option value. Overrides `placeholder` text. |
| `initial_index` | No | Int | — | Index of the initial option. `0` means no initial option displayed; `1` means the first option. Overrides `placeholder` text. |
| `placeholder` | No | Object | — | Placeholder text shown when no option is selected. |
| `placeholder.tag` | Yes | String | `plain_text` | Placeholder tag. Fixed value: `plain_text`. |
| `placeholder.content` | No | String | — | Content of the placeholder text. |
| `width` | No | String | `default` | Component width. `default`: default width. `fill`: fills parent container width. `[100,∞)px`: custom fixed width (min 100px; clips to parent width if exceeded). |
| `behaviors` | No | Array | — | Interaction behavior configuration. See Configuring card interactions. |
| `options` | No | Array | — | Array of option objects. |
| `options[].text` | Yes | Object | — | Option name. |
| `options[].text.tag` | Yes | String | `plain_text` | Option name tag. Fixed value: `plain_text`. |
| `options[].text.content` | Yes | String | — | Option name text. |
| `options[].icon` | No | Object | — | Prefix icon for the option. Supports custom icons or icons from the icon library. |
| `options[].icon.tag` | No | String | — | Icon type. `standard_icon`: icon library. `custom_icon`: custom image. |
| `options[].icon.token` | No | String | — | Icon library token. Effective when `tag` is `standard_icon`. |
| `options[].icon.color` | No | String | — | Icon color for linear/solid icons. Effective when `tag` is `standard_icon`. |
| `options[].icon.img_key` | No | String | — | Custom icon image key. Effective when `tag` is `custom_icon`. Obtain via the Upload Image API. |
| `options[].value` | Yes | String | — | Option callback value. Must be unique within the same select component. |
| `confirm` | No | Object | — | Secondary confirmation popup configuration. Only triggers when a button with a submit attribute is clicked. |
| `confirm.title` | Yes | Object | — | Popup title. |
| `confirm.title.tag` | Yes | String | `plain_text` | Popup title tag. Fixed value: `plain_text`. |
| `confirm.title.content` | Yes | String | — | Popup title content. |
| `confirm.text` | Yes | Object | — | Popup body text. |
| `confirm.text.tag` | Yes | String | `plain_text` | Popup body tag. Fixed value: `plain_text`. |
| `confirm.text.content` | Yes | String | — | Popup body content. |

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
        "tag": "select_static",
        "initial_option": "",
        "placeholder": {
          "tag": "plain_text",
          "content": "请选择"
        },
        "options": [
          {
            "text": {
              "tag": "plain_text",
              "content": "选项1"
            },
            "value": "1",
            "icon": {
              "tag": "standard_icon",
              "token": "signature_outlined"
            }
          },
          {
            "text": {
              "tag": "plain_text",
              "content": "选项2"
            },
            "value": "2",
            "icon": {
              "tag": "standard_icon",
              "token": "signature_outlined"
            }
          },
          {
            "text": {
              "tag": "plain_text",
              "content": "选项3"
            },
            "value": "3",
            "icon": {
              "tag": "standard_icon",
              "token": "signature_outlined"
            }
          }
        ],
        "type": "default",
        "width": "default",
        "disabled": false,
        "behaviors": [
          {
            "type": "callback",
            "value": {
              "select_static": "click"
            }
          }
        ],
        "margin": "0px 0px 0px 0px"
      }
    ]
  }
}
```

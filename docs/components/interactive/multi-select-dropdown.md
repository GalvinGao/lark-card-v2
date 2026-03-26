# Multi-Select Dropdown Menu

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/multi-select-dropdown-menu>

The multi-select dropdown menu component supports customizing the option text, icon, and return parameters. It is an interactive component and must be embedded within a form container.

This document describes the JSON 2.0 structure. For the historical JSON 1.0 structure, refer to Multi-Select Dropdown Menu.

## Notes

- The multi-select component must be placed inside a form container. Content is submitted via the form container's "Submit" button.
- When nested in a form container, the `name` field is required and must be unique within the card.
- Within the same select component, the `value` of each option must be unique.

## Nesting Rules

The multi-select dropdown component is only supported for use within a form container. Content is submitted through the form container's "Submit" button. For details, refer to the Form Container documentation.

## JSON Structure

```json
{
  "tag": "multi_select_static",
  "element_id": "custom_id",
  "margin": "0px 0px 0px 0px",
  "type": "default",
  "name": "multi_select_departments",
  "required": true,
  "disabled": false,
  "placeholder": {
    "tag": "plain_text",
    "content": "Default placeholder text"
  },
  "width": "default",
  "selected_values": [],
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
| `tag` | Yes | String | — | Component tag. Fixed value: `multi_select_static`. |
| `element_id` | No | String | — | Unique identifier for the component (JSON 2.0). Must be globally unique within the card. Only letters, numbers, and underscores allowed; must start with a letter; max 20 characters. |
| `margin` | No | String | `0` | Margin of the component (JSON 2.0). Value range is `[-99,99]px`. Accepts a single value (e.g. `"10px"`), double value (e.g. `"4px 0"`), or four values (e.g. `"4px 0 4px 0"`). |
| `type` | No | String | `default` | Border style. `default`: with border. `text`: plain text style without border. |
| `name` | Yes | String | — | Unique identifier within the form container. Required and must be card-unique when nested in a form container. |
| `placeholder` | No | Object | — | Placeholder text shown when no option is selected. |
| `placeholder.tag` | Yes | String | `plain_text` | Placeholder tag. Fixed value: `plain_text`. |
| `placeholder.content` | No | String | `Please select` | Placeholder text content. Supports up to 100 characters. |
| `width` | No | String | `default` | Component width. `default`: fixed 282px with border (`"type":"default"`), or content-adaptive without border (`"type":"text"`). `fill`: fills parent container width. `[100,∞)px`: custom fixed width (min 100px; clips to parent width if exceeded). |
| `required` | No | Boolean | `true` | Whether the options are required. Only effective in a form container. `true`: prompts "required items not filled" on submit. `false`: submits without selection. |
| `disabled` | No | Boolean | `false` | Whether to disable the component. `true`: displays placeholder/initial values, users cannot interact. `false`: component remains available. |
| `selected_values` | No | Array of String | — | Default selected options. Array values must correspond to `options[].value`. |
| `options` | No | Array | — | Option configuration. Options display in array order. |
| `options[].text` | Yes | Object | — | Option name. Displays blank if empty. |
| `options[].text.tag` | Yes | String | `plain_text` | Option name tag. Fixed value: `plain_text`. |
| `options[].text.content` | Yes | String | — | Option name text. |
| `options[].icon` | No | Object | — | Prefix icon for the option. Supports custom icons or icons from the icon library. |
| `options[].icon.tag` | No | String | — | Icon type. `standard_icon`: icon library. `custom_icon`: custom image. |
| `options[].icon.token` | No | String | — | Icon library token. Effective when `tag` is `standard_icon`. |
| `options[].icon.color` | No | String | — | Icon color for linear/solid icons. Effective when `tag` is `standard_icon`. |
| `options[].icon.img_key` | No | String | — | Custom icon image key. Effective when `tag` is `custom_icon`. Obtain via the Upload Image API. |
| `options[].value` | Yes | String | — | Option callback value. Must be unique within the same select component. |
| `behaviors` | No | Array | — | Interaction behavior configuration. See Configuring card interactions. |
| `confirm` | No | Object | — | Secondary confirmation popup configuration. Only triggers when a button with a submit attribute is clicked. |
| `confirm.title` | Yes | Object | — | Popup title. |
| `confirm.title.tag` | Yes | String | `plain_text` | Popup title tag. Fixed value: `plain_text`. |
| `confirm.title.content` | Yes | String | — | Popup title content. |
| `confirm.text` | Yes | Object | — | Popup body text. |
| `confirm.text.tag` | Yes | String | `plain_text` | Popup body tag. Fixed value: `plain_text`. |
| `confirm.text.content` | Yes | String | — | Popup body content. |

## Callback Structure

When the user clicks the submit button of the form container, the request address configured in the developer backend receives callback data. For more details, see the Form Container and Card Callback Communication documentation.

```json
{
  "schema": "2.0",
  "header": {
    "event_id": "f7984f25108f8137722bb63cee927e66",
    "token": "066zT6pS4QCbgj5Do145GfDbbagCHGgF",
    "create_time": "1603977298000000",
    "event_type": "card.action.trigger",
    "tenant_key": "xxxxxxx",
    "app_id": "cli_xxxxxxxx"
  },
  "event": {
    "operator": {
      "tenant_key": "xxxxxxx",
      "user_id": "xxxxxxx",
      "open_id": "ou_xxx"
    },
    "token": "c-xxxx",
    "action": {
      "value": {
        "key": "value"
      },
      "tag": "button",
      "name": "form_submit",
      "form_value": {
        "multi_select_departments": [
          "selectDemo1",
          "selectDemo2"
        ]
      }
    },
    "host": "im_message",
    "context": {
      "open_message_id": "om_xxx",
      "open_chat_id": "oc_xxx"
    }
  }
}
```

## Example

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "form",
        "elements": [
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
                    "tag": "multi_select_static",
                    "type": "default",
                    "name": "multi_select_departments",
                    "placeholder": {
                      "tag": "plain_text",
                      "content": "默认提示文本"
                    },
                    "width": "fill",
                    "required": true,
                    "disabled": false,
                    "selected_values": [],
                    "behaviors": [
                      {
                        "type": "callback",
                        "value": {
                          "select_static": "click"
                        }
                      }
                    ],
                    "options": [
                      {
                        "text": {
                          "tag": "plain_text",
                          "content": "选项1"
                        },
                        "icon": {
                          "tag": "standard_icon",
                          "token": "chat-forbidden_outlined",
                          "color": "orange"
                        },
                        "value": "selectDemo1"
                      },
                      {
                        "text": {
                          "tag": "plain_text",
                          "content": "选项2"
                        },
                        "icon": {
                          "tag": "standard_icon",
                          "token": "chat-forbidden_outlined",
                          "color": "orange"
                        },
                        "value": "selectDemo2"
                      }
                    ]
                  }
                ],
                "width": "weighted",
                "weight": 1
              }
            ]
          },
          {
            "tag": "column_set",
            "flex_mode": "none",
            "background_style": "default",
            "horizontal_spacing": "default",
            "columns": [
              {
                "tag": "column",
                "width": "auto",
                "vertical_align": "top",
                "elements": [
                  {
                    "tag": "button",
                    "text": {
                      "tag": "plain_text",
                      "content": "提交"
                    },
                    "type": "primary",
                    "action_type": "form_submit",
                    "name": "Button_lrztw8x3"
                  }
                ]
              },
              {
                "tag": "column",
                "width": "auto",
                "vertical_align": "top",
                "elements": [
                  {
                    "tag": "button",
                    "text": {
                      "tag": "plain_text",
                      "content": "取消"
                    },
                    "type": "default",
                    "action_type": "form_reset",
                    "name": "Button_lrztw8x4"
                  }
                ]
              }
            ]
          }
        ],
        "name": "Form_lrztw8x2"
      }
    ]
  }
}
```

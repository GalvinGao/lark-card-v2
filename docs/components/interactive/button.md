# Button

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/button>

The button component is an interactive component that supports multiple styles and sizes, and supports adding icons as prefix icons.

## Notes

- Card JSON 2.0 no longer supports attributes related to interactive modules (`"tag": "action"`). You can place the button directly in `elements` and configure appropriate component spacing (`vertical_spacing` and `horizontal_spacing`).

## Nesting Rules

The button component supports nesting within columns, form containers, folding panels, loop containers, and interactive containers.

## JSON Structure

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "button",
        "element_id": "custom_id",
        "margin": "0px 0px 0px 0px",
        "type": "primary",
        "size": "small",
        "width": "default",
        "text": {
          "tag": "plain_text",
          "content": "Confirm"
        },
        "icon": {
          "tag": "standard_icon",
          "token": "chat-forbidden_outlined",
          "color": "orange",
          "img_key": "img_v2_38811724"
        },
        "hover_tips": {},
        "disabled": false,
        "disabled_tips": {},
        "confirm": {
          "title": {
            "tag": "plain_text",
            "content": "title"
          },
          "text": {
            "tag": "plain_text",
            "content": "content"
          }
        },
        "behaviors": [
          {
            "type": "open_url",
            "default_url": "https://www.baidu.com",
            "android_url": "https://developer.android.com/",
            "ios_url": "lark://msgcard/unsupported_action",
            "pc_url": "https://www.windows.com"
          },
          {
            "type": "callback",
            "value": {
              "key": "value"
            }
          }
        ],
        "url": "https://open.larksuite.com",
        "multi_url": {
          "android_url": "https://open.larksuite.com",
          "ios_url": "https://open.larksuite.com",
          "pc_url": "https://open.larksuite.com"
        },
        "value": {
          "key_1": "value_1"
        }
      }
    ]
  }
}
```

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | The tag of the component. Fixed value: `button`. |
| `element_id` | No | String | Empty | Unique identifier for the operation component (JSON 2.0). Used to specify the component when calling related interfaces. Must be globally unique within the card. Only letters, numbers, and underscores are allowed; must start with a letter and not exceed 20 characters. |
| `margin` | No | String | `0` | Margin of the component (JSON 2.0). Value range: `[-99,99]px`. Supports single value (e.g. `"10px"`), double value (e.g. `"4px 0"`), or four values (e.g. `"4px 0 4px 0"`) separated by spaces. |
| `type` | No | String | `default` | The type of the button. Options: `default` (black font, border), `primary` (blue font, border), `danger` (red font, border), `text` (black font, no border), `primary_text` (blue font, no border), `danger_text` (red font, no border), `primary_filled` (blue background, white font), `danger_filled` (red background, white font), `laser` (laser button). |
| `size` | No | String | `medium` | The size of the button. Options: `tiny` (24px PC / 28px mobile), `small` (28px PC / 28px mobile), `medium` (32px PC / 36px mobile), `large` (40px PC / 48px mobile). |
| `width` | No | String | `default` | The width of the button. Options: `default` (default width), `fill` (maximum card width), or `[100,...)px` (custom width, e.g. `120px`; capped at the card's maximum width). |
| `text` | No | Struct | — | The text on the button. Supports plain text and custom text. |
| `icon` | No | Struct | — | Prefix icon. Supports standard icons and custom icons. |
| `hover_tips` | No | Struct | — | Tooltip text when the user hovers over the button on PC. Default is empty. |
| `disabled` | No | Boolean | `false` | Whether to disable the button. |
| `disabled_tips` | No | Struct | — | Tooltip text when the button is disabled and the user hovers over it on PC. When effective, `hover_tips` is ignored. |
| `confirm` | No | Struct | — | Secondary confirmation dialog configuration. **Note:** The `title` field is required when configuring `confirm`; otherwise, older Lark clients may have unresponsive button clicks. |
| `behaviors` | No | Array | — | Interaction behaviors. Supports `open_url` and `callback` types. |

### Form Container Attributes

Buttons embedded in form containers have these additional attributes:

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `name` | Yes | String | Empty | Unique identifier for the component within the form container. Used to identify which component the submitted data belongs to. Must be globally unique within the card. |
| `form_action_type` | Yes | String | Empty | Interaction type for the button in a form container. `submit`: triggers form submission with all filled values. `reset`: resets all form components to initial values. |

## Callback Structure

After configuring interaction for the button, callback data is sent to the request address configured in the developer console when users interact with the button.

- New card action trigger (`card.action.trigger`): refer to Card Callback Interaction.
- Old card action trigger (`card.action.trigger_v1`): refer to Message Card Callback Interaction (Old).

## Example

```json
{
  "schema": "2.0",
  "header": {
    "template": "blue",
    "title": {
      "content": "Buttons",
      "tag": "plain_text"
    }
  },
  "body": {
    "elements": [
      {
        "tag": "column_set",
        "flex_mode": "flow",
        "background_style": "default",
        "columns": [
          {
            "tag": "column",
            "width": "auto",
            "weight": 1,
            "vertical_align": "top",
            "elements": [
              {
                "tag": "button",
                "text": {
                  "tag": "plain_text",
                  "content": "镭射按钮"
                },
                "behaviors": [
                  {
                    "type": "open_url",
                    "default_url": "https://open.larksuite.com/document",
                    "android_url": "https://developer.android.com/",
                    "ios_url": "lark://msgcard/unsupported_action",
                    "pc_url": "https://www.windows.com"
                  }
                ],
                "type": "laser",
                "hover_tips": {
                  "tag": "plain_text",
                  "content": "hover提示"
                },
                "value": {
                  "key": "value"
                }
              }
            ]
          },
          {
            "tag": "column",
            "width": "auto",
            "weight": 1,
            "vertical_align": "top",
            "elements": [
              {
                "tag": "button",
                "type": "laser",
                "text": {
                  "tag": "plain_text",
                  "content": "镭射禁用按钮"
                },
                "disabled": true,
                "disabled_tips": {
                  "tag": "plain_text",
                  "content": "禁用 hover 提示"
                },
                "behaviors": [
                  {
                    "type": "open_url",
                    "default_url": "https://open.larksuite.com/document",
                    "android_url": "https://developer.android.com/",
                    "ios_url": "lark://msgcard/unsupported_action",
                    "pc_url": "https://www.windows.com"
                  }
                ]
              }
            ]
          }
        ]
      },
      {
        "tag": "column_set",
        "flex_mode": "flow",
        "background_style": "default",
        "columns": [
          {
            "tag": "column",
            "width": "auto",
            "weight": 1,
            "vertical_align": "top",
            "elements": [
              {
                "tag": "button",
                "text": {
                  "tag": "plain_text",
                  "content": "primary"
                },
                "url": "https://open.larksuite.com/document",
                "type": "primary",
                "hover_tips": {
                  "tag": "plain_text",
                  "content": "我是 primary button"
                },
                "value": {
                  "key": "value"
                }
              }
            ]
          },
          {
            "tag": "column",
            "width": "auto",
            "weight": 1,
            "vertical_align": "top",
            "elements": [
              {
                "tag": "button",
                "type": "default",
                "text": {
                  "tag": "plain_text",
                  "content": "default"
                },
                "hover_tips": {
                  "tag": "plain_text",
                  "content": "我是 default 按钮"
                },
                "behaviors": [
                  {
                    "type": "open_url",
                    "default_url": "https://open.larksuite.com/document",
                    "android_url": "https://developer.android.com/",
                    "ios_url": "lark://msgcard/unsupported_action",
                    "pc_url": "https://www.windows.com"
                  }
                ]
              }
            ]
          },
          {
            "tag": "column",
            "width": "auto",
            "weight": 1,
            "vertical_align": "top",
            "elements": [
              {
                "tag": "button",
                "type": "danger",
                "text": {
                  "tag": "plain_text",
                  "content": "我是 danger 按钮"
                },
                "hover_tips": {
                  "tag": "plain_text",
                  "content": "我是 danger 按钮"
                },
                "behaviors": [
                  {
                    "type": "open_url",
                    "default_url": "https://open.larksuite.com/document",
                    "android_url": "https://developer.android.com/",
                    "ios_url": "lark://msgcard/unsupported_action",
                    "pc_url": "https://www.windows.com"
                  }
                ]
              }
            ]
          },
          {
            "tag": "column",
            "width": "auto",
            "weight": 1,
            "vertical_align": "top",
            "elements": [
              {
                "tag": "button",
                "type": "danger",
                "text": {
                  "tag": "plain_text",
                  "content": "我是 disabled 按钮"
                },
                "disabled": true,
                "disabled_tips": {
                  "tag": "plain_text",
                  "content": "我是 disabled 按钮，我被禁用了"
                },
                "behaviors": [
                  {
                    "type": "open_url",
                    "default_url": "https://open.larksuite.com/document",
                    "android_url": "https://developer.android.com/",
                    "ios_url": "lark://msgcard/unsupported_action",
                    "pc_url": "https://www.windows.com"
                  }
                ]
              }
            ]
          }
        ]
      }
    ]
  }
}
```

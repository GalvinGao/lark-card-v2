# Input

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/input>
> Last updated on 2025-06-27

The input box component facilitates simple text content collection from users, such as reasons, evaluations, or remarks. To use the input box together with a button, embed both within a form container.

## Notes

- This document covers the JSON 2.0 structure. For the legacy JSON 1.0 structure, refer to the Input Box documentation.
- When the input box is embedded in a form container, data is submitted asynchronously: after the user completes all form items and clicks the submit button, all form data (including input box data) is sent back to the developer's server in a single callback.

## Nesting Rules

The input box component supports nesting within columns, form containers, folding panels, loop containers, and interactive containers.

## JSON Structure

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "input",
        "element_id": "custom_id",
        "margin": "0px 0px 0px 0px",
        "name": "input1",
        "required": false,
        "placeholder": {
          "tag": "plain_text",
          "content": "Please enter"
        },
        "default_value": "demo",
        "disabled": false,
        "disabled_tips": {
          "tag": "plain_text",
          "content": "User disabled tips"
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
        "max_length": 5,
        "input_type": "multiline_text",
        "rows": 1,
        "auto_resize": true,
        "max_rows": 5,
        "show_icon": false,
        "label": {
          "tag": "plain_text",
          "content": "Please enter text:"
        },
        "label_position": "left",
        "value": {
          "k": "v"
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
    ]
  }
}
```

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | Tag of the input box. Fixed value: `input`. |
| `element_id` | No | String | — | Unique identifier for the component (JSON 2.0). Used to specify the component in related interface calls. Must be globally unique within the card; only letters, numbers, and underscores allowed, must start with a letter, max 20 characters. |
| `margin` | No | String | `0` | Margin of the component (JSON 2.0). Value range: `[-99,99]px`. Accepts a single value (`"10px"`), double value (`"4px 0"`), or four values (`"4px 0 4px 0"`). |
| `name` | No | String | — | Unique identifier of the input box. Effective when embedded in a form container to identify which input box the submitted text belongs to. **Required** when nested within a form container and must be globally unique within the card. |
| `required` | No | Boolean | `false` | Whether input is required. Only effective inside a form container. `true`: frontend prompts "Required item not filled" and blocks submission. `false`: data is submitted even if empty. |
| `disabled` | No | Boolean | `false` | Whether to disable the input box. Requires Lark v7.4+. `true`: disabled. `false`: enabled. |
| `disabled_tips` | No | text object | — | Tooltip displayed when the component is disabled and the user hovers over it. Uses `{"tag": "plain_text", "content": "..."}`. |
| `placeholder` | No | text object | — | Placeholder text inside the input box. |
| `placeholder.tag` | No | String | `plain_text` | Tag for the placeholder text. Fixed value: `plain_text`. |
| `placeholder.content` | No | String | `Please enter` | Content of the placeholder text, up to 100 characters. |
| `default_value` | No | String | — | Pre-filled content for the user. Displays as user-entered text pending submission. |
| `width` | No | String | `default` | Width of the input box. Values: `default` (default width), `fill` (maximum card width), or a custom width `[100,∞)px`. |
| `behaviors` | Yes | Struct | — | Configures the interaction type and specific behaviors. See Configuring Card Interactions. |
| `max_length` | No | Number | `1000` | Maximum text length (range 1--1000). Displays an error if exceeded. |
| `input_type` | No | String | `text` | Input type. `text`: ordinary text. `multiline_text`: multiline text (line breaks returned as `\n`). `password`: entered text displayed as dots. Note: not yet supported in the card builder tool. |
| `show_icon` | No | Boolean | `true` | Whether to display a prefix icon when `input_type` is `password`. Only effective for password inputs. |
| `rows` | No | Number | `5` | Default number of display rows. Only effective when `input_type` is `multiline_text`. |
| `auto_resize` | No | Boolean | `false` | Whether the input box height adapts to the text height. Only effective on PC when `input_type` is `multiline_text`. `true`: height adapts. `false`: height fixed to `rows`. |
| `max_rows` | No | Number | — | Maximum number of display rows. Only effective when `auto_resize` is `true`. Must be an integer >= 1. If empty, there is no limit (up to the front-end rendering maximum). |
| `label` | No | text object | — | Text label describing the input box, prompting the user on what to enter. Mostly used in form containers. |
| `label.tag` | No | String | `plain_text` | Tag for the label text. Fixed value: `plain_text`. |
| `label.content` | No | String | — | Label description content. |
| `label_position` | No | String | `top` | Position of the text label. `top`: above the input box. `left`: to the left. On mobile / narrow screens, the label auto-adjusts to `top`. |
| `value` | No | String or Object | — | Custom data passed back during interaction events. Supports string or object types. |
| `confirm` | No | Struct | — | Secondary confirmation dialog configuration. Triggered only when the user clicks a button with the `submit` attribute. Provides confirm/cancel buttons by default; you only configure the title and content. |
| `confirm.title` | Yes | Struct | — | Title of the confirmation dialog. |
| `confirm.title.tag` | Yes | String | `plain_text` | Tag for the title text. Fixed value: `plain_text`. |
| `confirm.title.content` | Yes | String | — | Content of the confirmation dialog title. |
| `confirm.text` | Yes | Struct | — | Body text of the confirmation dialog. |
| `confirm.text.tag` | Yes | String | `plain_text` | Tag for the body text. Fixed value: `plain_text`. |
| `confirm.text.content` | Yes | String | — | Content of the confirmation dialog body. |

## Callback Structure

To use the input box, enable interactivity for your card (see Configuring Card Interactions). When a user clicks the submit button, the interaction event is passed back as shown below. If the input box is embedded in a form container, refer to the Form Container callback structure.

```json
{
  "schema": "2.0",
  "header": {
    "event_id": "f7984f25108f8137722bb63c*****",
    "token": "066zT6pS4QCbgj5Do145GfDbbag*****",
    "create_time": "1603977298000000",
    "event_type": "card.action.trigger",
    "tenant_key": "2df73991750*****",
    "app_id": "cli_a5fb0ae6a4******"
  },
  "event": {
    "operator": {
      "tenant_key": "2df73991750*****",
      "user_id": "867*****",
      "open_id": "ou_3c14f3a59eaf2825dbe25359f15*****"
    },
    "token": "c-295ee57216a5dc9de90fefd0aadb4b1d7d******",
    "action": {
      "value": {
        "key": "value"
      },
      "tag": "input",
      "input_value": "Zhang Min",
      "name": "Input_lf4fmxwfrd9"
    },
    "host": "im_message",
    "context": {
      "open_message_id": "om_574d639e4a44e4dd646eaf628e2*****",
      "open_chat_id": "oc_e4d2605ca917e695f54f11aaf56*****"
    }
  }
}
```

## Example

A form container with three input boxes (username, password, and multiline address) plus submit/cancel buttons:

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "form",
        "elements": [
          {
            "tag": "input",
            "element_id": "username",
            "margin": "0px 0px 0px 0px",
            "placeholder": {
              "tag": "plain_text",
              "content": "请输入"
            },
            "default_value": "",
            "width": "default",
            "label": {
              "tag": "plain_text",
              "content": "用户名："
            },
            "name": "Input_31q6mtuvdx9"
          },
          {
            "tag": "input",
            "element_id": "password",
            "margin": "0px 0px 0px 0px",
            "input_type": "password",
            "placeholder": {
              "tag": "plain_text",
              "content": "请输入"
            },
            "default_value": "",
            "width": "default",
            "label": {
              "tag": "plain_text",
              "content": "密码："
            },
            "label_position": "top",
            "name": "Input_5hez3q41fck"
          },
          {
            "tag": "input",
            "element_id": "address",
            "margin": "0px 0px 0px 0px",
            "input_type": "multiline_text",
            "rows": 4,
            "auto_resize": true,
            "placeholder": {
              "tag": "plain_text",
              "content": "请输入"
            },
            "default_value": "",
            "width": "default",
            "label": {
              "tag": "plain_text",
              "content": "收货地址："
            },
            "name": "Input_u2k3lbrokvd"
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
                    "name": "Button_lrocopxs"
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
                    "name": "Button_lrocopxt"
                  }
                ]
              }
            ],
            "margin": "0px"
          }
        ],
        "name": "Form_lrocopxr"
      }
    ]
  }
}
```

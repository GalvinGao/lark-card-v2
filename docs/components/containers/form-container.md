# Form Container

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/containers/form-container>

When using cards to collect content, there may be scenarios where users need to submit multiple form items. The form container allows users to locally enter a batch of form items at the frontend, and by clicking the Submit button once, all locally cached form content is asynchronously called back to the developer's server, achieving the effect of asynchronously submitting multiple form item data.

This document introduces the JSON 2.0 structure of the form container. To view the historical JSON 1.0 structure, refer to Form Container.

## Notes

Container type components support up to five levels of nesting. It is recommended to avoid multiple layers of nesting within the form container. Excessive nesting compresses the display space and affects the presentation of the card. If you wish to handle more complex form content, it is recommended to implement form capabilities via card links to H5 or mini-programs.

## Nesting Rules

In the Card JSON 2.0 Structure:

- Form containers do not support embedding table (table) and form container components.
- Form container components cannot be embedded within other components and can only be placed under the root node of the card.

## JSON Structure

The following is the JSON 2.0 structure of the form container card. In this structure, a form container embeds an input field component and a submit button bound to a submit event:

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "form",
        "element_id": "custom_id",
        "direction": "horizontal",
        "horizontal_spacing": "8px",
        "vertical_spacing": "8px",
        "horizontal_align": "left",
        "vertical_align": "top",
        "padding": "4px 0px 4px 0px",
        "margin": "0px 0px 0px 0px",
        "name": "form_1",
        "elements": [
          {
            "tag": "input",
            "name": "reason",
            "required": true
          },
          {
            "tag": "column_set",
            "columns": [
              {
                "tag": "column",
                "width": "auto",
                "elements": [
                  {
                    "tag": "button",
                    "type": "primary",
                    "text": {
                      "tag": "plain_text",
                      "content": "Submit"
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
                    "form_action_type": "submit",
                    "name": "Button_m8pn9lbf"
                  }
                ]
              },
              {
                "tag": "column",
                "width": "auto",
                "elements": [
                  {
                    "tag": "button",
                    "type": "default",
                    "text": {
                      "tag": "plain_text",
                      "content": "Cancel"
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
                    "form_action_type": "reset",
                    "name": "Button_m8pn9lbg"
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

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | Tag of the form container. Fixed value is `form`. |
| `element_id` | No | String | Empty | Unique identifier of the operation component. Newly added attribute in JSON 2.0. Used to specify the component when calling related component interfaces. The value of this field must be globally unique within the same card. Only letters, numbers, and underscores are allowed, must start with a letter, and must not exceed 20 characters. |
| `direction` | No | String | `vertical` | Arrangement direction of components within the container. Optional values: `vertical` (vertical arrangement), `horizontal` (horizontal arrangement). |
| `margin` | No | String | `0px` | Outer margin of the container. The value range is [-99,99]px. Optional values: single value such as `"10px"` (all four margins); double value such as `"4px 0"` (top/bottom and left/right); multiple values such as `"4px 0 4px 0"` (top, right, bottom, left). Separated by spaces. |
| `padding` | No | String | `0px` | Inner padding of the container. The value range is [-99,99]px. Optional values: single value such as `"10px"` (all four sides); double value such as `"4px 0"` (top/bottom and left/right); multiple values such as `"4px 0 4px 0"` (top, right, bottom, left). Separated by spaces. |
| `horizontal_spacing` | No | String | `8px` | Horizontal spacing between components within the container. Optional values: `small` (4px), `medium` (8px), `large` (12px), `extra_large` (16px), or a specific value such as `"20px"`. Value range is [0,99]px. |
| `horizontal_align` | No | String | `left` | Horizontal alignment of components within the container. Optional values: `left`, `center`, `right`. |
| `vertical_align` | No | String | `top` | Vertical alignment of components within the container. Optional values: `top`, `center`, `bottom`. |
| `vertical_spacing` | No | String | `12px` | Vertical spacing between components within the container. Optional values: `small` (4px), `medium` (8px), `large` (12px), `extra_large` (16px), or a specific value such as `"20px"`. Value range is [0,99]px. |
| `name` | Yes | String | — | Unique identifier of the form container. Used to identify which form container the user-submitted data belongs to. The value of this field must be globally unique within the same card. |
| `elements` | Yes | Array\<element\> | `[]` | Child nodes of the form container. Other container components and display/interactive components can be nested, but table, chart, and form container components cannot be nested. |
| `elements[].tag` | Yes | String | — | The tag of the component embedded in the form container. Supports all components except tables and form containers. Note: The form container must contain a button component for submitting the form. |
| `elements[].name` | Yes | String | — | Unique identifier of the component within the form container. Used to identify which component the user-submitted data belongs to. Note: This field is required for all interactive components in the form container and must be globally unique within the card; otherwise, the data submission will fail. |
| `elements[].required` | No | Boolean | `false` | Whether the content of the component is required. When `true`, if the user clicks "Submit" without filling this component, the frontend will prompt "There are required fields not filled out" and will not initiate a callback request. When `false`, unfilled components will still be included in the submission. |
| `elements[].form_action_type` | Yes | String | — | The interaction type of the button embedded in the form container. `submit`: binds the button to the submit event, triggering the form container's submit and asynchronously submitting all filled form items. `reset`: binds the button to the cancel/reset event, resetting all form component input values to their initial values. |
| `elements[].action_type` (deprecated) | No | String | — | Legacy interaction type for the button. Values: `link` (link jumps only), `request` (callback interactions only), `multi` (both link jumps and callbacks), `form_submit` (submit event), `form_reset` (reset event). Note: The form container must include a button with `action_type` set to `form_submit` for form submission. |

## Callback Structure

When a user clicks the submit button of the form container, your configured backend URL will receive callback data. If you added the new version card interaction callback (`card.action.trigger`), the structure is as follows. For detailed parameter descriptions, refer to Card Interaction Callback.

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
            "tag": "button",
            "timezone": "Asia/Shanghai",
            "form_value": {
                "DatePicker_bpqdq5puvn4": "2024-04-01 +0800",
                "DateTimePicker_ihz2d7a74i": "2024-04-29 07:07 +0800",
                "Input_lf4fmxwfrd9": "1234",
                "PersonSelect_2ejys7ype7m": "ou_3c14f3a59eaf2825dbe25359f1595b00",
                "Select_a2d5b7l3zd": "1",
                "TimePicker_7ecsf6xkqsq": "00:00 +0800"
            },
            "name": "Button_lvkepfu3"
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

The following JSON 2.0 example demonstrates a form container with person selection, priority dropdown, comment input, and submit/cancel buttons:

```json
{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "markdown",
                "content": "**项目名称**：业务做大做强",
                "text_align": "left",
                "text_size": "normal",
                "icon": {
                    "tag": "standard_icon",
                    "token": "add-app_outlined",
                    "color": "grey"
                }
            },
            {
                "tag": "form",
                "elements": [
                    {
                        "tag": "column_set",
                        "horizontal_spacing": "8px",
                        "horizontal_align": "left",
                        "columns": [
                            {
                                "tag": "column",
                                "width": "weighted",
                                "elements": [
                                    {
                                        "tag": "markdown",
                                        "content": "**经办人**<font color='red'>*</font>",
                                        "text_align": "left",
                                        "text_size": "normal",
                                        "icon": {
                                            "tag": "standard_icon",
                                            "token": "member_outlined",
                                            "color": "light_grey"
                                        }
                                    }
                                ],
                                "padding": "0px 0px 0px 0px",
                                "vertical_spacing": "8px",
                                "weight": 1
                            },
                            {
                                "tag": "column",
                                "width": "weighted",
                                "elements": [
                                    {
                                        "tag": "select_person",
                                        "placeholder": {
                                            "tag": "plain_text",
                                            "content": "请选择"
                                        },
                                        "options": [],
                                        "width": "fill",
                                        "type": "default",
                                        "required": true,
                                        "name": "PersonSelect_rg0ml5mh"
                                    }
                                ],
                                "padding": "0px 0px 0px 0px",
                                "vertical_spacing": "8px",
                                "weight": 4
                            }
                        ],
                        "margin": "16px 0px 16px 0px"
                    },
                    {
                        "tag": "column_set",
                        "horizontal_spacing": "8px",
                        "horizontal_align": "left",
                        "columns": [
                            {
                                "tag": "column",
                                "width": "weighted",
                                "elements": [
                                    {
                                        "tag": "markdown",
                                        "content": "**优先级**<font color='red'>*</font>",
                                        "text_align": "left",
                                        "text_size": "normal",
                                        "icon": {
                                            "tag": "standard_icon",
                                            "token": "msgcard-rectangle_outlined",
                                            "color": "grey"
                                        }
                                    }
                                ],
                                "padding": "0px 0px 0px 0px",
                                "vertical_spacing": "8px",
                                "weight": 1
                            },
                            {
                                "tag": "column",
                                "width": "weighted",
                                "elements": [
                                    {
                                        "tag": "select_static",
                                        "placeholder": {
                                            "tag": "plain_text",
                                            "content": "请选择"
                                        },
                                        "options": [
                                            {
                                                "text": {
                                                    "tag": "plain_text",
                                                    "content": "P0"
                                                },
                                                "value": "1",
                                                "icon": {
                                                    "tag": "standard_icon",
                                                    "token": "sheet-iconsets-increase_filled"
                                                }
                                            },
                                            {
                                                "text": {
                                                    "tag": "plain_text",
                                                    "content": "P1"
                                                },
                                                "value": "P1",
                                                "icon": {
                                                    "tag": "standard_icon",
                                                    "token": "sheet-iconsets-stable_filled"
                                                }
                                            },
                                            {
                                                "text": {
                                                    "tag": "plain_text",
                                                    "content": "P2"
                                                },
                                                "value": "3",
                                                "icon": {
                                                    "tag": "standard_icon",
                                                    "token": "expand-down_filled"
                                                }
                                            }
                                        ],
                                        "type": "default",
                                        "width": "fill",
                                        "required": true,
                                        "name": "Select_01taxkgaqc6c"
                                    }
                                ],
                                "padding": "0px 0px 0px 0px",
                                "vertical_spacing": "8px",
                                "weight": 4
                            }
                        ],
                        "margin": "16px 0px 16px 0px"
                    },
                    {
                        "tag": "column_set",
                        "horizontal_spacing": "8px",
                        "horizontal_align": "left",
                        "columns": [
                            {
                                "tag": "column",
                                "width": "weighted",
                                "elements": [
                                    {
                                        "tag": "markdown",
                                        "content": "**项目评论**",
                                        "text_align": "left",
                                        "text_size": "normal",
                                        "icon": {
                                            "tag": "standard_icon",
                                            "token": "chat_outlined",
                                            "color": "grey"
                                        }
                                    }
                                ],
                                "padding": "0px 0px 0px 0px",
                                "vertical_spacing": "8px",
                                "weight": 1
                            },
                            {
                                "tag": "column",
                                "width": "weighted",
                                "elements": [
                                    {
                                        "tag": "input",
                                        "placeholder": {
                                            "tag": "plain_text",
                                            "content": "请输入"
                                        },
                                        "default_value": "",
                                        "width": "fill",
                                        "name": "Input_0bqcy75cxklr",
                                        "fallback": {
                                            "tag": "fallback_text",
                                            "text": {
                                                "tag": "plain_text",
                                                "content": "仅支持在 V6.8 及以上版本使用"
                                            }
                                        }
                                    }
                                ],
                                "padding": "0px 0px 0px 0px",
                                "vertical_spacing": "8px",
                                "weight": 4
                            }
                        ],
                        "margin": "16px 0px 16px 0px"
                    },
                    {
                        "tag": "column_set",
                        "flex_mode": "bisect",
                        "horizontal_spacing": "8px",
                        "horizontal_align": "right",
                        "columns": [
                            {
                                "tag": "column",
                                "width": "auto",
                                "elements": [
                                    {
                                        "tag": "button",
                                        "text": {
                                            "tag": "plain_text",
                                            "content": "提交"
                                        },
                                        "type": "primary_filled",
                                        "width": "default",
                                        "icon": {
                                            "tag": "standard_icon",
                                            "token": "thumbsup_outlined"
                                        },
                                        "form_action_type": "submit",
                                        "name": "Button_lq544v6r"
                                    }
                                ],
                                "padding": "0px 0px 0px 0px",
                                "vertical_spacing": "8px"
                            },
                            {
                                "tag": "column",
                                "width": "auto",
                                "elements": [
                                    {
                                        "tag": "button",
                                        "text": {
                                            "tag": "plain_text",
                                            "content": "取消"
                                        },
                                        "type": "default",
                                        "width": "default",
                                        "form_action_type": "reset",
                                        "name": "Button_lq544v6s"
                                    }
                                ],
                                "padding": "0px 0px 0px 0px",
                                "vertical_spacing": "8px"
                            }
                        ],
                        "margin": "0px 0px 0px 0px"
                    }
                ],
                "name": "Form_lq544v6q",
                "fallback": {
                    "tag": "fallback_text",
                    "text": {
                        "tag": "plain_text",
                        "content": "仅支持在 V6.6 及以上版本使用"
                    }
                }
            }
        ]
    }
}
```

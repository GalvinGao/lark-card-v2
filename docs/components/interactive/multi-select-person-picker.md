# Multi-Select Person Picker

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/multi-select-user-picker>

The multi-select person picker component allows adding specified users as multi-choice options. It is an interactive component that must be embedded within a form container.

## Notes

- The multi-select person picker must be nested inside a form container. Content is submitted through the form container's "Submit" button.
- When nested in a form container, the `name` field is required and must be globally unique within the card.
- To understand form containers and their interactive configurations, refer to the Form Container documentation.

## Nesting Rules

The multi-select person picker is only supported within a form container. Refer to the Form Container documentation for details.

## JSON Structure

```json
{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "form",
                "elements": [
                    {
                        "tag": "multi_select_person",
                        "element_id": "custom_id",
                        "margin": "0px 0px 0px 0px",
                        "type": "text",
                        "name": "multi_select_users",
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
                        "required": true,
                        "disabled": false,
                        "selected_values": [
                            "ou_48d0958ee4b2ab3eaf0b5f6c968xxxxx"
                        ],
                        "options": [
                            {
                                "value": "ou_48d0958ee4b2ab3eaf0b5f6c968xxxxx"
                            },
                            {
                                "value": "ou_f9d24af786a14340721288cda6axxxxx"
                            }
                        ],
                        "confirm": {
                            "title": {
                                "tag": "plain_text",
                                "content": "Popup title"
                            },
                            "text": {
                                "tag": "plain_text",
                                "content": "Popup body text"
                            }
                        }
                    },
                    {
                        "tag": "button",
                        "text": {
                            "tag": "plain_text",
                            "content": "Submit"
                        },
                        "type": "primary",
                        "form_action_type": "submit"
                    }
                ],
                "name": "Form_example"
            }
        ]
    }
}
```

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | Component tag. Fixed value `multi_select_person`. |
| `element_id` | No | String | — | Unique identifier for the component (JSON 2.0). Used to specify the component in related interface calls. Must be globally unique within the card. Only letters, numbers, and underscores allowed; must start with a letter; max 20 characters. |
| `margin` | No | String | `0` | Component margin (JSON 2.0). Value range is `[-99,99]px`. Accepts a single value (e.g. `"10px"`), double value (e.g. `"4px 0"`), or four values (e.g. `"4px 0 4px 0"`) for top, right, bottom, left. |
| `type` | No | String | `default` | Border style. `default`: with border. `text`: plain text style without border. |
| `name` | Yes | String | — | Unique identifier for the component within the form container. Used to identify which component the submitted data belongs to. Must be globally unique within the card when nested in a form container. |
| `placeholder` | No | Object | — | Placeholder text displayed inside the picker when no option is selected. |
| `placeholder.tag` | Yes | String | `plain_text` | Fixed value `plain_text`. |
| `placeholder.content` | No | String | `Please select` | Placeholder text content. Max 100 characters. |
| `width` | No | String | `default` | Component width. `default`: fixed at 282px with border (`"type":"default"`), or content-adaptive without border (`"type":"text"`). `fill`: fills parent container width. `[100,...)px`: custom fixed width (min 100px; caps at parent container width). |
| `required` | No | Boolean | `true` | Whether selection is required. Only takes effect inside a form container. When `true`, submitting without a selection shows a "required items not filled in" prompt and prevents submission. |
| `disabled` | No | Boolean | `false` | Whether to disable the component. `true`: disabled, shows placeholder text or initial values. `false`: enabled. |
| `selected_values` | No | Array of strings | — | Default selected options. Array values must correspond to `options[].value`. |
| `options` | No | Array of objects | — | Option configuration. Options are displayed in array order. |
| `options[].value` | Yes | String | — | The candidate user's `open_id`. Within the same component, each option's value must be unique. |
| `confirm` | No | Object | — | Secondary confirmation popup configuration. Only triggers when the user clicks a button with the submit attribute. |
| `confirm.title` | Yes | Object | — | Popup title. |
| `confirm.title.tag` | Yes | String | `plain_text` | Fixed value `plain_text`. |
| `confirm.title.content` | Yes | String | — | Title text content. |
| `confirm.text` | Yes | Object | — | Popup body text. |
| `confirm.text.tag` | Yes | String | `plain_text` | Fixed value `plain_text`. |
| `confirm.text.content` | Yes | String | — | Body text content. |

## Callback Structure

When the user clicks the submit button of the form container, the configured request address receives callback data with the following structure. For more information, see the Form Container and Card Callback Communication documentation.

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
                "multi_select_users": [
                    "ou_48d0958ee4b2ab3eaf0b5f6c968xxxxx",
                    "ou_f9d24af786a14340721288cda6axxxxx"
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

Replace the `open_id` values with actual user IDs.

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
                                        "tag": "multi_select_person",
                                        "type": "default",
                                        "name": "multi_select_users",
                                        "placeholder": {
                                            "tag": "plain_text",
                                            "content": "请选择"
                                        },
                                        "width": "fill",
                                        "required": true,
                                        "disabled": false,
                                        "selected_values": [
                                            "ou_48d0958ee4b2ab3eaf0b5f6c968xxxxx"
                                        ],
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
                                                "value": "ou_48d0958ee4b2ab3eaf0b5f6c968xxxxx"
                                            },
                                            {
                                                "value": "ou_f9d24af786a14340721288cda6axxxxx"
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

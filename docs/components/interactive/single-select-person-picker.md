# Single Select Person Picker

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/single-select-user-picker>

The single select person picker component allows adding specified users as single-choice options. It is an interactive component.

## Notes

- The `initial_option` property is not currently supported in the Card Builder Tool.
- When the `options` array is empty or all `value` entries are invalid, the candidate options default to all members in the conversation where the card is located.

## Nesting Rules

The single select person picker supports nesting within columns, form containers, folding panels, loop containers, and interactive containers. In the Card Builder Tool, nesting within interactive containers is not yet supported.

## JSON Structure

```json
{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "select_person",
                "element_id": "custom_id",
                "margin": "0px 0px 0px 0px",
                "type": "text",
                "required": true,
                "disabled": false,
                "initial_option": "ou_48d0958ee4b2ab3eaf0b5f6c968xxxxx",
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
            }
        ]
    }
}
```

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | Component tag. Fixed value `select_person`. |
| `element_id` | No | String | — | Unique identifier for the component (JSON 2.0). Used to specify the component in related interface calls. Must be globally unique within the card. Only letters, numbers, and underscores allowed; must start with a letter; max 20 characters. |
| `margin` | No | String | `0` | Component margin (JSON 2.0). Value range is `[-99,99]px`. Accepts a single value (e.g. `"10px"`), double value (e.g. `"4px 0"`), or four values (e.g. `"4px 0 4px 0"`) for top, right, bottom, left. |
| `type` | No | String | `default` | Border style. `default`: with border. `text`: plain text style without border. |
| `required` | No | Boolean | `false` | Whether selection is required. Only takes effect inside a form container. When `true`, submitting the form without a selection shows a "required items not filled in" prompt and prevents submission. |
| `disabled` | No | Boolean | `false` | Whether to disable the component. `true`: disabled. `false`: enabled. |
| `initial_option` | No | String | — | The `open_id` of the initially selected person. Must match one of the `open_id` values in `options`; otherwise it has no effect. |
| `placeholder` | No | Object | — | Placeholder text displayed inside the picker when no option is selected. |
| `placeholder.tag` | Yes | String | `plain_text` | Fixed value `plain_text`. |
| `placeholder.content` | No | String | — | Placeholder text content. Max 100 characters. |
| `width` | No | String | `default` | Component width. `default`: default width. `fill`: maximum card width. `[100,...)px`: custom width (displays at max width if it exceeds the card). |
| `options` | No | Array of objects | — | Option configuration. Options are displayed in array order. |
| `options[].value` | No | String | — | The candidate user's `open_id`. See [How to get different user IDs](https://open.larksuite.com/document). |
| `confirm` | No | Object | — | Secondary confirmation popup configuration. Only triggers when the user clicks a button with the submit attribute. |
| `confirm.title` | Yes | Object | — | Popup title. |
| `confirm.title.tag` | Yes | String | `plain_text` | Fixed value `plain_text`. |
| `confirm.title.content` | Yes | String | — | Title text content. |
| `confirm.text` | Yes | Object | — | Popup body text. |
| `confirm.text.tag` | Yes | String | `plain_text` | Fixed value `plain_text`. |
| `confirm.text.content` | Yes | String | — | Body text content. |

## Callback Structure

After configuring interactions for the component, user interactions will send callback data to your configured request address.

- If you use the new card callback interaction (`card.action.trigger`), refer to the Card Callback Interaction documentation for the callback structure.
- If you use the old card callback interaction (`card.action.trigger_v1`), refer to the Message Card Callback Interaction (Old) documentation.

## Example

Replace the `open_id` values with actual user IDs.

```json
{
    "schema": "2.0",
    "body": {
        "direction": "vertical",
        "padding": "12px 12px 12px 12px",
        "elements": [
            {
                "tag": "select_person",
                "placeholder": {
                    "tag": "plain_text",
                    "content": "请选择"
                },
                "options": [
                    {
                        "value": "ou_449b53ad6aee526f7ed311b216aabcef"
                    },
                    {
                        "value": "ou_449b53ad6aee526f7ed311b216aabcef"
                    }
                ],
                "width": "default",
                "type": "default",
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

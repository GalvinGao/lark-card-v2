# Date Time Picker

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/date-time-picker>
> Last updated on 2025-06-27

The date time picker component is an interactive component used to provide combined date and time selection options. This document covers the JSON 2.0 structure. For the legacy JSON 1.0 structure, refer to the date-time picker documentation.

## Notes

When using the date time picker, remind users to select the timezone corresponding to the current scenario. For example, when booking overseas hotels, use the hotel location's timezone; for scheduling, use the user's current timezone. The open platform returns the user's current timezone as a reference, but this does not mean the user has selected that timezone.

## Nesting Rules

The date time picker supports nesting within column sets, form containers, folding panels, loop containers, and interactive containers. In the builder tool, the date time picker does not yet support nesting within interactive containers.

## JSON Structure

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "picker_datetime",
        "element_id": "custom_id",
        "margin": "0px 0px 0px 0px",
        "name": "picker_datetime1",
        "required": false,
        "disabled": false,
        "width": "default",
        "behaviors": [
          {
            "type": "callback",
            "value": {
              "key": "value"
            }
          }
        ],
        "initial_datetime": "2024-01-01 11:30",
        "placeholder": {
          "tag": "plain_text",
          "content": "Please select"
        },
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
    ]
  }
}
```

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | Tag of the component. Fixed value: `picker_datetime`. |
| `element_id` | No | String | — | Unique identifier for the component (JSON 2.0). Used to specify the component in related interface calls. Must be globally unique within the card; only letters, numbers, and underscores allowed, must start with a letter, max 20 characters. |
| `margin` | No | String | `0` | Margin of the component (JSON 2.0). Value range: `[-99,99]px`. Accepts a single value (`"10px"`), double value (`"4px 0"`), or four values (`"4px 0 4px 0"`). |
| `name` | No | String | — | Unique identifier of the date time picker. Effective when embedded in a form container to identify which component the submitted data belongs to. **Required** when nested within a form container and must be globally unique within the card. |
| `required` | No | Boolean | `false` | Whether the date time is required. Only effective inside a form container. `true`: frontend prompts "Required items are not filled in" and blocks submission. `false`: data is submitted even if empty. |
| `disabled` | No | Boolean | `false` | Whether to disable the date time picker. Requires Lark v7.4+. `true`: disabled. `false`: enabled. |
| `initial_datetime` | No | String | — | Initial date time value. Format: `yyyy-MM-dd HH:mm`. Overrides the placeholder text when set. |
| `placeholder` | No | Object | — | Placeholder text inside the date time picker. When `initial_datetime` is not set, this field is required. When `initial_datetime` is set, this field has no effect. |
| `placeholder.tag` | Yes | String | `plain_text` | Placeholder tag. Fixed value: `plain_text`. |
| `placeholder.content` | No | String | — | Content of the placeholder text, up to 100 characters. |
| `width` | No | String | `default` | Width of the date time picker. Values: `default` (default width), `fill` (maximum card width), or a custom width `[100,∞)px`. |
| `value` | Yes | JSON | — | Return data for interaction. When the user selects a date time, this value is sent to the callback server. Only supports key-value JSON structure with String keys. |
| `confirm` | No | Struct | — | Secondary confirmation dialog configuration. Triggered only when the user clicks a button with the `submit` attribute. Provides confirm/cancel buttons by default; you only configure the title and content. |
| `confirm.title` | Yes | Struct | — | Title of the confirmation dialog. |
| `confirm.title.tag` | Yes | String | `plain_text` | Tag for the title text. Fixed value: `plain_text`. |
| `confirm.title.content` | Yes | String | — | Content of the confirmation dialog title. |
| `confirm.text` | Yes | Struct | — | Body text of the confirmation dialog. |
| `confirm.text.tag` | Yes | String | `plain_text` | Tag for the body text. Fixed value: `plain_text`. |
| `confirm.text.content` | Yes | String | — | Content of the confirmation dialog body. |

## Callback Structure

After configuring interaction for the component, user interactions trigger callback data sent to the configured request address.

- New version callback (`card.action.trigger`): refer to Card Callback Interaction.
- Old version callback (`card.action.trigger_v1`): refer to Message Card Callback Interaction (Old).

## Example

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "picker_datetime",
        "placeholder": {
          "tag": "plain_text",
          "content": "请选择"
        },
        "width": "default"
      },
      {
        "tag": "picker_datetime",
        "placeholder": {
          "tag": "plain_text",
          "content": "请选择"
        },
        "width": "default",
        "initial_datetime": "2024-01-01 08:00"
      }
    ]
  }
}
```

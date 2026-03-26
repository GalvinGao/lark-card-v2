# Divider

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/divider>

You can add a divider component to a card to visually separate content.

## JSON Structure

```json
{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "hr",
                "element_id": "custom_id",
                "margin": "0px 0px 0px 0px"
            }
        ]
    }
}
```

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | Component tag. Fixed value: `hr`. |
| `element_id` | No | String | — | Unique identifier for the component (JSON 2.0). Must be globally unique within the card. Only letters, numbers, and underscores are allowed; must start with a letter and not exceed 20 characters. |
| `margin` | No | String | `"0"` | Margin of the component (JSON 2.0). Value range is [-99,99]px. Accepts a single value (e.g. `"10px"`), double value (e.g. `"4px 0"`), or four values (e.g. `"4px 0 4px 0"`) for top, right, bottom, left. |

## Example

```json
{
    "schema": "2.0",
    "body": {
        "direction": "vertical",
        "padding": "12px 12px 12px 12px",
        "elements": [
            {
                "tag": "div",
                "text": {
                    "tag": "plain_text",
                    "content": "普通文本示例",
                    "text_size": "normal",
                    "text_align": "left",
                    "text_color": "default"
                },
                "margin": "0px 0px 0px 0px"
            },
            {
                "tag": "hr",
                "margin": "0px 0px 0px 0px"
            },
            {
                "tag": "button",
                "text": {
                    "tag": "plain_text",
                    "content": "查看更多"
                },
                "type": "primary",
                "width": "default",
                "size": "medium",
                "margin": "0px 0px 0px 0px"
            }
        ]
    }
}
```

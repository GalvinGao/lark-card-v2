# Person

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/user-profile>

The person component displays a user's name and avatar. You can use this component by passing in the person's `open_id`, `user_id`, or `union_id`.

If you are using a specified application to send a card containing the person component, you must ensure that the application has access to the user IDs. Otherwise, the person component will not display any person information.

## JSON Structure

```json
{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "person",
                "element_id": "custom_id",
                "margin": "0px 0px 0px 0px",
                "size": "extra_small",
                "user_id": "ou_4a136bca010747fc3bd7b6f8f4cabcef",
                "show_avatar": true,
                "show_name": false,
                "style": "normal"
            }
        ]
    }
}
```

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | `"person"` | Component tag. Fixed value: `person`. |
| `element_id` | No | String | — | Unique identifier for the component (JSON 2.0). Must be globally unique within the card. Only letters, numbers, and underscores; must start with a letter and not exceed 20 characters. |
| `margin` | No | String | `"0"` | Margin of the component (JSON 2.0). Value range is [-99,99]px. Accepts a single value (e.g. `"10px"`), double value (e.g. `"4px 0"`), or four values (e.g. `"4px 0 4px 0"`) for top, right, bottom, left. |
| `size` | No | String | `"medium"` | Size of the person's avatar. Values: `extra_small`, `small`, `medium`, `large`. |
| `user_id` | Yes | String | — | ID of the person. Accepts an Open ID, Union ID, or User ID. |
| `show_avatar` | No | Boolean | `true` | Whether to display the user's avatar. |
| `show_name` | No | Boolean | `false` | Whether to display the user's name. |
| `style` | No | String | `"normal"` | Display style. Values: `normal` (default style) or `capsule` (capsule style). |

### User ID Types

- **Open ID** -- Identifies a user within a specific application. The same user has different Open IDs in different applications.
- **Union ID** -- Identifies a user under a specific app developer. The same user has the same Union ID across apps under the same developer but different Union IDs under different developers.
- **User ID** -- Identifies a user within a specific tenant. Within the same tenant, a user's User ID remains consistent across all applications (including store apps).

## Example

Replace the `user_id` values below with actual user IDs.

```json
{
  "schema": "2.0",
  "header": {
      "template": "blue",
      "title": {
          "content": "人员示例",
          "tag": "plain_text"
      }
  },
  "body": {
      "elements": [
          {
              "tag": "markdown",
              "content": "**extra_small 尺寸，默认样式**"
          },
          {
              "tag": "person",
              "size": "extra_small",
              "user_id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef",
              "show_avatar": true,
              "show_name": true,
              "style": "normal"
          },
          {
              "tag": "markdown",
              "content": "**small 尺寸，胶囊样式**"
          },
          {
              "tag": "person",
              "size": "small",
              "user_id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef",
              "show_avatar": true,
              "show_name": true,
              "style": "capsule"
          },
          {
              "tag": "markdown",
              "content": "**medium 尺寸，默认样式**"
          },
          {
              "tag": "person",
              "size": "medium",
              "user_id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef",
              "show_avatar": true,
              "show_name": true,
              "style": "normal"
          },
          {
              "tag": "markdown",
              "content": "**large 尺寸，胶囊样式**"
          },
          {
              "tag": "person",
              "size": "large",
              "user_id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef",
              "show_avatar": true,
              "show_name": true,
              "style": "capsule"
          }
      ]
  }
}
```

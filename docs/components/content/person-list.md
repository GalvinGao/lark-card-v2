# Person List

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/user-list>

The person list component displays multiple users' names and avatars. You can use this component by passing in a person's `open_id`, `user_id`, or `union_id`.

If you send a card containing a person list component using a specified application, you must ensure that the application has access to the user IDs. Otherwise, the person list component will not display person information.

## JSON Structure

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "person_list",
        "element_id": "custom_id",
        "margin": "0px 0px 0px 0px",
        "drop_invalid_user_id": false,
        "lines": 1,
        "show_name": true,
        "show_avatar": true,
        "size": "large",
        "persons": [
          {
            "id": "ou_0fdb0e7663af7128e7d9f8adeb2abcef"
          },
          {
            "id": "ou_0fdb0e7663af7128e7d9f8adeb2abcef"
          }
        ],
        "icon": {
          "tag": "standard_icon",
          "token": "chat-forbidden_outlined",
          "color": "orange",
          "img_key": "img_v2_38811724"
        },
        "ud_icon": {
          "token": "chat-forbidden_outlined",
          "style": {
            "color": "red"
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
| `tag` | Yes | String | `"person_list"` | Component tag. Fixed value: `person_list`. |
| `element_id` | No | String | — | Unique identifier for the component (JSON 2.0). Must be globally unique within the card. Only letters, numbers, and underscores; must start with a letter and not exceed 20 characters. |
| `margin` | No | String | `"0"` | Margin of the component (JSON 2.0). Value range is [-99,99]px. Accepts a single value (e.g. `"10px"`), double value (e.g. `"4px 0"`), or four values (e.g. `"4px 0 4px 0"`) for top, right, bottom, left. |
| `drop_invalid_user_id` | No | Boolean | `false` | Whether to ignore invalid user IDs. When `false`, an error is raised and the list of invalid IDs is returned. |
| `lines` | No | Int | — | Maximum number of display lines. Default is unlimited. |
| `show_name` | No | Boolean | `true` | Whether to display user names. |
| `show_avatar` | No | Boolean | `false` | Whether to display user avatars. When `show_name` is `false` and multiple IDs are in `persons`, the list is shown in a stacked avatar ("hulu string") style. |
| `size` | No | String | `"medium"` | Size of person avatars. Values: `extra_small`, `small`, `medium`, `large`. |
| `persons` | Yes | Array | — | List of persons. |
| `persons[].id` | Yes | String | — | ID of the person. Accepts an Open ID, Union ID, or User ID. |
| `icon` | No | Object | — | Prefix icon. Supports custom or icon library usage. |
| `icon.tag` | No | String | — | Icon type. Values: `standard_icon` (icon library) or `custom_icon` (custom image). |
| `icon.token` | No | String | — | Icon token from the icon library. Effective when `icon.tag` is `standard_icon`. |
| `icon.color` | No | String | — | Icon color. Supports line and surface icons (tokens ending in `outlined` or `filled`). Effective when `icon.tag` is `standard_icon`. |
| `icon.img_key` | No | String | — | Image key for a custom prefix icon. Effective when `icon.tag` is `custom_icon`. Obtain via the Upload Image API. |
| `ud_icon` | No | Object | — | Prefix icon from the icon library. If both `icon` and `ud_icon` are configured, only `icon` takes effect. |
| `ud_icon.token` | No | String | — | Icon token from the icon library. |
| `ud_icon.style` | No | Object | — | Icon style. Supports custom icon colors. |
| `ud_icon.style.color` | No | String | — | Icon color. Supports outlined and filled icons. The customization tool currently does not support custom icon colors. |

### User ID Types

- **Open ID** -- Identifies a user within a specific application. The same user has different Open IDs in different applications.
- **Union ID** -- Identifies a user under a specific app developer. The same user has the same Union ID across apps under the same developer but different Union IDs under different developers.
- **User ID** -- Identifies a user within a specific tenant. Within the same tenant, a user's User ID remains consistent across all applications (including store apps).

## Example

Replace the `id` values below with actual user IDs.

```json
{
  "schema": "2.0",
  "header": {
    "template": "blue",
    "title": {
      "content": "人员列表示例",
      "tag": "plain_text"
    }
  },
  "body": {
    "elements": [
      {
        "tag": "markdown",
        "content": "仅名字："
      },
      {
        "tag": "person_list",
        "persons": [
          {
            "id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef"
          },
          {
            "id": "ou_f9d24af786a14340721288cda6aabcef"
          },
          {
            "id": "ou_b824f85713725c632e78887dc7fabcef"
          }
        ]
      },
      {
        "tag": "markdown",
        "content": "名字+头像："
      },
      {
        "tag": "person_list",
        "show_name": true,
        "show_avatar": true,
        "lines": 2,
        "size": "small",
        "persons": [
          {
            "id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef"
          },
          {
            "id": "ou_f9d24af786a14340721288cda6aabcef"
          },
          {
            "id": "ou_b824f85713725c632e78887dc7fabcef"
          }
        ]
      },
      {
        "tag": "markdown",
        "content": "名字+头像+icon："
      },
      {
        "tag": "person_list",
        "show_name": true,
        "show_avatar": true,
        "lines": 2,
        "size": "small",
        "persons": [
          {
            "id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef"
          },
          {
            "id": "ou_f9d24af786a14340721288cda6aabcef"
          },
          {
            "id": "ou_b824f85713725c632e78887dc7fabcef"
          }
        ],
        "icon": {
          "tag": "standard_icon",
          "token": "group_outlined",
          "color": "blue"
        }
      },
      {
        "tag": "markdown",
        "content": "名字葫芦串："
      },
      {
        "tag": "person_list",
        "persons": [
          {
            "id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef"
          },
          {
            "id": "ou_f9d24af786a14340721288cda6aabcef"
          },
          {
            "id": "ou_b824f85713725c632e78887dc7fabcef"
          },
          {
            "id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef"
          },
          {
            "id": "ou_f9d24af786a14340721288cda6aabcef"
          },
          {
            "id": "ou_b824f85713725c632e78887dc7fabcef"
          }
        ],
        "size": "small",
        "show_avatar": true,
        "show_name": false
      }
    ]
  }
}
```

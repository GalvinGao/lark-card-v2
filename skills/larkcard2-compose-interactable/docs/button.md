# Button

**Tag:** `button`

## Minimal Example

```json
{
  "tag": "button",
  "text": { "tag": "plain_text", "content": "Click Me" },
  "type": "primary",
  "size": "medium",
  "width": "default",
  "icon": { "tag": "standard_icon", "token": "thumbsup_outlined", "color": "white" },
  "behaviors": [
    { "type": "callback", "value": { "action": "approve" } }
  ],
  "disabled": false,
  "margin": "0px"
}
```

## Key Fields

- **text**: `{ "tag": "plain_text", "content": "..." }` — button label (required)
- **type**: Visual style
  - Bordered: `"default"` | `"primary"` | `"danger"`
  - Text-only: `"text"` | `"primary_text"` | `"danger_text"`
  - Filled: `"primary_filled"` | `"danger_filled"`
  - Special: `"laser"` (animated highlight effect)
- **size**: `"tiny"` (24/28px) | `"small"` (28/28px) | `"medium"` (32/36px) | `"large"` (40/48px) — values are PC/mobile heights
- **width**: `"default"` | `"fill"` | `"[100,∞)px"`
- **icon**: Optional prefix icon — `{ "tag": "standard_icon", "token": "icon_token", "color": "color_name" }`
- **behaviors**: Array of `open_url` and/or `callback` actions
- **confirm**: Optional confirmation dialog before executing behaviors
- **hover_tips**: `{ "tag": "plain_text", "content": "..." }` — tooltip on hover (PC only)
- **disabled**: Boolean to grey out the button
- **disabled_tips**: `{ "tag": "plain_text", "content": "..." }` — tooltip when disabled
- **margin**: Spacing, e.g. `"0px"`, `"4px 0px"`

## In Form Context

When inside a `form` container, add:

- **name**: Unique identifier within the form
- **form_action_type**: `"submit"` | `"reset"` — determines button behavior on click

## Notes

- A button can have both `open_url` and `callback` behaviors simultaneously.
- The `text` field only supports `plain_text`, not `lark_md`.
- Icon color should contrast with the button type for visibility (e.g., `"white"` for `"primary"`).

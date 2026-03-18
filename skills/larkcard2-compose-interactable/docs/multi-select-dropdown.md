# Multi-Select Dropdown

**Tag:** `multi_select_static`

**Must be inside a `form` container.**

## Minimal Example

```json
{
  "tag": "multi_select_static",
  "name": "multi_1",
  "placeholder": { "tag": "plain_text", "content": "Select..." },
  "options": [
    { "text": { "tag": "plain_text", "content": "A" }, "value": "a" },
    { "text": { "tag": "plain_text", "content": "B" }, "value": "b" }
  ],
  "selected_values": ["a"],
  "type": "default",
  "width": "fill",
  "required": true,
  "margin": "0px"
}
```

## Key Fields

- **name**: Required — unique identifier within the form
- **type**: `"default"` (bordered) | `"text"` (borderless)
- **placeholder**: `{ "tag": "plain_text", "content": "..." }` — hint when nothing selected
- **options**: Array of selectable items (required)
  - **text**: `{ "tag": "plain_text", "content": "..." }` — display label
  - **value**: String — unique value for this option
  - **icon**: Optional icon — `{ "tag": "standard_icon", "token": "..." }`
- **selected_values**: Array of option `value` strings to pre-select
- **width**: `"default"` | `"fill"` | `"[100,∞)px"`
- **required**: Boolean — at least one selection required on form submit
- **disabled**: Boolean to grey out the dropdown
- **confirm**: Optional confirmation dialog

## Notes

- This component **requires** a `form` container parent — it cannot be used standalone.
- The `name` field is mandatory and must be unique within the form.
- `selected_values` entries must match existing option `value` strings.

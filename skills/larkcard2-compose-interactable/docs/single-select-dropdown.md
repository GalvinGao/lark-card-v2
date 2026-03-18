# Single-Select Dropdown

**Tag:** `select_static`

## Minimal Example

```json
{
  "tag": "select_static",
  "placeholder": { "tag": "plain_text", "content": "Select..." },
  "options": [
    { "text": { "tag": "plain_text", "content": "Option 1" }, "value": "1", "icon": { "tag": "standard_icon", "token": "check_outlined" } },
    { "text": { "tag": "plain_text", "content": "Option 2" }, "value": "2" }
  ],
  "initial_option": "1",
  "type": "default",
  "width": "default",
  "behaviors": [{ "type": "callback", "value": { "key": "val" } }],
  "margin": "0px"
}
```

## Key Fields

- **type**: `"default"` (bordered) | `"text"` (borderless)
- **placeholder**: `{ "tag": "plain_text", "content": "..." }` — hint when nothing selected
- **options**: Array of selectable items (required)
  - **text**: `{ "tag": "plain_text", "content": "..." }` — display label
  - **value**: String — unique value for this option (required)
  - **icon**: Optional icon — `{ "tag": "standard_icon", "token": "..." }`
- **initial_option**: String matching an option's `value` to pre-select
- **initial_index**: Integer — `0` = none selected, `1` = first option, `2` = second, etc.
- **width**: `"default"` | `"fill"` | `"[100,∞)px"`
- **behaviors**: Array of `callback` actions
- **confirm**: Optional confirmation dialog

## In Form Context

When inside a `form` container:

- **name**: Unique identifier within the form
- **required**: Boolean — mandatory selection on submit
- **disabled**: Boolean to grey out the dropdown

## Notes

- Use either `initial_option` or `initial_index`, not both.
- Each option's `value` must be unique within the options array.

# Single-Select Person Picker

**Tag:** `select_person`

## Minimal Example

```json
{
  "tag": "select_person",
  "placeholder": { "tag": "plain_text", "content": "Select person" },
  "options": [
    { "value": "ou_xxxx" },
    { "value": "ou_yyyy" }
  ],
  "initial_option": "ou_xxxx",
  "type": "default",
  "width": "default",
  "behaviors": [{ "type": "callback", "value": { "key": "val" } }],
  "margin": "0px"
}
```

## Key Fields

- **type**: `"default"` (bordered) | `"text"` (borderless)
- **placeholder**: `{ "tag": "plain_text", "content": "..." }` — hint when no one selected
- **options**: Array of candidate users
  - **value**: `open_id` of a user (e.g., `"ou_xxxx"`)
  - If `options` is an empty array `[]`, all members of the current chat are shown as candidates
- **initial_option**: `open_id` string to pre-select a user
- **width**: `"default"` | `"fill"` | `"[100,∞)px"`
- **behaviors**: Array of `callback` actions
- **confirm**: Optional confirmation dialog

## In Form Context

When inside a `form` container:

- **name**: Unique identifier within the form
- **required**: Boolean — mandatory selection on submit
- **disabled**: Boolean to grey out the picker

## Notes

- Person options only need `value` (the `open_id`); display name and avatar are resolved automatically by Lark.
- Pass an empty `options` array to let users pick from all chat members.

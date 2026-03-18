# Multi-Select Person Picker

**Tag:** `multi_select_person`

**Must be inside a `form` container.**

## Minimal Example

```json
{
  "tag": "multi_select_person",
  "name": "person_multi",
  "placeholder": { "tag": "plain_text", "content": "Select people" },
  "options": [
    { "value": "ou_xxxx" },
    { "value": "ou_yyyy" }
  ],
  "selected_values": [],
  "type": "default",
  "width": "fill",
  "required": true,
  "margin": "0px"
}
```

## Key Fields

- **name**: Required — unique identifier within the form
- **type**: `"default"` (bordered) | `"text"` (borderless)
- **placeholder**: `{ "tag": "plain_text", "content": "..." }` — hint when no one selected
- **options**: Array of candidate users
  - **value**: `open_id` of a user (e.g., `"ou_xxxx"`)
  - If `options` is an empty array `[]`, all members of the current chat are shown as candidates
- **selected_values**: Array of `open_id` strings to pre-select
- **width**: `"default"` | `"fill"` | `"[100,∞)px"`
- **required**: Boolean — at least one person required on form submit
- **disabled**: Boolean to grey out the picker

## Notes

- This component **requires** a `form` container parent — it cannot be used standalone.
- The `name` field is mandatory and must be unique within the form.
- Person display names and avatars are resolved automatically by Lark from the `open_id`.
- Pass an empty `options` array to let users pick from all chat members.

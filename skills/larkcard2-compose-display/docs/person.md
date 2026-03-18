# Person — `person`

Tag: `person`

## JSON Example

```json
{
  "tag": "person",
  "user_id": "ou_xxxx",
  "size": "medium",
  "show_avatar": true,
  "show_name": true,
  "style": "normal",
  "margin": "0px"
}
```

## Key Fields

- **user_id**: Required. Supports `open_id`, `user_id`, or `union_id`.
- **size**: `"extra_small"` | `"small"` | `"medium"` | `"large"`
- **show_avatar**: Boolean, default `true`. Show the user's avatar.
- **show_name**: Boolean, default `false`. Show the user's display name.
- **style**: `"normal"` (plain text) | `"capsule"` (pill/badge style)
- **margin**: Component margin

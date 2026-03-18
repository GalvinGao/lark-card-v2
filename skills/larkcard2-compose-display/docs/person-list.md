# Person List — `person_list`

Tag: `person_list`

## JSON Example

```json
{
  "tag": "person_list",
  "persons": [
    { "id": "ou_xxxx" },
    { "id": "ou_yyyy" }
  ],
  "size": "small",
  "show_name": true,
  "show_avatar": true,
  "lines": 2,
  "icon": { "tag": "standard_icon", "token": "group_outlined", "color": "blue" },
  "margin": "0px"
}
```

## Key Fields

- **persons**: Required. Array of `{ "id": "..." }` objects. Supports `open_id`, `user_id`, or `union_id`.
- **size**: `"extra_small"` | `"small"` | `"medium"` | `"large"`
- **show_name**: Boolean, default `true`. Show display names.
- **show_avatar**: Boolean, default `false`. Show avatars.
- **lines**: Max display lines (integer). Overflow is truncated.
- **icon**: Optional prefix icon (standard_icon or custom_icon).
- **drop_invalid_user_id**: Boolean, default `false`. When `false`, invalid user IDs cause an error. When `true`, invalid IDs are silently dropped.
- **margin**: Component margin

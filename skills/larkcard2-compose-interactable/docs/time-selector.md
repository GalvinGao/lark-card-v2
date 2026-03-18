# Time Selector

**Tag:** `picker_time`

## Minimal Example

```json
{
  "tag": "picker_time",
  "placeholder": { "tag": "plain_text", "content": "Select time" },
  "initial_time": "09:00",
  "width": "default",
  "behaviors": [{ "type": "callback", "value": { "key": "val" } }],
  "margin": "0px"
}
```

## Key Fields

- **placeholder**: `{ "tag": "plain_text", "content": "..." }` — hint when no time selected
- **initial_time**: Pre-selected time in `"HH:mm"` format (e.g., `"09:00"`)
- **width**: `"default"` | `"fill"` | `"[100,∞)px"`
- **behaviors**: Array of `callback` actions
- **confirm**: Optional confirmation dialog
- **disabled**: Boolean to grey out the selector
- **value**: Custom callback value object
- **margin**: Spacing, e.g. `"0px"`

## In Form Context

When inside a `form` container:

- **name**: Unique identifier within the form
- **required**: Boolean — mandatory selection on submit

## Notes

- Time format is strictly `HH:mm` (24-hour).

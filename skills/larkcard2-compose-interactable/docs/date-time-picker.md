# Date-Time Picker

**Tag:** `picker_datetime`

## Minimal Example

```json
{
  "tag": "picker_datetime",
  "placeholder": { "tag": "plain_text", "content": "Select date and time" },
  "initial_datetime": "2024-01-01 08:00",
  "width": "default",
  "behaviors": [{ "type": "callback", "value": { "key": "val" } }],
  "margin": "0px"
}
```

## Key Fields

- **placeholder**: `{ "tag": "plain_text", "content": "..." }` — hint when nothing selected
- **initial_datetime**: Pre-selected value in `"yyyy-MM-dd HH:mm"` format (e.g., `"2024-01-01 08:00"`)
- **width**: `"default"` | `"fill"` | `"[100,∞)px"`
- **behaviors**: Array of `callback` actions
- **confirm**: Optional confirmation dialog
- **disabled**: Boolean to grey out the picker
- **value**: Custom callback value object
- **margin**: Spacing, e.g. `"0px"`

## In Form Context

When inside a `form` container:

- **name**: Unique identifier within the form
- **required**: Boolean — mandatory selection on submit

## Notes

- Date-time format is strictly `yyyy-MM-dd HH:mm` (24-hour time).
- Combines both date and time selection in a single component.

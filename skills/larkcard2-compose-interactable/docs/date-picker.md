# Date Picker

**Tag:** `date_picker`

## Minimal Example

```json
{
  "tag": "date_picker",
  "placeholder": { "tag": "plain_text", "content": "Select date" },
  "initial_date": "2024-01-01",
  "width": "default",
  "behaviors": [{ "type": "callback", "value": { "key": "val" } }],
  "margin": "0px"
}
```

## Key Fields

- **placeholder**: `{ "tag": "plain_text", "content": "..." }` — hint when no date selected
- **initial_date**: Pre-selected date in `"yyyy-MM-dd"` format (e.g., `"2024-01-01"`)
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

- The callback payload includes the selected date along with the user's timezone.
- Date format is strictly `yyyy-MM-dd`.

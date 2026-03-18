# Input

**Tag:** `input`

## Minimal Example

```json
{
  "tag": "input",
  "placeholder": { "tag": "plain_text", "content": "Enter text..." },
  "default_value": "",
  "max_length": 1000,
  "input_type": "text",
  "width": "default",
  "label": { "tag": "plain_text", "content": "Name:" },
  "label_position": "top",
  "behaviors": [{ "type": "callback", "value": { "key": "val" } }],
  "margin": "0px"
}
```

## Key Fields

- **input_type**: `"text"` (single line) | `"multiline_text"` (multi-line) | `"password"` (masked)
- **placeholder**: `{ "tag": "plain_text", "content": "..." }` — hint text when empty
- **default_value**: Pre-filled text string
- **max_length**: 1–1000 (default 1000)
- **rows**: Default display rows for `multiline_text` (default 5)
- **auto_resize**: Boolean — auto-grow height for multiline (PC only)
- **max_rows**: Maximum rows when `auto_resize` is true
- **show_icon**: Boolean — show a prefix icon for `password` type
- **label**: `{ "tag": "plain_text", "content": "..." }` — field label
- **label_position**: `"top"` | `"left"` — label placement relative to the input
- **width**: `"default"` | `"fill"` | `"[100,∞)px"`
- **behaviors**: Array of `callback` actions (triggered on input change/submit)
- **confirm**: Optional confirmation dialog
- **disabled**: Boolean to grey out the input
- **disabled_tips**: `{ "tag": "plain_text", "content": "..." }` — tooltip when disabled
- **value**: Custom callback value object
- **margin**: Spacing, e.g. `"0px"`

## In Form Context

When inside a `form` container:

- **name**: Required, unique identifier within the form
- **required**: Boolean — mark field as mandatory on form submit

## Notes

- `auto_resize` and `max_rows` only apply to `multiline_text` and only work on PC.
- `show_icon` only applies to `password` type.
- The `label` supports only `plain_text`.

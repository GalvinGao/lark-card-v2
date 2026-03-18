# Form Container — `form`

Tag: `form`

## JSON Example

```json
{
  "tag": "form",
  "name": "form_1",
  "elements": [
    {
      "tag": "input",
      "name": "field1",
      "required": true,
      "placeholder": { "tag": "plain_text", "content": "Enter..." }
    },
    {
      "tag": "button",
      "text": { "tag": "plain_text", "content": "Submit" },
      "type": "primary",
      "form_action_type": "submit",
      "name": "btn_submit"
    }
  ]
}
```

## Key Fields

- **name**: Required. Unique form identifier string.
- **elements**: Array of nested components. Can include display and interactive components.
- **direction**: `"vertical"` (default) | `"horizontal"`
- **padding**: Inner padding
- **margin**: Component margin
- **horizontal_spacing**: Spacing between elements horizontally
- **vertical_spacing**: Spacing between elements vertically
- **horizontal_align**: `"left"` | `"center"` | `"right"`
- **vertical_align**: `"top"` | `"center"` | `"bottom"`

## Rules

- Forms can **only** be placed at card root level (directly in `body.elements`).
- **No nested forms** — a form cannot contain another form.
- **No tables** inside forms.
- All interactive components inside a form **must** have a `name` field (globally unique across the card).
- Buttons inside a form must use `form_action_type`:
  - `"submit"` — submits the form (collects all named field values)
  - `"reset"` — resets all fields to defaults
- Interactive components can set `required: true` for client-side validation on submit.

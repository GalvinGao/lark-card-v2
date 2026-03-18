# Checker

**Tag:** `checker`

## Minimal Example

```json
{
  "tag": "checker",
  "name": "task_1",
  "checked": false,
  "text": { "tag": "plain_text", "content": "Complete the report" },
  "overall_checkable": true,
  "checked_style": {
    "show_strikethrough": true,
    "opacity": 0.5
  },
  "button_area": {
    "pc_display_rule": "on_hover",
    "buttons": [
      {
        "tag": "button",
        "type": "text",
        "size": "small",
        "text": { "tag": "plain_text", "content": "" },
        "icon": { "tag": "standard_icon", "token": "forward-com_outlined", "color": "grey-500" },
        "behaviors": [{ "type": "callback", "value": { "key": "share" } }]
      }
    ]
  },
  "behaviors": [{ "type": "callback", "value": { "key": "check_task" } }],
  "padding": "2px",
  "margin": "0px"
}
```

## Key Fields

- **checked**: Boolean — initial checked state
- **text**: `{ "tag": "plain_text", "content": "..." }` or `{ "tag": "lark_md", "content": "..." }` — the label
- **overall_checkable**: Boolean (default `true`) — enables hover shadow effect on the entire row; when true, clicking anywhere on the row toggles the check
- **checked_style**: Visual changes when checked
  - **show_strikethrough**: Boolean — strike-through text when checked
  - **opacity**: Number 0–1 — dim the item when checked
- **button_area**: Action buttons displayed alongside the checker
  - **pc_display_rule**: `"always"` | `"on_hover"` — when to show buttons on PC
  - **buttons**: Array of up to 3 button elements, each with:
    - **type**: `"text"` | `"primary_text"` | `"danger_text"`
    - **size**: Button size (typically `"small"`)
    - **text**: Button label (can be empty `""` for icon-only)
    - **icon**: Optional icon
    - **behaviors**: Button-specific callback actions
- **behaviors**: Array of `callback` actions for the check toggle itself (no `open_url` support on the checker)
- **confirm**: Optional confirmation dialog before toggling
- **hover_tips**: `{ "tag": "plain_text", "content": "..." }` — tooltip on hover
- **disabled**: Boolean to grey out the checker
- **disabled_tips**: `{ "tag": "plain_text", "content": "..." }` — tooltip when disabled
- **padding**: Inner padding, e.g. `"2px"`
- **margin**: Outer spacing, e.g. `"0px"`

## In Form Context

When inside a `form` container:

- **name**: Required — unique identifier within the form

## Notes

- Requires **Lark V7.9+** on the client.
- The checker's own `behaviors` only support `callback`, not `open_url`. Buttons inside `button_area` can use both.
- `button_area.buttons` supports a maximum of 3 buttons.
- Use `checked_style` to give clear visual feedback for completed items (strikethrough + reduced opacity is the common pattern).

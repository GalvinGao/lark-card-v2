---
name: larkcard2-compose-interactable
description: "Lark Card V2 interactive component reference — buttons, inputs, dropdowns, person pickers, date/time pickers, overflow menus, image pickers, and checkers. Use whenever composing or modifying Lark/Feishu message card JSON v2 with interactive elements that accept user input or trigger callbacks. Requires familiarity with larkcard2-compose-display for card structure and layout."
---

# Lark Card V2 — Interactive Components

This skill covers interactive components for Lark Card V2: elements that accept user input, trigger callbacks, or open URLs.

**Prerequisite:** For card structure basics, display components, and containers, see the `larkcard2-compose-display` skill.

## Component Index

| Tag | Description | Doc |
|-----|-------------|-----|
| `button` | Clickable button with icon, styles, URL/callback behaviors | [docs/button.md](docs/button.md) |
| `input` | Text input field with single-line, multi-line, and password modes | [docs/input.md](docs/input.md) |
| `overflow` | Collapsed menu that expands to show a list of options | [docs/overflow.md](docs/overflow.md) |
| `select_static` | Single-select dropdown with static option list | [docs/single-select-dropdown.md](docs/single-select-dropdown.md) |
| `multi_select_static` | Multi-select dropdown with static options (form only) | [docs/multi-select-dropdown.md](docs/multi-select-dropdown.md) |
| `select_person` | Single-select person picker from org or chat members | [docs/single-select-person-picker.md](docs/single-select-person-picker.md) |
| `multi_select_person` | Multi-select person picker for choosing multiple users (form only) | [docs/multi-select-person-picker.md](docs/multi-select-person-picker.md) |
| `date_picker` | Date picker returning yyyy-MM-dd formatted date | [docs/date-picker.md](docs/date-picker.md) |
| `picker_time` | Time selector returning HH:mm formatted time | [docs/time-selector.md](docs/time-selector.md) |
| `picker_datetime` | Combined date-time picker returning yyyy-MM-dd HH:mm | [docs/date-time-picker.md](docs/date-time-picker.md) |
| `select_img` | Image picker for single or multi image selection | [docs/image-picker.md](docs/image-picker.md) |
| `checker` | Checkbox/task item with checked state and action buttons | [docs/checker.md](docs/checker.md) |

## Common Interactive Patterns

### `behaviors` Array

Most interactive components accept a `behaviors` array defining what happens on interaction:

```json
"behaviors": [
  {
    "type": "open_url",
    "default_url": "https://example.com",
    "pc_url": "https://example.com/pc",
    "ios_url": "https://example.com/ios",
    "android_url": "https://example.com/android"
  },
  {
    "type": "callback",
    "value": { "key": "value" }
  }
]
```

- `open_url`: Opens a URL. `default_url` is required; platform-specific URLs are optional overrides.
- `callback`: Sends a callback to your server with the `value` object as payload.

### `confirm` Dialog

Attach a confirmation dialog before the action executes:

```json
"confirm": {
  "title": { "tag": "plain_text", "content": "Are you sure?" },
  "text": { "tag": "plain_text", "content": "This action cannot be undone." }
}
```

### Form Integration

When a component is placed inside a `form` container:

- Add `"name": "unique_field_name"` — must be unique within the form.
- Set `"required": true` to make the field mandatory on submit.
- Buttons use `"form_action_type": "submit"` or `"reset"` to control form behavior.

### `disabled` + `disabled_tips`

```json
"disabled": true,
"disabled_tips": { "tag": "plain_text", "content": "You don't have permission" }
```

Greys out the component and shows a tooltip explaining why.

### `hover_tips`

```json
"hover_tips": { "tag": "plain_text", "content": "Click to approve" }
```

Shows a tooltip on hover (PC only).

### Common `width` Values

- `"default"` — auto-width based on content
- `"fill"` — expand to fill available space
- `"[100,∞)px"` — fixed pixel width, e.g. `"200px"` (minimum 100)

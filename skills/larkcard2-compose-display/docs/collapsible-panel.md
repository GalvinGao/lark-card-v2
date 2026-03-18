# Collapsible Panel — `collapsible_panel`

Tag: `collapsible_panel`

## JSON Example

```json
{
  "tag": "collapsible_panel",
  "expanded": false,
  "header": {
    "title": { "tag": "plain_text", "content": "Panel Title" },
    "vertical_align": "center",
    "icon": { "tag": "standard_icon", "token": "down-small-ccm_outlined", "size": "16px 16px" },
    "icon_position": "right",
    "icon_expanded_angle": -180
  },
  "border": { "color": "grey", "corner_radius": "5px" },
  "background_color": "grey",
  "padding": "8px",
  "vertical_spacing": "8px",
  "elements": []
}
```

## Key Fields

- **expanded**: Boolean, default `false`. Whether the panel starts expanded.

### header

- **header.title**: Required. `{ "tag": "plain_text", "content": "..." }` or markdown tag.
- **header.background_color**: Color enum for the header bar.
- **header.vertical_align**: `"top"` | `"center"` | `"bottom"`
- **header.padding**: Header inner padding
- **header.width**: `"fill"` | `"auto"` | `"auto_when_fold"` (auto width when collapsed, fill when expanded)
- **header.icon**: Expand/collapse icon. Standard or custom icon, with optional `size` (e.g., `"16px 16px"`).
- **header.icon_position**: `"left"` | `"right"` | `"follow_text"`
- **header.icon_expanded_angle**: Rotation when expanded: `-180` | `-90` | `90` | `180`

### Panel body

- **border**: `{ "color": "grey", "corner_radius": "5px" }` — panel border styling
- **background_color**: Color enum for the panel body background
- **direction**: `"vertical"` (default) | `"horizontal"`
- **padding**: Panel body inner padding
- **margin**: Component margin
- **horizontal_spacing**: Spacing between elements horizontally
- **vertical_spacing**: Spacing between elements vertically
- **horizontal_align**: `"left"` | `"center"` | `"right"`
- **vertical_align**: `"top"` | `"center"` | `"bottom"`
- **elements**: Nested components. `form` containers are **not** allowed inside collapsible panels.

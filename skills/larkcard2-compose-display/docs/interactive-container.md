# Interactive Container — `interactive_container`

Tag: `interactive_container`

## JSON Example

```json
{
  "tag": "interactive_container",
  "width": "fill",
  "height": "auto",
  "background_style": "default",
  "has_border": true,
  "border_color": "grey",
  "corner_radius": "8px",
  "padding": "4px 12px",
  "behaviors": [
    { "type": "callback", "value": { "key": "value" } }
  ],
  "elements": []
}
```

## Key Fields

- **width**: `"fill"` | `"auto"` | `"[16,999]px"`
- **height**: `"auto"` | `"[10,999]px"`
- **background_style**: `"default"` | `"laser"` | color enum
- **has_border**: Boolean. Show a border around the container.
- **border_color**: Color enum (only applies when `has_border` is `true`)
- **corner_radius**: `"[0,...]px"` or `"[0,100]%"`
- **padding**: Inner padding
- **margin**: Component margin
- **behaviors**: Required. Array of behavior objects:
  - `{ "type": "open_url", "default_url": "https://...", "pc_url": "...", "ios_url": "...", "android_url": "..." }` — opens a URL on click
  - `{ "type": "callback", "value": { ... } }` — triggers a server callback with the given payload
- **disabled**: Boolean. Disable click interaction.
- **disabled_tips**: Tooltip shown when disabled and hovered.
- **hover_tips**: Tooltip shown on hover (when enabled).
- **confirm**: Confirmation dialog before executing behavior. `{ "title": { "tag": "plain_text", "content": "..." }, "text": { "tag": "plain_text", "content": "..." } }`
- **elements**: Nested components. Supports all components **except** `form` and `table`.

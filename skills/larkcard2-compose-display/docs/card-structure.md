# Card Structure

Overall JSON structure for Lark Card v2.

## Minimal JSON Skeleton

```json
{
  "schema": "2.0",
  "config": {
    "update_multi": true,
    "width_mode": "default"
  },
  "header": {
    "title": { "tag": "plain_text", "content": "Card Title" },
    "template": "blue"
  },
  "body": {
    "elements": []
  }
}
```

## Top-Level Fields

- **schema**: Must be `"2.0"`.
- **config**: Optional card-level configuration.
- **card_link**: Optional click-through link for the entire card.
- **header**: Card title area (see `title.md`).
- **body**: Contains the `elements` array with all card components.

## config

| Field | Type | Description |
|---|---|---|
| `streaming_mode` | boolean | Enable streaming update mode |
| `enable_forward` | boolean | Allow card to be forwarded |
| `update_multi` | boolean | Update card for all recipients (not just sender) |
| `width_mode` | string | `"default"` (fixed), `"fill"` (full width), `"compact"` (auto-shrink) |
| `style` | object | Global styles (see below) |

### config.style — Custom Text Sizes and Colors

Define per-platform text sizes and light/dark custom colors in `config.style`:

```json
"style": {
  "text_size": {
    "cus-heading": {
      "default": "medium",
      "pc": "heading-1",
      "mobile": "heading-2"
    }
  },
  "color": {
    "cus-brand": {
      "light_mode": "rgba(20,86,240,0.8)",
      "dark_mode": "rgba(117,164,255,0.8)"
    }
  }
}
```

Then use `"cus-heading"` in any `text_size` field, and `"cus-brand"` in any color field. See `colors.md` for the full color enum table and RGBA usage.

## card_link

Makes the entire card clickable. Fields:

| Field | Type | Description |
|---|---|---|
| `url` | string | Default URL |
| `pc_url` | string | Desktop-specific URL |
| `ios_url` | string | iOS-specific URL |
| `android_url` | string | Android-specific URL |

## body

| Field | Type | Description |
|---|---|---|
| `direction` | string | `"vertical"` (default) or `"horizontal"` |
| `padding` | string | e.g. `"12px"`, `"4px 12px"`, `"4px 12px 4px 12px"` |
| `horizontal_spacing` | string | Spacing between horizontal child elements |
| `vertical_spacing` | string | Spacing between vertical child elements |
| `horizontal_align` | string | `"left"`, `"center"`, `"right"` |
| `vertical_align` | string | `"top"`, `"center"`, `"bottom"` |
| `elements` | array | Array of card components |

## Common Component Attributes

These fields are available on most components:

- **element_id**: Globally unique identifier. Rules: letters, numbers, underscores only; must start with a letter; max 20 characters.
- **margin**: Spacing around the component. Format: `"Xpx"` (all sides), `"Xpx Ypx"` (top-bottom, left-right), `"Xpx Ypx Xpx Ypx"` (top, right, bottom, left). Supports negative values.

## Constraints

- `schema` must be `"2.0"`.
- Maximum **200 elements** per card.
- Tables and charts can only be placed at the card root level (inside `body.elements`), not nested inside other containers.

# Image Picker

**Tag:** `select_img`

## Minimal Example

```json
{
  "tag": "select_img",
  "name": "img_select_1",
  "multi_select": false,
  "layout": "bisect",
  "options": [
    { "img_key": "img_v3_xxxx", "value": "img1" },
    { "img_key": "img_v3_yyyy", "value": "img2" }
  ],
  "width": "default",
  "margin": "0px"
}
```

## Key Fields

- **multi_select**: Boolean — `false` for single-select, `true` for multi-select
- **layout**: `"bisect"` (2-column grid) | `"trisect"` (3-column grid)
- **options**: Array of image choices (required)
  - **img_key**: Image key uploaded to Lark (e.g., `"img_v3_xxxx"`)
  - **value**: String — unique callback value for this image
- **width**: `"default"` | `"fill"` | `"[100,∞)px"`
- **style**: Optional styling object
  - **aspect_ratio**: Image aspect ratio
  - **corner_radius**: Border radius for images
- **margin**: Spacing, e.g. `"0px"`

## In Form Context

When inside a `form` container:

- **name**: Unique identifier within the form
- **required**: Boolean — at least one image selection required
- **disabled**: Boolean to grey out the picker

## Notes

- Requires **Lark V7.6+** on the client.
- Each option's `value` must be unique within the options array.
- Images are referenced by `img_key`, which must be uploaded via the Lark image upload API first.

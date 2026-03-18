# Multi-Image Layout — `img_combination`

Tag: `img_combination`

## JSON Example

```json
{
  "tag": "img_combination",
  "combination_mode": "double",
  "img_list": [
    { "img_key": "img_v3_xxxx" },
    { "img_key": "img_v3_yyyy" }
  ],
  "corner_radius": "12px",
  "combination_transparent": false,
  "margin": "0px"
}
```

## Key Fields

- **combination_mode**: Required. Layout arrangement:
  - `"double"` — 2 images side by side
  - `"triple"` — 3 images (1 large + 2 small)
  - `"bisect"` — 2-column grid, supports up to 6 images
  - `"trisect"` — 3-column grid, supports up to 9 images
- **img_list**: Required. Array of image objects:
  - `img_key`: Required image key
  - `transparent`: Optional boolean
- **corner_radius**: Optional. `"[0,...]px"` or `"[0,100]%"`
- **combination_transparent**: Optional boolean. Apply transparency to all images.
- **margin**: Component margin

## Notes

- Excess images beyond the mode's capacity are hidden (not displayed).
- Missing images leave blank slots in the grid.

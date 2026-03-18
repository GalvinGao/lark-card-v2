# Image — `img`

Tag: `img`

## JSON Example

```json
{
  "tag": "img",
  "img_key": "img_v3_xxxx",
  "alt": { "tag": "plain_text", "content": "Description" },
  "title": { "tag": "plain_text", "content": "Title" },
  "scale_type": "crop_center",
  "size": "stretch",
  "corner_radius": "5px",
  "transparent": false,
  "preview": true,
  "margin": "0px"
}
```

## Key Fields

- **img_key**: Required. Image key obtained from the Upload Image API.
- **alt**: Required. Hover/accessibility description. `{ "tag": "plain_text", "content": "..." }`
- **title**: Optional. Image title displayed below the image.
- **scale_type**: How the image is cropped/scaled:
  - `"crop_center"` — crop from center (default)
  - `"crop_top"` — crop from top
  - `"fit_horizontal"` — fit width, no cropping
- **size**: Image display size. Only applies when `scale_type` is `"crop_center"` or `"crop_top"`:
  - `"large"` | `"medium"` | `"small"` | `"tiny"` | `"stretch"`
  - Custom: `"[1,1000]px [1,1000]px"` (width height)
- **corner_radius**: `"[0,...]px"` or `"[0,100]%"`
- **transparent**: Boolean, default `false`. Renders with transparent background.
- **preview**: Boolean, default `true`. Allows click to enlarge.
- **margin**: Component margin

## Notes

- v2.0 no longer supports `stretch_without_padding`. To achieve a similar edge-to-edge effect, use negative margin instead: `"margin": "4px -12px"`.

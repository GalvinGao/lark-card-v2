# Image

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/image>

Lark cards support image components. You can call the Upload Image API or upload an image in the building tool's image component to obtain the image key, enriching the card's content.

This document introduces the JSON 2.0 structure of the image component. To view the historical JSON 1.0 structure, refer to Image.

## Notes

- In the JSON 2.0 structure, the `size` attribute of the image component no longer supports `stretch_without_padding` to achieve a full-width effect. Instead, set the `margin` attribute to a negative value:

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "img",
        "img_key": "img_v3_0238_073f1823-df2b-4377-86c6-e293f183622j",
        "scale_type": "crop_center",
        "margin": "4px -12px"
      }
    ]
  }
}
```

- To ensure clarity in the chat window, uploaded images should adhere to the following specifications:
  - Image dimensions within the range of 1500 x 3000 px.
  - Image size no more than 10 MB.
  - Image height-to-width ratio not exceeding 16:9.

## JSON Structure

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "img",
        "element_id": "custom_id",
        "margin": "0px 0px 0px 0px",
        "img_key": "img_v3_0238_073f1823-df2b-4377-86c6-e293f18abcef",
        "alt": {
          "tag": "plain_text",
          "content": ""
        },
        "title": {
          "tag": "plain_text",
          "content": ""
        },
        "corner_radius": "5px",
        "scale_type": "crop_top",
        "size": "100px 100px",
        "transparent": false,
        "preview": false,
        "mode": "large",
        "custom_width": "300px",
        "compact_width": false
      }
    ]
  }
}
```

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | Component tag. Fixed value: `img`. |
| `element_id` | No | String | Empty | Unique identifier for the component (JSON 2.0). Used to specify the component when calling related interfaces. Must be globally unique within the card. Only letters, numbers, and underscores are allowed; must start with a letter and not exceed 20 characters. |
| `margin` | No | String | `0` | Margin of the component (JSON 2.0). Value range is `[-99,99]px`. Accepts a single value (e.g. `"10px"` for all sides), double value (e.g. `"4px 0"` for top/bottom and left/right), or four values (e.g. `"4px 0 4px 0"` for top, right, bottom, left). |
| `img_key` | Yes | String | — | Key of the image resource. Obtain by calling the Upload Image API or by uploading images in the building tool. |
| `alt` | Yes | Struct | — | Description displayed when hovering over the image. Example: `{"tag": "plain_text", "content": "Hover description"}`. |
| `title` | No | Struct | — | Image title. Example: `{"tag": "plain_text", "content": "Image Title"}`. |
| `corner_radius` | No | String | — | Corner radius of the image. Values follow the format `[0,∞]px` or `[0,100]%`. |
| `scale_type` | No | String | `crop_center` | Image cropping mode, activated when the size ratio differs from the image. Values: `crop_center` (center cropping), `crop_top` (top cropping), `fit_horizontal` (display fully without cropping). |
| `size` | No | String | — | Dimensions of the image. Only effective when `scale_type` is `crop_center` or `crop_top`. Values: `large` (multiple images), `medium` (cover images), `small` (avatars), `tiny` (icons), `stretch` (aspect ratio less than 16:9), or custom `[1,1000]px [1,1000]px`. Note: JSON 2.0 no longer supports `stretch_without_padding`; use a negative `margin` instead. |
| `transparent` | No | Boolean | `false` | Whether the background is transparent. Default is `false` (white background). |
| `preview` | No | Boolean | `true` | Whether to enlarge the image on click. `true`: enlarges in a viewer. `false`: responds to the card's interaction event without enlarging. Tip: if a `card_link` redirect is configured, set this to `false` so clicking the image triggers the redirect. |

## Legacy Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `mode` | No | String | — | Image display mode. Values: `crop_center` (center cropping, limits height for long images), `fit_horizontal` (tiling mode, stretches width to display entire image), `stretch` (adaptive, stretches width; crops when height:width exceeds 16:9). Note: setting this overrides `custom_width`. |
| `custom_width` | No | int | — | Custom maximum display width of the image, adjustable within 278px to 580px. By default, the image width matches the component area. Effective in Lark V4.0 and above. |
| `compact_width` | No | Boolean | `false` | Whether to display the image in compact form. If `true`, the maximum display width is 278px. |

## Example

```json
{
  "schema": "2.0",
  "body": {
    "direction": "vertical",
    "padding": "12px 12px 12px 12px",
    "elements": [
      {
        "tag": "img",
        "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg",
        "preview": true,
        "transparent": false,
        "scale_type": "crop_center",
        "size": "stretch",
        "alt": {
          "tag": "plain_text",
          "content": "示例图片"
        },
        "corner_radius": "5%",
        "margin": "0px 0px 0px 0px",
        "element_id": "demoimg01"
      }
    ]
  }
}
```

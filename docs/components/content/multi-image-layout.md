# Multi-Image Layout

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/multi-image-laylout>

Lark cards support multi-image combination components. You can call the Upload Image API or upload images in the Lark card building tool, and enter the obtained image key into the multi-image combination component to enrich the card content.

This document introduces the JSON 2.0 structure of the multi-image layout component. To view the historical JSON 1.0 structure, refer to multi-image layout.

## Use Cases

In content delivery scenarios, you may need to organize and arrange multiple images within a card. Use the multi-image combination component to choose the image combination mode and quickly build a multi-image style.

Supported layouts: Double Image Combination, Triple Image Combination, Quad Grid, Six Grid, and Nine Grid.

## Notes

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
        "tag": "img_combination",
        "element_id": "custom_id",
        "margin": "0px 0px 0px 0px",
        "combination_mode": "double",
        "combination_transparent": false,
        "corner_radius": "12px",
        "img_list": [
          {
            "img_key": "img_v3_0239_8347760e-3173-4072-b1aa-e4e7c835741j",
            "transparent": false
          },
          {
            "img_key": "img_v3_0239_d9a9b734-57f8-4247-baf3-ae178b55f96j"
          }
        ]
      }
    ]
  }
}
```

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | Component tag. Fixed value: `img_combination`. |
| `element_id` | No | String | Empty | Unique identifier for the component (JSON 2.0). Used to specify the component when calling related interfaces. Must be globally unique within the card. Only letters, numbers, and underscores are allowed; must start with a letter and not exceed 20 characters. |
| `margin` | No | String | `0` | Margin of the component (JSON 2.0). Value range is `[-99,99]px`. Accepts a single value (e.g. `"10px"` for all sides), double value (e.g. `"4px 0"` for top/bottom and left/right), or four values (e.g. `"4px 0 4px 0"` for top, right, bottom, left). |
| `combination_mode` | Yes | String | — | Mode of the image combination. Values: `double` (up to 2 images), `triple` (up to 3 images), `bisect` (evenly divided double column, up to 6 images in 3 rows), `trisect` (evenly divided triple column, up to 9 images in 3 rows). If uploaded images exceed the limit, only the earliest images are shown. If fewer images are uploaded, the remaining slots are blank. |
| `combination_transparent` | No | Boolean | `false` | Whether the background is transparent. Default is `false` (white background). |
| `corner_radius` | No | String | — | Corner radius for the images in the combination, in pixels. Values follow the format `[0,∞]px` or `[0,100]%`. |
| `img_list` | Yes | Object[] | — | Array of image resources, in the order of image arrangement. |
| `img_list[].img_key` | Yes | String | — | Key of the image resource. Obtain by calling the Upload Image API or by uploading images in the building tool. |

## Examples

### Double Image Layout

```json
{
  "schema": "2.0",
  "body": {
    "direction": "vertical",
    "padding": "12px 12px 12px 12px",
    "elements": [
      {
        "tag": "img_combination",
        "combination_mode": "double",
        "img_list": [
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          }
        ],
        "combination_transparent": false,
        "margin": "0px 0px 0px 0px"
      }
    ]
  }
}
```

### Triple Image Layout

```json
{
  "schema": "2.0",
  "body": {
    "direction": "vertical",
    "padding": "12px 12px 12px 12px",
    "elements": [
      {
        "tag": "img_combination",
        "combination_mode": "triple",
        "img_list": [
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          }
        ],
        "combination_transparent": false,
        "margin": "0px 0px 0px 0px"
      }
    ]
  }
}
```

### Evenly Divided Double Column Layout

```json
{
  "schema": "2.0",
  "body": {
    "direction": "vertical",
    "padding": "12px 12px 12px 12px",
    "elements": [
      {
        "tag": "img_combination",
        "combination_mode": "bisect",
        "img_list": [
          { "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg" },
          { "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg" },
          { "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg" },
          { "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg" },
          { "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg" },
          { "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg" }
        ],
        "combination_transparent": false,
        "margin": "0px 0px 0px 0px"
      }
    ]
  }
}
```

### Evenly Divided Triple Column Layout

```json
{
  "schema": "2.0",
  "body": {
    "direction": "vertical",
    "padding": "12px 12px 12px 12px",
    "elements": [
      {
        "tag": "img_combination",
        "combination_mode": "trisect",
        "img_list": [
          { "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg" },
          { "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg" },
          { "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg" },
          { "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg" },
          { "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg" },
          { "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg" },
          { "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg" },
          { "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg" },
          { "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg" }
        ],
        "combination_transparent": false,
        "margin": "0px 0px 0px 0px"
      }
    ]
  }
}
```

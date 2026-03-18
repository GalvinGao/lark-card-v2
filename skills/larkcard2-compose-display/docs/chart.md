# Chart — `chart`

Tag: `chart`

## JSON Example

```json
{
  "tag": "chart",
  "chart_spec": {
    "type": "line",
    "data": [
      { "values": [{ "x": "Mon", "y": 10 }, { "x": "Tue", "y": 20 }] }
    ],
    "xField": "x",
    "yField": "y"
  },
  "aspect_ratio": "16:9",
  "color_theme": "brand",
  "preview": true,
  "height": "auto",
  "margin": "0px"
}
```

## Key Fields

- **chart_spec**: Required. A VChart JSON specification object. Supported chart types: line, area, bar, pie, doughnut, combo, funnel, scatter, radar, progress, word cloud.
- **aspect_ratio**: `"1:1"` | `"2:1"` | `"4:3"` | `"16:9"` (default)
- **color_theme**: `"brand"` | `"rainbow"` | `"complementary"` | `"converse"` | `"primary"`
- **preview**: Boolean, default `true`. Click to view enlarged chart.
- **height**: `"auto"` or `"[100,999]px"`
- **margin**: Component margin

## Constraints

- Maximum **5 charts** per card.
- No JavaScript syntax in `chart_spec` — pure JSON only.
- Charts **cannot** be nested inside other components (column_set, form, etc.). They must be placed directly in `body.elements`.

# Table — `table`

Tag: `table`

## JSON Example

```json
{
  "tag": "table",
  "page_size": 5,
  "row_height": "low",
  "header_style": {
    "text_align": "left",
    "text_size": "normal",
    "background_style": "grey",
    "text_color": "default",
    "bold": true,
    "lines": 1
  },
  "columns": [
    { "name": "col1", "display_name": "Name", "data_type": "text", "width": "auto" },
    { "name": "col2", "display_name": "Amount", "data_type": "number", "format": { "symbol": "$", "precision": 2, "separator": true } }
  ],
  "rows": [
    { "col1": "Item A", "col2": 100.50 }
  ],
  "margin": "0px"
}
```

## Key Fields

- **page_size**: Rows per page, 1-10, default `5`
- **row_height**: `"low"` | `"middle"` | `"high"` | `"auto"` | `"[32,124]px"`
- **header_style**: Header row styling:
  - `text_align`: `"left"` | `"center"` | `"right"`
  - `text_size`: text size enum
  - `background_style`: `"grey"` or color enum
  - `text_color`: `"default"` or color enum
  - `bold`: boolean
  - `lines`: max header text lines
- **columns**: Required. Up to 50 columns. Each column:
  - `name`: Column identifier (used as key in row objects)
  - `display_name`: Column header text
  - `data_type`: `"text"` | `"lark_md"` | `"options"` | `"number"` | `"persons"` | `"date"` | `"markdown"`
  - `width`: `"auto"` | `"[64,640]px"` | percentage string
  - `horizontal_align`: `"left"` | `"center"` | `"right"`
  - `vertical_align`: `"top"` | `"center"` | `"bottom"`
  - `format`: Type-specific formatting (e.g., `{ "symbol": "$", "precision": 2, "separator": true }` for numbers)
- **rows**: Array of objects. Each object is keyed by column `name` values.
- **margin**: Component margin

## Constraints

- Tables **cannot** be nested inside other components. They must be placed directly in `body.elements`.
- Maximum **5 tables** per card.

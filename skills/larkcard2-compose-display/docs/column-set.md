# Column Set — `column_set` / `column`

Tag: `column_set` (container) with `column` children

## JSON Example

```json
{
  "tag": "column_set",
  "flex_mode": "none",
  "horizontal_spacing": "8px",
  "columns": [
    {
      "tag": "column",
      "width": "weighted",
      "weight": 1,
      "vertical_align": "top",
      "vertical_spacing": "8px",
      "padding": "8px",
      "background_style": "default",
      "elements": []
    },
    {
      "tag": "column",
      "width": "weighted",
      "weight": 1,
      "elements": []
    }
  ]
}
```

## column_set Fields

- **flex_mode**: Column layout behavior:
  - `"none"` — columns use their declared widths
  - `"stretch"` — columns stretch to fill available space
  - `"flow"` — columns wrap to next line when overflowing
  - `"bisect"` — force 2 equal columns
  - `"trisect"` — force 3 equal columns
- **horizontal_spacing**: Gap between columns:
  - Named: `"small"` (4px) | `"medium"` (8px) | `"large"` (12px) | `"extra_large"` (16px)
  - Custom: `"[0,99]px"`
- **background_style**: `"default"` or color enum
- **action.multi_url**: Optional click-through link for the entire column set
- **margin**: Component margin

## column Fields

- **width**: Column width:
  - `"auto"` — fit content
  - `"weighted"` — proportional (use with `weight`)
  - `"[16,600]px"` — fixed pixel width
- **weight**: 1-5 (only when `width` is `"weighted"`)
- **vertical_align**: `"top"` | `"center"` | `"bottom"`
- **direction**: `"vertical"` (default) | `"horizontal"`
- **vertical_spacing**: Spacing between child elements vertically
- **horizontal_spacing**: Spacing between child elements horizontally
- **horizontal_align**: `"left"` | `"center"` | `"right"`
- **padding**: Inner padding (same format as body padding)
- **margin**: Column margin
- **background_style**: `"default"` or color enum
- **action.multi_url**: Optional click-through link for this column
- **elements**: Nested components (no `form` or `table` allowed inside columns)

## Constraints

- Maximum **5 nesting levels** of column_set.
- `form` and `table` components cannot be nested inside columns.

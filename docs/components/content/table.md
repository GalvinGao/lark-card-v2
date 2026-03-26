# Table

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/table>

Lark cards support table components and allow the addition of plain text, Markdown, option tags, personnel lists, date, and numeric content within the tables.

## Notes

- Each card can accommodate up to five table components. If the card is configured for multiple languages, each language can accommodate up to five table components.
- When there is insufficient space to display the full content within a cell, the content will be truncated. On the client side, users can view the truncated content by hovering over the cell or clicking on it.

### Nesting Rules

- The table component cannot be embedded within other components and can only be placed under the card's root node.
- The table component does not support embedding other components.

## JSON Structure

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "table",
        "element_id": "custom_id",
        "margin": "0px 0px 0px 0px",
        "page_size": 5,
        "row_height": "low",
        "row_max_height": "50px",
        "freeze_first_column": true,
        "header_style": {
          "text_align": "left",
          "text_size": "normal",
          "background_style": "none",
          "text_color": "grey",
          "bold": true,
          "lines": 1
        },
        "columns": [
          {
            "name": "customer_name",
            "display_name": "Customer Name",
            "width": "auto",
            "data_type": "text",
            "vertical_align": "top",
            "horizontal_align": "left"
          },
          {
            "name": "customer_link",
            "display_name": "Related Links",
            "data_type": "lark_md"
          },
          {
            "name": "customer_arr",
            "display_name": "ARR",
            "data_type": "number",
            "format": {
              "symbol": "$",
              "precision": 2,
              "separator": true
            },
            "width": "120px"
          },
          {
            "name": "customer_scale",
            "display_name": "Customer Scale",
            "data_type": "options"
          },
          {
            "name": "customer_poc",
            "display_name": "Contact Person",
            "data_type": "persons"
          },
          {
            "name": "meeting_date",
            "display_name": "Meeting Date",
            "data_type": "date",
            "date_format": "YYYY/MM/DD"
          },
          {
            "name": "company_image",
            "display_name": "Company Image",
            "data_type": "markdown"
          }
        ],
        "rows": [
          {
            "customer_name": "Lark Tech",
            "customer_scale": [{ "text": "S2", "color": "blue" }],
            "customer_arr": 168,
            "customer_poc": ["ou_user_id_1", "ou_user_id_2"],
            "customer_link": "[Lark Tech](https://example.com)"
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
| `tag` | Yes | String | — | The tag of the component. Fixed value: `table`. |
| `element_id` | No | String | — | Unique identifier for the component. New in JSON 2.0. Must be globally unique within the same card. Only letters, numbers, and underscores; must start with a letter and not exceed 20 characters. |
| `margin` | No | String | `0` | Margin of the component. New in JSON 2.0. Value range is `[-99,99]px`. Accepts a single value (`"10px"`), double value (`"4px 0"`), or four values (`"4px 0 4px 0"`) for top, right, bottom, left. |
| `page_size` | No | Number | `5` | Maximum number of data rows displayed per page. Supports integers `[1,10]`. |
| `row_height` | No | String | `low` | Row height. Values: `low`, `middle`, `high`, `auto` (adapts to content, new in JSON 2.0, Lark v7.33+), or `[32,124]px` for custom height. |
| `row_max_height` | No | String | `124px` | Maximum row height when `row_height` is `auto`. Value range is `[32,999]` pixels. New in JSON 2.0, Lark v7.33+. |
| `header_style` | No | Object | — | Header style settings. See `header_style` fields below. |
| `freeze_first_column` | No | Boolean | `false` | Whether to freeze the first column. When `true`, the first column remains fixed during horizontal scrolling. |
| `columns` | Yes | column[] | `[]` | Array of column objects. Up to 50 columns supported. See `column` fields below. |
| `rows` | Yes | JSON | `[]` | Array of row objects. Data corresponding to column definitions, defined as `"name": VALUE`. |

### `header_style` Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `text_align` | No | String | `left` | Text alignment. Values: `left`, `center`, `right`. Note: not currently supported in the card builder tool. |
| `text_size` | No | String | `normal` | Text size. Values: `normal` (14px), `heading` (16px). |
| `background_style` | No | String | `none` | Background color. Values: `grey`, `none`. |
| `text_color` | No | String | `default` | Text color. Values: `default` (black in light mode, white in dark mode), `grey`. |
| `bold` | No | Boolean | `true` | Whether header text is bold. |
| `lines` | No | Number | `1` | Number of text lines for the header. Supports integers >= 1. |

### `column` Fields

A maximum of 50 columns can be added; content beyond 50 columns will not be displayed.

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `name` | Yes | String | — | Custom column identifier. Used to match row data to the correct cell. |
| `display_name` | No | String | — | Column name displayed in the header. If empty, the column name is not displayed. |
| `width` | No | String | `auto` | Column width. Values: `auto`, custom pixels (`120px`, range `[80,600]px`), or percentage (`25%`, range `[1%,100%]`). |
| `vertical_align` | No | String | `center` | Vertical alignment. Values: `top`, `center`, `bottom`. |
| `horizontal_align` | No | String | `left` | Horizontal alignment. Values: `left`, `center`, `right`. |
| `data_type` | Yes | String | `text` | Column data type. See `data_type` values below. |
| `format` | No | Object | — | Only effective when `data_type` is `number`. Sets decimal places, currency units, and thousands separator. |
| `format.precision` | No | Int | — | Decimal places (0-10). `0` rounds to integer. |
| `format.symbol` | No | String | — | Currency unit before the number (one character, e.g. `"$"`). |
| `format.separator` | No | Boolean | `false` | Whether to apply thousands separator comma style. |
| `date_format` | No | String | — | Only effective when `data_type` is `date`. Placeholders: `YYYY` (year), `MM` (month), `DD` (day), `HH` (hour), `mm` (minute), `ss` (second). Recommended formats: `YYYY/MM/DD`, `YYYY/MM/DD HH:mm`, `YYYY-MM-DD`, `YYYY-MM-DD HH:mm`, `DD/MM/YYYY`, `MM/DD/YYYY`. Default: RFC 3339 standard. |

### `data_type` Values

| Value | Supported Version | Description | Data Structure Example |
|-------|-------------------|-------------|----------------------|
| `text` | Lark v7.4+ | Plain text without formatting (default). | `"name": "plain text"` |
| `lark_md` | Lark v7.10+ | Text supporting partial Markdown format. | `"name": "[Link](https://example.com)"` |
| `options` | Lark v7.4+ | Option tags. Supports custom tag color via `color` parameter (default: `blue`). Keep text short to avoid truncation. | Single: `"name": "option"`. Multiple: `"name": [{"text": "S1", "color": "red"}, {"text": "S2", "color": "green"}]` |
| `number` | Lark v7.4+ | Numbers, displayed right-aligned by default. Use the `format` field to configure number formatting. | `"name": 26.58` |
| `persons` | Lark v7.4+ | List of people (user name + avatar). Specify by user_id, open_id, union_id, or lark_id. Invalid IDs display as "Unknown User". | Multiple: `"name": ["user_id_1", "user_id_2"]`. Single: `"name": "user_id"` |
| `date` | Lark v7.6+ | Date and time. Input as Unix millisecond timestamp. Displayed in user's local time zone. Use `date_format` field to configure format. | `"name": 1606101072000` |
| `markdown` | Lark v7.14+ | Text supporting full Markdown syntax. | `"name": "![image.png](img_key)"` |

## Example

```json
{
  "schema": "2.0",
  "header": {
    "template": "blue",
    "title": {
      "content": "Table Component Example",
      "tag": "plain_text"
    }
  },
  "body": {
    "elements": [
      {
        "tag": "table",
        "page_size": 5,
        "row_height": "auto",
        "header_style": {
          "text_align": "left",
          "text_size": "normal",
          "background_style": "none",
          "text_color": "grey",
          "bold": true,
          "lines": 1
        },
        "columns": [
          {
            "name": "customer_name",
            "display_name": "Customer Name",
            "data_type": "text",
            "horizontal_align": "left",
            "vertical_align": "top",
            "width": "auto"
          },
          {
            "name": "customer_scale",
            "display_name": "Customer Scale",
            "data_type": "options",
            "horizontal_align": "left",
            "vertical_align": "top",
            "width": "auto"
          },
          {
            "name": "customer_arr",
            "display_name": "ARR",
            "data_type": "number",
            "format": {
              "symbol": "$",
              "precision": 2,
              "separator": true
            },
            "width": "auto"
          },
          {
            "name": "customer_poc",
            "display_name": "Contact",
            "data_type": "persons",
            "horizontal_align": "left",
            "vertical_align": "top",
            "width": "auto"
          },
          {
            "name": "customer_date",
            "display_name": "Sign Date",
            "data_type": "date",
            "date_format": "YYYY/MM/DD",
            "width": "auto"
          },
          {
            "name": "customer_link",
            "display_name": "Related Links",
            "data_type": "lark_md",
            "width": "auto"
          },
          {
            "name": "company_image",
            "display_name": "Company Image",
            "data_type": "markdown"
          }
        ],
        "rows": [
          {
            "customer_name": "Lark Tech",
            "customer_date": 1699341315000,
            "customer_scale": [{ "text": "S2", "color": "blue" }],
            "customer_arr": 168,
            "customer_poc": ["ou_14a32f1a02e64944cf19207aa43abcef", "ou_e393cf9c22e6e617a4332210d2aabcef"],
            "customer_link": "[Lark Tech](https://example.com)",
            "company_image": "![image.png](img_v3_02cc_bf88cdee-6650-4b39-987c-f8e87c3227fg)"
          },
          {
            "customer_name": "Lark Tech 01",
            "customer_date": 1606101072000,
            "customer_scale": [{ "text": "S1", "color": "red" }],
            "customer_arr": 168.23,
            "customer_poc": "ou_14a32f1a02e64944cf19207aa43abcef",
            "customer_link": "[Lark Tech 01](https://example.com)",
            "company_image": "![image.png](img_v3_02cc_bf88cdee-6650-4b39-987c-f8e87c3227fg)"
          },
          {
            "customer_name": "Lark Tech 02",
            "customer_date": 1606101072000,
            "customer_scale": [{ "text": "S3", "color": "orange" }],
            "customer_arr": 168.23,
            "customer_poc": "ou_14a32f1a02e64944cf19207aa43abcef",
            "customer_link": "[Lark Tech 02](https://example.com)",
            "company_image": "![image.png](img_v3_02cc_bf88cdee-6650-4b39-987c-f8e87c3227fg)"
          },
          {
            "customer_name": "Lark Tech 03",
            "customer_date": 1606101072000,
            "customer_scale": [{ "text": "S2", "color": "blue" }],
            "customer_arr": 168.23,
            "customer_poc": "ou_14a32f1a02e64944cf19207aa43abcef",
            "customer_link": "[Lark Tech 03](https://example.com)",
            "company_image": "![image.png](img_v3_02cc_bf88cdee-6650-4b39-987c-f8e87c3227fg)"
          }
        ]
      }
    ]
  }
}
```

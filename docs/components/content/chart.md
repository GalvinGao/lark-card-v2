# Chart

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/chart>

Lark cards provide chart components based on VChart definitions, supporting various data visualization methods such as line charts, area charts, bar charts, pie charts, and word clouds. These help you visualize various types of information and enhance communication efficiency.

## Notes

- It is recommended to place a maximum of five chart components on a single card.
- The chart component does not support JavaScript syntax.
- The following VChart-related properties are not supported on mobile devices. If these properties are specified in the chart component, the chart will fail to load on mobile:
  - Texture property (`barChart.bar.style.texture`)
  - Conical gradient property, i.e., when `gradient` is set to `conical`
  - Shape word cloud based on grid pixel layout, i.e., when `wordCloudChart.wordCloudConfig.layoutMode` is set to `grid`
  - Repeat property of extensionMark images, i.e., `extensionMark-image.style.repeatX` or `extensionMark-image.style.repeatY`
  - Bar background (`barChart.bar.style.background`) does not support SVG
- To enhance the display effect of charts on different terminals, the platform automatically appends media query configurations to the `chart_spec` you provide. To control the adaptive display logic yourself, set `"media":[]` in the `chart_spec` to disable the default configurations.

### Nesting Rules

The chart component can be used under the card's root node or nested within column containers and accordion panels.

### Features

- **Interactive charts:** Users can display data labels by clicking on the chart, filter data by clicking on legends, and select data by dragging the thumbnail axis.
- **Adaptive styling:** Supports the presentation of various chart styles and adapts well across different devices and color modes.
- **Enlargement:** On PC, charts support viewing in a separate window; on mobile, charts support full-screen viewing upon clicking.

## JSON Structure

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "element_id": "custom_id",
        "margin": "0px 0px 0px 0px",
        "aspect_ratio": "16:9",
        "color_theme": "brand",
        "chart_spec": {},
        "preview": false,
        "height": "auto"
      }
    ]
  }
}
```

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | The tag of the component. Fixed value: `chart`. |
| `element_id` | No | String | — | Unique identifier for the component. New in JSON 2.0. Must be globally unique within the same card. Only letters, numbers, and underscores are allowed; must start with a letter and not exceed 20 characters. |
| `margin` | No | String | `0` | Margin of the component. New in JSON 2.0. Value range is `[-99,99]px`. Accepts a single value (`"10px"`), double value (`"4px 0"`), or four values (`"4px 0 4px 0"`) for top, right, bottom, left. |
| `aspect_ratio` | No | String | PC: `16:9`, Mobile: `1:1` | Aspect ratio of the chart. Supported ratios: `1:1`, `2:1`, `4:3`, `16:9`. |
| `color_theme` | No | String | `brand` | Theme style of the chart. Adjusts color style when multiple colors exist. Ineffective if style attributes are declared in `chart_spec`. Values: `brand` (default, matches Lark theme), `rainbow` (monochromatic rainbow), `complementary`, `converse` (contrast colors), `primary`. |
| `chart_spec` | Yes | VChart spec | — | Chart definition based on VChart. See [VChart Official Documentation](https://www.visactor.io/vchart). VChart version support: v7.1-7.6 supports 1.2.2; v7.7-7.9 supports 1.6.6; v7.10-7.15 supports 1.8.3; v7.16-7.26 supports 1.10.1; v7.27+ supports 1.12.3. |
| `preview` | No | Boolean | `true` | Whether the chart can be viewed in a separate window. `true`: on PC opens in a new Lark window, on mobile opens full-screen. `false`: no separate viewing. |
| `height` | No | String | `auto` | Height of the chart component. `auto`: calculated from `aspect_ratio`. `[1,999]px`: custom fixed height (makes `aspect_ratio` ineffective). Only effective in Lark 7.10+. |

## Chart Types and Examples

The chart component is based on VChart and currently supports 13 types of charts. This section lists JSON 2.0 structure examples for each type. For detailed property descriptions, refer to the VChart configuration documentation.

### Line Chart

Line charts are generally used to display trends in data over time.

**JSON Template:**

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "line",
          "title": {
            "text": "Line Chart"
          },
          "data": {
            "values": "mock_data"
          },
          "xField": "time",
          "yField": "value"
        }
      }
    ]
  },
  "header": {
    "template": "purple",
    "title": {
      "content": "Card Title",
      "tag": "plain_text"
    }
  }
}
```

**Mock Data:**

```json
[
  { "time": "2:00", "value": 8 },
  { "time": "4:00", "value": 9 },
  { "time": "6:00", "value": 11 },
  { "time": "8:00", "value": 14 },
  { "time": "10:00", "value": 16 },
  { "time": "12:00", "value": 17 },
  { "time": "14:00", "value": 17 },
  { "time": "16:00", "value": 16 },
  { "time": "18:00", "value": 15 }
]
```

### Area Chart

Area charts are similar to line charts and can be used to display trends in data over time. The filled area below the line emphasizes the overall cumulative trend.

**JSON Template:**

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "area",
          "title": {
            "text": "Area Chart"
          },
          "data": {
            "values": "mock_data"
          },
          "xField": "time",
          "yField": "value"
        }
      }
    ]
  },
  "header": {
    "template": "purple",
    "title": {
      "content": "Card Title",
      "tag": "plain_text"
    }
  }
}
```

**Mock Data:**

```json
[
  { "time": "2:00", "value": 8 },
  { "time": "4:00", "value": 9 },
  { "time": "6:00", "value": 11 },
  { "time": "8:00", "value": 14 },
  { "time": "10:00", "value": 16 },
  { "time": "12:00", "value": 17 },
  { "time": "14:00", "value": 17 },
  { "time": "16:00", "value": 16 },
  { "time": "18:00", "value": 15 }
]
```

### Bar Chart

Bar charts are often used to compare data between different groups or categories, clearly showing the differences among them.

**JSON Template:**

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "bar",
          "title": {
            "text": "Bar Chart"
          },
          "data": {
            "values": "mock_data"
          },
          "xField": ["year", "type"],
          "yField": "value",
          "seriesField": "type",
          "legends": {
            "visible": true,
            "orient": "bottom"
          }
        }
      }
    ]
  },
  "header": {
    "template": "purple",
    "title": {
      "content": "Card Title",
      "tag": "plain_text"
    }
  }
}
```

**Mock Data:**

```json
[
  { "type": "Autoc", "year": "1930", "value": 129 },
  { "type": "Autoc", "year": "1940", "value": 133 },
  { "type": "Autoc", "year": "1950", "value": 130 },
  { "type": "Autoc", "year": "1960", "value": 126 },
  { "type": "Autoc", "year": "1970", "value": 117 },
  { "type": "Autoc", "year": "1980", "value": 114 },
  { "type": "Democ", "year": "1930", "value": 22 },
  { "type": "Democ", "year": "1940", "value": 13 },
  { "type": "Democ", "year": "1950", "value": 25 },
  { "type": "Democ", "year": "1960", "value": 29 },
  { "type": "Democ", "year": "1970", "value": 38 },
  { "type": "Democ", "year": "1980", "value": 41 }
]
```

### Bar Chart (Horizontal)

Horizontal bar charts are similar to column charts but displayed horizontally. They are commonly used to compare data across different categories and are easier to read when data labels are long or there are many categories.

**JSON Template:**

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "bar",
          "title": {
            "text": "Horizontal Bar Chart"
          },
          "data": {
            "values": "mock_data"
          },
          "direction": "horizontal",
          "xField": "value",
          "yField": "name"
        }
      }
    ]
  },
  "header": {
    "template": "purple",
    "title": {
      "content": "Card Title",
      "tag": "plain_text"
    }
  }
}
```

**Mock Data:**

```json
[
  { "name": "Apple", "value": 214480 },
  { "name": "Google", "value": 155506 },
  { "name": "Amazon", "value": 100764 },
  { "name": "Microsoft", "value": 92715 },
  { "name": "Coca-Cola", "value": 66341 },
  { "name": "Samsung", "value": 59890 },
  { "name": "Toyota", "value": 53404 },
  { "name": "Mercedes-Benz", "value": 48601 }
]
```

### Doughnut Chart

Doughnut charts represent the relative proportions of parts within a whole. They are suitable for displaying percentage distributions and emphasizing the structure of the whole.

**JSON Template:**

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "pie",
          "title": {
            "text": "Doughnut Chart"
          },
          "data": {
            "values": "mock_data"
          },
          "valueField": "value",
          "categoryField": "type",
          "outerRadius": 0.9,
          "innerRadius": 0.3,
          "label": {
            "visible": true
          },
          "legends": {
            "visible": true
          }
        }
      }
    ]
  },
  "header": {
    "template": "purple",
    "title": {
      "content": "Card Title",
      "tag": "plain_text"
    }
  }
}
```

**Mock Data:**

```json
[
  { "type": "oxygen", "value": "46.60" },
  { "type": "silicon", "value": "27.72" },
  { "type": "aluminum", "value": "8.13" },
  { "type": "iron", "value": "5" },
  { "type": "calcium", "value": "3.63" },
  { "type": "potassium", "value": "2.59" },
  { "type": "others", "value": "3.5" }
]
```

### Pie Chart

Pie charts represent the relative proportions of parts within a whole, but are usually suitable for displaying data for a few segments. They are applicable for presenting percentages or shares.

**JSON Template:**

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "aspect_ratio": "4:3",
        "chart_spec": {
          "type": "pie",
          "title": {
            "text": "Pie Chart"
          },
          "data": {
            "values": "mock_data"
          },
          "valueField": "value",
          "categoryField": "type",
          "outerRadius": 0.9,
          "legends": {
            "visible": true,
            "orient": "right"
          },
          "padding": {
            "left": 10,
            "top": 10,
            "bottom": 5,
            "right": 0
          },
          "label": {
            "visible": true
          }
        }
      }
    ]
  }
}
```

**Mock Data:**

```json
[
  { "type": "S1", "value": "340" },
  { "type": "S2", "value": "170" },
  { "type": "S3", "value": "150" },
  { "type": "S4", "value": "120" },
  { "type": "S5", "value": "100" }
]
```

### Combo Chart

Combo charts combine multiple chart types to present different kinds of data simultaneously. For example, a combination of a line chart and a bar chart can display both trends and totals at the same time.

**JSON Template:**

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "common",
          "title": {
            "text": "Combo Chart"
          },
          "data": [
            {
              "values": "mock_data_1"
            },
            {
              "values": "mock_data_2"
            }
          ],
          "series": [
            {
              "type": "bar",
              "dataIndex": 0,
              "label": { "visible": true },
              "seriesField": "type",
              "xField": ["x", "type"],
              "yField": "y"
            },
            {
              "type": "line",
              "dataIndex": 1,
              "label": { "visible": true },
              "seriesField": "type",
              "xField": "x",
              "yField": "y"
            }
          ],
          "axes": [
            { "orient": "bottom" },
            { "orient": "left" }
          ],
          "legends": {
            "visible": true,
            "orient": "bottom"
          }
        }
      }
    ]
  },
  "header": {
    "template": "purple",
    "title": {
      "content": "Card Title",
      "tag": "plain_text"
    }
  }
}
```

**Mock Data 1 (bar series):**

```json
[
  { "x": "Mon", "type": "Breakfast", "y": 15 },
  { "x": "Mon", "type": "Lunch", "y": 25 },
  { "x": "Tue", "type": "Breakfast", "y": 12 },
  { "x": "Tue", "type": "Lunch", "y": 30 },
  { "x": "Wed", "type": "Breakfast", "y": 15 },
  { "x": "Wed", "type": "Lunch", "y": 24 },
  { "x": "Thu", "type": "Breakfast", "y": 10 },
  { "x": "Thu", "type": "Lunch", "y": 25 },
  { "x": "Fri", "type": "Breakfast", "y": 13 },
  { "x": "Fri", "type": "Lunch", "y": 20 },
  { "x": "Sat", "type": "Breakfast", "y": 10 },
  { "x": "Sat", "type": "Lunch", "y": 22 },
  { "x": "Sun", "type": "Breakfast", "y": 12 },
  { "x": "Sun", "type": "Lunch", "y": 19 }
]
```

**Mock Data 2 (line series):**

```json
[
  { "x": "Mon", "type": "Drinks", "y": 22 },
  { "x": "Tue", "type": "Drinks", "y": 43 },
  { "x": "Wed", "type": "Drinks", "y": 33 },
  { "x": "Thu", "type": "Drinks", "y": 22 },
  { "x": "Fri", "type": "Drinks", "y": 10 },
  { "x": "Sat", "type": "Drinks", "y": 30 },
  { "x": "Sun", "type": "Drinks", "y": 50 }
]
```

### Funnel Chart

Funnel charts represent the reduction in data through a series of steps or stages. They are suitable for presenting conversion rates and depicting sales funnels.

**JSON Template:**

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "funnel",
          "title": {
            "text": "Funnel Chart"
          },
          "data": {
            "values": "mock_data"
          },
          "categoryField": "name",
          "valueField": "value",
          "isTransform": true,
          "label": { "visible": true },
          "transformLabel": { "visible": true },
          "outerLabel": { "visible": false }
        }
      }
    ]
  },
  "header": {
    "template": "purple",
    "title": {
      "content": "Card Title",
      "tag": "plain_text"
    }
  }
}
```

**Mock Data:**

```json
[
  { "value": 5676, "name": "Sent" },
  { "value": 3872, "name": "Viewed" },
  { "value": 1668, "name": "Clicked" },
  { "value": 565, "name": "Purchased" }
]
```

### Scatter Chart

Scatter charts display the relationship between two variables, showing correlation, trends, or outliers.

**JSON Template:**

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "scatter",
          "title": {
            "text": "Scatter Chart"
          },
          "data": {
            "values": "mock_data"
          },
          "xField": "milesPerGallon",
          "yField": "horsepower",
          "axes": [
            {
              "title": { "visible": true, "text": "Horse Power" },
              "orient": "left",
              "range": { "min": 0 },
              "type": "linear"
            },
            {
              "title": { "visible": true, "text": "Miles Per Gallon" },
              "orient": "bottom",
              "range": { "min": 10 },
              "type": "linear"
            }
          ]
        }
      }
    ]
  },
  "header": {
    "template": "purple",
    "title": {
      "content": "Card Title",
      "tag": "plain_text"
    }
  }
}
```

**Mock Data:**

```json
[
  { "name": "chevrolet woody", "milesPerGallon": 24.5, "cylinders": 4, "horsepower": 60 },
  { "name": "vw rabbit", "milesPerGallon": 29, "cylinders": 4, "horsepower": 70 },
  { "name": "honda civic", "milesPerGallon": 33, "cylinders": 4, "horsepower": 53 },
  { "name": "dodge aspen se", "milesPerGallon": 20, "cylinders": 6, "horsepower": 100 },
  { "name": "buick opel isuzu deluxe", "milesPerGallon": 30, "cylinders": 4, "horsepower": 80 },
  { "name": "renault 5 gtl", "milesPerGallon": 36, "cylinders": 4, "horsepower": 58 },
  { "name": "plymouth arrow gs", "milesPerGallon": 25.5, "cylinders": 4, "horsepower": 96 },
  { "name": "datsun f-10 hatchback", "milesPerGallon": 33.5, "cylinders": 4, "horsepower": 70 },
  { "name": "chevrolet caprice classic", "milesPerGallon": 17.5, "cylinders": 8, "horsepower": 145 },
  { "name": "oldsmobile cutlass supreme", "milesPerGallon": 17, "cylinders": 8, "horsepower": 110 },
  { "name": "dodge monaco brougham", "milesPerGallon": 15.5, "cylinders": 8, "horsepower": 145 },
  { "name": "mercury cougar brougham", "milesPerGallon": 15, "cylinders": 8, "horsepower": 130 },
  { "name": "chevrolet concours", "milesPerGallon": 17.5, "cylinders": 6, "horsepower": 110 },
  { "name": "buick skylark", "milesPerGallon": 20.5, "cylinders": 6, "horsepower": 105 },
  { "name": "plymouth volare custom", "milesPerGallon": 19, "cylinders": 6, "horsepower": 100 },
  { "name": "ford granada", "milesPerGallon": 18.5, "cylinders": 6, "horsepower": 98 },
  { "name": "pontiac grand prix lj", "milesPerGallon": 16, "cylinders": 8, "horsepower": 180 },
  { "name": "chevrolet monte carlo landau", "milesPerGallon": 15.5, "cylinders": 8, "horsepower": 170 },
  { "name": "chrysler cordoba", "milesPerGallon": 15.5, "cylinders": 8, "horsepower": 190 },
  { "name": "ford thunderbird", "milesPerGallon": 16, "cylinders": 8, "horsepower": 149 },
  { "name": "volkswagen rabbit custom", "milesPerGallon": 29, "cylinders": 4, "horsepower": 78 },
  { "name": "pontiac sunbird coupe", "milesPerGallon": 24.5, "cylinders": 4, "horsepower": 88 },
  { "name": "toyota corolla liftback", "milesPerGallon": 26, "cylinders": 4, "horsepower": 75 },
  { "name": "ford mustang ii 2+2", "milesPerGallon": 25.5, "cylinders": 4, "horsepower": 89 },
  { "name": "saab 99gle", "milesPerGallon": 21.6, "cylinders": 4, "horsepower": 115 },
  { "name": "ford country squire (sw)", "milesPerGallon": 15.5, "cylinders": 8, "horsepower": 142 },
  { "name": "chevrolet malibu classic (sw)", "milesPerGallon": 19.2, "cylinders": 8, "horsepower": 125 },
  { "name": "chrysler lebaron town @ country (sw)", "milesPerGallon": 18.5, "cylinders": 8, "horsepower": 150 },
  { "name": "vw rabbit custom", "milesPerGallon": 31.9, "cylinders": 4, "horsepower": 71 },
  { "name": "maxda glc deluxe", "milesPerGallon": 34.1, "cylinders": 4, "horsepower": 65 },
  { "name": "dodge colt hatchback custom", "milesPerGallon": 35.7, "cylinders": 4, "horsepower": 80 },
  { "name": "amc spirit dl", "milesPerGallon": 27.4, "cylinders": 4, "horsepower": 80 },
  { "name": "mercedes benz 300d", "milesPerGallon": 25.4, "cylinders": 5, "horsepower": 77 },
  { "name": "cadillac eldorado", "milesPerGallon": 23, "cylinders": 8, "horsepower": 125 },
  { "name": "peugeot 504", "milesPerGallon": 27.2, "cylinders": 4, "horsepower": 71 },
  { "name": "oldsmobile cutlass salon brougham", "milesPerGallon": 23.9, "cylinders": 8, "horsepower": 90 },
  { "name": "plymouth horizon", "milesPerGallon": 34.2, "cylinders": 4, "horsepower": 70 },
  { "name": "plymouth horizon tc3", "milesPerGallon": 34.5, "cylinders": 4, "horsepower": 70 },
  { "name": "datsun 210", "milesPerGallon": 31.8, "cylinders": 4, "horsepower": 65 },
  { "name": "fiat strada custom", "milesPerGallon": 37.3, "cylinders": 4, "horsepower": 69 },
  { "name": "buick skylark limited", "milesPerGallon": 28.4, "cylinders": 4, "horsepower": 90 },
  { "name": "chevrolet citation", "milesPerGallon": 28.8, "cylinders": 6, "horsepower": 115 },
  { "name": "oldsmobile omega brougham", "milesPerGallon": 26.8, "cylinders": 6, "horsepower": 115 },
  { "name": "pontiac phoenix", "milesPerGallon": 33.5, "cylinders": 4, "horsepower": 90 },
  { "name": "vw rabbit", "milesPerGallon": 41.5, "cylinders": 4, "horsepower": 76 },
  { "name": "toyota corolla tercel", "milesPerGallon": 38.1, "cylinders": 4, "horsepower": 60 },
  { "name": "chevrolet chevette", "milesPerGallon": 32.1, "cylinders": 4, "horsepower": 70 },
  { "name": "datsun 310", "milesPerGallon": 37.2, "cylinders": 4, "horsepower": 65 },
  { "name": "chevrolet citation", "milesPerGallon": 28, "cylinders": 4, "horsepower": 90 },
  { "name": "ford fairmont", "milesPerGallon": 26.4, "cylinders": 4, "horsepower": 88 },
  { "name": "amc concord", "milesPerGallon": 24.3, "cylinders": 4, "horsepower": 90 },
  { "name": "dodge aspen", "milesPerGallon": 19.1, "cylinders": 6, "horsepower": 90 },
  { "name": "audi 4000", "milesPerGallon": 34.3, "cylinders": 4, "horsepower": 78 },
  { "name": "toyota corona liftback", "milesPerGallon": 29.8, "cylinders": 4, "horsepower": 90 },
  { "name": "mazda 626", "milesPerGallon": 31.3, "cylinders": 4, "horsepower": 75 },
  { "name": "datsun 510 hatchback", "milesPerGallon": 37, "cylinders": 4, "horsepower": 92 },
  { "name": "toyota corolla", "milesPerGallon": 32.2, "cylinders": 4, "horsepower": 75 },
  { "name": "mazda glc", "milesPerGallon": 46.6, "cylinders": 4, "horsepower": 65 },
  { "name": "dodge colt", "milesPerGallon": 27.9, "cylinders": 4, "horsepower": 105 },
  { "name": "datsun 210", "milesPerGallon": 40.8, "cylinders": 4, "horsepower": 65 },
  { "name": "vw rabbit c (diesel)", "milesPerGallon": 44.3, "cylinders": 4, "horsepower": 48 },
  { "name": "vw dasher (diesel)", "milesPerGallon": 43.4, "cylinders": 4, "horsepower": 48 },
  { "name": "audi 5000s (diesel)", "milesPerGallon": 36.4, "cylinders": 5, "horsepower": 67 },
  { "name": "mercedes-benz 240d", "milesPerGallon": 30, "cylinders": 4, "horsepower": 67 },
  { "name": "honda civic 1500 gl", "milesPerGallon": 44.6, "cylinders": 4, "horsepower": 67 },
  { "name": "renault lecar deluxe", "milesPerGallon": 40.9, "cylinders": 4, "horsepower": 0 },
  { "name": "subaru dl", "milesPerGallon": 33.8, "cylinders": 4, "horsepower": 67 },
  { "name": "vokswagen rabbit", "milesPerGallon": 29.8, "cylinders": 4, "horsepower": 62 },
  { "name": "datsun 280-zx", "milesPerGallon": 32.7, "cylinders": 6, "horsepower": 132 },
  { "name": "mazda rx-7 gs", "milesPerGallon": 23.7, "cylinders": 3, "horsepower": 100 },
  { "name": "triumph tr7 coupe", "milesPerGallon": 35, "cylinders": 4, "horsepower": 88 },
  { "name": "ford mustang cobra", "milesPerGallon": 23.6, "cylinders": 4, "horsepower": 0 },
  { "name": "honda Accelerationord", "milesPerGallon": 32.4, "cylinders": 4, "horsepower": 72 },
  { "name": "plymouth reliant", "milesPerGallon": 27.2, "cylinders": 4, "horsepower": 84 },
  { "name": "buick skylark", "milesPerGallon": 26.6, "cylinders": 4, "horsepower": 84 },
  { "name": "dodge aries wagon (sw)", "milesPerGallon": 25.8, "cylinders": 4, "horsepower": 92 },
  { "name": "chevrolet citation", "milesPerGallon": 23.5, "cylinders": 6, "horsepower": 110 },
  { "name": "plymouth reliant", "milesPerGallon": 30, "cylinders": 4, "horsepower": 84 }
]
```

### Radar Chart

Radar charts compare multiple variables across different dimensions and show the relative relationships among multiple indicators.

**JSON Template:**

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "radar",
          "title": {
            "text": "Radar Chart"
          },
          "data": {
            "values": "mock_data"
          },
          "categoryField": "key",
          "valueField": "value",
          "area": { "visible": true },
          "outerRadius": 0.8,
          "axes": [
            {
              "orient": "radius",
              "label": {
                "visible": true,
                "style": { "textAlign": "center" }
              }
            }
          ]
        }
      }
    ]
  },
  "header": {
    "template": "purple",
    "title": {
      "content": "Card Title",
      "tag": "plain_text"
    }
  }
}
```

**Mock Data:**

```json
[
  { "key": "Strength", "value": 5 },
  { "key": "Speed", "value": 5 },
  { "key": "Range", "value": 3 },
  { "key": "Durability", "value": 5 },
  { "key": "Precision", "value": 5 },
  { "key": "Potential", "value": 5 }
]
```

### Bar Progress

Bar progress charts display the progress of one or more indicators, such as task completion or goal achievement.

**JSON Template:**

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "aspect_ratio": "2:1",
        "chart_spec": {
          "type": "linearProgress",
          "title": {
            "text": "Bar Progress"
          },
          "data": {
            "values": "mock_data"
          },
          "direction": "horizontal",
          "xField": "value",
          "yField": "type",
          "seriesField": "type",
          "axes": [
            {
              "orient": "left",
              "domainLine": { "visible": false }
            }
          ]
        }
      }
    ]
  },
  "header": {
    "template": "purple",
    "title": {
      "content": "Card Title",
      "tag": "plain_text"
    }
  }
}
```

**Mock Data:**

```json
[
  { "type": "Tradition Industries", "value": 0.795, "text": "79.5%" },
  { "type": "Business Companies", "value": 0.25, "text": "25%" }
]
```

### Circular Progress

Circular progress charts, similar to bar progress charts, are presented in a circular format. They emphasize overall progress and highlight the completion of specific segments.

**JSON Template:**

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "circularProgress",
          "title": {
            "text": "Circular Progress"
          },
          "data": {
            "values": "mock_data"
          },
          "valueField": "value",
          "categoryField": "type",
          "seriesField": "type",
          "radius": 0.7,
          "innerRadius": 0.4,
          "cornerRadius": 20,
          "progress": {
            "style": { "innerPadding": 5, "outerPadding": 5 }
          },
          "indicator": {
            "visible": true,
            "trigger": "hover",
            "title": {
              "visible": true,
              "field": "type",
              "autoLimit": true
            },
            "content": [
              { "visible": true, "field": "text" }
            ]
          },
          "legends": {
            "visible": true,
            "orient": "bottom",
            "title": { "visible": false }
          }
        }
      }
    ]
  },
  "header": {
    "template": "purple",
    "title": {
      "content": "Card Title",
      "tag": "plain_text"
    }
  }
}
```

**Mock Data:**

```json
[
  { "type": "Industries", "value": 0.795, "text": "79.5%" },
  { "type": "Companies", "value": 0.25, "text": "25%" }
]
```

### Word Cloud

Word clouds display the relative frequency of terms in text data. They can be used to show the importance of keywords or themes.

**JSON Template:**

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "wordCloud",
          "title": {
            "text": "Word Cloud"
          },
          "data": {
            "values": "mock_data"
          },
          "nameField": "challenge_name",
          "valueField": "sum_count",
          "seriesField": "challenge_name"
        }
      }
    ]
  },
  "header": {
    "template": "purple",
    "title": {
      "content": "Card Title",
      "tag": "plain_text"
    }
  }
}
```

**Mock Data:**

```json
[
  { "challenge_id": 1658490688121879, "challenge_name": "Topic A", "sum_count": 128 },
  { "challenge_id": 1640007327696910, "challenge_name": "Topic B", "sum_count": 103 },
  { "challenge_id": 1557656100811777, "challenge_name": "Topic C", "sum_count": 76 },
  { "challenge_id": 1553513807372289, "challenge_name": "Topic D", "sum_count": 70 },
  { "challenge_id": 1599321527572563, "challenge_name": "Topic E", "sum_count": 69 },
  { "challenge_id": 1588489879306259, "challenge_name": "Topic F", "sum_count": 54 },
  { "challenge_id": 1558589039423489, "challenge_name": "Topic G", "sum_count": 52 },
  { "challenge_id": 1565489422066689, "challenge_name": "Topic H", "sum_count": 36 },
  { "challenge_id": 1572618705886286, "challenge_name": "Topic I", "sum_count": 34 },
  { "challenge_id": 1626948076237836, "challenge_name": "Topic J", "sum_count": 32 }
]
```

# Last updated on 2025-06-27

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/chart>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsDisplay componentsChart
Chart
Copy Page
Last updated on 2025-06-27
The contents of this article
Precautions
Nesting rules
Functional features
Component properties
JSON structure
Field descriptions
Chart types and examples
Line chart
Area chart
Bar chart
Bar chart (horizontal)
Doughnut chart
Pie chart
Combo chart
Funnel chart
Scatter chart
Radar chart
Bar progress
Circular progress
Word cloud

Lark cards provide chart components based on VChart definitions, supporting various data visualization methods such as line charts, area charts, bar charts, pie charts, and word clouds. These help you visualize various types of information and enhance communication efficiency.

This document introduces the JSON 2.0 structure of the chart component. To view the historical JSON 1.0 structure, refer to Chart.

Precautions

It is recommended to place a maximum of five chart components on a single card.

The chart component does not support JavaScript syntax.

The following VChart-related properties are not supported on mobile devices. If these properties are specified in the chart component, the chart will fail to load on mobile:

Texture property (barChart.bar.style.texture)
Conical gradient property, i.e., when gradient is set to conical
Shape word cloud based on grid pixel layout, i.e., when wordCloudChart.wordCloudConfig.layoutMode is set to grid
Repeat property of extensionMark images, i.e., extensionMark-image.style.repeatX or extensionMark-image.style.repeatY
Bar background (barChart.bar.style.background) does not support SVG

To enhance the display effect of charts on different terminals and optimize the user experience, the platform automatically appends media query configurations to the chart definition (chart_spec) you provide by default. If you wish to control the adaptive display logic of the chart yourself, you can set "media":[] in the chart definition (chart_spec) to disable the default appended configurations.

Nesting rules

The chart component can be used under the card's root node or nested within column containers and accordion panels.

Functional features

Charts drawn using the chart component support the following features:

Interactive charts: Users can display data labels by clicking on the chart, filter data by clicking on legends, and select data by dragging the thumbnail axis.
Adaptive styling: Supports the presentation of various chart styles, and adapts well across different devices and color modes;
Support for enlargement: On PC, charts support viewing in a separate window; on mobile devices, charts support full-screen viewing upon clicking.
Component properties
JSON structure

The complete JSON 2.0 structure of the chart component is as follows:

{
  "schema": "2.0", // Version of the card JSON structure. The default is 1.0. To use the JSON 2.0 structure, 2.0 must be explicitly declared.
  "body": {
    "elements": [
      // Attributes supported by Lark client version 7.1 and later
      {
        "tag": "chart", // Tag of the component.
        "element_id": "custom_id", // Unique identifier for the operation component. New attribute in JSON 2.0. Used to specify the component when calling related interfaces. Customizable by developers.
        "margin": "0px 0px 0px 0px", // Margin of the component. New attribute in JSON 2.0. Default value is "0", supported range is [-99,99]px.
        "aspect_ratio": "16:9", // Aspect ratio of the chart.
        "color_theme": "brand", // Theme of the chart. Default value is brand.
        "chart_spec": {}, // Chart definition based on VChart. For detailed usage, refer to the VChart official documentation.
        "preview": false, // Whether to support viewing in an independent window. Default value is true.
        "height": "auto" // Height of the chart component. Default value is auto, which means it is automatically calculated based on the aspect ratio.
      }
    ]
  }
}
Field descriptions

The field descriptions for the chart component are as follows.

Field Name	Required	Type	Default Value	Description


tag

	

Yes

	

String

	

None

	

The tag of the component, the fixed value for the chart component is chart.




element_id

	

No

	

String

	

Empty

	

Unique identifier for the operation component. New attribute in JSON 2.0. Used to specify the component when calling related component interfaces. This value must be globally unique within the same card. Only letters, numbers, and underscores are allowed, must start with a letter, and must not exceed 20 characters.




margin

	

No

	

String

	

0

	

Margin of the component. New attribute in JSON 2.0. The value range is [-99,99]px. Optional values:

Single value, such as "10px", representing a margin of 10 px on all four sides of the component.
Double value, such as "4px 0", representing a margin of 4 px on the top and bottom, and 0 px on the left and right. Separated by space (unit can be omitted when the margin is 0).
Multiple values, such as "4px 0 4px 0", representing a margin of 4px, 12px, 4px, and 12px on the top, right, bottom, and left respectively. Separated by space.



aspect_ratio

	

No

	

String

	
PC: 16:9
Mobile: 1:1
	

The aspect ratio of the chart. Supports the following ratios:

1:1
2:1
4:3
16:9



color_theme

	

No

	

String

	

brand

	

The theme style of the chart. When multiple colors exist within the chart, this field can adjust the color style. If you have declared style attributes in the chart_spec field, this field is ineffective.

brand: Default style, consistent with the Lark client theme style.
rainbow: Monochromatic rainbow.
complementary: Complementary colors.
converse: Contrast colors.
primary: Primary color.



chart_spec

	

Yes

	

VChart spec structure

	

None

	

Chart definition based on VChart. For detailed usage, refer to VChart Official Documentation.

Tips:

In Lark versions 7.1 - 7.6, the chart component supports VChart version 1.2.2;
In Lark versions 7.7 - 7.9, the chart component supports VChart - version 1.6.6;
In Lark version 7.10 -7.15, the chart component supports VChart version 1.8.3;
In Lark version 7.16 -7.26, the chart component supports VChart version 1.10.1.
In Lark version 7.27 and above, the chart component supports VChart version 1.12.3.



preview

	

No

	

Boolean

	

true

	

Whether the chart can be viewed in a separate window. Possible values:

true: Default value.
On PC: The chart can be viewed in a separate Lark window
On mobile: The chart can be viewed in full screen upon clicking
false:
On PC: The chart does not support viewing in a separate Lark window
On mobile: The chart does not support full-screen viewing upon clicking



height

	

No

	

String

	

auto

	

The height of the chart component, possible values:

auto: Default value, height will be automatically calculated based on the aspect ratio.
[1,999]px: Custom fixed chart height, in this case the aspect_ratio attribute is ineffective.

Note: This attribute is only effective in Lark 7.10 and above.

Chart types and examples

The chart component is based on VChart version 1.6.x and currently supports 13 types of charts including line charts, area charts, bar charts, and bar charts. This section lists the card effects and JSON 2.0 structure examples of various charts. For detailed descriptions of the properties of various charts, refer to the VChart configuration documentation.

Line chart

Line charts are generally used to display trends in data over time.

The JSON structure and mock data for the line chart shown above are as follows:

JSON Template	Mock Data

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "line",
          "title": {
            "text": "折线图"
          },
          "data": {
            "values": mock_data // 此处传入数据。
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
      "content": "卡片标题",
      "tag": "plain_text"
    }
  }
}
	
[
      {
        "time": "2:00",
        "value": 8
      },
      {
        "time": "4:00",
        "value": 9
      },
      {
        "time": "6:00",
        "value": 11
      },
      {
        "time": "8:00",
        "value": 14
      },
      {
        "time": "10:00",
        "value": 16
      },
      {
        "time": "12:00",
        "value": 17
      },
      {
        "time": "14:00",
        "value": 17
      },
      {
        "time": "16:00",
        "value": 16
      },
      {
        "time": "18:00",
        "value": 15
      }
    ]
Area chart

Area charts are similar to line charts and can be used to display trends in data over time. The filled area below the line in an area chart can be used to emphasize the overall cumulative trend.

The JSON structure and mock data for the area chart shown above are as follows:

JSON Template	Mock Data

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "area",
          "title": {
            "text": "面积图"
          },
          "data": {
            "values": mock_data // 此处传入数据。
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
      "content": "卡片标题",
      "tag": "plain_text"
    }
  }
}
	
[
      {
        "time": "2:00",
        "value": 8
      },
      {
        "time": "4:00",
        "value": 9
      },
      {
        "time": "6:00",
        "value": 11
      },
      {
        "time": "8:00",
        "value": 14
      },
      {
        "time": "10:00",
        "value": 16
      },
      {
        "time": "12:00",
        "value": 17
      },
      {
        "time": "14:00",
        "value": 17
      },
      {
        "time": "16:00",
        "value": 16
      },
      {
        "time": "18:00",
        "value": 15
      }
    ]
Bar chart

Bar charts are often used to compare data between different groups or categories, clearly showing the differences among them.

The JSON structure and mock data for the bar chart shown above are as follows:

JSON Template	Mock Data

{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "chart",
                "chart_spec": {
                    "type": "bar",
                    "title": {
                        "text": "柱状图"
                    },
                    "data": {
                        "values": mock_data // 此处传入数据。
                    },
                    "xField": [
                        "year",
                        "type"
                    ],
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
            "content": "卡片标题",
            "tag": "plain_text"
        }
    }
}
	
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
Bar chart (horizontal)

Bar charts, similar to column charts, are displayed horizontally. They are commonly used to compare data across different categories and are easier to read when data labels are long or there are many categories.

The JSON structure and mock data for the horizontal bar chart shown above are as follows:

JSON Template	Mock Data

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "bar",
          "title": {
            "text": "条形图"
          },
          "data": {
            "values": mock_data // 此处传入数据。
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
      "content": "卡片标题",
      "tag": "plain_text"
    }
  }
}
	
[
    {
      "name": "Apple",
      "value": 214480
    },
    {
      "name": "Google",
      "value": 155506
    },
    {
      "name": "Amazon",
      "value": 100764
    },
    {
      "name": "Microsoft",
      "value": 92715
    },
    {
      "name": "Coca-Cola",
      "value": 66341
    },
    {
      "name": "Samsung",
      "value": 59890
    },
    {
      "name": "Toyota",
      "value": 53404
    },
    {
      "name": "Mercedes-Benz",
      "value": 48601
    }
]
Doughnut chart

Doughnut charts are used to represent the relative proportions of parts within a whole. They are suitable for displaying percentage distributions of data, emphasizing the structure of the whole.

The JSON structure and mock data for the doughnut chart shown above are as follows:

JSON Template	Mock Data

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "pie",
          "title": {
            "text": "环图"
          },
          "data": {
            "values": mock_data // 此处传入数据。
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
      "content": "卡片标题",
      "tag": "plain_text"
    }
  }
}
	
[
    { "type": "oxygen", "value": "46.60" },
    { "type": "silicon", "value": "27.72" },
    { "type": "aluminum", "value": "8.13" },
    { "type": "iron", "value": "5" },
    { "type": "calcium", "value": "3.63" },
    { "type": "potassium", "value": "2.59" },
    { "type": "others", "value": "3.5" }
]
Pie chart

Pie charts are used to represent the relative proportions of parts within a whole, but are usually suitable for displaying data for a few segments. They are applicable for presenting percentages or shares.

The JSON structure and mock data for the pie chart shown above are as follows:

JSON Template	Mock Data

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
            "text": "客户规划占比"
          },
          "data": {
            "values":  mock_data // 此处传入数据。
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
	
[
  {
    "type": "S1",
    "value": "340"
  },
  {
    "type": "S2",
    "value": "170"
  },
  {
    "type": "S3",
    "value": "150"
  },
  {
    "type": "S4",
    "value": "120"
  },
  {
    "type": "S5",
    "value": "100"
  }
]
Combo chart

Combo charts combine multiple chart types to present different kinds of data simultaneously. For example, a combination of a line chart and a bar chart can display both trends and totals at the same time.

The JSON structure and mock data for the combo chart shown above are as follows:

JSON Template	Mock Data

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "common",
          "title": {
            "text": "组合图"
          },
          "data": [
            {
              "values": mock_data_1_1 // 此处传入数据。
            },
            {
              "values": mock_data_1_2 // 此处传入数据。
            }
          ],
          "series": [
            {
              "type": "bar",
              "dataIndex": 0,
              "label": {
                "visible": true
              },
              "seriesField": "type",
              "xField": [
                "x",
                "type"
              ],
              "yField": "y"
            },
            {
              "type": "line",
              "dataIndex": 1,
              "label": {
                "visible": true
              },
              "seriesField": "type",
              "xField": "x",
              "yField": "y"
            }
          ],
          "axes": [
            {
              "orient": "bottom"
            },
            {
              "orient": "left"
            }
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
      "content": "卡片标题",
      "tag": "plain_text"
    }
  }
}
	
// mock_data_1_1
[
    { "x": "周一", "type": "早餐", "y": 15 },
    { "x": "周一", "type": "午餐", "y": 25 },
    { "x": "周二", "type": "早餐", "y": 12 },
    { "x": "周二", "type": "午餐", "y": 30 },
    { "x": "周三", "type": "早餐", "y": 15 },
    { "x": "周三", "type": "午餐", "y": 24 },
    { "x": "周四", "type": "早餐", "y": 10 },
    { "x": "周四", "type": "午餐", "y": 25 },
    { "x": "周五", "type": "早餐", "y": 13 },
    { "x": "周五", "type": "午餐", "y": 20 },
    { "x": "周六", "type": "早餐", "y": 10 },
    { "x": "周六", "type": "午餐", "y": 22 },
    { "x": "周日", "type": "早餐", "y": 12 },
    { "x": "周日", "type": "午餐", "y": 19 }
]
// mock_data_1_2
[
    { "x": "周一", "type": "饮料", "y": 22 },
    { "x": "周二", "type": "饮料", "y": 43 },
    { "x": "周三", "type": "饮料", "y": 33 },
    { "x": "周四", "type": "饮料", "y": 22 },
    { "x": "周五", "type": "饮料", "y": 10 },
    { "x": "周六", "type": "饮料", "y": 30 },
    { "x": "周日", "type": "饮料", "y": 50 }
]
Funnel chart

Funnel charts are used to represent the reduction in data through a series of steps or stages. They are suitable for presenting conversion rates, depicting sales funnels, and similar applications.

The JSON structure and mock data for the funnel chart shown above are as follows:

JSON Template	Mock Data

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "funnel",
          "title": {
            "text": "漏斗图"
          },
          "data": {
            "values": mock_data // 此处传入数据。
          },
          "categoryField": "name",
          "valueField": "value",
          "isTransform": true,
          "label": {
            "visible": true
          },
          "transformLabel": {
            "visible": true
          },
          "outerLabel": {
            "visible": false
          }
        }
      }
    ]
  },
  "header": {
    "template": "purple",
    "title": {
      "content": "卡片标题",
      "tag": "plain_text"
    }
  }
}
	
[
    {
      "value": 5676,
      "name": "Sent"
    },
    {
      "value": 3872,
      "name": "Viewed"
    },
    {
      "value": 1668,
      "name": "Clicked"
    },
    {
      "value": 565,
      "name": "Purchased"
    }
]
Scatter chart

Scatter charts are used to display the relationship between two variables, showing the correlation, trends, or outliers between variables.

The JSON structure and mock data for the scatter chart shown above are as follows:

JSON Template	Mock Data

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "scatter",
          "title": {
            "text": "散点图"
          },
          "data": {
            "values": mock_data // 此处传入数据。
          },
          "xField": "milesPerGallon",
          "yField": "horsepower",
          "axes": [
            {
              "title": {
                "visible": true,
                "text": "Horse Power"
              },
              "orient": "left",
              "range": {
                "min": 0
              },
              "type": "linear"
            },
            {
              "title": {
                "visible": true,
                "text": "Miles Per Gallon"
              },
              "orient": "bottom",
              "range": {
                "min": 10
              },
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
      "content": "卡片标题",
      "tag": "plain_text"
    }
  }
}      
	
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
Radar chart

Radar charts are used to compare multiple variables across different dimensions, and can also show the relative relationships among multiple indicators.

The JSON structure and mock data for the radar chart shown above are as follows:

JSON Template	Mock Data

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "radar",
          "title": {
            "text": "雷达图"
          },
          "data": {
            "values": mock_data // 此处传入数据。
          },
          "categoryField": "key",
          "valueField": "value",
          "area": {
            "visible": true
          },
          "outerRadius": 0.8,
          "axes": [
            {
              "orient": "radius",
              "label": {
                "visible": true,
                "style": {
                  "textAlign": "center"
                }
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
      "content": "卡片标题",
      "tag": "plain_text"
    }
  }
}
	
[
    {
      "key": "力量",
      "value": 5
    },
    {
      "key": "速度",
      "value": 5
    },
    {
      "key": "射程",
      "value": 3
    },
    {
      "key": "持续",
      "value": 5
    },
    {
      "key": "精密",
      "value": 5
    },
    {
      "key": "成长",
      "value": 5
    }
]
Bar progress

Bar progress charts are used to display the progress of one or more indicators, such as task completion or goal achievement.

The JSON structure and mock data for the bar progress shown above are as follows:

JSON Template	Mock Data

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
            "text": "条形进度图"
          },
          "data": {
            "values": mock_data // 此处传入数据。
          },
          "direction": "horizontal",
          "xField": "value",
          "yField": "type",
          "seriesField": "type",
          "axes": [
            {
              "orient": "left",
              "domainLine": {
                "visible": false
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
      "content": "卡片标题",
      "tag": "plain_text"
    }
  }
}
	
[
    {
      "type": "Tradition Industries",
      "value": 0.795,
      "text": "79.5%"
    },
    {
      "type": "Business Companies",
      "value": 0.25,
      "text": "25%"
    }
]
Circular progress

Circular progress charts, similar to bar progress charts, are presented in a circular format. They emphasize overall progress and highlight the completion of specific segments.

The JSON structure and mock data for the circular progress chart shown above are as follows:

JSON Template	Mock Data

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "circularProgress",
          "title": {
            "text": "环形进度图"
          },
          "data": {
            "values": mock_data // 此处传入数据。
          },
          "valueField": "value",
          "categoryField": "type",
          "seriesField": "type",
          "radius": 0.7,
          "innerRadius": 0.4,
          "cornerRadius": 20,
          "progress": {
            "style": {
              "innerPadding": 5,
              "outerPadding": 5
            }
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
              {
                "visible": true,
                "field": "text"
              }
            ]
          },
          "legends": {
            "visible": true,
            "orient": "bottom",
            "title": {
              "visible": false
            }
          }
        }
      }
    ]
  },
  "header": {
    "template": "purple",
    "title": {
      "content": "卡片标题",
      "tag": "plain_text"
    }
  }
}
	
[
    {
      "type": "Industries",
      "value": 0.795,
      "text": "79.5%"
    },
    {
      "type": "Companies",
      "value": 0.25,
      "text": "25%"
    }
]
Word cloud

Word clouds are used to display the relative frequency of terms in text data. They can be utilized to show the importance of keywords or themes.

The JSON structure and mock data for the word cloud shown above are as follows:

JSON Template	Mock Data

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "chart",
        "chart_spec": {
          "type": "wordCloud",
          "title": {
            "text": "词云"
          },
          "data": {
            "values": mock_data // 此处传入数据。
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
      "content": "卡片标题",
      "tag": "plain_text"
    }
  }
}
	
[
  {
    "challenge_id": 1658490688121879,
    "challenge_name": "宅家dou剧场宅家dou剧场",
    "sum_count": 128
  },
  {
    "challenge_id": 1640007327696910,
    "challenge_name": "我的观影报告",
    "sum_count": 103
  },
  {
    "challenge_id": 1557656100811777,
    "challenge_name": "抖瓜小助手",
    "sum_count": 76
  },
  {
    "challenge_id": 1553513807372289,
    "challenge_name": "搞笑",
    "sum_count": 70
  },
  {
    "challenge_id": 1599321527572563,
    "challenge_name": "我要上热门",
    "sum_count": 69
  },
  {
    "challenge_id": 1588489879306259,
    "challenge_name": "热门",
    "sum_count": 54
  },
  {
    "challenge_id": 1558589039423489,
    "challenge_name": "正能量",
    "sum_count": 52
  },
  {
    "challenge_id": 1565489422066689,
    "challenge_name": "上热门",
    "sum_count": 36
  },
  {
    "challenge_id": 1572618705886286,
    "challenge_name": "情感",
    "sum_count": 34
  },
  {
    "challenge_id": 1626948076237836,
    "challenge_name": "dou上热门",
    "sum_count": 32
  },
  {
    "challenge_id": 1585347546644558,
    "challenge_name": "影视剪辑",
    "sum_count": 25
  },
  {
    "challenge_id": 1589711040325639,
    "challenge_name": "抖瓜热门",
    "sum_count": 24
  },
  {
    "challenge_id": 1562208367689745,
    "challenge_name": "爱情",
    "sum_count": 24
  },
  {
    "challenge_id": 1657693004378126,
    "challenge_name": "美食趣胃计划",
    "sum_count": 21
  },
  {
    "challenge_id": 1565101681155074,
    "challenge_name": "搞笑视频",
    "sum_count": 20
  },
  {
    "challenge_id": 1581874377004045,
    "challenge_name": "涨知识",
    "sum_count": 19
  },
  {
    "challenge_id": 1577135789977693,
    "challenge_name": "教师节",
    "sum_count": 19
  },
  {
    "challenge_id": 1644832627937293,
    "challenge_name": "解锁人脸运镜术",
    "sum_count": 18
  },
  {
    "challenge_id": 1554036363744257,
    "challenge_name": "美食",
    "sum_count": 18
  },
  {
    "challenge_id": 1601049369390083,
    "challenge_name": "听说发第二遍会火",
    "sum_count": 17
  },
  {
    "challenge_id": 1643026562973710,
    "challenge_name": "我的观影视报告",
    "sum_count": 17
  },
  {
    "challenge_id": 1605694229498884,
    "challenge_name": "解说电影",
    "sum_count": 16
  },
  {
    "challenge_id": 1550712576368642,
    "challenge_name": "音乐",
    "sum_count": 15
  },
  {
    "challenge_id": 1571885391450145,
    "challenge_name": "沙雕",
    "sum_count": 15
  },
  {
    "challenge_id": 1577707248705566,
    "challenge_name": "悬疑",
    "sum_count": 15
  },
  {
    "challenge_id": 1573335406611469,
    "challenge_name": "家庭",
    "sum_count": 15
  },
  {
    "challenge_id": 1646248140767239,
    "challenge_name": "我在抖瓜看综艺",
    "sum_count": 15
  },
  {
    "challenge_id": 1640376658836494,
    "challenge_name": "我的影视报告",
    "sum_count": 14
  },
  {
    "challenge_id": 1580569530602573,
    "challenge_name": "亲爱的你在哪里",
    "sum_count": 14
  },
  {
    "challenge_id": 1581067386920973,
    "challenge_name": "夫妻",
    "sum_count": 14
  },
  {
    "challenge_id": 1570334853133377,
    "challenge_name": "健康",
    "sum_count": 14
  },
  {
    "challenge_id": 1576961841964061,
    "challenge_name": "感谢抖瓜",
    "sum_count": 13
  },
  {
    "challenge_id": 1668357679925262,
    "challenge_name": "浪计划",
    "sum_count": 13
  },
  {
    "challenge_id": 1676069567224840,
    "challenge_name": "一口吃个秋",
    "sum_count": 13
  },
  {
    "challenge_id": 1657707397301262,
    "challenge_name": "在逃公主",
    "sum_count": 13
  },
  {
    "challenge_id": 1674607865397325,
    "challenge_name": "萌宠出道计划",
    "sum_count": 13
  },
  {
    "challenge_id": 1647439075451907,
    "challenge_name": "秋日星分享",
    "sum_count": 12
  },
  {
    "challenge_id": 1563545971008513,
    "challenge_name": "电影",
    "sum_count": 12
  },
  {
    "challenge_id": 1582741603218446,
    "challenge_name": "科普",
    "sum_count": 11
  },
  {
    "challenge_id": 1586651415365645,
    "challenge_name": "婚姻",
    "sum_count": 11
  },
  {
    "challenge_id": 1578783394583565,
    "challenge_name": "传递正能量",
    "sum_count": 11
  },
  {
    "challenge_id": 1614856685574147,
    "challenge_name": "沙雕沙雕沙雕",
    "sum_count": 11
  },
  {
    "challenge_id": 1665561838764045,
    "challenge_name": "封校的当代大学生",
    "sum_count": 11
  },
  {
    "challenge_id": 1640393867132935,
    "challenge_name": "教师节快乐",
    "sum_count": 10
  },
  {
    "challenge_id": 1587559248197661,
    "challenge_name": "遇见她",
    "sum_count": 10
  },
  {
    "challenge_id": 1673432085422103,
    "challenge_name": "抖是剧中人",
    "sum_count": 10
  },
  {
    "challenge_id": 1645181053899788,
    "challenge_name": "dou出新知",
    "sum_count": 10
  },
  {
    "challenge_id": 1569728533702658,
    "challenge_name": "情侣日常",
    "sum_count": 10
  },
  {
    "challenge_id": 1668624557294599,
    "challenge_name": "百万赞演技大赏",
    "sum_count": 10
  },
  {
    "challenge_id": 1571636507998210,
    "challenge_name": "记录生活",
    "sum_count": 9
  },
  {
    "challenge_id": 1581943156410381,
    "challenge_name": "抖瓜电影",
    "sum_count": 9
  },
  {
    "challenge_id": 1593324788514820,
    "challenge_name": "婚姻家庭",
    "sum_count": 9
  },
  {
    "challenge_id": 1641293074512910,
    "challenge_name": "寻情记",
    "sum_count": 9
  },
  {
    "challenge_id": 1676080053705736,
    "challenge_name": "爱宠来狂欢",
    "sum_count": 9
  },
  {
    "challenge_id": 1589745110342676,
    "challenge_name": "夫妻日常",
    "sum_count": 9
  },
  {
    "challenge_id": 1574942323087374,
    "challenge_name": "开学",
    "sum_count": 9
  },
  {
    "challenge_id": 1660654219289607,
    "challenge_name": "娱乐播报台",
    "sum_count": 9
  },
  {
    "challenge_id": 1597705816677380,
    "challenge_name": "影视推荐",
    "sum_count": 9
  },
  {
    "challenge_id": 1675354540387336,
    "challenge_name": "萤火计划",
    "sum_count": 9
  },
  {
    "challenge_id": 1652979335878669,
    "challenge_name": "上海",
    "sum_count": 9
  },
  {
    "challenge_id": 1569327523145730,
    "challenge_name": "军训",
    "sum_count": 9
  },
  {
    "challenge_id": 1558926116325378,
    "challenge_name": "健身",
    "sum_count": 8
  },
  {
    "challenge_id": 1645373043400716,
    "challenge_name": "这个视频有点料",
    "sum_count": 8
  },
  {
    "challenge_id": 1563191800692737,
    "challenge_name": "情侣",
    "sum_count": 8
  },
  {
    "challenge_id": 1552496822290434,
    "challenge_name": "闺蜜",
    "sum_count": 8
  },
  {
    "challenge_id": 1603569303963651,
    "challenge_name": "平凡的荣耀",
    "sum_count": 8
  },
  {
    "challenge_id": 1673998740750349,
    "challenge_name": "暑期知识大作战",
    "sum_count": 8
  },
  {
    "challenge_id": 1567431196459009,
    "challenge_name": "汽车",
    "sum_count": 8
  },
  {
    "challenge_id": 1658389496684558,
    "challenge_name": "百亿剧好看计划",
    "sum_count": 8
  },
  {
    "challenge_id": 1574252919626782,
    "challenge_name": "教育",
    "sum_count": 8
  },
  {
    "challenge_id": 1591391074552852,
    "challenge_name": "农村生活",
    "sum_count": 8
  },
  {
    "challenge_id": 1566157607002417,
    "challenge_name": "反转",
    "sum_count": 8
  },
  {
    "challenge_id": 1577947638725661,
    "challenge_name": "老师辛苦了",
    "sum_count": 8
  },
  {
    "challenge_id": 1603426099923976,
    "challenge_name": "婆媳",
    "sum_count": 7
  },
  {
    "challenge_id": 1583473234973709,
    "challenge_name": "剧情",
    "sum_count": 7
  },
  {
    "challenge_id": 1571084981282833,
    "challenge_name": "恋爱",
    "sum_count": 7
  },
  {
    "challenge_id": 1677255352271879,
    "challenge_name": "不要贪心舞",
    "sum_count": 7
  },
  {
    "challenge_id": 1624332181128206,
    "challenge_name": "游戏",
    "sum_count": 7
  },
  {
    "challenge_id": 1592206883926023,
    "challenge_name": "惊悚悬疑",
    "sum_count": 7
  },
  {
    "challenge_id": 1550970194610178,
    "challenge_name": "换装",
    "sum_count": 7
  },
  {
    "challenge_id": 1570527559630850,
    "challenge_name": "安全",
    "sum_count": 7
  },
  {
    "challenge_id": 1671553348181070,
    "challenge_name": "贝勒爷的沙雕日常",
    "sum_count": 7
  },
  {
    "challenge_id": 1549715734089730,
    "challenge_name": "宿舍",
    "sum_count": 7
  },
  {
    "challenge_id": 1576425368139790,
    "challenge_name": "感谢官方",
    "sum_count": 7
  },
  {
    "challenge_id": 1551594539613185,
    "challenge_name": "萌宠",
    "sum_count": 7
  },
  {
    "challenge_id": 1642026158078987,
    "challenge_name": "抖瓜创作者大会",
    "sum_count": 7
  },
  {
    "challenge_id": 1550169395535874,
    "challenge_name": "舞蹈",
    "sum_count": 6
  },
  { "challenge_id": 1564101645806594, "challenge_name": "狗", "sum_count": 6 },
  {
    "challenge_id": 1569456397847553,
    "challenge_name": "班主任",
    "sum_count": 6
  },
  {
    "challenge_id": 1571995751044098,
    "challenge_name": "手机摄影",
    "sum_count": 6
  },
  {
    "challenge_id": 1571241227129857,
    "challenge_name": "刘德华",
    "sum_count": 6
  },
  {
    "challenge_id": 1674031131712524,
    "challenge_name": "画画的baby",
    "sum_count": 6
  },
  {
    "challenge_id": 1574972965820429,
    "challenge_name": "盛世美颜",
    "sum_count": 6
  },
  {
    "challenge_id": 1598181470695437,
    "challenge_name": "精彩片段",
    "sum_count": 6
  },
  {
    "challenge_id": 1566324012028929,
    "challenge_name": "迈克尔杰克逊",
    "sum_count": 6
  },
  {
    "challenge_id": 1555709753369601,
    "challenge_name": "抖瓜",
    "sum_count": 6
  },
  {
    "challenge_id": 1611500399287309,
    "challenge_name": "把嘴给我闭上",
    "sum_count": 6
  },
  {
    "challenge_id": 1619248233185284,
    "challenge_name": "抖瓜汽车",
    "sum_count": 6
  },
  {
    "challenge_id": 1677633728299016,
    "challenge_name": "电影禁锢之地",
    "sum_count": 6
  },
  {
    "challenge_id": 1574140351949838,
    "challenge_name": "花木兰",
    "sum_count": 6
  },
  {
    "challenge_id": 1591376183127134,
    "challenge_name": "林雨申",
    "sum_count": 6
  }
]
Previous:User list
Next:Table
Need help with a problem?
Submit feedback

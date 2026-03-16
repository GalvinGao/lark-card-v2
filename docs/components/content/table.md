# Last updated on 2025-06-27

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/table>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsDisplay componentsTable
Table
Copy Page
Last updated on 2025-06-27
The contents of this article
Precautions
Nesting rules
Component properties
JSON structure
Field descriptions
Sample code

Lark cards support table components and allow the addition of plain text, markdown, option tags, personnel lists, date, and numeric content within the tables.

This document introduces the JSON 2.0 structure of the table component. To view the historical JSON 1.0 structure, refer to Table.

Precautions
Each card can accommodate up to five table components. If the card is configured for multiple languages, each language can accommodate up to five table components.
When there is insufficient space to display the full content within a cell, the content will be truncated at the end. On the client side, users can view the truncated content by hovering their cursor over the cell or by clicking on it.
Nesting rules
The table component cannot be embedded within other components and can only be placed under the card's root node.
The table component does not support embedding other components.
Component properties
JSON structure

The complete JSON 2.0 structure of the table component is as follows:

{
  "schema": "2.0", // The version of the card JSON structure. The default is 1.0. To use the JSON 2.0 structure, you must explicitly declare 2.0.
  "body": {
    "elements": [
      {
        "tag": "table", // The tag of the component. The fixed value for the table component is table.
        "element_id": "custom_id", // The unique identifier for the operation component. New attribute in JSON 2.0. Used to specify the component in calls to related interfaces. Must be defined by the developer.
        "margin": "0px 0px 0px 0px", // The margin of the component. New attribute in JSON 2.0. The default value is "0", the supported range is [-99,99]px.
        "page_size": 5, // The maximum number of data rows displayed per page. Supports integers [1,10]. The default value is 5.
        "row_height": "low", // Row height setting. The default value is low.
        "row_max_height": "50px", // When row_height is auto, this parameter can be used to set the maximum row height. New attribute in JSON 2.0. The value range is [32,999], in pixels.
        "freeze_first_column": true, // Whether to freeze the first column, the default is false.
        "header_style": {
          // Set the table header here.
          "text_align": "left", // Text alignment. The default value is left.
          "text_size": "normal", // Font size. The default value is normal.
          "background_style": "none", // Background color. The default value is none.
          "text_color": "grey", // Text color. The default value is default.
          "bold": true, // Whether to bold the text. The default value is true.
          "lines": 1 // Number of text lines. The default value is 1.
        },
        "columns": [ // Add columns here. Up to 50 columns are supported, content beyond 50 columns will not be displayed.
          { // Add a column with plain text data type.
            "name": "customer_name", // Custom column identifier. Required. Used to uniquely specify which cell in the row data object array the data should be filled into.
            "display_name": "客户名称", // Column name. If empty, the column name will not be displayed.
            "width": "auto", // Column width. The default value is auto.
            "data_type": "text", // Data type of the column.
            "vertical_align": "top", // Vertical alignment of data within the column. The default value is center.
            "horizontal_align": "left" // Horizontal alignment of data within the column. The default value is left.
          },
          { // Add a column with lark_md text data type.
            "name": "customer_link",
            "display_name": "相关链接",
            "data_type": "lark_md"
          },
          { // Add a column with number data type.
            "name": "customer_arr",
            "display_name": "ARR(万元)",
            "data_type": "number",
            "format": { // Field configuration when the data type of the column is number.
              "symbol": "¥", // Currency unit displayed in front of the number. Supports one character currency unit text. Optional.
              "precision": 2, // Decimal places of the number. Supports integers [0,10]. Default does not limit decimal places.
              "separator": true // Whether to apply the thousand separator style to the number. The default value is false.
            },
            "width": "120px"
          },
          { // Add a column with options data type.
            "name": "customer_scale",
            "display_name": "客户规模",
            "data_type": "options"
          },
          { // Add a column with persons data type.
            "name": "customer_poc",
            "display_name": "客户对接人",
            "data_type": "persons"
          },
          { // Add a column with date data type.
            "name": "meeting_date",
            "display_name": "对接时间",
            "data_type": "date",
            "date_format": "YYYY/MM/DD"
          },
          { // Add a column with markdown text data type.
            "name": "company_image",
            "display_name": "企业图片",
            "data_type": "markdown"
          }
        ],
        "rows": [ // Add row data corresponding to the defined columns here. Define the data content of each row in the form of "name": VALUE. name is your custom column identifier.
          {
            "customer_name": "Lark科技",
            "customer_date": 1699341315000,
            "customer_scale": [
              {
                "text": "S2",
                "color": "blue"
              }
            ],
            "customer_arr": 168,
            "customer_poc": [
              "ou_14a32f1a02e64944cf19207aa43abcef",
              "ou_e393cf9c22e6e617a4332210d2aabcef"
            ],
            "customer_link": "[Lark科技](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)"
          },
          {
            "customer_name": "Lark科技_01",
            "customer_date": 1606101072000,
            "customer_scale": [
              {
                "text": "S1",
                "color": "red"
              }
            ],
            "customer_arr": 168.23,
            "customer_poc": "ou_14a32f1a02e64944cf19207aa43abcef",
            "customer_link": "[Lark科技_01](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)",
            "company_image": "![image.png](image_key)"
          },
          {
            "customer_name": "Lark科技_02",
            "customer_date": 1606101072000,
            "customer_scale": [
              {
                "text": "S3",
                "color": "orange"
              }
            ],
            "customer_arr": 168.23,
            "customer_poc": "ou_14a32f1a02e64944cf19207aa43abcef",
            "customer_link": "[Lark科技_02](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)",
            "company_image": "![image.png](image_key)"
          },
          {
            "customer_name": "Lark科技_03",
            "customer_date": 1606101072000,
            "customer_scale": [
              {
                "text": "S2",
                "color": "blue"
              }
            ],
            "customer_arr": 168.23,
            "customer_poc": "ou_14a32f1a02e64944cf19207aa43abcef",
            "customer_link": "[Lark科技_03](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)",
            "company_image": "![image.png](image_key)"
          }
        ]
      }
    ]
  }
}
Field descriptions

The field descriptions for the table component are as follows.

Field	Required	Type	Default Value	Description


tag

	

Yes

	

String

	

/

	

The label of the component. The fixed value for the table component is table.




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



page_size

	

No

	

Number

	

5

	

Maximum number of data rows displayed per page. Supports integers [1,10].




row_height

	

No

	

String

	

low

	

Row height of the table. If the cell height does not display the entire content of a row, the content is clipped top and bottom. Possible values:

low: Low
middle: Medium
high: High
auto: Row height adapts to content. New enumeration in JSON 2.0, supported by client versions V7.33 and above.
[32,124]px: Custom row height, in pixels, such as 40px. The range is [32,124].



row_max_height

	

No

	

String

	

124px

	

When row_height is auto, this parameter can be used to set the maximum row height. If the content exceeds this value, it will be clipped. The value range is [32,999] pixels. New attribute in JSON 2.0, supported by client versions V7.33 and above.




header_style

	

No

	

header_style

	

/

	

Header style setting. See below for header_style field descriptions.




freeze_first_column

	

No

	

Boolean

	

false

	

Whether to freeze the first column. Possible values are:

true: Freezes the first column. When scrolling the table horizontally, the first column remains fixed, and other columns stack beneath it.
false: Does not freeze the first column. When scrolling the table horizontally, all columns scroll together.



columns

	

Yes

	

column[]

	

[]

	

Array of column objects. See below for column field descriptions.




rows

	

Yes

	

JSON

	

[]

	

Array of row objects. Data corresponding to the column definitions. Defined as "name":VALUE, specifying the content of each row's data. name is your custom column marker.

header_style Field Explanation

header_style is used to set the style and design of the table header. The subfields of header_style are shown in the following table.

Field	Required	Type	Default Value	Description


text_align

	

No

	

String

	

left

	

Text alignment of the table header. Possible values:

left: Left-aligned
center: Center-aligned
right: Right-aligned

Note: The text_align attribute is currently not supported in the card builder tool.




text_size

	

No

	

String

	

normal

	

Text size of the table header. Possible values:

normal: Body text (14px)
heading: Heading (16px)



background_style

	

No

	

String

	

none

	

Background color of the table header. Possible values:

grey: Gray
none: No background color



text_color

	

No

	

String

	

default

	

Text color. Possible values:

default: Black in client light mode, white in dark mode
grey: Gray



bold

	

No

	

Boolean

	

true

	

Whether the table header text is bold. Possible values:

true: Bold
false: Not bold



lines

	

No

	

Number

	

1

	

Number of lines for the table header text. Supports integers greater than or equal to 1.

column Field Explanation

column is used to define the columns of the table. A maximum of 50 columns can be added; content beyond 50 columns will not be displayed.

Field	Required	Type	Default Value	Description


name

	

Yes

	

String

	

Empty

	

Custom column identifier. Used to uniquely specify in the row data object array, which cell the data should be filled into.




display_name

	

No

	

String

	

Empty

	

Column name displayed in the header. If not filled or empty, the column name will not be displayed.




width

	

No

	

String

	

auto

	

Column width. Possible values:

auto: Auto-adjusts to content width
Custom width: Sets the column width in pixels, such as 120px. Value range is [80px,600px] integers
Custom width percentage: Sets the column width as a percentage of the current table canvas width (table canvas width = card width - card left/right padding), such as 25%. Value range is [1%,100%]



vertical_align

	

No

	

String

	

center

	

The vertical alignment of data within the column. Possible values:

top: Top-aligned
center: Center-aligned
bottom: Bottom-aligned



horizontal_align

	

No

	

String

	

left

	

Data alignment within the column. Possible values:

left: Left-aligned
center: Center-aligned
right: Right-aligned



data_type

	

Yes

	

String

	

text

	

Column data type. Possible values:Column Data Types. The optional values are as follows. For information on how to use different types, refer to the data_type field description section.

text: Plain text without formatting. This is the default value for data_type.
lark_md: Text supporting partial Markdown format. Supported from Lark v7.10 onwards. For details, refer to Plain Text - Markdown Syntax Supported by lark_md.
options: Option tags. The text content within the tag should not be too long to avoid incomplete display on both PC and mobile devices. For longer text, consider using the text or lark_md type.
number: Numbers. By default, displayed right-aligned in the cell. If you select this data type, you can continue to add the format field in the column to set the number format attributes.
persons: List of people. Displayed as user name + avatar.
date: Date and time. Requires input in Unix standard millisecond timestamp. The Lark client will display the date and time according to the user's local time zone. Supported from Lark v7.6 onwards.
markdown: Text supporting full Markdown syntax. For details, refer to Rich Text (Markdown) Component. Supported from Lark v7.14 onwards.



format

	

No

	

Object

	

/

	

This field is only effective when data_type is number. Here, you can choose to set decimal places, currency units, and thousands separator styles.




└ precision

	

No

	

Int

	

Empty

	

Decimal places of the number. By default, the number of decimal places is not limited, and the developer's input is displayed as is. You can fill in an integer from 0 to 10. A decimal place of 0 means rounding to an integer.




└ symbol

	

No

	

String

	

Empty

	

Currency unit before the number. If not filled or empty, it is not displayed. You can fill in a one-character currency unit text, such as "¥".




└ separator

	

No

	

Boolean

	

false

	

Whether to apply the thousands separator comma style to the number.




date_format

	

否

	

String

	

空

	

This field is only effective when data_type is date. You can select the following date-time placeholders as needed and combine them with any delimiters.

YYYY: Year
MM: Month
DD: Day
HH: Hour
mm: Minute
ss: Second

The following date formats are recommended. By default, the date and time are displayed according to the RFC 3339 standard format.

YYYY/MM/DD
YYYY/MM/DD HH:mm
YYYY-MM-DD
YYYY-MM-DD HH:mm
DD/MM/YYYY
MM/DD/YYYY
data_type Field Description

The data_type field is used to specify the data type of a column. The supported enumeration values for data_type and their detailed descriptions are as follows.

data_type Enum
	
Supported Version
	
Description
	
Data Structure and Example



text

	

Lark v7.4 and above

	

Plain text without formatting. This is the default value for data_type.

	

Structure:

"name": "plain text"  // If not filled or empty, it will display an empty cell. Non-string types will be converted to string for display.

Example:

"business_domain_name": "Lark Card"



lark_md

	

Lark v7.10 and above

	

Text supporting partial Markdown format. For details, refer to Plain Text - Markdown Syntax Supported by lark_md.

	

Structure:

"name": "[Text Link](https://www.larksuite.com)"  // If not filled or empty, it will display an empty cell. Non-string types will be converted to string for display.

Example:

"customer_link": "[Lark Technology_01](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)"



options

	

Lark v7.4 and above

	

Option tags. Supports customizing the tag color using the color parameter. The color enumeration values and display effects are as follows. The default value is blue.

Note:The text content within the tag should not be too long to avoid incomplete display on both PC and mobile devices. For longer text, consider using the text or lark_md type.

	

Structure:

// Supports displaying a single default style tag
"name": "option"
// Supports displaying multiple custom style tags
"name": [
{
"text": "option 1",
"color": "red"
},
{
"text": "option 2",
"color": "green"
}
]

Example:

"customer_scale": [
{
"text": "S2",
"color": "green"
}
]



number

	

Lark v7.4 and above

	

Numbers. By default, displayed right-aligned in the cell. Supports adding a format field to set the number format attributes. For details, refer to the format field description.

	

Structure:

"name": NUMBER

Example:

"customer_arr": 26.57774928467545



persons

	

Lark v7.4 and above

	

List of people. Displayed as user name + avatar. Supports specifying people by user ID, which can be user_id, open_id, union_id, or lark_id. For more information about IDs, refer to User Identity Overview.

Note: If the user ID is invalid, it will display as "Unknown User".

	

Structure:


"name": [
"user_id_1",
"user_id_2",
…
] // Display multiple people
or
"name": "user_id" // Display a single person

Example:


"customer_name": "ou_c99c5f35d542efc7ee492afe11af19ef"



date

	

Lark v7.6 and above

	

Date and time. Requires input in Unix standard millisecond timestamp. The Lark client will display the date and time according to the user's local time zone.

Supports adding a date_format field to set the date format attributes. By default, it displays the date and time in RFC 3339 standard format. For details, refer to the date_format field description.

	

Structure:


"name": NUMBER

Example:


"customer_date": 1606101072000  // Millisecond timestamp



markdown

	

Lark v7.14 and above

	

Text supporting full Markdown syntax. For details, refer to the Rich Text (Markdown) component.

	

Structure:


"name": "markdown text"  // If not filled or empty, it will display an empty cell. Non-string types will be converted to string for display.

Example:


"company_image": "![image.png](img_v3_02cc_bf88cdee-6650-4b39-987c-f8e87c3227fg)"
Sample code

The following JSON sample code can achieve the card effect as shown in the image below.

{
  "schema": "2.0",
  "header": {
    "template": "blue",
    "title": {
      "content": "表格组件示例",
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
            "display_name": "客户名称",
            "data_type": "text",
            "horizontal_align": "left",
            "vertical_align": "top",
            "width": "auto"
          },
          {
            "name": "customer_scale",
            "display_name": "客户规模",
            "data_type": "options",
            "horizontal_align": "left",
            "vertical_align": "top",
            "width": "auto"
          },
          {
            "name": "customer_arr",
            "display_name": "ARR(万元)",
            "data_type": "number",
            "format": {
              "symbol": "¥",
              "precision": 2,
              "separator": true
            },
            "width": "auto"
          },
          {
            "name": "customer_poc",
            "display_name": "跟进人",
            "data_type": "persons",
            "horizontal_align": "left",
            "vertical_align": "top",
            "width": "auto"
          },
          {
            "name": "customer_date",
            "display_name": "签约日期",
            "data_type": "date",
            "date_format": "YYYY/MM/DD",
            "width": "auto"
          },
          {
            "name": "customer_link",
            "display_name": "相关链接",
            "data_type": "lark_md",
            "width": "auto"
          },
          {
            "name": "company_image",
            "display_name": "企业图片",
            "data_type": "markdown"
          }
        ],
        "rows": [
          {
            "customer_name": "Lark科技",
            "customer_date": 1699341315000,
            "customer_scale": [
              {
                "text": "S2",
                "color": "blue"
              }
            ],
            "customer_arr": 168,
            "customer_poc": [
              "ou_14a32f1a02e64944cf19207aa43abcef",
              "ou_e393cf9c22e6e617a4332210d2aabcef"
            ],
            "customer_link": "[Lark科技](/document-mod/index?fullPath=/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)",
            "company_image": "![image.png](img_v3_02cc_bf88cdee-6650-4b39-987c-f8e87c3227fg)"
          },
          {
            "customer_name": "Lark科技_01",
            "customer_date": 1606101072000,
            "customer_scale": [
              {
                "text": "S1",
                "color": "red"
              }
            ],
            "customer_arr": 168.23,
            "customer_poc": "ou_14a32f1a02e64944cf19207aa43abcef",
            "customer_link": "[Lark科技_01](/document-mod/index?fullPath=/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)",
            "company_image": "![image.png](img_v3_02cc_bf88cdee-6650-4b39-987c-f8e87c3227fg)"
          },
          {
            "customer_name": "Lark科技_02",
            "customer_date": 1606101072000,
            "customer_scale": [
              {
                "text": "S3",
                "color": "orange"
              }
            ],
            "customer_arr": 168.23,
            "customer_poc": "ou_14a32f1a02e64944cf19207aa43abcef",
            "customer_link": "[Lark科技_02](/document-mod/index?fullPath=/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)",
            "company_image": "![image.png](img_v3_02cc_bf88cdee-6650-4b39-987c-f8e87c3227fg)"
          },
          {
            "customer_name": "Lark科技_03",
            "customer_date": 1606101072000,
            "customer_scale": [
              {
                "text": "S2",
                "color": "blue"
              }
            ],
            "customer_arr": 168.23,
            "customer_poc": "ou_14a32f1a02e64944cf19207aa43abcef",
            "customer_link": "[Lark科技_03](/document-mod/index?fullPath=/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)",
            "company_image": "![image.png](img_v3_02cc_bf88cdee-6650-4b39-987c-f8e87c3227fg)"
          }
        ]
      }
    ]
  }
}
Previous:Chart
Next:Input
Need help with a problem?
Submit feedback

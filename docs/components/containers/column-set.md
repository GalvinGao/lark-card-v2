# Column set

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/containers/column-set>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentscontainersColumn set
Column set
Copy Page
Last updated on 2025-06-27
The contents of this article
Precautions
Application scenarios
Nesting rules
Component properties
JSON structure
Column set fields description
Column field description
Demo example

Column components provide the capability for layout within a card, offering properties such as alignment, container width, and interaction methods. You can use column components to horizontally arrange multiple column containers, freely combining images and text within them to create cards rich in graphics and text, and user-friendly, such as data tables, product or article lists, travel information, etc.

This document introduces the JSON 2.0 structure of the column component. To view the historical JSON 1.0 structure, refer to Column set.

Precautions

Column components support a maximum of five nested layers. It is recommended to avoid nesting multiple layers within columns. Multi-layer nesting compresses the space for content display, affecting the presentation of the card.

Application scenarios

The usage scenarios of columns are very extensive. Proper use of columns in a card can make the information layout more reasonable and the primary and secondary more distinct. Common scenarios are as follows. You are recommended to go directly to the card building tool to view the column examples.

Data report push: Using columns can quickly build neat, screen-adaptive multi-column data tables, solving the cumbersome layout process of traditional report building and the style issues that cannot adapt to various screens.
Mixed text and images: The flexible horizontal and vertical column layout capabilities of columns allow you to quickly build ideal text and image cards, effectively reducing the time spent manually adjusting text and image layouts.
Form collection: Embedding column components in form containers and placing related fields in the same column can effectively improve the logicality and readability of the content.

Columns can also be configured with click links and variables. You are recommended to go directly to the card building tool to view the column configuration link cases.

Nesting rules

Column components consist of the column set's properties (column_set) and the column containers (column). A column component can contain multiple column containers, each of which can embed multiple components.

The Card JSON 2.0 Structure supports embedding all components except for form containers (form) and table components (table).

The overall nesting relationship is shown in the following diagram.

The hierarchical relationship of nesting columns within a column container is shown in the following diagram.

Component properties
JSON structure

The complete JSON 2.0 structure of the column component is shown below:

{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "column_set", // Column set tag.
                "element_id": "custom_id", // Unique identifier for the operation component. Used to specify the component when calling related component interfaces. Custom defined by the developer.
                "margin": "4px", // Outer margin of the container, default value "0", supported range [-99,99]px.
                "horizontal_spacing": "large", // Spacing between components within the container, optional values: "small"(4px), "medium"(8px), "large"(12px), "extra_large"(16px) or [0,99]px. Default 8px.
                "horizontal_align": "left", // Horizontal alignment of components within the container, optional values: "left", "center", "right", default value is "left".
                "flex_mode": "none", // Adaptive mode of each column on narrow screens of mobile and PC. Default value none.
                "background_style": "default", // Background style of the column set. Default value default.
                "action": { // Interaction configuration when clicking the column set.
                    "multi_url": {
                        "url": "https://open.larksuite.com",
                        "pc_url": "https://open.feishu.com",
                        "ios_url": "https://developer.apple.com/",
                        "android_url": "https://developer.android.com/"
                    }
                },
                "columns": [
                    // Column configuration
                    {
                        "tag": "column",
                        "element_id": "custom_id", // Unique identifier for the operation component. Used to specify the component when calling related component interfaces. Custom defined by the developer.
                        "background_style": "default", // Background style of the column. Default value default.
                        "width": "auto", // Width of the column. Default value auto.
                        "weight": 1, // Effective when width is set to weighted, indicating the width proportion of the current column.
                        "horizontal_spacing": "large", // Spacing between components within the column, optional values: "small"(4px), "medium"(8px), "large"(12px), "extra_large"(16px) or [0,99]px. Default 8px.
                        "horizontal_align": "left", // Horizontal alignment of components within the column, optional values: "left", "center", "right", default value is "left".
                        "vertical_align": "center", // Vertical alignment of components within the column, optional values: "top", "center", "bottom", default value is "top".
                        "vertical_spacing": "4px", // Vertical spacing of subcomponents within the column. Default value default (8px).
                        "direction": "vertical", // Arrangement direction of the column. Optional values: "vertical" (vertical arrangement), "horizontal" (horizontal arrangement). Default is "vertical".
                        "padding": "8px", // Inner padding of the column. Default value 0px. Supported range [0,99]px.
                        "margin": "4px", // Outer margin of the column, default value 0px. supported range [-99,99]px.
                        "action": {
                            // Interaction configuration when clicking the column.
                            "multi_url": {
                                "url": "https://www.baidu.com",
                                "pc_url": "https://www.baidu.com",
                                "ios_url": "https://www.google.com",
                                "android_url": "https://www.apple.com.cn"
                            }
                        },
                        "elements": [] // Components nested within the column container, does not support nested table and form containers.
                    }
                ]
            }
        ]
    }
}
Column set fields description

The description of the various attribute fields for the columns within a column set is shown in the following table.

Name	Required	Type	Default	Description


tag

	

Yes

	

String

	

/

	

The tag of the component. The fixed value for column_set component is column_set.




element_id

	

No

	

String

	

Empty

	

The unique identifier for the component. A new attribute added in JSON 2.0. It is used to specify the component when calling component-related interfaces. The value of this field must be globally unique within the same card. Only letters, numbers, and underscores are allowed, and it must start with a letter and not exceed 20 characters.




horizontal_spacing

	

No

	

String

	

8px

	

Horizontal spacing of components within the container. Optional values:

small: small spacing, 4px
medium: medium spacing, 8px
large: large spacing, 12px
extra_large: extra-large spacing, 16px
Specific values, such as 20px. The value range is [0,99]px



horizontal_align

	

No

	

String

	

left

	

Horizontal alignment of components within the container. Optional values:

left: left-aligned
center: center-aligned
right: right-aligned



margin

	

No

	

String

	

0px

	

External margin of the container. The value range is [-99,99]px. Optional values:

Single value, such as "10px", indicates that the four external margins of the container are all 10 px.
Double value, such as "4px 0", indicates that the top and bottom external margins of the container are 4 px, and the left and right external margins are 0 px. Use spaces to separate (units can be omitted when the margin is 0).
Multiple values, such as "4px 0 4px 0", indicate that the top, right, bottom, and left external margins of the container are 4px, 12px, 4px, and 12px, respectively. Use spaces to separate.



flex_mode

	

No

	

String

	

none

	

Adaptive methods for narrow screens on mobile and PC. Values:

none: No layout adaptation, compress column width proportionally on narrow screens.
stretch: Column layout changes to row layout, and each column (row) width is forcibly stretched to 100%, all columns adapt to stack vertically.
flow: Column flow layout (automatic line wrap), when a row cannot display a column, it automatically wraps to the next row.
bisect: Two-column equal layout.
trisect: Three-column equal layout.



background_style

	

No

	

String

	

default

	

Background color style of the column. Possible values:

default: Default white background style, black background in client dark theme.
Color enumeration values supported by the card and custom colors in RGBA syntax. Refer to Color Enumeration Values.

Note: When there is nested column layout, the color of the upper column overrides the color of the lower column.




action

	

No

	

Action

	

/

	

Set the interaction configuration when clicking the column. Currently, only jump interaction is supported. If there are interactive components within the layout container, the interaction defined by the interactive components takes precedence.




└ multi_url

	

No

	

Struct

	

Empty

	

Configure the link addresses for each end.




└└ url

	

No

	

String

	

Empty

	

Fallback jump link.




└└ android_url

	

No

	

String

	

Empty

	

Jump link for Android. Can be configured as lark://msgcard/unsupported_action to declare that jump is not allowed on this end.




└└ ios_url

	

No

	

String

	

Empty

	

Jump link for iOS. Can be configured as lark://msgcard/unsupported_action to declare that jump is not allowed on this end.




└└ pc_url

	

No

	

String

	

Empty

	

Jump link for PC. Can be configured as lark://msgcard/unsupported_action to declare that jump is not allowed on this end.




columns

	

Yes

	

column[]

	

Empty

	

Configuration of columns in the column layout. See below for details.

Column field description

The attributes of each column in the split column are described in the table below.

Name	Required	Type	Default value	Description


tag

	

Yes

	

String

	

/

	

The tag of the column, with a fixed value of column.




element_id

	

No

	

String

	

Empty

	

The unique identifier of the operation component. A new attribute added in JSON 2.0. It is used to specify the component when calling component-related interfaces. The value of this field must be globally unique within the same card. Only letters, numbers, and underscores are allowed, must start with a letter, and must not exceed 20 characters.




background_style

	

No

	

String

	

default

	

The background color style of the column. Possible values:

default: Default white background style, black background style in client dark theme
Color enumeration values supported by the card and custom colors in RGBA syntax. Refer to color enumeration values



width

	

No

	

String

	

auto

	

Column width. This attribute is effective only when flex_mode is none. Possible values:

auto: Column width is consistent with the width of the elements within the column
weighted: Column width is distributed according to the weight defined by the weight parameter
Specific values, such as 100px. The value range is [16,600]px. This enumeration is supported in version V7.4 and above



weight

	

No

	

Number

	

1

	

Effective when the width field is set to weighted, indicating the width proportion of the current column. The value range is an integer between 1 and 5.




horizontal_spacing

	

No

	

String

	

8px

	

Horizontal spacing of components within the column. Optional values:

small: Small spacing, 4px
medium: Medium spacing, 8px
large: Large spacing, 12px
extra_large: Extra large spacing, 16px
Specific values, such as 20px. The value range is [0,99]px



horizontal_align

	

No

	

String

	

left

	

Horizontal alignment of components within the column. Possible values:

left: Left alignment
center: Center alignment
right: Right alignment



vertical_align

	

No

	

String

	

top

	

Vertical alignment of components within the column. Possible values:

top: Top alignment
center: Center alignment
bottom: Bottom alignment



vertical_spacing

	

No

	

String

	

8px

	

Vertical spacing of components within the column. Optional values:

small: Small spacing, 4px
medium: Medium spacing, 8px
large: Large spacing, 12px
extra_large: Extra large spacing, 16px
Specific values, such as 20px. The value range is [0,99]px



direction

	

No

	

String

	

vertical

	

Direction of the column arrangement. Optional values:

vertical: Vertical arrangement
horizontal: Horizontal arrangement



padding

	

No

	

String

	

0px

	

Padding of the column. The value range is [0,99]px. Optional values:

Single value, such as "10px", indicating that all four margins of the column are 10px.
Double value, such as "4px 0", indicating that the top and bottom margins of the column are 4px, and the left and right margins are 0px. Use spaces to separate (units can be omitted when the margin is 0).
Multiple values, such as "4px 0 4px 0", indicating that the top, right, bottom, and left margins of the column are 4px, 12px, 4px, and 12px, respectively. Use spaces to separate.



margin

	

No

	

String

	

0px

	

External margin of the column. The value range is [-99,99]px. Optional values:

Single value, such as "10px", indicates that the four external margins of the container are all 10 px.
Double value, such as "4px 0", indicates that the top and bottom external margins of the container are 4 px, and the left and right external margins are 0 px. Use spaces to separate (units can be omitted when the margin is 0).
Multiple values, such as "4px 0 4px 0", indicate that the top, right, bottom, and left external margins of the container are 4px, 12px, 4px, and 12px, respectively. Use spaces to separate.



elements

	

No

	

Element or ColumnSet[]

	

Empty

	

Components embedded within the column. Refer to the nesting relationship mentioned above for embeddable components.




action

	

No

	

Action

	

/

	

Interaction configuration when clicking the column. Currently, only jump interactions are supported. If there are interactive components within the layout container, the interaction defined by the interactive components takes precedence.




└ multi_url

	

No

	

Struct

	

Empty

	

Configure the link address for each end.




└└ url

	

No

	

String

	

Empty

	

Fallback link address.




└└ android_url

	

No

	

String

	

Empty

	

Link address for the Android end. It can be configured as lark://msgcard/unsupported_action to declare that this end does not allow redirection.




└└ ios_url

	

No

	

String

	

Empty

	

Link address for the iOS end. It can be configured as lark://msgcard/unsupported_action to declare that this end does not allow redirection.




└└ pc_url

	

No

	

String

	

Empty

	

Link address for the PC end. It can be configured as lark://msgcard/unsupported_action to declare that this end does not allow redirection.

Demo example

The following sample code in JSON 2.0 structure can achieve the card effect as shown in the figure below:

{
    "schema": "2.0",
    "body": {
      "elements": [
        {
                "tag": "markdown",
                "content": ":YouAreTheBest:**个人效率总览** ",
                "text_align": "left",
                "text_size": "heading"
            },
            {
                "tag": "column_set",
                "flex_mode": "bisect",
                "horizontal_spacing": "",
                "horizontal_align": "center",
                "columns": [
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "top",
                        "vertical_spacing": "8px",
                        "direction": "horizontal",
                        "elements": [
                            {
                                "tag": "column_set",
                                "flex_mode": "none",
                                "horizontal_spacing": "8px",
                                "horizontal_align": "left",
                                "columns": [
                                    {
                                        "tag": "column",
                                        "width": "weighted",
                                        "vertical_align": "top",
                                        "vertical_spacing": "8px",
                                        "background_style": "grey",
                                        "padding": "8px",
                                        "elements": [
                                            {
                                                "tag": "column_set",
                                                "flex_mode": "none",
                                                "horizontal_spacing": "default",
                                                "background_style": "default",
                                                "columns": [
                                                    {
                                                        "tag": "column",
                                                        "elements": [
                                                            {
                                                                "tag": "div",
                                                                "text": {
                                                                    "tag": "plain_text",
                                                                    "content": "已审批单量",
                                                                    "text_size": "normal",
                                                                    "text_align": "center",
                                                                    "text_color": "grey"
                                                                }
                                                            }
                                                        ],
                                                        "width": "weighted",
                                                        "weight": 1
                                                    }
                                                ]
                                            },
                                            {
                                                "tag": "column_set",
                                                "flex_mode": "none",
                                                "horizontal_spacing": "default",
                                                "background_style": "default",
                                                "columns": [
                                                    {
                                                        "tag": "column",
                                                        "elements": [
                                                            {
                                                                "tag": "div",
                                                                "text": {
                                                                    "tag": "plain_text",
                                                                    "content": "29 单",
                                                                    "text_size": "heading",
                                                                    "text_align": "center",
                                                                    "text_color": "default"
                                                                }
                                                            }
                                                        ],
                                                        "width": "weighted",
                                                        "weight": 1
                                                    }
                                                ]
                                            },
                                            {
                                                "tag": "markdown",
                                                "content": "<text_tag color='blue'>高于部门 86%</text_tag>",
                                                "text_align": "center",
                                                "text_size": "normal"
                                            }
                                        ],
                                        "weight": 1
                                    }
                                ],
                                "margin": "0px 0px 0px 0px"
                            }
                        ],
                        "weight": 1
                    },
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "top",
                        "vertical_spacing": "8px",
                        "elements": [
                            {
                                "tag": "column_set",
                                "flex_mode": "none",
                                "horizontal_spacing": "8px",
                                "horizontal_align": "left",
                                "columns": [
                                    {
                                        "tag": "column",
                                        "width": "weighted",
                                        "vertical_align": "top",
                                        "vertical_spacing": "8px",
                                        "background_style": "grey",
                                        "padding": "8px",
                                        "elements": [
                                            {
                                                "tag": "column_set",
                                                "flex_mode": "none",
                                                "horizontal_spacing": "default",
                                                "background_style": "default",
                                                "columns": [
                                                    {
                                                        "tag": "column",
                                                        "elements": [
                                                            {
                                                                "tag": "div",
                                                                "text": {
                                                                    "tag": "plain_text",
                                                                    "content": "审批平均耗时",
                                                                    "text_size": "normal",
                                                                    "text_align": "center",
                                                                    "text_color": "grey"
                                                                }
                                                            }
                                                        ],
                                                        "width": "weighted",
                                                        "weight": 1
                                                    }
                                                ]
                                            },
                                            {
                                                "tag": "column_set",
                                                "flex_mode": "none",
                                                "horizontal_spacing": "default",
                                                "background_style": "default",
                                                "columns": [
                                                    {
                                                        "tag": "column",
                                                        "elements": [
                                                            {
                                                                "tag": "div",
                                                                "text": {
                                                                    "tag": "plain_text",
                                                                    "content": "0.9 小时",
                                                                    "text_size": "heading",
                                                                    "text_align": "center",
                                                                    "text_color": "default"
                                                                }
                                                            }
                                                        ],
                                                        "width": "weighted",
                                                        "weight": 1
                                                    }
                                                ]
                                            },
                                            {
                                                "tag": "markdown",
                                                "content": "<text_tag color='orange'>落后部门 61%</text_tag>",
                                                "text_align": "center",
                                                "text_size": "normal"
                                            }
                                        ],
                                        "weight": 1
                                    }
                                ],
                                "margin": "0px 0px 0px 0px"
                            }
                        ],
                        "weight": 1
                    },
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "top",
                        "vertical_spacing": "8px",
                        "elements": [
                            {
                                "tag": "column_set",
                                "flex_mode": "none",
                                "horizontal_spacing": "8px",
                                "horizontal_align": "left",
                                "columns": [
                                    {
                                        "tag": "column",
                                        "width": "weighted",
                                        "vertical_align": "top",
                                        "vertical_spacing": "8px",
                                        "background_style": "grey",
                                        "padding": "8px",
                                        "elements": [
                                            {
                                                "tag": "column_set",
                                                "flex_mode": "none",
                                                "horizontal_spacing": "default",
                                                "background_style": "default",
                                                "columns": [
                                                    {
                                                        "tag": "column",
                                                        "elements": [
                                                            {
                                                                "tag": "div",
                                                                "text": {
                                                                    "tag": "plain_text",
                                                                    "content": "代批率",
                                                                    "text_size": "normal",
                                                                    "text_align": "center",
                                                                    "text_color": "grey"
                                                                }
                                                            }
                                                        ],
                                                        "width": "weighted",
                                                        "weight": 1
                                                    }
                                                ]
                                            },
                                            {
                                                "tag": "column_set",
                                                "flex_mode": "none",
                                                "horizontal_spacing": "default",
                                                "background_style": "default",
                                                "columns": [
                                                    {
                                                        "tag": "column",
                                                        "elements": [
                                                            {
                                                                "tag": "div",
                                                                "text": {
                                                                    "tag": "plain_text",
                                                                    "content": "0%",
                                                                    "text_size": "heading",
                                                                    "text_align": "center",
                                                                    "text_color": "default"
                                                                }
                                                            }
                                                        ],
                                                        "width": "weighted",
                                                        "weight": 1
                                                    }
                                                ]
                                            },
                                            {
                                                "tag": "markdown",
                                                "content": "<text_tag color='green'>领先部门 100%</text_tag>",
                                                "text_align": "center",
                                                "text_size": "normal"
                                            }
                                        ],
                                        "weight": 1
                                    }
                                ],
                                "margin": "0px 0px 0px 0px"
                            }
                        ],
                        "weight": 1
                    }
                ],
                "margin": "16px 0px 0px 0px"
            },
            {
                "tag": "markdown",
                "content": ":STRIVE: **待优化的任务类型**",
                "text_align": "left",
                "text_size": "heading"
            },
            {
                "tag": "column_set",
                "flex_mode": "none",
                "horizontal_spacing": "8px",
                "horizontal_align": "left",
                "columns": [
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "center",
                        "vertical_spacing": "4px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<text_tag color='red'>1</text_tag> 加班申请",
                                "text_align": "left",
                                "text_size": "normal"
                            }
                        ],
                        "weight": 1
                    },
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "center",
                        "vertical_spacing": "8px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<font color='grey'>低于部门</font> <font color='red'>95% </font><font color='grey'>的审批人</font> ",
                                "text_align": "right",
                                "text_size": "notation"
                            }
                        ],
                        "weight": 1
                    }
                ],
                "margin": "16px 0px 0px 0px"
            },
            {
                "tag": "column_set",
                "flex_mode": "none",
                "horizontal_spacing": "8px",
                "horizontal_align": "left",
                "columns": [
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "center",
                        "vertical_spacing": "4px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<text_tag color='red'>2</text_tag> 休假申请",
                                "text_align": "left",
                                "text_size": "normal"
                            }
                        ],
                        "weight": 1
                    },
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "center",
                        "vertical_spacing": "4px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<font color='grey'>低于部门</font> <font color='red'>55% </font><font color='grey'>的审批人</font> ",
                                "text_align": "right",
                                "text_size": "notation"
                            }
                        ],
                        "weight": 1
                    }
                ],
                "margin": "16px 0px 0px 0px"
            },
            {
                "tag": "markdown",
                "content": ":CheckMark:**效率高的任务类型**",
                "text_align": "left",
                "text_size": "heading"
            },
            {
                "tag": "column_set",
                "flex_mode": "none",
                "horizontal_spacing": "8px",
                "horizontal_align": "left",
                "columns": [
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "center",
                        "vertical_spacing": "4px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<text_tag color='green'>1</text_tag> 数据权限申请",
                                "text_align": "left",
                                "text_size": "normal"
                            }
                        ],
                        "weight": 1
                    },
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "center",
                        "vertical_spacing": "4px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<font color='grey'>高于部门</font> <font color='green'>68% </font><font color='grey'>的审批人</font> ",
                                "text_align": "right",
                                "text_size": "notation"
                            }
                        ],
                        "weight": 1
                    }
                ],
                "margin": "16px 0px 0px 0px"
            },
            {
                "tag": "column_set",
                "flex_mode": "none",
                "horizontal_spacing": "8px",
                "horizontal_align": "left",
                "columns": [
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "center",
                        "vertical_spacing": "4px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<text_tag color='green'>2</text_tag> BOT推送消息",
                                "text_align": "left",
                                "text_size": "normal"
                            }
                        ],
                        "weight": 1
                    },
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "center",
                        "vertical_spacing": "4px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<font color='grey'>高于部门</font> <font color='green'>56% </font><font color='grey'>的审批人</font> ",
                                "text_align": "right",
                                "text_size": "notation"
                            }
                        ],
                        "weight": 1
                    }
                ],
                "margin": "16px 0px 0px 0px"
            }
        ]
    },
    "i18n_header": {
        "zh_cn": {
            "title": {
                "tag": "plain_text",
                "content": "我的近期审批效率"
            },
            "subtitle": {
                "tag": "plain_text",
                "content": "日期范围：2024.03.01 至 2024.03.31"
            },
            "template": "orange",
            "icon": {
                "tag": "standard_icon",
                "token": "approval_colorful"
            }
        }
    }
}
Previous:component JSON v2.0 overview
Next:Form container
Need help with a problem?
Submit feedback

# Interactive container

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/containers/interactive-container>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentscontainersInteractive container
Interactive container
Copy Page
Last updated on 2025-06-27
The contents of this article
Precautions
Nesting rules
Component attributes
JSON structure
Field description
Callback structure
Example Code

You can embed components within interactive containers based on business needs, flexibly combine multiple interactive containers, and uniformly define the styles and interaction capabilities of multiple interactive containers, achieving various combination effects and rich card interactions.

This document introduces the JSON 2.0 structure of the interactive container. To view the historical JSON 1.0 structure, refer to interactive container.

Precautions

Container components support a maximum of five layers of nested components. It is recommended to avoid nesting multiple layers in interactive containers. Multi-layer nesting compresses the display space and affects the presentation of the card.

Nesting rules

In the Card JSON 2.0 Structure, interactive containers can embed all components except for form containers (form) and table components (table).

Component attributes
JSON structure

The complete JSON 2.0 structure of the interactive container is as follows:

{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "interactive_container", // Tag of the interactive container.
                "width": "fill", // Width of the interactive container. Default value is fill.
                "height": "auto", // Height of the interactive container. Default value is auto.
                "element_id": "custom_id", // Unique identifier of the component. Used to specify the component when calling related interfaces. Customizable by the developer.
                "direction": "vertical", // Arrangement direction of components within the container. Optional values: "vertical" (vertical arrangement), "horizontal" (horizontal arrangement). Default is "vertical".
                "margin": "4px", // Outer margin of the container. Default value is "0", supporting range [-99,99]px.
                "horizontal_spacing": "large", // Spacing between components within the container. Optional values: "small"(4px), "medium"(8px), "large"(12px), "extra_large"(16px) or [0,99]px.
                "horizontal_align": "left", // Horizontal alignment of components within the container. Optional values: "left", "center", "right". Default value is left.
                "vertical_align": "center", // Vertical alignment of components within the container. Optional values: "top", "center", "bottom". Default value is "top".
                "vertical_spacing": "4px", // Vertical spacing between components within the container. Optional values: "small"(4px), "medium"(8px), "large"(12px), "extra_large"(16px) or [0,99]px. Default value is 4px.
                "background_style": "default", // Background color. Default value is default (no background color).
                "has_border": false, // Whether to display the border, with a fixed thickness of 1px. Default value is false.
                "border_color": "grey", // Border color of the interactive container, effective only when has_border is true.
                "corner_radius": "40px", // Corner radius of the interactive container. Optional.
                "padding": "10px 20px 10px 20px", // Padding of the interactive container. Default value is "4px 12px 4px 12px".
                "behaviors": [
                    {
                        "type": "open_url", // Declares the interaction type as a link jump interaction.
                        "default_url": "https://www.baidu.com", // Default jump URL.
                        "android_url": "https://developer.android.com/", // Android jump URL.
                        "ios_url": "lark://msgcard/unsupported_action", // iOS jump URL.
                        "pc_url": "https://www.windows.com" // Desktop jump URL.
                    },
                    {
                        "type": "callback", // Declares the interaction type as a callback interaction to send data to the server.
                        "value": {
                            // Callback interaction data
                            "key": "value"
                        }
                    }
                ],
                "disabled": false,
                "disabled_tips": {
                    "tag": "plain_text",
                    "content": "demo"
                },
                "confirm": {},
                "hover_tips": {
                    "tag": "plain_text",
                    "content": "demo"
                },
                "elements": [] // Subcomponents of the container, supporting embedding all components except for form containers (`form`) and table components (`table`).
            }
        ]
    }
}
Field description

The table below describes the fields of the interactive container.

Name	Required	Type	Default	Description


tag

	

Yes

	

String

	

/

	

Tag of the interactive container. Fixed value is interactive_container.




width

	

No

	

String

	

fill

	

Width of the interactive container. Possible values:
fill: Maximum supported width of the card
auto: Adaptive width
[16,999]px: Custom width, e.g., "20px". Minimum width is 16px




height

	

No

	

String

	

auto

	

Height of the interactive container. Possible values:
auto: Adaptive height
[10,999]px: Custom height, e.g., "20px"




element_id

	

No

	

String

	

None

	

Unique identifier of the component. Newly added in JSON 2.0. Used to specify the component when calling related interfaces. Must be globally unique within the same card. Only letters, numbers, and underscores are allowed, must start with a letter, and must not exceed 20 characters.




margin

	

No

	

String

	

0px

	

Outer margin of the container. Value range is [-99,99]px. Possible values:
Single value, e.g., "10px", means the four outer margins of the container are all 10px.
Double value, e.g., "4px 0", means the top and bottom margins are 4px, and the left and right margins are 0px. Separated by spaces (unit can be omitted if the margin is 0).
Multiple values, e.g., "4px 0 4px 0", means the top, right, bottom, and left margins are 4px, 12px, 4px, and 12px respectively. Separated by spaces.




direction

	

No

	

String

	

vertical

	

Arrangement direction of components within the container. Possible values:
vertical: Vertical arrangement
horizontal: Horizontal arrangement




horizontal_spacing

	

No

	

String

	

8px

	

Horizontal spacing between components within the container. Optional values:

small: small spacing, 4px
medium: medium spacing, 8px
large: large spacing, 12px
extra_large: extra large spacing, 16px
Specific values, such as 20px. The value range is [0,99]px



horizontal_align

	

No

	

String

	

left

	

Horizontal alignment of components within the container. Optional values:

left: left aligned
center: center aligned
right: right aligned



vertical_spacing

	

No

	

String

	

12px

	

Vertical spacing between components within the container. Optional values:

small: small spacing, 4px
medium: medium spacing, 8px
large: large spacing, 12px
extra_large: extra large spacing, 16px
Specific values, such as 20px. The value range is [0,99]px



vertical_align

	

No

	

String

	

top

	

Vertical alignment of components within the container. Optional values:

top: top aligned
center: center aligned
bottom: bottom aligned



background_style

	

No

	

String

	

default

	

Background style of the interactive container. Optional values:

default: default white background style, black background in dark mode on the client
laser: laser gradient color style
Colors supported by the card's color enumeration values and custom colors in RGBA syntax. Refer to color enumeration values



has_border

	

No

	

Boolean

	

false

	

Whether to display a border, with a fixed thickness of 1px.




border_color

	

No

	

String

	

grey

	

Border color, effective only when has_border is true. The enumeration values are the colors supported by the card's color enumeration values and custom colors in RGBA syntax. Refer to color enumeration values.




corner_radius

	

No

	

String

	

0px

	

Corner radius of the interactive container, in pixels (px) or percentage (%). The value should follow the format:

[0,∞]px, such as "10px"
[0,100]%, such as "30%"



padding

	

No

	

String

	

4px, 12px

	

Padding of the container. The value range is [-99,99]px. Optional values:

Single value, such as "10px", which means the four paddings of the container are all 10px.
Double value, such as "4px 0", which means the top and bottom paddings of the container are 4px, and the left and right paddings are 0px. Use a space to separate the values (no unit needed for 0).
Multiple values, such as "4px 0 4px 0", which means the top, right, bottom, and left paddings of the container are 4px, 12px, 4px, and 12px, respectively. Use a space to separate the values.



behaviors

	

Yes

	

[]

	

/

	

Set the interaction configuration when clicking the interactive container. If there are interactive components within the container, the interaction defined by the interactive components takes precedence. Interactive components support callback and open_url interactions. For details, refer to configuring card interactions.




hover_tips

	

No

	

Object

	

empty

	

Tooltip text when the user hovers the cursor over the interactive container on the PC. The default is empty.




└ tag

	

Yes

	

String

	

plain_text

	

The tag of the text. The fixed value is plain_text.




└ content

	

Yes

	

String

	

empty

	

The content of the text.




disabled

	

No

	

Boolean

	

false

	

Whether to disable the interactive container. Optional values:

true: disable the entire container
false: keep the container components available



disabled_tips

	

No

	

Object

	

empty

	

Tooltip text when the user triggers the interaction after the interactive container is disabled. The default is empty, meaning no tooltip.




└ tag

	

Yes

	

String

	

plain_text

	

The tag of the tooltip text. The fixed value is plain_text.




└ content

	

Yes

	

String

	

empty

	

The content of the tooltip text.




confirm

	

No

	

Struct

	

default not enabled

	

Secondary confirmation dialog configuration. This refers to a secondary confirmation dialog prompt that appears when the user submits; the input content is only submitted after the user clicks confirm. This field provides confirm and cancel buttons by default, and you only need to configure the title and content of the dialog.

Note: The confirm field only triggers the secondary confirmation dialog when the user clicks a button with the submit attribute.




└ title

	

Yes

	

Struct

	

/

	

Secondary confirmation dialog title.




└ └ tag

	

Yes

	

String

	

plain_text

	

The tag of the secondary confirmation dialog title text. The fixed value is plain_text.




└ └ content

	

Yes

	

String

	

/

	

The content of the secondary confirmation dialog title.




└ text

	

Yes

	

Struct

	

/

	

The text content of the secondary confirmation dialog.




└ └ tag

	

Yes

	

String

	

plain_text

	

The tag of the secondary confirmation dialog text. The fixed value is plain_text.




└ └ content

	

Yes

	

String

	

/

	

The specific content of the secondary confirmation dialog text.




elements

	

Yes

	

Array<element>

	

[]

	

Components embedded in the interactive container, supporting embedding all components except for form containers (form) and table components (table).

Callback structure

Once components are successfully configured for interaction, when users interact with the components, the callback data will be sent to the developer backend URL configured by you.

If you've added a new version of the card action callback (card.action.trigger), you can refer to Card Callback Communication to understand the callback structure.
If you've added an older version of the card action callback (card.action.trigger_v1), you can refer to Old Message Card Callback Interaction to understand the callback structure.
Example Code

The following JSON example code achieves the card effect shown in the image below:

{
    "schema": "2.0",
    "header": {
        "title": {
            "content": "交互容器示例",
            "tag": "plain_text"
        },
        "icon": {
            "tag": "standard_icon",
            "token": "chat_outlined",
            "color": "orange"
        }
    },
    "body": {
        "elements": [
            {
                "tag": "markdown",
                "content": "在「内容创作」话题下，我可以帮助你进行产品方案、营销文案、工作报告等内容的创作。"
            },
            {
                "tag": "column_set",
                "flex_mode": "none",
                "background_style": "default",
                "columns": [
                    {
                        "tag": "column",
                        "width": "weighted",
                        "weight": 1,
                        "vertical_align": "top",
                        "vertical_spacing": "8px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<font color='grey'>你可以对我说：</font>",
                                "text_size": "notation"
                            },
                            {
                                "tag": "interactive_container",
                                "width": "fill",
                                "height": "auto",
                                "horizontal_align": "left",
                                "background_style": "default",
                                "has_border": true,
                                "border_color": "grey",
                                "corner_radius": "8px",
                                "padding": "4px 12px 4px 12px",
                                "behaviors": [
                                    {
                                        "type": "callback",
                                        "value": {
                                            "key": "value"
                                        }
                                    }
                                ],
                                "elements": [
                                    {
                                        "tag": "markdown",
                                        "content": "帮我生成一篇产品方案的框架",
                                        "icon": {
                                            "tag": "standard_icon",
                                            "token": "frame-selection_outlined",
                                            "color": "orange"
                                        }
                                    }
                                ]
                            },
                            {
                                "tag": "interactive_container",
                                "width": "fill",
                                "height": "auto",
                                "horizontal_align": "left",
                                "background_style": "default",
                                "has_border": true,
                                "border_color": "grey",
                                "corner_radius": "8px",
                                "padding": "4px 12px 4px 12px",
                                "behaviors": [
                                    {
                                        "type": "callback",
                                        "value": {
                                            "key": "value"
                                        }
                                    }
                                ],
                                "disabled": false,
                                "elements": [
                                    {
                                        "tag": "markdown",
                                        "content": "帮我生成一篇产品文案",
                                        "icon": {
                                            "tag": "standard_icon",
                                            "token": "file-link-docx_outlined",
                                            "color": "orange"
                                        }
                                    }
                                ]
                            },
                            {
                                "tag": "interactive_container",
                                "width": "fill",
                                "height": "auto",
                                "horizontal_align": "left",
                                "background_style": "default",
                                "has_border": true,
                                "border_color": "grey",
                                "corner_radius": "8px",
                                "padding": "4px 12px 4px 12px",
                                "behaviors": [
                                    {
                                        "type": "callback",
                                        "value": {
                                            "key": "value"
                                        }
                                    }
                                ],
                                "disabled": false,
                                "elements": [
                                    {
                                        "tag": "markdown",
                                        "content": "帮我写一篇周报",
                                        "icon": {
                                            "tag": "standard_icon",
                                            "token": "pa-calibration-report_outlined",
                                            "color": "orange"
                                        }
                                    }
                                ]
                            }
                        ]
                    }
                ]
            },
            {
                "tag": "column_set",
                "flex_mode": "none",
                "background_style": "default",
                "columns": [
                    {
                        "tag": "column",
                        "width": "weighted",
                        "weight": 1,
                        "vertical_align": "top",
                        "vertical_spacing": "8px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<font color='grey'>或者继续之前的话题</font>",
                                "text_size": "notation"
                            },
                            {
                                "tag": "interactive_container",
                                "width": "fill",
                                "height": "auto",
                                "horizontal_align": "left",
                                "background_style": "default",
                                "has_border": true,
                                "border_color": "grey",
                                "corner_radius": "8px",
                                "padding": "4px 12px 4px 12px",
                                "behaviors": [
                                    {
                                        "type": "callback",
                                        "value": {
                                            "key": "value"
                                        }
                                    }
                                ],
                                "disabled": false,
                                "elements": [
                                    {
                                        "tag": "column_set",
                                        "flex_mode": "none",
                                        "background_style": "default",
                                        "columns": [
                                            {
                                                "tag": "column",
                                                "width": "weighted",
                                                "weight": 1,
                                                "vertical_align": "center",
                                                "elements": [
                                                    {
                                                        "tag": "markdown",
                                                        "content": "内容创作:创作暑假营销活动文案",
                                                        "icon": {
                                                            "tag": "standard_icon",
                                                            "token": "chat-history_outlined"
                                                        }
                                                    }
                                                ]
                                            },
                                            {
                                                "tag": "column",
                                                "width": "auto",
                                                "weight": 1,
                                                "vertical_align": "center",
                                                "elements": [
                                                    {
                                                        "tag": "markdown",
                                                        "content": "<font color='grey'>昨天</font>",
                                                        "text_size": "notation"
                                                    }
                                                ]
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "tag": "interactive_container",
                                "width": "fill",
                                "height": "auto",
                                "horizontal_align": "left",
                                "background_style": "default",
                                "has_border": true,
                                "border_color": "grey",
                                "corner_radius": "8px",
                                "padding": "4px 12px 4px 12px",
                                "behaviors": [
                                    {
                                        "type": "callback",
                                        "value": {
                                            "key": "value"
                                        }
                                    }
                                ],
                                "disabled": false,
                                "elements": [
                                    {
                                        "tag": "column_set",
                                        "flex_mode": "none",
                                        "background_style": "default",
                                        "columns": [
                                            {
                                                "tag": "column",
                                                "width": "weighted",
                                                "weight": 1,
                                                "vertical_align": "center",
                                                "elements": [
                                                    {
                                                        "tag": "markdown",
                                                        "content": "内容创作:生成了季度工作报告",
                                                        "icon": {
                                                            "tag": "standard_icon",
                                                            "token": "chat-history_outlined"
                                                        }
                                                    }
                                                ]
                                            },
                                            {
                                                "tag": "column",
                                                "width": "auto",
                                                "weight": 1,
                                                "vertical_align": "center",
                                                "elements": [
                                                    {
                                                        "tag": "markdown",
                                                        "content": "<font color='grey'>上周</font>",
                                                        "text_size": "notation"
                                                    }
                                                ]
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "tag": "interactive_container",
                                "width": "fill",
                                "height": "auto",
                                "horizontal_align": "left",
                                "background_style": "default",
                                "has_border": true,
                                "border_color": "grey",
                                "corner_radius": "8px",
                                "padding": "4px 12px 4px 12px",
                                "behaviors": [
                                    {
                                        "type": "callback",
                                        "value": {
                                            "key": "value"
                                        }
                                    }
                                ],
                                "disabled": false,
                                "elements": [
                                    {
                                        "tag": "markdown",
                                        "content": "更多历史话题",
                                        "icon": {
                                            "tag": "standard_icon",
                                            "token": "chat-history_outlined"
                                        }
                                    }
                                ]
                            }
                        ]
                    }
                ]
            },
            {
                "tag": "column_set",
                "flex_mode": "none",
                "background_style": "default",
                "columns": [
                    {
                        "tag": "column",
                        "width": "weighted",
                        "weight": 1,
                        "vertical_align": "top",
                        "vertical_spacing": "8px",
                        "elements": [
                            {
                                "tag": "div",
                                "text": {
                                    "tag": "plain_text",
                                    "text_size": "notation",
                                    "content": "本话题中已选择以下插件",
                                    "text_color": "grey"
                                }
                            },
                            {
                                "tag": "interactive_container",
                                "width": "auto",
                                "height": "auto",
                                "horizontal_align": "left",
                                "background_style": "grey",
                                "has_border": false,
                                "border_color": "grey",
                                "corner_radius": "40px",
                                "padding": "2px 8px 2px 4px",
                                "behaviors": [
                                    {
                                        "type": "callback",
                                        "value": {
                                            "key": "value"
                                        }
                                    }
                                ],
                                "disabled": false,
                                "elements": [
                                    {
                                        "tag": "column_set",
                                        "flex_mode": "none",
                                        "background_style": "default",
                                        "horizontal_spacing": "4px",
                                        "columns": [
                                            {
                                                "tag": "column",
                                                "width": "auto",
                                                "weight": 1,
                                                "vertical_align": "center",
                                                "vertical_spacing": "8px",
                                                "elements": [
                                                    {
                                                        "tag": "img",
                                                        "img_key": "img_v2_58e37110-6878-44ee-bce4-7a571c1bb70g",
                                                        "transparent": true,
                                                        "scale_type": "crop_center",
                                                        "size": "18px 18px",
                                                        "preview": false
                                                    }
                                                ]
                                            },
                                            {
                                                "tag": "column",
                                                "width": "weighted",
                                                "weight": 1,
                                                "vertical_align": "center",
                                                "vertical_spacing": "8px",
                                                "elements": [
                                                    {
                                                        "tag": "markdown",
                                                        "content": "妙记插件"
                                                    }
                                                ]
                                            }
                                        ]
                                    }
                                ]
                            }
                        ]
                    }
                ]
            }
        ]
    }
}
Previous:Form container
Next:Collapsible panel
Need help with a problem?
Submit feedback

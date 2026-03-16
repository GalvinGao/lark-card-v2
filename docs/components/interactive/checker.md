# Checker

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/checker>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsInteractive componentsChecker
Checker
Copy Page
Last updated on 2025-06-27
The contents of this article
Precautions
Nesting rules
Component properties
JSON structure
Field description
Callback Structure
Example code

The checker is an interactive component that supports configurable callback responses, mainly used in task-checking scenarios.

This document introduces the JSON 2.0 structure of the checker component. To view the historical JSON 1.0 structure, refer to checker.

Precautions

The checker can only be used by writing card JSON code; it is not yet supported for use in the card building tool.

Nesting rules

The checker component supports nesting within all container components (including form containers, interactive containers, columns, and collapsible panels).

Component properties
JSON structure

The following is the card JSON 2.0 data for the checker component:

{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "checker", // Component tag. The fixed value for the checker component is checker.
                "element_id": "custom_id", // Unique identifier for the operation component. Used to specify the component in related interface calls. Needs to be customized by the developer.
                "name": "check_1", // Unique identifier for the checker component. Used to identify which component the user-submitted data belongs to.
                "checked": false, // Initial checked state of the checker. Default value is false.
                "text": { // Plain text information within the checker component.
                    "tag": "plain_text", // Text type tag.
                    "content": "", // Text content. When the tag is lark_md, it supports text content with partial Markdown syntax.
                    "text_size": "normal", // Text size. Default value is normal.
                    "text_color": "default", // Text color. Only effective when the tag is plain_text. Default value is default.
                    "text_align": "left" // Text alignment. Default value is left.
                },
                "overall_checkable": true, // Whether the checker has a shadow effect when the cursor hovers over it. Default value is true.
                "button_area": { // Configuration of the button area. Optional.
                    "pc_display_rule": "always", // Display rule for the button within the checker on the PC end. Default value is always, meaning the button is always displayed.
                    "buttons": [ // Add and configure buttons within the checker. Up to three buttons can be configured.
                        {
                            "tag": "button", // Button tag, fixed value is button.
                            "type": "text", // Button type. Required.
                            "size": "small", // Button size. Default value is medium.
                            "text": { // Text on the button.
                                "tag": "plain_text",
                                "content": "text button"
                            },
                            "icon": { // Add an icon as a prefix to the button text. Supports custom icons or icons from the icon library.
                                "tag": "standard_icon", // Icon type.
                                "token": "chat-forbidden_outlined", // Token of the icon. Only effective when the tag is standard_icon.
                                "color": "orange", // Icon color. Only effective when the tag is standard_icon.
                                "img_key": "img_v2_38811724" // Key of the image. Only effective when the tag is custom_icon.
                            },
                            "disabled": false,
                            "behaviors": []
                        }
                    ]
                },
                "checked_style": { // Checked state style.
                    "show_strikethrough": true, // Whether to show a strikethrough line across the content area. Default value is false.
                    "opacity": 1 // Opacity of the content area. Default value is 1.
                },
                "margin": "0px", // Overall margin of the component, supports single or multiple values. Default value is 0px.
                "padding": "0px", // Overall padding of the component, supports single or multiple values. Default value is 0px.
                "confirm": {
                    // Secondary confirmation popup configuration.
                    "title": {
                        "tag": "plain_text",
                        "content": "title"
                    },
                    "text": {
                        "tag": "plain_text",
                        "content": "content"
                    }
                }, // Secondary confirmation popup configuration. The interaction declared in behaviors is executed after the user clicks confirm.
                "behaviors": [ // Configure interaction types and specific interaction behaviors. When behaviors are not configured, the terminal user can check, but it is only effective locally.
                    {
                        "type": "callback", // Declare the interaction type. Only supports callback request interaction.
                        "value": {
                            // Callback interaction data
                            "key": "value"
                        }
                    }
                ],
                "hover_tips": {}, // Text reminder when the user hovers the cursor over the checker on the PC end.
                "disabled": false, // Whether to disable the checker. Default value is false.
                "disabled_tips": {} // Text reminder when the user hovers the cursor over the checker on the PC end after it is disabled.
            }
        ]
    }
}
Field description

The description of each field in the checker component is as follows.

Field	Required	Type	Default Value	Description


tag

	

Yes

	

String

	

/

	

The tag of the component. The fixed value for the checker component is checker.




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



padding

	

No

	

String

	

0

	

The padding of the component. New attribute in JSON 2.0. The value range is [-99,99]px. Optional values:- Single value, such as "10px", means the padding on all four sides of the component is 10 px.- Double value, such as "4px 0", means the padding on the top and bottom is 4 px, and the padding on the left and right is 0 px. Use space to separate values (unit can be omitted when the padding is 0).- Multiple values, such as "4px 0 4px 0", means the padding on the top, right, bottom, and left are 4px, 12px, 4px, and 12px respectively. Use space to separate values.




name

	

No

	

String

	

None

	

A unique identifier for the checker component. Used to identify which component the user-submitted data belongs to.

Note: When the checker component is nested within a form container, this field is required and must be unique within the card.




checked

	

No

	

Boolean

	

false

	

The initial check state of the checker. Optional values:

true: Checked
false: Unchecked



text

	

No

	

Object

	

/

	

Ordinary text information within the checker component.




└ tag

	

Yes

	

String

	

plain_text

	

Text type tag. Optional values:

plain_text: Plain text content
lark_md: Text content that supports partial Markdown syntax. See lark_md Supported Markdown Syntax

Note: In the Lark card building tool, only the plain_text type of ordinary text component is supported. You can use a rich text component to add Markdown formatted text.




└ content

	

Yes

	

String

	

/

	

Text content. When tag is lark_md, the text content supports partial Markdown syntax. See lark_md Supported Markdown Syntax




└ text_size

	

No

	

String

	

normal

	

Text size. Optional values:

normal: Normal (14px)
heading: Heading (16px)
notation: Notation (12px)



└ text_color

	

No

	

String

	

default

	

Text color. Only effective when tag is plain_text. Optional values:

default: Black in light theme mode; white in dark theme mode.
Color enumeration value. See Color Enumeration Values



└ text_align

	

No

	

String

	

left

	

Text alignment. Optional values:

left: Left aligned
center: Center aligned
right: Right aligned



overall_checkable

	

No

	

Boolean

	

true

	

Whether the entire checker has a shadow effect when the cursor hovers over it.

Note: To cancel the shadow effect, you need to ensure overall_checkable is false and pc_display_rule is not on_hover.




button_area

	

No

	

Object

	

/

	

Button area configuration.




└ pc_display_rule

	

No

	

String

	

always

	

Button display rule in the checker on PC. Buttons are always displayed on mobile. Optional values:

always: Buttons are always displayed.
on_hover: Buttons are displayed and the checker has a shadow effect when the cursor hovers over it.



└ buttons

	

No

	

Array<Object>

	

[]

	

Add and configure buttons in the checker. Up to three buttons can be configured. See the next section for the description of the buttons field.




checked_style

	

No

	

Object

	

/

	

Check state style.




└ show_strikethrough

	

No

	

Boolean

	

false

	

Whether to show a strikethrough line in the content area.




└ opacity

	

No

	

Number

	

1

	

Opacity of the content area. The value range is a number between [0,1], with no decimal places restricted.




margin

	

No

	

String

	

0px

	

Outer margin of the component as a whole, supporting single or multiple values:

Single value: such as "4px", indicating that all four outer margins of the component are 4px
Multiple values: such as "4px 12px 4px 12px", indicating the inner margins at the top, right, bottom, and left of the container are 4px, 12px, 4px, and 12px, respectively. Four values are required, separated by spaces



padding

	

No

	

String

	

0px

	

Inner padding of the component as a whole, supporting single or multiple values:

Single value: such as "4px", indicating that all four inner paddings of the component are 4px
Multiple values: such as "4px 12px 4px 12px", indicating the inner paddings at the top, right, bottom, and left of the container are 4px, 12px, 4px, and 12px, respectively. Four values are required, separated by spaces



confirm

	

No

	

Struct

	

This attribute is disabled by default.

	

Secondary confirmation popup configuration. A secondary confirmation popup will appear when the user submits; only when the user clicks confirm will the input content be submitted. The field by default provides confirm and cancel buttons. You only need to configure the title and content of the popup.

Note: The confirm field will only trigger the secondary confirmation popup when the user clicks the button with the submit attribute.




confirm.title

	

Yes

	

Struct

	

/

	

Secondary confirmation popup title.




confirm.title.tag

	

Yes

	

String

	

plain_text

	

The tag for the title text of the secondary confirmation popup. The fixed value is plain_text.




confirm.title.content

	

Yes

	

String

	

/

	

The content of the secondary confirmation popup title.




confirm.text

	

Yes

	

Struct

	

/

	

The text content of the secondary confirmation popup.




confirm.text.tag

	

Yes

	

String

	

plain_text

	

The tag for the text of the secondary confirmation popup. The fixed value is plain_text.




confirm.text.content

	

Yes

	

String

	

/

	

The specific content of the secondary confirmation popup text.




behaviors

	

Yes

	

Struct

	

/

	

Configure interaction types and specific interaction behaviors. If behaviors is not configured, the end-user can select but it will only be effective locally. For details, refer to Configuring Card Interactions.




hover_tips

	

No

	

Object

	

Empty

	

Tooltip text that appears when the user hovers the cursor over the checker on the PC.

Note: When both hover_tips and disabled_tips are configured, disabled_tips will take effect.




└ tag

	

No

	

String

	

plain_text

	

The tag for the tooltip text. The fixed value is plain_text.




└ content

	

No

	

String

	

Empty

	

The content of the tooltip text.




disabled

	

No

	

Boolean

	

false

	

Whether to disable the checker. Optional values:

true: Disable
false: The checker component remains available



disabled_tips

	

No

	

Object

	

Empty

	

Tooltip text that appears when the user hovers the cursor over the disabled checker on the PC.




└ tag

	

Yes

	

String

	

plain_text

	

The tag for the disabled tooltip text. The fixed value is plain_text.




└ content

	

Yes

	

String

	

Empty

	

The content of the disabled tooltip text.

buttons Field Description

You can add and configure buttons in the checker using the buttons field. A maximum of three buttons can be configured.

Field Name	Required	Type	Default Value	Description


tag

	

Yes

	

String

	

button

	

The tag of the button, with a fixed value of button.




type

	

Yes

	

String

	

Empty

	

The type of button, optional values:

text: Black font button, no border
primary_text: Blue font button, no border
danger_text: Red font button, no border



size

	

No

	

String

	

medium

	

The size of the button, optional values:

tiny: Extra small size, 24px on PC; 28px on mobile
small: Small size, 28px on PC; 28px on mobile
medium: Medium size, 32px on PC; 36px on mobile
large: Large size, 40px on PC; 48px on mobile



text

	

No

	

Struct

	

Empty

	

Text on the button.




└ tag

	

Yes

	

String

	

plain_text

	

The tag of the text. The fixed value is plain_text.




└ content

	

Yes

	

String

	

Empty

	

The content of the text.




icon

	

No

	

Object

	

/

	

Add an icon as a prefix to the text. Supports custom icons or icons from the icon library.




└ tag

	

No

	

String

	

/

	

The tag of the icon type. Optional values:

standard_icon: Use icons from the icon library.
custom_icon: Use custom images as icons.



└ token

	

No

	

String

	

/

	

The token of the icon from the icon library. Takes effect when tag is standard_icon. Refer to the Icon Library for enumeration values.




└ color

	

No

	

String

	

/

	

The color of the icon. Supports setting the color for both outline and filled icons (icons ending with outlined or filled). Takes effect when tag is standard_icon. Refer to the Color Enumeration Values for enumeration values.




└ img_key

	

No

	

String

	

/

	

Custom prefix icon image key. Takes effect when tag is custom_icon.To obtain the icon key: Call the Upload Image API to upload an image for sending messages and get the image_key from the return value.




disabled

	

No

	

Boolean

	

false

	

Whether to disable the button. Optional values:

true: Disable the button
false: The button component remains available



behaviors

	

Yes

	

Struct

	

/

	

Configure interaction types and specific interaction behaviors. For details, refer to Configuring Card Interactions.

Callback Structure

After successfully configuring interactions for the components, the configured request address in the developer backend will receive callback data when the user interacts with the components.

If you have added the new card callback interaction (card.action.trigger), refer to Card Callback Communication to understand the callback structure.
If you have added the old card callback interaction (card.action.trigger_v1), refer to Message Card Callback (Old) to understand the callback structure.
Example code

The following JSON 2.0 example code can achieve the card effect as shown in the figure:

{
    "schema": "2.0",
    "header": {
        "template": "blue",
        "title": {
            "tag": "plain_text",
            "content": "勾选组件（依赖端版本 7.9+)"
        }
    },
    "body": {
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
                        "vertical_spacing": "1px",
                        "elements": [
                            {
                                "tag": "checker",
                                "name": "check_1",
                                "checked": false,
                                "text": {
                                    "tag": "lark_md",
                                    "content": "完成新品上市计划报告 💬[战略研讨会](https://open.larksuite.com)"
                                },
                                "overall_checkable": false,
                                "button_area": {
                                    "pc_display_rule": "always",
                                    "buttons": [
                                        {
                                            "tag": "button",
                                            "type": "text",
                                            "size": "large",
                                            "text": {
                                                "tag": "plain_text",
                                                "content": ""
                                            },
                                            "icon": {
                                                "tag": "standard_icon",
                                                "token": "forward-com_outlined",
                                                "color": "grey-500"
                                            },
                                            "disabled": false,
                                            "behaviors": [
                                                {
                                                    "type": "callback",
                                                    "value": {
                                                        "key": "btn1"
                                                    }
                                                }
                                            ]
                                        },
                                        {
                                            "tag": "button",
                                            "type": "text",
                                            "size": "large",
                                            "text": {
                                                "tag": "plain_text",
                                                "content": ""
                                            },
                                            "icon": {
                                                "tag": "standard_icon",
                                                "token": "tab-todo_outlined",
                                                "color": "grey-500"
                                            },
                                            "disabled": false,
                                            "behaviors": [
                                                {
                                                    "type": "open_url",
                                                    "default_url": "https://www.baidu.com",
                                                    "android_url": "https://developer.android.com/",
                                                    "ios_url": "lark://msgcard/unsupported_action",
                                                    "pc_url": "https://www.windows.com"
                                                }
                                            ]
                                        }
                                    ]
                                },
                                "checked_style": {
                                    "show_strikethrough": true,
                                    "opacity": 0.5
                                },
                                "padding": "2px 2px 2px 2px",
                                "behaviors": [
                                    {
                                        "type": "callback",
                                        "value": {
                                            "key": "todo1"
                                        }
                                    }
                                ]
                            },
                            {
                                "tag": "checker",
                                "name": "check_2",
                                "checked": false,
                                "text": {
                                    "tag": "lark_md",
                                    "content": "把材料提前给💬[业务数据共享群](https://open.larksuite.com)审阅"
                                },
                                "overall_checkable": true,
                                "button_area": {
                                    "pc_display_rule": "on_hover",
                                    "buttons": [
                                        {
                                            "tag": "button",
                                            "type": "text",
                                            "size": "large",
                                            "text": {
                                                "tag": "plain_text",
                                                "content": ""
                                            },
                                            "icon": {
                                                "tag": "standard_icon",
                                                "token": "forward-com_outlined",
                                                "color": "grey-500"
                                            },
                                            "disabled": false,
                                            "behaviors": [
                                                {
                                                    "type": "callback",
                                                    "value": {
                                                        "key": "btn2"
                                                    }
                                                }
                                            ]
                                        },
                                        {
                                            "tag": "button",
                                            "type": "text",
                                            "size": "large",
                                            "text": {
                                                "tag": "plain_text",
                                                "content": ""
                                            },
                                            "icon": {
                                                "tag": "standard_icon",
                                                "token": "tab-todo_outlined",
                                                "color": "grey-500"
                                            },
                                            "disabled": false,
                                            "behaviors": [
                                                {
                                                    "type": "open_url",
                                                    "default_url": "https://www.baidu.com",
                                                    "android_url": "https://developer.android.com/",
                                                    "ios_url": "lark://msgcard/unsupported_action",
                                                    "pc_url": "https://www.windows.com"
                                                }
                                            ]
                                        }
                                    ]
                                },
                                "checked_style": {
                                    "show_strikethrough": true,
                                    "opacity": 0.5
                                },
                                "padding": "2px 2px 2px 2px",
                                "confirm": {
                                    "title": {
                                        "tag": "plain_text",
                                        "content": "弹窗标题"
                                    },
                                    "text": {
                                        "tag": "plain_text",
                                        "content": "确认提交吗"
                                    }
                                },
                                "behaviors": [
                                    {
                                        "type": "callback",
                                        "value": {
                                            "key": "todo2"
                                        }
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
Previous:Image picker
Next:Card JSON 1.0 structure
Need help with a problem?
Submit feedback

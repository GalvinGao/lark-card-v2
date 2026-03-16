# Button

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/button>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsInteractive componentsButton
Button
Copy Page
Last updated on 2025-06-27
The contents of this article
Notes
Nesting rules
Component properties
JSON structure
Field descriptions
Callback structure
Sample code

The button component is an interactive component that supports multiple styles and sizes, and it supports adding icons as prefix icons.

This document introduces the JSON 2.0 structure of the button component. To view the historical JSON 1.0 structure, refer to Button.

Notes

Card JSON 2.0 structure no longer supports attributes related to interactive modules ("tag": "action"). You can directly place the button in elements and configure appropriate component spacing (vertical_spacing and horizontal_spacing) for use.

Nesting rules

In the JSON 2.0 structure, the button component supports nesting within columns, form containers, folding panels, loop containers, and interactive containers.

Component properties
JSON structure

The complete JSON 2.0 structure of the button component is as follows:

{
  "schema": "2.0", // Version of the card JSON structure. Default is 1.0. To use JSON 2.0 structure, you must explicitly declare 2.0.
  "body": {
    "elements": [
      {
        "tag": "button", // Tag of the component. The fixed value for the button component is button.
        "element_id": "custom_id", // Unique identifier for the operation component. Used to specify the component in related interface calls. Customizable by the developer.
        "margin": "0px 0px 0px 0px", // Margin of the component. Default value is "0", supported range is [-99,99]px.
        "type": "primary", // Type of the button. Default is default.
        "size": "small", // Size of the button. Default value is medium.
        "width": "default", // Width of the button. Default is default.
        "text": {
          // Text on the button.
          "tag": "plain_text",
          "content": "Confirm"
        },
        "icon": {
          // Prefix icon.
          "tag": "standard_icon", // Type of the icon.
          "token": "chat-forbidden_outlined", // Token of the icon. Effective only when tag is standard_icon.
          "color": "orange", // Color of the icon. Effective only when tag is standard_icon.
          "img_key": "img_v2_38811724" // Key of the image. Effective only when tag is custom_icon.
        },
        "hover_tips": {}, // Tooltip text when the user hovers the cursor over the button on PC. Default is empty.
        "disabled": false, // Whether to disable the button. Default value is false.
        "disabled_tips": {}, // Tooltip text when the button is disabled and the user hovers the cursor over the button on PC. When this field is effective, hover_tips is no longer effective.
        "confirm": {
          // Secondary confirmation popup configuration
          "title": {
            "tag": "plain_text",
            "content": "title"
          },
          "text": {
            "tag": "plain_text",
            "content": "content"
          }
        },
        "behaviors": [
          {
            "type": "open_url", // Declares that the interaction type is a URL redirection.
            "default_url": "https://www.baidu.com", // Fallback URL.
            "android_url": "https://developer.android.com/", // Android URL.
            "ios_url": "lark://msgcard/unsupported_action", // iOS URL. Can be configured as `lark://msgcard/unsupported_action` to declare that redirection is not allowed on this end.
            "pc_url": "https://www.windows.com" // Desktop URL.
          },
          {
            "type": "callback", // Declares that the interaction type is a callback to the server.
            "value": {
              // Callback interaction data. Supports object data type.
              "key": "value"
            }
          }
        ],
        // Historical attributes
        "url": "https://open.larksuite.com",
        "multi_url": {
          "android_url": "https://open.larksuite.com",
          "ios_url": "https://open.larksuite.com",
          "pc_url": "https://open.larksuite.com"
        },
        "value": {
          "key_1": "value_1"
        }
      }
    ]
  }
}
Field descriptions

The descriptions of each field of the button component are shown in the table below:

Field Name	Required	Type	Default Value	Description


tag

	

Yes

	

String

	

/

	

The tag of the component. The fixed value for the button component is button.




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



type

	

No

	

String

	

default

	

The type of the button. Options include:

default: black font button with border
primary: blue font button with border
danger: red font button with border
text: black font button without border
primary_text: blue font button without border
danger_text: red font button without border
primary_filled: blue background white font button
danger_filled: red background white font button
laser: laser button



size

	

No

	

String

	

medium

	

The size of the button. Options include:

tiny: ultra-small size, 24 px for PC and 28 px for mobile
small: small size, 28 px for PC and 28 px for mobile
medium: medium size, 32 px for PC and 36 px for mobile
large: large size, 40 px for PC and 48 px for mobile



width

	

No

	

String

	

default

	

The width of the button. Supports the following enumerated values:

default: default width
fill: maximum supported width of the card
[100,∞)px: custom width, such as 120px. When exceeding the width of the card, it will be displayed according to the maximum supported width



text

	

No

	

Struct

	

/

	

The text on the button. Supports plain text and custom text.




icon

	

No

	

Struct

	

/

	

Prefix icon. Supports standard icons and custom icons.




hover_tips

	

No

	

Struct

	

/

	

Text reminder when the user hovers over the button on the PC side. Default is empty.




disabled

	

No

	

Boolean

	

false

	

Whether to disable the button. Default is false.




disabled_tips

	

No

	

Struct

	

/

	

Text reminder when the button is disabled and the user hovers over the button on the PC side. When this field is effective, hover_tips is no longer effective.




confirm

	

No

	

Struct

	

/

	

Double confirmation dialog box configuration.

Note:

To configure the confirm popup, the title field is required. Otherwise, the historical version of the Lark client may have the problem of unresponsive button clicks.




behaviors

	

No

	

Array

	

/

	

Interaction behavior. Supports multiple interaction types, including open_url and callback.

Buttons embedded in form containers have added name and form_action_type attributes. Detailed explanations are shown in the table below.

Attribute Name	Required	Type	Default Value	Description


name

	

Yes

	

String

	

Empty

	

The unique identifier for components within the form container. Used to identify which component the submitted data belongs to by the user.


Note: This field is required and must be unique within the card globally.




form_action_type

	

Yes

	

String

	

Empty

	

The interaction type of the button embedded in the form container. Enumeration values include:

submit: Binds the current button with the submit event. When the user clicks, it triggers the submission event of the form container and asynchronously submits all filled form items' contents
reset: Binds the current button with the cancel submission event. When the user clicks, it triggers the cancellation submission event of the form container and resets the input values of all form components to their initial values
Callback structure

After successfully configuring interaction for the button component, when users interact with the button component, the callback data will be sent to the request address configured in the developer console.

If you have added the callback for the new version card action trigger (card.action.trigger), you can refer to Card Callback Interaction for the callback structure.
If you have added the callback for the old version card action trigger (card.action.trigger_v1), you can refer to Message Card Callback Interaction (Old) for the callback structure.
Sample code

The following JSON 2.0 sample code can achieve the button effect as shown in the figure.

{
  "schema": "2.0",
  "header": {
    "template": "blue",
    "title": {
      "content": "Buttons",
      "tag": "plain_text"
    }
  },
  "body": {
    "elements": [
      {
        "tag": "column_set",
        "flex_mode": "flow",
        "background_style": "default",
        "columns": [
          {
            "tag": "column",
            "width": "auto",
            "weight": 1,
            "vertical_align": "top",
            "elements": [
              {
                "tag": "button",
                "text": {
                  "tag": "plain_text",
                  "content": "镭射按钮"
                },
                "behaviors": [
                  {
                    "type": "open_url",
                    "default_url": "https://open.larksuite.com/document",
                    "android_url": "https://developer.android.com/",
                    "ios_url": "lark://msgcard/unsupported_action",
                    "pc_url": "https://www.windows.com"
                  }
                ],
                "type": "laser",
                "hover_tips": {
                  "tag": "plain_text",
                  "content": "hover提示"
                },
                "value": {
                  "key": "value"
                }
              }
            ]
          },
          {
            "tag": "column",
            "width": "auto",
            "weight": 1,
            "vertical_align": "top",
            "elements": [
              {
                "tag": "button",
                "type": "laser",
                "text": {
                  "tag": "plain_text",
                  "content": "镭射禁用按钮"
                },
                "disabled": true,
                "disabled_tips": {
                  "tag": "plain_text",
                  "content": "禁用 hover 提示"
                },
                "behaviors": [
                  {
                    "type": "open_url",
                    "default_url": "https://open.larksuite.com/document",
                    "android_url": "https://developer.android.com/",
                    "ios_url": "lark://msgcard/unsupported_action",
                    "pc_url": "https://www.windows.com"
                  }
                ]
              }
            ]
          }
        ]
      },
      {
        "tag": "column_set",
        "flex_mode": "flow",
        "background_style": "default",
        "columns": [
          {
            "tag": "column",
            "width": "auto",
            "weight": 1,
            "vertical_align": "top",
            "elements": [
              {
                "tag": "button",
                "text": {
                  "tag": "plain_text",
                  "content": "primary"
                },
                "url": "https://open.larksuite.com/document",
                "type": "primary",
                "hover_tips": {
                  "tag": "plain_text",
                  "content": "我是 primary button"
                },
                "value": {
                  "key": "value"
                }
              }
            ]
          },
          {
            "tag": "column",
            "width": "auto",
            "weight": 1,
            "vertical_align": "top",
            "elements": [
              {
                "tag": "button",
                "type": "default",
                "text": {
                  "tag": "plain_text",
                  "content": "default"
                },
                "hover_tips": {
                  "tag": "plain_text",
                  "content": "我是 default 按钮"
                },
                "behaviors": [
                  {
                    "type": "open_url",
                    "default_url": "https://open.larksuite.com/document",
                    "android_url": "https://developer.android.com/",
                    "ios_url": "lark://msgcard/unsupported_action",
                    "pc_url": "https://www.windows.com"
                  }
                ]
              }
            ]
          },
          {
            "tag": "column",
            "width": "auto",
            "weight": 1,
            "vertical_align": "top",
            "elements": [
              {
                "tag": "button",
                "type": "danger",
                "text": {
                  "tag": "plain_text",
                  "content": "我是 danger 按钮"
                },
                "hover_tips": {
                  "tag": "plain_text",
                  "content": "我是 danger 按钮"
                },
                "behaviors": [
                  {
                    "type": "open_url",
                    "default_url": "https://open.larksuite.com/document",
                    "android_url": "https://developer.android.com/",
                    "ios_url": "lark://msgcard/unsupported_action",
                    "pc_url": "https://www.windows.com"
                  }
                ]
              }
            ]
          },
          {
            "tag": "column",
            "width": "auto",
            "weight": 1,
            "vertical_align": "top",
            "elements": [
              {
                "tag": "button",
                "type": "danger",
                "text": {
                  "tag": "plain_text",
                  "content": "我是 disabled 按钮"
                },
                "disabled": true,
                "disabled_tips": {
                  "tag": "plain_text",
                  "content": "我是 disabled 按钮，我被禁用了"
                },
                "behaviors": [
                  {
                    "type": "open_url",
                    "default_url": "https://open.larksuite.com/document",
                    "android_url": "https://developer.android.com/",
                    "ios_url": "lark://msgcard/unsupported_action",
                    "pc_url": "https://www.windows.com"
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
Previous:Input
Next:Overflow
Need help with a problem?
Submit feedback

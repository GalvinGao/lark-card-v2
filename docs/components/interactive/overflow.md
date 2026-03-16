# Overflow

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/overflow>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsInteractive componentsOverflow
Overflow
Copy Page
Last updated on 2025-06-27
The contents of this article
Nesting rules
Component properties
JSON structure
Field descriptions
Callback Structure
Sample code

The folding button group supports adding multiple buttons and folding them. Clicking on the button group will display all buttons within the group. This document introduces the JSON structure and related properties of the folding button group component.

This document introduces the JSON 2.0 structure and related attributes of the folding button group component. To understand the 1.0 structure and attributes, refer to Folding Button Group.

Nesting rules

In the JSON 2.0 structure, the folding button group component supports nesting within form containers, folding panels, loop containers, interactive containers, and column components. In the building tool, the folding button group component does not support nesting within interactive containers.

Component properties
JSON structure

The complete JSON 2.0 structure of the folding button group component is as follows:

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "overflow",
        "element_id": "custom_id", // The unique identifier for the operation component. It is used to specify the component when calling related interfaces. It needs to be customized by the developer.
        "margin": "0px 0px 0px 0px", // The margin of the component, default value is "0", with a supported range of [-99,99]px.
        "width": "fill", // The width of the folding button group. The default value is default.
        "options": [
          // Add option buttons in the folding button group here.
          { // Add text to the button.
            "text": {
              "tag": "plain_text", // The tag of the text. The fixed value is plain_text.
              "content": "This is a link jump" // The content of the text, supporting up to 100 characters.
            },
            "multi_url": { // Add a jump link to the button.
              "url": "/ssl:ttdoc/home/index", // The fallback jump address.
              "pc_url": "",
              "ios_url": "",
              "android_url": ""
            },
            "value": "document" // The return parameter value of the button. When the user clicks on the option, the application will return this value to the card request address.
          }
        ],
        "value": {
          // The callback data of the entire component.
          "key_1": "value_1"
        },
        "confirm": {
          // Secondary confirmation pop-up configuration
          "title": {
            "tag": "plain_text",
            "content": "title"
          },
          "text": {
            "tag": "plain_text",
            "content": "content"
          }
        }
      }
    ]
  }
}
Field descriptions

The field descriptions of the folding button group component are shown in the table below:

Field Name	Required	Type	Default Value	Description


tag

	

Yes

	

String

	

/

	

Tag of the folding button group. Fixed value is overflow.




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



width

	

No

	

String

	

default

	

Width of the folding button group. Supports the following enum values:

default: Default widthfill: Maximum supported width of the card[100,∞)px: Custom width. When exceeding the card width, it will be displayed according to the maximum supported width




options

	

Yes

	

Struct[]

	

Empty

	

Option buttons in the folding button group. See the options field description below for details.




value

	

No

	

JSON

	

Empty

	

Callback data for the entire component. When users click on the folding button of the folding button group, the value will be returned to the server receiving the callback data. You can subsequently process business with the received value from the server.This field value only supports key-value form JSON structure, and the key is of type String. Example value is as follows:

"value":{
    "key-1":Object-1,
    "key-2":Object-2,
    "key-3":Object-3,
    ······
}



confirm

	

No

	

Struct

	

This attribute is not effective by default.

	

Secondary confirmation dialog configuration. It prompts a secondary confirmation dialog when the user submits; only after the user clicks confirm, the entered content is submitted. This field provides confirmation and cancel buttons by default, and you only need to configure the title and content of the dialog.

Note: The confirm field will only trigger a secondary confirmation dialog when a button containing a submit attribute is clicked by the user.




└ title

	

Yes

	

Struct

	

/

	

Title of the secondary confirmation dialog.




└ └ tag

	

Yes

	

String

	

plain_text

	

Tag of the text for the secondary confirmation dialog title. Fixed value is plain_text.




└ └ content

	

Yes

	

String

	

/

	

Content of the secondary confirmation dialog title.




└ text

	

Yes

	

Struct

	

/

	

Content of the secondary confirmation dialog text.




└ └ tag

	

Yes

	

String

	

plain_text

	

Tag of the text for the secondary confirmation dialog content. Fixed value is plain_text.




└ └ content

	

Yes

	

String

	

/

	

Specific content of the secondary confirmation dialog text.

Description of options Field

You can add and configure multiple buttons in the option field. The related field descriptions are shown in the table below.

Field Name	Required	Type	Default Value	Description


text

	

No

	

Struct

	

Empty

	

Text on the button.




└ tag

	

No

	

String

	

plain_text

	

Tag of the text. Fixed value is plain_text.




└ content

	

No

	

String

	

Please input

	

Content of the text. Supports up to 100 characters.




multi_url

	

No

	

Struct

	

Empty

	

Jump links for multiple platforms.




└ url

	

No

	

String

	

Empty

	

Default jump link.




└ android_url

	

No

	

String

	

Empty

	

Jump link for Android platform. Can be configured as lark://msgcard/unsupported_action to declare that the current platform does not allow jumping.




└ ios_url

	

No

	

String

	

Empty

	

Jump link for iOS platform. Can be configured as lark://msgcard/unsupported_action to declare that the current platform does not allow jumping.




└ pc_url

	

No

	

String

	

Empty

	

Jump link for PC platform. Can be configured as lark://msgcard/unsupported_action to declare that the current platform does not allow jumping.




value

	

No

	

String

	

Empty

	

Return parameter value of this button. When users click on the option, the application will return this value to the card request address.

Callback Structure

After successfully configuring interaction for the component, when users interact with the component, the request address configured in the developer backend will receive callback data.

If you added a new card callback interaction (card.action.trigger), you can refer to Card Callback Interaction to understand the callback structure.
If you added an old card callback interaction (card.action.trigger_v1), you can refer to Message Card Callback Interaction (Old) to understand the callback structure.
Sample code

The following JSON 2.0 sample code achieves the card effect shown in the figure below. The card consists of three button components and one folding button group component, which folds two buttons within the group:

{
  "schema": "2.0",
  "body": {
    "direction": "vertical",
    "padding": "12px 12px 12px 12px",
    "elements": [
      {
        "tag": "column_set",
        "columns": [
          {
            "tag": "column",
            "width": "auto",
            "elements": [
              {
                "tag": "button",
                "text": {
                  "tag": "plain_text",
                  "content": "按钮 1"
                },
                "type": "default",
                "size": "medium"
              },
              {
                "tag": "button",
                "text": {
                  "tag": "plain_text",
                  "content": "按钮 2"
                },
                "type": "default",
                "size": "medium"
              },
              {
                "tag": "button",
                "text": {
                  "tag": "plain_text",
                  "content": "按钮 3"
                },
                "type": "default",
                "size": "medium"
              },
              {
                "tag": "overflow",
                "width": "",
                "options": [
                  {
                    "text": {
                      "tag": "plain_text",
                      "content": "按钮 4"
                    }
                  },
                  {
                    "text": {
                      "tag": "plain_text",
                      "content": "按钮 5"
                    }
                  }
                ]
              }
            ],
            "direction": "horizontal",
            "vertical_align": "top"
          }
        ],
        "margin": "0px 0px 0px 0px"
      }
    ]
  }
}
Previous:Button
Next:Single select dropdown menu
Need help with a problem?
Submit feedback

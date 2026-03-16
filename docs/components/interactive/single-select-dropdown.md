# Single select dropdown menu

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/single-select-dropdown-menu>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsInteractive componentsSingle select dropdown menu
Single select dropdown menu
Copy Page
Last updated on 2025-06-27
The contents of this article
Nesting rules
Component properties
JSON structure
Field descriptions
Callback structure
Sample code

The Dropdown Select - Single Selection Component supports customizing the option text, icon, and callback parameters of the single selection menu, making it an interactive component.

This document introduces the JSON 2.0 structure of the dropdown select-single component. To view the historical JSON 1.0 structure, refer to Dropdown Select-Single.

Nesting rules

The dropdown select-single component supports nesting within columns, form containers, folding panels, loop containers, and interactive containers. In the card building tool, the dropdown select-single component does not currently support nesting within interactive containers.

Component properties
JSON structure

The complete JSON 2.0 structure of the dropdown select-single component is as follows:

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "select_static", // The tag of the dropdown select-single component.
        "element_id": "custom_id", // The unique identifier for the operation component. It is used to specify the component when calling related interfaces. It needs to be customized by the developer.
        "margin": "0px 0px 0px 0px", // The margin of the component, default value is "0", with a supported range of [-99,99]px.
        "type": "text", // The border style of the component. The default value is default.
        "name": "select_static1", // The unique identifier of the dropdown select-single component. When the dropdown select-single component is nested in a form container, this attribute is effective and used to identify which dropdown select-single component the submitted text belongs to.
        "required": false, // Whether the content of the dropdown select-single component is required. The default value is false. When the dropdown select-single component is nested in a form container, this attribute is available. Other situations will report an error or be ineffective.
        "disabled": false, // Whether to disable this single select component. The default value is false.
        "initial_option": "Option 1", // The initial content displayed by the option. The default is empty.
        "placeholder": {
          // Placeholder text inside the dropdown select component.
          "tag": "plain_text",
          "content": "Default placeholder text"
        },
        "width": "default", // The width of the dropdown select component.
        "behaviors": [
          { // Configure the callback interaction for the dropdown select component.
            "type": "callback",
            "value": {
              // Callback interaction data. Supports string or object data types. The open platform SDK only supports object type callback interaction data.
              "key": "value"
            }
          }
        ],
        "options": [
          // Option configuration
          {
            "text": {
              // Option name
              "tag": "plain_text",
              "content": "I am an interactive component"
            },
            "icon": {
              // Add an icon as the prefix icon of the option. Supports custom icons or icons from the icon library.
              "tag": "standard_icon", // Icon type.
              "token": "chat-forbidden_outlined", // The token of the icon. Only effective when the tag is standard_icon.
              "color": "orange", // Icon color. Only effective when the tag is standard_icon.
              "img_key": "img_v2_38811724" // Key of the image. Only effective when the tag is custom_icon.
            },
            "value": "selectDemo1" // Option callback value, supports string type data.
          }
        ],
        "confirm": {
          // Secondary confirmation pop-up configuration
          "title": {
            "tag": "plain_text",
            "content": "Pop-up title"
          },
          "text": {
            "tag": "plain_text",
            "content": "Pop-up body text"
          }
        }
      }
    ]
  }
}
Field descriptions

The field descriptions for the dropdown-select component are as follows:

Field	Required	Type	Default	Description


tag

	

Yes

	

string

	

/

	

The label of the component. For the dropdown-select component, it takes a fixed value of select_static.




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

	

string

	

default

	

The border style of the component. Possible values:

default: with border style
text: pure text style without borders



name

	

No

	

String

	

empty

	

The unique identifier of the radio component. This attribute is effective when the radio component is embedded in a form container to identify which radio component the user-submitted text belongs to.

Note: When the radio component is nested in a form container, this field is required and must be unique within the card scope.




required

	

No

	

Boolean

	

false

	

Whether the content of the radio component is required. This attribute is available when the component is embedded in a form container. Other situations will result in errors or ineffective. Possible values:

true: The radio component is required. When the user clicks "Submit" in the form container without filling in the radio component, the front-end prompts "Required items are not filled in", and no callback request is sent to the developer's server.
false: The radio component is optional. When the user clicks "Submit" in the form container without filling in the radio component, the data in the form container is still submitted.



disabled

	

No

	

Boolean

	

false

	

Whether to disable the radio component. This attribute is only supported by Lark clients version 7.4 and above. Possible values:

true: Disable the radio component
false: The radio component remains available



initial_option

	

No

	

String

	

empty

	

The initial option content of the dropdown component. This configuration will override the placeholder text configuration.




initial_index

	

No

	

Int

	

Empty

	

The index of the initial option for the dropdown component. This configuration will override the placeholder text set by the placeholder configuration. 0 indicates no initial option will be displayed. 1 indicates the first option will be displayed.




placeholder

	

No

	

Object

	

/

	

The placeholder text inside the dropdown select component.




└ tag

	

Yes

	

String

	

plain_text

	

Placeholder prompt label. Fixed value is plain_text.




└ content

	

No

	

String

	

empty

	

The content of the placeholder text.




width

	

No

	

String

	

default

	

The width of the radio component. Supports the following enumerated values:

default: Default width
fill: Maximum supported width of the card
[100,∞)px: Custom width. When exceeding the width of the card, it will be displayed according to the maximum supported width



behaviors

	

No

	

Array

	

/

	

Interaction behavior. For details, see Configuring card interactions.




options

	

No

	

Array of objects

	

/

	

Configuration of options.




└ text

	

Yes

	

Object

	

/

	

The name of the option.




└ └ tag

	

Yes

	

String

	

plain_text

	

The tag of the option name. Fixed value is plain_text.




└ └ content

	

Yes

	

String

	

empty

	

The text of the option name.




└ icon

	

No

	

Object

	

/

	

Add an icon as a prefix icon for the text. Supports custom or icons from the icon library.




└└ tag

	

No

	

String

	

/

	

Tag of the icon type. Possible values:

standard_icon: Use icons from the icon library.
custom_icon: Use custom images as icons.



└└ token

	

No

	

String

	

/

	

Token of the icons in the icon library. Effective when tag is standard_icon. Enumeration values refer to Icon Library.




└└ color

	

No

	

String

	

/

	

Color of the icon. Supports setting colors for linear and solid icons (i.e., icons with outlined or filled at the end of the token). Effective when tag is standard_icon. Enumeration values refer to Enumeration Values for Colors.




└└ img_key

	

No

	

String

	

/

	

Key of the custom prefix icon image. Effective when tag is custom_icon.

Method of obtaining the icon key: Call the Upload Image interface to upload an image for sending messages, and obtain the image_key of the image in the return value.




└ value

	

Yes

	

String

	

/

	

Custom option callback value. When the user clicks an option of the interactive component, the value of value will be returned to the server receiving the callback data. Later, you can process business based on the value received by the server.

Note: Within the same selection component, the value of each option must be unique. Otherwise, there will be an interaction exception on the user side and the server will not be able to recognize which option the user clicked on.




confirm

	

No

	

Struct

	

This attribute does not take effect by default.

	

Configuration of the confirmation popup. It prompts the confirmation popup when the user submits; only when the user clicks confirm will the entered content be submitted. This field provides confirmation and cancel buttons by default, and you only need to configure the title and content of the popup.

Note: The confirm field only triggers the confirmation popup when the button containing the submit attribute is clicked by the user.




confirm.title

	

Yes

	

Struct

	

/

	

Title of the confirmation popup.




confirm.title.tag

	

Yes

	

String

	

plain_text

	

Tag of the confirmation popup title text. Fixed value is plain_text.




confirm.title.content

	

Yes

	

String

	

/

	

Content of the confirmation popup title.




confirm.text

	

Yes

	

Struct

	

/

	

Text content of the confirmation popup.




confirm.text.tag

	

Yes

	

String

	

plain_text

	

Tag of the confirmation popup text. Fixed value is plain_text.




confirm.text.content

	

Yes

	

String

	

/

	

Specific content of the confirmation popup text.

Callback structure

After successfully configuring interactions for the component, the callback data will be received at the request address you configured in the developer backend when the user interacts with the component.

If you added a new version of card callback interaction (card.action.trigger), refer to Card Callback Interaction for callback structure.
If you added an old version of card callback interaction (card.action.trigger_v1), refer to Message Card Callback Interaction (Old) for callback structure.
Sample code

The following JSON 2.0 sample code achieves the card effect shown in the image below:

{
  "schema": "2.0",
  "body": {
    "direction": "vertical",
    "padding": "12px 12px 12px 12px",
    "elements": [
      {
        "tag": "select_static",
        "initial_option": "",
        "placeholder": {
          "tag": "plain_text",
          "content": "请选择"
        },
        "options": [
          {
            "text": {
              "tag": "plain_text",
              "content": "选项1"
            },
            "value": "1",
            "icon": {
              "tag": "standard_icon",
              "token": "signature_outlined"
            }
          },
          {
            "text": {
              "tag": "plain_text",
              "content": "选项2"
            },
            "value": "2",
            "icon": {
              "tag": "standard_icon",
              "token": "signature_outlined"
            }
          },
          {
            "text": {
              "tag": "plain_text",
              "content": "选项3"
            },
            "value": "3",
            "icon": {
              "tag": "standard_icon",
              "token": "signature_outlined"
            }
          }
        ],
        "type": "default",
        "width": "default",
        "disabled": false,
        "behaviors": [
          {
            "type": "callback",
            "value": {
              "select_static": "click"
            }
          }
        ],
        "margin": "0px 0px 0px 0px"
      }
    ]
  }
}
Previous:Overflow
Next:Multi-select dropdown menu
Need help with a problem?
Submit feedback

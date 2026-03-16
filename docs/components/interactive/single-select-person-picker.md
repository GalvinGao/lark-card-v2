# Single select user picker

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/single-select-user-picker>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsInteractive componentsSingle select user picker
Single select user picker
Copy Page
Last updated on 2025-06-27
The contents of this article
Nesting rules
Component properties
JSON structure
Field description
Callback Structure
Sample code

The Personnel Selection - Single Choice Component supports adding specified personnel as single-choice options, and it is an interactive component.

This document introduces the JSON 2.0 structure of the single-select user picker component. To view the historical JSON 1.0 structure, refer to Single-select User Picker.

Nesting rules

The single-select user picker component supports nesting within columns, form containers, folding panels, loop containers, and interactive containers. In the builder tool, the single-select user picker component does not yet support nesting within interactive containers.

Component properties
JSON structure

The complete JSON 2.0 structure of the single-select user picker component is as follows:

{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "select_person", // Component tag.
                "element_id": "custom_id", // Unique identifier for the operation component. Used to specify the component in related interface calls. Needs to be customized by the developer.
                "margin": "0px 0px 0px 0px", // Component margin, default value is "0", supported range is [-99,99]px.
                "type": "text", // Component border style. Default value is default.
                "required": true, // Whether the option is required.
                "disabled": false, // Whether the option is disabled.
                "initial_option": "ou_48d0958ee4b2ab3eaf0b5f6c968xxxxx", // Initial content displayed by the option. Default is empty.
                "placeholder": {
                    // Placeholder text inside the user picker component.
                    "tag": "plain_text",
                    "content": "Default placeholder text"
                },
                "width": "default", // Width of the dropdown select component.
                "behaviors": [
                    { // Configure callback interaction for the component.
                        "type": "callback",
                        "value": {
                            // Callback interaction data. Supports string or object data types. The open platform SDK only supports object type callback interaction data.
                            "key": "value"
                        }
                    }
                ],
                "options": [
                    // Option configuration, only supports adding candidate users' open_id.
                    {
                        "value": "ou_48d0958ee4b2ab3eaf0b5f6c968xxxxx" // Candidate user's open_id.
                    },
                    {
                        "value": "ou_f9d24af786a14340721288cda6axxxxx" // Candidate user's open_id.
                    }
                ],
                "confirm": {
                    // Secondary confirmation popup configuration.
                    "title": {
                        "tag": "plain_text",
                        "content": "Popup title"
                    },
                    "text": {
                        "tag": "plain_text",
                        "content": "Popup body text"
                    }
                }
            }
        ]
    }
}
Field description

The field description of the Personnel Selection - Single Choice Component is as follows.

Field	Required	Type	Default Value	Description


tag

	

Yes

	

string

	

/

	

The tag of the component. The Personnel Selection - Single Choice Component takes the fixed value select_person.




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

	

The border style of the component. Optional values:

default: With border style
text: Pure text style without border



required

	

No

	

Boolean

	

false

	

Whether the content of the single-choice component is required. This property takes effect when the component is nested in a form container. Possible values:

true: The single-choice component is required. When the user clicks "Submit" of the form container, if the single-choice component is not filled in, the front end prompts "There are required items not filled in", and no back-pass request will be initiated to the developer's server.
false: The single-choice component is optional. When the user clicks "Submit" of the form container, if the single-choice component is not filled in, the data in the form container is still submitted.



disabled

	

No

	

Boolean

	

false

	

Whether to disable this single-choice component. Optional values:

true: Disable the single-choice component
false: The single-choice component remains available



initial_option

	

No

	

String

	

empty

	

The initial personnel's open_id displayed in the option. The value of open_id must match the open_id of one of the candidate users in options; otherwise, the option will not take effect.

Note: The initial_option property is currently not supported in the Card Construction Tool.




placeholder

	

No

	

object

	

/

	

The placeholder text in the Personnel Selection Component.




└ tag

	

Yes

	

string

	

plain_text

	

Placeholder prompt tag. The fixed value is plain_text.




└ content

	

No

	

String

	

/

	

The content of the placeholder text, supporting up to 100 characters.




width

	

No

	

String

	

default

	

The width of the Personnel Selection Component. Supports the following enumeration values:

default: Default width
fill: Maximum supported width of the card
[100,∞)px: Custom width. When it exceeds the card width, it will be displayed at the maximum supported width



options

	

No

	

Array of objects

	

/

	

Option value configuration. The option content is displayed in the order of the option array.




└ value

	

No

	

String

	

Empty

	

Option configuration, only supports adding the open_id of the candidate user. Learn more, refer to How to get different user IDs.

Note: When the options array is empty, or all the values of value are invalid, the candidate options are displayed as all member options in the conversation where the card is located.




confirm

	

No

	

Struct

	

This property does not take effect by default.

	

Configuration of the secondary confirmation pop-up window. It refers to the secondary confirmation pop-up prompt that pops up when the user submits; only after the user clicks confirm, the entered content is submitted. This field provides confirm and cancel buttons by default, you only need to configure the title and content of the pop-up window.

Note: The confirm field only triggers the secondary confirmation pop-up window when the user clicks the button containing the submit attribute.




confirm.title

	

Yes

	

Struct

	

/

	

Title of the secondary confirmation pop-up window.




confirm.title.tag

	

Yes

	

String

	

plain_text

	

Label of the text of the secondary confirmation pop-up window title. The fixed value is plain_text.




confirm.title.content

	

Yes

	

String

	

/

	

Content of the secondary confirmation pop-up window title.




confirm.text

	

Yes

	

Struct

	

/

	

Text content of the secondary confirmation pop-up window.




confirm.text.tag

	

Yes

	

String

	

plain_text

	

Label of the text of the secondary confirmation pop-up window. The fixed value is plain_text.




confirm.text.content

	

Yes

	

String

	

/

	

Specific content of the text of the secondary confirmation pop-up window.

Callback Structure

After successfully configuring the interaction for the component, when the user interacts with the component, the request address you configured in the developer backend will receive callback data.

If you add a new card back-pass interaction callback(card.action.trigger), you can refer to Card Back-pass Interaction to understand the callback structure.
If you add an old card back-pass interaction callback(card.action.trigger_v1), you can refer to Message Card Back-pass Interaction (Old) to understand the callback structure.
Sample code

The following JSON 2.0 structure example code can achieve the card effect shown in the figure. You need to fill in the actual user's user_id, union_id, or open_id.

{
    "schema": "2.0",
    "body": {
        "direction": "vertical",
        "padding": "12px 12px 12px 12px",
        "elements": [
            {
                "tag": "select_person",
                "placeholder": {
                    "tag": "plain_text",
                    "content": "请选择"
                },
                "options": [
                    {
                        "value": "ou_449b53ad6aee526f7ed311b216aabcef"
                    },
                    {
                        "value": "ou_449b53ad6aee526f7ed311b216aabcef"
                    }
                ],
                "width": "default",
                "type": "default",
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
Previous:Multi-select dropdown menu
Next:Multi-select user picker
Need help with a problem?
Submit feedback

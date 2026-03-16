# Time-selector

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/time-selector>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsInteractive componentsTime-selector
Time-selector
Copy Page
Last updated on 2025-06-27
The contents of this article
Notes
Nesting rules
Component properties
JSON Structure
Field descriptions
Callback structure
Example code

The time picker component is an interactive component used to provide time options. This document introduces the JSON structure and related properties of the time picker component.

This document introduces the JSON 2.0 structure of the time picker component. To view the historical JSON 1.0 structure, refer to time picker.

Notes

When using the time picker, you need to remind users to select timezone information corresponding to the current time scenario. For example, in the scenario of booking overseas hotels, the timezone of the hotel location is generally used; in the scheduling scenario, the timezone of the user's current location is generally used. The open platform will return the user's current timezone as a reference, but it does not mean that the user has selected that timezone.

Nesting rules

The time picker component supports nesting within column sets, form containers, folding panels, loop containers, and interactive containers. In the card building tool, the time picker component does not yet support nesting within interactive containers.

Component properties
JSON Structure

The JSON 2.0 structure of the time picker component is as follows:

{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "picker_time", // Tag of the time picker component.
                "element_id": "custom_id", // Unique identifier for the operation component. Used to specify the component in related interface calls. Needs to be customized by the developer.
                "margin": "0px 0px 0px 0px", // Component margin, default value is "0", supported range is [-99,99]px.
                "name": "picker_time1", // Unique identifier for the time picker component. This field is required when the component is nested within a form container.
                "required": false, // Whether the time is required. Default value is false.
                "disabled": false, // Whether the time picker component is disabled. Default value is false.
                "width": "default", // Width of the time picker.
                "behaviors": [
                    { // Configure callback interaction for the component.
                        "type": "callback",
                        "value": {
                            // Callback interaction data. Supports string or object data types. The open platform SDK only supports object type callback interaction data.
                            "key": "value"
                        }
                    }
                ],
                "initial_time": "11:30", // Initial time value.
                "placeholder": {
                    // Placeholder text inside the time picker component.
                    "tag": "plain_text",
                    "content": "Please select"
                },
                "value": {
                    // Callback data.
                    "key_1": "value_1"
                },
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
                }
            }
        ]
    }
}
Field descriptions

The field descriptions of the time picker component are as follows.

Field	Required	Type	Default Value	Description


tag

	

Yes

	

string

	

/

	

The tag of the component. For the time picker component, it takes a fixed value of picker_time.




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



name

	

No

	

String

	

Empty

	

The unique identifier of this time picker component. Used to identify which component the user-submitted data belongs to.

Note: This field is required when the time picker component is nested within a form container, and it must be unique within the card globally.




required

	

No

	

Boolean

	

false

	

Whether the time is required. This property can be used when the component is nested within a form container. Otherwise, it will result in an error or not take effect. It can take the values:

true: Time is required. When the user clicks "Submit" in the form container without filling in the time, the front end prompts "Required items are not filled in", and no callback request is sent to the developer's server.
false: Time is optional. When the user clicks "Submit" in the form container without filling in the time, the data in the form container is still submitted.



disabled

	

No

	

Boolean

	

false

	

Whether to disable this time picker. This property is only supported by Lark clients version 7.4 and above. It can take the values:

true: Disable the time picker component
false: The time picker component remains enabled



initial_time

	

No

	

String

	

Empty

	

Initial option value of the time picker component. Format is HH:mm. This configuration will override the placeholder text configuration.




placeholder

	

No

	

object

	

/

	

Placeholder text within the time picker component.

Note:

When the initial_time field is not configured to set the initial option value, this field is required.
When the initial_time field is configured to set the initial option value, this field will not take effect.



└ tag

	

Yes

	

String

	

plain_text

	

Placeholder tag. Fixed value is plain_text.




└ content

	

No

	

String

	

/

	

Content of the placeholder text, supporting up to 100 characters.




width

	

No

	

String

	

default

	

Width of the time picker component. Supports the following enumerated values:

default: Default width
fill: Maximum supported width of the card
[100,∞)px: Custom width. When exceeding the width of the card, it will be displayed according to the maximum supported width



value

	

Yes

	

JSON

	

/

	

Sets the return data for interaction. When the user clicks an option of the interactive component, the value will be returned to the server that receives the callback data. Subsequently, you can perform business logic based on the received value from the server.

This field value only supports JSON structure in key-value form, and the key is of type String. Example:

"value":{
    "key-1":Object-1,
    "key-2":Object-2,
    "key-3":Object-3,
    ······
}



confirm

	

No

	

Struct

	

Not applicable by default.

	

Configuration for a confirmation dialog. It prompts the user with a confirmation dialog upon submission, and only submits the entered content after the user clicks confirm. This field defaults to providing confirmation and cancel buttons, and you only need to configure the title and content of the dialog.Note: The confirm field is triggered only when the user clicks a button that includes the submit attribute.




confirm.title

	

Yes

	

Struct

	

/

	

The title of the confirmation dialog.




confirm.title.tag

	

Yes

	

String

	

plain_text

	

The tag for the confirmation dialog title text. Fixed value is plain_text.




confirm.title.content

	

Yes

	

String

	

/

	

The content of the confirmation dialog title.




confirm.text

	

Yes

	

Struct

	

/

	

The text content of the confirmation dialog.




confirm.text.tag

	

Yes

	

String

	

plain_text

	

The tag for the confirmation dialog text. Fixed value is plain_text.




confirm.text.content

	

Yes

	

String

	

/

	

The specific content of the confirmation dialog text.

Callback structure

After successfully configuring interaction for the component, when users interact with the component, the request address configured in your developer backend will receive callback data.

If you added the new version card callback interaction (card.action.trigger), you can refer to Card Callback Interaction to understand the callback structure.
If you added the old version card callback interaction (card.action.trigger_v1), you can refer to Message Card Callback Interaction (Old) to understand the callback structure.
Example code

The following JSON 2.0 example code can achieve the card effect as shown in the image below:

{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "picker_time",
                "placeholder": {
                    "tag": "plain_text",
                    "content": "请选择"
                },
                "width": "default",
                "initial_time": "09:00"
            },
            {
                "tag": "picker_time",
                "placeholder": {
                    "tag": "plain_text",
                    "content": "请选择"
                },
                "width": "default"
            }
        ]
    }
}
Previous:Date picker
Next:Date time picker
Need help with a problem?
Submit feedback

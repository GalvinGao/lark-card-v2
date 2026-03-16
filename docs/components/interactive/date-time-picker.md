# Date time picker

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/date-time-picker>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsInteractive componentsDate time picker
Date time picker
Copy Page
Last updated on 2025-06-27
The contents of this article
Notes
Nesting rules
Component properties
JSON structure
Field descriptions
Callback Structure
Sample code

The date time picker component is an interactive component used to provide options for both time and date selection.

This document introduces the JSON 2.0 structure of the date-time picker component. To view the historical JSON 1.0 structure, refer to date-time picker.

Notes

When using the date time picker, you need to remind users to select the time zone information corresponding to the current scenario. For example, in the scenario of booking overseas hotels, the time zone of the hotel location is generally used; in the scheduling scenario, the time zone of the user's current location is generally used. The open platform will return the user's current time zone as a reference, but it does not mean that the user has selected that time zone.

Nesting rules

The date picker component supports nesting within column sets, form containers, folding panels, loop containers, and interactive containers. In the building tool, the date-time picker component does not yet support nesting within interactive containers.

Component properties
JSON structure

The complete JSON 2.0 structure of the date-time picker component is as follows:

{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "picker_datetime", // Tag of the date-time picker component.
                "element_id": "custom_id", // Unique identifier for the operation component. Used to specify the component in related interface calls. Needs to be customized by the developer.
                "margin": "0px 0px 0px 0px", // Component margin, default value is "0", supported range is [-99,99]px.
                "name": "picker_datetime1", // Unique identifier for the date-time picker component. This field is required when the component is nested within a form container.
                "required": false, // Whether the date-time is required. Default value is false.
                "disabled": false, // Whether the date-time picker component is disabled. Default value is false.
                "width": "default", // Width of the date-time picker.
                "behaviors": [
                    { // Configure callback interaction for the component.
                        "type": "callback",
                        "value": {
                            // Callback interaction data. Supports string or object data types. The open platform SDK only supports object type callback interaction data.
                            "key": "value"
                        }
                    }
                ],
                "initial_datetime": "2024-01-01 11:30", // Initial date-time value. Default is empty.
                "placeholder": {
                    // Placeholder text inside the date-time picker component.
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

The field descriptions of the date time picker component are shown in the following table.

Field	Required	Type	Default Value	Description


tag

	

Yes

	

String

	

/

	

The tag of the component. For the date time picker component, it has a fixed value of picker_datetime.




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

	

The unique identifier of the date time picker component. This field takes effect when the date time picker is embedded within a form container and is used to identify which component the submitted data belongs to.

Note: When the date time picker component is nested within a form container, this field is required and must be unique within the global card.




required

	

No

	

Boolean

	

false

	

Whether the date time is required. This property is available when the component is embedded within a form container. Otherwise, it will either cause an error or not take effect. Possible values:

true: The date time is required. When the user clicks "Submit" in the form container without filling in the date time, a frontend prompt "Some required items are not filled in" will be displayed, and no callback request will be sent to the developer's server.
false: The date time is optional. When the user clicks "Submit" in the form container without filling in the date time, the data in the form container is still submitted.



disabled

	

No

	

Boolean

	

false

	

Whether to disable the date time picker. This property is only supported in Lark client version 7.4 and above. Possible values:

true: Disable the date time picker component.
false: Keep the date time picker component enabled.



initial_datetime

	

No

	

String

	

Empty

	

The initial selection value of the date time picker component. The format is yyyy-MM-dd HH:mm. This configuration will override the placeholder text configured in placeholder.




placeholder

	

No

	

Object

	

/

	

Placeholder text inside the date time picker component.

Note:

When the initial_datetime field is not configured to set the initial selection value, this field is required.
When the initial_datetime field is configured to set the initial selection value, this field will no longer take effect.



└ tag

	

Yes

	

String

	

plain_text

	

Placeholder tag. It has a fixed value of plain_text.




└ content

	

No

	

String

	

/

	

Content of the placeholder text, supporting up to 100 characters.




width

	

No

	

String

	

default

	

Width of the date time picker component. Supports the following enumeration values:

default: Default width
fill: Maximum supported width of the card
[100,∞)px: Custom width. When exceeding the width of the card, it will be displayed according to the maximum supported width.



value

	

Yes

	

JSON

	

/

	

Set the return data for interaction. When the user clicks an option in the interactive component, the value will be returned to the server receiving the callback data. Later, you can process business logic based on the received value from the server.

This field only supports key-value JSON structure, and the key is of type String.




confirm

	

No

	

Struct

	

This property does not take effect by default.

	

Configuration for the secondary confirmation dialog. It prompts a secondary confirmation dialog when the user submits, and only submits the entered content after the user clicks confirm. This field provides default confirm and cancel buttons, so you only need to configure the title and content of the dialog.

Note: The confirm field will only trigger the secondary confirmation dialog when the user clicks a button containing the submit attribute.




└ title

	

Yes

	

Struct

	

/

	

Title of the secondary confirmation dialog.




└ └ tag

	

Yes

	

String

	

plain_text

	

Tag for the text of the secondary confirmation dialog title. Fixed value is plain_text.




└ └ content

	

Yes

	

String

	

/

	

Content of the title of the secondary confirmation dialog.




└ text

	

Yes

	

Struct

	

/

	

Text content of the secondary confirmation dialog.




└ └ tag

	

Yes

	

String

	

plain_text

	

Tag for the text of the secondary confirmation dialog. Fixed value is plain_text.




└ └ content

	

Yes

	

String

	

/

	

Specific content of the text of the secondary confirmation dialog.

Callback Structure

After successfully configuring interactions for the component, when users interact with the component, the request address configured in the developer console will receive callback data.

If you have added the new version of the card callback interaction (card.action.trigger), you can refer to Card Callback Interaction for the callback structure.
If you have added the old version of the card callback interaction (card.action.trigger_v1), you can refer to Message Card Callback Interaction (Old) for the callback structure.
Sample code

The following JSON 2.0 sample code can achieve the card effect as shown in the image below:

{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "picker_datetime",
                "placeholder": {
                    "tag": "plain_text",
                    "content": "请选择"
                },
                "width": "default"
            },
            {
                "tag": "picker_datetime",
                "placeholder": {
                    "tag": "plain_text",
                    "content": "请选择"
                },
                "width": "default",
                "initial_datetime": "2024-01-01 08:00"
            }
        ]
    }
}
Previous:Time-selector
Next:Image picker
Need help with a problem?
Submit feedback

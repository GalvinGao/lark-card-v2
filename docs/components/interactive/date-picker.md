# Date picker

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/date-picker>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsInteractive componentsDate picker
Date picker
Copy Page
Last updated on 2025-06-27
The contents of this article
Notes
Nesting rules
Component properties
JSON structure
Field descriptions
Callback structure
Example code

The date picker component is an interactive component used to provide date options. This document introduces the JSON structure and related properties of the date picker component.

This document introduces the JSON 2.0 structure of the date picker component. To view the historical JSON 1.0 structure, refer to Date Picker.

Notes

When using the date picker, you need to remind users to select timezone information corresponding to the current date scenario. For example, in the scenario of booking overseas hotels, the timezone of the hotel location is generally used; in the scheduling scenario, the timezone of the user's current location is generally used. The open platform will return the user's current timezone as a reference, but it does not mean that the user has selected that timezone.

Nesting rules

The date picker component supports nesting within column containers, form containers, foldable panels, loop containers, and interactive containers. In the builder tool, the date picker component does not support nesting within interactive containers.

Component properties
JSON structure

The complete JSON 2.0 structure of the date picker component is as follows:

{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "date_picker",
                "element_id": "custom_id", // Unique identifier for the operation component. Used to specify the component in related interface calls. Needs to be customized by the developer.
                "margin": "0px 0px 0px 0px", // Component margin, default value is "0", supported range is [-99,99]px.
                "name": "date_picker1", // Unique identifier for the date picker component. This field is required when the component is nested within a form container.
                "required": false, // Whether the date is required, default value is false.
                "disabled": false, // Whether the date picker component is disabled. Default value is false.
                "width": "default", // Width of the date picker.
                "behaviors": [
                    { // Configure callback interaction for the component.
                        "type": "callback",
                        "value": {
                            // Callback interaction data. Supports string or object data types. The open platform SDK only supports object type callback interaction data.
                            "key": "value"
                        }
                    }
                ],
                "initial_date": "2024-01-01", // Initial date value.
                "placeholder": {
                    // Placeholder text inside the date picker component.
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

The field descriptions of the date picker component are as follows.

Field	Required	Type	Default Value	Description


tag

	

Yes

	

string

	

/

	

The tag of the component. For the date picker component, it takes a fixed value of date_picker.




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

	

The unique identifier of this date picker component. When the date picker is nested within a form container, this property takes effect to identify which component the user-submitted data belongs to.

Note: This field is required when the date picker component is nested within a form container, and it must be unique within the card globally.




required

	

No

	

Boolean

	

false

	

Whether the content of the date is required. This property can be used when the component is nested within a form container. Otherwise, it will result in an error or not take effect. It can take the values:

true: Date is required. When the user clicks "Submit" in the form container without filling in the date, the front end prompts "Required items are not filled in", and no callback request is sent to the developer's server.
false: Date is optional. When the user clicks "Submit" in the form container without filling in the date, the data in the form container is still submitted.



disabled

	

No

	

Boolean

	

false

	

Whether to disable this date picker. This property is only supported by Lark clients version 7.4 and above. It can take the values:

true: Disable the date picker component
false: The date picker component remains enabled



placeholder

	

No

	

object

	

/

	

Placeholder text within the date picker component.




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

	

Width of the date picker component. Supports the following enumerated values:

default: Default width
fill: Maximum supported width of the card
[100,∞)px: Custom width. When exceeding the width of the card, it will be displayed according to the maximum supported width



initial_date

	

No

	

String

	

Empty

	

Initial option value of the date picker component. Format is yyyy-MM-dd. This configuration will override the placeholder text configuration.




value

	

Yes

	

JSON

	

/

	

Sets the return data for interaction. When the user clicks an option of the interactive component, the value will be returned to the server that receives the callback data. Subsequently, you can perform business logic based on the received value from the server. This field value only supports JSON structure in key-value form, and the key is of type String.

Example:

"value":{
    "key-1":Object-1,
    "key-2":Object-2,
    "key-3":Object-3,
    ······
}



confirm

	

No

	

Struct

	

This property does not take effect by default.

	

Configuration for the secondary confirmation dialog. It prompts a secondary confirmation dialog when the user submits, and only submits the entered content after the user clicks confirm. This field provides default confirm and cancel buttons, so you only need to configure the title and content of the dialog.

Note: The confirm field will only trigger the secondary confirmation dialog when the user clicks a button containing the submit attribute.




confirm.title

	

Yes

	

Struct

	

/

	

Title of the secondary confirmation dialog.




confirm.title.tag

	

Yes

	

String

	

plain_text

	

Tag for the text of the secondary confirmation dialog title. Fixed value is plain_text.




confirm.title.content

	

Yes

	

String

	

/

	

Content of the title of the secondary confirmation dialog.




confirm.text

	

Yes

	

Struct

	

/

	

Text content of the secondary confirmation dialog.




confirm.text.tag

	

Yes

	

String

	

plain_text

	

Tag for the text of the secondary confirmation dialog. Fixed value is plain_text.




confirm.text.content

	

Yes

	

String

	

/

	

Specific content of the text of the secondary confirmation dialog.

Callback structure

After successfully configuring interaction for the component, when users interact with the component, the request address configured in your developer backend will receive callback data.

If you added a new version of card callback interaction (card.action.trigger), you can refer to Card Callback Interaction for callback structure.
If you added an old version of card callback interaction (card.action.trigger_v1), you can refer to Message Card Callback Interaction (Old) for callback structure.
Example code

The following JSON 2.0 example code can achieve the card effect shown in the image below:

{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "date_picker",
                "placeholder": {
                    "tag": "plain_text",
                    "content": "请选择"
                },
                "width": "default",
                "initial_date": "2024-01-01"
            },
            {
                "tag": "date_picker",
                "placeholder": {
                    "tag": "plain_text",
                    "content": "请选择"
                },
                "width": "default"
            }
        ]
    }
}
Previous:Multi-select user picker
Next:Time-selector
Need help with a problem?
Submit feedback

# Divider

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/divider>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsDisplay componentsDivider
Divider
Copy Page
Last updated on 2025-06-27
The contents of this article
JSON structure
Field descriptions
Example code

You can add a divider component in the card to make the content clearer.

JSON structure

The complete JSON 2.0 structure of the divider is as follows:

{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "hr",
                "element_id": "custom_id", // Unique identifier for the operation component. New attribute in JSON 2.0. Used to specify the component when calling related interfaces. Custom defined by the developer.
                "margin": "0px 0px 0px 0px" // Margin of the component. New attribute in JSON 2.0. Default value is "0", supported range is [-99,99]px.
            }
        ]
    }
}
Field descriptions

The field descriptions for the divider component are shown in the table below.

Parameter	Required	Type	Default value	Description


tag

	

Yes

	

String

	

Empty

	

The tag of the component. The fixed value for the divider component is hr.




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
Example code

The following JSON 2.0 example code can achieve the card effect shown in the image below:

{
    "schema": "2.0",
    "body": {
        "direction": "vertical",
        "padding": "12px 12px 12px 12px",
        "elements": [
            {
                "tag": "div",
                "text": {
                    "tag": "plain_text",
                    "content": "普通文本示例",
                    "text_size": "normal",
                    "text_align": "left",
                    "text_color": "default"
                },
                "margin": "0px 0px 0px 0px"
            },
            {
                "tag": "hr",
                "margin": "0px 0px 0px 0px"
            },
            {
                "tag": "button",
                "text": {
                    "tag": "plain_text",
                    "content": "查看更多"
                },
                "type": "primary",
                "width": "default",
                "size": "medium",
                "margin": "0px 0px 0px 0px"
            }
        ]
    }
}
Previous:Multi image laylout
Next:User profile
Need help with a problem?
Submit feedback

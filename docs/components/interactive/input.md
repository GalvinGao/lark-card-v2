# Last updated on 2025-06-27

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/input>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsInteractive componentsInput
Input
Copy Page
Last updated on 2025-06-27
The contents of this article
Nesting rules
Component attributes
JSON structure
Field explanation
Callback Structure
Sample code

In scenarios where you use cards for content collection, you might need to gather subjective content from users, such as reasons, evaluations, remarks, etc. In such cases, you can use the input box component to facilitate simple text content collection.This document introduces the JSON 2.0 structure of the input box component. To view the historical JSON 1.0 structure, refer to Input Box.

To use the input box component in conjunction with the button component, you need to embed both the input box and the button within a form container.

Nesting rules

The input box component supports nesting within columns, form containers, folding panels, loop containers, and interactive containers. In a form container, the data of the input box component is submitted asynchronously, meaning that after the user completes all form items and clicks the button bound to the submit event in the form container, all data including that of the input box component will be sent back to the developer's server in one callback.

Component attributes
JSON structure

The complete JSON 2.0 structure of the input box component is as follows:

{
  "schema": "2.0", // Version of the card JSON structure. Default is 1.0. To use JSON 2.0 structure, you must explicitly declare 2.0.
  "body": {
    "elements": [
      {
        "tag": "input", // Tag of the input box.
        "element_id": "custom_id", // Unique identifier for the operation component. New attribute in JSON 2.0. Used to specify the component in related interface calls. Customizable by the developer.
        "margin": "0px 0px 0px 0px", // Margin of the component. New attribute in JSON 2.0. Default value is "0", supported range is [-99,99]px.
        "name": "input1", // Unique identifier of the input box. When the input box is embedded in a form container, this attribute is effective and is used to identify which input box the submitted text belongs to.
        "required": false, // Whether the content of the input box is required. This attribute is available when the input box is embedded in a form container. Otherwise, it will cause an error or be ineffective.
        "placeholder": {
          // Placeholder text in the input box.
          "tag": "plain_text",
          "content": "Please enter"
        },
        "default_value": "demo", // Pre-filled content in the input box for the user.
        "disabled": false, // Whether to disable the input box component. Default value is false.
        "disabled_tips": { // Tips displayed when the component is disabled and the user hovers the cursor over the entire component.
          "tag": "plain_text",
          "content": "User disabled tips"
        },
        "width": "default", // Width of the input box.
        "behaviors": [
          { // Configure callback interaction for the component.
            "type": "callback",
            "value": {
              // Callback interaction data. Supports string or object data types. The open platform SDK only supports object type callback interaction data.
              "key": "value"
            }
          }
        ],
        "max_length": 5, // Maximum text length that the input box can accommodate. Default value is 1000.
        "input_type": "multiline_text", // Specifies the input type of the input box. Default is text, which is text type.
        "rows": 1, // Default number of rows displayed when the input type is multiline text.
        "auto_resize": true, // Whether the height of the input box adapts to the text height when the input type is multiline text. Effective only on PC.
        "max_rows": 5, // Maximum number of rows displayed in the input box. Effective only when `auto_resize` is true.
        "show_icon": false, // Whether to display a prefix icon when the input type is password.
        "label": {
          // Text label, which describes the input box and prompts the user for the content to be filled.
          "tag": "plain_text",
          "content": "Please enter text:"
        },
        "label_position": "left", // Position of the text label. Default value is top.
        "value": {
          // Returned data, supports string or object data types.
          "k": "v"
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
Field explanation

The input box component fields are explained in the following table:

Field name	Required	Type	Default value	Description


tag

	

Yes

	

String

	

Empty

	

Label of the input box. The fixed value is input.




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

	

Unique identifier of the input box. Effective when embedded in a form container, it identifies which input box the user-submitted text belongs to.

Note: When the input box component is nested within a form container, this field is required and must be unique globally within the card.




required

	

No

	

Boolean

	

false

	

Indicates whether the input box content is required. Available when embedded in a form container. In other cases, it will error out or be ineffective. Possible values:

true: Input box is required. If not filled when the user clicks "submit" in the form container, the frontend will prompt "Required item not filled," and no request will be made to the developer's server.
false: Input box is optional. Data is still submitted even if the input box is unfilled when the user clicks "submit" in the form container.



disabled

	

No

	

Boolean

	

false

	

Whether the input box is disabled. This attribute only supports Lark V7.4 and above. Possible values:

true: Input box is disabled
false: Input box remains enabled



placeholder

	

No

	

text structure

	

/

	

Placeholder text within the input box.




└ tag

	

No

	

String

	

plain_text

	

Label for the placeholder text. The fixed value is plain_text.




└ content

	

No

	

String

	

Please enter

	

The content of the placeholder text, supporting up to 100 characters. For example, "Please enter content."




default_value

	

No

	

String

	

This property is not activated by default.

	

Pre-filled content for the user in the input box. It displays as the style of text entered by the user and is pending submission.




width

	

No

	

String

	

default

	

Width of the input box. Supports the following values:

default: Default width
fill: Maximum width supported by the card
Custom width [100,∞)px: If it exceeds the card width, it will be displayed at the maximum supported width



behaviors

	

Yes

	

Struct

	

/

	

Configures the interaction type and specific interaction behaviors. For more details, refer to the field description of behaviors in Configuring Card Interactions.




max_length

	

No

	

Number

	

1,000

	

Maximum text length the input box can accommodate, within the range of 1 to 1,000. If the user's text exceeds this length, the component will display an error prompt.




input_type

	

No

	

String

	

text

	

Specifies the input type of the input box. The default is text, which means text type. Supports the following enumeration values:

text: ordinary text
multiline_text: multiline text, which allows entering multiline text content containing line breaks. Line breaks are returned as \n in the callback.
password: password. The text entered by the user will be displayed as "•".

Note: The input_type attribute is currently not supported in the card builder tool.




show_icon

	

No

	

Boolean

	

true

	

Whether to display the prefix icon shown below when the input type is password. Only effective when input_type is password.




rows

	

No

	

Number

	

5

	

When the input type is multiline text, the default number of display rows of the input box. Only effective when input_type is multiline_text.




auto_resize

	

No

	

Boolean

	

false

	

When the input type is multiline text, whether the height of the input box is automatically adjusted to the height of the text. Only effective on PC. Only effective when input_type is multiline_text. Optional values:

true: The height of the input box adapts to the height of the text in the input box.
false: The height of the input box is fixed to the height specified by the rows attribute and does not change with the content in the input box.



max_rows

	

No

	

Number

	

None

	

The maximum number of display rows of the input box. Only effective when auto_resize is true.Note:

The value should be an integer greater than or equal to 1. If it is less than 1, it will automatically be set to 1. If it is not an integer, it will be rounded.
If the value is empty, there is no limit to the maximum display height of the input box (default value), but the maximum height that the input box can display during front-end rendering does not exceed x rows.



label

	

No

	

text structure

	

This property is not activated by default.

	

Text label, i.e., the description of the input box, used to prompt the user on what content to fill in. Mostly used in input box components embedded within form containers.




└ tag

	

No

	

String

	

plain_text

	

Label for the input box description. The fixed value is plain_text.




└ content

	

No

	

String

	

/

	

Description content.




label_position

	

No

	

String

	

top

	

Position of the text label. Possible values:

top: Text label is above the input box
left: Text label is to the left of the input box

Note:In mobile and other narrow screen scenarios, the text label will automatically adjust to be fixed above the input box.




value

	

No

	

String or Object

	

Empty

	

You can customize the data to be passed back during interaction events, supports string or object data types.




confirm

	

No

	

Struct

	

This property is not activated by default.

	

Configuration for the confirmation popup. Refers to a popup that appears upon submission; content is submitted only after the user clicks confirm. This field provides confirm and cancel buttons by default, you only need to configure the popup's title and content.

Note: The confirm field is only triggered when the user clicks a button containing the submit property.




└ title

	

Yes

	

Struct

	

/

	

Title of the confirmation popup.




└ └ tag

	

Yes

	

String

	

plain_text

	

Label for the confirmation popup's title text. The fixed value is plain_text.




└ └ content

	

Yes

	

String

	

/

	

Content of the confirmation popup's title.




└ text

	

Yes

	

Struct

	

/

	

Content of the confirmation popup's text.




└ └ tag

	

Yes

	

String

	

plain_text

	

Label for the confirmation popup text. The fixed value is plain_text.




└ └ content

	

Yes

	

String

	

/

	

Specific content of the confirmation popup text.

Callback Structure

To use the input box component, you need to enable interactivity for your card. For more details, refer to Configuring Card Interactions. After configuration, when a user clicks the submit button in the input box, the interaction event will be passed back as shown below. If the input box is embedded in a form container, you can refer to the callback structure of the Form Container to understand the callback of the input box. You can also refer to Card Callback Communication for more explanations of the parameters.

{
    "schema": "2.0", // Version of the callback
    "header": { // Basic information of the callback
        "event_id": "f7984f25108f8137722bb63c*****", // Unique identifier of the callback
        "token": "066zT6pS4QCbgj5Do145GfDbbag*****", // Application's Verification Token
        "create_time": "1603977298000000", // Time when the callback was sent, close to the time when the event occurred
        "event_type": "card.action.trigger", // Type of the callback, fixed for card interaction scenarios as "card.action.trigger"
        "tenant_key": "2df73991750*****", // Tenant key to which the application belongs, i.e., the unique identifier of the tenant
        "app_id": "cli_a5fb0ae6a4******" // App ID of the application
    },
    "event": { // Detailed information of the callback
        "operator": { // Information of the callback initiator
            "tenant_key": "2df73991750*****", // Tenant key of the callback initiator, i.e., the unique identifier of the tenant
            "user_id": "867*****", // User ID of the callback initiator. This parameter returns when the application enables the "Obtain User ID" permission
            "open_id": "ou_3c14f3a59eaf2825dbe25359f15*****" // Open ID of the callback initiator
        },
        "token": "c-295ee57216a5dc9de90fefd0aadb4b1d7d******", // Token for updating the card, valid for 30 minutes, can be updated up to 2 times
        "action": { // Data passed back by the user operating the interaction component
            "value": { // Custom data passed back in the interaction event, corresponding to the value attribute in the component
                "key": "value"
            },
            "tag": "input", // Label of the input box component
            "input_value": "Zhang Min",  // Data submitted by the user in the input box
            "name": "Input_lf4fmxwfrd9" // Name of the input box component, i.e., the component ID in the building tool, customizable
        },
        "host": "im_message", // Scenario where the card is displayed
        "context": { // Information related to the card display scenario
            "open_message_id": "om_574d639e4a44e4dd646eaf628e2*****", // Message ID where the card is located
            "open_chat_id": "oc_e4d2605ca917e695f54f11aaf56*****" // Chat ID where the card is located
        }
    }
}
Sample code

The following JSON sample code can achieve the card effect as shown in the image below. The card consists of a form container with three embedded input box components:

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "form",
        "elements": [
          {
            "tag": "input",
            "element_id": "username",
            "margin": "0px 0px 0px 0px",
            "placeholder": {
              "tag": "plain_text",
              "content": "请输入"
            },
            "default_value": "",
            "width": "default",
            "label": {
              "tag": "plain_text",
              "content": "用户名："
            },
            "name": "Input_31q6mtuvdx9"
          },
          {
            "tag": "input",
            "element_id": "password",
            "margin": "0px 0px 0px 0px",
            "input_type": "password",
            "placeholder": {
              "tag": "plain_text",
              "content": "请输入"
            },
            "default_value": "",
            "width": "default",
            "label": {
              "tag": "plain_text",
              "content": "密码："
            },
            "label_position": "top",
            "name": "Input_5hez3q41fck"
          },
          {
            "tag": "input",
            "element_id": "address",
            "margin": "0px 0px 0px 0px",
            "input_type": "multiline_text",
            "rows": 4,
            "auto_resize": true,
            "placeholder": {
              "tag": "plain_text",
              "content": "请输入"
            },
            "default_value": "",
            "width": "default",
            "label": {
              "tag": "plain_text",
              "content": "收货地址："
            },
            "name": "Input_u2k3lbrokvd"
          },
          {
            "tag": "column_set",
            "flex_mode": "none",
            "background_style": "default",
            "horizontal_spacing": "default",
            "columns": [
              {
                "tag": "column",
                "width": "auto",
                "vertical_align": "top",
                "elements": [
                  {
                    "tag": "button",
                    "text": {
                      "tag": "plain_text",
                      "content": "提交"
                    },
                    "type": "primary",
                    "action_type": "form_submit",
                    "name": "Button_lrocopxs"
                  }
                ]
              },
              {
                "tag": "column",
                "width": "auto",
                "vertical_align": "top",
                "elements": [
                  {
                    "tag": "button",
                    "text": {
                      "tag": "plain_text",
                      "content": "取消"
                    },
                    "type": "default",
                    "action_type": "form_reset",
                    "name": "Button_lrocopxt"
                  }
                ]
              }
            ],
            "margin": "0px"
          }
        ],
        "name": "Form_lrocopxr"
      }
    ]
  }
}
Previous:Table
Next:Button
Need help with a problem?
Submit feedback

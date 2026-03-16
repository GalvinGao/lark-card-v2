# Multi-select user picker

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/multi-select-user-picker>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsInteractive componentsMulti-select user picker
Multi-select user picker
Copy Page
Last updated on 2025-06-27
The contents of this article
Nesting rules
Component properties
JSON structure
Field descriptions
Callback Structure
Sample code

The dropdown select - multi-select component supports customizing the option text, icon, and return parameters of the multi-select menu. It is an interactive component and needs to be embedded within a form container to be used.

This document introduces the JSON 2.0 structure of the multi-select user picker component. To view the historical JSON 1.0 structure, refer to Multi-select user picker.

Nesting rules

The dropdown select - multi-select component is only supported for use embedded within a form container, and the content is submitted through the "Submit" button of the form container. To understand form containers and their interactive configurations, refer to the Form Container.

Component properties
JSON structure

The JSON 2.0 data structure of the dropdown select - multi-select component is as follows:

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "form", // The tag of the form container. The person selection - multi-select component must be placed in the form container for use.
        "elements": [
          {
            "tag": "multi_select_person", // The tag of the person selection - multi-select component.
            "element_id": "custom_id", // The unique identifier of the component. It is used to specify the component when calling related component interfaces. It needs to be customized by the developer.
            "margin": "0px 0px 0px 0px", // The margin of the component, default value is "0", supported range is [-99,99]px.
            "type": "text", // The border style of the component. Default value is "default".
            "name": "multi_select_users", // The custom identifier of the component in the form container, used to identify which component's data is submitted by the user.
            "placeholder": {
              // Placeholder text inside the person selection component.
              "tag": "plain_text",
              "content": "Default placeholder text"
            },
            "width": "default", // The width of the dropdown selection component.
            "behaviors": [
              { // Configure callback interactions for the component.
                "type": "callback",
                "value": {
                  // Callback interaction data. Supports string or object data types. The open platform SDK only supports object type callback interaction data.
                  "key": "value"
                }
              }
            ],
            "required": true, // Whether the option is required.
            "disabled": false, // Whether the option is disabled.
            "selected_values": [
              "ou_48d0958ee4b2ab3eaf0b5f6c968xxxxx"
            ], // The default selected options of the component. The values of the array items need to correspond to options.value.
            "options": [
              // Option configuration, only supports adding candidate users' open_id.
              {
                "value": "ou_48d0958ee4b2ab3eaf0b5f6c968xxxxx"
              },
              {
                "value": "ou_f9d24af786a14340721288cda6axxxxx"
              }
            ],
            "confirm": {
              // Secondary confirmation popup configuration
              "title": {
                "tag": "plain_text",
                "content": "Popup title"
              },
              "text": {
                "tag": "plain_text",
                "content": "Popup body text"
              }
            }
          },
          { // Below are the properties of embedded components such as columns, buttons, etc. in the form container, for details refer to the form container documentation.
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
                      "content": "Submit"
                    },
                    "type": "primary",
                    "width": "default",
                    "form_action_type": "submit",
                    "name": "Button_m8gyoz81"
                  }
                ],
                "vertical_align": "top"
              },
              {
                "tag": "column",
                "width": "auto",
                "elements": [
                  {
                    "tag": "button",
                    "text": {
                      "tag": "plain_text",
                      "content": "Cancel"
                    },
                    "type": "default",
                    "width": "default",
                    "form_action_type": "reset",
                    "name": "Button_m8gyoz82"
                  }
                ],
                "vertical_align": "top"
              }
            ]
          }
        ],
        "padding": "4px 0px 4px 0px",
        "margin": "0px 0px 0px 0px",
        "name": "Form_m8gyoz80"
      }
    ]
  }
}
Field descriptions

The field descriptions for the dropdown-select-multi-select component are as follows:

Field	Required	Type	Default Value	Description


tag

	

Yes

	

String

	

/

	

The label of the component. For the dropdown-select-multi-select component, it has a fixed value of multi_select_static.




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

	

The border style of the component. Possible values:

default: with border style
text: plain text style without border



name

	

Yes

	

string

	

Empty

	

The unique identifier of the component in the form container. When the multi-select component is embedded in a form container, this attribute is effective for identifying which component the user-submitted data belongs to.

Note: When the multi-select component is nested in a form container, this field is required and must be unique within the card globally.




placeholder

	

No

	

Object

	

Empty

	

The placeholder text inside the dropdown select component when no option is selected.




└ tag

	

Yes

	

String

	

plain_text

	

The tag for the placeholder. Fixed value is plain_text.




└ content

	

No

	

String

	

Please select

	

The content of the placeholder text, supporting up to 100 characters.




width

	

No

	

String

	

default

	

The width of the dropdown select component. Supports the following enumerations:

default: default width:
When the component has a border (i.e., "type":"default"), the default width value is fixed at 282 px.
When the component does not have a border (i.e., "type":"text"), the component width adapts to the content width of the selector.
fill: the component width will fill the width of the parent container.
[100,∞)px: custom fixed numerical width, such as 200px. The minimum value is 100px. When exceeding the width of the parent container, it will be displayed according to the width of the parent container.



required

	

No

	

Bool

	

true

	

Whether the options of the multi-select component are required. This property is available when the component is embedded in a form container. In other cases, it will result in an error or have no effect. Possible values:

true: the options of the multi-select component are required. When the user clicks the "Submit" button of the form container without selecting multiple options, the frontend prompts "required items are not filled in", and no callback request is sent to the developer's server.
false: the options of the multi-select component are optional. When the user clicks the "Submit" button of the form container without selecting multiple options, the data in the form container is still submitted.



disabled

	

No

	

Bool

	

false

	

Whether to disable the multi-select component. Possible values:

true: disable the multi-select component, displaying custom placeholder text or initial values of options, and end users cannot modify the interaction.
false: the multi-select component remains available.



selected_values

	

No

	

Array of string

	

Empty

	

The options selected by default in the multi-select component. The values of array items need to correspond to options.value.




options

	

No

	

Array of objects

	

/

	

Option value configuration. Display options content in the order of the option array.




└ text

	

Yes

	

Object

	

Empty

	

Option name. Display blank option when empty. The JSON structure is as follows, described using plain text objects:

"text": {
// Option name tag. Fixed value is plain_text.
        "tag": "plain_text", 
// Option name text.
        "content": "I am an option"
}



└ icon

	

No

	

Object

	

/

	

Add an icon as a prefix icon for the option. Supports custom or using icons from the icon library.




└ └ tag

	

No

	

String

	

/

	

The tag type of the icon. Possible values:

standard_icon: use icons from the icon library.
custom_icon: use a custom image as the icon.



└ └ token

	

No

	

String

	

/

	

The token of the icon in the icon library. Effective when tag is standard_icon. Refer to the enumeration values for icons.




└ └ color

	

No

	

String

	

/

	

The color of the icon. Supports setting the color of linear and face icons (i.e., icons with outlined or filled at the end of the token). Effective when tag is standard_icon. Refer to the enumeration values for colors.




└ └ img_key

	

No

	

String

	

/

	

The image key of the custom prefix icon. Effective when tag is custom_icon.How to obtain the icon key: call the Upload Image interface to upload an image for sending messages, and obtain the image key in the return value.




└ value

	

Yes

	

String

	

/

	

Custom option callback value. When the user clicks on the option of the interactive component, the value of value will be returned to the server receiving the callback data. You can subsequently process business logic based on the received value from the server.

Note: Within the same selection component, the value of each option must be unique, otherwise the option clicked by the user cannot be identified.

Callback Structure

When the user clicks the submit button of the form container, the request address you configured in the developer background will receive the callback data of the following structure. For more information, see Form container and Card callback communication (callback structure).

  {
      "schema": "2.0",
      "header": { 
          "event_id": "f7984f25108f8137722bb63cee927e66",
          "token": "066zT6pS4QCbgj5Do145GfDbbagCHGgF",
          "create_time": "1603977298000000",
          "event_type": "card.action.trigger",
          "tenant_key": "xxxxxxx",
          "app_id": "cli_xxxxxxxx"
      },
      "event":{
          "operator": {
              "tenant_key": "xxxxxxx", 
              "user_id": "xxxxxxx",    
              "open_id": "ou_xxx"     
          },
          "token": "c-xxxx",
          "action": { // The value of the return interaction configured by the "submit" button of the form container itself
              "value": {
                  "key": "value"
               },
              "tag": "button",
              "name":"form_submit", // In the form container, the return attribute of the "submit" button component itself
              "form_value":  { // The return value of each component in the form container
              "multi_select_departments":[ // The custom identifier of the drop-down selection-multi-select component in the form container
                     "selectDemo1", // The value of options.value, used to judge which option the user chooses
                     "selectDemo2"               
                 ] 
              }
          },
          "host": "im_message", 
          "context": {
              "open_message_id":"om_xxx",
              "open_chat_id":"oc_xxx"           
          }  
      }
  }
Sample code

By replacing the value in the following JSON 2.0 structure example code with the actual user's open_id, the card effect as shown in the figure can be achieved:

{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "form",
                "elements": [
                    {
                        "tag": "column_set",
                        "flex_mode": "none",
                        "horizontal_spacing": "default",
                        "background_style": "default",
                        "columns": [
                            {
                                "tag": "column",
                                "elements": [
                                    {
                                        "tag": "multi_select_person",
                                        "type": "default",
                                        "name": "multi_select_users",
                                        "placeholder": {
                                            "tag": "plain_text",
                                            "content": "请选择"
                                        },
                                        "width": "fill",
                                        "required": true,
                                        "disabled": false,
                                        "selected_values": [
                                            "ou_48d0958ee4b2ab3eaf0b5f6c968xxxxx"
                                        ],
                                        "behaviors": [
                                            {
                                                "type": "callback",
                                                "value": {
                                                    "select_static": "click"
                                                }
                                            }
                                        ],
                                        "options": [
                                            {
                                                "value": "ou_48d0958ee4b2ab3eaf0b5f6c968xxxxx"
                                            },
                                            {
                                                "value": "ou_f9d24af786a14340721288cda6axxxxx"
                                            }
                                        ]
                                    }
                                ],
                                "width": "weighted",
                                "weight": 1
                            }
                        ]
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
                                        "name": "Button_lrztw8x3"
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
                                        "name": "Button_lrztw8x4"
                                    }
                                ]
                            }
                        ]
                    }
                ],
                "name": "Form_lrztw8x2"
            }
        ]
    }
}
Previous:Single select user picker
Next:Date picker
Need help with a problem?
Submit feedback

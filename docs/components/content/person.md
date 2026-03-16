# User profile

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/user-profile>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsDisplay componentsUser profile
User profile
Copy Page
Last updated on 2025-06-27
The contents of this article
Precautions
JSON structure
Field descriptions
Example code

The person component supports displaying a person's username and avatar. You can use this component by passing in the person's open_id, user_id, or union_id.

This document introduces the JSON 2.0 structure of the personnel component. To view the historical JSON 1.0 structure, refer to Person.

Precautions

If you are using a specified application to send a card containing the person component, you must ensure that the application has access to the user IDs. Otherwise, the person component on the card will not display any person information.

JSON structure

The complete JSON 2.0 structure of the person component is as follows:

{
    "schema": "2.0", // Version of the card JSON structure. The default is 1.0. To use the JSON 2.0 structure, 2.0 must be explicitly declared.
    "body": {
        "elements": [
            {
                "tag": "person",
                "element_id": "custom_id", // Unique identifier for the operation component. New attribute in JSON 2.0. Used to specify the component when calling related interfaces. Customizable by developers.
                "margin": "0px 0px 0px 0px", // Margin of the component. New attribute in JSON 2.0. Default value is "0", supported range is [-99, 99]px.
                "size": "extra_small", // Size of the personnel avatar. Default value is medium.
                "user_id": "ou_4a136bca010747fc3bd7b6f8f4cabcef", // Personnel's ID.
                "show_avatar": true, // Whether to display the personnel's avatar. Default is true.
                "show_name": false, // Whether to display the personnel's username. Default is false.
                "style": "normal" // Display style of the personnel component. Optional values are normal (default style) and capsule (capsule style).
            }
        ]
    }
}
Field descriptions

The field descriptions for the person component are as follows.

Parameter	Required	Type	Default Value	Description


tag

	

Yes

	

String

	

person

	

The label of the component, fixed to person for the person component.




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



size

	

No

	

String

	

medium

	

Size of the person's avatar. Available options:

extra_small: Extra small size
small: Small size
medium: Medium size
large: Large size



show_avatar

	

No

	

Boolean

	

true

	

Whether to display the user's avatar.




show_name

	

No

	

Boolean

	

false

	

Whether to display the user's name.




style

	

No

	

String

	

normal

	

The display style of the user component. Available options are:

normal: Default style
capsule: Capsule style



user_id

	

Yes

	

String

	

Empty

	

ID of the person. Available options include:The ID of the person. Available options are:

Person's Open ID: Identifies a user's identity within a specific application. The same user has different Open IDs in different applications. For details, refer to How to Obtain Open ID
Person's Union ID: Identifies a user's identity under a specific app developer. The same user has the same Union ID across all apps under the same developer but different Union IDs under different developers. With the Union ID, app developers can link the same user's identity across multiple apps. For details, refer to How to Obtain Union ID
Person's User ID: Identifies a user's identity within a specific tenant. The same user has different User IDs in Tenant A and Tenant B. Within the same tenant, a user's User ID remains consistent across all applications (including store apps). User ID is primarily used to synchronize user data across different applications. For details, refer to How to Obtain User ID
Example code

Replace the user_id in the following JSON 2.0 structure example code with the actual user ID to achieve the card effect shown in the example below:

{
  "schema": "2.0",
  "header": {
      "template": "blue",
      "title": {
          "content": "人员示例",
          "tag": "plain_text"
      }
  },
  "body": {
      "elements": [
          {
              "tag": "markdown",
              "content": "**extra_small 尺寸，默认样式**"
          },
          {
              "tag": "person",
              "size": "extra_small",
              "user_id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef",
              "show_avatar": true,
              "show_name": true,
              "style": "normal"
          },
          {
              "tag": "markdown",
              "content": "**small 尺寸，胶囊样式**"
          },
          {
              "tag": "person",
              "size": "small",
              "user_id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef",
              "show_avatar": true,
              "show_name": true,
              "style": "capsule"
          },
          {
              "tag": "markdown",
              "content": "**medium 尺寸，默认样式**"
          },
          {
              "tag": "person",
              "size": "medium",
              "user_id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef",
              "show_avatar": true,
              "show_name": true,
              "style": "normal"
          },
          {
              "tag": "markdown",
              "content": "**large 尺寸，胶囊样式**"
          },
          {
              "tag": "person",
              "size": "large",
              "user_id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef",
              "show_avatar": true,
              "show_name": true,
              "style": "capsule"
          }
      ]
  }
}
Previous:Divider
Next:User list
Need help with a problem?
Submit feedback

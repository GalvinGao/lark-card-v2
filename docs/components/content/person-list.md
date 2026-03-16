# User list

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/user-list>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsDisplay componentsUser list
User list
Copy Page
Last updated on 2025-06-27
The contents of this article
Precautions
JSON structure
Field descriptions
Example code

The person list component supports displaying multiple users' usernames and avatars. You can use this component by passing in a person's open_id, user_id, or union_id.

This document introduces the JSON 2.0 structure of the personnel list component. To view the historical JSON 1.0 structure, refer to Person List.

Precautions

If you want to send a card containing a person list component using a specified application, you must ensure that the application has access to the user IDs. Otherwise, the person list component in the card will not be able to display person information.

JSON structure

The complete JSON 2.0 structure of the personnel list component is as follows:

{
  "schema": "2.0", // Version of the card JSON structure. The default is 1.0. To use the JSON 2.0 structure, 2.0 must be explicitly declared.
  "body": {
    "elements": [
      {
        "tag": "person_list",
        "element_id": "custom_id", // Unique identifier for the operation component. New attribute in JSON 2.0. Used to specify the component when calling related interfaces. Customizable by developers.
        "margin": "0px 0px 0px 0px", // Margin of the component. New attribute in JSON 2.0. Default value is "0", supported range is [-99,99]px.
        "drop_invalid_user_id": false, // Whether to ignore invalid user IDs in the personnel list. Default is false, meaning that an error will be reported and the list of invalid user IDs will be returned if there are invalid user IDs.
        "lines": 1, // Maximum number of display lines. By default, there is no restriction on the maximum number of display lines.
        "show_name": true, // Whether to display the usernames corresponding to the personnel.
        "show_avatar": true, // Whether to display the avatars corresponding to the personnel.
        "size": "large", // Size of the personnel avatars.
        "persons": [
          // Personnel list. The IDs of personnel support open_id, user_id, union_id
          {
            "id": "ou_0fdb0e7663af7128e7d9f8adeb2abcef"
          },
          {
            "id": "ou_0fdb0e7663af7128e7d9f8adeb2abcef"
          }
        ],
        "icon": {
          // Prefix icon.
          "tag": "standard_icon", // Type of icon.
          "token": "chat-forbidden_outlined", // Token of the icon. Only effective when tag is standard_icon.
          "color": "orange", // Color of the icon. Only effective when tag is standard_icon.
          "img_key": "img_v2_38811724" // Key of the image. Only effective when tag is custom_icon.
        },
        "ud_icon": {
          // Prefix icon from the icon library. When both icon and ud_icon are set, icon takes precedence.
          "token": "chat-forbidden_outlined", // Token of the icon.
          "style": {
            "color": "red" // Color of the icon. Supports setting the color of both line and solid icons (i.e., icons with `outlined` or `filled` at the end of the token).
          }
        }
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

	

person_list

	

The tag of the component, fixed to person_list for the person list component.




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



drop_invalid_user_id

	

No

	

Boolean

	

false

	

Indicates whether to ignore invalid user IDs when there are invalid IDs in the user list. Defaults to false, meaning that if invalid user IDs are present, an error will be raised and the list of invalid IDs will be returned.




lines

	

No

	

Int

	

/

	

Maximum number of display lines, default is unlimited.




show_name

	

No

	

Boolean

	

true

	

Whether to display the person's username.




show_avatar

	

No

	

Boolean

	

false

	

Whether to display the person's avatar.

Note:
If the user's name is not displayed, and when there are multiple IDs in the persons field, the user list will be shown in a "hulu string" style.




size

	

No

	

String

	

medium

	

Size of the person's avatar. Possible values:

extra_small: Extra small sizesmall: Small sizemedium: Medium sizelarge: Large size




persons

	

Yes

	

Array

	

/

	

List of persons.




└ id

	

Yes

	

String

	

empty

	

ID of the person. Possible values include:

Person's Open ID: Identifies a user's identity within an application. The same user has different Open IDs in different applications. See How to Obtain Open ID
Person's Union ID: Identifies a user's identity under an application developer. The same user has the same Union ID across applications under the same developer, but different Union IDs under different developers. Through Union ID, app developers can link the same user's identities across multiple apps. See How to Obtain Union ID
Person's User ID: Identifies a user's identity within a tenant. The same user has different User IDs in tenant A and tenant B. Within the same tenant, a user's User ID remains consistent across all applications (including store apps). User ID is mainly used to bridge user data across different apps. See How to Obtain User ID



icon

	

No

	

Object

	

/

	

Add an icon as a text prefix. Supports custom or icon library usage.




└ tag

	

No

	

String

	

/

	

Icon type label. Possible values:

standard_icon: Use an icon from the icon library.
custom_icon: Use a custom image as the icon.



└ token

	

No

	

String

	

/

	

The token of the icon from the icon library. Effective when tag is standard_icon. See enumeration values in Icon Library.




└ color

	

No

	

String

	

/

	

The color of the icon. Supports setting colors for line and surface icons (i.e., tokens ending in outlined or filled). Effective when tag is standard_icon. See enumeration values in Color Enumeration.




└ img_key

	

No

	

String

	

/

	

The image key for a custom prefix icon. Effective when tag is custom_icon.To obtain the icon key: Call the Upload Image API, upload an image for messaging, and retrieve the image key from the response.




ud_icon

	

No

	

Object

	

/

	

Adds an existing icon from the icon library.

Note: Only one icon can be configured for a user component. If both icon and ud_icon are configured, only icon will be effective.




└ token

	

No

	

String

	

/

	

The token of the icon from the icon library. For enumerated values, see Icon Library.




└ style

	

No

	

Object

	

/

	

The style of the icon. Supports custom icon colors.




└└ color

	

No

	

String

	

/

	

The color of the icon. Supports setting the color for both outlined and filled icons (i.e., icons with "outlined" or "filled" at the end of the token). For enumerated values, see Color Enumerations.

Note: The customization tool currently does not support custom icon colors.

Example code

Replace the user_id in the following example code with the actual user ID to achieve the card effect shown in the example image:

{
  "schema": "2.0",
  "header": {
    "template": "blue",
    "title": {
      "content": "人员列表示例",
      "tag": "plain_text"
    }
  },
  "body": {
    "elements": [
      {
        "tag": "markdown",
        "content": "仅名字："
      },
      {
        "tag": "person_list",
        "persons": [
          {
            "id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef"
          },
          {
            "id": "ou_f9d24af786a14340721288cda6aabcef"
          },
          {
            "id": "ou_b824f85713725c632e78887dc7fabcef"
          }
        ]
      },
      {
        "tag": "markdown",
        "content": "名字+头像："
      },
      {
        "tag": "person_list",
        "show_name": true,
        "show_avatar": true,
        "lines": 2,
        "size": "small",
        "persons": [
          {
            "id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef"
          },
          {
            "id": "ou_f9d24af786a14340721288cda6aabcef"
          },
          {
            "id": "ou_b824f85713725c632e78887dc7fabcef"
          }
        ]
      },
      {
        "tag": "markdown",
        "content": "名字+头像+icon："
      },
      {
        "tag": "person_list",
        "show_name": true,
        "show_avatar": true,
        "lines": 2,
        "size": "small",
        "persons": [
          {
            "id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef"
          },
          {
            "id": "ou_f9d24af786a14340721288cda6aabcef"
          },
          {
            "id": "ou_b824f85713725c632e78887dc7fabcef"
          }
        ],
        "icon": {
          "tag": "standard_icon",
          "token": "group_outlined",
          "color": "blue"
        }
      },
      {
        "tag": "markdown",
        "content": "名字葫芦串："
      },
      {
        "tag": "person_list",
        "persons": [
          {
            "id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef"
          },
          {
            "id": "ou_f9d24af786a14340721288cda6aabcef"
          },
          {
            "id": "ou_b824f85713725c632e78887dc7fabcef"
          },
          {
            "id": "ou_48d0958ee4b2ab3eaf0b5f6c968abcef"
          },
          {
            "id": "ou_f9d24af786a14340721288cda6aabcef"
          },
          {
            "id": "ou_b824f85713725c632e78887dc7fabcef"
          }
        ],
        "size": "small",
        "show_avatar": true,
        "show_name": false
      }
    ]
  }
}
Previous:User profile
Next:Chart
Need help with a problem?
Submit feedback

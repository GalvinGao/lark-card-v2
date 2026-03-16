# Plain text

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/plain-text>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsDisplay componentsPlain text
Plain text
Copy Page
Last updated on 2025-06-27
The contents of this article
JSON structure
Field description
Field descriptions
Example Code
plain_text type example
lark_md type example
Markdown Syntax Supported by lark_md
Defining Different Font Sizes for Mobile and Desktop

The plain text component of the card supports adding plain text and prefix icons, and sets the display styles such as text size, color, alignment, etc.

This document introduces the JSON 2.0 structure of the plain text component. To view the historical JSON 1.0 structure, refer to plain text.

JSON structure

The complete JSON 2.0 structure of the plain text component is as follows:

{
  "schema": "2.0", // The version of the card JSON structure. The default is 1.0. To use the JSON 2.0 structure, you must explicitly declare 2.0.
  "body": {
    "elements": [
      {
        "tag": "div",
        "element_id": "custom_id", // Unique identifier for the operation component. New attribute in JSON 2.0. Used to specify the component in the related interface calls. Needs to be customized by the developer.
        "margin": "0px 0px 0px 0px", // Margin of the component, default value is "0". New attribute in JSON 2.0. Supported range is [-99,99]px.
        "width": "fill", // Text width. New attribute in JSON 2.0. Supports "fill", "auto", "[16,999]px". The default value is fill.
        "text": { // Configures plain text information.
          "tag": "plain_text", // Tag for text type. Possible values: plain_text and lark_md.
          "content": "", // Text content. When the tag is lark_md, it supports text content with partial Markdown syntax.
          "text_size": "normal", // Text size. The default value is normal. Supports customizing different font sizes on mobile and desktop.
          "text_color": "default", // Text color. Only effective when the tag is plain_text. The default value is default.
          "text_align": "left", // Text alignment. The default value is left.
          "lines": 2 // The maximum number of lines to display for the content. Content exceeding the set number of lines will be truncated with ...
        },
        "icon": {
          // Prefix icon.
          "tag": "standard_icon", // Icon type.
          "token": "chat-forbidden_outlined", // Token of the icon. Only effective when the tag is standard_icon.
          "color": "orange", // Icon color. Only effective when the tag is standard_icon.
          "img_key": "img_v2_38811724" // Key of the image. Only effective when the tag is custom_icon.
        }
      }
    ]
  }
}
Field description

The field description for the plain text component is as follows:

Field descriptions

The field descriptions for the plain text component are shown in the table below.

Name	Required	Type	Default	Description


tag

	

Yes

	

String

	

/

	

Tag of the component. The tag for the plain text component is div.




element_id

	

No

	

String

	

Empty

	

Unique identifier for the operation component. Used to specify the component in the calls to the component-related interface. The value of this field must be globally unique within the same card. Only letters, numbers, and underscores are allowed, must start with a letter, and must not exceed 20 characters.




margin

	

No

	

String

	

0

	

Margin of the component. New attribute in JSON 2.0. The value range is [-99,99]px. Optional values:

Single value, such as "10px", indicates that all four margins of the component are 10 px.
Double values, such as "4px 0", indicate that the top and bottom margins of the component are 4 px and the left and right margins are 0 px. Use spaces to separate values (units can be omitted when the margin is 0).
Multiple values, such as "4px 0 4px 0", indicate that the top, right, bottom, and left margins of the component are 4px, 12px, 4px, and 12px, respectively. Use spaces to separate values.



width

	

No

	

String

	

fill

	

Text width. New attribute in JSON 2.0. Optional values:

fill: The text width will be consistent with the component width, filling the component.
auto: The text width will adapt to the length of the text content itself.
[16,999]px: Custom text width.



text

	

No

	

Object

	

/

	

Configures the plain text information of the card.




└ tag

	

Yes

	

String

	

plain_text

	

Tag for text type. Optional values:

plain_text: Plain text content or emojis
lark_md: Text content that supports partial Markdown syntax. For details, refer to Markdown syntax supported by lark_md below.

Note: In the Lark card building tool, only the plain_text type of plain text component is supported. You can use the rich text component to add Markdown formatted text.




└ content

	

Yes

	

String

	

/

	

Text content. When the tag is lark_md, it supports text content with some Markdown syntax. See below for Markdown syntax supported by lark_md.




└ text_size

	

No

	

String

	

normal

	

Text size. Possible values are shown below. If you enter other values, the card will display the font size corresponding to the normal field. You can also define different font sizes for mobile and desktop versions; see below for Defining Different Font Sizes for Mobile and Desktop Versions.

heading-0: Extra large title (30px)
heading-1: First-level title (24px)
heading-2: Second-level title (20px)
heading-3: Third-level title (18px)
heading-4: Fourth-level title (16px)
heading: Title (16px)
normal: Text (14px)
notation: Supplementary information (12px)
xxxx-large: 30px
xxx-large: 24px
xx-large: 20px
x-large: 18px
large: 16px
medium: 14px
small: 12px
x-small: 10px



└ text_color

	

No

	

String

	

default

	

Text color. Effective only when tag is plain_text. Possible values:

default: Black in client light theme mode; white in client dark theme mode
Enumeration values for colors. See Color Enumeration Values



└ text_align

	

No

	

String

	

left

	

Text alignment. Possible values:

left: Left-aligned
center: Center-aligned
right: Right-aligned



└ lines

	

No

	

Int

	

/

	

Maximum number of lines to display content, content exceeding the set lines is abbreviated with ....




icon

	

No

	

Object

	

/

	

Adds an icon as a text prefix icon. Supports using custom or library icons.




└ tag

	

No

	

String

	

/

	

Label for icon type. Possible values:

standard_icon: Use icons from the library.
custom_icon: Use a custom image as an icon.



└ token

	

No

	

String

	

/

	

Token of the icon from the library. Effective when tag is standard_icon. Enumeration values see Icon Library.




└ color

	

No

	

String

	

/

	

The color of the icon. Supports setting colors for linear and filled icons (i.e., icons whose token ends with outlined or filled). Effective when tag is standard_icon. Enumeration values see Color Enumeration Values.




└ img_key

	

No

	

String

	

/

	

Custom prefix icon image key. Effective when tag is custom_icon.

Method to obtain the icon key: Call the Upload Image interface, upload an image for sending messages, and obtain the image_key from the return value.

Example Code
plain_text type example

The following example code of the JSON 2.0 structure can achieve the card effect shown in the figure below:

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "div",
        "text": {
          "tag": "plain_text",
          "content": "这是示例文本。",
          "text_size": "normal",
          "text_align": "center",
          "text_color": "default"
        },
        "icon": {
          "tag": "standard_icon",
          "token": "reply-cn_filled",
          "color": "blue"
        },
        "margin": "0px 0px 0px 0px",
        "element_id": "plaintext01"
      }
    ]
  }
}
lark_md type example

The following example code of the JSON 2.0 structure can achieve the card effect shown in the figure below:

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "div",
        "text": {
          "tag": "plain_text",
          "content": "text-lark_md",
          "lines": 1
        },
        "fields": [
          {
            "is_short": false,
            "text": {
              "tag": "lark_md",
              "content": "<a>https://open.larksuite.com</a>"
            }
          },
          {
            "is_short": false,
            "text": {
              "tag": "lark_md",
              "content": "ready\nnew line"
            }
          },
          {
            "is_short": false,
            "text": {
              "tag": "lark_md",
              "content": "*Italic*"
            }
          },
          {
            "is_short": false,
            "text": {
              "tag": "lark_md",
              "content": "**Bold**"
            }
          },
          {
            "is_short": false,
            "text": {
              "tag": "lark_md",
              "content": "~~delete line~~"
            }
          },
          {
            "is_short": false,
            "text": {
              "tag": "lark_md",
              "content": "<at id=all></at>"
            }
          }
        ]
      }
    ]
  }
}
Markdown Syntax Supported by lark_md
Feature	Syntax	Effect


New Line

	

First line\nSecond line

	

First line

Second line




Italic

	

*Italic*

	

Italic




Bold

	

**Bold** or __Bold__

	

Bold




Strikethrough

	

~~Strikethrough~~

	

Strikethrough




Text Link

	

[Text Link](https://www.larksuite.com)

	

Text Link




Hyperlink

	

<a href='https://open.larksuite.com'></a>

	

https://open.larksuite.com




@ Mention

	

<at id=all></at>

<at id={{open_id}}></at>

<at id={{user_id}}></at>

<at email=test@email.com></at>

	

@all@test




Colored Text

	

<font color=red>Red</font>Note: For color enumerations, see Color Enumeration Values.

	

Red




Emoji

	

😁😢🌞💼🏆❌✅

Tip: You can directly copy emojis. For more emojis, refer to Emoji.

	

😁😢🌞💼🏆❌✅




Lark Emoticons

	

:OK:

Tip: To view the list of emoticons, see Emoticon Descriptions.

	




Tags

	

<text_tag color='neutral'> neutral </text_tag>

Color enumerations include: neutral, blue, turquoise, lime, orange, violet, indigo, wathet, green, yellow, red, purple, carmine

	

Defining Different Font Sizes for Mobile and Desktop

In both plain text and rich text components, you can define different font sizes for the same text on mobile and desktop platforms. The related field descriptions are as follows:

Field	Required	Type	Default Value	Description


text_size

	

No

	

Object

	

/

	

Text size. Here you can customize different font sizes for mobile and desktop platforms.




└ custom_text_size_name

	

No

	

Object

	

/

	

Custom font size. You need to define the name of this field, such as cus-0, cus-1, etc.




└└ default

	

No

	

String

	

/

	

The font size attribute effective on old version of Lark clients where differentiated font sizes can't be configured. It's recommended to fill this field. Possible values are shown below.

heading-0: Extra large title (30px)
heading-1: First-level title (24px)
heading-2: Second-level title (20px)
heading-3: Third-level title (18px)
heading-4: Fourth-level title (16px)
heading: Title (16px)
normal: Text (14px)
notation: Supplementary information (12px)
xxxx-large: 30px
xxx-large: 24px
xx-large: 20px
x-large: 18px
large: 16px
medium: 14px
small: 12px
x-small: 10px



└└ pc

	

No

	

String

	

/

	

Desktop platform font size. Possible values are shown below.

heading-0: Extra large title (30px)
heading-1: First-level title (24px)
heading-2: Second-level title (20px)
heading-3: Third-level title (18px)
heading-4: Fourth-level title (16px)
heading: Title (16px)
normal: Text (14px)
notation: Supplementary information (12px)
xxxx-large: 30px
xxx-large: 24px
xx-large: 20px
x-large: 18px
large: 16px
medium: 14px
small: 12px
x-small: 10px



└└ mobile

	

No

	

String

	

/

	

Mobile platform text size. Possible values are shown below.

Note: Some mobile platform font size enumeration values differ from those on the PC platform; please note the distinctions when using.

heading-0: Extra large title (26px)
heading-1: First-level title (24px)
heading-2: Second-level title (20px)
heading-3: Third-level title (17px)
heading-4: Fourth-level title (16px)
heading: Title (16px)
normal: Text (14px)
notation: Supplementary information (12px)
xxxx-large: 26px
xxx-large: 24px
xx-large: 20px
x-large: 18px
large: 17px
medium: 14px
small: 12px
x-small: 10px

The specific steps are as follows:

In the global action settings of the card JSON code, configure the style field in the config section, and add custom font sizes:

{
  "config": {
    "style": { // Add and configure the style field here.
      "text_size": { // Add custom font sizes for both mobile and desktop platforms, including fallback sizes. This is used to set the font size property in the component JSON. Multiple custom font size objects can be added.
        "cus-0": {
          "default": "medium", // The font size attribute effective on old version of Lark clients where differentiated font sizes can't be configured. Optional.
          "pc": "medium", // Desktop platform font size.
          "mobile": "large" // Mobile platform font size.
        },
        "cus-1": {
          "default": "medium", // The font size attribute effective on old version of Lark clients where differentiated font sizes can't be configured. Optional.
          "pc": "normal", // Desktop platform font size.
          "mobile": "x-large" // Mobile platform font size.
        }
      }
    }
  }
}

Apply the custom font size in the text_size attribute of the plain text component or rich text component. Below is an example of applying a custom font size in a plain text component:

{
  "i18n_elements": {
    "zh_cn": [
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
                "tag": "div",
                "text": {
                  "tag": "plain_text",
                  "content": "This is an example of plain text.",
                  "text_size": "cus-0", // The custom font size is applied here.
                  "text_align": "center",
                  "text_color": "default"
                },
                "icon": {
                  "tag": "standard_icon",
                  "token": "app-default_filled",
                  "color": "blue"
                }
              }
            ],
            "width": "weighted",
            "weight": 1
          }
        ]
      }
    ],
    "i18n_header": {}
  }
}
Previous:Title
Next:Rich text (Markdown)
Need help with a problem?
Submit feedback

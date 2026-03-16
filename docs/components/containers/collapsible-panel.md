# Collapsible panel

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/containers/collapsible-panel>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentscontainersCollapsible panel
Collapsible panel
Copy Page
Last updated on 2025-06-27
The contents of this article
Precautions
Nesting rules
Component attributes
JSON Structure
Field Description
Example code

The collapsible panel allows for the folding of secondary information, such as notes and longer texts, within a card to highlight the primary information.

This document introduces the JSON 2.0 structure of the collapsible panel component. To view the historical JSON 1.0 structure, refer to collapsible panel.

Precautions
Collapsible panels can only be used by writing card JSON code, and are not yet supported in card-building tools.
Container components support a maximum of five layers of nested components. It is recommended to avoid nesting multiple layers in collapsible panels. Multi-layer nesting compresses the display space and affects the presentation of the card.
Nesting rules

Collapsible panels do not support embedding form (form) components.

Component attributes

This section introduces the attributes of the collapsible panel.

JSON Structure

The complete JSON 2.0 structure of the collapsible panel component is as follows:

{
  "schema": "2.0", // The version of the card JSON structure. The default is 1.0. To use the JSON 2.0 structure, you must explicitly declare 2.0.
  "body": {
    "elements": [
      {
        "tag": "collapsible_panel", // The tag of the collapsible panel.
        "element_id": "custom_id", // The unique identifier of the interactive component. New attribute in JSON 2.0. Used to specify the component in related API calls. Customizable by the developer.
        "direction": "vertical", // The arrangement direction of components within the panel. New attribute in JSON 2.0. Optional values: "vertical" (vertical arrangement), "horizontal" (horizontal arrangement). The default is "vertical".
        "vertical_spacing": "8px", // The vertical spacing between components within the panel. New attribute in JSON 2.0. Optional values: "small" (4px), "medium" (8px), "large" (12px), "extra_large" (16px), or [0,99]px.
        "horizontal_spacing": "8px", // The horizontal spacing between components within the panel. New attribute in JSON 2.0. Optional values: "small" (4px), "medium" (8px), "large" (12px), "extra_large" (16px), or [0,99]px.
        "vertical_align": "top", // The vertical alignment of components within the panel. New attribute in JSON 2.0. The default value is top.
        "horizontal_align": "left", // The horizontal alignment of components within the panel. New attribute in JSON 2.0. The default value is left.
        "padding": "8px 8px 8px 8px", // The padding of the collapsible panel. New attribute in JSON 2.0. The supported range is [0,99]px.
        "margin": "0px 0px 0px 0px", // The margin of the collapsible panel. New attribute in JSON 2.0. The default value is "0", and the supported range is [-99,99]px.
        "expanded": true, // Whether the panel is expanded. The default value is false.
        "background_color": "grey", // The background color of the collapsible panel. The default is transparent.
        "header": {
          // The title settings of the collapsible panel.
          "title": {
            // Title text settings. Supports plain_text and markdown.
            "tag": "markdown",
            "content": "**Panel title text**"
          },
          "background_color": "grey", // The background color of the title area. The default is transparent.
          "vertical_align": "center", // The vertical alignment of the title area.
          "padding": "4px 0px 4px 8px", // The padding of the title area.
          "position": "top", // The position of the title area.
          "width": "auto", // The width of the title area. The default value is fill.
          "icon": {
            // Prefix icon of the title
            "tag": "standard_icon", // The type of the icon.
            "token": "chat-forbidden_outlined", // The token of the icon in the icon library. Effective when the tag is standard_icon.
            "color": "orange", // The color of the icon. Effective when the tag is standard_icon.
            "img_key": "img_v2_38811724", // The image key of the custom prefix icon. Effective when the tag is custom_icon.
            "size": "16px 16px" // The size of the icon. The default value is 10px 10px.
          },
          "icon_position": "follow_text", // The position of the icon. The default value is right.
          "icon_expanded_angle": -180 // The rotation angle of the icon when the collapsible panel is expanded. Positive values indicate clockwise rotation, and negative values indicate counterclockwise rotation. The default value is 180.
        },
        "border": {
          // Border settings. The default is no border.
          "color": "grey", // The color of the border.
          "corner_radius": "5px" // The corner radius.
        },
        "elements": [
          // Here you can add the JSON structure of various components. Form components are not supported for now.
          {
            "tag": "markdown",
            "content": "A long text"
          }
        ]
      }
    ]
  }
}
Field Description

The table below describes the fields of the collapsible panel:

Name	Required	Type	Default	Description


tag

	

No

	

string

	

/

	

The tag of the component. The collapsible panel has a fixed value of collapsible_panel.




expanded

	

No

	

Boolean

	

false

	

Whether the panel is expanded. Optional values:

true: The panel is expanded

false: The panel is collapsed. The default is collapsed.



element_id

	

No

	

String

	

empty

	

The unique identifier of the operation component. A new attribute in JSON 2.0. Used to specify the component when calling component-related interfaces. The value of this field must be globally unique within the same card. Only letters, numbers, and underscores are allowed, must start with a letter, and must not exceed 20 characters.




margin

	

No

	

String

	

0px

	

The outer margin of the container. The value range is [-99,99]px. Optional values:

Single value, such as "10px", indicates that the four outer margins of the container are all 10 px.
Double value, such as "4px 0", indicates that the top and bottom outer margins of the container are 4 px, and the left and right outer margins are 0 px. Use spaces to separate (no unit required when the margin is 0).
Multiple values, such as "4px 0 4px 0", indicate that the top, right, bottom, and left outer margins of the container are 4px, 12px, 4px, and 12px, respectively. Use spaces to separate.



horizontal_spacing

	

No

	

String

	

8px

	

The horizontal spacing of components within the container. Optional values:

small: Small spacing, 4px
medium: Medium spacing, 8px
large: Large spacing, 12px
extra_large: Extra-large spacing, 16px
Specific values, such as 20px. The value range is [0,99]px



horizontal_align

	

No

	

String

	

left

	

The horizontal alignment of components within the container. Optional values:

left: Left-aligned
center: Center-aligned
right: Right-aligned



vertical_spacing

	

No

	

String

	

12px

	

Horizontal spacing between components within the container. Optional values:

small: small spacing, 4px
medium: medium spacing, 8px
large: large spacing, 12px
extra_large: extra large spacing, 16px
Specific values, such as 20px. The value range is [0,99]px



vertical_align

	

No

	

String

	

top

	

Vertical alignment of components within the container. Optional values:

top: align to the top
center: align to the center
bottom: align to the bottom



direction

	

No

	

String

	

vertical

	

Container arrangement direction. Optional values:

vertical: vertical arrangement
horizontal: horizontal arrangement



padding

	

No

	

String

	

0px

	

Padding of the container. The value range is [-99,99]px. Optional values:

Single value, such as "10px", indicating that the four outer margins of the container are all 10px.
Double value, such as "4px 0", indicating that the top and bottom outer margins of the container are 4px, and the left and right outer margins are 0px. Use spaces to separate (units can be omitted when the margin is 0).
Multiple values, such as "4px 0 4px 0", indicating that the top, right, bottom, and left outer margins of the container are 4px, 12px, 4px, and 12px respectively. Use spaces to separate.



background_color

	

No

	

String

	

None

	

Background color of the folding panel, default is transparent. For enumeration values, see Color Enumeration Values.




header

	

Yes

	

Object

	
	

Title settings of the folding panel.




└ title

	

No

	

Object

	
	

Title text settings.




└└ tag

	

Yes

	

String

	

None

	

Tag for text type. Optional values:

plain_text: plain text content
markdown: rich text content. For supported Markdown syntax, see Rich Text Component.



└└ content

	

No

	

String

	

Empty

	

The content of the collapsible panel title.




└ background_color

	

No

	

String

	

Empty

	

The background color setting of the collapsible panel title area, default is transparent. For enumeration values, see Color Enumeration.

Note: If you do not set this field, the background color of the collapsible panel title area will be determined by the background_color field.




└ width

	

No

	

String

	

fill

	

The width of the title element. New property in JSON 2.0. Supported in Lark client version 7.32 and above.

fill: Title and collapsible panel are of equal width
auto: Title adapts to text length
auto_when_fold: Title adapts to text length only when the collapsible panel is folded



└ vertical_align

	

No

	

String

	

center

	

The vertical alignment of the title area. Possible values:

top: Title area is vertically aligned to the top of the panel area

center: Title area is vertically aligned to the center of the panel area

bottom: Title area is vertically aligned to the bottom of the panel area



└ padding

	

No

	

String

	

0

	

The padding of the title area. The value range is [-99,99]px. Optional values:

Single value, such as "10px", means that all four margins of the title area are 10 px.
Double values, such as "4px 0", means the top and bottom margins of the title area are 4 px, and the left and right margins are 0 px. Use a space to separate (units are optional when the margin is 0).
Multiple values, such as "4px 0 4px 0", means the top, right, bottom, and left margins of the title area are 4px, 12px, 4px, and 12px respectively. Use a space to separate.



└ icon

	

No

	

Object

	

/

	

Add an icon as a prefix or suffix icon for the title. Supports custom or library icons. Example code is as follows:


"icon": {
  "tag": "standard_icon",
  "token": "down-small-ccm_outlined",
  "color": "",
  "size": "16px 16px"
}



└└ tag

	

No

	

String

	

/

	

The tag of the icon type. Possible values:


standard_icon: Use icons from the library

custom_icon: Use a custom image as the icon



└ └ token

	

No

	

String

	

/

	

The token of the icon in the icon library. Effective when tag is standard_icon. For enumeration values, see Icon Library.




└└ color

	

No

	

String

	

/

	

The color of the icon. Supports setting the color of linear and solid icons (i.e., icons with outlined or filled at the end of the token). Effective when tag is standard_icon. For enumeration values, see Color Enumeration.




└└ img_key

	

No

	

String

	

/

	

The image key of the custom prefix icon. Effective when tag is custom_icon. How to obtain the icon key: Call the Upload Image API to upload the image used for sending messages, and get the image_key from the return value.




└ └ size

	

No

	

String

	

10px 10px

	

The size of the icon. Supports "[1,999] [1,999]px".




└ icon_position

	

No

	

String

	

right

	

The position of the icon. Optional values:


left: Icon is on the far left side of the title area

right: Icon is on the far right side of the title area

follow_text: Icon is to the right of the text



└ icon_expanded_angle

	

No

	

Number

	

180

	

The angle at which the icon rotates when the collapsible panel is expanded. Positive values are clockwise, and negative values are counterclockwise. Optional values:

-180: Rotate 180 degrees counterclockwise

-90: Rotate 90 degrees counterclockwise

90: Rotate 90 degrees clockwise

180: Rotate 180 degrees clockwise



border

	

No

	

Object

	

Empty

	

Border settings. No border is displayed by default.




└└ color

	

No

	

String

	

/

	

The color of the icon. Supports setting the color for both linear and solid icons (icons with tokens ending in outlined or filled). Effective when tag is standard_icon. For enumeration values, refer to Color Enumeration Values.




└└ img_key

	

No

	

String

	

/

	

The image key for the custom prefix icon. Effective when tag is custom_icon. To obtain the icon key: call the Upload Image API, upload the image to be used in the message, and get the image_key from the response.




└ └ size

	

No

	

String

	

10px 10px

	

The size of the icon. Supports "[1,999] [1,999]px".




└ icon_position

	

No

	

String

	

right

	

The position of the icon. Optional values:


left: Icon on the far left of the title area

right: Icon on the far right of the title area

follow_text: Icon to the right of the text



└ icon_expanded_angle

	

No

	

Number

	

180

	

The angle of rotation for the icon when the panel is expanded. Positive values for clockwise, negative values for counterclockwise. Optional values:


-180: Rotate 180 degrees counterclockwise

-90: Rotate 90 degrees counterclockwise

90: Rotate 90 degrees clockwise

180: Rotate 180 degrees clockwise



border

	

No

	

Object

	

Empty

	

Border settings. No border is displayed by default.




└ color

	

No

	

String

	

grey

	

Border color settings. For enumeration values, refer to Color Enumeration Values.




└ corner_radius

	

No

	

String

	

5px

	

Corner radius settings.




elements

	

No

	

Array

	

Empty

	

JSON structure of each component. Form components are not supported yet.

Example code

The following JSON 2.0 example code can achieve the card effect as shown in the figure below:

{
  "schema": "2.0",
  "header": {
    "template": "yellow",
    "title": {
      "tag": "plain_text",
      "content": "折叠面板展示"
    }
  },
  "body": {
    "elements": [
      {
        "tag": "markdown",
        "content": "下面是一个 默认折叠 的折叠面板组件"
      },
      {
        "tag": "collapsible_panel",
        "expanded": false,
        "header": {
          "title": {
            "tag": "plain_text",
            "content": "面板标题文本"
          },
          "vertical_align": "center",
          "icon": {
            "tag": "standard_icon",
            "token": "down-small-ccm_outlined",
            "color": "",
            "size": "16px 16px"
          },
          "icon_position": "right",
          "icon_expanded_angle": -180
        },
        "border": {
          "color": "grey",
          "corner_radius": "5px"
        },
        "vertical_spacing": "8px",
        "padding": "8px 8px 8px 8px",
        "elements": [
          {
            "tag": "markdown",
            "content": "很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本"
          }
        ]
      },
      {
        "tag": "markdown",
        "content": "下面是一个 标题带背景色 且 默认展开 的折叠面板组件"
      },
      {
        "tag": "collapsible_panel",
        "expanded": true,
        "header": {
          "title": {
            "tag": "markdown",
            "content": "**<font color='white'>面板标题文本</font>**"
          },
          "background_color": "yellow",
          "vertical_align": "center",
          "icon": {
            "tag": "standard_icon",
            "token": "down-small-ccm_outlined",
            "color": "white",
            "size": "16px 16px"
          },
          "icon_position": "right",
          "icon_expanded_angle": -180
        },
        "border": {
          "color": "grey",
          "corner_radius": "5px"
        },
        "vertical_spacing": "8px",
        "padding": "8px 8px 8px 8px",
        "elements": [
          {
            "tag": "markdown",
            "content": "很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本"
          }
        ]
      },
      {
        "tag": "markdown",
        "content": "下面是一个无边框折叠面板组件"
      },
      {
        "tag": "collapsible_panel",
        "expanded": true,
        "header": {
          "title": {
            "tag": "markdown",
            "content": "**面板标题文本**"
          },
          "width": "auto_when_fold",
          "vertical_align": "center",
          "padding": "4px 0px 4px 8px",
          "icon": {
            "tag": "standard_icon",
            "token": "down-small-ccm_outlined",
            "color": "",
            "size": "16px 16px"
          },
          "icon_position": "follow_text",
          "icon_expanded_angle": -180
        },
        "vertical_spacing": "8px",
        "padding": "8px 8px 8px 8px",
        "elements": [
          {
            "tag": "markdown",
            "content": "很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本很长的文本"
          }
        ]
      }
    ]
  }
}
Previous:Interactive container
Next:Title
Need help with a problem?
Submit feedback

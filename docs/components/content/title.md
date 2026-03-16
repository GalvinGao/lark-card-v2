# Last updated on 2025-06-27

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/title>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsDisplay componentsTitle
Title
Copy Page
Last updated on 2025-06-27
The contents of this article
Precautions
Component Attributes
JSON Structure
Field Description
Demo example
Enumerations
Title theme style enumeration
Suffix tag color enumeration
Icon enumeration
Icon color enumeration

The title component of a card supports adding a main title, subtitle, suffix tags, and a title icon.

This document introduces the JSON 2.0 structure of the title component. To view the historical JSON 1.0 structure, refer to Title.

Precautions

Only one title component can be added to the same card.

Component Attributes
JSON Structure

The complete JSON 2.0 structure of the title component is as follows:

{
  "schema": "2.0", // Version of the card JSON structure. Default is 1.0. To use JSON 2.0 structure, you must explicitly declare 2.0.
  "header": {
    "title": {
      // Main title of the card. Required. To configure the title in multiple languages, refer to the multilingual card configuration document.
      "tag": "plain_text", // Text type tag. Optional values: plain_text and lark_md.
      "content": "Example Title" // Title content.
    },
    "subtitle": {
      // Subtitle of the card. Optional.
      "tag": "plain_text", // Text type tag. Optional values: plain_text and lark_md.
      "content": "Example Text" // Subtitle content.
    },
    "text_tag_list": [
      // Suffix tags for the title, up to 3 tags can be set, any excess will not be displayed. Optional.
      {
        "tag": "text_tag",
        "element_id": "custom_id", // Unique identifier for the operation element. Used to specify the element in component-related interface calls. Custom-defined by developers.
        "text": {
          // Tag content
          "tag": "plain_text",
          "content": "Tag 1"
        },
        "color": "neutral" // Tag color
      }
    ],
    "i18n_text_tag_list": {
      // Multilingual suffix tags for the title. Up to 3 tags can be set for each language environment, any excess will not be displayed. Optional. If both original and internationalized fields are configured, the multilingual configuration takes precedence.
      "zh_cn": [],
      "en_us": [],
      "ja_jp": [],
      "zh_hk": [],
      "zh_tw": []
    },
    "template": "blue", // Title theme style color. Supports "blue"|"wathet"|"turquoise"|"green"|"yellow"|"orange"|"red"|"carmine"|"violet"|"purple"|"indigo"|"grey"|"default". Default value is default.
    "icon": { // Prefix icon.
      "tag": "standard_icon", // Icon type.
      "token": "chat-forbidden_outlined", // Token of the icon. Only effective when tag is standard_icon.
      "color": "orange", // Icon color. Only effective when tag is standard_icon.
      "img_key": "img_v2_38811724" // Key of the image. Only effective when tag is custom_icon.
    },
    "padding": "12px 8px 12px 8px" // Padding of the title component. New attribute in JSON 2.0. Default value is "12px", supported range is [0,99]px.
  }
}
Field Description

The field descriptions for the title component are detailed in the following table.

Name	Required	Type	Description


title

	

Yes

	

Object

	

Configure the main title information of the card.

Note: If only the subtitle is configured, it will actually be displayed as the main title.




└ tag

	

Yes

	

String

	

Text type tag. Possible values:

plain_text: Plain text content or emoji
lark_md: Text content supporting partial Markdown syntax. For details, refer to lark_md supported Markdown syntax



└ content

	

No

	

String

	

Main title content of the card. To configure the title in multiple languages, refer to Configure card multilingual content.

Note: The main title content is up to four lines, content exceeding four lines will be truncated with ....




subtitle

	

No

	

Object

	

Configure the subtitle information of the card.

Note: If only the subtitle is configured, it will actually be displayed as the main title.




└ tag

	

Yes

	

String

	

Text type tag. Possible values:

plain_text: Plain text content or emoji
lark_md: Text content supporting partial Markdown syntax. For details, refer to lark_md supported Markdown syntax



└ content

	

No

	

String

	

Subtitle content of the card. To configure the title in multiple languages, refer to Configure card multilingual content.

Note: The subtitle content is up to one line, content exceeding one line will be truncated with ....




text_tag_list

	

No

	

TextTagList

	

Add suffix tags to the title. Up to 3 tags can be added, if the number of configured tags exceeds 3, the first 3 tags will be displayed. The display order of the tags is consistent with the array order.

Note:text_tag_list and i18n_text_tag_list can only configure one of them. If both are configured, only i18n_text_tag_list will take effect.




└ tag

	

Yes

	

String

	

Identifier of the suffix tag. Fixed value: text_tag.




└ element_id

	

No

	

String

	

Unique identifier of the operating component. A new attribute added in JSON 2.0. Used to specify the component when calling the component-related interface. The value of this field must be globally unique within the same card. Only letters, numbers, and underscores are allowed, must start with a letter, and must not exceed 20 characters.




└ text

	

No

	

Text Object

	

Content of the suffix tag. Define the content based on the plain_text mode of the text component.Example value:

"text": {
          "tag": "plain_text",
          "content": "这里是标签"
        }



└ color

	

No

	

String

	

Color of the suffix tag, default is blue. For optional values and example effects, refer to the suffix tag color enumeration below.




i18n_text_tag_list

	

No

	

Object

	

Configure the multilingual properties of the suffix tag by adding the complete suffix tag structure under the required language field. Up to 3 tag contents can be configured for each language. If the number of configured tags exceeds 3, the first 3 tags will be displayed. The display order of the tags is consistent with the array order. Supported multilingual enumeration values are as follows:

zh_cn: Simplified Chinese
en_us: English
ja_jp: Japanese
zh_hk: Traditional Chinese (Hong Kong)
zh_tw: Traditional Chinese (Taiwan)
id_id: Indonesian
vi_vn: Vietnamese
th_th: Thai
pt_br: Portuguese
es_es: Spanish
ko_kr: Korean
de_de: German
fr_fr: French
it_it: Italian
ru_ru: Russian
ms_my: Malay

Example configuration:

"i18n_text_tag_list": {
      "zh_cn": [
        {
          "tag": "text_tag",
          "text": {
            "tag": "plain_text",
            "content": "标签内容"
          },
          "color": "carmine"
        }
      ],
      "en_us": [
        {
          "tag": "text_tag",
          "text": {
            "tag": "plain_text",
            "content": "Tag content"
          },
          "color": "carmine"
        }
      ]
    }

Note:text_tag_list and i18n_text_tag_list can only configure one of them. If both fields are configured, the multilingual configuration takes precedence.




template

	

No

	

String

	

Configure the title theme color. For optional values and example effects, refer to the title theme style enumeration below.




icon

	

No

	

Object

	

Add an icon as a text prefix. Supports custom or icon library usage.




└ tag

	

No

	

String

	

Icon type label. Possible values:

standard_icon: Use an icon from the icon library.
custom_icon: Use a custom image as the icon.



└ token

	

No

	

String

	

The token of the icon from the icon library. Effective when tag is standard_icon. See enumeration values in Icon Library.




└ color

	

No

	

String

	

The color of the icon. Supports setting colors for line and surface icons (i.e., tokens ending in outlined or filled). Effective when tag is standard_icon. See enumeration values in Color Enumeration.




└ img_key

	

No

	

String

	

The image key for a custom prefix icon. Effective when tag is custom_icon.To obtain the icon key: Call the Upload Image API, upload an image for messaging, and retrieve the image key from the response.




padding

	

No

	

String

	

Padding of the title component. Defaults to 12px. The value range is [-99,99]px. Optional values:

Single value, such as "10px", indicates that the four margins of the container are all 10 px.
Double value, such as "4px 0", indicates that the top and bottom margins of the container are 4 px, and the left and right margins are 0 px. Use spaces to separate (units can be omitted when the margin is 0).
Multiple values, such as "4px 0 4px 0", indicate that the top, right, bottom, and left margins of the container are 4px, 12px, 4px, and 12px, respectively. Use spaces to separate.
Demo example

The following sample code of the JSON 2.0 structure can achieve the card effect shown in the figure below:

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "div",
        "text": {
          "tag": "plain_text",
          "content": "示例文本"
        }
      }
    ]
  },
  "header": {
    "title": {
      "tag": "lark_md",
      "content": ":Partying:卡片主标题 "
    },
    "subtitle": {
      "tag": "plain_text",
      "content": "卡片副标题"
    },
    "text_tag_list": [
      {
        "tag": "text_tag",
        "text": {
          "tag": "plain_text",
          "content": "标签 1"
        },
        "color": "blue"
      },
      {
        "tag": "text_tag",
        "text": {
          "tag": "plain_text",
          "content": "标签 2"
        },
        "color": "turquoise"
      },
      {
        "tag": "text_tag",
        "text": {
          "tag": "plain_text",
          "content": "标签 3"
        },
        "color": "orange"
      }
    ],
    "template": "blue",
    "icon": {
      "tag": "standard_icon",
      "token": "larkcommunity_colorful"
    },
    "padding": "12px"
  }
}
Enumerations
Title theme style enumeration

The template field in the title component determines the card's title theme style. You can refer to the table below to understand the enumeration values of template and their corresponding theme styles.

Enumeration	Theme Style Example
blue	
wathet	
turquoise	
green	
yellow	
orange	
red	
carmine	
violet	
purple	
indigo	
grey	
default	
Suffix tag color enumeration

The color style of the suffix tags in the title is defined by the color field in text_tag_list or i18n_text_tag_list. The supported enumeration values and example styles are as shown in the table below.

Enumeration Value	Color Effect
neutral	
blue	
turquoise	
lime	
orange	
violet	
indigo	
wathet	
green	
yellow	
red	
purple	
carmine	
Icon enumeration

For the enumeration of the icon.token field, refer to the Icon Library.

Icon color enumeration

You can set the color of icons through the icon.color. For the enumeration of icon.color, refer to Color Enumeration Values.

Previous:Collapsible panel
Next:Plain text
Need help with a problem?
Submit feedback

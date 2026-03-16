# Multi image laylout

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/multi-image-laylout>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsDisplay componentsMulti image laylout
Multi image laylout
Copy Page
Last updated on 2025-06-27
The contents of this article
Use cases
Precautions
JSON structure
Field descriptions
Example code
Double image mixed layout example
Triple image mixed layout example
Equally divided double column layout example
Equally divided triple column layout example

Lark cards support multi-image combination components. You can call the Upload Image API or upload images in the new version of the Lark card building tool, and enter the obtained image key into the multi-image combination component to enrich the card content.

This document introduces the JSON 2.0 structure of the multi-image layout component. To view the historical JSON 1.0 structure, refer to multi-image layout.

Use cases

In content delivery scenarios, you may need to organize and arrange multiple images within a card. In this case, you can use the multi-image combination component, choose the image combination mode, and quickly build a multi-image style.

Double Image Combination	Triple Image Combination	Quad Grid
		
Six Grid	Nine Grid	
	
Precautions

To ensure clarity in the chat window, it is recommended that the images uploaded in the component adhere to the following specifications:

Image dimensions within the range of 1500 × 3000 px.
Image size no more than 10 MB.
Image height:width ratio not exceeding 16:9.
JSON structure

The complete JSON 2.0 structure of the multi-image layout is as follows:

{
  "schema": "2.0", // Version of the card JSON structure. Default is 1.0. To use the JSON 2.0 structure, you must explicitly declare 2.0.
  "body": {
    "elements": [
      {
        "tag": "img_combination",
        "element_id": "custom_id", // Unique identifier for the operation component. New attribute in JSON 2.0. Used to specify the component when calling related interfaces. Custom defined by the developer.
        "margin": "0px 0px 0px 0px", // Margin of the component. New attribute in JSON 2.0. Default value is "0", supported range is [-99,99]px.
        "combination_mode": "double", // The mode of multi-image layout.
        "combination_transparent": false, // Whether the background is transparent. Default is false, meaning the image has a white background.
        "corner_radius": "12px", // Corner radius of the images in the multi-image layout, in pixels (px).
        "img_list": [
          // Array of image resources, in the order of image arrangement.
          {
            "img_key": "img_v3_0239_8347760e-3173-4072-b1aa-e4e7c835741j",
            "transparent": false // Whether the background is transparent. Default is false, meaning the image has a white background.
          },
          {
            "img_key": "img_v3_0239_d9a9b734-57f8-4247-baf3-ae178b55f96j"
          }
        ]
      }
    ]
  }
}
Field descriptions

The field descriptions for the multi-image combination component are listed in the table below.

Parameter	Required	Type	Default Value	Description


tag

	

Yes

	

String

	

/

	

Label of the multi-image combination component, fixed value: img_combination.




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



combination_mode

	

Yes

	

String

	

empty

	

Modes of the image combination, possible values:

double: Double image combination, can arrange up to two images.
triple: Triple image combination, can arrange up to three images.
bisect: Evenly divided double column combination, each row two equally-sized square images, up to three rows, i.e., six images.
trisect: Evenly divided triple column combination, each row three equally-sized square images, up to three rows, i.e., nine images.

Note:

If the number of uploaded images exceeds the limit of the combination mode, the system will display the images according to the order of upload, prioritizing the images arranged earlier. Images exceeding the limit will not be displayed.If the number of uploaded images is less than the limit of the combination mode, the unarranged parts will remain blank.




combination_transparent

	

No

	

Boolean

	

false

	

Whether the background is transparent. The default is false, meaning the image has a white background.




corner_radius

	

No

	

String

	

/

	

Corner radius for the images in the combination, in pixels (px). Values follow this format:[0,∞]px[0,100]%




img_list

	

Yes

	

Object

	

empty

	

Array of img_key for the image resources, in order with the arrangement of the images.




└ img_key

	

Yes

	

String

	

/

	

Key of the image resource. You can obtain this by calling the Upload Image API or by uploading images in the building tool.

Example code
Double image mixed layout example

Replace the img_key in the following example code with the actual image key to achieve the card effect shown in the example image:

{
  "schema": "2.0",
  "body": {
    "direction": "vertical",
    "padding": "12px 12px 12px 12px",
    "elements": [
      {
        "tag": "img_combination",
        "combination_mode": "double",
        "img_list": [
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          }
        ],
        "combination_transparent": false,
        "margin": "0px 0px 0px 0px"
      }
    ]
  }
}
Triple image mixed layout example

Replace the img_key in the following example code with the actual image key to achieve the card effect shown in the example image:

{
  "schema": "2.0",
  "body": {
    "direction": "vertical",
    "padding": "12px 12px 12px 12px",
    "elements": [
      {
        "tag": "img_combination",
        "combination_mode": "triple",
        "img_list": [
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          }
        ],
        "combination_transparent": false,
        "margin": "0px 0px 0px 0px"
      }
    ]
  }
}
Equally divided double column layout example

Replace the img_key in the following example code with the actual image key to achieve the card effect shown in the example image:

{
  "schema": "2.0",
  "body": {
    "direction": "vertical",
    "padding": "12px 12px 12px 12px",
    "elements": [
      {
        "tag": "img_combination",
        "combination_mode": "bisect",
        "img_list": [
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          }
        ],
        "combination_transparent": false,
        "margin": "0px 0px 0px 0px"
      }
    ]
  }
}
Equally divided triple column layout example

Replace the img_key in the following example code with the actual image key to achieve the card effect shown in the example image:

{
  "schema": "2.0",
  "body": {
    "direction": "vertical",
    "padding": "12px 12px 12px 12px",
    "elements": [
      {
        "tag": "img_combination",
        "combination_mode": "trisect",
        "img_list": [
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          },
          {
            "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg"
          }
        ],
        "combination_transparent": false,
        "margin": "0px 0px 0px 0px"
      }
    ]
  }
}
Previous:Image
Next:Divider
Need help with a problem?
Submit feedback

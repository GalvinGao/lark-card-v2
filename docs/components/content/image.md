# Last updated on 2025-06-27

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/content-components/image>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsDisplay componentsImage
Image
Copy Page
Last updated on 2025-06-27
The contents of this article
Precautions
JSON structure
Field descriptions
Legacy field descriptions
Example code

Lark cards support image components. You can call the Upload Image API or upload an image in the building tool's image component to obtain the image key, enriching the card's content.

This document introduces the JSON 2.0 structure of the image component. To view the historical JSON 1.0 structure, refer to Image.

Precautions

In the JSON 2.0 structure, the size attribute of the image component no longer supports passing stretch_without_padding to achieve full-width effect. You need to set the margin attribute to a negative value to achieve the full-width effect:

{
  "schema": "2.0", // Version of the card JSON structure. Default is 1.0. To use the JSON 2.0 structure, you must explicitly declare 2.0.
  "body": {
    "elements": [
      {
        "tag": "img",
        "img_key": "img_v3_0238_073f1823-df2b-4377-86c6-e293f183622j",
        "scale_type": "crop_center",
        "margin": "4px -12px"
      }
    ]
  }
}

To ensure clarity in the chat window, it is recommended that the images uploaded in the component adhere to the following specifications:

Image dimensions within the range of 1500 × 3000 px.
Image size no more than 10 MB.
Image height:width ratio not exceeding 16:9.
JSON structure

The complete JSON data for the image is shown below:

The complete JSON 2.0 structure of the image component is as follows:
```json
{
  "schema": "2.0", // Version of the card JSON structure. Default is 1.0. To use the JSON 2.0 structure, you must explicitly declare 2.0.
  "body": {
    "elements": [
      {
        "tag": "img",
        "element_id": "custom_id", // Unique identifier for the operation component. New attribute in JSON 2.0. Used to specify the component when calling related interfaces. Custom defined by the developer.
        "margin": "0px 0px 0px 0px", // Margin of the component. New attribute in JSON 2.0. Default value is "0", supported range is [-99,99]px.
        "img_key": "img_v3_0238_073f1823-df2b-4377-86c6-e293f18abcef", // Key of the image. Can be obtained by uploading an image through the upload image interface or in the building tool.
        "alt": {
          // Description displayed when the cursor hovers over the image.
          "tag": "plain_text",
          "content": ""
        },
        "title": {
          // Image title.
          "tag": "plain_text",
          "content": ""
        },
        "corner_radius": "5px", // Corner radius of the image in multi-image layout, in pixels (px).
        "scale_type": "crop_top", // Image cropping mode, triggered when the aspect ratio of the size field and the image are inconsistent.
        "size": "100px 100px", // Image size. Only effective when the scale_type field is crop_center or crop_top.
        "transparent": false, // Whether the background is transparent. Default is false, meaning the image has a white background.
        "preview": false, // Whether to enlarge the image when clicked. Default value is true.
        // Historical attributes
        "mode": "large", // Image size mode.
        "custom_width": "300px", // Custom maximum display width of the image.
        "compact_width": false // Whether to display the image in compact mode.
      }
    ]
  }
}
Field descriptions

The field descriptions for the image component are as follows.

Parameter	Required	Type	Default Value	Description


tag

	

Yes

	

String

	

None

	

The tag of the component, fixed to img for the image component.




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



img_key

	

Yes

	

String

	

/

	

Key of the image resource. You can obtain this by calling the Upload Image API or by uploading images in the building tool.




alt

	

Yes

	

Struct

	

/

	

Description displayed when hovering over the image. Example value:

"alt": {
        "tag": "plain_text",
        "content": "Description displayed when hovering over the image, pass empty if not needed"
      }



title

	

No

	

Struct

	

/

	

Image title. Example value:

"title": {
        "tag": "plain_text",
        "content": "Image Title"
      }



corner_radius

	

No

	

String

	

/

	

Corner radius of the image, in pixels (px). Values follow this format:

[0,∞]px
[0,100]%



scale_type

	

No

	

String

	

crop_center

	

Image cropping mode, activated when the size ratio differs from that of the image. Possible values:

crop_center: Center cropping
crop_top: Top cropping
fit_horizontal: Display fully without cropping



size

	

No

	

String

	

/

	

Dimensions of the image. Effective only when scale_type is crop_center or crop_top. Possible values:

large: Large, suitable for multiple images arranged together.

medium: Medium, suitable for cover images in mixed text and image.

small: Small, suitable for person avatars.

tiny: Tiny, suitable for icons or notes.

Note: The Card JSON 2.0 Structure no longer supports the stretch_without_padding property. You can achieve a full-width effect by setting the margin field to a negative value, such as "margin": "4px -12px". For more details, refer to Unified Component Layout Support.

stretch: Super large, suitable for images with an aspect ratio less than 16:9.

[1,1000]px [1,1000]px: Custom image dimensions, units in pixels, separated by a space.




transparent

	

No

	

Boolean

	

false

	

Whether the background is transparent. Default is false, meaning the image has a white background.




preview

	

No

	

Boolean

	

true

	

Whether to enlarge the image on click.

true: Enlarges the image in a viewer upon clicking.
false: Responds to the card's interaction event upon clicking, without enlarging.

Tip: If you configure a redirection link card_link parameter for the card, you can set this parameter to false, so that user clicks on the image will also trigger the card link redirection.

Legacy field descriptions
Parameter	Required	Type	Default Value	Description


mode

	

No

	

String

	

/

	

Image display mode. Values:

crop_center: Center cropping mode, limiting height for long images and displaying them after center cropping.
fit_horizontal: Tiling mode, stretching the width to display the entire uploaded image across the card.
stretch: Adaptive. The image width stretches across the card width, displaying the entire original image when the image's height:width is less than 16:9. When the image's height:width is more than 16:9, it aligns to the top, crops the image, and displays a Long Image label at the bottom.

Note: Setting this parameter overrides the custom_width parameter. For more information, see Message Card Design Specifications.




custom_width

	

No

	

int

	

/

	

Custom maximum display width of the image, adjustable within the range of 278px to 580px. By default, the image width matches the area of the image component.

Note: This parameter is effective in Lark V4.0 and above.




compact_width

	

No

	

Boolean

	

false

	

Whether to display the image in a compact form. If set to true, the maximum display width is 278px.

Example code

The following example code in JSON 2.0 structure can achieve the card effect shown in the figure:

{
  "schema": "2.0",
  "body": {
    "direction": "vertical",
    "padding": "12px 12px 12px 12px",
    "elements": [
      {
        "tag": "img",
        "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg",
        "preview": true,
        "transparent": false,
        "scale_type": "crop_center",
        "size": "stretch",
        "alt": {
          "tag": "plain_text",
          "content": "示例图片"
        },
        "corner_radius": "5%",
        "margin": "0px 0px 0px 0px",
        "element_id": "demoimg01"
      }
    ]
  }
}
Previous:Rich text (Markdown)
Next:Multi image laylout
Need help with a problem?
Submit feedback

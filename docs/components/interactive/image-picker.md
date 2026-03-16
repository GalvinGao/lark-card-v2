# Image picker

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/image-picker>

Developer GuidesMessage cardsBuild card with JSONCard JSON 2.0 version componentsInteractive componentsImage picker
Image picker
Copy Page
Last updated on 2025-06-27
The contents of this article
Considerations
Nesting rules
Component properties
JSON structure
Field description
Callback structure
Sample code
Immediate submission of single selection
Asynchronously submitting multiple selection options within a Form Container

The multi-image selection component is an interactive component used to provide options for images, supporting single or multiple image selection. This component is suitable for scenarios where images are the primary options, such as displaying product images, template images, or AI-generated images for user selection.

This document introduces the JSON 2.0 structure of the multi-image picker component. To view the historical JSON 1.0 structure, refer to multi-image picker.

Considerations

The multi-image selection component is only supported for use by writing card JSON code and is not currently supported for construction using card building tools.

Nesting rules

The multi-image selection component can be nested within the root node of a card, multi-column layout, form container component, and interactive containers (not supported by the building tool yet). In different nesting relationships, the multi-image selection component supports different interaction modes:

When the multi-image selection component is not nested within a form container, it only supports single image selection. After the end user clicks on an image option, it is immediately submitted, triggering a callback interaction. Multiple selections and asynchronous submissions are not supported.
When the multi-image selection component is nested within a form container, it supports both single and multiple selection interactions, as well as asynchronous submissions. That is, after the end user selects an image, they need to click the submit button of the form container to callback the locally cached form content to the developer's server, achieving asynchronous submission.
Component properties
JSON structure

The JSON 2.0 data structure of the multi-image selection component is as follows:

{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "select_img", // Component tag.
        "element_id": "custom_id", // Unique identifier for the operation component. Used to specify the component in related interface calls. Needs to be customized by the developer.
        "margin": "0px 0px 0px 0px", // Component margin, default value is "0", supported range is [-99,99]px.
        "multi_select": false, // Whether multiple selection is allowed.
        "layout": "bisect", // Layout mode of the options.
        "name": "choice_123", // Custom name of the multi-image picker component as a unique identifier. This field is effective and required when the component is nested within a form container, used to identify which component the user-submitted data belongs to.
        "required": false, // Whether the options in the multi-image picker are required. This property is available when the component is nested within a form container. Otherwise, it will report an error or not take effect.
        "can_preview": false, // Whether to pop up and enlarge the image when clicking the image option. This property is effective when the multi-image picker component is nested within a form container.
        "aspect_ratio": "16:9", // Aspect ratio of the images in the options.
        "disabled": false, // Whether to disable the entire selection component.
        "disabled_tips": { // Tips displayed when the user hovers the cursor over the entire component after it is disabled.
          "tag": "plain_text",
          "content": "User disabled tips"
        },
        "value": { // Custom callback parameters, supporting string or object structure composed of "key":"value".
          "key": "value"
        },
        // Array of options. Configure the properties of each image option in the multi-image picker component here.
        "options": [
          {
            "img_key": "xxxxxxxxxxxxxx", // Key of the image resource.
            "value": "picture1", // Custom callback parameter for each image option.
            "disabled": false, // Whether to disable the current image option.
            "disabled_tips": { // Tips displayed when the user hovers the cursor over or clicks the option after it is disabled.
              "tag": "plain_text",
              "content": "User disabled tips"
            },
            "hover_tips": { // Text reminder when the user hovers the cursor over the option on the PC end.
              "tag": "plain_text",
              "content": "First image"
            }
          }
        ],
        "behaviors": [
          { // Configure callback interaction for the component.
            "type": "callback",
            "value": {
              // Callback interaction data. Supports string or object data types. The open platform SDK only supports object type callback interaction data.
              "key": "value"
            }
          }
        ],
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
Field description

The field descriptions of the multi-image selection component are shown in the table below.

Field	Required	Type	Default Value	Description


tag

	

Yes

	

String

	

/

	

Component tag. For multi-image selection, fixed value is select_img.




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



multi_select

	

No

	

Boolean

	

false

	

Whether multiple images can be selected. Options:

true: Multiple selection, only supports asynchronous submission. The multi-image selection component must be embedded in a form container, otherwise the card JSON will report an error.
false: Single selection.
When the component is inside a form container, image options are displayed with radio buttons in an asynchronous submission style.
When the component is not inside a form container, image options are displayed without radio buttons in a synchronous submission style.



layout

	

No

	

String

	

bisect

	

Layout of image options. Options:

stretch: Image width of each option stretches to fill the width of the parent container, height scales proportionally according to the image size.
bisect: Divides into two equal parts, each option's image width occupies half of the parent container, height scales proportionally according to the image size.
trisect: Divides into three equal parts, each option's image width occupies one-third of the parent container, height scales proportionally according to the image size.



name

	

No

	

String

	

Empty

	

Custom name for the multi-image selection component as a unique identifier. Used to identify which component the submitted data belongs to.

Note: When the multi-image selection component is nested in a form container, this field is effective, required, and must be unique within the card scope.




required

	

No

	

Boolean

	

false

	

Whether options in multi-image selection are required. This attribute is only available when the component is embedded in a form container. Otherwise, it will result in an error or have no effect. Options:

true: Options are required. When the user clicks "Submit" in the form container without selecting any option, the frontend prompts "Some required fields are not filled in", and no request is sent to the developer's server.
false: Options are optional. When the user clicks "Submit" in the form container without selecting any option, the data in the form container is still submitted.



can_preview

	

No

	

Boolean

	

true

	

Whether to zoom in on the image after clicking. This property takes effect when the multi-image selection component is nested in a form container.

true: After clicking the image, a viewer pops up to zoom in on the clicked image.
false: After clicking the image, the interaction event of the card itself is responded to, and no viewer pops up.



aspect_ratio

	

No

	

String

	

16:9

	

Aspect ratio of images in options. Images stretch to fill the rendering container by the shortest side, and adaptively crop centeredly. Options:

1:1
16:9
4:3



disabled

	

No

	

Boolean

	

false

	

Whether to disable the entire selection component. Options:

true: Disable the entire selection component.
false: Selection component remains enabled.



disabled_tips

	

No

	

Object

	

Empty

	

Text displayed when the user hovers over or clicks on the entire component after it's disabled.




└ tag

	

No

	

String

	

plain_text

	

Tag for the disabled tips text. Fixed value is plain_text.




└ content

	

No

	

String

	

Empty

	

Content of the disabled tips text.




value

	

No

	

String or Object

	

Empty

	

Customizable parameters for passing back in interaction events, supporting either a string or an object structure composed of "key":"value".




options

	

Yes

	

Option object

	

/

	

Array of options used to configure the properties of each image option in the multi-image selection component.




L img_key

	

Yes

	

String

	

/

	

Key of the image resource. You can call the upload image interface or upload images in the building tool to get the image key.




L value

	

No

	

String

	

Empty

	

Custom passing parameter for each image option. The passing parameters specified in the interaction will be passed to the developer's server.




L disabled

	

No

	

Boolean

	

false

	

Whether to disable a specific image option. Options:

true: Disable the option.
false: Option remains enabled.



L disabled_tips

	

No

	

Object

	

Empty

	

Text displayed when the user hovers over or clicks on the option after it's disabled.




LL tag

	

No

	

String

	

plain_text

	

Tag for the disabled tips text. Fixed value is plain_text.




LL content

	

No

	

String

	

Empty

	

Content of the disabled tips text.




L hover_tips

	

No

	

Object

	

Empty

	

Text reminder when the user hovers over the multi-image selection on the PC end. Default is empty.




LL tag

	

No

	

String

	

plain_text

	

Tag for the hover tips text. Fixed value is plain_text.




LL content

	

No

	

String

	

Empty

	

Content of the hover tips text.

Callback structure

After successfully configuring interactions for the component, when users interact with the component, the request address configured in your developer backend will receive callback data.

If you added a new version card callback interaction (card.action.trigger), you can refer to Card Callback Interaction to understand the callback structure.
If you added an old version card callback interaction (card.action.trigger_v1), you can refer to Message Card Callback Interaction (Old) to understand the callback structure.
Sample code
Immediate submission of single selection

In the following example, the multi-image selection component is not nested within a form container and only supports single image selection. After the end user clicks on an image option, the data will be immediately submitted, triggering a callback interaction. You can replace the img_key in the sample code with the actual image key to achieve the card effect as shown in the image below:

{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "select_img",
                "name": "select_img-1",
                "layout": "bisect",
                "aspect_ratio": "16:9",
                "disabled": false,
                "disabled_tips": {
                    "tag": "plain_text",
                    "content": "用户禁用提示文案"
                },
                "options": [
                    {
                        "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg",
                        "value": "picture1",
                        "disabled": false,
                        "disabled_tips": {
                            "tag": "plain_text",
                            "content": "用户禁用提示文案1"
                        },
                        "hover_tips": {
                            "tag": "plain_text",
                            "content": "第一张图"
                        }
                    },
                    {
                        "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg",
                        "value": "picture2",
                        "disabled": false,
                        "disabled_tips": {
                            "tag": "plain_text",
                            "content": "用户禁用提示文案2"
                        },
                        "hover_tips": {
                            "tag": "plain_text",
                            "content": "第二张图"
                        }
                    },
                    {
                        "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg",
                        "value": "picture3",
                        "disabled": false,
                        "disabled_tips": {
                            "tag": "plain_text",
                            "content": "用户禁用提示文案3"
                        },
                        "hover_tips": {
                            "tag": "plain_text",
                            "content": "第三张图"
                        }
                    },
                    {
                        "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg",
                        "value": "picture4",
                        "disabled": false,
                        "disabled_tips": {
                            "tag": "plain_text",
                            "content": "用户禁用提示文案4"
                        },
                        "hover_tips": {
                            "tag": "plain_text",
                            "content": "第四张图"
                        }
                    },
                    {
                        "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg",
                        "value": "picture5",
                        "disabled": false,
                        "disabled_tips": {
                            "tag": "plain_text",
                            "content": "用户禁用提示文案5"
                        },
                        "hover_tips": {
                            "tag": "plain_text",
                            "content": "第五张图"
                        }
                    },
                    {
                        "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg",
                        "value": "picture6",
                        "disabled": false,
                        "disabled_tips": {
                            "tag": "plain_text",
                            "content": "用户禁用提示文案6"
                        },
                        "hover_tips": {
                            "tag": "plain_text",
                            "content": "第六张图"
                        }
                    }
                ]
            }
        ]
    }
}
Asynchronously submitting multiple selection options within a Form Container

In the following example, the multi-image selection component is nested within a form container and uses asynchronous submission for single selection interaction. After the end user selects an image, the data will be cached locally; when the user clicks the submit button of the form container, all locally cached form content will be callbacked to the developer's server at once, achieving asynchronous submission. You can replace the img_key in the sample code with the actual image key to achieve the card effect as shown in the image below:

{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "form",
                "name": "form1",
                "elements": [
                    {
                        "tag": "select_img",
                        "multi_select": false,
                        "name": "select_img-1",
                        "layout": "bisect",
                        "can_preview": false,
                        "aspect_ratio": "16:9",
                        "disabled": false,
                        "disabled_tips": {
                            "tag": "plain_text",
                            "content": "用户禁用提示文案"
                        },
                        "options": [
                            {
                                "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg",
                                "value": "picture1",
                                "disabled": false,
                                "disabled_tips": {
                                    "tag": "plain_text",
                                    "content": "用户禁用提示文案1"
                                },
                                "hover_tips": {
                                    "tag": "plain_text",
                                    "content": "第一张图"
                                }
                            },
                            {
                                "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg",
                                "value": "picture2",
                                "disabled": false,
                                "disabled_tips": {
                                    "tag": "plain_text",
                                    "content": "用户禁用提示文案2"
                                },
                                "hover_tips": {
                                    "tag": "plain_text",
                                    "content": "第二张图"
                                }
                            },
                            {
                                "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg",
                                "value": "picture3",
                                "disabled": false,
                                "disabled_tips": {
                                    "tag": "plain_text",
                                    "content": "用户禁用提示文案3"
                                },
                                "hover_tips": {
                                    "tag": "plain_text",
                                    "content": "第三张图"
                                }
                            },
                            {
                                "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg",
                                "value": "picture4",
                                "disabled": false,
                                "disabled_tips": {
                                    "tag": "plain_text",
                                    "content": "用户禁用提示文案4"
                                },
                                "hover_tips": {
                                    "tag": "plain_text",
                                    "content": "第四张图"
                                }
                            },
                            {
                                "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg",
                                "value": "picture5",
                                "disabled": false,
                                "disabled_tips": {
                                    "tag": "plain_text",
                                    "content": "用户禁用提示文案5"
                                },
                                "hover_tips": {
                                    "tag": "plain_text",
                                    "content": "第五张图"
                                }
                            },
                            {
                                "img_key": "img_v2_9dd98485-2900-4d65-ada9-e31d1408dcfg",
                                "value": "picture6",
                                "disabled": false,
                                "disabled_tips": {
                                    "tag": "plain_text",
                                    "content": "用户禁用提示文案6"
                                },
                                "hover_tips": {
                                    "tag": "plain_text",
                                    "content": "第六张图"
                                }
                            }
                        ]
                    },
                    {
                        "tag": "button",
                        "text": {
                            "tag": "plain_text",
                            "content": "提交"
                        },
                        "type": "primary",
                        "name": "button-submit",
                        "form_action_type": "submit",
                        "behaviors": [
                            {
                                "type": "callback",
                                "value": "form_callback"
                            }
                        ]
                    },
                    {
                        "tag": "button",
                        "text": {
                            "tag": "plain_text",
                            "content": "取消"
                        },
                        "name": "button-cancel",
                        "form_action_type": "reset"
                    }
                ]
            }
        ]
    }
}
Previous:Date time picker
Next:Checker
Need help with a problem?
Submit feedback

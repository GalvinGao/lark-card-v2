# Image Picker

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/image-picker>

The image picker component provides image-based options supporting single or multiple selection. It is suitable for scenarios where images are the primary options, such as displaying product images, template images, or AI-generated images for user selection.

## Notes

- The image picker can only be created by writing card JSON code; it is not supported in the Card Builder Tool.
- When not nested in a form container, only single selection with immediate submission is supported.
- When nested in a form container, both single and multiple selection with asynchronous submission are supported.
- When `multi_select` is `true`, the component must be embedded in a form container, otherwise the card JSON will error.

## Nesting Rules

The image picker can be nested within the card root node, multi-column layout, form container, and interactive containers (not yet supported in the Card Builder Tool).

- **Outside a form container**: Only single selection. Clicking an image option immediately submits and triggers a callback.
- **Inside a form container**: Supports single and multiple selection. After selecting, the user clicks the form's submit button to callback locally cached form data.

## JSON Structure

```json
{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "select_img",
                "element_id": "custom_id",
                "margin": "0px 0px 0px 0px",
                "multi_select": false,
                "layout": "bisect",
                "name": "choice_123",
                "required": false,
                "can_preview": false,
                "aspect_ratio": "16:9",
                "disabled": false,
                "disabled_tips": {
                    "tag": "plain_text",
                    "content": "User disabled tips"
                },
                "value": {
                    "key": "value"
                },
                "options": [
                    {
                        "img_key": "xxxxxxxxxxxxxx",
                        "value": "picture1",
                        "disabled": false,
                        "disabled_tips": {
                            "tag": "plain_text",
                            "content": "User disabled tips"
                        },
                        "hover_tips": {
                            "tag": "plain_text",
                            "content": "First image"
                        }
                    }
                ],
                "behaviors": [
                    {
                        "type": "callback",
                        "value": {
                            "key": "value"
                        }
                    }
                ],
                "confirm": {
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
```

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | Component tag. Fixed value `select_img`. |
| `element_id` | No | String | — | Unique identifier for the component (JSON 2.0). Must be globally unique within the card. Only letters, numbers, and underscores allowed; must start with a letter; max 20 characters. |
| `margin` | No | String | `0` | Component margin (JSON 2.0). Value range is `[-99,99]px`. Accepts a single value (e.g. `"10px"`), double value (e.g. `"4px 0"`), or four values (e.g. `"4px 0 4px 0"`) for top, right, bottom, left. |
| `multi_select` | No | Boolean | `false` | Whether multiple images can be selected. `true`: multi-select (must be inside a form container). `false`: single select (shows radio buttons in form container, no buttons outside). |
| `layout` | No | String | `bisect` | Layout of image options. `stretch`: each image fills the parent container width. `bisect`: two images per row. `trisect`: three images per row. Height scales proportionally. |
| `name` | No | String | — | Custom name as a unique identifier. Required and must be globally unique when nested in a form container. |
| `required` | No | Boolean | `false` | Whether selection is required. Only takes effect inside a form container. When `true`, submitting without a selection shows a "required fields not filled in" prompt and prevents submission. |
| `can_preview` | No | Boolean | `true` | Whether clicking an image opens a zoomed preview. Only takes effect inside a form container. `true`: opens viewer. `false`: responds to the card's interaction event instead. |
| `aspect_ratio` | No | String | `16:9` | Aspect ratio of option images. Images stretch to fill by shortest side and crop from center. Options: `1:1`, `16:9`, `4:3`. |
| `disabled` | No | Boolean | `false` | Whether to disable the entire selection component. |
| `disabled_tips` | No | Object | — | Tooltip shown when the user hovers over the disabled component. |
| `disabled_tips.tag` | No | String | `plain_text` | Fixed value `plain_text`. |
| `disabled_tips.content` | No | String | — | Tooltip text content. |
| `value` | No | String or Object | — | Custom callback parameters, supporting a string or an object of `"key":"value"` pairs. |
| `options` | Yes | Array of objects | — | Array of image options. |
| `options[].img_key` | Yes | String | — | Image resource key. Obtain by calling the Upload Image API or uploading in the Card Builder Tool. |
| `options[].value` | No | String | — | Custom callback parameter for each image option. Returned to the server on interaction. |
| `options[].disabled` | No | Boolean | `false` | Whether to disable this specific image option. |
| `options[].disabled_tips` | No | Object | — | Tooltip shown when hovering over or clicking the disabled option. |
| `options[].disabled_tips.tag` | No | String | `plain_text` | Fixed value `plain_text`. |
| `options[].disabled_tips.content` | No | String | — | Tooltip text content. |
| `options[].hover_tips` | No | Object | — | Tooltip shown when hovering over the option on PC. |
| `options[].hover_tips.tag` | No | String | `plain_text` | Fixed value `plain_text`. |
| `options[].hover_tips.content` | No | String | — | Tooltip text content. |
| `confirm` | No | Object | — | Secondary confirmation popup configuration. |
| `confirm.title` | Yes | Object | — | Popup title. |
| `confirm.title.tag` | Yes | String | `plain_text` | Fixed value `plain_text`. |
| `confirm.title.content` | Yes | String | — | Title text content. |
| `confirm.text` | Yes | Object | — | Popup body text. |
| `confirm.text.tag` | Yes | String | `plain_text` | Fixed value `plain_text`. |
| `confirm.text.content` | Yes | String | — | Body text content. |

## Callback Structure

After configuring interactions for the component, user interactions will send callback data to your configured request address.

- If you use the new card callback interaction (`card.action.trigger`), refer to the Card Callback Interaction documentation for the callback structure.
- If you use the old card callback interaction (`card.action.trigger_v1`), refer to the Message Card Callback Interaction (Old) documentation.

## Example

### Immediate single selection (outside form container)

Replace `img_key` values with actual image keys.

```json
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
```

### Asynchronous multi-selection (inside form container)

Replace `img_key` values with actual image keys.

```json
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
```

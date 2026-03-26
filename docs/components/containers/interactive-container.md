# Interactive Container

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/containers/interactive-container>

You can embed components within interactive containers based on business needs, flexibly combine multiple interactive containers, and uniformly define the styles and interaction capabilities of multiple interactive containers, achieving various combination effects and rich card interactions.

This document introduces the JSON 2.0 structure of the interactive container. To view the historical JSON 1.0 structure, refer to interactive container.

## Notes

Container components support a maximum of five layers of nested components. It is recommended to avoid nesting multiple layers in interactive containers. Multi-layer nesting compresses the display space and affects the presentation of the card.

## Nesting Rules

In the Card JSON 2.0 Structure, interactive containers can embed all components except for form containers (`form`) and table components (`table`).

## JSON Structure

The complete JSON 2.0 structure of the interactive container is as follows:

```json
{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "interactive_container",
                "width": "fill",
                "height": "auto",
                "element_id": "custom_id",
                "direction": "vertical",
                "margin": "4px",
                "horizontal_spacing": "large",
                "horizontal_align": "left",
                "vertical_align": "center",
                "vertical_spacing": "4px",
                "background_style": "default",
                "has_border": false,
                "border_color": "grey",
                "corner_radius": "40px",
                "padding": "10px 20px 10px 20px",
                "behaviors": [
                    {
                        "type": "open_url",
                        "default_url": "https://www.baidu.com",
                        "android_url": "https://developer.android.com/",
                        "ios_url": "lark://msgcard/unsupported_action",
                        "pc_url": "https://www.windows.com"
                    },
                    {
                        "type": "callback",
                        "value": {
                            "key": "value"
                        }
                    }
                ],
                "disabled": false,
                "disabled_tips": {
                    "tag": "plain_text",
                    "content": "demo"
                },
                "confirm": {},
                "hover_tips": {
                    "tag": "plain_text",
                    "content": "demo"
                },
                "elements": []
            }
        ]
    }
}
```

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | Tag of the interactive container. Fixed value is `interactive_container`. |
| `width` | No | String | `fill` | Width of the interactive container. `fill`: maximum supported width of the card. `auto`: adaptive width. `[16,999]px`: custom width, e.g., `"20px"`. Minimum width is 16px. |
| `height` | No | String | `auto` | Height of the interactive container. `auto`: adaptive height. `[10,999]px`: custom height, e.g., `"20px"`. |
| `element_id` | No | String | — | Unique identifier of the component. Newly added in JSON 2.0. Used to specify the component when calling related interfaces. Must be globally unique within the same card. Only letters, numbers, and underscores are allowed, must start with a letter, and must not exceed 20 characters. |
| `margin` | No | String | `0px` | Outer margin of the container. Value range is [-99,99]px. Single value (e.g., `"10px"`) applies to all four sides. Double value (e.g., `"4px 0"`) sets top/bottom and left/right. Multiple values (e.g., `"4px 0 4px 0"`) set top, right, bottom, and left respectively. |
| `direction` | No | String | `vertical` | Arrangement direction of components within the container. `vertical`: vertical arrangement. `horizontal`: horizontal arrangement. |
| `horizontal_spacing` | No | String | `8px` | Horizontal spacing between components within the container. `small` (4px), `medium` (8px), `large` (12px), `extra_large` (16px), or a specific value such as `"20px"`. Value range is [0,99]px. |
| `horizontal_align` | No | String | `left` | Horizontal alignment of components within the container. `left`, `center`, or `right`. |
| `vertical_spacing` | No | String | `12px` | Vertical spacing between components within the container. `small` (4px), `medium` (8px), `large` (12px), `extra_large` (16px), or a specific value such as `"20px"`. Value range is [0,99]px. |
| `vertical_align` | No | String | `top` | Vertical alignment of components within the container. `top`, `center`, or `bottom`. |
| `background_style` | No | String | `default` | Background style of the interactive container. `default`: white background (black in dark mode). `laser`: laser gradient color style. Also supports card color enumeration values and custom colors in RGBA syntax. |
| `has_border` | No | Boolean | `false` | Whether to display a border, with a fixed thickness of 1px. |
| `border_color` | No | String | `grey` | Border color, effective only when `has_border` is `true`. Supports card color enumeration values and custom colors in RGBA syntax. |
| `corner_radius` | No | String | `0px` | Corner radius of the interactive container, in pixels (px) or percentage (%). Format: `[0,inf]px` (e.g., `"10px"`) or `[0,100]%` (e.g., `"30%"`). |
| `padding` | No | String | `4px 12px` | Padding of the container. Value range is [-99,99]px. Single value (e.g., `"10px"`) applies to all four sides. Double value (e.g., `"4px 0"`) sets top/bottom and left/right. Multiple values (e.g., `"4px 0 4px 0"`) set top, right, bottom, and left respectively. |
| `behaviors` | Yes | Array | — | Interaction configuration when clicking the interactive container. If there are interactive components within the container, the interaction defined by those components takes precedence. Supports `callback` and `open_url` interactions. See configuring card interactions. |
| `hover_tips` | No | Object | — | Tooltip text when the user hovers the cursor over the interactive container on PC. |
| `hover_tips.tag` | Yes | String | `plain_text` | The tag of the text. Fixed value is `plain_text`. |
| `hover_tips.content` | Yes | String | — | The content of the tooltip text. |
| `disabled` | No | Boolean | `false` | Whether to disable the interactive container. `true`: disable the entire container. `false`: keep the container components available. |
| `disabled_tips` | No | Object | — | Tooltip text when the user triggers interaction after the container is disabled. |
| `disabled_tips.tag` | Yes | String | `plain_text` | The tag of the tooltip text. Fixed value is `plain_text`. |
| `disabled_tips.content` | Yes | String | — | The content of the tooltip text. |
| `confirm` | No | Struct | — | Secondary confirmation dialog configuration. A confirmation prompt appears when the user submits; input is only submitted after clicking confirm. Provides confirm and cancel buttons by default -- you only need to configure the title and content. Note: The `confirm` field only triggers when the user clicks a button with the submit attribute. |
| `confirm.title` | Yes | Struct | — | Secondary confirmation dialog title. |
| `confirm.title.tag` | Yes | String | `plain_text` | The tag of the title text. Fixed value is `plain_text`. |
| `confirm.title.content` | Yes | String | — | The content of the title. |
| `confirm.text` | Yes | Struct | — | The text content of the secondary confirmation dialog. |
| `confirm.text.tag` | Yes | String | `plain_text` | The tag of the dialog text. Fixed value is `plain_text`. |
| `confirm.text.content` | Yes | String | — | The specific content of the dialog text. |
| `elements` | Yes | Array\<element\> | `[]` | Components embedded in the interactive container. Supports all components except form containers (`form`) and table components (`table`). |

## Callback Structure

Once components are successfully configured for interaction, when users interact with the components, callback data will be sent to the developer backend URL you configured.

- If you have added the new version card action callback (`card.action.trigger`), refer to Card Callback Communication for the callback structure.
- If you have added the older version card action callback (`card.action.trigger_v1`), refer to Old Message Card Callback Interaction for the callback structure.

## Example

The following JSON example achieves an interactive container card with suggestion prompts, conversation history, and plugin badges:

```json
{
    "schema": "2.0",
    "header": {
        "title": {
            "content": "交互容器示例",
            "tag": "plain_text"
        },
        "icon": {
            "tag": "standard_icon",
            "token": "chat_outlined",
            "color": "orange"
        }
    },
    "body": {
        "elements": [
            {
                "tag": "markdown",
                "content": "在「内容创作」话题下，我可以帮助你进行产品方案、营销文案、工作报告等内容的创作。"
            },
            {
                "tag": "column_set",
                "flex_mode": "none",
                "background_style": "default",
                "columns": [
                    {
                        "tag": "column",
                        "width": "weighted",
                        "weight": 1,
                        "vertical_align": "top",
                        "vertical_spacing": "8px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<font color='grey'>你可以对我说：</font>",
                                "text_size": "notation"
                            },
                            {
                                "tag": "interactive_container",
                                "width": "fill",
                                "height": "auto",
                                "horizontal_align": "left",
                                "background_style": "default",
                                "has_border": true,
                                "border_color": "grey",
                                "corner_radius": "8px",
                                "padding": "4px 12px 4px 12px",
                                "behaviors": [
                                    {
                                        "type": "callback",
                                        "value": {
                                            "key": "value"
                                        }
                                    }
                                ],
                                "elements": [
                                    {
                                        "tag": "markdown",
                                        "content": "帮我生成一篇产品方案的框架",
                                        "icon": {
                                            "tag": "standard_icon",
                                            "token": "frame-selection_outlined",
                                            "color": "orange"
                                        }
                                    }
                                ]
                            },
                            {
                                "tag": "interactive_container",
                                "width": "fill",
                                "height": "auto",
                                "horizontal_align": "left",
                                "background_style": "default",
                                "has_border": true,
                                "border_color": "grey",
                                "corner_radius": "8px",
                                "padding": "4px 12px 4px 12px",
                                "behaviors": [
                                    {
                                        "type": "callback",
                                        "value": {
                                            "key": "value"
                                        }
                                    }
                                ],
                                "disabled": false,
                                "elements": [
                                    {
                                        "tag": "markdown",
                                        "content": "帮我生成一篇产品文案",
                                        "icon": {
                                            "tag": "standard_icon",
                                            "token": "file-link-docx_outlined",
                                            "color": "orange"
                                        }
                                    }
                                ]
                            },
                            {
                                "tag": "interactive_container",
                                "width": "fill",
                                "height": "auto",
                                "horizontal_align": "left",
                                "background_style": "default",
                                "has_border": true,
                                "border_color": "grey",
                                "corner_radius": "8px",
                                "padding": "4px 12px 4px 12px",
                                "behaviors": [
                                    {
                                        "type": "callback",
                                        "value": {
                                            "key": "value"
                                        }
                                    }
                                ],
                                "disabled": false,
                                "elements": [
                                    {
                                        "tag": "markdown",
                                        "content": "帮我写一篇周报",
                                        "icon": {
                                            "tag": "standard_icon",
                                            "token": "pa-calibration-report_outlined",
                                            "color": "orange"
                                        }
                                    }
                                ]
                            }
                        ]
                    }
                ]
            },
            {
                "tag": "column_set",
                "flex_mode": "none",
                "background_style": "default",
                "columns": [
                    {
                        "tag": "column",
                        "width": "weighted",
                        "weight": 1,
                        "vertical_align": "top",
                        "vertical_spacing": "8px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<font color='grey'>或者继续之前的话题</font>",
                                "text_size": "notation"
                            },
                            {
                                "tag": "interactive_container",
                                "width": "fill",
                                "height": "auto",
                                "horizontal_align": "left",
                                "background_style": "default",
                                "has_border": true,
                                "border_color": "grey",
                                "corner_radius": "8px",
                                "padding": "4px 12px 4px 12px",
                                "behaviors": [
                                    {
                                        "type": "callback",
                                        "value": {
                                            "key": "value"
                                        }
                                    }
                                ],
                                "disabled": false,
                                "elements": [
                                    {
                                        "tag": "column_set",
                                        "flex_mode": "none",
                                        "background_style": "default",
                                        "columns": [
                                            {
                                                "tag": "column",
                                                "width": "weighted",
                                                "weight": 1,
                                                "vertical_align": "center",
                                                "elements": [
                                                    {
                                                        "tag": "markdown",
                                                        "content": "内容创作:创作暑假营销活动文案",
                                                        "icon": {
                                                            "tag": "standard_icon",
                                                            "token": "chat-history_outlined"
                                                        }
                                                    }
                                                ]
                                            },
                                            {
                                                "tag": "column",
                                                "width": "auto",
                                                "weight": 1,
                                                "vertical_align": "center",
                                                "elements": [
                                                    {
                                                        "tag": "markdown",
                                                        "content": "<font color='grey'>昨天</font>",
                                                        "text_size": "notation"
                                                    }
                                                ]
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "tag": "interactive_container",
                                "width": "fill",
                                "height": "auto",
                                "horizontal_align": "left",
                                "background_style": "default",
                                "has_border": true,
                                "border_color": "grey",
                                "corner_radius": "8px",
                                "padding": "4px 12px 4px 12px",
                                "behaviors": [
                                    {
                                        "type": "callback",
                                        "value": {
                                            "key": "value"
                                        }
                                    }
                                ],
                                "disabled": false,
                                "elements": [
                                    {
                                        "tag": "column_set",
                                        "flex_mode": "none",
                                        "background_style": "default",
                                        "columns": [
                                            {
                                                "tag": "column",
                                                "width": "weighted",
                                                "weight": 1,
                                                "vertical_align": "center",
                                                "elements": [
                                                    {
                                                        "tag": "markdown",
                                                        "content": "内容创作:生成了季度工作报告",
                                                        "icon": {
                                                            "tag": "standard_icon",
                                                            "token": "chat-history_outlined"
                                                        }
                                                    }
                                                ]
                                            },
                                            {
                                                "tag": "column",
                                                "width": "auto",
                                                "weight": 1,
                                                "vertical_align": "center",
                                                "elements": [
                                                    {
                                                        "tag": "markdown",
                                                        "content": "<font color='grey'>上周</font>",
                                                        "text_size": "notation"
                                                    }
                                                ]
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "tag": "interactive_container",
                                "width": "fill",
                                "height": "auto",
                                "horizontal_align": "left",
                                "background_style": "default",
                                "has_border": true,
                                "border_color": "grey",
                                "corner_radius": "8px",
                                "padding": "4px 12px 4px 12px",
                                "behaviors": [
                                    {
                                        "type": "callback",
                                        "value": {
                                            "key": "value"
                                        }
                                    }
                                ],
                                "disabled": false,
                                "elements": [
                                    {
                                        "tag": "markdown",
                                        "content": "更多历史话题",
                                        "icon": {
                                            "tag": "standard_icon",
                                            "token": "chat-history_outlined"
                                        }
                                    }
                                ]
                            }
                        ]
                    }
                ]
            },
            {
                "tag": "column_set",
                "flex_mode": "none",
                "background_style": "default",
                "columns": [
                    {
                        "tag": "column",
                        "width": "weighted",
                        "weight": 1,
                        "vertical_align": "top",
                        "vertical_spacing": "8px",
                        "elements": [
                            {
                                "tag": "div",
                                "text": {
                                    "tag": "plain_text",
                                    "text_size": "notation",
                                    "content": "本话题中已选择以下插件",
                                    "text_color": "grey"
                                }
                            },
                            {
                                "tag": "interactive_container",
                                "width": "auto",
                                "height": "auto",
                                "horizontal_align": "left",
                                "background_style": "grey",
                                "has_border": false,
                                "border_color": "grey",
                                "corner_radius": "40px",
                                "padding": "2px 8px 2px 4px",
                                "behaviors": [
                                    {
                                        "type": "callback",
                                        "value": {
                                            "key": "value"
                                        }
                                    }
                                ],
                                "disabled": false,
                                "elements": [
                                    {
                                        "tag": "column_set",
                                        "flex_mode": "none",
                                        "background_style": "default",
                                        "horizontal_spacing": "4px",
                                        "columns": [
                                            {
                                                "tag": "column",
                                                "width": "auto",
                                                "weight": 1,
                                                "vertical_align": "center",
                                                "vertical_spacing": "8px",
                                                "elements": [
                                                    {
                                                        "tag": "img",
                                                        "img_key": "img_v2_58e37110-6878-44ee-bce4-7a571c1bb70g",
                                                        "transparent": true,
                                                        "scale_type": "crop_center",
                                                        "size": "18px 18px",
                                                        "preview": false
                                                    }
                                                ]
                                            },
                                            {
                                                "tag": "column",
                                                "width": "weighted",
                                                "weight": 1,
                                                "vertical_align": "center",
                                                "vertical_spacing": "8px",
                                                "elements": [
                                                    {
                                                        "tag": "markdown",
                                                        "content": "妙记插件"
                                                    }
                                                ]
                                            }
                                        ]
                                    }
                                ]
                            }
                        ]
                    }
                ]
            }
        ]
    }
}
```

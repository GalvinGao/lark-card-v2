# Column Set

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/containers/column-set>

Column components provide layout capabilities within a card, offering properties such as alignment, container width, and interaction methods. You can use column components to horizontally arrange multiple column containers, freely combining images and text to create cards rich in graphics — such as data tables, product or article lists, travel information, etc.

This document describes the JSON 2.0 structure of the column component. For the historical JSON 1.0 structure, refer to [Column set](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/containers/column-set).

## Notes

- Column components support a maximum of five nested layers. Avoid nesting multiple layers within columns, as multi-layer nesting compresses the content display space and affects card presentation.

## Application Scenarios

Column usage scenarios are extensive. Proper use of columns can make the information layout more reasonable and the hierarchy more distinct. Common scenarios include:

- **Data report push:** Using columns can quickly build neat, screen-adaptive multi-column data tables, solving cumbersome layout processes and style issues that cannot adapt to various screens.
- **Mixed text and images:** The flexible horizontal and vertical layout capabilities of columns allow you to quickly build ideal text-and-image cards, effectively reducing manual adjustment time.
- **Form collection:** Embedding column components in form containers and placing related fields in the same column can improve the logicality and readability of the content.

Columns can also be configured with click links and variables.

## Nesting Rules

Column components consist of the column set properties (`column_set`) and column containers (`column`). A column component can contain multiple column containers, each of which can embed multiple components.

The Card JSON 2.0 structure supports embedding all components except form containers (`form`) and table components (`table`).

## JSON Structure

The complete JSON 2.0 structure of the column component:

```json
{
    "schema": "2.0",
    "body": {
        "elements": [
            {
                "tag": "column_set",
                "element_id": "custom_id",
                "margin": "4px",
                "horizontal_spacing": "large",
                "horizontal_align": "left",
                "flex_mode": "none",
                "background_style": "default",
                "action": {
                    "multi_url": {
                        "url": "https://open.larksuite.com",
                        "pc_url": "https://open.feishu.com",
                        "ios_url": "https://developer.apple.com/",
                        "android_url": "https://developer.android.com/"
                    }
                },
                "columns": [
                    {
                        "tag": "column",
                        "element_id": "custom_id",
                        "background_style": "default",
                        "width": "auto",
                        "weight": 1,
                        "horizontal_spacing": "large",
                        "horizontal_align": "left",
                        "vertical_align": "center",
                        "vertical_spacing": "4px",
                        "direction": "vertical",
                        "padding": "8px",
                        "margin": "4px",
                        "action": {
                            "multi_url": {
                                "url": "https://www.baidu.com",
                                "pc_url": "https://www.baidu.com",
                                "ios_url": "https://www.google.com",
                                "android_url": "https://www.apple.com.cn"
                            }
                        },
                        "elements": []
                    }
                ]
            }
        ]
    }
}
```

## Column Set Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | The tag of the component. Fixed value: `column_set`. |
| `element_id` | No | String | — | Unique identifier for the component (JSON 2.0). Used to specify the component when calling component-related interfaces. Must be globally unique within the same card. Only letters, numbers, and underscores allowed; must start with a letter; max 20 characters. |
| `horizontal_spacing` | No | String | `8px` | Horizontal spacing of components within the container. Values: `small` (4px), `medium` (8px), `large` (12px), `extra_large` (16px), or a specific value in the range [0,99]px. |
| `horizontal_align` | No | String | `left` | Horizontal alignment of components within the container. Values: `left`, `center`, `right`. |
| `margin` | No | String | `0px` | Outer margin of the container. Range: [-99,99]px. Accepts single value (`"10px"`), double value (`"4px 0"`), or four values (`"4px 0 4px 0"`) for top/right/bottom/left. |
| `flex_mode` | No | String | `none` | Adaptive mode for narrow screens on mobile and PC. Values: `none` (compress column width proportionally), `stretch` (columns stack vertically at 100% width), `flow` (automatic line wrap), `bisect` (two-column equal layout), `trisect` (three-column equal layout). |
| `background_style` | No | String | `default` | Background color style of the column set. Values: `default` (white background, black in dark theme), or any color enumeration value / custom RGBA color. When columns are nested, the upper column color overrides the lower. |
| `action` | No | Action | — | Click interaction configuration. Currently only jump interaction is supported. If interactive components exist within the container, those interactions take precedence. |
| `action.multi_url` | No | Struct | — | Link addresses for each platform. |
| `action.multi_url.url` | No | String | — | Fallback jump link. |
| `action.multi_url.android_url` | No | String | — | Android jump link. Set to `lark://msgcard/unsupported_action` to disallow redirect on this platform. |
| `action.multi_url.ios_url` | No | String | — | iOS jump link. Set to `lark://msgcard/unsupported_action` to disallow redirect on this platform. |
| `action.multi_url.pc_url` | No | String | — | PC jump link. Set to `lark://msgcard/unsupported_action` to disallow redirect on this platform. |
| `columns` | Yes | column[] | — | Configuration of columns in the column layout. See Column Fields below. |

## Column Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | The tag of the column. Fixed value: `column`. |
| `element_id` | No | String | — | Unique identifier of the component (JSON 2.0). Must be globally unique within the same card. Only letters, numbers, and underscores allowed; must start with a letter; max 20 characters. |
| `background_style` | No | String | `default` | Background color style. Values: `default` (white background, black in dark theme), or any color enumeration value / custom RGBA color. |
| `width` | No | String | `auto` | Column width. Only effective when `flex_mode` is `none`. Values: `auto` (matches element width), `weighted` (distributed by `weight`), or a specific value in [16,600]px (supported in V7.4+). |
| `weight` | No | Number | `1` | Width proportion when `width` is `weighted`. Range: integer between 1 and 5. |
| `horizontal_spacing` | No | String | `8px` | Horizontal spacing of components within the column. Values: `small` (4px), `medium` (8px), `large` (12px), `extra_large` (16px), or a specific value in [0,99]px. |
| `horizontal_align` | No | String | `left` | Horizontal alignment. Values: `left`, `center`, `right`. |
| `vertical_align` | No | String | `top` | Vertical alignment. Values: `top`, `center`, `bottom`. |
| `vertical_spacing` | No | String | `8px` | Vertical spacing of components within the column. Values: `small` (4px), `medium` (8px), `large` (12px), `extra_large` (16px), or a specific value in [0,99]px. |
| `direction` | No | String | `vertical` | Arrangement direction. Values: `vertical`, `horizontal`. |
| `padding` | No | String | `0px` | Inner padding. Range: [0,99]px. Accepts single value (`"10px"`), double value (`"4px 0"`), or four values (`"4px 0 4px 0"`) for top/right/bottom/left. |
| `margin` | No | String | `0px` | Outer margin. Range: [-99,99]px. Accepts single value (`"10px"`), double value (`"4px 0"`), or four values (`"4px 0 4px 0"`) for top/right/bottom/left. |
| `elements` | No | Element[] or ColumnSet[] | — | Components embedded within the column. Supports all components except table and form containers. |
| `action` | No | Action | — | Click interaction configuration. Currently only jump interactions are supported. If interactive components exist within the container, those interactions take precedence. |
| `action.multi_url` | No | Struct | — | Link addresses for each platform. |
| `action.multi_url.url` | No | String | — | Fallback link address. |
| `action.multi_url.android_url` | No | String | — | Android link. Set to `lark://msgcard/unsupported_action` to disallow redirect. |
| `action.multi_url.ios_url` | No | String | — | iOS link. Set to `lark://msgcard/unsupported_action` to disallow redirect. |
| `action.multi_url.pc_url` | No | String | — | PC link. Set to `lark://msgcard/unsupported_action` to disallow redirect. |

## Example

The following JSON 2.0 example produces a personal efficiency overview card:

```json
{
    "schema": "2.0",
    "body": {
      "elements": [
        {
                "tag": "markdown",
                "content": ":YouAreTheBest:**个人效率总览** ",
                "text_align": "left",
                "text_size": "heading"
            },
            {
                "tag": "column_set",
                "flex_mode": "bisect",
                "horizontal_spacing": "",
                "horizontal_align": "center",
                "columns": [
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "top",
                        "vertical_spacing": "8px",
                        "direction": "horizontal",
                        "elements": [
                            {
                                "tag": "column_set",
                                "flex_mode": "none",
                                "horizontal_spacing": "8px",
                                "horizontal_align": "left",
                                "columns": [
                                    {
                                        "tag": "column",
                                        "width": "weighted",
                                        "vertical_align": "top",
                                        "vertical_spacing": "8px",
                                        "background_style": "grey",
                                        "padding": "8px",
                                        "elements": [
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
                                                                    "content": "已审批单量",
                                                                    "text_size": "normal",
                                                                    "text_align": "center",
                                                                    "text_color": "grey"
                                                                }
                                                            }
                                                        ],
                                                        "width": "weighted",
                                                        "weight": 1
                                                    }
                                                ]
                                            },
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
                                                                    "content": "29 单",
                                                                    "text_size": "heading",
                                                                    "text_align": "center",
                                                                    "text_color": "default"
                                                                }
                                                            }
                                                        ],
                                                        "width": "weighted",
                                                        "weight": 1
                                                    }
                                                ]
                                            },
                                            {
                                                "tag": "markdown",
                                                "content": "<text_tag color='blue'>高于部门 86%</text_tag>",
                                                "text_align": "center",
                                                "text_size": "normal"
                                            }
                                        ],
                                        "weight": 1
                                    }
                                ],
                                "margin": "0px 0px 0px 0px"
                            }
                        ],
                        "weight": 1
                    },
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "top",
                        "vertical_spacing": "8px",
                        "elements": [
                            {
                                "tag": "column_set",
                                "flex_mode": "none",
                                "horizontal_spacing": "8px",
                                "horizontal_align": "left",
                                "columns": [
                                    {
                                        "tag": "column",
                                        "width": "weighted",
                                        "vertical_align": "top",
                                        "vertical_spacing": "8px",
                                        "background_style": "grey",
                                        "padding": "8px",
                                        "elements": [
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
                                                                    "content": "审批平均耗时",
                                                                    "text_size": "normal",
                                                                    "text_align": "center",
                                                                    "text_color": "grey"
                                                                }
                                                            }
                                                        ],
                                                        "width": "weighted",
                                                        "weight": 1
                                                    }
                                                ]
                                            },
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
                                                                    "content": "0.9 小时",
                                                                    "text_size": "heading",
                                                                    "text_align": "center",
                                                                    "text_color": "default"
                                                                }
                                                            }
                                                        ],
                                                        "width": "weighted",
                                                        "weight": 1
                                                    }
                                                ]
                                            },
                                            {
                                                "tag": "markdown",
                                                "content": "<text_tag color='orange'>落后部门 61%</text_tag>",
                                                "text_align": "center",
                                                "text_size": "normal"
                                            }
                                        ],
                                        "weight": 1
                                    }
                                ],
                                "margin": "0px 0px 0px 0px"
                            }
                        ],
                        "weight": 1
                    },
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "top",
                        "vertical_spacing": "8px",
                        "elements": [
                            {
                                "tag": "column_set",
                                "flex_mode": "none",
                                "horizontal_spacing": "8px",
                                "horizontal_align": "left",
                                "columns": [
                                    {
                                        "tag": "column",
                                        "width": "weighted",
                                        "vertical_align": "top",
                                        "vertical_spacing": "8px",
                                        "background_style": "grey",
                                        "padding": "8px",
                                        "elements": [
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
                                                                    "content": "代批率",
                                                                    "text_size": "normal",
                                                                    "text_align": "center",
                                                                    "text_color": "grey"
                                                                }
                                                            }
                                                        ],
                                                        "width": "weighted",
                                                        "weight": 1
                                                    }
                                                ]
                                            },
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
                                                                    "content": "0%",
                                                                    "text_size": "heading",
                                                                    "text_align": "center",
                                                                    "text_color": "default"
                                                                }
                                                            }
                                                        ],
                                                        "width": "weighted",
                                                        "weight": 1
                                                    }
                                                ]
                                            },
                                            {
                                                "tag": "markdown",
                                                "content": "<text_tag color='green'>领先部门 100%</text_tag>",
                                                "text_align": "center",
                                                "text_size": "normal"
                                            }
                                        ],
                                        "weight": 1
                                    }
                                ],
                                "margin": "0px 0px 0px 0px"
                            }
                        ],
                        "weight": 1
                    }
                ],
                "margin": "16px 0px 0px 0px"
            },
            {
                "tag": "markdown",
                "content": ":STRIVE: **待优化的任务类型**",
                "text_align": "left",
                "text_size": "heading"
            },
            {
                "tag": "column_set",
                "flex_mode": "none",
                "horizontal_spacing": "8px",
                "horizontal_align": "left",
                "columns": [
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "center",
                        "vertical_spacing": "4px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<text_tag color='red'>1</text_tag> 加班申请",
                                "text_align": "left",
                                "text_size": "normal"
                            }
                        ],
                        "weight": 1
                    },
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "center",
                        "vertical_spacing": "8px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<font color='grey'>低于部门</font> <font color='red'>95% </font><font color='grey'>的审批人</font> ",
                                "text_align": "right",
                                "text_size": "notation"
                            }
                        ],
                        "weight": 1
                    }
                ],
                "margin": "16px 0px 0px 0px"
            },
            {
                "tag": "column_set",
                "flex_mode": "none",
                "horizontal_spacing": "8px",
                "horizontal_align": "left",
                "columns": [
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "center",
                        "vertical_spacing": "4px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<text_tag color='red'>2</text_tag> 休假申请",
                                "text_align": "left",
                                "text_size": "normal"
                            }
                        ],
                        "weight": 1
                    },
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "center",
                        "vertical_spacing": "4px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<font color='grey'>低于部门</font> <font color='red'>55% </font><font color='grey'>的审批人</font> ",
                                "text_align": "right",
                                "text_size": "notation"
                            }
                        ],
                        "weight": 1
                    }
                ],
                "margin": "16px 0px 0px 0px"
            },
            {
                "tag": "markdown",
                "content": ":CheckMark:**效率高的任务类型**",
                "text_align": "left",
                "text_size": "heading"
            },
            {
                "tag": "column_set",
                "flex_mode": "none",
                "horizontal_spacing": "8px",
                "horizontal_align": "left",
                "columns": [
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "center",
                        "vertical_spacing": "4px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<text_tag color='green'>1</text_tag> 数据权限申请",
                                "text_align": "left",
                                "text_size": "normal"
                            }
                        ],
                        "weight": 1
                    },
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "center",
                        "vertical_spacing": "4px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<font color='grey'>高于部门</font> <font color='green'>68% </font><font color='grey'>的审批人</font> ",
                                "text_align": "right",
                                "text_size": "notation"
                            }
                        ],
                        "weight": 1
                    }
                ],
                "margin": "16px 0px 0px 0px"
            },
            {
                "tag": "column_set",
                "flex_mode": "none",
                "horizontal_spacing": "8px",
                "horizontal_align": "left",
                "columns": [
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "center",
                        "vertical_spacing": "4px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<text_tag color='green'>2</text_tag> BOT推送消息",
                                "text_align": "left",
                                "text_size": "normal"
                            }
                        ],
                        "weight": 1
                    },
                    {
                        "tag": "column",
                        "width": "weighted",
                        "vertical_align": "center",
                        "vertical_spacing": "4px",
                        "elements": [
                            {
                                "tag": "markdown",
                                "content": "<font color='grey'>高于部门</font> <font color='green'>56% </font><font color='grey'>的审批人</font> ",
                                "text_align": "right",
                                "text_size": "notation"
                            }
                        ],
                        "weight": 1
                    }
                ],
                "margin": "16px 0px 0px 0px"
            }
        ]
    },
    "i18n_header": {
        "zh_cn": {
            "title": {
                "tag": "plain_text",
                "content": "我的近期审批效率"
            },
            "subtitle": {
                "tag": "plain_text",
                "content": "日期范围：2024.03.01 至 2024.03.31"
            },
            "template": "orange",
            "icon": {
                "tag": "standard_icon",
                "token": "approval_colorful"
            }
        }
    }
}
```

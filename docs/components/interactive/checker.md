# Checker

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/interactive-components/checker>

The checker is an interactive component that supports configurable callback responses, mainly used in task-checking scenarios.

## Notes

- The checker can only be used by writing card JSON code; it is not yet supported in the card building tool.

## Nesting Rules

The checker component supports nesting within all container components, including form containers, interactive containers, columns, and collapsible panels.

## JSON Structure

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "checker",
        "element_id": "custom_id",
        "name": "check_1",
        "checked": false,
        "text": {
          "tag": "plain_text",
          "content": "",
          "text_size": "normal",
          "text_color": "default",
          "text_align": "left"
        },
        "overall_checkable": true,
        "button_area": {
          "pc_display_rule": "always",
          "buttons": [
            {
              "tag": "button",
              "type": "text",
              "size": "small",
              "text": {
                "tag": "plain_text",
                "content": "text button"
              },
              "icon": {
                "tag": "standard_icon",
                "token": "chat-forbidden_outlined",
                "color": "orange",
                "img_key": "img_v2_38811724"
              },
              "disabled": false,
              "behaviors": []
            }
          ]
        },
        "checked_style": {
          "show_strikethrough": true,
          "opacity": 1
        },
        "margin": "0px",
        "padding": "0px",
        "confirm": {
          "title": {
            "tag": "plain_text",
            "content": "title"
          },
          "text": {
            "tag": "plain_text",
            "content": "content"
          }
        },
        "behaviors": [
          {
            "type": "callback",
            "value": {
              "key": "value"
            }
          }
        ],
        "hover_tips": {},
        "disabled": false,
        "disabled_tips": {}
      }
    ]
  }
}
```

## Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | The tag of the component. Fixed value: `checker`. |
| `element_id` | No | String | Empty | Unique identifier for the operation component (JSON 2.0). Used to specify the component when calling related interfaces. Must be globally unique within the card. Only letters, numbers, and underscores are allowed; must start with a letter and not exceed 20 characters. |
| `margin` | No | String | `0` | Margin of the component (JSON 2.0). Value range: `[-99,99]px`. Supports single value (e.g. `"10px"`), double value (e.g. `"4px 0"`), or four values (e.g. `"4px 0 4px 0"`) separated by spaces. |
| `padding` | No | String | `0` | Padding of the component (JSON 2.0). Value range: `[-99,99]px`. Supports single value (e.g. `"4px"`) or four values (e.g. `"4px 12px 4px 12px"`) separated by spaces. |
| `name` | No | String | — | Unique identifier for the checker component. Used to identify which component the submitted data belongs to. **Note:** Required and must be unique within the card when nested in a form container. |
| `checked` | No | Boolean | `false` | Initial check state. `true`: checked, `false`: unchecked. |
| `text` | No | Object | — | Plain text information within the checker component. |
| `text.tag` | Yes | String | `plain_text` | Text type tag. `plain_text`: plain text. `lark_md`: text with partial Markdown syntax support. |
| `text.content` | Yes | String | — | Text content. When `tag` is `lark_md`, supports partial Markdown syntax. |
| `text.text_size` | No | String | `normal` | Text size. `normal` (14px), `heading` (16px), `notation` (12px). |
| `text.text_color` | No | String | `default` | Text color. Only effective when `tag` is `plain_text`. `default`: black in light theme, white in dark theme. Also accepts color enumeration values. |
| `text.text_align` | No | String | `left` | Text alignment. `left`, `center`, or `right`. |
| `overall_checkable` | No | Boolean | `true` | Whether the entire checker has a shadow effect on hover. **Note:** To cancel the shadow, set `overall_checkable` to `false` and ensure `pc_display_rule` is not `on_hover`. |
| `button_area` | No | Object | — | Button area configuration. |
| `button_area.pc_display_rule` | No | String | `always` | Button display rule on PC. Buttons are always displayed on mobile. `always`: always displayed. `on_hover`: displayed with shadow effect on hover. |
| `button_area.buttons` | No | Array\<Object\> | `[]` | Buttons in the checker. Up to three buttons can be configured. See the **buttons Fields** section below. |
| `checked_style` | No | Object | — | Checked state style. |
| `checked_style.show_strikethrough` | No | Boolean | `false` | Whether to show a strikethrough line in the content area. |
| `checked_style.opacity` | No | Number | `1` | Opacity of the content area. Value range: `[0, 1]`. |
| `confirm` | No | Struct | Disabled | Secondary confirmation dialog configuration. A popup appears when the user submits; the action proceeds only after the user clicks confirm. |
| `confirm.title` | Yes | Struct | — | Title of the confirmation dialog. |
| `confirm.title.tag` | Yes | String | `plain_text` | Tag for the title text. Fixed value: `plain_text`. |
| `confirm.title.content` | Yes | String | — | Content of the confirmation dialog title. |
| `confirm.text` | Yes | Struct | — | Body text of the confirmation dialog. |
| `confirm.text.tag` | Yes | String | `plain_text` | Tag for the body text. Fixed value: `plain_text`. |
| `confirm.text.content` | Yes | String | — | Content of the confirmation dialog body text. |
| `behaviors` | Yes | Struct | — | Interaction types and behaviors. If not configured, the user can check locally only. See Configuring Card Interactions. |
| `hover_tips` | No | Object | Empty | Tooltip text on hover (PC only). **Note:** When both `hover_tips` and `disabled_tips` are configured, `disabled_tips` takes precedence. |
| `hover_tips.tag` | No | String | `plain_text` | Tag for the tooltip text. Fixed value: `plain_text`. |
| `hover_tips.content` | No | String | Empty | Content of the tooltip text. |
| `disabled` | No | Boolean | `false` | Whether to disable the checker. `true`: disabled, `false`: available. |
| `disabled_tips` | No | Object | Empty | Tooltip text on hover when the checker is disabled (PC only). |
| `disabled_tips.tag` | Yes | String | `plain_text` | Tag for the disabled tooltip text. Fixed value: `plain_text`. |
| `disabled_tips.content` | Yes | String | Empty | Content of the disabled tooltip text. |

### buttons Fields

Up to three buttons can be configured in the checker via the `buttons` array.

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | `button` | Tag of the button. Fixed value: `button`. |
| `type` | Yes | String | — | Button type. `text` (black font, no border), `primary_text` (blue font, no border), `danger_text` (red font, no border). |
| `size` | No | String | `medium` | Button size. `tiny` (24px PC / 28px mobile), `small` (28px PC / 28px mobile), `medium` (32px PC / 36px mobile), `large` (40px PC / 48px mobile). |
| `text` | No | Struct | Empty | Text on the button. |
| `text.tag` | Yes | String | `plain_text` | Tag of the text. Fixed value: `plain_text`. |
| `text.content` | Yes | String | Empty | Content of the text. |
| `icon` | No | Object | — | Prefix icon. Supports standard icons or custom icons. |
| `icon.tag` | No | String | — | Icon type tag. `standard_icon`: icon library. `custom_icon`: custom image. |
| `icon.token` | No | String | — | Icon token from the icon library. Effective when `tag` is `standard_icon`. |
| `icon.color` | No | String | — | Icon color. Supports outline and filled icons. Effective when `tag` is `standard_icon`. See Color Enumeration Values. |
| `icon.img_key` | No | String | — | Custom icon image key. Effective when `tag` is `custom_icon`. Obtain via the Upload Image API. |
| `disabled` | No | Boolean | `false` | Whether to disable the button. `true`: disabled, `false`: available. |
| `behaviors` | Yes | Struct | — | Interaction types and behaviors. See Configuring Card Interactions. |

## Callback Structure

After configuring interactions, callback data is sent to the request address configured in the developer console when users interact with the checker.

- New card action trigger (`card.action.trigger`): refer to Card Callback Communication.
- Old card action trigger (`card.action.trigger_v1`): refer to Message Card Callback (Old).

## Example

```json
{
  "schema": "2.0",
  "header": {
    "template": "blue",
    "title": {
      "tag": "plain_text",
      "content": "勾选组件（依赖端版本 7.9+)"
    }
  },
  "body": {
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
            "vertical_spacing": "1px",
            "elements": [
              {
                "tag": "checker",
                "name": "check_1",
                "checked": false,
                "text": {
                  "tag": "lark_md",
                  "content": "完成新品上市计划报告 💬[战略研讨会](https://open.larksuite.com)"
                },
                "overall_checkable": false,
                "button_area": {
                  "pc_display_rule": "always",
                  "buttons": [
                    {
                      "tag": "button",
                      "type": "text",
                      "size": "large",
                      "text": {
                        "tag": "plain_text",
                        "content": ""
                      },
                      "icon": {
                        "tag": "standard_icon",
                        "token": "forward-com_outlined",
                        "color": "grey-500"
                      },
                      "disabled": false,
                      "behaviors": [
                        {
                          "type": "callback",
                          "value": {
                            "key": "btn1"
                          }
                        }
                      ]
                    },
                    {
                      "tag": "button",
                      "type": "text",
                      "size": "large",
                      "text": {
                        "tag": "plain_text",
                        "content": ""
                      },
                      "icon": {
                        "tag": "standard_icon",
                        "token": "tab-todo_outlined",
                        "color": "grey-500"
                      },
                      "disabled": false,
                      "behaviors": [
                        {
                          "type": "open_url",
                          "default_url": "https://www.baidu.com",
                          "android_url": "https://developer.android.com/",
                          "ios_url": "lark://msgcard/unsupported_action",
                          "pc_url": "https://www.windows.com"
                        }
                      ]
                    }
                  ]
                },
                "checked_style": {
                  "show_strikethrough": true,
                  "opacity": 0.5
                },
                "padding": "2px 2px 2px 2px",
                "behaviors": [
                  {
                    "type": "callback",
                    "value": {
                      "key": "todo1"
                    }
                  }
                ]
              },
              {
                "tag": "checker",
                "name": "check_2",
                "checked": false,
                "text": {
                  "tag": "lark_md",
                  "content": "把材料提前给💬[业务数据共享群](https://open.larksuite.com)审阅"
                },
                "overall_checkable": true,
                "button_area": {
                  "pc_display_rule": "on_hover",
                  "buttons": [
                    {
                      "tag": "button",
                      "type": "text",
                      "size": "large",
                      "text": {
                        "tag": "plain_text",
                        "content": ""
                      },
                      "icon": {
                        "tag": "standard_icon",
                        "token": "forward-com_outlined",
                        "color": "grey-500"
                      },
                      "disabled": false,
                      "behaviors": [
                        {
                          "type": "callback",
                          "value": {
                            "key": "btn2"
                          }
                        }
                      ]
                    },
                    {
                      "tag": "button",
                      "type": "text",
                      "size": "large",
                      "text": {
                        "tag": "plain_text",
                        "content": ""
                      },
                      "icon": {
                        "tag": "standard_icon",
                        "token": "tab-todo_outlined",
                        "color": "grey-500"
                      },
                      "disabled": false,
                      "behaviors": [
                        {
                          "type": "open_url",
                          "default_url": "https://www.baidu.com",
                          "android_url": "https://developer.android.com/",
                          "ios_url": "lark://msgcard/unsupported_action",
                          "pc_url": "https://www.windows.com"
                        }
                      ]
                    }
                  ]
                },
                "checked_style": {
                  "show_strikethrough": true,
                  "opacity": 0.5
                },
                "padding": "2px 2px 2px 2px",
                "confirm": {
                  "title": {
                    "tag": "plain_text",
                    "content": "弹窗标题"
                  },
                  "text": {
                    "tag": "plain_text",
                    "content": "确认提交吗"
                  }
                },
                "behaviors": [
                  {
                    "type": "callback",
                    "value": {
                      "key": "todo2"
                    }
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

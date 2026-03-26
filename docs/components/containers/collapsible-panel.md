# Collapsible Panel

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/containers/collapsible-panel>

The collapsible panel allows folding of secondary information, such as notes and longer texts, within a card to highlight primary information.

This document describes the JSON 2.0 structure of the collapsible panel component. For the historical JSON 1.0 structure, refer to [collapsible panel](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/containers/collapsible-panel).

## Notes

- Collapsible panels can only be used by writing card JSON code and are not yet supported in card-building tools.
- Container components support a maximum of five layers of nested components. Avoid nesting multiple layers in collapsible panels, as multi-layer nesting compresses display space and affects card presentation.

## Nesting Rules

Collapsible panels do not support embedding form (`form`) components.

## JSON Structure

The complete JSON 2.0 structure of the collapsible panel component:

```json
{
  "schema": "2.0",
  "body": {
    "elements": [
      {
        "tag": "collapsible_panel",
        "element_id": "custom_id",
        "direction": "vertical",
        "vertical_spacing": "8px",
        "horizontal_spacing": "8px",
        "vertical_align": "top",
        "horizontal_align": "left",
        "padding": "8px 8px 8px 8px",
        "margin": "0px 0px 0px 0px",
        "expanded": true,
        "background_color": "grey",
        "header": {
          "title": {
            "tag": "markdown",
            "content": "**Panel title text**"
          },
          "background_color": "grey",
          "vertical_align": "center",
          "padding": "4px 0px 4px 8px",
          "position": "top",
          "width": "auto",
          "icon": {
            "tag": "standard_icon",
            "token": "chat-forbidden_outlined",
            "color": "orange",
            "img_key": "img_v2_38811724",
            "size": "16px 16px"
          },
          "icon_position": "follow_text",
          "icon_expanded_angle": -180
        },
        "border": {
          "color": "grey",
          "corner_radius": "5px"
        },
        "elements": [
          {
            "tag": "markdown",
            "content": "A long text"
          }
        ]
      }
    ]
  }
}
```

## Fields

### Panel Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `tag` | Yes | String | — | The tag of the component. Fixed value: `collapsible_panel`. |
| `expanded` | No | Boolean | `false` | Whether the panel is expanded. `true`: expanded, `false`: collapsed. |
| `element_id` | No | String | — | Unique identifier of the component (JSON 2.0). Must be globally unique within the same card. Only letters, numbers, and underscores allowed; must start with a letter; max 20 characters. |
| `margin` | No | String | `0px` | Outer margin of the container. Range: [-99,99]px. Accepts single value (`"10px"`), double value (`"4px 0"`), or four values (`"4px 0 4px 0"`) for top/right/bottom/left. |
| `horizontal_spacing` | No | String | `8px` | Horizontal spacing of components within the container. Values: `small` (4px), `medium` (8px), `large` (12px), `extra_large` (16px), or a specific value in [0,99]px. |
| `horizontal_align` | No | String | `left` | Horizontal alignment. Values: `left`, `center`, `right`. |
| `vertical_spacing` | No | String | `12px` | Vertical spacing of components within the container. Values: `small` (4px), `medium` (8px), `large` (12px), `extra_large` (16px), or a specific value in [0,99]px. |
| `vertical_align` | No | String | `top` | Vertical alignment. Values: `top`, `center`, `bottom`. |
| `direction` | No | String | `vertical` | Arrangement direction. Values: `vertical`, `horizontal`. |
| `padding` | No | String | `0px` | Padding of the container. Range: [0,99]px. Accepts single value (`"10px"`), double value (`"4px 0"`), or four values (`"4px 0 4px 0"`) for top/right/bottom/left. |
| `background_color` | No | String | — | Background color of the panel. Default is transparent. See Color Enumeration Values for options. |
| `header` | Yes | Object | — | Title settings of the collapsible panel. See Header Fields below. |
| `border` | No | Object | — | Border settings. No border is displayed by default. |
| `border.color` | No | String | `grey` | Border color. See Color Enumeration Values for options. |
| `border.corner_radius` | No | String | `5px` | Corner radius of the border. |
| `elements` | No | Array | — | JSON structure of nested components. Form components are not supported. |

### Header Fields

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `header.title` | No | Object | — | Title text settings. |
| `header.title.tag` | Yes | String | — | Text type tag. Values: `plain_text` (plain text), `markdown` (rich text; see Rich Text Component for supported syntax). |
| `header.title.content` | No | String | — | The content of the collapsible panel title. |
| `header.background_color` | No | String | — | Background color of the title area. Default is transparent. See Color Enumeration Values. If not set, the title area inherits the panel's `background_color`. |
| `header.width` | No | String | `fill` | Width of the title element (JSON 2.0, Lark V7.32+). Values: `fill` (equal width with panel), `auto` (adapts to text length), `auto_when_fold` (adapts to text length only when folded). |
| `header.vertical_align` | No | String | `center` | Vertical alignment of the title area. Values: `top`, `center`, `bottom`. |
| `header.padding` | No | String | `0` | Padding of the title area. Range: [-99,99]px. Accepts single, double, or four values. |
| `header.icon` | No | Object | — | Prefix/suffix icon for the title. Supports custom or library icons. |
| `header.icon.tag` | No | String | — | Icon type tag. Values: `standard_icon` (library icon), `custom_icon` (custom image). |
| `header.icon.token` | No | String | — | Icon token from the icon library. Effective when `tag` is `standard_icon`. See Icon Library for values. |
| `header.icon.color` | No | String | — | Icon color for linear/solid icons (tokens ending in `outlined` or `filled`). Effective when `tag` is `standard_icon`. See Color Enumeration Values. |
| `header.icon.img_key` | No | String | — | Image key of the custom icon. Effective when `tag` is `custom_icon`. Obtain via the Upload Image API. |
| `header.icon.size` | No | String | `10px 10px` | Icon size. Format: `"[1,999]px [1,999]px"`. |
| `header.icon_position` | No | String | `right` | Icon position. Values: `left` (far left of title area), `right` (far right of title area), `follow_text` (right of the text). |
| `header.icon_expanded_angle` | No | Number | `180` | Rotation angle of the icon when the panel is expanded. Positive values are clockwise, negative are counterclockwise. Values: `-180`, `-90`, `90`, `180`. |

## Example

The following JSON 2.0 example demonstrates collapsible panels in different configurations:

```json
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
```

# Card JSON v2.0 Breaking Changes & Release Notes

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-breaking-changes-release-notes>
>
> Last updated: 2025-06-27

This document introduces the differences in incompatibility and compatible optimizations between Card JSON 2.0 and 1.0 versions. For complete JSON 2.0 structure data, refer to [Card JSON 2.0 structure](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-structure).

## Notes

- The Card JSON 2.0 structure is supported by **Lark client version 7.20 and later**. When a card with the JSON v2.0 structure is sent to clients with versions lower than 7.20, the card title will be displayed normally, but the content will show a fallback upgrade prompt message.
- The Card JSON 2.0 structure currently only supports **shared cards** and does not support exclusive card configurations. That is, the `update_multi` parameter only supports being set to `true`.

---

## Incompatibility Changes

### Card Interaction Validity Change

| | 1.0 Structure | 2.0 Structure |
|---|---|---|
| Interactive time | 30 days | 14 days |
| Update time | 14 days (if interacted between day 14–30, update action won't take effect) | 14 days (unified) |

### Property Validation Change

In JSON 2.0 version, passing unsupported properties will result in an error.

| 1.0 Structure | 2.0 Structure |
|---|---|
| Passing unsupported properties will be ignored | Passing unsupported properties will result in an error |

### JSON Global Structure & Default Value Changes

#### Structure Changes

- JSON 2.0 adds a `body` field, with the `elements` attribute placed under the `body` level.
- JSON 2.0 no longer supports setting global multi-language through the `i18n_elements` field. Use component-level multi-language fields such as `i18n_content` instead.

#### `fallback` Field Changes

JSON 2.0 temporarily does **not** support using the `fallback` field to configure custom global downgrade rules.

#### Default Value Changes

In JSON 2.0, the default value of `update_multi` is changed to `true`.

**1.0 Structure:**

```json
{
  "schema": "1.0",
  "config": {
    "update_multi": true // Default: false
  },
  "card_link": {},
  "header": {},
  "i18n_header": {},
  "elements": [],
  "i18n_elements": {},
  "fallback": {}
}
```

**2.0 Structure:**

```json
{
  "schema": "2.0",
  "config": {
    "update_multi": true // Default: true
  },
  "card_link": {},
  "header": {},
  "body": {
    "elements": []
  },
  "fallback": {}
}
```

### Layout Property Default Value Changes

#### Form Container

The default values of `vertical_spacing` and `horizontal_spacing` changed from `16px` to `12px`, and developers can now customize these configurations.

| Property | 1.0 | 2.0 |
|---|---|---|
| `margin` | `"0"` | `"0"` |
| `padding` | `"0"` | `"0"` |
| `vertical_spacing` | `"16px"` | `"12px"` |
| `horizontal_spacing` | `"16px"` | `"12px"` |

#### Interactive Container

The default values for `vertical_spacing` and `horizontal_spacing` changed from `12px` to `4px` and `8px`, with support for custom configurations.

#### Collapsible Panel Padding

**With border or background color:**

| Property | 1.0 | 2.0 |
|---|---|---|
| `header.padding` | `"8px"` | `"4px 8px"` (top/bottom 4px, left/right 8px) |
| `padding` | `"8px"` | `"8px"` |
| `vertical_spacing` | `"8px"` | `"8px"` |
| `horizontal_spacing` | `"8px"` | `"8px"` |

**Without border or background color:**

| Property | 1.0 | 2.0 |
|---|---|---|
| `header.padding` | `"8px 0 8px 0"` | `"0"` |
| `padding` | `"0"` | `"8px 0 0 0"` (top 8px, others 0) |
| `vertical_spacing` | `"8px"` | `"8px"` |
| `horizontal_spacing` | `"8px"` | `"8px"` |

### `vertical_spacing` and `horizontal_spacing` Enumerations & Mapping Changes

| Enum | 1.0 Value | 2.0 Value |
|---|---|---|
| `small` | 4px | 4px |
| `medium` | 8px | 8px |
| `large` | 16px | 12px |
| `extra_large` | N/A | 16px |

### Header Component Configuration Changes

The `icon` configuration structure of the header component has been changed to align with other components.

**1.0 Structure:**

```json
{
  "header": {
    "title": {},
    "icon": {
      "img_key": "img_v2_38811724"
    },
    "ud_icon": {
      "token": "chat-forbidden_outlined",
      "style": {
        "color": "red"
      }
    }
  }
}
```

**2.0 Structure:**

```json
{
  "header": {
    "title": {},
    "icon": {
      "tag": "standard_icon",
      "token": "chat-forbidden_outlined",
      "color": "orange",
      "img_key": "img_v2_38811724"
    }
  }
}
```

### Image Component No Longer Supports Full Width Configuration

**1.0** supported `stretch_without_padding` to fill the card width:

```json
{
  "tag": "img",
  "img_key": "img_v3_0238_...",
  "size": "stretch_without_padding"
}
```

**2.0** no longer supports full width, but you can use negative margins:

```json
{
  "tag": "img",
  "img_key": "img_v3_0238_...",
  "size": "crop_center",
  "margin": "4px -12px"
}
```

### Markdown Component Deprecated Differential Link

The 2.0 structure no longer supports the following differential link syntax:

```json
{
  "tag": "markdown",
  "href": {
    "urlVal": {
      "url": "xxx",
      "pc_url": "xxx",
      "ios_url": "xxx",
      "android_url": "xxx"
    }
  },
  "content": "[Differential link]($urlVal)"
}
```

Use the `<link>` tag instead:

```
<link icon='chat_outlined' url='https://applink.larksuite.com/client/chat/xxxxx' pc_url='' ios_url='' android_url=''>chat</link>
```

### Fallback Height & Width Changes

| | 1.0 | 2.0 |
|---|---|---|
| Fallback height | 24px | 40px |
| Component width exceeds parent | Shrinks to fit (truncated only in interactive container) | Truncated |

### Deprecated Note Component & Action Module

The 2.0 structure no longer supports:

- **Note component** — Replace with a regular text component using notation font size, grey font color, and icon attributes.
- **Action module** (`"tag": "action"`) — Replace with a button or overflow button group component with appropriate `vertical_spacing` and `horizontal_spacing`.

---

## New Properties and Optimization Notes

### New `streaming_mode` Property for Streaming Updates

The 2.0 structure adds `streaming_mode` and `summary` fields to support card streaming updates.

```json
{
  "schema": "2.0",
  "config": {
    "streaming_mode": true,
    "summary": {
      "content": "Custom Content",
      "i18n_content": {
        "zh_cn": "",
        "en_us": "",
        "ja_jp": ""
      }
    }
  }
}
```

### New `element_id` Property for Component Operations

All components now include the `element_id` property as a unique identifier. Within the same card, the value must be **globally unique**, consist of letters/numbers/underscores only, begin with a letter, and cannot exceed 20 characters.

```json
{
  "tag": "button",
  "element_id": "button_1"
}
```

### Components Unified Support for Layout Related Capabilities

**Header and body level:**

```json
{
  "schema": "2.0",
  "header": {
    "title": {},
    "padding": "4px"
  },
  "body": {
    "vertical_spacing": "4px",
    "padding": "4px",
    "elements": []
  }
}
```

**Component level — new layout attributes:**

```json
{
  "tag": "xxxx",
  "margin": "4px",              // Outer margin, default "0", range [-99,99]px
  "padding": "4px",             // Inner margin, range [0,99]px
  "direction": "vertical",      // "vertical" | "horizontal", default "vertical"
  "horizontal_spacing": "3px",  // Range [0,99]px
  "vertical_spacing": "4px",    // Range [0,99]px
  "horizontal_align": "left",   // "left" | "center" | "right", default "left"
  "vertical_align": "center",   // "top" | "center" | "bottom", default "top"
  "elements": []
}
```

**Text component (`div`) now supports `width`:**

```json
{
  "tag": "div",
  "width": "fill" // "fill" | "auto" | "[16,999]px", default "fill"
}
```

### Rich Text Component Supports Standard Markdown Syntax

Card JSON 2.0 supports all standard Markdown syntax and some HTML syntax (except HTMLBlock). See [CommonMark Spec](https://spec.commonmark.org/) and [CommonMark playground](https://spec.commonmark.org/dingus/).

**Notes:**
- A single Enter key acts as a **soft break** (may be ignored depending on renderer).
- Two Enter keys act as a **hard break** (always displays as a new line).

**Supported HTML syntax:**
- Opening tags: `<br>`, `<hr>`
- Self-closing tags: `<br/>`, `<hr/>`
- Closing tags: `<person>`, `<local_datetime>`, `<at>`, `<a>`, `<text_tag>`, `<raw>`, `<link>`, `<font>` (supports nesting, e.g., `<font color=red>red<font color=green>green</font>again</font>`)

### Container Components Add Nested Component Types

In 2.0, **form containers**, **interactive containers**, **collapsible panels**, and **column components** can nest **all components** except:
- Form containers (`form`)
- Table components (`table`)

**1.0 limitations (now lifted):**
- Form containers: Could not nest tables, charts, form containers, or `div`-tagged components directly
- Interactive containers: Could only nest text, rich text, images, notes, columns, checkboxes, and interactive containers
- Collapsible panels: Could not nest form containers or tables
- Columns: Could not nest tables, forms, or multi-image combinations (`img_combination`)

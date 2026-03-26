# Component JSON v2.0 Overview

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-components/component-json-v2-overview>

Components in Lark cards can be categorized into container, display, and interactive components. Except for the recycling container, all components can be constructed using card JSON code. Except for collapsible panels, multi-image selection, and checkers, all components can be used with the card building tool. In the JSON structure, components declare their type by defining the `tag` field:

```json
{
    "tag": ""
}
```

This document summarizes the components built based on the card JSON 2.0 structure.

## Notes

The card JSON 2.0 structure supports Lark clients version 7.20 and above. When a card using the JSON 2.0 structure is sent to clients below version 7.20, the card title will be displayed normally, but the content will show a fallback upgrade prompt.

## Container Components

Container components are used for layout content or configuring interactive logic. Both display and interactive components can be added within container components.

| Component | Tag | Building Tool Support | Description |
|-----------|-----|----------------------|-------------|
| Column Set | `column_set` | Yes | Supports horizontal arrangement of multiple columns, freely combining image and text content within columns to build data tables, product or article lists, travel information, etc. |
| Loop Container | — | Yes | Supports embedding all display, interactive, and column components. Efficiently organizes similarly formatted content with different data. |
| Form Container | `form` | Yes | Allows users to enter a batch of form items locally and submit them in one callback to the developer's server via a submit button. |
| Interactive Container | `interactive_container` | No | Allows embedding components and flexibly combining multiple interactive containers to unify style and interactive capabilities for rich card interactions. |
| Collapsible Panel | `collapsible_panel` | No | Allows folding secondary information such as notes and longer texts to highlight the main information. |

## Display Components

Display components form the main content of the card and do not have interactive capabilities.

| Component | Tag | Building Tool Support | Description |
|-----------|-----|----------------------|-------------|
| Title | `header` | Yes | Used to build the title style and content, supporting main titles, subtitles, suffix tags, and title icons. |
| Plain Text | `div` | Yes | Supports adding plain text and prefix icons, and setting text size, color, alignment, and other display styles. |
| Rich Text | `markdown` | Yes | Supports rendering text, images, dividers, and other elements using Markdown syntax. |
| Image | `img` | Yes | Supports adding images by calling the Upload Image interface or uploading images in the card building tool. |
| Multi-Image Layout | `img_combination` | Yes | Supports adding multiple images by calling the Upload Image interface or uploading images in the card building tool. |
| Person | `person` | Yes | Displays usernames and avatars. Use by passing in the person's `open_id`, `user_id`, or `union_id`. |
| Person List | `person_list` | Yes | Displays multiple usernames and avatars. Use by passing in the `open_id`, `user_id`, or `union_id` of the individuals. |
| Chart | `chart` | No | Based on VChart chart definitions, supporting line charts, area charts, bar charts, pie charts, word clouds, and other data presentation methods. |
| Table | `table` | No | Supports adding plain text, option tags, person lists, and numeric content within the table. |
| Divider | `hr` | Yes | A horizontal line used to divide the content of the card for clearer presentation. |

## Interactive Components

Interactive components provide interactive capabilities for cards. When users receive cards that include interactive components, they can directly access links or handle business within the card.

| Component | Tag | Building Tool Support | Description |
|-----------|-----|----------------------|-------------|
| Input Box | `input` | Yes | Supports collecting unfixed text content such as reasons, evaluations, remarks, etc. |
| Button | `button` | Yes | Provides configured button callback interactions or link redirection capabilities, with multiple styles and sizes. |
| Overflow Button Set | `overflow` | Yes | Supports adding multiple buttons within the overflow set. By default collapsed; clicking displays all buttons inside. |
| Single-Select Dropdown | `select_static` | Yes | Supports customizing option text, icons, and callback parameters of the single-select menu. |
| Multi-Select Dropdown | `multi_select_static` | No | Supports customizing option text, icons, and callback parameters of the multi-select menu. |
| Single-Select Person Picker | `select_person` | Yes | Supports adding specific individuals as single-select options. |
| Multi-Select Person Picker | `multi_select_person` | No | Supports adding specific individuals as multi-select options. |
| Date Picker | `date_picker` | Yes | Supports providing date options. |
| Time Selector | `picker_time` | Yes | Supports providing time options. |
| Date-Time Picker | `picker_datetime` | No | Supports providing time and date options. |
| Image Picker | `select_img` | No | Supports providing image options, with single or multiple image selections. |
| Checker | `checker` | No | Supports configuring callback responses, mainly used for task checking scenarios. |

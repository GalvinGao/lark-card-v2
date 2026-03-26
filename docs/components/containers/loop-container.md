# Component Overview

> Source: <https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/containers/recycling-container>

Components in Lark Cards can be classified into container, display, and interactive types. The Card JSON declares a component by defining the `tag` field:

```json
{
  "tag": ""
}
```

## Container Components

Container components are used for layout content or configuring interactive logic. Both display and interactive components can be added within container components.

| Component | Client Version Requirement | Description |
|-----------|---------------------------|-------------|
| Column Set (`column_set`) | — | Column sets support horizontal arrangement of multiple columns, freely combining image and text content within columns to build data tables, product or article lists, travel information, etc., creating visually rich and interactive cards. |
| Form Container (`form`) | Lark V6.6 and above | The form container allows users to locally enter a batch of form items and, by clicking a submit button, send these locally cached form contents in one callback to the developer's server, achieving the effect of asynchronously submitting multiple form item data. |
| Interactive Container (`interactive_container`) | Lark V7.4 and above | The interactive container allows you to embed components based on business needs within the container and flexibly combine multiple interactive containers, unifying the style and interactive capabilities of multiple containers, to achieve various combination effects and rich card interactions. |
| Collapsible Panel (`collapsible_panel`) | Lark V7.9 and above | The collapsible panel allows you to fold secondary information, such as notes and longer texts, in the card to highlight the main information. |

## Display Components

Display components form the main content of the card and do not have interactive capabilities.

| Component | Client Version Requirement | Description |
|-----------|---------------------------|-------------|
| Title (`header`) | — | The title component is used to build the title style and content of the Lark card, supporting the addition of main titles, subtitles, suffix tags, and title icons. |
| Plain Text (`div`) | — | The plain text component supports adding plain text and prefix icons, and setting text size, color, alignment, and other display styles. |
| Rich Text (`markdown`) | — | The rich text (Markdown) component supports rendering text, images, dividers, and other elements. |
| Image (`img`) | — | The image component supports adding images by calling the Upload Image interface or uploading images in the new version of the Lark card building tool. |
| Multi-Image Layout (`img_combination`) | — | The multi-image layout component supports adding multiple images by calling the Upload Image interface or uploading images in the new version of the Lark card building tool. |
| Person (`person`) | — | The person component supports displaying usernames and avatars. You can use this component by passing in the person's open_id, user_id, or union_id. |
| Person List (`person_list`) | — | The person list component supports displaying multiple usernames and avatars. You can use this component by passing in the open_id, user_id, or union_id of the individuals. |
| Chart (`chart`) | Lark V7.1 and above | The chart component is based on VChart chart definitions, supporting line charts, area charts, bar charts, pie charts, word clouds, and other data presentation methods, helping you visualize various types of information and improving communication efficiency. |
| Table (`table`) | Lark V7.4 and above | The table component supports adding plain text, option tags, person lists, and numeric content within the table. |
| Divider (`hr`) | — | The divider component is a long horizontal line used to divide the content of the card, making the presentation clearer. |

## Interactive Components

Interactive components provide interactive capabilities for cards. When users receive cards that include interactive components, they can directly access links or handle business within the card.

| Component | Client Version Requirement | Description |
|-----------|---------------------------|-------------|
| Input Box (`input`) | Lark V6.8 and above | The input box component supports collecting unfixed text content, such as reasons, evaluations, remarks, etc. |
| Button (`button`) | — | The button component provides configured button callback interactions or link redirection capabilities, and supports multiple styles and sizes. |
| Overflow Button Set (`overflow`) | — | The overflow button set component supports adding multiple buttons within the overflow set. By default, the button set is in a collapsed state; clicking the set will display all buttons inside. |
| Single-Select Dropdown (`select_static`) | — | The single-select dropdown component supports customizing the option text, icons, and callback parameters of the single-select menu. |
| Multi-Select Dropdown (`multi_select_static`) | Lark V7.4 and above | The multi-select dropdown component supports customizing the option text, icons, and callback parameters of the multi-select menu. |
| Single-Select Person Picker (`select_person`) | — | The single-select person picker component supports adding specific individuals as single-select options. |
| Multi-Select Person Picker (`multi_select_person`) | Lark V7.4 and above | The multi-select person picker component supports adding specific individuals as multi-select options. |
| Date Picker (`date_picker`) | — | The date picker component supports providing date options. |
| Time Selector (`picker_time`) | — | The time selector component supports providing time options. |
| Date-Time Picker (`picker_datetime`) | — | The date-time picker component supports providing time and date options. |
| Image Picker (`select_img`) | Lark V7.6 and above | The image picker component supports providing image options, supporting single or multiple image selections. |
| Checker (`checker`) | Lark V7.9 and above | The checker supports configuring callback responses, mainly used for task checking scenarios. |

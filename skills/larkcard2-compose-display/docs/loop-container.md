# Loop Container

The loop container is a repeat/iteration component for rendering similarly formatted content with different data. It allows you to define a template layout once and have it repeated for each item in a data source.

## Overview

- Used to efficiently organize content that shares the same structure but displays different data values.
- Supports embedding all display components, interactive components, and column components inside the loop body.
- Each iteration renders the template with substituted data bindings.

## Usage

Loop containers are best constructed using the Lark Card Builder visual tool, which provides a drag-and-drop interface for configuring the loop template and data bindings.

For the full data binding syntax (variable references, iteration indices, conditional rendering within loops), refer to the official Lark/Feishu documentation:

- [Card JSON v2.0 — Loop Container](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/containers/loop-container)

## Notes

- The loop container is available in the card builder tools and through the JSON API.
- Data binding uses template variable syntax to reference fields from each item in the data array.
- Can be combined with other containers (column_set, collapsible_panel, etc.) for advanced layouts.

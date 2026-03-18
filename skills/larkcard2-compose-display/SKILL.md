---
name: larkcard2-compose-display
description: "Lark Card V2 display component reference — card structure, content components (text, markdown, image, divider, person, chart, table), and container components (column set, form, interactive container, collapsible panel, loop container). Use whenever composing or modifying Lark/Feishu message card JSON v2 with display-only or layout elements."
---

# Lark Card V2 — Display Components Reference

Reference for all display and layout components in Lark Card JSON v2 (`"schema": "2.0"`). Use these docs to compose valid card JSON without consulting external documentation.

## Table of Contents

### Card Structure & Globals

| Tag / Concept | Description | Doc |
|---|---|---|
| Card Structure | Overall JSON skeleton, config, body layout, and common attributes | [docs/card-structure.md](docs/card-structure.md) |
| Colors | Full color enum table (all hues × shades) + custom RGBA via config | [docs/colors.md](docs/colors.md) |
| Title (`header`) | Card header with title, subtitle, icon, color template, and tags | [docs/title.md](docs/title.md) |

### Content Components

| Tag | Description | Doc |
|---|---|---|
| Plain Text (`div`) | Text block with optional icon; supports plain text and lark_md | [docs/plain-text.md](docs/plain-text.md) |
| Rich Text (`markdown`) | Full CommonMark markdown block with Lark HTML extensions | [docs/rich-text.md](docs/rich-text.md) |
| Image (`img`) | Single image with scaling, corner radius, and preview options | [docs/image.md](docs/image.md) |
| Multi-Image Layout (`img_combination`) | Grid layout for 2-9 images in various arrangement modes | [docs/multi-image-layout.md](docs/multi-image-layout.md) |
| Divider (`hr`) | Simple horizontal line divider between card sections | [docs/divider.md](docs/divider.md) |
| Person (`person`) | Displays a single user's avatar and/or name by user ID | [docs/person.md](docs/person.md) |
| Person List (`person_list`) | Displays multiple users in a compact list with avatars | [docs/person-list.md](docs/person-list.md) |
| Chart (`chart`) | VChart-powered data visualization (line, bar, pie, etc.) | [docs/chart.md](docs/chart.md) |
| Table (`table`) | Structured data table with typed columns and pagination | [docs/table.md](docs/table.md) |

### Container Components

| Tag | Description | Doc |
|---|---|---|
| Column Set (`column_set`) | Multi-column layout container with weighted or fixed columns | [docs/column-set.md](docs/column-set.md) |
| Form Container (`form`) | Groups interactive inputs with submit/reset button actions | [docs/form-container.md](docs/form-container.md) |
| Interactive Container (`interactive_container`) | Clickable wrapper with callback or open_url behavior | [docs/interactive-container.md](docs/interactive-container.md) |
| Collapsible Panel (`collapsible_panel`) | Expandable/collapsible section with header and nested content | [docs/collapsible-panel.md](docs/collapsible-panel.md) |
| Loop Container | Repeat container for templated content with varying data | [docs/loop-container.md](docs/loop-container.md) |

---

> **Note:** For interactive components (buttons, inputs, pickers, selects), see the `larkcard2-compose-interactable` skill.

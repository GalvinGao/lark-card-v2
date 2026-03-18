# lark-card-v2

Go library for building [Lark Card JSON v2.0](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-structure) messages.

## Install

```bash
go get github.com/GalvinGao/lark-card-v2/lark
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/GalvinGao/lark-card-v2/lark"
)

func main() {
	card := lark.Card{
		Header: lark.Header{
			Title:    lark.PlainText("Deploy Notice"),
			Template: lark.Blue,
		},
		Body: lark.Body{
			Elements: lark.Elements{
				lark.Md("Service **user-api** deployed to production"),
				lark.Hr(),
				&lark.ColumnSet{
					Columns: []lark.Column{
						{
							Width:    "weighted",
							Weight:   1,
							Elements: lark.Elements{lark.Md("Left column")},
						},
						{
							Width:    "auto",
							Elements: lark.Elements{lark.Btn("Action")},
						},
					},
				},
			},
		},
	}

	data, _ := card.JSON()
	fmt.Println(string(data))
}
```

## Design

- **Struct-first** — Go structs with `json` tags map 1:1 to the Lark JSON spec. No builder pattern, no magic.
- **Tag injection** — The `Element` interface + `Elements` type auto-injects `"tag"` during JSON marshaling.
- **Zero dependencies** — stdlib only.
- **Shortcut constructors** — `Md()`, `PlainText()`, `MdText()`, `Btn()`, `Hr()`, `Img()` for convenience; they just return typed struct pointers.

## Components

All 25 Lark Card v2.0 component types are supported:

| Category | Components |
|---|---|
| Display | `Div`, `Markdown`, `Image`, `MultiImageLayout`, `Divider`, `Person`, `PersonList`, `Chart`, `Table` |
| Container | `ColumnSet`/`Column`, `Form`, `InteractiveContainer`, `CollapsiblePanel` |
| Interactive | `Button`, `Input`, `SelectStatic`, `MultiSelectStatic`, `SelectPerson`, `MultiSelectPerson`, `DatePicker`, `TimePicker`, `DateTimePicker`, `Overflow`, `ImagePicker`, `Checker` |

## Agent Skills

This repo ships [agent skills](https://skills.sh) that teach coding agents (Claude Code, Cursor, Codex, etc.) how to compose Lark Card v2 JSON.

### Install

```bash
npx skills add GalvinGao/lark-card-v2
```

### Available Skills

| Skill | Description |
|-------|-------------|
| `larkcard2-compose-display` | Card structure, content components (text, markdown, image, divider, person, chart, table), containers (column set, form, interactive container, collapsible panel), and color/text-size references |
| `larkcard2-compose-interactable` | Interactive components — buttons, inputs, dropdowns, person pickers, date/time pickers, overflow menus, image pickers, and checkers |

### Install a specific skill

```bash
npx skills add GalvinGao/lark-card-v2 --skill larkcard2-compose-display
```

### Install to a specific agent

```bash
npx skills add GalvinGao/lark-card-v2 -a claude-code
```

## License

MIT

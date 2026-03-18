# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Go library for building [Lark Card JSON v2.0](https://open.larksuite.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-json-v2-structure) messages. Zero dependencies (stdlib only). Module: `github.com/GalvinGao/lark-card-v2`.

## Commands

```bash
# Run all tests with race detection
go test ./... -race

# Run a single test
go test ./lark -run TestCardWithMarkdown -v

# Lint (matches CI)
golangci-lint run

# E2E test (sends real cards to Lark — requires .env with LARK_APP_ID, LARK_APP_SECRET, LARK_OPEN_ID)
go run ./internal/test/e2e/
```

## Architecture

### Tag Injection Pattern

The core design: Go structs map 1:1 to Lark JSON, but the `"tag"` field is **not** stored on structs. Instead, each component implements the `Element` interface (`cardTag() string`), and `Elements.MarshalJSON()` injects `"tag"` at serialization time.

```
Element interface  →  cardTag() returns "button", "markdown", etc.
Elements []Element →  MarshalJSON() injects {"tag": cardTag(), ...} for each element
Card.JSON()        →  wraps everything with {"schema": "2.0", ...}
```

**Where tag injection happens differently:** `Column`, `Button` (inside checker), and `TextTag` implement their own `MarshalJSON()` using the type-alias pattern because they appear in typed arrays (`[]Column`, `[]Button`, `[]TextTag`) rather than `Elements`.

### Key Files

- `lark/element.go` — `Element` interface + `Elements` custom marshaler
- `lark/json.go` — `Card.JSON()` schema injection
- `lark/helpers.go` — shortcut constructors (`Md()`, `PlainText()`, `Btn()`, `Hr()`, `Img()`)
- `lark/enums.go` — all color, size, spacing, and mode constants
- `lark/card_test.go` — 62 test cases covering all components

### Adding a New Component

1. Create a struct in `lark/` with `json` tags matching the Lark spec
2. Add `func (x *YourType) cardTag() string { return "your_tag" }` to satisfy `Element`
3. If the type appears in a typed slice (not `Elements`), add a custom `MarshalJSON()` using the type-alias pattern (see `column.go`)
4. Add a test case in `card_test.go`

## Lark Card v2 Docs

Component documentation from the Lark Open Platform lives in `docs/components/` (containers, content, interactive). Agent skills with cleaned references are in `skills/`.

## CI

GitHub Actions (`.github/workflows/ci.yml`): lint with golangci-lint v2.1.6, test with race detector. Pre-commit hook runs golangci-lint via prek.

# Overflow Menu

**Tag:** `overflow`

## Minimal Example

```json
{
  "tag": "overflow",
  "options": [
    {
      "text": { "tag": "plain_text", "content": "Option 1" },
      "value": "opt1",
      "multi_url": { "url": "https://example.com" }
    },
    {
      "text": { "tag": "plain_text", "content": "Option 2" },
      "value": "opt2"
    }
  ],
  "width": "default",
  "value": { "key": "overflow_click" },
  "margin": "0px"
}
```

## Key Fields

- **options**: Array of menu items (required)
  - **text**: `{ "tag": "plain_text", "content": "..." }` — display label
  - **value**: String — callback value for this option
  - **multi_url**: Optional URL action — `{ "url": "...", "pc_url": "...", "ios_url": "...", "android_url": "..." }` (only `url` is required, others are platform overrides)
- **width**: `"default"` | `"fill"` | `"[100,∞)px"`
- **value**: Custom callback value object sent when any option is selected
- **confirm**: Optional confirmation dialog before executing the action

## Notes

- Renders as a collapsed "..." button; clicking expands the full option list.
- Options with `multi_url` open the URL directly; options without it trigger a callback.
- Option `text` only supports `plain_text`, not `lark_md`.

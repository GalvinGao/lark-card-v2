# Title (Header)

Tag: placed in `header` at the card root level (not inside `body.elements`). One per card.

## JSON Example

```json
"header": {
  "title": { "tag": "plain_text", "content": "Main Title" },
  "subtitle": { "tag": "plain_text", "content": "Subtitle" },
  "text_tag_list": [
    { "tag": "text_tag", "text": { "tag": "plain_text", "content": "Tag 1" }, "color": "blue" }
  ],
  "template": "blue",
  "icon": { "tag": "standard_icon", "token": "chat-forbidden_outlined", "color": "orange" },
  "padding": "12px"
}
```

## Key Fields

- **title.tag**: `"plain_text"` or `"lark_md"`
- **title.content**: Main title text, max 4 lines
- **subtitle**: Optional subtitle, max 1 line. Same tag/content structure as title.
- **text_tag_list**: Up to 3 suffix tags displayed after the title. Each item:
  - `tag`: `"text_tag"`
  - `text`: `{ "tag": "plain_text", "content": "..." }`
  - `color`: color enum string
- **template**: Header background color theme. Values: `"blue"` | `"wathet"` | `"turquoise"` | `"green"` | `"yellow"` | `"orange"` | `"red"` | `"carmine"` | `"violet"` | `"purple"` | `"indigo"` | `"grey"` | `"default"`
- **icon**: Prefix icon before the title. Either:
  - `{ "tag": "standard_icon", "token": "icon_token", "color": "blue" }` — built-in icon
  - `{ "tag": "custom_icon", "img_key": "img_v3_xxxx" }` — uploaded image
- **padding**: Header padding, `"[0,99]px"`, default `"12px"`

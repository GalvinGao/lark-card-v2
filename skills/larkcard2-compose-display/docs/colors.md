# Color Enumerations & Custom RGBA Colors

## Properties That Support Color Enums / RGBA

- `text_color` in plain text (`div`) components
- `<font color='...'>` in `lark_md` and `markdown` content
- `icon.color` and `ud_icon.style.color`
- `background_style` in `column_set`, `column`, `interactive_container`
- `border_color` in `interactive_container`
- `background_color` and `border.color` in `collapsible_panel`
- `text_color` in `table` header_style
- `color` in `text_tag_list` (title suffix tags)

## Color Enumeration Table

Base color names (no suffix) are aliases for the `-600` shade.

| Enum | Hue | Light Mode | Dark Mode |
|------|-----|-----------|-----------|
| `blue` | Blue | #1456F0 | #75A4FF |
| `blue-50` | Blue | #F0F4FF | #152340 |
| `blue-100` | Blue | #E0E9FF | #173166 |
| `blue-200` | Blue | #C2D4FF | #194294 |
| `blue-300` | Blue | #94B4FF | #2655B6 |
| `blue-350` | Blue | #7AA2FF | #275FCE |
| `blue-400` | Blue | #5083FB | #3370EB |
| `blue-500` | Blue | #336DF4 | #4C88FF |
| `blue-600` | Blue | #1456F0 | #75A4FF |
| `blue-700` | Blue | #0442D2 | #8FB4FF |
| `blue-800` | Blue | #002F9E | #BDD2FF |
| `blue-900` | Blue | #002270 | #E0E9FF |
| `wathet` | Sky Blue | #076A94 | #25B2E5 |
| `wathet-50`..`wathet-900` | Sky Blue | #E7F8FE..#072B3D | #152830..#C7F0FF |
| `turquoise` | Turquoise | #067062 | #1AB7A1 |
| `turquoise-50`..`turquoise-900` | Turquoise | #E2F8F5..#02312A | #132926..#B7F7EF |
| `green` | Green | #1A7526 | #51BA43 |
| `green-50`..`green-900` | Green | #E4FAE1..#022C07 | #0E2B0A..#CBF5C6 |
| `lime` | Lime | #5C6D08 | #93AF04 |
| `lime-50`..`lime-900` | Lime | #F2F8D3..#262E00 | #212702..#E3F391 |
| `yellow` | Yellow | #865B03 | #FBCB46 |
| `yellow-50`..`yellow-900` | Yellow | #FBF4DF..#382201 | #30250A..#FFEAA3 |
| `sunflower` | Sunflower | #8F7C00 | #F5DF36 |
| `sunflower-50`..`sunflower-900` | Sunflower | #FFFFDB..#2C2502 | #29250A..#FFF59E |
| `orange` | Orange | #A44904 | #F3871B |
| `orange-50`..`orange-900` | Orange | #FFF3E5..#3B1A02 | #33210B..#FEE7CD |
| `red` | Red | #C02A26 | #F6827E |
| `red-50`..`red-900` | Red | #FEF0F0..#590603 | #3D1A19..#FEE3E2 |
| `carmine` | Carmine | #B82879 | #ED77BA |
| `carmine-50`..`carmine-900` | Carmine | #FEF0F8..#550C35 | #3A182B..#FFE0F2 |
| `violet` | Violet | #A630A6 | #E17FE1 |
| `violet-50`..`violet-900` | Violet | #FCEEFC..#520052 | #3B153B..#FCDFFC |
| `purple` | Purple | #7A35F0 | #B88FFE |
| `purple-50`..`purple-900` | Purple | #F5F0FE..#2F0080 | #2B194A..#EFE5FF |
| `indigo` | Indigo | #4752E6 | #9499F7 |
| `indigo-50`..`indigo-900` | Indigo | #F2F3FD..#151B70 | #1E204A..#E7E9FE |
| `grey` | Neutral | #646A73 | #A6A6A6 |
| `grey-00` | Neutral | #FFFFFF | #0A0A0A |
| `grey-50` | Neutral | #F5F6F7 | #1A1A1A |
| `grey-100` | Neutral | #F2F3F5 | #292929 |
| `grey-200` | Neutral | #EFF0F1 | #373737 |
| `grey-300` | Neutral | #DEE0E3 | #434343 |
| `grey-350` | Neutral | #D0D3D6 | #505050 |
| `grey-400` | Neutral | #BBBFC4 | #5F5F5F |
| `grey-500` | Neutral | #8F959E | #757575 |
| `grey-600` | Neutral | #646A73 | #A6A6A6 |
| `grey-650` | Neutral | #51565D | #CFCFCF |
| `grey-700` | Neutral | #373C43 | #E0E0E0 |
| `grey-800` | Neutral | #2B2F36 | #E8E8E8 |
| `grey-900` | Neutral | #1F2329 | #EBEBEB |
| `grey-950` | Neutral | #0F1114 | #F8F8F8 |
| `grey-1000` | Neutral | #000000 | #FFFFFF |
| `bg-white` / `white` | — | #FFFFFF | #1A1A1A |

## Custom RGBA Colors

Define custom colors in `config.style.color`, then reference by name:

```json
{
  "config": {
    "style": {
      "color": {
        "cus-0": {
          "light_mode": "rgba(5,157,178,0.52)",
          "dark_mode": "rgba(78,23,108,0.49)"
        }
      }
    }
  }
}
```

Then use `"cus-0"` anywhere a color enum is accepted:

```json
{ "tag": "div", "text": { "tag": "plain_text", "content": "Custom color", "text_color": "cus-0" } }
```

```json
{ "tag": "column", "background_style": "cus-0", ... }
```

```json
{ "tag": "interactive_container", "background_style": "cus-0", "border_color": "cus-1", ... }
```

```json
{ "icon": { "tag": "standard_icon", "token": "chat_outlined", "color": "cus-0" } }
```

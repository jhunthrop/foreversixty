# brand-consistency-reviewer: BA1 (addon and Android XML on the token source), round 1

Judged from the mock boards in `captures/`, the side-by-sides against the web (`side-by-side-*.png`), the Android captures and the code.

## Audit findings, one by one

| Audit finding | Now | Evidence |
|---|---|---|
| BLOCKER 1: navy and dull-gold palette | Fixed. Token umber ground; plates and hairlines match the web (`color.base.bg`, `raised`, `line`). | `side-by-side-panel.png`, `side-by-side-onboarding.png`; `Theme.HEX` is a key map onto `Tokens.color`. |
| BLOCKER 2: raw client rarity blues | Fixed. `Theme.qualityColor` reads `Tokens.color.rarity.*` with the `rareText` and `epicText` lifts. | `palette_spec.lua` asserts rare, epic and uncommon against the tokens. |
| SHOULD 1: Friz and Morpheus | Fixed. Marcellus and Open Sans 400/600/700 ship in `media/fonts` with their OFL texts; the client faces are fallbacks only. | `Tokens.lua` font block, `Theme.lua` FONT_ROLES. |
| SHOULD 2: forced all-caps | Fixed. No caps strings in `Locale.lua`; captures are sentence case. | |
| SHOULD 3: client-red buttons | Fixed. Secondary blue `114f72` / edge `1d6a96`, primary teal / edge `408076`, white label, as `web/src/lib/ui/button.ts`. | onboarding capture: Done teal, Back and Open my panel blue. |
| SHOULD 4: bevelled frames | Fixed. Hairline plates with flat heads and rules. | `panel-*.png` |
| Gate | A busted spec fails on any hex literal or `|cff` escape outside Tokens.lua. | `palette_spec.lua` |

## Android XML

`themes.xml`, `bg_live_answer*`, `widget_*` and `shortcut_ground` carry no hex; they read `@color/fs_*` from the generated `colors_gen.xml`. The splash sits on `fs_bg` with transparent system bars (`android-splash-and-bars.png`). Native idiom allowed: transparent bars, the shade grey `shadeAnswer`. The Compose Home screen's forced caps are out of scope (lane F).

## Remaining drift (NICE)

1. "+6.4 DPS" is one green run; the web mutes the unit.
2. "Your job:" target text looks like `clientGold` rather than `color.base.gold`.

VERDICT: CONSISTENT

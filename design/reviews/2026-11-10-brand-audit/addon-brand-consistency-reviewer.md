# brand-consistency-reviewer: in-game addon (2026-11-10)

This review is judged from the mock captures (`design/listing/gallery`, `2026-11-08-addon-ux`, `2026-11-10-addon-onboarding`, `2026-11-10-seal-flat`) together with the Lua code. The addon cannot run here.

## DRIFT

**BLOCKER**
1. **The panel palette is not the site's.** `Theme.HEX` at `addon/ForeverSixty/Theme.lua:15-98` claims to mirror `tokens.css`, but it carries the retired navy and gold system:

   | Key | Lua | Web token |
   |---|---|---|
   | background | 0d111a | `--color-bg` 100a06 |
   | gold | e5b955 | ffec71 |
   | body | e9e4d8 | e3dfd9 |
   | muted | 9a9484 | 8b8784 |

   A third palette (`umber`, `brass`, `bone`, and `teal` at 2f7a74 against the web's 0f6064) exists nowhere in tokens.json. Capture: `2026-11-08-addon-ux/panel-raider-60.png`, a navy panel, against the web's umber plates. Fix: a Lua target in build.mjs (see below).
2. **Rarity text on the dark panel uses the client's raw colours.** `Theme.qualityColor` (`Theme.lua:1259`) returns `ITEM_QUALITY_COLORS`, so rares render in dark blue 0070dd (gallery 03, 05). The web uses `rarity.rare-text` 3d94f0 and `epic-text` b866f5. The addon's own "bright" lift (`Theme.lua:62-68`: 6eaeec, cd92f5, ff8f1f) matches no token.

**SHOULD**
1. **Type.** The addon uses `Fonts\FRIZQT__.TTF` and `Fonts\MORPHEUS.TTF` (`Theme.lua:411-414`), and `media/` ships no fonts. The web uses Marcellus, Open Sans and Archivo, all OFL licensed. Ship Marcellus and Open Sans in `media/fonts`, with the client fonts as the fallback.
2. **Forced caps.** There are at least 13 all-caps eyebrows in `Locale.lua`: 120, 196, 245, 246, 251 and 381-422, for example "WHAT IS IN IT", "LOAD A BUILD" and "YOUR BUILD". The web uses sentence case.
3. **Buttons.** The onboarding captures show client-red buttons ("Skip tour", "Next") in `2026-11-10-addon-onboarding/onboarding-*.png`. Theme defines a teal primary (`tealTop` 3a8a83), but the web primary is teal 0f6064 with edge 408076. Check `Widgets.lua` button creation.
4. **Frame chrome.** The panel uses thick bevelled brown-gold frames. The web uses hairline `color.base.line` 3d342d and plates with radius 10.
5. **Figures.** "Forever Score 1840" has no thousands separator, and "+6.4 DPS" is one green run with no muted-unit anatomy. The web shows "1,102.8 DPS" with a white figure and a muted unit.
6. **Stale listing plates are shipping.** See the stale captures section below.

**NICE**
- `hover` 1f2739 and `raised` 161c2b are navy; the web uses 241b15 and 170e09.
- `notchHorde` b8261c against `--color-horde-bar` c0392b.
- `Theme.classColor` uses the client's `RAID_CLASS_COLORS`, not the web class tokens. The client table does not know Forever's new combinations.
- The parse table at `surfaces/Panel.lua:822-823` matches `parse.*`, except grey 9d9d9d, where the web's `parse.greyText` is 8c8c8c.
- Gallery 03 still shows "(needs bullets)", against the 2026-11-08 game-words rule.

## Theme.HEX vs tokens

| Key | Lua | Token | Name |
|---|---|---|---|
| background | 0d111a | 100a06 | color.base.bg |
| raised | 161c2b | 170e09 | color.base.raised |
| hover | 1f2739 | 241b15 | color.base.hover |
| border | 262e40 | 3d342d | color.base.line |
| body | e9e4d8 | e3dfd9 | color.base.text |
| muted | 9a9484 | 8b8784 | color.base.muted |
| gold | e5b955 | ffec71 | color.base.gold |
| goldHover | f2cd74 | f7b90a | color.base.goldHover |
| success | 6fcf8e | 7fd48a | color.outcome.kill |
| warning | ff6b5c | f0736b | color.outcome.wipe |
| rare/epic/legendaryBright | 6eaeec / cd92f5 / ff8f1f | 3d94f0 / b866f5 / ff8000 | rarity.*-text |
| teal / tealEdge | 2f7a74 / 173f3b | 0f6064 / 408076 | --color-teal / -edge |
| tooltipGround / Edge | 080a16 / 9095a3 | 060914 / 8b8f9c | --color-tooltip / -edge |
| card, inset, sidebar, track | navy | none | none |
| bone, ink, umber*, brass* | various | none | none |

Three keys already match: `clientGold` ffd100, `notchAlliance` and `shadow`.

About 11 keys are wrong and about 14 have no token at all.

## Hardcodes outside Theme.lua

There are few. The problem is that Theme.lua itself is the hardcode.
- Literal colours appear only in `surfaces/Panel.lua:822-823`, `surfaces/TourExample.lua:28,35` and `Panel.lua:1044`.
- There are 0 numeric `SetTextColor` or `SetVertexColor` literals.
- All `SetFont` calls go through Theme.

## Components

- **Item tooltip:** the client's `GameTooltip` is native and allowed. The added lines follow the web wording.
- **Crest:** nine ringed crests ship in `media/crests`. Classes without one fall back to the client `UI-Classes-Circles`.
- **Seal:** the minimap uses the flat seal at 32 and 20 px, which is correct.
- **Talent marks:** the gold comes from the stale e5b955.
- **Score tiers:** they use client rarity colours, not the web's `parse.*` and `rarity.*`.

## Stale captures

- `design/listing/gallery/*` was last committed 2026-11-03 (6b4e85898). It shows the navy panel and "needs bullets", and predates the umber panel, the v4 minimap, the tour and the quieter talent marks. These are the store images.
- `2026-11-08-addon-ux/*` predates the 11-10 restyle.
- `2026-11-10-addon-onboarding/*` is current.

## Proposed Lua target

1. Add `lua: addon/ForeverSixty/TokensGen.lua` to `OUTPUTS` in `design/tokens/build.mjs`. It emits `return { color = {...}, font = {...}, size = {...}, radius = {...} }`, with colours as six-hex strings so that `Theme.rgb` and `inlineColor` keep working.
2. Load `TokensGen.lua` in the .toc before `Theme.lua`.
3. Rewrite `Theme.HEX` as a thin key map onto `Tokens.color.*` so no caller changes. Promote the frame keys we keep (`bone`, `umber`, `brass`) into tokens.json with a `lua` extension, and fold or delete the navy keys.
4. Extend `tokens.test.mjs` to fail on a stale output, and add a busted test that fails on any hex literal in addon `*.lua` outside `TokensGen.lua`.

VERDICT: DRIFT

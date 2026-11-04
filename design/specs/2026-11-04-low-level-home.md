# Low-level home: a level 1-19 character gets a next step, never "Pick a spec"

Date 2026-11-04. Surfaces: home `/`, `/bis` landing, BiS spec page character card.
Reference: the in-game talent frame at level 10 (a glowing "points to spend" count, the next talent lit) and the Wowhead leveling guide's "talent order" strip. Borrow the order and the lit next point; nothing else.

## 1. Question and answer
"My character is level N. What do I do now?" The answer sits in the 56 px header answer box and the one gold button, visible at 360 and 1280 without scrolling. Level 1-9 and 10-19 answers differ; both name a real thing to do.

## 2. Rules
1. Below 20 nobody picks a spec. The spec is `levelingSpecFor(character)`: derived or addon spec if one exists, else the class entry of the new constant `LEVELING_SPEC_BY_CLASS` (druid feral, hunter beast-mastery, mage frost, paladin retribution, priest shadow, rogue combat, shaman enhancement, warlock affliction, warrior arms). Put it in `web/src/lib/leveling/leveling-spec.ts`; `specKeyForCharacter` falls back to it at every level.
2. At 20+ with no spec the same fallback applies: gear and upgrades compare against the leveling spec, and the upgrades note says "Compared as Beast Mastery". This is cleaner than a dead "Pick a spec" and costs nothing. `homeHeroCardsCopy.pickASpec`, `pickASpecLine` and `landingCopy.fallbackAction` are deleted.
3. One job per page: gear stays on /bis, talents in the planner, how to play in the guide. Cards link, never restate.
4. The header selector stays the only switcher. Section (f) needs nothing.
5. Talent order = first band (20-29) spend order of the leveling spec, computed in the browser from the BiS and talent files, same tooltips as guide chips (`data-game-tip`). Point k (1-based) is taken at level 9+k.

## 3. Resolved state (shared by all cards)
`spentTotal` and per-talent ranks come from `build.talents` when present, else unknown.
- Unknown export: assume the guide order was followed. Points owed = `level - 9`; the lit point is point `level - 9`.
- Known export: next = first talent in the order whose spent rank is below its order rank. Points to spend = `level - 9 - spentTotal`.
- Spent points below 20 define the spec (most points in one tree). If that tree differs from the leveling spec, use it: guide link, order and names follow the spent spec. Spec name is never asked for.
- Points to spend is 0 or less: next level for a point = `10 + spentTotal`.

## 4. Strings (verbatim; `{Spec}` = spec name, `{L}` = level, `{Talent}` = talent name)
### (a) Home header card (CharacterHeader, `homeAnswerSpec`, `homeAction`)
| Level | Answer (sentence, one line at 1280, max two at 360 inside 56 px) | Button (primary) | Href |
|---|---|---|---|
| 1-9 | `First talent point at level 10.` | `Level as {Spec} →` | `/guides/{class}/{spec}#leveling` |
| 10-19, point to spend | `Level {L} talent: {Talent}.` | `Talent order →` | `/planner?spec={specKey}&talents={band 20-29 string}` |
| 10-19, none to spend (known export) | `Next talent point at level {10+spentTotal}.` | `Level as {Spec} →` | guide `#leveling` |
| 20+ | unchanged | unchanged | unchanged |
Delete `landingCopy.under20` from the home header. Button text is sentence case in source; the existing uppercase CSS renders it.

### (b) Talents card (HomeHeroCards)
Head: `Talents`. Sub line under the head, muted 12 px: `Leveling as {Spec}`. Quiet text link right of it: `Other specs` to `/guides/{class}` (verify the route; if absent use `/guides`).
- 1-9 body: `Your first point is at level 10.` Below it, a muted line: `First pick: {Talent}.` (point 1 icon tile + name, tooltip). Action `ActionLink`: `Open the leveling guide` to `#leveling`.
- 10-19 body: row of up to 5 IconTiles (point to spend first, lit: gold 1 px ring, others 60% opacity), each with rank chip `1/5` and `data-game-tip`. Line under: `{n} to spend now` (n = points to spend; `1 to spend now` singular). If n is 0: `Next point at level {X}.` and the row shows the next 3 tiles dimmed. Action: `Open the planner` to the Talent order href above.
- Phone 390: five 40 px tiles with 8 px gap fit one row (232 px); never wrap, show 4 if the card is under 300 px.
- 20+: unchanged, with the leveling-spec fallback.
Testids: `home-talents-order`, `home-talents-tile-{talentId}`.

### (c) Simulator card
Below 20 the card stays in its grid cell (no shift) but is quiet: no figure, no `ActionLink`, muted text. Head `Simulator`. Body: `Sim numbers start at level 20.` One line only. 20+: unchanged.
Testid: `home-sim-locked`.

### (d) Your upgrades panel
Heading `Your upgrades` stays. Below 20 the table and the "PICK A SPEC" row are gone. Body: `Gear lists start at level 20.` Next line: `Until then, wear the best quest rewards you can use.` One `ActionLink`: `Open the leveling guide` to `#leveling`. Min height equals the loaded panel's reserve so nothing shifts.
Testid: `home-upgrades-locked`.

### (e) /bis landing and BiS spec page character card
Answer stays `Best in slot starts at 20.` (it is the true answer on a gear page). Button: `{Spec} gear for 20 to 29 →` using the leveling spec (rule 1), href is the 20-29 band of that page. A class page the visitor opened keeps its own spec. Under 10 add no second sentence. `Pick a spec →` never renders.

### (f) Header character menu
No change.

## 5. States (every new interactive element)
Button and ActionLink: existing gold/inline styles; hover, focus-visible ring, active as shipped. Talent tiles: hover shows the game tooltip, keyboard focus (`tabindex=0`) shows the same; touch tap shows it and a second tap dismisses. Loading: the talent row shows five Skeleton tiles at 40 px, the answer box keeps 56 px. Empty/error: talent or BiS file fails to load: answer `Open the leveling guide for your talent order.`, button `Level as {Spec} →`; the row is replaced by `LoadError` with `Talent order did not load.` Disabled: none.

## 6. Hierarchy
Header button is the one primary on the page. The Talents card link is secondary text. The Simulator and Upgrades locked cards are visibly muted (text-secondary token, no gold).

## 7. Performance
No layout shift: header 56 px box, card min-heights equal the loaded states. Talent icons use the preloaded icon sprite the guide chips use; no missing-icon flash (tile reserves 40 px square). Budgets in `web/lighthouserc.json` unchanged.

## 8. Acceptance screenshots (390 and 1440; 2000 for home)
Hunter, Alliance, with no export: levels 1, 9, 10, 15, 19, 20 on `/`. Level 12 with an export of 2 points spent (one to spend). Level 15 with export in Marksmanship (spec follows the points). Level 5 on `/bis` and on `/bis/hunter`. States: loading, load error. Attach the guide landing at `#leveling` after the button click.

# Paperdoll: one character-screen gear panel (experience spec)

Owner, 2026-10-30, on the sim's 4-column grid of flat slot rows ("HEAD · Item 288892 · Not in our item data yet"): "we should have this look like a character screen item window instead of flat rows - that's what users are used to."

This spec defines one shared component, `Paperdoll`, used wherever a full gear set is drawn. The native apps take the same layout later.

Mocks: `design/mocks/2026-10-31-paperdoll/` (`gen.py` renders them from build 1.60.1.70658, the active build in `web/src/data/active-build.json`; item ids, names, icons, qualities and item levels are real; the gains on the character board are illustrative).

| Board | Shows |
|---|---|
| `raid-1280.png`, `raid-1024.png`, `raid-2000.png` | Sim character card, Fury raid set (band 60 raid, Alliance) |
| `raid-390.png`, `raid-360.png` | The same at phone widths |
| `beta-1280.png`, `beta-390.png` | Character page gear: three beta items not in our data, a two-hander, an empty ranged, gains on six slots, one slot already best |
| `empty-1280.png` | A planner build with 3 of 17 slots filled |
| `loading-1280.png`, `loading-390.png` | Items file pending |
| `states.png` | Every slot state, plus the tooltip on hover |

---

## 1. The question and the answer in view

**Question:** "What is this character wearing?"

**Answer in view:** the crest and name, plus the head-to-wrist and hands-to-trinket columns, in one frame.

- At 1280 x 800 the crest, the name and slot rows 1 to 6 on both sides sit above the fold once the card head is in view.
- At 360 x 640 the identity row and slot rows 1 to 4 on both sides sit above the fold once the card head is in view.
- All 17 slots fit in one phone screen-height of scrolling (two columns, not seventeen rows).

## 2. Reference

| Reference | What we borrow |
|---|---|
| The client's character frame (PaperDollFrame) | The slot positions: head to wrist down the left, hands to trinket 2 down the right, weapons at the bottom. Sockets with quality borders. Empty sockets show the client's own slot silhouette (`UI-PaperDoll-Slot-*`). A two-hander leaves its icon greyed in the off hand. |
| Paperdoll-with-names addons (the client frame with item names beside each slot) | Names beside the socket, on the side toward the centre, in their quality colour. |
| Our Discord character sheet, `api/internal/sheet/render.go` | Slot order, the mirrored columns, top alignment when shirt and tabard are absent, and the weapons as main hand left, off hand right, ranged centred one row below. The centre holds the crest, the item level and the source. The site and the PNG show the same picture of a character. |
| Wowhead's tooltip, already our `ItemTooltip` | Hover, focus or tap opens the real item (tenet 2). |
| /account card (site alignment §3.2, §3.3.1) | The card surface and the compact head. |

We do not borrow the 3D model. We never use a Blizzard avatar. The centre is always the ringed HD crest (`ClassCrestRing`).

## 3. Layout

### 3.1 Component and call sites

New shared component: `web/src/components/character/Paperdoll.svelte`, plus a pure model, `web/src/lib/gear/paperdoll.ts`. This is a new shared component, so it gets its own design review (tenet 10).

The component draws the frame only. It renders no card and no head. Each call site wraps it in its own `Card` or existing card.

Props:

- `identity: { name; classSlug; race; spec; level; faction; source }`.
  - `source` is a string the caller already has: `sourcePill()` on the sim, "Addon export, 12 min ago" on the character page, "At pull" on a report.
- `slots: PaperdollSlot[]`, one per slot key. A `PaperdollSlot` is:
  - `{ slot, state: 'item' | 'empty' | 'unknown' | 'twohand-offhand', item?: { id; name; icon; quality; itemLevel }, detail?: Snippet }`
- `build: string`, for icons and the tooltip.
- `classSlug`, for `ItemHover`.
- `loading: boolean` (the items file is pending).
- `failed: boolean` (the items file failed).

Slot keys are `SLOTS` from `lib/planner/types.ts`, plus `shirt` and `tabard` when the data has them. Labels come from `SLOT_LABELS`. That file is the one label source. `gearCopy.slotLabels` ("Ring 1") is deleted in favour of the client's word "Finger 1".

Reused, not invented:

- `ItemHover` is the trigger and tooltip for every known item and every pick.
- `ClassCrestRing` draws the crest.
- The faction mark is the existing asset `web/public/icons/hd/faction/<faction>-logo-64.webp`, drawn at 16 px.
- The `.pill .pill-site` recipe is the source tag.
- `itemIdLabel`, `itemNotInDataPhrase` (`lib/items/unknown-item-copy.ts`) and `points()` (`lib/gear/copy.ts`) supply copy.
- `rarityClassFor` supplies quality colours.

### 3.2 Call sites (every surface that draws a full set)

| # | File | Today | Becomes |
|---|---|---|---|
| 1 | `web/src/components/sim/CharacterStrip.svelte` (live `/sim` via `SimView.svelte`, and `/sim/<id>` via `SavedSim.svelte`) | A 2/4-column grid of flat slot rows, its own copy of the layout | `Paperdoll` inside the card. The card head becomes the compact head: `user` tile, "CHARACTER", with "Change source" `button('secondary','sm')` in the right slot. The crest, name, descriptor and source pill leave the head row and go into the paperdoll's centre. The race picker, level note and talents row stay below the paperdoll, unchanged. |
| 2 | `web/src/components/character/CharacterGearPanel.svelte` (character page `#gear`) | 17 three-line rows: worn, arrow, pick, source, gap | `Paperdoll` with `detail` = the gain line (§6). The title line and its summary stay in the card head ("GEAR", meta "+57.1 points across 6 slots"). The footer's BiS link and its failed-file line stay. The `#gear` anchor and heading focus stay. |
| 3 | `web/src/components/report/GearList.svelte` (report Summary tab, "At pull") | A 1/2-column list with slot label, 18 px icon, name and item level | `Paperdoll` with `detail` = "Item level 66 · Enchanted". The centre source is "At pull". Retail ids keep the unknown state. |
| 4 | `apps/ios/App/Home/GearView.swift`, Android `home/` (later) | Gear list | The same layout, in the native lane, from this spec's §3 to §7 values. |

The following draw slots but are **not** call sites. Each has one job that is not "show a worn set" (one job per page):

- `sim/tools/SlotGrid.svelte` holds candidates per slot.
- `bis/GearRow.astro` / `BandPanelBody.astro` show picks with alternatives (tenet 4 keeps the alternatives visible).
- `sim/tools/NamedSets.svelte` lists set names.
- `guild/GuildReadiness.svelte` shows per-character missing-slot counts.
- `HomeUpgradesPanel` / `UpgradeRow` list upgrades.
- The planner has no gear panel.

### 3.3 Regions (container widths; the component uses a container query, not the viewport)

The component sets `container-type: inline-size`. The frame is `max-width: 960px; margin-inline: auto`, so it is centred in any card. The card head and foot span the full card.

**Frame, container at least 640 px (desktop):**

- Grid columns are `minmax(0,1fr) C minmax(0,1fr)`.
  - `C` = 168 px from 640 to 879.
  - `C` = 240 px from 880.
- Column gap 24 px. Row gap 8 px.
- The left column holds head, neck, shoulder, back, chest, shirt*, tabard*, wrist in grid rows 1 to n, top-aligned.
  - With shirt and tabard absent, wrist sits on row 6, level with Finger 2, exactly as the sheet does. Rows 7 and 8 on the left stay empty.
- The right column holds hands, waist, legs, feet, finger 1, finger 2, trinket 1, trinket 2 in rows 1 to 8.
- The centre spans rows 1 to 8, vertically centred.
- Main hand: row 9, column 1. Off hand: row 9, column 3. Each has an extra 8 px above it (16 px total from row 8).
- Ranged/relic: row 10, column 2.
- A row is 52 px. The frame body is about 584 px tall.

**Frame, container under 640 px (phone, §7):** two columns of `minmax(0,1fr)`, column gap 8 px.

- The identity sits on row 1 across both columns.
- Left slots go in column 1, right slots in column 2, from row 2.
- Main hand: row 10, column 1. Off hand: row 10, column 2. Each has 8 px extra above.
- Ranged: row 11, centred, at one column's width.
- A row is 48 px.

**Slot cell:**

- The cell is one element: a `<button>` (via `ItemHover`) for a known item, a `<div>` otherwise.
- `display: flex; align-items: center`. Padding 4 px. Radius 8 px. Min-height 52 px desktop, 48 px phone.
- Gap from socket to text: 12 px desktop, 8 px phone.
- In the right column the cell is `flex-direction: row-reverse; text-align: right`. The socket is always on the outer edge and the text faces the crest.
- In the bottom row, main hand and ranged follow the left column's direction, and off hand follows the right column's.

**Socket:**

- Size: 44 px desktop, 40 px phone.
- `border: 2px solid` in the quality colour token (`--color-rarity-*`, the game hues, not the `-text` variants).
- `border-radius: var(--radius-pill)` (6 px). Ground `#14171f`.
- `box-shadow: inset 0 0 0 1px rgba(0,0,0,.65), 0 1px 2px rgba(0,0,0,.5)`.
- The icon fills the socket inside the border (`object-fit: cover`) and is served from the build's `icons/<icon>.webp`.
- `width` and `height` attributes are set, so there is no shift.

**Text block (`min-width: 0`):**

- Name: Barlow 600, 14/18 desktop and 13/16 phone, in the quality text class (`rarityClassFor`, which uses `--color-rarity-rare-text` and `--color-rarity-epic-text`).
- The name clamps to 2 lines. It clamps to 1 line when a detail line is present.
- Detail (optional): 12/16 Barlow 400, `--color-muted`, 1 line with ellipsis.

**Centre, desktop, top to bottom:**

1. Crest: `ClassCrestRing` at 96 px (72 px when `C` is 168).
2. 14 px gap.
3. Name: `--font-heading` 800, 26/28 (22/24 at `C` 168), uppercase, `.01em`, class colour. It breaks anywhere rather than overflow.
4. Descriptor: the faction logo at 16 px, then "Human Fury Warrior", 14/18 `--color-text`.
5. "Level 60": 13/18 `--color-muted`.
6. 12 px gap.
7. Micro-label "ITEM LEVEL": 11/14 Barlow 700, `.14em`, `--color-muted`.
8. Figure: `--font-heading` 800, 28/32 (`--text-stat`), `--color-gold`, tabular.
9. When some worn items are unknown: "12 of 15 items", 12 px `--color-muted`.
10. 10 px gap.
11. Source tag: `.pill .pill-site`.

**Centre, phone:**

- A centred row: crest 56, 12 px gap, then a text block.
- The text block holds:
  - the name, 22/24 heading face;
  - the descriptor, 13/18;
  - one line "Level 60  Item level **65.2**" (13 px muted, figure gold 600, plus " 12 of 15 items" when partial);
  - the source tag, 6 px below.
- 16 px of space sits under the row.

**Item level:**

- The figure is the mean item level of the worn items the data knows, to one decimal.
- This is `averageItemLevel` in `api/internal/discordbot/cmd_char.go`, so the site, the Discord card and the sheet agree.
- The model computes it in `paperdoll.ts`.
- It is hidden, with its label, when no worn item is known.

**Note under the frame:**

- It is shown only when at least one slot is unknown, or when the items file failed.
- Margin 12 px top. 13/18 `--color-muted`. Left-aligned to the frame.

### 3.4 Tokens used (nothing else)

- Colour:
  - `--color-rarity-*` and the two `-text` variants;
  - `--color-class-*`;
  - `--color-gold`, `--color-muted`, `--color-text`, `--color-strong`;
  - `--color-line`, `--color-line-soft`;
  - `--row-hover`.
- Type: `--font-body`, `--font-heading`, `--text-stat`, `--text-meta`, `--text-micro`.
- Radius: `--radius-pill` for the socket, `--radius-control` for the cell.
- Card: `card-surface` and the compact head (site alignment §3.3.1).
- Body padding: 20 px at 768 and above, 16 px below.

## 4. States

Every state is drawn in `states.png`.

| Element | State | Exact treatment |
|---|---|---|
| Known slot | default | As §3.3. Cursor pointer. |
| | hover | Cell ground `--row-hover`. Icon `filter: brightness(1.12)`. `ItemHover` opens the tooltip. |
| | focus-visible | `outline: 2px solid var(--color-gold); outline-offset: 2px`, radius 8. The tooltip opens on focus. Esc closes it. |
| | active / tap | Ground `rgba(229,185,85,.10)`. A tap toggles the tooltip. A tap elsewhere closes it. Only one tooltip is open on the page (`ItemHover`'s rule). |
| | disabled | None. A slot is never disabled. |
| Pick inside a detail line (character page, desktop) | hover / focus | Its own `ItemHover` (inline, underline on hover) opens the pick's tooltip, which includes its source. Below 640 the pick is plain text inside the cell's single target (§7). |
| Empty slot | default | Socket `1px dashed var(--color-line)`, ground `#0b0e14`, padding 2. The client's silhouette for that slot at 45% opacity. The name line is the slot label ("Ranged") in `--color-muted` 600. It is not focusable and has no hover. On the character page the gain line can still sit under it. |
| Not in our item data yet | default | Socket `1px solid var(--color-line)`, padding 1. The game's `inv_misc_questionmark` at 55% opacity. Name "Item 288892", 400 tabular, `--color-muted`. No detail. Not focusable. `aria-label="Wrist: Item 288892, not in our item data yet"`. The phrase appears once, in the note under the frame, never on each line. |
| Off hand under a two-hander | default | Dashed socket as empty. Inside it, the main hand's icon `grayscale(1)` at 30% opacity (the client's greyed off hand). Name "Two-hand in main hand", `--color-muted` 600. Not focusable. The two-hand flag is the build's `two_hand`. While the items file is pending, the off hand shows loading. |
| Any slot | loading | The socket is a shimmer block (the report-skeleton gradient `#141925` to `#1c2230`). The name is a 12 px bar at 60% width, radius 4. Rows keep their height. The centre draws the crest, name, descriptor and source at once. The item-level figure is a 64 x 28 shimmer. Nothing moves when data arrives. |
| Any slot | error (items file failed) | The socket is plain (`1px solid var(--color-line)`, no icon). The name is in `--color-text`, because no quality is claimed. It is still an `ItemHover` trigger, since the tooltip fetches its own model. The item-level figure is hidden. The note under the frame shows. |
| Whole set | empty (no gear known) | The paperdoll is not drawn. The call site's existing state stays: `ExportNotRead` ("Your gear has not been read yet."), and `simCopy.savedNoGear` on a saved sim. Seventeen empty sockets never stand in for "not read". |

**Keyboard order:**

1. Left column, top to bottom.
2. Right column, top to bottom.
3. Main hand, off hand, ranged.
4. Then each slot's pick, if any, after its own item.

DOM order is the same order. Screen readers hear the list as `role="list"` named "Gear by slot", and each cell as a `listitem`.

**Touch:** every target is at least 44 px. The phone cell is 48 px, full column width.

## 5. Hierarchy

- **Frame:** the item names are the primary element. Quality colour is the one level of emphasis. Nothing else on a slot line is coloured, except the gold gain.
- **Centre:** the name is primary. The item-level figure is the only gold figure. Everything else is muted.
- **Secondary is visibly secondary:**
  - empty, unknown and two-hand slots are muted grey with dashed or dimmed sockets;
  - detail lines are 12 px muted;
  - the note is 13 px muted, below the frame.
- **Whitespace:** 24 px between the columns and the centre; 8 px between rows; 16 px above the weapons. The left column's rows 7 and 8 are left empty, as in the client and the sheet. They are never filled to look even.

## 6. Copy (verbatim)

| Where | Text |
|---|---|
| Sim card head title | `CHARACTER` |
| Character card head title | `GEAR` (meta: `gearCopy.summary`, e.g. "+57.1 points across 6 slots") |
| Empty slot name | the slot label: `Head`, `Neck`, `Shoulder`, `Back`, `Chest`, `Shirt`, `Tabard`, `Wrist`, `Hands`, `Waist`, `Legs`, `Feet`, `Finger 1`, `Finger 2`, `Trinket 1`, `Trinket 2`, `Main hand`, `Off hand`, `Ranged` (`Relic` for paladin, shaman and druid) |
| Unknown slot name | `Item 288892` (`itemIdLabel`) |
| Note, one unknown | `1 item is not in our item data yet. It is new in the beta, and it gets a name and icon when the client lists it.` |
| Note, several | `3 items are not in our item data yet. They are new in the beta, and they get names and icons when the client lists them.` |
| Note, items file failed | `Item icons and colours did not load. Reload the page to try again.` |
| Off hand under a two-hander | `Two-hand in main hand` |
| Centre micro-label | `Item level` (shown uppercase by CSS) |
| Centre partial count | `12 of 15 items` |
| Centre level | `Level 60` |
| Gain line (character page) | `+18.4 pts · Lionheart Helm` (gain gold 600, pick in its quality colour) |
| Gain line, slot already best | `Your best for level 60`; for bands under 60, `Your best for level 50 to 59` |
| Gain line, no pick for the band | none (no line) |
| Gain line, pick held while the two-hand answer is pending | none (no line), as `noteUnknown` today |
| Report detail | `Item level 66` and ` · Enchanted` when enchanted |
| List name (screen readers) | `Gear by slot` |
| Known slot accessible name | `Head: Lionheart Helm` |

## 7. Phone and desktop

- **Phone, container under 640** (`raid-390.png`, `raid-360.png`, `beta-390.png`): this spec chooses **two mirrored columns with the crest on top**, not a single compact list.
  - Why:
    - it keeps the frame's mirror, so a player finds trinkets bottom-right as in game;
    - it is half the height (11 rows of 48 px against 17 of 52 px);
    - it matches the Discord sheet and the coming native screens.
  - The cost is truncation. A name gets 2 lines of 13 px (about 15 characters a line at 360). With a detail line, the name gets 1 line. The full name is one tap away in the tooltip.
  - On the character page the gain line truncates the pick. The pick is not a separate tap target under 640, because a 16 px target fails touch parity. The full pick and source stay one tap down in the "BiS" link the card foot already has.
- **1024** (`raid-1024.png`): the desktop frame, `C` 240, inside the card. Names rarely clamp.
- **1280** (`raid-1280.png`, `beta-1280.png`): as 1024. The frame reaches its 960 px cap and is centred.
- **~2000** (`raid-2000.png`): the page column stops at `--page-max` 1344. The card fills it. The 960 px frame stays centred in the card, so the crest is centred on the screen. The frame never stretches its columns past 960.
- **Never collapses, at any width:**
  - the mirror (sockets on the outer edges);
  - the crest;
  - the item names;
  - the unknown and empty states;
  - the note.
- **Character page in the 440 px rail layout:** the gear card sits in the main column. At 1280 the main column is about 712 px, so `C` is 168 and each side is about 240 px, the sheet's own cell width.

## 8. Performance and polish

- No layout shift on hydration:
  - every socket, bar and figure box has fixed dimensions before data arrives;
  - CLS on `/sim`, `/sim/<id>` and the character page stays at or under the `web/lighthouserc.json` budget;
  - `/sim`'s CLS-tuned skeleton (`lib/sim/skeleton.ts`) is re-measured for the new frame height.
- No missing-icon flash:
  - names are held behind the loading bar until the items file answers (as `CharacterGearPanel` does today);
  - icons use `loading="lazy" decoding="async"` with width and height set;
  - a failed image keeps the plain socket and never shows a broken-image glyph.
- **Silhouettes:** the 18 client textures (`UI-PaperDoll-Slot-Head` to `-Relic`, back using `-Chest` as the client does) ship as faithful HD reproductions in `web/public/icons/hd/slots/<slot>.webp`.
  - The memory rule applies: game icons exactly, high-def, never reinterpreted.
  - The mock uses the Wowhead `inventoryslot_*` renders in `design/mocks/2026-10-31-paperdoll/slots/` only as stand-ins.
- No new font: Barlow and Barlow Condensed only. The `Paperdoll` island adds at most 4 KB gzip on top of `ItemHover`.
- Lighthouse budgets in `web/lighthouserc.json` stay unchanged and green.

## 9. Acceptance screenshots (the implementer attaches all)

Each at 360, 390, 1024, 1280 and 2000 unless noted, captured from the built site:

1. `/sim` with an addon-export Fury raid set: default.
2. `/sim/<id>` (a saved sim).
3. The character page `#gear` with gains, one slot already best, a two-hander, an empty slot and at least one unknown beta item, plus the note.
4. A report Summary tab "At pull".
5. Loading, with the items file throttled: 390 and 1280.
6. Items file failed (blocked): 390 and 1280.
7. Hover with the tooltip open (1280), focus-visible via Tab (1280), and the tooltip after a tap (390).
8. A same-scale side-by-side of the 1280 capture with `raid-1280.png`, and of the character page with `beta-1280.png`.
9. The Discord sheet PNG for the same character beside the 1280 capture: slot order identical.

**Checklist the review ticks:**

- [ ] One `Paperdoll` component; call sites 1 to 3 use it; `CharacterStrip`'s copied grid is deleted.
- [ ] Slot positions match `api/internal/sheet` (left 6 or 8 top-aligned, right 8, MH/OH row, ranged centred below).
- [ ] Sockets 44/40 px, 2 px game-hue quality border; names in the quality text colour, toward the centre.
- [ ] Crest is `ClassCrestRing`; no square icon, avatar or model.
- [ ] Item level equals the Discord card's figure for the same export.
- [ ] Empty, unknown, two-hand, loading and failed states as §4; the unknown phrase appears once.
- [ ] Every known item and pick opens the shared tooltip on hover, focus and tap; an e2e test covers each call site (tenet 9).
- [ ] Tab order as §4; 44 px targets; screen-reader names as §6.
- [ ] The frame is centred and capped at 960 at 2000.
- [ ] CLS and Lighthouse within budget; no icon or name flash.

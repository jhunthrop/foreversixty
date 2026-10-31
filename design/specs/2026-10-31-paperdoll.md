# Paperdoll: one character-screen gear panel (experience spec)

Owner, 2026-10-30, on the sim's 4-column grid of flat slot rows ("HEAD · Item 288892 · Not in our item data yet"): "we should have this look like a character screen item window instead of flat rows - that's what users are used to."

This spec defines one shared component, `Paperdoll`, used wherever a full gear set is drawn. The native apps take the same layout later.

Revision: fix round 1, 2026-10-31. It answers `design/reviews/2026-10-31-paperdoll/ux-verdict.md` and `player-verdict.md` (both HOLD). §10 maps every finding to its fix.

## 0. Mocks

The mocks live in `design/mocks/2026-10-31-paperdoll/`. `gen.py` renders them from build 1.60.1.70658, the active build in `web/src/data/active-build.json`. Every board is a 2x capture of CSS pixels, so all boards share one scale.

**Data rules for the mocks:**

- Every item, icon, quality, item level, required level and source line comes from the build's `items/warrior.json` and `loot.json`.
- **Launch-legal.** Each item has an ungated dungeon, crafted or world-drop source. The following are excluded, per the 2026-10-30 BiS phase audit:
  - Balzaphon drops;
  - the 1.10 summoned-boss items (22300 to 22450);
  - the late recipes (Titanic Leggings, Stormshroud Gloves);
  - every quest, reputation, PvP, Darkmoon Faire and Forever-new vendor source.
- **Gains are a real comparison.** A gain is the pick's score minus the worn item's score. A score is the item's stats times the spec's scale factors in `bis/warrior-arms.json`, with strength = 1.00. A percent stat is converted with the weight's own rating factor. Weapons carry no stat weight in that file, so the mock gives them no gain.

**Boards:**

| Board | Shows |
|---|---|
| `raid-1280.png`, `raid-1024.png`, `raid-2000.png`, `raid-390.png`, `raid-360.png` | Sim character card: a full launch-legal level-60 Fury set |
| `beta-1280.png`, `beta-390.png`, `beta-360.png` | Character page gear, with: three worn items not in our data (288892, 288887, 288893); Arcanite Champion (crafted, two-handed) with an empty off hand; an empty ranged slot; real gains on nine slots; two slots already best in slot; picks for unknown worn items, marked "Not scored" |
| `longest-1024.png`, `longest-390.png`, `longest-360.png` | Worst case: the longest launch-legal names in every slot. The longest is "Electromagnetic Gigaflux Reactivator" (36 characters). |
| `empty-1280.png` | A planner build with 3 of 17 slots filled |
| `loading-1280.png`, `loading-390.png` | Items file pending |
| `states-1280.png` | Every slot state, desktop cell |
| `tip-left-1280.png`, `tip-right-1280.png`, `tip-flip-1280.png` | Tooltip placement for the left column, the right column, and a cell near the viewport bottom |
| `tip-phone-390.png`, `tip-phoneflip-390.png`, `tip-phoneup-390.png` | Phone tooltip below a cell, flipped above near the bottom, and opened from an upgrade row |

## 1. The question and the answer in view

**Question:** "What is this character wearing, and what do I change next?" The second half applies only on the character page.

**Answer in view:**

- At 1280 x 800 the following sit above the fold once the card head is in view: the crest, the name, and slot rows 1 to 6 on both sides. On the character page the gain is on each of those lines.
- At 360 x 640 the identity row and the first two slot pairs sit above the fold. On the character page their upgrade rows show too.

## 2. Reference

| Reference | What we borrow |
|---|---|
| The Classic client's character frame (`Blizzard_CharacterFrame/Vanilla/PaperDollFrame.lua`, Gethe/wow-ui-source `classic_era` at 1.15.9 (70576)) | The slot positions and the quality-bordered sockets. An empty slot draws the client's own slot background texture (`PaperDollItemSlotButton_Update`, lines 839 to 852: `textureName = self.backgroundTextureName`). **The client has no two-hand case.** With a two-hander equipped, the off hand is simply empty and shows `UI-PaperDoll-Slot-SecondaryHand`. Nothing greys the two-hander into it. We draw exactly that, and add the words "Two-hand in main hand". |
| Paperdoll-with-names addons | Names beside the socket, on the side toward the centre, in their quality colour. |
| Our Discord character sheet, `api/internal/sheet/render.go` | Slot order, the mirrored columns, top alignment without shirt and tabard, and the weapons: main hand left, off hand right, ranged centred one row below. This is deliberate parity. The site and the PNG show the same picture of a character. The client's own bottom-centre weapon row is not used, because the two would then disagree. |
| Our shared `ItemTooltip` (`web/src/components/ItemTooltip.svelte`, model `lib/items/tooltip.ts`) | Hover, focus or tap opens the real item (tenet 2). The tooltip boards draw exactly what that model produces. |
| /account card (site alignment §3.2, §3.3.1) | The card surface and the compact head. |

We do not borrow the 3D model. We never use a Blizzard avatar. The centre is always the ringed HD crest (`ClassCrestRing`).

## 3. Layout

### 3.1 Component

New shared component: `web/src/components/character/Paperdoll.svelte`, plus a pure model, `web/src/lib/gear/paperdoll.ts`. This is a new shared component, so it gets its own design review (tenet 10).

The component draws the frame only. It renders no card and no head. Each call site wraps it in its own `Card`.

**Props:**

- `identity: { name; classSlug; race; spec; level; faction; source }`.
- `slots: PaperdollSlot[]`, one per slot key, where a `PaperdollSlot` is:
  - `slot`;
  - `state: 'item' | 'empty' | 'unknown'`;
  - `item?: { id; name; icon; quality; itemLevel; twoHand }`;
  - `upgrade?: { kind: 'gain'; points: number; pick: ItemRef } | { kind: 'best' } | { kind: 'unscored'; pick: ItemRef }`;
  - `detail?: string` (the report's "Item level 66 · Enchanted").
- `build`, `classSlug`, `loading`, `failed`.

The model derives the off hand's two-hand state from `item.twoHand` on the main hand.

**Slot keys and labels:**

- Slot keys are `SLOTS` from `lib/planner/types.ts`, plus `shirt` and `tabard` when the data has them.
- Labels come from `SLOT_LABELS`, the one label source. It uses the client's words: `FINGER0SLOT` is "Finger", so "Finger 1" and "Finger 2".
- `gearCopy.slotLabels` ("Ring 1") is deleted. These files say "Ring 1" or "Ring 2" today, and the build lane moves them, tests included:
  - `web/src/lib/gear/copy.ts`, `web/src/lib/gear/panel.ts`, `web/src/lib/gear/panel.test.ts`
  - `web/tests/e2e/character-gear.spec.ts`
  - `api/internal/sheet/copy.go`, `sheet.go`, `sheet_test.go`, `testdata/alts.json`
  - `api/internal/discordbot/commands.go`, `item_embed.go`, `cmd_bis_test.go`, `gear_panel_rules_test.go`
  - `web/tests/captures/discord-bot/payloads.json`, `payload-dump.go.txt`
  - `apps/shared/copy/en.json`, `apps/ios/App/Generated/Copy.swift`, `apps/ios/App/Localizable.xcstrings`, `apps/android/app/src/main/res/values/strings.xml`
  - `api/internal/bnetapi/testdata/era-kiloz/equipment.json` is Blizzard's own payload. It stays as is.

**Reused, not invented:**

- `ItemHover` is the trigger and tooltip for every item and pick. It gains an `align: 'start' | 'end'` option (§4.2).
- `ClassCrestRing` draws the crest.
- The faction mark is `web/public/icons/hd/faction/<faction>-logo-64.webp`, drawn at 16 px.
- `.pill .pill-site` is the source tag.
- `itemIdLabel` and `itemNotInDataPhrase` supply unknown-item copy.
- `points()` and `gearCopy.summary` supply gain copy.
- `rarityClassFor` supplies quality colours.

### 3.2 Call sites (audit, 2026-10-31)

This audit greps components and pages for `SLOTS`, `gear`, `worn`, slot labels and OG image routes.

| # | Surface | Draws a full set today? | Becomes |
|---|---|---|---|
| 1 | `web/src/components/sim/CharacterStrip.svelte`: live `/sim` (`SimView.svelte`) and saved `/sim/<id>` (`SavedSim.svelte`) | Yes, as flat rows (its own copy of the layout) | `Paperdoll`. The card head becomes the compact head: `user` tile (the mocks now draw the `user` glyph), "CHARACTER", and "Change source" `button('secondary','sm')`. The crest, name, descriptor and source move into the centre. The race picker, level note and talents row stay below. |
| 2 | `web/src/components/character/CharacterGearPanel.svelte` (character page `#gear`) | Yes, as 17 three-line rows | `Paperdoll` with `upgrade` per slot. The head meta keeps `gearCopy.summary` and turns `--color-gold` 600. The foot's BiS link and failed-file line stay. |
| 3 | `web/src/components/report/GearList.svelte` (report Summary, "At pull") | Yes, as a list | `Paperdoll` with `detail`. The centre source reads "At pull". |
| 4 | `apps/ios/App/Home/GearView.swift`, Android `home/` | Gear list | The same layout in the native lane, later. |
| 5 | Discord character sheet, `api/internal/sheet` (the /char and /bis share PNG) | Yes, its own Go renderer | It is not the Svelte component, but it must keep parity. Follow-ups for the sheet: the item-level rule of §3.5 (today `averageItemLevel` in `discordbot/cmd_char.go` averages known items only and must take the same rule), and drawing the off hand under a two-hander as empty with the words. Today `placements(twoHand)` drops the off hand. |

**Checked and not call sites:**

| Surface | Why not |
|---|---|
| Home character strip: `CurrentCharacterBar.svelte`, `character/HomeHeroCards.svelte`, `character/HomeUpgradesPanel.svelte` | Identity, the score cards and the top upgrades. No slot set is drawn. |
| Guild member card | No member hover card exists. `guild/GuildRosterTable.svelte` shows a gear-gap figure per member. `guild/GuildReadiness.svelte` shows missing-slot counts ("Main hand, Chest +2 more"). Neither draws a set. |
| Saved-sim share image | `/sim/[id].astro` sets no image of its own, so it inherits `Base.astro`'s default `og:image`. `web/src/pages/og/[...path].png.ts` draws title cards for static pages only. No set is drawn. |
| `sim/tools/SlotGrid.svelte`, `bis/GearRow.astro` / `BandPanelBody.astro`, `sim/tools/NamedSets.svelte` | Candidates, picks with alternatives, set names: one job each. |

### 3.3 Regions

The component sets `container-type: inline-size`. All breakpoints below are container widths.

**The stage:** an inset object behind the frame. The client frames its model the same way.

- Ground `#0b0e14`, 1 px `--color-line-soft`, radius 12 px.
- `max-width: 960px; margin-inline: auto`.
- Padding 20 px 24 px on desktop, 16 px 12 px on phone.
- It is flat, with no new asset.

**Desktop frame, container at least 640:**

- Grid columns are `minmax(0,1fr) C minmax(0,1fr)`.
  - `C` = 168 px from 640 to 879.
  - `C` = 240 px from 880.
- Column gap 24 px. Row gap 8 px.
- The left column holds head, neck, shoulder, back, chest, shirt\*, tabard\* and wrist in rows 1 to n, top-aligned. With shirt and tabard absent, rows 7 and 8 on the left stay empty, as on the sheet. This is deliberate (§5).
- The right column holds hands to trinket 2 in rows 1 to 8.
- The centre spans rows 1 to 8.
- Main hand: row 9, column 1. Off hand: row 9, column 3. Each has 8 px extra above it.
- Ranged/relic: row 10, column 2.
- Row height: 52 px, or 64 px when the frame carries upgrade or detail lines (the character page and the report).

**Phone frame, container under 640:**

- Two columns of `minmax(0,1fr)`. Column gap 8 px. Row gap 4 px.
- The identity sits on row 1 across both columns.
- Then 8 slot pairs: left slot and right slot on one row.
- **Upgrade rows** (§3.4) follow each pair: the left slot's first, then the right slot's.
- Then the main hand / off hand pair and its upgrade rows.
- Then ranged, centred at one column's width, and its upgrade row.
- Slot rows are 56 px.

**Slot cell:**

- One element: a `<button>` (via `ItemHover`) for a known item, a `<div>` otherwise.
- Flex, centred. Padding 4 px. Radius 8 px.
- Gap from socket to text: 12 px desktop, 8 px phone.
- In the right column the cell is `row-reverse` and `text-align: right`. The socket is always on the outer edge.

**Socket:**

- Size: 44 px desktop, 40 px phone, 36 px at containers under 340.
- `border: 2px solid` the game hue (`--color-rarity-*`). Radius 6 px. Ground `#14171f`.
- `box-shadow: inset 0 0 0 1px rgba(0,0,0,.65), 0 1px 2px rgba(0,0,0,.5)`.
- The icon fills the socket.

**Name: clamp rules.**

| Container | Type | Clamp |
|---|---|---|
| at least 640 | Barlow 600, 14/18 | 2 lines |
| 340 to 639 | 13/16 | 3 lines |
| under 340 (a 360 px phone: the 296 px container) | 12/15, socket 36 | 3 lines |

- Every clamp is `display: -webkit-box; -webkit-line-clamp: N; overflow: hidden; text-overflow: ellipsis; hyphens: auto; overflow-wrap: break-word`. The page sets `lang="en"`, so hyphenation works.
- A cut name always ends in "…". A single word longer than the line is hyphenated, never cut mid-word without a mark.
- Measured on the boards: the longest launch-legal names (36 characters) fit untruncated at 1024 (2 lines at most), 390 and 360 (3 lines at most), including "Electromagnetic Gigaflux Reactivator" and "Protector's Silvered Chain Leggings".
- **No two slots can collapse to the same visible string**, for three reasons:
  1. Names are never cut at today's lengths.
  2. The clamp leaves at least 40 visible characters.
  3. An unknown item always carries its slot word ("Waist · Item 288887"), and every phone upgrade row leads with its slot word.
- The build adds an e2e assertion on the worst-case fixture: no `.name` has `scrollHeight > clientHeight`.

**Detail and upgrade line, desktop:**

- A second line in the cell: 13/18, `display: flex; gap: 6px`. In the right column it is `justify-content: flex-end`.
- **Gain:** a 16 px pick icon with a 1 px quality border, then the pick name (quality colour, 1-line ellipsis, `min-width: 0`, its own `ItemHover`), then the gain "+18.4" in `--color-gold` 600 tabular. The gain never shrinks.
- **Best:** "Best in slot" in `--color-muted`. There is never a gain or a pick.
- **Unscored** (the worn item is unknown, so no comparison is possible): the pick, then "Not scored" in muted.

**Upgrade row, phone only:**

- A full-width `<button>` (the pick's `ItemHover`) under its slot's pair.
- Min-height 44 px. Padding 0 8 px. Radius 8 px. Ground `rgba(229,185,85,.035)`.
- Contents, left to right:
  1. the slot word, 12/16 600 muted, 64 px wide (56 px under 340);
  2. the pick's 24 px socket, 1 px quality border;
  3. the pick name, 13/16 600 in quality colour, 2-line clamp with ellipsis;
  4. the gain ("+18.4", gold 600) or "Not scored" (muted).
- A slot already at its best gets no upgrade row on phone.
- The desktop detail line is not drawn under 640, and the upgrade row is not drawn at 640 and above. Both render, and CSS hides one with `display: none`, so the reading order stays single.

**Centre:**

- Desktop:
  1. crest 96 (72 at `C` 168);
  2. 14 px gap;
  3. name, heading face 800, 26/28 (22/24), class colour;
  4. descriptor, 16 px faction mark plus "Human Fury Warrior", 14/18;
  5. "Level 60", 13 muted;
  6. item level (§3.5): micro-label "ITEM LEVEL" 11/14 700 `.14em` muted, then the figure in the heading face, 28/32, gold;
  7. the source pill.
- Phone: crest 56 with a text block beside it:
  - the name, 22/24;
  - the descriptor;
  - "Level 60  Item level **62.0**" on one line;
  - the pill.

**Note under the stage:**

- It shows only when a slot is unknown or the items file failed.
- 12 px above it. A 1 px `--color-line-soft` top rule and 12 px padding, so it reads as part of the panel, not the card foot. 13/18 muted.

### 3.4 Upgrades: when each kind shows

| Worn | Pick | Desktop line 2 | Phone upgrade row |
|---|---|---|---|
| known item | a different item scoring higher | pick + gain | yes |
| known item | the same item (the pick is worn; rings and trinkets matched either way round, as `wearsListedBest`) | "Best in slot" | none |
| known item | a different item scoring the same or lower | none | none |
| empty | any | pick + gain (the pick's whole score) | yes |
| unknown | any | pick + "Not scored" | yes, "Not scored" |
| any | none for the band | none | none |

The gain figure is the API's `gap` (`CharacterGearSlot.gap`). The model never shows a gain when `pick.id` equals the worn id. A unit test pins this, fixing the player's blocker 1.

### 3.5 Item level (ruling, 2026-10-31)

- **Counted slots:** every equipped slot except shirt and tabard.
  - An empty slot is not equipped, so it is not counted.
  - The empty off hand under a two-hander is not counted.
- **Shown only when every counted slot is known.** If any equipped item is not in our data, the label and the figure are both absent. There is no partial mean, no "12 of 15" and no dash.
- **The figure** is the mean item level of the counted items, to one decimal.
- **Why:** the unknown items are the newest and likely the highest level, so a partial mean is wrong in exactly the case that matters (player blocker 2).
- The note under the stage says so once: "Item level shows once every item is known."
- The Discord sheet adopts the same rule (§3.2, row 5).

### 3.6 Tokens used (nothing else)

- Colour:
  - `--color-rarity-*` and the `-text` variants;
  - `--color-class-*`;
  - `--color-gold`, `--color-muted`, `--color-text`, `--color-strong`;
  - `--color-line`, `--color-line-soft`;
  - `--row-hover`.
- Type: `--font-body`, `--font-heading`, `--text-stat`, `--text-meta`, `--text-micro`.
- Radius: `--radius-pill` for sockets, `--radius-control` for cells, `--radius-panel` for the stage.
- Card: `card-surface` and the compact head.

## 4. States

### 4.1 Every element

Every row below is drawn in `states-1280.png`.

| Element | State | Exact treatment |
|---|---|---|
| Known slot | default | As §3.3. |
| | hover | `--row-hover` ground on the 8 px cell. Icon `brightness(1.12)`. The tooltip opens. |
| | focus-visible | 2 px `--color-gold` outline, offset 2 px. The tooltip opens on focus. Esc closes it. |
| | active / tap | `rgba(229,185,85,.10)` ground. A tap toggles the tooltip. A tap elsewhere closes it. One tooltip at a time. |
| | disabled | None. A slot is never disabled. |
| Pick (desktop line 2, phone upgrade row) | hover / focus / tap | Its own `ItemHover`, opening the pick's tooltip. The phone upgrade row is a 44 px button with the same hover, focus and active treatments as a cell. |
| Empty slot | default | `1px dashed var(--color-line)` socket, ground `#090b10`. The client's slot texture at 45%. The slot word in muted 600. Not focusable. |
| Not in our item data yet | default | `1px solid var(--color-line)` socket, padding 1. The game's `inv_misc_questionmark` at 40% opacity and `saturate(.5)`. Name "Waist · Item 288887", 400 tabular, muted. Not focusable. `aria-label="Waist: Item 288887, not in our item data yet"`. |
| Off hand under a two-hander | default | The empty-slot treatment with the client's off-hand texture (§2). Words "Two-hand in main hand", muted 600. Not focusable. |
| Any slot | loading | Shimmer socket, and a 12 px bar at 60% width. Rows keep their height. The centre draws the identity at once. The item-level box is reserved as a 64 x 28 shimmer. If the loaded set turns out to have an unknown item, the box collapses, and only the centre's vertical centring moves. No slot moves. |
| Any slot | error (items file failed) | Plain socket. Name in `--color-text`, because no quality is claimed. Still an `ItemHover` trigger. No item level. The note shows. |
| Whole set | not read | No paperdoll. `ExportNotRead` or `simCopy.savedNoGear` stays, as today. |

**Keyboard and touch:**

- Keyboard order:
  1. the left column, top to bottom;
  2. the right column;
  3. main hand, off hand, ranged.
- On desktop each cell's pick follows its cell. On phone, the DOM order is the visual order: pair, then its upgrade rows.
- The list is `role="list"`, "Gear by slot".
- Every target is at least 44 px.
- The 8 px gap between phone columns belongs to no target: a tap in it hits nothing. Each cell's box ends at its column edge.

### 4.2 Tooltip placement

The rule extends `lib/items/tooltip-position.ts`. It is shown on the `tip-*` boards.

1. **Vertical.** The panel opens 8 px below the cell. It flips 8 px above when it does not fit below. This is the existing `tooltipVerticalPlacementFor`.
2. **Horizontal, desktop.**
   - **Left column and the ranged cell:** the panel's left edge meets the cell's left edge (`align: 'start'`, today's default).
   - **Right column:** the panel's right edge meets the cell's right edge (`align: 'end'`, new).
   - Both are then kept 8 px inside the viewport by the existing `tooltipShiftPx`, mirrored for `end`.
   - **Why not the client's rule?** The client anchors every slot's tooltip to the right (`GameTooltip:SetOwner(self, "ANCHOR_RIGHT")`, `PaperDollFrame.lua` line 870). Here that would cover the right column and the crest.
3. **Width.** `clamp(280px, 100vw - 32px, 360px)`, unchanged.
4. **Never covers the crest at 880 and above.** A 360 px panel reaches 20 px into the centre column (column 316 + gap 24 = 340), and the crest starts 72 px in.
   - Between 640 and 879 a panel can cover part of the crest while it is open. This is accepted, because it is transient and the crest repeats nothing a player needs while reading an item.
5. **Phone.** The panel is the viewport width minus 32 px, so it spans 16 px gutters whichever column was tapped. It opens below the tapped cell, or flips above near the bottom.
6. **The tooltip's content** is `ItemTooltip`'s, unchanged. Every number comes from the item record:
   - item level and required level are `item_level` and `required_level`;
   - Lionheart Helm is Item Level 61 and Requires Level 56 in the build, which matches the client's tooltip for 12640.
   - A crafted source reads "Crafted: Blacksmithing" (`craftedSourceLabel`).

**On "Requires Blacksmithing (300)":** the reviewer asked for this line. It is not on Lionheart Helm in the client. That line appears only on the recipe (Plans: Lionheart Helm), because wearing the helm needs no profession. Printing it on the helm would be false (tenet 8), so the tooltip keeps "Crafted: Blacksmithing".

**Found in the shared tooltip.** These are outside this lane. They are raised for the tooltip owner, and the boards show them as shipped:

- A drop chance under 0.5% prints "(0%)". `percentLabel` in `tooltip.ts` uses `Math.round`. It should print "(<1%)", as `gear/panel.ts` already does. The mock draws "<1%".
- An item whose stat is stored as a percent (`percent_stats`, e.g. Truestrike Shoulders `hit: 2`) prints "+2 Hit". The client says "Equip: Improves your chance to hit by 2%."

## 5. Hierarchy

- **Frame:** item names are primary. Quality colour is the one level of emphasis.
- **Character page:** the gain is the second thing seen. That means a gold head meta, a 13 px 600 gold gain on each line, and gold gains on the phone upgrade rows.
- **Centre:** the name is primary. The item-level figure is the only gold figure there.
- **Secondary is visibly secondary:**
  - empty, unknown and two-hand slots are muted, with dashed or dimmed sockets;
  - "Best in slot" and "Not scored" are muted;
  - the note sits in its own ruled strip.
- **Whitespace:** 24 px between the columns and the centre. 8 px between desktop rows. 16 px above the weapons.
  - The left column's rows 7 and 8 stay empty, as in the client without shirt and tabard, and as on the sheet. At 1024 that space sits under wrist, left of the trinkets (`raid-1024.png`). It is accepted, because filling it would break the slot positions players know.

## 6. Copy (verbatim)

| Where | Text |
|---|---|
| Sim card head | `CHARACTER` |
| Character card head | `GEAR`; meta `gearCopy.summary`, e.g. "+148.8 points across 9 slots" (gold) |
| Empty slot | the slot word: `Head`, `Neck`, `Shoulder`, `Back`, `Chest`, `Shirt`, `Tabard`, `Wrist`, `Hands`, `Waist`, `Legs`, `Feet`, `Finger 1`, `Finger 2`, `Trinket 1`, `Trinket 2`, `Main hand`, `Off hand`, `Ranged` (`Relic` for paladin, shaman and druid) |
| Unknown slot | `Waist · Item 288887` |
| Off hand under a two-hander | `Two-hand in main hand` |
| Gain (desktop line 2, phone row) | the pick's name, then `+18.4` |
| Already the best | `Best in slot` |
| Unknown worn item with a pick | the pick's name, then `Not scored` |
| Note, one unknown | `1 item is not in our item data yet. It is new in the beta, and it gets a name and icon when the client lists it. Item level shows once every item is known.` |
| Note, several | `3 items are not in our item data yet. They are new in the beta, and they get names and icons when the client lists them. Item level shows once every item is known.` |
| Note, items file failed | `Item icons and colours did not load. Reload the page to try again.` |
| Centre | `Level 60`, `Item level` (uppercase on desktop by CSS) |
| Report detail | `Item level 66`, plus ` · Enchanted` |
| Accessible names | list `Gear by slot`; cell `Head: Lionheart Helm`; desktop pick and phone row `Head upgrade: Lionheart Helm, 60 points to gain` (`gapLabel`) or `Chest upgrade: Dawn Armor, not scored` |

**On the unit:** the gain is a bare "+18.4" on each line. The unit is named once, in the head meta ("points"). This follows the stat-weights convention of one unit stated once.

## 7. Phone and desktop

- **Phone, container under 640:** two mirrored columns with the crest on top.
  - Why: it keeps the frame's mirror and roughly halves the height of a 17-row list. It matches the sheet and the coming native screens.
  - On the character page the upgrade rows make the pick readable and tappable at full width. The worn names keep the columns.
  - 360 (`raid-360`, `beta-360`, `longest-360`) uses the under-340 rule.
- **1024:** the desktop frame, `C` 240. Longest names take 2 lines (`longest-1024.png`).
- **1280:** the same frame, capped at 960 and centred.
- **About 2000:** the page column stops at `--page-max` 1344. The card fills it, and the 960 px stage is centred, so the crest is at screen centre. `raid-2000.png` is at the same 2x scale as every other board.
- **Never collapses, at any width:**
  - the mirror;
  - the crest;
  - item names (clamped, never hidden);
  - the gain;
  - the unknown and empty states;
  - the note.

## 8. Performance and polish

- **No layout shift on hydration:**
  - sockets, bars and the item-level box have fixed sizes;
  - rows have fixed heights (52/64 desktop, 56 phone, 44 upgrade rows);
  - CLS stays within `web/lighthouserc.json`;
  - `/sim`'s skeleton (`lib/sim/skeleton.ts`) is re-measured.
- **No missing-icon flash:** names wait behind the loading bar. Icons are lazy with set sizes. A failed image keeps the plain socket.
- **Slot textures:** the client's `UI-PaperDoll-Slot-*` textures ship as faithful HD reproductions in `web/public/icons/hd/slots/`. Back uses `-Chest`, as the client does. The Wowhead `inventoryslot_*` renders in `slots/` are mock stand-ins only.
- **No new font.** The island adds at most 4 KB gzip over `ItemHover`.

## 9. Acceptance screenshots (the implementer attaches all)

Each at 360, 390, 1024, 1280 and 2000 unless noted:

1. `/sim` with an addon export (a full set).
2. `/sim/<id>`.
3. The character page `#gear` with gains, best in slot, a two-hander, an empty slot and an unknown item.
4. The longest-names fixture.
5. Report "At pull".
6. Loading (throttled) and failed (blocked) at 390 and 1280.
7. Tooltip from the left column, the right column and a bottom cell at 1280, and from a cell, an upgrade row and a bottom cell at 390.
8. Side-by-sides with `raid-1280.png` and `beta-1280.png` at the same scale.
9. The Discord sheet for the same character beside the 1280 capture.

**Checklist:**

- [ ] One `Paperdoll`. Call sites 1 to 3 use it. `CharacterStrip`'s copied grid is deleted.
- [ ] Slot positions match `api/internal/sheet`.
- [ ] Clamp and ellipsis rules as §3.3. No clamped name on the worst-case fixture at any width.
- [ ] Phone upgrade rows: full width, 44 px, tappable, led by the slot word.
- [ ] Never a gain when the pick is worn. "Best in slot" instead.
- [ ] Item level only when every counted slot is known, matching §3.5. The Discord sheet follows.
- [ ] Two-hander: an empty off-hand socket with the client texture and "Two-hand in main hand". No greyed icon.
- [ ] Tooltip placement as §4.2. The tooltip never covers the crest at 880 and above.
- [ ] Every item and pick opens the shared tooltip on hover, focus and tap. There is an e2e test per call site.
- [ ] Tab order and screen-reader names as §4.1 and §6. 44 px targets.
- [ ] CLS and Lighthouse within budget. No icon or name flash.

## 10. Fix round 1: findings and fixes

| Finding | Fix |
|---|---|
| UX B1 / coordinator 1: phone gain row collides and the pick is untappable | Phone upgrade rows (§3.3): full width, 44 px, slot word first, pick name on 2 lines, tappable. `beta-390`, `beta-360`, `tip-phoneup-390`. |
| UX B2 / coordinator 2: clamp has no ellipsis, no 360 rule | §3.3 clamp table: 3 lines on phone, a 12/15 rule under 340, ellipsis and hyphenation. Worst-case boards use the longest launch-legal names. The Bronze Dragonflight signet is gone (audit `aq-wareffort`). |
| UX B3, player B4 / coordinator 3: Blight is not a two-hander, and the greyed icon is wrong | Arcanite Champion (`two_hand: true`). The client's behaviour is cited from `PaperDollFrame.lua`: the off hand is just empty. The mock draws exactly that. |
| UX B4, player B3 / coordinator 4: tooltip board broken | Six `tip-*` boards, with the panel as `ItemTooltip` renders it. Placement and flip rules are in §4.2. Every number is from the record. "Requires Blacksmithing (300)" is declined, with the reason given. |
| Player B1 / coordinator 5: a gain on a worn item | §3.4 table and a unit test. The mock's gains are a real stat comparison. |
| Player B2 / coordinator 6: partial item level | §3.5 ruling. Counted slots are defined. The figure is hidden when any counted slot is unknown. |
| Coordinator 7 / UX call sites: audit | §3.2. The home strip, guild member card and saved-sim image were checked: none draws a set. The Discord sheet is added as a parity call site. |
| UX should: stage | Added (§3.3). |
| UX should: `user` tile mismatch | The mocks draw `user`. |
| UX should: gains quieter than names | Gold head meta, 13/600 gold gains. |
| UX should: unknown too loud; note mistaken for the footer | 40% and half saturation. The note gets a ruled strip. |
| UX should: unknown plus gain | "Not scored" (§3.4). |
| UX should: 2000 at another scale | Every board is 2x. |
| UX should: "Ring 1" files | Listed in §3.1. |
| Player should: slot word on unknown slots | "Waist · Item 288887". |
| Player should: "pts" and no verb | The unit is named once in the head ("points"). Each line reads pick, then figure. The accessible name says "upgrade" and "points to gain". |
| Player should: weapons grouped centre, empty shirt and tabard | Declined for sheet parity (§2, §5). The owner can rule otherwise for both together. |
| Player should: data accuracy (Frozen Heart, Ravencrest/Shadowsong, The Purifier, PvP sets) | Those items are gone from the mocks: every set is launch-legal by the rule in §0. Frozen Heart of the Mountain stays (a crafted Enchanting trinket). Its Use is not scored, and the gain math scores stats only. |

# L3 Gear tab: info-density-reviewer

## Round 1: FIX

Page job: what do I wear, what beats it, where do I get it. Reviewed gear-hunter-60.png, gear-warrior-34.png against hub-gear.png, spec 4.3, GearTab.lua, GearTabView.lua, docs/contracts/addon-v2.md section 1 (bis `chance`, `alternatives`, `bisNames`).

What the tab already does well: every doll cell carries a gain, enchant state or "In your bags"; Upgrades has slot, source, standing and drop chance where the data has one; weights carry the DPS column; Sources counts, Off line and per-kind hover. The right rail is full and honest. Empty slots and the warrior's off hand read correctly.

### MUST

1. Left column dead space. The hunter render leaves about 300 px (roughly 4 rows) under the doll, the warrior about 220 px, while Upgrades stops at four of eight gaining slots (hunter: 12.0 of the 17.1 "in all" is shown; Nacreous +1.8, Painweaver +3.0 and Bloodmail +2.3 are only readable on the doll, with no source). Fill it with the rest of the ranking, from the same `GearTab` upgrade function (already computes every slot; `GearTab.UPGRADES = 4` is the only cut). Place it under the doll as a second block "Next upgrades" at the same row format as Upgrades (icon, name, +x DPS, slot, source via `ItemTooltip.sourceParts`, chance where `chance` exists), ranks 5 onward, capped to the free height. Source: `Data.lua` bis picks (`item_id, kind, source, chance`), `bisNames`. Why: the raider sees where the other 7 points come from without hovering 16 cells. This is the same list, not a new number, so the total and rows never disagree. Phone/narrow: the block collapses to ranks 5-6 or folds under Upgrades into one scroll list; doll first.
2. Missing enchants are a header count only ("2 enchants missing") and a red word on a cell. Add the slot names to the header figure tooltip at minimum; preferred, a one-line row at the foot of the left column ("No enchant: Chest, Feet") from `Inspect.missingEnchants(Inspect.readSelf())` (already called in GearTab.lua:60), each name clicking to flash the cell (`GearTabView.flash` exists). No invented enchant names; the contract carries none.

### SHOULD

3. Alternatives are unused. Contract section 1: every pick ships two `alternatives { item_id, kind, source }` (names in `bisNames`). Where a pick is not held and source filter shrinks the list, the doll's hover should list "Also: <alt1>, <alt2>" with source and kind, and the upgrade rows' hover likewise. This is the answer to "what if the boss does not drop it". Hover only, no new surface chrome.
4. Doll cell "Beats the pick" (warrior, 34) gives no number and no pick; confirm its hover shows the pick's name and source (spec says the hover does). Where the pick has a `chance`, the hover should show it ("1% drop"); the doll cell has no drop chance anywhere except the top four upgrades.
5. Sources counts do not say where the gain is. Optional: in the kind hover, append each item's gain next to its slot ("Chest: Blackened Defias Armor +2.5") using the existing gain function. Data exists; no new copy.
6. Sort: Stat weights is fine, but the first column should keep the DPS column alignment with sim vs shipped (the board shows a "±" column; shipped correctly omits it, per brief).

### Checked, not required

- No error on shipped weights, one list per band: respected. Do not fill the left column with a second band's list or sim figures (Rotation owns those).
- Haste line: shows only when the weights carry haste (hunter 60 shipped has none); correct.
- Raid-ready: the 0 on Raid is the spec's rule, fine.

### Needs data

- Drop chance for picks outside the top four where the ranker has none (hunter's vendor/rep items are not drops; no gap).
- Enchant names per slot (not in the contract); "No enchant" stays a slot flag.
- Per-slot stat comparison (what the worn item has that the pick lacks): needs item stat lines in `Data.lua`, absent today.

Verdict: FIX. Must items 1 and 2 (the left column ends early while the addon holds ranks 5+ and the missing-enchant slots).

## Round 2: PASS

Checked the re-made gear-hunter-60.png and gear-warrior-34.png.

- Must 1 met: "More upgrades" under the doll continues the ranking from the same list. Hunter shows Windreaver Greaves +2.3 and Mark of Fordring +1.8, matching the doll's Bloodmail +2.3 and Nacreous +1.8 gains; warrior shows four rows (+0.4 to +0.1). Each row has icon, name, slot, source and gain. The left column now ends level with the right.
- Must 2 met: "No enchant: Chest, Feet" (hunter) and "No enchant: Legs" (warrior) match the cell flags and the header count.
- Should 4 met: "Better than the list" replaces "Beats the pick" and carries the pick's tooltip on hover. Upgrade rows now show the drop chance where the data has one (1%, 90%, 20%).
- Deferred, acceptable: alternatives and per-kind gains. They stay on the optional list.

Optional: a long place is cut with an ellipsis on the narrow rows ("Kirtonos t...", "Bl...", "Leatherwor..."); the hover gives the full line, so no action.

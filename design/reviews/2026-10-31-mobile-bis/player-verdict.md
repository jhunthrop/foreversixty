# Player verdict: mobile BiS lists (iOS and Android), 2026-10-30

Seen: live captures (Fury warrior 50 to 59 Horde, both platforms), the Android picker and tooltip, the report. Not seen: any healer or tank on a device, the iOS picker or tooltip live (rendered tiles only), offline on a device.

VERDICT: HOLD

## Blockers

1. Launch gating is not proven for tonight's data. The captures and the app's fallback build are 1.60.1.70658 (BiS generated 2026-10-30 14:18Z). Build 1.60.1.70707 is in data/builds with no BiS files yet, so the re-rank is not in what I saw. I scanned all 56 spec files of 70658, bare bands 20 to 60: no Molten Core, Onyxia, BWL, ZG, AQ or Naxx drops, and no Kazzak or the Nightmare dragons. Those are clean. Re-run this scan on the re-ranked files and attach the result before shipping. The "60 · Phase 1 raid" chip also offers a raid list from 4 Nov while raids open 9 Dec. It needs "opens 9 Dec" on the chip or has to be hidden until then.
2. Tooltip, "The next best": "-0.2 DPS vs the pick" prints DPS for healer and tank bands (the report concedes tooltip.delta has no unit). A number with the wrong unit. The list row is correct (HPS / points); the sheet must use the same unit.
3. Tooltip footer: "Best in slot for Head · Fury Warrior 50 to 59 · Bare". "Bare" is the pipeline's preset name and means nothing to a player; drop it. On Android the footer also runs under the gesture bar.
4. The tooltip shows no stats at all: "Ornate Mithril Helm / Head / 383 Armor / Requires Level 44 / Item Level 49". A helm tooltip with no stats, no "Binds when picked up", no armor type is not an item tooltip, and BiS is the page where players tap to see it. Do not ship BiS ahead of the home lane's full sheet.
5. Tank deltas in "points" are not readable. Warrior-protection band 30 alternatives are -36.8 and -97.2 "points" against a pick scoring 13 to 21; the file's unit is tank_score, scaled to strength. A tank cannot tell whether -97 is a lot. Give the unit a one-line meaning, or show a relative figure, before tanks see it.

## Shoulds

- The runner-up on the row ("Next: Blood Guard's Plate Helm, same DPS") hides its source. Blood Guard's is a PvP rank 11 piece, unobtainable for a fresh level 50. Show the source on the runner-up line (Wowhead does) or skip PvP rank above 10 for "Next". The bare 20 to 50 files carry rank 12 alternatives (Legionnaire, "Scout's Blade", "Berserker Bracers"), and at 60 the file carries ranks 11 to 16. Legal as options, but the row presents them as the next step.
- PvP source text reads "PvP rank 12 · Legionnaire · Horde". Add the faction side only where it matters, and say what it costs (rank 12 is weeks of honor).
- No drop chance anywhere on a row or in the tooltip. "World drop" is fine; "Stratholme: Balzaphon" without a chance is not what Wowhead shows.
- "Crafted: Blacksmithing" does not say how to get the recipe, or the skill level needed. Add the skill and the trainer or recipe source in the tooltip.
- Level range chips: "60" next to "60 · Phase 1 raid" is unclear. Name it "60 · dungeons and PvP".
- Picker: no level-based hiding (a level 22 can open 50 to 59), no warning that a level 20 shaman's list has no mail (fine, the data handles it), but there is no "your level" marker beside the chip for a level 58. "Yours" on the title is good.
- iOS: floating tab bar covers the last row; leave bottom inset on the list.
- Fury 50 to 59 shows 17 rows, one weapon pair each for main hand and off hand: confirm that a shaman enhancement list shows no second weapon (not seen on a device).
- Units on the header line ("182.0 DPS at best in slot") are correct for DPS; for healers "HPS" needs a check on a live capture.

## Data I do not believe (for the accuracy reviewer)

- "Vendor: Lady Palanseer" for Stone Guard's Plate Armor at 50 to 59 (a Horde PvP-set-style name sold by a vendor): confirm source and the rank or reputation gate.
- "Prince Skaldrenox", "Mana Devourer", "Lord Skwol", "Lyn the Ignored" typed as kind "world": confirm these are Forever-new rare or world mobs obtainable at launch, not event or later-phase mobs.
- "Avatar of Hakkar" (Sunken Temple event) as a warlock main-hand source: confirm it is lootable in Forever at launch.
- Molten Heart of the Mountain appears as a pick or alternative in 17 spec lists at 50 and 60 while its Use is not modelled; the row correctly says "Gain not modelled yet" but check that the list does not rank it on stats alone.

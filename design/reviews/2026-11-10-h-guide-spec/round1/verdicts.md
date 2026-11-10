# h-guide-spec round 1 verdicts (captures: round1/captures, real-data build at cc7f7db37)

## ai-slop-reviewer
VERDICT: FAIL

- Rotation rails: stitched condition text ("On a party member below 30% Cast on a party member below 30%",
  "while mana holds Cast on a party member below 65%, when your mana is at least 90% of the share of the
  fight left", "mana % is over boss health %"); Greater Heal repeated per threshold.
- Priest 1440: row 18 and Cooldowns spilled under the prose without their list.
- Blank 40 px icon gutter on every "Trinkets, racials and potions" row; Two-Handed Weapon Specialization icon gutter.
- Professions poured into two columns breaks mid-sentence.
- Survival Rotation: Mark, Mana, Movement restate the rail; Stat priority dead space beside the table.
- Paladin: grey "Raid build only: needs a talent..." legend above the list; the hit-cap qualifier line.
- /sim/specs: healer and tank rows lead with dim crests; "Healers and tanks" a small grey label.
- Authentic: item names and sources, rank-tagged opener and priority with timings, the paperdoll tooltip,
  key talents with spent points, "Weight / DPS per point", player-voiced prose, phone folds.

## ux-designer
VERDICT: NOT SHIPPABLE

Must fix: (1) Professions splits mid-sentence in two columns, one column at the measure;
(2) Races at 2000: the Horde group splits across columns ("Berserking" orphaned), keep a race and a
faction whole; (3) Enchants and consumables at 2000: the consumables regroup across columns under the
wrong heading. Shoulds: the unlabelled consumables list under the prose at 1440; dead space beside Stat
priority; the Leveling tree using 450 of 980 px; Races Horde under the prose. Passes: sim-specs 390 one
column, 44 px rows; the stat lane; 390 guides; tighter prose.

## wow-player (diff-scoped accuracy, all 28 guides)
VERDICT: SHIP

Restores asked and made: druid/balance "Everywhere else, cast the newest rank."; rogue/assassination
Improved Slice and Dice "30% longer" and Improved Gouge "1.5 seconds"; mage/arcane the full Past 3
stacks rule; warrior/fury "Recklessness goes with the execute-phase Death Wish." Pre-existing
contradictions left as written (owner call): Affliction Malevolence 34/35 vs 26/27; Destruction Shadow
and Flame 35 to 38 vs full at 44; Retribution Consecration rank 1 to 53 vs to 52; Resto shaman Chain
Heal rank arrival levels.

## Round 2 changes (17f0d13f4)
Whole lists wrap with their labels (rotation, extras, races); consumables grouped by header; no lead on
a spell-less rotation line; raid-build legend removed; collapsed rotation rows drop generated words
that repeat their note; Leveling alone in two columns, Professions at the measure; sim-specs spec icons
and a real heading; Survival Mark and Movement cut. Left: "mana % is over boss health %" is curated
panel_when text that reaches the page through the generated addon-data.json (nightly owns it); the
hit-cap detail line is pinned by the stat lane contract.

# Sim-accuracy round 19: platform lane

Site branch `wt/r19-platform` (from int/r19 f91e9d3f9). Fork branch `lane/r19-platform` (from int/r19 6ea675594), one commit: the hand-synced Arcane APL copy. No engine code changed. No DPS moved for any spec.

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Arcane raid card: "Fireball: Fills in a raid once you stop blasting, with Heating Up".** The card appended every talent gate a row carries. Heating Up is a talent and its own aura, so the clause read as "wait for the stacks", on the spell that builds them. The same wording sat on Fire Blast and on the Pyroblast opener. Other cards had it too: Shadow's "Vampiric Embrace ... with Vampiric Embrace", Retribution's "with Unsealed Wrath", Demonology's "with/without Demonic Pact". | `withTalentGates`/`gateText` in web/src/lib/guides/rotation-lists.ts. Since round 18 a card lists only its own build's rows, so a talent gate on one of them is always true. Full before/after diff of every card: only those rows changed. | CONFIRMED-BUG (card) | old (b53e82298, round 15 web) | S 611d6713b |
| P2 | **Arcane raid card: one Frostbolt row mixed raid and leveling advice under "For the raid build."** The row read "In a raid when you are too low on mana for a Fireball, and as the filler once you stop blasting on the way to 60". It came from one ungated APL line that both builds share. | data/curated/apl/mage-arcane.json priority line 13 | CONFIRMED-BUG (card) | old | S 611d6713b, S fb42515b3, E 5f033ffaf |
| P3 | **Pre-regen hint lost by P1.** "with Heating Up" had been the only sign that a level 60 reader should skip Fire Blast and Fireball. | `markRaidBuildRows` tagged only a raid-only talent's own row. | follow-on to P1 | - | S ccec1214d |
| P4 | **Enhancement "Hit cap: reached".** The headline said hit was done, and its hover said hit past the cap is worth nothing. Earth Shock and the Maelstrom bolts still miss up to 16%, and the hit weight already counts them (`core.hitStepWindow`). Round 18 shaman X1. | sim/cmd/leveling-bis/hitcap.go published only the melee shape; hit-cap-line.ts and hit-cap.ts word only that. | CONFIRMED-BUG (wording + data) | old (round 17 weight fix left the line behind) | S ccec1214d |
| P5 | **Relic slot published empty below about 1.4 DPS.** Every Enhancement totem gains less than a DPS: Totem of the Storm about +0.8, Totem of Thunder +0.24. The relic gate is 2x the unpaired error of two 100-iteration runs, about 1.4 DPS. So the slot published empty, with Totem of Urgency (a healing totem) as the alternative. Round 18 shaman S3. | report.go `trinketLowGain`; weights.go `trinketGainSignificant` | CONFIRMED-BUG (ranker policy) | old | S 0c6a9483e, S 896d0defb, S f89701e49 |
| P6 | **Rank-swap harness (round 18 priest X4).** A probe that rewrites only the cast leaves a DoT condition watching the authored rank, so the line recasts every global. The production rewrite is safe within a line. A rank row **scoped to lines** rewrites only those lines, so a condition on another line could watch a rank nobody casts. None does today: the three scoped rows (Restoration druid Healing Touch, Elemental Lightning Bolt, Restoration shaman Healing Wave) have no outside watchers. | rotation_ranks.go `filterRankedLines`; scoped scan of data/curated/rank-choices.json against every curated APL | Tooling + guard (no live bug) | - | S 8bcae36ff |
| P7 | `tests/e2e/guides-sections.spec.ts` run alone (fixture data, desktop + mobile) on this branch at 8bcae36ff. | `E2E_PORT=4371 npx playwright test tests/e2e/guides-sections.spec.ts` | **58 passed**, 0 failed (1.2 min) | - | - |
| P8 | Fresh audit of shared plumbing round 18 touched: | | | | |
| | - engine-spell-names drops Berserking 26297. The engine names the one 70755 Berserking 20554. No curated APL, rank row, sim request or addon file references 26296 or 26297. | base-engine sim/core/racials.go:523-536; grep | Aligned | - | - |
| | - The addon export at 60 keeps the top ranks. The bare-preset rows at 60 (Retribution Consecration 4, Aimed Shot 1, Elemental Lightning Bolt 8, ...) do not reach the raid card or the level 60 addon entry, which matches the raid sim. Below 60 the export follows the bare rows, as the sim does. | addon-data.json level 60 entries; addonrotation.py `_band_ranks` | Aligned | - | - |
| | - Preset, buff, consume, kit, smoke and rank-choice tests in sim/request, plus the addon-rotation pytest: all pass. | `go test ./request -run 'Preset\|Buff\|Consum\|Kit\|Smoke\|RankChoice'`; `pytest tests/test_addonrotation.py` (48 passed) | Aligned | - | - |

### P1-P3: the Arcane card

- **Settled gates (P1).** `rotationListsFrom` takes `settled` (`buildSettlesGate`): a gate on one of the class's talents is settled by the described build, so it is not worded. Aura gates ("once Judgement of the Crusader is on the target") are still worded.
- **Raid card now:** "Pyroblast: 6 s before the pull", "Fire Blast: Every time it is ready", "Fireball: Fills in a raid once you stop blasting".
- **Level 60 card now:** "Frostbolt: 2.5 s before the pull". The Presence of Mind Frostbolt row ends at "...goes out instant".
- **Frostbolt split (P2).** The ungated filler is now two lines:
  - Heating Up known: "Frostbolt in a raid only when you are too low on mana for a Fireball. It costs a third less and still procs Missile Barrage 20% of the time, but it feeds no Ignite or Heating Up, so it only stretches the last of your mana."
  - Not known: the old leveling sentences, starting "Frostbolt as the filler once you stop blasting on the way to 60."
  - Every build still has exactly one Frostbolt filler, so nothing casts differently. Raid 60: 721.0 / 719.9 / 708.3 / 708.4, the same as before.tsv on all four rows. Bare 30 and 50 (horde, seed 1): identical.
  - The fork copy and sim/request copy were synced by hand; `make apl-check` matches.
  - The guide prose (arcane.md:74) already says the same, so no prose changed.
- **Collapse (fb42515b3).** A run of same-spell steps merged before the card's build filter ran, so the two Frostbolt steps would have merged into one row on both cards. The card's build filter now runs on the steps first (`collapseRuns` `keep`). The BiS page still merges every step.
- **Tags (P3).** Each row carries the spells its step needs known (`needs`). A row that needs a raid-only talent gets the "Raid build only" tag. The legend now reads "needs a talent the level 60 build does not take".
- **Reviews on the rendered card** ($R/platform/cap/arcane-rot-1440.png, -390.png):
  - wow-player: SHIP with SHOULDs.
  - ux-designer: SHIP with SHOULDs.
  - The captures used a local, uncommitted hand edit of addon-data so the split line rendered. The pipeline's own `build_rotations` produces the same two lines.

### P4: hit cap with spells

- **Ranker.** `hitToCap` gains `spell`: the distance to the 16% spell cap whenever the character casts damaging spells (`profile.Spell`). This is likely Retribution and the hunters too; the engine's profile decides.
- **Guide headline:** "Hit cap: reached (weapon attacks), 16% (spells) · 6% hit from BiS gear and buffs, 5% from Tidal Focus, 5% still to find for spells".
- **BiS rail:** "Hit to cap: 0% for specials, 5% for spells".
- **Hovers** say the spells keep gaining past the weapon cap.
- Files without the key read as before. The e2e helper `hitCapText` mirrors the new clause.
- **Reviews** (figures illustrative; $R/platform/cap/enh-hit-*.png, enh-bis-rail-1440.png): wow-player SHIP with SHOULDs, ux-designer SHIP with SHOULDs.

### P5: relic near tie

- **Rule (report.go `relicNearTie`).** A modelled relic publishes, with `near_tie: {gain, error}`, when both hold:
  - its gain over an empty relic slot is not exactly zero. Same seed, so an inert relic reproduces the empty slot to the last digit, as Urgency and Burning Totem do.
  - its gain is not significantly negative.
- **Ordering (relics.go `preferTouchingRelic`).** When the noisy ordering crowns an inert relic and the runner-up is a near tie, the runner-up takes the slot.
- **Web.** The row shows the existing quiet "best available" tag with a relic hover, and the line "+0.2 DPS over an empty relic slot, within the ±1.4 DPS margin of error".
- **Not done.** No ranker run. Option (a), a paired same-seed error, needs per-iteration data that `RunPlainDPSWithError` does not return.
- **wow-player:**
  - First capture: NO-SHIP. The pick showed no figure. Fixed in 896d0defb.
  - Re-review: SHIP. Its SHOULD wording ("within the margin of error") was adopted in f89701e49.

### P6: harness

- `request.ForceRotationRank` rewrites every ActionID naming the authored id (casts, conditions, every line, pre-pull) to the chosen rank. It fails on a rank not learned at the level.
- `spec-breakdown -rank "authoredID=rank[,...]"` applies it to `-rotation` or to the spec's own curated rotation, then runs the normal rank rewrite.
- Smoke test, priest-shadow raid 60 horde, 3000 iterations: `-rank 10894=8` gives 735.6, the same as no flag. `-rank 10894=4` gives 619.5, with the same 8.14 Shadow Word: Pain casts per iteration. The DoT condition follows the forced rank; the old harness showed -29% from recasting.
- Guard: `TestScopedRankRowsAreNotWatchedOutsideTheirLines`.

## Before/after

No DPS change for any spec. mage-arcane raid 60, guide raid build, frozen BiS, 10,000 iterations:
- alliance: 721.0 at seed 1 and 719.9 at seed 50001;
- horde: 708.3 at seed 1 and 708.4 at seed 50001;
- all four 0.00%.

Bare 30 and 50 (horde, seed 1, 3000 iterations): 0.00%.

## Files to regenerate

- **data/builds/1.60.1.70755/addon-data.json and addon/ForeverSixty/Data.lua (REQUIRED with this lane).** The curated Arcane APL gains a line. Until addon-data is regenerated, the card's step alignment (`stepsAlignedWith`) finds 3 curated Frostbolt steps against 2 published lines and drops every Arcane gate. The raid card would then show the leveling Frostbolt rows, and `guide-rotation.test.ts` "tags Arcane's raid rows that need Heating Up" stays red. Both are fixed by the regen; checked with the pipeline's own `build_rotations` output.
- The Arcane ladder golden should not move: the sims are identical. Re-run it with the regen anyway.
- BiS files: `hit_to_cap.spell` and `near_tie` appear with the next re-rank.

## Measure at refresh

- The first re-rank publishes `spell` on melee bands with damaging spells. Check that:
  - the Enhancement headline reads the real figure (about 16 minus the spell hit);
  - Retribution and the hunters read sensibly too;
  - Enhancement's "Hit: Not significant" row matches. It comes from older published weights; round 17's weight fix makes it non-zero.
- Relic slots after the re-rank:
  - Enhancement should publish Totem of the Storm or Totem of Thunder as a near tie.
  - Check that no healer relic wins a DPS spec's slot.
  - Check paladin and druid relic slots for new near ties.

## For other lanes

- **Mage:** the Arcane APL gained a line (priority 13 is now two lines, 13 and 14). Any line-scoped rank row added to mage-arcane after this must use the new indices. Nothing changes in the sim.
- **All lanes:** use `spec-breakdown -rank id=rank` to judge rank rows instead of hand-editing rotation copies.

## For the guide lane

None. No guide prose changed; arcane.md:74 already reads right for both Frostbolt lines.

## Still open (reviewer SHOULDs and pre-existing items, for the owner's list)

- **Relic slot label.** The BiS page labels the relic slot "Ranged" on paladin, druid and shaman pages, where the game says "Relic". The "NEW" pill also sits beside the near-tie tag. Both are pre-existing, site-wide, and outside this lane. wow-player calls "Relic" a must.
- **Arcane card:**
  - with the tag, 7 of 11 raid rows are "Raid build only";
  - the legend tells a raid reader to skip rows (wow-player);
  - "in a raid" repeats the tag on Fireball and Frostbolt (ux-designer);
  - the legend has no tag swatch;
  - "Fills in a raid once you stop blasting" could name the mana condition.
- **Enhancement headline:**
  - it wraps to 3 lines at 390; split the breakdown onto a muted second line;
  - align "weapon attacks" with the rail's "specials";
  - "5% from Tidal Focus" counts toward both caps, which a player may not expect.
- **Logs.** Live Forever logs were not needed for this lane's items and were not checked.

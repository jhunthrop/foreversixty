# Sim-accuracy round 18: mage

Specs: mage-arcane, mage-fire, mage-frost. Site wt/r18-mage (from int/r18 f7463a607), fork lane/r18-mage (from int/r18 8c24d6202). Client 1.60.1.70755 with the late re-layer.

Method:
- Seed items first, then the coordinator's two extra items, then the re-audit.
- Raid 60: spec-breakdown, 10,000 iterations, seeds 1 and 50001, both factions, `-bis-dir $R/bis-frozen`, guide raid builds (`ALLRT=1`).
- Leveling levels: the shipped leveling order (addon-builds.json) turned into a talent string per level (`$R/mage/tools/ladder.py`), bare preset, `-band`/`-level`.
- Every run went through `$R/sim.sh`.

**Confirmed and fixed: 2.**
- By origin: old 1 (Missiles-first stopped one level early, a guides-handoff from guide round 15, R15-1), and old 1 (a leveling line shown on the raid card; no DPS change).
- One number moves: Arcane at level 22 gains 18%. Raid 60, the bare bands and every other ladder level are unchanged.
- One engine change, in the mage package only.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Mana-gem order: the guide wording stays, and so does the engine.**<ul><li>The guide says to conjure a Ruby and a Citrine, use the Ruby once you are 1,200 mana down, and use the Citrine two minutes later (arcane.md:78, fire.md:78, frost.md:75).</li><li>No Forever source gives a gem order. The Forever guides only say to keep a gem conjured and to use it below a third of your mana (wowforeverbuilds Frost guide, 2026-09-17).</li><li>The Classic sources agree with the guide. The Vanilla wiki says to use the first charge early in a fight longer than two minutes, once you have spent enough mana not to over-restore. wowx5 says to carry a Ruby and a Citrine before a boss.</li><li>"Ruby at 1,200 down" is that same rule for the biggest gem.</li><li>Round 17's biggest-first engine run gave ±0.8% by band with no winner (round 17 X2), so the sim cannot settle it either.</li></ul> | Vanilla WoW Wiki, Conjure Mana Gem (https://vanilla-wow-archive.fandom.com/wiki/Conjure_Mana_Gem, read 2026-11-04); wowx5 conjured-items guide (https://wowx5.ru/en/guides/mage-conjured-items, read 2026-11-04); wowforeverbuilds Frost PvE guide (https://wowforeverbuilds.com/guide/frost-mage-pve-guide, 2026-09-17); WoW Forever DB utility page (https://wowfdb.com/classes/mage/utility/) | Aligned (decided on consensus) | - | - |
| S2 | **Regression targets from round 17 hold.**<ul><li>Arcane's "with the leveling build" text is at arcane.md:56, :72 and :84 and in both card notes. "Fire points" appears nowhere outside the design reviews (`git grep`).</li><li>The Missile Barrage line checks the talent.</li><li>`TestMageTalentAuraLinesCheckTheTalent`, `TestArcaneCastsMissilesBeforeMissileBarrage`, `TestArcanePresenceOfMindLinesCheckTheTalent`, `TestArcaneOpensWithPyroblast...` and `TestRound16MageRankRowsFollowTheLevel` all pass.</li><li>The fork contains 5f2852f74, and the three APL copies match the curated files.</li></ul> | tests in sim/request/mage_pull_timing_test.go; apl-check | Aligned | - | - |

## Arcane

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Missiles-first stopped one level too early.**<ul><li>The leveling order takes Missile Barrage at 22 and 23, so at 22 the mage has 1/2.</li><li>The Missiles-first line was gated on "Missile Barrage not known", so at 22 the rotation switched to Blast-first and waited for a proc that comes 20% of the time.</li><li>Leveling ladder, bare, 10,000 iterations, the same at seeds 1 and 50001:</li><li>Level 22: Missiles first 49.0 against 41.4 for Blast first (alliance), 46.6 against 39.3 (horde). The mage is out of mana 52 s of 180 s.</li><li>Level 23: Blast first wins, 51.0 against 50.2.</li><li>The APL has no talent-rank value, and the talent spell (400588) is the same id at both ranks. So the engine's Missile Barrage trigger aura now carries ActionID 400588 tagged with the rank, and the line checks `not auraIsKnown(400588 tag 2)`.</li><li>The card note and arcane.md:116 now say "20 to 22" and "from 23".</li></ul> | `TestMissileBarrageTalentAuraCarriesItsRank` (fork) and `TestArcaneCastsMissilesBeforeMissileBarrage` (site) failed, then passed; the mage package passes with `-tags=with_db`; `$R/mage/missiles-ladder.tsv` | CONFIRMED-BUG | old (round 17 A1 measured only 20, 21, 25, 30 and 40), guides-handoff (guide round 15 MAGE-R15-1) | E de926e22d, S e6c4f2730 |
| A2 | **A leveling line on the raid card (coordinator item, from the platform lane's wow-player review).**<ul><li>The Presence of Mind Frostbolt line had no gate that excludes the raid build. Line 5 (the Presence of Mind Pyroblast) always fires first on the raid build, so the line never ran there, but the raid card listed it.</li><li>It now checks `not auraIsKnown(400625)` (Heating Up: the raid build takes it, the leveling build never does). No number moves.</li></ul> | `TestArcanePresenceOfMindFrostboltExcludesTheRaidBuild` failed, then passed; raid 60 equals before.tsv on all 4 rows; bare 20, 40, 50 and 60 identical | CONFIRMED-BUG (card only) | old | S 60beaf102, E 434b54650 |
| A3 | **The Frostbolt filler is not leveling-only.**<ul><li>With the same gate, the raid build lost 0.12% on every row. It casts about 0.6 Frostbolts a fight when it cannot afford a 395-mana Fireball but can afford a 260-mana Frostbolt.</li><li>A/B on one engine puts the whole difference on this line.</li><li>Kept ungated (coordinator's call (b)). The card note now leads with the raid case. The raid loop (arcane.md:74) adds that Frostbolt costs a third less and feeds no Ignite or Heating Up, so it only stretches the last of your mana.</li></ul> | spec-breakdown A/B, `$R/mage/bd-sb*.txt` | Aligned (text added) | - | S 5c58ce8b4, S f2170d4ba, E 424026a9f, E e1aff602b |
| A4 | Re-read against the client:<ul><li>Arcane Blast 1239700 and its stack 400573: 12% of Intellect, +100% cost, 4 stacks, 18 s.</li><li>Missile Barrage 400589: −50% duration, the tick's −500 dummy, −100% cost, 15 s. The cast-complete procs are 20/40% for Blast and 10/20% for the others at 1/2 and 2/2.</li><li>Arcane Concentration 2% a rank with a 1 s proc cooldown; each missile can proc it, the channel cannot.</li><li>Presence of Mind and Arcane Power (+30% damage and cost, 15 s, 180 s).</li><li>Arcane Wrath 1326082: 451, coefficient 1.0, 5% of maximum mana, 3.5 s, 6 s cooldown. Accelerant takes 0.1 s off a rank.</li><li>Every rank text for Arcane Impact, Mind, Instability, Meditation and Subtlety matches the engine's per-rank constants.</li><li>Trainer ranks: Arcane Blast 20/30/40/50/60, Arcane Wrath 25/36/48/60.</li></ul> | spellconst/mage.json; talents/mage.json; trainables/mage.json; `talents.go:586-669`, `missile_barrage.go`, `arcane_missiles.go` | Aligned | - | - |

## Fire

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | Re-read against the client:<ul><li>Critical Mass 2% a rank; Ignite 40%; Improved Fire Blast −1 s a rank; Master of Elements 30% of base cost on Fire and Frost crits; Burning Soul 10% threat a rank.</li><li>Combustion 28682 (+10% a hit, 3 crits, 10 stacks).</li><li>Heating Up 400625: −25% a stack, 20 s.</li><li>Fire Power: 1% a rank per Fire effect, up to 4 effects. The Improved Scorch debuff counts; the mixed-school Frostfire dot does not.</li><li>Scorch, Fire Blast and Pyroblast rows match the generated tables, which `TestMageDamageLaddersMatchTheClient` checks.</li><li>The raid build's 4 surplus points (Flame Throwing 2, Burning Soul 2) are tolls. Combustion needs 30 Fire points and Arcane Meditation 15 Arcane points, so no point can move to Meditation 3/3 without dropping Precision (round 17 F2 still holds).</li></ul> | talents/mage.json; `talents.go`, `fire_power.go`; raid 60 breakdown `$R/mage/bd-fire.txt` | Aligned | - | - |
| F2 | Rank rows: Fire Blast 3 at 30-39 and 4 at 40-49, Scorch 1/2/3 at 30-39/40-45/46-49. Every rank is trainable at its band (Fire Blast 3 at 22, 4 at 30; Scorch 2 at 28, 3 at 34). | data/curated/rank-choices.json; spellranks.json | Aligned | - | - |

## Frost

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Cold Snap: the sim spends it at 0 s, the guide saves it, and an offensive use would gain 1.4%.**<ul><li>autocastOtherCooldowns casts Cold Snap on the pull, where it resets nothing. It has no GCD (12472 gcd_ms 0), so it costs nothing.</li><li>A line that casts Cold Snap while Fingers of Frost is up and Ice Lance is cooling down measured 705.7 against 696.0 (seed 1) and 705.8 against 696.1 (seed 50001), alliance: +1.4%.</li><li>Every Forever source treats Cold Snap as a defensive reset (a second Ice Block, Frost Nova or Ice Barrier), and so does frost.md:73. Wowcalc lists "Using Cold Snap with nothing on cooldown" as a mistake.</li><li>Spending it on damage gives up the second Ice Block. That is a survival-against-damage taste call, so it is not changed here.</li></ul> | `$R/mage/rot-coldsnap.json`; overgear (https://overgear.com/guides/wow-forever/frost-mage-how-to-play/), wowfdb (https://wowfdb.com/classes/mage/frost/rotation/), mmonster (https://mmonster.co/wow-forever/guides/mage/frost), wowcalc (https://wowcalc.net/en/forever/guides/classes/mage/frost), all read 2026-11-04 | OPEN (owner taste call) | - | - |
| R2 | Re-read:<ul><li>Ice Lance 1240047: 8 s category cooldown and +300% against Frozen.</li><li>Ice Barrier: the trainer gives 13033 at 58, not 1213278. It casts 3 times a fight at bands 40 and 50 (it does not recast while up).</li><li>Fingers of Frost 400669: 2 charges, 15 s.</li><li>Winter's Chill, Ice Shards 20% a rank, Frost Channeling −5% a rank, Improved Frostbolt −0.1 s a rank.</li><li>Rank rows: Frostbolt 4 at 26-29 and 5 at 32-33 are trainable.</li></ul> | spellconst; trainables/mage.json; band breakdowns | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | No new 70755 data for mage. The hotfix push touched no mage spell values, and the tables round 16 and 17 checked still hold. | hotfix-push.json | Aligned | - | - |
| X2 | The fork's mage package is green with `-tags=with_db`. Without the tag, TestP1Mage panics on item data, as round 17 noted. The new aura tag moves no fork golden. | `go test -tags=with_db ./sim/mage/` | Aligned | - | - |

## Before/after

Before: `$R/before.tsv` (raid 60), and `$R/mage/sb-nogate` (lane start, int/r18 engine and APL) for the ladder levels and bare bands. After: E e1aff602b on S f2170d4ba.

| Spec | Faction | Seed 1 before → after | % (s1 / s50001) |
|---|---|---|---|
| mage-arcane raid 60 | alliance / horde | 721.0 / 708.3, unchanged | 0 / 0 |
| mage-frost raid 60 | alliance / horde | 696.0 / 686.9, unchanged | 0 / 0 |
| mage-arcane level 22 (ladder, bare) | alliance | 41.4 → 49.0 | +18.4 / +18.4 |
| mage-arcane level 22 (ladder, bare) | horde | 39.3 → 46.6 | +18.6 / +18.6 |
| mage-arcane levels 20, 21, 23, 24, 25, 30 | both | unchanged | 0 / 0 |
| mage-arcane bare bands 20, 30, 40, 50, 60 | horde | unchanged | 0 |

No spec outside mage moves. The engine change is in sim/mage only.

## Files to regenerate

- addon-data.json and Data.lua: the Arcane card notes. The Frostbolt filler note, the Missiles-first note ("20 to 22", which still has no clause from round 17) and both new gates.
- Ladder goldens for mage-arcane: the level-22 rung now leads with Missiles.
- Coach starter data/curated/coach/mage-arcane.md, if it quotes the Missiles boundary.

## Measure at refresh

- Nothing new.

## For other lanes

- Platform: wow-player SHOULDs on the Arcane raid card. The Fireball clause "with Heating Up" reads oddly, since Fireball is what builds Heating Up. The "For the raid build" header sits above the shared Frostbolt row.
- Any lane: a talent's rank can now be read in an APL wherever an engine aura carries the rank as its ActionID tag (mage Missile Barrage is the first).

## For the guide lane

- Done here: arcane.md:74 (the raid loop's Frostbolt sentence), arcane.md:116 (Missiles 20 to 22, Blast from 23), and two card notes.
- Not done (outside this change, wow-player SHOULDs): the hit-weight unit "1 point of hit" (rendered lines 248 and 269); Brilliant Wand against Torch of Light in the level-60 list (gear, BiS refresh).

## Guide review

wow-player reviewed the rendered /guides/mage/arcane page at 1440 wide (astro dev on wt/r18-mage merged onto int/r18 b291a4b55, talent trees hydrated, card notes previewed by patching the review tree's addon-data). Captures: `$R/mage/shots/`.
- First pass: NO-SHIP. The capture had loading trees, the Frostbolt row led with leveling, and the reviewer questioned the ranks.
- Second pass: **SHIP, no MUSTs**, after the row was rewritten raid-first, the prose gained "feeds no Ignite or Heating Up", and the ranks were confirmed as Ahn'Qiraj books.

## Still open

- R1: offensive Cold Snap (+1.4% Frost raid), an owner taste call against the second Ice Block.
- Carried: Fire's dead Arcane tolls (round 17 F2); Ice Lance's coefficient; Arcane Subtlety's 15 resistance cut (SIM-R7-1); cost mods add (Arcane Power on stacked Blast and on a Barrage Missiles); the level-60 book ranks (pending owner); Warcraft Logs blocked.

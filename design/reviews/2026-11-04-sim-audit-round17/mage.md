# Sim-accuracy round 17: mage

Specs: mage-arcane, mage-fire, mage-frost. Site wt/r17-mage (from int/r17 d7fc51fe7), fork lane/r17-mage (from int/r17 f637ec32c). Client 1.60.1.70755 with the late re-layer.

Method:
- Seed items first, then the re-audit.
- Raid 60: spec-breakdown, 10,000 iterations, seeds 1 and 50001, both factions, `-bis-dir $R/bis-frozen`, guide raid builds (`ALLRT=1`). Exploration at 3,000. Every run went through `$R/sim.sh`.
- Client rows were fetched fresh from wago.tools for build 1.60.1.70755 (ItemEffect, SpellCategories, SpellCooldowns, SpellCategory, SpellEffect, SpellName) into `$R/mage/c755/`.

**Confirmed and fixed: 2.** By origin: old 1 (a phantom rotation line, no DPS change), R16-merge 1 (guide text that round 16 fixed and the guide-rounds merge put back). No mage number moves at raid 60 or at band 20. One engine change (mana gem order) was written, measured and dropped (X2).

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Mage Armor regeneration.** The guide says Mage Armor keeps half your mana regeneration running while you cast (index.md:30), and Arcane Meditation gives 17/33/50%, so three ranks with Mage Armor keep all of it and two ranks most of it (arcane.md:55, fire.md:64, arcane.md:98). The client agrees: every Mage Armor rank (6117 at 34, 22782 at 46, 22783 at 58) carries aura 134 (mana regen while casting) at 50; Arcane Meditation 18462 is 50 at three ranks and the talent text is 17/33/50%. The aura sums, so 50 + 50 = 100% and 50 + 33 = 83%. The engine adds 0.5 per Mage Armor (`armors.go:82`) and the Meditation table (`talents.go:665`). | SpellEffect 683277, 699445, 699868, 696606; spelltips/mage.json 22783, 6117; trainables/mage.json | Aligned | - | - |
| S2 | **The gem, rune and Robe sentence.** "Mana gems and Demonic Runes share one 2 minute cooldown ... The Robe of the Archmage has its own 5 minute cooldown, but using any of the three locks the other two out for a minute" (arcane.md:78, fire.md:78, frost.md:75). Client: ItemEffect gives Agate 5405, Jade 10052, Citrine 10057, Ruby 10058, Demonic Rune 16666 and Dark Rune 27869 SpellCategoryID 1153 with a 120,000 ms category cooldown. The Robe's 18385 has CoolDownMSec 300,000 and no item category. SpellCategories puts all of them, Robe included, in category 30 ("Item - Healing"), and SpellCooldowns gives each a 60,000 ms category recovery. The healthstones are in category 30 too, which the sentence does not need. The engine matches (`core/character.go:675`, `common/item_effects.go:3361`). | ItemEffect 97663, 97664, 98334, 98355, 99229, 99905, 102581; SpellCategories rows for the same spells; SpellCooldowns 53694 to 56032 | Aligned | - | - |
| S3 | **"Before you have the Fire points" was back on the page.** Round 16 changed it to "with the leveling build" (S 529a6bfca, wow-player SHIP), because the leveling build (31/4/16) never takes Pyroblast or Heating Up; its only Fire points are 4 in Incineration, at 41-44 (addon-builds.json). The merge of guide rounds 10-13 into int/r16 (c4483a3a6) took main's text and put the old wording back in arcane.md:56, :72 and :84 and in two Arcane card notes. Restored. The merge kept rounds 11-13's other fixes, and those stay. | `git diff 529a6bfca d7fc51fe7 -- web/src/content/guides/mage data/curated/apl/mage-arcane.json` | CONFIRMED (text) | R16 merge c4483a3a6 | S 88cacd3af |
| S4 | **Round 16 regression targets, all hold.** <ul><li>Arcane pre-pull Pyroblast at -6 s with Heating Up, Frostbolt at -2.5 s without: `TestArcaneOpensWithPyroblastOnTheRaidBuildAndFrostboltLeveling` passes; the prepull rows are in the curated file.</li><li>Line 4's talent gate: `TestArcanePresenceOfMindLinesCheckTheTalent` passes.</li><li>The per-level Scorch and Frostbolt rows: `TestRound16MageRankRowsFollowTheLevel` passes; no engine or rank change since round 16 touches them.</li><li>Winter's Chill: all six fork tests pass, including the rank penalty, the Frostbolt and Fingers-of-Frost Ice Lance stacking, and the client spell id.</li><li>Raid 60 on the lane's starting point equals before.tsv to the decimal on all 12 rows.</li></ul> | tests above; `$R/mage/after0.tsv` | Aligned | - | - |

## Arcane

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **A phantom line: band 20 Arcane cast Arcane Missiles on every global.** The Missile Barrage line was a bare `auraIsActive(400589)`. The engine drops an auraIsActive on an aura it never registered, and Missile Barrage registers only with the talent (`missile_barrage.go:42`), so with no Missile Barrage the line had no condition and fired on every global. Missile Barrage comes at 22, and the band 20 build has none, so every Arcane mage at 20 to 29 cast Missiles and never Blast (bare 20: 38 Missiles casts, 0 to 1.5 Blasts). This is the same kind of bug as round 16's A2. By chance, Missiles is also the right play there. At 20, rank 2 Missiles (4 missiles, 140 mana, 0.286 a missile, each missile can proc Clearcasting) beats rank 1 Blast and rank 4 Frostbolt, both per second and per mana. Measured at level 20, ladder build, 3,000 iterations, bare: <ul><li>Missiles first: 48.0</li><li>Blast to 3, then Missiles: 39.0</li><li>gate only (Blast and Frostbolt): 33.1</li><li>Frostbolt only: 33.2</li></ul>With Missile Barrage (25, 30, 40), Missiles as the filler loses: 59.1 against 59.7, 88.7 against 91.7, 170.0 against 173.5. Fix: the Barrage line now checks the talent, and a new line casts Missiles ahead of Blast while Missile Barrage is not known. The guide's leveling paragraph and a card note say so. | `TestMageTalentAuraLinesCheckTheTalent` and `TestArcaneCastsMissilesBeforeMissileBarrage` (sim/request/mage_pull_timing_test.go) failed, then passed. apl-check green. Bare 20 at levels 20, 21 and 25, both factions, both seeds: identical before and after (below). | CONFIRMED-BUG (latent, no DPS change) | old (the line predates round 15; round 16's A2 sweep read "a spell that only exists with the talent" as gated, and this one is not) | S 88cacd3af, E 5f2852f74 |
| A2 | The pre-pull Frostbolt at -2.5 s on the leveling build has no Improved Frostbolt below 45. So at bands 20 to 40 it is a 3.0 s cast (ranks 5 and up) that lands 0.5 s into the fight. Round 16 measured -2.5 s against -3 s at 60 as 717.7 against 719.3 (+0.2%, within error). Not worth a line. | spellconst 8406 and up (3000 ms); round 16 A1 | Aligned (below error) | - | - |
| A3 | The other lines re-read for talent gates: <ul><li>Line 9 (Arcane Wrath on 3 Blast stacks) and line 11 (Blast) read 400573, which only exists with Arcane Blast. Without it, Wrath and Blast are unknown spells and the actions drop.</li><li>Lines 2 and 3 (Arcane Power and Presence of Mind) cast talent spells.</li><li>Lines 4, 5, 8, 9 and 12 check Heating Up or Presence of Mind.</li></ul> | curated file | Aligned | - | - |

## Fire

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | Line 2 (Combustion on 5 Improved Scorch stacks) reads 12873 without a gate, but it casts Combustion, a talent spell, so the action drops without the talent. The leveling order takes Improved Scorch at 27-29, before Combustion at 40. Lines 3 and 4 check Improved Scorch and Heating Up. | curated file | Aligned | - | - |
| F2 | **Dead tolls in Fire's Arcane points (not adopted).** In the raid build, Improved Channeling 5 (pushback on Missiles, Blast and Wrath, none of which Fire casts) and Arcane Resilience 3 (armor) do nothing for Fire. Wand Specialization 2 does almost nothing (the wand fires only when out of mana). They are the 15 points Arcane Meditation needs. The candidate replacements in Arcane's first three rows are: <ul><li>Arcane Subtlety 2: the 15 resistance cut is real but unmodeled (SIM-R7-1);</li><li>Magic Absorption: resistances, not damage;</li><li>Missile Barrage: needs a Missiles line Fire does not have.</li></ul>None is a measured gain, and Classic 17/31/3 guides (no Forever source found) fill these rows the same way. Owner or guide-lane taste call, not a sim bug. | talents/mage.json; WebSearch 2026-11-04: Warcraft Tavern Classic PvE Fire builds (https://www.warcrafttavern.com/wow-classic/guides/pve-fire-mage-talents-builds/), Noob to Boss (https://noobtoboss.com/wow-classic-mage-talent-tree-builds/) | OPEN | - | - |

## Frost

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Fingers of Frost against the client: 15% on a Chill, 2 charges at rank 2, 15 s, and charges spent by the next direct-damage casts with Shatter's 50%. Ice Lance is +300% against Frozen. Ice Barrier is 15% for 1 min plus 5% while shielded (98% uptime). Cold Snap is a 10 minute reset with nothing to reset in this fight (a tier filler). Cone of Cold is never cast at raid range, so Winter's Chill is a toll, as the guide says. | talents/mage.json; `fingers_of_frost.go`; raid 60 breakdown | Aligned | - | - |
| R2 | The guide says "each Frostbolt then has a 40% chance to add a stack" (frost.md:63). The client also gives the Fingers of Frost Ice Lance the chance, and the engine models both. That is an omission in the guide, not an error, and Frost never casts Cone of Cold at range. | talents/mage.json 11180 | Aligned (guide lane may add it) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | No new 70755 data for mage since round 16 (no mage spellconst, trait curve or hotfix change; round 16 X1 still holds). | - | Aligned | - | - |
| X2 | **Mana gem order: written, measured, dropped.** A level 60 mage carries all four gems. The cooldown manager fires the smallest gem whose own small deficit is met, so the first window goes to an Agate at about 425 down, early in the fight, and a Ruby comes two minutes later. The guide's order (Ruby at 1,200 down, then Citrine) starts the shared two-minute timer later, and the second gem often does not fit before the fight ends. A biggest-first fix measured, 10,000 iterations, seeds 1 / 50001: <ul><li>raid 60: Arcane -0.04 to +0.08%, Fire -0.11 to -0.16%, Frost within ±0.07%</li><li>bare 50: Arcane +0.2 / +0.4%, Fire +0.6 / +0.5%, Frost +0.7 / +0.6%</li><li>bare 60: Arcane +0.3 / +0.4%, Fire -0.8 / -0.8%</li></ul>A Citrine threshold at its minimum instead of its maximum changed nothing. The results go both ways, so this is not a bug, and the commit was dropped. The guide's gem timing could start the timer earlier (see "Still open"). | `$R/mage/bare.txt`, `$R/mage/v2.sh` | Aligned (dropped) | - | - |
| X3 | Fork test runs: the mage package has 6 tests that panic on missing item data (TestP1Mage, both Robe tests, TestFrostSpellsCarryTheirMasksWhenRegistered and two more). They fail the same way on the round-start engine (integration open item 8). Each test was run alone: 113 pass, 6 item-data panics, the same before and after. | `$R/mage/pkg-before.txt`, `pkg-after.txt` | carried | - | - |

## Before/after

Before: `$R/before.tsv` (raid 60) and `$R/sb-before` on `$R/site-before` (bare 20). After: E 5f2852f74 on S 88cacd3af.

| Spec | Faction | Seed 1 | % (s1 / s50001) |
|---|---|---|---|
| mage-arcane raid 60 | alliance / horde | 721.0 / 708.3, unchanged | 0 / 0 |
| mage-fire raid 60 | alliance / horde | 735.6 / 721.0, unchanged | 0 / 0 |
| mage-frost raid 60 | alliance / horde | 696.0 / 686.9, unchanged | 0 / 0 |
| mage-arcane bare 20, levels 20, 21 and 25 | alliance | 47.8, 48.0, 56.1, unchanged | 0 / 0 |
| mage-arcane bare 20, levels 20, 21 and 25 | horde | 45.2, 45.5, 53.3, unchanged | 0 / 0 |

No spec outside mage moves. No shared code was changed.

## Files to regenerate

- addon-data.json and Data.lua: the Arcane card clauses. Line 5 still reads "before you have the Fire points", and the new pre-Barrage Missiles line (now line 7) has no clause.
- Coach starter data/curated/coach/mage-arcane.md, which carries the old Presence of Mind wording.
- Ladder goldens for mage-arcane: the rung casts do not change (A1 keeps Missiles first without Missile Barrage), but the line indexes after 6 move by one.

## Measure at refresh

- Nothing new. Arcane's band 20 BiS talents in the frozen file are the pre-round-15 walk (Arcane Resilience, not Arcane Blast, at 20). The refresh should write the ladder rung.

## For other lanes

- Platform or guide lane: the Arcane opener card's Frostbolt clause reads Improved Frostbolt from the raid tree, so it says 3 s for a cast the 60 leveling build makes in 2.5 s (wow-player SHOULD). The APL talent-gate test idea (any line that reads a talent's aura must also check the talent) could cover every class. Mine checks the four mage proc auras.

## For the guide lane

- Done here: arcane.md:56, :72, :84 ("with the leveling build") and the leveling paragraph (:116, Missiles at 20 and 21), plus two card notes and one new card note.
- Optional: frost.md:63 could add the Fingers of Frost Ice Lance to Winter's Chill's 40% (R2).

## Guide review

The wow-player agent reviewed the rendered /guides/mage/arcane page, served by astro dev at 1440 wide. The full-page capture and the page text are in `$R/mage/shots/`. **Verdict: SHIP, no MUSTs.** The three "with the leveling build" lines and the new leveling sentence are true, in a player's voice, and consistent with the page. SHOULDs, none blocking:
- The opener card's Frostbolt clause says "3 s before the pull", while the prose says 2.5 s. The clause logic (web guide-rotation) reads Improved Frostbolt from the raid tree, which has none, but the Frostbolt opener belongs to the leveling build, which has it at 60. At 45 and below, 3 s is right (A2). For the platform or guide lane.
- The "lead with Missiles" advice sits mid-paragraph and could be its own sentence.
- Card line 5 must read "with the leveling build" once addon-data regenerates (listed above).
- The Level 20 to 29 band card shows only the 11-point build. This is a band-card quirk, not mage text.

## Still open

- F2: Fire's dead Arcane tolls (owner or guide lane).
- Gem timing (X2): whether the guide should say to pop the first gem early, to start the two-minute timer sooner. The sim's spread is ±0.8% by band, and it needs real-fight evidence.
- Carried: Ice Lance's coefficient; Arcane Subtlety's 15 resistance cut (SIM-R7-1); Arcane Concentration from Shoot; the level-60 book ranks (pending owner); Warcraft Logs is blocked.

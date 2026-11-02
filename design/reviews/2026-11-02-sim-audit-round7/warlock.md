# Sim-accuracy round 7: warlock (2026-11-02)

Specs: Affliction, Demonology, Destruction and their demons. Client 1.60.1.70707. Pin b57765fb9, site origin/main 9cfb91c47. Lane engine `lane/r7-warlock` (fork), site branch `wt/r7-warlock`.

Method: spec-breakdown, raid preset, band 60, 10000 iterations at seeds 1 and 50001, horde troll and alliance gnome. "Before" is origin/main 9cfb91c47 on b57765fb9 (a detached worktree with its own `make simdb`); "after" is this branch on the lane engine. Client rows: the 70707 tables from wago.tools (SpellEffect, SpellMisc, SpellAuraOptions, SpellClassOptions, SpellCooldowns, SpellCategories, ItemSparse, Item, ItemEffect, ItemXItemEffect; fetched 2026-11-01). The server push was re-dumped from https://wago.tools/hotfixes on 2026-11-01: no row newer than 2026-10-31 10:59 UTC, so `r5/hotfix70707/diff-vs-base.tsv` is current. Searches: `talent-search -preset raid -faction horde -keep <raid-keeps.json>` and `rotation-search -preset raid -faction horde`.

Logs: still no level-60 Forever log (launch is 2026-11-04). forever-bugs #331 (demon crit) has not moved since 2026-10-14.

## Assigned item 1: side effects of round 6

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Wand shots in the wand's school. Torch of Light (Arcane) averages 127 a hit on a Shoot-only bare band-60 rotation, in all three specs (the midpoint less the level-63 partial resist; it was 81 as Physical). Bare band 60 shoots 0.38 times a fight (Affliction), 0 for the others; raid 0. Bands 20 and 40 shoot 0. The "out of mana" seconds the breakdown prints for bare Demonology and Destruction (3.9 and 4.5 s) are the engine's OOM clock, which starts at the first unaffordable cast even while cheaper spells keep casting; no Shoot is owed there. | spec-breakdown `-preset bare -band 20/40/60` and a Shoot-only `-rotation`; `core/mana.go:334` | Aligned | - |
| R2 | forever-bugs #419 (open, 2026-10-21, no reply): a player's shadow wand ignores Shadow Weaving, Darkness and Shadowform. The engine lets a Shadow wand take the warlock's Shadow multipliers (Improved Shadow Bolt's vulnerability, Demonic Sacrifice). Warlocks shoot 0% of a raid fight and the published wands are Arcane, so nothing published moves. | https://github.com/ClassicWoWCommunity/forever-bugs/issues/419 (read 2026-11-01) | OPEN (single source; no warlock number moves) | - |
| R3 | Dynamic DoTs. Corruption, Immolate's ticks, Siphon Life, Bane of Agony and Bane of Doom re-read spell power and the caster's multipliers each tick (`core/dot.go:265`). Shadow Mastery is a spell-level additive (`talents.go:357`), so every snapshot carries it. Curse of the Elements, Curse of Shadow and Improved Shadow Bolt are target-side: `TargetDamageMultiplier` (`core/spell_result.go:676`) applies them when the tick lands, including Shadow Vulnerability through `DamageDoneByCasterMultiplier` (warlock-only, aura 270, 4% a rank, 12 s, no charges: client 17794, DurationIndex 29, no ProcCharges). Siphon Life folds the target's modifiers into its snapshot and sets IgnoreTargetModifiers for the tick, so they count once. | engine code; client 17793/17794 | Aligned | - |
| R4 | Classic-db single-school spell damage reaches the sim on the committed simdb.bin: Dark Advisor's Pendant (+20 Shadow) 567.1 against the published neck's 564.4, Maleki's Footwraps 565.4, Skul's Ghastly Touch 566.8 (Affliction horde raid, seed 1, 5000 iterations); Eye of Flame (+43 Fire) 623.6 against 620.6 for Destruction. No published warlock item is classic-db sourced, so the fix moves no published set; these items stay out of the ranker as client-unconfirmed (round 6 C2, integrator policy). | spec-breakdown `-gear` | Aligned | - |
| R5 | Improved Scorch out of the raid preset: d1ddb012a removes only the `improved_scorch` entry. Nothing else the warlock preset carries went with it (Curse of the Elements, Curse of Recklessness, Judgement of the Crusader, Sunder, Faerie Fire, Hunter's Mark and the buffs are all still resolved). Forever's Curse of the Elements (1311680) is aura 87 +10% on mask 126 and aura 22 -75 resistance, as the engine's raid aura states. | `git show d1ddb012a`; client 1311680; tooltip 1311680 | Aligned | - |

## Assigned item 2: ledger rows

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| L1 | Improved Corruption, add or multiply. New client detail, not new evidence: Improved Corruption's damage (17810 effect 1) and Shadow Mastery's periodic half (18271 effect 1) are both aura 108 on op 22 (periodic damage), on masks that both cover Corruption. The engine multiplies them (1.10 x 1.10). A server that sums same-op percent mods gives 1.20, 0.8% less Corruption, about 0.1% DPS. Only a level-60 log settles it. | client SpellEffect 17810, 18271 | OPEN (blocked on logs) | - |
| L2 | Round 5's demon items (crit, haste, hit, level-60 sheets): nothing new. #331 open since 2026-10-14, #394 closed, no level-60 log. | forever-bugs #331, #394 | OPEN (blocked on logs) | - |

## Assigned item 3: Torch of Light out of the raid sets

bis run 41363895819 was cancelled by the owner, so the picks below come from a local ranker run on this branch (`leveling-bis -bands 60`, the committed simdb.bin), not committed.

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| B1 | Torch of Light is gone from every band-60 warlock set. Brilliant Wand replaces it in all six raid sets and in the bare Affliction and Destruction sets; bare Demonology takes Greater Eternal Wand, with Brilliant Wand at 0 delta for horde. That is right: a raid warlock never shoots, so the slot is decided by the wand's stats, and Brilliant Wand carries the most spell damage of the confirmed wands. Pyric Caduceus and the Eternal wands are within about 1 DPS. | local ranker output (scratch `r7/warlock/lb-out/`) | Aligned | - |
| B2 | Other slots move on the re-rank: Uther's Strength into trinket2 on most sets, Underworld Band on Affliction, three Dreadmist pieces on alliance Destruction raid (the client's 3-piece +18 spell damage, modeled as stats), Magister's pieces on alliance Demonology raid. These are the ranker's call on the new data. | same | for the integrator | - |

## Cross-cutting (fresh re-audit)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| C1 | **The warlock rotations held the Major Mana Potion to the execute phase.** All three curated rotations opened with a potion line whose condition was "remaining time at least 15 s AND mana at or below 75% AND execute phase (below 35%)". Its note said "at the top of the fight", but the AND with the execute phase meant the potion waited for the last 35%: 1.00 potion in a 180 s fight instead of 1.99, and two more Life Taps paid for it. No other spec's rotation names the potion; the engine drinks it on each cooldown when mana allows, which the raid preset's own note for Major Mana Potion promises. The rotation search flagged the line as wrong for all three. The line is removed (and from the fork's copies). Test `TestWarlockRaidRotationsDrinkAManaPotionOnEachCooldown` (failed at 1.00, then passed). | rotation-search probe "#1 (condition)" hurts; spec-breakdown item 13444 casts 1.00 to 1.99; `data/curated/apl/warlock-*.json` | CONFIRMED-BUG | S a089bcf0c, E 7a03d5dc7 |
| C2 | **Each warlock carried one school elixir, and Destruction the weaker one.** Destruction took only Elixir of Fire Power (+10 Fire, client 7844) because its rotation is "fire-led"; Affliction and Demonology only Elixir of Shadow Power (+40 Shadow, client 11474). Classic elixirs of different schools stack: the Battle and Guardian categories came in patch 2.1.0 (warcraft.wiki.gg Patch 2.1.0), and Forever's tooltips carry no category label (Wowhead Forever 21546). Every spec deals both schools: Destruction's Corruption, banes and Shadowburn are about 29% of its damage, so +40 Shadow alone beats +10 Fire for it (625.0 against 620.6), and both together are 628.9. Forever has no +40 Fire elixir: the Greater Firepower recipe (21547) makes the Elixir of Holy Power (21546, spell 1310077). Every warlock now carries both. Test: `TestRaidPresetConsumesFollowTheRole` (failed on all three, then passed). | client SpellEffect 7844, 11474, 26276; Item 21546; Wowhead Forever item 21546/21547; spec-breakdown `-consumes` | CONFIRMED-BUG | S 240c4cde0 |
| C3 | **No warlock spec weighed Spirit, and Affliction did not weigh Fire power.** Forever's Life Tap is (424 + Spirit) x Improved Life Tap at level 60 (client 11689 effect 0, 420 + 1 a level; text "Spirit increases the amount converted"), and every published set taps 12 to 18 times a fight, so Spirit buys back Life Tap globals: Divine Spirit's +40 is +1.07% (Affliction) and +1.09% (Destruction) on the bare band-60 sets. `weight_stats` (data/curated/specs.json) left Spirit out, so the ranker and stat weights scored it at zero; the ranker test's comment "warlocks, on Life Tap, measure neither" predates Forever's Spirit-scaled Life Tap. Affliction's Immolate is about 11% of its damage, but Fire power was not weighed. Test `TestEveryWarlockSpecWeighsSpiritAndBothOfItsSchools` (sim/specs; failed, then passed). The weights themselves are nightly-owned and regenerate. | client 11689; tooltip 11689; spec-breakdown `-preset bare -buffs divine_spirit` | CONFIRMED-BUG | S 289a0f8b1 |
| C4 | Alliance race: Human (Ruthless, +1% spell and attack crit, client race ability 664) beats the published Gnome by +0.50/+0.50% (Affliction), +0.29/+0.34% (Demonology), +0.49/+0.40% (Destruction), seeds 1/50001. Under the 1% bar, so Gnome stays. | spec-breakdown on a scratch bis dir with race set to human | Aligned (held by the adoption-rule bar) | - |
| C5 | Item stats on every published band-60 warlock piece match Wowhead Forever's tooltips (19 items read 2026-11-01). Kindling Stave, Spritecaster Cape and Satyrmane Sash have no 70707 ItemSparse row and are no longer on Wowhead Forever's feed; their stats come from saved Wowhead Forever records (versionNum 16001, `raw/wowhead/items.json`). Rune of the Guard Captain and Frozen Heart of the Mountain's "chance to hit" is stat 31, which `pipeline/simdb/ratings.py` maps to every school as Forever's unified hit. | https://nether.wowhead.com/forever/tooltip/item/<id> | Aligned (three items single-source, as the pipeline lane records) | - |
| C6 | Racials: Troll Berserking +10%, 15 s, 3 min; the Gnome's Eureka! per-class row; Expansive Mind +5% mana. No racial row in the push touches Troll or Gnome (the push changes Orc Command, Undead Touch of the Grave, Tauren, Human and Skyborne rows). | `core/racials.go`; diff-vs-base.tsv | Aligned | - |
| C7 | The push: apart from round 6's Attributes_14 bit on SpellMisc and round 5's Improved Drains, no warlock spell, talent or demon row changed. | diff-vs-base.tsv, re-dumped 2026-11-01 | Aligned | - |

## Affliction

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | Raid search: keeping every unmodeled guide talent, 1 Improved Drains to Improved Bane of Agony gives +4.2 ± 1.4 (+0.7%), below the bar; the larger winners drop Grim Reach. Held, as in rounds 4 to 6. | scratch `r7/warlock/ts-out/warlock-affliction-raid.md` | Aligned (held) | - |
| A2 | Rotation search: besides C1, swapping Siphon Life below Immolate measured +2.2 ± 1.3 at 800 iterations, but -0.2%/-0.1% at 10000 (571.4 to 570.4, 571.6 to 571.0). Kept. | rotation-search; spec-breakdown `-rotation` | Aligned | - |
| A3 | Fresh checks: Pandemic +33/67/100% crit bonus by rank on Corruption, the banes, the drains, Siphon Life and Wrack (`talents.go:937`), Malevolence +1% a rank on warlock Shadow spells, Improved Bane of Agony +10%, Life Tap (424 + Spirit) x 1.2, Bane of Agony's ramp at the per-tick read (#398), Siphon Life 41 + 0.05 a tick for 10 ticks (binary, no partial resist). | talent text in talents/warlock.json; client 11689, 18881; engine | Aligned | - |

## Demonology

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| D1 | Raid search: the guide build is the best (628.1 ± 1.4). | scratch `ts-out/warlock-demonology-raid.md` | Aligned | - |
| D2 | Rotation search: the potion line (C1) is the only change it found. | rotation-search | see C1 | - |
| D3 | Fresh checks: the prepull (Imp, Demonic Sacrifice, Succubus, Soul Link), Master Demonologist (Succubus) +10% Shadow for both, Burning Shadow +15% Shadow, Soul Link +3%, Demonic Knowledge +60, Decimation's Soul Fire, all against the tooltips. Demon melee glances at 40% against the boss, as Classic pets do. | Wowhead Forever 18788, 23785, 19028, 412732 (read 2026-11-01) | Aligned | - |

## Destruction

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | Raid search: the guide build is the best (618.3 ± 1.0). | scratch `ts-out/warlock-destruction-raid.md` | Aligned | - |
| S2 | Rotation search: besides C1, swapping Immolate below Conflagrate measured +3.2 ± 2.4 at 800 iterations and +0.2%/+0.1% at 10000 (635.5 to 636.9, 636.0 to 636.7). Under the bar; kept. | rotation-search; spec-breakdown `-rotation` | Aligned (held) | - |
| S3 | Fresh checks: Ruin +20% crit bonus a rank on Destruction spells (Shadowburn included), Aftermath 4/5 +40% on Immolate's hit, Pandemic 2/3 (+67%) on Corruption and the banes, Cataclysm -10% cost. | talent text; `talents.go:854`, `:937` | Aligned | - |

## Before/after

Raid preset, band 60, 10000 iterations, DPS. Before: origin/main 9cfb91c47 on b57765fb9. After: this branch (a089bcf0c, 240c4cde0, 289a0f8b1) on lane/r7-warlock 7a03d5dc7. C3 moves no DPS here; it changes what the ranker and weights measure.

| Spec | Horde s1 | Horde s50001 | Alliance s1 | Alliance s50001 |
|---|---|---|---|---|
| Affliction | 564.3 to 571.4 (+1.26%) | +1.24% | 563.0 to 569.9 (+1.23%) | +1.32% |
| Demonology | 628.5 to 638.7 (+1.62%) | +1.49% | 625.8 to 634.5 (+1.39%) | +1.41% |
| Destruction | 620.6 to 635.5 (+2.40%) | +2.40% | 617.8 to 632.5 (+2.38%) | +2.44% |

Of that, the potion (C1) is +1.1 to +1.4%; the elixirs (C2) are +1.3% for Destruction and +0.1 to +0.2% for the others.

The fork suite (`-tags=with_db ./sim/warlock/... ./sim/core/...`) passes; no golden moved. sim/request TestRotationLadder goldens are not regenerated, as the brief requires (they already differ at origin/main for twelve specs).

## Still open

- L1: Improved Corruption and Shadow Mastery's periodic half, add or multiply (about 0.1%).
- L2: demon crit, haste, hit and the level-60 sheets; Demonic Brand unmodeled; Siphon Life tick crit held for the owner.
- R2: forever-bugs #419, shadow wands and Shadow multipliers (no warlock number moves).
- Affliction's Improved Bane of Agony swap (+0.7%) and Destruction's Immolate/Conflagrate swap (+0.1 to +0.2%) are held.
- For the integrator: the local re-rank's picks (B1, B2) were not committed; bis is off until the zero round.
- Outside this group: Fire mage carries only Elixir of Fire Power (+10) with the same "no Greater Firepower" note; that is right for Fire alone, but any caster dealing two schools could take a second school elixir. Not measured here.

Sources (read 2026-11-01): https://wago.tools/db2/<Table>/csv?build=1.60.1.70707; https://wago.tools/hotfixes (70707); https://nether.wowhead.com/forever/tooltip/spell/11689 (and 1311680, 11735, 18788, 23785, 19028, 412732); https://nether.wowhead.com/forever/tooltip/item/21546 (and 21547, 19120, 249469 and the 19 published pieces); https://warcraft.wiki.gg/wiki/Patch_2.1.0; https://github.com/ClassicWoWCommunity/forever-bugs/issues/419, /331, /394.

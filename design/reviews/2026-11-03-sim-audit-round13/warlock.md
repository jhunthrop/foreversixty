# Sim-accuracy round 13: warlock (2026-11-03)

Specs: Affliction, Demonology, Destruction and their demons, bands 20 to 60, bare and raid presets. Client 1.60.1.70753. The hotfix push is still empty for every table (data/builds/1.60.1.70753/hotfix-push.json: 0 added, changed, removed or withheld), so the shipped client tables are the truth (round 11's copies in scratch `r11/c70753/builds/1.60.1.70753/raw`, plus TraitDefinitionEffectPoints and CurvePoint for talent ranks). Site int/r13 @ 93c5d2510, fork int/r13 @ 076a5a517. Lane branches: site `wt/r13-warlock`, fork `lane/r13-warlock` (no fork commit this round).

Method: the round-12 report first, then the owner's ruling 2, the round-12 holds, and a fresh pass over every spell the three APLs cast and every talent the published builds take. spec-breakdown at 10,000 iterations, seeds 1 and 50001, both factions (published races: Gnome and Troll).

Totals: **0 CONFIRMED-BUGs.** One owner-ruled build change (ruling 2), made. One new build proposal under the same precedent (Demonology, below 1%), left for the coordinator. Nothing outside the warlock group moves.

## Ruling 2: Improved Bane of Agony

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **The Affliction raid build moves its Improved Drains point to Improved Bane of Agony.** Old raid build 25520100135001050-0005200000000000000-0550004000000000 (guide `2552010013500105/00052/0550004`), new 25520010135001050-... (`2552001013500105/00052/0550004`). Improved Drains 403511 only raises Drain Life, Drain Soul and Drain Mana; the Affliction APL (data/curated/apl/warlock-affliction.json) casts 11661, 11668, 11672, 11689, 11713, 1311680, 18288, 18881 and 5019, so no drain. Improved Bane of Agony 18827 rank 1 is +5% Agony damage (talents/warlock.json; engine `talents.go:1047`). Both talents are tier 1 and tier 2 of Affliction, and the tier counts still hold (tier 0: 12 points, tiers 0 to 2: 19). Nothing the rotation casts is lost. Raid 60, 10,000 iterations: horde 571.9 to 576.0 (seed 1) and 571.7 to 575.8 (seed 50001), alliance 570.0 to 574.1 and 569.9 to 574.0; +0.72% at every seed and faction, error ±0.29. The guide's raid paragraph now says the drain point went to Improved Bane of Agony and why. `_prose.test.ts`, `builds.test.ts`, `_sections.test.ts`, `guide-builds.test.ts`, `key-talents.test.ts` pass. | client 403511, 18827; APL; spec-breakdown -talents | owner ruling, adopted | - | S 66fa0d042 |
| R2 | Which builds: the raid build only. Demonology and Destruction take no Improved Drains in either build. Affliction's leveling build (bare bands 30 to 60) keeps Soul Harvest 2 and Improved Drains 3: the bare APL casts no drain either, but the guide's leveling play casts Drain Life, Drain Soul and Wrack while questing, so the ruling's "does nothing" test is not met there. Left as is; see Still open. | guide frontmatter; bis/warlock-affliction.json bands | Aligned | - | - |

## Round-12 holds

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | Improved Bane of Agony: resolved by R1. | | done | - | S 66fa0d042 |
| H2 | Destruction: Conflagrate above the Immolate refresh (lines 6 and 7 swapped, round 12's `destro-swap.json`), re-measured on int/r13: horde 636.2 to 637.5 (+0.19%) and 636.3 to 637.6 (+0.20%), alliance 634.0 to 634.8 (+0.12%) and 633.5 to 635.0 (+0.23%), error ±0.3. Beyond error at three of four, under 1% at all. Still held. With Shadow and Flame 5 Conflagrate never eats Immolate (426316 effect 1, 100%), so the order only decides which instant goes first. | spec-breakdown -rotation | Aligned (held) | - | - |
| H3 | Round 12's other open rows (demon crit and sheets at 60, P2; additive versus multiplied percent modifiers; Ignition's rate; Siphon Life tick crit held for the owner; shadow wands #419): unchanged, no level-60 log exists yet (api.foreversixty.gg recent reports are all below 60). | | OPEN | - | - |

## Affliction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Per-cast costs read off the raid breakdown match the client: Shadow Bolt 370, Corruption 290, Immolate 370, Siphon Life 365, Bane of Agony 215 (SpellPower). Cast counts match durations and the 10.6% spell miss: Agony 8.3 casts (24 s), Corruption 10.8 (18 s), Siphon Life 6.7 (30 s, 55 of 67 ticks, the last cast cut by the fight's end). | spec-breakdown raid horde | Aligned | - | - |
| A2 | Life Tap 11689: "Converts ($m1 + Spirit) x (1 + Improved Life Tap)" with effect 0 bp 420, 1 a level from 56, coefficient 0. Engine `lifetap.go:27-44`: 424 at 60 plus Spirit, x1.2 with 2 points. Breakdown: Affliction 794 a tap (ILT 2), Demonology 662 (ILT 0), both (424 + 238 Spirit) times the talent. | SpellEffect 11689, 18182 | Aligned | - | - |
| A3 | Nightfall 18094 (leveling build only): 2% a rank, on Corruption, Drain Soul, Drain Life and Wrack ticks; Shadow Trance 17941 is aura 108 misc 10 -100% on Shadow Bolt (mask bit 0), 10 s. Engine `talents.go:301-350` rolls on every periodic tick of those four, makes the next Shadow Bolt instant and consumes the trance only on an instant bolt, so a trance gained mid-cast waits for the next bolt. | 18094, 17941 | Aligned | - | - |
| A4 | Bane of Agony ramp and crit: ticks 1-4 at half, 5-8 full, 9-12 one and a half (`curses.go:28-33`), Amplify Curse x1.5 snapshotted at the cast (#398). A DoT's crit multiplier is 1 + 0.5 x (1 + Pandemic): 1.5 base, 2.0 with Pandemic 3 (`core/spell.go:298`, `spell_outcome.go:1083`). | `curses.go`; core | Aligned | - | - |
| A5 | Talent build: every point checked against the APL. Grim Reach 2 (range) stays on the unmodeled keep rule; Demonic Aegis 2 is on the raid search's -keep list. No other point does nothing. | design/reviews/talent-search/warlock-affliction-raid.md | Aligned | - | - |

## Demonology

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | Decimation 440870: the per-rank values live in TraitDefinitionEffectPoints (trait 135428), not SpellEffect, whose effect 3 reads bp 0. CurvePoint: cooldown -45/-90% (curve 83196), cast -20/-40% (83197), damage at 35% +3/+6% (110390). Engine: `soul_fire.go:84`, `shadowbolt.go:75`, `searing_pain.go:71` use 90%, 40%, 6% at 2 points. Soul Fire 17924: 6.0 s, 335 mana, 60 s category cooldown, coefficient 1.0, speed 24 (`soul_fire.go`); 2.4 s with Bane and Decimation, 8.1 casts in the last 35%. | SpellEffect, TraitDefinitionEffectPoints, CurvePoint | Aligned | - | - |
| D2 | Bane of Doom 603: 1742 after 60 s, coefficient 4.0, 300 mana, 60 s cooldown, PERIODIC_CAN_CRIT (Attributes_8 0x200); breakdown 300 a cast. Cast only with a minute left, Agony otherwise. | 603 | Aligned | - | - |
| D3 | Demonic Sacrifice 18788 / Burning Shadow 18789 (aura 79 +15% Shadow, 2 h; effect 1 is the "any summon cancels it" script on the summon mask, which Demonic Pact 425464 lifts for the Succubus). Soul Link 19028 (20% base mana) and 25228 (aura 79 +3% every school, both units). Master Demonologist 23785: 2% a rank, Succubus +10% Shadow for both units (`talents.go:456-540`). Demonic Knowledge 412732: 100% of level. Summon Imp 80% and Succubus 100% base mana (SpellPower PowerCostPct), 10 s casts. | client rows; `talents.go` | Aligned | - | - |
| D4 | **Proposal, not adopted: Improved Voidwalker 1 to Improved Life Tap 1 in the raid build.** The raid build never has a Voidwalker out, so Improved Voidwalker 18705 does nothing, the same test as R1. Improved Life Tap is an Affliction tier-0 talent; Life Tap is cast 15 times a fight. Demonology tiers 0 to 3 keep 20 points, so Soul Link and Demonic Knowledge stay legal. 05000000000000000-2035113011220001351 to 15000000000000000-2035103011220001351: horde 639.3 to 643.8 (+0.71%) and 639.8 to 644.3 (+0.69%), alliance 635.5 to 639.6 (+0.64%) and 634.7 to 639.5 (+0.75%), error ±0.38. Moving Improved Health Funnel 2 into Improved Sayaad as well adds about +0.16%, but the guide tells the raid warlock to Health Funnel the Succubus (demonology.md:95), so that point has a use. The coordinator decides whether ruling 2's precedent covers this; if yes, also change the line at demonology.md:65 ("Improved Health Funnel, Demonic Aegis and Improved Voidwalker add no damage...") and the raidBuild. | spec-breakdown -talents | OPEN (coordinator call) | - | - |

## Destruction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Costs with Cataclysm 3 (17778, -10% on the Destruction mask): Incinerate 325 to 292.5, Immolate 370 to 333, Conflagrate 255 to 229.5 (breakdown 226.7, Eureka's -10% for 8 s every 2 min on the Gnome). Conflagrate 18932: 10 s category cooldown, 16.6 casts. Shadowburn 18871: 15 s category cooldown, 11 casts. Incinerate 1293813: 217 + 1.4 a level, 0.714, 2.5 s, speed 20, +25% on Immolate (effect 1). | SpellPower, SpellCooldowns, SpellEffect | Aligned | - | - |
| X2 | Shadow and Flame 426316: 10% Shadow after Conflagrate, 10% Fire after Shadowburn, 20 s; at 5 points Conflagrate keeps Immolate and Shadowburn refunds its shard (effect 1, 100%). Engine `conflagrate.go:84-87`. Uptimes 97% and 90%. Fire and Brimstone: Conflagrate crits 35.7% (10% + 25%). Ruin: crit 1 + 0.5 x (1 + 1.0) = 2.0x on the Destruction flag (`talents.go:896`). | 426316, 412751, 17959 | Aligned | - | - |
| X3 | Shadowburn costs a Soul Shard (Wowhead Forever tooltip 18871, read 2026-11-03), and the shard comes back at 5 points of Shadow and Flame. The engine keeps no shard count, which is right for this build. | nether.wowhead.com/forever/tooltip/spell/18871 | Aligned | - | - |
| X4 | Talent build: Destructive Reach 2 and Molten Skin 3 do nothing in the sim but are kept by the raid search's -keep list as unmodeled talents (range, damage taken), so they are not ruling-2 cases. | talent-search/warlock-destruction-raid.md | Aligned | - | - |

## Demons

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Succubus at raid 60: 90.4 swings (2.0 s), 5.9% miss, 6.5% dodge, 40% glancing (the 10 + 2 x 15 skill gap against a +3 boss; pets glance like players), no melee crit. Lash of Pain 11780: 50 flat, 0.429, 160 mana, 12 s category cooldown, 15.5 casts; Unholy Power applies (the tooltip names 18769), Improved Sayaad 18754 effect 0 misc 8 +30% at 3. Pet stats from the owner (`pet.go:208-290`, 17% AP and 10% spell damage, 8/17 of hit) as rounds 10 to 12 found. | 11780, 18754, 18769 | Aligned | - | - |
| P2 | Imp, Voidwalker, Felhunter: none is out in a published rotation (the Imp is sacrificed before the pull). Felguard 427733 is a Season of Discovery rune (no learn row): not in Forever. Inferno 1122 (1 h) and Ritual of Doom are not in any rotation; Bane of Doom's Doomguard summon needs the target to die from it. | trainables/warlock.json; SpellName | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | Published BiS (bis/warlock-*.json, all 36 entries): no item belongs to a set in sets.json, so no set bonus applies. Use and proc items: Frozen Heart of the Mountain (Chthonic Power 1249110, +29 Frost and Shadow, 15 s, 5 min, `item_effects.go:3732`), Kindling Stave (Ignition 1300781, 70 Fire, rate still open), Major Mana Potion, Demonic Rune. | ItemSet, sets.json, SpellEffect | Aligned | - | - |
| C2 | Racials of the published races: Gnome Expansive Mind 20591 (aura 178, +5% maximum mana, `racials.go:452`), Eureka! 1259821 (8 s, 2 min, -10% cost and +10% damage on the class mask, which leaves out Incinerate); Troll Berserking 20554 (10% cast speed, 15 s, 3 min; 1.48 casts because the fight can run past 180 s). | 20591, 1259821, 20554 | Aligned | - | - |
| C3 | Forever-only warlock spells in the casts: Bane of Agony and Doom (one bane slot, separate from curses), Curse of the Elements 1311680 (every magic school), Incinerate, Wrack (leveling only), Bane of Havoc (prerequisite only). The fork's damage tables pass against spellconst 70753 (`GOWORK=off go test -tags=with_db ./sim/warlock/...`, ok); the conformance copy is byte-identical to data/builds/1.60.1.70753/spellconst/warlock.json. | fork tests; cmp | Aligned | - | - |

## Before/after

10,000 iterations. Before: $R/before.tsv (int/r12 heads). After: wt/r13-warlock on the int/r13 engine. Only the Affliction raid entry moves.

| Spec, band, preset | Horde s1 | Horde % (s1 / s50001) | Alliance s1 | Alliance % (s1 / s50001) |
|---|---|---|---|---|
| Affliction 60 raid | 571.9 to 576.0 | +0.72 / +0.72 | 570.0 to 574.1 | +0.72 / +0.72 |

Every other warlock band and preset is unchanged (no engine or APL change).

## Files to regenerate

- data/builds/1.60.1.70753/addon-builds.json (`go run ./sim/cmd/addon-builds -repo-root .`): `TestCommittedAddonBuildsFollowTheGuides` fails until it is regenerated, as expected after a guide build change.
- Whatever reads addon-builds.json downstream (addon data / Data.lua) at the next pipeline run.
- Ladder goldens: not affected (the warlock ladders use the leveling build).

## Measure at refresh

- warlock-affliction raid entry: its talents string (bis json carries the old build), set DPS, weights and dps-per-point move by about +0.7%; picks should not move.

## For other lanes

- **Trainables:** data/builds/1.60.1.70753/trainables/warlock.json still lists Haunt (1293693, 1293694), Unstable Affliction (1242971, 1242972) and Dispel Magic (19476) as active warlock trainables, while round 12 (trainables A2, guides A5) found no learn row for Haunt or UA. Worth checking whether the pipeline's class-family fallback still lets them in.
- **Coordinator:** D4 (Demonology Improved Voidwalker to Improved Life Tap, +0.64 to +0.75%) is a ruling-2-style case waiting for a decision.

## Still open

- Affliction's leveling build (bare 30 to 60) carries Soul Harvest 2 and Improved Drains 3, which the bare APL never casts. The guide's leveling play uses drains, so this is not a ruling-2 case unless the owner wants the bare bands to sim the raid-style play.
- D4 (coordinator).
- From earlier rounds, unchanged: demon crit and sheets at 60 (no level-60 log), additive versus multiplied percent modifiers, Ignition's rate, Siphon Life tick crit (owner), shadow wands (#419), Destruction's held Conflagrate-first line (H2).

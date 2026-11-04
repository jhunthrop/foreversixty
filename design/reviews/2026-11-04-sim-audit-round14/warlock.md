# Sim-accuracy round 14: warlock (2026-11-03)

Specs: Affliction, Demonology, Destruction and their demons, bands 20 to 60, bare and raid presets. Client 1.60.1.70753 with the 2026-11-03 hotfix push layered. The push moved no warlock spell: spellconst/warlock.json is unchanged in 85204127f ("no other class moved"), and the push only touches warlock trainables, where the engine reads nothing. Site `wt/r14-warlock` from int/r14 72471fcb3; fork `lane/r14-warlock` from int/r14 5195a8889.

Method: round 13's report first, then a fresh pass that compares each talent effect's client class mask against the spells the engine applies it to. I listed every warlock talent effect (aura, misc op, base points, mask) and resolved each mask to spell names (SpellEffect plus SpellClassOptions for class set 5). I also checked every multi-rank talent's per-rank text against the engine table, read every spell file the three APLs cast, and checked cast-time and haste plumbing, the demons, pre-pulls, races and builds. spec-breakdown ran at 10,000 iterations, seeds 1 and 50001.

Totals: **3 CONFIRMED-BUGs, all origin old.** One is an engine bug: Immolate ignored haste. The other is a race pick: Orc should replace Troll for Affliction and Destruction. I also fixed a conformance note at the regcast lane's request; it is a documentation fix, not a sim bug. One build candidate goes to the data lane (Demonology bands 20 and 30). Nothing outside the warlock group moves.

## Affliction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Immolate's cast ignored haste** (all three specs; Affliction casts it too). `immolate.go:75-81` had its own `CastTime` hook, upstream SoD rune plumbing. It returned `DefaultCast.CastTime` without cast speed, so Berserking, Spellstone cast speed and haste gear never shortened Immolate. Every other warlock cast goes through the core default, `ApplyCastSpeedForSpell`. A Classic player would expect Berserking to speed up Immolate like any other cast. With 5 points of Bane, Immolate is a 1.5 s cast, the same as the unhasted global cooldown, so at raid 60 the fix moves nothing. It matters below Bane 5 (bands 20 to 50 for Affliction and Demonology, Troll). Test first: `TestImmolateCastTimeTakesCastSpeed` failed (2 s at double cast speed, want 1 s) and now passes. | `immolate.go`; core `spell.go:261` default; `cast.go:246` | CONFIRMED-BUG | old (1db608d39, upstream 2024; no round looked at cast-time hooks) | E c9e647908 |
| A2 | Talent masks against the engine: Shadow Mastery (18271 effect 0 damage: Shadow Bolt, Corruption, Death Coil, Shadowburn; effect 1 DoT: the banes, Corruption, the drains, Siphon Life, Wrack), Malevolence (1310949 aura 107 crit: every warlock Shadow spell plus Lash of Pain 0x2000, which the Succubus takes, `succubus.go`), Pandemic (427712 crit bonus: Corruption, banes, drains, Siphon Life, Wrack), Improved Corruption (17810: -2 s, +10% DoT), Improved Bane of Agony (18827), Wrack (aura 271 on 0x402, Corruption and Agony only) and Suppression (aura 54/55, hit 5). Each matches `talents.go`. Per-rank text for every multi-rank talent (talents/warlock.json) matches the engine tables. | SpellEffect, SpellClassOptions 70753 | Aligned | - | - |
| A3 | Shadow Vulnerability 17794: aura 270 school-mask damage taken from the caster, misc 32, 4% a rank, 12 s (DurationIndex 29), and no SpellAuraOptions row, so no charges (Classic had 4). Only the warlock's own Shadow damage takes it, so the demon and other casters do not. The proto `Debuffs.improved_shadow_bolt` raid option is read nowhere, and the raid preset does not set it, so no other spec gets it. Affliction raid (ISB 5) and Demonology raid (ISB 5) use it; Destruction takes 0. | 17794, 17793; `talents.go:782-841` | Aligned | - | - |
| A4 | Raid build 25520010135001050-0005200000000000000-0550004000000000. Every point is either used by the rotation or a keep: Grim Reach (range) and Demonic Aegis (Demon Armor, aura 108 on Demon Skin/Armor, defensive only). Amplify Curse is cast once before the pull and its 3 min cooldown is off the global (18288 RecoveryTime 180000, no StartRecoveryTime). | client rows; breakdown | Aligned | - | - |
| A5 | Life Tap: (424 + Spirit) × 1.2, as in rounds 2 to 13. The cached beta logs show only level-30 taps (Toxicity 274/275, Ruth 219/224, Dotts 256/274) and no Spirit readings, so they add nothing new. Report 2807 is a TBC Black Temple log (Life Tap 27222, 1,829 to 2,071 mana) and is not Forever data. | r3/casters/fl | Aligned | - | - |

## Demonology

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | Demon talents against the client: Unholy Power (18769 aura 107 misc 8 on Tamed Pet Passive: the demon's whole damage, `pet.go`), Fel Vitality (aura 107 on the pet passive plus aura 178 +15% mana), Improved Sayaad (aura 108 misc 8 on Lash of Pain's base, 10% a rank), Master Demonologist (four aura-4 dummies, 2% a rank), Soul Link 25228 (aura 79 +3% every school and aura 81 30% split, area pet aura), Burning Shadow 18789 (+15% Shadow; effect 1 the summon cancel), Touch of Fire 18791 (+15% Fire), Demonic Knowledge (100% of level) and Decimation (440870 cooldown -90% on Soul Fire; 440873 -40% cast and no reagent). All match. | SpellEffect 70753 | Aligned | - | - |
| D2 | **Round-13 hold, Improved Voidwalker 1 to Improved Sayaad 2**, re-measured: horde 639.3 to 639.8 (+0.08%) and 639.9 to 640.4 (+0.08%). Still within error and under 1%. Held. | spec-breakdown -talents | Aligned (held) | - | - |
| D3 | **Round-13 note, Soul Link at -2 s.** Moving it to -36 s, after the Succubus summon finishes at -38 s, measured horde 639.3 to 640.0 (+0.11%) and 639.9 to 640.0 (+0.02%), alliance 635.4 to 635.6 and 635.1 to 636.2 (+0.03 / +0.17%). Within error, under 1%. Held. Soul Link is a global and 20% of base mana (19028 SpellPower PowerCostPct 20, StartRecoveryTime 1500), as modelled. | spec-breakdown -rotation | Aligned (held) | - | - |
| D4 | **The leveling build's Improved Imp 3 did nothing; it is now Unholy Power 5** (coordinator ruling: the Shadow-build precedent, as round 13 adopted Raging Blows and Improved Bane of Agony). Every band runs the Succubus (`sim/request/request.go:631`), and the guide's leveling play uses the Voidwalker or the Succubus, never the Imp. Unholy Power had 3 open ranks on the same row. Guide `build:` 25522/233211310122000135/005 becomes 25522/203511310122000135/005. Row 0 still holds 10 points, so every gate holds; nothing the rotation casts is dropped; the raid build already takes Improved Imp 0. The walk follows: band 30 is 2035113101210001000 and still takes Soul Link. Horde, bare, seed 1 / 50001: band 20 +1.24 / +1.24%, 30 +1.10 / +1.10%, 40 +0.95 / +0.88%, 50 +0.72 / +0.72%, 60 +0.75 / +0.75% (347.0 to 349.6). The data lane measured +1.13 to +1.16% at 20 and +1.04 to +1.05% at 30 on both factions (data.md W2). Tests first: `TestDemonologyLevelingBuildTakesUnholyPowerOverImprovedImp` failed at every band and passes now; the band-30 want in `TestBandBuildsTakeTheCastTalentsThatFit` moved with it. The guide's talents sentence now says it. | spec-breakdown -talents; `sim/leveling` | CONFIRMED-BUG (dead points, coordinator-ruled) | old (68af1744d) | S c40362b6e |
| D5 | Raid build: Improved Health Funnel (Health Funnel is not registered: unmodeled, keep), Demonic Embrace and Demonic Aegis (defensive keeps), Master Summoner (summons; the guide's mid-fight resummon), Improved Voidwalker (D2). No other point is idle. | | Aligned | - | - |
| D6 | **Conformance note (regcast lane's request).** `sim/conformance/dispositions.go` and `warlock.golden.md` said "no rotation casts" Soul Link. Demonology's opener casts it. The note now says so; the conformance tests pass. This is a documentation fix, not a sim bug. | fork conformance | note fixed | - | E eb8a4a8b2 |

## Destruction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Destruction masks against the engine: Bane (17788 -0.5 s on Shadow Bolt, Immolate, Incinerate; -2 s Soul Fire), Cataclysm (17778 cost: Immolate, Incinerate, Conflagrate, Shadowburn, Searing Pain, Soul Fire, Shadow Bolt, Bane of Havoc, Rain of Fire, Hellfire), Ruin (17959 crit bonus: the same Destruction list, not Corruption) and Agonizing Flames (17927: +10% direct on the Destruction nukes, +10% DoT on Immolate and Rain of Fire, +10% crit on Searing Pain). The engine's `WarlockFlagDestruction` set (Shadow Bolt, Immolate, Incinerate, Conflagrate, Shadowburn, Searing Pain, Soul Fire, Rain of Fire, Bane of Havoc) covers exactly these. Fire and Brimstone gives Conflagrate 38.3% crit per landed hit (13.3% + 25%). | 70753 masks; breakdown | Aligned | - | - |
| X2 | Round-13 hold H2 (Conflagrate above Immolate's refresh): +0.12 to +0.23% in round 13. No engine change touches it, so it was not re-measured; still held. | round 13 H2 | Aligned (held) | - | - |
| X3 | Raid build: every point is used or a keep (Destructive Reach range, Molten Skin damage taken, Bane of Havoc as Incinerate's prerequisite). | | Aligned | - | - |

## Demons

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Succubus at raid 60: 2.0 s swings from behind (no parry), 6.5% dodge, 40% glancing (15 skill under a +3 boss), about 5.9% miss with 8/17 of the warlock's hit. Lash of Pain misses like the warlock (applyOwnerSpellHit). Lash crit is 11.4% per landed hit for Affliction (6.4% base + Malevolence 5) and 6.1% for the others. No player buffs (round 2/3 ruling). | breakdown; `pet.go` | Aligned | - | - |
| P2 | Imp, Voidwalker and Felhunter: none is out in any published rotation. Felguard is a Season of Discovery rune with no learn row. Haunt and Unstable Affliction have no SkillLineAbility row in 70753 (round 13 trainables X5), so they are not learnable. | SkillLineAbility | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **Horde race: Orc beats the published Troll beyond 1% for Affliction and Destruction.** Raid 60, published sets, lane engine: Affliction 571.9 to 578.3 and 572.0 to 579.1 (+1.12 / +1.24%); Destruction 636.9 to 643.5 and 636.7 to 643.9 (+1.04 / +1.13%); Demonology 639.3 to 644.6 and 639.9 to 645.0 (+0.83 / +0.80%). Error is about ±0.07%. The gain comes from Command (+5% demon damage) and Blood Fury (10% spell power, 15 s, 2 min). Undead also beats Troll (+1.0 / +1.1% Affliction, +0.88 / +0.80% Destruction, +0.61 / +0.42% Demonology), but less than Orc. Under the rule (race picks within 1% stay), Affliction and Destruction move to Orc and Demonology keeps Troll. `recommendedRaces` is now [gnome, orc] in affliction.md and destruction.md, and the race paragraphs follow without sim talk. On the Alliance side Human is +0.28 to +0.63% over Gnome, so Gnome stays. `_prose.test.ts` and the guide content tests pass (526), as do the `leveling-bis` tests. | spec-breakdown -race | CONFIRMED-BUG (stale pick) | old (round 9 saw "about 1.0%" and held it) | S 7014d2c48 |
| C2 | Published BiS (bis/warlock-*.json): no item belongs to a set, so there are no set bonuses. The engine's item effects for Rune of the Guard Captain are Authority 1318000 and 1287704 (attack power, nothing for a caster). The files are from engine 8baffe5fc (round 5) and still carry Torch of Light, which round 6 dropped from the raid sets. Destruction's set has about 5% less crit than Demonology's (Immolate crit 13.3% against 18.3% on near-identical buffs), though Ruin values crit most. That is for the refresh to settle; see Measure at refresh. | bis files; breakdown | Measure at refresh | - | - |
| C3 | Haste plumbing: Immolate was the only warlock spell with a custom cast-time hook (A1). Lash of Pain is IgnoreHaste because it is an instant. DoT ticks are not hasted. Berserking is MultiplyCastSpeed 1.1 for 15 s on a 3 min cooldown. | `grep CastTime/IgnoreHaste sim/warlock` | Aligned | - | - |
| C4 | Logs: `/v1/reports/recent` has no new warlock report, and no level-60 log exists. Demon crit at 60 stays open. WCL is blocked (gcloud login expired), so I did not try it. | API, 2026-11-03 | OPEN | - | - |

## Before/after

10,000 iterations. Before: $R/before.tsv (raid 60) and $R/sb-before (bare). After: lane engine c9e647908 on this branch.

| Spec, band, preset | Horde s1 | Horde % (s1 / s50001) | Alliance % (s1 / s50001) |
|---|---|---|---|
| Affliction 60 raid | 572.0 to 571.9 | -0.02 / +0.04 | 0.00 / 0.00 |
| Demonology 60 raid | 639.3 to 639.3 | 0.00 / 0.00 | 0.00 / 0.00 |
| Destruction 60 raid | 636.8 to 636.9 | +0.02 / -0.03 | 0.00 / 0.00 |
| Affliction bare 20 / 30 | 54.5 to 54.6 / 89.9 to 90.0 | +0.18 / +0.18, +0.11 / +0.11 | - |
| Demonology bare 20 / 30 / 40 | 48.3 to 48.4 / 82.0 to 82.1 / 147.6 to 147.7 | +0.21 / +0.21, +0.12 / +0.12, +0.07 / +0.14 | - |
| every other bare band | | within ±0.1 (noise) | - |

At raid 60 Immolate's 1.5 s cast already equals the global cooldown, so haste cannot shorten the cycle. The Orc pick (C1) reaches the published numbers at the refresh: about +1.1 to +1.2% on the Horde entries for Affliction and Destruction.

## Files to regenerate

- data/builds/1.60.1.70753/addon-builds.json (`go run ./sim/cmd/addon-builds -repo-root .`): `TestCommittedAddonBuildsFollowTheGuides` fails until it is regenerated, because the Demonology leveling build changed (D4). Data.lua and addon-data follow at the next pipeline run.
- Demonology BiS bands 20 to 60 (bare): the talents change; measure at the refresh.

- BiS / tiers / weights for warlock-affliction and warlock-destruction, Horde: the published race changes from Troll to Orc (guide `recommendedRaces`). This needs the ranker, so it waits for the refresh.
- Ladder goldens (sim/request TestRotationLadder) for the warlocks: may move slightly where the ladder character is a Troll below Bane 5 (A1). Not run.
- Web vitest suites that import `web/src/data/generated/*.json` (guide-bis, class-dps, gear-extras, rail-stats, GuideStatTable) fail in a fresh worktree for want of the generated fixtures, before and after my edit. I did not run the fixture sync.

## Measure at refresh

- Horde Affliction and Destruction on Orc: set DPS, picks and rank rows (the Troll-keyed rank rows no longer apply to them).
- Destruction's crit gear against Demonology's (C2), and Torch of Light in the old warlock sets.
- Haste weights for Affliction and Demonology below 60 (A1; small).

## For other lanes

- None. Demonology's Improved Imp row (D4) was adopted in this lane at the coordinator's request.

## Still open

- Demon crit and sheets at 60, additive versus multiplied percent modifiers (for example Improved Corruption is multiplied while Shadow Mastery is additive), Ignition's rate, Siphon Life tick crit (owner), shadow wands (#419): unchanged, with no level-60 log.
- Round 13's pending in-game checks are not reopened.

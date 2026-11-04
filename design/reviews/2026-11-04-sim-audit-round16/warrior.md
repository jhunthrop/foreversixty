# Round 16, lane warrior (Arms, Fury, Protection)

Site `wt/r16-warrior` on int/r16 c62d77ab8; engine `lane/r16-warrior` on fork int/r16 9f7f1f6e4 (no engine commit this round). Client 1.60.1.70755 with the late hotfix re-layer. Client rows: the layered 70755 raw tables (`scratchpad/c755/fetch/data/builds/1.60.1.70755/raw/`), `talents/warrior.json`, Wowhead Forever tooltips.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| 1 | **The talent ladder took a point back.** Protection Warrior holds Concussion Blow and Shield Slam at 40 (`...001201`), but at 41 the build is `...001310`: Shield Slam is gone and Bastion appears. No player can do that without a respec. Cause: `LadderTalentString` walked each level on its own. Its own-tree spine (the deepest talent and its chain, `ownTreeSpine`) is taken only when it keeps every talent the plain top-down walk would hold (`keepsEveryTalent`). At 41 the plain walk reaches Bastion's first rank, the spine trial would leave Bastion at 0, so the spine is refused and Shield Slam drops out. At 42 the spine is back. Balance druid's raid build does the same at 42 (Moonkin tree, `...01310` to `...01301`). Neither `ladder-first.json` nor `ladder-order.json` was involved. Fix: `LadderTalentString` now climbs the ladder one point a level from 10 (`climbLadder`): each level keeps the level below's ranks and adds the first legal point the walk wants, or, where the walk moved a rank, the first legal point toward the guide build. That is the algorithm `LadderPointOrder` already used for the addon, now shared, so the addon order and the bands come from one climb. Where the walk only adds points (every checkpoint level of every spec), the build is identical: a diagnostic over all 28 specs, leveling and raid builds, levels 10 to 60, found only three level/spec pairs that change (Protection Warrior leveling and raid at 41, Balance raid at 42). No checkpoint level (10, 20, 25, 30, 35, 38, 40, 45, 50, 55, 60) moves, so no band, ladder golden, dps-per-point or BiS number moves, and the addon point order is the same as before. The web port (`web/src/lib/guides/leveling-builds.ts`) climbs the same way. Tests: `TestEveryLadderOnlyAddsPoints` (sim/leveling, all 28 specs with ForSpec) and "every guide ladder only adds points" (web vitest) both failed on the old code at warrior-protection level 41 and pass now. `TestPointOrderNeverRemovesAPointWhereTheLadderMovesOne` now looks for a move in the one-level walk (`walkRanks`), since the ladder itself no longer moves. | sim/leveling/talents.go `LadderTalentString`, `spendOwnTree`; point_order.go | CONFIRMED-BUG | old (78efc0d3e, 2026-10-07, the spine; no round checked consecutive levels outside the point order) | S ad9c2ed9c |
| 2a | Arms raid build `30303213132215201/05050005012/1` (round 15): legal on 70755, casts Reckless Swings (11.3 a fight, the cap is 12) and Spearing Strike (6.4 a fight in Battle Stance, the 75% of the fight before the Berserker swap). Spearing Strike: 1310222 e0 effect 121 (normalized weapon damage), e1 75%, e2 1 (twice on Giants and Dragonkin); engine `spearing_strike.go` uses `MHNormalizedWeaponDamage` x 0.75 (x1.5 on those types), 15 rage, 20 s, Battle Stance only. | SpellEffect 1339572-3, 1339768; breakdown horde seed 1 | Aligned (regression clear) | - | - |
| 2b | Fury Improved Berserker Rage: Berserker Rage 5.8 casts a fight (cap 6), 10 rage each (57.8 a fight), 1.5 s global (SpellCooldowns 55298), Berserker Stance only (SpellShapeshift 40570 mask 262144), level 30 (SpellLevels 103013). Rank 2 = 10 rage (20500 e0, `talents/warrior.json`). | client rows | Aligned | - | - |
| 2c | Fury band 30 takes Death Wish (`TestBandBuildsTakeTheCastTalentsThatFit` green) and casts it 1.08 times a fight at band 30 bare. Reckless Swings comes at 31, Improved Berserker Rage at 34-35 (Berserker Rage is unusable before the 36 stance move, but the points are first spent at 34). | ladder order dump | Aligned | - | - |
| 2d | **Round 15's guide text never landed.** Guide round 10 (WARRIOR-R10-1 to R10-8) found the Arms and Fury prose still describing the old builds (Deflection, Blood Craze, Lingering Rage, Piercing Howl, Berserker Rage held on every fight) and three new APL lines with no card note. Written in this lane, guide voice: arms.md :71, :74, :79, :87; fury.md :67, :78, :86, :92, :128; APL notes for Arms Reckless Swings and Spearing Strike, Fury Berserker Rage, and the Arms Bloodrage note (R10-3: the raid build takes one point of Improved Bloodrage, 25%, not 50%; 12301 "by 25%"). `_prose.test.ts` and `test_apl.py` pass. wow-player review on the rendered pages: see "Rendered review". | guide round 10 report (scratchpad gr10/audit), round 15 lane "For the guide lane" | CONFIRMED (text) | R15 9ba13b069, c328e4192 (frontmatter changed without the prose), guides-handoff | S 4778ff1f0, S 24035b2b5, S (this commit) |

## Arms

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Weapon damage types against the client: Mortal Strike (21553 e1 effect 121 +160), Overpower (11585 effect 121 +35), Whirlwind (1680 effect 121), Spearing Strike (121) are normalized; Slam (11605 effect 17 +87), Heroic Strike (11567 effect 17 +138), Cleave (20569 effect 17 +50) are not. The engine matches every one (`*WeaponDamage(` calls in sim/warrior). | SpellEffect rows | Aligned | - | - |
| A2 | Raid breakdown (horde Orc, seed 1, 3000 iterations): Overpower crit 77% (Improved Overpower 50%), Mortal Strike 23.8 casts, Slam 10.5, Execute 14.7, Rend kept up (7.6 casts, ticks crit per the bleed ruling), Deep Wounds ticks 34 a fight, Hamstring only under Windfury, one Berserker swap and one Recklessness, Blood Fury twice, Mighty Rage Potion twice. Bloodrage 3.5 casts. | spec-breakdown | Aligned | - | - |
| A3 | Weaponmaster 1290261: e0 5 (axe/polearm crit), e1 15 (mace/staff armour), e2 5 (sword extra attack); engine 1% / 3% / 1% a point. Deep Wounds 60% of average weapon damage over 4 ticks of 3 s, refreshed by a new crit. | SpellEffect 1309067-9; deep_wounds.go | Aligned | - | - |
| A4 | Races, raid 60, guide raid build, 3000 iterations seed 1: Alliance Dwarf 824.2, Night Elf 821.7, Human 820.2, Gnome 819.3; Horde Orc 826.8, Undead 824.6, Troll 822.5, Tauren 818.2. Every gap is inside 1%, so the published races and the guide's Night Elf / Undead stay (race rule). | races.sh | Aligned | - | - |
| A5 | Leveling build: Deflection 3 (11-13) and Iron Will 5 (56-60) are leveling and world choices, not dead points (you are hit while leveling). The raid build drops them. | | Aligned | - | - |

## Fury

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | Off-hand hit: the BiS gear has 10.2% hit; main-hand white misses 18.7% (28% dual-wield table minus 10.2 minus skill), off-hand 8.6% with Furious Precision 3 (1323963 rank 3 = 10%, off-hand only). Matches the client text and Kaivax's 2026-10-02 post (Dual Wield Specialization no longer gives hit; Furious Precision is the off-hand hit). | hit_to_cap; breakdown; https://us.forums.blizzard.com/en/wow/t/2369360 (fetched 2026-11-04) | Aligned | - | - |
| F2 | Raid breakdown (horde Troll, seed 1): Bloodthirst 22.1 casts, Whirlwind 16.4 (8 s with Improved Whirlwind), Reckless Swings 10.6 (20 rage a cast), Death Wish 1.09, Recklessness 1, Bloodrage 29.7 rage a cast (Improved Bloodrage 2), Berserker Rage 10 a cast, Flurry up 84%. | spec-breakdown | Aligned | - | - |
| F3 | Races, raid 60, 3000 iterations seed 1: Human 880.6, Night Elf 880.1, Gnome 878.0, Dwarf 876.5; Undead 878.1, Orc 875.6, Troll 872.8, Tauren 866.8. Inside 1% on each side (Troll, the published Horde BiS race, is 0.6% under Undead); the guide's Night Elf / Undead stand. | races.sh | Aligned | - | - |
| F4 | Kaivax 2026-10-02 warrior post, line by line: crit rage 100% (rage.go, round 15 W2-2); Unbridled Wrath no double rage for two-handers (engine: 1 rage, any weapon); Lingering Rage, Furious Precision, Gore Drinker, Improved Berserker Rage row, Berserker Rage at 30, Bloodthirst 45% AP, Whirlwind both weapons without a talent, Blood Craze not from Bloodthirst. All followed. | forum post | Aligned | - | - |
| F5 | Heroic Strike and Cleave register with `ProcMaskMeleeMHSpecial | ProcMaskMeleeMHAuto` (upstream wowsims, 543752637), so white-swing procs (Unbridled Wrath, Reckless Swings charges, Flurry charges) also fire on a Heroic Strike hit. Unbridled Wrath's rage in the Fury breakdown (97 a fight) equals 0.6 x (white hits + Heroic Strike hits). The client gives Unbridled Wrath ProcTypeMask 4 (done melee auto attack), and whether a next-swing ability counts as one is a server rule no client row settles. Web search found no Classic or Forever test either way. | heroic_strike_cleave.go:84; SpellAuraOptions 124699 | OPEN (a Forever log with Heroic Strike hits and Unbridled Wrath energize events settles it) | - | - |

## Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Tank fight (alliance Dwarf, seed 1, 3000 iterations): 1097 threat/s, Shield Slam 30.2 (every 6 s), Revenge 30.0, Shield Block 78% up, Thunder Clap 7.2 (refresh under 3 s), Demoralizing Shout 5.1, Heroic Strike on 56 of 85 swings, chance of death 0%. Boss table: miss 6.9, dodge 12.4, parry 11.8, block 62.1, crush 1.8. | tank breakdown | Aligned | - | - |
| P2 | Dead points: none. Improved Rend 3 is Deep Wounds' prerequisite, Concussion Blow is Shield Slam's, Deflection 5 is parry. Last Stand and Shield Wall stay health-gated (O-items). | talents/warrior.json prereqs | Aligned | - | - |
| P3 | WARRIOR-13 (Improved Shield Wall) stays with the owner (round 15 2c). | | OPEN (owner) | guides-handoff | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Stance swaps: SpellCooldowns 52923/53448/53578 are a 1 s category cooldown with no global; engine `stances.go` uses a 1 s shared timer and no GCD. Tactical Mastery 10 + Improved Tactical Mastery 5 a rank. | client rows | Aligned | - | - |
| X2 | Battle Shout rank 6 (11551): 111 + 0.6 a level from 52, 20 yd, 10 rage; `core.BattleShoutRankTable` matches. Bloodrage 2687: 10 rage + 1 a second for 10 s, 1 min, no global; Recklessness 1719: 100% crit, +20% damage taken, 30 min, level 50; Death Wish 12328: +20% physical, fear immunity, +5% taken, 10 rage, 3 min; Hamstring 7373 rank 3: 45, 10 rage. | client rows | Aligned | - | - |
| X3 | Comment nit: `berserker_rage.go:9` says the client still has Berserker Rage at 32; 70755's SpellLevels 103013 says 30. No behaviour change, left for a later engine commit. | | note | - | - |
| X4 | The BiS trinkets are refresh-owned: Arms Alliance carries Thunderbrew's Boot Flask (8 Spirit and a 375-damage fire breath on a 5 min cooldown, modelled in `client_procs.go`), and both specs carry Frozen Heart of the Mountain for its 0.9% hit (its use is a Frost/Shadow spell damage buff, nothing for a warrior). | items/warrior.json; Wowhead tooltip 249469 | Measure at refresh | - | - |

## Before/after (raid 60, 10000 iterations, guide builds via `rt`, frozen BiS)

No commit moves a published number. Every warrior value at both seeds and factions is identical to $R/before.tsv:

| Spec | alliance s1 / s50001 | horde s1 / s50001 |
|---|---|---|
| warrior-arms | 820.4 / 820.4 (0.00%) | 826.4 / 826.5 (0.00%) |
| warrior-fury | 879.4 / 878.9 (0.00%) | 873.3 / 873.2 (0.00%) |
| warrior-protection (threat/s) | 1097.0 / 1096.3 (0.00%) | 1088.0 / 1088.0 (0.00%) |

The ladder fix changes the build only at Protection Warrior level 41 (both builds) and Balance's raid build at 42, levels no sim runs. **The ladder fix is shared code (sim/leveling and the web port) and touches every spec's ladder, but moves no checkpoint for any spec.**

## Files to regenerate

- None required by this lane's code. addon-builds.json is unchanged (the climb is the order it already wrote).
- The new and corrected APL notes reach the guide's rotation card one-liners only through the nightly data (the card reads `line.condition` from the published rotation data), so the next addon-data / BiS refresh carries them: warrior-arms (Bloodrage, Reckless Swings, Spearing Strike), warrior-fury (Berserker Rage). Notes are display-only, so `apl-check` needs no sync.

## Measure at refresh

- Arms and Fury race order on the round 15 builds (all inside 1% here; the refresh's race pick stays).
- Arms Alliance trinket2 (Thunderbrew's Boot Flask) and Frozen Heart picks.
- Carried from round 15: Arms and Fury picks and weights on the round 15 builds; the BiS bands still carry the old talent strings in the frozen copy.

## For other lanes

- Ranker/core: the `expertise_to_cap` block (round 15 2d) stands.
- Web (rotation card): a guide's rotation card shows the raid list under the level 60 build, so raid-only talent lines (Arms Reckless Swings, Spearing Strike) appear for a build without them; the notes now start "Raid build,".
- Core (all melee): F5, Heroic Strike/Cleave proc mask counting as a white swing, OPEN pending a log.

## For the guide lane

- Done in this lane (2d). If the guide lane carries the same round-10 rows, take this lane's text: arms.md :71, :74, :79, :87; fury.md :67, :78, :86, :92, :128.

## Rendered review

wow-player on the rendered Arms and Fury guide pages (astro dev, 1440 px full-page captures with both trees drawn): first pass Arms NOT SHIP (Reckless Swings promised four swings on a two-hander and no health cost; dropped talents called useless; Spearing Strike placement unclear; trees still loading in the capture), Fury SHIP with SHOULDs. All MUSTs fixed in S 24035b2b5. Second pass: **Arms SHIP, Fury SHIP.** The Arms leveling sentence (Improved Overpower at 22, Bloodthrill at 33) stays: it matches the shipped point order; the band tables show each band's first level. Reviewer notes passed on: the shared rotation card under the level 60 build lists the raid list's Reckless Swings and Spearing Strike rows (their notes now start "Raid build," so the nightly one-liners say so); the Fury BiS Ironwood Blade / Persuader pair is for the BiS refresh to confirm.

## Still open

- F5 (Heroic Strike as a white swing for procs): needs a Forever log.
- WARRIOR-13 (owner).
- Warcraft Logs blocked (gcloud login), not worked around.

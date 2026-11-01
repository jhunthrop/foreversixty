# Sim-accuracy round 4: warlock (2026-11-01)

Specs: Affliction, Demonology, Destruction. Client 1.60.1.70707. Pin 640da215f. Lane engine `lane/r4-warlock` (fork), site branch `wt/r4-warlock`.

Method: spec-breakdown, raid preset, band 60, horde troll and alliance gnome, 10000 iterations at seeds 1 and 50001. "Before" is the pin with origin/main's rotations. "After" is the lane engine with this branch. The client rows come from wago SpellEffect, SpellMisc, SpellClassOptions, SpellLevels and SpellPower for build 1.60.1.70707 (downloaded 2026-11-01). Talent searches use `sim/cmd/talent-search -preset raid -faction horde -keep <raid-keeps.json>` on the lane engine. The reports are in `warlock-ts/`; their header prints the pin sha, but GOWORK pointed them at the lane engine.

Logs: no level-60 Forever log exists. The beta capped at 30 and ended around 2026-10-23. foreverlogs.gg's API now refuses access, and our `/v1/reports/recent` holds no new warlock report. The checks use round 3's event-level timelines (24 reports, 2740 to 2811, cached in `r3/casters/fl`) and the ability and resource summaries in `fl/wl`.

## Pet scaling (assigned item 1)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Demons took Classic's 56.5% of spell damage as attack power and 15% as spell damage. Forever's figures are 17% and 10%.** The share is of the warlock's better Shadow or Fire spell damage. Resistances are 100%, not TBC's 40%. | (1) Client: `WARLOCK_PET_BONUS` in the Forever UI (Gethe/wow-ui-source branch `forever`, 1.60.1.70532, `Blizzard_UIPanels_Game/Camelot/PaperDollFrame.lua:54-63`) gives `PET_BONUS_SPELLDMG_TO_AP 0.17`, `_SPELLDMG_TO_SPELLDMG 0.1`, `_ARMOR 0.35` and `_RES 1.00`, and sets `_STAM` and `_INT` to 0 ("Pets no longer have primary attributes"). (2) wowsims/forever `sim/warlock/pets.go` (c3d9ea951, 2026-10-14) reads the beta tooltip "Increases your pet's Attack Power by 5 and Spell Damage by 3" at 30 spell damage. (3) Our own beta events: each warlock's spell power is solved from its own Corruption, Shadow Bolt, Immolate and Siphon Life hits. From that, Lash of Pain predicts 22.0 / 29.3 at 10% and 25.0 / 32.9 at 15%, against 22.05 logged (Toxicity, rank 1, report 2762) and 29.4 logged (Daddy, rank 2, report 2791). Firebolt rank 3 predicts 18.3 at 10% against 17.2 logged (Voidmire, report 2745). Test `TestDemonInheritsForeversShares`. | CONFIRMED-BUG | E 0ff39c41d |
| X2 | **Demon melee took the warlock's whole hit, Suppression included.** Forever gives a demon's spells the warlock's spell hit and its melee 8/17 of it. Lash of Pain and Firebolt now add the remainder at each cast. Demon melee misses rise from 2.4% to 5.9%. | forever-bugs #394, AidanZMoon for Blizzard (2026-10-20): "Warlocks will give spell hit equal to their spell hit, and melee hit equal to 8/17ths of their spell hit." wowsims/forever 9ce48b6ab models the same rule. No pet sheet has been read, so this is one primary source. Test `TestDemonSpellsTakeTheWarlocksWholeHitAndMeleeEightSeventeenths`. | CONFIRMED-BUG (single primary source, Blizzard) | E b35bda572 |
| X3 | Pet crit. Blizzard intends a demon to "mirror the spell critical strike chance of their master exactly" as both its melee and spell crit, with no crit of its own. That is not live: "Warlock pets currently only scale with their master's +% critical strike gear and have erroneous base crit." A Fierywind log at level 30 shows Succubus melee critting 0.70% of 1,718 swings and Lash 6.3% of 205. The engine keeps the demon's own crit, with melee at 2.8%. Blizzard's intended model would add +2.6% Affliction, +2.4% Demonology and +1.7% Destruction (lane run, horde, both seeds). | forever-bugs #331 (2026-10-13/14, issue still open); this lane's sensitivity runs `runs/crit-*` | OPEN. A launch-build log settles it: demon melee crit near 0.7% means it is still live, and near the warlock's spell crit means the change shipped. | - |
| X4 | Haste inheritance. The client has the channels (416189 aura 319/65). No source gives an amount, and this engine carries haste as speed multipliers. | client 416189 only | OPEN, blocked on level-60 logs: a demon's swing timer under Berserking or a haste potion on the warlock | - |
| X5 | Lash of Pain's coefficient reads the demon's spell damage, not the warlock's (round 3 P5). | Client 11780 has coefficient 0.429 and no owner flag. wowsims/forever fits "(16 + 0.429 x Spell Damage) x 1.10" to beta logs. The X1 log fits above use the demon's 10%. Kargoz's "42.9% of your spell power" fits none of these. | Aligned (two sources now) | - |
| X6 | Lash of Pain took no crit from Malevolence. Malevolence's class mask carries Lash's bit (0x2000), the same way Improved Sayaad's does. | client 1310949 effect 0 aura 107, mask0 550027 (0x8648B); class-mask map from SpellClassOptions set 5 | CONFIRMED-BUG (client only; +0.03% Affliction) | E 6f5d65cb6 |

The Stamina (75%) and Intellect (30%) shares stay. The client sheet says pets have no primary attributes, so these stand in for the health and mana the server grants. wowsims/forever reads health as 7.5 a point of the warlock's Stamina, which is what 75% gives, and reads mana regen as a share of Intellect. Neither moves DPS while the demon never runs out of mana.

## Affliction

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | Raid search after the fixes. The only winners beyond error drop Grim Reach, an unmodeled guide talent. The best variant that keeps it (1 Improved Drains to Improved Bane of Agony) gains +4.1 ± 1.4 (+0.7%), below the 1% bar. The build holds. Unholy Power keeps its 5 points. | `warlock-ts/warlock-affliction-raid.md` | Aligned | - |
| A2 | Fresh check of every cast spell against 70707. Shadow Bolt 11661: 251 + 1.8/level, coefficient 0.857, variance 0.108, 3 s cast. Corruption 11672: 57 a tick, 6 ticks at 3 s, 0.2 a tick. Bane of Agony 11713: 46 a tick, 12 ticks at 2 s, 0.133 a tick, with the ramp. Immolate 11668: 146 + 2.3/level at 0.2, then 52 a tick for 5 ticks at 0.13. Siphon Life 18881: 41 a tick, 10 ticks, 0.05. Life Tap 11689: 420 + 1/level + Spirit. The engine matches every row (`shadowbolt.go:34`, `corruption.go:29`, `curses.go:30-33`, `immolate.go:45-46`, `siphon_life.go:38`, `lifetap.go:28`). | client rows above; Wowhead Forever tooltips 6223, 6217, 11672 and 11713 (read 2026-11-01) give the same amounts | Aligned | - |
| A3 | Against the beta events: for each warlock, the spell power implied by Corruption, Shadow Bolt, Immolate, Searing Pain and Siphon Life agrees within the spread of its talents. Ruth: 150 / 152 / 158 to 165 / 152. Toxicity: 135 to 160 across all four, and exactly 135 once Improved Corruption's 10% is counted. So the coefficients' ratios hold. Bane of Agony's top tick (Dotts, rank 2, 34) matches 1.5 x base plus spell power, so the ramp holds. Life Tap matches base + Spirit within 3% in 19 of 27 logs, 15 of them to within 1%. The rest are Improved Life Tap (x1.2) or unspent talent points that shift the level estimate. | `ev.py` and `spsolve.py` over `r3/casters/fl`; `fl/wl/*.res.json` | Aligned | - |
| A4 | The ability summaries' "hits" undercount events (Searing Pain: 34 casts, 16 "hits"). Their per-hit averages run 1.5 to 3 times high, and the roster's gear stats miss most spell power. Only event-level data can check amounts. | foreverlogs.gg `character_abilities` against the timelines | Note (method) | - |
| A5 | Pandemic, Shadow Mastery, Nightfall and Improved Corruption masks re-read: they match round 2. Siphon Life tick crit stays held for the owner. | client 427712, 18271, 18094, 17810 | Aligned | - |

## Demonology

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| D1 | **Bane of Doom was recast with under a minute left, so it never fired.** Its single hit lands at 60 s. The fight cast 3.33 Dooms for 2.13 hits. Doom now needs 60 s remaining, and Bane of Agony covers the last minute. The guides already advised Agony for short fights; they now describe this. | client 603: one 60000 ms period, base 1742, coefficient 4.0; spec-breakdown cast and hit counts; `data/tests/test_apl.py` and `sim/request` (ladder skipped) pass | CONFIRMED-BUG (rotation) | S d8d8127f3, E a9274ae96 |
| D2 | Demonic Brand (assigned item 2). No guide build takes it, at band 60 or in raid: the Demonology raid string has 0 at Brand's slot (tier 3, column 3). Neither does any published BiS talent string. It stays unmodeled. wowsims/forever does sim Brand builds, so revisit it if a Forever guide adopts one. | talents json 1.60.1.70707; guide frontmatter; bis json | Aligned (moot) | - |
| D3 | Raid search after the pet fixes: the guide build is still the best (661.0 ± 1.4). | `warlock-ts/warlock-demonology-raid.md` | Aligned | - |
| D4 | Master Demonologist (2% a rank, client base 10 at 5/5), Soul Link (3% for both, 20% base-mana cast), Demonic Sacrifice (Imp 18789: aura 79, 15% Shadow, 2 h), Demonic Pact, Demonic Knowledge (100% of level at 3/3, warlock and demon) and Decimation (90% cooldown cut, 40% cast cut, 6% damage below 35%) all match the client and talent text. Improved Sayaad is ALL_EFFECTS (misc 8, +30%), so the engine's base-only scaling is right. | client 23785, 19028, 18789, 425464, 412732, 440870, 18754; `talents.go:557-740`, `soul_fire.go:102-160`, `succubus.go` | Aligned | - |

## Destruction

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | Bane of Doom in the last minute: same as D1. | as D1 | CONFIRMED-BUG (rotation) | S d8d8127f3, E a9274ae96 |
| S2 | Incinerate 1293813: 217 + 1.4/level, 0.714, variance 0.15, 2.5 s, +25% on Immolate. Conflagrate 18932: 282 + 1.3/level, 0.429, 10 s. Shadowburn 18871: 266 + 1.8/level, 0.429, 15 s. Ruin, Cataclysm and Agonizing Flames masks equal the engine's Destruction flag set (Shadow Bolt, Immolate, Rain of Fire, Soul Fire, Shadowburn, Searing Pain, Conflagrate, Incinerate). Agonizing Flames' periodic half covers Immolate and Rain of Fire. | client rows; Wowhead Forever 1293813 "213 to 246" (rendered at its max level, 69) | Aligned | - |
| S3 | Raid search: no variant beats the guide beyond error (+0.6 ± 1.7). | `warlock-ts/warlock-destruction-raid.md` | Aligned | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| C1 | **Shadow Vulnerability (Improved Shadow Bolt) raised every attacker's Shadow damage, the demon's Lash of Pain included.** It is aura 270 (school-mask damage taken from the caster) and reads "from your attacks". It is now a per-caster multiplier, the hook Wrack uses. Affliction and Demonology -0.12%. | client 17794 effect 0: aura 270, misc 32, base 4 a rank, 12 s, no charges; test `TestImprovedShadowBoltReachesOnlyTheWarlocksOwnShadowDamage` | CONFIRMED-BUG (client only) | E 501786aa8 |
| C2 | **Eureka! (Gnome, the published alliance race) is modeled as +10% to every magic spell for 3 costed casts.** The warlock row 1259821 limits its effects by class mask. The damage effect (misc 0) covers Shadow Bolt, Immolate's hit, Rain of Fire, Soul Fire, Shadowburn, Searing Pain, Death Coil and Conflagrate. The periodic effect (misc 22) covers Drain Life, Drain Soul and Wrack, and the cost effect covers a similar set. Corruption, Agony, Doom and Incinerate get nothing and should not spend a charge. My estimate is that the engine overstates gnome warlocks by about 0.5%: the opener's Corruption and Doom take the 10%. | client 1259821 effects 0-2 masks; `core/racials.go:459-505`; the racial review of 2026-10-07 already lists "the charge is still spent on any costed cast" | CONFIRMED (not fixed: a core change for every class's variant, outside this lane) | - |
| C3 | Racials. Troll Berserking (20554: 10%, 10 s, 3 min) matches. The Gnome entry is C2. Orc, Undead and Human are not published for warlocks. | client 20554; `core/racials.go:538-620` | Aligned | - |
| C4 | Set bonuses on the published BiS. Only the horde Affliction and Demonology sets carry one: Vestments of the Devout 3/8 (2: +8 resistances, 3: +18 spell damage), modeled as client stats. | `data/builds/1.60.1.70707/bis/warlock-*.json`, `sets.json`; `core/client_sets_gen.go:24` | Aligned | - |
| C5 | Round 3 side effects: demons still refuse player buffs (the auras table lists the player's only), Lash is 50 at rank 6, Firebolt uses the client range, and Unholy Power is pet-wide. | baseline runs `runs/base-*` | Aligned | - |

## Before/after (relative, raid preset, band 60, 10000 iterations)

| Spec | Step | Horde seed 1 | Horde seed 50001 |
|---|---|---|---|
| Affliction | shares + hit (E 0ff39c41d, b35bda572) | -3.01% | -3.01% |
| Affliction | + Shadow Vulnerability own-only (E 501786aa8) | -0.12% | -0.10% |
| Affliction | + Malevolence on Lash (E 6f5d65cb6) | +0.03% | +0.03% |
| Demonology | shares + hit / Shadow Vulnerability / Doom rotation (S d8d8127f3) | -3.09% / -0.12% / +2.55% | -3.06% / -0.12% / +2.50% |
| Destruction | shares + hit / Doom rotation | -2.57% / +1.38% | -2.59% / +1.44% |
| **All steps, horde** | Affliction / Demonology / Destruction | **-3.09% / -0.74% / -1.23%** | **-3.08% / -0.75% / -1.18%** |
| **All steps, alliance (gnome)** | Affliction / Demonology / Destruction | **-2.88% / -0.36% / -0.99%** | **-3.01% / -0.47% / -0.96%** |

Absolute DPS at seed 1, horde: Affliction 601.0 to 582.4, Demonology 664.1 to 659.2, Destruction 676.3 to 668.0. Alliance: 605.2 to 587.8, 661.1 to 658.7, 678.2 to 671.5. The order is still Destruction, then Demonology, then Affliction. The demon's share of damage is now about 12% (Affliction), 12% (Demonology) and 11% (Destruction), against 15%, 15% and 12% before.

## Still open

- Pet crit: Blizzard's stated intent is not live (X3; +1.7 to +2.6% if it ships). Pet haste (X4). A launch-build log settles both.
- Pet hit (X2) rests on one Blizzard statement and has not been read on a sheet.
- Eureka! class masks (C2) need a core fix for every class.
- The demon sheets at 60 (base attack power and weapon) are wowsims' Classic figures. wowsims/forever extrapolates a fixed 186 attack power for the Sayaad from levels 20 and 30, which is single-source and not read at 60.
- Demonic Brand stays unmodeled. Siphon Life tick crit stays held for the owner.
- The published BiS json still carries the pre-round-3 Destruction string (nightly-owned). sim/request TestRotationLadder goldens are not regenerated, as the brief requires.

Sources (read 2026-11-01): https://github.com/ClassicWoWCommunity/forever-bugs/issues/331, https://github.com/ClassicWoWCommunity/forever-bugs/issues/394, https://github.com/Gethe/wow-ui-source/tree/forever, https://github.com/wowsims/forever/blob/895cd67f9/sim/warlock/pets.go, https://github.com/bayov/forever-sim/pull/85, https://nether.wowhead.com/forever/tooltip/spell/603 (and 6223, 6217, 11672, 11713, 1293813, 412732, 18881), https://wago.tools/db2/SpellEffect/csv?build=1.60.1.70707.

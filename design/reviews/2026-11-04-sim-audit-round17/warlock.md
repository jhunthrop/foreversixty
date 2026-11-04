# Sim-accuracy round 17: warlock (2026-11-04)

Specs: Affliction, Demonology, Destruction, bands 20 to 60, bare and raid. Client 1.60.1.70755 with the late hotfix re-layer. Site `wt/r17-warlock` from int/r17 (d7fc51fe7); fork `lane/r17-warlock` from int/r17 (f637ec32c), no engine change. spec-breakdown at 10,000 iterations, seeds 1 and 50001, on `$R/bis-frozen` with each guide's raid build unless noted.

Totals: **1 CONFIRMED-BUG**, fixed on the site, no engine change.
- Affliction leveling build: Nightfall 2/2 does nothing from 40 to 60 (origin `old`: dead since round 15 made Wrack the filler; round 15 A4 looked and kept it).

## Logs (owner's standing order)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | **Still no launch log above level 30.** `GET api.foreversixty.gg/v1/reports/recent` returns the same single report (z2opt2x2gs5u, 2026-10-20, no boss). foreverlogs.gg `reports/public` newest is still 2811 (2026-10-24, Stockade, level 30). The only new id is 2814 ("Test", uploaded 2026-11-04 12:24 UTC, desktop client 0.1.58, zero encounters); 2815 to 2850, 2900 and 3000 are "Report not found". Warcraft Logs stays blocked (gcloud login), not tried. | foreverlogs.gg API and api.foreversixty.gg, 2026-11-04 ~07:45 local | OPEN (re-check next round) | - | - |

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Eureka! on Incinerate.** No log settles it (L1). Client: Eureka! 1259821 effect 1 (SPELLMOD_DAMAGE +10) has class mask word 0 0x801e5, word 1 0x800080; Incinerate (412758, 1293812, 1293813) is word 1 0x40, so the client leaves it out; effect 2 (SPELLMOD_DOT +10) reaches Drain Life, Drain Soul and Wrack. The engine (`sim/core/eureka.go`, `eureka_auto_gen.go`) matches. New web source: a player post on the Blizzard forums (PhoenixHa, 2026-10-07, https://us.forums.blizzard.com/en/wow/t/2375151) quotes the 1 October note "No longer benefits periodic effects at all. Channeled spells do not count as periodics", which agrees with the engine on DoTs and channels; the same post assumes Incinerate benefits, but that is the poster's reading of the tooltip ("non-periodic damaging abilities"), not a test. Icy Veins' nerf article (https://www.icy-veins.com/wow-forever/news/eureka-world-of-warcraft-forevers-strongest-racial-is-getting-a-massive-nerf/) returned 403. Tooltip and client still disagree on Incinerate. | wago SpellEffect / SpellClassOptions 70755 + push (`$R/warlock/tools/mask.py`); nether.wowhead.com/forever/tooltip/spell/1259821 | OPEN (a level 40+ Gnome Destruction log settles it) | - | - |
| S2 | Demon crit at 60 (C5 of round 16): no log. Lash of Pain crits 6.0% (Demonology) / 8.6% (Destruction) / 10.3% (Affliction) with Malevolence; demon melee 0.0% on a boss as logged at low level (#331). | spec-breakdown | OPEN (carried) | - | - |
| S3 | Soul Siphon's count (which effects it counts is server-side; 17804 is a dummy): no log. It cannot move a number where Wrack is cast (the three DoTs already reach the 36% cap). | round 16 A3 | OPEN (carried, no number at stake) | - | - |
| S4 | **Regression: Demonology's Sayaad point holds.** Raid build `05/2035103021220001351/0550005` still has Improved Voidwalker 0, Improved Imp 0, Improved Sayaad 2 (`TestDemonologyRaidBuildSpendsNothingOnDeadDemons` passes). Client 18754: aura 108 base 30 at 3/3, 10% a point (talents/warlock.json rank text); engine `succubus.go:116` multiplies Lash of Pain by 1 + 0.10 x rank. The raid fight summons the Imp, sacrifices it (Burning Shadow 18789: +15% Shadow, aura 79 misc 32), summons the Succubus at -44 s and casts Lash 15.5 times (12 s cooldown). Raid 60 matches `before.tsv` to the decimal (707.2 / 706.7 horde, 717.1 / 718.1 alliance). | spec-breakdown; wago 70755 | Aligned | - | - |
| S5 | **Regression: per-level Shadow Bolt rows hold.** Re-swept the row boundaries on the round-17 start engine, bare, 10,000 iterations, seeds 1 and 50001: Destruction 23 r3 61.9 > r4 61.6, 24 tie (65.2), 29 r5 72.6 > r4 72.3; Demonology 21 r3 55.6 > r4 55.3, 22 r3 56.2 vs r4 56.1 (within the 0.2% bar, default stands), 27 r4 64.4, 28 r4 66.7 > r5 66.0, 30 r4 91.7 > r5 91.3, 31 r4 92.4 vs r5 92.3 (within 0.2%, default stands). Identical to round 16 W1. `TestWarlockShadowBoltRowsFollowTheLevel` passes. | `$R/warlock/rowcheck.txt` | Aligned | - | - |

## Affliction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **The leveling build spent 2 points on Nightfall, which does nothing from 40 to 60.** Nightfall (18094, 2% a point on Corruption, Drain and Wrack ticks) makes the next Shadow Bolt instant (Shadow Trance 17941). Since round 15 Wrack is the filler from 40 and the rotation casts no Shadow Bolt, so the procs are never spent: bare 40 / 50 / 60 show 5.8 / 5.9 / 6.3 Shadow Trance procs a fight, up 27-29% of the time, and zero Shadow Bolts. Spending them does not help: an instant Shadow Bolt before Wrack measured 185.1 to 184.6 (40), 268.0 to 266.6 (50), 424.1 to 421.5 (60); gated on mana above 40% still -0.1% to -0.2%, above 70% exactly level (it never fires). The same-tier swap to Malevolence 2 (1% Shadow crit a point; Pandemic doubles the bonus on DoTs and Wrack): bare 60 423.9 / 423.7 to 430.0 / 430.0 (+1.44% / +1.49%, error ±0.05%); level 40 185.1 to 187.9 (+1.5%, both seeds), 45 211.9 to 215.1 / 215.0 (+1.5%); and below 40, where Shadow Bolt is still the filler and Nightfall does work, Malevolence still holds or wins: 35 +0.5% / +0.2%, 37 +0.2% / +0.3%, 39 +0.3% / +0.3%. Adoption rule met: keeps every unmodeled talent (Soul Harvest, Improved Drains, Improved Health Funnel, Improved Voidwalker), holds the spec tree, gains beyond 1% at both seeds, drops nothing the rotation casts; the ladder stays monotone (Malevolence takes Nightfall's 34-35 slot in the same tier). New leveling build `25522300132001351/200522/005`. Test first: `TestAfflictionLevelingBuildTakesMalevolenceOverNightfall` (sim/leveling) failed (Malevolence 0, Nightfall 2) and passes. Guide prose rewritten (affliction.md:48, :54, :59, :115); wow-player review below. | runs (`-talents`, `-level`, `-rotation` probes `$R/warlock/probe/aff-*.json`); talents/warlock.json | CONFIRMED-BUG | old (dead since round 15 A1 made Wrack the filler; round 15 A4 measured the Shadow Trance line and kept Nightfall) | S 8a4114ad9, S e18d0bb92 |
| A2 | Raid build unchanged and re-checked: every point acts on a cast spell or is a keep (round 16 A2). | dec.py | Aligned | - | - |
| A3 | Leveling keeps: Improved Drains 3 and Soul Harvest 2 (drains and Drain Soul kills, unmodeled questing tools), Improved Voidwalker 2, Improved Health Funnel 2, Demonic Aegis 2 (the guide levels behind the Voidwalker, affliction.md:115). Utility for a real leveler, not provably dead. | guide; round 15 A6 | Aligned | - | - |
| A4 | Curated APL `notes` and `engineering_notes` "5" still called Wrack the "#2 damage source behind Shadow Bolt" (carried by guide-accuracy rounds 10 to 13). Rewritten; not rendered, no rotation change; the sim copy carries no notes. | data/curated/apl/warlock-affliction.json | (housekeeping) | - | S 49ce5467d |

## Demonology

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | Bare 40 and 60 fights: Succubus out, Soul Link up, no sacrifice (no Demonic Pact before 41 points), Soul Fire only inside Decimation (7.6 casts at 60, 2.4 s with Bane and Decimation), Life Tap 24-28 a fight, out of mana 0 / 3.9 s. Ranks: highest learned rank of every spell other than Shadow Bolt (rows above). | spec-breakdown | Aligned | - | - |
| D2 | Leveling keeps: Improved Voidwalker 1, Demonic Energies 1 (pet heal; its pet mana never binds, Lash casts on cooldown), Demonic Embrace 3, Demonic Aegis 1, Improved Health Funnel 2, Master Summoner 2: survival and pet-tanking picks the guide explains (demonology.md:122-128). | guide | Aligned | - | - |

## Destruction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Rank sweep on the mana-bound spec** (Destruction Life Taps about 30 times a fight at every band, 1.3 to 4.7 s out of mana). Round 16 swept only Shadow Bolt. Highest rank against the one below, bare, 3,000 iterations: Incinerate 50 r2 278.9 / r1 274.9, 53 r2 287.7 / r1 282.4, 60 r3 458.1 / r2 447.9; Immolate 40 r5 169.9 / r4 169.8 (tie), 50 r6 278.9 / r5 272.8, 60 r7 458.1 / r6 451.6; Conflagrate 40 r3 169.9 / r2 169.2, 54 r5 297.6 / r4 294.0, 60 r6 458.1 / r5 454.8; Corruption 44 r5 175.8 / r4 174.6. The default wins everywhere; no row needed. | `$R/warlock/tools/rprobe.sh` | Aligned | - | - |
| X2 | Raid fight: Shadow and Flame up 95% (Conflagrate half) and 89% (Shadowburn half), Conflagrate 15.75 casts, Shadowburn 10.7, Incinerate 37.4. Client Incinerate 1293813: 217 + 1.4 a level, 0.15 wide, coefficient 0.714, MaxLevel 69, so 200.7-233.3 at 60 (`incinerate.go:37`); Wowhead's "213 to 246" is the same row evaluated at MaxLevel 69, which no player reaches. | spellconst/warlock.json; nether.wowhead.com/forever/tooltip/spell/1293813 | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | Round-16 engine changes that reach the warlock: the MaxLevel + 12 penalty (`sim/core/spell.go`, SpellLevelPenalty) and Wrack's comment-only change. Every rank the raid APLs cast is at or near its MaxLevel at 60 (no penalty); band ranks follow the rows (S5, X1). | fork diff 9f7f1f6e4..f637ec32c | Aligned | - | - |
| C2 | Orc: Blood Fury 20572 is Forever's +10% AP, RAP and spell power for 15 s (2 casts, 16.7% uptime), Command +5% pet damage and +5% on Lash of Pain: modeled (`sim/core/racials.go:220-235`). | races.json; engine | Aligned | - | - |
| C3 | Pending owner, not reopened: level-60 trainer book ranks (25307, 25309, 25311); Siphon Life tick crit; Demonology Horde race. | | pending owner | - | - |

## Before/after

Raid 60 (guide raid builds, frozen BiS): unchanged on all 12 runs (spec x faction x seed), identical to `before.tsv` to the decimal (Affliction 733.1 horde, Demonology 707.2, Destruction 776.4 at seed 1).

Affliction bare (leveling build, horde; seed 1 / 50001):

| level | before (Nightfall) | after (Malevolence 2) | % |
|---|---|---|---|
| 35 | 107.4 / 107.5 | 107.9 / 107.7 | +0.5 / +0.2 |
| 37 | 113.1 / 113.0 | 113.3 / 113.3 | +0.2 / +0.3 |
| 39 | 121.4 / 121.4 | 121.8 / 121.8 | +0.3 / +0.3 |
| 40 | 185.1 / 185.1 | 187.9 / 187.9 | +1.5 / +1.5 |
| 45 | 211.9 / 211.9 | 215.1 / 215.0 | +1.5 / +1.5 |
| 60 | 423.9 / 423.7 | 430.0 / 430.0 | +1.44 / +1.49 |

The published band numbers move only after addon-builds and the ladder goldens are regenerated (below). Nothing outside the warlock moves.

## Files to regenerate

- `data/builds/1.60.1.70755/addon-builds.json` (Affliction leveling order: Malevolence 2 at 34-35 in place of Nightfall); `TestCommittedAddonBuildsFollowTheGuides` (sim/cmd/addon-builds) fails until then.
- `addon-data.json` / Data.lua and the coach starters (Affliction leveling talents).
- Ladder golden `sim/request/testdata/ladder/warlock-affliction.golden.md` (rows 40, 50, 60 carry the old talent string).

## Measure at refresh

- Affliction bare 40, 50, 60 weights and leveling BiS on the new build (crit rises a little in value).

## For other lanes

- No shared code changed; nothing outside the warlock moves.
- Platform: the guide band lists show each band's floor build (see the wow-player review).

## For the guide lane

- Done in this lane: affliction.md:48, :54 (the Nightfall bullet is now Malevolence), :59, :115; wow-player review below.

## Still open

- Launch logs (L1): Eureka! on Incinerate (S1), demon crit (S2), Soul Siphon's count (S3), Wrack and Doom at 60.
- Pending owner items (C3).

## wow-player review (Affliction page, rendered at 1440)

**SHIP** on the second pass, no MUST. The first pass was NO-SHIP because the talent panels had not hydrated in the capture; the recapture shows the leveling tree with Malevolence 2/5 and Nightfall 0/2, the raid card with Malevolence 5 (35/7/9), and the "moves 9 points" line adding up. Two SHOULDs were done in this lane (S e18d0bb92): the Malevolence bullet now covers 34 to 39 (about level with Nightfall while Shadow Bolt is the filler) and says plainly that a Nightfall bolt still costs full mana.

For the platform lane (shared planner component, not this guide): the per-band lists show the build at the band's floor, so "Level 30 to 39" lists 21 points and a level-36 reader sees no Malevolence, Pandemic or Soul Siphon, though the prose says Malevolence comes at 34 and 35. This hits every guide's band lists.

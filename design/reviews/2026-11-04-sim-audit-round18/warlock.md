# Sim-accuracy round 18: warlock (2026-11-04)

Specs: Affliction, Demonology, Destruction, bands 20 to 60, bare and raid. Client 1.60.1.70755 with the late hotfix re-layer. Site `wt/r18-warlock` from int/r18 (4600b7d82); fork `lane/r18-warlock` from int/r18 (b2c4a2d05). No engine or site code changed. spec-breakdown on `$R/bis-frozen` with each guide's raid build (`$R/rt`) unless noted.

Totals: **0 CONFIRMED-BUGs.** Both round-17 regression targets hold. Every log-dependent item is still open, because no level-60 Forever log exists yet. Forever launches at 23:00 UTC; this lane last checked the logs at 16:18 UTC.

## Logs (owner's standing order)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | **Still no launch log.** `GET api.foreversixty.gg/v1/reports/recent` returns only z2opt2x2gs5u (2026-10-20, 42 fights, no boss). On foreverlogs.gg the newest `reports/public` entry is still 2811 (2026-10-24, Stockade, level 30). Ids 2812, 2813 ("shield", 2026-10-29) and 2814 ("Test", 2026-11-04 12:24) have zero encounters. 2810, 2815 to 2817, 2820, 2830, 2850 and 2900 return "Report not found". Checked at 16:08 and again at 16:18 UTC, before launch. | foreverlogs.gg API (`r3/casters/py/fl.sh`); api.foreversixty.gg | OPEN (re-check after launch) | - | - |
| L2 | **Warcraft Logs is blocked** because the gcloud login has expired. Per the brief, this lane did not try to work around it. | brief | recorded | - | - |

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Eureka! on Incinerate.** No log. A new web search found only guides and tier lists (Warcraft Tavern, Ranked Boost, gamerblurb, gamer-guides) that assume direct casts benefit. None of them tests Incinerate. One guide says its numbers come from "the BlizzCon floor build", not a patch note. The Blizzard forum thread (https://us.forums.blizzard.com/en/wow/t/2375151, 2026-10-07) is still the only near-source, and its Incinerate line is the poster's reading of the tooltip. The client mask (1259821 effect 1, word 1 0x800080) still leaves Incinerate (word 1 0x40) out, as round 17 S1 found. The engine follows the client row. | web search 2026-11-04; round 17 S1 | OPEN (a level-40+ Gnome Destruction log settles it) | - | - |
| S2 | **Demon crit.** No log. Lash of Pain crits 10.2% for Affliction and 5.9% for Demonology (raid 60, horde); Succubus melee crits 0.0% on a boss. Both are unchanged from round 17. | spec-breakdown | OPEN (carried) | - | - |
| S3 | **Soul Siphon's count.** No log. It still moves no number: Corruption, the bane and Siphon Life are up for every Wrack cast, so the 36% cap is reached either way. | round 16 A3 | OPEN (carried, no number at stake) | - | - |
| S4 | **Regression: Malevolence over Nightfall in the Affliction leveling build.** affliction.md:6 is `25522300132001351/200522/005` (Malevolence 2, Nightfall 0). `TestAfflictionLevelingBuildTakesMalevolenceOverNightfall` passes (sim/leveling). The ladder golden from round 17's regen carries the new string at 38 to 60. Note: spec-breakdown `-preset bare` with no `-talents` still reports Shadow Trance procs at 40, 50 and 60. That is because the published BiS band entries (`data/builds/1.60.1.70755/bis/warlock-affliction.md:69`, `:99`, `:129`) still carry the Nightfall string `...130201...`. The running BiS refresh rewrites those files, and they are not this lane's to touch. | go test; ladder golden; bis md | Aligned | - | - |
| S5 | **Regression: Heart of the Lion on demons.** `TestDemonReceivesOnlyHeartOfTheLionFromPlayerBuffs` passes (fork sim/warlock). On raid 60, Heart of the Lion 409583 is up 99.9% on the Destruction Succubus, and the bare preset has no buffs at all. All 12 raid-60 runs (three specs, both factions, seeds 1 and 50001) reproduce `before.tsv` to the decimal, for example Demonology horde 724.3 / 724.0. | `$R/warlock/start.tsv`; go test | Aligned | - | - |
| S6 | Round-17 code from other lanes that reaches warlocks was checked. Hit-weight profile (7c8f710d0) and the hunter pet opt-in (c6f2e362e) are covered by S5. The rank-choice `lines`/`prepull` scoping (7ea07b551) is covered by `TestWarlockShadowBoltRowsFollowTheLevel` and the other `Warlock` tests in sim/request, which all pass. No warlock row changed in round 17. | go test | Aligned | - | - |

## Affliction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Raid fight: Wrack 22.8 casts and 123 ticks (5.4 a cast, because a missed Wrack cast ends the channel), 47.5% of damage, 24.9% crit with Malevolence 5. Corruption, Doom, Siphon Life and Agony all run; Amplify Curse goes before the last-minute Agony; Life Tap is cast 2 times a fight. The vulnerability, Soul Siphon's per-tick read and the dynamic re-snapshot are as in round 15 A5 (`wrack.go`, `dot.go:265`, no compounding). | spec-breakdown; engine | Aligned | - | - |
| A2 | **Potion slot.** I replaced Major Mana Potion with Major Spellblasting Potion, which the mage and shadow priest rows adopted. Raid 60 horde, seeds 1 / 50001: Affliction 750.6 / 750.8 → 748.5 / 748.0 (-0.3%); Demonology 724.3 / 724.0 → 713.3 / 713.4 (-1.5%); Destruction 792.2 / 792.6 → 788.1 / 788.3 (-0.5%). The mana potion stays. | spec-breakdown `-consumes` | Aligned | - | - |
| A3 | Bare 30/40/50/60 fights show every learned rank and no unlearned spell (CoE ranks 1311676, 1311677, 1311680 by trainer level; Wrack only from 40). Nothing runs out of mana. | `$R/warlock/bd/` | Aligned | - | - |

## Demonology

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | Raid fight. Prepull: Imp at -60 s, sacrifice at -45 s, Succubus at -44 s, Soul Link at -2 s. In the fight: Decimation is up 41% and Soul Fire is skipped (the Demonic Pact gate, round 15 G3); Shadow Bolt is 48% of damage. | spec-breakdown | Aligned | - | - |
| D2 | **Downrank sweep** on the one mana-bound Demonology fight (bare 60 is out of mana for 3.9 s). Round 17 X1 swept only Destruction. Default rank against the one below, 3000 iterations, seed 1: Immolate r7 386.2 / r6 381.7 at 60, and 240.0 / r5 235.5 at 52; Corruption r6 386.2 / r5 381.1 at 60, and 171.9 / r4 170.0 at 45; Soul Fire r2 272.5 / r1 271.7 at 58. The default wins at every level tried, so no row is needed. | `$R/warlock/tools/rprobe.sh` | Aligned | - | - |
| D3 | Raid-build filler points (Improved Health Funnel 2, Demonic Embrace 3) are re-checked against Improved Corruption 5 (+10% Corruption and an instant cast, worth about +2.5%). They cannot leave the tree, because 26 Demonology points do not reach Demonic Pact. Taking Improved Corruption at the cost of Suppression gives up 5% hit, and costing Ruin gives up about 4.5%. Both are worse. The rounds 13/16 keeps stand. | talent gates; dec.py | Aligned | - | - |

## Destruction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Shadow and Flame values against the 70755 curve.** The rank text gives 4/8/12/16/20% Shadow after Conflagrate and Fire after Shadowburn for 20 s, with a 20 to 100% chance to keep Immolate. Engine `talents.go:959` gives 0.04 a rank, the 20 s Flame buff (426311 DurationIndex 18) and the Immolate keep. Raid uptimes are 94.9% (Shadow) and 89.3% (Fire). Older reports quote 10%, which is the pre-push value. | talents/warlock.json; 426311; engine | Aligned | - | - |
| X2 | Talent masks re-read bit by bit: Agonizing Flames (crit on Searing Pain; +10% direct on Shadow Bolt, Immolate, Conflagrate, Incinerate, Shadowburn, Soul Fire, Searing Pain and Rain of Fire; +10% periodic on Immolate and Rain of Fire), Ruin (the Destruction list without Bane of Havoc), Cataclysm (the same list plus Havoc), Fire and Brimstone (Conflagrate crit +25%; effect 0 is a dummy). Conflagrate ranks 1293817, 1293818, 17962, 18930, 18931 and 18932: mana 100/130/165/200/230/255, levels 25/32/40/48/54/60, 10 s cooldown, 0.429 coefficient. All of these match `conflagrate.go:20-38`. | wago 70755 layered (`tools/mask.py`, `tools/spell.py`) | Aligned | - | - |
| X3 | Raid fight: Conflagrate crits 40.0% (15% plus Fire and Brimstone's 25%). Incinerate 37.4 casts, Immolate 13.8, Shadowburn 10.7 (the shard comes back), Life Tap 14.8, 0.2 s out of mana. Prepull is Immolate at -1.5 s. | spec-breakdown | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | BiS trinket Frozen Heart of the Mountain (249469): use spell 1249110 Chthonic Power gives +29 Frost and Shadow (aura 13, misc 48) for 15 s on a 5 min cooldown. That matches `item_effects.go:3724-3753`. | wago SpellEffect 1252390 | Aligned | - | - |
| C2 | Human racials (published race): Ruthless +1% crit and The Human Spirit +10% Spirit. Spirit feeds Life Tap (`lifetap.go:44`, the client's $SPI*1). | races.json; racials.go:204 | Aligned | - | - |
| C3 | Life Tap per rank (424 at 60 on rank 6, +1 per level to MaxLevel, Improved Life Tap 10% a point on the whole amount) and Curse of the Elements per rank (-30/-45/-60/-75 resistance, +4/6/8/10%) match the client. | spell rows; tests | Aligned | - | - |
| C4 | Pending owner, not reopened: level-60 trainer book ranks (25307, 25309, 25311), Siphon Life tick crit, the Demonology Horde race (Troll's Berserking is the published pick), Heart of the Lion 280/276. Haunt and Unstable Affliction are still `class_spell` with no learn row (round 13 X5). | | pending owner | - | - |

## Before/after

No change in this lane. At round start, raid 60 equals `before.tsv` on all 12 runs. Seed 1 horde: Affliction 750.6, Demonology 724.3, Destruction 792.2.

## Files to regenerate

None from this lane. The published `data/builds/1.60.1.70755/bis/warlock-affliction.*` band entries still carry the pre-round-17 Nightfall talents, and the BiS refresh in flight rewrites them.

## Measure at refresh

- After the refresh, confirm that the Affliction bare 40, 50 and 60 BiS entries carry the Malevolence string. If they do, Shadow Trance disappears from `-preset bare` runs.

## For other lanes

- Nothing. No shared code was touched.

## For the guide lane

- Nothing. No build, APL, card or rank row changed.

## Guide text not done

- None.

## Still open

- Launch logs (L1): Eureka! on Incinerate (S1), demon crit (S2), Soul Siphon's count (S3), and Wrack and Doom at 60. The first level-40+ Destruction Gnome log and any level-60 Demonology log settle them.
- Pending owner items (C4).

## Sources

- Blizzard forums, "Gnome Warlock Eureka: The October 1 DoT Nerf Conflicts With Blizzard's October 7 Affliction Design", https://us.forums.blizzard.com/en/wow/t/2375151 (read via search, 2026-11-04)
- Warcraft Tavern Forever warlock guide, https://www.warcrafttavern.com/forever/guides/warlock/ (search, 2026-11-04)
- Ranked Boost, best race for warlock, https://rankedboost.com/world-of-warcraft/forever/tier-list/best-race-warlock/ (search, 2026-11-04)
- gamerblurb, https://gamerblurb.com/guides/wow-forever-best-race-for-warlock (search, 2026-11-04)
- gamer-guides racials, https://gamer-guides.com/warcraft-forever/racials/ (search, 2026-11-04)
- wago.tools 70755 layered SpellEffect, SpellClassOptions, SpellPower, SpellLevels, SpellCooldowns, SpellMisc (cached under `r17/warlock/wago`)

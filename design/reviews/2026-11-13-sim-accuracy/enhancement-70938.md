# Sim accuracy 2026-11-13: Enhancement on 70938, every band 20 to 60 and raid 60

Lane brief: `design/reviews/sim-accuracy/brief.md`. Site from origin/main a544d507a, fork from master 5603877e4 (the pin). Client on the site: 1.60.1.70938 with the corrected hotfix layer (273623f0f), which carries Blizzard's 10 November Enhancement changes as push 113722's rows. Scope: Enhancement shaman, bare 20 to 60 and raid 60: the engine against the data now that the data states the hotfix, the logs recorded after the changes went live (12 Nov 10:45 PST, 18:45 UTC), the imbue and totem choice, the rotation, the raid and leveling builds, and the owner's question "do Enhancement shamans not use Flame Shock?".

**Totals: 0 engine bugs. The engine's 10 November values were already the server's; the hand rows the data now carries are retired (Stormstrike's cooldown and cost read the generated constants, the conformance "server hotfix" verdict is gone because Stormstrike's 8 s is now a plain client match, and Stormstrike's 100%, the 20% stack, the five stacks, the 30 s and Mental Dexterity's top rank are checked against the client rows). Two card and build changes adopted under the adoption rule: Flame Shock after Stormstrike while it is off the target and mana is above 30% (bare 30 to 60 +0.7 to +1.8%, raid 60 +0.6 to +0.8%), and the raid build's Mindfulness 3 into Reverberation 5 and Ancestral Knowledge 4 (+1.2 to +1.6% at raid 60 on the new card). Raid 60 Alliance 873.4 to 890.0 (+1.9%), Horde 857.4 to 871.9 (+1.7%); bare 60 Alliance 336.6 to 341.9 (+1.6%), Horde 334.3 to 339.6 (+1.6%). Engine pin 5603877e4 to f17711960.**

## Method

1. **Engine against the data.** The fork's shaman files against `data/builds/1.60.1.70938/spellconst/shaman.json` and the push rows in `raw/wago-push/values.json` (push 113722). Fork tests on the box (`-tags with_db`).
2. **Logs** (`enhancement-70938/logs/`): foreverlogs.gg 2855 to 2857 (2854 ends before the deploy) and our own reports through the Cloud SQL proxy, read only; private reports in aggregate. Only fights that start after 18:45 UTC on 12 November.
3. **Sims** on the Windows box, 10,000 iterations, one at a time, nice 15, GOMAXPROCS 5: `base.sh` (the published entries), `jobs.sh` (imbues at every band, the four searches), `jobs2.sh` and `jobs3.sh` (the Flame Shock positions and gates, the raid build), `jobs4.sh` (the owner's three shock cards). Outputs beside each script.

## Engine against the data: the hand rows

Push 113722 in the 70938 data: SpellCooldowns 55241 (17364 RecoveryTime 8000), SpellEffect 1368308 (17364 effect 2, weapon percent, 100), SpellEffect 1137418 (408498 effect 1, 2), CurvePoint 251361 to 251364 (curve 85306, -20 at every rank), the new curve 128299 (TraitDefinitionEffectPoints 25235: 408498 effect 1 at 6/7/8/9/10 by rank), CurvePoint 251264 to 251266 (curve 85272, Mental Dexterity 33/66/100).

| row | data (push 113722) | engine before | after |
|---|---|---|---|
| Stormstrike cooldown | 8000 | 8 s by hand; conformance verdict "server hotfix" | reads `StormstrikeCooldownMS[0]`; the verdict and its whole mechanism (`server_hotfix.go`, the branch in `compare.go`) deleted, the golden reads "match" |
| Stormstrike cost | 8% base mana | 0.08 by hand | reads `StormstrikeManaCostPct[0]` |
| Stormstrike weapon percent | effect 2 = 100 | 1.0 | stays a constant (the generated arrays carry no effect-2 column); `TestStormstrikeIsTheClientRow` checks it, the charge's 20% and its 12 s against the client row |
| Maelstrom Weapon stack | 408505 -20 cast time and cost, 408498 effect 0 = 20, effect 2 = 5, 30 s | 20, 5, 30 s | checked against the rows (`TestMaelstromWeaponIsTheClientRow`) |
| Maelstrom Weapon rate | curve 128299: 6/7/8/9/10 | 6/7/8/9/10 a minute on the main hand's speed | stated (no generated table carries a trait curve); comment cites the push |
| Mental Dexterity | curve 85272: 33/66/100 (spellconst carries the top rank, 100) | 0.33/0.66/1.0 | rank 3 checked against 415140 effect 0; ranks 1 and 2 stated from the curve |
| imbues with weapon totems | Blizzard's post 6; the logs (below) | both apply | unchanged |
| Flametongue Totem rank 2 | 11 Nov fix: 8250 now triggers 8248 | the engine never modelled the bug | unchanged; the logs show it proc at the rank's amount |

Every published number is unchanged to the decimal on the retired rows (`base.sh` at 5603877e4 and at 72b7ad453, every band both factions). Red and green: `TestStormstrikeIsTheClientRow` fails at 16e5eb6bf ("cooldown = 6s, want the client's 8s") and passes at 72b7ad453 once the constants are regenerated from the 273623f0f data.

## Logs since the deploy

No Enhancement above level 46 has logged since 12 November 18:45 UTC. What exists: foreverlogs 2855 to 2857 (three Stormstrike shamans at about 20 to 30, 283 Stormstrikes), our public report ctxicu4cuyek (four Stormstrike shamans at 24 to 32) and one private report from 13 November (three shamans at 40 to 46, a patchy logger).

| measure | logs | engine | verdict |
|---|---|---|---|
| Stormstrike cooldown | 10th percentile of 179 gaps 8.00 s (lowest 7.98); five shorter gaps follow a dodge or parry (Improved Stormstrike) | 8 s, reset on dodge and parry | aligned |
| Stormstrike size | non-crit hit against the same shaman's white hit 1.07 on a 3.3 weapon, 1.14 on a 2.9; before armor within 1 to 3% of 100% normalized weapon damage | 100% normalized | aligned (150% would be 50% high) |
| Maelstrom Weapon rate | 66 stacks over 141 s of melee, about 28 a minute; 0.45 a white swing, 0.69 a Stormstrike; Windfury Weapon's strikes roll too | per melee hit at the rank's rate on weapon speed, about 26 a minute at raid 60 | aligned: the per-hit reading holds, a 10-a-minute cap is ruled out (closes the 2026-11-12 open row) |
| bolts at five stacks | 4 of 4 instant | the card casts at five | aligned |
| Windfury Weapon with Windfury Totem | both on the same shaman in the same fights, both from one swing four times; imbue 13 procs on 60 landed hits (0.22), shortest gap 2.1 s | 20% a hit, 1.5 s apart, both apply | aligned |
| Windfury Totem rank 2's hit | 6 landed hits on two shamans log as their own physical hit (10608) about 35% above an extra white swing with +179 attack power | an extra white swing | open (O1): six hits, and no rank 2 data from before the hotfix to tell the rank from the hotfix |
| Flametongue Totem rank 2 | procs at 781 x speed / 100 (28 on a 3.5 weapon), logged as 16368 | the same amount | aligned |
| Flame Shock | cast by every shaman at 28 and above (2.5 a minute for the busiest; 3 each for the two at 42 and 44) | not cast before this lane | see below |

## Imbues at every band (`jobs.sh`, seed 1)

Windfury Weapon from 30 at every band and in the raid: raid 60 873.4 against Rockbiter 773.5 and Flametongue 719.3 Alliance (+12.9%), 857.4 against 760.6 and 705.7 Horde; bare 60 +3.9 to +4.1% over the next, bare 50 +0.6 to +1.1% over Flametongue, bare 40 +7.4 to +7.6% (Fiery War Axe back at 3.6 s), bare 30 +8.5%. At 20, Flametongue Weapon over Rockbiter +1.3% both factions. The guide's imbue lines stand.

## Searches at 60 (`jobs.sh`, seed 7, the ranker's levels)

- Raid talents: Mindfulness 3 into Reverberation 2 and Ancestral Knowledge 1, +1.2% (+10.2 ± 3.3), every unmodeled guide talent kept.
- Bare talents: the leveling build holds (+0.6 ± 1.1).
- Rotations: one mutation in both, Flame Shock while it is not on the target: raid +0.9% (+8.0 ± 3.3), bare +1.8% (+6.1 ± 1.1).

## The owner's question: Flame Shock (`jobs4.sh`)

Three cards on the same gear and builds, 10,000 iterations, errors 0.7 at raid 60 and 0.1 to 0.2 bare. The engine's Forever rules for it: Flame Shock shares Earth Shock's shock cooldown (6 s, 1 s less with five points of Reverberation), its burn lasts 12 s and every tick can crit (the client's periodic crit flag), Stormstrike's 20% is for Lightning Bolt, Chain Lightning and Earth Shock only (client mask 0x100003), Clearcasting makes either shock free, and the fight is 3 minutes.

| cell | Earth Shock only (the old card) | Flame Shock kept up, Earth Shock on the other cooldowns (the new card) | Flame Shock only |
|---|---|---|---|
| raid 60, new raid build, Alliance seeds 1 / 50001 | 883.0 / 882.9 | 890.0 / 889.2 (+0.8 / +0.7%) | 810.1 / 808.9 (-8.3%) |
| raid 60, new raid build, Horde | 865.2 / 865.0 | 871.9 / 870.1 (+0.8 / +0.6%) | 791.8 / 791.8 (-8.5%) |
| raid 60, old raid build, Alliance / Horde seed 1 | 873.4 / 857.4 | 877.9 / 858.0 (+0.5 / +0.1%) | 807.5 / 790.0 |
| bare 60 Alliance / Horde seed 1 | 336.6 / 334.3 | 341.9 / 339.6 (+1.6%) | 293.6 / 292.0 (-12.7%) |
| bare 50 | 217.9 / 217.2 | 220.5 / 220.3 (+1.2 / +1.4%) | 194.5 / 194.7 (-10.7%) |
| bare 40 (dry seconds) | 144.5 / 144.3 (6 / 6) | 146.1 / 145.9 (+1.1%; 10 / 12) | 147.5 / 147.2 (+2.0%; 24 / 23) |
| bare 30 (dry seconds) | 85.1 / 84.5 (0 / 3) | 85.7 / 85.4 (+0.7 / +1.1%; 3 / 5) | 85.6 / 85.3 (+0.6 / +0.9%; 7 / 9) |

In game words: Earth Shock is still your best single shock, and Flame Shock alone throws away Stormstrike's 20% on every shock (it only reaches Earth Shock), so a Flame Shock-only card loses 8 to 13% at 50 and 60. But one Flame Shock burns for 12 seconds, two whole shock cooldowns, so putting it up once and spending every other shock on Earth Shock beats Earth Shock alone: about 1.5% at 60 out of a raid, 0.7% in one. What stopped it before was mana: kept up without a gate it leaves a level 30 or 40 shaman dry for 40 to 60 seconds of a 3-minute fight (`jobs2.txt`), so the line waits for a third of your mana (30% kept the gain and the dry time near the old card's; 50% and 70% gave less, `jobs3.txt`). Flame Shock only edges out at 30 and 40, but only by running dry two to four times as long. The guide's old line ("Flame Shock costs more mana than its damage is worth") was the Earth Shock-only answer and is replaced.

**Adopted** under the adoption rule (≥1% beyond error at bare 40 to 60, nothing the card casts dropped, Earth Shock still cast; raid gains less but gains): the card casts Flame Shock right after Stormstrike while it is off the target and mana is above 30%.

## The raid build

On the new card, Mindfulness 3 into Reverberation 5 and Ancestral Knowledge 4 (`0535001/054030310051202051/05`): Alliance 889.7 / 889.3 against 877.9 / 878.0 (+1.3%), Horde 871.5 / 870.0 against 858.0 / 859.4 (+1.2 to +1.6%), never dry either way. A raid shaman never runs dry (Judgement of Wisdom, Mana Spring), so regeneration while casting buys nothing, while a shorter shock cooldown is more Earth Shocks. Keeps every unmodeled guide talent (Improved Ghost Wolf) and everything the card casts: **adopted** (raid-searches had held it at +0.8% on the old card, 800 iterations). The leveling build holds.

## Before and after (published entries, seed 1)

| entry | before (5603877e4, old card and raid build) | after (f17711960, new card, new raid build) |
|---|---|---|
| raid 60 Alliance | 873.4 ± 0.7 | 890.0 ± 0.7 (+1.9%) |
| raid 60 Horde | 857.4 ± 0.7 | 871.9 ± 0.7 (+1.7%) |
| bare 60 Alliance / Horde | 336.6 / 334.3 | 341.9 / 339.6 (+1.6%) |
| bare 50 | 217.9 / 217.2 | 220.5 / 220.3 (+1.2 / +1.4%) |
| bare 40 | 144.5 / 144.3 | 146.1 / 145.9 (+1.1%) |
| bare 30 | 85.1 / 84.5 | 85.7 / 85.4 (+0.7 / +1.1%) |
| bare 20 | 38.3 / 38.5 | not re-measured: the line casts the band's own Flame Shock rank, and band 20 runs Flametongue Weapon; see MR108 |

Tier list: on raid 60 Enhancement moves from 873.4 to about 890 Alliance, between Destruction (888.0) and Rogue Combat (902.3) on the engine-70938 table; it stays below Fury, the hunters and the cat.

## Fixed, held, open

- **Retired** (refactor, value for value): the Stormstrike cooldown and cost hand values, the conformance server-hotfix entry and mechanism; Stormstrike's percent, Maelstrom Weapon's stack and Mental Dexterity's top rank now checked against the client rows.
- **Adopted:** the Flame Shock line; the raid build. Guide and card text in the same landing, wow-player verdict in `enhancement-70938-wow-player.md`.
- **Not a code bug, needs a data run:** the talent planner still shows the pre-hotfix text (Stormstrike 150%, Mental Dexterity 50/100/150, Maelstrom Weapon 4% a rank) because the talents come from Wowhead's live payload, which the 13 Nov 17:05 UTC data run fetched before Wowhead updated. The payload now carries 100%, 33/66/100 and 20%; the next data run fixes the planner (MR108).
- **Open O1:** Windfury Totem rank 2's hit logs about 35% above an extra white swing (6 hits); needs a level 50 to 60 log.
- **Open O2:** Maelstrom Weapon's per-rank rate and the curve's lower ranks of Mental Dexterity stay stated in the engine until a generated table carries trait curves.

## Shas

- Fork: 16e5eb6bf (refactor and tests, red), 72b7ad453 (constants and conformance golden from 273623f0f, green), f17711960 (the card's fork copy). Master fast-forwarded to f17711960.
- Site: the pin f17711960, the card, the guide, the ladder golden, this report.

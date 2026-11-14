# Sim accuracy 2026-11-13: warlocks on 70938 (Affliction, Demonology, Destruction and their demons)

Lane brief: `design/reviews/sim-accuracy/brief.md`. Site from origin/main 623667798, engine pin 037798db1 (unchanged). Client 1.60.1.70938 with its hotfix push layer as the data run b5bdd533b (14 Nov 01:05 UTC) committed it, newest push in the store 114082. The last warlock lanes were 2026-11-12 (`warlocks-logs.md`, `warlocks-3.md`); their closed items are not redone.

**Totals: 0 engine bugs, nothing fixed, the pin does not move. 0 new warlock data: between 70790 and 70938 every warlock spellconst row that moved is push 113766 (the 10 Nov class push the engine already carries: Wrack, Bane of Havoc's global, Firebolt ranks 1 to 4) or a float written exactly since the fractional-values lane; the talent and trainable files are identical; no push after 113766 touches a warlock spell; Blizzard's 13 Nov post has one warlock line (Siphon Life rank 1 can no longer be cast while silenced), which no sim fight reaches. The first logged Wrack: two Affliction warlocks at 40 in our own uploads (Desolace, 10 Nov; Feralas, 13 Nov) close the 12 Nov lanes' H2 and O3 on every number the engine models: the 1 s tick, six ticks in a full channel, the push's amount, Soul Siphon on Wrack at rank 3, Pandemic's 2.0x crit, and Siphon Life, Corruption and Bane of Agony ticking twice as fast while it channels with their tick count kept. The one item on the warlocks' gear, Channeler's Ring's stale Spirit, is the Fire and Frost lane's D1 (9e6cdc3c2, MR110), which now names the warlock bands too.**

## Method

1. **Client and push.** `warlocks-70938/client_diff.py` against `data/builds/1.60.1.70790` and `1.60.1.70938`: every warlock spellconst row that differs, the talent and trainable files whole, and every push entry from 113700 on whose record or spell id is a warlock spell (`results/client-diff.txt`). Blizzard's hotfix thread t/2403040 to its post of 14 Nov 01:54 UTC (the 13 Nov notes).
2. **Engine.** The fork at the pin, in a worktree (`/Users/jh/code/wowsims-forever-sa-warlock`, branch `lane/sa-warlock-70938` at 037798db1): `wrack.go`, `talents.go` (Soul Siphon, Pandemic, Nightfall), `drain_life.go`, `drain_soul.go`, `core/dot.go`'s clock rate, against the rows below. The fork's warlock, core and conformance tests on the box at the pin (`fork_tests.sh`, `results/box-fork-pin.txt`: all ok; the warlock conformance golden the crit-flags lane left red is green since ff2172004).
3. **Logs.**
   - Our reports (read only: the reports and fights tables through the Cloud SQL proxy, the fight files from the logs bucket; `logs/db.sh`, `logs/r2env.sh`, `logs/r2get.py`): eight reports, the newest ju2t73py6pib (Feralas, 13 Nov 18:38 UTC, units null, so every fight was read). Warlocks found by the warlock spells they cast; the real level floor is the highest spell level cast, since the advanced log's level field is the zone's scaled level (`logs/scan.py`, `results/scan.txt`). Private reports are cited by id and in mechanics only; no names.
   - foreverlogs.gg: the public list answers again (29 reports, newest 2857 on 13 Nov). 2855 to 2857 hold three warlocks at 20 to 34 (`logs/fetchw.py`, the 12 Nov lane's script). Nothing at 40 or above.
   - Warcraft Logs: no Forever partition (the 13 Nov hunter, rogue and tank lanes).
4. **No sims.** No engine change, so no before and after; nothing ran on the Mac.

## Client and push against the engine

| row | 70790 to 70938 | engine | result |
|---|---|---|---|
| Wrack 1316697 | 36 +2.5 a level, 0.286, effect 1 aura 271 at 10 became 25.2 +1.75, 0.2002, a dummy at 100 | `wrack.go` reads the generated row (MR41 retired by the fractional-values lane) | aligned, no change |
| Bane of Havoc 1225228 | gcd 0 to 1500 | native since engine-70938 (MR92) | aligned |
| Firebolt 3110, 7799 to 7801 | the push's 6/11/15/18 centres and widths | `FireboltDamage` reads the table (MR41 retired) | aligned |
| Health Funnel effect 3, the rank-0 NPC copies (Corruption 18376, 21068, 1213450, Curse of Agony 18671, 1233077, Firebolt 20801, Shadow Burst 28447) | whole amount written as its exact float | no reader | n/a |
| talents/warlock.json, trainables/warlock.json | identical but the build stamp | | 0 new data |
| push entries since 10 Nov on a warlock id | 113702 (Command), 113704 (The Quick and the Dead attributes), 113766 (the class push), 113837 (French names of Demonic Frenzy and Massive Destruction, two Season of Discovery runes no Forever warlock learns) | already in the engine or no effect | 0 new data |
| Blizzard 13 Nov: "Siphon Life Rank 1 can no longer be cast while silenced" | | no fight silences the warlock | no effect |

## Logs against the engine

Two Affliction warlocks cast Wrack, the first in any log (`logs/wrack.py`, `logs/clock.py`; `results/wrack.txt`, `results/clock.txt`):

- **Warlock A** (3sdrcwdgo2yi, Desolace, 10 Nov): spell power 119 on his advanced rows, Corruption rank 4, Bane of Agony rank 4, Siphon Life rank 1, three channels with 13 Wrack ticks.
- **Warlock B** (ju2t73py6pib, Feralas, 13 Nov): spell power 213, Siphon Life rank 2, Drain Life rank 4, three channels with 10 ticks.

Both cast nothing at 41 or higher (Corruption rank 5 is not trained), so both are 40 to 43.

| # | finding | evidence | engine | result |
|---|---|---|---|---|
| L1 | **Wrack ticks every 1.0 s, six ticks in a full channel.** Tick offsets from the channel start 0.92 to 1.06, 1.91 to 2.07, 2.94 to 3.04, 4.00 to 4.10, 5.02 and 6.06; the one full channel has six; the others end early on a Shadow Bolt or a death. Each cast costs 200 mana (`adv_power_cost`). | A and B | `WrackTickLength` 1 s, `WrackNumTicks` 6, cost from the client's 200 | aligned |
| L2 | **The tick is the push's amount.** A's non-crit ticks are 51 on every one of 12: 25.2 at the spell's level 40 plus 0.2002 of 119 spell power is 49.0, and 51 is that times 1.04 to 1.06 (his Curse of the Elements rank 2 at 6% on the same target fits). B's are 109 on three DoTs and 99 on two. | A and B | `WrackTickDamage` (25.2, 1.75 a level from 40), `WrackTickCoefficient` 0.2002 | aligned |
| L3 | **Soul Siphon reaches Wrack, at rank 3.** B's target with Corruption, Bane of Agony and Siphon Life took 109 a tick in two channels; the target whose Siphon Life had fallen off (Corruption and Agony left) took 99. 109/99 is 1.101; rank 3's 1.36/1.24 is 1.097, rank 2's 1.24/1.16 is 1.069. A's ticks stay at 51 with two or three DoTs on the target: no points in it. | B, fight 2 | `soulSiphonAffectedSpellCodes` includes Wrack; 12% a counted DoT, at most three, read at each tick | aligned |
| L4 | **Wrack crits at 2.0x with Pandemic.** A's crit 104 on a 51 tick, his Corruption 121 on 60 and Agony 86 on 42; B's crit tick 2.0x too; the Shadow Bolt crits 1.5x. 2 crits in 23 ticks. | A and B | `pandemicAffectedSpellCodes` includes Wrack | aligned |
| L5 | **The three DoTs tick twice as fast while Wrack channels.** Gaps between ticks on the channelled target: Bane of Agony median 1.012 s inside a channel (21) against 2.007 s outside (18); Corruption 1.51 s (14) against 3.006 s (13); Siphon Life 1.50 s (8) against 3.03 s (2). | A and B | `WrackDoTClockRate` 2 on Siphon Life, Corruption and Bane of Agony | aligned |
| L6 | **The DoT keeps its tick count.** A's Bane of Agony (fight 651) ran all 12 ticks, stepping through the ramp at ticks 1 to 4, 5 to 8 and 9 to 12 (30, 42, 55 a tick), in 18 s instead of 24, across two channels. | A | the clock rate rescales the expiry with the ticks | aligned |
| L7 | **The tick in flight when a channel starts or ends.** Blizzard's note (the engine's model) has the remaining time halve at the start and double at the end. Of the 11 edges where the two readings differ, Bane of Agony's in-flight tick kept its own time at both edges (5 of 5: 2.0 s after the last tick at a start, 1.0 s at an end); Corruption's split (one fits the note, one keeps its time, three fit neither); Siphon Life's one end kept 1.5 s. | A and B (`results/clock.txt`) | `core/dot.go` `SetClockRate` rescales the in-flight tick | **held (H1)**: 11 edges from two players, and either reading only moves a tick's time: averaged over channel starts and ends the shift is zero, so no DPS figure moves |
| L8 | **Nightfall from Wrack.** Shadow Trance came up on A during his channels and he spent it on the instant Shadow Bolt that ended the channel. | A | `nightfallTriggerSpellCodes` includes Wrack | aligned |
| L9 | **Demon crit, our reports.** Succubi 2 crits in 266 swings (0.8%, 16 owners at 20 to 36), Voidwalkers 6 in 363 (1.7%, 20 owners), Imp Firebolt 3 in 49, Lash of Pain 1 in 23. | eight reports (`results/demons.txt`) | `pet.go`: 0.7% plus the owner's gear crit; Firebolt 5.9% | consistent; the 12 Nov H1 (one Succubus at 3.9%) stays held for a level 60 Succubus on a boss |

Nothing at 50 or above exists in any source: our highest warlock is 42 (a Hellfire cast in Desolace), foreverlogs' 31 to 34, and Forever's raids open 9 Dec. Every 50 to 60 and raid check against a log stays open.

## Gear: the 13 Nov item hotfixes

Blizzard's 13 Nov notes change two rings every warlock band at 60 wears or lists (the ranking at 5060a7eb7 puts Channeler's Ring 272406 in a finger slot of Affliction raid, Demonology bare and raid, and Destruction raid, both factions; Destruction lists Blessed Band of Light 272407 as an alternative). Push 114062 gives Channeler's Ring Intellect, hit and spell power and takes its Spirit away, and removes the Band's hit. Data run b5bdd533b took the push, but the Wowhead equip fill put the removed stats back (Wowhead still lists the old rings). The Fire and Frost lane fixed the fill while this lane ran (9e6cdc3c2, MR110). For the warlocks the fill only adds 8 Spirit, worth 8 mana a Life Tap. The ring's new 9 spell power is the bigger change: the warlock sets were ranked at 00:07 UTC 14 Nov, before b5bdd533b gave it any. MR110 now names the warlock bands, and the re-rank after the next data run settles their ring slots.

## Confirmed and fixed

None.

## Held and open

- **H1** (L7): the tick in flight at a Wrack edge. Blizzard's note and the engine rescale it. Bane of Agony's logged tick keeps its time (5 of 5), and Corruption's evidence splits. No DPS effect on average. A level 60 Affliction log with 20 or more channels would settle it per DoT.
- **H2** (12 Nov H1): the Succubus's melee crit, unchanged (L9).
- **Open:** every 50 to 60 and raid check against a log; the 12 Nov leftovers (a Searing Pain and Demonic Brand leveling style for Demonology, the level-by-level Affliction talent list, the Execute idea).

## Measure at refresh

- MR110 (the Fire and Frost lane's) gains the warlock bands: Affliction raid 60, Demonology bare and raid 60, Destruction raid 60, both factions. Their ring slots, weights and set_dps follow the next data run (Channeler's Ring 9 spell power, no Spirit; Blessed Band of Light without hit).

## Sim changes in this landing

None. The engine pin stays 037798db1, and no guide, card, APL or golden changed, so no wow-player pass was due. The fork worktree `lane/sa-warlock-70938` sits at the pin with no commits; it was removed.

## Gates

- Fork, on the box at 037798db1: `go test -tags=with_db` for `sim/warlock/...`, `sim/core`, `sim/conformance`: ok (`results/box-fork-pin.txt`).
- Site: docs and lane scripts only. No code, data or golden changed. Rebased on origin/main and pushed once.

## Sources

- Our reports 3sdrcwdgo2yi, ju2t73py6pib and the six others, read only.
- foreverlogs.gg `/api/reports/public` and reports 2855 to 2857 (13 Nov).
- Blizzard, "WoW: Forever Hotfixes", https://us.forums.blizzard.com/en/wow/t/2403040 (posts to 14 Nov 01:54 UTC).
- Client 1.60.1.70938 and 1.60.1.70790: `spellconst/warlock.json`, `talents/warlock.json` (Soul Siphon, Pandemic, Nightfall and Wrack texts), `raw/wago-push/`.

# Sim accuracy 2026-11-13: the healers' rotations on the raid fight

Lane `healer-apl`, open item O2 of `healer-raidfight.md`: every healer's list waited for health thresholds written for the 300 s throughput fight (`heal-profile.json`), so on the raid-shaped fight (`heal-raid-profile.json`, 110 s, fitted to 33 Molten Core and Onyxia kills) each sat mostly idle and left a different share of the damage unhealed at the kill. The Holy paladin was last (-16.7%) where the logs put it level with the priest, and overheal ran from 2% (paladin) to 74% (druid) against the logs' 39 to 48%. Worktree from origin/main; sims on the box (`tools/remote-test.sh`, `nice -n 15`, `GOMAXPROCS=5`), one at a time. Measured on client 1.60.1.70790 (the site's active build; 70938's data is in the tree but not active), engine 8f61b9e6a for the trials and the lane engine b0a43f004 (on 70dd8eb24) for the confirmation; the two differ only by the Hammer of the Righteous fix, which no healer casts.

**Totals: 0 sim bugs; one rotation per healer retuned, right on both fights (no second raid rotation needed). Each healer's raid-fight figure moves beyond error toward the logged play; the 300 s fight moves by at most 0.4% at raid 60 and within or about its error at bare 60 and bare 40. Engine pin after: b0a43f004 (the rotation copies; no engine code). One Measure-at-refresh row (MR103).**

## 1. What the logged healers cast

`healer-apl/castmix.py` reads the healer-mana sample (`healer-mana/logs-kills.json`) through `wcl.py`'s cache only (every request one `collect.py` already made, no key): per class, each heal's casts a minute and ranks, its share of casts and of effective healing, who it lands on, and active time. Full tables in `healer-apl/logs-castmix.md`.

| class | casts a minute | active time | what they cast | overheal (median) |
|---|---|---|---|---|
| Paladin | 24.8 | 0.87 | Flash of Light 91% of casts and 89% of healing, rank 6 (97%); Holy Light 3%, Holy Shock 2% | 44% |
| Priest | 18.7 | 0.82 | Flash Heal 60% of casts (rank 7 40%, ranks 1 to 5 the rest), Renew 10%, Shield 9%, Heal 8%, Greater Heal 2%, Prayer of Healing under 1% | 44% |
| Shaman | 17.0 | 0.83 | Chain Heal 56% (rank 1 57%, rank 3 30%), Lesser Healing Wave 20%, Healing Wave 8%; Healing Stream 0.5% of healing | 39% |
| Druid | 18.7 | 0.89 | Rejuvenation 51% (rank 10 68%), Healing Touch 14% (ranks 3 and 4), Regrowth 14%, Swiftmend 3% | 48% |

A logged healer heals at small deficits all fight: the paladin presses Flash of Light every 2.4 s on whoever has a hole. The sim's lists before this lane (`healer-apl/sim-r1.txt`, the `base` jobs): the paladin cast Flash of Light 8.9 times a minute, only on the tank below 90% and the party below 80% (two to three of the group's 176 hits on a 5,000 bar), with 2.4% overheal; the shaman 1.0 Chain Heal a minute (anyone below 90%); the Holy priest kept Renew and Prayer of Mending rolling on a tank who was nearly always full (Renew healed 194 of 12,295 a fight), its single-target heals waiting for a member under 65%; the druid kept Regrowth and rank 10 Rejuvenation on the tank unconditionally (74% overheal) and Rejuvenated the party only below 70%. Era has no Prayer of Mending, Riptide, Penance, Bloom or Lightwell, so the comparison is by play (heal size against the hole, who gets it, how much is wasted), not by spell for spell.

## 2. Trials

Each trial is a curated list with thresholds moved or lines appended (`healer-apl/variants.py`, run with `spec-breakdown -heal -rotation`), raid 60 Alliance, seed 1, 2,000 iterations, on both fights (`jobs-r1.txt` to `jobs-r3.txt`, outputs `sim-r1.txt` to `sim-r3.txt`, summary `summary.txt`). Below 60 and bare were checked on the adopted ones.

| healer | trial | raid fight | overheal | 300 s fight | kept? |
|---|---|---|---|---|---|
| paladin | p1: tank 95%, party 90% | 129.8 | 21% | -0.2% | no, p2 closer to the logs |
| paladin | **p2: tank 97%, party 93%** | **134.8** | **46%** | -0.2% | **yes** |
| shaman | s1: Chain Heal 95%, Riptide 95% | 137.6 | 60% | 0.0% | no, overheal 21 points over the logs |
| shaman | s3: Chain Heal 93% | 135.8 | 46% | 0.0% | no |
| shaman | s2, s4: plus a Lesser Healing Wave top-off | as s1, s3 | | | no: never cast, Chain Heal fires first |
| shaman | **s5: Chain Heal 92%, Riptide 95%** | **135.0** | **41%** | 0.0% | **yes** |
| Holy | h1: Flash Heal top-off, pace 0.9 | 130.7 | 67% | -1.0% | no, costs the 300 s fight |
| Holy | h2: h1 plus party Renew | 135.3 | 70% | -1.0% | no, overheal and the 300 s cost |
| Holy | **h3: tank Renew below 95%, top-off on full pace** | **130.8** | **54%** | -0.1% | **yes** |
| Discipline | h1 / h2 | 127.5 / 133.4 | 50% / 57% | -1.1% / -1.4% | no, the 300 s cost |
| Discipline | **h3 (as Holy)** | **126.8** | **50%** | -0.1% | **yes** |
| druid | d1: party Rejuvenation 70% to 90% | 133.9 | 76% | -2.8% | no, the 300 s cost |
| druid | d2: d1 plus tank Regrowth below 90% | 133.5 | 67% | -2.6% | no |
| druid | d3: old lines kept, paced 90% party lines, tank Regrowth below 90% | 133.5 | 67% | -0.2% | no, d4 closer |
| druid | **d4: d3 plus tank Rejuvenation below 97%** | **133.2** | **63%** | -0.4% | **yes** |

The pacing rule is what makes one list right on both fights. A top-off line that waits on the mana pace (mana % at least boss health %, "while mana holds") never fires on the 300 s fight, where every healer runs dry and the bar is behind the clock from early on, and always fires on the raid kill, where every healer ends with 90 to 99% of its mana. The thresholds that moved without a pace guard (the paladin's Flash of Light, the shaman's Chain Heal and Riptide) were already passed by every hit of the 300 s fight's 2,645 tank hits and 690 pulses, so they cost it nothing.

## 3. Adopted (data/curated/apl, synced to the fork and sim/request/apl)

- **Holy paladin**: Flash of Light on the tank below 97% (was 90%) and on yourself and the party below 93% (was 80%). Rank 5 at raid 60 as before (rank-choices row); the logs' rank 6 was not tried, a rank row is its own judgement.
- **Restoration shaman**: Chain Heal when anyone in the party is below 92% (was 90%), Riptide on the tank below 95% (was 90%).
- **Holy and Discipline priest**: Renew on the tank only once he is below 95% and it is not on him (was whenever it was not on him); six appended Flash Heal lines (rank 7): the tank below 90%, yourself and the party below 85%, each only while mana holds.
- **Restoration druid**: the tank's Regrowth only below 90% and its rank 10 Rejuvenation only below 97% (both were unconditional); five appended rank 7 Rejuvenation lines on yourself and the party below 90%, while mana holds, beside the unpaced 70% lines (round 13's rule kept).

## 4. Before and after (lane engine b0a43f004, `sim-final.txt`)

Raid fight, raid 60, seed 1, 2,000 iterations; the 300 s fight's change as relative gains (prose never quotes nightly figures).

| healer | raid fight before (Alliance / Horde) | after (Alliance / Horde) | change | overheal before, after, logs | heal casts a minute after (logs) | 300 s fight raid 60 | bare 60 | bare 40 |
|---|---|---|---|---|---|---|---|---|
| Restoration shaman | 131.6 / 131.8 | 135.0 +/- 0.09 / 135.0 | +2.6% | 27.6%, 40.9%, 39% | 4.1 (17.0) | 0.0% | 0.0% | -0.3% |
| Holy paladin | 109.6 / 109.5 | 134.8 +/- 0.09 / 134.8 | +23.0% | 2.4%, 46.1%, 44% | 19.8 (24.8) | -0.2% | +0.1% | +0.1% |
| Restoration druid | 111.9 / 111.8 | 133.2 +/- 0.11 / 133.1 | +19.0% | 74.2%, 62.5%, 48% | 10.3 (18.7) | -0.4% | +0.1% | +0.4% |
| Holy priest | 125.6 / 125.6 | 130.8 +/- 0.14 / 130.8 | +4.1% | 65.4%, 53.9%, 44% | 8.0 (18.7) | -0.1% | +0.1% | -0.1% |
| Discipline priest | 119.2 / 119.4 | 126.8 +/- 0.14 / 127.1 | +6.4% | 44.9%, 49.9%, 44% | 7.2 (18.7) | -0.1% | 0.0% | -0.3% |

**The raid-fight order, against the logs.** Sim after: shaman 135.0, paladin 134.8 (-0.1%), druid 133.2 (-1.3%), Holy 130.8 (-3.1%), Discipline 126.8 (-6.1%); every gap beyond the error (the closest, paladin under shaman, is 0.2 against a combined 0.13). Logs (share of the kill's mean healer): shaman 1.07, priest 1.01 (-6%), paladin 1.00 (-7%), druid 0.84 (-22%). Before: shaman, Holy, Discipline, druid, paladin, the paladin last by 16.7%.

- The paladin is fixed: it plays the logged paladin's game (Flash of Light at every hole, 46% overheal against 44%) and comes up from last to level with the shaman.
- The shaman stays first; its overheal now matches the logs (41% against 39%).
- The priests come closer on overheal (Holy 65% to 54%) and keep their order; the logs put the priest level with the paladin, the sim puts it 3 to 6% under. Holy's single-target casting is still Prayer of Mending on cooldown where the logged priest is mostly Flash Heal: Era has no Prayer of Mending, so how Forever priests split the two is not in these logs.
- The druid is the miss that remains. It plays closer to the logs (overheal 74% to 63%, the party HoTs the logged druid spreads) and sits third, where the logs put it last by 22%. Every healer now heals 88 to 94% of the 143.5 a second before the kill; the fight's ceiling compresses the gaps, and the logged druid's deficit is a share of the raid's healing taken by others' faster heals, which a one-healer fight cannot show (Open O3).
- Casting: the sim's healers still cast less than the logged ones (Holy 8 heals a minute against 18.7). Forever's passive healing (Prayer of Mending bouncing, Healing Stream, Lightwell, Bloom) and the profile's six targets leave fewer holes than a 40-player raid; the paladin, with no passive healing, is the one whose cast rate comes close (19.8 against 22.6 Flash of Light a minute).

## 5. Gates

- Box, engine b0a43f004: `go test ./request/ ./internal/inproc/` (green after: `TestRestorationDruidPartyRejuvenationIsNotPaced` now holds the 70% lines unpaced and the 90% top-off lines paced; `TestPriestTankFlashHealWaitsOnTheManaPace` counts two paced tank Flash Heal lines; the five healer ladder goldens regenerated with `FOREVER_UPDATE_GOLDEN=1`, changes within 1% a rung); the fork's `sim/paladin/holy`, `sim/druid/restoration`, `sim/priest/healing`, `sim/shaman/restoration` with `-tags=with_db` from a writable module copy, green.
- Ladder strict-mode notes (informational, the test skips them): three new zero-cast rungs, Holy Light rank 6 at paladin 38 and Holy Shock at paladin 50 (paced lines the busier Flash of Light now pre-empts), Flash Heal at Holy 40.
- `make apl-check` green; data `pytest tests/test_addonrotation.py test_specs.py test_apl.py test_card_note_levels.py test_panel_when_levels.py` (175 passed; panel lines held to 60 characters).
- Box vitest `src/lib/bis/ src/lib/guides/ src/content/guides/ src/lib/app-data/` against a throwaway local `addon-data` regen (dropped before the push): 1,821 passed. Against the committed addon-data, `apl-level-gate.test.ts` fails for Discipline (the published 70790 lines predate the appended steps, so the alignment falls back and shows gated Heal rows) until data.yml's push job regenerates addon-data.json and Data.lua, as it is built to (data.yml's note on rounds 3 and 4). `rotation-view.test.ts` reads the frozen 70374 list, which never gains the six Flash Heal lines, so its Holy band 20 cases add the lines the next publish carries (`withRaidTopOff`). eslint and prettier on the touched files.

## 6. heal-builds (the raid-search lane's T1)

The rotation work used no `heal-builds` comparison. The tool runs the 300 s heal profile, on which these rotations move every healer by 0.4% or less at raid 60 and within about the error at bare 60, so a build comparison it makes does not move because of them. The comparisons themselves are hand-picked, and none is committed to re-run; re-running the healer guides' build choices on enchanted gear is still the raid-search lane's Open item, for a healer build lane.

## 7. Reviews

wow-player, `healer-apl/review/wow-player.md` (the prose diff and card changes in `review/`): pass 1 SHIP WITH FIXES for all five (overheal apology, the "Pace, not health bars" heading, Riptide and Chain Heal placement, "goes back" for Renew, the Mana and Flash Heal contradictions in the priest guides, the druid's pre-pull Rejuvenation, "on pace" not a game phrase); pass 2 SHIP for all five.

## Findings

| # | finding | evidence | status |
|---|---|---|---|
| A1 | The healers' lists waited for deficits the raid kill never reaches | sections 1 and 4 | fixed: section 3, every raid-fight move beyond error toward the logs |
| O3 | The druid sits third on the raid fight where the logs put it last by 22%; the priests 3 to 6% under the paladin where the logs have them level | section 4 | open: a one-healer fight cannot show healing taken by other healers; a multi-healer raid fight, or a sniping model, is its own lane |

## Shas

- Fork b0a43f004 (`lane/healer-apl`, fast-forwarded onto master): the five rotation copies.
- Site: the rotation, tests, goldens and guide commit; `engine: pin b0a43f004`; the docs commit with this report, the ledger row and MR103.

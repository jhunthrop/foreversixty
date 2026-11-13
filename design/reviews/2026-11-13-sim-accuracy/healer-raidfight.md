# Sim accuracy 2026-11-13: the raid-shaped healer fight

Lane `healer-raidfight`, the owner's ruling on O1 of `healer-mana.md` (option 2): keep `heal-profile.json` for the BiS ranking, where its job is to make gear separate, and rank the /tiers Healers tab on a second, raid-shaped fight fitted to the logs. Worktree from origin/main ce64e36f0; engine pin 79358b3c2 (unchanged). Sims on the box (`tools/remote-test.sh`, `nice -n 15`, `GOMAXPROCS=5`), one at a time.

**Totals: 0 sim bugs; a new profile, a new published metrics block and a page change. No engine, card, preset, BiS profile, guide or rank change; no BiS rank, data run or simdb rebuild. One Measure-at-refresh row (MR100): the next ranking publishes the figure and the page switches to it by itself.**

## 1. The profile (`data/curated/heal-raid-profile.json`, id `raid60-logs`)

Fitted to the log sample of `healer-mana.md` (Classic Era Hardcore Fresh, vanilla.warcraftlogs.com partition 8, every Molten Core boss and Onyxia, 33 kills, 342 healer-kills; `healer-mana/logs/collect.py`, `logs-kills.json`, `logs-summary.txt`). The file states the sample in `log_sample`, and the loader (`sim/request/raidhealprofile.go`) refuses a file whose fight length, tank hit or pulse is not the derivation of it.

| figure | log basis | profile |
|---|---|---|
| length | 90th percentile kill 107 s (median 61), rounded up to 10 s | 110 s |
| damage one healer answers | median over kills of (main tank taken + every non-tank player's taken) / healers (median 11): 143.7 a second | 143.5 a second |
| tank share | median 0.264 of that | 360 a hit (143.7 x 0.264 x 9.5 s), about 38 a second |
| tank spikiness | main tank's largest 5 s over his mean 5 s, median 2.33 | one hit every 9.5 s, 25% spread: 2.3 (`TestRaidHealProfileTankSpikeMatchesTheLogs`, collect.py's own measure on the engine's schedule) |
| the group | the rest, 106 a second; no member spikiness was sampled | 176 on 3 of the 5 group members (the healer among them) every 5 s, `heal-profile.json`'s cadence |
| health | the logs carry none | `heal-profile.json`'s 9,500 tank and 5,000 members |

It has no bands and no tuning: it is a level 60 raid kill, run only on band 60 entries.

## 2. The ranker

`sim/cmd/leveling-bis/score_heal.go`: the healer engine loads both profiles; `healerFieldsFor` runs every band 60 entry's final set (raid and bare, both factions) once more on the raid fight (2,000 iterations, seed 11, as the other metrics) and publishes `metrics.raid_fight` beside the existing figures: `profile`, `duration_sec`, `hps` (effective, a shield for what it soaked), `hps_error`, `raw_hps`, `overheal_pct`, `mana_left` and `mana_left_pct` (the bar when the kill ends: the pool from the engine's computed stats plus every point gained, none over the cap, less every point spent; `score.HealingResult.WithManaPool`). The picks, `hps`, `set_dps` and the bounds are untouched: they stay the BiS profile's. It runs inside `attachHealerFields`, so each healer class's leg of the per-class matrix (ac987b9ec) publishes it with no workflow change. `spec-breakdown -heal -raid-fight` prints the same fight, with overheal and mana left now on every `-heal` report.

Tests, fixture first: `TestAttachHealerFieldsRunsTheRaidFightAtBand60` (one spec, fake backend: the block, its request at 110 s exactly and 2,000 iterations, the BiS figures unchanged, the json key), `...SkipsTheRaidFightBelow60`, `TestWithHealerEngineLoadsTheProfileForAHealer`; `request`: derivation, refusals, spike, Attach; `inproc`: `TestHealingRunOnTheRaidFightLeavesOverhealAndMana` (Holy, 3 iterations).

## 3. The five healers on the new fight

Raid 60 sets as published (70790 files), seed 1, 2,000 iterations (`healer-raidfight/sim-raid-and-bis.txt`, `sim-horde-bare.txt`):

| healer | raid fight HPS (Alliance) | vs #1 | Horde | overheal | mana left | bare 60 | today's list (BiS fight, 300 s, as published on 70790) | vs #1 |
|---|---|---|---|---|---|---|---|---|
| Restoration shaman | 131.6 +/- 0.1 | top | 131.8 | 27.6% | 99.7% | 131.6 | 889.7 | -13.5% (=1, tied) |
| Holy priest | 125.6 +/- 0.2 | -4.5% | 125.6 | 65.4% | 95.5% | 123.8 | 871.3 | -15.3% (3rd) |
| Discipline priest | 119.2 +/- 0.3 | -9.4% | 119.4 | 44.9% | 95.7% | 117.3 | 604.2 | -41.2% (5th) |
| Restoration druid | 111.9 +/- 0.2 | -14.9% | 111.8 | 74.2% | 97.4% | 101.9 | 1,028.3 | top (=1) |
| Holy paladin | 109.6 +/- 0.2 | -16.7% | 109.5 | 2.4% | 99.5% | 103.3 | 719.6 | -30.0% (4th) |

Every gap is far beyond the error (the closest pair, druid over paladin, is 2.3 against a combined error of 0.27), so all five order (no tie; today's list ties the druid and shaman at the top). The order turns nearly over: the druid goes from first to fourth, Discipline from last to third. On this fight no healer is short of mana (95 to 100% left), so effective HPS is how much of the incoming damage each card heals before the kill: the shaman leaves about 12 a second of the 143.5 on the bars, the paladin about 34. That is set by the cards' thresholds, which were written for the 300 s fight: the paladin's list waits for a member well down before it casts (2% overheal, the bar nearly full), and the priests' and shaman's lists never single-target the healer itself (group heals only), so their own pulse damage waits for the next group heal. Open as O2.

## 4. Against the logs

The same 33 kills, each healer's effective healing over the kill's mean healer, by class (no Discipline in the sample: Era priests are all logged Holy):

| class | healer-kills | share of the kill's mean, mean (median) | overheal median | sim raid fight vs the shaman |
|---|---|---|---|---|
| Shaman | 57 | 1.07 (0.99) | 39% | top |
| Priest | 166 | 1.01 (0.97) | 44% | -4.5% |
| Paladin | 88 | 1.00 (0.98) | 44% | -16.7% |
| Druid | 31 | 0.84 (0.84) | 48% | -14.9% |

Logs: shaman, then priest and paladin level, then the druid about 22% behind the shaman. Sim: shaman, priest, druid, paladin. Three of the four classes sit in the logged order with the druid's gap close (15% against 22%); the paladin is the one miss (logged level with the priest, simmed last). Today's list puts the logged last class (the druid) first, so the raid fight is the closer of the two. Overheal: logged 39 to 48% for every class; the sim gives the druid 74%, Holy 65%, shaman 28% and the paladin 2%. One healer alone heals its share, so nobody snipes its heals, which keeps the reactive cards' overheal under the logs' and lets the HoT druid's run over it.

## 5. The page

`/tiers/healer` ranks on `metrics.raid_fight.hps` and its error once every healer band of both factions publishes it (`healerBasisOf`, `web/src/lib/tiers/tier-list.ts`), and on today's figure until then, never a mix. The raid fight has no tuning, so only its error ties a pair. The page says which: the lead opens with the fight ("Ranked on one healer's share of a Molten Core or Onyxia kill, so it reads lower than your meter." or, until the raid figure is published, "Ranked on a stand-in fight, not a boss, that runs every healer dry. A Molten Core or Onyxia kill replaces it with the next update."), then a short form of the ce64e36f0 sentence ("HPS is effective healing. Shields count for what they soak, and buffs on others are not counted."), whose long form with the Power Infusion example stays in the notes; the meta line owns the fight's length (110 s or 5 min); the notes columns and a tied row's hover explain the fight the list is on. The app tiers file reads the same copy. Tests: `tier-list.test.ts` (raid ranking, raid tie, fallback, one basis for both factions), `tier-load.test.ts`, `tier-copy.test.ts` (both leads, notes, tie reason, no dashes), `tests/e2e/tiers.spec.ts` (the lead names a fight).

Captures and verdicts: `design/reviews/2026-11-13-healer-tiers/raidfight/` (390, 1024, 1440, 2000; the page today, and the page with the lane's measured raid-fight figures laid into the box run's copy of the 70790 files, capture only, never committed).

Verdicts (`raidfight/ux-designer.md`, `raidfight/wow-player.md`; first passes kept as `*-pass1.md`): both SHIP WITH FIXES on the first copy (the lead ran six lines on a phone and stated the length twice in two units; the raid sentence did not name the raid or say the figure is a share; the fallback's "BiS fight" and "next ranking" were site talk and gave the four-figure HPS no reason). On the copy above, ux-designer SHIP and wow-player SHIP, polish only (a one-word last line at 390 in the raid state; the tab strip sits a little lower on Healers than on the leadless Damage and Tanks tabs).

## Findings

| # | finding | evidence | status |
|---|---|---|---|
| R1 | A raid-shaped fight fitted to the logs, published beside the BiS figures | sections 1 to 3 | built; publishes at the next ranking (MR100) |
| O2 | On the raid fight the order is the cards' thresholds: each heals a different share of the damage before the kill (shaman 92%, paladin 76%), because the lists were written for a fight that is never quiet; the paladin's waits for a member well down, the priests' and shaman's never single-target themselves. The logs put the paladin level with the priest | section 3, section 4 | open: a card pass on the raid fight (tank, group and self thresholds) is its own lane; the adoption rule applies |

## Gates

Box: `go test ./request ./score`, `./internal/inproc -run HealingRun`, `./cmd/leveling-bis -run "Heal|Healer|Incumbent|RaidFight"`, `go vet` on the touched packages; `npm test -- src/lib/tiers/ src/lib/app-data/ src/lib/bis/heal-view.test.ts` (one failure, `app-files.test.ts` "carries each guide's current build codes", identical on origin/main, not this lane), `npm run check` (0 errors), `npx playwright test tests/e2e/tiers.spec.ts` (58 passed), eslint and prettier on the touched files.

## Shas

- bbc07527e feat(sim): the raid-shaped healer fight, the ranker block, the page (on d8d421995).
- The docs commit that carries this report, the captures, the verdicts, the ledger row and MR100.

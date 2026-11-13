# Sim accuracy 2026-11-13: the raid 60 healer mana model

Lane `healer-mana`, from the Discipline lane's finding (`discipline.md`): under the raid 60 heal fight (300 s, `data/curated/heal-profile.json`) every healer runs dry, Discipline first. The question: is the raid 60 healer preset short of the mana a raiding healer really has, or is the fight the cause? Site from origin/main 05d80dada, engine pin f46237344 (unchanged). Sims on the box (`tools/remote-test.sh`, `nice -n 15`, `GOMAXPROCS=5`), one at a time.

**Totals: 0 bugs. The preset is right: every mana source a Forever raider can bring is in it, and the sim uses each at its cooldown. The fight is the cause. It asks one healer for about eight times the damage a logged Phase 1 raid healer faces, for five times as long, so every healer casts flat out until the bar is empty and almost nothing overheals. That is the heal profile's stated design (gear has to separate), not a model error, and changing it is the owner's call (O1). One page change: the healer tier list now says under its title what HPS does not count. No engine, card, preset, profile, guide or rank change; no Measure-at-refresh row.**

## 1. The preset against what a raiding healer has

`sim/request/raid-consumes.json`, every healer at 55 to 60 (`spec-breakdown -heal`, raid 60 Alliance seed 1, 2,000 iterations; `healer-mana/sim-raid60.txt`):

| source | in the preset | in the sim (per 300 s fight) | Forever client |
|---|---|---|---|
| Major Mana Potion (13444, 1,350 to 2,250, 2 min) | all five | 3.0 drunk, 5,400 mana (the cooldown allows 3) | item 13444, spell 17531 |
| Demonic Rune (12662, 900 to 1,500, own 2 min timer) | all five | 3.0 used, 3,600 mana | Item 12662 class 7 (trade goods, so not in `simconsumes.json`, which keeps class 0); spell 16666 |
| Dark Rune (20520) | no | | in the client (class 7), but it shares the rune cooldown, so a second rune type adds nothing |
| Flask of Distilled Wisdom (2,000 mana) | all five | in the pool | 13511 |
| Greater Mageblood Elixir (20 mp5, Forever's) and Elixir of Sages | all five | in mp5 | 250341, 250338 |
| Brilliant Mana Oil (12 mp5) | all five | in mp5 | 20748 |
| Food | Sage's Tea (44 healing) | | Nightfin Soup is 22 spell damage in this client (`TestNightfinSoupMatchesTheClient`), no mp5; no Forever food gives mp5 a healer would take over it |
| Blessing of Wisdom, improved Mana Spring Totem | all five | in mp5 (Disc 108, Holy 124, shaman 115, paladin 120, druid 110) | |
| Judgement of Wisdom | all five | 0: it pays the attacker, and a healer does not hit the boss | correct, and the logs agree (below) |
| own class mana | Innervate (druid, on itself, 1 a fight), Mana Tide (shaman, own, 5 ticks), Illumination, Litany of Light | all used | |
| Innervate from another druid, Mana Tide from another shaman, Power Infusion | no | | externals; the raid default carries none (round 5, `TestNoRaidDefaultCarriesPowerInfusion`), and the logs show why (below) |
| Five-second rule, Spirit regen | engine | Disc regains 48 a second casting and 9 more out of the rule; no healer leaves the rule for long, because every target is always hurt | |

Every consumable a Forever raider brings for mana is in the preset and is used as often as its cooldown allows: 3 potions and 3 runes in 300 s, 9,000 mana, about a third of each healer's income.

## 2. The logs

No level 60 Forever log exists (raids open 9 Dec; foreverlogs.gg tops out at 43; no Warcraft Logs Forever partition; our backfilled reports are open world under 40, see `discipline.md`). The nearest public analogue of a Forever Phase 1 raid is Classic Era Hardcore Fresh on vanilla.warcraftlogs.com (partition 8: Phase 1 gear, no world buffs). Sample (`healer-mana/logs/collect.py`): every Molten Core boss and Onyxia, 3 kills each from the hps rankings at pages 2, 5 and 10, one kill per report, hash-ordered, raid size 30 or more: 33 kills, 342 healer-kills. Summaries in `healer-mana/logs-summary.txt`, the kills in `logs-kills.json` (no player names).

**L1, fight length.** Median kill 61 s, 75th percentile 78 s, 90th 107 s; the longest 339 s (one Onyxia). Onyxia's median is 136 s, Majordomo's 105 s, every other boss under 80 s. The profile's 300 s is longer than 32 of the 33 kills.

**L2, damage a healer faces.** Main tank 446 damage a second (25th to 75th: 368 to 542), every other player 31 a second, 11 healers a kill: about 144 damage a second for each healer to answer. The profile gives one healer 1,215 a second (the tank 2,645 every 3.3 s, 801 a second, plus 414 a second of pulses), about 8 times as much. Logged healers' raw healing is a median 281 to 374 a second (90th percentile 459 to 510); the sim's healers put out 644 to 1,058 raw HPS for the whole fight.

**L3, spikiness.** The main tank's largest 5 s of damage is a median 2.33 times his mean 5 s (75th percentile 3.00). The profile's tank hit (2,645 with 25% spread every 3.3 s) peaks at about 1.65 times its mean. Real tank damage is spikier and lighter.

**L4, overheal.** Logged overheal, median by spec: Holy priest 44% (30 to 54), Holy paladin 44% (35 to 58), Restoration shaman 39% (27 to 47), Restoration druid 48% (36 to 60). The sim: 0.5 to 5.4% at raid 60 (raw HPS over effective: Holy 5.4, Discipline 3.0, druid 2.8, shaman 0.7, paladin 0.5). A real raid healer spends close to half her healing on targets that are already full; the profile has none.

**L5, consumables and mana by source**, per minute of boss time:

| healer | healer-kills | minutes | Major Mana Potions a min | runes a min | top logged mana sources a min |
|---|---|---|---|---|---|
| Holy priest | 166 | 197 | 0.22 | 0.11 | Major Mana Potion 389, Blessing of Wisdom 231, Dark Rune 94, Demonic Rune 40, Mana Tide 24 |
| Holy paladin | 88 | 101 | 0.01 | 0.01 | Illumination 1,015, Blessing of Wisdom 401, Major Mana Potion 18 |
| Restoration shaman | 57 | 69 | 0.25 | 0.10 | Major Mana Potion 457, Mana Tide 285, Dark Rune 84, Mana Spring 45 |
| Restoration druid | 31 | 40 | 0.13 | 0.13 | Major Mana Potion 236, Blessing of Wisdom 168, Demonic Rune 123 |

The sim drinks 0.6 potions and 0.6 runes a minute for every healer: about 2.5 times the logged priest's and shaman's potions, 5 to 6 times their runes, and 60 times the logged paladin's. Judgement of Wisdom gave the logged healers 1,416 mana in 407 healer-minutes, about 3 a minute. The logs carry no mana bar (`classResources` reads 0 of 0 on every event of every kill), so "how many healers end dry" cannot be read; the consumable rate says the answer: a real priest or shaman reaches for a potion about once in four minutes of boss time, the sim's once in 100 seconds.

**L6, externals** (`logs/externals.py`, the buffs table by target): in 197 priest-minutes, 101 paladin-minutes and 69 shaman-minutes, no healer but a druid received an Innervate (the druids' 10 are their own); priests received 7 Power Infusions. Leaving other players' Innervate and Mana Tide out of the default is what these raids do.

## 3. What the sim shows about the fight

Fight length against HPS and mana (raid 60 Alliance seed 1, 2,000 iterations, `-duration`; `healer-mana/sim-durations.txt`). Effective HPS:

| fight | Resto druid | Holy priest | Resto shaman | Holy paladin | Discipline |
|---|---|---|---|---|---|
| 60 s | 1,026.9 | 1,117.0 | 989.9 | 874.5 | 1,009.8 |
| 90 s | 1,092.7 | 1,149.3 | 994.7 | 895.8 | 959.2 |
| 120 s | 1,121.3 | 1,081.7 | 990.6 | 866.3 | 821.8 |
| 180 s | 1,112.1 | 1,009.8 | 991.7 | 829.7 | 734.7 |
| 240 s | 1,061.6 | 901.0 | 934.8 | 739.1 | 648.6 |
| 300 s (the profile) | 1,028.3 | 881.9 | 889.6 | 719.7 | 624.6 |

At the logged 90th percentile kill (about 107 s, between the 90 s and 120 s rows) Discipline is 17 to 24% behind Holy instead of 29%, and the order of the top three changes (Holy, druid, shaman at 90 s; druid, shaman, Holy at 300 s). The shaman, whose Water Shield pays for itself, is the only flat line; every other healer is ranked on how long its bar lasts.

## Before and after (raid 60, unchanged)

No sim input changed, so before and after are the same run: effective HPS and the time the mana lasts, raid 60 Alliance seed 1, 2,000 iterations, engine f46237344.

| healer | effective HPS | raw HPS | mana lasts |
|---|---|---|---|
| Restoration druid | 1,028.3 | 1,058.3 | 274 s |
| Restoration shaman | 889.6 | 895.5 | 229 s |
| Holy priest | 881.9 | 932.1 | 126 s |
| Holy paladin | 719.7 | 723.0 | 201 s |
| Discipline priest | 624.6 | 644.0 | 95 s |

## Findings

| # | finding | evidence | status |
|---|---|---|---|
| A1 | The raid preset carries every mana consumable and raid mana buff a Forever raider brings, and the sim uses each at its cooldown | section 1; `sim-raid60.txt` | aligned |
| A2 | Judgement of Wisdom gives a healer nothing | engine `JudgementOfWisdomAura` pays the attacker; logs 1,416 mana in 407 healer-minutes | aligned |
| A3 | Other players' Innervate and Mana Tide stay out of the default | L6: no non-druid healer received an Innervate in 367 healer-minutes; round 5's externals rule | aligned |
| O1 | **The heal profile is a throughput test, not a raid.** One healer faces about 8 times a logged healer's damage for about 5 times a logged kill's length; real overheal is 39 to 48%, the sim's 0.5 to 5.4%; the sim drinks about 2.5 times the potions and 5 times the runes a minute. So the healer list ranks sustained healing until the bar is empty, where a shield healer and a Holy paladin lose most. This is the profile's stated rule (`tuning`: scale damage until the strongest raid set covers about 80% and mana is a live constraint), which exists so that gear separates in the BiS ranking. | L1 to L5, section 3 | **owner's call** (see below) |
| P1 | The healer tier list never said what HPS leaves out | `/tiers/healer` | fixed: a lead line under the title |

**O1, for the owner.** Three ways, none taken here:
1. Keep the profile for the BiS ranking (it is what makes gear separate) and keep the tier page's line (shipped).
2. Rank the healer tier list on a second, raid-shaped fight: the logged 90th percentile kill (about 110 s), with damage sized to the logs' per-healer share and spikiness, so overheal and idle time exist. It needs a second heal profile in the ranker and a published second figure; a ranking run measures it.
3. Retune the one profile toward the logs. Gear would stop separating (why the profile was scaled up from its first, Onyxia-sized figures), so the BiS lists would need another way to rank healer gear.

## The tier page

`/tiers/healer` gets one lead line under the title (`PageHeader lead`), and the same sentence replaces the app notes' "Mana and raid utility are not ranked here" (which was wrong: mana decides most of this list):

> HPS is effective healing, overheal left out. A shield counts for what it soaks, not for the death it stopped by being up before the big hit, and buffs on others, like Power Infusion, are not counted.

It names no figure, so it holds whatever the nightly ranking publishes. `web/src/lib/tiers/tier-copy.ts` (`tiersCopy.lead`), `web/src/components/tiers/TierPage.astro`; tests `tier-copy.test.ts` (red first, then green) and `tests/e2e/tiers.spec.ts`. Captures (390, 1024, 1440, 2000) and verdicts: `design/reviews/2026-11-13-healer-tiers/`. The first copy ("HPS counts healing done. It misses a shield up before the big hit...") drew SHIP WITH FIXES from both: "healing done" includes overheal in a player's meter, and it read as if shields do not count. The wow-player's rewrite shipped: wow-player SHIP on the second pass; ux-designer SHIP WITH FIXES, one optional cut of the stopped-death clause, which it accepts keeping when that point must stay (it is the point this lane was asked to make).

## Gates

On the box, rebased on eefe068f0: `npm test -- src/lib/tiers/` (122 passed), `npx playwright test tests/e2e/tiers.spec.ts` (68 passed, desktop and mobile), eslint and prettier on the touched files, `npm run check` (0 errors). Before the rebase, `app-files.test.ts` "carries each guide's current build codes" failed identically on origin/main 05d80dada (not this lane).

## Files

- `healer-mana/jobs-base.txt`, `jobs-duration.txt` (run with `discipline/jobs.sh`), `sim-raid60.txt`, `sim-durations.txt`.
- `healer-mana/logs/` (`wcl.py`, the round 13 fetcher with its cache moved out of the repo; `collect.py`; `analyze.py`; `externals.py`), `logs-summary.txt`, `logs-kills.json`. The key is read from `WCL_V1_API_KEY` in Secret Manager into the environment and never written.

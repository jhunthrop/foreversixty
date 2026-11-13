# Hunters at 50+ in real logs (2026-11-13)

**Result: no hunter at level 50 or above in any readable log. n = 0 for every requested measurement (boss or trash).**
The highest hunter seen anywhere is level 43 (one Aspect of the Cheetah cast in Stormwind City); the highest hunter in combat is level 38 to 39.
No raid exists to log yet (Forever raids open 9 Dec), so "raid bosses first" has no data on any site.

## Sources checked

| source | reports | hunters (rows) | highest hunter level | highest hunter in combat | 50+ |
|---|---|---|---|---|---|
| Our DB + bucket, all reports (backfilled) | 7 (z2opt2x2gs5u 20 Oct Zephras Isle; gxrgvz4yzsi2 8 Nov Wetlands; 3sdrcwdgo2yi 10 Nov Desolace, 674 fights; edcgw6rodqs2 12 Nov Scarlet Monastery, 5 boss kills; eygg26pg4gr2 12 Nov Hillsbrad; ctxicu4cuyek 12 Nov Stormwind City; 4q5el2xsgoy7 12 Nov, 1 fight) | 412 hunter rows across 1,178 fights | 43 (adv_level, Stormwind, no combat) | 39 (SM, 6 casts), 38 (Hillsbrad, 23 casts), 36 (Desolace, 68 casts); best sample: a level-30 Beast Mastery hunter with Summon Hawk, 18 fights, 30.7 active min | none |
| foreverlogs.gg | one report newer than 2853: 2854 (12 Nov, Barrens: Razormane, Bael'dun; 55 segments, 0 kills). Ids 2855 to 2900 return 404; the public list (26 reports, hasMore false) ends at 2854 | 5 hunters in 2854 | about 20 to 24 by ranks (Aimed Shot 19434 r1, Arcane Shot 14282 r2, Serpent Sting 13550 r2, Multi-Shot 2643 r1) | same | none |
| Warcraft Logs | vanilla/fresh/classic/sod subdomains return HTTP 403 (Cloudflare) to non-browser fetches, and no source shows a Forever zone or partition on WCL. The v2 API needs a client token we do not have | n/a | n/a | n/a | not readable |

Level comes from the advanced-logging `adv_level` field on the hunter's own events (our logs) and from ranks cast (foreverlogs, which carries no level). The prior 8 Nov note that `adv_level` might be scaled does not apply here: the ranks cast match it (e.g. Aimed Shot 20901 rank 3 at 31 and 38, Arcane Shot 14284 rank 4 at 33 to 39).

## What the newest data does show (below 50, for context only)

Not for comparison with a level-60 sim. n is small; one player dominates.

- **Summon Hawk is a short-lived summon next to the pet, not a pet replacement.** Hunters cast it (1293241 rank 1; 1293525 rank 2 by 31) while their tamed pet stays out. Each cast summons a Hawk unit (summon spells 1293248 and 1312637) and deals its own hit (Summon Hawk SPELL_DAMAGE, median 55, n=168, levels 26 to 34). Hawk units: 147, median 4 swings each, first-to-last swing 6.7 s, swing gap median 2.24 s (p10 1.71, n=503), hit median 24. Cast cadence for the one heavy user: 3.2 casts a minute, shortest gap 5.95 s. Level-30 BM hunter (18 fights, 30.7 min, 57,709 damage): hawk 25.2% (hawk melee 9,334 + summon hits 5,200), tamed pet 26.2% (white 12,705, Claw 2,400), hunter's own 48.6% (Auto Shot 15,585 = 27.0%).
- **Auto Shot gap medians 2.24 to 2.26 s** (level 30 to 31, n=52 and 22 gaps under 5 s, 2 hunters); Auto Shot crit 3.6 to 5.1%; Auto Shot misses 6 of 129 (4.7%) for the level-30 hunter against level 25 to 32 mobs.
- **Bestial Wrath (19574) seen once in our logs:** a level-31 hunter, 3 casts, pet share 62% over 4.9 min of mostly trash.
- **Attack power on the hunter's own snapshots** (adv_attack_power, likely melee AP): 660 to 1,123 at levels 28 to 38.

Everything else asked (Aimed Shot cadence at 60, +3 miss rate, Sniper Shot 1310687/1310785/1310786, Strider Kick 1317257, Lacerate, Lacerating Strikes, Volley, Rapid Fire uptime, trinkets, mana): n = 0 at 50+. None of the Survival-only or Forever-only shots appear in any log at any level.

## Summon Hawk in detail (own logs, levels 17 to 36, follow-up)

Data: the 7 own reports already downloaded. 288 Hawk units (summon 1293248 or 1312637) from 37 hunters, 136 of them with at least one swing; 168 dive hits. Scripts `py/hawk2.py` (per unit, writes `data/hawk_units.json`), `py/hawk3.py` (lifetime), `py/hawk4.py` (swing size), dive table in `data/hawk_dive.txt`. "Alive" = the cast target still has events 18 s or more after the summon. Times are the log's own clock (the client's local time read as UTC; the Brazilian logs may run 3 h behind true UTC).

Engine model under test: 18 s, swing every 2.26 s (about 8 swings), swing mean 0.576 x hunter level (17.3 at 30), no AP, dive = flat + 5% RAP.

**1. Lifetime**
- **Hard cap of about 18 s.** No hawk swings later than 18.30 s after its summon (n=136). The longest runs have 8 swings over 17.6 to 18.3 s (7 hawks with 8 swings, 1 with 9).
- **Swing interval:** median 2.24 s, p10 1.78, min 1.56 (n=418 gaps). The first swing lands with the summon: under 0.1 s for 82% of hawks.
- **Swings per hawk (all 136):** 1: 12, 2: 15, 3: 29, 4: 26, 5: 26, 6: 14, 7: 6, 8: 7, 9: 1. The mode is 3 to 5, median last swing 8.5 s. Most of these are trash pulls where the target dies early.
- **Target alive for 18 s, no recast by the same hunter:** only n=7. Swings were 3, 3, 3, 4, 4, 5 and 8, last swing at 4.1, 4.5, 6.0, 6.7, 7.5, 17.6 and 18.3 s. Two ran the full ~18 s and five stopped at 4 to 7.5 s with the target still alive. n is tiny, and the log doesn't say why (range, CC or leash).
- **Target died first (n=114):** 80 hawks stopped at the death. 34 moved to another mob and kept swinging, all on a different target, until 8.6 to 18.3 s (most 14 to 18 s). So the hawk either retargets or vanishes, never past ~18 s. Only 5 hawks show a UNIT_DESTROYED event.
- **A recast does not dismiss the previous hawk.** For 64 hawks the same hunter cast again within 18 s (median 8.1 s later), and 50 of them kept swinging after that recast. Two or three hawks can be out at once.

**2. Swing size (landed, non-crit, not glancing; 447 hits, 21 hunters)**
- base_amount is the hit before armor. **base = 1.10 x hunter level** (median; p10 1.04, p90 1.25). That is about 1.9 times the engine's 0.576 x level. After armor, amount/base median 0.735, so the hit is 0.81 x level against level 26 to 41 mobs.

| hunter level | n | base min / median / max | non-crit amount median | crit % (landed) | hunter AP median | target levels |
|---|---|---|---|---|---|---|
| 26 | 16 | 26 / 32.5 / 35 | 22 | 9.1% (n=22) | 640 | 35, 37 |
| 27 | 41 | 26 / 30 / 42 | 24 to 28.5 | 0 to 6% | 605 to 863 | 30 to 36 |
| 28 | 68 | 29 / 32.5 / 38 | 24 | 2.5 to 8.8% | 614 to 699 | 26 to 41 |
| 30 | 272 | 30 / 33 / 38 | 24 | 9.3% (n=291) | 790 | 27 to 31 |
| 31 | 33 | 31 / 35 / 39 | 25 | 9.8% (n=41) | 656 | 31 to 37 |
| 33 to 34 | 5 | 31 / 35 / 37 | 27 to 31.5 | 0% | 719 to 744 | 24 to 31 |

- **AP:** base correlates with the hawk's own AP snapshot (r 0.61; hawk AP 290 to 690) and with the hunter's (r 0.60), but both rise with level (r 0.47 against level). The data can't separate an AP share from a per-level scale.
- **Glancing is flagged in the logs:** 18 glances in all, mostly a level-17 hunter's hawk against level-31 mobs (6 of 9) and level 23 to 26 hunters' hawks against level 35 to 37. Misses are 1 to 4% (plus dodges and parries).
- **The 10 Nov hotfix ("Hawk melee +80%") does not show.** base/level is 1.107 before 10 Nov 21:00 (n=26, logged 9 to 10 Nov), 1.111 in the last three hours before 11 Nov (n=44) and 1.100 after (n=377). Either the hotfix landed before 9 Nov or it changed something other than this hit.

**3. Dive (Summon Hawk SPELL_DAMAGE, 168 hits)**

| rank | hunter level | n (non-crit) | base min / median / max | amount median | hunter AP median |
|---|---|---|---|---|---|
| 1293241 | 23 to 29 | 34 | 70 / 79 / 85 | 49 to 69 | 507 to 784 |
| 1293241 | 30 | 84 | 70 / 74 / 77 | 50 to 54 | 758 to 790 |
| 1293241 | 31 to 36 | 6 | 78 / 79 / 90 | 48 to 58 | 656 to 1079 |
| 1293525 (rank 2) | 26 to 33 | 35 | 90 / 95 / 100 | 53.5 to 83 | 580 to 734 |

- **The dive is nearly flat.** Rank 1 base against the hunter's AP (after 11 Nov, n=106, AP 507 to 1079): slope 0.006 per AP point, intercept 70.5. A 5% share of AP would add about 28 across that range; the log shows about 3. The AP field is the advanced-log attack power, probably melee AP. Rank 2 is about 20 points higher.
- **It takes armor:** amount/base is 0.6 to 0.9.
- **Crits:** 8 of 168, at 1.43 to 1.46 times base, like the magic-table crits seen before. One crit reads 1.84.
- **Before and after 11 Nov:** rank-1 base is 74 to 82 before (n=19) and 74 to 79 after.

## Files

- Scripts: `py/q.sh` (read-only psql via the proxy; URL never printed), `py/fetch_own.sh` (public reports from foreversixty.gg/logs-data), `py/fetch_r2.py` (private reports, read-only GETs from the logs bucket; keys from GSM at runtime), `py/fetch_fl.py` (foreverlogs, hunters + pets), `py/scan_own.py` (every hunter row: level, casts), `py/hunter_own.py` (per-hunter damage, casts a minute, crit, misses, Auto Shot gaps, pet share), `py/hawk.py`.
- Data (gitignored, about 90 MB): `data/own/<report>/` (report.json, fights/<n>/summary.json and events.parquet), `data/scan_own.txt` (all 412 hunter rows, sorted by level), `data/fl/2854.json`, `data/fl_2854*.json`, `data/fl_public.json`.

## To get 50+ data

Re-run `py/scan_own.py` after new uploads (`py/q.sh -Atc "select id,log_zone,created_at from reports order by created_at"` shows them; private ones via `py/fetch_r2.py`). Level-60 dungeons (UBRS, Stratholme, Scholomance, Dire Maul) are the earliest source until raids open 9 Dec. `py/hunter_own.py data/own 50` prints the per-hunter table directly once a 50+ hunter exists.

# Sim-accuracy round 21: platform lane

Site branch `wt/r21-platform` (from int/r21 93c66ebc8). Fork branch `lane/r21-platform` (from int/r21 d5c708834). The lane owns sim/core, racials, the web, the addon and the live-log harvest. Client 1.60.1.70790.

Method: the live-log harvest first (`$R/platform/logs.md`, LOGS-READY at about 02:57 UTC), then the seed items in the brief's order, then the shared code round 20 changed. Every sim went through `$R/sim.sh` (frozen BiS, guide raid builds, 10,000 iterations at seeds 1 and 50001). Every CONFIRMED-BUG has a test that was red before its fix.

**Confirmed and fixed: 2** (both `old`), plus one comment correction. One change moves numbers, for Gnome warriors only (no published pick), by under 0.05%.

## Live logs

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | **No live Forever log exists anywhere yet.** api.foreversixty.gg `/v1/reports/recent` lists only `z2opt2x2gs5u` (20 Oct). foreverlogs.gg `/api/reports/public` is an empty list (also with `phase=2`, which the site now labels "WoW Forever"). Every report id probed (1 to 4000, the beta range 2700 to 2814 included) answers "Report not found", so the beta archive is gone from the API. Listings (`/api/reports?`, `/api/live`, `/api/analytics/*`) need a login token. Rankings and statistics are empty. Warcraft Logs v1 (key from GSM, not stored or printed) lists no Forever zone or partition on www, classic, fresh, vanilla or sod; the classic and vanilla zone lists are byte-identical to the 03:00 snapshot. Every log-gated item stays OPEN. | `$R/platform/logs.md`, `$R/platform/wcl/zones-*.json`; last looked 03:02 UTC 5 Nov | OPEN (no data) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Eureka! on Execute: the cost cut turned into extra Execute damage.** Execute uses all your rage and turns each point past its cost into damage. The engine paid Eureka!'s 10% cheaper cost and converted the 1.5 rage it saved, so a Gnome's Execute under Eureka! hit for 10% more plus the saved rage. On the server the saved rage converts nothing. The reporter's raw combat log on forever-bugs #439 (Zeyf, Gnome warrior, level 31, 16 Oct, client 1.60.1) shows it: Eureka! applied at 18:51:29.572, Execute (20647) at 18:51:30.569 with 105 rage (Expansive Mind's cap; the log's power field reads 1050/1050) hit for 434 raw. Unbuffed, the same Execute is 125 + (105 - 15) x 3 = 395, and 395 x 1.1 = 434.5. With the saved rage converted it would be 439.45 (the reporter's "expected" column). So Eureka! gave the 10% damage and nothing for the cost. Now the engine converts only the rage a flat cut saves (Improved Execute, as in Classic). Eureka! is the only percent cost modifier that reaches Execute. | https://github.com/ClassicWoWCommunity/forever-bugs/issues/439 and its attached `WoWCombatLog-101626_184451.txt` (cached in `$R/platform/i439/`); fork `sim/warrior/execute.go` `executeConvertedRage`; test `TestExecuteConvertsNoRageFromAPercentCostCut` (red: undefined, then green) | CONFIRMED-BUG (one player's raw log, flagged single-source; the log is the server's own record) | old (rounds 19 and 20 carried #439 as OPEN; nobody read the attached log) | E 91385c192 |
| P1b | **"Since 17 October Eureka! does nothing for Execute."** The reporter's later comment (30 Oct) shows a spreadsheet: on 17 Oct, Executes under Eureka! hit for 395, the same as without it. No log is attached for that day. The issue is open, with no Blizzard reply. On #96 (Revenge) Blizzard's AidanZMoon wrote "There are a few things Eureka! doesn't work on for each class", so a server-side exclusion is plausible. The 70790 client row still names Execute (5308, 20658 to 20662, 20647) in both of the warrior row's masks. The engine keeps the client's 10% damage. | #439 comment of 2026-10-30; #96 (Fixed); `sim/core/eureka_auto_gen.go` warrior rows | OPEN (a live Gnome warrior log with Execute both with and without Eureka! at the same rage) | - | - |
| P2 | **`spellIsKnown` was true below a spell's trainer level** (warrior round 20 X1). Some spells are registered at every level and gated only when cast (Whirlwind below 36). For those, the APL value read "known" at 30, so a `spellIsKnown` gate could not agree with a guide's trainables level gate. It now also needs the unit's level to reach the spell's `RequiredLevel`. Probe: a temporary build logged every `spellIsKnown` reference whose spell sits above the character's level, for all 12 specs that use the value (bare bands 10 to 50 at their floor and +9, plus raid 60). It found none, so no shipped figure moves. | fork `sim/core/apl_values_spell.go`; test `TestValueSpellIsKnownHonoursRequiredLevel` (red at level 35, then green); probe `$R/platform/probe.sh`; `go test ./sim/core` passes | CONFIRMED-BUG (engine) | old | E 671c6bac8 |
| P3 | **Arcanite Champion's proc rate.** Still 1 PPM (`sim/common/item_effects.go:227`, "assumed"). There is no rate in the client. Item 12790 spell 16916 is trigger 2 (chance on hit) with no rate. cmangos classic-db (`raw/classicdb/item_template.json`, commit ec4f5961) gives 16916 proc_chance 101 (the placeholder) and no ppm. No row in hotfix-push.json. The web has no measured rate either: the Blizzard Classic thread says "not sure of the PPM", and Wowhead and warcraft.wiki.gg list none. 1 PPM is the convention for a chance-on-hit weapon with no data. It stays, flagged. | `data/builds/1.60.1.70790/raw/classicdb/item_template.json`; https://us.forums.blizzard.com/en/wow/t/weaponsmith-champion-or-reaper/166147 ; https://warcraft.wiki.gg/wiki/Plans:_Arcanite_Champion_(Classic) (read 2026-11-05) | OPEN (a level-60 log with about 30 minutes of Arcanite Champion swings, counting the Strength of the Champion procs) | - | - |
| P4 | **Skyborne launch offsets (round 20 P1a).** #414 has nothing new (closed and Fixed; the last comment is Blizzard's of 21 Oct, "different stat offsets" per variant, values unpublished). A web search (Warcraft Tavern, Icy Veins, seemeta, rankedboost, metabot) finds racial abilities only, no stat sheet. No published pick is Skyborne. The beta row stays in `ProvisionalConstants`. | https://github.com/ClassicWoWCommunity/forever-bugs/issues/414 ; search 2026-11-05 03:00 UTC | OPEN (one naked Skyborne of each variant, Character panel at any level) | - | - |
| P5 | **The stale fork comment** in `sim/shaman/talents.go` named Elemental Alacrity's node 104765, which is Elemental Warding. It is 104771 (talents.json 70790; `talents_auto_gen.go` already maps it). Comment only. | talents.json:569, :655 | Fixed (comment) | - | E b0d2e74ca |
| P6 | **Bear threat constants** (Maul and Swipe 1.75, Primal Bite and Lacerate 2.0, druid round 20 B3). There is still no client column and no Forever source. The forever-bugs searches ("threat bear", "maul threat", "lacerate") return nothing. The web gives Classic/TBC values only (a 20% damage-to-threat Lacerate in TBC, which Forever's Lacerate does not share). No log carries threat. The druid lane's constants stay. | gh search on forever-bugs 2026-11-05; web search (Classic/TBC threads) | OPEN (a threat-meter capture or a blue post) | - | - |
| P7 | **Fresh audit of the shared code round 20 changed:** Skyborne base stats (cd6a2ebcb), the Blizzard-only proc share (8fb31536d, reverted 9079201c9), and the 70790 follow (1494a06e8: build stamps, Eureka! drops the rune row Mutilate 1310707, Lava Burst table unused). The Gnome rows (Eureka! 8 s, no charges, 2 min; Expansive Mind +5% mana, rage cap 105 and energy cap) match both the client and the #439 log's 105 rage cap. The Troll, Orc, Dwarf, Night Elf, Undead, Human and Tauren rows of the published picks are as round 20's tests pin them. `go test ./sim/core` passes. `go test ./sim/warrior` passes. | engine read; tests | Aligned | - | - |

## Before/after

Only P1 moves a number. Gnome warriors are no published pick (Fury is Human/Troll, Arms Human/Orc), so `before.tsv` does not change. Raid 60, alliance, `-race gnome`, guide raid build, frozen BiS, 10,000 iterations, seed 1 / seed 50001, before -> after:

| spec | seed 1 | seed 50001 |
|---|---|---|
| warrior-fury (gnome) | 916.2 -> 915.9 (-0.04%) | 917.5 -> 917.1 (-0.04%) |
| warrior-arms (gnome) | 852.3 -> 852.1 (-0.02%) | 851.1 -> 850.9 (-0.02%) |

P2 moves nothing. The probe found no shipped `spellIsKnown` reference to a spell above the character's level.

## Files to regenerate

- None. No fork golden was regenerated, and none should move (no golden sims a Gnome warrior, and the probe shows no `spellIsKnown` change on any shipped APL).

## Measure at refresh

- None from this lane. A Gnome warrior pick in the re-rank would carry P1's -0.04%.

## For other lanes

- **Warrior (touches your package):** `sim/warrior/execute.go` gained `executeConvertedRage`, and Execute's damage now uses it (E 91385c192). It is three lines in ApplyEffects plus a helper. If your lane also edits Execute, keep the helper call. Gnome only.
- **Every lane:** `spellIsKnown` now also needs the character's level to reach the spell's `RequiredLevel` (E 671c6bac8). A gate on a spell registered at every level (Whirlwind) now reads false below its trainer level. That is the point of X1, and no shipped line changes.
- **Every lane:** no live log exists yet (`$R/platform/logs.md`). Log-gated items stay OPEN with the log named.

## For the guide lane

- None. The Arms and Fury race paragraphs (arms.md:111, fury.md:114) describe Eureka! as "10% more damage and 10% less rage on your non-periodic damaging abilities", which is still the client tooltip. The Execute nuance is too small to change a race ranking.

## Still open

- P1b (Eureka! and Execute since 17 Oct), P3 (Arcanite Champion PPM), P4 (Skyborne launch offsets), P6 (bear threat constants), L1 (no live logs; last looked 03:02 UTC 5 Nov).
- Round 20 web SHOULDs, unchanged: relic-row body colour and example icon; "Hit is capped" wording beside "Hit cap: reached"; the Fury card's "with 30 seconds left".

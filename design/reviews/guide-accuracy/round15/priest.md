# Guide accuracy round 15: Priest

Audited on the round 15 audit tree (main 4600b7d82: fourteen guide rounds, sim round 17, launch state), client 1.60.1.70755, engine pin b2c4a2d05 (`sim/enginever/version.go`).
Files: `web/src/content/guides/priest/index.md`, `discipline.md`, `holy.md`, `shadow.md`; the card notes and opener rows in `data/curated/apl/priest-*.json`; every frontmatter and body link; the live BiS files `data/builds/1.60.1.70755/bis/priest-{shadow,discipline,holy}.json` against `items/priest.json`. Read first: the round 1 to 14 priest reports and fix logs, the sim round 17 priest report (`design/reviews/2026-11-04-sim-audit-round17/priest.md`), and the ledger.

## What moved since round 14

`git diff d7fc51fe7 4600b7d82` on the priest side:

- **discipline.md:64** (sim round 17, D1): "Renew the tank a second before he engages, so it is already ticking when the first hits land." The Discipline APL gained a pre-pull Renew rank 7 (10927) on the tank (player 5) at -1 s, after Lightwell at -2.5 s.
- **shadow.md:107** (guide round 14 fix of R14-1): the best-in-slot paragraph was rewritten.
- `addon-builds.json`: no priest hunk (rogue Subtlety and warlock Affliction only). `talents/`, `spelltips/priest.json`, `races.json` and `bis/priest-*` are unchanged.

The decoded builds are therefore the same as round 13 and 14 (Shadow leveling `02510503103/0/52542001201300051`, Shadow raid `00510503103/0/50052201211312251`, Discipline `00500503133001143/3350500303/0`, Holy `00500503103/33505003030121431/0`).

## Findings

| id | status | kind | file:line | claim | evidence | fix |
|---|---|---|---|---|---|---|
| PRIEST-R15-1 | CONFIRMED | gear | shadow.md:107 | "Judge it by that damage number and its Intellect, never by the healing: a healing piece with good Intellect can beat a thin caster piece." | This is round 14's own fix, and the page's Shadow table breaks it. At 60 bare Alliance the ranged slot is **Torch of Light** (6 spell damage, 11 Spirit, no Intellect), with **Brilliant Wand** (9 spell damage, 7 Intellect) as the runner-up at -1.03 (`bis/priest-shadow.json`, band 60 bare alliance, ranged; stats from `items/priest.json` 279246 and 249385). Judged "by that damage number and its Intellect", Brilliant Wand wins on both counts, yet the table picks the Torch on its Spirit. At 30 the back slot is the same story: Glowing Thresher Cape (3 damage, 8 Spirit, no Intellect) over Repairman's Cape (2 damage, 3 Intellect, 5 Spirit). The round 14 fix log itself noted that "Spirit can also tip a slot", then shipped a rule that leaves Spirit out. | "Forever puts a third of a healing item's bonus on it as spell damage, so a healer's cloth piece still does some damage. Judge it by that damage number, its Intellect and its Spirit, never by the healing: while you level, a healing piece with good Intellect or Spirit can beat a thin caster piece." |

Counts: CONFIRMED 1, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## SHOULDs (wording, not findings)

| id | file:line | note |
|---|---|---|
| PRIEST-R15-S1 | discipline.md:64 | "so it is already ticking when the first hits land." Cast at -1 s, Renew's first tick comes about 2 s into the fight (the sim round 17 wow-player note says the same), so on the first hit the HoT is on but has not ticked. Not false enough to file. Optional: "so it is already on him when the first hits land." |
| Carried | R14-S1 (shadow.md:103 haste reason), R13-S1 (index.md:48 Skyborne per faction), R13-S2 (shadow.md:91 15 s vs 6 s), R12-S1/S2 (healer "would overheal" card notes), R11-S2, R8-S1/S2, R9-S1 | Unchanged and still optional. |

## Verified clean this round

- **The new Discipline pre-pull line (discipline.md:64) against the APL.** The prepull list is Lightwell at -2.5 s then Renew 10927 rank 7 on player 5 at -1 s; Lightwell's 1.5 s cast ends at -1 s, so the Renew lands on the next global. The sentence order (Lightwell, Fear Ward if needed, Renew a second before) matches. Rank 7 is the rank the tank Renew row casts and the rank discipline.md:70 names ("rank 7 or higher"); rank 7 is learned at 44, inside the 17-level grace at 60, so it keeps its whole bonus as discipline.md:74 implies. Renew is a trainer spell (`spelltips/priest.json` 10927: 250 mana, 370 over 15 s). The line does not contradict holy.md:66, which keeps Holy's Renew "as he engages": sim round 17 H1 measured Holy's pre-pull Renew as no gain, and the Holy list has none.
- **The rest of shadow.md:107.** "A healing piece with good Intellect can beat a thin caster piece" holds where it says so: Papal Fez (17 Int, 8 damage) over Augural Shroud (10 Int, 11 damage) and Inquisitor's Shawl (13 Int) over Bloodmage Mantle (9 Int) at 40 both factions; Mooncloth Boots (14 Int, 13 Spirit) at 60 bare Alliance. The first clause still holds for every healing item in the table.
- **Everything round 14 verified**, re-read against unchanged data: talent text, spell values, both leveling blocks against `addon-builds.json` (no priest change), low-rank arithmetic, card notes, the Shadow opener, consumables, race prose against round3/race-sims.md and `recommendedRaces`, racials against `races.json`, the Skyborne line, stat paragraphs against the band-60 tables (explain, never rank), and the Discipline and Holy profession lines against their BiS rows.
- **Sim round 17's other priest items** touch no prose: the Shadow Word: Pain cutoff measured as a tie (10 s, 7.5 s and 6 s inside error), so shadow.md:89's 6 seconds stands; Eureka! masks stay open on the engine side and index.md:38 follows the client tooltip.
- **Source links** (fetched 2026-11-04): Blizzard Deep Dive recap 200; Power Up Gaming 200; Talents Forever data.json 200; foreversixty.gg /sim/weights, /planner, /planner?class=priest 200; Icy Veins talent calculator 403 (unknown, not dead, R5-D2).
- No em or en dashes in the changed lines or anywhere on the four pages; no site or sim words outside the planner links.

## Decisions

- **R15-D1.** shadow.md:107 is filed as CONFIRMED: it is a round 14 edit that states a rule the page's own live table breaks (Torch of Light over Brilliant Wand), which the round 14 and 15 notes say to file. The fix adds Spirit and scopes the beaten-caster case to leveling, where Spirit has weight; at raid 60 Spirit weighs nothing and no healing piece outranks a caster piece on Spirit there.
- **R15-D2.** The Discipline pre-pull wording is a SHOULD, not a finding: the Renew is on the tank when the first hits land, and "ticking" is loose, not false.

## Sims requested

Carried, no prose change waits on them:

- **Discipline race pick with the top-rank shield and the new pre-pull Renew**: band 60, published gear, Gnome against Human (Alliance), Undead against Troll (Horde), 10,000 iterations, seeds 1 and 50001.
- **Shadow race lines on the raid build** `50052201211312251`: Human, Night Elf, Dwarf, Gnome, Undead, Troll, same method.
- **BiS re-rank on the shipped builds** (post-launch). The Shadow bands still carry the pre-round-15 ladder talents (sim round 17 S4).

## Sim findings

None new. Still open: S-R2-1 (Power Infusion on self), S-R9-1 (Eureka! periodic mask, needs a Gnome Shadow log), X1 (low-rank penalty above 30 and the top-rank exemption), W1, and Lightwell Renew's crit source (sim round 17 D3).

## Sources

- Repo data (70755): `bis/priest-shadow.json` (band 60 bare alliance ranged; band 30 back; band 40 head, neck, shoulder), `items/priest.json` (279246 Torch of Light, 249385 Brilliant Wand, Glowing Thresher Cape, Repairman's Cape, Papal Fez, Augural Shroud, Inquisitor's Shawl, Bloodmage Mantle), `spelltips/priest.json` (10927), `addon-builds.json`; `data/curated/apl/priest-discipline.json` prepull
- Sim round 17 priest report: `design/reviews/2026-11-04-sim-audit-round17/priest.md` (D1, H1, S1, S2, S4)
- Race measurements: `design/reviews/guide-accuracy/round3/race-sims.md`
- Blizzard, Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Talents Forever data export: https://talentsforever.com/data.json; Power Up Gaming: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/; Icy Veins talent calculator news (403, unknown)

## Fix log (round 15)

- PRIEST-R15-1 fixed, shadow.md:107: Spirit added to the rule. Dropped "while you level", since Torch of Light is a level 60 pick too.
- PRIEST-R15-S1 left (optional).

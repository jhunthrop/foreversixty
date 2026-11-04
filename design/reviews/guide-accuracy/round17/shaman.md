# Guide accuracy round 17: Shaman

Auditor: shaman (class index, Elemental, Enhancement, Restoration). Audit tree main 919f7e8a7 (sixteen guide rounds plus sim round 18), client 1.60.1.70755, engine pin 6ea675594 (`sim/enginever/version.go`).
Files: `web/src/content/guides/shaman/{index,elemental,enhancement,restoration}.md`, card notes in `data/curated/apl/shaman-{elemental,enhancement,restoration}.json`, rank rows in `data/curated/rank-choices.json`, live BiS `data/builds/1.60.1.70755/bis/shaman-*.json`, generated card text in `addon-data.json` and `addon/ForeverSixty/Data.lua`. Read first: rounds 1 to 16 `shaman.md` (with fix logs), `round4/shaman-levels.md`, `ledger.md` through round 16, the round 1, 2 and 17 briefs, and the sim round 18 shaman report (`design/reviews/2026-11-04-sim-audit-round18/shaman.md`).

## What moved since round 16 (2c3135a2e to 919f7e8a7)

- **Sim round 18 (new work):** `shaman-enhancement.json` priorityList[7] gate 50% to 25% and its note; `enhancement.md:77` (loop item 6) and `:87` (Mana paragraph), b2a0fd5b2 and c97b6e11b. `addon-data.json` and `Data.lua` regenerated: every Enhancement Earth Shock row now reads "more than a quarter of your mana".
- **Round 16 fix:** `enhancement.md:93` adds Fire Nova to the past-9% list.
- **70755 data job:** `spelltips/shaman.json` fills Ancestral Spirit's mana (120/260/420/600/800, was `$q1`); `items/shaman.json` adds class/subclass/inventory fields and expertise on a few items (no bonding or armor-type change); `hotfix-push.json` adds one hunter spell. Nothing a shaman page quotes moved.
- **Unchanged:** `addon-builds.json`, `talents/shaman.json`, `trainables/shaman.json`, `spellconst/shaman.json`, `races.json`, `bis/shaman-*`, the shaman rows of `rank-choices.json`, the Elemental and Restoration APLs.

## Builds and shipped leveling order

Unchanged since round 16 and re-checked against `addon-builds.json`; the level-by-level orders in round 16's report stand, and every "Levels 1 to 19", "Levels 20 to 60" and talents paragraph still matches them and the header card. The leveling tree (land/leveling-tree) shows the same orders; no prose contradicts it.

## What checked out (no finding)

- **Loop item 6 and card note 7** (`enhancement.md:77`, APL priorityList[7]): "more than a quarter of your mana" matches the APL's `currentManaPercent > 25%`, and the rendered card text in `addon-data.json` matches. "Keep that last quarter for Stormstrike and your totems" is true: Stormstrike costs 8% of base mana (`spelltips` 17364 `cost_pct: 8.0`), and Searing r6 170, Strength of Earth r4 225, Grace of Air r2 250 cost mana.
- **"Stormstrike, which hits harder than the shock you skip"** (`:87`): Stormstrike is 150% weapon damage (17364); Earth Shock r7 is 531 base (10414) with Concussion. A raid two-hander at 150% plus normalized attack power is well past that, and the same holds at every leveling band with a level-appropriate two-hander. True.
- **Judgement of Wisdom on both factions:** Undead paladins are new in Forever (Deep Dive recap; `paladin/index.md:39`), so a Horde raid has the Judgement too.
- **Ancestral Spirit** (`restoration.md:86`): "brings a dead player back after the fight" holds on the filled 70755 text ("Cannot be cast when in combat").
- **Mana Tide** 12 seconds (`restoration.md:57`, `:84`, card .1): `${13 sec-1}` = 12 on all four ranks.
- **Every other line** re-read against the unchanged data: no change from round 16's verified state (spell numbers, low-rank arithmetic, bolt and heal ladders, talent text, gear and professions against both factions' BiS rows, races, consumables, card notes, voice).
- **Source links** (fetched 2026-11-04): Blizzard Deep Dive recap 200, Talents Forever data.json 200, Wowhead Skyborne First Look 200, foreversixty.gg /sim/weights, /planner, /bis/shaman/{elemental,enhancement,restoration} 200; Icy Veins calculator 403 (unknown, not dead). Labels match.
- **Voice:** no em or en dashes, no site, sim, APL, preset or band talk, no nightly numbers in the changed lines.

## Findings

| id | status | kind | file:line | quote | evidence | fix |
|---|---|---|---|---|---|---|
| SHAMAN-R17-1 | CONFIRMED | rotation (mana) | enhancement.md:87 | "In a raid, Judgement of Wisdom and Mana Spring keep you topped up, so Earth Shock never needs to wait." | Sim round 18's own before/after (same binary, paired seeds 1 and 50001, 10,000 iterations): lowering the gate from 50% to 25% raised **raid 60** DPS 855.3 to 859.1 (Horde, +0.45 / +0.41%) and 855.3 to 856.6 (Alliance, +0.16 / +0.13%). On fixed seeds the gate can only change the result if it binds, so a raid-buffed Enhancement shaman does fall below half mana during the fight: "topped up" is false. What round 18 measured is that the shaman never runs **out** (out of mana 0 s of 180). Whether the 25% gate ever binds in a raid ("never needs to wait") was not measured (see Sims requested 1). The sentence also reads as a raid exception to loop item 6, which has none. | Replace the first sentence with: "In a raid, Judgement of Wisdom and Mana Spring keep enough mana coming that you never run dry, so keep shocking down to that last quarter." If Sim 1 shows the 25% gate never fires in a raid, "Earth Shock never needs to wait" may stay after "never run dry"; "topped up" goes either way. |

Counts: 1 CONFIRMED, 0 LIKELY.

Wording SHOULDs, not findings:

| id | status | kind | file:line | text | note |
|---|---|---|---|---|---|
| SHAMAN-R14-S2 | SHOULD (carried) | wording | enhancement.md:52 | "9% melee crit for 10 seconds after a spell crits" | Elemental Devastation reads "critical strikes with direct damage spells"; "after a direct damage spell crits". |
| SHAMAN-R15-S2 | SHOULD (carried) | wording | restoration.md:76 and card .17 | two phrasings of the same zero-bonus fact | Both true. |
| SHAMAN-R12-S1, R9-S1 to S4 | SHOULD (carried) | wording | as filed | | Unchanged. |

## Decisions

- **D1. The 25% rule stands as one rule.** Round 18 measured band 40 preferring 50% by 0.3 to 0.85% (Maelstrom Weapon 3/5 still charges for the bolt) and chose one gate for the guide; the page gives one rule and states nothing about band 40, so nothing on the page is false. Not a finding.
- **D2. "Stormstrike ... hits harder than the shock you skip"** is true at every band (see above); round 18's wow-player HOLD on the earlier totem wording was the right call and the fix is correct.
- **D3. Carried, not re-filed:** Grace of Air rank 3 and Strength of Earth rank 5 book ranks (card .2), Stormstrike `RequiredLevel: 40` against `enhancement.md:121`, Restoration's statPriority leaving out spell haste (no shaman item carries haste).

## Sims requested

1. **shaman-enhancement, raid 60, guide raid build, frozen BiS, Horde and Alliance, seeds 1 and 50001, 10,000 iterations:** the 25% gate (shipped) against no gate (Earth Shock on cooldown whenever affordable), plus the time spent under 25% mana on the shipped rotation. If the two runs are identical, Earth Shock never waits in a raid and the fix may keep "Earth Shock never needs to wait"; if not, use the fix as written.
2. Carried, optional: shaman-elemental Lightning Bolt rank at 20 to 29 (round 16 Sims requested 1, unchanged).
3. Carried, at the refresh: Enhancement band 60 hit weight; Enhancement band 30 with Stormstrike; Enhancement relic slot (sim round 18 S3).

## Sim findings

- Carried: Chain Heal party-only jump rule absent from `chain_heal.go`; Stormstrike `RequiredLevel: 40`; haste and the global cooldown (SIM-R4-1); Grace of Air rank 3 and Strength of Earth rank 5 book ranks (N5); Lightning Overload ratio on the first level-60 log; Spellblasting potion on the Elemental pull; Totemic Call positioning (SIM-R5-1); self-dropped Windfury Totem plus Rockbiter not a shipped preset; Healing Wave rank 10 treated as a book rank by the Nature's Swiftness save.
- Sim round 18 for other lanes, not guide findings: relic slot publishes empty (S3, ranker), "Hit cap: reached" ignores the spell side (X1, platform), ladder golden header says Enhancement dual-wields (N5, text only).
- Data lane (already filed): Restoration band 60 Horde rows run as tauren; BiS race captions; the 7% spell-hit baseline for an Orc; `spellranks.json` still carries Lava Burst; Enhancement hit weight.

## Sources

- Client data 1.60.1.70755: `spelltips/shaman.json` (Stormstrike 17364, Earth Shock 10414, Searing Totem 10438, Strength of Earth 10442, Grace of Air 10627, Mana Spring 10497, Mana Tide 16190/17354/17359/1325256, Ancestral Spirit 2008/20609/20610/20776/20777, Totemic Call 437009), `races.json`, `addon-builds.json`, `items/shaman.json`, `addon-data.json`.
- `data/curated/apl/shaman-enhancement.json` priorityList[7]; `design/reviews/2026-11-04-sim-audit-round18/shaman.md` (N1 and the before/after table); commits b2a0fd5b2, c97b6e11b.
- Blizzard, "World of Warcraft Forever Deep Dive panel recap" (Undead Paladin among the six new pairs) https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, "Skyborne First Look" https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Talents Forever, data export (beta client data files) https://talentsforever.com/data.json
- Icy Veins, "Start Planning Your WoW Forever Character With Our New Talent Calculator" (403 to automated fetch; unknown, not dead) https://www.icy-veins.com/wow-forever/news/start-planning-your-wow-forever-character-with-our-new-talent-calculator/

## Fix log (round 17)

- SHAMAN-R17-1: fixed, `enhancement.md:87`. The Mana paragraph now opens with the one rule ("Stop shocking at a quarter of your mana ...") and closes with "In a raid, Judgement of Wisdom and Mana Spring bring in enough mana that you never run dry, so you can shock freely down to that quarter." Neither "topped up" nor "never needs to wait" is claimed, so Sims requested 1 is not needed for the page.

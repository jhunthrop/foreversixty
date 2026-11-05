# Guide accuracy round 17: Rogue

Audited on the round 17 audit tree (main 919f7e8a7: sixteen guide rounds plus sim round 18), client 1.60.1.70755, engine pin 6ea675594 (`sim/enginever/version.go`).

**Scope.**
- Pages: `web/src/content/guides/rogue/{index,assassination,combat,subtlety}.md`, every line.
- Card notes: every rogue line in `data/builds/1.60.1.70755/addon-data.json` (`rotations.rogue-*`: 23 Assassination bands, 22 Combat, 22 Subtlety), and the notes in `data/curated/apl/rogue-{assassination,combat,subtlety}.json`.
- Shipped orders: `addon-builds.json` (index 0 = level 10), both orders per spec.
- BiS: `bis/rogue-*.json` (unchanged since round 16).

**Read first:** the rogue reports for rounds 1 to 16 with their fix logs, `round3/race-sims.md` (still the only race file with rogue rows), `ledger.md` through round 16.

## What changed since round 16's tree (2c3135a2e)

Rogue-relevant diff:
- `subtlety.md:70` (sim round 18): the opener now reads "Then **Ghostly Strike** for a fifth point, **Rupture** right away, and **Death Mark** ...".
- `subtlety.md:80` (sim round 18): "Between marks, every point Rupture and Slice and Dice do not need waits for the next mark, never Eviscerate, even with Death Mark on cooldown."
- `rogue-subtlety.json` notes: Ghostly Strike (R16-1 fix), Garrote (R16-2 fix plus "ahead of Rupture at exactly four"), Rupture (rewritten), Death Mark (rewritten), Hemorrhage (rewritten). The APL actions did not change.
- `addon-data.json` regenerated: only `rogue-subtlety` changed among the rogue rotations.
- Data job: rogue `talents`, `spelltips`, `trainables`, `races.json` and `bis/rogue-*` are unchanged. `items/rogue.json` was regenerated; no page names an item whose drop it moved.
- `data/curated/coach/rogue-subtlety.md` changed (out of scope, see the end).

**Re-check of round 16's fixes.**
- R16-1: the Ghostly Strike note now says "40 against Hemorrhage's 35, or 30 with both points of Efficient Strikes". 14278 costs 40, 16511 costs 35, Efficient Strikes 2/2 is 5. True on every band where both spells exist. Holds.
- R16-2: the Garrote note now says "a bleed that Opportunity feeds, and Hemorrhage and Death Mark too once you have them". Opportunity from 13, Hemorrhage 30, Death Mark 40. Holds.

**Re-check of the sim round 18 prose.**
- `subtlety.md:70` opener against the APL (`prepullActions`: Stealth at -3 s, Premeditation at -2 s, Garrote at 0): Premeditation 2, Garrote 1 plus Initiative 3/3 (100%) makes 4; Shank waits (its line needs three or fewer); Ghostly Strike fires at four or fewer and makes 5; Garrote cannot fire at five; Rupture fires at four or more with no Rupture ticking. Energy: Garrote 50 minus Dirty Deeds 20 minus Poised 10 is 20, Ghostly Strike 40, Rupture 25, all inside the opening bar. Matches the APL and does not contradict :78 ("never at five points"; Garrote follows a Ghostly Strike only "at four points or fewer").
- `subtlety.md:80`: matches the Death Mark line (four or more, Rupture ticking), the Slice and Dice line (three or more while Rupture ticks) and the Eviscerate line (only before Death Mark is known). Holds.

## Builds decoded

Decoded positionally against `talents/rogue.json` (tier then column). Every build is unchanged from round 16 and is legal on 70755.

- **Subtlety `build` 12/8/31:** Malice 5, Ruthlessness 3, Improved Slice and Dice 2, Blood Splatter 1, Lethality 1 / Coup de Grace 3, Efficient Strikes 2, Precision 3 / Master of Deception 3, Opportunity 5, Dirty Deeds 2, Camouflage 3, Initiative 3, Premeditation 1, Serrated Blades 3, Elusiveness 2, Ghostly Strike 1, Preparation 1, Hemorrhage 1, Boot Knife 5, Death Mark 1.
- **Combat `build` 14/34/3 and `raidBuild` 12/34/5, Assassination `build` 38/11/2 and `raidBuild` 36/10/5:** as listed in `round16/rogue.md`.

**Shipped Subtlety order** (both orders identical): Master of Deception 10-12, Opportunity 13-17, Dirty Deeds 18-19, Premeditation 20, **Serrated Blades 21-23**, Camouflage 24, Elusiveness 25, Ghostly Strike 26, Camouflage 27-28, Initiative 29 and 31-32, **Hemorrhage 30**, Elusiveness 33, Preparation 34, Boot Knife 35-39, Death Mark 40, then Assassination 41-52, Coup de Grace 53-55, Efficient Strikes 56-57, Precision 58-60. This matters for R17-1.

## Findings

| id | status | kind | file:line | quote | evidence | fix |
|---|---|---|---|---|---|---|
| ROGUE-R17-1 | CONFIRMED | card note (sim round 18 edit) | `data/curated/apl/rogue-subtlety.json` priorityList[4] notes, rendered on the Subtlety Rupture line of every band 20 to 60 (`addon-data.json` rogue-subtlety) | "Rupture at four or more combo points whenever it is not ticking, so it goes up right after the opener: Hemorrhage and Serrated Blades both feed it, which makes it the main finisher." | Rupture trains at 20 (`trainables/rogue.json`, 1943), so the line renders from band 20. Hemorrhage is index 20 of both Subtlety orders (level 30) and Serrated Blades is indices 11 to 13 (21 to 23). On the band 20 card neither talent exists; on the bands 22, 24 and 28 cards Hemorrhage does not, so its debuff cannot feed anything. This is the exact pattern fixed on the Garrote note in R16-2 ("Hemorrhage ... too once you have them") and the R15 band rule (a note that names a talent must hold on every band it renders on). Round 16 missed this line; sim round 18 rewrote it and kept the clause, so it is new work. Both talents do feed Rupture at 60 (Hemorrhage 16511 "+15% Rupture and Garrote damage", Serrated Blades 3/3 "+30% Rupture"). | Note: "Rupture at four or more combo points whenever it is not ticking, so it goes up right after the opener: it is the main finisher, and Serrated Blades from 21 and Hemorrhage's debuff from 30 make it hit harder still." Regenerate `addon-data.json` and Data.lua. |

No page-prose finding. Every line of the four pages was re-derived against 70755 (list below).

## Shoulds (wording, not findings)

| id | file:line | note | suggested wording |
|---|---|---|---|
| ROGUE-R17-S1 | `rogue-subtlety.json` Hemorrhage note (sim round 18 rewrite) | "Hemorrhage, your builder ..., and it keeps Hemorrhage's debuff up" names the spell twice in one clause and reads like generated text. | "Hemorrhage is your builder for every global nothing above takes, and it keeps its own debuff up: one combo point a cast, and the debuff makes your Rupture and Garrote hit 15% harder. ..." (rest unchanged) |
| ROGUE-R17-S2 | subtlety.md:80, :117; assassination.md:115; combat.md:116; `rogue-combat.json` Sinister Strike note | Carried R16-S1 to S4 (and R15-S1 to S4). | As R16. |
| ROGUE-R17-S3 | `rogue-subtlety.json` Garrote note, bands 14 to 18 | "ahead of Rupture at exactly four" renders on cards before Rupture trains at 20. It is a relative-order clause, not a false claim, and the line also renders before Ghostly Strike (component lane, R15/R16), so not filed. | Fold into the component lane's whole-row fix. |

## Checked and right (re-derived this round)

**Spell texts** (`spelltips/rogue.json`, unchanged): Ghostly Strike 40 energy, 20 s, 125%/180% with a main-hand dagger, 15% dodge 7 s, next Ambush or Garrote within 7 s needs no Stealth. Hemorrhage 35, 145% with a dagger, +15% Rupture and Garrote for 15 s. Garrote 50 (30 with Dirty Deeds 2/2). Cheap Shot 60, 4 s, 2 points, trains 26. Kidney Shot 25 on 20 s, 6 s at five at rank 2 (50). Shank 5 energy, 15 back at 24 and 25 at 44, 10 s, cannot miss. Poised -10 at full energy out of combat, trains 18. Premeditation 2 points, 20 s, 2 min, no Stealth clause. Preparation 10 min. Blind and Vanish 5 min (3.5 with Elusiveness 2/2). Redirect at 1, 15 s. Adrenaline Rush 15 s on 5 min, Blade Flurry 25 energy, 20% speed, 15 s, 2 min. Cold Blood 3 min. Mutilate rank 3 90% plus 79.2 a hand, 20% more on a poisoned target. Instant Poison VI 20% chance (so Venom 20 plus Improved Poisons 10 makes half). Death Mark 25 energy, 10 s, 6 s, Garrote and Rupture at twice the rate, rank 4 at 60.

**Talent texts:** Initiative 100% at 3/3; Boot Knife 100% for two points at 5/5; Opportunity 4% a rank on Backstab, Garrote, Ambush, Mutilate; Serrated Blades 9% armor and 30% Rupture; Blood Splatter 6 s and 9 s; Efficient Strikes 3/5 on Sinister Strike and Hemorrhage; Composure 5/10; Flash of Steel 15 energy and 2 s; Hack and Slash 5% dagger crit, 15% mace armor; Lethality +20% crit bonus on Sinister Strike, Gouge, Backstab, Mutilate, Ghostly Strike, Hemorrhage (220%, not Ambush); Puncturing Wounds 30% Backstab crit, 15% Mutilate, 45% extra point; Relentless Strikes 20% a point for 25 (four in five at four points); Venom 30% and 20; Seal Fate 100%; Coup de Grace 20% Eviscerate and Death Mark; Master of Deception three levels; Remorseless Attacks 40%; Improved Gouge 1.5 s.

**Levels** (trainables): Backstab 4, Gouge 6, Evasion 8, Slice and Dice and Sprint 10, Kick 12, Garrote 14, Ambush and Poised 18, Instant Poison 20, Rupture 20, Shank 24, Cheap Shot 26, Deadly Poison 30, Kidney Shot 30. Every level and talent range the four Leveling sections name matches `addon-builds.json` (Assassination: Improved Gouge 10-12, Remorseless 13-14, Malice 15-19, Venom 30, Mutilate 40, Relentless Strikes 39 and 41-44; Combat: Coup de Grace 10-12, Efficient Strikes 13, Lightning Reflexes 14-18, Puncturing Wounds 19-21, Aggression 28-30, Improved Gouge 44-46, Remorseless 47-48; Subtlety as above).

**Build paragraphs:** all three swap and toll paragraphs recount. Assassination :52 trades 7 for 7. Combat :52 trades 5 for 5 in Assassination, 3 for 3 in Combat and 3 for 5 in Subtlety, and the :63 toll holds. Subtlety :51 names every Assassination and Combat point.

**Cards:** every other rogue note holds on every band it renders on with its talents present (Assassination Rupture "only with Blood Splatter (the raid build)", Shank "with Venom"; Combat Backstab "once you have Puncturing Wounds", Shank "every 8 with five points"; Subtlety Backstab "from 13", Shank "with Boot Knife", Death Mark from 40). The cut-off Death Mark and Hemorrhage notes are the component lane's.

**Races:** every race line matches `round3/race-sims.md:18-23` and `races.json` (Touch of the Grave 5% / 15 s, Blood Fury 10% AP 15 s / 2 min, Berserking 10% / 15 s / 3 min, Elune's Light 10% crit 15 s / 3 min, Ruthless 1% crit, Artisan's Precision 1% hit, Expansive Mind 5% energy, Eureka! non-periodic damaging on the rogue row, Will of the Forsaken 2 min). Skyborne only as the paid option.

**Gear and professions:** BiS unchanged since round 16; Assassination and Combat dagger claims, Subtlety's Teebu's Blazing Longsword and Persuader rows, the Hearts and Heartseeker all hold as round 16 recorded.

**Sources:** all six URLs returned 200 on 2026-11-04 (curl with a browser user agent): the two Blizzard articles, the Wowhead Skyborne article, foreversixty.gg /sim/weights, /planner and /planner?class=rogue. Each says what its label says.

**Voice:** no em or en dashes, no site, sim, APL or band talk in the pages, no nightly numbers.

## Decisions

- **R17-1 follows R16-2.** A sim-lane rewrite that keeps a level-60-only talent clause on a line that renders from 20 is a finding, filed for the Hemorrhage and Serrated Blades clause only. The row rendering itself stays with the component lane.
- **The new opener is right.** Premeditation, Garrote (4), Ghostly Strike (5), Rupture is what the APL does and what a top Subtlety rogue does: Ghostly Strike is on cooldown and off a 20 s timer, so spending it on the fifth point before the first Rupture loses nothing, and the Garrote it would otherwise let in is already ticking.
- No prose, build, order, `statPriority` or race change.

## Sims requested

None new. Carried, with sim round 18 / the data lane, not re-filed: Subtlety band 20 BiS mace, Subtlety BiS and weights on the Dirty Deeds string, race re-runs for all three specs on the current builds, Combat band 60 raid re-rank on the shipped `raidBuild`.

## Sim findings

None new. Open from earlier rounds and untouched by pin 6ea675594 as far as the pages are concerned: Death Mark storage and the Mutilate/Shank off-hand penalty (K1), Combat Eviscerate at four (K2/C1), Occult Poison learnability (X1), Hemorrhage's dagger hand (R16 S4). No page claim on these contradicts the client.

## Out of scope, handed on

- `data/curated/coach/rogue-subtlety.md:25` (coach persona, regenerated this round) still says Ghostly Strike "hits harder than Hemorrhage for ten more energy": the R16-1 claim the card note dropped, false below 58 (5 more at 30 to 55, 8 at 56). For the coach lane: regenerate from the current note.

## Sources

- "World of Warcraft: Forever Class Deep Dives - Rogue and Warlock", Blizzard: https://news.blizzard.com/en-us/article/24310968/world-of-warcraft-forever-class-deep-dives-rogue-and-warlock
- "World of Warcraft: Forever Deep Dive Panel Recap", Blizzard: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- "Skyborne First Look: New Neutral Race in World of Warcraft: Forever", Wowhead: https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- foreversixty.gg: /planner, /sim/weights, /planner?class=rogue
- Client data, `data/builds/1.60.1.70755/`: `talents/rogue.json`, `spelltips/rogue.json` (14278, 16511, 11290, 703, 1833, 408, 8643, 1329283, 1329549, 1328825, 14183, 14185, 1316709, 1325758, 438040, 13750, 13877, 14177, 1241584, 11340), `trainables/rogue.json`, `races.json`, `addon-builds.json`, `addon-data.json`, `bis/rogue-*.json`
- APLs: `data/curated/apl/rogue-{assassination,combat,subtlety}.json`
- Previous rounds: `round16/rogue.md`, `round15/rogue.md`, `round3/race-sims.md`, `ledger.md`

## Counts

CONFIRMED 1, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs: 3 (1 new, 1 carried group, 1 component note). New sim findings: 0.

Verdict: NOT zero for rogue. The one finding is a Subtlety card note (Rupture) that credits Hemorrhage and Serrated Blades on cards from level 20. The page prose of all four rogue pages, sim round 18's two Subtlety edits and round 16's two fixes are clean.

## Fix log (round 17)

- ROGUE-R17-1: fixed, `data/curated/apl/rogue-subtlety.json` Rupture notes: "it is the main finisher, and Serrated Blades from 21 and Hemorrhage's debuff from 30 make it hit harder still" (Serrated Blades indices 11 to 13 and Hemorrhage index 20 of both Subtlety orders in `addon-builds.json`). `sim/request/apl/rogue-subtlety.apl.json`, `addon-data.json` and Data.lua regenerated (notes only; no engine pin).
- ROGUE-R17-S1 (Hemorrhage note wording): not taken this round; the note is true.

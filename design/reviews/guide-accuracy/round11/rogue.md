# Guide accuracy round 11: Rogue

- Class: rogue. I audited every line of `web/src/content/guides/rogue/index.md`, `assassination.md`, `combat.md` and `subtlety.md` in the round 11 audit tree (main d54de48aa: guide rounds 1 to 10, sim round 15, leveling-order). I also read every published card note in `data/builds/1.60.1.70755/addon-data.json` (`rotations.rogue-*`, all seven bands) and the curated APLs `data/curated/apl/rogue-{assassination,combat,subtlety}.json`.
- Engine pin: 9f7f1f6e4 (`sim/enginever/version.go`). I read the engine at `scratchpad/gr10/engine`: `sim/rogue/venom.go`, `poisons.go`, `premeditation.go`, `talents.go` (Blade Flurry, Adrenaline Rush), `client_values.go` (Blood Splatter).
- Read first: the rogue reports for rounds 1 to 10 with their fix logs, `round4/rogue-levels.md`, `round3/race-sims.md` (the only race file with rogue rows) and `ledger.md`.

## What changed for rogue since round 10

Between 84460cec8 and d54de48aa, only the round 10 fix commit (eabba8f84) touched rogue:
- prose: `assassination.md`, `combat.md`, `index.md`, `subtlety.md`
- APL notes: `rogue-assassination.json`, `rogue-subtlety.json`
- `addon-data.json`

Rogue's rows in `addon-builds.json` are byte-identical. The leveling-order merge moved no rogue point, and `bis/rogue-*.json` still carry the nightly re-rank de044125c. Every round 10 fix is re-checked below.

## Decoded builds (re-verified)

Builds were decoded positionally against `talents/rogue.json` tab order (`fs1.ts`).

| Spec | build | raidBuild |
|---|---|---|
| Assassination | 38/11/2 night-elf | 36/10/5 (Vigor 1, Blood Splatter, Improved Slice and Dice 3, Precision 2, Opportunity 5) |
| Combat | 14/34/3 gnome (Efficient Strikes 1, Hack and Slash 5, Flash of Steel 4, Lightning Reflexes 5) | 12/34/5 (Composure 2, Lightning Reflexes 4, Hack and Slash 4, Flash of Steel 5, Lethality 2, Opportunity 5) |
| Subtlety | 12/8/31 night-elf, single build | none |

**Swaps recount exactly:**
- Assassination: 7 for 7.
- Combat: 12 for 12.
  - Raid side: Lethality 2, Improved Slice and Dice +2, Opportunity 5, Flash of Steel +1, Composure 2.
  - Level 60 side: Improved Gouge 3, Remorseless Attacks 2, Ruthlessness +1, Master of Deception 3, Efficient Strikes 1, Hack and Slash +1, Lightning Reflexes +1.

**Tier gates hold:**
- Combat raid: Composure sits in tier 2 behind 10 points. Tier 3 is reached with Coup de Grace 3, Lightning Reflexes 4, Puncturing Wounds 3, Precision 3 and Composure 2, which is 15.
- Subtlety: tiers 0 to 4 hold 25 only with the Improved Ambush point.

**Leveling orders** (`addon-builds.json`, index 0 = level 10) match every level the prose names:

| Spec | Order |
|---|---|
| Assassination | Improved Gouge 10-12, Remorseless Attacks 13-14, Malice 15-19, Venom 30, Relentless Strikes 39 and 41-44, Mutilate 40 |
| Combat | Coup de Grace 10-12, Efficient Strikes 13, Lightning Reflexes 14-18, Puncturing Wounds 19-21, Aggression 28-30, Improved Gouge 44-46, Remorseless Attacks 47-48 |
| Subtlety | Master of Deception 10-12, Opportunity 13-17, Dirty Deeds 18, Camouflage 19, Premeditation 20, Ghostly Strike 26, Hemorrhage 30, Boot Knife 35-39, Death Mark 40 |

## Findings

| id | status | kind | file:line | quote | evidence | fix |
|---|---|---|---|---|---|---|
| ROGUE-R11-1 | CONFIRMED | talent | assassination.md:56 and :78 | ":56 every Shank makes your poisons hit 30% harder and apply 20% more often for 6 seconds" / ":78 with Venom your poisons hit 30% harder and land 20% more often for the next 6 seconds" | See the evidence and arithmetic below the table. | :56: "**Venom** (1 point): every Shank adds 30% to your poison damage, on top of Vile Poisons, and 20 points to each poison's chance to land for 6 seconds, so Instant Poison procs on half your off-hand hits. It is a passive on Shank, with no finisher or global of its own, and Mutilate needs it." :78: "...it cannot miss, and with Venom, for the next 6 seconds your poisons do 30% more damage and Instant Poison lands on half your off-hand hits." |
| ROGUE-R11-2 | CONFIRMED | rotation (round 10 edit) | subtlety.md:129; `rogue-subtlety.json:316` (Eviscerate card note, bands 10 to 60) | ":129 and Eviscerate at five until Death Mark comes at 40; from then on, every five points go into Death Mark." / note: "from 40, every five points go into Death Mark." | See the evidence below the table. | :129: "...keep Slice and Dice up on anything that lives, and Eviscerate at five until Death Mark comes at 40. From then on, keep Rupture ticking on anything that will live, and put five points into Death Mark while it ticks; Eviscerate a mob that will die before the mark goes off." Note: "Eviscerate at three or more combo points while Rupture is ticking, only before you have Death Mark: from 40, five points go into Death Mark while Rupture ticks." |

### Evidence for ROGUE-R11-1

**Client rows.** Venom's buff (`spellconst/rogue.json` 1310703) has three effects:
- e0 and e1: aura 108, amount 30. These are percent modifiers on the Instant and Deadly Poison masks.
- e2: aura 107 misc 18, amount 20. This is a **flat** add to the proc chance.

Improved Poisons (14113) is the same flat kind: aura 107 misc 18, amount 10. The engine models both that way:
- `venom.go`: "The aura 108 modifiers share Vile Poisons' (16513) kind... so the two add, never multiply."
- `poisons.go`: `0.2 + improvedPoisons + additivePoisonBonusChance`.

**Base proc chances.** The Wowhead Forever tooltips give the base rates:
- Instant Poison VI (8928): "Each strike has a 20% chance". With Improved Poisons 5/5 that is 30%; with Venom it is 50%, i.e. two thirds more often, not 20% more.
- Deadly Poison V (20844): 30%. With Improved Poisons 5/5 that is 40%; with Venom it is 60%, half again.

**Damage.** Venom's 30% adds to Vile Poisons' 20%, so poisons go from 120% to 150% of base. That is 25% harder than without Venom, not 30%. This is the same "bonus on a bonus" reading as the Lethality lesson.

**The card note** ("land more often") gives no number and stays true.

### Evidence for ROGUE-R11-2

**The rule on the same page.**
- :76 says "Rupture at four or more combo points whenever it is not ticking. **Death Mark** at five points while Rupture ticks."
- :78 says "never press it into a lapsed bleed".

**The APL.** `rogue-subtlety.json` puts Rupture (priorityList[4], at four or more points with Rupture down) above Death Mark (priorityList[5], which needs `dotIsActive(Rupture)`). So a five-point set with Rupture down goes into Rupture, not Death Mark. The Slice and Dice line at three or more points also takes five-point sets whenever it is down.

**Real play.** Death Mark (1325758) deals what it stored "After 6 sec". A leveling mob that dies inside that window takes nothing.

**Origin.** This line was written in round 10. It shortened the round 10 S1 wording, which had kept the Eviscerate-the-rest clause.

### Re-checked and holding (fresh pass)

**Round 10 fixes**
- combat.md:52: the swap recounts 12 for 12 against both strings.
- combat.md:63: Efficient Strikes 1/2 takes 3 off Sinister Strike (and Hemorrhage).
  - With the dagger APL, Sinister Strike fires only when Backstab cannot.
  - Composure rank 2 is +10 on Poised's 10.
  - Garrote 11290 costs 50, so the raid opener costs 30.
  - The tier 3 toll arithmetic holds.
- combat.md:73: Poised takes 20 with Composure in the raid build.
- assassination.md:64, :78 and :80 match `rogue-assassination.json` priorityList[3]: Rupture at four or more, Slice and Dice up, Rupture down, `auraIsKnown(1329256)`.
- **Blood Splatter is 6 s and 9 s.** The talent spell's own base points in `spellconst` are 4000/6000. The engine comment (`client_values.go:104-111`) shows that the trait curves give 6000/9000, which is what the 70755 rank text and Rupture's `$<bloodS>` print. So the page's 6 s and 9 s are right.
- subtlety.md:76 and :78 match priorityList[1] to [8]:
  - Shank at three or fewer and Ghostly Strike at four or fewer.
  - Garrote re-applied after Ghostly Strike at four or fewer.
  - Rupture at four or more, Death Mark at five with Rupture ticking.
  - Eviscerate gated on `not spellIsKnown(Death Mark)`.
  - Slice and Dice at three or more when down or under 3 s.
  - Hemorrhage held when Ghostly Strike is within 1 s unless energy is at least 75.
- The Shank and Ghostly Strike card notes match their gates.
- index.md:37 and combat.md:112: "Expansive Mind makes your energy bar 5% bigger" matches 1259803.

**Talents, against the `talents/rogue.json` rank texts.** I re-read every one the prose names:
- Assassination: Improved Gouge 1.5 s; Remorseless Attacks 40%; Malice 5%; Ruthlessness 60%; Cold Blood; Relentless Strikes 20% a point for 25.
- Seal Fate 100% ("abilities that add Combo Points").
- Lethality: 20% of the crit bonus, so 220%.
- Puncturing Wounds: Backstab +30%, Mutilate +15% and a 45% second point with no crit condition. So combat.md:100's "a crit never adds a combo point" holds, as does assassination.md:59's "with Malice and Puncturing Wounds you crit a lot".
- Combat: Aggression 15% on Sinister Strike and Eviscerate; Flawless Execution 15; Weapon Expertise 2%; Dual Wield Specialization 50%; Hack and Slash 5%/5%/15%; Flash of Steel +15 and 2.0 s; Adrenaline Rush 100%.
- Subtlety: Opportunity 20% on Backstab, Garrote, Ambush and Mutilate.
- Dirty Deeds rank 1: 10 off Garrote and Cheap Shot, and Garrote from the front. So subtlety.md:33's "no facing requirement anywhere in the rotation" holds, Garrote included.
- Camouflage: speed in stealth. Initiative 100%. Premeditation 2 points for 20 s. Serrated Blades 9% and 30%. Elusiveness 90 s.
- Ghostly Strike 125%/180% and 7 s. Hemorrhage 100%/145% and 15% on Rupture and Garrote. Boot Knife 2 points at 100%.
- Death Mark: 12% a point, double for bleeds, 6 s, the cap scaled by `$mcrit`. Coup de Grace: storage and damage +20%.

**Spells against `trainables/rogue.json`, `spelltips/rogue.json` and `spellconst/rogue.json`**
- Every training level named on the four pages holds:
  - Sinister Strike, Eviscerate and Stealth at 1.
  - Backstab 4, Gouge 6, Evasion 8, Slice and Dice and Sprint 10, Kick 12, Garrote 14, Ambush and Poised 18.
  - Instant Poison 20, Shank 24 and 44, Cheap Shot 26, Deadly Poison 30, Kidney Shot 30 and 50.
- Cheap Shot gives 2 points and a 4 s stun.
- Shank: 5 energy, 15 or 25 back, 10 s.
- Cold Blood and Premeditation are off the global cooldown.
- Blade Flurry: 25 energy, 2 min. Adrenaline Rush: 5 min.
- Kidney Shot: 6 s at five points at 60.
- Blind: 5 min, so three and a half with Elusiveness.
- Thistle Tea 7676: 100 energy on a 5 minute cooldown.

**Engine use.**
- Premeditation's automatic use waits for three points or fewer (`premeditation.go`), which matches subtlety.md:70.
- Adrenaline Rush fires at 45 energy or less (`talents.go`), which matches combat.md:81's "low on energy, never at the cap".

**Stats, races, gear, professions, sources**
- `statPriority` matches the band 60 raid scale factors in all three `bis/rogue-*.json` files, both factions. Haste is outside the normalization.
- Each stat paragraph explains and never ranks.
- Race lines match `round3/race-sims.md:18-23`:
  - Combat: Orc 756.6, Troll 756.0 and Night Elf 755.9 are "within a hair".
  - Subtlety: Dwarf trails on Alliance.
  - Assassination: Night Elf, Dwarf and Human sit within 1.2.
  - Skyborne is named only as the paid option.
- Gear and professions are unchanged since round 9 and hold.
- Sources: all six links returned 200 on 2026-11-04:
  - the two Blizzard articles
  - the Wowhead Skyborne article
  - /planner, /planner?class=rogue and /sim/weights
- No em or en dashes, no site, sim or APL talk, and no quoted nightly numbers.

### Shoulds (wording, not findings)

| id | file:line | note | suggested wording |
|---|---|---|---|
| ROGUE-R11-S1 | subtlety.md:129 vs the Eviscerate card (bands 20 to 38) | The prose says "Eviscerate at five" for 20 to 40. The card says three or more while Rupture ticks. Both fit their context (leveling mobs without a Rupture, against a boss with one), but a reader sees two numbers. | "...and Eviscerate at five on mobs, or at three while Rupture ticks on anything bigger, until Death Mark comes at 40." |

## Decisions

- **Venom wording follows the client modifier kinds.** Proc chance is a flat add (aura 107 misc 18), like Improved Poisons. Damage is a percent add that sums with Vile Poisons (aura 108, same misc and masks). The page says what a player sees: half the off-hand hits proc Instant Poison inside the window. That needs no sim.
- **Blood Splatter stays at 6 s and 9 s.** The spellconst base points (4/6 s) are pre-trait. The trait curves, the rank text and Rupture's tooltip all give 6/9.
- **No build, race, `statPriority` or BiS change.** The BiS files did not move for rogue in this tree. Prose names no BiS row, so nothing is provisional.

## Sims requested

Carried from round 10 and still open:
1. Subtlety raid 60: the shipped APL against the APL with the old Eviscerate line.
2. Subtlety race re-run, both factions, on the Death Mark rotation.
3. Assassination race re-run with the Rupture line.

New: none.

## Sim findings

**SIM-R11-1 (APL, leveling band 10). Subtlety below 20 never spends combo points.**
- Eviscerate (priorityList[6]) and Slice and Dice (priorityList[7]) both require `dotIsActive(Rupture)`. Rupture trains at 20 (`trainables/rogue.json` 1943).
- So at the 10 band the list only builds with Sinister Strike. Its card still prints Eviscerate and Slice and Dice, whose conditions cannot hold below 20.
- Question for the sim lane: add `or not spellIsKnown(Rupture)` to both lines, or add a below-20 Eviscerate line, and measure band 10.

**For the component or data lane, not a guide finding.**
- The Subtlety cards print Ghostly Strike at bands 10 and 20, and Garrote "right after Ghostly Strike" at band 20. In the shipped order Ghostly Strike comes at 26.
- Cause: `trainables/rogue.json` gives talent spells level 0 (Ghostly Strike 14278, Cold Blood 14177, Preparation, Adrenaline Rush), so `apl-level-gate.ts` treats them as open from level 1.
- Assassination's Cold Blood shows at band 30 for a point taken at 33.

## Sources

- "World of Warcraft: Forever Class Deep Dives - Rogue and Warlock", Blizzard (200, 2026-11-04): https://news.blizzard.com/en-us/article/24310968/world-of-warcraft-forever-class-deep-dives-rogue-and-warlock
- "World of Warcraft: Forever Deep Dive Panel Recap", Blizzard (200): https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- "Skyborne First Look: New Neutral Race in World of Warcraft: Forever", Wowhead (200): https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Wowhead Forever tooltips (read 2026-11-04):
  - Instant Poison VI, "Each strike has a 20% chance": https://nether.wowhead.com/forever/tooltip/item/8928
  - Deadly Poison V, "30% chance": https://nether.wowhead.com/forever/tooltip/item/20844
  - Venom, the talent and the buff: https://nether.wowhead.com/forever/tooltip/spell/1329328 and https://nether.wowhead.com/forever/tooltip/spell/1310703
- foreversixty.gg/planner, /planner?class=rogue and /sim/weights (all 200)
- Client data, in `data/builds/1.60.1.70755/`:
  - `talents/rogue.json`, `trainables/rogue.json` and `addon-builds.json`
  - `spelltips/rogue.json`: 1833, 1329283, 1329549, 1328825, 1316709, 1325758, 14185, 438040, 11290, 11275, 408, 8643, 14177, 1241584, 16511, 14278, 1329328
  - `spellconst/rogue.json`: 1310703, 14113, 1329256, 1329286, 14177, 14183, 13877, 13750
  - `addon-data.json` (`rotations.rogue-*`) and `bis/rogue-*.json`
- APLs: `data/curated/apl/rogue-{assassination,combat,subtlety}.json`
- Engine 9f7f1f6e4: `sim/rogue/venom.go:10-22`, `poisons.go:24-38`, `premeditation.go:56-71`, `talents.go:451-522`, `client_values.go:104-111`
- Race runs: `design/reviews/guide-accuracy/round3/race-sims.md:18-23`

## Counts

CONFIRMED 2, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs 1. New sim findings: 1, plus one component item.

Verdict: not zero.
- ROGUE-R11-1 is the old "bonus on a bonus" misreading, on Venom's poison chance.
- ROGUE-R11-2 is a round 10 edit that overstated Death Mark's share of five-point finishers.

## Fix log (round 11)

- ROGUE-R11-1: fixed, assassination.md:56 and :78 with the report's text. Checked: Venom rank text (30% damage, 20% chance), Vile Poisons 20%, Instant Poison "Each strike has a 20% chance" in the 70755 data, Improved Poisons 10%: 20 + 10 + 20 = 50%; :88 puts Instant on the off hand.
- ROGUE-R11-2: fixed, subtlety.md:129 and the rogue-subtlety.json Eviscerate note (notes only): "from 40, five points go into Death Mark while Rupture ticks."
- ROGUE-R11-S1: applied in the same subtlety.md:129 sentence ("Eviscerate at five on mobs, or at three while Rupture ticks on anything bigger, until Death Mark comes at 40").
- SIM-R11-1 and the component item: not guide text, left for those lanes.

# Guide accuracy round 10: Rogue

- Class: rogue. I audited every line of `web/src/content/guides/rogue/index.md`, `assassination.md`, `combat.md` and `subtlety.md` in the round 10 audit tree (origin/main 84460cec8: guide rounds 1 to 9 plus sim round 15). I also read every card note as it is published in `data/builds/1.60.1.70755/addon-data.json` (`rotations.rogue-*`, all seven bands) and the curated APLs in `data/curated/apl/rogue-{assassination,combat,subtlety}.json`.
- Engine pin: 9f7f1f6e4 (`sim/enginever/version.go`). I read the round 15 rogue engine commits at that pin: `db91a347d` (Mutilate off-hand penalty), `a3146478d` (Blood Splatter aura) and `c9759342f` (Preparation resets), plus `sim/rogue/preparation.go`.
- Read first: the rogue reports for rounds 1 to 9 with their fix logs, `round4/rogue-levels.md`, the race-sims files (only round 3 has rogue rows) and `ledger.md`.

## What changed for rogue since round 9

Round 9 audited 645caa983. Since then, sim round 15 landed these rogue commits:
- `6e178531b`: the Combat raid build swaps Efficient Strikes 2/2 for Composure 2/2. The `raidBuild` string changed, and the prose did not.
- `22e8c28bb` and `4eba2d9d1`: Assassination now Ruptures at four or more points when Slice and Dice is up, Rupture is not ticking and Blood Splatter is known. Measured at +0.8 to +1.1% raid 60.
- `ca0e03899`: Assassination's Cold Blood moves to just before Eviscerate.
- `493a3d038`: Subtlety Eviscerates only before Death Mark is learned. It also gates Shank at three points or fewer and Ghostly Strike at four or fewer, and re-applies Garrote after Ghostly Strike. Raid 60 went from 635.8 to 679.1 with the Preparation fix.
- `96f9fd6a9`: rank labels for Shank and Mutilate.
- Engine `db91a347d`: Mutilate's off-hand strike now takes the 50% off-hand weight, and the flat bonus stays whole. This now matches assassination.md:54 and :103, so the carried owner item is closed.
- `addon-builds.json`: only the Combat raid order changed. All three leveling orders are byte-identical to round 9, so round 9's leveling table still holds.

## Decoded builds

Builds were decoded positionally against `talents/rogue.json` tab order (`fs1.ts`).

| Spec | build | raidBuild |
|---|---|---|
| Assassination | 38/11/2 (night-elf), unchanged | 36/10/5, unchanged |
| Combat | 14/34/3 (gnome), unchanged | 12/34/5: **Composure 2/2, Efficient Strikes 0/2** |
| Subtlety | 12/8/31 (night-elf), single build | none |

**Combat raid build, in full:**
- Assassination: Malice 5, Ruthlessness 2, Improved Slice and Dice 3, Lethality 2.
- Combat: Coup de Grace 3, Lightning Reflexes 4, Puncturing Wounds 3, Precision 3, Composure 2, Flawless Execution 1, Weapon Expertise 2, Aggression 3, Dual Wield Specialization 2, Blade Flurry 1, Hack and Slash 4, Flash of Steel 5, Adrenaline Rush 1.
- Subtlety: Opportunity 5.

**Tier gate:** tiers 0 to 2 of Combat hold 3 + 4 + 3 + 3 + 2 = 15, so the build is legal.

**Swap, level 60 against raid:**
- Assassination: +6/-4. Improved Gouge 3, Remorseless Attacks 2 and Ruthlessness +1, against Lethality 2 and Improved Slice and Dice -2.
- Combat: +3/-3. Efficient Strikes 1, Lightning Reflexes +1 and Hack and Slash +1, against Flash of Steel +1 and Composure 2.
- Subtlety: Master of Deception 3 against Opportunity 5.

**Composure (105759):** "Increases the Energy cost reduction of your Poised ability by 5" at rank 1, and by 10 at rank 2. Poised (1328825) is 10, so the raid opener is 20 cheaper.

## Findings

| id | status | kind | file:line | quote | evidence | fix |
|---|---|---|---|---|---|---|
| ROGUE-R10-1 | CONFIRMED | talent | combat.md:52 | "and the fifth points of Hack and Slash and Lightning Reflexes in place of the fifth Flash of Steel and the second Efficient Strikes" | The raid build decodes to Efficient Strikes 0/2 and Composure 2/2. The level 60 build has Efficient Strikes 1/2. So the raid build has no Efficient Strikes to trade, and the two Composure points are missing from the swap. The swap recounts only as shown above (`6e178531b`). | "...a third point of Ruthlessness, and a point of Efficient Strikes plus the fifth points of Hack and Slash and Lightning Reflexes in place of the fifth Flash of Steel and both points of Composure." |
| ROGUE-R10-2 | CONFIRMED | talent | combat.md:63 | "In the raid build its two points and four of Lightning Reflexes are the toll to reach Flawless Execution, Weapon Expertise and Aggression." | The raid build has 0 points in Efficient Strikes. The tier 3 gate is met by Coup de Grace 3, Lightning Reflexes 4, Puncturing Wounds 3, Precision 3 and Composure 2. Composure's rank text is about Poised, so it pays on every pull's opener. | "Efficient Strikes takes 3 energy off Sinister Strike at one point: your builder's discount in the level 60 build. The raid build never presses Sinister Strike, so it skips Efficient Strikes and puts two points in Composure instead: Poised takes 20 energy off your opener, not 10, so the opening Garrote costs 30. Those two points and four of Lightning Reflexes are the toll to reach Flawless Execution, Weapon Expertise and Aggression." |
| ROGUE-R10-3 | CONFIRMED | rotation | assassination.md:80 (with :78 and :64) | ":80 Rupture and Expose Armor do not make the list." / ":78 At four or more combo points, Slice and Dice first if it is down or about to fall off. Otherwise press Cold Blood if it is ready and Eviscerate straight after." / ":64 Blood Splatter keeps your opening Garrote ticking 9 seconds longer" | `rogue-assassination.json` priorityList[3] casts Rupture at four or more points when Slice and Dice is up, Rupture is not ticking and `auraIsKnown(Blood Splatter 1329256)`. That line sits above Cold Blood and Eviscerate. The raid build has Blood Splatter, so Rupture is in the raid rotation (`22e8c28bb`: +0.8 to +1.1% raid 60). The 60 card shows a Rupture line. Blood Splatter's rank text also adds 6 s to Rupture, which :64 leaves out. | :78: "At four or more combo points, Slice and Dice first if it is down or about to fall off. In the raid build, Rupture next whenever it is not ticking: Blood Splatter makes it last 6 seconds longer, and bleeds always crit in Forever. Otherwise press Cold Blood if it is ready and Eviscerate straight after." :80: "Expose Armor does not make the list: Forever's Expose Armor does not stack with Sunder, so leave it to your warriors. Without Blood Splatter, in the level 60 build, skip Rupture too." :64: "Blood Splatter keeps your opening Garrote ticking 9 seconds longer and every Rupture 6 seconds longer, which is what puts Rupture in your raid rotation." |
| ROGUE-R10-4 | CONFIRMED | rotation | subtlety.md:76 and :78 | ":76 While Rupture ticks and Death Mark is on cooldown, spend three or more points on Eviscerate, or on Slice and Dice when that is down or about to fall off." / ":78 the 10 second cooldown means one comes back every few finishers anyway" | `rogue-subtlety.json` priorityList[6] Eviscerate is gated on `not spellIsKnown(Death Mark)`, so it is never cast from 40. `493a3d038` explains why: with an Eviscerate line, a five-point Death Mark came only when a two-point jump landed on five (6.3 of about 17 marks a fight at 60). Death Mark costs 25 on a 10 s cooldown and stores 12% per point (`spelltips/rogue.json` 1325758). The budget shows there are no spare points. Over 20 s, Shank with Boot Knife gives 4, Ghostly Strike 1 and about four Hemorrhages 4. That is about 9 or 10 points, against 14 for one Rupture and two five-point marks. | :76: "While Rupture ticks, put three or four points into Slice and Dice only when it is down or about to fall off; otherwise keep building to five for the next Death Mark. Eviscerate is your finisher only until Death Mark at 40." :78: "Mark at five points, not four: the storage is per combo point, and the 10 second cooldown is back by the time you have built five again." |
| ROGUE-R10-5 | CONFIRMED | card note | `data/curated/apl/rogue-subtlety.json` priorityList[6].notes (Eviscerate card, bands 40, 50, 60) | "Eviscerate at three or more combo points while Rupture is ticking." | The action also needs `not spellIsKnown(Death Mark)`. `apl-level-gate.ts` treats a talent spell as open from its learn level, so the card still prints at 40 to 60, where the rotation never casts it. The note states a cast the card does not make. | "Eviscerate at three or more combo points while Rupture is ticking, only before you have Death Mark: from 40, every five points go into Death Mark." |
| ROGUE-R10-6 | CONFIRMED | card note | `rogue-subtlety.json` priorityList[1].notes (Shank) and [2].notes (Ghostly Strike), bands 10 to 60 | "Shank on cooldown: ..." / "Ghostly Strike on cooldown: ..." (the cards print "On cooldown") | The actions are gated by `currentComboPoints <= 3` (Shank) and `<= 4` (Ghostly Strike). subtlety.md:76 states the gates, so the cards contradict the prose under them. | Shank: "Shank on cooldown at three combo points or fewer: an off-hand stab that cannot miss, pays back more energy than it costs and, with Boot Knife, adds two combo points, so at four or five you finish first." Ghostly Strike: "Ghostly Strike on cooldown at four combo points or fewer: it hits harder than Hemorrhage for 10 more energy (Hemorrhage costs 30 with Efficient Strikes) and gives a combo point." |

### Re-checked and holding

**Round 9 decisions**
- Boot Knife 2 points.
- Premeditation needs no Stealth.
- Lethality is 220%.
- Eureka! covers non-periodic damaging abilities.

**Card notes for Combat (all bands)**
- Shank every 8 s with Flash of Steel 5.
- Eviscerate at five.
- Slice and Dice at one or two.
- Backstab with Opportunity and Lethality in the raid build.
- Sinister Strike held to 60.

**Card notes for Assassination**
- Cold Blood is now right before Eviscerate, which matches :78.
- Rupture's note names Blood Splatter, so the leveling-band Rupture card states its own gate truthfully.

**Mutilate off-hand split.** assassination.md:54 and :103 now match engine `db91a347d` and the SpellMisc REQUIRES_OFF_HAND_WEAPON flag.

**Preparation.** subtlety.md:42 and :86 ("2 + 2 + 1 = 5") hold against `preparation.go` (it resets Premeditation, Shank, Ghostly Strike and Death Mark) and the 14185 text.

**Subtlety pull and priority against the APL.** Pre-pull is Stealth at -3 s, Premeditation at -2 s and Garrote at 0. :76 Shank at three or fewer, Ghostly Strike never at five, Rupture at four or more and Death Mark at five with Rupture all match.

**Source links.** All six returned 200 on 2026-11-04:
- the two Blizzard articles
- the Wowhead Skyborne article
- /sim/weights, /planner and /planner?class=rogue

**Prose rules.** No em or en dashes, no site or sim talk, and no quoted nightly numbers.

### Shoulds (wording, not findings)

| id | file:line | note | suggested wording |
|---|---|---|---|
| ROGUE-R10-S1 | subtlety.md:129 | "and Eviscerate at five" covers 20 to 60. From 40, Death Mark takes the five-point finishes on anything that lives past 6 seconds. | "...keep Slice and Dice up on anything that lives, and Eviscerate at five until Death Mark comes at 40; after that, mark anything that will live 6 more seconds and Eviscerate the rest." |
| ROGUE-R10-S2 | `rogue-subtlety.json` top-level `notes` | It still says "Eviscerate and Slice and Dice spend the combo points that build while it runs". The note is not rendered, so this is not a finding. | Match the new gating. |
| ROGUE-R10-S3 (carried R9-S1) | index.md:37, combat.md:112 | "Expansive Mind gives you 5% more energy." | "**Expansive Mind** makes your energy bar 5% bigger." |
| ROGUE-R10-S4 (carried R9-S2) | assassination.md:80 | "from nothing to 2 or 3, then to 4 or 5" | As in round 9. |
| ROGUE-R10-S5 (carried R9-S3) | combat.md:104 | "only needs to be fast and hit hard" | As in round 9. |

## Decisions

- **Combat swap prose follows the shipped raid string (Composure 2/2).** Composure is not a dead pick: the raid opener is always a Poised Garrote. The sim lane measured the swap within error, so the prose follows the string.
- **Assassination prose takes Rupture into the raid priority.** Measured +0.8 to +1.1% at raid 60, both seeds and both factions. The level 60 build has no Blood Splatter, so leveling and level 60 play keep no Rupture, as the APL gate does.
- **Subtlety pools for Death Mark.** I side with the shipped rotation over the old prose: Death Mark's per-point storage on a 10 s cooldown leaves no spare points for Eviscerate. Sims requested item 1 confirms that isolated.
- **No build, race or `statPriority` change.** The BiS files did not change for rogue in this tree. Prose names no BiS row, so there is nothing provisional to file.

## Sims requested

1. Subtlety raid 60, both seeds: the shipped APL against the same APL with the old Eviscerate line restored (three or more points, Rupture ticking, Death Mark on cooldown). This confirms ROGUE-R10-4 in isolation from `493a3d038`'s other changes.
2. Subtlety race re-run (both factions) on the new rotation. Round 3's order (Night Elf, Human; Undead, Orc) predates a +6.8% rotation change built on Death Mark, whose cap scales with crit. This is relevant to subtlety.md:109 and index.md:37.
3. Assassination race re-run (both factions) with the Rupture line. Round 3 had Night Elf, Dwarf and Human within a hair, so the "play the one you like" call at assassination.md:111 should be re-checked. This can wait for the final best-in-slot re-rank like the others.

## Sim findings

None new at 9f7f1f6e4.
- Mutilate's off-hand penalty is fixed (`db91a347d`).
- Preparation is now used in the fight and resets every rogue cooldown (`c9759342f`).
- The Subtlety Shank and Ghostly Strike gating is in the APL.

## Sources

- "World of Warcraft: Forever Class Deep Dives - Rogue and Warlock", Blizzard (200, 2026-11-04): https://news.blizzard.com/en-us/article/24310968/world-of-warcraft-forever-class-deep-dives-rogue-and-warlock
- "World of Warcraft: Forever Deep Dive Panel Recap", Blizzard (200): https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- "Skyborne First Look: New Neutral Race in World of Warcraft: Forever", Wowhead (200): https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- foreversixty.gg/sim/weights, /planner, /planner?class=rogue (all 200)
- Client data, in `data/builds/1.60.1.70755/`:
  - `talents/rogue.json`: Composure 105759, Efficient Strikes 105741, Blood Splatter 105718
  - `spelltips/rogue.json`: Poised 1328825, Death Mark 1325758, Rupture 11275, Shank 1329549, Ghostly Strike 14278, Hemorrhage 16511, Mutilate 1241584, Preparation 14185
  - `addon-data.json` (`rotations.rogue-*`)
  - `addon-builds.json`
- APLs: `data/curated/apl/rogue-{assassination,combat,subtlety}.json`; commits 6e178531b, 22e8c28bb, 4eba2d9d1, ca0e03899, 493a3d038, 96f9fd6a9.
- Rendering: `web/src/lib/guides/rotation-lists.ts` (`clauseOf`) and `web/src/lib/bis/apl-level-gate.ts`.
- Engine 9f7f1f6e4: `sim/rogue/preparation.go`; commits db91a347d, a3146478d, c9759342f.
- Round 3 race runs: `design/reviews/guide-accuracy/round3/race-sims.md:18-23`

## Counts

CONFIRMED 6, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs 5, three of them carried from round 9.

Verdict: not zero. Every finding comes from sim round 15 changing a build string or rotation without the prose and card notes following it.

## Fix log (round 10)

Findings:
- ROGUE-R10-1: fixed, combat.md:52. The level 60 swap now trades a point of Efficient Strikes and the fifth Hack and Slash and Lightning Reflexes against the fifth Flash of Steel and both points of Composure.
- ROGUE-R10-2: fixed, combat.md:63, using the audit wording. The figures are checked against Poised 1328825 (10 energy), Composure rank 2 (+10) and Garrote 11290 (50 energy). The raid build has no Dirty Deeds, so the opener costs 30. combat.md:73 also says that Composure takes Poised to 20 in the raid build.
- ROGUE-R10-3: fixed. assassination.md:64 adds Blood Splatter's 6 s on Rupture, which is what puts Rupture in the raid rotation. assassination.md:78 adds Rupture after Slice and Dice in the raid build. assassination.md:80 drops "Rupture ... do not make the list"; Expose Armor stays off, and the level 60 build without Blood Splatter skips Rupture.
- ROGUE-R10-4: fixed. subtlety.md:76 no longer has an Eviscerate step from 40 on: Slice and Dice is at three or more only when it is down or about to fall off, otherwise you build to five for Death Mark. subtlety.md:78 has the audit's five-not-four line.
  - I added the Garrote re-application after Ghostly Strike to subtlety.md:76. It matches APL priorityList[3] and was drafted by the sim lane.
- ROGUE-R10-5: fixed, rogue-subtlety.json:316, the Eviscerate note. It is now gated "only before you have Death Mark".
- ROGUE-R10-6: fixed. rogue-subtlety.json:95 is the Shank note (three or fewer, finish first at four or five). rogue-subtlety.json:138 is the Ghostly Strike note (four or fewer).

Shoulds:
- S1: applied, subtlety.md:129. Eviscerate at five until Death Mark at 40, then mark anything that will live 6 more seconds.
- S2: applied, rogue-subtlety.json:4, the top-level notes. Death Mark (or Eviscerate before you have it) and Slice and Dice spend the points.
- S3: applied, index.md:37 and combat.md:112: "makes your energy bar 5% bigger".
- S4: applied, assassination.md:80, in round 9's Ruthlessness wording.
- S5: applied, combat.md:104, in round 9's wording (Instant Poison and every Shank).
- Sim lane item 7: applied, rogue-assassination.json:4, the top-level notes now carry the Blood Splatter Rupture line. The existing Rupture and Garrote card notes were reviewed for voice and are true as written, so I left them unchanged.

Stat priority rule:
- All three spec pages already explain rather than rank, and none of them names a ranking word. The prose test passes for rogue (`-t rogue`).
- No enchant or gear prose ranks by a nightly value.
- The race prose keeps the measured race order (round 7 rule).

Other:
- Decisions: all applied as written. Builds, actions, statPriority and BiS files are untouched.
- updated: was already 2026-11-04 on all four guide files.

SIM findings: none new. Sims requested 1 to 3 stay open. The subtlety.md:78 wording holds whatever item 1 shows, because it rests on the per-point storage.

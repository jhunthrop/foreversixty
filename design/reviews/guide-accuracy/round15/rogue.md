# Guide accuracy round 15: Rogue

Audited on the round 15 audit tree (main 4600b7d82: fourteen guide rounds, sim round 17, launch state), client 1.60.1.70755, engine pin b2c4a2d05 (`sim/enginever/version.go`). Sim round 17 changed no rogue engine code, so nothing the page says about mechanics moved with the pin.

**Scope.**
- Pages: `web/src/content/guides/rogue/{index,assassination,combat,subtlety}.md`, every line.
- Card notes: every rogue row in `data/builds/1.60.1.70755/addon-data.json` (`rotations.rogue-*`, every band), regenerated at 38b890465 (round 17 regen).
- APLs: `data/curated/apl/rogue-{assassination,combat,subtlety}.json`.
- Shipped orders: `addon-builds.json` (index 0 = level 10).
- BiS: `bis/rogue-*.json`, every band on both factions (weapons) and band 60 raid (stats, gear prose).

**Read first:**
- the rogue reports for rounds 1 to 14 with their fix logs, and `round4/rogue-levels.md`
- `round2` to `round5/race-sims.md` (`round3/race-sims.md` is still the only one with rogue rows)
- `ledger.md` through round 14
- `design/reviews/2026-11-04-sim-audit-round17/rogue.md` (sim round 17, which landed in this tree)

## What changed since round 14's tree

`git diff d7fc51fe7 HEAD` touches rogue only through sim round 17 (600c52c0f, 7cb1b23fc, 080b12930, regen 38b890465):
- **Subtlety build.** `035131031321001151` became `035230031321001151`: Improved Ambush 1 out, Dirty Deeds 2/2 in. The shipped walk now has Dirty Deeds at 18 and 19 and Camouflage from 24.
- **subtlety.md.** Changed lines:
  - Strengths ("level 60 rotation")
  - the talents paragraph (Dirty Deeds, "Improved Ambush is skipped")
  - the priority ("its 30 energy")
  - Levels 1 to 19 (Dirty Deeds at 18 and 19)
  - Levels 20 to 60 ("Camouflage, from 24"; Backstab until Hemorrhage)
- **assassination.md.** Two changes:
  - :78 now has round 14's "its ticks can crit"
  - :131 adds the group Backstab sentence
- **APLs.** Assassination and Subtlety build with Backstab before their talent builder. Sinister Strike is the fallback when Backstab cannot be cast. New Backstab and Sinister Strike notes.
- **addon-data.json.** Regenerated. Round 14's two note fixes are now on the cards: the Subtlety Eviscerate note ends "those points go into Death Mark instead", and the Assassination Rupture note says "its ticks can crit".

## Builds decoded

Decoded positionally against `talents/rogue.json` tab order (tier, then column; matches round 13's decode).

- **Subtlety `build` 12/8/31 (new):**
  - Assassination: Malice 5, Ruthlessness 3, Improved Slice and Dice 2, Blood Splatter 1, Lethality 1.
  - Combat: Coup de Grace 3, Efficient Strikes 2, Precision 3.
  - Subtlety: Master of Deception 3, Opportunity 5, **Dirty Deeds 2**, Camouflage 3, Initiative 3, Premeditation 1, Serrated Blades 3, Elusiveness 2, Ghostly Strike 1, Preparation 1, Hemorrhage 1, Boot Knife 5, Death Mark 1.
  - Legal. Subtlety tiers 1 and 2 hold 13, tiers 1 to 6 hold 30 for Death Mark. Assassination tier 2 holds 10 for Blood Splatter, Combat tier 1 holds 5 for Precision.
- **Unchanged:** Assassination `build` and `raidBuild`, and Combat `build` and `raidBuild`, are as round 13 decoded them.

**Shipped orders** (`addon-builds.json`), with the levels the prose names:
- **Subtlety:**
  - Master of Deception 10-12, Opportunity 13-17, Dirty Deeds 18-19.
  - Premeditation 20, Serrated Blades 21-23, Camouflage 24, Elusiveness 25, Ghostly Strike 26, Camouflage 27-28.
  - Initiative 29, Hemorrhage 30, Initiative 31-32, Elusiveness 33, Preparation 34.
  - Boot Knife 35-39, Death Mark 40.
  - Then Assassination 41-52 and Combat 53-60.
  - Leveling and raid orders are identical.
- **Assassination leveling:**
  - Improved Gouge 10-12, Remorseless Attacks 13-14, Malice 15-19.
  - Venom 30, Cold Blood 33.
  - Relentless Strikes 39 and 41-44, Mutilate 40.
- **Combat leveling:**
  - Coup de Grace 10-12, Efficient Strikes 13, Lightning Reflexes 14-18, Puncturing Wounds 19-21.
  - Aggression 28-30, Adrenaline Rush 40.
  - Improved Gouge 44-46, Remorseless Attacks 47-48.
  - Combat raid takes Puncturing Wounds 17-19.

Every level the prose names matches.

## Findings

| id | status | kind | file:line | quote | evidence | fix |
|---|---|---|---|---|---|---|
| ROGUE-R15-1 | CONFIRMED | card note (new in sim round 17) | `rogue-subtlety.json` Backstab notes, rendered on the Subtlety Backstab row of every band 10 to 60 in `addon-data.json` | "Backstab from behind with a dagger in your main hand is the builder until Hemorrhage: it hits far harder than Sinister Strike, and Opportunity adds 20% to it." | The shipped Subtlety order takes Opportunity at 13 to 17 (`addon-builds.json` rogue-subtlety indices 3 to 7 = [3,1,3]). On the cards where this row matters most, the player has no Opportunity at all (bands 10 and 12), 2 points (band 14, 8%) or 4 points (band 16, 16%). Opportunity's rank text is 4% a point ("by 4%" ... "by 20%", `talents/rogue.json`, spell 14057). The Combat note on the same spell says it conditionally ("in the raid build Opportunity and Lethality land on it too"). | Note: "Backstab from behind with a dagger in your main hand is the builder until Hemorrhage: it hits far harder than Sinister Strike, and every point of Opportunity, from 13, adds 4% to it." Regenerate `addon-data.json` and Data.lua. |
| ROGUE-R15-2 | CONFIRMED | card note | `rogue-combat.json` Backstab notes, rendered on the Combat Backstab row of every band 10 to 60 | "Backstab with a dagger in the main hand: Puncturing Wounds raises its crit chance and can give it a second combo point, and in the raid build Opportunity and Lethality land on it too." | Neither Combat order has Puncturing Wounds before 17: leveling takes it at 19 to 21, raid at 17 to 19 (`addon-builds.json` rogue-combat). On the band 10, 12, 14 and 16 cards the note credits Backstab with a talent the player does not have. The Opportunity and Lethality half is already conditional ("in the raid build"); the Puncturing Wounds half is not. This is the same pattern as R15-1, and earlier rounds did not file it. | Note: "Backstab with a dagger in the main hand: once you have Puncturing Wounds it crits more often and can give a second combo point, and in the raid build Opportunity and Lethality land on it too." Regenerate `addon-data.json` and Data.lua. |
| ROGUE-R15-3 | LIKELY | gear (prose against the page's own tables) | combat.md:132 (with :104) | ":132 Level as Combat with a sword, axe or mace and Sinister Strike: mobs face you while you level, and Backstab needs you behind the target." | Every Combat leveling band on the page's table holds a main-hand dagger on both factions:<br>- 20: Blackfang<br>- 30: Toxic Revenger<br>- 40 and 50: Gut Ripper<br><br>The page gives a reader picking band 20 to 50 a sword in the prose and a dagger in the table, with nothing to square them. Both are right in their own setting: the tables are ranked from behind the target (sim round 17 K2: 37 to 41 Backstabs a fight at every band). Sinister Strike on a sword is the solo questing play. Assassination got exactly this bridge in sim round 17 (assassination.md:131 "In a group, where the tank holds the mob ..."); Combat did not. | Add after the first sentence of :132: "In a dungeon, where the tank holds the mob and you stand at its back, a dagger and Backstab beat the sword at every level, which is what your best in slot table lists." |

## Re-check of sim round 17's text (new work)

- **subtlety.md:33.** "No facing requirement anywhere in the level 60 rotation" holds:
  - Hemorrhage, Ghostly Strike, Rupture, Shank and Death Mark have no positional.
  - With Dirty Deeds 2/2, Garrote "no longer requires you to be behind your target" (14082, both ranks).
  - The band 60 sets hold a sword (Teebu's Blazing Longsword) and a mace (Persuader), so Backstab and Ambush cannot be cast.
- **subtlety.md:62.** "Both points of Dirty Deeds take 20 energy off Garrote and Cheap Shot ... and let Garrote go in from the front": rank 2 reads "by 20, and your Garrote ability no longer requires you to be behind your target". "Improved Ambush is skipped" matches the decode. The gate still holds without it (tiers 1 and 2 hold 13).
- **subtlety.md:80.** "the Garrote never gets its 30 energy": Garrote 11290 costs 50, and Dirty Deeds 2/2 takes 20. Holds.
- **subtlety.md:127.** Dirty Deeds is at 18 and 19 in the shipped order. Rank 1 already removes the behind requirement; rank 2 makes it 20. Holds. Opportunity 13 to 17 and Master of Deception 10 to 12 hold.
- **subtlety.md:133.**
  - "Camouflage, from 24" holds (index 14).
  - "Until Hemorrhage, Backstab with a dagger ... Sinister Strike when you are not" matches the new Backstab and Sinister Strike cards. The Sinister Strike card's "or the target is facing you" matches the prose.
- **assassination.md:78.** "its ticks can crit" (round 14 R14-2) is now on both the page and the card.
- **assassination.md:131.** The group Backstab sentence matches the new Backstab card ("builder until Mutilate") and the Sinister Strike card ("only when Backstab cannot be cast: no dagger in the main hand, or the target is facing you"). It does not contradict "From 20 to 40, Sinister Strike is your everyday builder on mobs that face you".
- **Round 14's fixes on the cards.**
  - Subtlety Eviscerate: "from 40, those points go into Death Mark instead", which agrees with the Death Mark row (four or more) and subtlety.md:80, :82 and :133.
  - Assassination Rupture: "its ticks can crit in Forever".

## Shoulds (wording, not findings)

| id | file:line | note | suggested wording |
|---|---|---|---|
| ROGUE-R15-S1 | subtlety.md:80 | Carried from R14-S1: "every point Rupture does not need goes into the next mark" sits right before spending three or more on Slice and Dice. | "Once you have Death Mark, Eviscerate is off your bar, even with Death Mark on cooldown: the points Rupture and Slice and Dice do not need go into the next mark." |
| ROGUE-R15-S2 | subtlety.md:90 | Carried from R13-S1 and R14-S2: "Death Mark is your only damage cooldown" while Shank and Ghostly Strike are damage on cooldowns too. | "Death Mark is your burst, and it comes back every 10 seconds, so there is nothing to hold: ..." |
| ROGUE-R15-S3 | subtlety.md:117 | Carried from R14-S3: the Leatherworking reason names Devilsaur and Stormshroud. The Subtlety band 60 raid list uses neither; its leatherworking crafts are Outlaw's Collar and Dawn Armor. True, but it points at the wrong pieces. | "**Skinning** and **Leatherworking** pay for themselves: Outlaw's Collar, Dawn Armor and the Devilsaur set are leatherworking crafts." |
| ROGUE-R15-S4 | `rogue-combat.json` Sinister Strike note | Sim round 17's wow-player flagged it, and it was left for the guide lane. "held to 60 energy so a dagger rogue never spends Backstab's energy here" describes the APL's gate, not play. A sword rogue reading the band 10 to 38 card is told to pool to 60 before every Sinister Strike, with no reason given. It is not false (pooling under the cap costs nothing). | "Sinister Strike only when Backstab cannot be cast: no dagger in the main hand, or the target is facing you." (Assassination's and Subtlety's wording.) |

## Checked and right (re-derived this round)

**Spells and levels** (`trainables/rogue.json`, `spelltips/rogue.json`):
- Trainer levels:
  - Backstab 4, Gouge 6, Evasion 8, Slice and Dice and Sprint 10, Kick 12, Garrote 14, Ambush 18, Poised 18.
  - Instant Poison 20, Premeditation 20, Shank 24 (rank 2 at 44, 15 then 25 energy, 10 s, cannot miss), Cheap Shot 26.
  - Deadly Poison 30, Hemorrhage 30, Venom 30, Mutilate and Death Mark 40.
- Cheap Shot costs 60 and gives 2 combo points (3 with Initiative). Ambush costs 60 (50 with Poised).
- Ghostly Strike costs 40 and Hemorrhage 35 (30 with Efficient Strikes 2/2), so "10 more energy" holds.
- Kidney Shot rank 2 lasts up to 6 s at five. Redirect stores the points for 15 s.
- Blind and Vanish are on 5 min cooldowns, 3.5 min with Elusiveness 2/2.
- Preparation has a 10 min cooldown. Premeditation gives 2 points that last 20 s, on a 2 min cooldown.
- Mutilate rank 3: 90% plus 79.2 a hand.
- Instant Poison 20% plus Improved Poisons 10% plus Venom 20% is half the off-hand hits.

**Talent texts:**
- Efficient Strikes is 3 at one point and 5 at two.
- Composure 2/2 gives Poised +10 (opening Garrote 30).
- Flash of Steel 5/5 is +15 energy and -2 s (40 every 8 s).
- Lethality is +20% on the crit bonus (220%).
- Remorseless Attacks 2/2 gives 40%. Improved Gouge 3/3 gives 1.5 s. Vigor 1 gives 5 energy.
- Ruthlessness 60%, Relentless Strikes 20% a point, Puncturing Wounds 30% and 45%, Hack and Slash 5% at five.
- Dual Wield Specialization 50% at 2/2, Weapon Expertise 2%, Coup de Grace (Eviscerate and Death Mark 20%).
- Boot Knife, Serrated Blades 9% and 30%, Blood Splatter 6 s and 9 s.

**Build paragraphs.**
- Assassination's 7-for-7 swap recounts.
- Combat's 14/34/3 against 12/34/5 swap recounts: Assassination 5 for 5, Combat 3 for 3.

**Races.** Every line still matches `round3/race-sims.md:18-23`, and no newer measured rogue race result exists in the tree. Racial texts match `races.json`: Touch of the Grave, Elune's Light, Eureka!, Expansive Mind, Ruthless, Artisan's Precision, Wind Blessed and Will of the Forsaken. Skyborne is named only as the paid Heroic Pack option.

**Gear and stats.** The band 60 raid weapons and stat prose are as round 14 recorded. No BiS file moved since round 14.

**Sources.** All six frontmatter URLs returned 200 on 2026-11-04:
- the two Blizzard articles
- the Wowhead Skyborne article
- foreversixty.gg /sim/weights, /planner and /planner?class=rogue

**Voice.** No em or en dashes, no site, sim or APL talk, no quoted nightly numbers in the new round 17 prose.

## Decisions

- **Subtlety's band 20 table, not its prose, is wrong.** subtlety.md:123 and :133 say to keep a dagger in the main hand and Backstab until Hemorrhage. The band 20 table holds Diamond Hammer (a mace) in the main hand on both factions. It was ranked under the old Sinister Strike-only rotation, and sim round 17 measured a dagger set +5 to +10% at 20 to 29. The prose stands; the table goes to the data lane (Sims requested 1). Not filed as a guide finding.
- **Card notes that name a talent must hold on every band they render on.** R15-1 and R15-2 make the talent conditional, as the Combat note's "in the raid build" already does. This is separate from the component-lane issue of whole talent-spell rows showing before the talent (Ghostly Strike, Cold Blood, Garrote's Ghostly Strike note), which stays with that lane.
- **Combat leveling keeps the sword for solo play** and adds the dungeon caveat (R15-3). This is the same split Assassination now states.
- **No build, order, `statPriority` or race change.**

## Sims requested

1. Data lane, new: re-rank `bis/rogue-subtlety.json` band 20 (both factions) on the round 17 rotation. A dagger main hand should replace Diamond Hammer (sim round 17 S2, "Measure at refresh").
2. Carried: Subtlety BiS and weights on the new build (Dirty Deeds 2/2, no Improved Ambush). The band 60 tables were ranked on the old string. The stat order the prose explains is unaffected.
3. Carried: Subtlety race re-run on both factions, on the round 16 and 17 rotation and build.
4. Carried: Assassination race re-run with the Rupture line. Assassination bands 20 and 30 weights now that Backstab builds.
5. Carried: re-rank `bis/rogue-combat.json` band 60 raid on the shipped `raidBuild`.

## Sim findings

None new. Carried, not re-filed:
- Sim round 17 open items:
  - K1: Death Mark storage, and the Mutilate and Shank off-hand penalty. Needs a level 40+ log.
  - K2/C1: Combat Eviscerate at four by gear.
  - X1: Occult Poison learnability.
- Round 16 S4: Hemorrhage's dagger hand.
- The page makes no claim on any of these that the client contradicts.

## Sources

- "World of Warcraft: Forever Class Deep Dives - Rogue and Warlock", Blizzard (200, 2026-11-04): https://news.blizzard.com/en-us/article/24310968/world-of-warcraft-forever-class-deep-dives-rogue-and-warlock
- "World of Warcraft: Forever Deep Dive Panel Recap", Blizzard (200): https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- "Skyborne First Look: New Neutral Race in World of Warcraft: Forever", Wowhead (200): https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- foreversixty.gg: /sim/weights, /planner, /planner?class=rogue (all 200)
- Client data in `data/builds/1.60.1.70755/`:
  - `talents/rogue.json`: Dirty Deeds 14082, Opportunity 14057, Puncturing Wounds 1224716, Efficient Strikes, Composure, Flash of Steel, Vigor
  - `trainables/rogue.json`
  - `spelltips/rogue.json`: 1833, 11290, 703, 1329283, 1329549, 8643, 438040, 1328825, 14185, 14183, 1325758, 14278, 16511, 25300, 8676, 11340, 25351
  - `races.json`
  - `addon-builds.json`, `addon-data.json`, `bis/rogue-*.json`, `items.json`
- APLs: `data/curated/apl/rogue-{assassination,combat,subtlety}.json`
- Commits: 600c52c0f, 7cb1b23fc, 080b12930, 38b890465 (sim round 17 rogue and regen)
- Previous rounds: `round14/rogue.md`, `round13/rogue.md`, `round3/race-sims.md`, `2026-11-04-sim-audit-round17/rogue.md`, `ledger.md`

## Counts

CONFIRMED 2, LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs: 4. New sim findings: 0.

Verdict: NOT zero for rogue. The page prose is clean. The two confirmed errors are card notes that credit a talent on bands before the shipped order takes it. One came in with sim round 17's new Subtlety Backstab note; the other is an older Combat note no round caught.

## Fix log (round 15)

- ROGUE-R15-1 fixed: rogue-subtlety.json Backstab note now reads "every point of Opportunity, from 13, adds 4%".
- ROGUE-R15-2 fixed: rogue-combat.json Backstab note now reads "once you have Puncturing Wounds".
- ROGUE-R15-3 fixed, combat.md:132: added the dungeon dagger sentence.
- addon-data.json and Data.lua are not regenerated here. The integrator regenerates them.
- S1 to S4 left.

# Guide accuracy round 14: Rogue

Audited on the round 14 audit tree (main d7fc51fe7: thirteen guide rounds, sim round 16, launch state), client 1.60.1.70755, engine pin f637ec32c (`sim/enginever/version.go`; source read at `scratchpad/gr14/engine`, HEAD f637ec32c).

**Scope.**
- Pages: `web/src/content/guides/rogue/{index,assassination,combat,subtlety}.md`, every line.
- Card notes: every rogue row in `data/builds/1.60.1.70755/addon-data.json` (`rotations.rogue-*`, every band), regenerated at 6e57a9c60.
- APLs: `data/curated/apl/rogue-{assassination,combat,subtlety}.json`.
- BiS: `bis/rogue-*.json` (de044125c picks), band 60 raid entries on both factions, which are the rows the guide table opens on (`lib/bis/presets.ts` selectBand: raid unless asked).

**Read first:**
- the rogue reports for rounds 1 to 13 with their fix logs, `round4/rogue-levels.md`
- `round3/race-sims.md`, still the only race file with rogue rows
- `ledger.md` through round 13
- `design/reviews/2026-11-04-sim-audit-round16/rogue.md` (sim round 16, which landed in this tree)

## What changed since round 13's tree

`git diff 604b78a80 HEAD` touches rogue through sim round 16 and its merge with guide rounds 10 to 13 (c4483a3a6):
- `subtlety.md`: the round 16 text (two bleeds; Garrote after every Ghostly Strike; Hemorrhage held for it; Death Mark at four; "every Garrote 9"; Preparation "five combo points"; the leveling hand-off at 26 and 40).
- `rogue-subtlety.json`: Garrote line under Ghostly Strike, Hemorrhage held while 1329326 is up, Death Mark at four or more.
- `addon-data.json` regenerated from those APLs.

The merge resolved three rogue text conflicts in favour of the guide-round side, which undid two round 16 fixes. Both are filed below. Comparing 8160ad885 (round 16's text) with HEAD:
- Assassination Rupture note and assassination.md:78: round 16 wrote "bleeds can crit"; HEAD has "bleeds always crit" again.
- Subtlety Eviscerate note: round 16 wrote "only before you have Death Mark."; HEAD has the round 10 to 13 tail "from 40, five points go into Death Mark while Rupture ticks", which the same merge's Death Mark line (four or more) now contradicts.

No rogue build string, `addon-builds.json` rogue row or `bis/rogue-*.json` moved.

## Builds decoded (unchanged, re-verified)

Same strings and decodes as round 13 (all legal on 70755, tier gates hold, swap paragraphs recount 7 for 7 and 12 for 12, shipped orders match every level the prose names). Not repeated here; see `round13/rogue.md` "Builds decoded".

## Findings

| id | status | kind | file:line | quote | evidence | fix |
|---|---|---|---|---|---|---|
| ROGUE-R14-1 | CONFIRMED | rotation (card note contradicts page) | `rogue-subtlety.json` priorityList[6] notes, rendered on the Subtlety Eviscerate row of every band 10 to 60 in `addon-data.json` | "Eviscerate at three or more combo points while Rupture is ticking, only before you have Death Mark: from 40, five points go into Death Mark while Rupture ticks." | The Death Mark line two rows above (priorityList[5]) casts at `currentComboPoints >= 4`, and its own note says "Death Mark at four or more combo points". The page says "Death Mark at four or more points while Rupture ticks" (subtlety.md:80), "Mark at four, and do not hold it for a fifth point" (:82) and "put four or more points into Death Mark" (:133). Round 16 (S2, 8160ad885) had already cut the tail; the merge c4483a3a6 put the round 10 to 13 wording back. | Note: "Eviscerate at three or more combo points while Rupture is ticking, only before you have Death Mark: from 40, those points go into Death Mark instead." Regenerate `addon-data.json` and Data.lua. |
| ROGUE-R14-2 | CONFIRMED | mechanic (regression) | assassination.md:78; `rogue-assassination.json` Rupture notes (rendered on the Assassination Rupture row, bands 20 to 60) | ":78 Blood Splatter makes it last 6 seconds longer, and bleeds always crit in Forever." / note: "...it lasts 6 seconds longer, and bleeds always crit in Forever." | Every Rupture tick does not crit. The Forever rule is that bleeds are able to crit: the owner's words are "dots can crit in forever" (`forever-dots-crit.md`), and the engine's `periodicDamageCanCrit` (`sim/core/dot.go:444-453` at f637ec32c) lets bleed ticks roll the melee table under `Dot.OutcomeTick`, so a tick crits at your crit chance. The same site's subtlety.md:101 says it right ("your Rupture and Garrote ticks crit, and Death Mark's cap rises with your crit chance"). Round 16 (X1, 8160ad885) changed both to "can crit" after a wow-player MUST; the merge c4483a3a6 reverted them. | :78: "In the raid build, Rupture next whenever it is not ticking: Blood Splatter makes it last 6 seconds longer, and its ticks can crit in Forever." Note: "Rupture at four or more combo points while Slice and Dice is up and Rupture is not ticking, only with Blood Splatter (the raid build): it lasts 6 seconds longer, and bleeds can crit in Forever." |

## Re-check of round 13 and round 16 text that is on the page now

- **subtlety.md:41-42.** "two bleeds" (Garrote and Rupture) and "one extra round of combo points" both hold.
- **subtlety.md:62.** "every Garrote 9": Blood Splatter 1329256 reads "Rupture ability by 6 sec and your Garrote ability by 9 sec". Holds.
- **subtlety.md:70.**
  - Premeditation 2, then Garrote 1 plus Initiative 3/3 (100%) gives the fourth point. Rupture at four matches priorityList[4].
  - "Death Mark as soon as you have four or more again with Rupture ticking" matches priorityList[5].
  - Prepull (Stealth -3 s, Premeditation -2 s, Garrote 0 s) matches the APL.
- **subtlety.md:78.**
  - Ghostly Strike's rank text: "your next Ambush or Garrote cast within 7 sec to not require Stealth". Dirty Deeds 1/2 (14082) removes the behind requirement, so the Garrote goes in from the front.
  - "with Initiative it gives two combo points, which Hemorrhage does not": Garrote 1 + Initiative 1 against Hemorrhage's "Awards 1 Combo Points". Holds.
  - "At exactly four ... Garrote first" matches the APL order (Garrote line [3] above Rupture [4], both live at four).
- **subtlety.md:80.**
  - "the Garrote never gets its 40 energy": Garrote 11290 costs 50, Dirty Deeds 1/2 takes 10. Holds.
  - Hemorrhage's hold matches priorityList[8] (not while 1329326 is up; not when Ghostly Strike is under 1 s unless 75 energy or more).
- **subtlety.md:82, :90, :133.** Death Mark at four, Preparation's five combo points (Premeditation 2 + Boot Knife Shank 2 + Ghostly Strike 1), Garrote from 26 and Death Mark from 40 all hold.
- **Hemorrhage on Garrote** (:53 and the Garrote note "Hemorrhage, Opportunity and Death Mark all feed"): the 70755 rank text reads "15% increased Rupture and Garrote damage", and the server push 113483 sets the mask to Rupture + Garrote (round 16 S3). Holds.

## Shoulds (wording, not findings)

| id | file:line | note | suggested wording |
|---|---|---|---|
| ROGUE-R14-S1 | subtlety.md:80 | "every point Rupture does not need goes into the next mark, never into Eviscerate" is followed straight away by spending three or more on Slice and Dice. The point of the sentence (no Eviscerate) is right; the "every point" reads against the next sentence. | "Once you have Death Mark, Eviscerate is off your bar, even with Death Mark on cooldown: the points Rupture and Slice and Dice do not need go into the next mark." |
| ROGUE-R14-S2 | subtlety.md:90 | Carried from R13-S1: "Death Mark is your only damage cooldown" while Shank and Ghostly Strike are damage on cooldowns too. | "Death Mark is your burst, and it comes back every 10 seconds, so there is nothing to hold: ..." |
| ROGUE-R14-S3 | subtlety.md:117 | The Leatherworking reason names Devilsaur and Stormshroud, but the Subtlety band 60 raid list uses neither: its leatherworking crafts are Outlaw's Collar and Dawn Armor (both factions). Not false (all are bind on equip crafts), just pointing at the wrong pieces. | "**Skinning** and **Leatherworking** pay for themselves: Outlaw's Collar, Dawn Armor and the Devilsaur set are leatherworking crafts." |

## Checked and right (re-derived this round)

**Gear prose against the live BiS rows (band 60 raid, both factions).**
- **Assassination** ("Two daggers, always"): Electrified Dagger 19100 (Alliance) or Glacial Blade 19099 (Horde) in the main hand, Heartseeker 12783 in the off hand. All subclass 15 (dagger). Holds.
- **Combat** ("daggers at 60"): Electrified Dagger or Glacial Blade, off hand Sentinel's Blade 19546 or Scout's Blade 19542. All daggers. Holds. "when it is on your list" for Heartseeker is conditional and true (it is on the band 60 bare list, not the raid one).
- **Subtlety** ("dagger or not"): Teebu's Blazing Longsword 1728 (sword, subclass 7) and Persuader 22384. Persuader is a one-handed MACE (subclass 4), not a sword as round 13's and round 16's reports wrote; nothing on the page depends on it.
- **Hearts.** Molten and Frozen Heart of the Mountain are bind on pickup (`items/rogue.json`) and appear in every band 60 raid list except Horde Assassination and Horde Combat/Subtlety, which carry one Heart plus Rune of the Guard Captain. "Your best in slot table shows which of them your set wants" (assassination.md:115) holds on both factions.
- **Leatherworking crafts** (Stormshroud, Devilsaur, Outlaw's Collar, Dawn Armor) are all bind on equip, so "pay for themselves" (combat.md:116) holds.

**Stat prose against the live tables (band 60 raid scale factors, haste set aside).**
- Assassination: hit 1.99, agility 1, crit 0.82, strength 0.59, attack power 0.53 (both factions). Prose explains crit's Seal Fate and Lethality value; no ranking.
- Combat: hit 2.39, agility 1, strength 0.74, attack power 0.67, crit 0.52. Prose explains why crit is low (no Seal Fate, no bleed all fight).
- Subtlety: Alliance hit 1.05, agility 1, crit 0.83, strength 0.57, attack power 0.52; Horde hit 1.36, agility 1, crit 0.92, strength 0.52, attack power 0.48. "Crit is better than it looks" fits crit above strength and attack power.

**Races.** Every race line still matches `round3/race-sims.md:18-23`, and each page's `recommendedRaces` agrees with its prose: Assassination and Subtlety Night Elf and Undead, Combat Gnome and Undead; Orc before Troll on Horde in all three; Skyborne only as the paid Heroic Pack option. No newer measured rogue race result exists in the tree.

**Card notes** (all bands) against the prose, other than R14-1 and R14-2: Subtlety Shank at three or fewer, Ghostly Strike at four or fewer, Garrote after Ghostly Strike, Rupture at four, Death Mark at four; Assassination Slice and Dice at four, Cold Blood then Eviscerate, Shank with Venom; Combat Eviscerate at five, Slice and Dice at one or two, Backstab, Adrenaline Rush with racial and trinket. All agree.

**Spells and talents.** Re-read this round: Shank 1329283/1329549 (15/25 energy, 10 s, cannot miss), Garrote 11290 (50, stealthed and behind), Preparation 14185 (10 min, other Rogue abilities), Death Mark 1325758 (25, 10 s), Adrenaline Rush 13750 (+100%, 15 s, 5 min), Blade Flurry 13877 (25, 2 min, +20%, one extra target), Dirty Deeds, Camouflage, Elusiveness (90 s at 2/2), Initiative, Serrated Blades, Ghostly Strike, Hemorrhage, Boot Knife, Composure, Efficient Strikes, Opportunity, Blood Splatter. Everything else as round 13 listed.

**Sources.** All six frontmatter URLs returned 200 on 2026-11-04 (the two Blizzard articles, the Wowhead Skyborne article, foreversixty.gg /planner, /planner?class=rogue and /sim/weights).

**Voice.** No em or en dashes, no site, sim or APL talk, no quoted nightly numbers.

## Decisions

- **Bleeds "can crit", not "always crit".** The owner's ruling and the engine both make bleed ticks able to crit at your crit chance. Round 16's wording stands and the merge's revert is the error.
- **Death Mark at four everywhere.** The Eviscerate note must stop saying five.
- **Race lines stand.** None contradicts a measured result.
- **No build, order, `statPriority` or BiS change.**

## Sims requested

Carried (none new):
1. Subtlety race re-run, both factions, on the round 16 rotation (Garrote after Ghostly Strike, Death Mark at four: +4% raid 60, so the 2.9 and 1.4 DPS gaps in `round3/race-sims.md` may move).
2. Assassination race re-run with the Rupture line.
3. Data lane: re-rank `bis/rogue-combat.json` band 60 raid on the shipped `raidBuild` `005203002/3043030002012321451/005` (the rows were ranked on `005203002-3243030000012321451-005`). The stat order still matches the prose, so nothing on the page is false today.
4. Subtlety BiS and weights at the next refresh (round 16 "Measure at refresh": Garrote's share and Death Mark's frequency moved).

Round 13's item 1 (Subtlety shipped APL against the old Eviscerate line) is closed by round 16 S1/S2.

## Sim findings

None new. Carried, not re-filed:
- Round 16 open: K1 (what Death Mark stores), K2 (off-hand halving on Shank and Mutilate), S4 (Hemorrhage's dagger hand), C3 (Combat's Eviscerate threshold by weapon). The page makes no claim on any of them that the client contradicts.
- Component lane: talent-spell rows (Ghostly Strike, Garrote's Ghostly Strike note) showing on bands before the talent; the cut-off Hemorrhage note.

## Sources

- "World of Warcraft: Forever Class Deep Dives - Rogue and Warlock", Blizzard (200, 2026-11-04): https://news.blizzard.com/en-us/article/24310968/world-of-warcraft-forever-class-deep-dives-rogue-and-warlock
- "World of Warcraft: Forever Deep Dive Panel Recap", Blizzard (200): https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- "Skyborne First Look: New Neutral Race in World of Warcraft: Forever", Wowhead (200): https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- foreversixty.gg: /planner, /planner?class=rogue, /sim/weights (all 200)
- Client data in `data/builds/1.60.1.70755/`: `talents/rogue.json`, `spelltips/rogue.json` (1329283, 1329549, 11290, 14185, 1325758, 13750, 13877, 11300), `items.json` and `items/rogue.json` (1728, 22384, 12783, 19099, 19100, 19542, 19546, 249469, 249470, 252483, 279253, 15063), `bis/rogue-*.json`, `addon-data.json`
- APLs: `data/curated/apl/rogue-{assassination,combat,subtlety}.json`
- Engine f637ec32c: `sim/core/dot.go:30-38, 436-460` (periodicDamageCanCrit)
- Commits: 8160ad885 (round 16 rogue text), c4483a3a6 (merge that reverted it), 6e57a9c60 (addon-data regen)
- Memory: `forever-dots-crit.md`
- Previous rounds: `round13/rogue.md`, `round3/race-sims.md`, `2026-11-04-sim-audit-round16/rogue.md`, `ledger.md`

## Counts

CONFIRMED 2, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs: 3. New sim findings: 0.

Verdict: NOT zero for rogue. Both errors are round 16 fixes that the guide-round merge (c4483a3a6) undid.

## Fix log (round 14)

- ROGUE-R14-1: fixed, rogue-subtlety.json Eviscerate note: "from 40, those points go into Death Mark instead." Notes only. addon-data.json and Data.lua are not regenerated here; that is the integrator's job.
- ROGUE-R14-2: fixed, assassination.md:78 and the rogue-assassination.json Rupture note: "its ticks can crit in Forever".
- SHOULDs: not applied.

All files were already at updated: 2026-11-04. The prose and sections tests pass (145/145).

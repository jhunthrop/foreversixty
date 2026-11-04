# Guide accuracy round 4: Paladin

Auditor: paladin (class index, Holy, Protection, Retribution). Audit tree land/guides-r4 752c02e43, client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go). Read-only: no guide edits, no sims run.

Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, plus the rendered card notes in `data/curated/apl/paladin-*.json`. Before starting I read the round 1, 2 and 3 paladin reports, round4/paladin-levels.md, round2/race-sims.md, round3/race-sims.md and ledger.md.

## Builds decoded (talents[] order in talents/paladin.json; unchanged since round 3, all legal on 70755)

- **Retribution `build`** `050250002/502/0523330100130135` (14/7/30) and **`raidBuild`** `050050002/503/05233301001301351` (12/8/31): as listed in round 3.
- **Protection `build`** `05005/15302102513501031/5003` (10/33/8): as round 3.
- **Holy `build`** `50320215012101251/50322/0503` (31/12/8): as round 3.

## Leveling prose against the shipped point order (addon-builds.json, index 0 = level 10)

- **Retribution:** Benediction 10-14, Improved Judgement 15-16, Conviction 17-19, Sanctified Judgement 20-22, Vindication 23-25, Seal of Command 26, Sacred Arbiter 27, Two-Handed Weapon Specialization 28-30, Repentance 31, Vengeance 32-34, Champion of the Light 35-39, then Divine Strength 40-44, Reverence, Improved Seals, Holy Conduit, Divine Intellect and Precision. The new Levels 1 to 19 text at retribution.md:123 and the line at :74 both match.
- **Protection:** Divine Intellect 10, Redoubt 11-15, Precision 16-18, Sacred Duty 19, Templar's Bulwark 20, Sacred Duty 21, Shield Specialization 22-23, Toughness 24-28, Swift Judgement 29, Holy Shield 30, Shieldsworn 40. The text at protection.md:133 and :138 matches.
- **Holy:** Spiritual Focus 10-14, Healing Light 15-17, Reverence 18-19, Purifying Power 20-21, Divine Favor 22, Oath of Light 25 (the 16th point, the first one its row allows), Divine Precision 29, Infusion of Light 30-31, Holy Shock 32, Light's Vigil 40. The text at holy.md:116 and :121 matches.

## Levels 1 to 19 (new this round), checked line by line

Each claim was checked against trainables/paladin.json, spelltips/paladin.json and talents/paladin.json:

- **Trainer levels:** Devotion Aura, Holy Light and Seal of Righteousness at 1. Judgement and Blessing of Might at 4. Holy Strike, plus Holy Light rank 2, at 6. Hammer of Justice and Purify at 8. Lay on Hands, Righteous Fury and Seal of Fury at 10. Blessing of Wisdom and Holy Light rank 3 at 14.
- **Cooldowns:** Judgement and Holy Strike are both on 10 s cooldowns, and Judgement "Does not consume the Seal". Improved Judgement 2/2 brings Judgement to 8 s.
- **Seals:** Seal of Righteousness says "Slower weapons cause more Holy damage per swing". Seal of Fury says "While a shield is equipped, each attack also grants an absorb shield", and its Judgement deals Holy damage and taunts for 3 sec.
- **Talent text:**
  - Retribution: Benediction 10% on instants, Conviction 6%.
  - Holy: Spiritual Focus 70%, Healing Light 12%, Reverence 30% while casting.
  - Protection: Divine Intellect 2% at rank 1, Redoubt 30% block for 10 sec or 5 blocks, Precision 3%, Sacred Duty rank 1 30 sec off Divine Protection.
- **Gear and armor:** paladins cannot use wands. Mail until plate at 40 follows the owner's class rule.
- **Redemption at 12:** it has a trainables row at 12, and in Forever it is a level 12 class quest chain. Icy Veins' Forever paladin class quests page lists The Tome of Divinity for Alliance and A Lesson in Divinity through Tirisfal Glades for Undead.
- **Consistency:** none of the new subsections contradicts the Levels 20 to 60 text or the strip. Protection's merged sentence ("Righteous Fury and Seal of Fury are enough to hold...") reads correctly after the 1 to 19 bullet that names level 10.

## Round 3 fixes re-checked

- R3-1 (retribution.md:138): correct.
- R3-2 (retribution.md:72): now reads "keeps 30% of your mana regeneration running while you cast". The talent text is "Allows 30% of your Mana regeneration to continue while casting", so it is correct.
- R3-3 (holy.md:86), R3-4 (protection.md:113), R3-7 (retribution.md:66), R3-8 (protection.md:51) and R3-9 (holy.md:70): correct.
- R3-5 (Protection Holy Strike card note): correct.
- R3-6 (retribution.md:93): it now states facts only (Crusader's Judgement lasts 40 s and is refreshed by your melee; Judgement does not consume the Seal). It is correct, and it still waits on the sim.

## Checked and correct (no finding)

All other prose and card notes on the four pages were checked again. Examples:

- **Divine Shield card note:** "Twelve seconds immune" matches rank 2 (1020, 12 sec).
- **Index milestone talents:** the 11/16/21/31 rows hold for all three trees. Holy has Divine Favor, Oath of Light, Holy Shock and Light's Vigil. Protection has Templar's Bulwark, Swift Judgement, Holy Shield and Shieldsworn. Retribution has Seal of Command, Sacred Arbiter, Repentance and Unsealed Wrath.
- **Holy:** Concentration Aura is a trainer spell at 22 (35%).
- **Retribution:** Holy Conduit is 40% off the mana cost, the Two-Handed Weapon Specialization threat cut is 20%, and Champion of the Light is 60% of Intellect and 1.0 s off Hammer of Wrath's cast.
- **Light's Vigil:** about 10 top-rank Flash of Lights (1340 against 140).
- **Lay on Hands:** the drain wording matches the 10310 description.

There are no em or en dashes and no site or sim talk.

## Source links (all fetched 2026-11-03; every one returns 200)

| page | label | URL | says what the label says? |
|---|---|---|---|
| index | Blizzard, Deep Dive panel recap | news.blizzard.com/.../24303313 | yes (paladin section: Holy Strike, Seals not consumed, Seal of Fury, Undead Paladin) |
| index | Blizzard, What's Next panel recap | news.blizzard.com/.../24303862 | yes |
| index | Wowhead, Undead Paladins coming to Forever | wowhead.com/forever/news/...-382823 | yes (title matches) |
| index | Forever Sixty, which races can be paladins | foreversixty.gg/data/1.60.1.70755/combos.json | yes (real JSON, not an HTML fallback) |
| index | Power Up Gaming, Forever talent changes vs Classic | powerupgaming.co.uk/... | yes |
| ret/holy/prot | Forever Sixty, paladin talents | foreversixty.gg/data/1.60.1.70755/talents/paladin.json | yes (JSON, class_slug paladin, build 70755) |
| ret/holy/prot | Forever Sixty, races | foreversixty.gg/data/1.60.1.70755/races.json | yes |
| ret | Warcraft Tavern, Classic Retribution Paladin rotation | warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/ | yes ("PvE Retribution Paladin Rotations & Cooldowns - WoW Classic") |
| ret | Forever Sixty, what each stat is worth | foreversixty.gg/sim/weights | yes for Retribution |
| holy, prot | Forever Sixty, what each stat is worth | foreversixty.gg/sim/weights | **no**, see PALADIN-R4-2 |
| ret/holy/prot | Forever Sixty, raids by phase | foreversixty.gg/sources | **no**, see PALADIN-R4-1 |

Note: the Deep Dive recap says Holy Strike has a "12-second cooldown". The 70755 client says 10 s (spelltips 679/10333, trainables), and the guides follow the client, which is correct.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R4-1 | CONFIRMED | source link (round 4 edit) | retribution.md:26-28, holy.md:22-24, protection.md:36-38 | `label: 'Forever Sixty, raids by phase'` with `url: https://foreversixty.gg/sources` | /sources is web/src/content/pages/sources.md, "How sources are classified". It explains the source pills and lists panel recaps, and says nothing about raids or phases. Round 4's commit 752c02e43 re-pointed the dead data/curated/loot/forever-raid-phases.json link here, so this is a wrong round 4 edit. No site page lists raids by phase. Blizzard's What's Next recap does name the raids (Hyjal Summit 20-player, the Barrow Deeps 10-player) and their December 9 unlock. | Replace the entry on all three pages with `label: "Blizzard, What's Next panel recap"`, `url: https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap`, `kind: blizzard` (the raids and when they open). If the integrator prefers fewer sources, delete the entry instead, but never keep the /sources URL under a raid label. |
| PALADIN-R4-2 | CONFIRMED | source link (round 4 edit) | holy.md:19-21, protection.md:33-35 | `label: 'Forever Sixty, what each stat is worth'` with `url: https://foreversixty.gg/sim/weights` | /sim/weights is a damage-spec tool. The page copy (web/src/lib/sim/copy.ts:719, :880) reads "What one point of each stat is worth for your damage spec". isSimulatedSpec (web/src/lib/sim/spec-label.ts:96) admits only role `dps`, and /sim/specs lists healers and tanks under "Not simulated yet". A Holy or Protection paladin who follows the link gets no number for their spec. The link is correct on retribution.md. | holy.md and protection.md: delete this source entry. Their stat tables already render on the page. Keep it on retribution.md. |
| PALADIN-R4-3 | LIKELY | leveling (new 1 to 19 text) | retribution.md:124, holy.md:117 | "drop the starter weapon for the first slow two-hander you find" / "swap the starter for the first slow two-hander you find" | The paladin starter is already a two-hander: items.json 2361 Battleworn Hammer, class 2 subclass 5 (two-handed mace), inventory type 17, and it is already slow at 2.9 in Classic. The line reads as though you start with a one-hander. | Ret: "Take quest rewards over vendor gear, and replace your Battleworn Hammer with the first harder-hitting slow two-hander you find: Seal of Righteousness hits harder per swing the slower the weapon." Holy: "...and replace your Battleworn Hammer with the first better slow two-hander you find." |

## Decisions

- **Source fixes (R4-1, R4-2):** decided as written above, with no owner question. The raid-phase source becomes Blizzard's What's Next recap. Holy and Protection drop the weights link.
- **Retribution race:** Human stays (round 2 sims: raid tie, bare Human +1.2%).
- **Holy and Protection races:** the prose and frontmatter stay as they are until the sims below report. Holy is `[human, undead]`, Protection is `[dwarf, undead]` with a `dwarf` build string. The round 3 ledger row lists "Human (Ret, Holy), Dwarf (Protection)", but round3/race-sims.md has no paladin rows, so those two picks are not measured yet.

## Sims requested

1. **Protection paladin race, tank fight** (`spec-breakdown -tank`, guide build, published level 60 raid-ready set, curated boss, 10,000 iterations, seeds 1 and 50001). Run Dwarf, Human and Undead and report damage taken per second, TMI and threat per second. The engine never casts Stoneform, as the warrior round 3 caveat notes, so Dwarf's case on the page rests on Artisan's Precision plus an unmeasured button. If Human beats Dwarf beyond error on damage taken and TMI, the integrator should flip the build race and recommendedRaces to `[human, undead]` and swap the order of the first two race sentences.
2. **Holy paladin race, healer fight** (the same healer profile as the round 3 priest runs, guide build, 10,000 iterations, seeds 1 and 50001). Run Human, Dwarf and Undead and report effective HPS. Human stays unless Dwarf wins beyond error.
3. Carried, not re-filed: the round 3 Retribution Crusader opener sim (R3-6), and the Consecration leveling rank under the level-gap curve.

## Sim findings

None new. Seal of Fury's shield condition, Divine Shield's 12 s at rank 2, Lay on Hands' drain and Judgement of the Crusader's refresh all agree with the 70755 client text.

## Sources

- 70755 client data: talents/paladin.json, trainables/paladin.json, spelltips/paladin.json, items.json (2361 Battleworn Hammer), addon-builds.json, combos.json, races.json. Also data/curated/apl/paladin-*.json.
- Site code: web/src/content/pages/sources.md, web/src/lib/sim/copy.ts:700-720 and :878-880, web/src/lib/sim/spec-label.ts:96, web/scripts/sync-data.mjs:34-80 (talents, races and combos are published under /data/<build>/).
- Blizzard, World of Warcraft Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Blizzard, World of Warcraft Forever What's Next panel recap: https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap
- Icy Veins, WoW Forever Paladin Class Quests (Redemption at 12): https://www.icy-veins.com/wow-forever/paladin-class-quests
- Warcraft Tavern, Forever Paladin Leveling Guide: https://www.warcrafttavern.com/forever/guides/paladin-leveling/
- Wowhead, Undead Paladins Coming to World of Warcraft Forever: https://www.wowhead.com/forever/news/undead-paladins-coming-to-world-of-warcraft-forever-382823
- Power Up Gaming, New WoW Forever Talents: Changes vs Classic Guide: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Warcraft Tavern, PvE Retribution Paladin Rotations & Cooldowns (Classic): https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/

## Counts

CONFIRMED 2 (R4-1 on three pages, R4-2 on two pages; both are source links that round 4 re-pointed), LIKELY 1 (R4-3), NEEDS-SIM 0 separate (two race sims requested above), NEEDS-OWNER 0. Guide prose and card notes: 0 confirmed errors. The new Levels 1 to 19 text has 0 confirmed errors.

## Fix log (round 4)

- PALADIN-R4-1 fixed: the "Forever Sixty, raids by phase" entry (foreversixty.gg/sources) is replaced with "Blizzard, What's Next panel recap" (news.blizzard.com/en-gb/article/24303862, kind blizzard), which names the raids and their December 9 unlock: retribution.md:26-28, holy.md:19-21, protection.md:33-35.
- PALADIN-R4-2 fixed: "Forever Sixty, what each stat is worth" (/sim/weights) deleted from holy.md and protection.md (formerly holy.md:19-21, protection.md:33-35); kept on retribution.md:20-22.
- PALADIN-R4-3 fixed (verified: items.json 2361 Battleworn Hammer is a two-handed mace): retribution.md:124 "replace your Battleworn Hammer with the first harder-hitting slow two-hander you find"; holy.md:114 "replace your Battleworn Hammer with the first better slow two-hander you find".
- Race decisions (round4/race-sims.md), unchanged picks: Protection `[dwarf, undead]` with the dwarf build string, Holy `[human, undead]`, Retribution `[human, undead]`. Frontmatter already matched. Protection race prose (protection.md:118) now says Dwarf and Human take hits and hold threat about the same and Stoneform tips it; Holy and Retribution prose already matched.
- Sources check: no wago.tools links, no "Talents Forever" label, no /sources link left on any paladin page; no em or en dashes.
- Card notes: no change needed (none flagged).
- New SIM findings: none.

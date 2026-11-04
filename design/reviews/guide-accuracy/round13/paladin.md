# Guide accuracy round 13: Paladin

Auditor: paladin (class index, Holy, Protection, Retribution). Audit tree gr13/audit (main 604b78a80: twelve guide rounds, sim round 15, guide-header-build). Client 1.60.1.70755, engine pin 9f7f1f6e4 (sim/enginever/version.go; source read from gr10/engine). Read-only: no guide edits, no sims run.

Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, the card notes in `data/curated/apl/paladin-*.json`, `data/curated/rank-choices.json`, `data/builds/1.60.1.70755/addon-builds.json`. Before starting I read the round 12 report and fix log, the ledger, and the round 1 to 11 decisions it carries.

## What changed since round 12's audit (git diff a8b147a6f..604b78a80)

- **Guides:** only the R12-1 fix at retribution.md:133 (leveling Consecration ranks).
- **Card notes, rank-choices.json, addon-builds.json:** unchanged for paladin.

## Builds decoded (unchanged, all legal on 70755)

- **Retribution** `build` 050250002/502/0523330100130135 (14/7/30), `raidBuild` 050050002/503/05233301001301351 (12/8/31). The raid order adds Unsealed Wrath at 40 and takes Precision 3 for Reverence 2, as retribution.md:58 says.
- **Protection** `build` 05005/15302102513501031/5003 (10/33/8).
- **Holy** `build` 50320215012101251/50322/0503 (31/12/8).

Leveling orders (addon-builds.json, index 0 = level 10), decoded by name:
- Retribution: Benediction 10-14, Improved Judgement 15-16, Conviction 17-19, Sanctified Judgement 20-22, Vindication 23-25, Seal of Command 26, Sacred Arbiter 27, Two-Handed Weapon Specialization 28-30, Repentance 31, Vengeance 32-34, Champion of the Light 35-39, then Holy and Protection (raid: Unsealed Wrath 40). retribution.md:74 and :125 match.
- Protection: Divine Intellect 10, Redoubt 11-15, Precision 16-18, Sacred Duty 19, Templar's Bulwark 20, Sacred Duty 21, Shield Specialization 22-23, Toughness 24-28, Swift Judgement 29, Holy Shield 30, ... Shieldsworn 40. protection.md:133 and :139 (Holy Shield at 30, Bulwark before the second Sacred Duty) match.
- Holy: Spiritual Focus 10-14, Healing Light 15-17, Reverence 18-19, Illumination 20, Divine Favor 21, Illumination 22-24, Oath of Light 25, Illumination 26, Infusion of Light 27-28, Purifying Power 29-30, Divine Precision 31, Holy Shock 32, ... Light's Vigil 40. holy.md:117 and :123 match (Oath at 25 is the 16th point, the first one row 4 allows).

## Round 12 edit re-checked

- **retribution.md:133 (R12-1 fix):** holds. Rank 1 Consecration (learned at 20) keeps its full spell power share to 37 and loses 5% a level after (engine core.LevelPenalty; ForeverChanges' formula `1 - (L - 17 - T) x 0.05`), so at 46 to 48 it has lost 45% to 55%: "about half" is right, and "until your mid 40s" matches round 12's crossover of about 46. "About twice as hard for the same global" at 40 matches 197 against 100 per target (round 12). The opening clause "carries as much spell damage as any rank" is only exact to 37; the next sentences qualify it, so it is a SHOULD, not a finding. The damage-spell half of the level rule is the source's own assumption (it measured heals); Blizzard confirmed the effect without numbers. Kept as a decision, not a finding.

## Checked and correct (no finding)

- Every talent bullet on all three pages against talents/paladin.json at the page's rank (Spiritual Focus 70%, Healing Light 12%, Reverence 30%, Illumination, Divine Favor, Infusion of Light 1.0 s within 15 s, Boundless Light 30%, Holy Power (+50% of the bonus, 175%), Light's Vigil, Divine Precision 4%/10% at 1 point, Redoubt 30% for 10 s or 5 blocks, Precision, Sacred Duty 60 s/10 s (30 s at rank 1, protection.md:133), Templar's Bulwark, Shield Specialization 4% once every 3 s, Swift Judgement, Improved Righteous Fury 9%, Anticipation 20, Holy Shield 30%/4 charges/20% threat, Iron Creed, Shieldsworn, Benediction, Improved Judgement, Conviction, Sanctified Judgement 60%, Sacred Arbiter, Two-Handed Weapon Specialization 6%/20%, Repentance 6 s, Vengeance non-periodic, Champion of the Light, Unsealed Wrath 15 s/3 stacks).
- Trainer levels in every "Levels 1 to 19" paragraph (trainables/paladin.json): Seal of Righteousness, Holy Light and Devotion Aura 1; Judgement and Blessing of Might 4; Holy Strike 6; Hammer of Justice and Purify 8; Lay on Hands, Righteous Fury and Seal of Fury 10; Redemption 12; Holy Light ranks at 6 and 14; Blessing of Wisdom 14; Flash of Light 20. Index: Consecration and Blessing of Kings 20, Hammer of the Righteous 40, Holy Strike 10 s cooldown.
- Spell values (spelltips/spellconst 70755): Judgement 10 yd, off the global cooldown, 6% base mana; Swift Judgement free, off the global, 1 min; Seal of Fury taunt 3 s and shield absorb; Judgement of the Crusader 40 s, refreshed by melee; Holy Shield 10 s cooldown equal to its duration, ranks 30/40/50/60; Hammer of the Righteous target plus 3; Holy Strike rank 8 20 mana; Hammer of Wrath 1.0 s base cast, 6 s, 30 yd; Exorcism 15 s, Undead and Demon only; Holy Wrath 2 s cast; Lay on Hands 20 min, maximum health, 550 mana at rank 3, drains your mana; Righteous Fury 30 min, 30% base mana; Divine Shield 12 s; Light's Vigil 1,340 mana at 60; Voice of Truth 6 s, 3 min; Blessing of Protection 5 min less Guardian's Favor's 2; Blessings and Greater Blessings 60 min; Holy Shock rank 1 at 30 for 160 mana.
- Racials (spelltips/shared.json): Ruthless, The Human Spirit, Will to Survive, Stoneform 10% for 8 s / 3 min, Artisan's Precision, Big Game Hunter, Will of the Forsaken (no immunity), Cannibalize (health and mana), Touch of the Grave. Heart of the Lion is a hunter trainer spell (409580, 10% all stats). No Blessing of Light, Salvation or Sanctuary in trainables.
- Consecration ticks crit (memory forever-dots-crit names Consecration), so retribution.md:62 holds.
- rank-choices.json: Retribution Consecration rank 1 below 60, rank 4 bare at 60 (retribution.md:95); Holy Flash of Light rank 1 rows only below 60 (holy.md:78 and :127 say "at 60").
- Card notes: every paladin note re-read; nothing false on the rendered lines. The Protection top entry's note "Lay on Hands fires itself below 10% health" is on an `autocastOtherCooldowns` line, which rotation-lists.ts never renders and no shipped data file carries, so it is not a page statement (listed for the data lane).
- Source links: all nine distinct URLs returned 200 with the right content type (fetched 2026-11-04); labels match their pages.
- Voice: no em or en dashes, no site or sim talk, no stat ranking in the Stat priority sections, no nightly numbers.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R13-1 | CONFIRMED (contradiction) | talent | retribution.md:60 | "Hammer of Wrath is the exception: Champion of the Light makes it instant, but it still costs full price." | Both Retribution builds take Holy Conduit 2/2 (decoded above), and its 70755 text (1237268) reads "Reduces the mana cost of your Consecration, Holy Wrath, Exorcism, and Hammer of Wrath spells by 40%." The engine applies it (sim/paladin/paladin.go `PaladinSpellMaskHolyConduitCost` includes Hammer of Wrath). The page itself says so at :72 ("Holy Conduit takes 40% off Consecration, Exorcism and Hammer of Wrath"). Hammer of Wrath costs 255 of its 425 mana in this build, not full price. What is true is that Benediction does not discount it (hammer_of_wrath.go carries no benediction multiplier). | "**Benediction** (5 points): your Seals, Judgement, Holy Strike, Consecration and Exorcism cost 10% less. Hammer of Wrath is the exception: Champion of the Light makes it instant, but Benediction still does not discount it." |

## SHOULDs (wording, not findings)

- retribution.md:133: "Rank 1 at 135 mana carries as much spell damage as any rank" is exact only to level 37. Suggest "keeps almost all of its spell damage through your early 40s for a fraction of the cost".
- Carried: the Retribution Holy Strike card note's "Twenty mana" against Benediction's 18.
- Carried: retribution.md:111 "Human is ahead solo" (default target; Beasts unmeasured).

## For the component and data lanes (not guide findings)

- data/curated/apl/paladin-protection.json priorityList[0] note: "Lay on Hands fires itself below 10% health" describes the engine's autocast default (lay_on_hands.go:69, marked TODO), not the game. It does not render today; reword it if that line ever surfaces ("Cast it below 10% health on yourself...").
- Carried: the Holy APL engineering note ("no downrank penalty") and the Protection APL note ("+60 percent threat on Holy damage", where 25780 says 33% on all threat and a further 90% on Holy). Neither renders.
- Carried: trainables lists Holy Light 9, Blessing of Might 7 and Blessing of Wisdom 6 at 60 as skill_line_ability rows (AQ librams in 1.12).

## Decisions

- **R13-1:** decided from the Holy Conduit client text, the decoded builds and the page's own :72. No owner question.
- **Consecration level rule on damage spells:** kept. ForeverChanges measured the rule on heals and assumes it for damage; Blizzard confirmed the effect on October 1. No Forever source contradicts it, so the round 12 text stands.
- **Races:** unchanged. Retribution Human (Undead), Holy Human (Undead), Protection Dwarf (Undead). No paid race named, and no race line contradicts a measured result or the client.

## Sims requested

1. Carried (confirming R12-1, not needed): Retribution leveling Consecration on 3 targets, 20 s pack fight, bare gear from bis/paladin-retribution.json, rank 1 against the newest rank learned at levels 40, 45 and 50; damage per mana and per cast.
2. Carried: the Crusader opener on a 15 s single mob at bands 20, 30 and 40 (retribution.md:131).
3. Carried: Retribution Human against Dwarf on a Beast target (retribution.md:111); Protection threat by ability on 1 and 4 targets.
4. Race re-runs and the BiS re-rank on the shipped builds stay with the sim lane.
5. PROVISIONAL (BiS re-rank): none. No prose names a BiS piece or ranks a stat.

## Sim findings

None new.

## Sources

- 70755 client data, data/builds/1.60.1.70755: talents/paladin.json (1237268 Holy Conduit, every talent text above), trainables/paladin.json, spelltips/paladin.json, spellconst/paladin.json and spellconst/hunter.json (409580), spelltips/shared.json, addon-builds.json, trainables/hunter.json.
- data/curated/apl/paladin-{holy,protection,retribution}.json; data/curated/rank-choices.json; web/src/lib/guides/rotation-lists.ts (only castSpell lines render).
- Engine 9f7f1f6e4: sim/paladin/paladin.go (Holy Conduit cost mask), hammer_of_wrath.go, talents.go (benediction()), consecration.go, lay_on_hands.go:69, sim/core/spell.go:822 (LevelPenalty).
- Memory: forever-dots-crit.md (Consecration crits).
- ForeverChanges, Downranking in WoW Forever: https://foreverchanges.pro/downranking (formula, measured on heals 18 to 30, assumed for damage).
- Links fetched 2026-11-04, all 200: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap ; https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap ; https://www.wowhead.com/forever/news/undead-paladins-coming-to-world-of-warcraft-forever-382823 ; https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/ ; https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/ ; https://foreversixty.gg/planner?class=paladin ; https://foreversixty.gg/sim/weights ; https://foreversixty.gg/data/1.60.1.70755/races.json ; https://foreversixty.gg/data/1.60.1.70755/combos.json

## Counts

| status | count |
|---|---|
| CONFIRMED | 1 (R13-1) |
| LIKELY | 0 |
| NEEDS-SIM | 0 new (3 carried) |
| NEEDS-OWNER | 0 |

Round 12 edit: 0 errors. Levels 1 to 19: 0 errors. Source links: 0 errors. Card notes: 0 rendered errors.

## Fix log (round 13)

- PALADIN-R13-1 fixed (retribution.md:60): "Benediction still does not discount it" (Holy Conduit 2/2 still takes 40% off, as :72 says).

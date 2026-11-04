# Guide accuracy round 11: Paladin

Auditor: paladin (class index, Holy, Protection, Retribution). Audit tree gr11/audit (main d54de48aa: guide rounds 1 to 10, sim round 15, leveling-order). Client 1.60.1.70755, engine pin 9f7f1f6e4 (sim/enginever/version.go; source read from gr10/engine). Read-only: no guide edits, no sims run.

Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, the card notes in `data/curated/apl/paladin-*.json`, `data/curated/rank-choices.json`, `data/builds/1.60.1.70755/addon-builds.json`. Before starting I read the paladin reports for rounds 1 to 10 with their fix logs, round4/paladin-levels.md, the race-sims files, ledger.md and the sim round 15 paladin report.

## What changed since round 10's audit (git diff 84460cec8..HEAD)

- **Guides:** only the round 10 fixes (eabba8f84): retribution.md:78, :89, :93, :95, :111, :131, :133; holy.md:33, :78, :95, :127; protection.md:107, :111, :119.
- **Card notes:** the Retribution primary-Seal line (priorityList 2) and the Protection Stoneform line (priorityList 4) now have notes.
- **addon-builds.json and ladder-order.json:** unchanged for paladin. The leveling-order lane (ec60e6abf) changed only how the strip reads the order.
- **Rotation gates (b53e82298):** the Retribution opener now renders Seal of the Crusader alone, so round 10's component item 1 is closed.

## Builds decoded (unchanged, all legal on 70755)

- **Retribution** `build` 050250002/502/0523330100130135 (14/7/30), `raidBuild` 050050002/503/05233301001301351 (12/8/31). The raid build moves Reverence's 2 points into Precision 3 and Unsealed Wrath. retribution.md:58 matches.
- **Protection** `build` 05005/15302102513501031/5003 (10/33/8): Divine Strength 5, Improved Seals 5; Divine Intellect 1, Redoubt 5, Precision 3, Sacred Duty 2, Templar's Bulwark 1, Shield Specialization 2, Toughness 5, Swift Judgement 1, Improved Righteous Fury 3, Anticipation 5, Holy Shield 1, Iron Creed 3, Shieldsworn 1; Deflection 5, Conviction 3. protection.md:66 to :82 match.
- **Holy** `build` 50320215012101251/50322/0503 (31/12/8). Rows 1 to 5 of Holy hold exactly 25 points, so Holy Power's row opens. holy.md:51 to :67 match.

## Point orders (addon-builds.json, index 0 = level 10), re-derived

- **Retribution:** Benediction 10 to 14, Improved Judgement 15 and 16, Conviction 17 to 19, Sanctified Judgement 20 to 22, Vindication 23 to 25, Seal of Command 26, Sacred Arbiter 27, Two-Handed Weapon Specialization 28 to 30, Repentance 31, Vengeance 32 to 34, Champion of the Light 35 to 39. The raid order then takes Unsealed Wrath at 40. retribution.md:74 and :125 match.
- **Protection:** Divine Intellect 10, Redoubt 11 to 15, Precision 16 to 18, Sacred Duty 19, Templar's Bulwark 20, Sacred Duty 21, Holy Shield 30, Shieldsworn 40. protection.md:133 and :139 match.
- **Holy:** Spiritual Focus 10 to 14, Healing Light 15 to 17, Reverence 18 and 19, Oath of Light 25 (the first point in row 4: 15 Holy points by 24), Infusion of Light 27 and 28, Divine Precision 31. holy.md:117 and :123 match.

## Round 10 edits re-checked

- **retribution.md:78 (Crusader opener):** correct. Holy Strike is Holy school (engine holy_strike.go:32 and :141, client school_mask 2). Seal of Righteousness and Seal of Command procs are Holy, and so are Judgements and Consecration ticks, so "adds Holy damage to every Seal proc, Judgement, Holy Strike and Consecration tick" holds against 20308's "increasing Holy damage taken by up to 161". The client also says "Only one Judgement per Paladin can be active at any one time". Judgement of Righteousness and Judgement of Command are instant hits with no debuff, so judging your damage Seal leaves the Crusader debuff in place.
- **retribution.md:89, :93, :131:** consistent with :78. :131's single-mob questing claim is still unmeasured (carried sim 2).
- **retribution.md:95 (single target):** correct. rank-choices.json keeps rank 1 from 30 to 59, rank 4 at bare 60 and the authored rank 5 in the raid. The round 15 measurements back it (W3-4: +1.3% at 40, +3.0% at 50). Rank 1 costs 135 mana (26573).
- **retribution.md:133 (packs):** wrong. See PALADIN-R11-1.
- **retribution.md:111 (race):** "Dwarf matches Human on the raid build in a raid" holds: 783.2 against 780.5, error 0.52 (round5/race-sims.md). "Human is ahead solo" holds on the sim's default target: 283.1 against 280.0, error 0.15. Wording SHOULD below.
- **holy.md:33, :78, :95, :127:** correct. Flash of Light rank 5 (19942) is learned at 50 and heals 223 for 115 mana, against rank 6's 303 for 140. A gap of 10 levels keeps the full healing bonus. The cards are rank 5 on the tank (priorityList 11) and rank 4 on everyone else (12 to 16). Holy Shock's card is rank 4 (20930).
- **protection.md:107, :111, :119:** correct. Stoneform is 10% physical damage taken for 8 s on a 3 minute cooldown, and it costs a global (races.json; engine core/racials.go Stoneform, GCDDefault). The Protection 4P bonus of Tier 1 gives expertise (round 10).
- **New card notes:** paladin-retribution.json priorityList 2 and paladin-protection.json priorityList 4 are both accurate.

## Checked and correct (no finding)

Every line of the four pages and every card note, checked against spelltips/paladin.json, trainables/paladin.json, talents/paladin.json, races.json, combos.json and rank-choices.json:

- **Levels 1 to 19, all three pages:** every trainer level is right. Seal of Righteousness, Holy Light and Devotion Aura come at 1. Judgement and Blessing of Might come at 4. Holy Strike, Holy Light rank 2 and Divine Protection come at 6. Hammer of Justice and Purify come at 8. Lay on Hands, Righteous Fury and Seal of Fury come at 10, Redemption at 12, and Holy Light rank 3 and Blessing of Wisdom at 14. Seal of Fury's Judgement taunts for 3 s (20423). Spiritual Focus 5/5 gives 70%, Divine Intellect rank 1 gives 2%, and Sacred Duty rank 1 takes 30 s off.
- **Trainer levels in the 20 to 60 text:** Consecration, Exorcism, Flash of Light and Blessing of Kings come at 20. Holy Shield comes at 30, 40, 50 and 60 (1239570 at 60). Hammer of the Righteous comes at 40, Holy Wrath at 50, and Holy Shock at 30, 40, 48 and 56.
- **Numbers:** Light's Vigil rank 3 costs 1,340 mana and has a 1.5 s cast, so Benediction does not reduce it. Lay on Hands rank 3 restores 550 mana. Hammer of Wrath rank 3 has a 1.0 s base cast, which Champion of the Light makes instant, a 6 s cooldown and 30 yd range. Judgement and Swift Judgement reach 10 yd, and Judgement costs 6% of base mana. Righteous Fury costs 30% of base mana and lasts 30 min. Divine Shield rank 2 lasts 12 s. Holy Shield has 4 charges, and its 10 s cooldown matches its duration.
- **Talents:** Holy Power doubles the heal crit bonus from 50% to 75%, so a crit heals for 175% (holy.md:62). Divine Precision's single point gives 4% Holy spell hit and 10% shorter silences (holy.md:123). The paladin's 1-point milestones sit at 11, 16, 21 and 31 points in every tree (index.md:35).
- **Races:** paladins are Human, Dwarf or Undead (combos.json), and every racial line matches races.json. No paid race is named.
- **Source links:** all nine distinct URLs returned 200 with the right content types (fetched 2026-11-04), and every label matches its page. The source-link rule holds on all four pages.
- **Voice and rules:** no em or en dashes, no site or sim talk, no stat ranking in any Stat priority section, no nightly numbers. No BiS piece is named in the prose.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R11-1 | CONFIRMED (round 10 edit) | leveling / rotation | retribution.md:133 | "From level 20 Consecration is your AoE: two or three mobs on you is the time for it, and rank 1 at 135 mana is the rank to keep on your bar until 60: a rank loses its spell damage bonus once you are far above the level you learned it, but rank 1 has so little to lose that its low cost wins." | Round 10 measured only single targets on boss-length fights. R10-3 was LIKELY, and round 10's own sim request 1 asked for the pack case. This sentence is about packs, and on a pack the client numbers decide it without a sim. Per target, each cast deals (engine consecration.go, client companion rows 1280345 to 1280349, 8 ticks on the first 4 enemies): rank 1 48 for 135 mana, rank 2 (30) 80 for 235, rank 3 (40) 136 for 320, rank 4 (50) 224 for 435. That is 0.36, 0.34, 0.43 and 0.51 base damage per mana. The spell power share (0.042 a tick) is the same for every rank before the level-gap rule. The rule then takes 15% of it off rank 1 at 40 and 65% at 50, and nothing off the newest rank. So from 40 the newest rank beats rank 1 on damage per mana (+20% at 40, +45% at 50) and on damage per global (2.8x and 4.7x) on every mob it hits. On a pull you start with the mana for it, so rank 1 loses both ways. Rank 1 wins single target only because the pacing rule lets the cheap rank go out more often, as in :95. Through the 30s, ranks 1 and 2 deal about the same damage per mana. The text before round 10 ("from 40 move up to the newest rank") was right for packs. | "From level 20 Consecration is your AoE: two or three mobs on you is the time for it. Through your 30s, rank 1 at 135 mana does about as much for its mana as rank 2, so keep it. From 40, on a pack, cast the newest rank you have: each new rank hits every mob far harder for its mana, and rank 1 also loses its spell damage bonus as you outlevel it. On a single target, rank 1 stays the rank to use until 60, for the reason in the Consecration paragraph of the rotation." |

## SHOULDs (wording, not findings)

- retribution.md:111: "Human is ahead solo". This was measured on the sim's default target. Against Beasts, Big Game Hunter's 5% should put Dwarf ahead, but that is unmeasured (carried sim). The round 5 wording, "Human is ahead solo against everything but Beasts", is safer. It also fits the sentence that follows.
- Carried: the Holy Strike card's base cost of 20 against Benediction's 18 on Retribution; holy.md:117 naming Flash of Light at 10 to 14, although Flash of Light is learned at 20.

## For the component and data lanes (not guide findings)

- Round 10's item 1 (three ungated Seals in the Retribution opener) is fixed by b53e82298.
- Carried: the Holy APL's top-level engineering note still names rank 1 Holy Shock and rank 1 Flash of Light and says the engine has "no downrank penalty". It does not render.
- trainables/paladin.json lists Holy Light 9 (25292), Blessing of Might 7 (25291) and Blessing of Wisdom 6 (25290) at 60 as skill_line_ability rows. In 1.12 these are Ahn'Qiraj librams. The guide's "the rank 9 book has no known source in Forever yet" (holy.md:81, round 6) still holds, because no Forever source for them is known. If the data lane ever marks them trainable, holy.md:76 and :81 and the Holy Light card note must change.

## Decisions

- **R11-1:** decided from the client damage rows and the level-gap rule, with no owner question and no sim needed. It reverses a round 10 edit that took a single-target measurement into a pack sentence.
- **Races:** unchanged. Retribution Human (Undead), Holy Human (Undead), Protection Dwarf (Undead). No paid race named. The new Retribution APL (Crusader opener) has not been race-measured, and no race line contradicts a measured result.

## Sims requested

1. Confirm R11-1 (not needed to decide it): Retribution leveling Consecration on 3 targets, 20 s pack fight, bare, rank 1 against the newest rank learned at bands 40 and 50.
2. Carried: the Crusader opener on a 15 s single mob at bands 20, 30 and 40 (retribution.md:131).
3. Carried: Retribution Human against Dwarf on a Beast target (retribution.md:111 SHOULD); Protection threat by ability on 1 and 4 targets.
4. Race re-runs for Retribution on the round 15 APL (Crusader opener), raid 60 and bare 60.
5. PROVISIONAL (final BiS re-rank): none. The bis/paladin-*.json files in this tree are still the 03:30 to 03:50 UTC run on d1d9f2775, and no prose names a BiS piece or ranks a stat.

## Sim findings

None new. W3-1, W6-4 and W6-1 from sim round 15 agree with the client.

## Sources

- 70755 client data, data/builds/1.60.1.70755: spelltips/paladin.json (26573, 20116, 20922, 20923, 20924, 20308, 21082, 20271, 1310994, 10333, 20423, 19941, 19942, 19943, 1311595, 10310, 24239, 1239570, 20216, 20066, 642, 1020, 5573, 25780), trainables/paladin.json, talents/paladin.json, races.json, combos.json, addon-builds.json, bis/paladin-*.json.
- data/curated/apl/paladin-{holy,protection,retribution}.json; data/curated/rank-choices.json.
- Engine 9f7f1f6e4: sim/paladin/consecration.go (rank table, 0.042 coefficient), sim/paladin/holy_strike.go:32, :141, sim/core/racials.go (Dwarf, Stoneform GCD), sim/core/debuffs.go:298-330.
- design/reviews/2026-11-04-sim-audit-round15/paladin.md (W3-1, W3-4, W6-4); design/reviews/guide-accuracy/round5/race-sims.md, round4/race-sims.md.
- Commits eabba8f84, ec60e6abf, b53e82298.
- Links fetched 2026-11-04, all 200: Blizzard Deep Dive recap https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap ; Blizzard What's Next recap https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap ; Wowhead, Undead Paladins https://www.wowhead.com/forever/news/undead-paladins-coming-to-world-of-warcraft-forever-382823 ; Power Up Gaming https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/ ; Warcraft Tavern Retribution rotation https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/ ; foreversixty.gg planner, /sim/weights, races.json, combos.json.

## Counts

| status | count |
|---|---|
| CONFIRMED | 1 (R11-1, a round 10 edit) |
| LIKELY | 0 |
| NEEDS-SIM | 0 new (1 confirming, 4 carried) |
| NEEDS-OWNER | 0 |

Levels 1 to 19: 0 errors. Source links: 0 errors. Card notes: 0 errors. Round 10 edits: 1 error (retribution.md:133).

## Fix log (round 11)

- PALADIN-R11-1: fixed, retribution.md:133 with the report's text. Checked against spelltips/paladin.json: ranks 1 to 4 cost 135/235/320/435 and deal 16+32, 24+56, 48+88, 64+160 per target (0.36, 0.34, 0.43, 0.51 per mana); trainer levels 20/30/40/50/60.
- SHOULD holy.md:117: applied, "Holy Light, or Flash of Light once you learn it at 20" (trainables: Flash of Light 20).
- SHOULD retribution.md:111 ("against everything but Beasts"): left. It would imply Dwarf is ahead on Beasts, which is unmeasured (carried sim 3).
- Holy Strike card base-cost SHOULD: left, carried.

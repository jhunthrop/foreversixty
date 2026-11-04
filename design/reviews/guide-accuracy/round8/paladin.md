# Guide accuracy round 8: Paladin

Auditor: paladin (class index, Holy, Protection, Retribution). Audit tree gr8/audit (land/guides-r7, head 09b98d9e4). Client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go). Read-only: no guide edits, no sims run.

Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, plus the rendered card notes in `data/curated/apl/paladin-*.json`. Before starting I read the paladin reports from rounds 1 to 7 (with every fix log), round4/paladin-levels.md, the race-sims files for rounds 2 to 5, and ledger.md.

What changed since round 7's audit (git diff 394562f20..HEAD): only the round 7 fixes (retribution.md:70, :103, :111; holy.md:40, :81; protection.md:94; the Hammer of the Righteous card note in paladin-protection.json:251) and index.md's `updated:` stamp (2026-11-02 to 2026-11-03; the class page's content did not change). Builds, addon-builds.json, BiS files, rank-choices.json and the other APL lines are unchanged.

## Builds decoded (unchanged since round 1, all legal on 70755)

- **Retribution `build`** `050250002/502/0523330100130135` (14/7/30) and **`raidBuild`** `050050002/503/05233301001301351` (12/8/31). retribution.md:58 matches.
- **Protection `build`** `05005/15302102513501031/5003` (10/33/8). protection.md:65-81 matches.
- **Holy `build`** `50320215012101251/50322/0503` (31/12/8). holy.md:51-67 matches.
- Prerequisites as in round 7 (Unsealed Wrath needs Repentance, Vengeance needs Seal of Command, Swift Judgement needs Templar's Bulwark, Shieldsworn needs Holy Shield, Shield Specialization needs Redoubt, Illumination needs Reverence, Boundless Light and Light's Vigil need Holy Shock).

## Shipped point orders (addon-builds.json, index 0 = level 10)

Unchanged file. Retribution (Benediction 10-14, Improved Judgement 15-16, Conviction 17-19, ... Champion of the Light 35-39), Protection (Divine Intellect 10, Redoubt 11-15, Precision 16-18, Sacred Duty 19, Templar's Bulwark 20, Sacred Duty 21, ... Holy Shield 30, Shieldsworn 40) and Holy (Spiritual Focus 10-14, Healing Light 15-17, Reverence 18-19, ... Oath of Light 25, Divine Precision 29, Infusion of Light 30-31, Holy Shock 32, Light's Vigil 40). retribution.md:74, :125; protection.md:132, :138; holy.md:117, :123 all match. The 1 to 19 text and the 20 to 60 text still read right in page order.

## Round 7 fixes re-checked

- **R7-1** (retribution.md:103): correct. Holy Strike 10333 is "50% weapon damage plus an additional 93 as Holy damage" (weapon damage carries attack power), Seal of Command 20920 is "70% of normal weapon damage", Seal of Righteousness 20293 has no attack power term and "Slower weapons cause more Holy damage per swing".
- **R7-2** (retribution.md:111): correct. The site's request builder defaults the target to Humanoid (sim/request/request.go:539; sim/api/envelope.go:280), so "Human is ahead solo against Humanoids" is what round5/race-sims.md measured. Big Game Hunter: 5% to Beasts (races.json). The page makes no claim about Beasts beyond that.
- **R7-3** (holy.md:81): correct. Holy Shock rank 1 1311606 is a level 30 rank; 30 levels behind at 60 it keeps about 35% of your healing power.
- **R7-4** (protection.md:94 and the card note): correct. 407632: "Hammer the current target and up to 3 additional nearby targets, causing Holy damage"; Righteous Fury 25780 adds 90% more threat on Holy attacks.
- **R7-5** (retribution.md:70): Repentance half correct (Unsealed Wrath's prereq_talent_id is 105694, Repentance 20066 is 6 s on a 1 min cooldown, Humanoids only). The Vindication half has a wording slip, R8-1.
- **R7-6** (holy.md:40): correct (1340 against 10 × 140).

## Checked and correct (no finding)

Every line of the four pages and every rendered card note was checked again against talents/paladin.json, trainables/paladin.json, spelltips/paladin.json, races.json, combos.json, items.json and the APLs. New this round:

- **Seal of Fury** 1311649/20423: "causing melee attacks to deal an additional 35 Holy damage. While a shield is equipped, each attack also grants an absorb shield equal to 50% of the Holy damage dealt", Judgement 153 Holy and a 3 s taunt. So protection.md:79 "Improved Seals adds 15% to Seal of Fury and its Judgement" and :128 hold.
- **Judgement** 20271 and **Swift Judgement** 1310994: 10 yd, 6% base mana and free, 10 s and 60 s, "Does not consume the Seal".
- **Light's Vigil** rank 3 1311595: 1340 mana, 1.5 s, heals the target's party for 704 through the next Holy Shock; 30 s.
- **Holy Shock** range 20 yd; **Hammer of Wrath** and **Exorcism** 30 yd; **Holy Wrath** 20 yd, Undead and Demons, 60 s.
- **Divine Protection** 5573 (8 s, no attacking) and **Divine Shield** 1020 (12 s, half damage dealt), both Forbearance; Templar's Bulwark 1311015 the same lockout. protection.md:55 and :106 hold.
- **Blessings and Greater Blessings**: 60 min (index.md:33). Blessing of Kings is a trainer spell at 20 (index.md:31). Heart of the Lion 409580 is a hunter trainer spell at level 1 in Forever (trainables/hunter.json), so the Kings lines on all three spec pages hold.
- **Benediction mask** (engine benediction_mask_test.go, read from the client's SpellClassMask): Seals, Judgement, Consecration, Exorcism, Holy Strike, Holy Shock and Holy Shield are discounted, so holy.md:64 and retribution.md:125 hold. Hammer of Wrath is not (see R8-2).
- **Exorcism never crits** (retribution.md:62): still the engine's rule at the pin (exorcism.go, OutcomeMagicHit, beta-log comment); no newer Forever source says otherwise.
- **Elixir of Holy Power** 21546: in items.json, required level 40 (retribution.md:115, protection.md:122).
- **Races**: combos.json class 2 is Human, Dwarf and Undead only; neither Skyborne race (95, 96) has a paladin row, so index.md:39 holds after the Skyborne landing (715359838).
- Stat paragraphs: no ranking words, no numbers, every stated reason true on 70755. No gear piece named on any page.
- No em or en dashes, no site or sim talk in the prose.

## Source links (fetched 2026-11-04)

All nine distinct URLs return 200 (both Blizzard recaps, the Wowhead Undead paladin article, Power Up Gaming, the Warcraft Tavern Classic rotation page, the paladin planner, /sim/weights on Retribution only, combos.json and races.json as application/json). Labels match the pages. The source-link rule (ledger, round 4) holds on all four pages.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R8-1 | LIKELY | talent (round 7 edit, wording) | retribution.md:70 | "**Vindication** is filler on the way to Champion of the Light, since nothing else in that row adds damage." | Vindication sits on Retribution tier 2 (talents/paladin.json) with Unyielding Faith, **Seal of Command** and Pursuit of Justice. Seal of Command is in that row and adds damage: the same page makes it your main raid Seal at :78 and :91. Read literally the sentence contradicts the page; what it means is that the row's other open talents, Unyielding Faith and Pursuit of Justice, add none. | "**Vindication** is filler on the way to Champion of the Light: once Seal of Command has its point, the rest of that row is Unyielding Faith and Pursuit of Justice, and neither adds damage." |
| PALADIN-R8-2 | LIKELY | talent | retribution.md:60 (read with :67) | "**Benediction** (5 points): every instant costs 10% less, and nearly everything you press is instant." | :67 says Champion of the Light makes Hammer of Wrath instant, so the page tells you Hammer of Wrath costs 10% less. Its client data says it does not. Benediction 20101's SpellClassMask (word 1 0x26018242) leaves out Hammer of Wrath (word 1 0x80), and the engine reads it the same way (sim/paladin/retribution/benediction_mask_test.go at d1d9f2775: Hammer of Wrath 100%, the Seals, Judgement, Holy Strike, Consecration and Exorcism 90%). The client text "all instant cast spells and abilities" means spells that are instant by base, and Hammer of Wrath has a 1.0 s base cast (trainables 24275/24274/24239). It matters in the execute, where :97 already tells you to save a mana potion for it. | "**Benediction** (5 points): your Seals, Judgement, Holy Strike, Consecration and Exorcism cost 10% less. Hammer of Wrath is the exception: Champion of the Light makes it instant, but it still costs full price." |

No CONFIRMED findings.

## SHOULDs (wording, not findings)

- paladin-retribution.json Holy Strike note, "Twenty mana for a Holy weapon strike": with Benediction 5/5 it is 18, while the same card's twist note gives the Seals' cost after Benediction ("about 190"). Pick one convention on the card (round 3 left base costs on purpose; this only makes the card consistent).
- holy.md:117 (levels 10 to 14): "do not push back Holy Light or Flash of Light". Flash of Light comes at 20; "your Holy Light" is enough for those levels.

## Decisions

- **R8-1, R8-2:** decided on talents/paladin.json (row membership) and the client's Benediction class mask as the pinned engine reads it. Under the round 1 lesson the description wins where it says something the flag does not. Here they agree: "instant cast spells" is base-instant, and the mask confirms that Hammer of Wrath is outside it. No owner question.
- **Races:** unchanged. Retribution Human (Undead), Holy Human (Undead), Protection Dwarf (Undead). The open Beast-target sim from round 7 is not re-filed; the page states no result for it.
- **Oath of Light on low ranks** ("adds the same flat amount to every rank", holy.md:78 and the Holy card note): not filed. Nothing in the 70755 text scales Oath by rank, and no Forever source says it does (searched 2026-11-04, nothing found). It stays the round 1 L1 reading.

## Sims requested

None new. Still open from earlier rounds, not re-filed:
1. Retribution Human against Dwarf on a Beast target, bands 30, 40, 50 and 60 (round 7, sim 1).
2. Protection threat by ability on 1 and 4 targets (round 7, sim 2).
3. Where Dwarf's Retribution raid edge comes from; the Crusader opener (R3-6); the leveling Consecration rank under the level-gap curve.

## Sim findings

None new. The engine's Benediction mask, Exorcism (no crit), Seal of Fury (Holy damage per swing, absorb with a shield) and the Humanoid default target all agree with the 70755 client data and the page.

## Sources

- 70755 client data, data/builds/1.60.1.70755: talents/paladin.json (tier 2 row, prereq_talent_id), trainables/paladin.json and trainables/hunter.json (Heart of the Lion 409580), spelltips/paladin.json (1311649, 20423, 20271, 1310994, 10310, 25780, 10333, 407632, 1239570, 1311595, 1311606, 20929, 20930, 5573, 1020, 1311015, 20920, 20293, 20308, 24239, 10318, 10314), races.json, combos.json, items.json (21546, 2361, 21290), addon-builds.json.
- data/curated/apl/paladin-*.json (card notes).
- Engine d1d9f2775: sim/paladin/retribution/benediction_mask_test.go, sim/paladin/holy_shield.go:133, sim/paladin/exorcism.go.
- Site: sim/request/request.go:445-547, sim/api/envelope.go:280 (Humanoid default target).
- design/reviews/guide-accuracy/round5/race-sims.md, round4/race-sims.md.
- Blizzard, World of Warcraft Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Blizzard, World of Warcraft Forever What's Next panel recap: https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap
- Wowhead, Undead Paladins Coming to World of Warcraft Forever: https://www.wowhead.com/forever/news/undead-paladins-coming-to-world-of-warcraft-forever-382823
- Power Up Gaming, New WoW Forever Talents: Changes vs Classic Guide: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Warcraft Tavern, PvE Retribution Paladin Rotations & Cooldowns (Classic): https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/
- Mobalytics, Holy Paladin Guide, WoW Forever (searched for Oath of Light rank behaviour; it says nothing on it): https://mobalytics.gg/wow-forever/classes/holy-paladin-guide

## Counts

| status | count |
|---|---|
| CONFIRMED | 0 |
| LIKELY | 2 (R8-1 Vindication row wording, a round 7 edit; R8-2 Benediction and Hammer of Wrath) |
| NEEDS-SIM | 0 new (3 carried) |
| NEEDS-OWNER | 0 |

Levels 1 to 19 text: 0 errors. Source links: 0 errors. Round 7 edits: 0 confirmed errors (one LIKELY wording, R8-1).

## Fix log (round 8)

- PALADIN-R8-1 fixed, retribution.md:70: row named (talents/paladin.json: Vindication, Unyielding Faith, Seal of Command, Pursuit of Justice share the tier).
- PALADIN-R8-2 fixed, retribution.md:60: Benediction names its spells; Hammer of Wrath at full price.
- retribution.md:125 ("10% less mana on every instant" at 10 to 14) left: true at those levels, before Hammer of Wrath.
- SHOULDs (Holy Strike card cost, holy.md:117 Flash of Light) left: wording only.
- Prose check: paladin clean. No new SIM findings.

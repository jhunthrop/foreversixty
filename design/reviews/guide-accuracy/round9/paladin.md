# Guide accuracy round 9: Paladin

Auditor: paladin (class index, Holy, Protection, Retribution). Audit tree gr9/audit (land/guides-r8, head 645caa983). Client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go). Read-only: no guide edits, no sims run.

Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, plus the rendered card notes in `data/curated/apl/paladin-*.json`. Before starting I read the paladin reports from rounds 1 to 8 (with every fix log), round4/paladin-levels.md, the race-sims files and ledger.md.

What changed since round 8's audit (git diff 09b98d9e4..HEAD on the paladin guides, APLs, talents, addon-builds.json and the paladin BiS files): only the two round 8 fixes in retribution.md (:60 Benediction, :70 Vindication row). Builds, point orders, BiS files and card notes are unchanged.

## Builds decoded (unchanged, all legal on 70755)

- **Retribution `build`** `050250002/502/0523330100130135`: Divine Strength 5, Reverence 2, Improved Seals 5, Holy Conduit 2; Divine Intellect 5, Precision 2; Benediction 5, Improved Judgement 2, Conviction 3, Sanctified Judgement 3, Vindication 3, Seal of Command 1, Sacred Arbiter 1, Two-Handed Weapon Specialization 3, Repentance 1, Vengeance 3, Champion of the Light 5 (14/7/30).
- **Retribution `raidBuild`** `050050002/503/05233301001301351`: Reverence 0, Precision 3, Unsealed Wrath 1, the rest as above (12/8/31). retribution.md:58 matches.
- **Protection `build`** `05005/15302102513501031/5003`: Divine Strength 5, Improved Seals 5; Divine Intellect 1, Redoubt 5, Precision 3, Sacred Duty 2, Templar's Bulwark 1, Shield Specialization 2, Toughness 5, Swift Judgement 1, Improved Righteous Fury 3, Anticipation 5, Holy Shield 1, Iron Creed 3, Shieldsworn 1; Deflection 5, Conviction 3 (10/33/8). protection.md:65-81 matches, and :81's left-out list is exactly the four untaken Protection talents.
- **Holy `build`** `50320215012101251/50322/0503`: Spiritual Focus 5, Healing Light 3, Reverence 2, Purifying Power 2, Divine Favor 1, Illumination 5, Divine Precision 1, Infusion of Light 2, Oath of Light 1, Holy Shock 1, Boundless Light 2, Holy Power 5, Light's Vigil 1; Divine Intellect 5, Precision 3, Guardian's Favor 2, Sacred Duty 2; Benediction 5, Conviction 3 (31/12/8). holy.md:51-67 matches.

## Round 8 fixes re-checked

- **R8-1** (retribution.md:70): correct. Tier 2 of Retribution (talents/paladin.json) is Vindication, Unyielding Faith, Seal of Command and Pursuit of Justice; Unyielding Faith is fear and disorient duration, Pursuit of Justice is movement speed.
- **R8-2** (retribution.md:60): correct. Each named spell carries a family_mask bit inside Benediction 20101's mask (0x1be11fbf / 0x26018242 / 0x20212100 / 0x8800): Seals, Judgement 20271, Holy Strike, Consecration, Exorcism. Hammer of Wrath 24239 is mask word 1 0x80, outside it. The engine test agrees (benediction_mask_test.go at d1d9f2775). The same check reaches holy.md:64, which round 8 did not read against the mask (R9-1).

## Checked and correct (no finding)

Every line of the four pages and every rendered card note, against spellconst/paladin.json (cost, cost_pct, cast, gcd, cooldown, family_mask), spelltips/paladin.json and hunter.json, talents/paladin.json, trainables/paladin.json, spellranks.json, races.json, items.json, simconsumes.json and spelltips/shared.json. Points I re-derived this round:

- **Crit-bonus lesson:** Holy Power is "Increases the critical strike healing bonus of your Holy spells by 50%". The +50% heal bonus times 1.5 is +75%, so "175% instead of 150%" (holy.md:62) is right. Retribution's "a melee crit doubles the hit" (:103) is the base 200%. No other paladin talent raises a crit bonus, and paladins have no Eureka! line.
- **Off the global cooldown:** Judgement 20271 and Swift Judgement 1310994 have gcd 0 (retribution.md:82, protection.md:72, the cards). Hammer of Wrath has a 1.0 s global, so "you never stop swinging to cast it" holds.
- **Cooldowns:** Holy Shock 10 s, Divine Favor 120 s, Voice of Truth 180 s with "immunity to Silence and Interrupt effects for 6 sec", Lay on Hands 20 min (all ranks), Templar's Bulwark, Divine Shield and Divine Protection 5 min, Blessing of Protection 5 min (3 min with Guardian's Favor 2/2, holy.md:47), Repentance 60 s, Holy Shield 10 s against a 10 s duration, Hammer of the Righteous 6 s at 6% base mana, Righteous Fury 30% base mana for 30 min.
- **Lay on Hands rank 3** 10310: maximum health, 550 mana, "Drains all of the Paladin's remaining Mana when used, but does not interrupt Mana regeneration" (holy.md:91, protection.md:106).
- **Kings and Heart of the Lion:** Blessing of Kings 20217 and Greater Blessing of Kings 25898 are "total stats by 10% for 60 min"; Heart of the Lion 409580 is 10% total stats (plus 40 attack power) for nearby allies (protection.md:85).
- **Seals:** Seal of the Crusader rank 6 (306 attack power, 40% faster, Judgement 40 s refreshed by melee), Seal of Wisdom Judgement 40 s, Seal of Command Judgement doubled on a stunned or incapacitated target, Seal of Fury rank 1 1311649 at 10.
- **Elixir of Holy Power** 21546 (required level 40) uses 1310077: "Increases spell Holy damage by up to 40 for 30 min" (retribution.md:115).
- **Trainer levels (1 to 19 texts):** Seal of Righteousness, Holy Light, Devotion Aura at 1; Judgement and Blessing of Might at 4; Holy Strike, Divine Protection, Holy Light rank 2 at 6; Hammer of Justice and Purify at 8; Lay on Hands, Righteous Fury, Seal of Fury at 10; Redemption at 12; Blessing of Wisdom and Holy Light rank 3 at 14. Consecration, Exorcism, Blessing of Kings, Flash of Light and Voice of Truth at 20; Hammer of the Righteous 40; Holy Wrath 50; Holy Shield ranks 30/40/50/60.
- **Point orders** (addon-builds.json, unchanged): retribution.md:74, :125; protection.md:132, :138; holy.md:117, :123 all match, and each 1 to 19 block reads right ahead of its 20 to 60 block.
- **Races:** combos.json class 2 is Human, Dwarf, Undead; every racial line matches races.json. No paid race is recommended.
- Stat paragraphs rank nothing and every reason holds on 70755. No gear piece named. No em or en dashes, no site or sim talk.

## Source links (fetched 2026-11-04)

All nine distinct URLs return 200 (Blizzard Deep Dive and What's Next recaps, the Wowhead Undead paladin article, Power Up Gaming, Warcraft Tavern's Classic Retribution page, the paladin planner, /sim/weights, and combos.json and races.json as application/json). Labels unchanged and still match. The source-link rule holds on all four pages.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R9-1 | LIKELY | talent | holy.md:64 | "**Benediction** (5 points) and **Conviction** (3 points): Holy Shock and every instant 10% cheaper, and 6% crit on your heals and everything else you do." | This is the same miss round 8 fixed for Hammer of Wrath on the Retribution page. Two of Holy's own instants with a mana cost sit outside Benediction 20101's class mask (0x1be11fbf / 0x26018242 / 0x20212100 / 0x8800). **Oath of Light** 1328548 is instant at 8% of base mana, and **Divine Favor** 20216 is instant at 4% of base mana. Both have family_mask [0, 0, 0, 0] in spellconst/paladin.json, so no class-mask modifier can reach them. The engine prices both at full cost (sim/paladin/oath_of_light.go:26 and :56, sim/paladin/divine_favor.go:64-69, neither with a Benediction multiplier). Holy Shock (0x200000), Cleanse, Purify, every Blessing, the Seals and Judgement are inside the mask. The client text "all instant cast spells and abilities" reads the other way for these two base-instant spells, so this is LIKELY, not CONFIRMED. The fix below is true under either reading. | "**Benediction** (5 points) and **Conviction** (3 points): Holy Shock, Cleanse, your Blessings and your Seals and Judgement cost 10% less, and you get 6% crit on your heals and everything else you do." |

No CONFIRMED findings.

## SHOULDs (wording, not findings)

- retribution.md:121 (1 to 19): "it and Holy Strike both come back every 10 seconds" is the trainer baseline. Two lines later Improved Judgement makes Judgement 9 seconds at 15 and 8 at 16. The two read fine in order. "come back every 10 seconds until Improved Judgement" would read cleaner.
- protection.md:122: "Elixir of Holy Power adds to every Holy threat button you have". The elixir is Holy spell damage (1310077). Whether Hammer of the Righteous (3 × main-hand DPS) and Holy Strike's flat Holy part take spell damage is not in the client text. "adds to Seal of Fury, Judgement, Holy Shield and Consecration" would only name what the text supports.
- Carried from round 8, still open: the Holy Strike card's base cost against Benediction (Retribution card), and holy.md:117 naming Flash of Light at levels 10 to 14.

## Decisions

- **R9-1:** decided on the 70755 class masks and the pinned engine, as round 8 did for Hammer of Wrath. The description and the mask conflict for Oath of Light and Divine Favor, so the fix names only spells that both readings discount. No owner question.
- **Races:** unchanged. Retribution Human (Undead), Holy Human (Undead), Protection Dwarf (Undead). The open Beast-target sim is not re-filed; the page claims no result for it.

## Sims requested

None new. Still open, not re-filed:
1. Retribution Human against Dwarf on a Beast target, bands 30, 40, 50 and 60 (round 7, sim 1).
2. Protection threat by ability on 1 and 4 targets (round 7, sim 2).
3. Where Dwarf's Retribution raid edge comes from; the Crusader opener (R3-6); the leveling Consecration rank under the level-gap curve.

## Sim findings

None new. The engine's Benediction coverage (no discount on Hammer of Wrath, Oath of Light or Divine Favor) matches the client class masks.

## Sources

- 70755 client data, data/builds/1.60.1.70755: spellconst/paladin.json (family_mask, cost_pct, gcd_ms, cooldowns for 20101's mask, 1328548, 20216, 20271, 1310994, 24239, 407632, 1311606, 20473, 20929, 20930, 4987, 1152, the Blessings and Seals), spelltips/paladin.json (1310897, 10310, 20217, 25898, 20308, 20166, 20357, 20920, 20423, 1311649, 1239570, 20924, 26573, 19746), spelltips/hunter.json (409580), spelltips/shared.json (21546 to 1310077), talents/paladin.json, trainables/paladin.json, spellranks.json (25292 book), races.json, combos.json, items.json (21546, 2361), simconsumes.json, addon-builds.json.
- data/curated/apl/paladin-*.json (card notes), data/curated/book-ranks.json.
- Engine d1d9f2775: sim/paladin/retribution/benediction_mask_test.go, sim/paladin/oath_of_light.go, sim/paladin/divine_favor.go.
- Blizzard, World of Warcraft Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Blizzard, World of Warcraft Forever What's Next panel recap: https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap
- Wowhead, Undead Paladins Coming to World of Warcraft Forever: https://www.wowhead.com/forever/news/undead-paladins-coming-to-world-of-warcraft-forever-382823
- Power Up Gaming, New WoW Forever Talents: Changes vs Classic Guide: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Warcraft Tavern, PvE Retribution Paladin Rotations & Cooldowns (Classic): https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/

## Counts

| status | count |
|---|---|
| CONFIRMED | 0 |
| LIKELY | 1 (R9-1, Benediction on Oath of Light and Divine Favor) |
| NEEDS-SIM | 0 new (3 carried) |
| NEEDS-OWNER | 0 |

Levels 1 to 19 text: 0 errors. Source links: 0 errors. Round 8 edits: 0 errors.

## Fix log (round 9)

Applied by the integrator; see the round 9 ledger entry and the commit for the exact lines.

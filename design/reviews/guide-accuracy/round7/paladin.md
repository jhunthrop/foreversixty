# Guide accuracy round 7: Paladin

Auditor: paladin (class index, Holy, Protection, Retribution). Audit tree gr7/audit (land/guides-r6, head 394562f20). Client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go). Read-only: no guide edits, no sims run.

Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, plus the rendered card notes in `data/curated/apl/paladin-*.json`. Before starting I read the paladin reports from rounds 1 to 6 (with the round 5 and round 6 fix logs), round4/paladin-levels.md, the race-sims files for rounds 2 to 5, and ledger.md.

What changed since round 6's audit (git diff 367cbadb7..HEAD): only the round 6 fixes on holy.md, protection.md and retribution.md, and the Holy Light card note in paladin-holy.json. Builds, addon-builds.json, the BiS files and rank-choices.json are unchanged.

## Builds decoded (unchanged since round 1, all legal on 70755)

- **Retribution `build`** `050250002/502/0523330100130135` (14/7/30) and **`raidBuild`** `050050002/503/05233301001301351` (12/8/31): Reverence 2 out, Precision 3 and Unsealed Wrath 1 in. retribution.md:58 matches.
- **Protection `build`** `05005/15302102513501031/5003` (10/33/8). protection.md:65-81 matches.
- **Holy `build`** `50320215012101251/50322/0503` (31/12/8). holy.md:51-67 matches.
- Prerequisites (prereq_talent_id): Illumination needs Reverence, Boundless Light and Light's Vigil need Holy Shock, Shield Specialization needs Redoubt, Swift Judgement needs Templar's Bulwark, Shieldsworn needs Holy Shield, Vengeance needs Seal of Command, **Unsealed Wrath needs Repentance** (see R7-5).

## Shipped point orders (addon-builds.json, index 0 = level 10, decoded again)

- **Retribution:** Benediction 10-14, Improved Judgement 15-16, Conviction 17-19, Sanctified Judgement 20-22, Vindication 23-25, Seal of Command 26, Sacred Arbiter 27, Two-Handed Weapon Specialization 28-30, Repentance 31, Vengeance 32-34, Champion of the Light 35-39, then Holy and Protection (raid: Unsealed Wrath at 40). retribution.md:74 and :125 match.
- **Protection:** Divine Intellect 10, Redoubt 11-15, Precision 16-18, Sacred Duty 19, Templar's Bulwark 20, Sacred Duty 21, Shield Specialization 22-23, Toughness 24-28, Swift Judgement 29, Holy Shield 30, Improved Righteous Fury 31-33, Anticipation 34-38, Iron Creed 39, Shieldsworn 40, Iron Creed 41-42. protection.md:132 and :138 match.
- **Holy:** Spiritual Focus 10-14, Healing Light 15-17, Reverence 18-19, Purifying Power 20-21, Divine Favor 22, Illumination 23-24, Oath of Light 25, Illumination 26-28, Divine Precision 29, Infusion of Light 30-31, **Holy Shock 32**, Boundless Light 33-34, Holy Power 35-39, Light's Vigil 40. holy.md:117 and :123 match; holy.md:81 does not (R7-3).

## Round 6 fixes re-checked

- R6-1 / new stat rule (retribution.md:103, protection.md:110, holy.md:95): no ranking words, each ends by pointing at the table, which renders above the prose. Correct, except one reason in the Retribution paragraph that predates round 6 (R7-1).
- R6-2 (retribution.md:111): "slightly ahead on the raid build in a raid" matches round5/race-sims.md (Dwarf 783.39/782.97 against Human 780.79/780.26). "Human is ahead solo and while leveling" rests on a humanoid target only (R7-2).
- R6-3 (holy.md:53): matches the talent text ("a 70% chance to not lose casting time when you take damage").
- R6-4 (protection.md:132, holy.md:117): reads right in the new order.
- R6-5 (protection.md:110): matches the table's "Not significant" label for Block.
- R6-6 (holy.md:81 and the Holy Light card note): correct. spellranks.json marks Holy Light rank 9 (25292) `book: true`, data/curated/book-ranks.json lists it, and loot.json has no Forever source for Libram: Holy Light IX.

## Checked and correct (no finding)

Every line of the four pages and every rendered card note was checked against talents/paladin.json, trainables/paladin.json, spellranks.json, spelltips/paladin.json, races.json, combos.json, rank-choices.json, book-ranks.json and the APLs. Highlights:

- 1 to 19 trainer levels: Seal of Righteousness, Holy Light, Devotion Aura 1; Judgement and Blessing of Might 4; Holy Strike, Divine Protection, Holy Light rank 2 at 6; Hammer of Justice and Purify 8; Lay on Hands, Righteous Fury, Seal of Fury 10; Redemption 12; Blessing of Wisdom and Holy Light rank 3 at 14. Sacred Duty rank 1 is 30 s off Divine Protection.
- Seal of Fury (1311649 to 20423): absorb only with a shield, Judgement deals Holy damage and "taunts the target to attack you for 3 sec". Judgement 20271 and Swift Judgement 1310994: 10 yd, Seal kept. Seal of the Crusader 20308: 306 AP, 40% faster, Judgement 40 s refreshed by melee.
- Consecration rank 1 (26573) 135 mana; rank 4 bare and rank 5 raid at 60 match rank-choices.json. Hammer of Wrath 24239: 1.0 s cast, 6 s, 30 yd, below 20%. Holy Wrath 50 and 60, 20 yd, Undead and Demons. Exorcism 15 s, Undead and Demons.
- Holy Strike rank 8: 20 mana, 10 s (6 s with Shieldsworn). Holy Shield rank 4 (1239570): 240 mana, 10 s, four charges. Hammer of the Righteous 407632: 6% base mana, 6 s, target plus 3. Righteous Fury 25780: 30% base mana, 30 min, 33% plus 90% on Holy.
- Holy ranks: Flash of Light rank 4 at 42 (18 levels behind at 60, so nearly all of the healing bonus under the level-gap rule), Holy Shock rank 3 at 48, rank 1 160 mana against rank 4 325. Light's Vigil rank 3 1340 mana, 1.5 s. Lay on Hands rank 3: maximum health, 550 mana, drains your mana, regeneration keeps running. Divine Favor 2 min. Voice of Truth 3 min, 6 s.
- Talent texts: Holy Power (engine 75% bonus, round 5), Illumination's 50% casting regeneration (round 5, 20272), Divine Precision rank 1 4% and 10%, Boundless Light 30%, Infusion of Light 1.0 s within 15 s, Iron Creed, Shieldsworn, Shield Specialization 4% once every 3 s, Sacred Duty, Improved Righteous Fury 9%, Anticipation 20, Vengeance non-periodic crits, Unsealed Wrath 15 s and 3 stacks.
- Index: Holy Strike 6 at 10 s, Consecration and Blessing of Kings 20, Hammer of the Righteous 40, Seal of Fury and Righteous Fury 10, milestone rows at tiers 2/3/4/6 (11/16/21/31) in all three trees, Holy Shield at tier 4, no Light, Salvation or Sanctuary in trainables, races Human/Dwarf/Undead.
- Racial texts (races.json): Ruthless 1% crit on spells and attacks, The Human Spirit 10%, Will to Survive, Stoneform 10% physical for 8 s, Artisan's Precision 1% hit, Big Game Hunter 5% to Beasts, Will of the Forsaken without immunity, Cannibalize health and mana, Touch of the Grave up to 5% of maximum health once every 15 s.
- Stat paragraphs: no ranking words, no numbers, and no gear piece named anywhere on the four pages. No other section ranks stats or gear by a nightly value.
- No em or en dashes, no site or sim talk.

## Source links (fetched 2026-11-04)

All nine distinct URLs return 200: both Blizzard recaps, the Wowhead Undead paladin article, Power Up Gaming, the Warcraft Tavern Classic rotation page, the paladin planner, /sim/weights (Retribution only), and combos.json and races.json (application/json). Labels match the pages, as in round 6. The source-link rule (ledger, round 4) holds on all four pages.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R7-1 | CONFIRMED | stats | retribution.md:103 | "Strength is attack power on every swing and Seal proc, and Divine Strength raises it by 10%." | Seal of Righteousness procs do not use attack power. Client text 20293: "${(1786/87+$<1h>)*$<mult>} to ${(1786/25+$<2h>)*$<mult>} Holy damage. Slower weapons cause more Holy damage per swing". There is no AP term. The engine at d1d9f2775 agrees, sim/paladin/sor.go: the proc is `value * 1.2 * SwingSpeed` with a two-hander, plus a spell power coefficient, and nothing from AP. Seal of Command's proc does use AP: sim/paladin/soc.go:128 is `MHWeaponDamage(sim, spell.MeleeAttackPower(target))`. The same page says at :89 that Righteousness "scales with spell damage". It is your only Seal while leveling and half of the raid twist, so "every Seal proc" is false. | "Strength is attack power on every swing, every Holy Strike and every Seal of Command proc, and Divine Strength raises it by 10%. Seal of Righteousness works differently: it scales with your weapon's speed and your spell damage, not attack power." |
| PALADIN-R7-2 | LIKELY (NEEDS-SIM) | consumable/race (round 2 decision, round 6 wording) | retribution.md:111 | "**Dwarf** comes out slightly ahead on the raid build in a raid, but Human is ahead solo and while leveling, which is why Human stays the pick." | Human's solo lead is 283.1 against 280.0, about 1.1% (round5/race-sims.md). That was measured against the engine's default Humanoid target. Big Game Hunter (races.json; engine sim/core/racials.go:302-307, `applyMobTypeDamageMultiplier` 1.05 on Beast targets) adds 5% to everything Dwarf deals to Beasts. The engine models it, but no paladin race sim used a Beast target. Against a Beast, the arithmetic is 280.0 × 1.05 = 294 against Human's 283, so Dwarf comes out ahead by roughly 4%. Beasts make up a large share of leveling kills, so "ahead while leveling" is not established. Round 2 set Big Game Hunter aside ("only works on Beasts") without measuring it. | Until the sim below reports: "**Dwarf** comes out slightly ahead on the raid build in a raid, and Big Game Hunter adds 5% to everything you deal to Beasts. Against everything else Human is ahead solo and in dungeons, and Will to Survive is a stun break Dwarf has no answer to." Then follow the Decision below. |
| PALADIN-R7-3 | CONFIRMED | leveling / spell | holy.md:81 | "Rank 1 costs half as much, but you learned it at 30, so at 60 it keeps only part of your healing power." | Holy Shock is a tier 4 talent (talents/paladin.json, tier 4, so it needs 20 Holy points). The shipped Holy order takes it at 32 (addon-builds.json, paladin-holy leveling index 22). The page also tells most Holy paladins to level as Retribution and respec, and they learn it at 60. Nobody learns it at 30. The penalty depends on the rank's own level: rank 1 (1311606) is a level 30 rank (spellranks.json), 30 levels behind you at 60. | "Rank 1 costs half as much, but it is a level 30 rank, so at 60 it keeps only part of your healing power." |
| PALADIN-R7-4 | LIKELY | rotation (prose and card note) | protection.md:94, and the Hammer of the Righteous card note at paladin-protection.json:251 | "It hits your target and the enemies next to it, and it is your biggest single threat button." | The rotation's own tank-fight measurement (paladin-protection.json:4, 10,000 iterations, each line removed in turn): without Holy Shield, 21% less threat; without Hammer of the Righteous, 17% less. Per press, Hammer is 3 × main-hand DPS in Holy damage on one target (407632). Holy Shield rank 4 is 221 Holy per block across four charges, at 20% extra threat, and Shieldsworn fires it again on every Holy Strike. On a single boss, Holy Shield makes more threat than Hammer. The page itself says at :47 that "Holy Shield adds more on every block". | protection.md:94: "It hits your target and up to three enemies next to it, for Holy damage that Righteous Fury multiplies, so it is where your threat on a pack comes from." Card note: "On cooldown, once Holy Shield's mana is safe. It hits your target and up to three enemies next to it, and after Holy Shield it is your biggest threat on a pack." |
| PALADIN-R7-5 | LIKELY | talent | retribution.md:70 | "**Vindication** and one point of **Repentance** are filler on the way to Champion of the Light, since nothing else in the tree adds damage there." | That holds for the leveling build. On the raid build, Repentance is the required point for the capstone: Unsealed Wrath's prereq_talent_id is 105694, Repentance (talents/paladin.json). The sentence before it already names Seal of Command as "a required point on the way to Vengeance", so the page notes one prerequisite and leaves out the other. | "**Vindication** is filler on the way to Champion of the Light, since nothing else in that row adds damage. **Repentance** is the point Unsealed Wrath requires, and it still earns its place in a dungeon: a 6 second incapacitate on a Humanoid, on a one minute cooldown." (This replaces the last two sentences of the paragraph.) |
| PALADIN-R7-6 | LIKELY | voice / number | holy.md:40 | "only one party heal, Light's Vigil, which costs as much as ten top-rank Flash of Lights" | Light's Vigil rank 3 (1311595) costs 1340 mana. Flash of Light rank 6 (19943) costs 140, so ten would be 1400. The Holy card note already says "almost ten". | "only one party heal, Light's Vigil, which costs almost as much as ten top-rank Flash of Lights" |

## Decisions

- **Retribution Alliance race (R7-2).** The prose fix above goes in now. The pick is then decided by sim 1: if Dwarf's lead on a Beast target at the leveling bands is larger than Human's lead on a Humanoid target, and Dwarf is also ahead in a raid, the frontmatter becomes `recommendedRaces: [dwarf, undead]` with `dwarf` in both build strings. The race line then reads "**Dwarf** is the Alliance pick: it is slightly ahead in a raid, and Big Game Hunter adds 5% against Beasts while you level. **Human** is ahead against Humanoids solo and brings Will to Survive, a stun break." Otherwise Human stays. Round 2 reasoned "Big Game Hunter only works on Beasts", which cannot settle a leveling claim when much of leveling is Beasts, so this decision supersedes it.
- **R7-4:** settled on the rotation's own measurement and the client texts. The new wording ranks nothing on a single target.
- **Holy and Protection races:** unchanged (round 4 sims). Big Game Hunter does not change a healer or tank pick that rests on healing and damage-taken numbers.

## Sims requested

1. **Retribution, Human against Dwarf against a Beast target.** Leveling build at bands 30, 40, 50 and 60, bare preset, with the encounter target's MobType set to Beast, then the same runs on the default Humanoid target. Published gear, 10,000 iterations, seeds 1 and 50001. Report DPS ± error per race. This settles R7-2 and the Retribution race pick.
2. **Protection, threat by ability on the tank fight.** spec-breakdown -tank, guide build, Dwarf, raid band 60, published raid gear, 10,000 iterations, seeds 1 and 50001. Report threat per second from Holy Shield (including the Shieldsworn procs), Hammer of the Righteous, Holy Strike, Judgement, Swift Judgement and Consecration, on 1 target and on 4 targets. This confirms R7-4's wording.
3. Carried from round 6, not re-filed: where Dwarf's Retribution raid edge comes from (Artisan's Precision off and Ruthless off), the Crusader opener (R3-6), and the leveling Consecration rank under the level-gap curve (prose "from 40 move up" against rank-choices.json rank 1 from 30 to 59).

## Sim findings

None new. The engine's Seal of Righteousness (weapon speed plus spell power, no AP), Seal of Command (weapon damage with AP plus a 0.29 spell power coefficient), Big Game Hunter (1.05 against Beasts) and Judgement on the melee table all match the 70755 client text. One gap: every race comparison runs on a Humanoid target, so Big Game Hunter never shows up in any race pick. That affects every Dwarf comparison, not just paladins.

## Sources

- 70755 client data, data/builds/1.60.1.70755:
  - talents/paladin.json (tiers, prereq_talent_id, rank text), trainables/paladin.json, spellranks.json (book flag on 25292), spelltips/paladin.json (20293, 20920, 1311606, 19943, 1311595, 407632, 1239570)
  - addon-builds.json, races.json (Big Game Hunter), combos.json, loot.json
- data/curated: apl/paladin-*.json (card notes and the tank-fight measurement in paladin-protection.json:4), rank-choices.json, book-ranks.json.
- Engine d1d9f2775: sim/paladin/sor.go (proc damage), sim/paladin/soc.go:125-129, sim/core/racials.go:151-172 and :302-307.
- design/reviews/guide-accuracy/round5/race-sims.md, round4/race-sims.md, round2/paladin.md (Decisions).
- Blizzard, World of Warcraft Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Blizzard, World of Warcraft Forever What's Next panel recap: https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap
- Wowhead, Undead Paladins Coming to World of Warcraft Forever: https://www.wowhead.com/forever/news/undead-paladins-coming-to-world-of-warcraft-forever-382823
- Power Up Gaming, New WoW Forever Talents: Changes vs Classic Guide: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Warcraft Tavern, PvE Retribution Paladin Rotations & Cooldowns (Classic): https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/

## Counts

| status | count |
|---|---|
| CONFIRMED | 2 (R7-1 Strength and Seal of Righteousness; R7-3 Holy Shock "learned it at 30") |
| LIKELY | 4 (R7-2 race against Beasts, also NEEDS-SIM; R7-4 Hammer of the Righteous as the biggest threat button; R7-5 Repentance; R7-6 Light's Vigil cost) |
| NEEDS-SIM | 2 new requests (race on a Beast target, threat by ability), 3 carried |
| NEEDS-OWNER | 0 |

Levels 1 to 19 text: 0 errors. Source links: 0 errors. Round 6 edits: 0 errors (R7-2's wording problem goes back to the round 2 decision).

## Fix log (round 7)

- **R7-1** fixed, retribution.md:103. Strength now covers swings, Holy Strike and Seal of Command procs; Seal of Righteousness is said to scale with weapon speed and spell damage, not attack power (client 20293, engine sor.go/soc.go).
- **R7-2** fixed per integrator, retribution.md:111. Human stays the pick. The line now says Human is ahead solo against Humanoids and that Dwarf's Big Game Hunter adds 5% to damage dealt to Beasts (races.json client text). It makes no claim about who wins on Beasts. Frontmatter and build strings unchanged.
- **R7-3** fixed, holy.md:81. "you learned it at 30" is now "it is a level 30 rank".
- **R7-4** fixed per integrator, protection.md:94 and the card note at paladin-protection.json:251 (notes only, action unchanged). "Your biggest single threat button" is gone. The lasting reason is that it hits the target and up to three nearby enemies (client 407632) for Holy damage that Righteous Fury multiplies, so it carries your threat on a pack. Nothing is ranked against Holy Shield.
- **R7-5** fixed, retribution.md:70. Vindication is filler; Repentance is named as the point Unsealed Wrath requires (prereq 105694), and the dungeon incapacitate is kept.
- **R7-6** fixed, holy.md:40. "costs as much as" is now "costs almost as much as" (1340 against 10 × 140).
- **Decision (Retribution race):** Human kept until sim 1 reports. The decision in the audit stands for after that run.
- **Stat priority rule:** the Stat priority prose on holy.md, protection.md and retribution.md already ranks nothing and points to the table. The R7-1 edit adds no ranking words. Prose test `-t paladin` passes. The class index is unchanged, so its updated date stays at 2026-11-02. The three spec pages were already dated 2026-11-04.
- **SIM findings:** none new. Still open: sim 1 (Retribution Human against Dwarf on a Beast target, bands 30/40/50/60) and sim 2 (Protection threat by ability on 1 and 4 targets). The three items carried from round 6 are still open too.

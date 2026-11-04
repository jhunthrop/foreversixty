# Guide accuracy round 5: Paladin

Auditor: paladin (class index, Holy, Protection, Retribution). Audit tree gr5/audit (land/guides-r4, head 42f10dda9, which carries 752c02e43, 77bd7c12b and a901bae80). Client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go). Read-only: no guide edits, no sims run.

Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, plus the rendered card notes in `data/curated/apl/paladin-*.json`. Before starting I read the paladin reports from rounds 1 to 4 (including the round 4 fix log and round4/paladin-levels.md), round2/race-sims.md, round3/race-sims.md, round4/race-sims.md and ledger.md.

## Builds decoded (talents[] order in talents/paladin.json; unchanged since round 1, all legal on 70755)

- **Retribution `build`** `050250002/502/0523330100130135` (14/7/30) and **`raidBuild`** `050050002/503/05233301001301351` (12/8/31): as in round 3. Vengeance needs Seal of Command and Unsealed Wrath needs Repentance (prereq_talent_id), so retribution.md:70 holds.
- **Protection `build`** `05005/15302102513501031/5003` (10/33/8): as in round 3. Swift Judgement needs Templar's Bulwark (protection.md:70), Shieldsworn needs Holy Shield, and Shield Specialization needs Redoubt 5.
- **Holy `build`** `50320215012101251/50322/0503` (31/12/8): as in round 3. Illumination needs Reverence 2.

## Leveling prose against the shipped point order (addon-builds.json, index 0 = level 10, decoded again)

- **Retribution:** Benediction 10-14, Improved Judgement 15-16, Conviction 17-19, Sanctified Judgement 20-22, Vindication 23-25, Seal of Command 26, Sacred Arbiter 27, Two-Handed Weapon Specialization 28-30, Repentance 31, Vengeance 32-34, Champion of the Light 35-39, Divine Strength 40-44, Reverence 45-46, Improved Seals 47-51, Holy Conduit 52-53, Divine Intellect 54-58, Precision 59-60. The raid order takes Unsealed Wrath at 40. retribution.md:74 and :125 match.
- **Protection:** Divine Intellect 10, Redoubt 11-15, Precision 16-18, Sacred Duty 19, Templar's Bulwark 20, Sacred Duty 21, Shield Specialization 22-23, Toughness 24-28, Swift Judgement 29, Holy Shield 30, Improved Righteous Fury 31-33, Anticipation 34-38, Iron Creed 39, Shieldsworn 40, Iron Creed 41-42, then Holy and Retribution. protection.md:132 and :138 match.
- **Holy:** Spiritual Focus 10-14, Healing Light 15-17, Reverence 18-19, Purifying Power 20-21, Divine Favor 22, Illumination 23-24, Oath of Light 25, Illumination 26-28, Divine Precision 29, Infusion of Light 30-31, Holy Shock 32, Boundless Light 33-34, Holy Power 35-39, Light's Vigil 40. holy.md:115 and :121 match.

## Levels 1 to 19, checked paragraph by paragraph

Checked against trainables/paladin.json and spelltips/paladin.json:

- **Retribution (:121-127):**
  - Trainer levels: Seal of Righteousness and Holy Light 1, Judgement and Blessing of Might 4, Holy Strike 6, Hammer of Justice 8, Lay on Hands 10.
  - Judgement "Does not consume the Seal". Judgement 20271 and Holy Strike 679 are both on 10 s cooldowns.
  - Redemption: the trainables row is at 12.
  - Talent text: Benediction 10%, Improved Judgement 2 s off (so 8 s), Conviction 6% at 3/3.
  - Gear: Battleworn Hammer 2361 is a two-handed mace. Seal of Righteousness: "Slower weapons cause more Holy damage per swing". Paladins have no wand skill. Plate at 40 follows the class rule.
  - Wording note: see R5-3.
- **Holy (:111-117):**
  - Trainer levels: Holy Light ranks at 1, 6 and 14. Purify 8, Lay on Hands 10, Blessing of Wisdom 14.
  - Talent text: Spiritual Focus 70%, Healing Light 12%, Reverence "30% of your Mana regeneration to continue while casting".
- **Protection (:128-134):**
  - Trainer levels: Devotion Aura 1, Judgement 4, Holy Strike 6, Hammer of Justice 8. Righteous Fury (25780) and Seal of Fury (1311649) both at 10.
  - Seal of Fury 1311649: "While a shield is equipped, each attack also grants an absorb shield ... inflicts 23 Holy damage to an enemy and taunts the target to attack you for 3 sec".
  - Talent text: Divine Intellect 2% at rank 1; Redoubt 30% for "10 sec or 5 blocks"; Precision 3%. Sacred Duty rank 1 takes 30 s off Divine Protection, and Divine Protection rank 2 is learned at 18.
- **Consistency:** none of the three contradicts its Levels 20 to 60 text or the shipped order.

## Round 4 fixes re-checked

- **R4-1:** the What's Next recap URL returns 200, and the page names Hyjal Summit (20-player), the Barrow Deeps (10-player) and the December 9 unlock. Correct.
- **R4-2:** /sim/weights is now only on retribution.md. Correct.
- **R4-3:** the Battleworn Hammer wording is correct for Human and Dwarf. No Forever source or data file names the Undead paladin's starting weapon, so it is not filed.
- **Race decisions** (round4/race-sims.md: Protection Dwarf 362.3 against Human 361.0 damage taken per second, a tie; Holy Human 676.8 against Undead 662.2 and Dwarf 661.0): protection.md:118 and holy.md:101 match.

## Source links (fetched 2026-11-03; every one returns 200)

| page | label | URL | says what the label says? |
|---|---|---|---|
| all four | Blizzard, Deep Dive panel recap | news.blizzard.com/en-us/article/24303313 | yes (paladin section: Holy Strike at 6, Seals not consumed, Seal of Fury, Reverence, Swift Judgement, Undead Paladin) |
| all four | Blizzard, What's Next panel recap | news.blizzard.com/en-gb/article/24303862 | yes |
| index | Wowhead, Undead Paladins coming to Forever | wowhead.com/forever/news/...-382823 | yes (the title matches) |
| index | Forever Sixty, which races can be paladins | foreversixty.gg/data/1.60.1.70755/combos.json | yes (application/json) |
| index | Power Up Gaming, Forever talent changes vs Classic | powerupgaming.co.uk/2026/09/14/... | yes (the paladin table and Holy Strike entry) |
| ret, holy, prot | Forever Sixty, paladin talents | foreversixty.gg/planner?class=paladin | yes (round 4 rule: talent-tree sources open the planner) |
| ret | Forever Sixty, what each stat is worth | foreversixty.gg/sim/weights | yes (a damage spec, as the rule allows) |
| ret | Warcraft Tavern, Classic Retribution Paladin rotation | warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/ | yes |
| ret, holy, prot | Forever Sixty, races | foreversixty.gg/data/1.60.1.70755/races.json | yes (application/json) |

Source-link rule (ledger, round 4):
- No paladin page links /sources.
- /sim/weights is on Retribution only.
- None of the pages has a boss-drop entry, a "how <spec> plays" entry, a Talents Forever entry or a wago.tools link.
- Compliant.

## Checked and correct (no finding)

I checked every line of the four pages again against talents/paladin.json, trainables/paladin.json, spelltips/paladin.json, spelltips/shared.json and the bis weights. Highlights:

- **Hammer of the Righteous** costs 6% of base mana (spelltips 407632). So protection.md:94 and :100 ("once Holy Shield's mana is safe", "cut ... Hammer of the Righteous") are right, even though the trainables cost field reads 0.
- **Blessings:** Might, Kings, Wisdom and the Greater ranks are all 60 min (index.md:33). Blessing of Protection is 300 s base, 3 min with Guardian's Favor (holy.md:47).
- **Seal of the Crusader 20308:** 306 attack power, "attacks 40% faster, but deals less damage", and its Judgement lasts 40 s and is refreshed by your melee (retribution.md:93).
- **Holy Power:** "175% instead of 150%" matches the engine's crit-bonus modifier (healing_talents.go:38-44, a 50% bonus becomes 75%).
- **Racials** (spelltips/shared.json):
  - Touch of the Grave: a 20% chance to drain health. The 15 s limit is in the engine.
  - Stoneform: 10% physical damage taken for 8 s.
  - Will of the Forsaken: removes Charm, Fear and Sleep, with no immunity.
  - Cannibalize: health and mana.
- **Stat priorities** against the band 60 raid weights:
  - Retribution: Hit 2.16 > Int 1.07 > Str 1.00 > Crit 0.99.
  - Protection: Defense > Block > Dodge > Parry ≈ Agility > Block value > Stamina > Hit.
  - Holy: MP5 2.44 > Int 1.00 > Healing 0.72 > Crit 0.67 > Spirit 0.41.
- **Rotation prose** matches all three APLs step for step:
  - Retribution: Judgement, Exorcism, Hammer of Wrath at E20, the Seal lines, Holy Strike, then Consecration (the 0.8 × time-left gate).
  - Protection: Seal of Fury, Holy Shield, Holy Strike, Hammer of the Righteous, Swift Judgement, Judgement, Exorcism, Consecration above 50%. Bulwark below 35%.
  - Holy: thresholds 40/70/90/80 and the -5 point pace.
- **Light's Vigil:** 1340 mana against Flash of Light 140, so "ten top-rank Flash of Lights" holds.
- **Flash of Light rank 4** (learned 42) keeps about 95% of your healing power at 60, and **Holy Shock rank 1** (learned 30) about 35%, under the level-gap rule. holy.md:78, :81 and :93 follow it.
- **Style:** no em or en dashes, and no site or sim talk.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R5-1 | CONFIRMED | stats / talent | holy.md:93 | "Spirit only works outside the five second rule, apart from Reverence's 30%." | Illumination's mana return (spell 20272) also puts a 15 second buff on you. Wowhead Forever's tooltip for 20272 reads "Regenerating at 50% of normal rate while casting spells. 15 seconds remaining" (https://nether.wowhead.com/forever/tooltip/spell/20272). The engine applies the same effect from the client's SpellEffect row (sim/paladin/healing_talents.go:149-163, SpiritRegenRateCasting +0.50 for 15 s on every Illumination proc). With 5/5 Illumination, every critical heal gives you half your Spirit regeneration while casting, on top of Reverence. The line says Reverence is the only exception, which is false. | ":93: Spirit only works outside the five second rule, apart from Reverence's 30% and the 50% Illumination gives you for 15 seconds after every critical heal, which is one more reason crit is a mana stat for you." Optionally, at :56 add after the first sentence: "It also keeps half your Spirit regeneration running while you cast for 15 seconds." |
| PALADIN-R5-2 | LIKELY | leveling / stats (round 3 edit) | retribution.md:141 | "Ignoring hit at 60. Below the cap, a point of hit is worth more than a point of Strength once you are there; while leveling Strength still comes first." | bis/paladin-retribution.json scale factors, alliance bare. Band 20: Intellect 1.39 > Strength 1.00 > hit 0.66. Band 30: spell power 1.03 > Strength 1.00 > hit 0.79. Band 40: spell power 1.01 > Strength 1.00 > hit 0.97. Band 50: Intellect 1.84 > spell power 1.13 > hit 1.01 ≈ Strength 1.00. "Strength still comes first" reads as "Strength is your best stat while leveling", which is false at every band. The comparison it means, Strength against hit, is true at 20 to 40 and a tie at 50. "Below the cap ... once you are there" also contradicts itself inside one clause. | "Ignoring hit at 60. Until you reach the cap, a point of hit is worth more than a point of Strength there; while leveling, Strength is still worth more than hit." |
| PALADIN-R5-3 | LIKELY | leveling (round 4 text) | retribution.md:125 | "Benediction: 10% less mana on every instant, which is your whole bar." | Holy Light (ranks at 1, 6 and 14) is a 2.5 s cast (trainables cast_time_ms 2500) and is on every leveling paladin's bar for healing between pulls. Benediction is "all instant cast spells and abilities". Seals, Judgement, Holy Strike, Blessing of Might and Hammer of Justice are instant, so it is nearly the whole bar, not all of it. | "...go into Benediction: 10% less mana on every instant, which is every button you fight with." |
| PALADIN-R5-4 | LIKELY | talent | holy.md:121 | "Divine Precision only helps your Holy spells hit enemies." | talents/paladin.json Divine Precision 1310904, rank 1: "Improves your chance to hit with Holy spells by 4%, and reduces the duration of Silence effects on you by 10%." It does two things, and the second matters to a healer. | "...and Divine Precision only adds Holy spell hit and a little silence reduction, nothing for your heals." |

## Decisions

- **R5-1:** decided on the client tooltip and the engine, with no owner question. Wowhead Forever's 20272 buff text is client data.
- **R5-2 to R5-4:** wording only. Each fix keeps the claim the page meant to make.
- **Divine Shield and Divine Protection as tank defensives** (protection.md:106): not filed. The 70755 texts say Divine Shield halves your damage and Divine Protection stops you attacking, so both cost threat. The page already ranks them behind Templar's Bulwark and calls them a last resort. Whether a boss turns off an immune tank is a server behaviour no Forever source documents.
- **Voice of Truth** (trainer 20, 1310897: "immunity to Silence and Interrupt effects for 6 sec", 3 min): the Holy page never names it. A top Holy player keeps it bound for silence and interrupt casters. It is an omission rather than a false claim, so it is not counted. A lane may add to holy.md:87: "Voice of Truth makes you immune to silences and interrupts for 6 seconds, every 3 minutes: save it for the caster pack or boss that would lock you out."
- **Races:** unchanged. Retribution is Human, with Undead for the Horde. Holy is Human, with Undead for the Horde. Protection is Dwarf, with Undead for the Horde. Rounds 2 and 4 measured these.

## Sims requested

None new. Carried from round 3 (not re-filed):

1. The Retribution Crusader opener at bare 60 and band 50.
2. The Retribution leveling Consecration rank under the level-gap curve.

## Sim findings

None. The engine's Illumination (mana return plus 50% casting regeneration for 15 s), Holy Power (75% crit bonus), Lay on Hands drain and autocast below 10%, and Seal of the Crusader refresh all match the 70755 client text.

## Sources

- 70755 client data:
  - data/builds/1.60.1.70755: talents/paladin.json, trainables/paladin.json, spelltips/paladin.json, spelltips/shared.json, addon-builds.json, combos.json, races.json, items.json (2361, 21290, 21546), loot.json, bis/paladin-{retribution,protection,holy}.json.
  - data/curated: apl/paladin-*.json and book-ranks.json.
- Engine d1d9f2775: sim/paladin/healing_talents.go and sim/paladin/lay_on_hands.go.
- Wowhead Forever tooltip, Illumination (20272): https://nether.wowhead.com/forever/tooltip/spell/20272
- Wowhead Forever tooltip, Illumination talent (20210): https://nether.wowhead.com/forever/tooltip/spell/20210
- Blizzard, World of Warcraft Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Blizzard, World of Warcraft Forever What's Next panel recap: https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap
- Wowhead, Undead Paladins Coming to World of Warcraft Forever: https://www.wowhead.com/forever/news/undead-paladins-coming-to-world-of-warcraft-forever-382823
- Power Up Gaming, New WoW Forever Talents: Changes vs Classic Guide: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Warcraft Tavern, PvE Retribution Paladin Rotations & Cooldowns (Classic): https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/
- Warcraft Tavern, Undead Paladin Overview (searched for the Undead starting weapon; none named): https://www.warcrafttavern.com/forever/guides/undead-paladin/

## Counts

| status | count |
|---|---|
| CONFIRMED | 1 (R5-1) |
| LIKELY | 3 (R5-2 to R5-4) |
| NEEDS-SIM | 0 (two carried sim items) |
| NEEDS-OWNER | 0 |

The Levels 1 to 19 text has 0 confirmed errors (one LIKELY wording, R5-3). The source links have 0 errors.

## Fix log (round 5)

- **PALADIN-R5-1 (CONFIRMED): fixed.** holy.md:95 now names the 50% casting regeneration Illumination gives for 15 seconds after every critical heal, next to Reverence's 30%; holy.md:56 adds the same effect to the Illumination talent line. Verified at the engine pin (sim/paladin/healing_talents.go:149-163, aura 134 at 50 for 15 s from SpellEffect 1237075).
- **PALADIN-R5-2 (LIKELY): fixed, differently from the proposal.** The nightly bis regenerated after the audit (de044125c, 2026-11-04). On the current file, hit beats Strength per point at bare 40, 50 and 60 and loses only at 20 and 30, so the proposed "while leveling, Strength is still worth more than hit" would be false. retribution.md:141 now reads "Ignoring hit at 60. Until you reach the cap, a point of hit is worth more than a point of Strength." The leveling clause is gone because it moves with the nightly numbers.
- **PALADIN-R5-3 (LIKELY): fixed.** retribution.md:125 "which is every button you fight with". Benediction's client text is "all instant cast spells and abilities", and Holy Light is a 2.5 s cast.
- **PALADIN-R5-4 (LIKELY): fixed.** holy.md:123 now gives Divine Precision's Holy spell hit and silence reduction (talents/paladin.json 1310904 text).
- **Decision, Voice of Truth: added.** holy.md:89 gets a new "Silences." paragraph. Checked against trainables (learned at 20, 3 min cooldown) and spelltips 1310897 ("immunity to Silence and Interrupt effects for 6 sec").
- **Decisions, Divine Shield/Protection and races:** no change, as decided.
- **Lay on Hands:** not touched. It drains all your mana, as the client text says.
- **Card notes:** no paladin APL note says anything false about Illumination, Benediction or Divine Precision. Unchanged.
- **New observation for the integrator (not a guide edit in this lane):** the regenerated bis changed the band 60 raid weights since the audit. Retribution raid now has hit at 0 (insignificant), and Strength 1.00 > crit 0.93 > spell power 0.73 ≈ Intellect 0.71. Holy raid has spell haste at the top, then mp5 > healing ≈ Intellect ≈ crit > Spirit. Retribution statPriority (Hit, Intellect, Strength, ...) and retribution.md:103 follow the audited weights, not these. Re-check after the refresh settles. If the numbers hold, fix it in a later round.
- **SIM findings:** none new.

## Fix log (round 5, review pass)

Orders read from data/builds/1.60.1.70755/bis/paladin-*.json, band 60 raid scale factors.

1. MUST retribution stat prose: fixed, retribution.md:103. Now Strength, crit, spell power and Intellect close together, Agility; hit matters only up to the cap, the raid gear already reaches it, nothing past it; below the cap take hit (Precision and Artisan's Precision count).
2. MUST retribution hit lines: fixed, retribution.md:72 ("hit is worth taking up to the cap and adds nothing past it") and retribution.md:141 (gear for hit below the cap, stop once you reach it).
3. Dwarf line at retribution.md:111: left unchanged per the integrator ruling (Dwarf even in a raid, Human wins solo).
4. MUST protection stat prose: fixed, protection.md:110. Defense, Stamina, then Dodge, Agility, block value, Parry; more block chance is no gain at 60 but blocks still power Holy Shield and Shield Specialization's mana; armor never goes to zero; hit sits low once Precision is in.
SHOULD holy stat prose: fixed, holy.md:95. MP5 first, then healing power, Intellect and crit close together.
Not touched: retribution.md:9 frontmatter statPriority (component-rendered, not in the review).

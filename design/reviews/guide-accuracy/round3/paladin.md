# Guide accuracy round 3: Paladin

Auditor: paladin (class index, Holy, Protection, Retribution). Audit tree land/guides-r2 7a2165cb0, client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go). Read-only: no guide edits, no sims run.

Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, plus the rendered card notes (`data/curated/apl/paladin-*.json`). Round 1 and round 2 reports, round2/race-sims.md, round3/race-sims.md (no paladin rows) and ledger.md read first.

## Builds decoded (talents[] order in talents/paladin.json; unchanged since round 2, all legal on 70755)

- **Retribution `build`** `050250002/502/0523330100130135` (14/7/30). Holy: Divine Strength 5, Reverence 2, Improved Seals 5, Holy Conduit 2. Protection: Divine Intellect 5, Precision 2. Retribution: Benediction 5, Improved Judgement 2, Conviction 3, Sanctified Judgement 3, Vindication 3, Seal of Command 1, Sacred Arbiter 1, Two-Handed Weapon Specialization 3, Repentance 1, Vengeance 3, Champion of the Light 5.
- **Retribution `raidBuild`** `050050002/503/05233301001301351` (12/8/31): minus Reverence 2, plus Precision 3 and Unsealed Wrath 1.
- **Protection `build`** `05005/15302102513501031/5003` (10/33/8). Holy: Divine Strength 5, Improved Seals 5. Protection: Divine Intellect 1, Redoubt 5, Precision 3, Sacred Duty 2, Templar's Bulwark 1, Shield Specialization 2, Toughness 5, Swift Judgement 1, Improved Righteous Fury 3, Anticipation 5, Holy Shield 1, Iron Creed 3, Shieldsworn 1. Retribution: Deflection 5, Conviction 3.
- **Holy `build`** `50320215012101251/50322/0503` (31/12/8). Holy: Spiritual Focus 5, Healing Light 3, Reverence 2, Purifying Power 2, Divine Favor 1, Illumination 5, Divine Precision 1, Infusion of Light 2, Oath of Light 1, Holy Shock 1, Boundless Light 2, Holy Power 5, Light's Vigil 1. Protection: Divine Intellect 5, Precision 3, Guardian's Favor 2, Sacred Duty 2. Retribution: Benediction 5, Conviction 3.

## Leveling prose against the shipped point order (addon-builds.json, index 0 = level 10)

- **Retribution:** Benediction 10-14, Improved Judgement 15-16, Conviction 17-19, Sanctified Judgement 20-22, Vindication 23-25, Seal of Command 26, Sacred Arbiter 27, Two-Handed Weapon Specialization 28-30, Repentance 31, Vengeance 32-34, Champion of the Light 35-39, then Holy and Protection. retribution.md:74 matches.
- **Protection:** Sacred Duty 19, Templar's Bulwark 20, Sacred Duty 21, Swift Judgement 29, Holy Shield 30, Shieldsworn 40. protection.md:129 matches.
- **Holy:** Oath of Light 25, Divine Precision 29, Infusion of Light 30-31, Holy Shock 32, Light's Vigil 40. holy.md:112 matches.

## Round 2 fixes and decisions re-checked

- R2-1 and R2-2 card notes (paladin-holy.json:584, :255): now state nothing false about the rank the card casts. Correct.
- R2-3 (paladin-retribution.json:148, "about 190 mana"): Seal of Command 210 and Seal of Righteousness 200 base, 189 and 180 with Benediction 5/5. Correct.
- R2-4 (holy.md:36): correct.
- Retribution race decision, retribution.md:111: matches the decided text and race-sims.md (raid tie, bare Human +1.2%). Correct.
- Lay on Hands, holy.md:92 and protection.md:109: match the 70755 description of 10310 ("restores 550 of their mana. Drains all of the Paladin's remaining Mana when used, but does not interrupt Mana regeneration"). Correct.
- Source labels: player words, no paths. Correct.

## Checked and correct (no finding)

Every talent line against talents/paladin.json; every spell level, rank, cost, cooldown and range against trainables/paladin.json and spelltips/paladin.json (Holy Strike 6 and 20 mana at rank 8, Judgement 4 at 6% base mana and 10 yd, Swift Judgement free on 60 s, Seal of Fury and Righteous Fury 10 with RF at 30% base mana for 30 min, Consecration 20 with rank 1 at 135, Exorcism 20, Hammer of the Righteous 40 at 6 s and 3 extra targets, Hammer of Wrath 44 at 1.0 s and 6 s, Holy Wrath 50 at 20 yd, Holy Shield 30/40/50/60, Holy Shock 30/40/48/56, Repentance 6 s on 1 min, Templar's Bulwark 5 min, Blessing of Protection 5 min, 3 with Guardian's Favor, Light's Vigil 1340). Index facts (milestone rows at 11/16/21/31 for all three trees, Holy Shield on Protection's 21 point row, no Light/Salvation/Sanctuary trainer rows, Human/Dwarf/Undead only per combos.json). Racials against spelltips/shared.json, Touch of the Grave's 15 s internal cooldown against engine racials.go:377 (SpellAuraOptions 15000 ms). Hammer of the Righteous as Protection's biggest single threat button (hammer_of_the_righteous.go: 3x weapon DPS incl. attack power, against Holy Strike's 50% of normalized damage plus 93). Stat priorities against bis band 60 raid orders. Rotation prose against all three APLs. Card note costs quote the tooltip base costs (Holy Strike "Twenty mana", Judgement "6% of base mana"), which is what the player sees; not filed. No em or en dashes, no site or sim talk.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R3-1 | CONFIRMED | leveling / stats | retribution.md:129 | "Ignoring hit. Below the cap, a point of hit beats a point of Strength." (in the Leveling section's common mistakes) | bis/paladin-retribution.json scale factors, hit against Strength: band 20 0.66, band 30 0.79, band 40 0.97, band 50 1.01 (tie), bare 60 1.84, raid 60 2.16. While leveling Strength is ahead of hit until 50. Round 1 PALADIN-7 fixed the same claim at :72 to "With raid buffs at 60"; this line still says it unconditionally. | "Ignoring hit at 60. Below the cap, a point of hit is worth more than a point of Strength once you are there; while leveling Strength still comes first." |
| PALADIN-R3-2 | CONFIRMED | talent | retribution.md:72 | "**Reverence** keeps 30% of your regeneration running in combat" | talents/paladin.json Reverence 1310899: "Allows 30% of your Mana regeneration to continue while casting." It works inside the five second rule after a cast, not across combat. holy.md:58 has it right ("while you cast"), so the two pages disagree. | "**Reverence** keeps 30% of your Spirit regeneration running after you cast, which matters more while leveling than in a raid with Blessing of Wisdom on you." |
| PALADIN-R3-3 | CONFIRMED | healing | holy.md:86 | "Every crit refunds half its base cost through Illumination, which is why Holy Shock on the party (where Boundless Light's crit applies) pays for itself." | Illumination 20210: "gain Mana equal to 50% of the base cost of the spell" on a crit. The most any Holy Shock gets back is half its cost, and only when it crits, so it never pays for itself. | "Every crit refunds half its base cost through Illumination, which is why Holy Shock on the party, where Boundless Light adds 30% crit, gets a large part of its cost back." |
| PALADIN-R3-4 | CONFIRMED | talent | protection.md:113 | "Block also feeds Holy Shield's damage, Redoubt and Shield Specialization's mana." | Redoubt 20127: "Damaging melee attacks against you have a 10% chance to increase your chance to block by 30%. Lasts 10 sec or 5 blocks." Getting hit procs it, and each block spends one of its charges. Block chance does not feed it. Holy Shield's damage and Shield Specialization's mana do come from blocks. | "Block also feeds Holy Shield's damage and Shield Specialization's mana." |
| PALADIN-R3-5 | CONFIRMED | rotation (card note) | data/curated/apl/paladin-protection.json, Holy Strike row (10333), rendered on the Protection page | "Your cheapest threat at 20 mana" | trainables/paladin.json Swift Judgement: cost 0. The same card's Swift Judgement row says "free and off the global cooldown", so the card contradicts itself. | "On cooldown. Twenty mana for strong threat; with Iron Creed it also cuts your damage taken for 6 seconds, and Shieldsworn brings it back every 6 seconds." |
| PALADIN-R3-6 | LIKELY | rotation | retribution.md:93 (and :119) | "On your own, skip Crusader: one Seal and one Judgement spent on it are worth more as damage." | Forever's Judgement does not consume the Seal, and Judgement of the Crusader lasts 40 s and is refreshed by your own melee (spelltips 20308; core/debuffs.go:314-330). So one Seal of the Crusader cast and one Judgement give +161 Holy damage taken for the whole fight. Engine coefficients on that bonus: Judgement of Righteousness 0.429, Holy Strike 0.429, Seal of Righteousness proc 0.1 (x1.1), Seal of Command proc 0.29, Consecration 0.042 a tick. That comes to about 25 to 30 DPS before crit and Vengeance at bare 60 (264 DPS), against one global, about 144 mana and one Judgement of Righteousness, once a fight. Mobalytics' Forever Retribution guide opens every pull with Seal of the Crusader and a Judgement on the first swing, then moves to Seal of Righteousness. The APL never casts Crusader, so the guide's claim was never measured. | If the sim below confirms a gain beyond error: ":93 last sentence: On your own, open with Seal of the Crusader, Judge it on your first swing, then put up Seal of Righteousness: your swings keep the debuff up, so one Judgement adds Holy damage to every Seal proc, Judgement, Holy Strike and Consecration tick for the rest of the fight." and ":119: Pull with Seal of the Crusader and a Judgement on your first swing, then Seal of Righteousness and Holy Strike on cooldown." Then add the prepull Crusader to the bare APL. |
| PALADIN-R3-7 | LIKELY | talent | retribution.md:66 | Vengeance: "3% more Physical and Holy damage for 30 seconds after a crit" | Vengeance 20049: "after landing a non-periodic critical strike". The engine agrees (talents.go:154-160, OnSpellHitDealt only). Line :62 has just listed Consecration's ticks among your crits, so the reader assumes they stack Vengeance. They do not. | "...for 30 seconds after a crit from anything but Consecration's ticks, stacking five times to 15%." |
| PALADIN-R3-8 | LIKELY | talent | protection.md:51 | "Holy Shield, Redoubt and Shieldsworn keep your block chance high" | Shieldsworn 1328544 adds no block chance. It cuts Holy Strike's cooldown, fires Holy Shield and Shield Specialization without spending charges, and lets you block spells for 6 s. | "Holy Shield and Redoubt keep your block chance high, and Shieldsworn fires Holy Shield's damage on every Holy Strike and lets you block spells" |
| PALADIN-R3-9 | LIKELY | talent | holy.md:70 | "**Precision** and one point of **Divine Precision** are filler on the way." | Holy's 12 Protection points end in Protection's tier 1 (Divine Intellect 5, Precision 3, Guardian's Favor 2, Sacred Duty 2), so Precision leads nowhere. It is the only tier 1 talent left once the healing ones are taken. Divine Precision is on the way: it is the 25th Holy point that opens Holy Power's row. | "**Divine Precision** is a filler point to reach Holy Power, and **Precision** fills out the Protection points, since nothing else in those rows heals." |

## Decisions

- **Retribution solo opener (PALADIN-R3-6):** the decision goes to the measurement below. If Crusader-then-Righteousness gains beyond error at bare 60, the guide and the bare APL adopt it. If it ties or loses, :93 stays and :119 is unchanged. Forever source: Mobalytics' Forever Retribution guide uses it. Classic and TBC leveling guides agree, but there the Judgement consumed the Seal, which made it cost more than it does in Forever.
- **Holy build's Precision points:** keep the build. The claim is reworded only (R3-9). Healing sims cannot measure a utility swap (Pursuit of Justice), and no Forever source names one.
- **Retribution race:** Human stays (round 2 decision, re-checked).

## Sims requested

1. **Retribution, bare 60 and band 50, guide `build`, published gear, 10,000 iterations, seeds 1 and 50001.** Today's APL against the same APL with:
   - a prepull Seal of the Crusader rank 6 (20308) at -3s;
   - the first Judgement (20271) cast under it;
   - then Seal of Righteousness rank 8 (20293) before the normal priority.

   Report DPS ± error. Also run a 30 s fight length (the leveling pull), where the one-time cost weighs most.
2. Unchanged from round 1 L1(c): the Retribution leveling Consecration rank with the level-gap curve on. Consecration's only spell power part is 0.042 a tick on rank 1's 6 a second base (consecration.go:40-67), so most of what rank 1 does at 40+ is the part the curve takes away. retribution.md:95 and :121 ("from 40 move up") follow the curve; rank-choices.json (rank 1 through 59) predates it.

## Sim findings

None new. Engine Vengeance (non-periodic only), Touch of the Grave (15 s), Lay on Hands drain and Judgement of the Crusader (40 s, refreshed by the caster's melee) all match the 70755 client text.

## Sources

- 70755 client data: talents/paladin.json, trainables/paladin.json, spelltips/paladin.json, spelltips/shared.json, races.json, combos.json, addon-builds.json, bis/paladin-{retribution,protection,holy}.json, items.json; data/curated/apl/paladin-*.json, rank-choices.json, book-ranks.json.
- Engine d1d9f2775: sim/paladin/{sotc.go, sor.go, soc.go, consecration.go, holy_strike.go, hammer_of_the_righteous.go, talents.go, lay_on_hands.go}, sim/core/debuffs.go, sim/core/spell_result.go:674-678, sim/core/racials.go:377.
- design/reviews/guide-accuracy/round2/race-sims.md.
- Mobalytics, Retribution Paladin Guide, WoW Forever: https://mobalytics.gg/wow-forever/classes/retribution-paladin-guide (read 2026-11-04 through the search index; direct fetch returned 403): "Put on Seal of the Crusader as you run up to an enemy and use Judgement on them as you do your first Auto Attack. Follow this up by casting Seal of Righteousness and hitting them with a Holy Strike."
- Icy Veins, TBC Classic Retribution Paladin Leveling Guide: https://www.icy-veins.com/tbc-classic/retribution-paladin-leveling-guide (Crusader for the Holy damage debuff, then Righteousness; pre-Forever context only).

## Counts

CONFIRMED 5 (R3-1 to R3-5), LIKELY 4 (R3-6 to R3-9, R3-6 waiting on a sim), NEEDS-SIM 0 separate (R3-6's sim is listed above), NEEDS-OWNER 0.

## Fix log (round 3)

- PALADIN-R3-1 fixed: retribution.md:129, hit is ahead of Strength at 60, Strength first while leveling.
- PALADIN-R3-2 fixed: retribution.md:72, Reverence keeps 30% of Spirit regeneration running after you cast (matches holy.md:58).
- PALADIN-R3-3 fixed: holy.md:86, Holy Shock on the party gets a large part of its cost back, no longer "pays for itself".
- PALADIN-R3-4 fixed: protection.md:113, Redoubt dropped from what block feeds.
- PALADIN-R3-5 fixed: paladin-protection.json Holy Strike note (line 213), "Twenty mana for strong threat"; no longer calls it the cheapest. Action unchanged.
- PALADIN-R3-6 softened, decision pending the sim: retribution.md:93 no longer tells a solo player to skip Crusader as settled fact. It states what each Seal and its Judgement does (Crusader's Judgement: more Holy damage taken for 40 s, refreshed by your melee, per spelltips 20308) and that Judgement does not consume the Seal (spelltips 20271), so a Crusader opener costs one extra Seal and one Judgement. retribution.md:119 unchanged. The Crusader-opener measurement (sim request 1) stays a sim-round item; if it gains beyond error, :93 and :119 take the auditor's text and the bare APL adds the prepull Crusader.
- PALADIN-R3-7 fixed: retribution.md:66, Vengeance procs from any crit except Consecration's ticks (spelltips 20049 "non-periodic"; talents.go:154-160).
- PALADIN-R3-8 fixed: protection.md:51, block chance from Holy Shield and Redoubt; Shieldsworn fires Holy Shield's damage on Holy Strike while it is up and lets you block spells.
- PALADIN-R3-9 fixed: holy.md:70, Divine Precision is the point that opens Holy Power's row; Precision fills out the Protection points. Build kept (Decisions).
- Decision, Retribution race: Human kept, no change.
- Left: Consecration leveling rank under the level-gap curve (sim request 2) stays a sim-round item; retribution.md:95 and :121 unchanged.
- New SIM findings: none.

# Guide accuracy round 15: Druid

Auditor: druid persona. Tree: gr15/audit at 4600b7d82 (main after fourteen guide rounds and sim round 17; READ-ONLY). Client 1.60.1.70755, engine pin b2c4a2d05 (sim/enginever/version.go; source read from gr15/engine). Before starting I read the round 1 to 14 druid reports and fix logs, round4/druid-levels.md, the race-sims files, ledger.md and the sim round 17 druid report (design/reviews/2026-11-04-sim-audit-round17/druid.md).

Specs audited: class page (`index.md`), Balance (`balance.md`), Feral cat (`feral.md`), Feral bear (`feral-bear.md`), Restoration (`restoration.md`), and the card notes in all four druid APLs. I checked the live best in slot rows (band 60, bare and raid, both factions) against every gear, professions and stat line.

## What moved since round 14 (d7fc51fe7 to 4600b7d82)

- feral.md:105, Professions (R14-3 fix and the rendered-review MUST).
- restoration.md:101, Professions (R14-2 fix).
- restoration.md:117, the Levels 20 to 60 pointer to the Ranks paragraph (sim round 17 S7: d479bbf99, 59affb460).
- data/curated/apl/druid-restoration.json: the priorityList[8] note and the top-level notes (R14-1 and sim round 17 R5).
- Engine pin f637ec32c to b2c4a2d05. Sim round 17 found 0 druid sim bugs and changed no druid action, rank row or build.

These did not change for any druid spec:
- the rank-choices.json druid rows (the hunks in that file are hunter and shaman)
- addon-builds.json druid orders
- talents/druid.json, trainables, spelltips, races.json
- bis/druid-*.json

## Decoded builds (unchanged; all legal, 51 points)

- Balance leveling `5232220115101151/0/555301`, raid `5232220115501351/0/505003`.
- Feral leveling `050022/50212032123012123001/05501`, raid `050022/50230032023012125001/055`.
- Bear `0/5123033202310212551/055`.
- Restoration `05302001/0/555300515311321`.

Tier positions come from talents/druid.json `tier`/`column`. Both Primal Bite and Shifting Power sit at tier 6, so the bear page and the cat page can each call its own point "the capstone".

## Findings

| id | status | kind | file:line | claim | evidence | fix |
|---|---|---|---|---|---|---|
| DRUID-R15-1 | CONFIRMED | professions (cross-page, made by the R14-3 fix) | feral-bear.md:107 | "Skinning and Leatherworking, as for the cat: leather is your armor, and Leatherworking makes it." | Round 14 rewrote the cat's line (feral.md:105). It now opens "Enchanting is the one profession your raid set asks for" and says "The Leatherworking pieces in your set bind when equipped, so buy them". So the bear's "as for the cat" points at a page that recommends Enchanting and says not to train Leatherworking for gear. The bear line gives the opposite reason: Leatherworking because it "makes" your armor. That is also against round 14's decision 2: a craft that binds when equipped is bought, not trained for. The bear's own live set, band 60 raid and bare on both factions (bis/druid-feral-bear.json, items/druid.json bonding), has four crafts: Outlaw's Collar 279253 (LW), Warbear Harness 15064 (LW), Hivethrasher's Girdle 275614 (LW) and Fel Cape 279269 (Tailoring), the cape only in the raid rows. All four bind when equipped. No bind-on-pickup craft is in any bear row, so no profession is needed for the set. | "Nothing in your set binds on pickup, so no profession is needed for your gear: the Leatherworking and Tailoring pieces bind when equipped and can be bought. Skinning and Leatherworking pay for themselves, and Alchemy keeps you in Stoneshield and Mighty Rage Potions." |

No LIKELY or NEEDS-SIM findings.

## Checks

**Round 14 fixes re-checked**
- R14-1, druid-restoration.json [8] note: "rank 2 from 20 to 37, rank 4 from 38 to 49, rank 6 from 50 to 52 and rank 8 at 58 and 59". It matches the four authored-9841 rows in rank-choices.json. It also matches restoration.md:81 and :117. The top-level notes say the same (not rendered).
- R14-2, restoration.md:101, against both raid rows of bis/druid-restoration.json:
  - Bind on pickup: Dormant Heart of the Mountain 249473 (Enchanting) and Mender's Leather Gauntlets 252551 (Leatherworking).
  - Every other craft binds when equipped: Living Crown, Mooncloth Shoulders, Mooncloth Vest, Wisdom of the Timbermaw, Elderwild Pants, Mooncloth Boots and Idol of Swiftness.
  - "Every other craft ... binds when equipped" holds.
- R14-3, feral.md:105, against the raid rows:
  - The Alliance row wears Molten Heart 249470 and Frozen Heart 249469.
  - The Horde row wears Molten Heart and Rune of the Guard Captain.
  - Both hearts are Enchanting crafts that bind on pickup. The Leatherworking pieces bind when equipped: Outlaw's Collar, Timbermaw Tunic, Might of the Timbermaw, and Warbear Woolies on the Horde.
  - The sentence holds as worded.
- restoration.md:117 (sim round 17 S7):
  - "from rank 2 at 20 to rank 4 at 38" matches the rows.
  - "Rank 1 stays on the party throughout" matches the 8910 row at 10 to 59, and the link target exists.
  - It does not contradict :81.

**Rank prose at the new pin**
- `core.LevelPenalty` at b2c4a2d05 is unchanged: MaxLevel is learn + 5, then a 12-level grace and 5% a level.
- Rejuvenation ranks come from trainables: r7 at 40, r8 at 46, r9 at 52, r10 at 58, and r11 at 60, which is the Ahn'Qiraj book.
- At 60 that gives rank 8 the full bonus, rank 7 85% and rank 1 0. restoration.md:81 holds.
- balance.md:135 matches the three druid-balance rows. Moonfire ranks are r2 10, r3 16, r4 22. Wrath ranks are r3 14, r4 22, r5 30. Insect Swarm r2 is at 30.

**Leveling against addon-builds.json** (decoded through talents/druid.json tier and column, index 0 = level 10)
- Balance: Improved Wrath 10-14, Genesis 15-16, Moonglow 17-19, Insect Swarm 27, Moonkin Form 40 in both builds. Matches balance.md:127 and :135.
- Feral leveling:
  - Ferocity 10-14, Feral Swiftness 15-16, Feral Instinct 17, Brutal Impact 18-19.
  - Shredding Attacks 20-22, Feral Charge 25, Shifting Power 40.
  - Then Genesis 41-45, Nature's Majesty 46-47, Nature's Reach 48-49, Heart of the Wild 50-54, Naturalist 55-59, Natural Shapeshifter 60.
  - Matches feral.md:55, :113 and :119 ("31 Feral points are done at 40" is 31 levels from 10 to 40).
- Bear: Ferocity 10-14, Furor 15, Feral Swiftness 16-17, Feral Instinct 18-19, Lacerate 25 (after Savage Fury 20-21), Primal Bite 40. Matches feral-bear.md:115 and :121.
- Restoration: Heart of the Wild 10-14, Reflection 15-17, Nature's Focus 18-19 (28%, "more than a quarter"). Matches restoration.md:109.
- Trainer levels (trainables/druid.json) are all as stated:
  - Wrath 1/6/14, Moonfire 4, Entangling Roots 8, Nature's Grasp 10, Faerie Fire 18.
  - Healing Touch 1, Rejuvenation 4, Regrowth 12, Revive 12, Cure Poison 14.
  - Bear Form, Maul, Growl and Demoralizing Roar 10. Enrage 12, Bash 14, Swipe 16, Maul r2 18.
  - Cat Form 20, Rake and Ferocious Bite 24, Poised and Ravage 32, Pounce 36, Frenzied Regeneration 36.
  - Dire Bear Form 40, Innervate 40, Omen of Clarity 20.

**Client values** (talents/druid.json, spelltips/druid.json, 70755)
- Cat:
  - Shifting Power: 55% base mana on an 8 s cooldown, 40 energy.
  - Shred 60 (42 with Shredding Attacks 3). Claw 45 and Rake 40 (5 less with Ferocity).
  - Ravage and Pounce read "Must be prowling and behind the target".
  - Poised is 10 off at max energy, out of combat.
  - Rip 6 is 12 s. Sim round 17 S3 notes that the BiS Idol of the Dream adds a tick. No page states Rip's length, so no line is affected.
- Bear:
  - Enrage: 10 rage plus 20 over 10 s, and an armor cut for the duration.
  - Frenzied Regeneration: 10 rage a second, 1% health a rage, 10 s.
  - Barkskin: 20% for 15 s, 60 s cooldown.
  - Lacerate 15 rage (12 with Shredding Attacks). Primal Bite 20 rage (15 with Ferocity), 6 s cooldown. Maul 15 (10). Swipe 20 rage, 3 targets.
- Feral talents:
  - Savage Fury: Claw, Rake, Shred, Maul and Swipe +10%. The cat line lists its three and the bear line its two.
  - Feral Swiftness: 4% dodge and 30% Cat Form speed.
  - Feral Instinct: Swipe 30% plus the Prowl clause.
  - Thick Hide: base armor plus Defense above five times your level, multiplied by the form.
  - Natural Reaction: 5% and 5 rage. Blood Frenzy: 5 rage, and a combo point on builder crits.
  - Natural Instinct: +20% crit bonus. Rend and Tear: 15% on bleeding targets. Predatory Strikes: 150% of level.
  - Heart of the Wild: Intellect 10%, Bear Stamina 20%, Cat Strength 10%. Furor rank 1: 20%.
- Balance and Restoration talents:
  - Balance: Eclipse is 2 charges at 0.5 s, stores 4. Improved Starfire is -0.5 s at 5. Nature's Grace is 10% for 3 s on non-periodic spell crits. Vengeance is +100% to the bonus.
  - Restoration: Nature's Splendor is +3/+6/+2. Improved Regrowth is 50% at 5. Improved Tranquility cuts the cooldown 60%. Tranquility is 5 min, Wild Growth 3.0 s on a 6 s cooldown for 1050 mana. Swiftmend is 20% of base mana on 15 s.
  - Every page line that quotes these numbers matches.

**Card notes**: every note in the four APLs matches its action and the values above. Examples:
- Balance: [1]/[2] end of fight, [5] Eclipse.
- Feral: [5] Shifting Power at 60, [6]/[7] Rip under 2 s and Bite at 5 s, [8]/[9] Claw 40 against Shred 54/48.
- Bear: Frenzied Regeneration at 50% with 20 rage, Lacerate stacks and the 3 s refresh, Demoralizing Roar 204, Maul at 15, Swipe above 40.
- Restoration: Innervate under 30%, Nature's Swiftness under 35%, Swiftmend under 75%, Healing Touch under 55%, party Rejuvenation under 70%.

**Best in slot and stats** (band 60, both presets, both factions; unchanged since round 14)
- Every feral and bear row wears The Unstoppable Force (a two-hand mace). That fits "a high-DPS two-hander" and the bear's weapon line.
- Wolfshead Helm 8345 is item level 45, needs level 40, and is in no row. The prose does not claim it is.
- Balance professions: the only bind-on-pickup craft is Frozen Heart (Enchanting), and Enchanting is named. The Tailoring and Leatherworking pieces bind when equipped. Not false (see Shoulds).
- Stat prose on all four pages explains and does not rank, and no reason contradicts the 70755 client or the tables.

**Races**: Elune's Light, Quickness, Endurance, War Stomp, Wind Blessed and Read Leyline (including the 60 min Attuned doubling, spelltips/shared.json) read as the pages say. The bear line matches round4/race-sims.md. Skyborne is named only as the paid option on every page.

**Sources**: every frontmatter URL returned 200 on 2026-11-04 (curl, browser user agent). That covers both Blizzard articles, Wowhead Skyborne First Look, talentsforever data.json, /sim/weights, /sim/drops, /bis/druid/restoration and /bis/druid/feral-bear. Labels match the pages, and all follow the round 4 source-link rule.

**Voice**: no em or en dashes. No site, sim, APL, band or preset words in prose. No quoted nightly numbers.

## Shoulds (wording, not findings)

- feral.md:89: the stat paragraph never explains melee haste, the stat the frontmatter selects and the raid table puts first. Suggested: "Haste is faster paw swings, and every swing can proc Omen of Clarity."
- restoration.md:101: "Herbalism or Alchemy keeps you in mana potions" reads as if Herbalism makes potions. Better: "Herbalism and Alchemy keep you in mana potions and elixirs".
- Carried from round 14:
  - feral.md:67, say what the raid build's Feral Instinct points do (Prowl detection on the walk in to your Ravage).
  - feral.md:73, "let it run to its last tick" should read "let it run into its last two seconds".
  - restoration.md:81, drop "(not 3)".
  - balance.md:119, name Frozen Heart of the Mountain as the bind-on-pickup reason for Enchanting.
  - balance.md:105, explain spell haste.
- Carried from rounds 12 and 13:
  - balance.md:105, "a resisted Wrath banks no Eclipse charges" (sim round 17 L1 keeps this open).
  - feral.md:77, "(55 with Wolfshead Helm)".
  - feral.md:85, "your crits".
  - balance.md:91, "almost always".

## Decisions

1. Round 14's three fixes and sim round 17's Restoration pointer sentence are right and stay.
2. The cat and bear profession lines must give the same reason. A craft that binds when equipped is bought (round 14 decision 2), so the bear line no longer sends you to Leatherworking for your armor (R15-1).
3. Rounds 1 to 14 decisions stand. These include:
   - Swiftmend keeps the HoT.
   - Gift of Nature misses Tranquility.
   - Pounce and Ravage need you behind the target.
   - Bear Furor is 1 point.
   - Feral Charge (Cat) is the 30 s leap.
   - Nature's Reach covers melee hit.
   - Rank 11 Rejuvenation and Healing Touch are Ahn'Qiraj book ranks.

## Sims requested

None new. Carried: druid-feral Bite energy pooling under Predatory Discipline.

## Sim findings

- Carried, open for a launch log:
  - Sim round 17 L1, Eclipse on a resisted Wrath.
  - S1/S2, the Cat Form spread and the unarmed paw.
  - X2, Read Leyline stacking with Innervate.
- Carried for the BiS lane: Balance raid trinket 2 is Uther's Strength on the Alliance (sim round 16 L3). The BiS talent strings are not the shipped builds (known, post-launch re-rank).

## Sources

- Client 1.60.1.70755:
  - talents/druid.json (texts, tier and column)
  - trainables/druid.json
  - spelltips/druid.json and spelltips/shared.json (racials)
  - races.json
  - items.json (8345, 16608, 220606, 260816)
  - items/druid.json (bonding of every crafted BiS piece named above)
  - addon-builds.json
  - bis/druid-balance.json, druid-feral.json, druid-feral-bear.json, druid-restoration.json (band 60, bare and raid, both factions)
- data/curated/rank-choices.json and data/curated/apl/druid-*.json.
- Engine b2c4a2d05: sim/core/spell.go `LevelPenalty` and `LevelPenaltyFromMaxLevel`.
- design/reviews/2026-11-04-sim-audit-round17/druid.md (S3, S5, S7, L1, R5); guide-accuracy round4/race-sims.md; round14/druid.md.
- Frontmatter URLs fetched 2026-11-04 (all 200):
  - https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid
  - https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
  - https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
  - https://talentsforever.com/data.json
  - https://foreversixty.gg/sim/weights
  - https://foreversixty.gg/sim/drops
  - https://foreversixty.gg/bis/druid/restoration
  - https://foreversixty.gg/bis/druid/feral-bear

## Fix log (round 15)

- DRUID-R15-1 fixed, feral-bear.md:107. Checked the bonding: Outlaw's Collar, Warbear Harness, Hivethrasher's Girdle and Fel Cape all bind when equipped (items/druid.json).

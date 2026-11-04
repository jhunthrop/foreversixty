# Guide accuracy round 14: Paladin

Auditor: paladin (class index, Holy, Protection, Retribution). Audit tree gr14/audit (main d7fc51fe7: thirteen guide rounds, sim round 16, launch state). Client 1.60.1.70755, engine pin f637ec32c (sim/enginever/version.go; source read from gr14/engine). Read-only: no guide edits, no sims run.

Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, `data/curated/apl/paladin-*.json` (card notes), `data/curated/rank-choices.json`, `data/builds/1.60.1.70755/addon-builds.json`, the live `bis/paladin-*.json` (de044125c picks), `races.json`, `spelltips/shared.json`. Before starting I read the round 13 report and fix log, the ledger, round 2 to 5 race-sims.md, and the sim round 16 paladin report (design/reviews/2026-11-04-sim-audit-round16/paladin.md).

## What changed since round 13's audit (git diff 604b78a80..d7fc51fe7)

- **Guides:** retribution.md:60 (the R13-1 fix) and two sim-lane edits, retribution.md:95 and :133 (Consecration ranks on the MaxLevel + 12 rule, sim round 16 W3-2). Holy, Protection and the class page are unchanged.
- **rank-choices.json:** sim round 16 re-cut the Holy Light, Flash of Light, Retribution Consecration and Protection Holy Strike rows level by level.
- **Card notes, addon-builds.json (paladin orders), BiS files:** unchanged. The BiS files are de044125c, which is what ships.

## Builds decoded (unchanged, all legal on 70755)

- **Retribution** `build` 050250002/502/0523330100130135 (14/7/30), `raidBuild` 050050002/503/05233301001301351 (12/8/31).
- **Protection** `build` 05005/15302102513501031/5003 (10/33/8).
- **Holy** `build` 50320215012101251/50322/0503 (31/12/8).

Leveling orders (addon-builds.json, index 0 = level 10) are the ones round 13 decoded; retribution.md:74/:125, protection.md:133/:139 and holy.md:117/:123 still match them.

## The live tables each page renders (band 60, raid preset, first recommended race's faction)

- **Retribution** (alliance human): melee haste (not on items), Strength, crit, spell power, Intellect, Agility, attack power; Hit "Not significant", hit cap "reached" (baseline 12.9%). BiS crafted rows: Lionheart Helm, Champion's Legplates, Depleted Thorium Sabatons (Blacksmithing, bind on equip), Timbermaw Tunic (Leatherworking), Frozen Heart of the Mountain (Enchanting, **bind on pickup**, in every band 50 and 60 row, both factions, both presets). Bare 50/60 rows also carry Libram of Invocation (Enchanting, bind on pickup).
- **Protection** (alliance dwarf): Expertise, Defense, Stamina, Dodge, Agility, block value, Parry, then Strength, crit, armor, attack power; Block and Hit "Not significant". BiS crafted rows are all Blacksmithing (Enchanted Thorium Helm, Heavy Thorium Gauntlets, Depleted Thorium Sabatons, Greenhammer, Sentinel's Wristguards bind on pickup).
- **Holy** (alliance human): spell haste (not on items), MP5, healing power, Intellect, crit, Spirit. BiS: Dormant Heart of the Mountain (Enchanting, **bind on pickup**) in every band 60 row; Blacksmithing, Leatherworking and Tailoring pieces are bind on equip.

## Checked and correct (no finding)

- **R13-1 fix (retribution.md:60):** holds. Benediction's class mask does not cover Hammer of Wrath; Holy Conduit 2/2 still discounts it (:72).
- **Stat prose against the live tables:** every reason is true on 70755 and none ranks. Retribution: hit "matters up to the cap and nothing past it" agrees with the raid table (cap reached, Hit not significant); Seal of Righteousness scales with weapon speed and spell damage (R1, 0.1 coefficient); Champion of the Light 60% of Intellect. Protection: "More block chance adds nothing you can measure at 60" agrees with Block "Not significant"; Expertise reaches Holy Strike (holy_strike.go melee special table). Holy: "Spell damage does nothing for a heal" (no spell power row); Reverence 30%, Illumination 50% Spirit regeneration for 15 s.
- **Race lines against races.json and spelltips/shared.json:** Ruthless 1% crit, The Human Spirit 10% Spirit, Will to Survive (stuns), Stoneform 10% physical for 8 s on 3 min, Artisan's Precision 1% hit, Big Game Hunter 5% on Beasts, Will of the Forsaken (no immunity), Cannibalize 7% health and mana, Touch of the Grave at most once every 15 s. combos.json: paladin is Human, Dwarf or Undead (Undead new), no Skyborne (index.md:39). No paid race named. Recommended races match the prose picks on every page. "Dwarf matches Human on the raid build in a raid" (retribution.md:111) agrees with the last measurement on this gear and build (round 5 race-sims: Dwarf 783.2 against Human 780.5); see Sim findings for why it needs a re-run.
- **Gear and professions lines against both factions' BiS rows:** Retribution and Protection "Mining and Blacksmithing" agree with the Blacksmithing plate, weapons and shield in the rows. No prose names a BiS piece. Battleworn Hammer, Major Mana Potion, Demonic Rune, Flask of Distilled Wisdom and Elixir of Holy Power (21546, spell 1310077) exist on 70755.
- **New rank rows against prose:** retribution.md:95 ("rank 1 ... for most of your leveling; only from about 44 to 49 is rank 2 worth its cost"; rank 4 bare and rank 5 raid at 60) matches the rows 30-43 rank 1, 44-49 rank 2, 50-59 rank 1, 60 bare rank 4. holy.md:78/:81/:95/:127 speak only to level 60 (Flash of Light rank 5 learned at 50 keeps its whole bonus at 60; rank 4 at 42 keeps 95%; Holy Shock rank 1 at 30 keeps 35%), so the new Holy rows below 60 change nothing on the page. Protection names no leveling rank.
- **Spell numbers re-derived at 60:** Light's Vigil 1,340 against ten Flash of Light rank 6 at 140 (1,400); Holy Shock rank 1 160 mana against 275/325; Holy Light rank 8 (10329, 54) top trainer rank, rank 9 (25292) no known source (round 1 decision); Consecration ranks 135/235/320/435/565 mana, 48/80/136/224/312 to each of the first four targets.
- **Levels 1 to 19** on all three pages: unchanged since round 13 and still right (trainer levels, Benediction "all instant cast spells and abilities", Spiritual Focus 70%, Divine Intellect 2%, Redoubt 30%/10 s/5 blocks, Sacred Duty rank 1 30 s).
- **Source links:** all nine distinct URLs return 200 (curl, 2026-11-04) and their labels match.
- **Voice:** no em or en dashes, no site or sim talk, no stat ranking, no nightly numbers.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R14-1 | LIKELY (contradiction) | leveling | retribution.md:133 | "From your mid to late 40s the newest rank wins on mana too, because rank 1 has lost about half of its spell damage bonus by then. On a long fight against one target your mana runs out first: keep rank 1 there into your early 40s, take rank 2 from about 44, and drop back to rank 1 through your 50s." | The two sentences give opposite answers to the same question. If the newest rank does more damage per mana from the late 40s (true: at 50 with rank 1 at 35% of its bonus, rank 4 does about 0.71 damage per mana to rank 1's 0.57, from the 70755 costs and amounts above and core.LevelPenalty), then a fight where "your mana runs out first" should take the newest rank, not rank 1. The rows (sim round 16 W3-2) pick rank 1 at 50-59 because over a paced single-target fight the cheap rank goes out on every cooldown while the big one waits on the mana rule; at 53 to 59 the two are within 0.5%. The stated reason ("mana runs out first") is not why rank 1 wins, and it contradicts the sentence before it. | Replace the last sentence with: "Against one target Consecration is filler, and a cheap rank you can put down every time it is ready beats a big one you have to hold back: keep rank 1 into your early 40s, take rank 2 from about 44, and go back to rank 1 through your 50s." |

## SHOULDs (wording, not findings)

- **retribution.md:115 (Professions):** every band 50 and 60 BiS row, both factions and both presets, lists Frozen Heart of the Mountain, a bind-on-pickup Enchanting trinket, and the bare rows add Libram of Invocation, also bind on pickup from Enchanting. The section never names Enchanting and ends "Pick two that pay for your consumables", so a reader cannot wear two of the table's picks. Nothing it says is false (in the raid row the trinket's only paladin stat is hit, which the table scores at zero), so this is a SHOULD. Suggested sentence after the Engineering one: "Enchanting is the only way to the Frozen Heart of the Mountain and the Libram of Invocation, both bind on pickup; the trinket's hit and the libram's cheaper Seals help most before you are hit capped."
- **holy.md:107 (Professions):** the same for Holy. Dormant Heart of the Mountain (bind on pickup, Enchanting) is in every band 60 row; its Holy Light proc is for you. Suggested: "Enchanting gives you the healing enchants yourself, and it is the only way to the Dormant Heart of the Mountain, a bind-on-pickup trinket that turns your Holy Lights into more healing."
- Carried: retribution.md:133 "Rank 1 at 135 mana carries as much spell damage as any rank" is exact only to level 37.
- Carried: the Retribution Holy Strike card note's "Twenty mana" against Benediction's 18.

## For the component and data lanes (not guide findings)

- **Retribution raid BiS trinket2:** Frozen Heart of the Mountain's only paladin stat is 9 hit (its spell damage is Frost and Shadow), and the same raid band scores hit at zero with the cap reached. A zero-value bind-on-pickup trinket heads the slot for both factions; the ranker should prefer a sourced trinket with any significant stat, or show the slot as open.
- Carried: the Protection APL top entry's "Lay on Hands fires itself below 10% health" note (does not render), the Holy APL "no downrank penalty" note and the Protection Righteous Fury note (neither renders), and the trainables rows for Holy Light 9, Blessing of Might 7 and Blessing of Wisdom 6.

## Decisions

- **R14-1:** decided from the 70755 Consecration costs and amounts, core.LevelPenalty at f637ec32c and the sim round 16 W3-2 measurements. LIKELY rather than CONFIRMED because each sentence alone is defensible; together they contradict.
- **Retribution Dwarf line kept:** the round 14 rule says a race line must not contradict the page's own tables. On paper it now does (Dwarf's only boss racial is 1% hit, and the page's raid table says hit past the reached cap adds nothing, while Human's Ruthless adds crit, about 0.76% DPS at the table's crit value), but the last measurement on this exact gear and raid build says Dwarf matches Human. A measurement beats an inference from a weight, so the line stands and the disagreement goes to the sim lane (Sims requested 1, Sim findings).
- **Professions omissions:** SHOULDs, not findings, because no sentence is false.
- **Races:** unchanged. Retribution Human (Undead), Holy Human (Undead), Protection Dwarf (Undead). No paid race named.

## Sims requested

1. **Retribution Human against Dwarf, raid 60, guide raid build, live BiS (de044125c alliance raid row), engine f637ec32c, 10,000 iterations, seeds 1 and 50001.** If Human now leads beyond error, retribution.md:111 becomes: "**Human** is the Alliance pick, in a raid and solo. Raid gear takes you past the hit cap, so Dwarf's Artisan's Precision adds nothing there, while Ruthless's crit always counts." If Dwarf still matches, the weights' hit row needs the engine check below.
2. Same comparison on the bare 60 row (expect Human ahead, as the page says).
3. Carried: Retribution Human against Dwarf on a Beast target; Protection threat by ability on 1 and 4 targets; the Crusader opener on a 15 s single mob at bands 20, 30 and 40.

## Sim findings

- **The Retribution raid weights and the race measurement disagree.** On the alliance raid row the weights give hit exactly 0 (hit cap reached, baseline 12.9%), yet round 5 measured Dwarf (+1% hit) 0.35% ahead of Human (+1% crit) on the same gear and build, while the same row prices 1% crit at about 0.76% DPS. One of the two is wrong: either some Retribution damage (Seal procs, Judgement, Consecration, Exorcism) still rolls against a higher spell hit cap that the weights' hit step does not reach, or Ruthless or Artisan's Precision is applied wrongly. Check which spells Artisan's Precision reaches in the engine and re-run the hit weight with the step taken on spell hit as well.

## Sources

- 70755 client data, data/builds/1.60.1.70755: bis/paladin-{holy,protection,retribution}.json (band 60 weights, hit_to_cap, slots), items/paladin.json (bonding: 249469, 249442, 249473 bind on pickup), spelltips/paladin.json (Consecration 26573/20116/20922/20923/20924, Flash of Light, Holy Light, Holy Shock), spellranks.json, talents/paladin.json, spelltips/shared.json (racials), races.json, combos.json, addon-builds.json, simconsumes.json (21546).
- data/curated/rank-choices.json (paladin rows); data/curated/apl/paladin-*.json.
- Engine f637ec32c: sim/core/spell.go:777-852 (SpellLevelPenalty, MaxLevel + 12, 5% a level).
- design/reviews/2026-11-04-sim-audit-round16/paladin.md (W3-1, W3-2, W3-3); design/reviews/guide-accuracy/round5/race-sims.md (Retribution Human against Dwarf); round4/race-sims.md (Protection, Holy).
- web/src/lib/guides/class-dps.ts and web/src/lib/bis/presets.ts (the stat table reads band 60, raid preset, the first recommended race's faction).
- Links fetched 2026-11-04, all 200: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap ; https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap ; https://www.wowhead.com/forever/news/undead-paladins-coming-to-world-of-warcraft-forever-382823 ; https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/ ; https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/ ; https://foreversixty.gg/planner?class=paladin ; https://foreversixty.gg/sim/weights ; https://foreversixty.gg/data/1.60.1.70755/races.json ; https://foreversixty.gg/data/1.60.1.70755/combos.json

## Counts

| status | count |
|---|---|
| CONFIRMED | 0 |
| LIKELY | 1 (R14-1) |
| NEEDS-SIM | 1 new (Retribution raid race), 3 carried |
| NEEDS-OWNER | 0 |

Round 13 edit: 0 errors. Sim round 16 guide edits: 1 likely (R14-1). Levels 1 to 19: 0 errors. Race lines: 0 errors (1 sim request). Gear and professions lines: 0 errors (2 SHOULDs). Source links: 0 errors. Card notes: 0 rendered errors.

## Fix log (round 14)

- PALADIN-R14-1 (likely) fixed with the report wording (retribution.md:133).
- SHOULD applied: Retribution and Holy professions now name the bind-on-pickup Enchanting trinkets (Frozen Heart, Dormant Heart), both present in the live BiS rows.

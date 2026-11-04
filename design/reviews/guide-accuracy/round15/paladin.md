# Guide accuracy round 15: Paladin

Auditor: paladin (class index, Holy, Protection, Retribution). Audit tree gr15/audit (main 4600b7d82: fourteen guide rounds, sim round 17, launch state). Client 1.60.1.70755, engine pin b2c4a2d05 (sim/enginever/version.go; source read from gr15/engine). Read-only: no guide edits, no sims run.

Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, the card notes in `data/curated/apl/paladin-*.json`, `data/curated/rank-choices.json`, `data/builds/1.60.1.70755/addon-builds.json`, the live `bis/paladin-*.json` (de044125c picks), `spelltips/{paladin,shared}.json`, `spellconst/paladin.json`, `talents/paladin.json`, `trainables/paladin.json`, `combos.json`, `races.json`. Before starting I read the round 14 report and fix log, the round 13 report, the ledger, the race-sims files of rounds 3 to 5 and the sim round 17 paladin report (design/reviews/2026-11-04-sim-audit-round17/paladin.md).

## What changed since round 14's audit (git diff d7fc51fe7..4600b7d82)

- **Guides:** only the round 14 fixes: retribution.md:115 (Enchanting and Frozen Heart of the Mountain), retribution.md:133 (R14-1, single-target Consecration ranks) and holy.md:107 (Enchanting and Dormant Heart of the Mountain).
- **Sim round 17:** zero paladin changes (no engine, APL, build string, rank row or guide text). Its rank-choices.json and addon-builds.json edits are hunter, rogue and warlock rows only. The paladin BiS files, stat weights and card notes are unchanged.

## Builds decoded (unchanged, all legal on 70755)

- **Retribution** `build` 050250002/502/0523330100130135 (14/7/30), `raidBuild` 050050002/503/05233301001301351 (12/8/31). The raid order adds Unsealed Wrath at 40 and moves Reverence's 2 points into Precision 3 and Unsealed Wrath (retribution.md:58).
- **Protection** `build` 05005/15302102513501031/5003 (10/33/8).
- **Holy** `build` 50320215012101251/50322/0503 (31/12/8).

Leveling orders (addon-builds.json, index 0 = level 10), decoded by name, match every leveling line:
- Retribution: Benediction 10-14, Improved Judgement 15-16, Conviction 17-19, Sanctified Judgement 20-22, Vindication 23-25, Seal of Command 26, Sacred Arbiter 27, Two-Handed Weapon Specialization 28-30, Repentance 31, Vengeance 32-34, Champion of the Light 35-39 (retribution.md:74, :125).
- Protection: Divine Intellect 10, Redoubt 11-15, Precision 16-18, Sacred Duty 19, Templar's Bulwark 20, Sacred Duty 21, ... Holy Shield 30, ... Shieldsworn 40 (protection.md:133, :139).
- Holy: Spiritual Focus 10-14, Healing Light 15-17, Reverence 18-19, Illumination 20, Divine Favor 21, Illumination 22-24, Oath of Light 25 (the 16th point, the first row 4 allows; Divine Precision and Infusion of Light share that row), ... Holy Power 35-39, Light's Vigil 40 (holy.md:117, :123). Divine Precision is one of exactly 25 Holy points under Holy Power's row, so holy.md:67 holds.

## Round 14 fixes re-checked

- **retribution.md:133 (R14-1):** holds. "Against one target Consecration is filler, and a cheap rank you can put down every time it is ready beats a big one you have to hold back" no longer contradicts the sentence before it, and the ranks named (rank 1 to the early 40s, rank 2 from about 44, rank 1 through the 50s) are the rank-choices.json rows (30-43 rank 1, 44-49 rank 2, 50-59 rank 1). Rank 2 at 44 is the middle of three ranks known then (rank 3 at 40), so "cheap" still fits.
- **retribution.md:115:** holds. Frozen Heart of the Mountain is Enchanting, bind on pickup (items/paladin.json `bonding: pickup`), in every band 50 and 60 row of both factions and both presets, so only an enchanter can wear it.
- **holy.md:107:** holds. Dormant Heart of the Mountain is Enchanting, bind on pickup, in every band 60 row of both factions and presets.

## Checked and correct (no finding)

- **Talent lines** on all three pages against talents/paladin.json at the page's rank, including prerequisites (Vengeance needs Seal of Command, Unsealed Wrath needs Repentance, Swift Judgement needs Templar's Bulwark, Shieldsworn needs Holy Shield, Shield Specialization needs Redoubt 5, Illumination needs Reverence 2). Holy Power 175% (50% more of the heal crit bonus). Shieldsworn's spell block for 6 s and its uncharged Holy Shield and Shield Specialization triggers.
- **Trainer levels** in every Levels 1 to 19 paragraph and in the index (trainables/paladin.json): Devotion Aura, Holy Light, Seal of Righteousness 1; Judgement, Blessing of Might 4; Holy Strike, Holy Light 2, Divine Protection 6; Hammer of Justice, Purify 8; Lay on Hands, Righteous Fury, Seal of Fury 10; Redemption 12; Holy Light 3, Blessing of Wisdom 14; Consecration, Exorcism, Flash of Light, Blessing of Kings 20; Holy Shield ranks 30/40/50/60; Hammer of the Righteous 40; Divine Shield 34 (so Sacred Duty at 19 reaches only Divine Protection, protection.md:133); Hammer of Wrath 44; Holy Wrath 50.
- **Spell values (spelltips, spellconst 70755):** Judgement and Swift Judgement off the global cooldown (gcd 0), Judgement 6% base mana and 10 yd; Hammer of the Righteous 6% base mana, 6 s, target plus 3; Seal of Fury absorb with a shield and a 3 s taunt; Righteous Fury 30 min, 30% base mana, 33% and a further 90% on Holy; Holy Shield 10 s cooldown and duration; Light's Vigil 1,340 at rank 3 against ten Flash of Light rank 6 at 140; Holy Shock rank 1 160 against 275/325; Lay on Hands 20 min, 550 mana at rank 3, drains your mana; Divine Shield rank 2 12 s at half damage; Voice of Truth 6 s, 3 min; Repentance 6 s, 1 min; Seal of Command 210 and Seal of Righteousness 200 mana, about 190 after Benediction (card note "about 190 mana").
- **Coefficients behind consumable and Seal claims (spellconst):** Seal of Command proc 20424 carries 0.29, Judgements 0.429, Holy Strike 0.429, Consecration ticks 0.042, so "Elixir of Holy Power reaches every Seal proc, Judgement and Consecration tick" (retribution.md:115) holds on the engine's model (see Sim findings for Seal of Righteousness rank 8).
- **Racials (spelltips/shared.json, engine racials.go):** Ruthless, The Human Spirit, Will to Survive, Stoneform 10% physical for 8 s on 3 min on the global cooldown, Artisan's Precision, Big Game Hunter, Will of the Forsaken (no immunity), Cannibalize 7% health and mana, Touch of the Grave 20% chance with a 15 s internal cooldown. combos.json: Human, Dwarf, Undead (Undead new), no Skyborne. No paid race is named anywhere.
- **Race lines against the race table:** the RaceTable component lists races and racials only; no prose pick contradicts it. Recommended races match the prose picks on every page.
- **Gear and professions against both factions' BiS rows:** Mining and Blacksmithing for Retribution (Lionheart Helm, Champion's Legplates, Depleted Thorium Sabatons, Forest Defender's Axe, Thorium Greatmace) and Protection (Enchanted Thorium Helm, Heavy Thorium Gauntlets, Greenhammer, Evergreen Shield, Sentinel's pieces) agree. No prose names a nightly BiS piece for a moving reason.
- **Stat prose against the live tables:** every reason is true and none ranks. Protection Expertise "which only a few raid sets carry" (settled round 10). Retribution hit "up to the cap and nothing past it" agrees with the raid row (hit 0, insignificant).
- **APL lines against prose and card notes:** Holy pacing "-0.05" (above 45% halfway), 40%/70%/90%/80% thresholds, Flash of Light ranks 5 and 4; Retribution Consecration pacing (0.8 of the fight left, Seal + Judgement + Holy Strike reserve), twist gate (3 stacks, 30% mana, 1,200 attack power); Protection Seal of Fury in its last 2 s with Holy Shield's mana kept, Bulwark below 35%, Divine Shield 25%, Divine Protection 20%, Stoneform 70%, Consecration above half mana.
- **Source links:** all nine distinct URLs return 200 (curl, 2026-11-04) and their labels match.
- **Voice:** no em or en dashes, no site or sim talk, no stat ranking, no nightly numbers.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R15-1 | CONFIRMED | consumable | protection.md:123 | "With Alchemy, Elixir of Holy Power adds to every Holy threat button you have." | Elixir of Holy Power (21546, spell 1310077) is "Increases spell Holy damage by up to 40" (spelltips/shared.json; aura 13), so it reaches a spell only through its spell power coefficient. Hammer of the Righteous (407632) has one damage effect, `amount 1, sp_coefficient 0.0` (spellconst/paladin.json), and the engine gives it no BonusCoefficient (hammer_of_the_righteous.go:72-81: flat 1 plus 3 times weapon DPS). The page itself calls Hammer of the Righteous the button "where your threat on a pack comes from" (protection.md:95), so "every Holy threat button" leaves out the biggest one. The elixir does reach Holy Shield (0.08), Holy Strike (0.429), Seal of Fury's proc (0.1), Judgement of Fury (0.429) and Consecration (0.042 a tick). | "With Alchemy, Elixir of Holy Power adds to Holy Shield, Holy Strike, Seal of Fury, your Judgements and Consecration. Hammer of the Righteous hits off your weapon, so the elixir does nothing for it." |
| PALADIN-R15-2 | NEEDS-SIM | talent | protection.md:82 | "Left out: Reckoning (extra attacks after a block or a crit, worth less than Conviction)" | The only measurement of Reckoning (sim round 16 P1) is against Anticipation, not Conviction: Reckoning 5 in place of Anticipation 5 gave +2.5% threat in a raid at both seeds, and at bare 60 it raised the chance of death from about 27% to about 48%; it was rejected on survival. Reckoning against Conviction was never run, so the stated reason is unsupported, and the measured reason (it buys threat with Defense) is not on the page. | Whatever Sims requested 1 shows, this is true now on the measured data: "Left out: Reckoning (extra attacks after a block or a crit; it adds a little threat, but its five points have to come out of Anticipation's Defense, and that trade gets you killed on hard pulls), Stoicism, Guardian's Favor and Improved Hammer of Justice." If the sim shows Reckoning beating Conviction 3 + Deflection 2, the build question goes to the sim lane instead. |

## SHOULDs (wording, not findings)

- holy.md:81: "Rank 1 costs half as much" is exact against rank 4 (160 against 325) and 58% of rank 3 (275). "about half as much" covers both ranks the sentence names.
- Carried: retribution.md:133 "Rank 1 at 135 mana carries as much spell damage as any rank" is exact only to level 37.
- Carried: the Retribution Holy Strike card note's "Twenty mana" against Benediction's 18 (round 3 decision: card notes quote tooltip base cost).

## For the component and data lanes (not guide findings)

- Carried from round 14: the Retribution raid trinket2 Frozen Heart of the Mountain (its only paladin stat is hit, scored at zero on the same row); the Protection APL "Lay on Hands fires itself below 10% health" note (does not render); the Holy APL "no downrank penalty" and Protection "+60 percent threat" engineering notes (do not render); trainables rows for Holy Light 9, Blessing of Might 7, Blessing of Wisdom 6.

## Decisions

- **R15-1:** decided from the client spellconst row of 407632 (sp_coefficient 0.0 on its only damage effect), the elixir's "up to 40" text and the engine. No owner question.
- **R15-2:** NEEDS-SIM rather than CONFIRMED because the Conviction comparison is unmeasured; the replacement text rests on the round 16 measurement and is true either way.
- **Oath of Light and low ranks (holy.md:33, :59, :78 and the card note at paladin-holy.json:584):** not filed. Sim round 17 O1/O4 leaves it OPEN behind the owner's in-game check, and the rows stand. The text changes only if that check shows the Oath penalized.
- **Retribution Dwarf line (retribution.md:111):** kept, as in round 14. The last measurement (round 5) has Dwarf level with Human on the raid build, and sim round 17 did not re-run it. Still requested below.
- **Exorcism never crits (retribution.md:62):** kept (round 1 beta-log decision; sim round 17 X1 still open for a level 60 log).
- **Races:** unchanged. Retribution Human (Undead), Holy Human (Undead), Protection Dwarf (Undead). No paid race named.

## Sims requested

1. **Protection talents, raid 60 and bare 60, alliance dwarf, 10,000 iterations, seeds 1 and 50001:** guide build against Reckoning 5 in place of Conviction 3 and Deflection 2 (Retribution tree Deflection 3 only; legal: Reckoning's row needs 20 Protection points, the build has 28 before it). Report threat, damage taken and chance of death. Decides PALADIN-R15-2's reason; the fix text above stands either way unless Reckoning wins on both threat and survival.
2. Carried from round 14: Retribution Human against Dwarf, raid 60, guide raid build, live BiS (de044125c alliance raid row), engine b2c4a2d05, 10,000 iterations, seeds 1 and 50001; and the same on the bare 60 row. If Human leads beyond error in the raid, retribution.md:111 becomes round 14's wording ("Human is the Alliance pick, in a raid and solo...").
3. Carried: Retribution Human against Dwarf on a Beast target; Protection threat by ability on 1 and 4 targets; the Crusader opener on a 15 s single mob at bands 20, 30 and 40.

## Sim findings

- **Seal of Righteousness rank 8's spell power coefficient.** spellconst/paladin.json (70755) gives the Seal's dummy aura effect `sp_coefficient 0.1` on ranks 1 to 7 (20154, 21084, 20287 to 20292) and `0.0` on rank 8 (20293); every proc id (25713, 25735 to 25742) carries 0.0, so the scripted scaling reads the aura's value. The engine uses 0.1 for every rank (sor.go:49-56, `coeff: 0.1` on 20293). If rank 8 really is 0 on the server, level 58+ Retribution loses spell power scaling on its main Seal, the band 60 spell power and Intellect weights are too high, and retribution.md:89 and :103 ("it scales with your weapon's speed and your spell damage") would need "until rank 8". Check the 70755 SpellEffect row for 20293 (EffectBonusCoefficient) against the Era row and, if it is 0 in both, against a level 60 log (SoR proc damage against the paladin's spell power). Seal of Fury's aura values (20419 at 0.9, 20423 at 0.0) do not matter: its proc ids (20415 to 20418) carry 0.1 themselves.
- Carried: the Retribution raid weights (hit 0) against the round 5 race measurement (Dwarf level with Human).

## Sources

- 70755 client data, data/builds/1.60.1.70755: spellconst/paladin.json (407632 effects; 20154-20293 and 25713-25742 coefficients; 20424; Holy Shield, Seal of Fury, Judgement of Fury, Consecration, Holy Strike; gcd_ms of 20271, 1310994, 1311015), spelltips/paladin.json, spelltips/shared.json (racials; elixir_of_holy_power 1310077), talents/paladin.json, trainables/paladin.json, addon-builds.json, bis/paladin-{holy,protection,retribution}.json (band 50 and 60 rows, both factions and presets), items/paladin.json (bonding), combos.json, races.json, simconsumes.json (21546).
- data/curated/apl/paladin-{holy,protection,retribution}.json; data/curated/rank-choices.json; web/src/components/guides/RaceTable.astro.
- Engine b2c4a2d05: sim/paladin/hammer_of_the_righteous.go:9-81, sor.go:11-56 and :112-130, healing_talents.go:29-34 (Illumination's 50% Spirit regeneration for 15 s), sim/core/racials.go:244-280 (Stoneform, global cooldown) and :379-381 (Touch of the Grave 15 s), sim/core/consumes.go:626.
- design/reviews/2026-11-04-sim-audit-round16/paladin.md (P1, Reckoning against Anticipation); design/reviews/2026-11-04-sim-audit-round17/paladin.md (O1, O4, X1); design/reviews/guide-accuracy/round5/race-sims.md.
- Links fetched 2026-11-04, all 200: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap ; https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap ; https://www.wowhead.com/forever/news/undead-paladins-coming-to-world-of-warcraft-forever-382823 ; https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/ ; https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/ ; https://foreversixty.gg/planner?class=paladin ; https://foreversixty.gg/sim/weights ; https://foreversixty.gg/data/1.60.1.70755/races.json ; https://foreversixty.gg/data/1.60.1.70755/combos.json

## Counts

| status | count |
|---|---|
| CONFIRMED | 1 (R15-1) |
| LIKELY | 0 |
| NEEDS-SIM | 1 new (R15-2), 3 carried |
| NEEDS-OWNER | 0 |

Round 14 fixes: 0 errors. Sim round 17 guide edits: none for paladin. Levels 1 to 19: 0 errors. Race lines: 0 errors. Gear and professions lines: 1 error (R15-1). Source links: 0 errors. Card notes: 0 rendered errors.

## Fix log (round 15)

- PALADIN-R15-1 fixed with the report wording (protection.md:123).
- PALADIN-R15-2 fixed with the measured reason (protection.md:82), true whatever the Reckoning-vs-Conviction sim shows.

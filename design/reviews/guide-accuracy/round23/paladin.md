# Guide accuracy round 23: Paladin

Audited on the round 23 audit tree (24ee457d0: main after guide round 22; the log has no sim round 21 commit and there is no `*sim-audit-round21*` report, so sim round 21 had not landed). Engine pin d5c708834 (`sim/enginever/version.go`; source read at `scratchpad/gr23/engine`). Active client 1.60.1.70790 under `data/builds/1.60.1.70790/`.

Files audited:
- `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, line by line.
- Every card note in `data/curated/apl/paladin-{holy,protection,retribution}.json`.
- Every rendered row of `addon-data.json` (`rotations/paladin-retribution` 18 bands, `paladin-protection` 18 bands, `paladin-holy` 16 bands).
- Both shipped point orders per spec in `addon-builds.json`.
- `talents/paladin.json`, `spelltips/paladin.json`, `trainables/paladin.json`, `data/curated/rank-choices.json`, `bis/paladin-*.json`, `items/paladin.json`, `races.json`, `combos.json`.

Read first: the round 22 paladin report and the ledger's Round 19 to 22 sections.

Lenses:
- Round 20's band-by-band note sweep, on both orders.
- Round 21's sweep of every superlative, comparison and quoted cost, duration and cooldown.
- Round 22's shapes: a spell or talent named in a leveling paragraph before the player has it; a later note clause naming a spell or rank before it trains; a "cheap" or "best" list containing its own top; a claim true only among one spell's ranks; "every" claims a talent covers only in part; rounded costs stated as exact.
- Round 23's addition: both sides of every comparison across all bands, including talents later in the order.

## Changes since round 22's tree (987d08a7c)

- d3320fa41 and 5cc46bbfa changed three lines, all prose:
  - holy.md:33, "every Holy Light, Flash of Light and Holy Shock on the tank" (R22-4).
  - holy.md:95, "so at 60 your Flash of Light is rank 4 or 5 and your Holy Shock rank 3 or 4, never rank 1" (R22-1, as reworded by the rendered review).
  - retribution.md:133, "it is the Consecration rank that gives you the most damage for your mana" and "From 31, Repentance" (R22-2, R22-3).
- Paladin `rotations` in addon-data.json, the three paladin APLs, `rank-choices.json` and the paladin orders in `addon-builds.json` are identical to round 22's tree (compared in code against `git show 987d08a7c:...`). All card notes are byte-identical.

## Findings

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R23-1 | CONFIRMED | card note (comparison fails on a band; round 20 shape 2, both sides not checked) | `paladin-holy.json` priorityList[3], renders at bands 20 to 24 | "Holy Light is the only heal that outheals a tank by itself." | At 20 to 24 the row casts Holy Light rank 1 (band 20: spell 635; `rank-choices.json` Holy Light 20-24 = rank 1; `addon-data.json`). Client: rank 1 heals 43 for 35 mana in 2.5 s (635). The other heal on the same band is Flash of Light rank 1, 46 for 35 mana in 1.5 s (19750). Rank 1's MaxLevel is 5 (`spell_max_level_auto_gen.go`), so it keeps 85% to 65% of its 0.714 at 20 to 24, while Flash of Light rank 1 keeps all of its 0.429. Worked heal sizes: level 22 at 0 / 10 / 25 / 50 healing power is Holy Light 43 / 48 / 56 / 70 against Flash of Light 46 / 50 / 57 / 67. So the row's Holy Light is no bigger than Flash of Light on a low-level character, on a slower cast, and the Oath that later pulls it clear does not exist until 25. From 25 the claim holds (rank 2 with the Oath's share: 233 against 94 at 25, and 218 against 102 at 29). Round 22's sweep checked 47 to 49 and missed this band. | "On the tank below 40%, right behind Divine Favor (from 21). At 60 rank 8 is your top rank: the trainer stops there, and the rank 9 book has no known source in Forever yet. From 25, Holy Light is the only heal that outheals a tank by itself." |
| PALADIN-R23-2 | CONFIRMED | page prose ("every" wider than the talent, and contradicts the page; side effect of R22-4's rewording) | holy.md:33 and holy.md:59 | :33 "Oath of Light adds a flat amount to every Holy Light, Flash of Light and Holy Shock on the tank". :59 "it adds a flat amount to every Holy Light, Flash of Light and Holy Shock you put on them." | The same page says otherwise at :78 ("adds the same flat amount to every rank close to your level ... the Oath's bonus goes with it") and :127 ("It has lost your healing power and the Oath's bonus with it"). Card rows 1 and 11 already carry the limit ("as long as the rank is close to your level"). The rendered rows cast ranks where the amount is not whole: Flash of Light rank 1 on the tank at 40 to 42 keeps 85% to 75%, Holy Light rank 4 at 47 to 49 keeps 60% to 50% (`rank-choices.json` H1 note), and Flash of Light rank 1 keeps none from 57. Round 22 listed this clause as a SHOULD. R22-4 then rewrote :33 to name the spells and kept "every", so the line now names exactly what the rank rule cuts. This is the brief's "every" shape, and :59 is the page's definition of the talent. See D1. | :33 "Oath of Light adds a flat amount to every Holy Light, Flash of Light and Holy Shock close to your level on the tank, so a Flash of Light a rank or two below..." (rest unchanged). :59 "...it adds a flat amount to every Holy Light, Flash of Light and Holy Shock you put on them, as long as the rank is close to your level. It took Blessing of Light's place, and it is the biggest point in the tree." |

Counts: CONFIRMED 2, LIKELY 0, NEEDS-SIM 0 new (2 carried, already routed), NEEDS-OWNER 0.

## SHOULDs (wording, not findings)

- **retribution.md:43** "Unsealed Wrath turns Seal swapping into real damage once you have raid buffs": the APL's own measurement is about +1.3% in the raid run. Not false, but "real damage" is generous; "a small gain" would match it. No nightly number belongs on the page, so wording only.
- **Holy row 12** "a small rank is the best balance": at 50 to 54 the party row casts rank 4 (151) while the tank row casts rank 3 (101), so the "small" rank is larger than the staple. Nothing stated is false, since rank 4 is still below the top rank 5 there. See the optional sim below.
- Carried, still true as SHOULDs:
  - Retribution row 6 "180 to 190 at 60" (exact: 180 and 189; a range containing both).
  - retribution.md:72 Holy Conduit omits Holy Wrath (incomplete, not false).
  - retribution.md:131 "Sanctified Judgement refunds most of each Seal" (3/3 only from 22). See D3.
  - retribution.md:121 "both come back every 10 seconds"; retribution.md:36 "Judge every 8 seconds"; retribution.md:93's repeated last sentence.
  - Holy row 12 "rank 1 for most of your leveling" (rank 2 renders at 44 to 49).
  - holy.md:81 "half as much" (rank 1 is 49% of rank 4, 58% of rank 3).
  - protection.md:111 block chance.
  - Holy Shield ranks 2 and 3 tooltips missing "4 charges" (data lane).

## Checked and right this round

- **Round 22's fixes.**
  - holy.md:95: rows cast Flash of Light rank 4 (party, bare tank) or 5 (raid tank) and Holy Shock rank 3 (50 to 55) or 4 (56 on) at the top of leveling and at 60. At 60 no row casts rank 1 of either. "Lose most or all of it": at 60 Flash of Light rank 1 keeps 0% and Holy Shock rank 1 keeps 35%. Consistent with :78, :81 and :127.
  - retribution.md:133, "the Consecration rank that gives you the most damage for your mana": recomputed per mana on both sides at each level. Per target over 8 s the damage is 8 x (tick + first-targets tick) + 8 x 0.042 x (spell damage + Crusader bonus) x level share. With Crusader on the target (92 at 32, 127 at 42) and about 50 to 70 spell damage: rank 1 against rank 3 is 0.66 against 0.57 at 40, 0.66 against 0.62 at 44, and 0.62 against 0.64 at 46. Rank 2 never passes rank 1 before 47. So "until your mid 40s" and "from your mid to late 40s" both hold. On a three-mob pack with one judged target it stays about level at 40 (1.58 against 1.53); not false.
  - retribution.md:133, "From 31, Repentance": index 21 on both orders.
  - holy.md:33, the three named spells match the talent text and :59 (the remaining "every" is R23-2).
- **Later-talent side effects (both sides of each comparison).**
  - Holy Conduit at 52 (leveling) and 51 (raid) cuts every Consecration rank by the same share, so the rank comparisons on retribution.md:95 and :133 do not move. "About 120 mana after Benediction" sits in the sentence about your 30s, before Holy Conduit.
  - Benediction at 53 to 57 on the Holy order discounts Holy Shock but not Flash of Light (1.5 s cast). The page quotes no per-mana comparison between them. holy.md:40's Light's Vigil has a 1.5 s cast, so no discount, 1,340 against ten at 140.
  - Guardian's Favor at 49 to 50 gives "Blessing of Protection every 3 minutes" only at 60's frame; holy.md:47 is in the raid section.
  - Protection has no Benediction or Holy Conduit: Holy Strike 20, Consecration 135 to 565.
  - protection.md:99 and row 12: Hammer of the Righteous per mana against rank 1 (40 to 58) and rank 5 (60), both sides, still far ahead.
- **Flash of Light rank comparison, holy.md:33 and :78, every band.** Rank below top against top with the Oath's 1.9167 a level and 0.429: at 42 a rank down heals 78%, at 50 about 77%, at 60 rank 5 heals 85% and rank 4 about 71%. Per mana rank 4 passes the top rank at about 50 healing power and rank 5 at about 100 (58 to 60, where the margin is small). With no healing power at all the top rank is level or slightly ahead, but a healer's 60 gear is far past both points, so the line holds as written.
- **Oath values.** The tooltip prints "up to 0" in `spelltips`; the engine's 6.667, 1.917 and 3.833 a level give 400, 115 and 230 at 60. The page quotes no number.
- **Quoted costs, durations and cooldowns.**
  - Retribution: Judgement 6% base mana; Holy Strike 18 at 60 with Benediction; Consecration rank 1 121.5; Seal of Righteousness rank 8 200, Seal of Command rank 5 210; Hammer of Wrath 6 s on all three ranks; Repentance 6 s, one minute; Crusader debuff 40 s.
  - Protection: Holy Strike 20; Righteous Fury 30% base mana for 30 minutes; Templar's Bulwark 8 s; Holy Shield 10 s and 4 charges; Divine Shield 10 and 12 s; Swift Judgement one minute, free; Stoneform 8 s, 3 minutes; Judgement of Fury's taunt 3 s from 10 yards.
  - Holy: Light's Vigil 1,340; Oath 5 minutes; Divine Favor 2 minutes; Holy Shock 10 s; Lay on Hands 20 minutes and 550 mana; Voice of Truth 6 s every 3 minutes; Infusion of Light 15 s and 1.0 s.
- **Superlatives and "only" claims.** holy.md:87 only instant heal outside Lay on Hands; holy.md:40 only party heal; protection.md:80 only parry talent; retribution.md:133 "about twice as hard" (1.9 to 2.2x at 40 across 100 to 200 effective spell damage, more at later bands); Holy row 3 at 25 on (R23-1 is 20 to 24).
- **Levels 1 to 19 and 20 to 60 paragraphs** against the orders and trainers: every level on all three pages matches. Oath of Light at 25 is the 16th Holy point; Sacred Duty rank 1 at 19 is 30 s; Templar's Bulwark at 20; Holy Shield 30/40/50/60; Hammer of the Righteous 40; Exorcism and Consecration 20; Holy Wrath 50; Redemption 12; Blessing of Wisdom 14.
- **Build lines.** Holy 31/12/8, Protection 10/33/8, Retribution 14/7/30, raid 12/8/31, and "moves Reverence's two points" all match the decoded orders and strings. index.md's one-point talents at 11, 16, 21 and 31 match all three trees. Prerequisites: Swift Judgement needs Templar's Bulwark; Vengeance needs Seal of Command; Unsealed Wrath needs Repentance.
- **Talent texts** re-read against every talent bullet on the three pages, and right: Spiritual Focus, Healing Light, Reverence, Illumination (20272's regeneration per the engine), Divine Favor, Infusion of Light, Boundless Light, Holy Power 175%, Redoubt, Shield Specialization, Iron Creed, Shieldsworn, Sacred Duty, Conviction, Vengeance, Two-Handed Weapon Specialization, Champion of the Light, Unsealed Wrath, Purifying Power, Guardian's Favor, Divine Precision.
- **Seal of Righteousness share:** 0.12 against 0.0765 is 57% more, "about 60%". Procs carry no level penalty, so "every rank takes the same share" holds.
- **Gear and professions against live rows.** Dormant Heart (Holy trinket1 at 60, all rows) and Frozen Heart (Retribution trinket2 at 50 and 60, all rows) are bind on pickup and Enchanting. Elixir of Holy Power 21546 exists at level 40.
- **Racials against `races.json`:** Ruthless, The Human Spirit, Will to Survive, Stoneform, Artisan's Precision, Big Game Hunter 5%, Will of the Forsaken without immunity, Cannibalize both bars, Touch of the Grave at most once every 15 s. index.md's race list matches `combos.json`.
- **Source links** (curl -L, 2026-11-04): all nine distinct frontmatter URLs answered 200. Unchanged since round 22's label check.
- **Voice:** no em or en dashes in the four pages, no site or sim talk, no nightly numbers, no stale 70755 reference.

## Decisions

- **D1. R23-2 moves from SHOULD to CONFIRMED.** Round 22 carried it as a SHOULD before its own rewrite of :33. The line now lists the exact spells the rank rule cuts, the page says the opposite 19 lines later, and the rendered rows cast partial-share ranks. That is a false "every" and a page contradiction, not a preference. The Oath's level rule is the engine's reading of Blessing of Light (`oath_of_light.go`, no Forever log at 25+), but the client's "up to" wording fits it and the page states it as fact, so the page must agree with itself. If later logs overturn the rule, :78 and :127 change, not just these two.
- **D2. R23-1 is a note finding, not a card-action finding.** Casting rank 1 at 20 to 24 is a measured action (accepted). The sentence that calls that cast the only heal that outheals a tank is what is false there.
- **D3. retribution.md:131 stays a SHOULD.** At 20 Sanctified Judgement is 33% of 20%, about 7% a Judgement, and at 21 it is 66% of 40%, about 26%. With a Judgement every 8 s (Improved Judgement from 16) against a 30 s Seal, that is about 25% of a Seal at 20, about all of it at 21 and more than all of it from 22. Only level 20 falls short, the line tells the player nothing to do at that level, and the same paragraph is general. Filing it would not meet this round's bar.
- **D4. Retribution row 6's "180 to 190" stays a SHOULD.** It is a range containing both exact values.
- **D5. Not re-filed, per the brief:**
  - The opener's "and raid buffs" and row 2's ungated Seal of Command (ledger Round 20).
  - Consecration ranks 1 to 4 at 50 to 53 (ledger Round 19).
  - Holy emergency Holy Light rank at 47 to 59 (ledger Round 19).
  - Whole rows rendering before their talent: Seal of Command rows at 20 to 25, Swift Judgement at 10 to 28, Divine Favor at 20, Holy Shock at 30 and 31, and the Unsealed Wrath release row at 10 to 19.
  - BiS race captions; card actions casting measured ranks; the Holy Light rank 9 book; Seal of Righteousness proc coefficient; Exorcism never crits; Benediction not discounting Hammer of Wrath.
  - The Retribution raid trinket2 Frozen Heart at zero hit value (with the re-rank since round 14).

## Sims requested

- Holy emergency Holy Light rank at 20 to 24 (`spec-breakdown`, both factions, rank 1 against ranks 3 and 4 from 22). Optional. The row casts 43 points on a tank below 40% while rank 3 (155) is trained, which is the reason R23-1 exists. The note fix stands either way.
- Holy party Flash of Light rank 3 against rank 4 at 50 to 54, where the party row casts a bigger rank than the tank row. Optional.
- Carried and already routed: Retribution row 2 Seal of Command with and without the 1200 attack power / 30% mana check (ledger Round 20); Retribution Consecration ranks 1 to 4 at 50 to 53 (ledger Round 19).

## Engine questions

- None new. Consecration (`sim/paladin/consecration.go`), Righteous Fury (1.33 x 1.9 on Holy, `righteous_fury.go`), Holy Conduit (`talents.go:331-335`), the level-gap rule (`sim/core/spell.go:777-850`, Holy Light rank 1 MaxLevel 5 in `spell_max_level_auto_gen.go`) and Oath of Light (`oath_of_light.go`) agree with the page's lines.
- Housekeeping, not a guide finding: the Protection APL's top-level `notes` still says Righteous Fury is "worth +60 percent threat" where the client and engine say 33% and then 90% on Holy. That field does not render.

## Sim findings

- None new.

## Sources

No web research was needed this round: the client data and engine decided every call.

- **Client data 1.60.1.70790** (`data/builds/1.60.1.70790/`): `talents/paladin.json`, `spelltips/paladin.json`, `trainables/paladin.json`, `addon-builds.json`, `addon-data.json` (`rotations/paladin-*`, every band), `bis/paladin-*.json`, `items/paladin.json`, `races.json`, `combos.json`, `simconsumes.json`.
- **Curated:** `data/curated/apl/paladin-{holy,protection,retribution}.json`, `data/curated/rank-choices.json`.
- **Engine d5c708834:** `sim/core/spell.go:777-850`, `sim/core/spell_max_level_auto_gen.go`, `sim/core/debuffs.go:302-322`, `sim/paladin/consecration.go`, `oath_of_light.go`, `holy_light.go`, `flash_of_light.go`, `sor.go`, `righteous_fury.go`, `talents.go`, `healing_talents.go`.
- **Git:** d3320fa41 and 5cc46bbfa (round 22 paladin fixes), 3d5260aa1 (addon-data; paladin rotations identical).
- **Reports:** `design/reviews/guide-accuracy/round22/paladin.md`, `ledger.md` Rounds 19 to 22, round 14 to 18 paladin (Heart trinkets, racials, Elixir of Holy Power).
- **Links fetched 2026-11-04, all 200:**
  - https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
  - https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap
  - https://www.wowhead.com/forever/news/undead-paladins-coming-to-world-of-warcraft-forever-382823
  - https://foreversixty.gg/data/1.60.1.70790/combos.json
  - https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
  - https://foreversixty.gg/planner?class=paladin
  - https://foreversixty.gg/sim/weights
  - https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/
  - https://foreversixty.gg/data/1.60.1.70790/races.json

## Counts

| status | count |
|---|---|
| CONFIRMED | 2 (R23-1 Holy row 3 card note, R23-2 holy.md:33 and :59) |
| LIKELY | 0 |
| NEEDS-SIM | 0 new, 2 carried (routed), 2 optional measurements requested |
| NEEDS-OWNER | 0 |
| Card notes with a finding | 1 of 33 (Retribution 11, Protection 13, Holy 9) |

## Note sweep

One line per row note: the spec, the row's action, the bands it renders on, each named talent, spell, rank or value with the level the player has it (leveling / raid order where they differ), and the verdict. All notes are byte-identical to round 22. Each line was re-read against the rows, the orders and round 23's check of both sides of every comparison.

**Retribution** (bands 10, 12, 18, 20, 26, 28, 30, 34, 36, 38, 40, 42, 44, 50, 52, 54, 58, 60; the raid order differs only from 40, and no note names a talent that moves)

- **Row 1 Judgement, 10-60:** trains at 4; off the global cooldown; 6% of base mana. "A tenth less once Benediction is full at 14": 10-14 on both orders. Never consumes the Seal (20271). "From 22 Sanctified Judgement": 3/3 at 22, 100% of 60%. OK.
- **Row 2 Seal of Command (opener, Unsealed Wrath), 20-60:** conditional on Unsealed Wrath (raid 40). Seal of the Crusader 6, Judgement 4, Seal of Command 26 (render at 20-25 is D5). Three swings store three stacks (1310735). "Your swings keep the debuff up": every Crusader rank's text, "Your melee strikes will refresh". OK; the raid-buff gate is a routed sim.
- **Row 3 Seal of Righteousness (opener, no Unsealed Wrath), 10-60:** "top rank you have trained": every band renders its trainer rank (20287 at 10 to 20293 at 58). Seal of the Crusader 6. OK.
- **Row 4 Exorcism, 20-60:** Exorcism 20; Undead and Demons only. OK.
- **Row 5 Hammer of Wrath, 44-60:** trains at 44; 6 s on all three ranks; Champion of the Light 5/5 at 39 on both orders takes 1.0 s off a 1.0 s cast. OK.
- **Row 6 Seal of Righteousness (Unsealed Wrath release), 10-60:** Unsealed Wrath conditional (raid 40); 30% mana and 1200 attack power; "180 to 190 at 60" is 180 and 189 (SHOULD, D4). OK.
- **Row 7 Seal of Command (back after one swing), 20-60:** Unsealed Wrath restarts at 15 s. OK.
- **Row 8 Seal of Command (keep up), 20-60:** Unsealed Wrath, raid buffs, more than 30% mana, last second. OK.
- **Row 9 Seal of Righteousness (otherwise), 10-60:** top rank trained (as row 3); last second; under 30% mana. OK.
- **Row 10 Holy Strike, 10-60:** trains at 6; "18 mana at 60 with Benediction": 10333 is 20 less 10%, exact; Sacred Arbiter conditional (27 on both orders). OK.
- **Row 11 Consecration, 20-60:** Consecration 20; reserve of a Seal, a Judgement and a Holy Strike; four-fifths pacing; Seal lapsing before the execute (Hammer of Wrath's own row starts at 44, and the clause is about pacing, not a named spell). No cost or rank claim. OK.

**Protection** (bands 10, 12, 18, 20, 26, 28, 30, 34, 36, 38, 40, 42, 43, 44, 50, 52, 58, 60; raid order identical)

- **Row 0 auto cooldowns** (not a rendered row): Lay on Hands 10, below 10%. OK.
- **Row 1 Templar's Bulwark, 20-60:** talent at 20; 100% of maximum health for 8 s; Forbearance; Divine Protection 6; "Divine Shield from 34". OK.
- **Row 2 Divine Shield, 34-60:** trains at 34; Templar's Bulwark from 20; 10 s on rank 1 (642, rows 34-44); "twelve from rank 2 at 50" (1020, rows 50-60); half damage dealt. OK.
- **Row 3 Divine Protection, 10-60:** 498 (6) and 5573 (18); Forbearance. OK.
- **Row 4 Stoneform (Dwarf)**, race line: 10% physical for 8 s, every 3 min, on the global cooldown. OK.
- **Row 5 Seal of Fury, 10-60:** Seal of Fury 10; last 2 s; "Holy Shield at 30"; 30 s on all seven ranks; Judgement does not consume it. OK.
- **Row 6 Holy Shield, 30-60:** talent 30; 10 s cooldown and duration; four charges (20925, 1239570; the rank 2/3 tooltips are with the data lane). OK.
- **Row 7 Holy Strike, 10-60:** "Twenty at 60": 10333, and the Protection build has no Benediction. Iron Creed conditional (39, 41-42). "From 40 Shieldsworn" (10 s less 4 s). "Little mana" holds on every rendered rank, including rank 1 at 40-42. OK.
- **Row 8 Hammer of the Righteous, 40-60:** trains at 40; Holy Shield 30; target plus three (407632); Righteous Fury 10. OK.
- **Row 9 Swift Judgement, 10-60:** one minute, free, its own cooldown (1310994); Seal of Fury taunt. OK (the render before 29 is D5).
- **Row 10 Judgement, 10-60:** "Holy Shield at 30"; off the global cooldown, 6% of base mana; Seal of Fury taunt 10. OK.
- **Row 11 Exorcism, 20-60:** Exorcism 20; Holy Shield at 30; Undead and Demons only. OK.
- **Row 12 Consecration, 20-60:** Consecration 20; half mana; "from 30, Holy Shield". "From 40 Hammer of the Righteous makes far more threat for its mana" holds against rank 1 (40-58) and rank 5 (60) on the rendered rows. OK.

**Holy** (bands 10, 14, 20, 25, 30, 38, 40, 43, 44, 46, 47, 50, 54, 55, 56, 60; raid order identical; Benediction at 53-57 changes no quoted value)

- **Row 0 auto cooldowns** (not a rendered row): mana potions and Demonic Runes; Divine Favor kept off the list. OK.
- **Row 1 Oath of Light, 25-60:** talent at 25; a flat amount on Holy Light, Flash of Light and Holy Shock (the three the talent names), while the rank is close to your level; 5 min. OK.
- **Row 2 Divine Favor, 20-60:** talent at 21; sure crit; 2 min (20216). OK (the render at 20 is D5).
- **Row 3 Holy Light, tank below 40%, 10-60:** "Divine Favor (from 21)". "At 60 rank 8" (10329, trains at 54); rank 9 book (D5). "Only heal that outheals a tank": 10-19 no other heal; **20-24 rank 1 (43) against Flash of Light rank 1 (46): PALADIN-R23-1**; 25-29 rank 2 with the Oath's share 233 to 218 against 94 to 102; 30 on larger on every band, including rank 4 at 40-42 and 47-49 against Holy Shock rank 1's 114 plus its share.
- **Row 4 Holy Light, tank below 70% (paced), 10-60:** five-point pacing, 45% example. OK.
- **Row 5 Holy Shock, below 70%, 30-60:** Holy Shock at 32 (the render at 30-31 is D5). "From 34": Boundless Light at 30%. Illumination 5/5 at 26; Oath target from 25. OK.
- **Row 6 Holy Shock, self, 30-60:** "Yourself included". OK.
- **Rows 7 to 10 Holy Shock, party:** no notes.
- **Row 11 Flash of Light, tank below 90%, 20-60:** trains at 20. "From 25": Oath, ranks close to your level. "At 60 without raid buffs, rank 4 ... rank 5": scoped to 60, where bare renders rank 4 and raid renders rank 5 (19942, 115 mana, against 90). OK.
- **Row 12 Flash of Light, everyone else below 80%, 20-60:** rank 1 at 20-43; rank 2 at 44-49 (SHOULD). "Rank 4 from 50" (19941, trains at 42; the row from 50). "Best balance" is a measured rank pick, not a per-mana superlative. OK.
- **Rows 13 to 16 Flash of Light, party:** no notes.

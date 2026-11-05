# Guide accuracy round 22: Paladin

Audited on the round 22 audit tree (987d08a7c: main after guide round 21; the log has no sim round 21 commit and there is no `*sim-audit-round21*` report, so sim round 21 had not landed). Engine pin d5c708834 (`sim/enginever/version.go`; source read at `scratchpad/gr22/engine`). Active client 1.60.1.70790 under `data/builds/1.60.1.70790/`.

Files audited:
- `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`.
- Every card note in `data/curated/apl/paladin-{holy,protection,retribution}.json`.
- Every rendered row of `addon-data.json` (`rotations/paladin-retribution` 18 bands, `paladin-protection` 18 bands, `paladin-holy` 16 bands).
- Both shipped point orders per spec in `addon-builds.json`.
- `talents/paladin.json`, `spelltips/paladin.json`, `trainables/paladin.json`.

Read first: the round 21 paladin report and the ledger's Round 21 section.

Lenses:
- Round 20's band-by-band note sweep, on both orders.
- Round 21's sweep of every superlative, comparison and quoted cost, duration and cooldown, checked against every spell the player has at that level.
- Round 22's four shapes, checked on every page and note:
  - a spell named in a "Levels 20 to 60" paragraph without its level;
  - a rank comparison on bands before the rank trains;
  - a "less than your best" list that includes the best itself;
  - a rounded cost quoted as exact.

## Changes since round 21's tree (e039e1490)

- c0f2e01f9 changed two lines, both in retribution.md:
  - :95 opens with "On one target it is filler, and running dry for your Seal costs far more than a skipped Consecration, so pace it". This replaces the "most expensive button" line (R21-2).
  - :133 now reads "Rank 1, about 120 mana after Benediction, keeps almost all of its spell damage through your 30s for a fraction of the cost" (R21-1).
- 93973f2f9 and 39050340f regenerated addon-data. The `rotations/paladin-*` blocks are byte-identical to round 21's tree, so the paladin card notes and rows did not change.
- `addon-builds.json` paladin orders are unchanged. They are decoded in the round 21 report. On the Retribution leveling and raid orders, Repentance is index 21 = level 31 (tier 5, col 2).

## Findings

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R22-1 | CONFIRMED | page prose (list of "cheap" heals includes the top rank; contradicts the page) | holy.md:95 | "so at 60 your cheap heals are Flash of Light rank 4 or 5 and Holy Shock rank 3 or 4, not rank 1." | Holy Shock rank 4 (20930, trained at 56) is the top rank and costs 325 mana for 320 healing. Rank 3 (20929) costs 275 for 230. Both cost more per point healed than any Flash of Light rank: rank 4 is 90 mana for 151 and rank 5 is 115 for 223 (`spelltips/paladin.json`). The page treats Holy Shock as a spend-when-ahead heal: :77 says to cast it "only with mana to spare" and to "heal with Flash of Light alone" when behind, :83 says the same, and card rows 5 and 6 are paced on mana. Calling rank 4, the most expensive Holy Shock, a "cheap heal" is the round 22 shape: a "cheap" list that includes the top rank itself. | "...so at 60 heal with Flash of Light rank 4 or 5 and Holy Shock rank 3 or 4, not rank 1." |
| PALADIN-R22-2 | CONFIRMED | page prose (per-mana superlative false against other spells the player has; round 21 shape) | retribution.md:133 | "Rank 1, about 120 mana after Benediction, keeps almost all of its spell damage through your 30s for a fraction of the cost, so until your mid 40s it gives you the most damage for your mana." | The claim is absolute, but it holds only among Consecration ranks. Damage per mana for rank 1 at 20, with almost no spell damage: 16 + 32 = 48 Holy per target over 8 s (26573), three mobs = 144, for about 121 mana = about 1.2 damage per mana. Holy Strike rank 3 at the same level (1866) is 12 mana, 10.8 after Benediction, for 32% weapon damage plus 19 Holy, and 20% more with Sacred Arbiter from 27. A leveling two-hander lands that for about 60, so about 5 or more damage per mana. A 30 s Seal of Righteousness (60 to 90 mana, less 10%) fires on every swing. Round 21 filed the same shape elsewhere (Riptide and Penance per mana beside Healing Stream Totem and Lightwell). | "...so until your mid 40s it is the Consecration rank that gives you the most damage for your mana." |
| PALADIN-R22-3 | CONFIRMED | page prose (a spell named in "Levels 20 to 60" without its level; round 22 shape) | retribution.md:133 | "Repentance takes one Humanoid out of a bad pull, and Hammer of Justice stops a runner or a caster." | Repentance is a talent (20066). On both Retribution orders it is index 21, level 31. The same paragraph gives levels for Consecration (20), Exorcism (level 20) and Holy Wrath (level 50). A reader from 20 to 30 is told to use a button they cannot have. This is the same shape as round 21's "Soul Harvest named without its level", which was filed CONFIRMED. | "From 31, Repentance takes one Humanoid out of a bad pull, and Hammer of Justice stops a runner or a caster." |
| PALADIN-R22-4 | CONFIRMED | page prose ("every" wider than the talent; contradicts :59) | holy.md:33 | "Oath of Light adds a flat amount to every heal on the tank" | Oath of Light (1328548; talent text) covers only "your Holy Light ... your Flash of Light ... and your Holy Shock". holy.md:59 lists exactly those three. Lay on Hands, which the page sends to the tank or anyone about to die (:91), gets no Oath bonus: its amount is "equal to the Paladin's maximum health" (633, 2800, 10310). | "Oath of Light adds a flat amount to every Holy Light, Flash of Light and Holy Shock on the tank, so..." (rest unchanged) |

Counts: CONFIRMED 4, LIKELY 0, NEEDS-SIM 0 new (2 carried, already routed), NEEDS-OWNER 0.

## SHOULDs (wording, not findings)

- **Retribution row 6 note** "every Seal costs a global and its mana, 180 to 190 at 60": the exact costs are 180 (Seal of Righteousness rank 8, 200 less 10%) and 189 (Seal of Command rank 5, 210 less 10%). The range contains both, so nothing false is stated, but "180 or 189" would be exact. This was checked as the "rounded cost quoted as exact" shape and not filed, because a range quotes no single exact figure.
- **retribution.md:72** "Holy Conduit takes 40% off Consecration, Exorcism and Hammer of Wrath": the client text also lists Holy Wrath, which :99 and :133 tell the player to use. Incomplete, not false.
- Carried, still true as SHOULDs:
  - retribution.md:131 "Sanctified Judgement refunds most of each Seal" (3/3 only from 22);
  - retribution.md:121 "both come back every 10 seconds";
  - retribution.md:36 "Judge every 8 seconds";
  - retribution.md:93's repeated last sentence;
  - Holy row 12 "rank 1 for most of your leveling" (rank 2 renders at 44 to 49);
  - holy.md:33/:59 "close to your level" clause;
  - holy.md:81 "half as much" (rank 1 is 49% of rank 4, 58% of rank 3);
  - protection.md:111 block chance;
  - Holy Shield ranks 2 and 3 tooltips missing "4 charges" (data lane).

## Checked and right this round

- **Round 21's fixes.**
  - retribution.md:133, "Rank 1, about 120 mana after Benediction": 135 less 10% is 121.5, and Benediction is 5/5 from 14 on both orders, so "about" is correct.
  - retribution.md:133, "keeps almost all of its spell damage through your 30s": the engine's level gap (`sim/core/spell.go` SpellLevelPenalty) applies to rank 1 (learned 20, MaxLevel 25). It keeps 100% to 37, 95% at 38 and 90% at 39.
  - The later rank-1 claims in the same paragraph also hold: "about half" at 46 to 48 (55% to 45%) and "nearly all" lost by 55 (10% left).
  - retribution.md:95: "running dry for your Seal costs far more than a skipped Consecration" makes no cost or rank claim. The rest of the paragraph matches `rank-choices.json` and the rendered rows: rank 1 at 20 to 43 and 50 to 53, rank 2 at 44 to 49, rank 4 at 54 to 59 and bare 60, rank 5 at raid 60.
- **Quoted costs, durations and cooldowns at the ranks the rows cast**, with every talent discount in the build. All match:
  - Retribution: Judgement 6% base mana, a tenth less at 14. Holy Strike 18 at 60 with Benediction.
  - Protection: Holy Strike 20 (no Benediction in the Protection build). Righteous Fury 30% base mana for 30 minutes (25780).
  - Holy: Light's Vigil 1,340 (1311595; 1.5 s cast, so no Benediction) against ten Flash of Light rank 6 at 140. Blessing of Protection every 3 minutes (5 min less Guardian's Favor 2/2 at 49 to 50).
  - Lay on Hands: 20 minutes, 550 mana at rank 3.
  - Voice of Truth: 6 s every 3 minutes.
  - Divine Shield: 10 s, and 12 s from rank 2 at 50.
  - Templar's Bulwark: 8 s.
  - Holy Shield ranks at 30/40/50/60.
- **Superlatives and comparisons** checked against every spell held at the level, and right:
  - protection.md:99 and Protection row 12: Hammer of the Righteous does far more threat per mana than the Consecration rank the rows cast (rank 1 at 40 to 58, 48 base per target for 135; rank 5 at 60, 312 for 565). Hammer of the Righteous is 3x weapon DPS per target for 6% base mana.
  - holy.md:87 "only instant heal outside Lay on Hands": Light's Vigil has a 1.5 s cast.
  - protection.md:80 "the only parry talent": Deflection is the only parry talent in `talents/paladin.json`.
  - holy.md:40 "only one party heal".
  - retribution.md:133 "about twice as hard": rank 3 against rank 1 at 40 is 136 against 48 base per target plus the same 0.042 a tick coefficient, about 2.2 to 2.5x across 50 to 150 spell damage.
- **Rank comparisons on early bands:**
  - Protection row 2 ("twelve from rank 2 at 50"), Holy row 11 ("At 60 ..."), Holy row 12 ("rank 4 from 50") and Retribution row 6 ("at 60") are all scoped.
  - No note names a rank before it trains.
- **Levels 20 to 60 paragraphs:**
  - Retribution: Consecration, Exorcism and Holy Wrath have their levels. Seal of the Crusader (6) and Hammer of Justice (8) are held before 20. Repentance is R22-3.
  - Holy: Flash of Light 20, and Oath of Light at 25 as the 16th Holy point.
  - Protection: Templar's Bulwark 20 and Sacred Duty 19/21 on the order; Holy Shield 30/40/50/60; Hammer of the Righteous 40.
- **Talent texts** re-read against every talent bullet on the three pages, and right: Benediction, Sanctified Judgement 100% chance of 60%, Vengeance 3% x5, Champion of the Light, Holy Power +50% bonus = 175%, Shieldsworn 4 s off, Iron Creed, Redoubt, Sacred Duty, Purifying Power.
- **Source links** (curl -L, 2026-11-04): all nine distinct frontmatter URLs answered 200 and match their labels.
- **Voice:** no em or en dashes, no site or sim talk, no nightly numbers.

## Decisions

- **D1. R22-1 is CONFIRMED, not a SHOULD.** The sentence names Holy Shock rank 4 among "your cheap heals". Rank 4 is the most expensive Holy Shock, and the page paces Holy Shock on spare mana twice (:77, :83). That is a contradiction on the page, and it is the "cheap list includes the top itself" shape.
- **D2. R22-2 is filed even though the round 21 rewrite passed the rendered review.** The brief asks for every "best per mana" claim to be checked against every spell held at the level. The paragraph's subject is Consecration ranks, and the fix keeps that meaning.
- **D3. R22-4 is filed as CONFIRMED.** "Every heal" is false for Lay on Hands, and the page's own talent line (:59) names the three spells.
- **D4. Retribution row 6's "180 to 190" stays a SHOULD.** It is a range containing both exact values, not a single rounded figure.
- **D5. Not re-filed, per the brief:**
  - The opener's "and raid buffs" and row 2's ungated Seal of Command (ledger Round 20, Sims routed 1).
  - Consecration ranks 1 to 4 at 50 to 53 (ledger Round 19).
  - Holy emergency Holy Light rank at 47 to 59.
  - Whole rows rendering before their talent: Seal of Command rows at 20 to 25, Swift Judgement at 10 to 28, Divine Favor at 20, Holy Shock at 30 and 31, and the Unsealed Wrath release row at 10 to 19.
  - BiS race captions.
  - Card actions casting measured ranks.
  - The Holy Light rank 9 book (25292; settled round 6/7 as a book with no Forever source).
  - Seal of Righteousness proc coefficient (round 16).
  - Exorcism never crits (round 1, X1).

## Sims requested

- None new. Carried and already routed:
  - Retribution row 2, Seal of Command with and without the 1200 attack power / 30% mana check (ledger Round 20).
  - Retribution Consecration ranks 1 to 4 at 50 to 53 (ledger Round 19).

## Engine questions

- None. The engine's level-gap rule (`sim/core/spell.go:777-850`) and Benediction on Consecration (`sim/paladin/consecration.go:102-105`) agree with the page's Consecration lines.

## Sim findings

- None new.

## Sources

- **Client data 1.60.1.70790** (`data/builds/1.60.1.70790/`):
  - `talents/paladin.json`: Oath of Light, Holy Conduit, Repentance, Benediction, Sanctified Judgement, Guardian's Favor, Purifying Power, every max-rank text.
  - `spelltips/paladin.json`: 26573 to 20924, 1866, 679 to 10333, 1311606, 20473, 20929, 20930, 19750 to 19943, 633, 2800, 10310, 1311595, 20066, 20293, 20920, 25780, 1022 to 10278, 1310897, 642, 1020, 1311015, 407632.
  - `trainables/paladin.json`, `addon-builds.json` (Retribution order index 21 = Repentance), `addon-data.json` (`rotations/paladin-*`, every band).
- **Curated:** `data/curated/apl/paladin-{holy,protection,retribution}.json`, `data/curated/rank-choices.json`.
- **Engine d5c708834:** `sim/core/spell.go:777-850` (SpellLevelPenalty), `sim/paladin/consecration.go:37-115`.
- **Git:** c0f2e01f9 (round 21 paladin fixes), 93973f2f9 and 39050340f (addon-data; paladin rotations byte-identical).
- **Reports:** `design/reviews/guide-accuracy/round21/paladin.md`, `ledger.md` Round 21, round 6/7/10/11 paladin (Holy Light rank 9, Flash of Light rank lines).
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
| CONFIRMED | 4 (R22-1 holy.md:95, R22-2 retribution.md:133, R22-3 retribution.md:133, R22-4 holy.md:33; all page prose) |
| LIKELY | 0 |
| NEEDS-SIM | 0 new, 2 carried (routed) |
| NEEDS-OWNER | 0 |
| Card notes with a finding | 0 of 33 (Retribution 11, Protection 13, Holy 9) |

## Note sweep

One line per row note: the spec, the row's action, the bands it renders on, each named talent, spell, rank or value with the level the player has it (leveling / raid order where they differ), and the verdict. All notes are byte-identical to round 21. Each line was re-read against the rows, the orders and round 22's shapes: a rank compared before it trains, a spell named before it trains, a rounded exact cost, and a superlative.

**Retribution** (bands 10, 12, 18, 20, 26, 28, 30, 34, 36, 38, 40, 42, 44, 50, 52, 54, 58, 60; the raid order differs only from 40, and no note names a talent that moves)

- **Row 1 Judgement, 10-60:**
  - Judgement trains at 4; it is off the global cooldown and costs 6% of base mana.
  - "A tenth less once Benediction is full at 14": Benediction is 10-14 on both orders.
  - Never consumes the Seal (20271).
  - "From 22 Sanctified Judgement": 3/3 at 22, 100% chance of 60%.
  - OK.
- **Row 2 Seal of Command (opener, Unsealed Wrath), 20-60:**
  - Conditional on Unsealed Wrath (raid 40).
  - Seal of the Crusader 6, Judgement 4, Seal of Command 26 (the render at 20-25 is D5).
  - Three swings store three stacks (1310735).
  - OK; the raid-buff gate is a routed sim.
- **Row 3 Seal of Righteousness (opener, no Unsealed Wrath), 10-60:**
  - "Top rank you have trained": every band renders its trainer rank (20287 at 10 to 20293 at 58).
  - Seal of the Crusader 6.
  - OK.
- **Row 4 Exorcism, 20-60:** Exorcism 20; Undead and Demons only. OK.
- **Row 5 Hammer of Wrath, 44-60:**
  - Trains at 44; 6 s on all three ranks.
  - Champion of the Light 5/5 at 39 on both orders takes 1.0 s off a 1.0 s cast.
  - OK.
- **Row 6 Seal of Righteousness (Unsealed Wrath release), 10-60:**
  - Unsealed Wrath is conditional (raid 40); "three stored Commands" is inside that condition.
  - 30% mana and 1200 attack power.
  - "180 to 190 at 60" is 180 and 189 (SHOULD, D4).
  - OK.
- **Row 7 Seal of Command (back after one swing), 20-60:** Unsealed Wrath restarts at 15 s. OK.
- **Row 8 Seal of Command (keep up), 20-60:** Unsealed Wrath, raid buffs, more than 30% mana, last second. OK.
- **Row 9 Seal of Righteousness (otherwise), 10-60:** top rank trained (as row 3); last second; under 30% mana. OK.
- **Row 10 Holy Strike, 10-60:**
  - Holy Strike trains at 6.
  - "18 mana at 60 with Benediction": 10333 is 20 less 10%, exact.
  - Sacred Arbiter is conditional (27 on both orders).
  - OK.
- **Row 11 Consecration, 20-60:** Consecration 20; reserve of a Seal, a Judgement and a Holy Strike; four-fifths pacing; Seal lapsing before the execute. No cost or rank claim. OK.

**Protection** (bands 10, 12, 18, 20, 26, 28, 30, 34, 36, 38, 40, 42, 43, 44, 50, 52, 58, 60; raid order identical)

- **Row 0 auto cooldowns** (not a rendered row): Lay on Hands 10, below 10%. OK.
- **Row 1 Templar's Bulwark, 20-60:** talent at 20; 100% of maximum health for 8 s; Forbearance; Divine Protection 6; "Divine Shield from 34". OK.
- **Row 2 Divine Shield, 34-60:**
  - Trains at 34; Templar's Bulwark from 20.
  - 10 s on rank 1 (642, rows 34-44); "twelve from rank 2 at 50" (1020, rows 50-60).
  - Half damage dealt.
  - OK.
- **Row 3 Divine Protection, 10-60:** 498 (6) and 5573 (18); Forbearance. OK.
- **Row 4 Stoneform (Dwarf)**, race line: 10% physical for 8 s, every 3 min, on the global cooldown. OK.
- **Row 5 Seal of Fury, 10-60:** Seal of Fury 10; last 2 s; "Holy Shield at 30"; 30 s on all seven ranks; Judgement does not consume it. OK.
- **Row 6 Holy Shield, 30-60:** talent 30; 10 s cooldown and duration; four charges (20925, 1239570; the rank 2/3 tooltips are with the data lane). OK.
- **Row 7 Holy Strike, 10-60:**
  - "Twenty at 60": 10333, and the Protection build has no Benediction.
  - Iron Creed is conditional (39, 41-42).
  - "From 40 Shieldsworn" (10 s less 4 s).
  - OK.
- **Row 8 Hammer of the Righteous, 40-60:** trains at 40; Holy Shield 30; target plus three (407632); Righteous Fury 10. OK.
- **Row 9 Swift Judgement, 10-60:** one minute, free, its own cooldown (1310994); Seal of Fury taunt. OK (the render before 29 is D5).
- **Row 10 Judgement, 10-60:** "Holy Shield at 30"; off the global cooldown, 6% of base mana; Seal of Fury taunt 10. OK.
- **Row 11 Exorcism, 20-60:** Exorcism 20; Holy Shield at 30; Undead and Demons only. OK.
- **Row 12 Consecration, 20-60:**
  - Consecration 20; half mana; "from 30, Holy Shield".
  - "From 40 Hammer of the Righteous makes far more threat for its mana" holds against rank 1 (40-58) and rank 5 (60) on the rendered rows.
  - OK.

**Holy** (bands 10, 14, 20, 25, 30, 38, 40, 43, 44, 46, 47, 50, 54, 55, 56, 60; raid order identical; Benediction at 53-57 changes no quoted value)

- **Row 0 auto cooldowns** (not a rendered row): mana potions and Demonic Runes; Divine Favor kept off the list. OK.
- **Row 1 Oath of Light, 25-60:** talent at 25; a flat amount on Holy Light, Flash of Light and Holy Shock (the three the talent names), while the rank is close to your level; 5 min. OK.
- **Row 2 Divine Favor, 20-60:** talent at 21; sure crit; 2 min (20216). OK (the render at 20 is D5).
- **Row 3 Holy Light, tank below 40%, 10-60:**
  - "Divine Favor (from 21)".
  - "At 60 rank 8" (10329, trains at 54); rank 9 book (D5).
  - Only heal that outheals a tank: Holy Light is the largest heal on every band, including rank 4's 287 at 47-49 against Holy Shock rank 1's 114.
  - OK.
- **Row 4 Holy Light, tank below 70% (paced), 10-60:** five-point pacing, 45% example. OK.
- **Row 5 Holy Shock, below 70%, 30-60:**
  - Holy Shock at 32 (the render at 30-31 is D5).
  - "From 34": Boundless Light at 30%.
  - Illumination 5/5 at 26; Oath target from 25.
  - OK.
- **Row 6 Holy Shock, self, 30-60:** "Yourself included". OK.
- **Rows 7 to 10 Holy Shock, party:** no notes.
- **Row 11 Flash of Light, tank below 90%, 20-60:**
  - Flash of Light trains at 20.
  - "From 25": Oath, ranks close to your level.
  - "At 60 without raid buffs, rank 4 ... rank 5": scoped to 60, where bare renders rank 4 and raid renders rank 5 (19942, 115 mana, against 90).
  - OK.
- **Row 12 Flash of Light, everyone else below 80%, 20-60:**
  - Rank 1 at 20-43; rank 2 at 44-49 (SHOULD).
  - "Rank 4 from 50" (19941, trains at 42; the row from 50).
  - "Best balance" is a measured rank pick, not a per-mana superlative.
  - OK.
- **Rows 13 to 16 Flash of Light, party:** no notes.

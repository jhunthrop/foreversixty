# Guide accuracy round 20: Paladin

Audited on the round 20 audit tree (4234643f5: main after guide round 19 and sim round 20; engine pin d5c708834, source read at `scratchpad/gr20/engine`), against the ACTIVE client 1.60.1.70790 under `data/builds/1.60.1.70790/`. Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`; every card note in `data/curated/apl/paladin-{holy,protection,retribution}.json`; every rendered row of `addon-data.json` (`rotations/paladin-retribution` 18 bands, `paladin-protection` 18 bands, `paladin-holy` 16 bands, read band by band); the shipped point orders in `addon-builds.json`; `talents/paladin.json`, `spelltips/paladin.json`, `trainables/paladin.json`, `races.json`; `data/curated/rank-choices.json`. Read first: round 19 paladin report, sim round 20 `paladin.md` and `integration.md`, the ledger's round 19 section. Engine read for the cost question: `sim/paladin/holy_strike.go`, `judgement.go` (`paladin.benediction()`), `lay_on_hands.go:69`.

Main lens: every sentence of every card note at every band its row renders on, against the shipped order and trainer levels; then every page line.

## Changes since round 19's tree (6beb919c2)

- `addon-builds.json`: paladin orders unchanged (only Arcane raid and Elemental orders moved). Orders as round 19 decoded them: Retribution leveling Benediction 10-14, Improved Judgement 15-16, Conviction 17-19, Sanctified Judgement 20-22, Vindication 23-25, Seal of Command 26, Sacred Arbiter 27, Two-Handed Weapon Specialization 28-30, Repentance 31, Vengeance 32-34, Champion of the Light 35-39, then Divine Strength from 40 (raid: Unsealed Wrath 40); Protection Divine Intellect 10, Redoubt 11-15, Precision 16-18, Sacred Duty 19, Templar's Bulwark 20, Sacred Duty 21, ..., Swift Judgement 29, Holy Shield 30, ..., Iron Creed 39, Shieldsworn 40, Iron Creed 41-42; Holy Spiritual Focus 10-14, Healing Light 15-17, Reverence 18-19, Illumination 20, Divine Favor 21, Illumination 22-24, Oath of Light 25, Illumination 26, Infusion of Light 27-28, Purifying Power 29-30, Divine Precision 31, Holy Shock 32, Boundless Light 33-34, Holy Power 35-39, Light's Vigil 40.
- Card notes: round 19's eleven fixes (Protection 7, Holy 3, Retribution 3 including the merged ones), and sim round 20's new Retribution opener lines (rows 2 and 3), the "180 to 190" Seal line, "the top rank you have trained" on row 9, Holy rows 1 and 11.
- Pages: holy.md lines 33, 78, 127 (sim round 20 H3); retribution.md lines 78 (see R20-1), 89, 91, 95, 103, 133.
- Rendered rows (`addon-data.json`, regenerated at the pin): Retribution Seal of Righteousness rank 8 (20293) from band 58; Holy tank Flash of Light rank 2 at 43-49, rank 3 at 50-54, rank 4 at 55-59, rank 5 at 60 (raid); group Flash of Light rank 2 at 44-49, rank 4 from 50; Holy Light rank 4 at 47-49; Holy Shock rank 1 to 49. They match `rank-choices.json` and sim round 20 H1.

## Findings

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R20-1 | CONFIRMED | page prose (merge regression, contradicts the card and the page) | retribution.md:78 | "Judge it on your first swing, then put up Seal of Righteousness, rank 7 even after rank 8 arrives at 58. While leveling you stay in it. With Unsealed Wrath and raid buffs, switch to Seal of Command as soon as your first swing under Righteousness lands: it stores a Righteousness copy, your first Command swing releases it, and the twist is under way." | Sim round 20 S2 measured rank 8 over rank 7 from 58 (+0.43 to +0.45% raid, +3.4 to +3.8% bare) and the card now casts 20293 rank 8 on every Seal of Righteousness row from band 58; R4 measured the split opener (Seal of Command straight after the Crusader Judgement when Unsealed Wrath is known) over this Righteousness-first line (+0.36/+0.40%) and shipped it as card rows 2 and 3 and the pre-pull. The lane report says land/guides-r17's version of this paragraph "describes the losing line and the rank-7 reading" and "int/r20's versions of those paragraphs must win"; merge ac7638941 kept main's. `git show 5797b0792:` has "then put up your damage Seal: Seal of Righteousness while leveling, Seal of Command with Unsealed Wrath and raid buffs"; HEAD has the old text. The paragraph now contradicts retribution.md:89 ("Train rank 8 at 58"), card row 2's note ("then go straight to Seal of Command") and row 3's ("at the top rank you have trained"). | "Put a Blessing of Might on yourself if nobody has given you one. If no other paladin has Judgement of the Crusader on the target, put up Seal of the Crusader about a second and a half before the pull, Judge it on your first swing, then put up your damage Seal: Seal of Righteousness at the top rank you have trained while leveling, and Seal of Command straight away with Unsealed Wrath and raid buffs, so three swings store three Commands and the twist is under way. Your swings and Holy Strike keep the Crusader debuff up, ..." (rest of the paragraph unchanged) |
| PALADIN-R20-2 | CONFIRMED | card note (rank credited early, new in sim round 20) | `paladin-holy.json` priority 11 (Flash of Light, tank below 90%), third sentence; renders on bands 20 to 60 | "Without raid buffs, rank 4 does the same job for less mana." | Flash of Light rank 4 (19941) trains at 42. The row casts rank 1 (19750) at 20 to 42, rank 2 (19939) at 43 to 49, rank 3 (19940) at 50 to 54, rank 4 at 55 to 59 and rank 5 (19942) at 60 (rank-choices.json, addon-data). On bands 20 to 40 the reader has no rank 4; on 43 to 54 rank 4 costs more (90 mana) than the rank the row casts (50, 70), so it is not "less mana"; on 55 to 59 rank 4 is the row itself. The sentence is true only at 60, against rank 5 (bare 60 row, rank-choices note at line 448). | "On the tank below 90%: your staple. From 25, Oath of Light adds the same flat amount to every rank close to your level that you cast on the tank. At 60 without raid buffs, rank 4 does the same job as rank 5 for less mana." |
| PALADIN-R20-3 | CONFIRMED | card note (talent credited early; round 19 passed it on a band misread) | `paladin-holy.json` priority 5 (Holy Shock, below 70%), second sentence; renders on bands 30 to 60 | "On anyone but your Oath target, Boundless Light adds 30% crit, which Illumination turns back into mana." | Band 30 covers levels 30 to 37. The Holy order takes Holy Shock at 32 (index 22) and Boundless Light at 33 and 34 (indices 23-24): rank 1 is 15%, rank 2 30% (`talents/paladin.json` 1328550). At 32 the reader has Holy Shock and no Boundless Light, at 33 half of it. Round 19 checked this note as true "38 and up", which skips levels 32 to 37 of band 30. Illumination is 5/5 at 26, so the clause after it holds. | "On anyone below 70%, with the same mana to spare. From 34, on anyone but your Oath target, Boundless Light adds 30% crit, which Illumination turns back into mana." |
| PALADIN-R20-4 | CONFIRMED | card note (value; contradicts the page and the same card) | `paladin-retribution.json` priority 10 (Holy Strike), bands 10 to 60, and priority 1 (Judgement), bands 10 to 60 | Holy Strike: "A cheap Holy weapon strike, twenty mana at 60, ..."; Judgement: "Judgement is off the global cooldown, costs 6% of base mana and never consumes your Seal, ..." | Benediction (Retribution order 10 to 14) "Reduces the Mana cost of all instant cast spells and abilities by 10%" (`talents/paladin.json` 20101); retribution.md:60 says it covers "your Seals, Judgement, Holy Strike, Consecration and Exorcism"; the engine applies `paladin.benediction()` to both spells (`holy_strike.go`, `judgement.go:45`). A Retribution paladin at 60 pays 18 mana for Holy Strike rank 8 (10333, 20 base) and 5.4% of base mana for Judgement. Row 6 of the same card quotes the Seals after Benediction ("180 to 190 at 60" = 200 and 210 less 10%), so the card states one discount and drops it on the next row. Round 19's Holy Strike fix wrote the base cost. Protection's "twenty at 60" is right (no Benediction in 05005/15302102513501031/5003). | Holy Strike: "On cooldown. A cheap Holy weapon strike, 18 mana at 60 with Benediction, and with Sacred Arbiter it refreshes every Judgement on the target." Judgement: "On cooldown. Judgement is off the global cooldown, costs 6% of base mana less Benediction's discount and never consumes your Seal, and from 22 Sanctified Judgement hands back 60% of the judged Seal's cost." |

Counts: CONFIRMED 4, LIKELY 0, NEEDS-SIM 1 new (Sims requested 1), NEEDS-OWNER 0.

## SHOULDs (wording, not findings)

- **holy.md:33 and :59** "adds a flat amount to every heal on the tank" / "to every Holy Light, Flash of Light and Holy Shock you put on them": since sim round 20 O1 a rank learned far below your level loses the Oath bonus with its healing power, which holy.md:78 and card row 1 now say. The sim lane judged :59 "stays true" (H3); a "close to your level" clause on both would make the page agree with itself.
- **Holy row 12** names rank 1 "for most of your leveling" and rank 4 from 50; the row beside it casts rank 2 at 44 to 49. "rank 1 into your 40s, rank 2 from 44, rank 4 from 50" would match every band.
- **retribution.md:121** "it and Holy Strike both come back every 10 seconds" sits in Levels 1 to 19, where Improved Judgement makes it 9 seconds at 15 and 8 from 16 (said two paragraphs later); **retribution.md:36** "Judge every 8 seconds" (sim round 20's guide-lane SHOULD).
- Sim round 20's other guide-lane SHOULDs stand: the last sentence of "Judgements in a raid" (retribution.md:93) repeats the paragraph; holy.md steps 2 and 3 read as if Holy Light were emergency only.
- Carried: the Unsealed Wrath rows rendering on the leveling build; holy.md:81 "half as much"; retribution.md:133 "as much spell damage as any rank"; protection.md:111 block chance; Holy Shield ranks 2 and 3 tooltips missing "4 charges" (data lane).

## Checked and right this round

- **Round 19's eleven fixes**, each against the rendered rows: Protection Seal of Fury, Judgement, Exorcism ("once you have Holy Shield at 30"; Holy Shield is index 20, its row from band 30), Holy Strike ("twenty at 60", no Benediction on the Protection build; Shieldsworn index 30 = 40; Iron Creed conditional), Divine Shield (642 10 s on the rank 1 rows at 34 to 44, 1020 12 s from 50), Templar's Bulwark (Divine Shield from 34), Consecration (Holy Shield from 30, Hammer of the Righteous from 40); Holy Light ("Divine Favor (from 21)", "At 60 rank 8"); group Flash of Light (rank 4 from 50, see the SHOULD); tank Flash of Light first two sentences ("From 25"); Retribution Judgement ("from 22"), Seal line ("180 to 190 at 60").
- **Sim round 20's new work:** Retribution rows 2 and 3 (Seal of the Crusader 6, Judgement 4, the Crusader debuff refreshed by melee per 20308's tooltip, Unsealed Wrath's 15 s and three stacks), row 9 "top rank you have trained" (every band's Seal of Righteousness row is the rank its band trains: 20287 at 10, 20288 18, 20289 26, 20290 34, 20291 42, 20292 50, 20293 58); retribution.md:89 (rank 8 at 58, 20293), :91 (rank 7 removed), :103 (0.12 against 0.0765 = 57% more, "about 60%"); Holy row 1 ("as long as the rank is close to your level", O1); holy.md:33 ("a rank or two below your top one": rank 5 raid and rank 4 bare against rank 6), :78 (rank 5 learned at 50 keeps all of its power at 60; rank 4 the bare-60 row), :127.
- **Every other page line** on all four pages re-derived on 70790: talent bullets against the max-rank texts (Benediction, Improved Judgement 2 s, Conviction 6%, Sanctified Judgement 60%, Sacred Arbiter, Two-Handed Weapon Specialization, Vengeance, Champion of the Light 60% and 1.0 s, Unsealed Wrath; Redoubt, Precision, Sacred Duty, Templar's Bulwark, Shield Specialization, Swift Judgement, Improved Righteous Fury 9%, Anticipation 20, Holy Shield tier 4, Iron Creed, Shieldsworn, Toughness, Deflection; Spiritual Focus 70%, Healing Light 12%, Reverence 30%, Illumination 50% of base cost, Divine Favor, Infusion 1.0 s, Holy Shock 10 s, Boundless Light 30%, Holy Power 175%, Light's Vigil, Guardian's Favor, Divine Precision, Purifying Power), trainer levels (Judgement 4, Holy Strike 6, Hammer of Justice 8, Lay on Hands 10, Redemption 12, Consecration 20, Exorcism 20, Holy Wrath 50, Hammer of the Righteous 40, Holy Shield ranks 30/40/50/60, Blessing of Wisdom 14, Purify 8), spell values (Light's Vigil 1,340 against ten 140-mana Flash of Lights; Lay on Hands 550 mana and 20 min; Voice of Truth 6 s and 3 min; Righteous Fury 30% base mana; Seal of Fury and Seal of the Crusader Judgements 3 s taunt and 40 s), the Levels 1 to 19 paragraphs against the orders, races against races.json (Ruthless, The Human Spirit, Will to Survive, Stoneform 10% 8 s 3 min, Artisan's Precision 1%, Big Game Hunter 5%, Touch of the Grave 15 s, Will of the Forsaken without immunity, Cannibalize mana), the Consecration rank levels against the rendered rows (rank 2 at 44 to 49, rank 1 at 50 to 53, rank 4 from 54, rank 5 at 60).
- **Voice:** no em or en dashes, no site or sim talk, no nightly numbers.
- **Source links** (curl -L, 2026-11-04): all nine frontmatter links 200 and say what their labels say (Wowhead's Undead paladin news answered 200 this round).

## Decisions

- **D1. R20-1 is the worst.** The opener is the first instruction in the Retribution rotation section and now teaches the measured-losing line and the rank sim round 20 retired, against the card under it and the page's own line 89. It is a landing error (the integration note says "Retribution keeps round 20's split opener" for the card; the paragraph went the other way), so the fix restores the lane's text with the top-rank wording of line 89.
- **D2. R20-4 is filed although Judgement's "6% of base mana" passed nineteen rounds.** The page states Benediction's discount on both spells (retribution.md:60), the engine charges it, and row 6 of the same card prices the Seals after it; a reader comparing two rows of one card gets two rules. Lowest severity of the four.
- **D3. R20-3 re-opens a round 19 "checked and right" line** because a band covers levels to the next band, and band 30 runs to 37.
- **D4. holy.md:59 stays a SHOULD**, per sim round 20 H3, because holy.md:78 states the rank rule where the reader picks a rank.
- **D5. Not filed, per the brief:** rows rendering before their talent (Swift Judgement 10 to 28, Seal of Command rows 20 to 25, Divine Favor at 20, Holy Shock at 30 and 31), BiS race captions, the leveling tree's rendering, card actions casting measured ranks, the Holy Light rank 9 book (trainables lists 25292 at 60 as a skill-line row; rounds 1 to 19 settled it as the AQ book).

## Sims requested

1. **New:** Retribution row 2 and the matching pre-pull line cast Seal of Command whenever Unsealed Wrath is known, with no 1200 attack power or 30% mana term, while rows 6 to 8 and retribution.md:91 keep Seal of Command to raid buffs. Measure a fresh boss at bare 60 on the raid build (Human and Undead, 10,000 iterations, seeds 1 and 50001): row 2 as shipped against row 2 with `currentAttackPower >= 1200` (falling through to row 3 otherwise). If the gate wins, row 2's note gains "and raid buffs" (R20-1's fix already has it); if the shipped row wins, R20-1's fix drops "and raid buffs".
2. Carried: Retribution Consecration ranks 1 to 4 at 50 to 53 (sim round 20 S4 measured 58 to 60 only).
3. Carried, for logs: Seal of Righteousness rank 8 procs at two spell damage values; Oath of Light on a low rank (sim round 20 "Still open").

## Engine questions

- None new. The engine charges Benediction on Holy Strike and Judgement as the client says (R20-4 is a note error, not an engine one).

## Sim findings

- Card row 2 has no raid-buff gate (Sims requested 1); the rest of the card agrees with its notes once R20-4's costs are fixed.

## Sources

- **Client data 1.60.1.70790** (`data/builds/1.60.1.70790/`): `talents/paladin.json` (Benediction 20101, Sanctified Judgement, Unsealed Wrath 1310735, Champion of the Light, Boundless Light 1328550, Illumination, Holy Power, Iron Creed, Shieldsworn, Holy Shield, Templar's Bulwark, every max-rank text), `spelltips/paladin.json` (20271, 679, 10333, 19750 to 19943, 1311606, 20473, 20929, 20930, 26573 to 20924, 407632, 20925, 1239570, 642, 1020, 498, 5573, 20216, 1310994, 24275, 24239, 20287, 20293, 20920, 20308, 21082, 1311649, 20423, 879, 10314, 1311015, 25780, 2812, 1310911, 1311595, 1310897, 633, 10310, 635, 10329, 25292), `trainables/paladin.json`, `addon-builds.json`, `addon-data.json` (`rotations/paladin-*`, every band), `races.json`.
- **Curated:** `data/curated/apl/paladin-{holy,protection,retribution}.json`; `data/curated/rank-choices.json` (paladin-holy, -protection, -retribution).
- **Engine d5c708834:** `sim/paladin/holy_strike.go` (cost table, Benediction), `sim/paladin/judgement.go:36-45`, `sim/paladin/lay_on_hands.go:69`.
- **Git:** `git show 5797b0792:web/src/content/guides/paladin/retribution.md` (the sim lane's opener) against HEAD and merge ac7638941.
- **Reports:** `design/reviews/2026-11-04-sim-audit-round20/paladin.md` (S1, S2, S4, O1, H1, H3, R3, R4, "For the guide lane"), `integration.md`; `design/reviews/guide-accuracy/round19/paladin.md`; `ledger.md` (round 19).
- **Links fetched 2026-11-04, all 200:** https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap ; https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap ; https://foreversixty.gg/planner?class=paladin ; https://foreversixty.gg/sim/weights ; https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/ ; https://foreversixty.gg/data/1.60.1.70790/races.json ; https://foreversixty.gg/data/1.60.1.70790/combos.json ; https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/ ; https://www.wowhead.com/forever/news/undead-paladins-coming-to-world-of-warcraft-forever-382823

## Counts

| status | count |
|---|---|
| CONFIRMED | 4 (R20-1 page prose; R20-2, R20-3, R20-4 card notes) |
| LIKELY | 0 |
| NEEDS-SIM | 1 new, 2 carried |
| NEEDS-OWNER | 0 |

## Note sweep

One line per row note: spec, row action, bands it renders on, each named talent, spell, rank or value with the level the player has it, verdict.

**Raid order check (follow-up).** Only retribution.md ships a `raidBuild`. Its raid order matches the leveling order point for point through 39 and differs only from 40 (Unsealed Wrath 40, Divine Strength 41-45, Improved Seals 46-50, Holy Conduit 51-52, Divine Intellect 53-57, Precision 58-60; Reverence never). Protection and Holy have identical raid and leveling orders. Every talent a Retribution note names (Benediction 10-14, Sanctified Judgement 20-22, Sacred Arbiter 27, Champion of the Light 35-39, Unsealed Wrath as a stated condition) sits at the same level on both orders, and no note names a talent that moves. Result: no new findings on the raid order; the Retribution lines below hold on both orders.

**Retribution** (bands 10, 12, 18, 20, 26, 28, 30, 34, 36, 38, 40, 42, 44, 50, 52, 54, 58, 60)
- Row 1 Judgement, 10-60: Judgement (4); off the global cooldown (client); 6% of base mana (less Benediction from 10, 5/5 at 14): **R20-4**; never consumes the Seal (20271 tooltip); Sanctified Judgement 3/3 at 22 (orders 20-22), 60% (client): OK.
- Row 2 Seal of Command (opener, Unsealed Wrath), 20-60: Unsealed Wrath conditional (raid 40); Seal of the Crusader (6); Judgement (4); Seal of Command (talent 26); three swings, three stacks (1310735); swings keep the debuff up (20308 "Your melee strikes will refresh"): OK (whole-row render before 26 is D5). Gate question: Sims requested 1.
- Row 3 Seal of Righteousness (opener, no Unsealed Wrath), 10-60: top rank trained (20287 at 10 to 20293 at 58, each its band's trainer rank); Seal of the Crusader (6); Judgement (4): OK.
- Row 4 Exorcism, 20-60: Exorcism (20); Undead and Demons only (879 to 10314): OK.
- Row 5 Hammer of Wrath, 44-60: trains 44; 6 s on all ranks (24275 to 24239); Champion of the Light 5/5 at 39, 1.0 s off a 1.0 s cast: OK.
- Row 6 Seal of Righteousness (Unsealed Wrath release), 10-60 (fires from 40 raid): Unsealed Wrath conditional; three stored Commands; 30% mana and 1200 attack power (the condition); "180 to 190 at 60" (200 and 210 less Benediction's 10%); Seal of Command below Seal of Righteousness at low attack power (round 13 measurement): OK.
- Row 7 Seal of Command (back after one swing), 20-60: Unsealed Wrath 15 s restart (client): OK.
- Row 8 Seal of Command (keep up), 20-60: Unsealed Wrath, raid buffs, 30% mana, last second (condition ≤ 1 s): OK.
- Row 9 Seal of Righteousness (otherwise), 10-60: top rank trained (as row 3); last second; under 30% mana: OK.
- Row 10 Holy Strike, 10-60: Holy Strike (6), ranks 679 at 10 to 10333 at 60; "twenty mana at 60" (18 after Benediction): **R20-4**; Sacred Arbiter conditional (27): OK otherwise.
- Row 11 Consecration, 20-60: Consecration (20); the reserve of a Seal, a Judgement and a Holy Strike and the 0.8 pacing (condition); Seal lapsing before the execute: OK.

**Protection** (bands 10, 12, 18, 20, 26, 28, 30, 34, 36, 38, 40, 42, 43, 44, 50, 52, 58, 60)
- Row 0 auto cooldowns (Lay on Hands), not a rendered row: Lay on Hands (10); fires below 10% (lay_on_hands.go:69): OK.
- Row 1 Templar's Bulwark, 20-60: talent (20); 100% of maximum health for 8 s (1311015); Forbearance; Divine Protection (6); Divine Shield from 34: OK.
- Row 2 Divine Shield, 34-60: trains 34; Templar's Bulwark (20); 10 s on rank 1 (rows 34-44 cast 642), 12 s from rank 2 at 50 (1020); half damage dealt: OK.
- Row 3 Divine Protection, 10-60: 498 (6), 5573 (18); Forbearance: OK.
- Row 4 Stoneform (Dwarf race line): 10% physical for 8 s, 3 min (races.json), on the global cooldown (engine racials): OK.
- Row 5 Seal of Fury, 10-60: Seal of Fury (10); last 2 s (condition); Holy Shield at 30 (index 20); 30 s on all seven ranks; Judgement does not consume it: OK.
- Row 6 Holy Shield, 30-60: talent (30); 10 s cooldown and duration; four charges (20925, 1239570); first claim on mana: OK.
- Row 7 Holy Strike, 10-60: twenty at 60 (10333, no Benediction on this build); Iron Creed conditional (39, 41-42), 6% for 6 s at 3/3; Shieldsworn from 40 (index 30): OK.
- Row 8 Hammer of the Righteous, 40-60: trains 40; Holy Shield (30); target plus three (407632); Righteous Fury (10): OK.
- Row 9 Swift Judgement, 10-60: one minute, free, own cooldown (1310994); Seal of Fury taunt: OK (whole-row render before 29 is D5).
- Row 10 Judgement, 10-60: Holy Shield at 30; off the global cooldown, 6% of base mana (no Benediction); Seal of Fury taunt (10): OK.
- Row 11 Exorcism, 20-60: Exorcism (20); Holy Shield at 30; Undead and Demons only: OK.
- Row 12 Consecration, 20-60: Consecration (20); half mana (condition); Holy Shield from 30; Hammer of the Righteous from 40: OK.

**Holy** (bands 10, 14, 20, 25, 30, 38, 40, 43, 44, 46, 47, 50, 54, 55, 56, 60)
- Row 0 auto cooldowns, not a rendered row: mana potions and Demonic Runes; Divine Favor kept off the list (credits nothing): OK.
- Row 1 Oath of Light, 25-60: talent (25); flat amount on Holy Light, Flash of Light, Holy Shock while the rank is close to your level (sim round 20 O1); 5 min: OK.
- Row 2 Divine Favor, 20-60: talent (21); sure crit on the next Holy Light; 2 min (20216): OK (whole-row render at 20 is D5).
- Row 3 Holy Light, tank below 40%, 10-60: Divine Favor from 21; "At 60 rank 8" (10329 at 54, the 60 row); rank 9 book (D5); only heal that outheals a tank: OK.
- Row 4 Holy Light, tank below 70% (paced), 10-60: the 5-point pacing and the 45% example (condition): OK.
- Row 5 Holy Shock, tank below 70%, 30-60: Holy Shock (32; 30 and 31 are D5); Boundless Light 30% (33 and 34): **R20-3**; Illumination 5/5 at 26: OK.
- Row 6 Holy Shock, self, 30-60: "Yourself included": OK.
- Rows 7 to 10 Holy Shock, party: no notes.
- Row 11 Flash of Light, tank below 90%, 20-60: Flash of Light (20); Oath of Light from 25, ranks close to your level; "rank 4 without raid buffs" (rank 4 at 42; row ranks 1/2/3/4/5): **R20-2**.
- Row 12 Flash of Light, self below 80%, 20-60: rank 1 at 20-43, rank 2 at 44-49 (unnamed, SHOULD), rank 4 from 50 (19941 at 42, the row from 50): OK.
- Rows 13 to 16 Flash of Light, party: no notes.

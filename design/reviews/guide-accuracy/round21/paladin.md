# Guide accuracy round 21: Paladin

Audited on the round 21 audit tree (e039e1490: main after guide round 20; no sim round 21 commit in the log, no `*sim-audit-round21*` report, so nothing from sim round 21 had landed), engine pin d5c708834 (`sim/enginever/version.go`; source read at `scratchpad/gr21/engine`), against the ACTIVE client 1.60.1.70790 under `data/builds/1.60.1.70790/`. Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`; every card note in `data/curated/apl/paladin-{holy,protection,retribution}.json`; every rendered row of `addon-data.json` (`rotations/paladin-retribution` 18 bands, `paladin-protection` 18 bands, `paladin-holy` 16 bands, read band by band); both shipped point orders per spec in `addon-builds.json`; `talents/paladin.json`, `spelltips/paladin.json`, `trainables/paladin.json`, `races.json`, `combos.json`, `bis/paladin-*.md`; `data/curated/rank-choices.json`. Read first: round 20 paladin report, the ledger's Round 19 and Round 20 sections, earlier paladin rounds for carried items.

Main lens: every sentence of every card note at every band it renders on, on both orders; then every page line, with round 20's three shapes (raid-order credit, comparisons and superlatives that fail on a lower-rank band, quoted values a talent in the build changes).

## Changes since round 20's tree (4234643f5)

- Paladin commits: e37ecc0d4 (retribution.md:78 opener; Holy rows 5 and 11; Retribution rows 1 and 10), a5c0e7496 (Retribution row 1 wording "a tenth less once Benediction is full at 14"), a54761d7b and 9832b5feb (addon-data regenerated, note text only).
- `addon-builds.json` paladin orders unchanged. Decoded (index 0 = level 10):
  - Retribution leveling: Benediction 10-14, Improved Judgement 15-16, Conviction 17-19, Sanctified Judgement 20-22, Vindication 23-25, Seal of Command 26, Sacred Arbiter 27, Two-Handed Weapon Specialization 28-30, Repentance 31, Vengeance 32-34, Champion of the Light 35-39, Divine Strength 40-44, Reverence 45-46, Improved Seals 47-51, Holy Conduit 52-53, Divine Intellect 54-58, Precision 59-60. Raid: identical to 39, then Unsealed Wrath 40, Divine Strength 41-45, Improved Seals 46-50, Holy Conduit 51-52, Divine Intellect 53-57, Precision 58-60.
  - Protection (both orders identical): Divine Intellect 10, Redoubt 11-15, Precision 16-18, Sacred Duty 19, Templar's Bulwark 20, Sacred Duty 21, Shield Specialization 22-23, Toughness 24-28, Swift Judgement 29, Holy Shield 30, Improved Righteous Fury 31-33, Anticipation 34-38, Iron Creed 39, Shieldsworn 40, Iron Creed 41-42, Divine Strength 43-47, Improved Seals 48-52, Deflection 53-57, Conviction 58-60.
  - Holy (both orders identical): Spiritual Focus 10-14, Healing Light 15-17, Reverence 18-19, Illumination 20, Divine Favor 21, Illumination 22-24, Oath of Light 25, Illumination 26, Infusion of Light 27-28, Purifying Power 29-30, Divine Precision 31, Holy Shock 32, Boundless Light 33-34, Holy Power 35-39, Light's Vigil 40, Divine Intellect 41-45, Precision 46-48, Guardian's Favor 49-50, Sacred Duty 51-52, Benediction 53-57, Conviction 58-60.
- Round 20's four fixes re-checked against rows and orders: all four hold (see "Checked and right").

## Findings

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R21-1 | CONFIRMED | page prose (quoted cost a talent in the build changes; same shape as R20-4) | retribution.md:133 | "Rank 1 at 135 mana carries as much spell damage as any rank for a fraction of the cost" | Consecration trains at 20 (26573, 135 mana base). Benediction is 5/5 from 14 on both Retribution orders, "Reduces the Mana cost of all instant cast spells and abilities by 10%" (`talents/paladin.json` 20101), and retribution.md:60 names Consecration among the spells it discounts; the engine charges it (`sim/paladin/consecration.go:102-105`, `Multiplier: paladin.benediction()`). So a Retribution paladin never pays 135 for rank 1: it is about 121 from the day Consecration is trained, and from 52 (leveling order; 51 raid) Holy Conduit takes 40% more off. Round 20 fixed the identical base-cost quote on Holy Strike and Judgement (R20-4) and card row 6 already prices the Seals after Benediction, so the page now states one rule on the card and another in the prose. Round 13's carried SHOULD on "as much spell damage as any rank" (exact only to 37) is folded into the fix. | "From level 20 Consecration is your AoE: two or three mobs on you is the time for it. Rank 1, about 120 mana after Benediction, keeps almost all of its spell damage through your 30s for a fraction of the cost, so until your mid 40s it gives you the most damage for your mana." (rest of the paragraph unchanged) |
| PALADIN-R21-2 | CONFIRMED | page prose (superlative that fails on bands where a lower rank renders; round 20 shape 2) | retribution.md:95 | "**Consecration.** It is the most expensive button in your priority, so pace it." | The same paragraph and the card cast rank 1 at 20 to 43 and 50 to 53, rank 2 at 44 to 49, rank 4 at 54 to 59 and bare 60 (`rank-choices.json` paladin-retribution, `addon-data.json` rows), and every one of those is cheaper than another button in the priority above it: at 36 to 43 Exorcism rank 3 (5615, 180 less 10% = 162) costs more than Consecration rank 1 (135 less 10% = about 121), and at 42 so does Seal of Righteousness rank 6 (140 less 10% = 126); at 44 to 49 Hammer of Wrath rank 1 (24275, 295, no Benediction per retribution.md:60) against rank 2 (235 less 10% = about 212); at 50 to 53 rank 1 is among the cheapest buttons you press; at bare 60, rank 4 with Benediction and Holy Conduit (435, 50% off, about 218) against Hammer of Wrath rank 3 (425, 40% off = 255). Only raid 60's rank 5 (565, about 283) is the most expensive. | "**Consecration.** It is the one button in your priority you can always afford to skip, so pace it." |

Counts: CONFIRMED 2, LIKELY 0, NEEDS-SIM 0 new (2 carried, already routed), NEEDS-OWNER 0.

## SHOULDs (wording, not findings)

- **retribution.md:131** "Sanctified Judgement refunds most of each Seal" opens Levels 20 to 60, where Sanctified Judgement is 1/3 at 20 (33% chance of 20%) and 2/3 at 21; it is 3/3 from 22. "From 22, Sanctified Judgement refunds most of each Seal" matches the order. The section is a span and the talent completes two levels in, so this stays a SHOULD.
- Carried from round 20 and earlier, still true as SHOULDs: retribution.md:121 "it and Holy Strike both come back every 10 seconds" (Improved Judgement 15-16 is two paragraphs later); retribution.md:36 "Judge every 8 seconds"; retribution.md:93's repeated last sentence; Holy row 12 "rank 1 for most of your leveling" (rank 2 renders at 44 to 49); holy.md:33/:59 "close to your level" clause; holy.md:81 "half as much" (rank 1 is 49% of rank 4, 58% of rank 3); protection.md:111 block chance; Holy Shield ranks 2 and 3 tooltips missing "4 charges" (data lane).

## Checked and right this round

- **Round 20's fixes.**
  - retribution.md:78 (R20-1) now matches card rows 2 and 3 and :89 (top rank trained; Seal of Command straight after the Crusader Judgement), apart from "and raid buffs", which is the routed sim (ledger Round 20, Sims routed 1).
  - Retribution row 1 (R20-4): "costs 6% of base mana, a tenth less once Benediction is full at 14". Benediction 10-14 on both orders, 2% a rank, so 10% at 14, and the engine charges it on Judgement.
  - Retribution row 10 (R20-4): "18 mana at 60 with Benediction". 10333 is 20 base; with 10% off, 18.
  - Holy row 5 (R20-3): "From 34". Boundless Light is 33-34 on both Holy orders, 30% at rank 2 (1328550).
  - Holy row 11 (R20-2): "At 60 without raid buffs, rank 4 does the same job as rank 5 for less mana". The bare 60 row is rank 4 per rank-choices "preset: bare" (rank 4 beats rank 5 by 9.8 to 9.9%); the raid row is rank 5 (19942, 115 mana) against rank 4 (19941, 90).
- **Benediction on the Holy order (53-57).** No Holy note or line quotes the cost of an instant spell. The only Holy cost figure is Light's Vigil's 1,340 (1.5 s cast, not discounted), and holy.md:81 is a ratio that the discount leaves unchanged.
- **Every other page line** on all four pages against 70790. Specific checks:
  - Talent texts and point counts, both builds per spec. Ret leveling 14/7/30 and raid 12/8/31 (:58), Holy 31/12/8, Protection 10/33/8.
  - Trainer levels: Seal of Righteousness and Holy Light 1, Judgement and Blessing of Might 4, Holy Strike, Seal of the Crusader and Divine Protection 6, Hammer of Justice and Purify 8, Lay on Hands, Righteous Fury and Seal of Fury 10, Redemption 12, Blessing of Wisdom 14, Consecration, Exorcism, Blessing of Kings and Voice of Truth 20, Concentration Aura 22, Divine Shield 34/50, Hammer of the Righteous 40, Hammer of Wrath 44/52/60, Holy Wrath 50/60, Holy Shield 30/40/50/60, Seal of Righteousness rank 8 at 58.
  - Ranges: Judgement and Swift Judgement 10 yd, Hammer of Wrath 30 yd.
  - The index milestones: tiers at 10/15/20/30 points, so the 11th, 16th, 21st and 31st points.
  - Holy Power's row: the 25 Holy points before it at 35 include Divine Precision.
  - Oath of Light at 25 is the 16th Holy point.
  - Prot leveling: Sacred Duty 19, Templar's Bulwark 20, Sacred Duty 21.
  - Races against races.json and combos.json (Human, Dwarf, Undead; no Skyborne).
  - Frozen Heart and Dormant Heart of the Mountain both appear in their specs' BiS files.
  - Consecration rank prose (retribution.md:95, :133) against `rank-choices.json`.
- **Voice:** no em or en dashes, no site or sim talk, no nightly numbers.
- **Source links** (curl -L, 2026-11-04): all nine distinct frontmatter URLs answered 200 and match their labels.

## Decisions

- **D1. R21-1 is filed although "135 mana" passed twenty rounds.** Round 20 set the rule for this page (R20-4: a Retribution cost is stated after Benediction, because the page says Benediction covers the spell and the engine charges it), and round 20's fix applied it to the card but not to the prose two paragraphs away. Low severity, but it is the exact shape the brief names (shape 3).
- **D2. R21-2 is filed as CONFIRMED, not a SHOULD.** It is a superlative that is false on most bands where the paragraph's own ranks render (shape 2). It is true only at raid 60.
- **D3. Not re-filed, per the brief:**
  - The opener's "and raid buffs" against row 2's ungated Seal of Command. This is ledger Round 20, Sims routed 1. The same sim decides retribution.md:70 "Below 60 you never cast it" for a raid-order reader with Unsealed Wrath from 40.
  - Consecration ranks 1 to 4 at 50 to 53 (ledger Round 19).
  - Holy emergency Holy Light rank at 47 to 59 (ledger Round 19).
  - Whole rows rendering before their talent: Seal of Command rows at 20 to 25, Swift Judgement at 10 to 28, Divine Favor at 20, Holy Shock at 30 and 31.
  - BiS race captions.
  - Card actions casting measured ranks.
  - The Holy Light rank 9 book.

## Sims requested

- None new. Carried and already routed: Retribution row 2 Seal of Command with against without the 1200 attack power / 30% mana check (ledger Round 20); Retribution Consecration ranks 1 to 4 at 50 to 53 (ledger Round 19).

## Engine questions

- None. The engine applies Benediction to Consecration (`consecration.go:102-105`), Judgement and Holy Strike as the client text says. R21-1 and R21-2 are page errors.

## Sim findings

- None new.

## Sources

- **Client data 1.60.1.70790** (`data/builds/1.60.1.70790/`):
  - `talents/paladin.json`: Benediction 20101, Holy Conduit 1237268, Sanctified Judgement 1311074, Boundless Light 1328550, Unsealed Wrath 1310735, every max-rank text.
  - `spelltips/paladin.json`: costs and ranges for 26573, 20116, 20922, 20923, 20924, 879 to 10314, 24275, 24274, 24239, 20287 to 20293, 20375 to 20920, 679 to 10333, 19750 to 19943, 1311606 to 20930, 1310911 to 1311595, 20271, 1310994.
  - `trainables/paladin.json`, `addon-builds.json`, `addon-data.json` (`rotations/paladin-*`, every band), `races.json`, `combos.json`, `bis/paladin-*.md`.
- **Curated:** `data/curated/apl/paladin-{holy,protection,retribution}.json`; `data/curated/rank-choices.json` (paladin-holy, -protection, -retribution).
- **Engine d5c708834:** `sim/paladin/consecration.go:90-115`, `sim/paladin/talents.go:65-67` (benediction()), `:237-243`, `:331-335` (Holy Conduit cost mod).
- **Git:** e37ecc0d4, a5c0e7496 (round 20 paladin fixes).
- **Reports:** `design/reviews/guide-accuracy/round20/paladin.md`; `ledger.md` Round 19 and Round 20; round 10 to 16 paladin reports (Consecration history, carried SHOULD).
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
| CONFIRMED | 2 (R21-1, R21-2; both retribution.md prose) |
| LIKELY | 0 |
| NEEDS-SIM | 0 new, 2 carried (routed) |
| NEEDS-OWNER | 0 |
| Card notes with a finding | 0 of 33 (Retribution 11, Protection 13, Holy 9) |

## Note sweep

One line per row note: spec, row action, bands it renders on, each named talent, spell, rank or value with the level the player has it (leveling / raid order where they differ), verdict.

**Retribution** (bands 10, 12, 18, 20, 26, 28, 30, 34, 36, 38, 40, 42, 44, 50, 52, 54, 58, 60; raid order differs only from 40, and no note names a talent that moves)

- Row 1 Judgement, 10-60. Judgement (4); off the global cooldown; 6% of base mana; "a tenth less once Benediction is full at 14" (Benediction 10-14 both orders); never consumes the Seal (20271); Sanctified Judgement 3/3 at 22 both orders, 60%. OK.
- Row 2 Seal of Command (opener, Unsealed Wrath), 20-60. Conditional on Unsealed Wrath (raid 40); Seal of the Crusader (6); Judgement (4); Seal of Command (26; render at 20-25 is D3); three swings, three stacks (1310735); swings refresh the debuff (Crusader "Your melee strikes will refresh"). OK. The raid-buff gate is a routed sim.
- Row 3 Seal of Righteousness (opener, no Unsealed Wrath), 10-60. "Top rank you have trained": every band's rank is its trainer rank (20287 at 10, 20288 at 18, 20289 at 26, 20290 at 34, 20291 at 42, 20292 at 50, 20293 at 58). Seal of the Crusader (6). OK.
- Row 4 Exorcism, 20-60. Exorcism (20); Undead and Demons only (879-10314). OK.
- Row 5 Hammer of Wrath, 44-60. Trains at 44; 6 s on all three ranks; Champion of the Light 5/5 at 39 both orders, 1.0 s off a 1.0 s cast. OK.
- Row 6 Seal of Righteousness (Unsealed Wrath release), 10-60. Unsealed Wrath conditional; three stored Commands; 30% mana and 1200 attack power (both in the condition); "180 to 190 at 60" (200 and 210 less 10%); Seal of Command below Seal of Righteousness at low attack power. OK.
- Row 7 Seal of Command (back after one swing), 20-60. Unsealed Wrath restarts at 15 s. OK.
- Row 8 Seal of Command (keep up), 20-60. Unsealed Wrath, raid buffs (1200 attack power in the condition), more than 30% mana, last second. OK.
- Row 9 Seal of Righteousness (otherwise), 10-60. Top rank trained (as row 3); last second; under 30% mana. OK.
- Row 10 Holy Strike, 10-60. Holy Strike (6); "18 mana at 60 with Benediction" (10333 20 less 10%); Sacred Arbiter conditional (27 both orders). OK.
- Row 11 Consecration, 20-60. Consecration (20); reserve of a Seal, a Judgement and a Holy Strike; four-fifths pacing (condition); Seal lapsing before the execute. OK.

**Protection** (bands 10, 12, 18, 20, 26, 28, 30, 34, 36, 38, 40, 42, 43, 44, 50, 52, 58, 60; raid order identical)

- Row 0 auto cooldowns (Lay on Hands), not a rendered row. Lay on Hands (10), below 10%. OK.
- Row 1 Templar's Bulwark, 20-60. Talent (20); 100% of maximum health for 8 s (1311015); Forbearance; Divine Protection (6); "Divine Shield from 34". OK.
- Row 2 Divine Shield, 34-60. Trains 34; Templar's Bulwark (20); 10 s on rank 1 (642, rows 34-44), 12 s from rank 2 at 50 (1020); half damage dealt. OK.
- Row 3 Divine Protection, 10-60. 498 (6), 5573 (18); Forbearance. OK.
- Row 4 Stoneform (Dwarf), race line. 10% physical for 8 s, 3 min (races.json), on the global cooldown. OK.
- Row 5 Seal of Fury, 10-60. Seal of Fury (10); last 2 s; Holy Shield at 30; 30 s on all seven ranks; Judgement does not consume it. OK.
- Row 6 Holy Shield, 30-60. Talent (30); 10 s cooldown and duration; four charges (20925, 1239570). OK.
- Row 7 Holy Strike, 10-60 (rank 1 at 40-42 by rank choice). "Twenty at 60" (10333, no Benediction on this build); Iron Creed conditional (39, 41-42; 6% for 6 s with Righteous Fury); Shieldsworn from 40 (10 s less 4 s = 6 s). OK.
- Row 8 Hammer of the Righteous, 40-60. Trains 40; Holy Shield (30); target plus three (407632); Righteous Fury (10). OK.
- Row 9 Swift Judgement, 10-60. One minute, cost 0, own cooldown (1310994); Seal of Fury taunt. OK (render before 29 is D3).
- Row 10 Judgement, 10-60. Holy Shield at 30; off the global cooldown, 6% of base mana (no Benediction on the Protection build); Seal of Fury taunt (10). OK.
- Row 11 Exorcism, 20-60. Exorcism (20); Holy Shield at 30; Undead and Demons only. OK.
- Row 12 Consecration, 20-60. Consecration (20); half mana; "from 30, Holy Shield"; "From 40 Hammer of the Righteous". OK.

**Holy** (bands 10, 14, 20, 25, 30, 38, 40, 43, 44, 46, 47, 50, 54, 55, 56, 60; raid order identical; Benediction 53-57 changes no quoted value)

- Row 0 auto cooldowns, not a rendered row. Mana potions and Demonic Runes; Divine Favor kept off the list. OK.
- Row 1 Oath of Light, 25-60. Talent (25); flat amount on Holy Light, Flash of Light and Holy Shock while the rank is close to your level; 5 min. OK.
- Row 2 Divine Favor, 20-60. Talent (21); sure crit; 2 min (20216). OK (render at 20 is D3).
- Row 3 Holy Light, tank below 40%, 10-60. "Divine Favor (from 21)"; "At 60 rank 8" (10329, L54); rank 9 book (D3); the only heal that outheals a tank (Holy Light is the largest heal at every band, including rank 4's 287 at 47-49). OK.
- Row 4 Holy Light, tank below 70% (paced), 10-60. Five-point pacing, 45% example. OK.
- Row 5 Holy Shock, below 70%, 30-60. Holy Shock (32; 30-31 render is D3); "From 34" Boundless Light 30%; Illumination 5/5 at 26. OK.
- Row 6 Holy Shock, self, 30-60. "Yourself included". OK.
- Rows 7 to 10 Holy Shock, party: no notes.
- Row 11 Flash of Light, tank below 90%, 20-60. Flash of Light (20); "From 25" Oath, ranks close to your level; "At 60 without raid buffs, rank 4 ... rank 5" (bare 60 rank 4, raid 60 rank 5). OK.
- Row 12 Flash of Light, everyone else below 80%, 20-60. Rank 1 20-43, rank 2 44-49 (SHOULD), rank 4 from 50 (19941, L42, row from 50). OK.
- Rows 13 to 16 Flash of Light, party: no notes.

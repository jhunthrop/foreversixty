# Guide accuracy round 20: Warrior

Audited on the round 20 audit tree (4234643f5: main after guide round 19 and sim round 20, engine pin d5c708834 in `sim/enginever/version.go`) against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/warrior/{index,arms,fury,protection}.md`, every note in `data/curated/apl/warrior-{arms,fury,protection}.json`, every rendered row of `addon-data.json` (`rotations/warrior-arms` 20 bands, `warrior-fury` 19 bands, `warrior-protection` 24 bands), both shipped orders per spec in `addon-builds.json`, and `talents/warrior.json`, `spelltips/warrior.json`, `trainables/warrior.json`. I read the round 19 warrior report, the sim round 20 warrior report and `integration.md`, and the ledger's Round 18 and Round 19 sections, and checked the earlier rounds for each candidate (round 18 R18-1/R18-2 in particular).

How a note renders (re-read this round): the addon row carries the note up to the first ". " at every band (`data/pipeline/addonrotation.py` `_one_line`; band chosen as the highest `band.level <= level`, `addon/ForeverSixty/Rotation.lua` `bandFor`, so a band covers every level up to the next band). The guide page's card prints the clause up to the first ": ", ". " or "; " (`web/src/lib/guides/rotation-lists.ts` `clauseOf`). The whole note, every sentence, is shown in the sim's rotation drawer (`web/src/lib/sim/rotations.ts`), which carries no level, so every sentence counts at every level the row renders.

## Changes since round 19's tree (6beb919c2)

`git diff 6beb919c2 HEAD`: Fury Death Wish note reworded (platform dd169626d); Fury Berserker Rage note "From 36, ..." with a `currentLevel >= 36` gate; Fury Overpower gated `currentLevel < 36`; Protection Bloodrage note (R19-2 fix); fury.md:134 Berserker Stance / Berserker Rage sentence. `addon-data.json` regenerated: Fury Berserker Rage now renders from band 36, Overpower on bands 12 to 32 only. Warrior orders in `addon-builds.json` unchanged. All re-checked below.

**The guides-handoff (Fury Berserker Rage and Overpower level gates):** fixed and right. The row starts at band 36, Improved Berserker Rage is the order's 34 and 35 points, Whirlwind trains at 36 and the engine's default stance moves there, so "From 36, under 80 rage: Improved Berserker Rage makes it 10 rage, back every 30 seconds" is true on every band it renders (spelltip 18499: 30 s cooldown). Overpower "Only before level 36" now renders on 12 to 32 only (band 32 covers 32 to 35). fury.md:134 "Berserker Stance and Berserker Rage train at 30, but Berserker Rage only works in that stance, so leave it on the bar until 36" agrees with the card, the order and `trainables` (18499 at 30, 1680 at 36). Not re-filed.

## Orders (index 0 = level 10), decoded against tree, tier and column

- Arms leveling: Improved Heroic Strike 10; Deflection 11-13; Improved Rend 14-16; Improved Charge 17-18; Improved Tactical Mastery 19-21; Improved Overpower 22-23; Anger Management 24; Deep Wounds 25-27; Two-Handed Spec 28-30; Impale 31-32; Bloodthrill 33-34; Sweeping Strikes 35; Weaponmaster 36-39; Mortal Strike 40; Weaponmaster 41; Improved Slam 42-43; Cruelty 44-48; Unbridled Wrath 49-53; Improved Bloodrage 54-55; Iron Will 56-60.
- Arms raid (a full 10 to 60 order a player can level with): Improved Heroic Strike 10-12; Improved Rend 13-15; Improved Tactical Mastery 16-18; Improved Overpower 19-20; Anger Management 21; Deep Wounds 22-24; Spearing Strike 25; Two-Handed Spec 26-28; Impale 29-30; Bloodthrill 31-32; Sweeping Strikes 33; Weaponmaster 34-38; Improved Slam 39; Mortal Strike 40; Improved Slam 41; Cruelty 42-46; Unbridled Wrath 47-51; Enrage 52-56; Reckless Swings 57; Improved Execute 58-59; Improved Bloodrage 60.
- Fury leveling = Fury raid: Cruelty 10-14; Unbridled Wrath 15-19; Furious Precision 20-22; Enrage 23-24; Dual Wield Spec 25-29; Death Wish 30; Reckless Swings 31; Improved Execute 32-33; Improved Berserker Rage 34-35; Improved Whirlwind 36-37; Flurry 38-39; Bloodthirst 40; Flurry 41; Improved Heroic Strike 42-44; Improved Rend 45-47; Improved Charge 48-49; Improved Tactical Mastery 50-52; Anger Management 53; Deep Wounds 54-56; Impale 57-58; Improved Bloodrage 59-60.
- Protection leveling = Protection raid: Improved Bloodrage 10-11; Shield Specialization 12-16; Anticipation 17-21; Improved Revenge 22-24; Improved Thunder Clap 25-27; Last Stand 28; Master of Defense 29-30; Defiance 31-33; Improved Sunder Armor 34-36; Concussion Blow 37; Focused Rage 38-39; Shield Slam 40; Focused Rage 41; Bastion 42-46; Improved Heroic Strike 47-49; Deflection 50-54; Improved Rend 55-57; Deep Wounds 58-60.

Trainer levels as round 19 (Rend 4, Overpower 12, Execute 24, Slam 20, Berserker Stance / Berserker Rage 30, Whirlwind 36, Recklessness 50, Revenge and Demoralizing Shout 14, Shield Block 16, Shield Wall 28).

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| WARRIOR-R20-1 | CONFIRMED | card note (rank value) | `warrior-arms.json`:207 Overpower, first sentence, rendered on the addon row at every band 12 to 60 | "...; with Improved Overpower it crits most of the time." | Improved Overpower is two ranks: 12290 rank 1 "by 25%", rank 2 "by 50%" (`talents/warrior.json`). The leveling order's first point is level 22 and the second 23, so a level 22 player on band 22 has one point; the raid order's first point is 19 (band 16 covers 16 to 19), the second 20. With one point Overpower crits at 25% plus a level 19 to 22 warrior's sheet crit (5% base plus Agility, well under 10%), about one in three: not "most of the time". The conditional is met at one point and the value is rank 2's, the R19-2 / D2 shape. Round 18 (R18-2) called this clause "conditional and holds" without splitting the ranks. | "The moment it lights up, even ahead of Mortal Strike from 40: a dodge opens it for a few seconds, and a Bloodthrill proc does too once the talent is in; with both points of Improved Overpower it crits most of the time." |
| WARRIOR-R20-2 | CONFIRMED | card note (rank value) | `warrior-fury.json`:74 Bloodrage, the whole note, rendered at every band 10 to 60 | "Whenever rage is under 80: with Improved Bloodrage, 15 rage at once and 15 more over 10 seconds, back every minute." | Improved Bloodrage is the Fury order's 59 and 60 points (both orders). Band 56 covers 56 to 59 (`Rotation.lua` `bandFor`), so a level 59 Fury warrior reads this row with one point: rank 1 (12301) "by 25%", 12.5 and 12.5, not 15 and 15. The conditional is met and the number is rank 2's: the R19-2 shape (the Protection twin of this note was fixed in round 19 for exactly this). Round 19 left the Fury note a Should because no band shows one point, but the band covers level 59. Sim round 20 handed it back to the guide lane as optional. | "Whenever rage is under 80: 10 rage at once and 10 more over 10 seconds, back every minute, a quarter more from your first point of Improved Bloodrage at 59 and half again from 60." |
| WARRIOR-R20-3 | CONFIRMED | card note (value vs build) | `warrior-arms.json`:138 Bloodrage, first sentence, rendered on the addon row at every band; the talent sentence is the second and never reaches the row | "Whenever rage is under 80: 10 rage at once and 10 more over 10 seconds, back every minute." | The level 60 (leveling) order takes Improved Bloodrage at 54 and 55, the raid order at 60. On bands 54, 56 and 60 (leveling) and band 60 (raid) the row the player reads quotes the untalented 10 and 10 while their build makes it 12.5 or 15 each (12301: 25% / 50%). The quarter / half again correction sits only in the second sentence, which the addon never shows. The round 20 rule: a quoted value must match the talents in the build at the band it renders. | "Whenever rage is under 80: 10 rage at once and 10 more over 10 seconds, back every minute, a quarter more for each point of Improved Bloodrage. The level 60 build takes two points at 54 and 55, the raid build one at 60." |
| WARRIOR-R20-4 | CONFIRMED | card note (false mechanic; round 18 regression) | `warrior-arms.json`:184 Rend, first sentence, rendered at every band 10 to 60 | "Keep it ticking: the bleed ignores armor, and once Bloodthrill is in, only a target bleeding from your Rend lights Overpower." | A dodge lights Overpower on any target, bleeding or not (7384: "Only useable after the target dodges"); Bloodthrill (1289682) only adds a chance on main-hand hits "against enemies afflicted by your Rend". The same card's Overpower row says "a dodge opens it", and arms.md:87 has it right ("Bloodthrill only lights Overpower on a bleeding target"). Round 18's R18-1 rewrite turned the old, correct "Bloodthrill only lights Overpower on a target that is bleeding" into "only a target bleeding from your Rend lights Overpower", which credits Rend with every Overpower from 33 on. A wrong earlier-round fix is a confirmed finding. | "Keep it ticking: the bleed ignores armor, and once Bloodthrill is in, your main-hand hits on a target bleeding from your Rend can light Overpower too." |
| WARRIOR-R20-5 | CONFIRMED | card note (later sentence) | `warrior-arms.json`:239 Execute, second sentence, rendered with the row on bands 24 to 38 (levels 24 to 39) | "Below 20%. Your first one spends the full bar you walked in with, so arrive with rage; after that it shares the globals with Mortal Strike." | Execute trains at 24; Mortal Strike is the 40 point in both Arms orders. On every level 24 to 39 the note pairs Execute with a button the player does not have. Round 18 fixed this exact credit on the Overpower note ("from 40") and round 19 left this one as "second-sentence text"; the round 20 rule covers every sentence. | "Below 20%. Your first one spends the full bar you walked in with, so arrive with rage; after that it shares the globals with Mortal Strike from 40." |
| WARRIOR-R20-6 | CONFIRMED | card note (later sentence) | `warrior-arms.json`:327 Heroic Strike, second sentence, rendered with the row on every band 10 to 60 | "Below that, keep the rage for Mortal Strike, Overpower and Execute, which lose damage when they wait." | Overpower trains at 12, Execute at 24, Mortal Strike is the 40 point. At 10 the player has none of the three, and Mortal Strike is missing on every level to 39. | "Below that, keep the rage for the attacks that lose damage when they wait: Overpower from 12, Execute from 24 and Mortal Strike from 40." |
| WARRIOR-R20-7 | CONFIRMED | card note (later sentence) | `warrior-arms.json`:324 Hamstring, second sentence, rendered with the row on every band 10 to 60 | "Without the totem it is rage Mortal Strike and Overpower need." | Hamstring renders from band 10 (trains at 8). Overpower is 12, Mortal Strike 40, so levels 10 to 39 read a reason built on a button they lack. | "Without the totem it is rage your other attacks need." |
| WARRIOR-R20-8 | CONFIRMED | card note (raid order) | `warrior-arms.json`:255 Slam, first sentence, rendered from band 20 | "Between Mortal Strikes, only with Improved Slam: the talent keeps your swing timer running through the cast." | On the raid order Improved Slam's first point is 39 and Mortal Strike 40. At 39 (band 38) the card's own condition (`spellCastTime` under 1.5 s; rank 1 takes 0.25 s off) fires Slam, and the note places it between Mortal Strikes the player does not have. One level, raid order only (the leveling order has Mortal Strike at 40 before Improved Slam at 42), the case the round 20 notes ask to check on both orders. | "Only with Improved Slam, and between Mortal Strikes from 40: the talent keeps your swing timer running through the cast." |
| WARRIOR-R20-9 | CONFIRMED | card note (level 60 figure on lower ranks) | `warrior-protection.json`:278 Demoralizing Shout, first sentence, rendered on bands 14 to 60 | "Keep it on the boss: 204 less melee attack power at 60 takes a real slice off every swing." | The rows cast rank 1 (1160, 49) at 14 to 22, rank 2 (6190) at 24 to 32, rank 3 (11554) at 34 to 42, rank 4 (11555) at 44 to 52, and rank 5 (11556) only from 54. The round 20 rule says plainly: "quote no level-60 figure (cost, damage, duration, ...) on a band where a lower rank ... render[s]". Round 19 held it as a Should because the note names its level; the round 20 lens rules that out, as round 19 D2 did for the Protection Bloodrage note. The page prose (protection.md:93) is the level 60 paragraph and keeps its number. | "Keep it on the boss: less attack power on the boss takes a real slice off every swing. One is enough, since it does not stack with a druid's Demoralizing Roar. Refresh it in its last 3 seconds." |

Every fix needs `addon-data.json` and `addon/ForeverSixty/Data.lua` regenerated (note text only; no cast moves).

Counts: CONFIRMED 9, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. Page prose on all four pages: 0 errors.

## Shoulds (wording, not findings)

- **Arms Reckless Swings note** (:229, "your white swings for the next 8 seconds each bring 5 more rage"): the talent is "Your next 4 melee attacks ... Lasts 8 sec." With a slow two-hander four swings never fit into 8 seconds, so it is not false, but with Windfury Totem extra attacks the cap can bite. "your next four white swings within 8 seconds each bring 5 more rage" matches the Fury note and the tooltip.
- **Protection Shield Block note** (:182, "75% more block chance for 7 seconds"): the tooltip adds "but will only block 2 attacks", which protection.md:127 says. "75% more block chance for up to two attacks in 7 seconds" would match.
- **Integration Should, Fury Execute "First"**: "Before any other attack" reads better on the 24 to 35 bands where the card sits in Battle Stance. Not false.
- Carried from round 19: Protection Defensive Stance note (10 and 12 name Revenge, Shield Block, Shield Wall); Arms Hamstring "never from 25% down" before Execute (sim round 20 G4: no Windfury Totem exists before 32, so it never matters); fury.md:112 "edges Night Elf by a hair"; protection.md:113 cross-faction "no race's health swings less"; arms.md:125 skipping the level 19 point; fury.md:74 "Fury never casts Rend" next to the leveling Rend advice; arms.md:115 "other pair"; fury.md:54 "debuff"; fury.md:100 Edgemaster wording; the "From 36" lead on the level 60 page card (wow-player nit in sim round 20).

## Checked and right this round

- **Every quoted value against the current builds' talents.** Execute "10 rage" with Improved Execute (fury.md:69, arms.md:79: 15 minus 5). Heroic Strike "cost 3 less" (Improved Heroic Strike 3 in the raid Arms, Fury and Protection builds; the Arms level 60 build has one point and arms.md:125 says 1). Charge 21 rage with Improved Charge (arms.md:71, the level 60 build only; fury.md:78 "6 extra"). Tactical Mastery 25 with Improved Tactical Mastery 3 (Arms note :79, arms.md:70 and :91, both Arms orders by 21 / 18). Revenge "nearly free" (5 minus Focused Rage 3, protection.md:76). Thunder Clap "6 less" and the Improved Shield Wall swap "4 more" (protection.md:75, :83). Shield Wall "15 minute cooldown" (spelltip 871 900000 ms; the Protection build has no Improved Shield Wall). Whirlwind "every 8 seconds instead of 10" (Improved Whirlwind 2). Reckless Swings 15 s cooldown, 25% of base health (spelltip 1329555: `cooldown_ms` 15000, `cost_pct` 25). Berserker Rage 30 s, 10 s immunity (18499). Death Wish 30 s, 3 min, 10 rage (12328). Bloodrage 60 s (2687). Last Stand 30% for 20 s (12975). Shield Block 75%, 7 s, 5 s (2565). Thunder Clap 20% slower attacks, 4 targets (6343, every rank). Battle Shout 10 rage, 20 yards, 3 min (6673 to 11551). No quoted value is a base value the build changes, apart from R20-2, R20-3 and R20-1.
- **Fury notes, every band 10 to 60** (see the Note sweep): Death Wish's new wording matches its card (`remainingTime >= 210s` or E20 or `<= 30s`) and fury.md:92. Execute, Bloodthirst, Whirlwind, Reckless Swings, Berserker Rage, Overpower, Heroic Strike hold on both (identical) orders.
- **Protection notes**: the R19-2 fix is right (12301 rank 1 at 10, rank 2 at 11; the row starts at 10). Every other note as in the sweep.
- **Page prose, line by line**: index.md:22 to :39; arms.md:42 to :131 (the talent list against the raid and level 60 builds, the five rotation paragraphs against the card, stats with no ranking words, races, professions, both leveling blocks against the leveling order: Improved Overpower 22, Bloodthrill 33 after Deep Wounds 25 to 27 and Impale 31 to 32, Sweeping Strikes 35 and Whirlwind 36 "once you have both"); fury.md:41 to :134 (including the new :134 sentence, Furious Precision at 20 with Dual Wield, Cruelty 10 to 14 and Unbridled Wrath 15 to 19); protection.md:49 to :135 (Shield Slam the 31st point at 40, Focused Rage's last point 41, Bastion from 42, then Improved Heroic Strike, Deflection, Improved Rend, Deep Wounds). No page names 70755; no em or en dashes on any page or note.
- **Arms raid order against every Arms note** (the second order a player can level with): Spearing Strike from 25 and Reckless Swings from 57 ("raid build only", after Mortal Strike), Bloodthrill 31 and 32 (Rend and Overpower notes conditional), Improved Tactical Mastery by 18 (Berserker Stance note from 50), Improved Execute 58 and 59 (no cost quoted), Improved Bloodrage at 60 (R20-3), Improved Overpower 19 (R20-1), Improved Slam 39 (R20-8).

## Decisions

- **D1. R20-1 and R20-2 follow round 19 D2**: a conditional met at one point while the note quotes the second rank's value is false at that level. Round 18 and round 19 checked only whether the conditional was met, not by which rank.
- **D2. R20-3 is filed** because the talent sentence never reaches the addon row: the row the player reads at 54 to 60 states a value their build no longer has. The round 20 notes name this check (quoted values against the talents in the current builds).
- **D3. R20-5 to R20-7 are filed under the every-sentence rule.** The full note renders in the sim drawer with no level attached, and round 18 already fixed the identical "Mortal Strike before 40" credit in the Overpower note's first sentence.
- **D4. R20-9 is filed under the round 20 rule as written**, overturning round 19's Should. The note names "at 60", but the rule forbids the level 60 figure on lower-rank bands, not only an unlabelled one.
- **D5. Rows that render before their talent or spell exists stay with the component lane** (Fury Death Wish and Reckless Swings at 10 to 30, Arms Reckless Swings and Spearing Strike on the leveling order, Protection Last Stand at 10 to 26), as in rounds 15 to 19.
- **D6. The guides-handoff is closed**: Fury Berserker Rage and Overpower read right on every band, on the page and in fury.md:134.

## Sims requested

- None new. The data settles every finding (no cast moves).
- Carried, only if the wording is kept: Fury Alliance Skyborne against Night Elf at 30,000 iterations (fury.md:112); Protection race TMI on one shared raid preset (protection.md:113). WARRIOR-13 (Improved Shield Wall) stays with the owner.

## Engine questions

- None new. Carried for the sim round 21 integrator: X1 (`spellIsKnown` true below RequiredLevel; Whirlwind at 30), P5 threat constants, AC Arcanite Champion PPM, C1 Eureka! with Execute (#439), C2 Deep Wounds' double attacker multiplier, and whether an outright-missed Heroic Strike spends a Reckless Swings charge.

## Sim findings

- None new. Sim round 20's G1 to G5 are in and read right on the regenerated rows.

## Sources

- Client 1.60.1.70790: `talents/warrior.json` (Improved Overpower 12290, Improved Bloodrage 12301, Improved Slam 12862, Bloodthrill 1289682, Reckless Swings 1329555 rank text), `spelltips/warrior.json` (7384, 2687, 1160, 11556, 871, 2565, 1329555, 18499, 12328, 6343, 6673), `trainables/warrior.json`, `addon-builds.json` (both orders per spec), `addon-data.json` (`rotations/warrior-*`).
- Curated: `data/curated/apl/warrior-{arms,fury,protection}.json`.
- Site and addon code: `web/src/lib/guides/rotation-lists.ts` (`clauseOf`), `web/src/lib/sim/rotations.ts` (drawer shows the whole note), `data/pipeline/addonrotation.py` (`_one_line`), `addon/ForeverSixty/Rotation.lua` (`bandFor`).
- Reviews: `design/reviews/guide-accuracy/round19/warrior.md`, `round18/warrior.md` (R18-1, R18-2), `ledger.md` (Rounds 18 and 19), `design/reviews/2026-11-04-sim-audit-round20/warrior.md` (G1 to G5, S1) and `integration.md`.

## Note sweep

One line per row note: spec, row, bands it renders, each named talent / spell / rank / value with the level the player has it (L = leveling order, R = raid order), verdict.

Arms

- Berserker Stance (bands 30 to 60): "From 50" (Recklessness 50); Improved Tactical Mastery 3 (L 21, R 18) "25 rage"; Mortal Strike (40) and Execute (24) "still work there", Overpower and Rend do not. All after 50. OK.
- Bloodrage (10 to 60): 10 + 10 over 10 s, 60 s; Improved Bloodrage L 54 / 55, R 60, only in sentence 2. R20-3.
- Battle Shout (10 to 60): 10 rage, 20 yards, raid (every rank). OK.
- Rend (10 to 60): bleed ignores armor (from 4); Bloodthrill L 33, R 31, conditional; "only a target bleeding ... lights Overpower". R20-4.
- Overpower (12 to 60): Mortal Strike "from 40" (L/R 40); dodge opens it, a few seconds (5 s); Bloodthrill "once the talent is in"; Improved Overpower rank 1 L 22, R 19, rank 2 L 23, R 20, "crits most of the time". R20-1.
- Mortal Strike (40 to 60): biggest hit outside Execute; 50% healing cut (12294 text). OK.
- Reckless Swings (10 to 60, component row; talent R 57): "raid build only", after Mortal Strike (40), 8 s, 5 rage per white swing, a quarter of base health (`cost_pct` 25). OK (Should on the four-attack cap).
- Execute (24 to 60): below 20% (5308 at 24); first one spends the bar (converts extra rage); "shares the globals with Mortal Strike" before 40. R20-5.
- Slam (20 to 60): Improved Slam rank 1 L 42, R 39 (0.25 s off, "no longer interrupts" at rank 1, so the 1.5 s gate fires); "between Mortal Strikes" at R 39 without Mortal Strike. R20-8.
- Spearing Strike (10 to 60, component row; talent R 25): raid build only, Battle Stance, 75% weapon damage, plus 75% on Giants and Dragonkin. OK.
- Hamstring (10 to 60): Windfury Totem (shaman 32), 10 rage (1715), "never from 25% down" (Should, carried; no totem before 32); sentence 2 "Mortal Strike and Overpower" before 12 / 40. R20-7.
- Heroic Strike (10 to 60): above 90 rage (card `>= 90`); sentence 2 "Mortal Strike, Overpower and Execute" before 12 / 24 / 40. R20-6.

Fury (leveling and raid orders identical)

- Bloodrage (10 to 60): "with Improved Bloodrage, 15 and 15", rank 1 at 59 (band 56), rank 2 at 60. R20-2.
- Death Wish (10 to 60; talent 30, rows 10 to 28 component): pull if a second fits, execute phase, last 30 s; matches the card. OK.
- Recklessness (50 to 60): execute or last 15 s; once a fight (30 min). OK.
- Battle Shout (10 to 60): as Arms. OK.
- Execute (24 to 60): first below 20%; takes every point of rage; Bloodthirst "from 40" (40). OK.
- Bloodthirst (40 to 60): on cooldown. OK.
- Whirlwind (36 to 60): four targets, both weapons (Dual Wield 20); behind Bloodthirst "from 40". OK.
- Reckless Swings (10 to 60; talent 31): behind Whirlwind "once you have it" (36); four white swings, 5 rage, a quarter of base health. OK.
- Berserker Rage (36 to 60): "From 36"; Improved Berserker Rage rank 2 at 35, 10 rage; 30 s; fears on Magmadar, Onyxia, Nefarian (18499 Fear immunity). OK.
- Overpower (12 to 32, levels 12 to 35): Battle Stance before 36; Whirlwind 36. OK.
- Heroic Strike (10 to 60): above 60; 25% stop only once Execute (24) is known, matching the `spellIsKnown(20662)` branch; queued, no global. OK.

Protection (leveling and raid orders identical)

- Defensive Stance (10 to 60): threat bonus (+30%); Revenge 14, Shield Block 16, Shield Wall 28 "all need it". OK (Should, carried).
- Last Stand (10 to 60; talent 28, rows 10 to 26 component): below 35%; 30% for 20 s, lost after. OK.
- Shield Wall (28 to 60): below 40%; 60% for 12 s; 15 min (no Improved Shield Wall in the build). OK.
- Bloodrage (10 to 60): 10 + 10; Improved Bloodrage rank 1 at 10 "a quarter", rank 2 from 11 "half again"; under 60 rage. OK.
- Shield Block (16 to 60): 75% for 7 s, off the global; every block lights Revenge (14). OK (Should on the two-attack cap).
- Revenge (14 to 60): dodge, parry or block opens it; threat per rage (5 rage, 2 with Focused Rage from 41). OK.
- Shield Slam (40 to 60): biggest hit; adds block value (23922). OK.
- Thunder Clap (10 to 60): attack speed slow (20% at every rank from rank 1, 6343); refresh under 3 s. OK.
- Demoralizing Shout (14 to 60): "204 ... at 60" on ranks 1 to 4 (bands 14 to 52); sentence 2 druid Roar does not stack; refresh under 3 s. R20-9.
- Battle Shout (10 to 60): only when nobody else's is up; same aura at 60. OK.
- Sunder Armor (10 to 60): five stacks (7386 "up to 5 times"), then threat. OK.
- Heroic Strike (10 to 60): above 40 rage, queued, no global. OK.

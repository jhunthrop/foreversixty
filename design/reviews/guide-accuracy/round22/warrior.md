# Guide accuracy round 22: Warrior

Audited on the round 22 audit tree (987d08a7c: main after guide round 21; no sim round 21 commit in the log, so there is no sim round 21 class report or integration.md to read; engine pin d5c708834 in `sim/enginever/version.go`, engine source at `gr22/engine`) against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/warrior/{index,arms,fury,protection}.md`, every note in `data/curated/apl/warrior-{arms,fury,protection}.json`, every rendered row of `addon-data.json` (`rotations/warrior-arms` 20 bands, `warrior-fury` 19, `warrior-protection` 24), both shipped orders per spec in `addon-builds.json` (re-decoded against tier and column), `talents/warrior.json`, `spelltips/warrior.json`, `trainables/warrior.json`, `races.json` and the engine's warrior spell code. I read the round 21 warrior report and the ledger's Round 19, 20 and 21 sections.

Rendering as rounds 20 and 21 documented: the addon row carries the note up to the first ". " at every band; the guide card prints up to the first ": ", ". " or "; "; the sim drawer shows the whole note with no level, so every sentence counts at every band.

## Changes since round 21's tree (e039e1490)

`git diff e039e1490 HEAD` for warrior: the Arms Mortal Strike note (`warrior-arms.json`, and the `sim/request` copy), arms.md:42, and the regenerated `addon-data.json` / `Data.lua` text. Orders in `addon-builds.json` unchanged.

**Round 21's fixes, re-checked as new work:**

- R21-1 Arms Mortal Strike note "On cooldown: a big hit every 6 seconds, and it cuts the healing the target takes in half.": right at every band 40 to 60. Every rank's cooldown is 6 s (12294 `cooldown_ms` 6000, trainer row 6000 for all four ranks) and no talent in either Arms build touches it; 50% healing cut for 10 s at every rank. "a big hit" makes no comparison. Addon rows carry the new text at all ten bands; no row still says "biggest hit outside Execute" (`addon-data.json` and `Data.lua`: 0 matches).
- R21-2 arms.md:42 "a full bar of rage spent on one huge Execute": right. It matches arms.md:48 and :91 (first Execute in Battle Stance spends the bar) and makes no cross-spec claim.

## Orders (index 0 = level 10)

Unchanged from round 21: Arms leveling Improved Overpower 22-23, Anger Management 24, Deep Wounds 25-27, Two-Handed Specialization 28-30, Impale 31-32, Bloodthrill 33-34, Sweeping Strikes 35, Weaponmaster 36-39 and 41, Mortal Strike 40, Improved Slam 42-43, Improved Bloodrage 54-55; Arms raid Improved Overpower 19-20, Spearing Strike 25, Bloodthrill 31-32, Sweeping Strikes 33, Improved Slam 39 and 41, Mortal Strike 40, Enrage 52-56, Reckless Swings 57, Improved Execute 58-59, Improved Bloodrage 60; Fury (both identical) Cruelty 10-14, Unbridled Wrath 15-19, Furious Precision 20-22, Enrage 23-24, Dual Wield Specialization 25-29, Death Wish 30, Reckless Swings 31, Improved Execute 32-33, Improved Berserker Rage 34-35, Improved Whirlwind 36-37, Flurry 38-39 and 41, Bloodthirst 40; Protection (both identical) Improved Bloodrage 10-11, Shield Specialization 12-16, Anticipation 17-21, Improved Revenge 22-24, Improved Thunder Clap 25-27, Last Stand 28, Master of Defense 29-30, Defiance 31-33, Improved Sunder Armor 34-36, Concussion Blow 37, Focused Rage 38-39 and 41, Shield Slam 40, Bastion 42-46, then Improved Heroic Strike, Deflection, Improved Rend, Deep Wounds.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| WARRIOR-R22-1 | CONFIRMED | card note (superlative fails at 60 against another spell on the same card) | `warrior-protection.json` Shield Slam note, the whole note, rendered on the addon row at every band 40 to 60 and in the sim drawer | "On cooldown: your biggest hit, and it grows with your block value." | Revenge carries a quarter of attack power on every rank in Forever: Wowhead Forever tooltip for the rank the card casts at 54 to 60 (11601) reads "[109.07 + ((0.25 * 1) * Attack Power)] to [132.93 + ((0.25 * 1) * Attack Power)] damage"; engine `revenge.go` `revengeAttackPowerCoefficient = 0.25` (client 1.60.1.70707 APRatio). Both Protection orders hold Improved Revenge 3 from 24: 60% more Revenge damage (`talents/warrior.json`; engine `talents.go:286`, applied to the whole hit). Shield Slam at 60 (23925) is 640 to 670 plus block value with no attack power term (`spelltips` 23925, engine `shield_slam.go:84`). Revenge outhits it once (121 + 0.25 AP) x 1.6 > 655 + block value, that is attack power above about 1,150 + 2.5 x block value. The engine's own level 60 Phase 1 tank (`tank_warrior/TestP1TankWarrior.results`, CharacterStats) has 1,713 attack power and 77 block value: Revenge about 880 a hit against Shield Slam's 732. The page's raid preset stacks Battle Shout, Blessing of Might, Strength of Earth Totem, Trueshot Aura, Mark of the Wild and Smoked Desert Dumplings (`bis/warrior-protection.json` presets), so a raid tank at 60 is past the threshold and Revenge, not Shield Slam, is the biggest hit on the card. Bastion multiplies both, crits scale both alike. Below 60 Shield Slam stays ahead (54: rank 3 580 + BV against rank 5 needs about 1,150 AP; 40 to 53 needs more), so the note is false exactly where the page plays. Round 21's Should on this note looked only at Execute and used the flat Revenge base. | "On cooldown: a hard hit with a very high amount of threat, and it grows with your block value." |

Text only. Needs `addon-data.json` and `addon/ForeverSixty/Data.lua` regenerated (note text only; no cast moves).

Counts: CONFIRMED 1, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## Shoulds (wording, not findings)

- **Levels 20 to 60 dates (round 22's warlock shape, checked on all three pages; none is false, so none is filed):**
  - arms.md:131 "Sweeping Strikes into Whirlwind once you have both": the paragraph dates Improved Overpower (22), Bloodthrill (33) and Mortal Strike (40) but not these. Gated by "once you have both", so it credits nothing early. Optional: "Sweeping Strikes (35) into Whirlwind (36) for pulls of two or three." Deep Wounds and Impale are named only as "before 33", which is true (25 to 27, 31 to 32).
  - fury.md:134 "Fury levels fastest once Bloodthirst and Whirlwind are in": Whirlwind is dated (36), Berserker Stance and Berserker Rage (30) too, Bloodthirst is not. Gated by "once ... are in", so not false. Optional: "once Whirlwind (36) and Bloodthirst (40) are in".
  - protection.md:135: every talent dated or ordered (Shield Slam 40, Focused Rage 41, Bastion, then the Arms points in order). Clean.
  - Levels 1 to 19 (arms.md:121 to :127, fury.md:124 to :130, protection.md:125 to :131): every spell dated against `trainables` (Charge 4, Rend 4, Thunder Clap 6, Hamstring 8, Bloodrage 10, Defensive Stance / Sunder Armor / Taunt 10 by quest, Shield Bash 12, Overpower 12, Revenge 14, Demoralizing Shout 14, Tactical Mastery 14, Shield Block 16, Dual Wield 20 as decided in round 4/16). Clean.
- **protection.md:69** "Shield Slam ... Your biggest button": reads as importance, not size, so not filed; if the note changes for R22-1, "your key button" keeps the two in step.
- **protection.md:75** "These are the three buttons you press most after Shield Slam": Sunder Armor on every free global is pressed more often than Shield Slam's 6 s cooldown. "the three buttons you lean on most beside Shield Slam" is safer.
- Carried from round 21 / 20: Arms Reckless Swings note "white swings for the next 8 seconds" (client: next 4 melee attacks within 8 s); Protection Shield Block note's two-attack cap; Fury Execute "First"; Protection Defensive Stance note at 10 to 27 names Revenge, Shield Block and Shield Wall before they train (not a comparison; true of the spells); Arms Hamstring "never from 25% down" before Execute; fury.md:112 "edges Night Elf by a hair"; protection.md:113 "no race's health swings less"; arms.md:125 skipping the level 19 point; fury.md:74 "Fury never casts Rend"; arms.md:115 "other pair"; fury.md:54 "debuff"; fury.md:100 Edgemaster wording; "From 36" lead on the level 60 card.

## Checked and right this round

- **Superlatives and comparisons, every one on all four pages, against every spell the player has at the level (including Execute, Recklessness, Death Wish, Spearing Strike):** arms.md:46 / note "a big hit every 6 seconds" (no comparison); arms.md:48 and :91 "one huge Execute", "that first Execute is your biggest hit of the fight" (rank 5 Execute with a full bar: 600 + 15 a point against a Mortal Strike crit at 220%); arms.md:55 "fewer swings and lumpier rage than a dual-wielder"; fury.md:52 "gears up slower than Arms"; fury.md:64 "the button you never delay" and fury.md:96 / note "the button you lose the most by delaying" (Bloodthirst 45% AP + 30 on 6 s against Whirlwind with Improved Whirlwind on 8 s: per cooldown second Bloodthirst is ahead); fury.md:74 "the best use of the Arms points" (measured build); fury.md:84 Bloodrage with Improved Bloodrage 15 + 15 "more rage than a Charge" (21 with Improved Charge); protection.md:51 "huge threat for almost no rage"; Revenge note "most threat per rage you have" (5 rage, 2 with Focused Rage, against Sunder 15/12/9, Shield Slam 20/17, Thunder Clap 20/14/11); protection.md:113 Night Elf "less damage than any other Alliance race" (measured, carried). Only the Shield Slam note fails (R22-1).
- **Every quoted cost, duration and cooldown at the ranks the rows cast, with the build's discounts:** Mortal Strike 6 s every rank; Execute 15, 10 with Improved Execute 2 (Fury 32, Arms raid 58); Heroic Strike 3 less (Improved Heroic Strike 3 on Fury, Arms raid, Protection from 49) and 1 less in the Arms level 60 build; Spearing Strike 15 rage, 20 s, Battle Stance only (engine `spearing_strike.go:60`, ShapeshiftMask Battle Stance); Reckless Swings 15 s, a quarter of base health (`cost_pct` 25), 8 s; Death Wish 10 rage, 30 s, 3 min; Recklessness 15 s, 30 min; Berserker Rage 30 s, 10 rage with Improved Berserker Rage 2; Bloodrage 60 s, 10 + 10, 25% / 50% per point; Victory Rush 10% health, 30 s cooldown, usable 20 s after a kill (402927); Charge 15 at rank 3 (46), 21 with Improved Charge; Tactical Mastery 10 plus 15; Slam 15 s, 10 s and 1 s cast with Improved Slam 2; Shield Block 75%, 7 s, 5 s, two attacks; Shield Wall 60%, 12 s, 15 min (4 min with Improved Shield Wall 2); Last Stand 30%, 20 s; Challenging Shout 6 s, 10 min; Demoralizing Shout 204 at 60; Thunder Clap 6 less with Improved Thunder Clap 3, 4 more with one point; Sunder Armor 3 less; Focused Rage 3 less; Hamstring 10 rage (Arms has no discount on it); Anger Management 10 rage every 30 s.
- **Notes that name another spell, against its trainer level on every band:** Arms Rend (10 to 60) names Overpower and Bloodthrill only behind "once Bloodthrill is in" (31 / 33); Arms Overpower, Execute, Slam and Heroic Strike name Mortal Strike, Execute and Overpower with "from 40 / 24 / 12"; Arms Reckless Swings names Mortal Strike, raid build only (talent 57); Arms Berserker Stance "From 50" names Mortal Strike and Execute (both trained by 50); Fury Execute and Whirlwind name Bloodthirst "from 40"; Fury Reckless Swings names Whirlwind "once you have it"; Fury Heroic Strike names Execute "once you have Execute"; Fury Overpower names Whirlwind as the reason it stops (36); Protection Shield Block names Revenge (14 < 16). Protection Defensive Stance (Should, carried).
- **index.md** every line: Tactical Mastery 14 and 10 rage (1310185); Victory Rush 20; Slam 15 s and Improved Slam 10 s / 1 s; Thunder Clap in Defensive Stance; Shield Block 7 s / 5 s; Shield Wall 60% / 12 s / 15 min and Improved Shield Wall to 4 min; Death Wish 5% more damage taken; Battle Shout 20 yards; the named new talents exist in `talents/warrior.json`.
- **Page prose, line by line:** arms.md:42 to :131, fury.md:41 to :134, protection.md:47 to :135 against the cards, both orders, `trainables`, talent text and race rows. No page names 70755; no em or en dashes.

## Decisions

- **D1. R22-1 is CONFIRMED on the client tooltip and the engine's own level 60 tank stats, not a sim.** The Revenge formula (tooltip, 0.25 x AP) and Improved Revenge's 60% are client values; Shield Slam has no attack power term. The comparison reduces to the tank's attack power against about 1,150 + 2.5 x block value, and the engine's Phase 1 tank sits at 1,713 AP and 77 block value with fewer raid buffs than the page's preset. Round 21 filed this note's superlative as a Should on the strength of Revenge's flat base alone; the attack power share overturns that at 60.
- **D2. The round 22 leveling-paragraph shape (a spell or talent named without its level where the paragraph dates the rest) is found on arms.md:131 and fury.md:134, but both are gated ("once you have both", "once ... are in") and credit nothing before it is had,** so they are Shoulds, not findings. The warlock case was filed because it credited Death Coil across the whole range.
- **D3. Rows that render before their talent or spell exists stay with the component lane** (Fury Death Wish and Reckless Swings at 10 to 30, Arms Reckless Swings and Spearing Strike on the leveling order, Protection Last Stand at 10 to 26), as in rounds 15 to 21.

## Sims requested

- Optional confirmation of R22-1: Protection raid preset at 60, average landed Revenge (11601) against average landed Shield Slam (23925) from the ability breakdown. The data already decides the call; this only records the margin.
- Carried, only if the wording is kept: Fury Alliance Skyborne against Night Elf at 30,000 iterations (fury.md:112); Protection race TMI on one shared raid preset (protection.md:113).

## Engine questions

- **New, W22-1:** Revenge's rank at 60. `trainables/warrior.json` lists rank 6 (25288) at 60 from the trainer (`source: skill_line_ability`), but the engine (`revenge.go:8`) treats 25288 as an Ahn'Qiraj book rank and casts rank 5 at 60 with `IncludeAQ` off, and the card names 11601. Confirm which the Forever trainer gives; if rank 6 trains at 60, the card and engine leave about 30 damage a Revenge on the table (and R22-1's margin grows). Not a guide finding.
- Carried: W21-1 Classic weapon speeds in the engine's item database; X1 (`spellIsKnown` below RequiredLevel), P5 threat constants, AC Arcanite Champion PPM, C1 Eureka! with Execute (#439), C2 Deep Wounds' double attacker multiplier, whether an outright-missed Heroic Strike spends a Reckless Swings charge.

## Sim findings

- None new.

## Sources

- Wowhead Forever tooltip, Revenge rank 5: https://nether.wowhead.com/forever/tooltip/spell/11601 ("[109.07 + ((0.25 * 1) * Attack Power)] to [132.93 + ((0.25 * 1) * Attack Power)] damage").
- Client 1.60.1.70790: `talents/warrior.json` (every talent's rank text, incl. Improved Revenge 20/40/60%), `spelltips/warrior.json` (402927, 12328, 1329555, 1310222, 2565, 871, 12294, 1680, 18499, 2458, 12292, 1719, 12975, 1161, 2687, 100, 6178, 11578, 1310185, 23881, 5308, 20662, 1715, 6343, 11581, 1160, 11556, 7386, 72, 6572, 25288, 23922, 23925, 1240193, 11605, 12809), `trainables/warrior.json`, `races.json`, `addon-builds.json`, `addon-data.json` (`rotations/warrior-*`), `bis/warrior-protection.json` (raid preset buffs).
- Engine d5c708834: `sim/warrior/revenge.go` (0.25 AP, rank cap), `shield_slam.go:84`, `talents.go:286`, `spearing_strike.go:60`, `constants_auto_gen.go` (Revenge and Shield Slam rolls), `tank_warrior/TestP1TankWarrior.results` (CharacterStats: AP 1712.62, block value 77), `proto/common.proto` (stat indices).
- Curated: `data/curated/apl/warrior-{arms,fury,protection}.json`.
- Reviews: `design/reviews/guide-accuracy/round21/warrior.md`, `round22/warlock.md` (the leveling shape), `ledger.md` (Rounds 19 to 21).

## Note sweep

One line per row note: spec, row, bands it renders, each named talent / spell / rank / value with the level the player has it (L = leveling order, R = raid order), verdict.

Arms (20 bands, 10 to 60)

- Berserker Stance (30 to 60): "From 50" (Recklessness 50; card asks `spellIsKnown` 1719); Improved Tactical Mastery 3 (L 21, R 18) 25 rage; Mortal Strike (40) and Execute (24) work there, Overpower and Rend do not. OK.
- Bloodrage (10 to 60): 10 + 10 over 10 s, 60 s (2687); a quarter per point (12301); L 54 / 55, R 60. OK.
- Battle Shout (10 to 60, ranks 6673 to 11551): 10 rage every rank, 20 yards, raid. OK.
- Rend (10 to 60): bleed ignores armor; Overpower and Bloodthrill named only "once Bloodthrill is in" (L 33, R 31); main-hand hits on your Rend (1289682). OK.
- Overpower (12 to 60): Mortal Strike "from 40"; dodge opens it; Bloodthrill "once the talent is in"; Improved Overpower 2 (L 23, R 20) "most of the time". OK.
- Mortal Strike (40 to 60, ranks 12294 to 21553): 6 s every rank; healing cut 50%, 10 s. OK (round 21 fix holds).
- Reckless Swings (10 to 60, component row; talent R 57): raid build only, after Mortal Strike (40); 8 s; 5 rage; quarter of base health. OK (Should on the four-attack cap).
- Execute (24 to 60): below 20%; first spends the bar; Mortal Strike "from 40". OK.
- Slam (20 to 60): Improved Slam L 42, R 39; Mortal Strikes "from 40"; swing timer kept. OK.
- Spearing Strike (10 to 60, component row; talent R 25): raid build only, Battle Stance (engine ShapeshiftMask), 75%, twice on dragons and giants. OK.
- Hamstring (10 to 60, ranks 1715 to 7373): Windfury Totem gate; 10 rage every rank; never from 25% down (Should); rage other attacks need. OK.
- Heroic Strike (10 to 60): at 90 or more; Overpower 12, Execute 24, Mortal Strike 40. OK.

Fury (19 bands; leveling and raid orders identical)

- Bloodrage (10 to 60): 10 + 10, 60 s; Improved Bloodrage rank 1 at 59, rank 2 at 60. OK.
- Death Wish (10 to 60; talent 30, rows 10 to 28 component): pull if a second fits, execute phase, last 30 s. OK.
- Recklessness (50 to 60): execute phase or last 15 s; once a fight (1800000 ms). OK.
- Battle Shout (10 to 60): as Arms. OK.
- Execute (24 to 60): first below 20%; every point of rage; Bloodthirst "from 40". OK.
- Bloodthirst (40 to 60): "the button you lose the most by delaying". OK.
- Whirlwind (36 to 60): four targets, both weapons (1680; Dual Wield 20); Bloodthirst "from 40". OK.
- Reckless Swings (10 to 60; talent 31): behind Whirlwind "once you have it" (36); four white swings, 5 rage, quarter of base health. OK.
- Berserker Rage (36 to 60): "From 36"; Improved Berserker Rage 2 at 35, 10 rage; 30 s (18499); Magmadar, Onyxia, Nefarian. OK.
- Overpower (12 to 32, levels 12 to 35): Battle Stance before 36; Whirlwind 36. OK.
- Heroic Strike (10 to 60): above 60; 25% stop once Execute (24) is known. OK.

Protection (24 bands; leveling and raid orders identical)

- Defensive Stance (10 to 60): threat bonus; Revenge 14, Shield Block 16, Shield Wall 28 "all need it". OK (Should carried).
- Last Stand (10 to 60; talent 28, rows 10 to 26 component): below 35%; 30% for 20 s, lost after (12975). OK.
- Shield Wall (28 to 60): below 40%; 60% for 12 s; 15 min (871; no Improved Shield Wall in the build). OK.
- Bloodrage (10 to 60): before the pull, under 60; 10 + 10; Improved Bloodrage rank 1 at 10, rank 2 from 11. OK.
- Shield Block (16 to 60): 75% for 7 s, off the global; every block lights Revenge (14). OK (Should carried).
- Revenge (14 to 60, ranks 6572 to 11601): dodge, parry or block; most threat per rage (5, 2 with Focused Rage 38, 39, 41). OK.
- Shield Slam (40 to 60, ranks 23922 to 23925): "your biggest hit" fails at 60 in raid buffs against Revenge (0.25 x AP, Improved Revenge 60%). WARRIOR-R22-1. "grows with your block value" OK.
- Thunder Clap (10 to 60, 6343 to 11581): 20% attack speed slow every rank; refresh under 3 s. OK.
- Demoralizing Shout (14 to 60, 1160 to 11556): attack power stripped at every rank; does not stack with Demoralizing Roar; refresh under 3 s. OK.
- Battle Shout (10 to 60): only when nobody else's is up; same aura at 60. OK.
- Sunder Armor (10 to 60, 7386 to 11597): five stacks, then threat. OK.
- Heroic Strike (10 to 60): above 40, queued, no global. OK.

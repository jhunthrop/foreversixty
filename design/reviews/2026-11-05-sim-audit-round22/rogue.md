# Sim-accuracy round 22: rogue

Lane: site `wt/r22-rogue` (from int/r22 f9088dba6), fork `lane/r22-rogue` (from int/r22 170a59c39). Three confirmed bugs, all `old`: rogues could not use one-hand axes anywhere in the data, the Assassination raid build left 2.6% on the table, and Subtlety pressed Hemorrhage at five combo points. The level-60 book ranks (Backstab 9, Eviscerate 9, Feint 5, Deadly Poison V) are aligned with the client and stay with the owner's pending trainer check.

Measurements: spec-breakdown, 10000 iterations, seeds 1 and 50001, frozen BiS (`$R/bis-frozen`), guide builds, both factions. "Before" is this branch's own baseline (`$R/rogue/sb-base`, built on int/r22 with its 3% crit change), not `$R/before.tsv`. Raw outputs: `$R/rogue/ba.tsv`; rotation variants `$R/rogue/*.json`; tools `$R/rogue/tools/` (`ab.sh`, `bands.sh`, `dec.py`, `enc.py`).

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| K1 | **Rogues train one-hand axes in Forever; the pipeline, the BiS ranker's safety net and the addon all refused them.** `proficiency.py` listed the rogue bit on Axes as "vestigial" (beside the druid's polearm), so no axe reached `items/rogue.json`, the ranker or the addon's upgrade check. | One-Handed Axes (spell 196, skill line 44) has ClassMask **79** in the 70009, 70755 and 70790 clients and **71** in Era 1.15.9.70576: the rogue bit (8) is the only difference, so Forever added it on purpose. Two-Handed Axes (197) is 71 in both. SkillRaceClassInfo 629/630 (skill 44, ClassMask 9 = warrior + rogue). Hack and Slash (13960) reads "Axe/Sword" as one group. The druid polearm bit (spell 200, 1031) is in Era's client too, so it stays vestigial. Main's guild lane reached the same answer for its equip table (647bf7c5a). Fixed: `data/pipeline/proficiency.py` (fallback table, client derivation, docstring), `sim/cmd/leveling-bis/weapon_requirements.go` (static table), `addon/ForeverSixty/Rules.lua` (ROGUE row), `.claude/agents/wow-player.md` fact. Tests first, each seen failing: `test_rogue_has_one_hand_axe_but_no_two_hand_axe`, the vestigial-bit test, `TestClassWeaponSubclassesMatchesForeverDocumentedFacts`, `rules_spec.lua`. The engine already gave Hack and Slash's extra attack to axes (`talents.go` isAxeOrSword); `TestHackAndSlashAxeRegistersExtraAttackAura` now pins it with Bone Slicing Hatchet in the main hand. | CONFIRMED | old (ac42eb40b, 2026-09-30; round 8 re-checked paladin axes and kept the rogue exclusion) | S 4efed6f52, E 635416bf9 |
| K2 | Trainer ranks at 60 the cards do not cast: Backstab 9 (25300), Eviscerate 9 (31016), Feint 5 (25302), Deadly Poison V (25347/25351) | Same shape as the warrior lane's Revenge 6 / Battle Shout 7 / Heroic Strike 9. Their SkillLineAbility rows (13238, 14489, 13239, 13240) are byte-identical in Era 1.15.9.70576 and Forever 70790, where they are Ahn'Qiraj / Naxxramas book ranks. The books are in the 70790 ItemSparse and live on Wowhead Forever: 21300 Handbook of Backstab IX, 24102 Manual of Eviscerate IX, 21303 Handbook of Feint V, 21302 Handbook of Deadly Poison V ("Requires Level 60. Use: Teaches ..."). `spellranks.json` marks all four `book: true`; `trainables/rogue.json` lists 25300, 31016 and 25302 with `source: skill_line_ability` (same platform note as the warrior's). The engine caps all four under `core.IncludeAQ` (`backstab.go`:17, `eviscerate.go`:16-27, `poisons.go`:159); the three cards cast Backstab 11281, Eviscerate 11300 and Deadly Poison IV. Engine, cards and client agree. | Aligned; trainer question **pending owner** ("level-60 trainer book ranks") | - | - |
| K3 | **Assassination raid build: the MythicSim talent lead (+2.6%) is real and passes the adoption rule.** | Theirs `00530311251051031-30230302020122` against the guide's `00530311251551031-302302-005`: drops Improved Poisons 5 and Opportunity 5 for Precision +1, Composure 2, Endurance 2, Flawless Execution 1, Weapon Expertise 2 (needs Precision 3), Aggression 2. Raid 60, guide build, frozen BiS: Alliance 757.8 / 757.8 -> 777.5 / 777.1 (+2.6%), Horde 760.2 / 760.9 -> 775.4 / 775.0 (+2.0%). Every point it adds is modeled (Composure: `poised.go`; Endurance is the row toll, as Lightning Reflexes 2 already was); it drops nothing the rotation casts; it keeps 31 points in Assassination for Mutilate. Alternatives tried at 3000: Improved Poisons for Vile Poisons -0.2%, for Relentless Strikes -9%, Aggression 3 / Weapon Expertise 1 -0.5%. With the cheaper Eviscerate, Slice and Dice uptime fell from 93% to 76%; a three-point Slice and Dice once it has dropped (gate tests at 1, 2, 3 points; refresh windows 6 s and 10 s lose) adds +0.4% at both seeds (777.5 / 777.1 -> 780.6 / 780.6) and ties or gains at bare 20 to 60 (20 tie; 30 +0.4 to +0.5%; 40 +0.3 to +0.5%; 50 +0.3%; 60 +0.3 to +0.6%). Guide build line, talent prose, key-talent rows, priority prose and the Slice and Dice note rewritten; `TestAssassinationSliceAndDiceAtThreeOnceDown`. | CONFIRMED (adoption rule; rotation follows) | old (3d4712981, the 70755 activation build) | S a0a2818ed, 9f0adec67; E 31d269ec3, d7437239a |
| K4 | Subtlety MythicSim talent lead (+2.2%) | Theirs `00530301-320003-005203131321021151` on this branch: 791.2 / 791.8 -> 792.7 / 792.7 (+0.2%). Their gain was Improved Ambush 3 behind Ambush casts; our Subtlety holds a sword (Teebu's Blazing Longsword), so Ambush never casts. Dagger plus Ambush after every Ghostly Strike: 787.1 (-0.5% against the sword); Ambush opener 780.0. The rest is Improved Slice and Dice 3 for Lethality 1 and swaps between talents that do nothing on a boss (Master of Deception, Camouflage, Heightened Senses, Dirty Tricks). Below the 1% bar | Not adopted (below the bar) | - | - |
| K5 | Combat MythicSim rotation lead (Slice and Dice first at any combo point when under 1 s, Eviscerate at 4: +0.5%) | Eviscerate at four is the owner's pending "Combat Eviscerate at five": not reopened. The Slice and Dice half alone: raid 60 +0.2% Alliance (780.2 / 780.7 -> 782.0 / 782.0), +0.3% Horde, but -1.5 to -3% at every bare band 20 to 60 (60: 289.7 -> 283.6) and -0.6% on the raid build without raid buffs. No APL gate separates the builds: none of the raid-only Combat talents (Lethality, Opportunity, Composure) registers an aura `auraIsKnown` can read (gated variant = baseline exactly) | Not adopted (loses on every leveling band, no clean gate) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| X1 | Every rogue line in Blizzard's Forever notes | Thread 2360696 re-fetched 5 Nov 08:30 UTC: still 10 posts, last 4 Nov 19:35 UTC. Re-read the rogue lines (Expose Armor, Setup, poisons vs attackers, Sap, spec-swap energy, Cold Blood on Mutilate, Poised 18, Shank, poison AP 0.009/0.01, Blood Splatter, Relentless Strikes, Venom, Seal Fate, Mutilate, Coup de Grace, Efficient Strikes, Composure, Flawless Execution, Aggression, Flash of Steel, the dagger Hemorrhage/Ghostly Strike fix, Setup, Opportunity, Camouflage, Ghostly Strike's opener window, Hemorrhage's bleed share, Boot Knife, Death Mark) and Eureka's 10% energy discount: each matches the engine (as round 21 X1) | Aligned | - | - |
| X2 | Backstab's flat bonus: the live Forever tooltip reads "150% weapon damage plus 140"; the engine deals 1.5 x (weapon + 140) = 1.5 x weapon + 210 | Client 11281: effect 0 is 121 (normalized weapon damage, +140), effect 1 is 31 (150% weapon percent); the server applies weapon effects in index order, so the flat is multiplied, which is why Era's own text was hard-coded "plus 210" (Era Spell.csv 1.15.9.70576). Forever re-templated the text to "plus $s1", but kept Ambush's hard-coded "plus 290" (2.5 x 116) and spells out Mutilate's as `${m1*m2/100}` (79): the same client keeps the multiplied flat elsewhere. Round 5 recorded "150% plus 140" as aligned without this check | Aligned with the client rows; display only. A 60 Backstab log with a known dagger would settle it (OPEN on a log) | - | - |
| X3 | Live logs | Last looked 08:38 UTC 5 Nov: api.foreversixty.gg `/v1/reports/recent` lists only the 20 Oct report; foreverlogs.gg `/api/reports/public` 403; Warcraft Logs v1 `/v1/zones` on www/classic/fresh/vanilla/sod (66/57/21/29/12) has no Forever zone | OPEN | - | - |

## Assassination

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| A1 | Raid build and Slice and Dice at three | K3 | CONFIRMED | old | S a0a2818ed, 9f0adec67; E 31d269ec3, d7437239a |
| A2 | Spell by spell (fresh eyes) | Mutilate strikes are effect 121 (+88) then 31 (90%): 0.9 x (normalized weapon + 88), the off hand halved, as `mutilate.go`:139-185 and the tooltip's `${m1*m2/100}`. Ruthlessness 20% a rank (`talents.go`:133; client "60% at 3"). Malice 5% on every table including poisons (aura 290). Vigor +5 a point. Improved Poisons +2% a rank to each poison's base chance (`poisons.go`:50). Opportunity 4% a rank on Backstab, Garrote, Ambush, Mutilate (`client_values.go`:82). Weapon Expertise 1% a point as expertise. Aggression 5% a rank on Sinister Strike and Eviscerate. Composure only touches the opener (Poised breaks on the first energy ability; `core/poised.go`). Book ranks K2 | Aligned | - | - |

## Combat

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| C1 | Hack and Slash on axes | K1: the engine already treats axes like swords for the extra attack; Orc has no axe racial in Forever (`racials.go`:220-240), so race does not change the axe/sword choice | Aligned (engine), data fixed under K1 | - | E 635416bf9 (test) |
| C2 | Slice and Dice rotation lead | K5 | Not adopted | - | - |
| C3 | Spell by spell | Backstab (X2), Sinister Strike (effect 121 only, normalized), Eviscerate at five (owner), Dual Wield Specialization on white off-hand swings, Shank and Mutilate's off-hand strike (`talents.go`:106, `shank.go`:123, `mutilate.go`:185), Blade Flurry, Adrenaline Rush, Flash of Steel: as round 21 C1 | Aligned | - | - |

## Subtlety

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| S1 | **Hemorrhage was pressed at five combo points** (guide r24 sim request). With Rupture ticking and Death Mark cooling, a fifth-point Hemorrhage spent 35 energy on a point with nowhere to go | Gated at four or fewer: raid 60 guide build Alliance 791.2 / 791.8 -> 795.8 / 796.4 (+0.6%), Horde 784.6 / 784.6 -> 789.0 / 788.9 (+0.6%); bare 30, 40, 50 identical; bare 60 374.2 / 374.0 -> 374.5 / 374.6. Note, prose and `TestSubtletyHemorrhageNeverAtFiveComboPoints` (fails on the old card) | CONFIRMED | old (f89d5bc98, round 16), guides-handoff | S dd8f2408f, 9f0adec67; E 707736aa6, d7437239a |
| S2 | Talent lead | K4 | Not adopted | - | - |
| S3 | "Most of your damage comes from Rupture, Garrote, Death Mark, Shank and your white swings" (guide r23 optional request) | Raid 60 breakdown on the sword: Death Mark 19%, white swings 35%, Rupture 11%, Garrote 9%, Shank 2% (76% together); Hemorrhage 12%, Ghostly Strike 4.5% | Aligned | - | - |

## Before/after

Raid 60, guide builds, frozen BiS, seeds 1 / 50001, against this branch's own baseline:

| spec | faction | before | after | change |
|---|---|---|---|---|
| rogue-assassination | alliance | 757.8 / 757.8 | 780.6 / 780.6 | +3.0 / +3.0% |
| rogue-assassination | horde | 760.2 / 760.9 | 778.8 / 778.4 | +2.4 / +2.3% |
| rogue-combat | both | 780.2 / 780.7, 781.6 / 781.4 | unchanged | 0 |
| rogue-subtlety | alliance | 791.2 / 791.8 | 795.8 / 796.4 | +0.6 / +0.6% |
| rogue-subtlety | horde | 784.6 / 784.6 | 789.0 / 788.9 | +0.6 / +0.5% |

Bare bands: Assassination 30 to 60 +0.3 to +0.6% (the three-point Slice and Dice; 20 tie); Subtlety 30 to 50 identical, 60 +0.1%. No spec outside the rogue moves: every engine commit is a rogue APL copy or a rogue test.

## Files to regenerate

- `data/builds/1.60.1.70790/items/rogue.json` (pipeline normalize: one-hand axes enter the rogue list) and everything that reads it (simitems for rogue, the rogue BiS).
- `addon-data.json`, `addon-builds.json`, `Data.lua`: Assassination's raid order (new raid build), the Assassination Slice and Dice row and note, the Subtlety Hemorrhage row and note. The page's rotation card shows the old notes until then (the reviewer flagged it).
- Ladder goldens: `sim/request/testdata/ladder/rogue-assassination.golden.md`, `rogue-subtlety.golden.md`.
- Rogue stat weights / dps-per-point and coach starters if they read the raid build or the rotations.

## Measure at refresh

- **Rogue BiS picks move at the next re-rank** (no ranker run here). One-hand axes the ranker will now see, by band (the strongest by weapon DPS and rogue stats): 20 Butcher's Cleaver (1292), Shoni's Disarming Tool (9608), Gutterblade (17046); 30 Pronged Reaver (6692), Vibroblade (9485), Bearded Boneaxe (2878); 40 Axe of the Ebon Drake (10744), Digmaster 5000 (9465), Force of the Hippogryph (9684); 50 Dawn's Edge (12774), Grizzle's Skinner (11702), Flurry Axe (871), Axe of Rin'ji (13014); 60 Bone Slicing Hatchet (18737), Well Balanced Axe (18347), Zulian Hacker (19921), Crul'shorukh (19363) and Deathbringer (17068) when their content opens, the Grand Marshal's / High Warlord's Hacker (235481 / 235476) at their PvP rank. Combat (Hack and Slash, Axe/Sword) and Subtlety (weapon-agnostic) can take them; Assassination needs daggers.
- The Assassination raid build's BiS (`bis/rogue-assassination.json` still carries the old talents string) and K3's +2.6% on the re-ranked gear.
- Subtlety's Hemorrhage gate and round 21's 50-energy Shank threshold on the re-ranked BiS.
- Carried: rogue races at the re-rank; Subtlety band 20 weapon (round 20 K4).

## For other lanes

- **Platform / data:** `trainables/rogue.json` lists the book ranks 25300, 31016 and 25302 beside trainer ranks (same as the warrior's); `book-ranks.json` and `spellranks.json` already mark them.
- **Platform (addon):** `Rules.WEAPONS.DRUID` still carries `polearm = true`, which `proficiency.py`, the ranker table and the guild table all refuse (Era's client carries the same vestigial druid bit).
- **Warrior:** `web/src/content/guides/_sections.test.ts` fails on int/r22 for warrior/protection ("strengths: expected 5 to be less than or equal to 4"); not a rogue file, left alone.
- `.claude/agents/wow-player.md` now says rogues train one-handed axes; every lane's reviewer reads it.

## For the guide lane

- Done here: assassination.md build line, Strengths bullet, Talents and builds (31/20/0), key-talent rows (Aggression, Flawless Execution, Weapon Expertise in place of Improved Poisons), the Combat-points paragraph, the priority line (three-point Slice and Dice, on main's "any damage racial" wording); subtlety.md priority line (Hemorrhage never at five, on main's wording). `_prose.test.ts` and `builds.test.ts` pass.
- combat.md:56 "swords and axes" for Hack and Slash is now true for the BiS too; the Combat leveling paragraph's "sword, axe or mace" (round 4 lead) is now backed by the data.

## Guide text not done

- None.

## Rendered review

wow-player on the rendered Assassination and Subtlety pages (astro dev on this worktree after `npm run sync`, 1440 px, `$R/rogue/shots`), plus the curated note text: **SHIP with fixes**, all taken except the two below. Taken: Slice and Dice note "a three-point cast ... getting the attack speed back" (not "20%": rank 1 is 20%, rank 2 30%); prose "a three-point Slice and Dice"; Weapon Expertise "dodges or parries your attacks 2% less"; Subtlety note "At five the extra point is wasted, so save the energy for the finisher"; prose "but never at five combo points, where its point is wasted". Checked and left: the generated raid-build summary ("Opportunity (1)" is right: the level-60 build takes one point); the rotation card's old note (regeneration, listed above). Not taken: a Precision row in Key talents (the old page did not list it either).

## Still open

- Death Mark storage and the Mutilate / Shank off-hand penalty (40+ log); the doubled poison coefficients on a live log; Backstab's flat bonus (X2) on a 60 log. Combat Eviscerate at five and the level-60 book ranks stay with the owner.

## Sources

- wago SkillLineAbility 1.60.1.70790 / 70755 / 70009 and 1.15.9.70576 (Era); SkillRaceClassInfo 70755; Spell.csv 70790 and Era 1.15.9.70576; SpellEffect 70790; ItemSparse 70790 (cached under `scratchpad/b70790`, `scratchpad/raw70755`, `$R/warrior/era`, `$R/rogue/era`).
- Wowhead Forever tooltips (5 Nov, about 08:15 UTC): https://nether.wowhead.com/forever/tooltip/item/21300, /item/21302, /item/21303, /item/24102, /spell/11281.
- Blizzard, WoW Forever Beta Development Notes: https://us.forums.blizzard.com/en/wow/t/2360696 (10 posts, read 5 Nov 08:30 UTC).
- MythicSim comparison: `scratchpad/competitor/lanes/physical.md` sections 3 to 5 (another agent's data; every figure above re-measured on this branch).
- Guide-accuracy rounds 22 to 24 rogue reports (origin/main); round 21 rogue report; main's guild fix 647bf7c5a.

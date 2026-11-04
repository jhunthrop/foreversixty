# Sim audit round 14: regression lane, melee, hybrid, tank and hunter side (regmelee)

Scope: every change rounds 12 and 13 and the land2 landing made to Arms, Fury, Protection warrior, the three rogues, the three hunters, cat, bear, Retribution, Protection paladin and Enhancement. For each: what it claimed, what the code and data do now, and a quick test or sim that shows it.

- **Branches:** fork `lane/r14-regmelee` from int/r14 (5195a8889), site `wt/r14-regmelee` from int/r14 (72471fcb3).
- **Client rows:** the layered 70753 tables with the 2026-11-03 14:14 UTC push (`scratchpad/r13/hotfix/root/1.60.1.70753/raw/`, the copy the landing layered), read with `regmelee/tools/q.py`. A tooltip and its client row count as one source.
- **Measurements:** spec-breakdown, seeds 1 and 50001, 10,000 iterations for every reported figure; 1,000 to 3,000 iterations for the checks marked "explore". Before = `$R/sb-before` from `$R/site-before/sim`, or `$R/before.tsv` at raid 60.

**Result: 2 confirmed bugs, both fixed.** One is a sim bug: Beast Mastery's Arcane Shot downrank rows lose after the hotfix push (new-data, missed by round 13's check). The other is guide copy: round 13's Survival line says Aspect of the Falcon has 10 more melee attack power than Beast (R13). Everything else re-checks as aligned. There are also three stale code comments; they are fixed but not counted.

## The landing (0bff3751b..72471fcb3, fork 919616e88..5195a8889)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | **85204127f hotfix layer: the pushed hunter values are exactly what spellconst and the engine use.** I regenerated each class's constants into a scratch copy of the package from the committed `data/builds/1.60.1.70753/spellconst/<class>.json` (`go run ./sim/core/spellconst/gen`). For hunter, warrior, rogue, druid, paladin and shaman the output is byte-identical to the fork's `constants_auto_gen.go`, apart from the Source header line. The spellconst diff in 85204127f matches the push row by row: Aimed Shot 40/70/115/185/260/350, Arcane Shot 25/35/50/85/120/175/220/280, Serpent Sting 3/7/14/26/40/58/75/100 (rank 9 135, AQ), Sniper Shot 320/450/590, Hawk 20/35/50/70/90/110, Trueshot Aura 20/35/50/75/100 (e.g. SpellEffect 697764, 691069, 690655, 692138, 697683 in `r13/hotfix/pushdiff.txt`). The hand tables pass `TestHotfixPush70753HunterAmounts`, `TestHotfixPush70753TrueshotAura` and `TestAspectOfTheHawkAttackPowerFollowsThe70753Push`. The committed push record has one more row than the hotfix lane's copy: one SpellEquippedItems add. No sim reader uses it. In the sim, MM raid 60 shows Aspect of the Hawk 14322 and Trueshot Aura 20906 (100 ranged attack power), Aimed Shot 20904. Only hunter player spells changed (hotfix lane X1). | scratch `regmelee/sc/`; fork tests; MM breakdown (explore) | Aligned | - | - |
| L2 | **fc90c1fcc apl-sync.** `make apl-check ENGINE_DIR=<lane engine>`: for all 28 specs, the fork copy and the site copy run the same rotation as `data/curated/apl` (parsed JSON, notes stripped). No mismatch line. | apl-check output | Aligned | - | - |
| L3 | **f9b264a4f addon-builds and ladder goldens.** `TestCommittedAddonBuildsFollowTheGuides` passes. addon-builds and the ladder share one walk, `leveling.LadderPointOrder` with `ForSpec` reservations, so the committed band builds are the guide and walk builds. For my 14 specs, `TestRotationLadder` matches every golden at the pin. The goldens show the round-13 walk rows: Arms 40 casts Mortal Strike 12294, Survival 30 Strider Kick 1317257, Assassination 30 Mutilate 1310707, Subtlety 20 Ghostly Strike 14278 and 30 Hemorrhage 16511, bear 30 Lacerate 1328617 and 40 Primal Bite 407995, Fury 38+ Raging Blows (`05153105122010...`). Arms 30 to 40 report "zero_casts 2458": Berserker Stance is gated on Recklessness, which is learned at 50 (round 12 ef4b989bb). That is intended, and it is informational only (non-strict). | `go test ./cmd/addon-builds ./request -run TestRotationLadder/...` | Aligned | - | - |
| L4 | **32fb15c5f addon-data, coach, dps-per-point.** `tools/coach/test_persona_starter.py` passes (53). The starters say what the rotations do: Arms Charge then shout, MM Lone Wolf and Serpent Sting, BM Aimed Shot, Kidney Shot at 25 Energy. The bear coach rule lists Primal Bite 1238073 between ranks 2 and 3. That id has no SkillLineAbility row (no trainer teaches it), so a player never casts it; harmless. addon-data's MM rotation carries the Aimed Shot and Serpent Sting lines. Copy note: its first MM line ("In a raid, dismiss your pet before the pull") shows at every level from 10. It is correct but early; sent to guides. | coach test; addon-data.json `rotations` | Aligned | - | - |
| L5 | **3e3a14ff9 (the dps-per-point tool change only).** `formIgnoresWeapon` zeroes the cat's `weapon_dps` at 60 and at every band (`TestTheCatHasNoWeaponTerm`). That is right for the game: Cat Form ignores weapon damage. The tool still writes `weapon_dps_error` (0.0004 to 0.0007) beside a forced 0. No reader uses that field (grep of sim, pipeline, web, addon, tools), so this is cosmetic. Whether the engine lets weapon DPS reach Cat Form is the hybrids lane's question, not this row's. | `sim/cmd/dps-per-point/main.go:162-178`; dps-per-point.json | Aligned | - | - |
| L6 | **The review lane's Arms opener (5dc93c822 / S review).** Charge 11578 at 0 s, then Battle Shout at 0 s (skipped under the raid's shout), with Bloodrage in the priority list. Raid 60: Charge 1.00 a fight, no shout cast. Bare 20: Charge 100 for 9 rage plus Improved Charge's 6, then Battle Shout 5242 (1.5 casts in 180 s at a 2 minute duration). Raid 60 at the landing: horde 731.2 → 729.7, alliance 719.0 → 717.5 (before.tsv), as the review lane measured (-0.2%). | breakdowns (explore) | Aligned | - | - |
| L7 | **The landing moved nothing else at raid 60 in this group.** before.tsv equals round 13's after-table for all 14 specs, except Arms (L6). Fury's Raging Blows and MM's Lone Wolf build reach the published numbers at the refresh; through `-talents`, the guide MM raid build reads 737.9 with no pet (round 13: 737.7). | before.tsv vs round-13 integration.md | Aligned | - | - |
| L8 | f9a1644b4 (MM summary and the Surefooted sentence), 2b0ec9937 (e2e), the merges into int/land2: none touches a rotation, build or engine file in this group. The merge-side diffs into my files are only main's coach MM wording (72cecbabb). | `git diff <merge>^2 <merge>` | Aligned | - | - |

## Round 13 changes

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| T1 | **Serrated Blades and Hack and Slash's mace as a share of armor (E b08a1c83c).** Both write `PseudoStats.ArmorIgnorePercent` (9% at 3/3, 15% at rank 5), applied in `GetArmorDamageModifier` as armor × (1 - pct) before flat penetration (`core/spell_resistances.go:84-88`). The warrior's Weaponmaster mace clause shares the path. The only other flat ArmorPenetration writer is an item stack effect. Subtlety raid 60 in before.tsv: 544.0 / 544.7, as round 13. | code; `TestHackAndSlashMaceIgnoresArmorPercent` passes | Aligned | - | - |
| T2 | **Demoralizing Shout threat 43 + 0.8 a level (E 1f440ade7).** Client: effect 63 on 1160/6190/11554/11555/11556 at 11/19/27/35/43, RealPointsPerLevel 0.8 (SpellEffect 1361983-1361987). Engine: `demoralizingShoutThreat` capped at the rank's max level; FlatThreatBonus is not multiplied by the 0.4 damage multiplier (`core/spell_result.go:96`). The comment above it still said the client has no threat effect; fixed (comment only). | test passes; client rows | Aligned (comment fixed) | - | E f62646f06 |
| T3 | **The hotfix push in the engine (E a3b2af6d4) and the Aimed Shot / Serpent Sting lines (E bcb7c2912, S 35534af90).** See L1. Aimed Shot: 2 s cast (CastingTimeIndex 5), 6 s category cooldown shared with Multi-Shot (SpellCooldowns 55521 / 53416), 310 mana at rank 6. The engine matches (`aimed_shot.go`). The hunter Hawk comment still gave the 70707 ranks; fixed (comment only). | client rows; code | Aligned (comment fixed) | - | E f62646f06 |
| T4 | **The hotfix lane's H8: "the hunter rank rows still win".** It removed every hunter row at once. Measured row by row, the Aimed Shot rows win but the Arcane Shot rows lose: see **H1**. | below | see H1 | - | - |
| T5 | **MM raid build with Lone Wolf (S daaf8314d).** The guide's raidBuild through `-talents`: no pet rows, 737.9 at 3,000 iterations, seed 1, horde. addon-builds carries the raid order, legal point by point (row 6 opens after 25 tree points). | breakdown (explore); addon-builds.json | Aligned | - | - |
| T6 | **Fury Raging Blows (S cb1325223).** Guide build `30305013002/05153105122010501/2`. The walk and the ladder take Raging Blows from 38. | ladder golden | Aligned | - | - |
| T7 | **Bear bands 30/40 (hybrids B1) and Templar's Bulwark level 0 (E 0c7ed03ba).** Ladder: Lacerate at 30, Primal Bite 407995 at 40. `TestTemplarsBulwarkMatchesTheClient` passes. The ladder's "learned but unused Primal Bite 1238073" at 60 comes from a spellranks.json row with no trainer row. The rewrite only steps down from the authored 1238070, so it is never cast. The spellranks.go override comment still said the engine registers 1238073; fixed (comment only). | ladder; `sim/internal/spellranks/spellranks.go:411-422`; SkillLineAbility (no 1238073 row) | Aligned (comment fixed) | - | S 9adfa903e |
| T8 | **The walk's rows in this group (S 340048375):** Arms 40, Fury 30, Survival 30, Assassination 30, Subtlety 20/30. Each band casts its reserved talent (L3). | ladder goldens | Aligned | - | - |
| T9 | **Protection paladin Holy Strike rank 1 at 40 (S 34583ae09).** The ladder at 40 casts 679 (rank 1) as its top action. The row is unchanged since the ranks lane measured it, and no paladin data moved since. | ladder; rank-choices.json | Aligned | - | - |
| T10 | **Searing Totem cadence 2.437 s (E 5ca93cfa3) reaches Enhancement.** `fire_totems.go` attackInterval is shared by both shaman specs. The Enhancement golden was regenerated for it. | code | Aligned | - | - |
| T11 | **Mana regen phase (E e88a33c97) and the trainables fixes.** The regen phase moves the mana users of this group within ±0.3% (round 13). The trainables changes for my classes: Battle Stance and Auto Shot added; talent tags; Aimed Shot 2 s; Thunder Clap 6 s; triggered spells dropped. The sim does not read trainables (only the web rotation lists do). Where trainables now disagree with spellconst (Thunder Clap 4000 ms in spellconst), the engine already uses the client's 6 s category cooldown (`thunder_clap.go:31`; SpellCooldowns 54504). | diff of trainables/*.json at 85204127f; code | Aligned | - | - |
| T12 | **Guide copy round 13 wrote (13497f823): Survival's aspect line.** See **S1**. | below | see S1 | - | - |

## Round 12 changes (re-checked at the landed pin)

Round 13 re-checked these (regmelee G1-G16). I re-ran each change's own fork test on 5195a8889 and checked its effect in a breakdown or golden. All still hold:
- Poised for rogues and the cat;
- Improved Charge's rage with a Charge;
- the Thunder Clap raid debuff at 20%;
- Windfury Totem's bonus ends with the totem;
- missile speeds and Hammer of Wrath's flight;
- Arcane Shot, Serpent Sting and Volley with no spell power (still pending the owner's tooltip check, ruling 5);
- Enrage at 16% in Dire Bear;
- Primal Bite rank 3 at 60;
- Elixir of Holy Power;
- the faction/race rank rows;
- Command;
- the Arms execute swap.

Tests: `TestImprovedChargeRageComesWithTheCharge`, `TestPoised*`, `TestCatPoised*`, `TestEnrageArmorPenaltyFollowsTheForm`, `TestHunterShotsTakeNoSpellPower`, `TestHunterShotsTravelAtTheClientsSpeed`, `TestHammerOfWrathFliesAtTheClientsSpeed`, `TestWindfuryTotemBonusEndsWithTheTotem`, `TestThunderClapRaidDebuffSlowsTwentyPercent`, `TestElixirOfHolyPowerMatchesTheClient`, `TestCommand*`, `TestManaTick*`: all pass.

## Hunters

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Beast Mastery's Arcane Shot downranks now lose.**<br>**What was wrong:** BM cast Arcane Shot rank 3 (14282) at 40-49 and rank 1 (3044) at 50-59 and at bare 60. Those rows were adopted in round 12, when the top rank hit for 217 and Aimed Shot was only the opener.<br>**What changed under them:** the 3 November push raised Arcane Shot to 280 at rank 8 (rank 5 94 → 120, rank 6 134 → 175). Round 13 made Aimed Shot share Multi-Shot's cooldown. Round 13 re-measured the hunter rows only all together; the Aimed Shot rows' gain (+2.8 to +4.4%) hid the Arcane Shot loss.<br>**Row by row now** (spec-breakdown built with only the Arcane rows removed, 10,000 iterations, error 0.1 DPS; seed 1 / 50001):<br>- horde: 40 167.4/167.5 → 167.9/168.1; 50 229.8/229.8 → 229.9/229.9; bare 60 327.2/327.1 → 329.5/329.5;<br>- alliance: 40 167.7/167.6 → 168.1/168.2; 50 227.9/227.7 → 228.1/227.9; bare 60 327.2/327.1 → 329.7/329.7.<br>So the rows lose beyond error at 40 and bare 60 and tie at 50. Under the rank-row rule, none stays.<br>**Fix:** the three Arcane Shot rows are removed, so BM casts its top learned rank. The Aimed Shot rows stay. | `TestRound14BeastMasteryArcaneShotCastsTheTopLearnedRank` (sim/request) failed first (14282 / 3044 cast), then passed. Round 12's two Arcane cases moved into it. `TestEmbeddedRankChoicesMatchCurated`, sim/rankchoices, api coach rank tests and test_addonrotation pass. | CONFIRMED-BUG | new-data (the 11-03 push; rows from round 12; round 13 H8 checked them only together) | S 717364c9c |
| H2 | Aimed Shot rank-1 rows, checked row by row: BM 40 162.8 → 167.4 (+2.8%), BM bare 60 315.4 → 327.3 (+3.7%), MM bare 60 316.2 → 330.2 (+4.4%) (3,000 iterations, both seeds, horde). They hold. The fights are mana-bound (61 s out of mana at BM 40). | `sb-noaimed` vs `sb` | Aligned | - | - |
| H3 | Aspect of the Falcon in the engine follows the pushed Hawk: 110 melee and 110 ranged at 60, the same melee attack power as Beast rank 4 (1299447, 110). No rotation casts Falcon. The engine lets Deadly Aspects proc under Falcon (by Falcon's text, 469145). It does not model Falcon's own 5% Quick Strikes (SpellAuraOptions 239200, ProcChance 5). Deadly Aspects' own text names only Hawk and Beast. Moot while nobody casts Falcon. | client rows; `aspect_of_the_falcon.go`, `talents.go:430-470` | OPEN (latent; no published number) | - | - |

## Survival guide

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **"Aspect of the Falcon gives 10 more melee attack power at 60" is wrong.**<br>**What the client says:** Falcon (469145) gives "the same amount as their highest rank of Aspect of the Hawk". At 60 that rank is 14322: 55 before the push, 110 after. Rank 7's 120 (25296) needs the Ahn'Qiraj book. Beast rank 4 (1299447) is 110.<br>**So:** Falcon gave 55 less before the push, and gives the same now. It never gave 10 more.<br>**Fix:** the sentence now says Falcon gives the same melee attack power at 60 because it copies your Aspect of the Hawk. The Deadly Aspects half of the sentence is unchanged (see H3). | SpellEffect 692138 (push), Spell 469145 text, `TestAspectOfTheFalconGrantsHawkAttackPowerToBothAttackKinds`; `_prose.test.ts` + `_sections.test.ts` 142/142, coach starters 53/53 | CONFIRMED-BUG (copy) | R13 13497f823 | S 9adfa903e |

## Warriors, rogues, cat, bear, Retribution, Protection paladin, Enhancement

No new finding beyond the rows above. What I re-checked with new eyes this round:
- Arms raid and bare 20 openers (L6);
- Demoralizing Shout's client rows: effect 63 threat, and aura 99 -196 base, -1.4 a level, 204.4 at 60 (T2);
- Thunder Clap's 6 s client cooldown (T11);
- Raging Blows on Fury's ladder (T6);
- armor ignore on the rogues and Weaponmaster (T1);
- the bear's and cat's Primal Bite and Poised rows (T7, round 12 list);
- Searing Totem's cadence for Enhancement (T10);
- Holy Strike rank 1 and the Consecration rows for the paladins, unchanged since round 13's measurement with no data move since (T9).

Every spell and talent these specs' APLs name was audited spell by spell in rounds 1 to 13 (grep of the reports: Hamstring 29 reports, Seal of Command 76, Rage of the Farseer 14, Sacred Arbiter 9, and so on). I did not redo those tables; nothing in the push or the landing touches them (hotfix X1: only hunter player spells changed).

## Before/after

Only H1 moves a number. Beast Mastery, bare (horde race of the published entry; alliance likewise), 10,000 iterations, seed 1 / 50001:

| band | horde % | alliance % | seed 1 horde / alliance |
|---|---|---|---|
| bare 40 | +0.30 / +0.36 | +0.24 / +0.36 | 167.4 → 167.9 / 167.7 → 168.1 |
| bare 50 | +0.04 / +0.04 | +0.09 / +0.09 | 229.8 → 229.9 / 227.9 → 228.1 |
| bare 60 | +0.70 / +0.73 | +0.76 / +0.79 | 327.2 → 329.5 / 327.2 → 329.7 |

Raid 60 and every other spec are unchanged: the rows are below 60 or bare-only, and the fork commit changes only comments. No spec outside this group moves.

## Files to regenerate

- `sim/request` TestRotationLadder golden for hunter-beast-mastery (bands 40 and 50 cast the top Arcane Shot).
- addon-data.json and Data.lua (BM's rank rows feed the addon's priority card).
- BM BiS and weights at bands 40, 50 and bare 60, at the refresh.

## Measure at refresh

- BM bands 40 to 60 picks and weights with the top Arcane Shot.
- Another Arcane Shot rank was not searched (rank 2, or rank 4 at 50). A future row needs a gain beyond error at both seeds.

## For other lanes

- **ranks / data (rank-choice rows):** H1 removed BM's three Arcane Shot rows. Please measure rows one at a time, not all removed together; a winning row can hide a losing one.
- **hybrids:** 3e3a14ff9 zeroes the cat's weapon term in the tool. Whether the engine lets weapon DPS reach Cat Form (the sweep read 0.8 to 2.3 DPS per weapon DPS) is yours. If it does, the cat's sims and BiS picks value weapon DPS.
- **guides:** Survival's aspect sentence is edited (S1). Two notes, no change made: the MM priority card's first line ("In a raid, dismiss your pet before the pull") shows from level 10 (L4). Survival's "Deadly Aspects does nothing under Falcon" disagrees with the engine's Falcon handling (H3); it is moot while no rotation casts Falcon.
- **physical / hunter owner:** H3 (Falcon's own 5% proc, Deadly Aspects under Falcon) is latent.

## Still open

- Arcane Shot's spell power (ruling 5, E 05b72a3b5): still pending the owner's tooltip check.
- H3: Falcon's proc model (latent).
- Carried from round 13: Deep Wounds' double multiplier, Raptor Strike's queue and white miss split, normalized rage's off-hand factor, Lacerating Strikes' refresh, Holy Shield crit. All need level 40+ logs; WCL is blocked (gcloud login expired).

## Origin totals

2 confirmed: **new-data 1** (H1), **R13 1** (S1, 13497f823). R12 0, landing 0, old 0.

Heads: fork `lane/r14-regmelee` f62646f06 (comments only); site `wt/r14-regmelee` (this report on top of 717364c9c).

Sources (read 2026-11-03): the layered 70753 raw tables with the 11-03 push (SpellEffect, SpellMisc, SpellCooldowns, SpellPower, SpellLevels, SpellAuraOptions, SkillLineAbility, Spell); `scratchpad/r13/hotfix/pushdiff.txt`; data/builds/1.60.1.70753/{spellconst,trainables,talents}/; the published BiS files.

# Sim-accuracy round 11: ranks lane

Scope: the per-band rank-choice mechanism (`sim/request/rotation_ranks.go`, `rank_choices.go`, `sim/rankchoices`), every non-healer row of `data/curated/rank-choices.json` and the bare-60 healer rows (the sub-60 healer rows belong to the heal profile lane). Where it was measured: site `wt/r11-ranks` (from origin/main 8f047a62f), engine at the pin cfcb76e80 (no engine change in this lane). Horde, each band's published gear and preset, 10,000 iterations, seeds 1 and 50001, unless marked. Tools and raw output are in scratchpad `r11/ranks/` (`reconfirm.py`, `explore.py`, `sweep.py`, `coefcheck.py`, `learn.py`, `addoncheck2.py`, `coachcheck.py`, `out/`). The scratch Go files that let spec-breakdown read a substitute table (`RC_FILE`) and dump spells and ranks are in `r11/ranks/scratchgo/` and are not committed.

Error bar: a DPS row passes when its gain at each seed beats twice the combined standard error. Healer rows (effective HPS) and tank rows (threat per second) print no error. They pass at more than 0.2% with the two seeds agreeing, as in round 10. In the table, "0.00" in the error column means either a healer or tank row, or a DPS figure printed to one decimal. The one DPS row that matters, Shadow Bolt at 30, was re-run at full precision: its error is 0.07%.

## Owner ruling 3 (recorded)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| O1 | Retribution at bare 20 loses 0.3 to 0.5% under round 10's paced Consecration line. Owner ruling: accepted, within error, not reopened. No rank choice applies at 20: only Consecration rank 1 is learned there. This round re-measured the Retribution rows (R1 below) without changing the line. | round 10 ranks.md "Still open"; brief ruling 3 | ruled (no change) | - | - |

## Rank rows: re-confirmed at the pin

Each row against the same table with that row taken out (the default rule), then against the ranks one either side of it. Columns: gain at seeds 1 / 50001 (%), error (%), the neighbouring ranks' change against the row.

| spec | spell | band | rank | row vs default | err | neighbours vs row |
|---|---|---|---|---|---|---|
| druid-balance | Moonfire | 30 | 3 | +4.45 / +4.75 | 0.29 | r2 -2.48/-2.61, r4 +0.14/-0.14 |
| druid-balance | Wrath | 30 | 4 | +6.44 / +6.59 | 0.29 | r3 -10.59/-10.71, r5 -6.05/-6.18 |
| druid-balance | Insect Swarm | 30 | 1 | +2.25 / +2.39 | 0.28 | r2 -2.20/-2.33 |
| druid-restoration | Rejuvenation | 60 bare | 10 | -1.05 / -0.96 | 0.00 | r9 -0.17/-0.19, r11 +1.07/+0.97 |
| hunter-beast-mastery | Aimed Shot | 40 | 1 | +0.39 / +0.39 | 0.18 | r2 -0.26/-0.26 |
| hunter-beast-mastery | Aimed Shot | 60 bare | 1 | +0.38 / +0.34 | 0.10 | r2 -0.17/-0.17 |
| hunter-marksmanship | Aimed Shot | 60 bare | 1 | +0.32 / +0.25 | 0.10 | r2 -0.07/-0.04 |
| mage-arcane | Frostbolt | 30 | 4 | +1.34 / +1.46 | 0.34 | r3 -3.73/-3.84, r5 -1.32/-1.44 |
| mage-fire | Fire Blast | 40 | 4 | +0.72 / +0.79 | 0.29 | r3 -1.89/-2.21, r5 -0.72/-0.78 |
| mage-fire | Scorch | 30 | 1 | +1.00 / +1.15 | 0.41 | r2 -0.99/-1.14 |
| mage-fire | Scorch | 40 | 1 | +5.58 / +5.63 | 0.20 | r2 -0.98/-1.17 |
| mage-fire | Scorch | 50 | 1 | +6.27 / +6.38 | 0.32 | r2 -2.63/-2.68 |
| paladin-holy | Holy Light | 60 bare | 4 | +4.91 / +4.70 | 0.00 | r3 -2.31/-2.11, r5 -4.92/-4.69 |
| paladin-protection | Consecration | 40 | 1 | +1.32 / +1.11 | 0.00 | r2 -1.88/-1.83 |
| paladin-protection | Consecration | 50 | 1 | +1.96 / +2.10 | 0.00 | r2 -1.40/-1.50 |
| paladin-protection | Consecration | 60 bare | 1 | +3.07 / +3.05 | 0.00 | r2 -1.05/-1.10 |
| paladin-retribution | Consecration | 30 | 1 | +1.28 / +1.13 | 0.00 | r2 -1.26/-1.12 |
| paladin-retribution | Consecration | 40 | 1 | +2.22 / +2.21 | 0.00 | r2 -1.98/-2.07 |
| paladin-retribution | Consecration | 50 | 1 | +7.69 / +7.69 | 0.18 | r2 -5.26/-5.32 |
| paladin-retribution | Consecration | 60 bare | 4 | +1.61 / +1.73 | 0.12 | r3 -0.41/-0.45, r5 -1.58/-1.71 |
| priest-discipline | Heal | 60 bare | 4 | +2.61 / +2.50 | 0.00 | r3 -0.17/-0.07, r5 -2.54/-2.44 |
| priest-discipline | Binding Heal | 60 bare | 1 | +3.44 / +3.33 | 0.00 | r2 -0.71/-0.61 |
| priest-holy | Flash Heal | 60 bare | 7 | +8.80 / +8.77 | 0.00 | r6 -2.53/-2.49, r8 -8.09/-8.06 |
| priest-holy | Prayer of Healing | 60 bare | 1 | +0.29 / +0.29 | 0.00 | r2 -0.03/-0.06 |
| priest-holy | Lightwell | 60 bare | 1 | +1.29 / +1.33 | 0.00 | r2 +0.06/+0.13 |
| priest-shadow | Shadow Word: Pain | 30 | 3 | +3.55 / +3.37 | 0.00 | r2 -3.60/-3.60, r4 -3.43/-3.26 |
| priest-shadow | Mind Blast | 40 | 4 | +1.90 / +1.82 | 0.22 | r3 -2.02/-1.94, r5 -0.54/-0.54 |
| priest-shadow | Mind Blast | 50 | 2 | +4.14 / +4.13 | 0.15 | r1 -3.17/-3.17, r3 -2.06/-2.01 |
| priest-shadow | Devouring Plague | 30 | 1 | +1.39 / +1.39 | 0.00 | r2 -1.37/-1.37 |
| priest-shadow | Devouring Plague | 40 | 1 | +3.21 / +3.21 | 0.23 | r2 -1.24/-1.24 |
| priest-shadow | Devouring Plague | 50 | 1 | +4.63 / +4.74 | 0.15 | r2 -1.26/-1.26 |
| priest-shadow | Devouring Plague | 60 bare | 1 | +2.46 / +2.52 | 0.16 | r2 -0.45/-0.48 |
| priest-shadow | Shadow Word: Death | 50 | 2 | +1.22 / +1.27 | 0.14 | r1 -1.11/-1.16, r3 -1.21/-1.26 |
| priest-shadow | Shadow Word: Death | 60 bare | 2 | +1.46 / +1.49 | 0.16 | r1 -0.68/-0.62, r3 -0.45/-0.48 |
| shaman-elemental | Lightning Bolt | 30 | 4 | +8.53 / +8.51 | 0.00 | r3 -5.73/-5.93, r5 -7.86/-7.84 |
| shaman-elemental | Lightning Bolt | 40 | 4 | +17.69 / +17.82 | 0.35 | r3 -4.80/-4.69, r5 -9.38/-9.48 |
| shaman-elemental | Lightning Bolt | 50 | 4 | +36.39 / +36.45 | 0.25 | r3 -12.52/-12.57, r5 -6.85/-6.84 |
| shaman-elemental | Lightning Bolt | 60 bare | 4 | +31.81 / +31.93 | 0.20 | r3 -9.55/-9.66, r5 -6.03/-5.93 |
| shaman-elemental | Flame Shock | 40 | 1 | +5.16 / +5.15 | 0.32 | r2 -1.17/-1.17 |
| shaman-elemental | Flame Shock | 50 | 2 | +2.29 / +2.29 | 0.19 | r1 -2.57/-2.63, r3 -0.13/-0.20 |
| shaman-elemental | Lava Burst | 40 | 1 | +1.30 / +1.29 | 0.30 | r2 -1.28/-1.28 |
| shaman-elemental | Lava Burst | 50 | 1 | +2.64 / +2.56 | 0.19 | r2 +0.92/+0.99 |
| shaman-elemental | Lava Burst | 60 bare | 2 | +1.65 / +1.60 | 0.15 | r1 -1.00/-1.05, r3 -0.79/-0.84 |
| shaman-restoration | Healing Wave | 60 bare | 3 | +3.73 / +3.73 | 0.00 | r2 -2.23/-2.20, r4 -3.08/-3.06 |
| shaman-restoration | Riptide | 60 bare | 2 | +2.84 / +2.90 | 0.00 | r1 -0.24/-0.21, r3 -2.76/-2.81 |
| warlock-destruction | Shadow Bolt | 30 | 4 | +0.57 / +0.57 | 0.00 | r3 -5.93/-5.93, r5 -0.57/-0.57 |
Source: `r11/ranks/out/reconfirm.jsonl`.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Rows that hold.** Every non-healer row and every bare-60 healer row but one still beats the default rule beyond error at both seeds, and no neighbouring rank beats it, with two exceptions below. Moonfire rank 4 ties rank 3 at 30 (+0.14 / -0.14), so rank 3 stays. Flame Shock rank 3 at 50 is within error of rank 2 (-0.13 / -0.20). Lightwell rank 2 at bare 60 is within the healer bar of rank 1 (+0.06 / +0.13). | table above | Aligned | - | - |
| R2 | **Lava Burst at 50 to 59 cast the wrong rank.** Round 9's row chose rank 1 (408490). Rank 2 (1238299, learned at 40) beats it by +0.92% / +1.00%, with an error of about 0.15%, and beats the default rank 3 by +3.58% / +3.57%. The same comparison on round 10's baseline (bb619a188, engine 4d3129cae) gives the same numbers to the hundredth. So rank 2 was already better when round 9 chose rank 1. Round 10's re-confirmation only compared each row with no row, never the neighbouring rank. A Classic player would cast the rank-2 Lava Burst at 50: it costs 148 mana against rank 1's 122 and rank 3's 207, it lands more damage, and the shaman still runs dry only at the very end (out of mana 4.2 s against 1.4 s at rank 1, of 180 s). After the change, every other Elemental row still holds (re-run: `out/reconfirm-ele2.jsonl`). Test: `TestReMeasuredRankChoicesCastTheWinningRank` (failed first). | `sweep.py`; scratchpad `out/ele-lvb-*.json` on `r10/sb-before` | CONFIRMED-BUG | R9 S 5f1cbfacf | S ba2b209e9 |
| R3 | **Restoration druid's bare-60 Rejuvenation row went stale.** Round 9 adopted rank 10 (9841) at +5.2%. Rank 10 now loses -1.05% / -0.96% effective HPS to the rotation's own rank 4 (2090), and every rank from 5 to 10 loses to rank 4. The cause: round 10 made Swiftmend leave the heal over time up and read Regrowth first, from beta logs (fork 9badd243f). Before that, a bigger Rejuvenation fed a bigger Swiftmend. With 9badd243f reverted in a scratch engine, rank 10 wins again (+5.8%). The row was not re-measured after that fix. The fight still runs dry without the row (mana lasts 238 s of 300), so dropping it does not break the healer mana rule. Test: same as R2 (failed first). | `sweep.py` on `sb` and on `sb-exp` (9badd243f reverted) | CONFIRMED-BUG | R10 E 9badd243f | S d938ec8f4 |
| R4 | **Destruction's Shadow Bolt rank 4 at 30, re-checked on the warlock lane's branches** (site a5606e906 merged with this branch, fork 18cfcb100), where every warlock now casts Curse of the Elements before the bane. Rank 4 against the default rank 5: +0.60% / +0.55%, error 0.07%; before the merge it was +0.57% / +0.57%. Rank 3 loses 5.2%. The row holds beyond error. It is under the adoption rule's 1% (see Still open). | `sweep.py` with SITE/SB on the merged tree | Aligned | - | - |

## Search for missed rows

For every mana-bound spec and band (out of mana for more than a few seconds of the bare fight, `out/survey.jsonl`), I tried each ranked spell cast at least twice a fight at one, two and three ranks below its current rank, and at rank 1. That came to 457 runs at 3,000 iterations, seed 1 (`out/explore.jsonl`). The specs and bands: Balance 20/40/50, Beast Mastery and Marksmanship 20 to bare 60, Survival 40, all three mages 20 to bare 60 (plus raid 60 for Arcane and Fire), Shadow 20 to bare 60, Elemental 20 to bare 60, Destruction 20/40/50/bare 60, Demonology 50/bare 60, Retribution 50/bare 60. Every candidate above +0.25% was then confirmed at 10,000 iterations on both seeds.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | Beast Mastery's Arcane Shot rank 1 at 50: +0.32% / +0.32%, beyond error but under 1%. Rank 3 at 40: +0.32% / +0.32%. Not adopted (adoption rule). | `sweep.py` | not adopted | - | - |
| M2 | Elemental's Flame Shock rank 1 at 30: +0.42% / +0.42%, under 1%. Not adopted. | `sweep.py` | not adopted | - | - |
| M3 | Fire's Fire Blast rank 5 at 50: +0.16% / +0.05%, within error. Rank 4: -0.5%. | `sweep.py` | no row | - | - |
| M4 | Fire's Scorch at bare 60 (out of mana 24 s): every lower rank loses (rank 6 -0.3%, rank 1 -6.5%). Fire Blast and Pyroblast lower ranks lose too. No bare-60 row. | `explore.jsonl` | Aligned | - | - |
| M5 | Frost's Ice Barrier lower ranks gain +0.45 to +0.64% at 50 and bare 60. Ice Barrier is a defensive cooldown, which is never a rank choice (the table's notes). | `explore.jsonl` | not a rank choice | - | - |
| M6 | Every other lower rank tried loses: Balance's Wrath, Starfire, Moonfire and Insect Swarm at 20/40/50; Frost's Frostbolt at every band; Arcane's Frostbolt, Arcane Missiles and Arcane Blast; Shadow's Mind Flay and Shadow Word: Pain at 40 to bare 60; Elemental's Flame Shock at bare 60; Destruction's Incinerate, Immolate and Shadow Bolt; Demonology's Shadow Bolt; Marksmanship's Arcane Shot (rank 5 at 50 -0.5%). | `explore.jsonl` | Aligned | - | - |
| M7 | Multi-Shot has one rank in Forever (trainables: 2643, rank 0, level 18; spellranks lists no other player rank), so it has no rank choice. | `trainables/hunter.json` | Aligned (Forever change) | - | - |

## Is each row's rank learnable, and are its numbers the client's?

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | Every row's rank is in its class's `trainables/<class>.json` under the same spell name and rank number, with a learn level at or below the band's first level. None is a rune row: Lava Burst 408490 and 1238299 are SkillLineAbility AcquireMethod 0 trainer rows. The coach's `TestRankChoicesJoinTheTrainableFamilies` holds the names. | `learn.py` (every row OK); 70707 SkillLineAbility (scratchpad `r6/pipeline/layered70707`) | Aligned | - | - |
| L2 | **Coefficients and learn levels per rank.** I built each band's character in the engine (scratch dump) and compared every rank of every row's spell with spellconst: RequiredLevel against spell_level, BonusCoefficient and the DoT coefficient against the client sp_coefficient times the sub-20 level penalty, 1 - (20 - level) x 0.0375. There were 0 mismatches over 33 spec-bands and about 280 rank entries. For example, Flame Shock rank 1 (level 10) is 0.1338 direct and 0.0625 per tick (0.214 and 0.1 x 0.625); Moonfire rank 3 (16) is 0.1275 / 0.1105; Healing Wave rank 3 (12) is 0.4998. The penalty itself (fork 87743552b) applies only below level 20, as the cmangos and vmangos Classic cores do; Classic has no TBC-style penalty for ranks far below the caster's level. Forever keeps that rule (every row rank is learned at 10 or later). | `out/coefcheck.txt`; `sim/core/spell.go:777-809` | Aligned | - | - |
| L3 | **Costs.** Base costs match spellconst wherever no talent applies (Rejuvenation, Aimed Shot at 40, Frostbolt, Scorch, Fire Blast, Holy Light, Consecration for Protection, Mind Blast, Shadow Word: Pain). Where they differ, the difference is a taken talent: Balance's Wrath at 24% of 50 (Improved Wrath 50% plus Moonglow); Elemental and Restoration shaman -10% and -5% (Convection, Tidal Focus); Retribution's Consecration -10% at 30 to 50 and -50% at bare 60. The talents themselves belong to their class lanes. | `out/coefcheck.txt` | Aligned | - | - |
| L4 | The hotfix push carries no row for any id of any row's spell (233 ids checked against `hotfix-push.json`'s tables). | scratchpad check | Aligned | - | - |

## The mechanism

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **A chosen rank that is not learned yet.** `rankResolver` takes the row's rank only when `SpellIDOfRank` says it is learned at the level. Otherwise it falls back to the highest learned rank not above the authored one, which is always below the chosen rank. Every row's rank is learned by its first level (L1), so the fallback never fires today. | `sim/request/rotation_ranks.go:183-195`; `spellranks.go:307-365` | Aligned | - | - |
| X2 | **/sim, the API, the ranker and the ladder resolve the same rank.** All of them reach `rotation()` → `rewriteRotationRanks` with `specRankChoices(spec, preset)`: /sim and the API through `request.Build` (`rankPreset`: the request's preset, else the level's published one); the ranker sends `Preset` (raid pass raid, the rest bare; `leveling-bis/preset.go:68`); spec-breakdown sends `-preset` through `RewriteRotationRanks`; the ladder calls `rankResolver` directly. There is one function, so they cannot drift. | code read | Aligned | - | - |
| X3 | **The addon's rank table** (`data/pipeline/addonrotation.py`). Below the list's rank, where a rank carries a rune copy at the same level, it named the copy: Fire Blast 400618 to 400622 at 10 to 50, Renew 425268 at 10, Rejuvenation 417058/417059 at 10/20. SkillLineAbility marks 400618 to 400622 and 425268 AcquireMethod 3, and 417058/417059 have no row. The sim casts the trainer ids. I wrote the fix (same id band as the list, highest rank at a tied level, as `sameBandOrFirst` does; Counterattack's ranks 1 and 2 share level 30). Then I found that the platform lane had already landed the same logic, so I dropped my commit to avoid a conflict. Cross-check of that logic against the sim's resolver over every cast line of every spec at every addon band below 60: 1203 lines, 1 differs (X5). | `addoncheck2.py` | CONFIRMED-BUG | old (aeac1a64f, 2026-09-28; round 9's a65bd8111 kept it for a list's lower ranks) | platform S 45f7f5065 |
| X4 | **The coach (F5)** counted Ahn'Qiraj book ranks as "the highest trainable rank" (the trainables files list them at 60). At 60 it called the site's own rank low on 28 rotation lines across 16 specs: every warlock's Shadow Bolt, Corruption and Immolate; the mages' Frostbolt, Fireball and Arcane Missiles; the rogues' Eviscerate; Heroic Strike, Revenge and Battle Shout; Starfire; Serpent Sting; and more. Again the platform lane landed the fix first (it also covers Elemental's raid Lightning Bolt rank 9), and I dropped my duplicate. After removing book ranks, the only lines F5 would still flag: Elemental's raid Lightning Bolt (platform's fix covers it), Survival's Counterattack at 30 to 40 (resolved once X3 lands), and Protection's Divine Protection at 10 (a defensive). The coach and the sim read the same rank-choice bytes (`sim/rankchoices`). | `coachcheck.py` | CONFIRMED-BUG | old (84d4522ab, 2026-10-27) | platform S 49c70a7ef |
| X5 | The addon's Water Shield line at 10 names 409941 (a level-1 copy). The sim leaves 408510 as written, because Water Shield has no ranked chain. This is not a rank choice and predates this round; it goes to the platform lane. | `addoncheck2.py` | OPEN (other lane) | old | - |

## Before/after

| spec, band | metric | before, seed 1 | after, seed 1 | change seed 1 / 50001 |
|---|---|---|---|---|
| shaman-elemental, bare 50 | DPS | 151.8 | 153.2 | +0.92% / +1.00% |
| druid-restoration, bare 60 | effective HPS | 412.5 | 416.9 | +1.07% / +0.97% (mana lasts 224 s to 238 s of 300) |

"Before" is `$R/sb-before` from `$R/site-before/sim`, and "after" is this branch on the pin. No other spec or band moves: the two commits change only those two rows.

## For other lanes

- **Integrator:** `TestRotationLadder/shaman-elemental` moves (the 50 rung casts Lava Burst rank 2), and so may `druid-restoration` at 60 bare. Not regenerated here. The ranker's bare-60 Restoration druid sets and the ranks of the two bands should be re-measured at the pin. Both commits touch only `data/curated/rank-choices.json` and its embed (`make rank-choices-sync`).
- **Heal profile lane:** this lane dropped Restoration druid's bare-60 Rejuvenation band and left your 50 to 59 band untouched. If you rewrite Rejuvenation's rows, take my deletion.
- **Platform:** X3 and X4 are fixed on your branch (45f7f5065, 49c70a7ef). I independently confirmed both and the 1203-line cross-check above; X5 (Water Shield 409941 at 10) is yours.
- **Warlock lane:** R4. The Shadow Bolt row holds on your branches.
- **Hybrids:** Feral (cat) runs out of mana for 128 to 136 s of the bare fight at 40 to bare 60 (`out/survey.jsonl`); no rank choice touches it, but it looked worth a look.

## Still open

- **Which bar do rank rows need?** The owner's round-9 law says a row needs "a gain beyond error at seeds 1 and 50001". The round brief says the adoption rule (beyond error and at least 1%) applies to rank choices too. Six kept rows clear error but not 1%: Aimed Shot rank 1 (BM 40, BM bare 60, MM bare 60; +0.25 to +0.39%), Fire Blast rank 4 at 40 (+0.72 / +0.79%), Shadow Bolt rank 4 at 30 (+0.55 to +0.60%), Prayer of Healing rank 1 at bare 60 (+0.29%). I kept them, as my brief says to drop only rows that no longer win beyond error and the integrator agreed (the round-9 law is the specific rule for rank rows), and did not adopt new sub-1% candidates (M1, M2). One owner line settles both.
- Lava Burst rank 2 over rank 1 at 50 is +0.92% at seed 1, a hair under 1%. I adopted it because it corrects an existing row's rank, and the row itself gains +3.6% over the default. If the owner reads the 1% bar as applying to a change of rank too, revert ba2b209e9.

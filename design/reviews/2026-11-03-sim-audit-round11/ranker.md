# Sim audit round 11: ranker (2026-11-02)

This lane owns `sim/cmd/leveling-bis`, the BiS pipeline code and `data/curated/content-phases.json`. It made no engine change.

**Setup**
- Site: `wt/r11-ranker` from origin/main 8f047a62f. Engine: the pin cfcb76e80 (`$R/base-engine`, read only).
- A local copy of the held data run in `scratchpad/r11/ranker/root`, never committed: `fetch-wago-hotfixes`, `normalize --allow-shrink`, `loot --allow-shrink` and `simdb` on 1.60.1.70707, then the committed talents put back.
- Re-rank: 10 of 28 specs, all bands, both factions (120 rows): Balance, Feral, Bear, Restoration druid, the three hunters, the three mages. Then the owner stopped every BiS ranking run for the rest of the round (CPU), so the other 18 specs, the healer re-rank on the healprofile branch and the after-fix re-rank are listed under "Measure at refresh". Nothing of the lane is still running.

**In short**
- 2 CONFIRMED-BUGs, both `old`, both fixed with tests first: the PvP rank cap and the PvP source line read the client's RequiredPVPRank as the in-game rank.
- 1 CONFIRMED race pick outside the adoption bar, handed off: Holy priest's guide races (gnome, troll) lose to human and undead.
- The 165 "gone" ids are not gone; no pick or candidate drops.

## Confirmed bugs

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| K22 | **The PvP rank cap sat at rank 6, not rank 10.** The owner's ruling (2026-09-29, kept in round 9): "rank 10 is a few weeks of casual battlegrounds, rank 11+ needs sustained top-bracket play". The ranker compared that 10 with loot.json's `rank` / `title_rank`. Those are the client's RequiredPVPRank, which runs 5 (Private / Scout, rank 1) to 18 (Grand Marshal / High Warlord, rank 14). So the cap really stopped at Knight / Stone Guard, rank 6. Knight-Lieutenant / Blood Guard gloves and boots (rank 7), Knight-Captain / Legionnaire chest and legs (rank 8), and Lieutenant Commander / Champion helms and shoulders (rank 10) were never a default pick. | - ItemSparse 70707 RequiredPVPRank on the PvP rewards takes the values 5 to 18.<br>- The pipeline's own ladder `data/pipeline/loot/pvp_faction.py` `rank_title` uses index rank − 5, so field 10 = Knight, field 14 = Lieutenant Commander. `web/src/lib/items/source-labels.ts` already said "these ranks + 4".<br>- warcraft.wiki.gg "Honor system (pre-2.0)" (fetched 2026-11-02): rank 6 Knight / Stone Guard, rank 7 Knight-Lieutenant / Blood Guard (boots, gloves), rank 8 (chest, legs), rank 10 Lieutenant Commander / Champion (helm, shoulders); rank 11 needs the top 1.4%. That matches the ruling's own wording for 10 and 11.<br>- On the local re-rank (120 rows), the highest PvP pick is field 10 (rank 6) in 11 rows; nothing above it.<br>- Fix: `inGamePvpRank` (field − 4) before the cap (band.go). 333 armor and weapon items at ranks 7 to 10 with an open source become default-eligible, nearly all at required level 50 to 60.<br>- Tests first: `TestThePvpRankCapReadsTheInGameRank` failed, then passed. The old fixtures that paired titles with raw fields (`TestPvpRankExceedsCap`, `TestARivalAboveThePvpRankCapDoesNotCoverTheSlot`, the title_rank test) now use the client scale. | CONFIRMED-BUG (picks move at the refresh) | old: S 46fd13bba (2026-09-29). Rounds 8 and 9 reviewed the cap and title_rank and missed the scale. | S 61d4a79c2 |
| K23 | **PvP source lines named the client field as the rank.** The BiS page, the item tooltips and the ranker's report read "PvP rank 18 · Grand Marshal · Alliance" and "PvP rank 10 · Stone Guard · Horde". Grand Marshal is rank 14 and Stone Guard rank 6. 3,000+ source lines in the published BiS files carry it. | Same evidence as K22.<br>Fix: the label names the in-game rank, in Go (`pvpSourceLabel`) and on the page (`web/src/lib/items/source-labels.ts`, used by the BiS source cell and the tooltip). A field below 5 reads "PvP reward".<br>Tests first: `TestThePvpLabelNamesTheInGameRank` failed; the vitest cases in source-labels, copy, source-cell and tooltip failed 8 of 85 on the old code, then passed. | CONFIRMED-BUG (copy; the page moves at deploy, the BiS files at the refresh) | old: the third and fourth wow-player sweeps (2026-09-29) | S 61d4a79c2, S 3ae1922fc |

## Race picks (round 10 checklist: Undead, Touch of the Grave)

Spec-breakdown, committed data, published 60 gear with only the race swapped, 3,000 iterations, seeds 1 / 50001, under `nice` one run at a time. Healers are effective HPS on the heal profile.

| # | spec, faction, preset | races | result | verdict | origin | fix sha |
|---|---|---|---|---|---|---|
| P1 | Shadow, Horde, raid 60 | undead (guide) / troll | 718.1 / 717.9 against 711.4 / 710.7 DPS (+0.9 / +1.0%) | Aligned: undead stays | - | - |
| P2 | Shadow, Alliance, raid 60 | gnome (guide) / human / dwarf / night elf | 711.2 / 711.0; 713.8 / 713.3; 712.9 / 712.7; 713.8 / 713.5 | Aligned: human +0.35%, inside the 1% bar, so gnome stays | - | - |
| P3 | Discipline, Horde, raid 60 | undead (guide) / troll | 625.4 / 625.4 against 598.5 / 598.7 | Aligned | - | - |
| P4 | Discipline, Alliance, raid 60 | human (guide) / gnome | 623.0 / 622.7 against 616.8 / 616.7 | Aligned | - | - |
| P5 | **Holy priest, Horde** | troll (guide) / undead | raid 60: 693.7 / 693.5 against 712.0 / 712.6 (**undead +2.6%**). Bare 60: 312.9 / 313.0 against 340.6 / 340.6 (**+8.8%**). The gain is Dark Sacrifice (1277328, the Forever undead priest racial): about 2,060 mana a fight in 5 casts, "mana lasts" 455 s → 696 s. Troll's Berserking adds no throughput on a mana-bound healer. | **CONFIRMED: the race pick fails the adoption bar.** Not changed here: it rests on Dark Sacrifice's engine amounts, which only round 2 looked at (its race mask). Handed off (below). | old (no round measured Holy's race) | - |
| P6 | **Holy priest, Alliance**, raid 60 | gnome (guide) / human / dwarf / night elf | 706.2 / 706.5; **720.5 / 720.2 (+2.0%)**; 695.1 / 694.8; 697.8 / 697.6 | **CONFIRMED: human beats gnome beyond 1%** (Divine Grace and the human Spirit bonus). Handed off with P5. | old | - |
| P7 | Retribution, Protection and Holy paladin, Horde | undead only (combos.json: Horde paladins are undead) | No race choice. Their gear picks need the re-rank after round 10's Touch of the Grave change. | Measure at refresh | - | - |

## The 165 ids (items lane cross-check)

| # | finding | evidence | verdict |
|---|---|---|---|
| G1 | **None of the 165 is gone, and none of them would drop a pick.** My first local run, like round 10's, lacked `raw/wago-hotfixes/` and lost exactly the items lane's 165 ids (same set). With `fetch-wago-hotfixes` layered in, `normalize` keeps all 25,853 items, and the class files lose none (paladin gains its 257 axes). | Published BiS files: 0 picks and 0 alternatives among the 165. loot.json: only 2 of the 165 have any source (9399 Precision Arrow, Uldaman; 23192 Tabard of the Scarlet Crusade), neither a ranked slot. 110 are gear, mostly item level 30 / required level 1 copies of raid items (Thunderfury 289294, Benediction 289296 and others) with no source, so never candidates. | Aligned with the items lane (S 3484405dd on wt/r11-items): the refresh must fetch the item hotfix layer and should NOT accept a 165-item drop. `--allow-shrink` is still needed, for wowhead's stats gate (289 items take wowhead stats against 689 before, because the client layer now carries them). |

## Fresh audit of the ranker

Checked on the 120 local rows with `audit11.py`: round 10's checks plus armor by class and level, weapon subclass by class, two-hander with an off hand, required level, the client's RequiredSkill and RequiredPVPRank on every pick, Unique and LimitCategory pairs, the 165 ids. **0 violations.** Every row was also screened for a listed alternative measured more than 1% above its pick: none.

| # | area | what I checked | verdict |
|---|---|---|---|
| A1 | Armor by level | `armorAvailableLevel`: mail 40 for hunter and shaman, plate 40 for warrior and paladin (memory forever-class-rules). 0 sub-40 mail picks on hunters. | Aligned |
| A2 | Weapons by class | `classWeaponSubclasses` against client SkillLineAbility 70707. The client masks also give rogues one-hand axes (196, mask 79) and druids polearms (200, mask 1031); both stay out on round 8's trainer evidence (K1 there). Shamans: fists (473), axes and maces both hands; dual wield 674 for shaman is AcquireMethod 3 (a rune), not learnable. | Aligned |
| A3 | Dual wield | Shamans never get an off-hand weapon (`offHandPool`); hunters and warriors dual wield from 20, the lowest band. | Aligned |
| A4 | Professions | The ranker has no profession gate; the design shows crafted items with their profession as a note. On the client, 53 Engineering items need the skill to equip. Engineering is the only equip-gated profession picked (9 of 120 rows, goggles); no row needs more than two primary professions. | Aligned (design), recorded |
| A5 | Phases, side, standing | Raids, Dire Maul, The Watchers and the Undermine traders closed; battlegrounds, Darkspear and Powderfuse open; 0 wrong-side, 0 revered or exalted below 60. World drops (`world_drop`) are picked on purpose, as auction-house items (band.go's sourceKindPriority doc). | Aligned |
| A6 | Iterations and error bars | Verify 300 iterations, swap only beyond 1% and the combined error; trinkets ordered at 100, winner refined to stderr < gain/4 (cap 2000), ceiling re-measure of the top 3 at 2000 when the winner has no gain; the healer scale 7 now stops at `api.MaxIterations` (round 10 K20). No "between 1 and 10000" refusal in the 10 logs. Trinket candidates share one seed, so the ordering is common-random-number paired. | Aligned |
| A7 | Preset per pass | Bare at every band (the leveling kit, the /sim bare set), raid only at 60; each pass sends its preset for the rank choices. Encounter is the type-neutral default, target level by band. | Aligned |
| A8 | Frost after Fingers of Frost | Re-ranked all 12 rows on the current engine; every row changes against the published file (which predates rounds 9 and 10). Band 20's neck is empty (`no_dps_value`): the only sourced necks carry stamina and spirit, which weigh nothing for Frost there. | Aligned (by design); published at the refresh |
| A9 | Healer slots | Restoration druid band 30 and 40: the healing-power sweep is not significant on today's sub-60 profile (−0.005 ± 0.014 per point), so the ranker ranks by the ladder sweep. This is the healprofile lane's ruling 1. Balance bare 60 leaves the idol empty (no gain beyond error); raid 60 takes Idol of the Moon (+2.9 ± 1.0 at 2000). | Measure at refresh (healprofile) |
| A10 | Tanks | Bear 60: Insignia, Talisman of Arathor and Defiler's Talisman picks, no idol. No change to tank scoring; round 9's two-seed spread stays open. | Aligned (one spec looked at) |

## Before/after

No sim DPS moves from this lane. K22 and K23 move BiS picks and source lines only, at the refresh. K23 moves the page's PvP line at the next deploy.

| change | measured | result |
|---|---|---|
| K22 | data only (no ranking runs allowed) | 333 rank-7-to-10 armor and weapon items gain default eligibility, required level 50 to 60 (2 at 25). Rows at bands 50 and 60 can move; nothing below. |
| K23 | label tests | "PvP rank 18 · Grand Marshal" → "PvP rank 14 · Grand Marshal" |

## Measure at refresh

Each needs a full re-rank, which the owner held this round.
- **Every spec, bands 50 and 60, every slot a PvP reward can fill (head, shoulder, chest, legs, hands, feet, weapons):** K22 lets ranks 7 to 10 compete.
- **The 18 specs not re-ranked locally, every band:** Retribution, Protection paladin, Holy paladin (Libram of Infusion in the relic slot, E ac11976ed), the priests, rogues, shamans, warlocks, warriors. Not audited on output this round.
- **Holy paladin, Retribution, Protection paladin, Horde, every band, every slot:** Touch of the Grave changed in round 10 and Horde paladins are undead.
- **Five healers, bands 20 to 50, every slot, and the trinket and relic slots at 60:** on the healprofile lane's pull multiplier (wt/r11-healprofile eaa2b2285). I built the ranker with it but could not run it.
- **Holy priest, both factions, every band:** if P5/P6 are adopted the race changes and every pick is re-measured.

## For other lanes

- **Healers lane / integrator: Holy priest's races (P5, P6).** On today's engine undead beats troll by 2.6% (raid 60) and 8.8% (bare 60), and human beats gnome by 2.0% (raid 60), at 3,000 iterations, seeds 1 and 50001. Before adopting, check Dark Sacrifice (1277328) against the 70707 client: health taken, mana given, cooldown, and when the rotation casts it. Then set `recommendedRaces: [human, undead]` and the build string's race in web/src/content/guides/priest/holy.md. Discipline already uses human and undead.
- **Integrator, data run.** Run `fetch-wago-hotfixes` (S 3484405dd folds it into `fetch`) before `normalize`. With it no item drops, but `--allow-shrink` is still needed for the wowhead-stats gate (289 against 689). Do not record a 165-item drop.
- **Items lane.** Agreed on the 165; the census here found no pick, alternative or sourced candidate among them.

## Still open

- From round 10: The Watchers' gate; `TestRepSideMatchesTheBuildsFactionIDs` still skips until loot.json is rebuilt (held output).
- From round 9: the two-seed tank spread; Glimmering Staff's spell power; the dagger rule for Combat; the off-hand weapon-DPS term; two copies of a non-unique ring.
- `data/pipeline/loot/pvp_faction.py` and `models.py` docstrings still give "PvP rank 9 · Master Sergeant" as the example; comment only, no output reads it.

# Sim-accuracy round 12: ranks lane

Scope: the per-band rank-choice mechanism (`sim/request/rotation_ranks.go`, `rank_choices.go`, `sim/rankchoices`), its new faction key, and every non-healer row of `data/curated/rank-choices.json`. The bare-60 healer rows were re-measured too, because the faction key splits three of them (the heal profile lane owns the sub-60 healer rows, which this lane did not touch). Measured on site `wt/r12-ranks` from the landing c6bdc855c, engine at the pin 2606bc05e (no engine change in this lane), client 1.60.1.70753. Each band's published gear and race, 10,000 iterations, seeds 1 and 50001, both factions, unless marked. Scripts and raw output are in scratchpad `r12/ranks/` (`reconfirm.py`, `sweep.py`, `explore.py`, `survey.py`, `firegrid.py`, `ba.sh`, `tools/`, `out/`). The scratch Go override that lets spec-breakdown read a substitute table (`RC_FILE`) is not committed.

Error bar (as in round 11): a DPS or threat row passes when its gain at each seed beats twice the combined standard error. The healer fight prints no error, so a healer row passes at more than 0.2% with the two seeds agreeing.

Owner rulings applied: a rank row is adopted when it gains beyond error at seeds 1 and 50001, with no 1% bar. A row names a faction only where Alliance and Horde differ beyond error. This lane reads "differ" as: one faction's row gains beyond error while the other's does not (it ties or loses), or the two factions' best ranks differ.

## The faction key (mechanism)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | **Rank rows had no faction.** Added an optional `"faction": "alliance" \| "horde"` on a row. A row without one applies to both factions. An Alliance row and a Horde row may share a band and preset. The loader refuses an unknown faction, and refuses two rows that could apply to the same request (it now checks every earlier row, not only the one before). `request.Build` takes the faction from the character's race (`FactionOfRace`; Skyborne by its two slugs). `RewriteRotationRanks` takes the faction, so spec-breakdown (`-faction`) and rotation-search (the BiS band's faction) cast what /sim does. The ladder follows its race. talent-search and the ranker reach it through `request.Build`. Tests: `TestRankChoicesFollowTheFaction`, `TestFactionOfRaceCoversEveryRace`, `TestBuildCastsTheFactionsRankChoice`. All failed first. | `sim/request/rank_choices.go`, `faction.go`, `rotation_ranks.go:76`, `request.go:222-230` | owner ruling | - | S 6230f7a28 |
| F2 | **Every reader of the table.** The coach (`rankchoices.ChosenRanks`) takes a faction's row for either faction, as it already took either preset. A log's rank is no mistake where the site casts it for one faction (`TestChosenRanksReadsTheCuratedRowsAtALevel`, Mind Blast at 55 is now [2 3]). The addon export (`data/pipeline/addonrotation.py`) carries one rotation per band for both factions. Where a faction row changes a line's rank, the line keeps the both-faction id and names each faction's id in `faction_spell_ids`. Data.lua gets `factionSpellId = { alliance = .., horde = .. }`, and `Rotation.lua` shows the player's own id (`Compat.playerFaction`) on a copy of the band. The first commit made the export refuse a faction row below 60. That refusal would have failed the regen on this round's rows, so the second commit replaced it. Tests: `test_a_faction_rank_choice_below_sixty_names_each_factions_rank`, `test_the_lua_names_each_factions_id_only_where_they_differ`, `rotation_spec.lua` "Rotation with a faction's rank". All failed first. | `api/internal/coach/tables.go:194`, `sim/rankchoices/lookup.go`, `addonrotation.py`, `addonlua.py`, `addon/ForeverSixty/Rotation.lua` | owner ruling | - | S 75cb822f3 |

## Re-judgement of every row (landing pin, both factions)

Every row's gain against the same table with that row removed, then each neighbouring rank against the row. The table shows Horde / Alliance at seeds 1 / 50001. Source: `out/reconfirm-{horde,alliance}.jsonl`, and `out/final-*.jsonl` after this lane's changes.

| spec | spell | band | rank | Horde row vs default | Alliance row vs default | err | verdict |
|---|---|---|---|---|---|---|---|
| druid-balance | Moonfire | 30 | 3 | +4.52 / +4.63 | +4.32 / +4.44 | 0.23 | holds (rank 4 ties: +0.14 / +0.12 A) |
| druid-balance | Wrath | 30 | 4 | +6.44 / +6.45 | +6.58 / +6.60 | 0.21 | holds |
| druid-balance | Insect Swarm | 30 | 1 | +2.30 / +2.29 | +2.43 / +2.45 | 0.20 | holds |
| hunter-beast-mastery | Aimed Shot | 40 | 1 | +0.34 / +0.32 | +0.24 / +0.25 | 0.10 | holds (under 1%; kept by the ruling) |
| hunter-beast-mastery | Aimed Shot | bare 60 | 1 | +0.57 / +0.56 | +0.46 / +0.41 | 0.11 | holds (under 1%) |
| hunter-marksmanship | Aimed Shot | bare 60 | 1 | +0.42 / +0.41 | +0.35 / +0.40 | 0.12 | holds (under 1%) |
| mage-arcane | Frostbolt | 30 | 4 | +0.84 / +1.08 | +0.67 / +0.81 | 0.25 | holds |
| mage-fire | Fire Blast | 40 | 4 | +1.29 / +1.44 | +0.70 / +0.88 | 0.30 | holds (under 1% for Alliance) |
| mage-fire | Scorch | 30 / 40 / 50 | 1 | +2.23, +7.78, +3.22 | +0.98, +5.70, +2.55 | 0.30 | holds |
| paladin-protection | Consecration | 40 / 50 / bare 60 | 1 | +1.32, +1.96, +3.45 (threat) | +1.49, +2.12, +3.83 | - | holds |
| paladin-retribution | Consecration | 30 / 40 / 50 | 1 | +1.26, +2.35, +7.67 | +1.09, +2.06, +7.53 | 0.16 | holds |
| paladin-retribution | Consecration | bare 60 | 4 | +1.57 / +1.65 | +1.76 / +1.87 | 0.14 | holds |
| priest-shadow | Shadow Word: Pain | 30 | 3 | +1.86 / +1.87 | +3.63 / +3.65 | 0.15 | holds |
| priest-shadow | Mind Blast | 40 | 4 | +1.35 / +1.40 | +1.35 / +1.39 | 0.21 | holds |
| priest-shadow | Mind Blast | 50 | 2 | **-0.06 / -0.11** | +2.43 / +2.49 | 0.21 | **R1: split by faction** |
| priest-shadow | Devouring Plague | 30 / 40 / 50 | 1 | +0.46, +3.99, +2.71 | +1.24, +3.90, +4.15 | 0.21 | holds |
| priest-shadow | Shadow Word: Death | 50 / bare 60 | 2 | +0.45, +0.62 | +1.28, +1.27 | 0.20 | holds (under 1% for Horde) |
| shaman-elemental | Lightning Bolt (15207) | 30 / 40 / 50 / bare 60 | 4 | +8.0, +16.3, +36.3, +30.2 | +7.8, +15.9, +35.1, +31.1 | 0.28 | holds |
| shaman-elemental | Flame Shock | 40 / 50 | 1 / 2 | +5.42, +4.01 | +5.43, +3.42 | 0.26 | holds |
| shaman-elemental | Lava Burst | 40 / 50 / bare 60 | 1 / 2 / 2 | +1.11, +4.05, +1.83 | +1.32, +3.60, +1.70 | 0.25 | holds (rank 1 at 50 is now within error of rank 2 after this round's rows: -0.16 / -0.12 H; kept) |
| warlock-destruction | Shadow Bolt | 30 | 4 | +0.60 / +0.55 | +0.53 / +0.58 | 0.12 | holds (under 1%) |
| priest-discipline (healer) | Binding Heal | bare 60 | 1 | **-0.03 / -0.13** | +0.51 / +0.48 | - | **R5: Alliance only** |
| priest-holy (healer) | Prayer of Healing | bare 60 | 1 | +0.29 / +0.29 | **-0.15 / -0.09** | - | **R6: Horde only** |
| priest-holy (healer) | Lightwell | bare 60 | 1 | +1.29 / +1.33 | **-0.06 / +0.06** | - | **R6: Horde only** |
| priest-holy (healer) | Flash Heal | bare 60 | 7 | +8.80 / +8.77 | +8.05 / +8.15 | - | holds |
| paladin-holy (healer) | Holy Light | bare 60 | 4 | +4.60 / +4.37 | +4.77 / +4.68 | - | holds. Horde rank 3 heals 4.2% more but ends with mana unspent (healer mana rule, as round 11 H2) |
| shaman-restoration (healer) | Healing Wave / Riptide | bare 60 | 2 / 2 | +19.5, +3.4 | +21.4, +3.4 | - | holds |

The six rows round 11 left under 1% all gain beyond error at both seeds for both factions, so the ruling keeps all six: the three Aimed Shot rows, Fire Blast rank 4 at 40, Shadow Bolt rank 4 at 30, and Prayer of Healing at bare 60 (Horde; see R6). Lava Burst rank 2 at 50 beats the default by +4.4% / +4.0%, so it also stays.

## Findings

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Horde Shadow at 50 to 59 cast the wrong Mind Blast.** The rank-2 row (8102) no longer beats the top learned rank for an Undead priest (-0.06 / -0.11%). Rank 3 (8103, 28 mana more a cast) beats the default by +0.41 / +0.36% (error 0.21) and beats rank 2 by +0.46 / +0.47%. For Alliance, rank 2 still wins by +2.4% and rank 3 loses 2.0%. Row split: Horde rank 3, Alliance rank 2. On int/r11 (site f263dbcc8, engine 47df4a89d), rank 2 still won for Horde by +1.6%. The landing's Shadowform 20 s before the pull, plus the launch-day Shadow builds (Mental Strength), leave Horde with more mana (out of mana 12.7 s instead of 17.9 s), so a bigger Mind Blast pays. A Classic player with Dark Sacrifice and mana to spare casts the bigger rank. Test: `TestRound12RankChoicesCastTheWinningRank` (failed first). | `out/reconfirm-*.jsonl`, `out/mb50.txt`, `out/batch5.txt` | CONFIRMED-BUG | landing f726a437b (and the Shadow build change) | S 86ca0f6fe |
| R2 | **Elemental's paced top-rank Lightning Bolt line had no rank row.** The landing added a second Lightning Bolt line, 15208 (rank 10), cast while mana is ahead of the fight. Its note says "rank 9 (or the band's rank choice below 60) after". But rows key on the authored id, so at 50 to 59 and at bare 60 that line cast the band's top rank (rank 9 at 50, rank 10 at bare 60), while the filler line cast rank 4. Rank 7 (10391, 135 mana against 190 and 220) on the paced line gains +0.43 / +0.40% Horde and +0.52 / +0.51% Alliance at 50 (error 0.18), and +0.41 / +0.41 and +0.40 / +0.40 at bare 60. Ranks 6 and 8 lose to rank 7. At 30 and 40 the line almost never fires (0.01 casts a fight), and no rank moves it. Rows: 15208 rank 7 at 50-59 and bare 60. Raid 60 is untouched. Test: same, plus `TestThePublishedRotationTakesTheCuratedRankChoices` (updated: the paced line now casts rank 7). | `out/lb15208*.txt`, `out/final-*.jsonl` | CONFIRMED-BUG | landing b2e4a3bc5 | S 86ca0f6fe |
| R3 | **Horde Fire at 30 to 39 casts Fire Blast rank 3.** The landing's opener (Pyroblast before the pull) changed the Fire fight at 30. A joint search over Fire Blast {3, default}, Pyroblast {1, 2, default} and Scorch {1, default} at 10,000 iterations (`out/firegrid.jsonl`) found the best confirmed set: Fire Blast rank 3 with the existing Scorch rank 1. That set beats the current table by +1.70 / +1.69% (error 0.28). Pyroblast rank 2 on top adds +0.25 / +0.22%, within error, so it is not adopted. For Alliance, Fire Blast rank 3 is +0.02 / -0.01% and Pyroblast rank 1 loses 1.2%, so the row is Horde only. On int/r11, Fire Blast rank 3 for Horde was +0.35 / +0.28% (error 0.27: not beyond error at seed 50001). Test: same. | `out/b6-*.txt`, `out/firegrid.jsonl`, `out/r11-origin.txt` | CONFIRMED-BUG | landing 61619dbe7 | S 86ca0f6fe |
| R4 | **Beast Mastery at bare 60 casts Arcane Shot rank 1.** +0.34 / +0.25% Horde and +0.28 / +0.15% Alliance (error 0.11). Rank 2 is in between. On int/r11 it was +0.02 / -0.01%. The landing's Hunter's Mark re-mark spends mana, and bare BM runs dry for 87 s of 180. Test: same. | `out/b6-*.txt`, `out/r11-origin.txt` | CONFIRMED-BUG | landing c63fc42ac | S 86ca0f6fe |
| R5 | **Horde Discipline's bare-60 Binding Heal row stopped gaining.** Rank 1 is -0.03 / -0.13% for Horde, and the same on int/r11, where round 11 measured it +3.4% only on the healers branch before the merge. For Alliance it still gains +0.51 / +0.48%. Row now Alliance only. (Bare-60 healer row: the heal profile lane should review.) Test: same. | `out/reconfirm-*.jsonl`, `out/r11-reconfirm-disc-shadow.jsonl` | CONFIRMED-BUG | R11 (already lost on int/r11; Lesser Heal rank 3 b3690c448 or the heal coefficient fix 074edfadf, not isolated) | S 86ca0f6fe |
| R6 | **Holy priest's bare-60 Prayer of Healing and Lightwell rank 1 gain only for Horde.** Horde: +0.29% and +1.3%. Alliance (human since round 11): -0.15 / -0.09% and -0.06 / +0.06%, within the healer bar. Both rows are now Horde only. (Bare-60 healer rows: for the heal profile lane to review.) Test: same. | same | faction key (ruling) | - | S 86ca0f6fe |
| R7 | **Alliance Shadow at bare 60: Devouring Plague rank 2, not the owner's rank 1.** The "known case" changed under the landing. Horde now loses heavily to any lower rank (rank 1 -5.8%, rank 5 -0.8%), so it casts the top rank, as round 11 left it. Alliance gains with every lower rank: rank 1 +0.40%, rank 2 +0.77 / +0.72% (error 0.15), rank 3 +0.69%, rank 4 +0.51%. Row: Alliance rank 2 at bare 60. Raid 60 is untouched (test case). Test: same. | `out/dp60-*.txt` | faction key (ruling) | - | S 86ca0f6fe |
| R8 | **Rows round 11 held only for being under 1%, now adopted.** Beast Mastery's Arcane Shot (M1): rank 3 at 40-49 is +0.47 / +0.41% H and +0.36 / +0.35% A (error 0.10; ranks 2 and 4 lose to it). Rank 1 at 50-59 is +0.32 / +0.24% H and +0.31 / +0.33% A (error 0.09). Elemental's Flame Shock rank 1 at 30-39 (M2): +1.37 / +1.36% H and +2.59 / +2.59% A (error 0.25); round 11 measured it at +0.42% before the landing's rotation (Mana Spring before the pull). Test: same. | `out/m1-*.txt`, `out/m2-*.txt`, `out/final-*.jsonl` | ruling (round 11 M1, M2) | - | S 86ca0f6fe |
| R9 | **Horde Destruction at 40 to 49 casts Corruption rank 3.** +0.17 / +0.17% (error 0.14). Alliance is -0.08 / -0.06%, so the row is Horde only. Rank 4 loses to rank 3. Int/r11 gives the same +0.17%, so this predates round 12. Round 11's exploration at 3,000 iterations missed it. Under the round-9 law it would already have been a row. Test: covered by `TestCuratedRankChoicesAreCastableDownranks`; no per-row case. | `out/b6-*.txt`, `out/r11-origin.txt` | CONFIRMED-BUG | old (round 11's exploration looked and missed it) | S 86ca0f6fe |
| R10 | Destruction's Immolate rank 4 at 40: +0.11 / +0.20% H (error 0.13), not beyond at seed 1. Not adopted. | `out/b6-*.txt` | not adopted | - | - |
| R11 | Enhancement's Strength of Earth Totem rank 1 at bare 60 (round 9, held because Enhancement is not mana-bound). That rejection was not for being under 1%, so the ruling does not reopen it. | round-9 ranks.md | not reopened | - | - |

## Search for missed rows

The survey covered every non-healer spec at 20/30/40/50/bare 60/raid 60 at 3,000 iterations (`out/survey.jsonl`). On the 45 spec-bands out of mana for more than 3 s, the search tried each ranked spell cast at least twice a fight at one, two and three ranks below its rank, and at rank 1, on every authored id of its family. That was 406 runs at 3,000 iterations (`out/explore.jsonl`). Candidates above +0.15% were confirmed at 10,000 on both seeds and both factions (R3, R4, R9, R10). The rest:
- Ice Barrier (Frost 50 and bare 60) is a defensive and never a rank choice.
- Lightning Bolt 15208 at 20 and 30 (+0.2%) is noise on a line that does not fire.
- Everything else loses or sits within noise.

Feral (cat) runs dry for 128 to 136 s of the bare fight at 40 to bare 60. Its ranked casts cost energy, so no rank row applies (the same as round 11).

## Is each row's spell still there, still cast, still the client's?

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | Every row's rank is a trainer rank (`trainables/<class>.json`, 70753) learned by the row's first level, and none is a rune copy. Fire Blast rank 4 has a rune twin, 400620; the sim casts 8412. | `tools/learn.py` | Aligned | - | - |
| L2 | Every row's spell is still cast by its rotation (`tools/aplcheck.py`). The Aimed Shot rows act on the one pre-pull Aimed Shot, the rotation's only Aimed Shot (it was so before the landing too). Holy paladin's and Holy priest's rotations name a second rank of Flash of Light and Flash Heal (19750, 10917); those are healer rows below 60 and belong to the heal profile lane. The only non-healer line that bypassed its spell's rows was Elemental's 15208 (R2). | APL walk | Aligned except R2 | - | - |
| L3 | Client data: spellranks and spellconst for every rank of every row's spell are byte-identical between 70707 and 70753 (`tools/datacheck.py`). The committed 70753 hotfix push carries no Spell, SpellEffect, SpellMisc or SpellPower row. wago.tools showed no 70753 push for those tables on 2026-11-03 at 08:13 CST (the 70707 search returns 1,276 SpellEffect entries, so the search works). | `hotfix-push.json`; wago.tools/hotfixes | Aligned | - | - |
| L4 | Amounts, levels and costs: the fork's client-damage tests for every row class pass on the pin (druid, hunter, mage, paladin, priest, shaman, warlock: `Wrath`, `Moonfire`, `Insect Swarm`, `Arcane Shot`, `Aimed Shot`, the mage and priest ladders, `Lightning Bolt`, `Flame Shock`, `Lava Burst`, `Consecration`, the Shadow Bolt table). The vendored client file equals data/builds/1.60.1.70753/spellconst for all seven classes. A scratch dump of each band's character (354 rank entries over 44 spec-bands) agrees with spellconst on spell level and coefficient (including the below-20 penalty). The exceptions are talent-explained costs: Moonglow 25% (Forever's three ranks are 8/17/25%) plus Improved Wrath; Efficiency on Aimed Shot; Mental Agility; Convection; Retribution's Consecration talents (all as round 11 L3). | `out/engine-client.txt`, `out/coefcheck.txt` | Aligned | - | - |
| L5 | Arcane Shot: the engine gives it a spell-power coefficient (0.204 / 0.3 / 0.429 by rank, `sim/hunter/arcane_shot.go:14`). The client effect's `sp_coefficient` is 0.0 (spellconst 3044 to 14287, effect 2). Bare hunters carry no spell power, so the rows do not move, but gear with spell damage would. Single source (client table; the engine's numbers are Classic's). | spellconst/hunter.json | OPEN (hunter owner: does Forever's Arcane Shot scale with spell power? A tooltip with spell damage gear, or a log, settles it) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | The engine names Skyborne's faction Unknown (`sim/core/racials.go:616` lists only the eight old races), while the site and BiS files place High Order and Windshaper Skyborne. The rank rows follow the site's race map, so they are unaffected. Engine faction checks (buff faction, item restrictions) would see Unknown for a Skyborne character. | engine code read | OPEN (auras lane) | old | - |

## Before/after

Spec-breakdown, bare preset, 10,000 iterations. "Before" is `$R/sb-before` from `$R/site-before/sim`; "after" is this branch. Raid 60 does not move (no row this lane touched applies to the raid preset), so `before.tsv` is unchanged.

| spec, band, faction | metric | before, seed 1 | after, seed 1 | change seed 1 / 50001 |
|---|---|---|---|---|
| priest-shadow 50 Horde | DPS | 238.5 | 239.6 | +0.46% / +0.47% |
| priest-shadow 50 Alliance | DPS | 216.1 | 216.1 | 0.00% |
| priest-shadow bare 60 Alliance | DPS | 359.4 | 362.1 | +0.77% / +0.72% |
| hunter-beast-mastery 40 H / A | DPS | 159.9 / 160.3 | 160.7 / 160.9 | +0.48 / +0.41%; +0.36 / +0.36% |
| hunter-beast-mastery 50 H / A | DPS | 222.2 / 220.4 | 222.9 / 221.1 | +0.32 / +0.24%; +0.31 / +0.33% |
| hunter-beast-mastery bare 60 H / A | DPS | 301.3 / 301.3 | 302.3 / 302.1 | +0.34 / +0.25%; +0.28 / +0.15% |
| shaman-elemental 30 H / A | DPS | 48.1 / 47.9 | 48.8 / 49.1 | +1.37 / +1.35%; +2.57 / +2.59% |
| shaman-elemental 50 H / A | DPS | 155.0 / 152.6 | 155.7 / 153.4 | +0.43 / +0.40%; +0.51 / +0.51% |
| shaman-elemental bare 60 H / A | DPS | 191.2 / 198.0 | 191.9 / 198.8 | +0.41 / +0.41%; +0.39 / +0.39% |
| mage-fire 30 Horde | DPS | 69.9 | 71.0 | +1.69% / +1.70% |
| warlock-destruction 40 Horde | DPS | 147.2 | 147.4 | +0.17% / +0.16% |
| priest-discipline bare 60 Horde | effective HPS | 318.1 | 318.2 | +0.03% / +0.13% |
| priest-holy bare 60 Alliance | effective HPS | 324.9 | 325.1 | +0.06% / 0.00% |

## Files to regenerate

- `sim/request/testdata/ladder/*.golden.md` (TestRotationLadder): priest-shadow (50), hunter-beast-mastery (40, 50), shaman-elemental (30, 50), mage-fire (30 if the ladder race is Horde), warlock-destruction (40 if Horde). The ladder follows its race's faction now.
- addon-data: `data/builds/1.60.1.70753/addon-data.json` and `addon/ForeverSixtyData/Data.lua` (`pipeline addon-data`). The rotation bands at 30 to 50 for Beast Mastery, Elemental, Shadow, Fire and Destruction move, and the Shadow 50, Fire 30 and Destruction 40 lines gain `factionSpellId`.
- Coach starters: no change expected. The coach reads the embedded table at run time. Re-render them only if the starters embed rotation ids per band.
- `sim/rankchoices/rank-choices.json` is synced in the commit (`make rank-choices-sync`); nothing else is generated from it.

## Measure at refresh

- Bare-band BiS picks for the specs whose bare rotation changed: Beast Mastery 40, 50 and bare 60 (both factions); Elemental 30, 50 and bare 60 (both); Shadow 50 Horde and bare 60 Alliance; Fire 30 Horde; Destruction 40 Horde; Discipline bare 60 Horde; Holy priest bare 60 Alliance.

## For other lanes

- **Heal profile:** R5 and R6 key three bare-60 healer rows by faction (Discipline's Binding Heal for Alliance only; Holy priest's Prayer of Healing and Lightwell for Horde only). This lane did not touch the sub-60 healer rows. If you rewrite those rows, keep the new `faction` keys and the notes text.
- **Platform / addon:** `Rotation.lua` now reads an optional `factionSpellId` on a rotation line (S 75cb822f3). Any other surface that reads `data.rotations` directly should go through `Rotation.model`. None does today.
- **Hunter owner:** L5, Arcane Shot's spell-power coefficient against the client's 0.0.
- **Auras:** C1, Skyborne's engine faction.
- **Shaman:** the paced Lightning Bolt line's note says "or the band's rank choice below 60". That is now true at 50 and bare 60 (rank 7 on that line).

## Still open

- L5 and C1 above.
- Lava Burst rank 1 against rank 2 at 50 is now within error (-0.16 / -0.12% Horde, -0.33 / -0.22% Alliance). Rank 2 stays as round 11 set it.

# Sim-accuracy round 22: priest

Lane priest. Site wt/r22-priest from int/r22 219efc56f. Fork lane/r22-priest from int/r22 595bd23d1. Client 1.60.1.70790.

All sims ran through `$R/sim.sh` on the frozen BiS (`$R/bis-frozen`), 10,000 iterations, seeds 1 and 50001. Band levels were judged on the ladder leveling build at each level (`$R/priest/tools/ladder.py`), raid 60 on the guide build (`$R/rt`). Before is this lane's own baseline binary (`$R/priest/sb-base`, int/r22 with the shared 3% crit change), not `$R/before.tsv`.

**Error.** Healing in this profile is deterministic to 0.1 HPS (the two seeds differ by at most 0.1). A heal rank row is taken on 0.2 HPS or more at both seeds, as in round 21. Shadow prints its own error (0.3 DPS at raid 60).

## Seeds

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Heal rank 1 was never tried against the later ranks** (guides r24 sim 1, "Holy Heal ranks at 22-29").<br>- What was wrong: the healer rank method never tried a rank learned below 20 (the APL header's "never one learned below level 20"), a rule that only made sense with Era's sub-20 cut. Forever has no such cut (platform P1: three independent sources), and the engine already gives Heal rank 1 (2054, level 16, MaxLevel 21) its whole coefficient until 33. Rank 1 heals 280 + 0.857 x healing for 155 mana; rank 3 heals 543 + 0.857 x healing for 255, so rank 1 is ahead per mana above 149 healing, and band gear has 144 to 251.<br>- Client: Wowhead Forever tooltips 2054 (260 to 300, 155 mana) and 2055 (364 to 415, 205 mana) match `heal_ranks.go` exactly.<br>- Measured on every race, 22 to 39, both specs (`$R/priest/heal22.txt`, `heal30.txt`, `heal30b.txt`; picks by `tools/pick.py`):<br>  - **Discipline (the tank's Heal):** rank 1 wins at 22 to 27 (+0.6% to +1.3% over rank 2) and 28 to 34 (+0.3% to +1.6% over rank 3, ties for a Human or Gnome at 28 and a Dwarf or Night Elf at 29), and at 35 for a Human (+0.2%) and Undead (+0.8%), ties for the rest. From 36 rank 3 wins again. Round 21's Undead rank 2 at 30 to 33 loses to rank 1 by 0.9% to 1.3%.<br>  - **Holy:** rank 1 wins for every race at 22 to 26 (+1.0% to +2.8%; ties for an Undead at 23 and a Gnome at 24 to 26), 29 (+1.3% to +7.5%) and 33 (+0.7% to +1.6%). Race by race: 27 Dwarf, Undead, Troll; 28 Undead (+0.9% over round 21's rank 2); 30 and 31 the four Alliance races (+3.9% to +10.1%); 32 Troll (+0.7%); 34 Troll (+0.4%). Holy at 30 to 32 swings with talent points (Improved Healing at 29 to 31, Holy Binding at 32 changes how often the pace gate lets the filler through), so the race split there follows the measurement level by level, as the rule asks.<br>  - Rank 1 was confirmed cast (2054 in the mana table) on every adopted row; `sb-after` reproduces the predicted rows (`$R/priest/rows-check.txt`).<br>- **Fix:** Heal 6064 rows rewritten for both specs (below). | `TestRound22PriestHealRankRows` failed (Discipline Gnome 22, Undead 30, Human 35 and the Holy rows), then passed | CONFIRMED (rank rows) | old (round 21 P-H23/D1 measured ranks 2 to 4 only), guides-handoff | S e7093d9b9 |
| P2 | **Holy rows around round 21's** (the brief's "check the rows below and around them"). Ranks 2 and 3 against the row at 34 to 39:<br>- A Troll heals +0.7% at 38 and +0.5% at 39 with rank 2 over ranks 3 and 4.<br>- An Undead heals +0.2% at 39 with rank 2 over round 21's rank 3 (both seeds, the bar).<br>- At 32 a Dwarf heals +0.6% with rank 2 over rank 3.<br>- Round 21's rank 3 at 34 to 37 holds for every race but a Troll at 34 (rank 1); a Human keeps rank 2 at 28.<br>**Fix:** rows rank 2 at 38-39 Troll, 39 Undead, 32 Dwarf. | same files | CONFIRMED (rank rows) | old | S e7093d9b9 |
| P3 | **Guides r24 sim 2: why rank 2 wins at 28 for a Human.** Rank 3 heals more per mana at 28 (guides R24-3). The Human row is now the only rank 2 hold left at 28; it wins by 0.2 HPS against rank 1 and rank 3 (45.7 against 45.5), on the overheal side: rank 3's raw healing is 1.4 HPS higher and its effective healing lower. The note and prose no longer give a reason for the 28 hold. | `heal22.txt` (raw vs effective) | Aligned (reason dropped, row kept) | - | - |
| P4 | **Guides r24 sim 3: Discipline's tank had no filler below 16.** Shield and Renew only on the tank at 10 to 15, while the party already got Lesser Heal rank 3 (2053, trained at 10). Two tank lines now mirror the tank's Heal lines (paced under 80%, and under 50%), scoped below 16. On band-20 gear (round 21's throwaway `-level` build, same binary both sides): 12 Human 34.2 to 40.5, Undead 31.8 to 37.8; 15 Human 35.8 to 42.0, Undead 33.3 to 39.3 effective HPS (+18%), equal at both seeds. | `$R/priest/disc-lh.txt`; `TestRound22DisciplineTankLesserHealBeforeHeal` failed, then passed; `make apl-check` green | CONFIRMED | old, guides-handoff | S e7093d9b9, E afea66e0c |
| P5 | **Shadow +4.0% without Vampiric Embrace upkeep (MythicSim lead).** Verified: the raid list without the Vampiric Embrace line scores 756.1 / 756.4 (Human) and 745.2 / 745.7 (Undead) against 727.4 / 727.5 and 717.5 / 717.2 with it, **+3.9%**. The cost is the global: six casts in a 180-second fight take 9 seconds of casting, about 5% of it, and Mind Flay gets the time back.<br>**Decision: keep it.** Vampiric Embrace is cast for its raid value: 20% of every Shadow spell heals the whole group, and in Forever it is cheap (30-second cooldown, 40 mana, can't be resisted). The Forever guides keep it up: Warcraft Tavern ("you should just keep it up at all times"), wow.gg ("invaluable during high raid damage phases") and Icy Veins' Forever Shadow guide. The guide already says what it costs and when to skip it (a fight where the group takes nothing). | `$R/priest/ve-raid.txt`; sources below | Aligned (rotation choice kept) | - | - |
| P6 | **The platform lane's finding (no sub-20 penalty in Forever for heals): did any priest rank row rely on it?** No row was measured with a sub-20 cut: the engine's `SpellLevelPenalty` has had only the level-gap rule since round 11. The cut lived only in the healer method's exclusion of sub-20 ranks, which is P1. The Renew rank 1 rows (Holy 14 to 32, Discipline 14 to 29) were measured with rank 1 in the race and stand. Shadow's rows (Shadow Word: Pain, Mind Blast, Devouring Plague, Shadow Word: Death) compare ranks learned from 18 up, measured on the same engine. | engine `sim/core/spell.go` `SpellLevelPenalty`; rank-choices.json | Aligned | - | - |

### The Heal rows now (authored 6064)

| spec | levels | rank | races |
|---|---|---|---|
| Holy | 22-26 | 1 | all |
| Holy | 27 | 1 | Dwarf, Undead, Troll (others rank 2) |
| Holy | 28 | 2 | Human (Undead rank 1 row; others rank 3) |
| Holy | 28 | 1 | Undead |
| Holy | 29 | 1 | all |
| Holy | 30-31 | 1 | Human, Dwarf, Gnome, Night Elf (Horde rank 3) |
| Holy | 32 | 2 / 1 | Dwarf / Troll (others rank 3) |
| Holy | 33 | 1 | all |
| Holy | 34 | 1 | Troll |
| Holy | 34-37 | 3 | all but a Troll at 34 |
| Holy | 38 | 3 / 2 | Undead / Troll (Alliance rank 4) |
| Holy | 39 | 2 | Undead, Troll (Alliance rank 4) |
| Discipline | 22-35 | 1 | all |
| Discipline | 36-39 | 3 | all |

## Discipline

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | P1 and P4 above. | | CONFIRMED (2) | | |
| D2 | Weakened Soul 15 s (SpellDuration 8), Power Word: Shield 10901 cooldown 4 s (SpellCooldowns 54449) against `power_word_shield.go`. Renewed Hope (10% and 5 s at 5/5; engine 2% and 1 s a rank), Divine Aegis 15% capped at 20 x level, Improved Healing 5% a rank on Lesser Heal, Heal, Greater Heal, Penance and Prayer of Mending, Divine Fury 0.1 s a rank on Smite, Holy Fire, Heal and Greater Heal: all match the 70790 talent text. | client tables, talents/priest.json, `healing_talents.go` | Aligned | - | - |

## Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | P1 and P2 above. | | CONFIRMED (2) | | |
| H2 | Litany of Light compares heal families (SpellCode), not ranks, so a Heal rank 1 after a Heal rank 3 gets nothing, as "a different spell" reads. Holy Binding is Flash Heal on others only, 50%. Spiritual Guidance 25% of Spirit as healing at 5/5; Spiritual Healing 10%. Twilight Focus (pushback) is an unmodeled keep. | `healing_talents.go:216-275`; talent text | Aligned | - | - |

## Shadow

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Vampiric Embrace could miss.** Blizzard's 16 October notes: "Changed Vampiric Embrace so that it can't be resisted." The 70790 client carries that as a new attribute on 15286: SpellMisc Attributes_3 is 67371136 where Era's is 67108992, the difference being 0x40000, the ignore-hit-result flag. The engine rolled the spell hit table, so about 9% of casts missed on a level-63 boss (raid run: 6.03 casts, 5.48 landed, 9.1% miss). A missed cast still started the 30-second cooldown, so the group lost a sixth of its Vampiric Embrace healing. Round 21 X1 listed this note as already in the engine; it was not.<br>**Fix:** `OutcomeAlwaysHit`. Now 6.03 casts, 6.03 landed. No damage change (raid 60 within 0.04%): the cast count and the global are the same. | client SpellMisc 70790 vs Era; Blizzard 16 Oct; `TestVampiricEmbraceCannotMiss` failed on cast 6, then passed | CONFIRMED | old (round 21 X1 read the note and missed it) | E c1360df15 |
| S2 | Every other priest spell whose 70790 SpellMisc gains the ignore-hit flag over Era: only Shadow Weaving 15258 (the self buff, which always applies already). | scan of SpellMisc 70790 vs Era over SpellClassSet 6 | Aligned | - | - |
| S3 | Inner Focus: 3-minute cooldown (SpellCooldowns 54951) started when the buff is spent, as in the engine; 1.34 casts a 180-second fight comes from longer fights. Mind Blast 8 s, Shadow Word: Death 18 s, Devouring Plague 24 s, Vampiric Embrace 30 s cooldowns match SpellCooldowns 54580, 100966, 55344, 55127. Shadow Focus 1% a rank (5% at 5/5) gives 9.5% miss against 63 with the BiS hit, consistent across Mind Blast, Shadow Word: Death and Mind Flay (per cast). | client SpellCooldowns 70790; raid breakdown | Aligned | - | - |
| S4 | P5 above (Vampiric Embrace upkeep kept). | | Aligned | | |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Blizzard's Forever notes, every priest line (24 Sep to 4 Nov, round 21's raw copy `r21/priest/notes/rawall.txt`), re-read against the engine. One gap: the 16 October "can't be resisted" line (S1). Every other line is as round 21 X1 lists. The platform lane found no post after 8 October and the 3 November edit already handled. | Blizzard thread 2360696 | S1 only | - | - |
| X2 | `TestHealerFiresASpellPowerTrinket` (round 21 X3, item 18820 missing) passes on this head: `go test -tags with_db ./sim/priest/...` is green. The priest conformance golden still differs on arm64 in the fourth decimal (round 21 X2); regenerating it before and after S1 gives identical files, so S1 moves no conformance cell. | fork tests | Aligned | - | - |
| X3 | Other classes' spells with the new ignore-hit flag (for those lanes): mage Winter's Chill 12579 and Fire Vulnerability 22959, warlock Shadow Vulnerability 17794. The engine applies these as procs with no hit roll (`sim/mage/talents.go:576`), so nothing to fix that I could see. | SpellMisc scan | Aligned (for mage, warlock) | - | - |

## Before/after

Before = `$R/priest/sb-base` (int/r22 595bd23d1 engine, int/r22 site); after = lane engine afea66e0c + site e7093d9b9 (`$R/priest/before.tsv`, `after.tsv`, ALLRT, frozen BiS).

**Raid 60:** Discipline 530.4 / 526.0 and Holy 771.2 / 788.0 (Alliance / Horde, seed 1) are unchanged to the decimal at both seeds. Shadow 724.6 / 717.7 to 724.3 / 717.5 (seed 1) and 724.8 / 717.1 to 724.8 / 717.2 (seed 50001), within error.

**Bands (rows' levels):** Discipline +0.2% to +1.6% at 22 to 35; Holy +0.4% to +10.1% where a row was added (largest at 30 and 31 for the Alliance races, +5% to +10%); Discipline +18% at 10 to 15 (tank Lesser Heal).

**Other classes:** none. E c1360df15 touches `sim/priest` only; E afea66e0c is the Discipline list copy; the site rows are priest-only.

## Files to regenerate

- `data/builds/1.60.1.70790/addon-data.json`, then `addon/ForeverSixty/Data.lua`: the Holy and Discipline Heal rows (race_spell_ids), the Discipline tank Lesser Heal lines, and the Holy and Discipline Heal notes.
- TestRotationLadder goldens for priest-holy and priest-discipline (bands 10/20/30), if they print these lines or rows.
- Coach starters `data/curated/coach/priest-{holy,discipline}.{md,json}`, if they carry the Heal notes or the band-10 lines.
- Copied by hand, not regenerated: `sim/rankchoices/rank-choices.json`, `sim/request/apl/priest-{holy,discipline}.apl.json`, fork `ui/healing_priest/apls/forever_discipline.apl.json` (rotation only; apl-check ignores notes).

## Measure at refresh

- Every Heal row on re-ranked band-20 and band-30 gear. Rank 1's lead is per mana, so more healing on re-ranked gear widens it (above 149 healing it out-heals rank 3 per mana), but the Holy 30 to 32 race split rides on the pace gate and may move.
- Carried: the Shadow Word: Death order (round 21), Discipline Gnome against Human, the Devouring Plague 28-29 row.

## For other lanes

- **Paladin and every healer lane:** the same "never one learned below level 20" exclusion is in `data/curated/apl/paladin-holy.json`'s header. If any leveling rank row was chosen without trying a rank learned below 20 (Holy Light 1 to 3, Healing Touch, Healing Wave, Rejuvenation, Regrowth), it is the same gap as P1: Forever keeps those ranks whole until MaxLevel + 12.
- **Mage, warlock:** X3 (no action found).
- **Integrator:** holy.md, discipline.md and priest-holy.json were taken from origin/main (guide round 24) before this lane's edits, so the merge of origin/main into int/r22 meets identical text there.

## For the guide lane

- Written in this lane: holy.md Levels 20 to 60 (Heal rank 1 and the race rows), discipline.md Levels 1 to 19 (tank Lesser Heal) and 20 to 60 (Heal rank 1 through 35, rank 3 from 36), the Holy Heal note, the Discipline tank Heal note and the new tank Lesser Heal note. **wow-player: SHIP** on the rendered pages (captures `$R/priest/shots/`, no MUSTs). Its SHOULDs applied: the Holy paragraph split into Heal, Renew and Greater Heal paragraphs and the Human-at-28 clause reworded; the Holy note's first part now says some races cast rank 2 or 3; the Discipline tank note puts the ranks before the colon; "When you heal a group from 10"; "From 40, Greater Heal". Re-captured in `$R/priest/shots2/`. Not applied: "every cast costs mana regeneration" wording (older text) and the build-number footer (the standard page stamp).
- The card was not captured: it hydrates from `addon-data.json`, which waits on the regen.
- Still out of scope (guides r24): `data/curated/coach/priest-holy.md` says Shadowfiend and "a mid rank of Heal".

## Still open

- Round 20's P2 (heals scaling with the target's level), P3 (Eureka! on a live Gnome Shadow log), P4 (Lightwell Renew's crit), S2 (Dark Inspiration's Spirit base).
- **Logs:** last checked 09:00 UTC 5 November. api.foreversixty.gg `/v1/reports/recent` lists only the 20 October report; foreverlogs.gg `reports/public` is empty. P3 and P4 need a live log from a level-40+ character.

## Sources

- Blizzard, WoW Forever Beta Development Notes: https://us.forums.blizzard.com/en/wow/t/2360696 (16 October priest list: "Changed Vampiric Embrace so that it can't be resisted.")
- Client 70790: `scratchpad/b70790/70790/{SpellMisc,SpellCooldowns,SpellDuration}.csv`; Era: `scratchpad/wago/SpellMisc_era.csv`
- Wowhead Forever tooltips 2054, 2055: https://nether.wowhead.com/forever/tooltip/spell/2054 (read 5 Nov)
- Vampiric Embrace upkeep: https://www.warcrafttavern.com/forever/guides/shadow-priest/ ; https://wow.gg/guides/priest-shadow-forever-overview ; https://www.icy-veins.com/wow-forever/shadow-priest-ranged-dps-pve-guide (read 5 Nov)
- Platform round 22 P1 (no sub-20 cut): `design/reviews/2026-11-05-sim-audit-round22/platform.md`

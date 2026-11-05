# Sim-accuracy round 21: priest

Lane priest. Site wt/r21-priest from int/r21 da6152025 (guides rounds 19 and 20 merged, plus the platform, warlock, mage, hunter and druid lanes). Fork lane/r21-priest from int/r21. Client 1.60.1.70790.

All sims ran through `$R/sim.sh` on the frozen BiS (`$R/bis-frozen`) at 10,000 iterations, seeds 1 and 50001. Band levels were judged on the ladder leveling build at each level (`$R/priest/tools/ladder.py`, which decodes `addon-builds.json`; it gives the same strings as round 20's helper). Raid 60 used the guide build (`$R/rt`). Exploration ran at 3,000 iterations.

**Error.** In this heal profile, healing comes out deterministic to 0.1 HPS: seeds 1 and 50001 differ by at most 0.1. A heal rank row is taken only on a gain of 0.2 HPS or more at both seeds. Shadow prints its own error, which is 0.0 to 0.3 DPS at these levels.

## Seeds

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P-H23 | **Holy Heal rank 2 against rank 3 at 28 to 33** (guide r19). This was measured on all six races at 28 and 29, and on Human and Undead from 30 to 33.<br>- At 28, rank 2 wins for a Human (45.7 against 45.5, +0.4%) and an Undead (43.4 against 42.9, +1.2%). A Gnome ties. A Dwarf, Night Elf or Troll loses 0.6% to 1.5%.<br>- At 29, only an Undead keeps rank 2 (+1.6% / +1.4%).<br>- From 30 (band-30 gear), rank 3 wins for Human and Undead by 1.0% to 1.5%.<br>Rank 2 was confirmed cast (2055 in the mana table).<br>**Fix:** rank rows for Heal (authored 6064): rank 2 at 28 for Human and Undead, and at 29 for Undead. | `$R/priest/heal23.txt`, `heal23b.txt`; `TestRound21PriestHealRankRows` failed, then passed | CONFIRMED (rank row) | old, guides-handoff | S 8dcb3404d |
| P-SWD | **Shadow Word: Death ahead of Mind Blast at 32 to 51** (guide r20). Swapping the two on-cooldown lines gives mixed results.<br>- Human, the page's race: ties at 32, 35, 38 and 51. It loses 0.1% to 0.3% at 41, 47 and 50. It gains 1.5% at 44.<br>- Undead: gains 0.2% to 0.6% at most levels and loses 0.2% at 44.<br>- 53 to 59: gains 0.1% to 0.5%.<br>- Bare 60: +0.3% Human, +0.15% Undead.<br>- Raid 60: -0.1% for both factions, beyond error.<br>Every gain comes from fights that run out of mana, where Death's damage per mana wins. Death is not the better hit per cast time. A level- and race-split order would only fit the noise.<br>**Decision:** keep Mind Blast first. No change. | `$R/priest/swd-1.txt`, `swd-2.txt`, `swd-raid.txt`, `swd-explore.txt` | Aligned (no change) | - | - |

## Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Heal rank 3 against rank 4 at 34 to 39.** No round compared these ranks; round 15 H1 only listed what was trainable. Rank 4 arrives at 34 and costs 305 mana against rank 3's 255, and the fight is mana-bound on band-30 gear.<br>- Rank 3 wins for every race at 34 (+0.4% to +0.9%) and 35.<br>- At 36 and 37 it wins or ties for every race (Undead +0.8%, Troll +0.3%; Human -0.1% at 37 is within the 0.1 resolution).<br>- At 38, Undead (+0.5%) and Troll (+0.2%) gain; the Alliance races tie or lose up to 0.2%.<br>- At 39, only Undead gains (+0.4% / +0.5%); the Alliance races lose 0.2% to 0.3%.<br>**Rows:** rank 3 at 34-37 for every race, 38 for Undead and Troll, 39 for Undead. | `$R/priest/holyheal34.txt`; `TestRound21PriestHealRankRows` | CONFIRMED (rank row) | old | S 8dcb3404d |
| H2 | **Holy had no heal filler at 10 to 15.** The list's first heal after Renew was Heal, which trains at 16, so the band-10 card showed Renew alone. The engine also cast only Renew there, although Lesser Heal rank 3 is trained at 10 (trainables). A Holy priest in a dungeon at 10 to 15 casts Lesser Heal.<br>**Fix:** ten Lesser Heal rank 3 lines that mirror the Heal lines' targets and mana pace, scoped below level 16 as Discipline's are scoped below 40. On band-20 gear (a throwaway build that allows -level 12 and 15), effective HPS goes 14.7 to 43.2 at 12 and 16.7 to 46.2 at 15, the same at both seeds. Level 20 is unchanged (41.9). | `$R/priest/holy-lh.txt`; `TestRound21HolyLesserHealBeforeHeal` and `TestRound21HolyLesserHealStopsAtSixteen` failed, then passed; `make apl-check` green | CONFIRMED | old | S 2e3a7d7a7, E 6677b9697 |
| H3 | The raid-60 list, Prayer of Mending, Renew, Greater Heal, Prayer of Healing and Lightwell were all checked against Wowhead's Forever tooltips (this round). The amounts are centred rolls at the rank's MaxLevel, as `clientdamage.Range` computes them: Greater Heal 4 gives 1722-1917, Heal 3 gives 510-576, Prayer of Mending 3 gives 413 + 0.429, Prayer of Healing 4 gives 581-612, Renew 7 gives 370 over 15 s, Lightwell gives 1600. The rows match 70790 SpellEffect exactly (`heal_ranks.go`). | `$R/priest/wh/*.json`; scratchpad b70790/70790/SpellEffect.csv | Aligned | - | - |

## Discipline

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | **The tank's Heal at 28 to 39.** Rank 3 wins at 28 and 29 for every race, so no row is needed there.<br>- At 30 to 33, rank 2 gains +0.2% to +0.4% for an Undead only; every other race ties within 0.1.<br>- At 34 to 38, rank 3 beats rank 4 for every race (+0.3% to +0.7%), and at 39 it wins or ties.<br>**Rows:** rank 2 at 30-33 for Undead; rank 3 at 34-39 for every race. | `$R/priest/discheal.txt`, `discheal34.txt` | CONFIRMED (rank row) | old | S 8dcb3404d |
| D2 | **Round 20's changes re-read.** Renewed Hope once per Penance: `healTarget` cuts Weakened Soul and `healPulse` keeps the crit bonus only, as tracker #288 says. The Lesser Heal gate (`currentLevel < 40`) is present on all eight lines, and the rendered rows end at 38. Penance ranks 1 to 4 match the Forever tooltips: 134 / 209 / 279 / 425 at their MaxLevel, costing 150 / 220 / 270 / 385. Power Word: Shield 10901 gives 950 at MaxLevel 65 (928 at 60). | penance.go, heals.go:94-112; tooltips 402174, 1240720, 1240721, 1316995, 10901 | Aligned | - | - |

## Shadow

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Twin Disciplines did not reach Starshards.** The talent's periodic modifier (1225132 eff1, aura 108 misc 22, mask 0xa08060/0/0/0x10) names Starshards' class mask bit (0x200000 on 10797 and 19296 to 19305). Blizzard's 16 October notes say Twin Disciplines "now correctly benefits Shadow Word: Death, Mind Flay, Starshards and Penance". The engine's Starshards carried no class mask at all.<br>**Fix:** a Starshards mask bit, added to Twin Disciplines.<br>No published number moves, because no rotation casts Starshards. A Night Elf Shadow priest who weaves Starshards in ahead of Mind Flay loses 2.2% at raid 60 (745.2 to 729.1), and the +5% does not close that gap. | client SpellEffect/SpellClassOptions 70790; Blizzard 16 Oct; `TestTwinDisciplinesRaisesStarshards` failed (undefined mask), then passed | CONFIRMED | old (no round read eff1 for Starshards) | E 0e210f6fd |
| S2 | **Live tooltips against the engine** (Wowhead Forever, fetched 04:10 UTC):<br>- Mind Blast 9: 485-511, 350 mana, 8 s.<br>- Shadow Word: Death 4: 447-474, 340 mana, 18 s, 30% backlash.<br>- Mind Flay 6: 390 over 3 s, 205 mana.<br>- Shadow Word: Pain 8: 762 over 18 s.<br>- Devouring Plague 6: 848 over 24 s, 985 mana, 24 s cooldown; ranks 1 and 2 at 128 and 232.<br>- Vampiric Embrace: 20%, 30 s, 40 mana.<br>- Shadowform: 40% base mana, +10%, -50% cost, +100% crit bonus.<br>- Inner Focus: non-periodic only.<br>All match round 20's S1 values. | `$R/priest/wh/` | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Every priest line in Blizzard's Forever notes from 24 September to 4 November, read against the engine:**<br>- *Already in the engine:* Renew can crit (3 Nov). Power Word: Shield overwrite (3 and 4 Nov; round 20 D3). Shadow Word: Death backlash is 30% and none on a miss (4 Nov). Vampiric Embrace healing cannot crit. Holy Binding replaces Binding Heal. Inner Focus does not touch periodic effects (2 Oct). Shadow Word: Death uses Early Demise only below the threshold. Shadow Weaving always applies. Devouring Plague crits. The 16 Oct list: Devouring Plague and Shadow Word: Death cooldowns 24 s and 18 s, Vampiric Embrace 30 s and unresistable, Improved Mind Flay compounding, Dark Inspiration, Mental Strength. The 8 Oct Penance and Heal/Lesser Heal base points arrive through the client rows.<br>- *Contingency Plan can crit (3 Nov):* its heal over time is a non-channelled HoT, so it resolves through `Dot.OutcomeTick`, the same path that lets Renew's ticks crit.<br>- *Racials (4 Nov):* Human Spirit 10%, Dwarf Artisan's Precision 1% hit, Troll Berserking 15 s, Undead Touch of the Grave 15 s internal cooldown. All four are in sim/core already.<br>- *Changes with no sim effect:* Starshards pushback, the Devouring Contagion critter rule, Cure Disease in Shadowform, Dark Sacrifice and crowd control.<br>**Unlike the druid lane, no priest note carries a number the cached 70790 rows lack.** The live tooltips (S2, H3, D2) agree with the rows. | Blizzard thread 2360696 (raw fetched 04:05 UTC, posts 1 to 10, to 4 Nov 19:35 UTC) | Aligned (S1 is the one gap) | - | - |
| X2 | **Fork conformance goldens "stale at d5c708834".** Regenerating `sim/core/testdata/conformance/priest.golden.md` here moves only four penalty-share cells in the fourth decimal: Heal 1 0.5570 to 0.5571, Heal 4 0.4714 to 0.4713, and Mind Blast 5 0.2360 to 0.2359 on both presets. The golden's last commit (abb4e6b8e) says the goldens are "written on amd64, as CI checks them", and this Mac is arm64. **This hides no priest change.** The golden was reverted, not committed. | `FOREVER_UPDATE_GOLDEN=1 go test ./sim/conformance -run TestConformance$/priest` | Aligned (platform rounding) | - | - |
| X3 | **`TestHealerFiresASpellPowerTrinket` (sim/priest/healing) fails on the base engine too**, with "No item with id: 18820" from the local item db. It predates this lane, and I left it alone. | base-engine, GOWORK=off | Note (platform) | - | - |

## Before/after

Before is `$R/before.tsv`. After is this branch on lane engine 0e210f6fd (`$R/priest/after.tsv`, ALLRT, frozen BiS).

**Raid 60:** every priest spec is unchanged to the decimal, at both factions and both seeds. Discipline is 539.9 / 536.2, Holy 780.6 / 797.8, and Shadow 742.6 / 735.5 (seed 1, Alliance / Horde).

**Bands:** the row levels only, all relative to the top rank:

| level | spec | change |
|---|---|---|
| 28 | Holy | Human +0.4%, Undead +1.2% |
| 29 | Holy | Undead +1.5% |
| 30-33 | Discipline | Undead +0.2% to +0.4% |
| 34-37 | Holy | +0.4% to +0.9% (every race) |
| 34-39 | Discipline | +0.3% to +0.7% |
| 10-15 | Holy | from Renew alone to a real filler (about three times the healing on band-20 gear) |

**Other classes:** none. E 0e210f6fd touches `sim/priest` only, and E 6677b9697 is the Holy list copy.

## Files to regenerate

- `data/builds/1.60.1.70790/addon-data.json`, then `addon/ForeverSixty/Data.lua`. This picks up:
  - the Holy band-10 Lesser Heal row and its note;
  - the Holy Heal note;
  - the Heal rank rows for Holy (28, 29, 34-39) and Discipline (30-39).
- TestRotationLadder goldens for priest-holy (bands 10, 20, 30) and priest-discipline (band 30), if they print those lines or rows.
- Coach starters `data/curated/coach/priest-holy.{md,json}`, if they carry the Heal note or the band-10 lines.
- Copied by hand, not regenerated: `sim/rankchoices/rank-choices.json` and `sim/request/apl/priest-holy.apl.json`.

## Measure at refresh

- Every Heal row on re-ranked band-20 and band-30 gear. The margins are 0.2% to 1.6%, and the gains come from mana, so better gear moves them.
- The Shadow Word: Death order, if re-ranked gear changes how mana-bound the 40s are.
- Carried from round 20: Discipline Gnome against Human, and the Devouring Plague 28-29 row.

## For other lanes

- **Platform:**
  - X2: the conformance goldens differ on arm64 only in the fourth decimal. Regenerate on amd64 or leave them.
  - X3: the item-18820 db failure in `sim/priest/healing`.
  - `spec-breakdown -level` cannot go below a band floor of 20 for healers, because there is no band-10 BiS entry. The Holy 10-15 check needed a throwaway build.
- **Integrator:** the new test file `sim/request/rank_choices_priest_r21_test.go` is separate from `rank_choices_remeasured_test.go` (which the hunter lane also appends to), so the two will not conflict.

## For the guide lane

- Written in this lane with a wow-player SHIP on the second pass (captures in `$R/priest/shots/`):
  - holy.md levels 1 to 19 (Lesser Heal rank 3 as the filler before 16);
  - holy.md levels 20 to 60 (Heal rank 2 at 28, rank 3 at 29, rank 3 from 34 to 37);
  - discipline.md leveling (Heal rank 3 from 34 until Greater Heal at 40);
  - the Holy Heal and Lesser Heal card notes.
- Not applied: the reviewer's optional SHOULD to add "then rank 3 at 29" to the Holy Heal card note. An Undead keeps rank 2 at 29, and the card rows follow race.
- The card itself was not captured. It hydrates from `addon-data.json`, which waits on the regen.

## Still open

- Round 20's items are still open: P2 (heals scaling with the target's level), P3 (Eureka! on a live Gnome Shadow log), P4 (Lightwell Renew's crit) and S2 (Dark Inspiration's Spirit base).
- **Logs:** last checked at 04:27 UTC on 5 November. There are none to use.
  - api.foreversixty.gg `/v1/reports/recent` lists only the 20 October report.
  - foreverlogs.gg `reports/public` is empty.
  - The platform harvest (`$R/platform/logs.md`) shows no Warcraft Logs zone for Forever.
  - P3 and P4 need a live log from a level-40+ character.

## Sources

- Blizzard, WoW Forever Beta Development Notes, posts 1 to 10 (24 Sep to 4 Nov 2026): https://us.forums.blizzard.com/en/wow/t/2360696 (raw copy at `$R/priest/notes/rawall.txt`)
- Wowhead Forever tooltips: https://nether.wowhead.com/forever/tooltip/spell/<id> for 10947, 1309636, 18807, 10894, 19280, 2944, 19276, 402174, 1240720, 1240721, 1316995, 10965, 2060, 10917, 10927, 401859, 1240826, 1240827, 2055, 6063, 6064, 10901, 15286, 15473, 14751, 10961, 27871
- Client 70790: scratchpad `b70790/70790/` SpellEffect, SpellClassOptions, SpellName; `data/builds/1.60.1.70790/{talents,trainables,addon-builds,addon-data}.json`
- forever-bugs #288 (Renewed Hope once per Penance)

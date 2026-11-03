# Sim-accuracy round 12: integration (2026-11-03)

Round 12 on client 1.60.1.70753, from site c6bdc855c (the combined landing of round 11 and the launch-day guide lanes) and engine pin 2606bc05e. Sixteen lane reports sit beside this file: regmelee.md and regcast.md (the regression lanes over round 11 and the landing), warlock.md, casters.md, physical.md, hybrids.md, healers.md, healprofile.md, auras.md, itemscrit.md, holypower.md (split off itemscrit's Elixir of Holy Power finding), ranks.md, trainables.md, guides.md and platform.md.

**Not landed.** Following the owner's batch-landing practice, nothing is pushed or pinned. Everything is on site `int/r12` and fork `int/r12`; one integrator lands it with one pin, one regen and one gate run (see "For the integrator").

## Owner rulings this round, and where they landed

| ruling | outcome | where |
|---|---|---|
| Rank-choice rows are adopted beyond error; the 1% bar is for builds | The six sub-1% rows round 11 left still beat error at both seeds and both factions, so all stay. Rows rejected earlier only for missing 1% are adopted: BM Arcane Shot rank 3 at 40 and rank 1 at 50, Elemental Flame Shock rank 1 at 30. | ranks.md |
| A faction key only where factions differ beyond error | Built through the sim, spec-breakdown, rotation-search, the ladder, the coach and the addon export (`factionSpellId`). Eight rows are keyed. Each one turns out to depend on the published race, not the faction (Dark Sacrifice, Gnome, Troll); each row's note names its race. A race key is recommended, not built. | ranks.md R12 |
| Heal profile tank-versus-group split from data, with a stated rule | Below 60, the tank's share is the median share of group damage taken by the tank over logged foreverlogs.gg beta trash fights in the band's dungeons. Bands with fewer than 30 logged fights pool every beta dungeon. Shares: 20 0.679, 30 0.687, 40 and 50 0.677 (pooled); 60 stays 0.659 (no readable level-60 logs). Line of sight stays out of the pull rule: logged pulls are already bigger than the rule's. | healprofile.md |

## Confirmed and fixed, by lane, with origin

Origin key:
- **landing:** caused by the combined landing c6bdc855c's guide lanes or main's launch-day guide rewrites.
- **R9 / R10 / R11:** caused by that round's fix.
- **old:** predates round 9; no round had found it.
- **new-data:** 70753 or a hotfix push changed under us.

Bugs two lanes found are counted once:
- Juju Might is counted under itemscrit (regmelee L1).
- Enrage's Dire Bear armor is counted under regmelee (hybrids G1).
- Item haste is counted under itemscrit (found by casters X2).
- Hammer of Wrath, hunter and wand travel come from the warlock lane's relay and are counted where fixed.

| lane | confirmed | landing | R11 | R10 | R9 | old | new-data | bugs |
|---|---|---|---|---|---|---|---|---|
| guides | 11 | 11 | | | | | | Unstable Affliction / Haunt line, rogue Races panel, Kidney Shot costs 25 energy, Eureka!, Scatter Shot is a talent, three build point counts, Ruin, shaman crit pools, Revive's cost |
| regmelee | 7 | 3 | | | | 4 | | pre-pull Bloodrage let warriors recast the raid's shout; Arms' Charge after Bloodrage (the client's in-combat effect); shout copy; Arms' execute swap before Recklessness is learned (bare 30 +2.4 to +2.6, bare 40 +3.5 to +3.6); Improved Charge rage with no Charge; Dire Bear Enrage 27% to 16% armor; Fury's dead Improved Charge and Deflection points |
| ranks | 6 | 4 | 1 | | | 1 | | Elemental's new Lightning Bolt line had no row; Shadow Mind Blast and Devouring Plague by race; Fire Blast rank 3 at 30 (Orc); BM Arcane Shot rank 1 at bare 60; Destruction Corruption rank 3 at 40 (old); Discipline's bare-60 Binding Heal row (R11) |
| auras | 5 | | | | | 4 | 1 | Windfury Totem's bonus ends with the totem; Blood Fury raises healing; Orc Command: guardians out, Firebolt/Lash +5% (new-data), additive with Unholy Power and Unleashed Fury |
| healers | 5 | | | | | 5 | | Vampiric Embrace 20% (not 5%), Shadow ticks only, own damage only; Mental Agility on Power Infusion; Idol of Swiftness registered; Reflection 17/33% |
| healprofile | 5 | 2 | | 1 | | 2 | | the fixed 60 split at every band (R10); Discipline's 20/30 Shield rows and its 40 list (landing); Holy paladin 40 Flash of Light pair; Restoration shaman band 20 takes Water Shield |
| hybrids | 5 | 1 | | | | 2 | 2 | cat Poised (new-data); bear Primal Bite rank 3 at 60, rank 4 removed by push 113283 (new-data); Protection paladin band 30 without Holy Shield; Hammer of Wrath travel; Shred/Claw note |
| trainables | 5 | | | | | 5 | | a filter dropping live baseline passives; 124 talents untagged; 16 talents missing; duplicate ranks; no-learn-row copies taken as ranks (Aimed Shot's 3 s cast, Arcane Blast, Freezing Trap, Thunder Clap, seven costs) |
| physical | 4 | | | | | 3 | 1 | rogue Poised at 18 (new-data); Thunder Clap raid debuff at Era's 10/15%; hunter and wand missile speeds; Arcane Shot, Serpent Sting and Volley take no spell power (BM raid -1.5%) |
| casters | 4 | | | | 1 | 3 | | Evocation's lost tick; a cast ending on the global completed a global late (core); Inner Fire skipped under any slowing; dps-per-point measured band 60 on raid ranks with bare gear (R9). Together these were the negative haste readings. |
| itemscrit | 3 | 1 | | | | 2 | | Juju Might back to every hunter (landing 252455d15: BM +1.0, MM +1.1); Lesser Mana Potion, Recombobulator and healthstone amounts; item haste counted twice (custom gear only) |
| platform | 3 | | | 1 | | 2 | | leveling talent strips in point-spend order; weights headline in the SimC form; the coach flagged ranks between the site's pick and the top (R10) |
| holypower | 2 | | | | | 2 | | Elixir of Holy Power (21546) modelled and in the paladin raid presets (Ret +1.6, Prot +2.2 threat); the old Greater Firepower id inert |
| warlock | 1 | | | | | 1 | | Shadow Bolt and Incinerate travel at 20 yards a second |
| regcast | 1 | | | | | 1 | | the healing priest's Inner Fire took a global during the pre-pull Lightwell |
| coordinator | 2 | 1 | | | | | 1 | guide-cooldowns.json left out Reincarnation and Inferno (found by regcast); the paladin class landing listed Voice of Truth and Instrument of Law as talents (found by trainables and the wow-player) |
| **total** | **69** | **23** | **1** | **2** | **1** | **37** | **5** | |

Not counted: the flaky sim-island test (fixed), the dps_rogue fixture database (a test fix), comment and note fixes (Berserking's 15 s, Trueshot Aura, Endurance, Heart of the Lion's 276), and the wow-player's wording fixes.

**Is the count falling? No: 55 in round 11, 69 in round 12.** The rise has three causes:
- **Two new lanes found 16.** Guides (11) and trainables (5) audited ground no earlier round covered.
- **The landing caused 23.** Eleven were guide prose and twelve were rotations, builds and rank rows.
- **That leaves 46 sim bugs** outside guides, trainables and the landing. Of those, 37 are old, 5 are new-data and 4 come from rounds 9 to 11.

Earlier rounds' own fixes caused only 4 this round, down from 13 in round 11.

## Merge notes

| # | conflict | resolution |
|---|---|---|
| M1 | Juju Might fixed by itemscrit and regmelee; two identical tests | regmelee's preset reason text kept; one test kept (S 27529aa3d) |
| M2 | Enrage's Dire Bear armor fixed by regmelee and hybrids | hybrids' form-tier version kept, regmelee's duplicate test dropped (E 9e7421f2f) |
| M3 | the healing priest's pre-pull Inner Fire fixed by regcast and casters | regcast's buff-before-the-pull kept (no global, no cast slot); casters' castWhenFree dropped as dead code; both lanes' tests pass (E 1214cded7) |
| M4 | TestP1TankWarrior and TestEnhancement goldens from three lanes | regenerated once on the merged engine (E 690028efd) |
| M5 | rank-choices.json notes (ranks and healprofile) | both sentences kept |
| M6 | test_addonrotation.py (ranks and healprofile) | both tests kept |
| M7 | poised.go in physical and hybrids | byte-identical |

Coordinator commits on int/r12:
- the cat leveling build line and raid-keeps record (S 454ae42b9);
- the custom panel's Holy Power elixir group and inert elixirs (S bc08b8540);
- the wow-player copy fixes (S 9efecab6f);
- the guide cooldowns (S 157e893b0);
- the fork goldens (E 690028efd).

## Before/after

Before: pin 2606bc05e with c6bdc855c. After: fork int/r12 with site int/r12. Raid preset, band 60, 10,000 iterations, both factions; healers in effective HPS on the heal profile, tanks in threat per second. The bear and cat ran with their APL copies synced to the curated files (the integrator's apl-sync); without it the bear loses Primal Bite.

| Spec | Faction | seed 1 | % (seed 1 / 50001) |
|---|---|---|---|
| druid-balance | alliance | 566.9 to 566.9 | +0.00 / +0.00 |
| druid-balance | horde | 572.2 to 572.2 | +0.00 / +0.00 |
| druid-feral | alliance | 665.0 to 665.0 | +0.00 / -0.02 |
| druid-feral | horde | 660.0 to 660.0 | +0.00 / +0.00 |
| druid-feral-bear | alliance | 1020.3 to 1020.1 | -0.02 / -0.02 |
| druid-feral-bear | horde | 1011.5 to 1010.7 | -0.08 / -0.02 |
| druid-restoration | alliance | 834.6 to 834.6 | +0.00 / +0.00 |
| druid-restoration | horde | 833.7 to 833.7 | +0.00 / +0.00 |
| hunter-beast-mastery | alliance | 623.0 to 619.4 | -0.58 / -0.47 |
| hunter-beast-mastery | horde | 626.9 to 623.5 | -0.54 / -0.54 |
| hunter-marksmanship | alliance | 610.7 to 617.3 | +1.08 / +0.98 |
| hunter-marksmanship | horde | 615.7 to 622.5 | +1.10 / +0.99 |
| hunter-survival | alliance | 741.1 to 741.1 | +0.00 / +0.00 |
| hunter-survival | horde | 744.7 to 744.7 | +0.00 / +0.00 |
| mage-arcane | alliance | 613.7 to 613.7 | +0.00 / +0.00 |
| mage-arcane | horde | 619.8 to 619.8 | +0.00 / +0.00 |
| mage-fire | alliance | 665.2 to 665.2 | +0.00 / +0.00 |
| mage-fire | horde | 672.9 to 673.0 | +0.01 / -0.01 |
| mage-frost | alliance | 627.3 to 627.3 | +0.00 / +0.00 |
| mage-frost | horde | 632.0 to 632.0 | +0.00 / +0.00 |
| paladin-holy | alliance | 656.5 to 656.5 | +0.00 / +0.00 |
| paladin-holy | horde | 645.1 to 645.1 | +0.00 / +0.00 |
| paladin-protection | alliance | 1088.3 to 1113.1 | +2.28 / +2.26 |
| paladin-protection | horde | 1101.0 to 1125.7 | +2.24 / +2.15 |
| paladin-retribution | alliance | 693.0 to 704.3 | +1.63 / +1.63 |
| paladin-retribution | horde | 690.1 to 701.5 | +1.65 / +1.61 |
| priest-discipline | alliance | 649.8 to 652.4 | +0.40 / +0.38 |
| priest-discipline | horde | 653.5 to 655.9 | +0.37 / +0.41 |
| priest-holy | alliance | 708.9 to 710.5 | +0.23 / +0.20 |
| priest-holy | horde | 696.7 to 698.9 | +0.32 / +0.40 |
| priest-shadow | alliance | 712.7 to 712.7 | +0.00 / +0.00 |
| priest-shadow | horde | 720.0 to 720.0 | +0.00 / +0.00 |
| rogue-assassination | alliance | 625.1 to 626.4 | +0.21 / +0.16 |
| rogue-assassination | horde | 621.0 to 622.1 | +0.18 / +0.08 |
| rogue-combat | alliance | 625.5 to 627.0 | +0.24 / +0.22 |
| rogue-combat | horde | 626.8 to 628.2 | +0.22 / +0.24 |
| rogue-subtlety | alliance | 549.6 to 550.8 | +0.22 / +0.18 |
| rogue-subtlety | horde | 549.0 to 550.2 | +0.22 / +0.18 |
| shaman-elemental | alliance | 471.6 to 471.6 | +0.00 / +0.00 |
| shaman-elemental | horde | 479.6 to 479.6 | +0.00 / +0.00 |
| shaman-enhancement | alliance | 585.4 to 585.4 | +0.00 / +0.00 |
| shaman-enhancement | horde | 586.0 to 586.0 | +0.00 / +0.00 |
| shaman-restoration | alliance | 842.8 to 842.8 | +0.00 / +0.00 |
| shaman-restoration | horde | 840.2 to 840.2 | +0.00 / +0.00 |
| warlock-affliction | alliance | 570.2 to 570.0 | -0.04 / -0.09 |
| warlock-affliction | horde | 572.1 to 571.9 | -0.03 / -0.09 |
| warlock-demonology | alliance | 636.5 to 635.5 | -0.16 / -0.44 |
| warlock-demonology | horde | 641.4 to 639.3 | -0.33 / -0.25 |
| warlock-destruction | alliance | 635.0 to 634.0 | -0.16 / -0.27 |
| warlock-destruction | horde | 638.0 to 636.2 | -0.28 / -0.30 |
| warrior-arms | alliance | 719.5 to 719.0 | -0.07 / +0.01 |
| warrior-arms | horde | 731.7 to 731.2 | -0.07 / +0.01 |
| warrior-fury | alliance | 754.9 to 755.3 | +0.05 / +0.17 |
| warrior-fury | horde | 757.2 to 758.2 | +0.13 / +0.07 |
| warrior-protection | alliance | 1080.0 to 1081.2 | +0.11 / +0.06 |
| warrior-protection | horde | 1069.2 to 1068.8 | -0.04 / +0.05 |

Where the moves over 1% at raid 60 come from:
- **Protection paladin +2.2 to +2.3% threat, Retribution +1.6%:** Elixir of Holy Power in the raid presets.
- **Marksmanship +1.0 to +1.1%:** Juju Might returns (+1.15%), less Arcane Shot's lost spell power (-0.13%).
- **Beast Mastery -0.5%:** Juju Might (+1.0%) against Arcane Shot's lost spell power (-1.5%). Its published set carries spell power, which now does nothing; the re-rank is "measure at refresh".

Under 1% at raid 60:
- Discipline +0.4% and Holy +0.2 to +0.4%: the pre-pull Inner Fire and Lightwell timing.
- Rogues +0.1 to +0.24%: Poised.
- Warlocks -0.03 to -0.44%: bolt travel.
- Every other spec within 0.2%.

Moves over 1% below 60 and on builds (lane reports have every band):
- **Bands:**
  - Protection paladin band 30: +38 to +40% threat (Holy Shield in the build).
  - Restoration shaman band 20: +12 to +16% (Water Shield).
  - Holy paladin 40: +2.8 to +3.1% (Flash of Light pair).
  - Arms bare 30: +2.4 to +2.6%; bare 40: +3.5 to +3.6% (the execute swap).
  - Fire bare 30 to 60: +1 to +2% (Evocation).
  - Holy and Discipline bare 50/60: +0.6 to +1.6%.
  - Elemental Flame Shock rank 1 at 30: +1.4 to +2.6%.
  - Horde Fire Blast rank 3 at 30: +1.7%.
- **Builds** (they reach published numbers at the refresh):
  - Cat leveling build: +1.1 to +1.3% bare 60, about +1.5% band 50 (adopted under the rule).
  - Fury build: +0.7 to +1.0% raid, +1.7% bare 60 (points moved out of talents with no Fury effect, the Shadow-build precedent).
  - Frost leveling build, kept on audit: +5.9 to +6.3% bare 60.

## wow-player review

- **SHIP:** warlock, regcast, healers, itemscrit, trainables, the custom panel elixirs.
- **SHIP after fixes, every fix done:**
  - **guides:**
    - feral's Genesis is a Balance point, and Natural Shapeshifter is 3 in both builds;
    - the Scatter Shot line became Freezing Trap;
    - the paladin class landing (S 9efecab6f);
    - the rogue Kidney Shot and Holy Shadowfiend coach lines are generated from the fixed prose and change at the coach-starter regen.
  - **platform:**
    - the Shadow strip prose says what the strips show (S 9efecab6f);
    - the weights caption is already on the run page (WeightsScaleTable.svelte:19).
  - **regmelee:**
    - the Arms opener charges before Bloodrage (W5);
    - Fury's dead points moved (W6);
    - "skip the shout under a raid shout" in the guides and notes (W7).
  - **hybrids:**
    - the cat build line and prose (S 454ae42b9);
    - raid-keeps (S 454ae42b9);
    - the bear coach id and trainables' Primal Bite (trainables T7; the coach rules follow at regen).
  - **healprofile:** levels 13 to 30 named, pooled bands say none were logged, and the Holy paladin 40 export is checked (H9).
  - **ranks:**
    - each faction row names its race (R12);
    - the coach no longer flags a rank between the site's pick and the top (platform C6).
  - **holypower:** the panel groups it as a battle elixir and hides the two inert elixirs (S bc08b8540).
  - **casters:** the regcast lane's guide-cooldowns gap was fixed (S 157e893b0).
- **Two reviewer calls left for the owner:**
  - **Tidal Focus at bare 50/60.** The reviewer would adopt it: 5% spell hit for Improved Healing Wave, +5.6 to +5.9%. The casters lane held it because it drops an unmodeled guide talent.
  - **Arcane Shot's spell power.** One in-game check settles it before launch: Arcane Shot's tooltip with Greater Arcane Elixir up. If the number moves, revert E 05b72a3b5.
- **Release-note lines** (the reviewers'):
  - "Rogues get Poised from 18 and cats from 32: at full energy out of combat, your next ability costs 10 less."
  - "Arcane Shot, Serpent Sting and Volley no longer count spell power."
  - "Beast Mastery and Marksmanship again assume Juju Might, which adds 40 ranged attack power."
  - "Retribution and Protection paladins use Elixir of Holy Power (+40 Holy spell damage) in raids."
  - "Arms warriors under 50 stay in Battle Stance through Execute; Arms opens with Charge, then Battle Shout, then Bloodrage."
  - "Warriors in a raid skip their own Battle Shout when another warrior's is up."
  - "Protection paladins at 30 take Holy Shield."
  - "Cats level with Genesis, three Heart of the Wild and Natural Shapeshifter from 50."
  - "Bears use Primal Bite rank 3, the rank druids learn; Dire Bear Enrage costs 16% armor."
  - "Windfury Totem's bonus ends with the totem."
  - "Shadow Bolt and Incinerate fly to the target; hunter shots and wands fly at their own speeds."
  - "Vampiric Embrace heals the party for 20% of your Shadow damage."
  - "Restoration shamans at 20 take Water Shield; sub-60 healer figures use the tank-and-group split from beta logs."
  - "Some rank choices now differ by faction, where the races differ."
  - "Leveling talent strips follow the order you spend points."
  - "Guides: Kidney Shot costs 25 energy, there is no Unstable Affliction or Haunt in Forever, and more fixes."

## For the integrator

**Branches.**
- Site `int/r12` merges all 15 site lane branches plus the coordinator commits (head f0eddb375 before this line; the next commit records it).
- Fork `int/r12` merges 11 lane branches plus the goldens commit, head 076a5a517.
- The fork has no master fast-forward.

**Pin first.** The site must pin the fork's int/r12 head:
- data/proto/ENGINE_SHA and sim/enginever/version.go;
- web/src/lib/sim/version.ts and the web sim fixtures, as at every pin;
- data/proto/common.proto follows the engine (holy_power_buff = 38);
- data/pipeline/simproto/common_pb2.py: `python -m pipeline simproto --engine <fork at the pin>`.

**Regenerate, in this order.** Each of these was left by a lane on purpose:
1. **apl-sync** (`make apl-sync`).
   - Required: sim/request/apl/druid-feral-bear.apl.json still names Primal Bite 1238073, which the engine no longer registers. TestEveryWrittenRotationRunsInTheEngine fails until it is synced, and the bear drops Primal Bite.
   - druid-feral.apl.json: notes only.
   - Then `make apl-check`.
2. **Trainables:** `python -m pipeline trainables --build 1.60.1.70753`.
   - Trainables changes: 124 rows tagged talent, 16 talents added, copies dropped as ranks.
   - Then the readers: addon-data.json and the coach rules (pipeline.coachrules; this fixes the bear coach's 1238073).
3. **The committed build copy of curated data:** data/builds/1.60.1.70753/classes.json (the paladin landing line; test_curated.py::test_the_committed_build_matches_the_curated_facts).
4. **addon-builds.json:** `go run ./sim/cmd/addon-builds -repo-root .`. This covers:
   - the cat leveling build;
   - Protection band 30;
   - Restoration shaman band 20;
   - the Fury build.

   TestCommittedAddonBuildsFollowTheGuides and web leveling-builds.test.ts (druid-feral, paladin-protection, shaman-restoration, warrior-fury) fail until then.
5. **sim/request TestRotationLadder goldens** at the pin. These move:
   - the warlocks, warriors, hunters, rogues and mages;
   - priest-holy, priest-discipline, priest-shadow;
   - shaman-elemental;
   - paladin-protection (band 30) and paladin-retribution;
   - druid-feral and druid-feral-bear;
   - the five healers below 60;
   - Primal Bite's id in the druid unused lists.
6. **addon-data:** `pipeline addon-data`, which writes data/builds/1.60.1.70753/addon-data.json and addon/ForeverSixtyData/Data.lua. This covers:
   - rank-choice bands and `factionSpellId`;
   - the warrior and druid rotations;
   - the Holy paladin 40 pair;
   - trainables' rank lists.
7. **Coach starters:** `python tools/coach/persona_starter.py`. They change for:
   - the three rogues (Kidney Shot 25 energy);
   - priest-holy (the stale Shadowfiend lines);
   - the three warriors (the shout).
8. **dps-per-point.json** band 60 for every spec, then Data.lua. Casters H4: the tool now uses the band's preset.
9. **Nightly-owned, at the refresh:**
   - simbuffs.json: the Mender's potions, elixir_of_holy_power;
   - BiS files: haste scale factors, Fury and mage-frost talents, the hunters' Juju note.
10. **Fork, at the next engine follow:** the conformance trainables copies and goldens (trainables).

Already done on the branches, not to redo: sim/request/raid-consumes.json and IDS.md (S 0206582a6, S 408d35cde), the fork goldens (E 690028efd), sim/rankchoices/rank-choices.json.

**Gates.** No lane ran a full suite. On int/r12 the coordinator ran:
- the fork's `go test -tags=with_db ./sim/...`: green except sim/web (setup fails at the base pin too) and the goldens regenerated in E 690028efd;
- site sim/request: red only on TestRotationLadder and TestEveryWrittenRotationRunsInTheEngine (regens 1 and 5), plus sim/rankchoices and sim/leveling;
- the guide vitest set;
- the touched data pytests.

The full gate set runs once, at the landing.

## Measure at refresh (added this round)

- **MR14:** Beast Mastery and Marksmanship gear without spell power (Arcane Shot H4); Juju Might in their sets.
- **MR15:** Retribution and Protection paladin with Elixir of Holy Power in the preset (picks, weights, tiers).
- **MR16:**
  - the cat at bands 50 and 60 on the new leveling build, and raid 60 on the c0f5432b2 raid build;
  - Protection paladin band 30 with Holy Shield;
  - Enhancement on b2e4a3bc5's builds.
- **MR17:** every sub-60 healer pick on the logged split; Restoration shaman band 20 and Holy paladin 40 above all; the druid idol slot at 60 (Idol of Swiftness).
- **MR18:** Fury's build at bare and raid 60; mage-frost's bare talents.
- **MR19:**
  - haste scale factors for every spec (H4 and the core cast-timing fix);
  - the guide builds against the stale published builds (Holy +2.2%, Discipline +1.8%, Restoration druid +0.7%).

## Still open

- **Owner decisions:**
  - Tidal Focus at bare 50/60 (the reviewer would adopt it).
  - A race key on rank rows: every faction row depends on one race (ranks R12).
  - Randomising the regen tick's phase per fight (removes small ties at 0% haste; moves every mana class's fork goldens).
  - The level-60 heal split needs a warcraftlogs v2 API key.
  - Improved Bane of Agony: +0.72%, held. The point it replaces, Improved Drains, touches nothing the rotation casts, so the Shadow-build precedent could apply.
- **Verify in game or on logs:**
  - Arcane Shot's tooltip with spell power up (H4);
  - Bloodrage's in-combat flag against Charge;
  - Prayer of Mending's hop travel;
  - Blood Fury on heals;
  - Lightwell and Binding Heal rank 1 casts;
  - Holy Shield crit;
  - Swiftmend with Nature's Splendor;
  - Illumination's fourth effect;
  - Blessing of Kings stacking with Heart of the Lion;
  - round 11's log list.
- **Held under the adoption rule:**
  - Improved Bane of Agony;
  - Destruction's Immolate line (+0.2%);
  - Marksmanship (+0.9%);
  - two cat Claw/Shred lines (+0.7%, Shred at 22);
  - round 11's holds.
- **Not done:**
  - the mobile apps carry their own copy of the old talent-order note (apps/);
  - six Era talent passives listed at level 1 in trainables (no reader);
  - stray rune ids in trainables.
- **The round was not zero:** round 13 follows the landing.

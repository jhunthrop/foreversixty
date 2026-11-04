# Sim-accuracy round 19: integration (2026-11-04)

Round 19 started from the landed round 18 (site f91e9d3f9, fork 6ea675594) on client 1.60.1.70755. One lane per class plus platform, at most five at once. The lane reports sit beside this file.

**Not a zero round:** about 77 bugs in round 15, about 43 in round 16, about 15 in round 17, about 26 in round 18 and about 39 in round 19. Every round 19 bug is `old`: none was caused by round 18, the 70755 activation or new data. Round 19 found more because three method changes widened the search:
- rank rows judged level by level under the low-rank penalty;
- client coefficients of 0 checked against logs;
- every downranked spell checked to be really cast.

Rogue is the first class with a zero round. Round 20 is needed.

## Branches and regen

- **Site `int/r19`:** every lane, on f91e9d3f9.
- **Fork `int/r19`:** at **da54b7b06**, pushed. Fork master is fast-forwarded at landing.
- **Pin:** cd836f546. The proto is re-vendored for the new `currentLevel` APL value.
- **Regen once at the pin (a2231e7aa):**
  - ladder goldens (every spec, because the header text changed for Enhancement);
  - addon-data and Data.lua;
  - coach starters (hunter-survival, priest-discipline, rogue-combat, warrior-arms);
  - dps-per-point.
- addon-builds and the planner vectors came out unchanged, because no guide build changed. apl-sync was clean and apl-check passes on every copy.
- **Not regenerated, left to the BiS re-rank:** the Frost BiS, which moves with Ice Lance, and the warrior BiS, which moves with Deep Wounds. Relic near-tie publishing and the spell-side hit cap start at the next re-rank.
- **Checks:**
  - Go: gofmt and vet, plus the sim request, leveling, rankchoices, cmd and internal packages, and api sims and coach;
  - data: test_apl, test_addonrotation and test_addondata, with addon-data and coach-rules `--check` clean;
  - web vitest: 8756 of 8757 pass. The one failure, spec-icon "has art for every icon in the published build", reads the local fixture icon folder, which lacks Starfall art. Round 19 changes no spec icon; CI is the gate.
- **Stale test fixed by the integrator (e9c8c6253):** the rotation smoke test now explains the Enhancement gate's Maelstrom Weapon rank-5 aura (408498 tag 5) on a build without the talent.

## Before/after, raid 60

Before: round-start engine 6ea675594 on site f91e9d3f9. After: int/r19 at the pin. Both runs use:
- the same frozen published BiS;
- each side's guide raid builds;
- 10,000 iterations at seeds 1 and 50001, both factions.

The table lists the specs moving more than 1%: the seed 1 value, then the change at seed 1 / seed 50001.

| Spec | Faction | Seed 1 | % |
|---|---|---|---|
| mage-frost | alliance | 696.0 to 632.3 | -9.15 / -9.12 |
| mage-frost | horde | 686.9 to 624.1 | -9.14 / -9.12 |
| warrior-arms | alliance | 820.4 to 853.0 | +3.97 / +3.85 |
| warrior-arms | horde | 826.4 to 855.5 | +3.52 / +3.38 |
| warrior-fury | alliance | 879.4 to 917.6 | +4.34 / +4.48 |
| warrior-fury | horde | 873.3 to 909.8 | +4.18 / +4.27 |

- **Protection warrior** threat is down 0.8% (the Heroic Strike threat per rank and Deep Wounds).
- **Every other spec** moves less than 0.1% at raid 60.

### Frost: Ice Lance's coefficient

- **Client:** SpellEffect EffectBonusCoefficient and BonusCoefficientFromAP are 0 on every effect of every Ice Lance rank (1312002, 400640, 1240044 to 1240047). The rows are identical in 70755 and 70790.
- **Old value:** the engine's 0.4286 was the generator's default for instant spells, not a client value.
- **Logs:** nine players in seven foreverlogs beta reports (among them 2727, 2730, 2747, 2762, 2763 and 2786; the full list is in mage.md), levels 24 to 35. Each player's spell power was read off their own Frostbolt, Cone of Cold and Frost Nova hits in the same fights; those three agree within about 10%. Per-player coefficients run 0.09 to 0.13, and the hit-weighted mean is 0.1045 over 304 non-crit hits. The frozen-to-unfrozen ratio of 4.0 also confirms the whole-hit x4.
- **New value: 0.107.** This meets the two-independent-log rule.
- **Effect:** the Ice Lance line still gains +2.5% over dropping it, so the rotation and the guide text are unchanged.
- **Uncertainty:** every log is level 35 or below, so the value goes on the owner's list.

### Warriors: this supersedes round 17's Heroic Strike call

Round 17 kept Heroic Strike counting as a white swing for procs, on Season of Discovery log evidence. Forever's own beta logs (builds 1.60.1 and 70450) show 228 landed Heroic Strikes with no Unbridled Wrath proc and no Flurry or Reckless Swings charge spent. Forever's own logs outrank Season of Discovery evidence, so the coordinator ruled to ship the change (E af3bc587b).

Two more fixes came with it:
- **Deep Wounds** rolls over, keeps its tick timer and ignores attack power. Blizzard says so on the community bug tracker, issue #234 (E c647736e6).
- **Heroic Strike threat** now follows the Classic table per rank, from one lineage of emulator data (E fae61cd3d).

The gain, fix by fix. Raid 60, seed 1 / seed 50001, % against the pre-lane fork 1789ffd5e:

| Spec | Heroic Strike procs | Threat | Deep Wounds (alone) | All three |
|---|---|---|---|---|
| Arms, Alliance | -0.24 / -0.29 | 0 | +4.12 / +4.07 | +3.97 / +3.85 |
| Arms, Horde | -0.27 / -0.25 | 0 | +3.53 / +3.42 | +3.52 / +3.38 |
| Fury, Alliance | +0.05 / +0.11 | 0 | +4.22 / +4.23 | +4.34 / +4.48 |
| Fury, Horde | +0.25 / +0.26 | 0 | +3.98 / +4.09 | +4.18 / +4.27 |

The fixes multiply to within 0.1% of the combined result, so the order they land in doesn't matter.

**Why removing a proc raises Fury's DPS:**
- **Fury:** Heroic Strike no longer spends Flurry charges, so white swings keep them. Flurry uptime goes from 83.9% to 86.8%. That outweighs the Unbridled Wrath rage lost (97.6 down to 79.1).
- **Arms:** it has no Flurry, so it only loses rage.
- **Deep Wounds:** almost all of the gain. The rollover keeps every crit's bleed instead of overwriting it. Fury goes from 24.3 ticks a fight (11.1 DPS) to 58.8 (46.1 DPS), and Arms from 34.3 (29.3 DPS) to 58.2 (59.0 DPS). That more than makes up for each crit's bleed losing attack power.

**Guide text:** Unbridled Wrath now procs on white swings only (arms.md:75, fury.md:67, fury.md:128), and the race lines are ties. wow-player gave SHIP on Arms, Fury and Protection.

### Below 60

- **Discipline priest:**
  - no Lesser Heal rank 3 from 40 to 59, gated on the new `currentLevel`: +7.7% at 40, up to +46% / +39% at 50 and +44% / +38% at 59;
  - a tank Heal before 40: +7.4% / +9.1% at bare 20;
  - Renew rows re-judged by race.
- **Restoration druid** (+0.3 to +2.7% at those levels):
  - tank Rejuvenation at 44-57;
  - party Rejuvenation rank 2 at 42;
  - Healing Touch filler ranks at 34-39 and 45-49.
- **Balance druid:** Wrath rank 6 at 46-48, and Insect Swarm rank 3 at 54-58.
- **Restoration shaman:** Chain Heal rank 1 at 46-59 and rank 2 at bare 60 (+1.0%); tank Healing Wave rank 7 from 56.
- **Elemental shaman:** Lightning Bolt rank 3 at 26-27, +3.7 to +4.9%.
- **Enhancement shaman:** the Earth Shock mana gate is half while Maelstrom Weapon is below 5/5, and a quarter otherwise: +0.3 to +1.2% at 40-49.
- **Holy paladin:** 40-49 rows, up to +1.5%.
- **Retribution paladin:** Consecration rank 1 ends at 53, up to +2.2% at 54-59.
- **Holy priest and Shadow priest:** rank rows.
- **Survival hunter:** Raptor Strike rank 5 at 46-47.
- **Destruction warlock:** Immolate rank 4 at 40, and Corruption rank 3 at 41.

## Confirmed by lane

| Lane | Fixed | Main ones |
|---|---|---|
| priest | 9 | Discipline Lesser Heal 40-59 and tank Heal before 40; pre-pull timing; Holy and Shadow rank rows |
| platform | 8 | card "with X" tags; Arcane Frostbolt one line per build; Enhancement spell hit cap; relic near-tie gate; "Relic" label; phone contents bar covering the heading; Discipline card's 24 repeated rows; plus new tools: `currentLevel`, `spec-breakdown -rank`, the all-ranks-registered guard |
| druid | 7 | Restoration and Balance rank rows held back by a wrongly applied 1% bar |
| shaman | 4 | Restoration Chain Heal and Healing Wave rows; Elemental bolt at 26-27; Enhancement Earth Shock gate by Maelstrom rank |
| hunter | 3 | Serpent Sting and Raptor Strike registered only their top rank; Raptor Strike rank 5 row |
| warrior | 3 | Heroic Strike and Cleave procs; Heroic Strike threat per rank; Deep Wounds |
| mage | 2 | Ice Lance coefficient; "1 point of hit" on every guide's stat table |
| paladin | 2 | Holy 40-49 rows; Retribution Consecration |
| warlock | 1 | Destruction 40-41 rows |
| rogue | 0 | aligned (Combat card notes reworded; Occult Poison closed as not learnable) |
| **total** | **about 39** | all `old` |

## Method changes this round (now in the brief)

1. **Rank-choice rows adopt on ANY gain beyond error at seeds 1 and 50001.** The 1% bar is for talent builds only. Earlier rounds had kept rows "under 1%".
2. **Client coefficient 0 means check the logs.** The engine's generated constant may be a generator default. Settle it on two or more independent foreverlogs players, or record it OPEN. Druid, shaman, paladin, priest, warlock, hunter and rogue are clean; mage was not (Ice Lance).
3. **When a lower rank "loses", confirm it was cast.** An unregistered rank silently drops the line. `TestEveryLearnedRankOfACastSpellIsRegistered` (platform) now guards every downranked spell. Its one known finding is Primal Bite "rank 4" (1238073), an unlearnable copy of rank 3. Most melee and hunter spells still register only their top rank, so any future downrank row on them needs the engine change first.

## Owner's post-launch list

1. **Discipline at 60:** drop Lesser Heal rank 3 at 60 too, by removing the `< 60` bound. That gives +30% / +24% at bare 60 and +6.0% / +7.5% at raid 60, and waits on the pending low-rank heal check.
2. **Ice Lance's coefficient at 60:** the logs top out at 35.
3. **Seal of Righteousness rank 7 vs 8 at 60** (round 18).
4. **Empty relic slot copy:** "Relic effects aren't simulated yet" in a dashed box (wow-player MUSTs; the copy predates this round).
5. **Carried from round 18:** Oath of Light, the low-rank penalty parameters, Heart of the Lion on the pet's bar, `/dump GetItemStats("item:23072")`, a naked level-60 Skyborne.

## Open, for round 20

1. **Launch logs** (live from 23:00 UTC, none at level 40+ yet):
   - Death Mark and the Mutilate/Shank off-hand;
   - Eureka! on Incinerate and Mind Flay, demon crit, Soul Siphon, Wrack and Doom;
   - Eclipse on a missed Wrath, the cat form spread;
   - Ranged Weapon Specialization on Serpent Sting, Lacerating Strikes;
   - the Seal of Righteousness one-hand/two-hand split, the second Judgement coefficient;
   - Revenge and Shield Slam threat, Arcanite Champion's proc rate.
2. **Unmodelled:** Arcane Power's cost increase on stacked Blast and Barrage Missiles; heals scaling with target level; the low-rank proc curve.
3. **Waiting on the BiS refresh:** Subtlety band 20 on a mace.
4. **Guide lane:** Fury Death Wish timing ("last 30 seconds" against the card's "only if a second one fits").
5. **Platform:** whether the heading's focus ring shows after a finger tap (ux-designer reviewed a capture only); tag density and the Enhancement headline wrapping at 390 (SHOULDs).
6. **The fork's TS UI proto** is not regenerated for `currentLevel`. Only the fork's own UI reads it.

# Sim-accuracy round 16: integration (2026-11-04)

Round 16 started from the landed round 15 (site c62d77ab8, fork pin 9f7f1f6e4) on client 1.60.1.70755 with the late hotfix re-layer. There was one lane per class plus platform, and a second wave for rank rows and talent-gated lines. The lane reports sit beside this file.

**Not a zero round.** About 43 confirmed bugs were fixed (see the table). Round 17 is needed.

**New process rule (owner, from this round on).** A change to a guide's build line, an APL action or a card rank ships with its guide prose and card notes in the same landing. That text must pass the prose test and get a wow-player SHIP on the rendered page. Every lane that changed guide-visible text did this.

## Branches and regen

- **Site `int/r16`.** Every lane, plus origin/main 1b7671ce0, merged by the coordinator's rule (c4483a3a6):
  - APL actions and ranks are round 16's, with main's notes, except where round 16 changed that line.
  - Guide text is guide rounds 10-13 (main), except sentences describing a build or rotation round 16 changed.
  - Warrior text is main's throughout.
  - Kept from round 16: Feral Finishers (Rip under 2 s), Restoration druid tank Rejuvenation ranks, the Arcane pre-pull Pyroblast, Retribution Consecration rank 2 at 44-49, Subtlety Garrote / Hemorrhage hold / Death Mark at four, and the shaman rank text with notes 7, 8 and 17.
  - wow-player gave SHIP or SHIP WITH FIXES on the merged pages; the MUSTs were fixed.
- **Fork `int/r16`** at **f637ec32c**, pushed to origin. Fork master is not fast-forwarded (that happens at landing).
- **Pin:** e2ca9fd00.
- **Regen, once, at the pin (6e57a9c60):**
  - ladder goldens (12);
  - addon-builds, addon-data and Data.lua, spellicons;
  - coach rules and starters;
  - dps-per-point;
  - the app guide contract fixtures (mage-fire, shaman-elemental, warrior-protection codes were still on 70753);
  - the planner golden vectors.
- **Checks:** apl-check is clean on all copies.
  - Web: the full vitest passes (8729), with astro check, eslint and prettier clean.
  - Go: vet and the touched sim packages pass.
  - Data: test_apl and test_addonrotation pass.
- **Three stale tests fixed by the integrator (3ebd56d0d):**
  - the priest Renew test follows the per-level rows;
  - the Subtlety Death Mark test was flaky, because protojson randomizes its whitespace;
  - Read Leyline (1259705) is an ungrouped racial id, like Stoneform.

## Before/after, raid 60

Before is the round-start engine 9f7f1f6e4 on site c62d77ab8. After is int/r16 at the pin. Both run on the same frozen copy of the published BiS (the BiS re-rank was in flight) and on each side's guide raid builds, at 10,000 iterations, seeds 1 and 50001, both factions.

Specs moving more than 1% (seed 1; % at seed 1 / 50001):

| Spec | Faction | Seed 1 | % |
|---|---|---|---|
| druid-feral | alliance | 871.6 to 886.1 | +1.66 / +1.61 |
| druid-feral | horde | 856.8 to 870.6 | +1.61 / +1.50 |
| rogue-subtlety | alliance | 722.3 to 751.3 | +4.01 / +3.93 |
| rogue-subtlety | horde | 716.8 to 746.0 | +4.07 / +4.02 |

- **Subtlety:** Garrote after Ghostly Strike was round 15's own line, and it almost never fired because Hemorrhage took the energy. Hemorrhage now waits during the Ghostly Strike opener buff. Death Mark marks at four points.
- **Cat:** Rip was recast under 6 s, clipping ticks; it now recasts under 2 s, and Bite needs 5 s on Rip.
- **Every other spec is within 1% at raid 60.** Most of round 16's change is below 60:
  - **Rank rows:** split level by level for shaman, priest, druid, paladin, mage and warlock, with gains up to +54% at the levels a row now covers.
  - **Ladder:** the Protection warrior ladder fix.
  - **Platform:** the addon and web fixes.

## The low-rank penalty, refined

The penalty now keys on a rank's client **MaxLevel + 12**, not learn level + 17 (priest E 18815822c).
- **Sources:**
  - The HealRankAdvisor addon (GitHub) measured Lesser Heal 1 at about 24% at level 30.
  - ForeverChanges' Healing Touch 1 and Holy Light 1 run one level early.
  - Ozgar's "16".
  - All three outliers are ranks whose MaxLevel is below learn level + 5, and the new key fits every measured rank.
- **Effect:** no raid-60 spec moves.
- **Still assumed:** the top-rank exemption, damage spells, and levels above 30. There are no launch logs yet.
- **Owner test:** the exact in-game test is in priest.md section 1. On a level-60 priest, swap Spell Healing gear and read Greater Heal 4 against Prayer of Healing 1/2, Lesser Heal 3 and Mind Flay 3. The expected values for each rule are listed.

## Settled this round

- **Hunter pets and player buffs:** the engine is right. The 70755 client marks every group buff "can only target players" (Attributes_3 0x100). Round 15's BM +26% upper bound was never real.
- **Aspect of the Hawk** procs on Auto Shot only (32/32 beta procs).
- **Wrack's 10%:** the mask is 0x402 (Corruption and Agony only), which matches the engine.
- **The 70783 build** was retail (Midnight). Wago hotfix rows must be read for region 91 only.
- **Ladder monotonicity:** each level keeps the level below's points plus one. It is tested for all 28 specs and every guide.
- **Expertise:** the items JSON now exports expertise, and the BiS ranker converts rating to percent.

## Confirmed by lane (approximate; the lane reports have the tables)

| Lane | Fixed | Main ones |
|---|---|---|
| warrior | 1 | Protection ladder dropped Shield Slam at 41 (now monotone for every spec) |
| rogue | 3 | Subtlety Garrote/Hemorrhage (R15), Death Mark at four, guide text |
| shaman | 4 | Lightning Bolt and Healing Wave per-level rows; NS save fired without the talent below 44; NS save took filler rows |
| priest | 8 | penalty key (MaxLevel + 12); conformance zero-share harness (R15); Holy/Discipline/Shadow stale guide text (R15); Renew and Devouring Plague per-level rows |
| druid | 7 | cat Rip clipping; Genesis on Lacerate; Moonfire/Wrath/Insect Swarm/tank Rejuvenation per-level rows |
| paladin | 6 | cast Kings with Heart of the Lion; Holy Light and tank Flash of Light per-level rows; Retribution Consecration rank 2 at 44-49; Protection Holy Strike row |
| warlock | 2 | Demonology dead Improved Voidwalker point; Shadow Bolt per-level rows |
| mage | 5 | Arcane pre-pull Frostbolt became a 3 s cast (R15); Arcane line 4 talent gate; Scorch and Frostbolt per-level rows; Winter's Chill penalty |
| hunter | 0 | aligned; pets and Hawk settled |
| platform | 7 | items expertise; Leveling strip order; coach Battle Shout rank (R14); red "?" icons; raid-only rows; addon bands at every change level; stale web tests |
| **total** | **about 43** | |

## Open, for round 17 (top first)

1. **Oath of Light and the low-rank penalty.** Holy paladin's cheap Holy Light ranks (rank 1 at 47-49, rank 2 at 50-59) win only because Oath's flat bonus skips the penalty in the engine. A log or in-game tooltip settles it; if it's penalized, re-pick those rows.
2. **The low-rank penalty's remaining parameters:** top-rank exemption, damage spells, levels above 30. The owner test is in priest.md.
3. **Addon gear score:** the site (score.ts) and addon (Gear.lua) multiply rating points by per-1% weights for every rating stat. The two must change together.
4. **Skyborne base stats:** the profile API's /statistics endpoint (not called by our import; the live profile returned 404 on 11-04) or `/dump UnitStat` on a level-60 Skyborne.
5. **Summon Hawk:** the owner's two-minute check (ten swings at 60) is in hunter.md.
6. **Logs, once Forever launch logs exist:** Exorcism/Holy Shield crit, Death Mark storage, Mutilate/Shank off-hand, Heroic Strike as a white swing for procs, Eureka! on Incinerate.
7. **Combat Eviscerate at four:** +0.7% raid but -0.3 to -1.8% bare. Not adopted (it loses at some bands), and the rotation cannot test for a dagger.
8. **Fork tests panicking on missing item ids** (TestP1Mage, hunter sharpening-stone / ranged-hit tests). They already failed on the round-start engine; the test item data is missing items 1728 and 2098.
9. Restoration druid: the Leveling section should point at the Ranks paragraph (wow-player suggestion).

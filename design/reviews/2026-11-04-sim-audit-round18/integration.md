# Sim-accuracy round 18: integration (2026-11-04)

Round 18 started from the landed round 17 (site 4600b7d82, fork b2c4a2d05) on client 1.60.1.70755. One lane per class plus platform, at most five at once. The lane reports sit beside this file.

**Not a zero round:** about 77 bugs in round 15, about 43 in round 16, about 15 in round 17, and about 26 in round 18. Most of this round's finds are priest rank rows. They were measured before round 15's low-rank penalty and were re-judged level by level this round. Only one raid-60 number moves by more than 1%. Round 19 is needed.

## Branches and regen

- **Site `int/r18`:** every lane, plus origin/main through 357a634fd.
- **Fork `int/r18`:** at **6ea675594**, pushed. Fork master is not fast-forwarded until landing.
- **Pin:** c77533d8b. It also drops Berserking 26297 from `engine-spell-names.json`, because the engine now uses 20554 (E 4742b0075).
- **Regen once at the pin (3b0afe9cf):**
  - ladder goldens;
  - addon-builds, addon-data and Data.lua;
  - spellicons;
  - coach rules and starters;
  - dps-per-point;
  - app fixtures and planner vectors.
- **Not regenerated, left to the data job:**
  - **spelltips:** the `$q` fill (Rebirth, Revive, Redemption, Resurrection, Ancestral Spirit) needs the current layered raw;
  - **items/<class>.json:** for the expertise key.
- **Checks:**
  - apl-check is clean;
  - web vitest passes in full (736 files, 8730 tests), with astro check and prettier clean;
  - Go: gofmt and vet, plus the sim packages and api sims/coach;
  - data: test_apl and test_addonrotation pass.
- **Stale tests fixed by the integrator (67d1621a8):**
  - the spec-breakdown race test reads the re-judged Shadow Mind Blast rows;
  - the rotation smoke test explains the Arcane Missiles-first gate's rank-2 Missile Barrage aura (400588 tag 2) on a build without the talent.

## Before/after, raid 60

Before: round-start engine b2c4a2d05 on site 4600b7d82. After: int/r18 at the pin. Both use the same frozen published BiS and each side's guide raid builds, at 10,000 iterations, seeds 1 and 50001, both factions. The table lists specs moving more than 1%: seed 1 value, then % at seed 1 / 50001.

| Spec | Faction | Seed 1 | % |
|---|---|---|---|
| rogue-subtlety | alliance | 765.4 to 788.4 | +3.00 / +2.99 |
| rogue-subtlety | horde | 760.1 to 782.7 | +2.97 / +2.99 |

- **Subtlety:** Hemorrhage's 15% now reaches bleed ticks.
- **Under 1% at raid 60:**
  - Holy priest +0.5 to +1.0% (rank and priority rows);
  - Assassination +0.4 to +0.5% (Deadly Poison takes Venom's bonus at each re-application);
  - Enhancement +0.1 to +0.4% (Earth Shock gate at a quarter mana);
  - Retribution -0.4% (Seal of Righteousness rank 8 has a spell power coefficient of 0 on the client, so Retribution keeps rank 7 from 58).
- **Every other spec** moves less than 0.1%.
- **Below 60:**
  - Restoration druid Healing Touch and party rows, +1.2 to +3.3% at 30 to 44;
  - Shadow Mind Blast and Devouring Plague rows at 40 to 51;
  - Holy and Discipline Renew and Greater Heal rows at 40 to 59;
  - Arcane Missiles first until both Missile Barrage ranks;
  - Marksmanship Aimed Shot rank 1 at 52 to 59.

## Confirmed by lane

| Lane | Fixed | Main ones |
|---|---|---|
| priest | 10 | Shadow Mind Blast and Devouring Plague rows at 40-51; Holy and Discipline Renew and Greater Heal rows; Holy Prayer of Healing at bare 60 |
| druid | 4 | Read Leyline and Innervate no longer stack (shared core, every Alliance Skyborne); Attuned in raids (guide text); Restoration rows |
| rogue | 3 | Hemorrhage's 15% on bleeds; Deadly Poison's Venom snapshot; Thistle Tea above 100 energy (shared core) |
| platform | 3 | iOS/Android expertise line; Arcane opener card (gate-aware cards for every spec); the leveling strip (fixed, then superseded by land/leveling-tree) |
| paladin | 2 | Seal of Righteousness ranks 1-3 coefficient; rank 8 coefficient 0 (Retribution keeps rank 7) |
| mage | 2 | Arcane Missiles first at 20-22 with Missile Barrage 1/2; the raid Frostbolt row note |
| hunter | 2 | Summon Hawk's 18 s assault modelled (no guide build takes it: +0.6 to 0.7%, under the adoption bar); Aimed Shot rank 1 at 52-59 |
| shaman | 1 | Enhancement Earth Shock gate half to quarter mana |
| warrior | 0 | aligned; the racials audit moved Berserking to client id 20554 and fixed three racial comments (no number moves) |
| warlock | 0 | aligned; every log-dependent item is still open |
| **total** | **about 26** | |

## Calls decided this round

| Call | Decision | Reason |
|---|---|---|
| Seal of Righteousness rank 7 vs 8 at 58-60 | Keep rank 7 | The client gives rank 8 a coefficient of 0 (client-data rule). The owner checks it in game (list below). |
| Arcane opener card | Gate-aware cards (platform) | A card no longer shows a cast whose talent gate excludes the described build. This applies to every spec, with a test over every raid and leveling card. |
| Frost Cold Snap for Ice Lance (+1.4% raid) | Keep Cold Snap | Every Forever source and the guide keep it for Ice Block (the strong-player consensus rule). |
| Enhancement Earth Shock gate | Quarter mana | It wins at every band but 40-49. At 40-49 (Maelstrom Weapon 3/5) it costs 0.3 to 0.85%. No single gate keeps that band level, and the guide gives one rule. |
| Reckoning vs Conviction and Deflection | Keep the guide | Reckoning loses on threat and on survival. |

## Owner's post-launch list

1. **Seal of Righteousness rank 7 vs rank 8 at 60:** compare procs with and without spell damage gear. Rank 7 stays in the spellbook.
2. **Oath of Light** (round 17 item 1, still open).
3. **Low-rank penalty parameters** (round 16 priest.md, section 1): the top-rank exemption, damage spells, levels above 30.
4. **Pet buff bar:** check that Heart of the Lion shows on the pet. Also check that the pet copies the hunter's crit, including Elune's Light. That rests on guides and beta logs, and there is no client row for it.
5. **`/dump GetItemStats("item:23072")`** and **a naked level-60 Skyborne's base stats** (round 17 items 5 and 6).

## Data job (coordinator)

- Spelltips and talent text rebuild (the raw `$q` resurrection tooltips).
- items JSON with the expertise key.
- After launch:
  - Subtlety level 20 mace vs dagger (+5 to 10%);
  - Frost 60 BiS on the old build;
  - Enhancement hit weight rebuild.

## Open, for round 19

1. Everything on the owner list above.
2. **Log-dependent items,** once level-60 Forever logs exist (none at 16:18 UTC):
   - Exorcism/Holy Shield crit, Death Mark storage, Mutilate/Shank off-hand;
   - Heroic Strike procs;
   - Eureka! on Incinerate, demon crit, Soul Siphon;
   - Eclipse on a missed Wrath, the cat form spread;
   - Heart of the Lion on the pet.
3. **Priest, unmodelled:**
   - low ranks proc less (the curve is unknown);
   - heals scale with target level (one forum thread).
4. **Holy priest Priority card:** re-capture it after landing (wow-player condition).
5. **Enhancement:**
   - the hit-cap line ignores spell miss;
   - the ladder golden header says Enhancement dual-wields;
   - enhancement.md:93 wording on hit past the cap.
6. **Mage card SHOULDs (platform):**
   - the Fireball clause "with Heating Up" wording;
   - a "For the raid build" header over a shared row.

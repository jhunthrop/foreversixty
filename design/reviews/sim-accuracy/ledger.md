# Sim-accuracy ledger

The owner's goal (2026-10-31): repeat audit and fix rounds over every spec until a whole round finds zero confirmed sim bugs. Each round has three or more audit groups, and every confirmed bug is fixed end to end before the next round starts.

| Round | Dates | Reports | Confirmed bugs | Engine pin after | Zero-bug round? |
|---|---|---|---|---|---|
| 1 | 2026-10-31 | design/reviews/2026-10-31-theorycraft-audit/ (casters, hybrids, physical) | every row in the reports' Fixed columns; the rest held under the adoption rule or open | 17d31ecfa (via 31fbe10e2, c99fb94b4, f923e1ad1, 09ca529f3 and the lane follows) | no |
| 2 | 2026-10-31 to 2026-11-01 | design/reviews/2026-11-01-sim-audit-round2/ (warlock, casters, hybrids, physical) | about 35 fixed; 2 held (warlock demons refusing player buffs, and the Destruction build that depends on it) | c8d29ea33 | no |

## Round 2 at a glance

Confirmed and fixed, by group:
- **Warlock (8).**
  - Wrack is a 6 s channel, and its +10% reaches only Corruption and Agony.
  - Nightfall and Shadow Mastery follow the client masks.
  - Affliction adopts Malevolence and Ruin.
  - Decimation applies only below 35% health, its cast cut is 40% of 4 s, and it no longer cuts mana.
  - Blood Fury is off the global cooldown.
- **Priest and mage (7, plus Shadowfiend).**
  - Darkness covers all Shadow damage.
  - Shadow Word: Death adds Shadow Weaving, takes its 30% backlash, and is cast on cooldown.
  - Dark Sacrifice is Undead-only.
  - Prayer of Mending spends a charge when its target is hit.
  - Ignite no longer double-dips target modifiers.
  - Shadowfiend is added below 30% mana.
- **Hybrids (5).**
  - The Feral form attack uses the weapon's DPS, for cat and bear.
  - The Ravage opener counts.
  - Fire Nova uses the Forever totem damage.
  - Retribution takes Unsealed Wrath and swaps seals.
  - Two client amounts are corrected: Heart of the Lion is 280 AP at 60, and Judgement of the Crusader is 161.
- **Physical (11).**
  - Mutilate needs daggers in the BiS ranker.
  - A stone sits beside a poison on rogues.
  - Revenge and Sunder scale with attack power.
  - The pre-pull cast lands at 0 and counts.
  - Lacerating Strikes no longer double-dips.
  - Subtlety adopts Quietus; Beast Mastery adopts Careful Aim 5.

Held:
- Demons refusing player buffs (lane E dfbb63ae3). Hunter pets already refuse them, and the client's warlock pet scaling spell matches the hunter one, but demons inherit no owner stats yet, so landing it alone would push pet damage further from the logs. It lands with a pet-inheritance model.
- Destruction's Pandemic and Cataclysm build (lane S 5d013a9c0). It passes the adoption rule only once demons refuse buffs.

Open going into round 3 (each report's "Still open" has the detail):
- Warlock pet inheritance.
- Siphon Life never crit in 772 low-level logged ticks, against the owner's DoT crit rule.
- Kings stacking with Heart of the Lion.
- Judgement of the Crusader in the raid preset.
- Omen of Clarity's rate.
- Fury cooldowns held for the execute phase.
- Venom with Vile Poisons.
- Murder's target type.
- Revenge rank 6.
- The rage-on-crit factor.
- Penance and Power Word: Shield coefficients.
- Ice Lance's cooldown on lower ranks.
- Shadowfiend's learnability at launch.
- talent-search and spec-breakdown give different baselines for Marksmanship.
- Elemental's Improved Reincarnation hold.

Real logs: no level-60 Forever logs exist before launch (2026-11-04). The 29 warlock logs at levels 13 to 30 on foreverlogs.gg agree that DoT ticks crit and that Life Tap scales with Spirit. Our API holds one public beta trash report. Round 3 starts on the first level-60 logs after launch, and warlocks are checked against real logs every round.

## Round 2 wow-player review

SHIP: Demonology, Destruction, Fire, Arcane, Frost, Protection warrior, Combat, Subtlety, the hunters, Feral, Elemental, Enhancement, Retribution and Protection paladin. The reviewer held three, and each was settled before the push:
- **Affliction, "Wrack is never cast".** The client data settles it. Wrack (1316697) is 36 a tick at 0.143 a tick over a 6 s channel, 0.143 coefficient a second, against Shadow Bolt's 0.343. Every Wrack rotation measured below the shipped one, wowtbc's Forever rotation verbatim included (-2.4%). Row A14 of warlock.md has the numbers.
- **Shadow, "Shadow Word: Death on cooldown kills the priest".** Two Forever guides (Warcraft Tavern Forever, wowtbc Forever) rank Shadow Word: Death above Mind Flay. The 30% backlash costs the healers, not the priest's damage. The rotation stays.
- **Assassination, "about 260 DPS on the page".** The ranker now needs daggers for Mutilate. The bis.yml run dispatched after the push re-ranks the set (576 on daggers in the lane run).
